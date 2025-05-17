package mainpage

import (
	"context"
	"github.com/viktor8881/helloworld/generated"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) MainPage(ctx context.Context, in *generated.EmptyRequest) (*generated.AnyTextResponse, error) {
	var dest generated.AnyTextResponse
	dest.Mess = "Hello, world!"
	return &dest, nil
}
