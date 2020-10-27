# 批量推向共有或私有仓库
#!/bin/bash
docker images | grep registry.cn-hangzhou.aliyuncs.com/dreamlu/common | grep gt-crud | awk '{print "docker push "$1":"$2}' | sh

# 删除空镜像
docker images|grep none|awk '{print $3 }'|xargs docker rmi

# 删除停止的容器
#docker rm `docker ps -a|grep Exited|awk '{print $1}'`
# 自动更新，取消注释，配置ssh免密
#ssh root@ip "cd shop/docker-compose/;./update.sh;exit"