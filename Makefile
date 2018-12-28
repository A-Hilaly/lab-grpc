GO111MODULE=on

compile-protos:
	protoc \
		--proto_path=api/proto/v1 \
		--proto_path=third_party \
		--go_out=plugins=grpc:pkg/api/v1 \
		example.proto

run-server:
	go run cmd/server/main.go

build-server:
	go build -o server cmd/server/main.go 

build-client:
	go build -o client cmd/client/main.go

build-server-image:
	GO111MODULE=on GOOS=linux go build -o server cmd/server/main.go
	docker build \
		-t grpc-server \
		.

go-tidy:
	GO111MODULE=on go mod tidy

go-verify:
	GO111MODULE=on go mod verify

go-vendor:
	GO111MODULE=on go mod vendor

go-download:
	GO111MODULE=on go mod download