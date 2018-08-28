#This is a sample Image
FROM golang:1.10

RUN go get github.com/gorilla/websocket
ADD . /go/src/
WORKDIR src/broker
RUN go build
CMD ./broker