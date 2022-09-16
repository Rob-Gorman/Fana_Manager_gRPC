# Build static dash files
FROM node as builder
WORKDIR /reactapp
COPY manager-dash .
RUN npm i
RUN npm run build

# Build manager binary
FROM golang:latest as managerbuild

WORKDIR /go/src/gomanager
COPY . .
COPY --from=builder /reactapp/build ./build
RUN go mod download
RUN CGO_ENABLED=0 go build -o /managerexe

# Deployment Image
FROM scratch as manager

COPY --from=managerbuild /managerexe /managerexe
EXPOSE 3000
ENTRYPOINT [ "/managerexe" ]