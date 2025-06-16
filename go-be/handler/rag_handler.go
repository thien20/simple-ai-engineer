package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-be/constant"
	"go-be/requests"
	"go-be/response"
	"go-be/utils"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type RagHandler struct {
}

// NewRagHandler initializes a new RagHandler instance - aka constructor
func NewRagHandler() *RagHandler {
	return &RagHandler{}
}

func (rag *RagHandler) RagRequest(ctx fiber.Ctx) error {
	var input requests.RagRequest
	if err := utils.BodyParser(ctx, &input); err != nil {
		return err
	}
	// Call the RAG service with the user input
	retrieverReq, _ := json.Marshal(map[string]string{"user_input": input.UserInput})
	log.Print("Sending request to retriever service with input: ", input.UserInput)
	retrieverResp, err := http.Post(constant.RetrieveApi, "application/json", bytes.NewBuffer(retrieverReq))
	if err != nil {
		return err
	}
	defer retrieverResp.Body.Close()

	var retriever response.RetrieverResponse
	if err := utils.BodyParser(ctx, &retriever); err != nil {
		return err
	}

	prompt := fmt.Sprintf("%s: %s\n\n User Query: %s\nAnswer:", constant.SystemPrompt, retriever.Context, input.UserInput)
	ollamaReq := requests.OllamaRequest{Model: "gemma", Prompt: prompt, Stream: false}
	reqBytes, _ := json.Marshal(ollamaReq)

	// call the LLM service with the retriever response
	llmResp, err := http.Post("http://ollama:11434/api/generate", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	defer llmResp.Body.Close()

	var result response.OllamaResponse
	if err := utils.BodyParser(ctx, &result); err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"response": result.Response,
	})
}
