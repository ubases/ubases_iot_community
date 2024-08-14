FROM openjdk:8-jdk

WORKDIR /opt/bat

RUN mkdir /opt/bat/conf
RUN mkdir /opt/bat/conf/open
RUN mkdir /opt/bat/temp

#安装依赖项
#RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.15/main/ > /etc/apk/repositories
#RUN sed -i 's/https/http/' /etc/apk/repositories
#RUN apk add curl
#RUN apk update && apk add tzdata
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && apt-get install -y ntpdate \
    && rm -rf /var/lib/apt/lists/* \
    && ntpdate -q pool.ntp.org

#拷贝配置、程序、并赋予程序可执行权限
COPY ../bin/conf/.env /opt/bat/conf
COPY ../bin/conf/global.yml /opt/bat/conf
COPY ../bin/conf/open/iot_app_oem_service.yml  /opt/bat/conf/open
COPY ../bin/pepk.jar /opt/bat/
COPY ../bin/iot_app_oem_service /opt/bat/

#挂载目录
VOLUME ../bin/temp:/opt/bat/temp

RUN chmod ugo+x ./iot_app_oem_service


EXPOSE 0-65535

CMD ["./iot_app_oem_service"]
