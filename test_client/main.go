package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	service "oppo16u-arc-remote-server/protos/go"
	"oppo16u-arc-remote-server/security"
)

func main() {
	dialOpts, err := security.GetDialDialOptions()
	if err != nil {
		panic(err.Error())
	}
	clientConn, err := grpc.Dial("127.0.0.1:9090", dialOpts...)
	if err != nil {
		panic(err.Error())
	}
	defer func(clientConn *grpc.ClientConn) {
		err := clientConn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(clientConn)
	client := service.NewSayHelloClient(clientConn)
	response, err := client.SayHello(context.Background(), &service.HelloRequest{
		RequestMessage: "Hello, this is client",
		Name:           "ZZQ",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(response.GetResponseMessage())
}
