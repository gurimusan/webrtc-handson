#!/bin/sh

docker build -t webrtc-handson-golang .
docker run --rm -p 3001:3001 webrtc-handson-golang
