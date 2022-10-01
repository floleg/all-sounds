## Build
FROM golang:1.19 AS build

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

ADD ./configs ./configs

ADD ./cmd ./cmd
ADD ./pkg ./pkg
ADD ./internal ./internal

RUN go build -v -o /usr/local/bin/server ./cmd/server

## Deploy
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=build /usr/local/bin/server /server
COPY --from=build /usr/src/app/configs /configs

EXPOSE 8080

CMD ["/server"]