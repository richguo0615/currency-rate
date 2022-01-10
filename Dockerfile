# build stage
FROM golang:1.16-alpine AS builder

ARG VERSION
ENV GOOS=linux
ENV GOARCH=amd64
ARG GOPROXY
ENV GOPROXY=${GOPROXY}

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./app ./main.go

# final stage
FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /build/app .

ENTRYPOINT ["./app"]
