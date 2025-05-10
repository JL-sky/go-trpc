package main

import (
	"context"
	"errors"
	"fmt"

	pb "woa.com/cheersjiang/pb"

	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)

func Greeter() {
	// 创建客户端，并请求8080端口的服务
	c := pb.NewGreeterClientProxy(client.WithTarget("ip://127.0.0.1:8000"))
	// 向服务端发送请求: world
	rsp, err := c.Hello(context.Background(), &pb.HelloRequest{Msg: "world"})
	if err != nil {
		log.Error(err)
	}
	// 打印服务端返回结果
	log.Info(rsp.Msg)
}

func Add(num1, num2 int32) (int32, error) {
	c := pb.NewAddClientProxy(client.WithTarget("ip://127.0.0.1:8000"))
	rsp, err := c.Add(context.Background(), &pb.AddRequest{A: num1, B: num2})
	if err != nil {
		log.Error("Add service call error:%v", err)
		return -1, fmt.Errorf("Add service call error:%w", err)
	}

	// 防御性检查：确保响应不为空
	if rsp == nil {
		err := errors.New("加法服务返回空响应")
		log.Error(err)
		return 0, err
	}

	return rsp.Sum, nil
}

func main() {
	// Greeter()
	res, err := Add(1, 2)
	if err != nil {
		log.Fatalf("计算失败: %v", err)
	}
	fmt.Printf("计算结果: %d\n", res)
}
