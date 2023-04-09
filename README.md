# oppo16u-arc-remote-server

# 生成grpc代码

1. 进入到`protos`文件夹
2. 执行命令`protoc --go_out=. hello.proto`
3. 执行命令`protoc --go-grpc_out=. hello.proto`

> 需要提前安装protoc

# 安全认证

1. 生成私钥`openssl genrsa -out server.key 2048`
2. 生成证书`openssl req -new -x509 -key server.key -out server.crt -days 36500`
3. 生成csr`openssl req -new -key server.key -out server.csr`
4. 生成证书私钥test.key`openssl genpkey -algorithm RSA -out test.key`
5. 通过私钥test.key生成证书请求文件test.csr`openssl req -new -nodes -key test.key -out test.csr -days 3650 -subj "/C=cn/OU=org/CN=name" -config ./openssl.cfg -extensions v3_req`
6. 生成SAN证书 pem`openssl x509 -req -days 365 -in test.csr -out test.pem -CA server.crt -CAkey server.key -CAcreateserial -extfile ./openssl.cfg -extensions v3_req` 