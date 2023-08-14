#!/bin/sh

# this script builds the image locally then
# pushes it to the docker hosts local registry

docker build -t utils:latest . &&
docker save utils:latest | gzip > utils.tar.gz &&
scp utils.tar.gz zwojcik@192.168.128.19:~/utils.tar.gz &&
ssh zwojcik@192.168.128.19 'docker load -i ~/utils.tar.gz'