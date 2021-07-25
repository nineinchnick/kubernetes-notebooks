#!/bin/bash

TAG=nineinchnick/config-printer:0.1

docker build --tag $TAG .

docker push $TAG
kind load docker-image $TAG


TAG=nineinchnick/config-printer:0.2

docker build --tag $TAG -f Dockerfile2 .

docker push $TAG
kind load docker-image $TAG
