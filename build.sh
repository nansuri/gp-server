#!/bin/sh
git pull origin master
sudo docker build -t gp-server .
sudo docker run --name gp-server --restart=always -p 3025:3025 gp-server:latest