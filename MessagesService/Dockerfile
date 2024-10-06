FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o=./bin/messages_service ./cmd/messages_service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/bin/messages_service .

EXPOSE 8080

CMD ./messages_service