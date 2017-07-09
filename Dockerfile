FROM golang:latest

MAINTAINER showers suyulingxm@163.com

RUN go get github.com/lifei6671/gocaptcha

COPY app.go /home/gocaptcha
COPY fonts /home/gocaptcha

EXPOSE 8080

ENTRYPOINT go run /home/gocaptcha/app.go




# RUN go run app.go



