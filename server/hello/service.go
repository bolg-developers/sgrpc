package hello

import (
	"context"
	"log"
)

type Service struct {
}

func (s *Service) GetHello(context.Context, *Empty) (*GetHelloRes, error) {
	log.Println("サーバーのGetHello関数が実行されました！")
	return &GetHelloRes{Msg: "夢に、嚙みつけ。"}, nil
}
