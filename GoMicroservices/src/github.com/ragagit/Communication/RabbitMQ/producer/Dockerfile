FROM golang:alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get github.com/streadway/amqp

ENV SOURCES /go/src/github.com/ragagit/Communication/RabbitMQ/
COPY . ${SOURCES}

RUN cd ${SOURCES}producer/ && CGO_ENABLED=0 go build

ENV BROKER_ADDR amqp://guest:guest@localhost:5672/

WORKDIR ${SOURCES}producer/
CMD ${SOURCES}producer/producer
