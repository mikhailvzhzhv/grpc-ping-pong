package mapper

import (
	"github.com/mikhailvzhzhv/grpc-ping-pong/pong_service/internal/domain/models"
	pb "github.com/mikhailvzhzhv/grpc-ping-pong/proto/generated/pingpong"
)

func ToPing(req *pb.PingRequest) *models.Ping {
	return &models.Ping{
		Message: req.Message,
	}
}

func ToPongResponse(pong *models.Pong) *pb.PongResponse {
	return &pb.PongResponse{
		Message: pong.Message,
	}
}
