package pingpong

import (
	"context"

	pingpong_domain "github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/domain/pingpong"
	"github.com/mikhailvzhzhv/grpc-ping-pong/ping_service/internal/service/pingpong/mapper"
	"github.com/mikhailvzhzhv/grpc-ping-pong/proto/generated/pingpong"
)

type Service struct {
	pingpongerClient pingpong.PingPongerClient
}

func New(client pingpong.PingPongerClient) *Service {
	return &Service{pingpongerClient: client}
}

func (s *Service) DoPingPong(ctx context.Context, ping *pingpong_domain.Ping) (*pingpong_domain.Pong, error) {
	pbPingpongRequest := mapper.ToPingRequest(ping)

	pbPingpongResponse, err := s.pingpongerClient.PingPong(ctx, pbPingpongRequest)
	if err != nil {
		return nil, err
	}

	pong := mapper.ToPong(pbPingpongResponse)

	return pong, nil
}
