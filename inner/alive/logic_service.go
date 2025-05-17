package alive

import (
	"context"
	"github.com/viktor8881/helloworld/generated"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Alive(ctx context.Context, in *generated.EmptyRequest) (*generated.EmptyResponse, error) {
	var dest generated.EmptyResponse
	return &dest, nil
}
