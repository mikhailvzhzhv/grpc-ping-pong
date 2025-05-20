package grpc

import (
	"fmt"
	"log"
	"net"

	pb "github.com/mikhailvzhzhv/grpc-ping-pong/proto/generated/pingpong"
	"google.golang.org/grpc"
)

type App struct {
	port       string
	grpcServer *grpc.Server
}

func New(srv pb.PingPongerServer, port string) *App {
	port = fmt.Sprintf(":%s", port)

	grpcServer := grpc.NewServer()
	pb.RegisterPingPongerServer(grpcServer, srv)

	return &App{
		port:       port,
		grpcServer: grpcServer,
	}
}

func (app *App) Run() error {
	listener, err := net.Listen("tcp", app.port)
	if err != nil {
		return err
	}

	log.Printf("run service on port: %s", app.port)

	if err = app.grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}
