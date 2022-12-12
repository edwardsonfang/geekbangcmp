# 云原生训练营6期作业Module3
## 第一部分 功能介绍
## 第二部分 创建镜像
root@cncamp:~/geekbangcmp/module3# docker build -f ./Dockerfile -t tigerhttpserver:v0.1 .
Sending build context to Docker daemon  4.608kB
Step 1/11 : FROM golang:1.19 AS builder
1.19: Pulling from library/golang
Digest: sha256:04f76f956e51797a44847e066bde1341c01e09054d3878ae88c7f77f09897c4d
Status: Downloaded newer image for golang:1.19
 ---> 180567aa84db
Step 2/11 : MAINTAINER tigerfang
 ---> Running in ec33418dc6ce
Removing intermediate container ec33418dc6ce
 ---> fbb875bfa881
Step 3/11 : RUN go env -w GO111MODULE=on
 ---> Running in da7e424b2ac0
Removing intermediate container da7e424b2ac0
 ---> 16a53baff516
Step 4/11 : RUN go env -w GOPROXY=https://goproxy.cn,direct
 ---> Running in 546b326c6e66
Removing intermediate container 546b326c6e66
 ---> e8cdbffba2a6
Step 5/11 : WORKDIR /tiger-httpserver
 ---> Running in d6116ab896b9
Removing intermediate container d6116ab896b9
 ---> 59565abe4af0
Step 6/11 : COPY ./main.go /tiger-httpserver/
 ---> ae9a7f3a1a07
Step 7/11 : RUN go build -v -o tigerhttpserver ./main.go
 ---> Running in eb9d2aa125ce
command-line-arguments
Removing intermediate container eb9d2aa125ce
 ---> 7697d1226d49
Step 8/11 : FROM alpine
latest: Pulling from library/alpine
c158987b0551: Pull complete 
Digest: sha256:8914eb54f968791faf6a8638949e480fef81e697984fba772b3976835194c6d4
Status: Downloaded newer image for alpine:latest
 ---> 49176f190c7e
Step 9/11 : WORKDIR /tiger-httpserver/
 ---> Running in e1b928febbc5
Removing intermediate container e1b928febbc5
 ---> 44ab270e5b0c
Step 10/11 : EXPOSE 80
 ---> Running in 3a9e9ba9b4bb
Removing intermediate container 3a9e9ba9b4bb
 ---> 871e1dcfa23c
Step 11/11 : CMD ["./tigerhttpserver"]
 ---> Running in 3e25c1428b03
Removing intermediate container 3e25c1428b03
 ---> bdb910800794
Successfully built bdb910800794
Successfully tagged tigerhttpserver:v0.1