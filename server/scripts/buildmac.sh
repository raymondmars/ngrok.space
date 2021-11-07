#!/bin/sh
GOBIN= GOOS=darwin GOARCH=amd64 go build -tags release -ldflags "-s -w" -a -o ../deploy/bin/ngrokd . 

