

ALTER TABLE `iot_system`.`t_sys_user`
    MODIFY COLUMN `address` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系地址' AFTER `is_admin`;

update iot_app.t_uc_home set sid='1' where sid='0';

ALTER TABLE `iot_product`.`t_pm_control_panels`
    MODIFY COLUMN `desc` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '面板描述' AFTER `lang`;

ALTER TABLE `iot_device`.`t_iot_device_fault`
    ADD COLUMN `tenant_id` VARCHAR (10) NULL COMMENT '租户ID' AFTER `fault_name`;

CREATE TABLE `t_cloud_app_build_auth` (
  `id` bigint(20) NOT NULL COMMENT '唯一Id',
  `name` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '申请名称',
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `auth_key` varchar(255) NOT NULL COMMENT '授权Key',
  `tenant_id` varchar(8) CHARACTER SET utf8 NOT NULL COMMENT '租户ID',
  `updated_at` timestamp NULL DEFAULT NULL,
  `update_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='云构建授权';


-- 产品授权APP关联表
CREATE TABLE iot_product.`t_opm_product_app_relation` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `product_id` bigint(20) NOT NULL COMMENT '产品ID(t_pm_product.id)',
  `product_key` varchar(20) NOT NULL COMMENT '产品Key',
  `app_key` varchar(50) NOT NULL COMMENT 'APP KEY',
  `app_name` varchar(100) DEFAULT NULL COMMENT 'APP名称',
  `created_by` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品授权APP使用关联表';


-- 私有云客户平台管理
CREATE TABLE iot_system.`t_sys_cloud_platform` (
   `id` bigint(20) unsigned NOT NULL,
   `type` smallint(1) NOT NULL COMMENT '私有云类型（=1 企业版私有云 =2 社区版私有云）',
   `code` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '平台编码',
   `name` varchar(60) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '平台名称',
   `platform_url` varchar(100) CHARACTER SET utf8mb4 NULL DEFAULT '' COMMENT '站点地址',
   `company_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '公司名称',
   `company_address` varchar(255) CHARACTER SET utf8mb4 NULL DEFAULT '' COMMENT '公司地址',
   `contact_person` varchar(50) COLLATE utf8mb4_bin NULL DEFAULT '' COMMENT '联系人',
   `contact_phone` varchar(50) CHARACTER SET utf8mb4 NULL DEFAULT '' COMMENT '联系电话',
   `describe` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '描述信息',
   `created_by` bigint(20) unsigned DEFAULT '0' COMMENT '创建人',
   `updated_by` bigint(20) DEFAULT NULL COMMENT '修改人',
   `status` smallint(1) DEFAULT NULL COMMENT '状态（=1 启动 =2禁用）',
   `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
   `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='私有云平台表';


-- 设备三元组增加所属平台
ALTER TABLE `iot_device`.`t_iot_device_triad`
    ADD COLUMN `platform_code` varchar(20) NULL COMMENT '归属平台编码' AFTER `export_time`;


-- 三元组增加是否其它平台三元组
ALTER TABLE `iot_device`.`t_iot_device_triad`
    ADD COLUMN `is_other_platform` smallint(1) NULL COMMENT '是否其它私有化部署平台' AFTER `platform_code`;

-- 设备信息增加激活所属区域
ALTER TABLE `iot_device`.`t_iot_device_info`
    ADD COLUMN `region_server_id` varchar(20) NULL COMMENT '区域服务器Id';

ALTER TABLE `iot_device`.`t_iot_device_triad`
    ADD COLUMN `export_time_list` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '导出时间列表' AFTER `is_other_platform`;

-- 增加app升级提醒方式
ALTER TABLE `iot_app_build`.`t_oem_app_custom_record`
    ADD COLUMN `remind_mode` smallint(1) NULL COMMENT '更新提醒模式 =1 红点提醒 =2 弹框提醒 =3 强制升级提醒' AFTER `launch_markets`;

-- 增加协议提醒方式
ALTER TABLE `iot_app_build`.`t_oem_app_introduce`
    ADD COLUMN `remind_mode` smallint(1) NULL COMMENT '更新提醒模式 =1 不提醒 =2 弹框提醒';


update `iot_app_build`.`t_oem_app_custom_record`  set remind_mode = 1;
update `iot_app_build`.`t_oem_app_introduce`  set remind_mode = 1;


-- 邮件和短信通知记录
CREATE TABLE iot_message.`t_ms_notice_record` (
                                                  `id` bigint(20) NOT NULL COMMENT '主键ID',
                                                  `lang` varchar(10) DEFAULT NULL COMMENT '语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语]',
                                                  `platform` smallint(1) DEFAULT NULL COMMENT '平台 =1 云管&开放 =2 APP =3 其它',
                                                  `account` varchar(50) DEFAULT NULL COMMENT '账号手机号或者邮箱',
                                                  `method` int(11) NOT NULL COMMENT '通知方式[1:邮件，2：短信]',
                                                  `sms_supplier` int(2) DEFAULT NULL COMMENT '短信服务供应商（短信服务提供商, 默认0）',
                                                  `type` int(11) DEFAULT NULL COMMENT '消息类型 1-注册 2-忘记密码 3-修改密码 4-注销账号 5-绑定第三方账号 6-绑定手机或邮箱 7-验证码登录',
                                                  `thirdpary_code` varchar(64) DEFAULT NULL COMMENT '第三方模板编码（当method=2的时候，值为对应短信平台的编码）',
                                                  `notice_tempate_id` bigint(20) DEFAULT NULL COMMENT '通知模板Id',
                                                  `title` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '标题',
                                                  `content` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '发送内容',
                                                  `app_key` varchar(50) DEFAULT NULL COMMENT 'app key',
                                                  `tenant_id` varchar(20) DEFAULT NULL COMMENT '开发者租户编号',
                                                  `status` int(1) DEFAULT NULL COMMENT '发送状态 =1 发送成功 =2 发送失败',
                                                  `send_err_msg` varchar(20) DEFAULT NULL COMMENT '发送结果描述，失败原因展示',
                                                  `created_by` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
                                                  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='短信、邮件消息记录';


-- 附件表
CREATE TABLE iot_system.`t_sys_attachment` (
                                               `id` bigint(20) NOT NULL COMMENT '主键Id',
                                               `file_name` varchar(100) NOT NULL COMMENT '文件名称（包括目录和文件名称，例如：dir/xxxxx.png）',
                                               `file_size` int(10) NOT NULL COMMENT '文件大小',
                                               `file_type` varchar(50) NOT NULL COMMENT '文件类型',
                                               `file_url` varchar(200) NOT NULL COMMENT '文件地址',
                                               `file_md5` varchar(50) DEFAULT NULL COMMENT '文件MD5',
                                               `oss_platform` int(2) NOT NULL COMMENT 'oss归属平台，1 阿里云oss、2 亚马逊s3、3：七牛',
                                               `source_table` varchar(50) NULL COMMENT '来源的表，待定',
                                               `source_row_id` varchar(50) NULL COMMENT '来源的行id，待定',
                                               `status` smallint(1) NULL COMMENT '业务表单提交成功，数据标记为有效，否则需要定时清理 ',
                                               `allow_clear` smallint(1) NULL COMMENT '是否容许清理 =1 可清理 =2 不可清理 ',
                                               `created_by` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
                                               `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                               `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间（标记删除时间的数据，将被列入到回收站，等待定时任务清理）',
                                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='附件表';


-- APP升级增加提醒描述
ALTER TABLE `iot_app_build`.`t_oem_app_custom_record`
    ADD COLUMN `remind_desc` varchar(200) NULL COMMENT '提醒描述中文' AFTER `remind_mode`,
ADD COLUMN `remind_desc_en` varchar(200) NULL COMMENT '提醒描述英文' AFTER `remind_desc`;


-- 用户信息增加协议弹框提醒
ALTER TABLE `iot_app`.`t_uc_user`
    ADD COLUMN `agreement_flag` smallint(1) NULL COMMENT '协议标记 =1 弹框提醒  其它不弹框' AFTER `user_salt`;

-- APP 构建记录增加构建来源
ALTER TABLE `iot_app_build`.`t_oem_app_build_record`
    ADD COLUMN `build_origin` smallint(1) NULL COMMENT '构建来源 = 2 私有云平台打包申请，如果' AFTER `app_template_version`;

-- 设备配置网线初始区域服务Id
update iot_device.t_iot_device_info d inner join iot_app.t_uc_user u on d.active_user_id=u.id
    set d.region_server_id = u.region_server_id where d.active_user_id is not null ;

-- ALTER TABLE `iot_device`.`t_iot_device_triad`
-- ADD COLUMN `region_server_id` varchar(20) NULL COMMENT '区域服务器Id';

ALTER TABLE `iot_app`.`t_uc_user`
    MODIFY COLUMN `region_server_id` bigint(20) NULL DEFAULT NULL COMMENT '注册区域服务器编号' AFTER `birthday`,
    ADD COLUMN `register_region_id` bigint(20) NULL DEFAULT NULL COMMENT '注册区域Id' AFTER `agreement_flag`;

-- 授权数量修改时间字段，不自动更新
ALTER TABLE `iot_open_system`.`t_open_auth_quantity`
    MODIFY COLUMN `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间,有值表示已删除' AFTER `updated_at`;

-- 面板增加面板编码
ALTER TABLE iot_product.`t_pm_control_panels`
    ADD COLUMN `code`  varchar(50) NULL COMMENT '面板编码';

ALTER TABLE iot_product.`t_opm_panel`
    ADD COLUMN `code`  varchar(50) NULL COMMENT '面板编码' ;

-- 场景增加实例编码
ALTER TABLE iot_device.`t_scene_intelligence`
    ADD COLUMN `instance_code`  varchar(20) NULL COMMENT '实例编码' AFTER `region_server_id`;

ALTER TABLE `iot_message`.`t_ms_notice_record`
    MODIFY COLUMN `content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '发送内容' AFTER `title`;


ALTER TABLE `iot_message`.`t_ms_notice_record`
    MODIFY COLUMN `send_err_msg` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '发送结果描述，失败原因展示' AFTER `status`;
ALTER TABLE `iot_message`.`t_ms_notice_record`
    MODIFY COLUMN `content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '发送内容' AFTER `title`;


ALTER TABLE `iot_message`.`t_ms_notice_record`
    MODIFY COLUMN `send_err_msg` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '发送结果描述，失败原因展示' AFTER `status`;
-- 初始化数据
INSERT INTO `iot_open_system`.`t_open_auth_quantity` (`id`, `user_id`, `company_id`, `tenant_id`, `auth_code`, `auth_quantity`, `auth_date`, `status`, `created_by`, `created_at`, `updated_by`, `updated_at`)
select a.id as id, a.id as user_id, b.company_id as company_id, b.tenant_id as tenant_id, 'default' as auth_code,
       (case
            when a.account_type=1 and c.status=3 then 300
            when a.account_type=1 and c.status<>3 then 100
            else 50 end) as auth_quantity, a.created_at as auth_date, 1 as status, a.created_by, a.created_at, a.updated_by, a.updated_at from `iot_open_system`.t_open_user as a
                                                                                                                                                   left join  `iot_open_system`.t_open_user_company as b on a.id=b.user_id and user_type=1
                                                                                                                                                   left join  `iot_open_system`.t_open_company as c on c.id=b.company_id
where not exists (select null from `iot_open_system`.`t_open_auth_quantity` as cc where cc.id=a.id);


-- 菜单
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1667294086220054528, 4100852806750142464, 'CloudPackage', 'OEM APP云构建服务', 'cloud-download', '', '', 1, 9, 1, 1, '/appDevelop/cloudPackage/index', '', '/appDevelop/cloudPackage/index', 2, 'sys_admin', 0, 2, 2, '2024-06-13 10:57:49', '2024-06-21 10:05:19', NULL, 0, 0);

-- 菜单
INSERT INTO `iot_system`.`t_sys_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`) VALUES (1345501258991108096, 238, 'AppApiService', 'API服务管理', 'font-colors', '', '', 1, 4, 1, 1, '/system/appApiService/index', '', '/system/appApiService/index', 2, 'sys_admin', 0, 1, 2, '2024-06-14 13:51:43', '2024-06-24 11:22:10', NULL);
INSERT INTO `iot_system`.`t_sys_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`) VALUES (8258072103354466304, 5325694820637638656, 'MailRecord', '邮箱发送记录', 'mail', '', '', 1, 8, 1, 1, '/data-center/mail-record', '', '/dataCenter/mailRecord', 2, 'sys_admin', 0, 1, 2, '2024-06-13 16:38:08', '2024-06-13 16:40:03', NULL);
INSERT INTO `iot_system`.`t_sys_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`) VALUES (8573705225020473344, 5325694820637638656, 'SmsRecord', '短信发送记录', 'message', '', '', 1, 7, 1, 1, '/data-center/sms-record', '', '/dataCenter/smsRecord', 2, 'sys_admin', 0, 1, 2, '2024-06-13 16:36:53', '2024-06-13 16:39:51', NULL);

-- 字典数据
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1028795541917630464, 2, '禁用', '2', 'is_enabled', '', '', 0, 0, '', 'JINYONG', 'J', '', 0, 0, '2024-06-17 09:04:31', '2024-06-17 09:04:31', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1059419057394122752, 1, '启用', '1', 'is_enabled', '', '', 0, 0, '', 'QIYONG', 'Q', '', 0, 0, '2024-06-17 09:04:24', '2024-06-17 09:04:24', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2421243976733523968, 3, '全部', '0', 'export_status', '', '', 0, 0, '', 'QUANBU', 'Q', '', 0, 0, '2024-06-14 16:14:03', '2024-06-14 16:14:08', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2454918262048260096, 2, '已导出', '1', 'export_status', '', '', 0, 0, '', 'YIDAOCHU', 'Y', '', 0, 0, '2024-06-14 16:13:55', '2024-06-14 16:13:55', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2492631077240274944, 1, '未导出', '2', 'export_status', '', '', 0, 0, '', 'WEIDAOCHU', 'W', '', 0, 0, '2024-06-14 16:13:46', '2024-06-14 16:13:46', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5365612575331549184, 2, '发送失败', '2', 'sms_send_status', '', '', 0, 0, '', 'FASONGSHIBAI', 'F', '', 0, 0, '2024-06-13 16:49:38', '2024-06-13 16:49:38', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5498016148265992192, 1, '发送成功', '1', 'sms_send_status', '', '', 0, 0, '', 'FASONGCHENGGONG', 'F', '', 0, 0, '2024-06-13 16:49:06', '2024-06-13 16:49:06', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (6365614519187767296, 2, '国际短信', '2', 'sms_notice_type', '', '', 0, 0, '', 'GUOJIDUANXIN', 'G', '', 0, 0, '2024-06-13 16:45:40', '2024-06-25 11:58:35', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (6394600187930705920, 1, '短信', '1', 'sms_notice_type', '', '', 0, 0, '', 'DUANXIN', 'D', '', 0, 0, '2024-06-13 16:45:33', '2024-06-25 14:13:21', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2730522856919040000, 5, '已关闭', '5', 'work_order_status', '', '', 0, 0, '', 'YIGUANBI', 'Y', '', 0, 0, '2024-06-12 15:42:29', '2024-06-12 15:42:29', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2623068874110566400, 3, '线下回访中', '3', 'work_order_status', '', '', 0, 0, '', 'XIANXIAHUIFANGZHONG', 'X', '', 0, 0, '2024-06-12 15:42:04', '2024-06-12 15:42:04', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1463359995612594176, 0, '技术类', '2', 'work_order_category', '', '', 0, 0, '', 'JISHULEI', 'J', '', 0, 0, '2024-06-12 15:25:49', '2024-06-12 15:25:49', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1493216274455560192, 0, '业务类', '1', 'work_order_category', '', '', 0, 0, '', 'YEWULEI', 'Y', '', 0, 0, '2024-06-12 15:25:42', '2024-06-12 15:25:42', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (7092812905971613696, 3, '强制升级提醒', '3', 'app_shelf_remind_mode', '', '', 0, 0, '', 'QIANGZHISHENGJITIXING', 'Q', '', 0, 0, '2024-06-11 11:53:54', '2024-06-11 11:53:54', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (7036910552987107328, 2, '弹框提醒', '2', 'app_shelf_remind_mode', '', '', 0, 0, '', 'DANKUANGTIXING', 'D', '', 0, 0, '2024-06-11 11:53:41', '2024-06-11 11:53:41', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (7006470218452992000, 1, '红点提醒', '1', 'app_shelf_remind_mode', '', '', 0, 0, '', 'HONGDIANTIXING', 'H', '', 0, 0, '2024-06-11 11:53:34', '2024-06-11 11:53:34', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2090315212409176064, 2, '弹窗提醒', '2', 'protocol_remind_type', '', '', 0, 0, '', 'DANCHUANGTIXING', 'D', '', 0, 0, '2024-06-06 17:57:03', '2024-06-06 17:57:03', NULL);
INSERT INTO `iot_config`.`t_config_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `remark`, `pinyin`, `firstletter`, `listimg`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2013320667906080768, 1, '不提醒', '1', 'protocol_remind_type', '', '', 0, 0, '', 'BUTIXING', 'B', '', 0, 0, '2024-06-06 17:56:45', '2024-06-06 17:57:06', NULL);

-- 字典类型
INSERT INTO `iot_config`.`t_config_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `value_type`, `remark`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `is_system`) VALUES (8793833814057975969, '启用状态', 'is_enabled', 0, 1, '', 0, 0, '2024-06-17 09:04:11', '2024-06-17 09:04:11', NULL, 0);
INSERT INTO `iot_config`.`t_config_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `value_type`, `remark`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `is_system`) VALUES (8793833814057975968, '导出状态', 'export_status', 0, 1, '', 0, 0, '2024-06-14 16:12:29', '2024-06-14 16:12:29', NULL, 0);
INSERT INTO `iot_config`.`t_config_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `value_type`, `remark`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `is_system`) VALUES (8793833814057975967, '短信发送状态', 'sms_send_status', 0, 1, '', 0, 0, '2024-06-13 16:48:46', '2024-06-13 16:48:46', NULL, 0);
INSERT INTO `iot_config`.`t_config_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `value_type`, `remark`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `is_system`) VALUES (8793833814057975966, '短信通知类型', 'sms_notice_type', 0, 1, '', 0, 0, '2024-06-13 16:43:30', '2024-06-13 16:43:30', NULL, 0);
INSERT INTO `iot_config`.`t_config_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `value_type`, `remark`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `is_system`) VALUES (8793833814057975965, '工单分类', 'work_order_category', 0, 1, '', 0, 0, '2024-06-12 15:24:37', '2024-06-12 15:24:37', NULL, 0);
INSERT INTO `iot_config`.`t_config_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `value_type`, `remark`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `is_system`) VALUES (8793833814057975964, 'APP上架提醒方式', 'app_shelf_remind_mode', 0, 1, '', 0, 0, '2024-06-11 11:50:51', '2024-06-11 11:50:51', NULL, 0);
INSERT INTO `iot_config`.`t_config_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `value_type`, `remark`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `is_system`) VALUES (8793833814057975963, '协议更新提醒方式', 'protocol_remind_type', 0, 1, '', 0, 0, '2024-06-06 17:55:45', '2024-06-06 17:55:45', NULL, 0);
INSERT INTO `iot_config`.`t_config_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `value_type`, `remark`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`, `is_system`) VALUES (8793833814057975971, '规则条件类型', 'rule_condition_type', 0, 1, '', 0, 0, '2024-07-18 14:48:39', '2024-07-18 14:48:39', NULL, 0);

-- 部分字典图片地址修改
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 0, `dict_label` = '绿色主题', `dict_value` = '1', `dict_type` = 'skin_code', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'LVSEZHUTI', `firstletter` = 'L', `listimg` = 'https://osspublic.iot-aithings.com//opt/bat/temp/dic//08d0efc2-7b2c-4a93-a073-62a1bfa70c35.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-11-11 15:13:35', `updated_at` = '2022-11-14 10:54:12', `deleted_at` = NULL WHERE `dict_code` = 4597521919623725056;
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 0, `dict_label` = '紫色主题', `dict_value` = '2', `dict_type` = 'skin_code', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'ZISEZHUTI', `firstletter` = 'Z', `listimg` = 'https://osspublic.iot-aithings.com//opt/bat/temp/dic//2239d32e-4edf-49d1-afd4-dea657c25355.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-11-11 15:14:02', `updated_at` = '2022-11-14 10:54:18', `deleted_at` = NULL WHERE `dict_code` = 4708155170741977088;
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 0, `dict_label` = '授权管理', `dict_value` = 'Authority', `dict_type` = 'open_dashboard_icons', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'SHOUQUANGUANLI', `firstletter` = 'S', `listimg` = 'https://osspublic.iot-aithings.com//opt/bat/temp/dic//3ae367b9-54e4-4a6e-a7e7-46e97c0b38a1.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-08-25 17:14:45', `updated_at` = '2022-08-25 17:15:19', `deleted_at` = NULL WHERE `dict_code` = 7647267563906367488;
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 0, `dict_label` = '营销', `dict_value` = 'Marketing', `dict_type` = 'open_dashboard_icons', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'YINGXIAO', `firstletter` = 'Y', `listimg` = 'https://osspublic.iot-aithings.com//opt/bat/temp/dic//542fdd59-b3f2-451f-8346-665cf1d20bdf.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-08-25 17:14:26', `updated_at` = '2022-08-25 17:15:31', `deleted_at` = NULL WHERE `dict_code` = 7725224209445650432;
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 0, `dict_label` = '数据管理', `dict_value` = 'DataManage', `dict_type` = 'open_dashboard_icons', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'SHUJUGUANLI', `firstletter` = 'S', `listimg` = 'https://osspublic.iot-aithings.com/dic/7819a629-c269-43e1-b793-24357fec9655.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-08-25 17:13:58', `updated_at` = '2024-05-13 11:08:44', `deleted_at` = NULL WHERE `dict_code` = 7844142060369182720;
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 0, `dict_label` = '产品开发', `dict_value` = 'Product', `dict_type` = 'open_dashboard_icons', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'CHANPINKAIFA', `firstletter` = 'C', `listimg` = 'https://osspublic.iot-aithings.com//opt/bat/temp/dic//a4fb92cc-b360-4519-b00e-309b7c9b6b97.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-08-25 17:13:34', `updated_at` = '2022-09-07 10:15:18', `deleted_at` = NULL WHERE `dict_code` = 7942520399491858432;
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 0, `dict_label` = 'APP开发', `dict_value` = 'AppDevelop', `dict_type` = 'open_dashboard_icons', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'APPKAIFA', `firstletter` = 'A', `listimg` = 'https://osspublic.iot-aithings.com//opt/bat/temp/dic//88dc4895-042a-4501-a4b1-7e18d001c766.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-08-25 17:12:52', `updated_at` = '2022-08-25 17:15:15', `deleted_at` = NULL WHERE `dict_code` = 8122090702550171648;
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 3, `dict_label` = '设备控制问题', `dict_value` = '3', `dict_type` = 'feedback_question_type', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'SHEBEIKONGZHIWENTI', `firstletter` = 'S', `listimg` = 'https://osspublic.iot-aithings.com//opt/bat/temp/dic//f9cbe49e-23ba-4f65-a7f7-01d9270a9aec.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-08-10 15:35:17', `updated_at` = '2022-11-10 23:56:41', `deleted_at` = NULL WHERE `dict_code` = 8260227614233427968;
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 0, `dict_label` = '数据中心', `dict_value` = 'data_center', `dict_type` = 'open_dashboard_icons', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'SHUJUZHONGXIN', `firstletter` = 'S', `listimg` = 'https://osspublic.iot-aithings.com/dic/fb2122c3-ee40-44ae-9624-9627da4ecdad.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-08-25 17:11:46', `updated_at` = '2024-05-13 11:08:21', `deleted_at` = NULL WHERE `dict_code` = 8398005706606870528;
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 0, `dict_label` = '云服务管理', `dict_value` = 'cloud_manage', `dict_type` = 'open_dashboard_icons', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'YUNFUWUGUANLI', `firstletter` = 'Y', `listimg` = 'https://osspublic.iot-aithings.com//opt/bat/temp/dic//779e4dee-1e58-46b7-9d60-212a41baee16.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-08-25 17:11:19', `updated_at` = '2022-08-25 17:11:19', `deleted_at` = NULL WHERE `dict_code` = 8511663738758201344;
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 0, `dict_label` = 'APP管理', `dict_value` = 'app_manage', `dict_type` = 'open_dashboard_icons', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'APPGUANLI', `firstletter` = 'A', `listimg` = 'https://osspublic.iot-aithings.com//opt/bat/temp/dic//6dbdd6b1-761d-4d0b-87af-9dd706e8148e.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-08-25 17:10:54', `updated_at` = '2022-08-25 17:10:54', `deleted_at` = NULL WHERE `dict_code` = 8614544802798731264;
UPDATE `iot_config`.`t_config_dict_data` SET `dict_sort` = 0, `dict_label` = '产品管理', `dict_value` = 'product_manage', `dict_type` = 'open_dashboard_icons', `css_class` = '', `list_class` = '', `is_default` = 0, `status` = 0, `remark` = '', `pinyin` = 'CHANPINGUANLI', `firstletter` = 'C', `listimg` = 'https://osspublic.iot-aithings.com//opt/bat/temp/dic//fbab2e7c-3be9-4aff-948a-2cf5a0f3aa95.png', `created_by` = 0, `updated_by` = 0, `created_at` = '2022-08-25 17:10:23', `updated_at` = '2023-01-29 14:20:19', `deleted_at` = NULL WHERE `dict_code` = 8746678868611530752;

INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (2100407803599486976, 4100852806750142464, 'OEMApp', 'OEM App开发', 'icon-APPkaifa', '', '', 1, 1, 1, 1, '/appDevelop/OEMApp/index', '', '/appDevelop/OEMApp/index', 2, 'sys_admin', 0, 2, 2, '2022-05-27 10:21:12', '2024-06-27 11:23:21', NULL, 0, 0);