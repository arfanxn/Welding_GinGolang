# ---- build stage ----
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git build-base
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o /server ./main.go

# ---- final stage ----
FROM alpine:3.19

# Install dependencies and create app user with a home directory
RUN apk add --no-cache ca-certificates tzdata && \
    adduser -D -h /app app

# Set working directory and copy the binary
WORKDIR /app
COPY --from=builder /server .
COPY --from=builder /src/database/migrations  ./database/migrations

# Create logs directory and set permissions
RUN mkdir -p /app/logs && \
    chown -R app:app /app

# Switch to non-root user
USER app

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./server", "serve"]
