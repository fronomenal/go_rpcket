build:
	protoc --go_out=./protos/v1/ --go-grpc_out=./protos/v1/ protos/v1/rocket.proto