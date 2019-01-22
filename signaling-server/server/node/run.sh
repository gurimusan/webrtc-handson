#!/bin/sh

docker build -t webrtc-handson-node .
docker run --rm -p 3001:3001 webrtc-handson-node &
