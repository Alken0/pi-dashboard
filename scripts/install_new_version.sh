#!/usr/bin/env bash

# this binary gets referenced by:
# /etc/systemd/system/pidashboard.service

go build cmd/server/main.go
sudo chown root:root main
sudo chmod 755 main
