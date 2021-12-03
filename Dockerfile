# syntax=docker/dockerfile:1

FROM golang:1.16-alpine3.14

WORKDIR /api-app

COPY . .

RUN go mod download

RUN go build -o mainfile

EXPOSE 80

CMD ["./mainfile"]
