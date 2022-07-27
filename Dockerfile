FROM golang:latest

WORKDIR /go/src/gomanager
COPY ./gomanager .
RUN go mod tidy
# EXPOSE 6000
RUN go run main.go
# CMD ["go", "run", "main.go"]

FROM node:16-alpine
WORKDIR /app/dashboard
COPY ./manager-dash .
RUN npm install
CMD npm start
EXPOSE 3000