cd ..
# 更新：不带/root/gt-crud/docker/, 后的docker/可直接覆盖
# but: 初次复制时服务器后缀需要加上docker下级目录(/root/gt-crud/docker/),必须确保copy的上级目录存在
scp -r docker/ root@ip:/root/gt-crud/