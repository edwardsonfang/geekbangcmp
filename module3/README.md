# 云原生训练营6期作业Module3
## 第一部分 功能介绍
>本次作业主要包含两个文件：main.go和Dockerfile
## 第二部分 构建镜像
    root@cncamp:~/geekbangcmp/module3# docker build -f ./Dockerfile -t tigerhttpserver:v0.1 .
    Sending build context to Docker daemon  6.656kB
    Step 1/9 : FROM golang:1.19 AS builder
     ---> 180567aa84db
    Step 2/9 : MAINTAINER tigerfang
     ---> Using cache
     ---> fbb875bfa881
    Step 3/9 : RUN go env -w GO111MODULE=on
     ---> Using cache
     ---> 16a53baff516
    Step 4/9 : RUN go env -w GOPROXY=https://goproxy.cn,direct
     ---> Using cache
     ---> e8cdbffba2a6
    Step 5/9 : WORKDIR /tiger-httpserver
     ---> Using cache
     ---> 59565abe4af0
    Step 6/9 : COPY *.go /tiger-httpserver/
     ---> Using cache
     ---> ae9a7f3a1a07
    Step 7/9 : RUN go build -v -o tigerhttpserver ./main.go
     ---> Using cache
     ---> 7697d1226d49
    Step 8/9 : EXPOSE 80
     ---> Running in 74e61624fa8a
    Removing intermediate container 74e61624fa8a
     ---> f3136cd517bf
    Step 9/9 : CMD ["/tiger-httpserver/tigerhttpserver"]
     ---> Running in bee1106b42fe
    Removing intermediate container bee1106b42fe
     ---> 29d1e9fa6ac4
    Successfully built 29d1e9fa6ac4
    Successfully tagged tigerhttpserver:v0.1
## 第三部分 测试镜像内应用
    root@cncamp:~/geekbangcmp/module3# docker run -d tigerhttpserver:v0.1
    87b35bd873c46325713b4d56556ff2784f35662f473acd0a6adba990435d3232
    root@cncamp:~/geekbangcmp/module3# docker ps -a
    CONTAINER ID   IMAGE                  COMMAND                  CREATED          STATUS                    PORTS     NAMES
    87b35bd873c4   tigerhttpserver:v0.1   "/tiger-httpserver/t…"   10 seconds ago   Up 9 seconds              80/tcp    mystifying_einstein
    89cd365a9914   nginx                  "/docker-entrypoint.…"   2 days ago       Exited (0) 23 hours ago             kind_colden
    root@cncamp:~/geekbangcmp/module3# curl -H 'Host:tiger.local.ubuntuvm' -H 'Accept-Language: es'  -v  http://172.17.0.2/healthz
    *   Trying 172.17.0.2:80...
    * TCP_NODELAY set
    * Connected to 172.17.0.2 (172.17.0.2) port 80 (#0)
    > GET /healthz HTTP/1.1
    > Host:tiger.local.ubuntuvm
    > User-Agent: curl/7.68.0
    > Accept: */*
    > Accept-Language: es
    > 
    * Mark bundle as not supporting multiuse
    < HTTP/1.1 200 OK
    < Accept: */*
    < Accept-Language: es
    < User-Agent: curl/7.68.0
    < Version: Not Available!
    < Date: Wed, 14 Dec 2022 09:22:40 GMT
    < Content-Length: 2
    < Content-Type: text/plain; charset=utf-8
    < 
    * Connection #0 to host 172.17.0.2 left intact