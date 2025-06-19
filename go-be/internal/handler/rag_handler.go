package handler

import (
	"bytes"
	"encoding/json"
	"fmt"

	"go-be/internal/constant"
	"go-be/internal/requests"
	"go-be/internal/response"
	"go-be/internal/service"
	"go-be/utils"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type RagHandler struct {
	ragSvc *service.RagService
}

func NewRagHandler() *RagHandler {
	return &RagHandler{}
}

func (rag *RagHandler) RagRequest(ctx fiber.Ctx) error {
	var input requests.RagRequest
	if err := utils.BodyParser(ctx, &input); err != nil {
		return err
	}
	ragResp, err := rag.ragSvc.GetDocuments(input)
	if err != nil {
		return err
	}
	// retrieverReq, err := json.Marshal(input)
	// if err != nil {
	// 	log.Print("Error marshalling retriever request: ", err)
	// }
	// log.Print("Sending request to retriever service with input: ", retrieverReq)
	// retrieverResp, err := http.Post(constant.RetrieveApi, "application/json", bytes.NewBuffer(retrieverReq))
	// if err != nil {
	// 	return err
	// }
	// defer retrieverResp.Body.Close()

	// var retrieverResponse response.RetrieverResponse
	// if err := json.NewDecoder(retrieverResp.Body).Decode(&retrieverResponse); err != nil {
	// 	return err
	// }

	prompt := fmt.Sprintf("%s: %s\n\n User Query: %s\nAnswer:", constant.SystemPrompt, ragResp, input.UserInput)
	ollamaReq := requests.OllamaRequest{
		Model:  "gemma:2b-instruct-q4_0",
		Prompt: prompt,
		Stream: false,
	}
	reqBytes, err := json.Marshal(ollamaReq)
	if err != nil {
		log.Print("Error marshalling Ollama request: ", err)
	}

	// call the LLM service with the retriever response
	llmResp, err := http.Post(constant.OllamaApi, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	defer llmResp.Body.Close()

	var result response.OllamaResponse
	if err := json.NewDecoder(llmResp.Body).Decode(&result); err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"response": result.Response,
	})
}
