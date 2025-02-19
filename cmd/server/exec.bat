@echo off
set PB_DIR=./../../pb

if not exist "%PB_DIR%" mkdir "%PB_DIR%"
protoc ./../../proto/*.proto --go_out=./../../ --go-grpc_out=./../../ --proto_path=./../../proto