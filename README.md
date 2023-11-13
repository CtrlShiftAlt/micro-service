#### 下载安装 proto 编译器
* https://github.com/protocolbuffers/protobuf/releases 

#### 安装 golang 插件
* go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
* go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

#### 编写 helloworld/helloworld.proto 文件
```
* protoc 
* --proto_path/-I=pb proto 文件目录
* --go_out=service service 为 生成的目录，以上截图可以看到最终生成的service下存在一个.go文件
* --go-grpc_out=
* --go-grpc_opt=
* pb/service.proto
```
#### 执行，
* .\protoc-25.0-win64\bin\protoc.exe --proto_path=pb --go_out=service --go_opt=paths=source_relative --go-grpc_out=service --go-grpc_opt=paths=source_relative pb\service.proto

#### 生成了 helloworld_grpc.pb.go 和 helloworld.pb.go 两个文件

### 服务端

### 客户端


#### PHP
* .\protoc-25.0-win64\bin\protoc.exe --proto_path=pb --php_out=service pb/service.proto
