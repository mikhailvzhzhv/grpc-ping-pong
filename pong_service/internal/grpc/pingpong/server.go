package pingpong

import (
	"context"

	"github.com/mikhailvzhzhv/grpc-ping-pong/pong_service/internal/grpc/pingpong/mapper"
	pb "github.com/mikhailvzhzhv/grpc-ping-pong/proto/generated/pingpong"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedPingPongerServer
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) PingPong(ctx context.Context, in *pb.PingRequest) (*pb.PongResponse, error) {
	msgPing := in.GetMessage()

	if msgPing == "" {
		return nil, status.Error(codes.DataLoss, "message is empty")
	}

	ping := mapper.ToPing(in)

	pong, err := h.service.PingPong(ctx, ping)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return mapper.ToPongResponse(pong), nil
}
