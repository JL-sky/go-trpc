package main

import (
	"context"

	"trpc.group/trpc-go/tnet/log"
	"trpc.group/trpc-go/trpc-go"
	pb "woa.com/cheersjiang/pb"
)

type Greeter struct{}

// Hello API
// 1. 接受client请求并打印
// 2. 拼接Hello后作为响应返回给client
func (g *Greeter) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Infof("got hello request: %s", req.Msg)
	return &pb.HelloReply{Msg: "Hello " + req.Msg + "!"}, nil
}

type AddService struct{}

func (this *AddService) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddReply, error) {
	return &pb.AddReply{Sum: req.A + req.B}, nil
}

func main() {
	//设置server配置文件路径，默认在./trpc_go.yaml
	// trpc.ServerConfigPath = "/home/pcl/projects/test/go/server/trpc_go.yaml"
	s := trpc.NewServer()
	// pb.RegisterGreeterService(s, &Greeter{})
	pb.RegisterAddService(s, &AddService{})
	if err := s.Serve(); err != nil {
		log.Error(err)
	}
}
