# 批量拉取镜像
#!/bin/bash
cat docker-compose.yaml  | grep :gt-crud | awk '{print "sudo docker pull "$2}' | sh

# 删除空镜像
docker images|grep none|awk '{print $3 }'|xargs docker rmi