FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY ./gateway ./gateway/
COPY ./libs/contracts ./libs/contracts/

WORKDIR /app/gateway

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
     go build -o /gateway ./cmd/main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /gateway .

WORKDIR /app
EXPOSE 8080

ENTRYPOINT ["./gateway", "-b"]
