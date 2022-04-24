# Build stage
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.15
RUN apk add --no-cache bash
WORKDIR /app
COPY ./db/migrations ./db/migrations
COPY ./wait-for-it.sh .
COPY --from=builder /app/main .

EXPOSE 3000
CMD /app/main
