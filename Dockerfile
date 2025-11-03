FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o /go-crypto-service ./cmd
FROM alpine:3.18
COPY --from=builder /go-crypto-service /go-crypto-service
EXPOSE 8080 50051
CMD ["/go-crypto-service"]