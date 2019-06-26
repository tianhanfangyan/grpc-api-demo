
RELEASE?=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "latest")
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD?=$(shell date +"%Y-%m-%dT%H:%M:%S%:z")
PROJECT=github.com/tianhanfangyan/grpc-api-demo

.PHONY: vendor
vendor:
	go mod vendor

grpc-gen:
	protoc --proto_path=api/proto \
        		--proto_path=third_party \
        		--go_out=plugins=grpc:api/proto \
        		api/proto/*.proto

grpc-gateway:
	protoc --proto_path=api/proto \
               		--proto_path=third_party \
               		--grpc-gateway_out=logtostderr=true:api/proto \
               		api/proto/*.proto


grpc-swagger:
	protoc --proto_path=api/proto \
		    --proto_path=third_party \
		    --swagger_out=logtostderr=true:api/swagger \
		    api/proto/*.proto

grpc: grpc-gen grpc-gateway grpc-swagger

run:
	go build \
            -mod=vendor \
            -ldflags "-s -w \
            -X ${PROJECT}/version.Release=${RELEASE} \
            -X ${PROJECT}/version.Commit=${COMMIT} \
            -X ${PROJECT}/version.Build=${BUILD}" \
            -o ./app \
            .
	./app

