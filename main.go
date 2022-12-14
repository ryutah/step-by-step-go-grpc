package main

import (
	"context"
	"fmt"
	"net"

	"github.com/ryutah/step-by-step-go-grpc/helloworld"
	"google.golang.org/grpc"
)

// リクエストを受け取るサーバの実装をする

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	// 最高にクールな実装をここにする
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("Hello: %s", req.GetName()),
	}, nil
}

func (s *server) mustEmbedUnimplementedGreeterServer() {}

// バックエンドを起動するメイン関数を定義する

func main() {
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
