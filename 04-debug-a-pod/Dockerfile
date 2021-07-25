FROM golang:1.16-buster AS build

WORKDIR /go/src/app
COPY ./*.go .
COPY ./*.mod .

RUN go build -v ./...

FROM debian:buster

COPY --from=build /go/src/app/config-printer /config-printer

CMD ["/config-printer"]
