# 批量拉取镜像
#!/bin/bash
cat docker-compose.yaml  | grep deercoder-gin-demo- | awk '{print "sudo docker pull "$2}' | sh