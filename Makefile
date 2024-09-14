gen-cal:
	protoc proto/user.proto --go-grpc_out=.
	protoc proto/user.proto --go_out=.
run-server:
	go run internal/intergation/user_server.go