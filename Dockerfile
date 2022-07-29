FROM golang:latest

WORKDIR /go/src/gomanager
COPY . .
RUN go mod tidy

FROM node

WORKDIR /go/src/gomanager/manager-dash
RUN npm i
RUN npm run build
RUN cp -r build ../build && cd ..
EXPOSE 3000
CMD ["go", "run", "main.go"]