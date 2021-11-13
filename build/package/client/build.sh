#!/usr/bin/env bash

echo "build client for $1" 
cd $2
case $1 in  
    linux|LINUX)  
      GOBIN= GOOS=linux GOARCH=amd64 go build -tags release -ldflags "-s -w" -a -o $3/linux_amd64/ngrok $2
      ;;  
    macos|MACOS)  
      GOBIN= GOOS=darwin GOARCH=amd64 go build -tags release -ldflags "-s -w" -a -o $3/macos_amd64/ngrok $2 . .
      ;;  
    windows|WINDOWS)
      GOBIN= GOOS=windows GOARCH=amd64 go build -tags release -ldflags "-s -w" -a -o $3/windows_amd64/ngrok.exe $2 .
      ;;
    linux-arm|LINUX-ARM)
      GOBIN= GOOS=linux GOARCH=arm GOARM=5 go build -tags release -ldflags "-s -w" -a -o $3/linux_arm/ngrok $2 .
      ;;
esac


