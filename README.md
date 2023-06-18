# Modern_API_GPRC_GOLANG


PS D:\GRPC\Modern_API_GPRC_GOLANG> protoc --version
libprotoc 3.12.4

PS D:\GRPC\Modern_API_GPRC_GOLANG> go get google.golang.org/protobuf/cmd/protoc-gen-go
PS D:\GRPC\Modern_API_GPRC_GOLANG> go get google.golang.org/grpc/cmd/protoc-gen-go-grpc




command to create pb, pb2 grpc files 

 protoc -Igreet/proto --go_out=. --go-grpc_out=. greet/proto/dummy.proto

grpc.pb.go
pb.go 

files created 


One more method to add pb2 inside proto folder

protoc -Igreet/proto --go_out=. --go_opt=module=github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG  --go-grpc_out=. --go-grpc_opt=module=github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG greet/proto/dummy.proto 

---> files generated inside proto folder



Added Makefile 
make greet

