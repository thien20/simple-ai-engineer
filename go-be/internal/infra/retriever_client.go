package infra

import (
	"bytes"
	"encoding/json"
	"go-be/internal/constant"
	"go-be/internal/requests"
	"go-be/internal/response"
	"net/http"
)

type RetrieverClient interface {
	GetDocuments(input requests.RagRequest) ([]string, error)
}

func NewRetrieverClient() RetrieverClient {
	return &retrieverClient{}
}

type retrieverClient struct {
}

func (r *retrieverClient) GetDocuments(input requests.RagRequest) ([]string, error) {
	retrieveReq, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	retrieveResp, err := http.Post(constant.RetrieveApi, "application/json", bytes.NewBuffer(retrieveReq))
	if err != nil {
		return nil, err
	}
	defer retrieveResp.Body.Close()
	var retrieverResponse response.RetrieverResponse
	if err := json.NewDecoder(retrieveResp.Body).Decode(&retrieverResponse); err != nil {
		return nil, err
	}
	return retrieverResponse.Documents, nil
}
