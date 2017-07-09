FROM golang:latest

MAINTAINER showers suyulingxm@163.com

RUN go get github.com/lifei6671/gocaptcha

WORKDIR /go/src/gocaptcha

COPY app.go /go/src/gocaptcha/
COPY fonts/* /go/src/gocaptcha/fonts/


EXPOSE 3000

 ENTRYPOINT go run app.go


# DOCKER_OPTS="--registry-mirror=suyulin.com:50000"



