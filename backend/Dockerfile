# Build stage
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod ./
# go.sum will be created after dependencies are installed
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/main.go
RUN go build -o seeder ./cmd/seeder/main.go

# Run stage
FROM alpine:3.18

WORKDIR /app

# Copy binaries from builder
COPY --from=builder /app/server .
COPY --from=builder /app/seeder .

EXPOSE 8080

CMD ["./server"]
