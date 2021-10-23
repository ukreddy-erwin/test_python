#!/bin/bash

protoc --go_out=. --go_opt=paths=source_relative excel.proto
protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative excel.proto