# ETAPA 1: Constructor (Builder)
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# Compilamos el binario (go:embed incluirá api/static automáticamente)
RUN CGO_ENABLED=0 GOOS=linux go build -o main api/main.go

# ETAPA 2: Ejecución (Runtime)
FROM alpine:latest
WORKDIR /root/

# SOLO necesitamos el binario. 
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]