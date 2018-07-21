FROM golang:alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

#RUN go get -u github.com/micro/micro && \
#    go get github.com/micro/protobuf/proto && \
#    go get -u github.com/micro/protobuf/protoc-gen-go

RUN go get github.com/micro/micro
RUN go get github.com/micro/protobuf/proto
RUN go get -u github.com/micro/protobuf/protoc-gen-go

ENV SOURCES /go/src/github.com/ragagit/Communication/Go-Micro/
COPY . ${SOURCES}

RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}server/
CMD ${SOURCES}server/server
