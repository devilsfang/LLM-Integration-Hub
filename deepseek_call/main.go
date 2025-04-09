package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	httpClient := http.DefaultClient

	llm, err := openai.New(
		openai.WithToken(""),
		//openai.WithModel("deepseek-ai/DeepSeek-R1"),
		//openai.WithModel("Qwen/QwQ-32B"),
		openai.WithModel("deepseek-ai/DeepSeek-V3"),
		openai.WithBaseURL("https://api.siliconflow.cn"),
		openai.WithHTTPClient(httpClient),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "`你是一个网络安全专家。请分析 {{.cve}} 漏洞并返回规范的JSON格式数据，包含以下字段：\n\t\t- cve_id (必须)\n\t\t- name (漏洞名称)\n\t\t- description (详细技术描述)\n\t\t- cvss_score (CVSSv3 评分)\n\t\t- severity (高危/中危/低危)\n\t\t- affected_versions (受影响版本数组)\n\t\t- solution (修复方案)\n\t\t确保使用双引号且不包含注释`"),
		llms.TextParts(llms.ChatMessageTypeHuman, "`你是一个网络安全专家。请分析 CVE-2024-12345 漏洞并返回规范的JSON格式数据，包含以下字段：\n\t\t- cve_id (必须)\n\t\t- name (漏洞名称)\n\t\t- description (详细技术描述)\n\t\t- cvss_score (CVSSv3 评分)\n\t\t- severity (高危/中危/低危)\n\t\t- affected_versions (受影响版本数组)\n\t\t- solution (修复方案)\n\t\t确保使用双引号且不包含注释`"),
	}

	if _, err := llm.GenerateContent(ctx, content,
		llms.WithMaxTokens(1024),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		})); err != nil {
		log.Fatal(err)
	}
}
