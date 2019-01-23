FROM golang:1.11-alpine

COPY signaling.go /go/src/signaling.go

RUN apk add --update --no-cache --virtual build-deps \
    git \
    && go get golang.org/x/net/websocket \
    && go build -o /go/src/signaling /go/src/signaling.go \
    && apk del build-deps

CMD ["/go/src/signaling"]
