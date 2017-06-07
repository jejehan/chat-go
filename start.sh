#! /bin/bash
clear

echo "1. Stop container chat"
docker stop chat

echo "2. Remove container chat"
docker rm chat

echo "3. Run container chat"
docker run -ti --name chat -v $(pwd):/go/src/chat-go -p 9090:9090 chat-go
