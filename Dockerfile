FROM golang:1.13.5-alpine3.10 as builder
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/release
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o zhxf main.go

EXPOSE 8888
CMD ["/zhxf"]