#!/usr/bin/env bash

# 开发模式 dev/prod
# 此处修改模式
# 执行该脚本
devMode=${1}
version=1.0
project=${2}

# 必须要echo才能赋值version ???
command -v git >/dev/null 2>&1 || echo "使用git版本号"; version=`git describe --tags --abbrev=0` #version=`git rev-parse --short HEAD`
#echo "版本$version"

# 默认dev
if [[ ${devMode} = '' ]]; then
   devMode=dev
fi

# conf配置修改
# 替换源文件第三行内容
# 行首添加两个空格
sed -i '3c \  \devMode: '${devMode} conf/app.yaml

# prod 模式自动构建docker
# 可注释, 通过docker.sh执行构建
# shell awk参考:https://www.cnblogs.com/mfryf/p/3564779.html
if [[ ${devMode} = 'prod' ]]; then
    echo "prod 开始构建docker"
    ./docker.sh -p ${project} -v ${version}
fi