package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	_ "google.golang.org/grpc/status"
	service "grpc-test/server/proto"
	"net"
)

type server struct {
	service.TestHelloServer
}

func (s *server) TestHello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {

	// 获取元数据信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("获取元数据信息失败")
	}
	var appId, appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	if appId != "" {
		fmt.Printf("appId: %s, appKey: %s", appId, appKey)
	}

	return &service.HelloResponse{Message: "REE===EEE " + req.Name}, nil
}

func main() {
	//cerds, err := credentials.NewServerTLSFromFile("./cert/test.pem", "./cert/test.key")

	//if err != nil {
	//	panic(err)
	//}
	// 监听端口
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// 创建grpc server
	//grpcServer := grpc.NewServer(grpc.Creds(cerds))
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	// 注册服务
	service.RegisterTestHelloServer(grpcServer, &server{})

	fmt.Println("rRPC 服务启动成功！")
	// 启动服务
	grpcServer.Serve(lis)

}
