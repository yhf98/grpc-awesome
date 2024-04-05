package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	service "grpc-test/client/proto"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "angular",
		"appkey": "123456",
	}, nil
}
func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	//creds, _ := credentials.NewClientTLSFromFile("./cert/test.pem", "*.aikezc.com")
	//连接服务
	//conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	var opts []grpc.DialOption

	//opts = append(opts, grpc.WithTransportCredentials(creds))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, err := grpc.Dial("127.0.0.1:8080", opts...)
	defer conn.Close()

	if err != nil {
		fmt.Println("连接失败" + err.Error())
	}

	// 创建客户端
	client := service.NewTestHelloClient(conn)
	// 创建请求
	req := &service.HelloRequest{
		Name: "test",
		Age:  18,
	}
	// 调用
	rep, err := client.TestHello(context.Background(), req)

	if err != nil {
		fmt.Println("调用失败" + err.Error())
		return
	}

	fmt.Printf("调用成功,返回值:%s", rep.Message)

}
