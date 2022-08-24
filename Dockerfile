# This file describes the standard way to build GoAdmin develop env image, and using container
#
# Usage:
#
# # Assemble the code dev environment, get related database tools in docker-compose.yml, It is slow the first time.
# docker build -t goadmin:1.0 .
#
# # Mount your source code to container for quick developing:
# docker run -v `pwd`:/home/goadmin --name -d goadmin:1.0
# docker exec -it goadmin /bin/bash
# # if your local code has been changed ,you can restart the container to take effect
# docker restart goadmin
#  

FROM golang:latest
MAINTAINER josingcjx
COPY . /home/goadmin
ENV GOPATH=$GOPATH:/home/goadmin/ GOPROXY=https://mirrors.aliyun.com/goproxy,https://goproxy.cn,direct
RUN apt-get update --fix-missing && \
    apt-get install -y zip vim postgresql mysql-common default-mysql-server && \
    tar -C / -xvf /home/goadmin/tools/godependacy.tgz 
    #if install dependacy tools failed, you can copy local's to remote
    #mkdir -p /go/bin  && \
    #mv /home/goadmin/tools/{gotest,goimports,golint,golangci-lint,adm} /go/bin
    #go get golang.org/x/tools/cmd/goimports && \
    #go get github.com/rakyll/gotest && \
    #go get -u golang.org/x/lint/golint && \
    #go install github.com/GoAdminGroup/adm@latest && \
    #go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
WORKDIR /home/goadmin