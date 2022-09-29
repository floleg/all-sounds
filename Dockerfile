## Build
FROM golang:1.19 AS build

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY ./db ./db
COPY ./server ./server

RUN go build -v -o /usr/local/bin/server ./server

## Deploy
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=build /usr/local/bin/server /server

CMD ["/server"]