# 该工具的用途：
从mysql数据库表生成数据库表增、删、改、查、列表等接口代码，以及proto文件。
配置文件是config.yaml，配置项说明请参考配置文件的注释。

* 首次创建数据库后，建议整库生成。
* 后续修改或新增表，建议单个表生成。

生成的文件默认放在gen文件夹，请务必在生成前清空该文件夹，以防和上次的混淆。
```
生成整个库:
sgen curd --config ./config.yaml

生成某个表，后边加 --tableName {table name}
如生成表t_sys_apis的代码，则执行以下命令:
sgen curd --config ./config.yaml --tableName t_config_translate_language
```
生成的文件夹解释说明：
* convert文件夹，微服务专用，数据库表model结构体和protobuf结构体数据转换函数
* entitys文件夹，API专用，数据库表操作相关的api实体结构体、查询结构体，以及实体结构体和protobuf结构体数据转换函数，可节省大量开发时间。请拷贝到api工程下合适的位置。
* handler文件夹，微服务专用，微服务处理入口
* service文件夹，微服务专用，默认都是数据库表操作逻辑，请根据实际业务自行修改。
* proto文件夹，proto文件，请务必拷贝到统一存放proto的位置，再生成go文件。

生成工具生成的文件，不能过度依赖，请根据实际逻辑自行修改代码。








