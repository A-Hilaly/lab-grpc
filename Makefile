compile-protos:
	protoc \
		--proto_path=api/proto/v1 \
		--proto_path=third_party \
		--go_out=plugins=grpc:pkg/api/v1 \
		todo-service.proto

run-server:
	go run cmd/server/main.go

build-server:
	go build -o server cmd/server/main.go 

build-client:
	go build -o server cmd/server/client.go

build-server-image:
	docker build \
		-t grpc-server \
		.