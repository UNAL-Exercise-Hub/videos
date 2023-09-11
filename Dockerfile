# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build
RUN go build -o /golang_videos_ms

EXPOSE 8080

CMD ["/golang_videos_ms"]