package pingpong

import pb "github.com/mikhailvzhzhv/grpc-ping-pong/proto/generated/pingpong"

type Handler struct {
	service pb.PingPongerServer
}

func New(service pb.PingPongerServer) *Handler {
	return &Handler{service: service}
}

func (h *Handler) PingPong() {
	h.service.PingPong()
}
