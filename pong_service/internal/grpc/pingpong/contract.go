package pingpong

import (
	"context"

	"github.com/mikhailvzhzhv/grpc-ping-pong/pong_service/internal/domain/models"
)

type Service interface {
	PingPong(ctx context.Context, ping *models.Ping) (*models.Pong, error)
}
