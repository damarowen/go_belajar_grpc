.PHONY: protos run client

protos:
	 protoc proto/user.proto --go-grpc_out=. --go_out=.

run_server:
	cd server/go run main.go

run_client:
	cd client/go run main.go
