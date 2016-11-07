#!/usr/bin/env bash

rm -f ./bin/GameFramework
go build -o ./bin/GameFramework ./src/main.go ./src/js_vm.go ./src/room.go ./src/server.go ./src/http.go ./src/redis.go
./bin/GameFramework