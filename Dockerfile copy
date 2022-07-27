FROM golang:latest

WORKDIR /go/src/gomanager

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "main.go"]