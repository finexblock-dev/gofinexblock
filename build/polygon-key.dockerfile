FROM golang:1.20-alpine as build
ENV TZ=Asia/Seoul
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add gcc
RUN apk --no-cache --update add build-base
RUN apk update && apk add --no-cache bash cmake gcc git make

RUN export GOROOT=/usr/local/go
RUN export GOPATH=$HOME/go
RUN export PATH=$PATH:$(go env GOPATH)/bin

WORKDIR /build
COPY . /build

RUN go mod download
RUN go mod vendor

RUN make polygon-key

RUN chmod +x /build/init/polygon_key

ENTRYPOINT ["init/polygon_key"]

EXPOSE 50051