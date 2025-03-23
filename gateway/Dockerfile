FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY gateway/go.mod gateway/go.sum ./gateway/
COPY libs ./libs

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
     go build -o /app/bin/gateway ./gateway/cmd/main.go

FROM alpine:lates AS runner

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/bin/gateway .

WORKDIR /app
EXPOSE 8080

ENTRYPOINT ["./gateway", "-b"]
