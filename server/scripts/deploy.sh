#!/bin/sh
GOBIN= GOOS=linux GOARCH=amd64 go build -tags release -ldflags "-s -w" -a -o ngrokd . && docker-compose build && docker-compose push && rm ./ngrokd

