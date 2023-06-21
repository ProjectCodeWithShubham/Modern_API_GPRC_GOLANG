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





PS D:\GRPC\Modern_API_GPRC_GOLANG> .\bin\calculator\server.exe
2023/06/19 19:43:16 Listening on 0.0.0.0:50051
2023/06/19 19:43:32 Sum Function was invoked with first_number:1  second_number:3


PS D:\GRPC\Modern_API_GPRC_GOLANG> .\bin\calculator\client.exe
2023/06/19 19:43:32 Client is running.......
2023/06/19 19:43:32 doSum was invoked and it is client file 
2023/06/19 19:43:32 Sum: 4




CLient Streaming 

 make greet
protoc -Igreet/proto --go_opt=module=github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG --go_out=. --go-grpc_opt=module=github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG --go-grpc_out=. greet/proto/*.proto
go build -o bin/greet/server.exe ./greet/server
go build -o bin/greet/client.exe ./greet/client
PS D:\GRPC\Modern_API_GPRC_GOLANG>


.\bin\greet\server.exe
2023/06/21 15:24:01 Listening on port address in server main file 0.0.0.0:50051
2023/06/21 15:24:09 Long-Greet function was invoked


.\bin\greet\client.exe
2023/06/21 15:24:08 Client is running.......
2023/06/21 15:24:08 doLongGreet Was invoked
2023/06/21 15:24:09 Sending req:first_name:"Shubham_client_streaming" 
2023/06/21 15:24:11 Sending req:first_name:"ujwal_client_streaming"

2023/06/21 15:24:12 Sending req:first_name:"samyak_client_streaming"

2023/06/21 15:24:13 LongGreet: Hello Shubham_client_streaming!
Hello utkarsh_client_streaming!
Hello ujwal_client_streaming!
Hello samyak_client_streaming!


