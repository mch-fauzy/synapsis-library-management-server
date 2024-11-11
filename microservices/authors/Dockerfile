# Stage 1: Build stage
FROM golang:1.22-alpine as builder

RUN go install github.com/rubenv/sql-migrate/...@latest

# Install dockerize to wait docker container
RUN wget -qO- https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz | tar xvz -C /usr/local/bin

WORKDIR /app

# Install Go dependencies (this will be cached unless go.mod or go.sum changes)
COPY go.mod go.sum ./ 
RUN go mod tidy

COPY . .

RUN go build -o main .

# Stage 2: Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy executables files from the builder stage to specific directory
COPY --from=builder /usr/local/bin/dockerize /usr/local/bin/
COPY --from=builder /go/bin/sql-migrate /usr/local/bin/

# Copy files from the builder stage to current working directory
COPY --from=builder /app/main .
COPY --from=builder /app/start.sh .
COPY --from=builder /app/.env .
COPY --from=builder /app/.env.docker .
COPY --from=builder /app/dbconfig.yml .

# Copy files inside migrations (builder stage) to migrations folder in current working directory
COPY --from=builder /app/migrations ./migrations

RUN chmod +x start.sh

ENTRYPOINT ["./start.sh"]
