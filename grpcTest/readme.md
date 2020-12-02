# 步骤
1. 创建 search.proto 文件
2. 运行` protoc --go_out=plugins=grpc:./ ./proto/search.proto` 生成search.pb.go
    > protoc --go_out=plugins=grpc:{输出目录}  {proto文件}
3. 创建 server.go 实现search.proto 里面的方法
4. 创建 client.go 去使用search.proto 里面的方法

