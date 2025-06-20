package service

import (
	"go-be/internal/builder"
	"go-be/internal/constant"
	"go-be/internal/infra"
	"go-be/internal/requests"
	"go-be/internal/response"
)

type RagService struct {
	RetrieverClient infra.RetrieverClient
	LLMClient       infra.LLMClient
	PromptBuilder   builder.PromptBuilder
}

func NewRagService() *RagService {
	return &RagService{
		RetrieverClient: infra.NewRetrieverClient(),
		LLMClient:       infra.NewLLMClient(),
		PromptBuilder:   &builder.DefaultSystemPrompt{},
	}
}

func (s *RagService) GenerateAns(input requests.RagRequest) (response.OllamaResponse, error) {
	documents, err := s.RetrieverClient.GetDocuments(input)
	if err != nil {
		return response.OllamaResponse{}, err
	}

	if len(documents) == 0 {
		return response.OllamaResponse{}, nil
	}

	prompt := s.PromptBuilder.BuildPrompt(constant.SystemPrompt, documents[0], input.UserInput)

	return s.LLMClient.GetAnswer(prompt)
}
