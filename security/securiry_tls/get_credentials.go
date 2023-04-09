package securiry_tls

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func GetClientDialOptions() (opts []grpc.DialOption, err error) {
	transportCredentials, err := credentials.NewClientTLSFromFile(
		"C:\\Users\\ZZQ\\Documents\\GolandProjects\\grpc_test\\tls\\test.pem",
		"*.zzqcn.com")
	opts = append(opts, grpc.WithTransportCredentials(transportCredentials))
	return
}

func GetServerCredentials() (transportCredentials credentials.TransportCredentials, err error) {
	transportCredentials, err = credentials.NewServerTLSFromFile(
		"C:\\Users\\ZZQ\\Documents\\GolandProjects\\grpc_test\\tls\\test.pem",
		"C:\\Users\\ZZQ\\Documents\\GolandProjects\\grpc_test\\tls\\test.key")
	return
}

func VerifyMetaData(metaDataMap metadata.MD) (ok bool, errMessage string) {
	return true, ""
}
