# Building the binary of the App
FROM golang:latest AS builder
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/video-api /app/cmd/video-api/main.go

# Moving the binary to the 'final Image' to make it smaller
FROM alpine
WORKDIR /app
RUN apk add --no-cache nano git curl
# COPY --from=builder /feedback-api/internal/configs/dev.env .env
COPY --from=builder bin/video-api video-api

CMD ["./video-api"]

EXPOSE 80
