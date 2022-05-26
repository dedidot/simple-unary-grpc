run-protoc:
	protoc -Iproto/bookpb --go_out=. --go_opt=module=github.com/dedidot/grpc-book --go-grpc_out=. --go-grpc_opt=module=github.com/dedidot/grpc-book proto/bookpb/book.proto

build:
	go build -o .