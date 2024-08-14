FROM centos:7

WORKDIR /opt/bat

#拷贝配置、程序、并赋予程序可执行权限
COPY ./conf/.env /opt/bat/conf/
COPY ./conf/ca.crt /opt/bat/conf/
COPY ./conf/ca.pem /opt/bat/conf/
COPY ./conf/device.blacklist /opt/bat/conf/
COPY ./conf/device.specified /opt/bat/conf/
COPY ./conf/product.blacklist /opt/bat/conf/
COPY ./conf/product.specified /opt/bat/conf/
COPY ./conf/dev/iot_device_simulator.yml /opt/bat/conf/dev/
COPY ./iot_device_simulator /opt/bat/
RUN chmod ugo+x ./iot_device_simulator

CMD ["./iot_device_simulator"]
