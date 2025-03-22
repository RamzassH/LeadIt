FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.work .
COPY notificationService ./notificationService
COPY authService ./authService
COPY organizationService ./organizationService
COPY libs/contracts ./libs/contracts
COPY libs/redis ./libs/redis
COPY libs/kafka ./libs/kafka
COPY gateway ./gateway

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth-service ./organizationService/cmd/organization/main.go

RUN apk update && apk add --no-cache git

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
RUN mkdir -p ./organizationService/config

COPY --from=builder /app/organization-service .
COPY --from=builder /app/organizationService/config/local.yaml /app/organizationService/config/config.yaml

EXPOSE 57444

CMD ["./organization-service", "--config", "/app/organizationService/config/config.yaml"]