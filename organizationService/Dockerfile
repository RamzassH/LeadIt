FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /workspace

COPY ../go.work .
COPY . ./
COPY ../libs/contracts/ ./libs/contracts/
COPY ../libs/redis/ ./libs/redis/
COPY ../libs/kafka/ ./libs/kafka/

RUN go work use ./organizationService && \
    go work use ./libs/contracts && \
    go work use ./libs/redis && \
    go work use ./libs/kafka && \
    go mod download


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /organization-service ./organizationService/cmd/organization/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /organization-service /app/
COPY organizationService/config/local.yaml /app/config.yaml

WORKDIR /app
EXPOSE 57444

CMD ["/app/organization-service", "--config", "/app/config.yaml"]