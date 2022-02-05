FROM golang:1.17 as builder

RUN GOCACHE=OFF

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN make build-linux

FROM alpine:latest

RUN mkdir -p /app

WORKDIR /app

COPY --from=builder /app/bin/api-gateway ./api-gateway

RUN chmod +x ./api-gateway

CMD ["./api-gateway"]