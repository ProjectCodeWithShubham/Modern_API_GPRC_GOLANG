# Modern_API_GPRC_GOLANG


PS D:\GRPC\Modern_API_GPRC_GOLANG> protoc --version
libprotoc 3.12.4

 go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
 grpc

 PS D:\GRPC\Modern_API_GPRC_GOLANG> go mod init github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG        
go: creating new go.mod: module github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG

go get -u google.golang.org/grpc

PS D:\GRPC\Modern_API_GPRC_GOLANG> go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go: added google.golang.org/protobuf v1.28.1
PS D:\GRPC\Modern_API_GPRC_GOLANG> go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go: added google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0



go mod init github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG
go: D:\GRPC\Modern_API_GPRC_GOLANG\go.mod already exists

protoc -Igreet/proto --go_out=. --go-grpc_out=. greet/proto/dummy.proto
-----> crated file inside github folder pb and grpc.pb


second way to create in proto folder only

By using this command we can create both pb files in proto folder
-->  protoc -Igreet/proto --go_out=. --go_opt=module=github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG  --go-grpc_out=. --go-grpc_opt=module=github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG greet
/proto/dummy.proto

proto ----> dummy_grpc_pb.go , dummy.pb.go

make clean
make greet ------>  create files automatically 

choco install make

make help


make greet

go mod tidy
-----> import expected packages 

after make greet start server by using generated binary file executable 

.\bin\greet\server.exe
2023/06/18 23:31:44 Listening on 0.0.0.0:50051

Server has started now client requried 

Client setup

after client file -----> binary client file created 

PS D:\GRPC\Modern_API_GPRC_GOLANG> .\bin\greet\client.exe
2023/06/18 23:44:48 Failed to connect: grpc: no transport security set (use grpc.WithTransportCredentials(insecure.NewCredentials()) explicitly or set credentials) 

add

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))


2023/06/18 23:49:39 Client is running.......
