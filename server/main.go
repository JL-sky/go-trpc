package main

import (
	"context"
	"github.com/go-demo/go-trpc/pb"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	//设置server配置文件路径，默认在./trpc_go.yaml
	trpc.ServerConfigPath = "/home/pcl/projects/test/go/server/trpc_go.yaml"
	s := trpc.NewServer()
	pb.RegisterGreeterService(s, &Greeter{})
	if err := s.Serve(); err != nil {
		log.Error(err)
	}
}

type Greeter struct{}

// Hello API
// 1. 接受client请求并打印
// 2. 拼接Hello后作为响应返回给client
func (g Greeter) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Infof("got hello request: %s", req.Msg)
	return &pb.HelloReply{Msg: "Hello " + req.Msg + "!"}, nil
}

