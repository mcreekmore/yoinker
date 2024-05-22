FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod ./
# COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o yoinker ./cmd/main.go

# FROM alpine:latest
FROM python:3.12.3-alpine3.19

RUN apk add --no-cache wget ca-certificates tar xz

RUN wget -q https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -O /usr/local/bin/yt-dlp \
    && chmod a+rx /usr/local/bin/yt-dlp

RUN wget -q https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz -O /tmp/ffmpeg.tar.xz \
    && mkdir -p /tmp/ffmpeg \
    && tar -xf /tmp/ffmpeg.tar.xz -C /tmp/ffmpeg --strip-components=1 \
    && mv /tmp/ffmpeg/ffmpeg /usr/local/bin/ \
    && chmod a+rx /usr/local/bin/ffmpeg \
    && rm -rf /tmp/ffmpeg* /tmp/ffmpeg.tar.xz

COPY --from=builder /app/. /app/.

EXPOSE 8080

WORKDIR /app/cmd

CMD ["/app/yoinker"]
