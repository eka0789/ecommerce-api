FROM golang:1.24.5-alpine

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main ./cmd/main.go
RUN go build -o seeder ./cmd/seeder.go

EXPOSE 8080
CMD ["./main"]
