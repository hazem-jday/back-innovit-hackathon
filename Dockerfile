FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /app/
WORKDIR /app/
RUN go mod download
COPY . /app/
ENTRYPOINT ["go","run","main.go"]