FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o update_service .

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/update_service /app/update_service

EXPOSE 8082

CMD ["./update_service"]
