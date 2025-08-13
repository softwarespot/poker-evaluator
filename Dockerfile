# See URL: https://hub.docker.com/_/golang
# Use the Go image to build the binary only
FROM golang:1.25.0 AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
WORKDIR /go/src/poker-evaluator/
COPY . .

RUN make

# See URL: https://hub.docker.com/_/alpine
# Use this image (~50MB) to run the "poker-evaluator", as the Go image contains too much bloat,
# which isn't needed for running the application in production and the image which can be uploaded
# to a public/private Docker register is then small
FROM alpine:3.20.3

COPY --from=builder /go/src/poker-evaluator/bin/* ./
ENTRYPOINT ["./poker-evaluator"]
