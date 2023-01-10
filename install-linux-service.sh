sudo CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

sudo cp telegram-webhook /usr/local/bin/

sudo cp ./telegram-webhook.service /usr/lib/systemd/system/

sudo systemctl start telegram-webhook.service

sudo systemctl enable telegram-webhook.service
