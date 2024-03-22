# 社区版docker单机版部署

##    一、目录说明：
   1、service: 后台iot相关的服务
   2、third: 后台引用的第三方服务，如MySQL、Redis等
   3、simulator: 设备模拟器，模拟设备接收指令和上报数据
   4、web: web前端服务，前端代码在该路径的html下。

##    二、系统准备：

   1、服务器硬件级参数：2~4核CPU，8G以上内存
   2、服务器操作系统CentOS 7或以上版本，64位系统，其它64位的linux也可以，但未测试
   3、确保服务器时区正确，确保同步到了最新时间。

##    三、部署步骤：

###    1、配置docker网络

   在欲要部署的服务器上查看是否有docker_iot_network网络，没有则创建网络。
   查看docker网络：
   docker network ls
       NETWORK ID     NAME                 DRIVER    SCOPE
       03c964e72c05   bridge               bridge    local
       512e9aece1a3   host                 host      local
       ed991be6cb82   none                 null      local
   创建docker网络：
   docker network create --driver bridge --subnet 172.16.0.0/16 --gateway 172.16.0.1 docker_iot_network
   windows下如果常见网络后，docker启动失败，则需要重启winnat服务，linux下不需要。
   关闭：net stop winnat
   开启：net start winnat

###    2、拷贝部署文件

   将bin和deploy目录压缩，拷贝到centos服务器解压。   
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

###    3、第三方服务配置

####    (1)、对象存储服务(OSS)配置

   注册七牛的对象存储服务账号，阿里云或亚马逊的也可以。整个系统保持一致就行，比如用七牛的，所有服务就用七牛的存储服务。账号申请完后，直接配置到以下文件中：iot_app_api_service.yml、iot_app_oem_service.yml、iot_cloud_api_service.yml

####    (2)、APP消息推送账号配置

   极光推送
       注册极光推送账号。在iot_message_service.yml中填入appKey和secret。
   gorush配置
       参考gorush官方文件进行配置，注意gorush的ios证书配置，ios.key_path和ios.password。另外，将gorush的访问地址配置到iot_message_service.yml对应配置项中。

####    (3)、 IP地址相关服务账号配置

   阿里云IP地址服务：系统采用了阿里云的IP地址服务，请注册和购买阿里云的IP地址服务。取得APPCODE，填入iot_weather_service.yml服务。
   MaxMind服务：作为备用的IP地址服务，需要开通MaxMind服务。到MaxMind上开通IP地址数据库更新服务，取得license key。
   配置到iot_weather_service.yml文件的geo下的licensekey中。

####    (4)、阿里短信服务

   APP手机号码注册使用，收取验证码。若不配置，则不支持手机号码注册、登录。
   目前支持阿里云SMS服务。请到阿里云开通和配置短信服务，短信模板请参考云管平台的配置。阿里短信开通和配置好后，需要改以下2个地方：
       1 修改通知模板。注意修改对应的模板ID，要修改正确方可正常发送短信。云管平台或数据库修改。
       2 修改iot_message_service.yml的SMS配置。

####    (5)、邮件发送服务

   APP邮箱注册使用，收取验证码。若不配置，则不支持邮箱注册、登录。  
   注册SMTP邮箱，开通第三方软件登录和发送邮件服务。然后到iot_message_service.yml的smtp进行配置。

####    (6)、开通天气数据相关服务

   开通openweathermap服务:
   到openweathermap开通天气服务，取得apikey，配置到iot_weather_service.yml的weather下的apikey。
   开通apicn服务:
   到aqicn开通服务，取得token，配置到iot_weather_service.yml的weather下的aqicnToken。
   开通易客云天气服务:
   到易客云开通天气服务，获取appid和appsecret，分别配置到iot_weather_service.yml的yiketianqi下的appid和appsecret。

####    (7)、haproxy证书配置

   若要开启MQTT TLS连接，则需要生成和配置pem证书。具体可参考haproxy官方要求。

###    5、启动依赖的服务

​   进入third目录执行：
​   chmod ugo+w vernemq
​   chmod ugo+w redis
​   chmod ugo+w haproxy
​   chmod ugo-w ./mysql/conf/my.cnf
​   mkdir ./mysql/data
​   chmod ugo+w ./mysql/data

   启动third下所有的容器
   docker-compose up -d
   执行docker ps查看容器运行情况。
   如果某个容器运行失败，则执行
   docker-compose up {container_name}，查看控制台报错信息，根据报错信息提示修复。

###    6、启动iot服务

​   进入service目录启动所有
   docker-compose up -d

###    7、启动web服务
   进入web目录启动nginx服务   
   docker-compose up -d
   web目录下的html目录说明:
   iot-open-web:开放平台前端
   iot-platform-web:云管平台前端
   如果前端有更新，请记得重新打包前端，删除旧的内容，再覆盖新的内容。

###    8、启动虚拟设备服务
   如果需要使用虚拟调试，则需要启动虚拟设备服务。
   进入simulator目录，执行
   docker-compose up -d

## 四、系统默认账号

###   1、云管平台

   云管平台访问地址：http://{IP}:2888
   登录账号：admin  
   登录密码：Admin123
   其中的IP为服务器的IP地址，配置为开发机器能够访问到的IP地址。下同。

###  2、开放平台

   开放平台访问地址：http://{IP}:2887
   登录账号：opensource@dev.com
   登录密码：Admin123

###  3、任务调度平台(xxl-job)账号

   系统使用了任务调度中心xxl-job调度执行后台数据统计。xxl-job的访问地址为：
   http://{IP}:8088/xxl-job-admin/
   用户名：admin
   默认密码：123456

4、链路跟踪zipkin

​   微服务开发用，跟踪后端服务间的调用链路及各服务执行状况。
   访问地址：http://{IP}:9411/zipkin/

### 4、APP使用

​   社区版目前只支持安卓版，从app目录获取安卓版apk安装包进行安装。
​   安装后，打开APP，会要求输入参数配置。
​   服务器地址：http://{IP}:2886
​   租户ID：iotcode
![image-20240322101549253](https://osspublic.iot-aithings.com/docs/image-20240322101549253.png)



   ```