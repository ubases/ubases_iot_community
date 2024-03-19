@echo off
rem 执行本脚本之前，请将本目录下的google文件夹拷贝到GOPATH下的include目录下，若原来有google文件夹，请直接覆盖
for %%i in ("./*.proto") do (
    echo %%i
    echo %%~ni
    protoc --proto_path=./ --micro_out=../protosService --go_out=../protosService %%i
    protoc-go-inject-tag -input ../protosService/%%~ni.pb.go
    protoc --proto_path=./ --go_out=../protosService  --micro_out=../protosService  %%i

    rem 生成grpc网关文件
    rem protoc --go_out=./gateway %%i
    rem protoc-go-inject-tag -input ./gateway/%%~ni.pb.go
    rem protoc --go-grpc_out=./gateway --swagger_out=./gateway --grpc-gateway_out=./gateway %%i
)
echo "ext 生成完成"
