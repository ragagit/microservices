#FROM golang:1.10.3
FROM golang:alpine
MAINTAINER ragadocker
ENV SOURCES /go/src/github.com/ragagit
ADD . ${SOURCES}
RUN cd ${SOURCES}/ragamicro && CGO_ENABLE=0 go install
RUN go install github.com/ragagit/ragamicro
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ragamicro

#=========== Improved ============
#FROM alpine
#MAINTAINER ragadocker

#COPY ./ragamicro/ragamicro /app/ragamicro
#RUN chmod +x /app/ragamicro

#ENV PORT 8080
#EXPOSE 8080

#ENTRYPOINT /app/ragamicro
