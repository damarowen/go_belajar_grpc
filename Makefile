.PHONY: protos run client

protoc:
	 protoc protos/user.proto --go-grpc_out=. --go_out=.
	 @echo "Protoc compile done ...!"

run_server:
	cd server/cmd && go run main.go

run_client:
	cd client && go run main.go
