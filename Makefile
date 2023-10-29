generate-go:
	protoc -I./proto --go-grpc_out=. --go_out=. ./proto/info.proto