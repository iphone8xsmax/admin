FROM golang:1.13.5-alpine3.10 as builder
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/release
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o myapp main.go

FROM scratch as prod
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /go/release/gin_demo /
COPY --from=builder /go/release/conf ./conf
EXPOSE 8888
CMD ["/myapp"]