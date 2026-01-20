FROM golang:1.24-alpine AS builder

WORKDIR /app

# 1. Copiar y descargar dependencias (Capa cacheable)
COPY go.mod go.sum ./
RUN go mod download

# 2. Copiar SOLO el código compartido (pkg)
COPY pkg/ ./pkg/

# 3. Copiar SOLO el código de ESTE servicio específico
COPY services/pick-age/ ./services/pick-age/

# 4. Compilar
RUN CGO_ENABLED=0 GOOS=linux go build -o /server services/pick-age/cmd/main.go

# Etapa Final (Scratch)
FROM scratch
WORKDIR /
COPY --from=builder /server /server
ENTRYPOINT ["/server"]

