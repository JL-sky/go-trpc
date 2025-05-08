# go-trpc
go trpc test

# 环境安装

## 安装trpc

### 安装

```bash
go install trpc.group/trpc-go/trpc-cmdline/trpc@latest
```

### 安装验证

```bash
trpc version
```

## 安装protoc

### 安装

```bash
git clone https://github.com/protocolbuffers/protobuf.git
```

```bash
cd protobuf/
./autogen.sh
./configure --prefix=/usr/local/protobuf
make
sudo make install
sudo ldconfig 
```

### 安装验证

```bash
protoc --version
```

# trpc实战

## 初始化项目

初始化go.mod文件

```bash
go mod init github.com/go-demo/go-trpc
```

## 编写proto文件

```protobuf
syntax = "proto3";

package trpc.helloworld;
option go_package="github.com/go-demo/go-trpc/pb";

service Greeter {
  rpc Hello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string msg = 1;
}

message HelloReply {
  string msg = 1;
}
```

编译proto并生成go文件

```bash
trpc create -p helloworld.proto --rpconly --nogomod --mock=false
```

安装依赖

```bash
go mod tidy
```

## 编写服务端

server/main.go

```go
package main

import (
	"context"
	"github.com/go-demo/go-trpc/pb"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	//设置server配置文件路径，默认在./trpc_go.yaml
	trpc.ServerConfigPath = "E:\\Go\\GoPro\\src\\go_code\\ziyifast-code_instruction\\go-demo\\go-trpc\\server\\trpc_go.yaml"
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

```

服务端配置文件 server/trpc_go.yaml

```bash
server:
  service:
    - name: trpc.helloworld
#      监听地址
      ip: 127.0.0.1
#      服务监听端口
      port: 8000
```



## 编写客户端

client/main.go

```go
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
```

