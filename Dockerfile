FROM golang:1.25.6-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/web 

FROM alpine:latest AS worker
WORKDIR /app
RUN adduser -D appuser
COPY --from=builder /app/app .
COPY --from=builder /app/tls ./tls/
RUN chown -R appuser:appuser /app
USER appuser
EXPOSE 4000
ENTRYPOINT ["./app"]
