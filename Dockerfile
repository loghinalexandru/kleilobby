#syntax=docker/dockerfile:latest

FROM golang:1.22-alpine AS build

WORKDIR /app

COPY ../ .

RUN go mod download
RUN go build -o / .

FROM golang:1.20-alpine

WORKDIR /

COPY --link --from=build /kleilobby /kleilobby

EXPOSE 3002

ENTRYPOINT ["/kleilobby"]
