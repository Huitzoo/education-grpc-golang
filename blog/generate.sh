#!/bin/bash

protoc ./blogpb/blog.proto --go_out=plugins=grpc:.

rm -r go.mod
rm -r go.sum

go mod init blog
go mod vendor
rm -r vendor
go build .