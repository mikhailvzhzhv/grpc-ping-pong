package main

import (
	"log"

	grpcapp "github.com/mikhailvzhzhv/grpc-ping-pong/pong_service/internal/app/grpc"
	"github.com/mikhailvzhzhv/grpc-ping-pong/pong_service/internal/config"
	handler_pingpong "github.com/mikhailvzhzhv/grpc-ping-pong/pong_service/internal/grpc/pingpong"
	service_pingpong "github.com/mikhailvzhzhv/grpc-ping-pong/pong_service/internal/service/pingpong"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("load config error: %v", err)
	}

	pingpongService := service_pingpong.New()
	pingpongHandler := handler_pingpong.New(pingpongService)

	app := grpcapp.New(pingpongHandler, conf.PORT)

	if err = app.Run(); err != nil {
		log.Fatalf("app run error: %v", err)
	}
}
