FROM golang:1.21-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o ServiceLoyalty ./Cmd/LoyaltyService/main.go

CMD ["./ServiceLoyalty"]












#FROM golang:1.21.6-alpine AS builder
#
#WORKDIR /app
#
#COPY ./ ./
#RUN go mod download
#
#COPY *.go ./
#COPY Migrations ./Migrations
#
#RUN go build -o ServiceLoyalty/Cmd/LoyaltyService/main.go
#
#FROM alpine:latest
#
#RUN apk add --no-cache sqlite
#
#WORKDIR /app
#
#COPY --from=builder app/Cmd/LoyaltyService/main .
#COPY Migrations ./Migrations
#
#RUN mkdir -p /data
#
#CMD ["/ServiceLoyalty/Cmd/LoyaltyService/main"]