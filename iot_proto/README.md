## proto

```go
# protos生成
# device 
protoc --proto_path=./ext --micro_out=./ --go_out=./ casbin.ext.proto

google/protobuf/timestamp.proto
使用方法：


protoc --proto_path=./ext --micro_out=./protosService --go_out=./protosService t_iot_device_home.proto
rotoc --proto_path=./ext --micro_out=./protosService --go_out=./protosService t_iot_device_home.proto


%PROTOBUF_HOME%\Release\protoc -I=. -I=D:/WorkCode/protobuf-3.0.0-beta-1/src –grpc_out=./generate –plugin=protoc-gen-grpc=%GRPC_HOME_RELEASE%\grpc_cpp_plugin.exe ./DataAccessSerivceInterface.proto
protoc -I=. -I=D:\go\google\protobuf-3.20.0\src --proto_path=../../../protos --micro_out=../../../protos --go_out=../../../protos gen/iot_device/t_iot_device_home.proto

```


## 如果提示“google/api/xxx.proto” Not found
1、检查GO环境，GOPATH环境，检查GOPATH下是否存在include/google
2、如果环境配置没问题，但是还是提示报错，可以使用快捷方式导入的方式
```cmd
mklink /D {$projectPath}\iot_proto\protos\gen\iot_app_oem\google {$GOPATH}\include\google

# 例如
mklink /D G:\code\work\bat\cloud_platform\iot_proto\protos\gen\iot_app_oem\google D:\gopath\include\google
mklink /D G:\code\work\bat\cloud_platform\iot_proto\protos\gen\iot_system\google D:\gopath\include\google
# google目录放在\docs\开发指南\第三方包


```