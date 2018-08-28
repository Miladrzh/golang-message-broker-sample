#Dockerfile for subscriber
FROM golang:1.10

RUN go get github.com/gorilla/websocket
ADD . /go/src/MessageQueue
WORKDIR src/MessageQueue/client
RUN go build
WORKDIR /go/src/MessageQueue/subscriber
RUN go build
CMD ./subscriber