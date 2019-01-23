#!/bin/sh

docker build -t webrtc-handson-python .
docker run --rm -p 3001:3001 webrtc-handson-python
