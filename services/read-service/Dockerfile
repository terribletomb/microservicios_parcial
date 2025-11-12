FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o read_service .

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/read_service /app/read_service

EXPOSE 8081

CMD ["./read_service"]
