FROM golang:1.23.5-alpine3.21 AS go-builder

RUN apk update && apk add git

WORKDIR /app

COPY ./backend ./

RUN go mod download \
&& go build -o main /app/cmd/main.go

FROM alpine:3.21
WORKDIR /app
COPY --from=go-builder /app/main .

EXPOSE 4000

CMD ["/app/main"]