# 拉镜像
```
docker pull --platform linux/arm64 ubuntu:22.04
```
# 运行容器
```
docker run -itd -p 0.0.0.0:10000:22 --name myubuntu ubuntu:22.04 /bin/bash -v ~/.ssh:/root/.ssh
```
# 进入容器
```
docker exec -it myubuntu /bin/bash
```
# 安装软件
```
apt-get update
apt-get install openssh-server
apt-get install vim, wget, curl, git,net-tools, iputils-ping, iproute2
```
# 配置ssh
```
vim /etc/ssh/sshd_config
```
```
PermitRootLogin yes
PasswordAuthentication yes
```
```
service ssh restart
```
# 配置root密码
```
passwd root
```
# 配置ssh免密登录
```
ssh-keygen -t rsa
```
```
ssh-copy-id -i ~/.ssh/id_rsa.pub root@
```

# 备份这个容器
```
docker commit myubuntu myubuntu:22.04
```
# 保存这个容器
```
docker save -o myubuntu.tar myubuntu:22.04
```

# 从镜像导入容器
```
docker run -itd -p 0.0.0.0:10000:22 --name ubuntu myubuntu:22.04 /usr/sbin/sshd -D
```