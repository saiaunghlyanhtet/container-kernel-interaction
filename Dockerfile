FROM golang:1.22 as builder

WORKDIR /app

COPY main.go .

RUN go mod init syscall-demo || true

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o time-app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/time-app .

RUN chmod +x ./time-app

EXPOSE 8080

CMD ["./time-app"]

