FROM golang:1.23-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY go.sum ./

COPY . .

RUN go build -o main .

EXPOSE 50051

CMD ["./main"]