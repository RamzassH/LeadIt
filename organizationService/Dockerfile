FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY ./organizationService ./organizationService/
COPY ./libs/ ./libs/

WORKDIR /app/organizationService

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /organization-service ./cmd/organization/main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates

WORKDIR /app

RUN mkdir -p ./organizationService/config

COPY --from=builder /organization-service .
COPY --from=builder /app/organizationService/config/local.yaml /app/organizationService/config/config.yaml

EXPOSE 57444

CMD ["/app/organization-service", "--config", "/app/organizationService/config/config.yaml"]
