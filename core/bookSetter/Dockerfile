FROM golang:alpine

WORKDIR /app
COPY ./go.* /app/
RUN go mod download
RUN go mod tidy

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main


CMD ["/app/main"]