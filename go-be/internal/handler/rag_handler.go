package handler

import (
	"go-be/internal/requests"
	"go-be/internal/service"
	"go-be/utils"

	"github.com/gofiber/fiber/v3"
)

type RagHandler struct {
	ragSvc *service.RagService
}

func NewRagHandler(ragSvc *service.RagService) *RagHandler {
	return &RagHandler{
		ragSvc: service.NewRagService(),
	}
}

func (rag *RagHandler) RagRequest(ctx fiber.Ctx) error {
	var input requests.RagRequest
	if err := utils.BodyParser(ctx, &input); err != nil {
		return err
	}

	result, err := rag.ragSvc.GenerateAns(input)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"response": result.Response,
	})
}
