FROM golang:latest AS build-base
WORKDIR /go/src/github.com/sasimpson/testkube
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download


FROM build-base AS builder
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GARCH=arm go build -o service service.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/sasimpson/testkube/service .
RUN chmod +x service
CMD ["./service"]
