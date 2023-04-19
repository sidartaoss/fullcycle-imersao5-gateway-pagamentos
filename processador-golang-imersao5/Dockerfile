FROM golang:latest

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

WORKDIR /go/src

COPY . .

WORKDIR /go/src/cmd

RUN ls -la && GOOS=linux go build -o main .

CMD ["./main"]