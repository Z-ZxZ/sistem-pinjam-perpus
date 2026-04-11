# Build stage
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod ./
# go.sum will be created after dependencies are installed
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server/main.go

# Run stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
