社区版docker单机版部署（可用于调试、测试）

一、目录说明：
1、service: 后台iot相关的服务
2、third: 后台引用的第三方服务，如MySQL、Redis等
3、simulator: 设备模拟器，模拟设备接收指令和上报数据
4、web: web前端服务，前端代码在该路径的html下。

二、系统准备：
1、服务器硬件级参数：2~4核CPU，8G以上内存
2、服务器操作系统CentOS 7或以上版本，64位系统，其它64位的linux也可以，但未测试
3、确保服务器时区正确，确保同步到了最新时间。

三、部署步骤：
1、创建自定义网络docker_iot_network
查看是否有docker_iot_network网络，没有则创建网络。

查看：
docker network ls
NETWORK ID     NAME                 DRIVER    SCOPE
03c964e72c05   bridge               bridge    local
512e9aece1a3   host                 host      local
ed991be6cb82   none                 null      local
创建：
docker network create --driver bridge --subnet 172.16.0.0/16 --gateway 172.16.0.1 docker_iot_network
windows下如果常见网络后，docker启动失败，则需要重启winnat服务，linux下不需要。
关闭：net stop winnat
开启：net start winnat

2、将bin和deploy目录压缩，拷贝到centos服务器解压。
(1)、修改配置和脚本的IP地址为本次部署的IP地址(不正确会引起设备无法连接和设备虚拟调试不可用的问题)
yml配置修改：
修改iot_cloud_api_service.yml配置，将webmqtt节的addr改为正确的MQTT地址。
可以进到bin/conf/open目录下，执行以下命令：
   sed -i 's/127.0.0.1:8883/192.168.5.56:8883/g' iot_cloud_api_service.yml
sql脚本修改：
修改iot_config_write.sql文件，将所有的127.0.0.1改为正确的IP地址，局域网或外网可访问的。(手机、设备能够访问到的IP地址)
可以进到third/mysql/docker-entrypoint-initdb.d目录，执行以下命令:
 例如，把iot_config_write.sql中的127.0.0.1改为你的IP地址，下边演示改为192.168.5.56
    sed -i 's/127.0.0.1/192.168.5.56/g' iot_config_write.sql
(2)、修改MySQL、Clickhouse、redis、nats等密码，不要用仓库配置文件中的初始密码，存在安全隐患
(3)、docker-compose中的MySQL、Clickhouse密码结合(2)对应修改

3、进入third目录执行：
  mkdir ./vernemq/data
  mkdir ./vernemq/log
  chmod -R ugo+w vernemq
  chmod -R ugo+w redis
  chmod -R ugo+w haproxy
  mkdir ./mysql/data
  chmod -R ./mysql/data
  chmod ugo-w ./mysql/conf/my.cnf

  启动third下所有的容器
  docker-compose up -d
  执行docker ps查看容器运行情况。
  如果某个容器运行失败，则执行
  docker-compose up {container_name}，查看控制台报错信息，根据报错信息提示修复。非管理员账号进入执行部署可能存在权限问题，注意出错服务目录下权限。

  初始化nats stream，进入deploy目录，执行jetstream_tool初始化stream，其中addrs参数为nats链接地址，跟服务配置文件的nats地址保持一致即可。
  示例：
  ./jetstream_tool --addrs "nats://iLmz8sCXjkTYuh@127.0.0.1:4222"

4、进入service目录启动所有iot服务
docker-compose up -d

5、进入web目录启动nginx服务
docker-compose up -d
web目录下的html目录说明:
iot-open-web:开放平台前端
iot-platform-web:云管平台前端
如果前端有更新，请记得重新打包前端，删除旧的内容，再覆盖新的内容。

6、如果需要使用虚拟调试，则需要开启模拟器。
进入simulator目录，执行
docker-compose up -d


