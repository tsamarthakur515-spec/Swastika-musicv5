FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /swastika-music .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates ffmpeg python3 py3-pip \
    && pip3 install --break-system-packages --no-cache-dir yt-dlp
WORKDIR /app
COPY --from=builder /swastika-music .
CMD ["./swastika-music"]
