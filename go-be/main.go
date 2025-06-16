package main

import (
	"go-be/handler"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	ragHandler := handler.NewRagHandler()
	app.Post("/rag", ragHandler.RagRequest)
	app.Listen(":8000")
}

// rag system:
// 1. user (userInput) --> 2. knowledge base (userInput, system prompt) --> knowledge
// 3. service_model_inference (userInput, knowledge, system prompt 2) --> response

// workflow:
// go-be: 1 api (rag) 	  --> embeded with retrieve api (python_client_vectorDB) --> inference api
// python-be: 1 api (rag) --> embeded with retrieve api (python_client_vectorDB) --> inference api
