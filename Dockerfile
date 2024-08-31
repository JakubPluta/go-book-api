FROM golang:alpine

WORKDIR /app
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/app/bin/api"]
EXPOSE 8080