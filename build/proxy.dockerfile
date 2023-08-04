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

RUN make proxy

RUN chmod +x /build/init/proxy && \
    chmod +x /build/deployments/proxy/run.sh


FROM scratch as release

COPY --from=build /build/init/proxy .
COPY --from=build /build/deployments/proxy/run.sh .

ENTRYPOINT ["run.sh"]

EXPOSE 50051