FROM golang:1.24-alpine AS builder
LABEL authors="maria-radeva"

ENTRYPOINT ["top", "-b"]

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o schrodinger-db .
FROM alpine:latest
RUN apk --no-cache add postgresql-client
WORKDIR /root/
COPY --from=builder /app/schrodinger-db .
COPY sample.env .env

# future API port exposure
# EXPOSE 8080

ENTRYPOINT ["./schrodinger-db"]

