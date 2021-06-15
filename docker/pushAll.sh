# 批量推向共有或私有仓库
#!/bin/bash
# 私有/共有镜像地址
echo "开始推送"
dockerUrl=registry.cn-hangzhou.aliyuncs.com/dreamlu
project=${1}
# 批量推向私有仓库
cmd=`docker images | grep ${dockerUrl} | grep ${project} | awk '{print "docker push "$1":"$2}'`
echo -e "推送命令如下:\n${cmd}"
docker images | grep ${dockerUrl} | grep ${project} | awk '{print "docker push "$1":"$2}' | sh
# 删除空镜像
docker images|grep none|awk '{print $3 }'|xargs docker rmi
# 删除停止的容器
#docker rm `docker ps -a|grep Exited|awk '{print $1}'`
# 自动更新，取消注释，配置ssh免密
#ssh root@ip "cd /root/txgc/docker/;./update.sh;exit"