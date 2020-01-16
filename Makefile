
RELEASE?=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "latest")
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD?=$(shell date +"%Y-%m-%dT%H:%M:%S%:z")
PROJECT=github.com/tianhanfangyan/grpc-api-demo

GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

.PHONY: vendor
vendor:
	go mod vendor

fmt:
	@gofmt -w ${GOFILES}

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

.PHONY: server
server:
	go build \
            -mod=vendor \
            -ldflags "-s -w \
            -X ${PROJECT}/version.Release=${RELEASE} \
            -X ${PROJECT}/version.Commit=${COMMIT} \
            -X ${PROJECT}/version.Build=${BUILD}" \
            -o ./server-app \
            ./server/server.go
	./server-app

.PHONY: client
client:
	go build \
            -mod=vendor \
            -ldflags "-s -w \
            -X ${PROJECT}/version.Release=${RELEASE} \
            -X ${PROJECT}/version.Commit=${COMMIT} \
            -X ${PROJECT}/version.Build=${BUILD}" \
            -o ./client-app \
            ./client/client.go
	./client-app

doc: grpc-swagger
	docker run -d --name grpc-api-doc \
           -p 8080:8080 \
           -v ${PWD}/api/swagger:/data \
           -e "SWAGGER_JSON=/data/api-demo.swagger.json" \
           -e "VALIDATOR_URL=null" \
           -e 'DEFAULT_MODELS_EXPAND_DEPTH=-1' \
           swaggerapi/swagger-ui:v3.22.0
