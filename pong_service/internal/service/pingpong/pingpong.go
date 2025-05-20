package pingpong

import (
	"context"
	"fmt"
	"strings"

	"github.com/mikhailvzhzhv/grpc-ping-pong/pong_service/internal/domain/models"
)

const (
	pingPrefix = "ping"
	pongPrefix = "pong"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) PingPong(ctx context.Context, ping *models.Ping) (*models.Pong, error) {
	msgPing := ping.Message

	if !strings.HasPrefix(msgPing, pingPrefix) {
		return nil, fmt.Errorf("message(%s) has not prefix(%s)", msgPing, pingPrefix)
	}

	msgPing, _ = strings.CutPrefix(msgPing, pingPrefix)
	msgPong := fmt.Sprint(pongPrefix, msgPing)

	return &models.Pong{Message: msgPong}, nil
}
