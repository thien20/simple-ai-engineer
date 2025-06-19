package service

import (
	"bytes"
	"encoding/json"
	"go-be/internal/constant"
	"go-be/internal/infra"
	"go-be/internal/requests"
	"go-be/internal/response"
	"net/http"
)

type RagService struct {
	RetrieverClient infra.RetrieverClient
	LLMClient       infra.LLMClient
}

func (s *RagService) GetDocuments(input requests.RagRequest) ([]string, error) {
	retrieverReq, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	retrieverResp, err := http.Post(constant.RetrieveApi, "application/json", bytes.NewBuffer(retrieverReq))
	if err != nil {
		return nil, err
	}
	defer retrieverResp.Body.Close()

	var retrieverResponse response.RetrieverResponse
	if err := json.NewDecoder(retrieverResp.Body).Decode(&retrieverResponse); err != nil {
		return nil, err
	}

	return retrieverResponse.Documents, nil
}
