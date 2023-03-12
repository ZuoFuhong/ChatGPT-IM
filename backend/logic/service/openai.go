package service

import (
	"ChatGPT-IM/backend/consts"
	"ChatGPT-IM/backend/pkg/util"
	"bytes"
	"encoding/json"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Choice struct {
	Message      Message `json:"message"`
	Index        int32   `json:"index"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int32 `json:"prompt_tokens"`
	CompletionTokens int32 `json:"completion_tokens"`
	TotalTokens      int32 `json:"total_tokens"`
}

type ChatCompletionResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// ProxyRobotPost 代理请求 OpenAI Chat completion 接口
func ProxyRobotPost(content string) (string, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + consts.APIKey,
		"Content-Type":  "application/json",
	}
	req := &ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
			{
				Role:    "user",
				Content: content,
			},
		},
	}
	reqBytes, _ := json.Marshal(req)
	rspBody, err := util.DefaultClient.DoReq("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBytes), headers)
	if err != nil {
		return "", err
	}
	rsp := &ChatCompletionResponse{}
	if err := json.Unmarshal(rspBody, rsp); err != nil {
		return "", err
	}
	var answer string
	for _, choice := range rsp.Choices {
		answer += choice.Message.Content
	}
	return answer, nil
}
