FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.work ./

COPY libs ./libs
COPY notificationService ./notificationService
COPY authService ./authService
COPY gateway ./gateway

WORKDIR /app/gateway
RUN go mod download

WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/gateway ./gateway/cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/gateway .

EXPOSE 8080

ENTRYPOINT ["./gateway", "-b"]
