SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- 创建数据库
use iot_config;

-- 字典动态地址
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (122196223254953984, 0, 'open', 'http://127.0.0.1:2887/api/v1/platform/web/open/oem', 'oem_app_package_domain', '', '', 0, 0, '', 'PROD', 'P', '', 0, 0, '2022-07-12 16:59:29', '2022-07-12 16:59:29', NULL);

INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1539834052192665600, 0, 'open', 'http://127.0.0.1:2888/buildPackageQrCode.html', 'oem_app_default_download_url', '', '', 0, 0, '', 'OPEN', 'P', '', 0, 0, '2022-07-21 17:20:19', '2022-11-07 15:37:06', NULL);

INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4149303595445092352, 0, 'open', 'http://127.0.0.1:2887/api/v1/platform/web/open/oem/app/buildFinishNotify', 'oem_app_build_notify', '', '', 0, 0, '', 'PROD', 'P', '', 0, 0, '2022-07-12 14:16:52', '2022-07-12 14:16:52', NULL);

INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4276582131666354176, 0, 'open', 'http://127.0.0.1:2886', 'oem_app_build_region_server_url', '', '', 0, 0, '', 'PROD', 'P', '', 0, 0, '2022-07-12 14:16:22', '2022-07-12 17:05:47', NULL);


-- 区域数据
INSERT INTO `iot_config`.`t_sys_region_server` (`id`, `sid`, `limited_count`, `binded_count`, `mqtt_server`, `http_server`, `describe`, `en_describe`, `enabled`, `lat`, `lng`, `country`, `province`, `city`, `district`, `created_by`, `area_phone_number`, `websocket_server`, `mqtt_port`) VALUES (1, '1', 10000, 10000, '127.0.0.1', 'http://127.0.0.1:2886', '中国', 'China', 1, NULL, NULL, NULL, NULL, NULL, NULL, 0, '86', 'ws://127.0.0.1:8883/mqtt', '1883');


SET FOREIGN_KEY_CHECKS = 1;






