package routes

import (
	"github.com/gofiber/fiber/v3"
	handler_pingpong "github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/handlers/pingpong"
)

const (
	pingpongPath = "/pingpong"
)

func NewPingPongRouter(app *fiber.App, handler *handler_pingpong.Handler) {
	app.Post(pingpongPath, handler.DoPingPong)
}
