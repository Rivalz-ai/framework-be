#build stage
FROM golang:1.23-alpine AS builder
ENV CGO_ENABLED=1
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init
WORKDIR /go/src/app
COPY . .
RUN go build -o /go/bin/main -v ./main.go

#final stage
FROM alpine:3.21.2
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/main /main
ENTRYPOINT /main
EXPOSE 30000
