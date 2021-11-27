FROM golang:alpine AS builder

WORKDIR /app/

COPY go.mod .
COPY go.sum .

RUN go env -w GOPROXY='https://goproxy.cn,goproxy.io,direct'
RUN go mod download

COPY . .

RUN go build -o /cache-client

FROM alpine:latest

WORKDIR /app/

COPY --from=BUILDER /cache-client .
COPY config.yaml .

EXPOSE 8080

ENTRYPOINT ["./cache-client"]

CMD [""]