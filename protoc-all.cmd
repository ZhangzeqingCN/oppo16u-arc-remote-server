@echo on
setlocal enabledelayedexpansion

set "PROTO_DIR=%CD%\protos\src"
set "OUTPUT_DIR=%CD%\protos\go"
set "PROTOC_PATH=protoc.exe"

for /r "%PROTO_DIR%" %%a in (*.proto) do (
    @rem echo "%PROTOC_PATH%" --proto_path="%PROTO_DIR%" --python_out="%OUTPUT_DIR%" "%%a"
    "%PROTOC_PATH%" --proto_path="%PROTO_DIR%" --go-grpc_out="%OUTPUT_DIR%" --go_out="%OUTPUT_DIR%" "%%a"
)

echo Done!
@rem pause
