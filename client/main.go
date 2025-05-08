package main

import (
	"context"
	"github.com/go-demo/go-trpc/pb"

	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	//创建客户端，并请求8080端口的服务
	c := pb.NewGreeterClientProxy(client.WithTarget("ip://127.0.0.1:8000"))
	//向服务端发送请求: world
	rsp, err := c.Hello(context.Background(), &pb.HelloRequest{Msg: "world"})
	if err != nil {
		log.Error(err)
	}
	//打印服务端返回结果
	log.Info(rsp.Msg)
}

