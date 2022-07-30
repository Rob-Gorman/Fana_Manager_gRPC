FROM node as builder
WORKDIR /reactapp
COPY manager-dash .
RUN npm i
RUN npm run build

FROM golang:latest as manager

WORKDIR /go/src/gomanager
COPY . .
COPY --from=builder /reactapp/build ./build
RUN go mod tidy
EXPOSE 3000
CMD ["go", "run", "main.go"]