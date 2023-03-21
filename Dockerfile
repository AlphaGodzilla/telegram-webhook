FROM golang:1.19.4-alpine as builder
WORKDIR /telegram-webhook
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine:3.17.2
COPY --from=builder /telegram-webhook/telegram-webhook /
RUN apk --no-cache add ca-certificates && update-ca-certificates
CMD ["/telegram-webhook"]
