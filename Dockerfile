FROM golang:1.15.7 AS builder
WORKDIR /go/src/github.com/larkintuckerllc/exporter/cmd/exporter
ADD . /go/src/github.com/larkintuckerllc/exporter
RUN CGO_ENABLED=0 GOOS=linux go install .

FROM alpine:latest  
WORKDIR /root
COPY --from=builder /go/bin/exporter .
ENTRYPOINT /root/exporter

