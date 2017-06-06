FROM golang:1.8.3-alpine

ENV SRC_DIR=/go/src/chat-go

# install aplikasi yang kira kira dibutukan, contoh git, curl, vim
RUN set -ex \
        && apk add --no-cache git curl vim nano bash
# buat directory
RUN mkdir -p /go/src/chat-go

# tambahkan seluruh file , ke directory container
ADD . $SRC_DIR

# default ketika akses masuk bash ke container
WORKDIR  $SRC_DIR

# EXPOSING / Port yang akan digunakan
EXPOSE 9090

# Build dulu chatnya
RUN cd $SRC_DIR; go build -o chat

# running aplikasinya
RUN chmod +x chat

ENTRYPOINT ["./chat"]
