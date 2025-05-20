package pingpong

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/mikhailvzhzhv/grpc-ping-pong/pong_service/pkg/generated/pingpong"
)

const (
	pingPrefix = "ping"
	pongPrefix = "pong"
)

type Service struct{}

func (s *Service) PingPong(ctx context.Context, req *pb.PingRequest) (*pb.PongResponse, error) {
	msg := req.GetMessage()

	if !strings.HasPrefix(msg, pingPrefix) {
		return nil, fmt.Errorf("message has not pefix: %s", pingPrefix)
	}

	msg, _ = strings.CutPrefix(msg, pingPrefix)
	msg = fmt.Sprint(pongPrefix, msg)

	return &pb.PongResponse{Message: msg}, nil
}
