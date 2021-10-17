#!/bin/sh

ssh root@cat "cd ~/up && docker-compose pull ngrok && docker stop ngrok && docker-compose up -d ngrok"