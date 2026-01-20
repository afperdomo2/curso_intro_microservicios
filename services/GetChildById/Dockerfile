# Multistage build for small image size
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY pkg ./pkg
COPY services/GetChildById ./services/GetChildById

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /server services/GetChildById/main.go

# Final stage
FROM scratch
COPY --from=builder /server /server
EXPOSE 8080
CMD ["/server"]
