
-- 自定义固件增加英文名称
ALTER TABLE `iot_product`.`t_opm_firmware`
    ADD COLUMN `name_en` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '固件英文名称' AFTER `firmware_key`;

-- 将英文名称字段默认赋值为中文
update `iot_product`.t_opm_firmware set `name_en` = `name`;


-- 增加honor厂商推送配置
ALTER TABLE `iot_app_build`.`t_oem_app_push_cert`
    ADD COLUMN `honor` text NULL COMMENT 'honer推送配置，存json' AFTER `oppo`;

-- 设备相关更新脚本
ALTER TABLE `iot_device`.`t_iot_device_fault`
    ADD COLUMN `base_product_id` bigint(20) NULL DEFAULT NULL COMMENT '产品类型ID' AFTER `device_name`;

ALTER TABLE `iot_device`.`t_iot_device_fault`
    ADD COLUMN `fault_identifier` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '功能标识符' AFTER `product_name`;

ALTER TABLE `iot_device`.`t_iot_device_fault`
    ADD COLUMN `fault_dpid` smallint(6) NULL DEFAULT NULL COMMENT '功能dpid' AFTER `fault_identifier`;

ALTER TABLE `iot_device`.`t_iot_device_fault`
    MODIFY COLUMN `fault_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '故障编号,对应故障值' AFTER `fault_dpid`;

ALTER TABLE `iot_device`.`t_iot_device_fault`
    MODIFY COLUMN `fault_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '故障名称,对应故障名称' AFTER `fault_code`;


-- 推送脚本，增加描述
ALTER TABLE `iot_message`.`t_app_push_token` MODIFY COLUMN `app_push_platform` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'APP推送注册平台 ios、android、huawei\r\n1=huawei =2 xiaomi =3 oppo =4 vivo 5=honor  6=meizu 7=ios 8=android 9=wechat' AFTER `app_push_id`;


-- 是否已引导字段
ALTER TABLE `iot_open_system`.`t_open_user`
    ADD COLUMN `has_guided` smallint NULL COMMENT '是否已引导' AFTER `deleted_at`;


-- 产品相关更新脚本


ALTER TABLE `iot_product`.`t_opm_panel`
    ADD COLUMN `base_product_id` bigint(20) NULL DEFAULT NULL COMMENT '产品类型Id（基础产品Id、品类Id）' AFTER `source_zip`;

ALTER TABLE `iot_product`.`t_opm_panel`
    ADD COLUMN `product_id` bigint(20) NULL DEFAULT NULL COMMENT '开发者产品Id' AFTER `base_product_id`;


ALTER TABLE `iot_product`.`t_opm_panel`
    ADD COLUMN `lang_file_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '翻译文件名称' AFTER `theme_json`;

ALTER TABLE `iot_product`.`t_opm_panel`
    ADD COLUMN `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '面板描述' AFTER `lang_file_name`;


ALTER TABLE `iot_product`.`t_opm_product`
    ADD COLUMN `is_demo_product` smallint(6) NULL DEFAULT NULL COMMENT '是否demo产品' AFTER `tsl_updated_at`;


ALTER TABLE `iot_product`.`t_opm_product`
    MODIFY COLUMN `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '产品品类名称' AFTER `product_key`;



ALTER TABLE `iot_product`.`t_pm_product_type` CHARACTER SET = utf8mb4, COLLATE = utf8mb4_general_ci;

ALTER TABLE `iot_product`.`t_pm_product_type`
    MODIFY COLUMN `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '产品分类名称' AFTER `parent_id`;

ALTER TABLE `iot_product`.`t_pm_product_type`
    MODIFY COLUMN `name_en` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '产品分类名称（英文）' AFTER `name`;


-- APP面板状态描述更新
ALTER TABLE `iot_product`.`t_opm_panel`
    MODIFY COLUMN `status` smallint(6) NOT NULL COMMENT '面板状态 1 草稿（待启用）、2 开发中（待启用）、3 已发布（待启用）、4 已禁用' AFTER `panel_type`;



ALTER TABLE `iot_app`.`t_uc_user`
    ADD COLUMN `user_salt` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码盐';


