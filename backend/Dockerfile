FROM golang:1.22-alpine as builder
RUN apk --no-cache add ca-certificates git build-base
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./bin/main ./cmd/server/main.go

CMD ["./bin/main", "-config=./config/config.yml"]