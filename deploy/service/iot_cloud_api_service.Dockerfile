FROM alpine

WORKDIR /opt/bat

RUN mkdir /opt/bat/conf
RUN mkdir /opt/bat/conf/open
RUN mkdir /opt/bat/temp

#安装依赖项
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.15/main/ > /etc/apk/repositories
RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add curl
RUN apk update && apk add tzdata

#拷贝配置、程序、并赋予程序可执行权限
COPY ../bin/conf/.env /opt/bat/conf
COPY ../bin/temp /opt/bat/temp
COPY ../bin/conf/cloud_api_msg.ini /opt/bat/conf
COPY ../bin/conf/panel_translate.ini /opt/bat/conf
COPY ../bin/conf/open/iot_cloud_api_service.yml  /opt/bat/conf/open
COPY ../bin/iot_cloud_api_service /opt/bat/

#挂载目录
VOLUME ../bin/temp:/opt/bat/temp

RUN chmod ugo+x ./temp/mcu_sdk/pull_mcu_sdk_template.sh
RUN chmod ugo+x ./iot_cloud_api_service

EXPOSE 0-65535

CMD ["./iot_cloud_api_service"]
