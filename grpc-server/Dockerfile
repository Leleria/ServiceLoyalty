FROM golang:1.22-alpine AS builder

RUN apk add --no-cache --virtual build-deps gcc musl-dev linux-headers

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .


RUN CGO_ENABLED=1 go build -o loyalty-service ./Cmd/LoyaltyService/main.go

RUN apk del build-deps

FROM alpine:latest

RUN apk add --no-cache sqlite

WORKDIR /app

COPY --from=builder /app/loyalty-service .
COPY Config/config_local.yaml /app/grpc-server/Config/
COPY Migrations/*.sql ./Migrations
COPY Migrations/*.sql ./Tests/Migrations
COPY Storage/ServiceLoyaltyDB.db /app/Storage/ServiceLoyaltyDB.db




EXPOSE 44044

CMD ["/app/loyalty-service"]
