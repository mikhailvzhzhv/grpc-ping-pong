package pingpong

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/dto"
	pingpong_dto "github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/dto/pingpong"
	"github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/handlers/pingpong/mapper"
)

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) DoPingPong(ctx fiber.Ctx) error {
	req := new(pingpong_dto.PingRequest)

	if err := ctx.Bind().JSON(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{Message: "cannot parse json"})
	}

	ping := mapper.ToPing(req)

	pong, err := h.service.DoPingPong(ctx.Context(), ping)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&dto.ErrorResponse{Message: err.Error()})
	}

	res := mapper.ToPongResponse(pong)

	return ctx.Status(fiber.StatusOK).JSON(res)
}
