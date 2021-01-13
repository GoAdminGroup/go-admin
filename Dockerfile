# This file describes the standard way to build GoAdmin image, and using container
#
# Usage:
#
# # Assemble the full dev environment. This is slow the first time.
# docker build -t goadmin:1.0 .
#
# # Mount your source code to container for quick developing:
# docker run -v `pwd`:/home/goadmin -p 3307:3306 --name goadmin -e MYSQL_ROOT_PASSWORD=root -d goadmin:1.0
# docker exec -it goadmin /bin/bash
# # if your local code has been changed ,you can restart the container to take effect
# docker restart goadmin
#  

FROM mysql:latest
MAINTAINER 72326219@qq.com
COPY . /home/goadmin
ENV GOPATH=/root/go:/home/goadmin/ GOPROXY=https://mirrors.aliyun.com/goproxy,https://goproxy.cn,direct PATH=$PATH:/usr/local/go/bin:/root/go/bin
RUN apt-get update  && \
    # apt-get upgrade  && \
    apt-get install -y gcc make gcc-c++ openssl-devel wget git zip nano vim gcc && \
    # current path has already included go1.15.6.linux-amd64.tar and godependacy.tgz
    #mv /home/goadmin/go1.15.6.linux-amd64.tar /tmp && \
    #mv /home/goadmin/godependacy.tgz /tmp && \
    #tar -C /usr/local -xvf /tmp/go1.15.6.linux-amd64.tar && \
    #tar -C /root -xvf /tmp/godependacy.tgz
    go get golang.org/x/tools/cmd/goimports && \
    go get github.com/rakyll/gotest && \
    go get -u golang.org/x/lint/golint && \
    go get -u github.com/golangci/golangci-lint/cmd/golangci-lint && \
    # need port to expose, to do
    #apt-get install -y postgresql && \
    #cd /home/goadmin/ && make fmt && make test;exit 0
WORKDIR /home/goadmin