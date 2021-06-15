#!/usr/bin/env bash
# 拉取最新
git pull
# project
project=$(basename `pwd`)
# 修改开发模式
./devMode.sh prod ${project}

cd docker
./pushAll.sh ${project}
cd ..
# 还原开发模式
./devMode.sh dev