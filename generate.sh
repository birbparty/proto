#!/bin/sh

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    birb/birb_server.proto

sed -i '' 's/mustEmbedUnimplementedBirbServer/MustEmbedUnimplementedBirbServer/g' birb/birb_server_grpc.pb.go
