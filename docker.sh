#!/usr/bin/env bash

# -tags netgo apline构建golang编译问题
GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -tags netgo -o main

# docker build
# default version
version=0.1
project=gt-crud
# 参数处理
# :需要参数
while getopts ":p:v:h" opt
do
    case ${opt} in
        p)
        project=$OPTARG
        echo "项目的值${project}"
        ;;
        v)
        version=$OPTARG
        echo "版本号version的值${version}"
        ;;
        h)
        echo -e "-v 版本号id\n-h 帮助\n"
        exit 1
        ;;
        ?)
        echo "未知参数"
        exit 1;;
    esac
done

# docker build
# 版本记录
docker build -f ./Dockerfile -t registry.cn-hangzhou.aliyuncs.com/dreamlu/common:${project}-${version} .
# 最新版本 :project
docker tag registry.cn-hangzhou.aliyuncs.com/dreamlu/common:${project}-${version}  registry.cn-hangzhou.aliyuncs.com/dreamlu/common:${project}

# remove build
rm -rf main