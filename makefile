build-proto:
	protoc --go_out=./protos/v1/ --go-grpc_out=./protos/v1/ protos/v1/rocket.proto

test-client:
	go test .\httpd\client\test\ -tags=e2e -v

run-client: run-server client

run-server:
	go run .\httpd\server\main.go

test-server:
	go test .\httpd\server\test\ -tags=server -v

client:
	go run .\httpd\client\main.go
 
inval-cache:
	go clean -testcache
