FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
ENV GOARCH=amd64
RUN go build -o app .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
ENTRYPOINT ["./app"]
