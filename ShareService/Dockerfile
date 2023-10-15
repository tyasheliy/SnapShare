FROM golang:1.21-bullseye AS builder

ENV GOOS=linux \
    CGO_ENABLED=0

WORKDIR /build

COPY . .

RUN go build -o app cmd/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /build/app /app/app

COPY --from=builder /build/config/config.json /app/config.json

EXPOSE 4000

CMD ["./app"]