FROM golang:latest

WORKDIR /go/src/gomanager
COPY . .
RUN go mod tidy
EXPOSE 3000
CMD ["go", "run", "main.go"]