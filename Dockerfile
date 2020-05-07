FROM golang:1.13.3-alpine

RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go env -w GO111MODULE=on

WORKDIR $GOPATH/src
COPY . .
RUN go build -o $GOPATH/bin/main main.go

EXPOSE 8080

WORKDIR $GOPATH
RUN cp /opt/src/index.html /opt/
CMD bin/main
