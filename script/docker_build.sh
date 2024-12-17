#!/usr/bin/env bash
ImageTag=$1
if [ ! -n "$ImageTag" ]; then
  echo "image tag not set, usage: ./build_docker.sh 1.01"
  exit
fi

buildctl build --frontend=dockerfile.v0 --local context=. --local dockerfile=. --output type=image,name=registry.cn-hangzhou.aliyuncs.com/wxg-uat/micro-user-service:"$ImageTag",push=true

