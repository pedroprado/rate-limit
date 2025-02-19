# Build phase
FROM golang:1.23-alpine AS builder

LABEL maintainer="Bexs Digital"

ENV CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app

COPY . .

RUN apk update

RUN apk add git


RUN go mod download

RUN go mod tidy

WORKDIR /app/src

RUN go build -o notification-service

# Copy phase
FROM alpine

WORKDIR /app

COPY --from=builder /app/src/notification-service .

EXPOSE 8182

CMD ["./notification-service"]