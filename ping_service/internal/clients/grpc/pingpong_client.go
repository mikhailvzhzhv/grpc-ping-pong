package grpc

import (
	"fmt"

	"github.com/mikhailvzhzhv/grpc-ping-pong/proto/generated/pingpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewPingPongClient(host string, port string) (pingpong.PingPongerClient, error) {
	addrs := fmt.Sprintf("%s:%s", host, port)

	conn, err := grpc.NewClient(addrs, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return pingpong.NewPingPongerClient(conn), nil
}
