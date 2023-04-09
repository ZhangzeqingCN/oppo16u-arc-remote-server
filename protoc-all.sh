@rem 1. 进入到`protos`文件夹
@rem 2. 执行命令`protoc --go_out=. hello.proto`
@rem 3. 执行命令`protoc --go-grpc_out=. hello.proto`

@rem protoc.exe

for file in *.proto; do
    protoc --proto_path=. --python_out=path/to/output/dir "$file"
done
