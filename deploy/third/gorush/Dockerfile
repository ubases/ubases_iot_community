FROM alpine

WORKDIR /opt/gorush

#拷贝配置、程序、并赋予程序可执行权限
COPY credentials /opt/gorush/credentials
COPY config.yml /opt/gorush/
COPY push_dev.p12 /opt/gorush/
COPY gorush /opt/gorush/

RUN chmod ugo+x ./gorush

EXPOSE 0-65535

CMD ["./gorush","-c","./config.yml"]

