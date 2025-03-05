FROM golang:1.24.0-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o dbs2

EXPOSE 8081

CMD ["./dbs2"]