# basic image
FROM golang:1.19 AS builder

# maintainer
MAINTAINER tigerfang

# configure go env
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# set workdir
WORKDIR /tiger-httpserver

# copy code to workdir
COPY ./main.go /tiger-httpserver/

# comply code to bin
RUN go build -v -o tigerhttpserver ./main.go

# build image
FROM alpine

# set work dir
WORKDIR /tiger-httpserver/

# expose port
EXPOSE 80

# start my httpserver
CMD ["./tiger-httpserver/tigerhttpserver"]