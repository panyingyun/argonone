#!/bin/bash
go build -o argonone

sudo cp argonone /usr/local/bin/
sudo mkdir -p  /etc/argonone/
sudo cp prod.yml /etc/argonone/
sudo cp argonone.service /lib/systemd/system/
sudo chmod 644 /lib/systemd/system/argonone.service

sudo systemctl daemon-reload
sudo systemctl enable argonone.service
sudo systemctl restart  argonone.service