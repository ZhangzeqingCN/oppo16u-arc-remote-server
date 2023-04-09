package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	service "oppo16u-arc-remote-server/protos/go"
	"oppo16u-arc-remote-server/security"
)

type SayHelloServer struct {
	service.UnimplementedSayHelloServer
}

func (server *SayHelloServer) SayHello(ctx context.Context, request *service.HelloRequest) (*service.HelloResponse, error) {
	metaData, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, errors.New("no token")
	}

	ok, message := security.VerifyMetaData(metaData)

	if !ok {
		return &service.HelloResponse{
			ResponseMessage: fmt.Sprintln("Error", message),
		}, nil
	}

	return &service.HelloResponse{
		ResponseMessage: fmt.Sprintf("Hello %v, this is server", request.Name),
	}, nil
}

func main() {
	transportCredentials, err := security.GetServerCredentials()
	if err != nil {
		panic(err.Error())
	}
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err.Error())
	}
	grpcServer := grpc.NewServer(grpc.Creds(transportCredentials))
	service.RegisterSayHelloServer(grpcServer, &SayHelloServer{})
	if err := grpcServer.Serve(listen); err != nil {
		panic(err.Error())
	}
}
