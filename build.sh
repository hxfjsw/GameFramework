#!/usr/bin/env bash

rm -f ./bin/GameFramework
go build -o ./bin/GameFramework ./src/main.go ./src/room.go ./src/server.go