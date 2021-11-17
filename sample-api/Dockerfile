FROM golang:1.17.3 as builder

WORKDIR /sample-api
COPY . /sample-api

RUN go mod download
WORKDIR /sample-api/app
RUN go build -o /bin/sample-api -ldflags '-s -w'

FROM debian:stable-slim as runner

COPY --from=builder /bin/sample-api /app/sample-api

CMD ["/app/sample-api"]
