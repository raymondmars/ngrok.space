#!/usr/bin/env bash

cd $1/linux_amd64/

docker build -f $2/Dockerfile -t $3 .