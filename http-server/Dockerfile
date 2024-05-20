FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o http-server main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/http-server .

EXPOSE 50051

CMD ["/app/http-server"]
