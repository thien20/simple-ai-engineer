package infra

import (
	"bytes"
	"encoding/json"
	"go-be/internal/constant"
	"go-be/internal/requests"
	"go-be/internal/response"
	"net/http"
)

type LLMClient interface {
	GetAnswer(input string) (response.OllamaResponse, error)
}

func NewLLMClient() LLMClient {
	return &llmClient{}
}

type llmClient struct{}

func (l *llmClient) GetAnswer(input string) (response.OllamaResponse, error) {
	reqLLM := requests.OllamaRequest{
		Model:  constant.LLMModel,
		Prompt: input,
		Stream: false,
	}

	reqBytes, err := json.Marshal(reqLLM)
	if err != nil {
		return response.OllamaResponse{}, err
	}
	llmResp, err := http.Post(constant.OllamaApi, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return response.OllamaResponse{}, err
	}
	defer llmResp.Body.Close()
	var response response.OllamaResponse
	if err := json.NewDecoder(llmResp.Body).Decode(&response); err != nil {
		return response, err
	}
	return response, nil
}
