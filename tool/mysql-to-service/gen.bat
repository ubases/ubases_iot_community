
rem 生成整个库
sgen curd --config ./config.yaml --tableName t_opm_panel_image_asset
@REM sgen curd --config ./config.yaml --tableName t_scene_template_app_relation
@REM sgen curd --config ./config.yaml --tableName t_scene_template_condition
@REM sgen curd --config ./config.yaml --tableName t_scene_template_task

rem --tableName t_oem_app_android_cert
rem 生成某个表，后边加 --tableName {table name}
rem sgen curd --config ./config.yaml --tableName t_tpl_testcase_template



