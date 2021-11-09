#!/usr/bin/env bash

echo $1
cd $1
docker-compose push

#ssh root@cat "cd ~/up && docker-compose pull ngrok && docker stop ngrok && docker-compose up -d ngrok"