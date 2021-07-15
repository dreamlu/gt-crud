#!/usr/bin/env bash
# 拉取最新,多人合作需要
#git pull
# project
work=`pwd`
project=$(basename `pwd`)
# 终止信号捕捉,还原开发模式
trap 'cd ${work};echo 目录:`pwd`; ${work}/devMode.sh dev;exit' 2
# 修改开发模式
./devMode.sh prod ${project}

cd docker
./pushAll.sh ${project}
cd ..
# 还原开发模式
./devMode.sh dev