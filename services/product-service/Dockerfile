FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o /bin/product-service ./cmd/product-service

FROM alpine:3.19
COPY --from=builder /bin/product-service /bin/product-service
EXPOSE 9082
ENTRYPOINT ["/bin/product-service"]

