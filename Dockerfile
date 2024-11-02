# STAGE 1
FROM golang:alpine AS builder

ENV GO111MODULE=on
WORKDIR /go/codespacex
COPY go.mod ./

RUN go mod download
RUN go clean --modcache
RUN apk add --no-cache make

COPY . .
RUN go build -o main ./main.go

# STAGE 2
FROM alpine:3.19.0

RUN apk add --no-cache curl

WORKDIR /root/

COPY --from=builder /go/codespacex/main .
COPY --from=builder /go/codespacex/.env .

EXPOSE 7777
CMD ["nohup", "./main"]