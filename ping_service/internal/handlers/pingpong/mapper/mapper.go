package mapper

import (
	pingpong_domain "github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/domain/pingpong"
	pingpong_dto "github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/dto/pingpong"
)

func ToPing(req *pingpong_dto.PingRequest) *pingpong_domain.Ping {
	return &pingpong_domain.Ping{Message: req.Message}
}

func ToPongResponse(pong *pingpong_domain.Pong) *pingpong_dto.PongResponse {
	return &pingpong_dto.PongResponse{Message: pong.Message}
}