-- 推送描述更新
ALTER TABLE `iot_message`.`t_app_push_token` MODIFY COLUMN `app_push_platform` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'APP推送注册平台 ios、android、huawei\r\n1=huawei =2 xiaomi =3 oppo =4 vivo 5=honor  6=meizu 7=ios 8=android 9=wechat' AFTER `app_push_id`;



-- 初始化数据

-- 将所有产品修改为非默认产品
update `iot_product`.`t_opm_product`  set is_demo_product = 2;


-- 字典数据同步
INSERT INTO `iot_config`.`t_config_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `value_type`, `remark`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `is_system`) VALUES (8793833814057975962, '工单状态', 'work_order_status', 0, 1, '', 0, 0, '2024-04-10 04:04:34', '2024-04-10 04:04:34', NULL, 0);
INSERT INTO `iot_config`.`t_config_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `value_type`, `remark`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `is_system`) VALUES (8793833814057975961, '面板状态', 'panel_status', 0, 1, '', 0, 0, '2024-04-08 10:31:45', '2024-04-08 10:31:45', NULL, 0);
INSERT INTO `iot_config`.`t_config_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `value_type`, `remark`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `is_system`) VALUES (8793833814057975960, '面板类型', 'panel_type', 0, 1, '', 0, 0, '2024-02-22 02:58:09', '2024-02-22 02:58:09', '2024-02-22 04:13:38', 0);


INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (50368649885745152, 4, '已关闭', '4', 'work_order_status', '', '', 0, 0, '', 'YIGUANBI', 'Y', '', 0, 0, '2024-04-10 04:05:22', '2024-04-10 07:30:35', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (120057602192277504, 2, '已回复', '2', 'work_order_status', '', '', 0, 0, '', 'YIHUIFU', 'Y', '', 0, 0, '2024-04-10 04:05:06', '2024-04-10 04:05:14', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (151424412237791232, 1, '待回复', '1', 'work_order_status', '', '', 0, 0, '', 'DAIHUIFU', 'D', '', 0, 0, '2024-04-10 04:04:58', '2024-04-10 04:05:10', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (232436372294500352, 0, '已禁用', '4', 'panel_status', '', '', 0, 0, '', 'YIJINYONG', 'Y', '', 0, 0, '2024-04-08 10:34:16', '2024-04-10 01:53:38', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (187846759072825344, 0, '已启用', '3', 'panel_status', '', '', 0, 0, '', 'YIQIYONG', 'Y', '', 0, 0, '2024-04-08 10:34:05', '2024-04-10 01:53:35', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (137028941951959040, 0, '待启用', '1', 'panel_status', '', '', 0, 0, '', 'DAIQIYONG', 'D', '', 0, 0, '2024-04-08 10:33:53', '2024-04-10 17:19:05', NULL);
# INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2332057701959237632, 7, '故障型（Fault）', 'FAULT', 'data_type', '', '', 0, 0, '', 'GUZHANGXINGFAULT', 'G', '', 0, 0, '2024-03-21 11:27:26', '2024-03-21 11:27:26', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2257349934489436160, 0, '定制面板', '2', 'panel_type', '', '', 0, 0, '', 'DINGZHIMIANBAN', 'D', '', 0, 0, '2024-02-22 02:58:37', '2024-02-22 02:58:37', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2215203809888468992, 0, '通用面板', '1', 'panel_type', '', '', 0, 0, '', 'TONGYONGMIANBAN', 'T', '', 0, 0, '2024-02-22 02:58:26', '2024-02-22 02:58:26', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3199327908233379840, 4, '其它', '4', 'device_mode_type', '', '', 0, 0, '', 'QITA', 'Q', '', 0, 0, '2024-01-31 14:16:07', '2024-01-31 14:16:07', NULL);

INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (7000682639950708736, 3232128938264133632, 'OfflinePanelDetail', '线下开发面板详情', 'copy', '', '', 1, 4, 1, 2, '/panel/offlinePanel/detail/index', '', '/panel/offlinePanel/detail/index', 2, 'sys_admin', 0, 2, 2, '2024-03-19 16:04:49', '2024-03-19 16:05:02', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (343706670901133312, 3232128938264133632, 'OfflinePanel', '线下开发面板管理', 'copy', '', '', 1, 3, 1, 1, '/panel/offlinePanel/index', '', '/panel/offlinePanel/index', 2, 'sys_admin', 0, 1, 2, '2024-03-19 14:07:24', '2024-03-19 16:05:08', NULL, 0, 0);

-- 云管平台菜单数据更新
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 7290673875191234560, `name` = 'DataLanguage', `title` = '数据多语言', `icon` = 'smile', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 1, `status` = 1, `always_show` = 1, `path` = '/multi-language/data-language', `jump_path` = '', `component` = '/multiLanguage/dataLanguage', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-07-11 19:30:04', `updated_at` = '2024-04-23 13:51:54', `deleted_at` = NULL WHERE `id` = 7432193774773501952;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 4904777198126137344, `name` = 'AppLog', `title` = 'APP用户日志', `icon` = 'setting', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 1, `status` = 1, `always_show` = 1, `path` = '/log/app-log', `jump_path` = '', `component` = '/log/appLog/index', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-07-11 19:21:23', `updated_at` = '2024-04-23 13:51:39', `deleted_at` = NULL WHERE `id` = 5244520642086076416;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 242, `name` = 'DictionariesList', `title` = '字典集列表', `icon` = 'dash', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 1, `status` = 1, `always_show` = 1, `path` = '/system/dictionaries/list', `jump_path` = '', `component` = '/system/dictionaries/index', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-04-29 09:27:09', `updated_at` = '2024-04-23 13:51:11', `deleted_at` = NULL WHERE `id` = 1591328788963033088;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 258, `name` = 'DeveloperAuditList', `title` = '开发者认证列表', `icon` = 'dash', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 1, `status` = 1, `always_show` = 1, `path` = '/authority/developer-audit/list', `jump_path` = '', `component` = '/authority/developerAudit/index', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-04-27 13:49:58', `updated_at` = '2024-04-23 13:50:32', `deleted_at` = NULL WHERE `id` = 7039214281463595008;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 2509984231870857216, `name` = 'DeveloperList', `title` = '开发者账号列表', `icon` = 'dash', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 1, `status` = 1, `always_show` = 1, `path` = '/authority/developer/list', `jump_path` = '', `component` = '/authority/developer/index', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-04-26 17:38:55', `updated_at` = '2024-04-23 13:50:10', `deleted_at` = NULL WHERE `id` = 2584332897636941824;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 244, `name` = 'UserList', `title` = '系统账号列表', `icon` = 'vertical-left', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 1, `status` = 1, `always_show` = 1, `path` = '/authority/user/list', `jump_path` = '', `component` = '/authority/user/index', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-05-19 09:48:43', `updated_at` = '2024-04-23 13:49:38', `deleted_at` = NULL WHERE `id` = 5206708362141401088;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 5325694820637638656, `name` = 'DeviceFaultData', `title` = '设备故障数据', `icon` = 'shopping-cart', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 6, `status` = 1, `always_show` = 1, `path` = '/data-center/device-fault-data', `jump_path` = '', `component` = '/dataCenter/deviceFaultData', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-06-06 16:37:57', `updated_at` = '2024-04-23 13:48:27', `deleted_at` = NULL WHERE `id` = 8792431261814521856;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 5325694820637638656, `name` = 'DeveloperData', `title` = '开发者数据', `icon` = 'link', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 4, `status` = 1, `always_show` = 1, `path` = '/data-center/developer-data', `jump_path` = '', `component` = '/dataCenter/developerData', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-06-06 16:35:35', `updated_at` = '2024-04-23 13:47:51', `deleted_at` = NULL WHERE `id` = 8198880678188580864;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 5325694820637638656, `name` = 'AppData', `title` = 'App数据', `icon` = 'share-alt', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 2, `status` = 1, `always_show` = 1, `path` = '/data-center/app-data', `jump_path` = '', `component` = '/dataCenter/appData', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-06-06 16:28:22', `updated_at` = '2024-04-23 13:47:18', `deleted_at` = NULL WHERE `id` = 6383903012183506944;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 1279949607547273216, `name` = 'Message', `title` = '消息模板', `icon` = 'inbox', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 3, `status` = 1, `always_show` = 1, `path` = '/template/message/index', `jump_path` = '', `component` = '/template/message/index', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-06-06 16:13:32', `updated_at` = '2024-04-23 13:47:02', `deleted_at` = NULL WHERE `id` = 2647934226384453632;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 1279949607547273216, `name` = 'Agreement', `title` = 'APP协议', `icon` = 'rocket', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 3, `status` = 1, `always_show` = 1, `path` = '/template/agreement/index', `jump_path` = '', `component` = '/template/agreement/index', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-04-27 16:59:39', `updated_at` = '2024-04-23 13:47:00', `deleted_at` = NULL WHERE `id` = 565449152205520896;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 1279949607547273216, `name` = 'Test', `title` = '测试用例', `icon` = 'tablet', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 2, `status` = 1, `always_show` = 1, `path` = '/template/test/index', `jump_path` = '', `component` = '/template/test/index', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-04-27 16:58:55', `updated_at` = '2024-04-23 13:46:57', `deleted_at` = NULL WHERE `id` = 750973458883641344;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 1279949607547273216, `name` = 'Inform', `title` = '通知模板', `icon` = 'customer-service', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 1, `status` = 1, `always_show` = 1, `path` = '/template/inform/index', `jump_path` = '', `component` = '/template/inform/index', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-04-27 16:57:55', `updated_at` = '2024-04-23 13:46:54', `deleted_at` = NULL WHERE `id` = 1002304094779572224;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 5597739597732872192, `name` = 'TemplateManage', `title` = 'app模板管理', `icon` = 'appstore', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 6, `status` = 1, `always_show` = 1, `path` = '/oem-app/template-manage', `jump_path` = '', `component` = '/oemApp/templateManage', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-10-31 16:57:44', `updated_at` = '2024-04-23 13:46:01', `deleted_at` = NULL WHERE `id` = 7855240489086320640;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 5597739597732872192, `name` = 'AssistPutaway', `title` = '辅助上架', `icon` = 'appstore', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 4, `status` = 2, `always_show` = 1, `path` = '/oem-app/assist-putaway', `jump_path` = '', `component` = '/oemApp/assistPutaway', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-10-20 13:57:57', `updated_at` = '2024-04-23 13:45:47', `deleted_at` = NULL WHERE `id` = 473174410618896384;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 5597739597732872192, `name` = 'HelpCenter', `title` = '帮助中心', `icon` = 'inbox', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 3, `status` = 1, `always_show` = 1, `path` = '/oem-app/help-center', `jump_path` = '', `component` = '/oemApp/helpCenter', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-07-05 14:11:47', `updated_at` = '2024-04-23 13:45:35', `deleted_at` = NULL WHERE `id` = 3506504086125969408;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 5597739597732872192, `name` = 'AppLang', `title` = 'App语言包', `icon` = 'meh', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 2, `status` = 1, `always_show` = 1, `path` = '/oem-app/app-lang', `jump_path` = '', `component` = '/oemApp/appLang', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-06-06 16:16:15', `updated_at` = '2024-04-23 13:45:19', `deleted_at` = NULL WHERE `id` = 3331615539670712320;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 246, `name` = 'ProductList', `title` = '产品列表', `icon` = 'dash', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 1, `status` = 1, `always_show` = 2, `path` = '/product/product/index', `jump_path` = '', `component` = '/product/product/index', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-04-25 11:24:57', `updated_at` = '2024-04-23 13:43:52', `deleted_at` = NULL WHERE `id` = 2084296393052028928;
UPDATE `iot_system`.`t_sys_auth_rule` SET `pid` = 252, `name` = 'ControlPanelList', `title` = '控制面板列表', `icon` = 'dash', `condition` = '', `remark` = '', `menu_type` = 1, `weigh` = 1, `status` = 1, `always_show` = 1, `path` = '/product/control-panel/list', `jump_path` = '', `component` = '/product/controlPanel/index', `is_frame` = 2, `module_type` = 'sys_admin', `model_id` = 0, `is_cache` = 1, `is_hide_child_menu` = 2, `created_at` = '2022-04-27 10:14:08', `updated_at` = '2024-04-23 13:43:47', `deleted_at` = NULL WHERE `id` = 8063493896488779776;