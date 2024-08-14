SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- 创建数据库
DROP database IF EXISTS `iot_message`;
create database iot_message;
use iot_message;

-- ----------------------------
-- Table structure for t_app_push_token
-- ----------------------------
DROP TABLE IF EXISTS `t_app_push_token`;
CREATE TABLE `t_app_push_token`  (
                                   `id` bigint(20) NOT NULL COMMENT '主键ID',
                                   `app_token` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'APP生成的推送Token',
                                   `app_push_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'APP生成的手机唯一Id',
                                   `app_push_platform` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'APP推送注册平台 ios、android、huawei',
                                   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `updated_at` timestamp NOT NULL DEFAULT now() COMMENT '修改时间',
                                   `last_push_time` timestamp NULL DEFAULT NULL COMMENT '最后推送时间',
                                   `app_key` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'APP Key',
                                   `app_packet_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'APP包名',
                                   `lang` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '登录APP所属语言',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_app_push_token
-- ----------------------------

-- ----------------------------
-- Table structure for t_app_push_token_user
-- ----------------------------
DROP TABLE IF EXISTS `t_app_push_token_user`;
CREATE TABLE `t_app_push_token_user`  (
                                        `id` bigint(20) NOT NULL COMMENT '主键ID',
                                        `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户Id',
                                        `app_push_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'APP生成的手机唯一Id',
                                        `app_key` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'APP Key',
                                        `tenant_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '开发者租户编号',
                                        `region_id` bigint(20) NULL DEFAULT NULL COMMENT '区域服务Id',
                                        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_app_push_token_user
-- ----------------------------
INSERT INTO `t_app_push_token_user` VALUES (231934105493602304, 165935358154801152, 'undefined', 'iot-open-source-app', 'ioqp4r', 0);
INSERT INTO `t_app_push_token_user` VALUES (1996226041160499200, 3410520306123440128, 'undefined', 'iot-open-source-app', 'ioqp4r', 0);

-- ----------------------------
-- Table structure for t_mp_message
-- ----------------------------
DROP TABLE IF EXISTS `t_mp_message`;
CREATE TABLE `t_mp_message`  (
                               `id` bigint(20) NOT NULL,
                               `push_type` int(11) NOT NULL COMMENT '推送类型[0:站内且站外，1：站内， 2：站外]',
                               `message_type` int(11) NOT NULL COMMENT '消息类型[0：系统提醒，1：设备提醒，....自定义]',
                               `push_to` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '接收对象[数据字典配置， 比如全用户：all、user、home、device]',
                               `push_mode` int(11) NOT NULL COMMENT '发送模式[0:实时发送， 1：定时发送，2：轮询发送]',
                               `push_status` int(11) NOT NULL COMMENT '发送状态[0:已发送， 1：待发送， 2：发送失败， 3：已删除]',
                               `agent_type` int(11) NOT NULL COMMENT '终端类型[0:所有终端， 1：IOS端， 2:android端]',
                               `push_time` date NOT NULL COMMENT '推送时间',
                               `expire_hour` int(11) NULL DEFAULT NULL COMMENT '有效时间(单位小时)',
                               `action_type` int(11) NOT NULL COMMENT '行为类型[0:主动发送， 1:被动发送]',
                               `tpl_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '消息模板编号',
                               `push_params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '消息模板的参数',
                               `target_ids` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '推送的目标编号集合(push_to = user  为userid集合）,使用逗号分割',
                               `did` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                               `product_key` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '推送目标产品',
                               `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '消息内容',
                               `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
                               `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for t_mp_message_black
-- ----------------------------
DROP TABLE IF EXISTS `t_mp_message_black`;
CREATE TABLE `t_mp_message_black`  (
                                     `id` bigint(20) NOT NULL COMMENT '主键ID',
                                     `type` int(11) NOT NULL COMMENT '类型[0：全部，1：站外拦截，2：站内拦截]',
                                     `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户ID(t_cloud_uc_user.id)',
                                     `did` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '设备DID',
                                     `expire_start_time` date NULL DEFAULT NULL COMMENT '有效起始时间',
                                     `expire_end_time` date NULL DEFAULT NULL COMMENT '有效截止时间',
                                     `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
                                     `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
                                     `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                     `updated_at` timestamp NULL DEFAULT NULL COMMENT '修改时间',
                                     `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '消息黑名单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_mp_message_black
-- ----------------------------

-- ----------------------------
-- Table structure for t_mp_message_red_dot
-- ----------------------------
DROP TABLE IF EXISTS `t_mp_message_red_dot`;
CREATE TABLE `t_mp_message_red_dot`  (
                                       `id` bigint(20) NOT NULL COMMENT '唯一主键',
                                       `user_id` bigint(20) NOT NULL COMMENT '用户编号',
                                       `home_msg` smallint(1) NULL DEFAULT 0 COMMENT '家庭消息读取情况',
                                       `system_msg` smallint(1) NULL DEFAULT NULL COMMENT '系统消息读取情况',
                                       `system_msg_id` bigint(20) NULL DEFAULT NULL,
                                       `device_msg` smallint(1) NULL DEFAULT 0 COMMENT '设备消息读取情况',
                                       `public_msg` bigint(20) NULL DEFAULT NULL COMMENT '公共消息读取情况（最新公共消息编号）',
                                       `feedback_msg` smallint(1) NULL DEFAULT 0 COMMENT '反馈消息读取情况',
                                       `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
                                       `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
                                       `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                       `updated_at` timestamp NULL DEFAULT NULL COMMENT '修改时间',
                                       `system_msg_num` int(11) NULL DEFAULT NULL COMMENT '系统消息未读数量',
                                       `device_msg_num` int(11) NULL DEFAULT NULL COMMENT '设备消息未读数量',
                                       `public_msg_num` int(11) NULL DEFAULT NULL COMMENT '公共消息未读数量',
                                       `home_msg_num` int(11) NULL DEFAULT NULL COMMENT '家庭消息未读数量',
                                       `feedback_msg_num` int(11) NULL DEFAULT NULL COMMENT '反馈消息未读数量',
                                       PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ---------------------------
-- Table structure for t_mp_message_template
-- ----------------------------
DROP TABLE IF EXISTS `t_mp_message_template`;
CREATE TABLE `t_mp_message_template`  (
                                        `id` bigint(20) NOT NULL COMMENT '主键ID',
                                        `tpl_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '模板编码',
                                        `tpl_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '模板名称',
                                        `tpl_content` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '模板内容',
                                        `tpl_params` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '模板参数',
                                        `push_type` int(11) NOT NULL DEFAULT 0 COMMENT '消息平台[0:APP消息，1：云管平台消息2， 2：开放平台消息]',
                                        `message_type` int(11) NOT NULL COMMENT '消息类型[0：系统提醒，1：设备提醒，....自定义]',
                                        `agent_type` int(11) NOT NULL DEFAULT 0 COMMENT '接收终端类型[0:所有终端， 1：IOS端， 2:android端]',
                                        `lang` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语]',
                                        `expire_hour` int(11) NULL DEFAULT NULL COMMENT '有效时间(单位小时)',
                                        `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
                                        `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
                                        `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                        `updated_at` timestamp NULL DEFAULT NULL COMMENT '修改时间',
                                        `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_mp_message_template
-- ----------------------------
INSERT INTO `t_mp_message_template` VALUES (22164327880359936, 'QuitHomeMessage', '成员退出家庭（不删）', '“{{.userName}}”退出了家庭“{{.homeName}}”', '', 1, 1, 0, '', 0, 0, 0, '2022-07-18 15:56:15', '2022-07-18 15:56:15', NULL);
INSERT INTO `t_mp_message_template` VALUES (80035263098748928, 'RemoveMembersMessage', '被移除家庭（不删）', '您被家庭管理员从家庭“{{.homeName}}”中移除', '', 1, 1, 0, '', 0, 0, 0, '2022-07-18 15:55:50', '2022-07-18 15:55:50', NULL);
INSERT INTO `t_mp_message_template` VALUES (137406513831378944, 'RemoveDeviceMessage', '移除设备（不删）', '“{{.userName}}”将家庭“{{.homeName}}”中 “{{.deviceName}}”移除了', '', 1, 1, 0, '', 0, 0, 0, '2022-06-26 23:30:34', '2022-06-26 23:32:34', NULL);
INSERT INTO `t_mp_message_template` VALUES (347710295163961344, 'RemoveHomeMessage', '家庭被移除（不删）', '“{{.userName}}”移除了家庭“{{.homeName}}”', '', 1, 1, 0, '', 0, 0, 0, '2022-06-26 23:31:24', '2022-07-08 16:39:00', NULL);
INSERT INTO `t_mp_message_template` VALUES (431031441455742976, 'JoinHomeMessage', '新成员加入（不删）', '“{{.inviteUser}}”邀请用户“{{.invitedUser}}”加入 “{{.homeName}}”家庭', '', 1, 1, 0, '', 0, 0, 0, '2022-06-26 23:31:44', '2022-06-29 16:10:42', NULL);
INSERT INTO `t_mp_message_template` VALUES (1174916000357908480, 'SceneIntelligenceNotice', '智能场景通知任务（不删）', '“{{.title}}”任务执行中', '', 1, 4, 0, '', 0, 0, 0, '2022-06-10 10:22:36', '2022-06-28 08:36:33', NULL);
INSERT INTO `t_mp_message_template` VALUES (1241164354670395392, 'DeviceOffline', '设备离线消息（不删）', '设备已离线', '', 1, 3, 0, '', 0, 0, 0, '2022-08-12 10:48:42', '2022-08-12 10:50:56', NULL);
INSERT INTO `t_mp_message_template` VALUES (2241111354270395322, 'DeviceOnline', '设备在线消息（不删）', '设备已上线', '', 1, 3, 0, '', 0, 0, 0, '2022-08-12 10:48:42', '2023-04-13 14:56:31', NULL);
INSERT INTO `t_mp_message_template` VALUES (3468366293747269632, 'SendAddSharedMessage', '共享者发送邀请给接受者(不删)', '共享者“{{.userName}}”发送了设备共享邀请', '', 1, 3, 0, '', 0, 0, 0, '2022-11-01 16:55:27', '2022-11-01 16:55:27', NULL);
INSERT INTO `t_mp_message_template` VALUES (3538279681415348224, 'SendCancelSharedMessage', '共享者取消共享给接受者(不删)', '共享者“{{.userName}}”取消设备共享', '', 1, 3, 0, '', 0, 0, 0, '2022-11-01 16:55:44', '2022-11-01 16:55:44', NULL);
INSERT INTO `t_mp_message_template` VALUES (3606349707393007616, 'SendRefuseReceiveSharedMessage', '接受者取消接受设备，拒绝接受设备(不删)', '接受者“{{.userName}}”取消接受设备共享', '', 1, 3, 0, '', 0, 0, 0, '2022-11-01 16:56:00', '2022-11-01 16:56:00', NULL);
INSERT INTO `t_mp_message_template` VALUES (3687383696198762496, 'SendReceiveSharedMessage', '接受者已接受设备(不删)', '接受者“{{.userName}}”已接受设备共享邀请', '', 1, 3, 0, '', 0, 0, 0, '2022-11-01 16:56:19', '2022-11-01 16:56:19', NULL);
INSERT INTO `t_mp_message_template` VALUES (4893604863314853888, 'UpdatePasswordMessage', '修改密码（不删）', '账号“{{.account}}”修改密码成功', '', 1, 2, 0, '', 0, 0, 0, '2022-06-25 16:38:04', '2022-06-26 23:30:49', NULL);
INSERT INTO `t_mp_message_template` VALUES (6989277402417430528, 'FeedbackReplyNotice', '反馈状态消息内容（不删）', '您提交的问题反馈有收到回复，请查看。', '', 1, 2, 0, '', 0, 0, 0, '2022-08-12 10:25:52', '2022-08-12 10:26:58', NULL);
INSERT INTO `t_mp_message_template` VALUES (7003514157047644160, 'DeviceActiveMessage', '设备添加（不删）', '“{{.userName}}”在家庭“{{.homeName}}”中添加的“{{.deviceName}}”可以使用了', '', 1, 1, 0, '', 0, 0, 0, '2022-06-25 16:12:02', '2022-07-25 10:24:50', NULL);

-- ----------------------------
-- Table structure for t_mp_message_user_in
-- ----------------------------
DROP TABLE IF EXISTS `t_mp_message_user_in`;
CREATE TABLE `t_mp_message_user_in`  (
                                       `id` bigint(20) NOT NULL COMMENT '主键ID',
                                       `lang` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语]',
                                       `message_type` int(11) NULL DEFAULT NULL COMMENT '消息类型[1：系统提醒，2：设备提醒，3：家庭消息....自定义]',
                                       `child_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '消息扩展类型',
                                       `action_type` int(11) NULL DEFAULT NULL COMMENT '行为类型[0：主动发送，1：被动发送]',
                                       `push_mode` int(11) NULL DEFAULT NULL COMMENT '发送模式[0:实时发送， 1：定时发送，2：轮询发送]',
                                       `tpl_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '模板编码',
                                       `message_id` bigint(20) NULL DEFAULT NULL COMMENT '消息ID(mp_message.id)',
                                       `is_public` smallint(1) NULL DEFAULT NULL COMMENT '是否公共消息',
                                       `did` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '设备唯一DID',
                                       `product_key` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '产品key',
                                       `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户ID(t_cloud_uc_user.id)',
                                       `home_id` bigint(20) NULL DEFAULT NULL COMMENT '家庭编号',
                                       `msg_tag` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标签类型 0：alias 、1：tag、 2：regid',
                                       `icon_url` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图标',
                                       `push_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
                                       `push_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '发送内容',
                                       `read_flag` int(11) NULL DEFAULT NULL COMMENT '读取标识[0:未读， 1：已读]',
                                       `un_set_expire` smallint(1) NULL DEFAULT NULL COMMENT '不设置超时 1',
                                       `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
                                       `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                       `source_table` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '来源的表',
                                       `source_row_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '来源的行id',
                                       `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '外链地址',
                                       `app_key` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'APP Key',
                                       `tenant_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '开发者租户编号',
                                       `params` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '推送参数',
                                       PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_mp_message_user_in
-- ----------------------------

-- ----------------------------
-- Table structure for t_mp_message_user_out
-- ----------------------------
DROP TABLE IF EXISTS `t_mp_message_user_out`;
CREATE TABLE `t_mp_message_user_out`  (
                                        `id` bigint(20) NOT NULL COMMENT '主键ID',
                                        `lang` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语]',
                                        `message_type` int(11) NOT NULL COMMENT '消息类型[0：系统提醒，1：设备提醒，....自定义]',
                                        `action_type` int(11) NOT NULL COMMENT '行为类型[0：主动发送，1：被动发送]',
                                        `push_mode` int(11) NOT NULL COMMENT '发送模式[0:实时发送， 1：定时发送，2：轮询发送]',
                                        `tpl_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '模板编码',
                                        `message_id` bigint(20) NULL DEFAULT NULL COMMENT '消息ID(t_cloud_message.id)',
                                        `did` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '设备唯一DID',
                                        `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户ID(t_cloud_uc_user.id)',
                                        `link_type` int(11) NULL DEFAULT NULL COMMENT '跳转类型',
                                        `link_function_code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '跳转功能编号',
                                        `link_url` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '跳转链接',
                                        `dynamic` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '透传业务参数',
                                        `push_title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
                                        `push_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '发送内容',
                                        `external_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '外部消息ID(极光ID)',
                                        `external_status` int(11) NULL DEFAULT NULL COMMENT '外部发送状态',
                                        `external_error_message` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '外部失败原因描述',
                                        `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
                                        `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
                                        `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                        `updated_at` timestamp NULL DEFAULT NULL COMMENT '修改时间',
                                        `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_mp_message_user_out
-- ----------------------------

-- ----------------------------
-- Table structure for t_ms_notice_template
-- ----------------------------
DROP TABLE IF EXISTS `t_ms_notice_template`;
CREATE TABLE `t_ms_notice_template`  (
                                       `id` bigint(20) NOT NULL COMMENT '主键ID',
                                       `tpl_subject` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '通知主题',
                                       `tpl_code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '模板编码',
                                       `tpl_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '模板名称',
                                       `tpl_content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '模板内容',
                                       `tpl_params` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '模板参数',
                                       `thirdpary_code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '第三方模板编码',
                                       `sms_supplier` int(2) NULL DEFAULT NULL COMMENT '短信服务供应商（短信服务提供商, 默认0）',
                                       `method` int(11) NOT NULL COMMENT '通知方式[0:邮件，1：短信，2：MQ]',
                                       `lang` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语]',
                                       `tpl_type` int(11) NOT NULL DEFAULT 0 COMMENT '1:验证码;2:注册结果通知;',
                                       `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
                                       `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
                                       `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                       `updated_at` timestamp NULL DEFAULT NULL COMMENT '修改时间',
                                       `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                       PRIMARY KEY (`id`) USING BTREE,
                                       UNIQUE INDEX `uniq_notice_template_lang_type_method`(`sms_supplier`, `method`, `lang`, `tpl_type`, `deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_ms_notice_template
-- ----------------------------
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (7534941256053522432, '某某科技', 'EMAILCODE', '忘记密码验证码', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n        <p>您的验证码为{{.Code}}，有效时间10分钟。</p>\n        <p>您正在使用本邮箱进行验证，如非本人操作，请忽略。</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', '', 'EMAILCODE', 1, 2, 'zh', 2, 1, 1, '2022-12-14 14:04:17', '2022-12-14 14:13:06', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (6327450088234909696, 'Moumou', 'SMS_12345678', '绑定手机或邮箱验证码短信_en', 'The verification code for binding mobile phone/email is {{.Code}}, valid for 10 minutes.', '', 'SMS_261025110', 1, 3, 'en', 6, 1, 1, '2022-11-14 17:01:31', '2022-11-29 11:37:43', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5727028904548270080, 'MOUMOU Cloud', 'EMAIL_CODE', '绑定邮箱验证码', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n          <p>Hi</p>\n          <p>Your verification code is {{.Code}}, valid for 10 minutes.</p>\n          <p>You are using this phone to verify, please ignore if you are not doing it yourself.</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', '', 'EMAIL_CODE', 1, 2, 'en', 6, 1, 1, '2022-12-14 16:10:17', '2023-03-21 10:32:44', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5720208345542000640, '某某科技', 'SMS_12345678', '注册验证码', '您的注册验证码为${code}，有效时间10分钟。您正在使用本手机进行验证，如非本人操作，请忽略。', 'code', 'SMS_257831396', 1, 1, 'zh', 1, 0, 1, '2022-04-21 13:04:01', '2022-11-28 16:53:18', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5526625417563308032, '某某科技', 'EMAILCODE', '注册验证码', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n        <p>感谢注册！</p>\n        <p>您的注册验证码为{{.Code}}，有效时间10分钟。</p>\n        <p>您正在使用本邮箱进行验证，如非本人操作，请忽略。</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', 'code;userName', 'EMAILCODE', 1, 2, 'zh', 1, 0, 1, '2022-04-21 13:08:14', '2022-12-14 14:43:08', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5521253235974635520, 'MOUMOU Cloud', 'EMAIL_CODE', '绑定第三方账号验证码_en', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n          <p>Hi</p>\n          <p>Your verification code is {{.Code}}, valid for 10 minutes.</p>\n          <p>You are using this phone to verify, please ignore if you are not doing it yourself.</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', '', 'EMAIL_CODE', 1, 2, 'en', 5, 1, 0, '2022-12-14 16:09:28', '2022-12-14 16:09:28', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5511025109038628864, '某某科技', 'EMAILCODE', '注册验证码', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n        <p>感谢！</p>\n        <p>您的验证码为{{.Code}}，有效时间10分钟。</p>\n        <p>您正在使用本邮箱进行验证，如非本人操作，请忽略。</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', 'code;userName', 'EMAILCODE', 1, 2, 'zh', 7, 0, 1, '2022-04-21 13:08:14', '2023-02-17 11:54:55', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5296657197420347392, '某某科技', 'SMS_12345678', '注册验证码短信', '您的注册验证码为{{.Code}}，有效时间10分钟。如非本人操作，请忽略。', '', 'SMS_258660017', 1, 3, 'zh', 1, 1, 1, '2022-11-28 17:03:20', '2022-11-29 11:37:35', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5245044940513116160, 'MOUMOU Cloud', 'EMAIL_CODE', '注销账号验证码_en', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n          <p>Unfortunately, we can no longer provide you with services and experiences.</p>\n          <p>Your verification code is {{.Code}}, valid for 10 minutes.</p>\n          <p>You are using this phone to verify, please ignore if you are not doing it yourself.</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', '', 'EMAIL_CODE', 1, 2, 'en', 4, 1, 0, '2022-12-14 16:08:22', '2022-12-14 16:08:22', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5207608909153533952, '某某科技', 'SMS_12345678', '绑定手机或邮箱验证码短信', '您的验证码为${code}，有效时间10分钟。您正在使用本手机进行绑定手机/邮箱的服务，如非本人操作，请忽略。', '', 'SMS_257751454', 1, 1, 'zh', 6, 1, 1, '2022-11-14 17:05:58', '2022-11-28 16:52:07', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5040546389180907520, '某某科技', 'EMAIL_CODE', '修改密码验证码', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n        <p>您的验证码为{{.Code}}，有效时间10分钟。</p>\n        <p>您正在使用本邮箱进行验证，如非本人操作，请忽略。</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', '', 'EMAIL_CODE', 1, 2, 'zh', 3, 1, 0, '2022-12-14 14:14:12', '2022-12-14 14:14:12', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5017588199329792000, 'MOUMOU Cloud', 'EMAIL_CODE', '修改密码验证码_en', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n          <p>Hi</p>\n          <p>Your verification code is {{.Code}}, valid for 10 minutes.</p>\n          <p>You are using this phone to verify, please ignore if you are not doing it yourself.</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', '', 'EMAIL_CODE', 1, 2, 'en', 3, 1, 0, '2022-12-14 16:07:28', '2022-12-14 16:07:28', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5017326905083396096, '某某科技', 'SMS_12345678', '忘记密码验证码短信', '您的重置密码的验证码为{{.Code}}，有效时间10分钟。如非本人操作，请尽快修改密码。', '', 'SMS_258815013', 1, 3, 'zh', 2, 1, 1, '2022-11-28 17:04:27', '2022-11-29 11:37:29', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4767794427764768768, '某某科技', 'SMS_12345678', '修改密码验证码短信', '您的重置密码的验证码为{{.Code}}，有效时间10分钟。如非本人操作，请尽快修改密码。', '', 'SMS_258815013', 1, 3, 'zh', 3, 1, 1, '2022-11-28 17:05:26', '2022-11-29 11:37:20', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4696897020076392448, 'MOUMOU Cloud', 'EMAIL_CODE', '忘记密码验证码_en', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n          <p>Hi</p>\n          <p>Your verification code is {{.Code}}, valid for 10 minutes.</p>\n          <p>You are using this phone to verify, please ignore if you are not doing it yourself.</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', '', 'EMAIL_CODE', 1, 2, 'en', 2, 1, 0, '2022-12-14 16:06:11', '2022-12-14 16:06:11', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4641886828910182400, '某某科技', 'EMAIL_CODE', '注销账号验证码', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n        <p>很抱歉，未能给你带来更好的体验！</p>\n        <p>您的验证码为{{.Code}}，有效时间10分钟。</p>\n        <p>您正在使用本邮箱进行验证，如非本人操作，请忽略。</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', '', 'EMAIL_CODE', 1, 2, 'zh', 4, 1, 0, '2022-12-14 14:15:47', '2022-12-14 14:15:47', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4534313866414358528, '某某科技', 'SMS_12345678', '注销账号验证码短信', '您的注销账号的验证码为{{.Code}}，有效时间10分钟。如非本人操作，请忽略。', '', 'SMS_258630017', 1, 3, 'zh', 4, 1, 1, '2022-11-28 17:06:22', '2022-11-29 11:37:14', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4475073133874675712, 'Moumou', 'SMS_12345678', '绑定第三方账号验证码短信_en', '\nThe verification code of the bound account is {{.Code}}, valid for 10 minutes.', '', 'SMS_260940131', 1, 3, 'en', 5, 1, 1, '2022-11-14 15:17:51', '2022-11-29 11:37:53', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4404793502181785600, 'MOUMOU Cloud', 'EMAIL_CODE', '注册验证码_en', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n          <p>Thanks for registering！</p>\n          <p>Your registration verification code is {{.Code}}, valid for 10 minutes.</p>\n          <p>You are using this phone to verify, please ignore if you are not doing it yourself.</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', '', 'EMAIL_CODE', 1, 2, 'en', 1, 1, 0, '2022-12-14 16:05:02', '2022-12-14 16:05:02', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4376046017510998016, '某某科技', 'SMS_12345678', '绑定第三方账号验证码短信', '您的绑定账号的验证码为{{.Code}}，有效时间10分钟。如非本人操作，请忽略。', '', 'SMS_258745018', 1, 3, 'zh', 5, 1, 1, '2022-11-28 17:06:59', '2022-11-29 11:37:07', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4277607634409783296, 'Moumou', 'SMS_12345678', '注销账号验证码短信_en', 'The verification code for canceling the account is {{.Code}}, valid for 10 minutes.', '', 'SMS_261105070', 1, 3, 'en', 4, 1, 1, '2022-11-14 15:17:04', '2022-11-29 11:38:01', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4251553458605359104, '某某科技', 'EMAIL_CODE', '绑定第三方账号验证码', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n        <p>感谢注册！</p>\n        <p>您的验证码为{{.Code}}，有效时间10分钟。</p>\n        <p>您正在使用本邮箱进行验证，如非本人操作，请忽略。</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', '', 'EMAIL_CODE', 1, 2, 'zh', 5, 1, 0, '2022-12-14 14:17:20', '2022-12-14 14:17:20', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4121467601387683840, 'Moumou', 'SMS_12345678', '修改密码验证码短信_en', 'The verification code for resetting the password is {{.Code}}, valid for 10 minutes.', '', 'SMS_261165087', 1, 3, 'en', 3, 1, 1, '2022-11-14 15:16:27', '2022-11-29 11:38:09', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3942397983244320768, 'Moumou', 'SMS_12345678', '忘记密码验证码短信_en', 'The verification code for resetting the password is {{.Code}}, valid for 10 minutes.', '', 'SMS_261165087', 1, 3, 'en', 2, 1, 1, '2022-11-14 15:15:44', '2022-11-29 11:38:19', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3803131127003185152, '某某科技', 'EMAIL_CODE', '绑定邮箱验证码', '<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n</head>\n<body>\n  <div class=\"container\">\n    <div class=\"content\">\n      <div class=\"verification\">\n        <p>您的验证码为{{.Code}}，有效时间10分钟。</p>\n        <p>您正在使用本邮箱进行验证，如非本人操作，请忽略。</p>\n      </div>\n    </div>\n  </div>\n</body>\n</html>', '', 'EMAIL_CODE', 1, 2, 'zh', 6, 1, 0, '2022-12-14 14:19:07', '2022-12-14 14:19:07', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3799921762402140160, '某某科技', 'SMS_12345678', '绑定手机或邮箱验证码短信', '您的绑定手机/邮箱的验证码为{{.Code}}，有效时间10分钟。如非本人操作，请忽略。', '', 'SMS_258705017', 1, 3, 'zh', 6, 1, 1, '2022-11-28 17:09:17', '2022-11-29 11:37:01', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3747707393828356096, 'Moumou', 'SMS_12345678', '注册验证码短信_en', 'The registration verification code is {{.Code}}, valid for 10 minutes.', '', 'SMS_260950122', 1, 3, 'en', 1, 1, 1, '2022-11-14 15:14:57', '2022-11-29 11:38:25', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3489607401969975296, '某某科技', 'SMS_12345678', '绑定第三方账号验证码短信', '您的验证码为${code}，有效时间10分钟。您正在使用本手机进行绑定账号的服务，如非本人操作，请忽略。', '', 'SMS_257751453', 1, 1, 'zh', 5, 1, 1, '2022-11-14 15:13:56', '2022-11-28 16:52:51', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3354134749097394176, '某某科技', 'SMS_12345678', '注销账号验证码短信', '您的验证码为${code}，有效时间10分钟。您正在使用本手机进行注销账号的服务，如非本人操作，请忽略。', '', 'SMS_257851443', 1, 1, 'zh', 4, 1, 1, '2022-11-14 15:13:24', '2022-11-28 16:52:57', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3134029889368129536, '某某科技', 'SMS_12345678', '修改密码验证码短信', '您的验证码为${code}，有效时间10分钟。您正在使用本手机进行重置密码的服务，如非本人操作，请尽快修改密码。', '', 'SMS_257706413', 1, 1, 'zh', 3, 1, 1, '2022-11-14 15:12:31', '2022-11-28 16:53:05', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2956359534144028672, '某某科技', 'SMS_12345678', '注册验证码', '您的注册验证码为${code}，有效时间10分钟。您正在使用本手机进行验证，如非本人操作，请忽略。', 'code', 'SMS_257831396', 1, 1, 'zh', 7, 0, 1, '2022-04-21 13:04:01', '2023-07-25 16:06:55', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2914112916292534272, '某某科技', 'SMS_12345678', '忘记密码验证码短信）', '您的验证码为${code}，有效时间10分钟。您正在使用本手机进行重置密码的服务，如非本人操作，请尽快修改密码。', '', 'SMS_257706413', 1, 1, 'zh', 2, 1, 1, '2022-11-14 15:11:39', '2022-11-28 16:53:12', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2392966253582057472, 'Moumou', 'SMS_12345678', '注册验证码短信_en', 'The registration verification code is ${code}, valid for 10 minutes.', '', 'SMS_260950122', 1, 1, 'en', 1, 1, 0, '2022-11-28 17:14:52', '2022-11-28 17:14:52', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2080348402658410496, 'Moumou', 'SMS_12345678', '忘记密码验证码短信_en', 'The verification code for resetting the password is ${code}, valid for 10 minutes.', '', 'SMS_261165087', 1, 1, 'en', 2, 1, 0, '2022-11-28 17:16:07', '2022-11-28 17:16:07', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1871003265550876672, 'Moumou', 'SMS_12345678', '修改密码验证码短信_en', 'The verification code for resetting the password is ${code}, valid for 10 minutes.', '', 'SMS_261165087', 1, 1, 'en', 3, 1, 0, '2022-11-28 17:16:57', '2022-11-28 17:16:57', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1693035753333948416, 'Moumou', 'SMS_12345678', '注销账号验证码短信_en', 'The verification code for canceling the account is ${code}, valid for 10 minutes.', '', 'SMS_261105070', 1, 1, 'en', 4, 1, 0, '2022-11-28 17:17:39', '2022-11-28 17:17:39', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1549792638289412096, 'Moumou', 'SMS_12345678', '绑定第三方账号验证码短信_en', 'The verification code of the bound account is ${code}, valid for 10 minutes.', '', 'SMS_260940131', 1, 1, 'en', 5, 1, 0, '2022-11-28 17:18:13', '2022-11-28 17:18:13', NULL);
INSERT INTO `iot_message`.`t_ms_notice_template` (`id`, `tpl_subject`, `tpl_code`, `tpl_name`, `tpl_content`, `tpl_params`, `thirdpary_code`, `sms_supplier`, `method`, `lang`, `tpl_type`, `created_by`, `updated_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1406869786020904960, 'Moumou', 'SMS_12345678', '绑定手机或邮箱验证码短信_en', 'The verification code for binding mobile phone/email is ${code}, valid for 10 minutes.', '', 'SMS_261025110', 1, 1, 'en', 6, 1, 0, '2022-11-28 17:18:47', '2022-11-28 17:18:47', NULL);

SET FOREIGN_KEY_CHECKS = 1;
