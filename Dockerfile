FROM golang:1.16-alpine AS builder
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    git \
    && update-ca-certificates

COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -o nifheim -a .

FROM alpine:latest as prod

COPY --from=builder /app/nifheim /usr/local/bin/nifheim
EXPOSE 5000

ENTRYPOINT ["/usr/local/bin/nifheim"]