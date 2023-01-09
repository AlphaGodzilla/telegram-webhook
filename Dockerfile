FROM golang:1.19.4-alpine as builder
WORKDIR /telegram-webhook
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch
COPY --from=builder /telegram-webhook/telegram-webhook /
CMD ["/telegram-webhook"]
