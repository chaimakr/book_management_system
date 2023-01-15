FROM golang:alpine

WORKDIR /app
COPY ./go.* /app/
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main


CMD ["/app/main"]