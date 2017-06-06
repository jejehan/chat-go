#! /bin/bash
clear

echo "1. Stop container chat"
docker stop chat

echo "2. Remove container chat"
docker rm chat

echo "3. Run container chat"
docker run -ti --name chat -v /Users/macbook/Documents/Docker/chat-go:/src/go/chat-go -p 9090:9090 chat-go
