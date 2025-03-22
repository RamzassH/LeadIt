FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY ../go.work .

COPY ../libs ./libs
COPY . ./


RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
     go build -o /app/bin/gateway ./gateway/cmd/main.go

FROM scratch

COPY --from=builder /app/bin/gateway .

WORKDIR /app
EXPOSE 8080

ENTRYPOINT ["./gateway", "-b"]
