# Golang RPCket

A golang microservice implemented using gRPC consisting of a client services and a server service.

## Technologies
* grpc
* sqlite
* make

### Stack
Project is created with: 
* golang

### Binaries
The following binaries were required for generating mocks and creating protobuf
* github.com/golang/mock/mockgen
* google.golang.org/protobuf/cmd/protoc-gen-go@latest
* google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

### Packages
Project uses the following packages: 
* gomock
* mockgen
* jmoiron/sqlx
* stretchr/testify
* go-migrate
* mattn/go-sqlite3

## Setup

### Make
1. Navigate to the root directory
2. Run the ff make commands for the desired outcome:
 * `run-client -j2`: provisions the client(:8080) and server(:515151) services
 * **optional**: after running above command, run `test-client` in **another** terminal for e2e test
3. Client service is now exposed and can be queried

### Manual
* In the root directory
  * The main package for the client service is available at ./httpd/client/
  * The main package for the server service is available at ./httpd/server/
* Run or Build the main src files. **client service depends on server service**
* Tests can also be run from above dirs.

### Client service API doc
Listens on port 8080

| resource                    | endpoint | method | params                             | response type              |
|-----------------------------|----------|--------|------------------------------------|----------------------------|
| Retrieve a rocket via id    | /rocket  | get    | query: id<int>                     | json{default}; html; plain |
| Store a new rocket          | /rocket  | post   | name<str>; type<str>; flights<int> | json{default}; html; plain |
| Remove a rocket via id      | /rocket  | delete | query: id<int>                     | json                       |
