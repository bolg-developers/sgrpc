package hello

import "context"

type Service struct {
}

func (s *Service) GetHello(context.Context, *Empty) (*GetHelloRes, error) {
	return &GetHelloRes{Msg: "夢に、嚙みつけ。"}, nil
}
