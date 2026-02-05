FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o service main.go

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/service /
ENV TESTKUBE_STATICDIR=/static
CMD ["./service"]
