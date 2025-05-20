package mapper

import (
	pingpong_domain "github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/domain/pingpong"
	"github.com/mikhailvzhzhv/grpc-ping-pong/proto/generated/pingpong"
)

func ToPingRequest(ping *pingpong_domain.Ping) *pingpong.PingRequest {
	return &pingpong.PingRequest{Message: ping.Message}
}

func ToPong(res *pingpong.PongResponse) *pingpong_domain.Pong {
	return &pingpong_domain.Pong{Message: res.Message}
}
