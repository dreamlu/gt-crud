#!/bin/bash
./pullAll.sh
docker-compose up --build -d
# 删除空镜像
docker images|grep none|awk '{print $3 }'|xargs docker rmi