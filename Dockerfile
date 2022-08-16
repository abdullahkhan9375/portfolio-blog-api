# syntax=docker/dockerfile:1
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go mod tidy

RUN go build -o github.com/abdullahkhan9375/portfolio-blog-api

EXPOSE 8080

CMD [ "/github.com/abdullahkhan9375/portfolio-blog-api" ]