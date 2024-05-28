FROM alpine:3.18

# 这个关键词的意思是复制的意思，可以将宿主机中的内容复制到容器中
# 命令 左边是宿主机的目录，右边是容器目录
RUN mkdir /user && mkdir /user/bin && mkdir /user/conf

# 复制编译后的二进制文件
COPY bin/user-api /user/bin/
# 复制配置文件
COPY api/etc/user.yaml /user/conf/

# 为二进制提供执行权限
RUN chmod +x /user/bin/user-api

# 该命令指定容器会默认进入那个目录，如我们每次进入服务器的时候会自动进入root目录一样的作用
WORKDIR /user

# 这个命令可以让我们的docker容器在启动的时候就执行下面的命令
# 与CMD不同之处是，在docker run 后跟的命令不能替换它，它仍然会启动的时候执行
ENTRYPOINT ["bin/user-api", "-f","/user/conf/user.yaml"]