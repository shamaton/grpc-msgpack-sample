FROM golang:1.17-buster

RUN apt-get update \
 && apt-get install -y protobuf-compiler \
 && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

WORKDIR /work
COPY protobuf-go protobuf-go
COPY schema schema

RUN cd protobuf-go \
 && CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/protoc-gen-go ./cmd/protoc-gen-go/

RUN protoc -I./schema --go_out=/home --go_opt=paths=source_relative --go-grpc_out=/home --go-grpc_opt=paths=source_relative  schema/*.proto