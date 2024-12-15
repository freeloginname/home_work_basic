FROM golang:1.23.1 as builder

WORKDIR /build/hw16_docker/

COPY .dockerignore ./

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server_api ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /build/hw16_docker/server_api ./

ENV CONFIG_ENV=".env"

EXPOSE 8080
EXPOSE 5432

CMD ["./server_api"]
