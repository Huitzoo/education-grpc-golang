#!/bin/bash

protoc ./calculatorpb/calculator.proto --go_out=plugins=grpc:.

rm -r go.mod
rm -r go.sum

go mod init calculator
go mod vendor
go build .