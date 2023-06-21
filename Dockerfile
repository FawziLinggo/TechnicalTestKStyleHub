FROM golang:1.18
WORKDIR /app

COPY . .
RUN go build -o server/http/main server/http/main.go

EXPOSE 3000
WORKDIR /app/server/http

CMD ["./main"]