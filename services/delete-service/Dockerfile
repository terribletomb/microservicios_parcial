FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o delete_service .

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/delete_service /app/delete_service

EXPOSE 8083

CMD ["./delete_service"]
