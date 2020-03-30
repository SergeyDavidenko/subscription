#!/bin/bash

HASH=`git rev-parse origin/master`
SAVER_TAG=`date +"%Y%m%d-%H%M"`-${HASH:0:7}

docker pull golang:latest
echo "Start build"
docker build -t subscription:$SAVER_TAG .
echo "Build docker container subscription:$SAVER_TAG"
# docker push subscription:$SAVER_TAG
# echo "Build push container subscription:$SAVER_TAG"