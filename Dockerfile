# ETAPA 1: Construcción (El taller)
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main api/main.go

# ETAPA 2: Ejecución (El cohete ligero)
FROM scratch
WORKDIR /root/
# Solo copiamos el binario y los archivos visuales (Máxima seguridad y mínimo peso)
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
EXPOSE 8080
CMD ["./main"]