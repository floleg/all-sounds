## Build
FROM golang:1.19-alpine3.16 AS build
WORKDIR /cmd

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./db ./db
COPY ./server ./server

WORKDIR /cmd/server

RUN go build -o ./server

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /cmd/server /server

ENTRYPOINT ["/server"]