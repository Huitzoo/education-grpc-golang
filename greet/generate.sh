#!/bin/bash

protoc ./greetpb/greet.proto --go_out=plugins=grpc:.

rm -r go.mod
rm -r go.sum
rm -r vendor

go mod init greet
go mod vendor

rm -r vendor

