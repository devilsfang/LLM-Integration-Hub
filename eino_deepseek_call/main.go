package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino/schema"
	"log"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
)

func main() {
	ctx := context.Background()
	//apiKey := os.Getenv("DEEPSEEK_API_KEY")
	//if apiKey == "" {
	//	log.Fatal("DEEPSEEK_API_KEY environment variable is not set")
	//}

	// 创建 deepseek 模型
	cm, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		BaseURL:   "https://api.siliconflow.cn",
		APIKey:    "sk-qerbnmnahexfbjhcopejqeourtohhhcmelpvzmcdpsgyindz",
		Model:     "deepseek-ai/DeepSeek-R1",
		MaxTokens: 2000,
	})
	if err != nil {
		log.Fatal(err)
	}

	messages := []*schema.Message{
		{
			Role:    schema.System,
			Content: "You are a helpful AI assistant. Be concise in your responses.",
		},
		{
			Role:    schema.User,
			Content: "What is the capital of France?",
		},
	}

	resp, err := cm.Generate(ctx, messages)
	if err != nil {
		log.Printf("Generate error: %v", err)
		return
	}

	reasoning, ok := deepseek.GetReasoningContent(resp)
	if !ok {
		fmt.Printf("Unexpected: non-reasoning")
	} else {
		fmt.Printf("Resoning Content: %s\n", reasoning)
	}
	fmt.Printf("Assistant: %s\n", resp.Content)
	if resp.ResponseMeta != nil && resp.ResponseMeta.Usage != nil {
		fmt.Printf("Tokens used: %d (prompt) + %d (completion) = %d (total)\n",
			resp.ResponseMeta.Usage.PromptTokens,
			resp.ResponseMeta.Usage.CompletionTokens,
			resp.ResponseMeta.Usage.TotalTokens)
	}
}
