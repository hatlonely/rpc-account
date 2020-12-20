FROM centos:centos7

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" >> /etc/timezone

COPY docker/ /var/docker/go-rpc-account
RUN mkdir -p /var/docker/go-rpc-account/log

WORKDIR /var/docker/go-rpc-account
CMD [ "bin/account", "-c", "config/go-rpc-account.json" ]