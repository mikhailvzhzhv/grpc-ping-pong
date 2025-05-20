package pingpong

import (
	"context"

	pingpong_domain "github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/domain/pingpong"
)

type Service interface {
	DoPingPong(ctx context.Context, ping *pingpong_domain.Ping) (*pingpong_domain.Pong, error)
}
