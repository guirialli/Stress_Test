# Dockerfile
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

# Copiar arquivos para o container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar o binário
RUN go build -o stress_test ./cmd/stress_test/main.go

# Imagem final mais leve
FROM alpine:latest
WORKDIR /app

# Copiar o binário do estágio de build
COPY --from=builder /app/stress_test /app/stress_test

# Definir o ENTRYPOINT para aceitar argumentos dinâmicos
ENTRYPOINT ["/app/stress_test"]
