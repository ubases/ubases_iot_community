FROM alpine

WORKDIR /opt/bat

RUN mkdir /opt/bat/conf
RUN mkdir /opt/bat/conf/open

#安装依赖项
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.15/main/ > /etc/apk/repositories
RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add curl
RUN apk update && apk add tzdata

#拷贝配置、程序、并赋予程序可执行权限
COPY ../bin/conf/.env /opt/bat/conf
COPY ../bin/conf/GeoLite2-City.mmdb /opt/bat/conf
COPY ../bin/conf/open/iot_weather_service.yml  /opt/bat/conf/open
COPY ../bin/iot_weather_service /opt/bat/
RUN chmod ugo+x ./iot_weather_service

EXPOSE 0-65535

CMD ["./iot_weather_service"]
