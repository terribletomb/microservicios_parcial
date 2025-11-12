# Etapa 1: compilación
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o create_service .

# Etapa 2: imagen final
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/create_service /app/create_service

# El contenedor siempre expone el puerto interno donde escucha la app Go
EXPOSE 8080

# Arrancar la app
CMD ["./create_service"]
