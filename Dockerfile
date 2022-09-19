FROM golang:1.17-alpine

RUN apk add --no-cache git
WORKDIR /app/portfolio-blog-api

COPY go.mod .
COPY go.sum .
COPY portfolio-blog-api-key.json .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/portfolio-blog-api .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/portfolio-blog-api"]