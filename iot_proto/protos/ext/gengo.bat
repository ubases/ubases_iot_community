@echo off

for %%i in ("./*.proto") do (
    echo %%i
    echo %%~ni
    protoc --proto_path=./ --micro_out=../protosService --go_out=../protosService %%i
    rem protoc-go-inject-tag -input ../protosService/%%~ni.pb.go
    protoc --proto_path=./ --go_out=../protosService  --micro_out=../protosService  %%i
)

pause

