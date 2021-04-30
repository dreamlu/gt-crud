# 懒人必备: 拿到服务器ip+密码的软件安装、目录创建、文件上传一体化脚本
# 配置免密登录
#1、在主机A生成秘钥对
#ssh-keygen -t rsa
#一路回车，最后会生成秘钥对：
#Your identification has been saved in /home/lu/.ssh/id_rsa.
#Your public key has been saved in /home/lu/.ssh/id_rsa.pub.
ip=
name=gt-crud
#2、将公钥复制到主机B, 已存在则可注释，已存在会有警告, 无影响
# =============================================
ssh-copy-id -p 22 -i ~/.ssh/id_rsa.pub root@${ip}
# =============================================
# apt install
ssh root@${ip} "[ -x "$(command -v docker-compose)" ] || apt update & apt install docker docker-compose"
# mkdir
ssh root@${ip} "[ -d /root/${name} ] && echo ok || mkdir -p /root/${name}/docker"
cd ..
# 更新：不带/root/${name}/docker/, 后的docker/可直接覆盖
# but: 初次复制时服务器后缀需要加上docker下级目录(/root/${name}/docker/),必须确保copy的上级目录存在
# 最新：由上面的mkdir直接创建,无需关心
scp -r docker/ root@${ip}:/root/${name}/