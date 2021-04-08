#!/usr/bin/env bash
# 参数
# 备份文件目录
bakDir=
# 数据信息
USER=
PASSWORD=
DATABASE=
# 容器名或id, 不传备份本机
CONTAINER=

# 运行docker-compose, 启动程序
# docker-compose up -d
# 安装备份脚本
#./mysql/setup.sh -c 容器名或id -u 数据库账号 -p 数据库密码 -d 备份的数据库名 -b 备份目录(/root/bak)
cd mysql
./setup.sh -c ${CONTAINER} -u ${USER} -p ${PASSWORD} -d ${DATABASE} -b ${bakDir}
cd ..

# 验证定时备份脚本
cat /etc/cron.d/crontab
