#!/usr/bin/env bash

# change version from 1.25.2 to any version
# downloads and replaces the user-go-version
wget https://go.dev/dl/go1.25.2.linux-armv6l.tar.gz 
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.25.2.linux-armv6l.tar.gz
rm go1.25.2.linux-armv6l.tar.gz

# only run the first time
# echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
# source ~/.profile
