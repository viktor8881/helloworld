FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .
COPY config.yaml.dist .
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

EXPOSE 8080
ENTRYPOINT ["./entrypoint.sh"]
