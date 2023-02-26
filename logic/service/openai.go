package service

import (
	"bytes"
	"encoding/json"
	"go-IM/consts"
	"go-IM/pkg/util"
)

type CompletionsRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature"`
	N           int8    `json:"n"`
	Stream      bool    `json:"stream"`
	MaxTokens   int32   `json:"max_tokens"`
}

type Choice struct {
	Text         string `json:"text"`
	Index        int32  `json:"index"`
	FinishReason string `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int32 `json:"prompt_tokens"`
	CompletionTokens int32 `json:"completion_tokens"`
	TotalTokens      int32 `json:"total_tokens"`
}

type CompletionsResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// ProxyRobotPost 代理请求 Openai API 接口
func ProxyRobotPost(prompt string) (string, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + consts.APIKey,
		"Content-Type":  "application/json",
	}
	req := &CompletionsRequest{
		Model:       "text-davinci-003",
		Prompt:      prompt,
		Temperature: 0.7,
		N:           1,
		Stream:      false,
		MaxTokens:   100,
	}
	reqBytes, _ := json.Marshal(req)
	rspBody, err := util.DefaultClient.DoReq("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(reqBytes), headers)
	if err != nil {
		return "", err
	}
	rsp := &CompletionsResponse{}
	if err := json.Unmarshal(rspBody, rsp); err != nil {
		return "", err
	}
	var answer string
	for _, choice := range rsp.Choices {
		answer += choice.Text
	}
	return answer, nil
}
