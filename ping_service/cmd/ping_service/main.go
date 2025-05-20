package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	fiber_app "github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/app/fiber"
	"github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/clients/grpc"
	"github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/config"
	pingpong_handler "github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/handlers/pingpong"
	"github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/routes"
	pingpong_service "github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/service/pingpong"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("config load: %v", err)
	}

	pingpongClient, err := grpc.NewPingPongClient(conf.GRPC_HOST, conf.GRPC_PORT)
	if err != nil {
		log.Fatalf("pingpong grpc client: %v", err)
	}

	pingpongService := pingpong_service.New(pingpongClient)
	pingpongHandler := pingpong_handler.New(pingpongService)

	fiberApp := fiber.New()
	routes.NewPingPongRouter(fiberApp, pingpongHandler)

	app := fiber_app.New(fiberApp, conf.HTTP_PORT)

	if err = app.Run(); err != nil {
		log.Fatalf("run app: %v", err)
	}
}
