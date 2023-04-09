package security_token

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
)

const (
	testAppId  = "zzqcnTestApp"
	testAppKey = "123"
)

type ClientTokenCredentials struct {
	credentials.PerRPCCredentials
}

func (c ClientTokenCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	/* 获取客户端的元数据，context用于控制超时和取消，uri是请求入口的uri */
	return map[string]string{
		"appId":  testAppId,
		"appKey": testAppKey,
	}, nil
}

func (c ClientTokenCredentials) RequireTransportSecurity() bool {
	/* 是否要加上tls验证 */
	return false
}
