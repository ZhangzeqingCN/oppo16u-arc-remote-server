package security_token

import (
	"fmt"
	"google.golang.org/grpc/metadata"
)

func VerifyMetaData(metaDataMap metadata.MD) (ok bool, errMessage string) {
	var appKey string
	var appId string

	if v, ok := metaDataMap["appid"]; ok {
		appId = v[0]
	}

	if v, ok := metaDataMap["appkey"]; ok {
		appKey = v[0]
	}

	if appId != testAppId || appKey != testAppKey {
		return false, fmt.Sprintf("Wrong appId %v or appKey %v", appId, appKey)
	}

	return true, ""
}
