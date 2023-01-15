FROM golang:alpine

WORKDIR /app
COPY ./go.* /app/
RUN go get -t -v ./...

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main

CMD ["/app/main"]