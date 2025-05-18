FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git gettext
# Configure Git to use HTTPS with GitHub token from build arg
ARG GITHUB_TOKEN
RUN if [ -n "$GITHUB_TOKEN" ]; then \
    git config --global url."https://${GITHUB_TOKEN}:@github.com/".insteadOf "https://github.com/" && \
    echo "machine github.com login ${GITHUB_TOKEN}" > ~/.netrc; \
    fi
# Set GOPRIVATE for private repositories
ENV GOPRIVATE=github.com/viktor8881
RUN go mod tidy
RUN go mod vendor
RUN GOOS=linux GOARCH=amd64 go build -mod=vendor -o app .

FROM alpine:3.19
WORKDIR /app
RUN apk add --no-cache gettext
COPY --from=builder /app/app .
COPY config.yaml.dist .
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

EXPOSE 8080
ENTRYPOINT ["./entrypoint.sh"]
