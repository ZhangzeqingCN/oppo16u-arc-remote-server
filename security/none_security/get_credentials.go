package none_security

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func GetClientDialOptions() (opts []grpc.DialOption, err error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return
}

func GetServerCredentials() (transportCredentials credentials.TransportCredentials, err error) {
	transportCredentials = insecure.NewCredentials()
	return
}
