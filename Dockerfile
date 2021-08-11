FROM golang:1.16.6-alpine3.14 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build .

CMD /app/mail-notification

FROM alpine:3.14.0

COPY --from=builder /app/mail-notification /usr/local/bin/mail-notification

CMD mail-notification