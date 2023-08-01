FROM golang:1.20-alpine as build
ENV TZ=Asia/Seoul
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add gcc
RUN apk --no-cache --update add build-base
RUN apk update && apk add --no-cache bash cmake gcc git make




FROM build as release

WORKDIR /home

COPY . /home/go/src/finexblock

WORKDIR /home/go/src/finexblock

RUN export GOROOT=/usr/local/go
RUN export GOPATH=$HOME/go
RUN export PATH=$PATH:$GOROOT/bin:/usr/local/bin:$GOPATH/bin
RUN export PATH=$PATH:$(go env GOPATH)/bin

RUN go mod download
RUN go mod vendor

RUN make polygon

ENTRYPOINT ["init/polygond"]

EXPOSE 50051