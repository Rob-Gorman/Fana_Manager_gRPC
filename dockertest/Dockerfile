FROM golang:latest

WORKDIR /go/flagserver

COPY . .

RUN go build dockertest.go

EXPOSE 3000

CMD ["go", "run", "dockertest.go"]