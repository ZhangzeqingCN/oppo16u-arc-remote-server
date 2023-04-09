package security

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	securityMethods "grpc_test/security/security_token"
)

func GetDialDialOptions() (opts []grpc.DialOption, err error) {
	return securityMethods.GetClientDialOptions()
}

func GetServerCredentials() (transportCredentials credentials.TransportCredentials, err error) {
	return securityMethods.GetServerCredentials()
}

func VerifyMetaData(metaDataMap metadata.MD) (ok bool, errMessage string) {
	return securityMethods.VerifyMetaData(metaDataMap)
}
