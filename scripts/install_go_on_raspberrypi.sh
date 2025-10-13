#!/usr/bin/env bash

# change version from 1.25.2 to any version
wget https://go.dev/dl/go1.25.2.linux-armv6l.tar.gz 
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.25.2.linux-armv6l.tar.gz

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
source ~/.profile
