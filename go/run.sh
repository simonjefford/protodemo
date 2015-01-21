#!/bin/bash
mkdir -p fourth_people
protoc --go_out=fourth_people --proto_path=../proto/ ../proto/*.proto
go run *.go
