SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- 创建数据库
DROP database IF EXISTS `iot_open_system`;
create database iot_open_system;
use iot_open_system;

-- ----------------------------
-- Table structure for t_open_auth_quantity
-- ----------------------------
DROP TABLE IF EXISTS `t_open_auth_quantity`;
CREATE TABLE `t_open_auth_quantity`  (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `user_id` bigint(20) NULL DEFAULT NULL COMMENT '主用户ID',
  `company_id` bigint(20) NOT NULL COMMENT '公司ID',
  `tenant_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '租户ID',
  `auth_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '授权码,自动生成',
  `auth_quantity` int(11) NULL DEFAULT 50 COMMENT '授权数量',
  `auth_date` timestamp NULL DEFAULT NULL COMMENT '授权日期',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态;1: 正常; 2:禁用;',
  `created_by` bigint(20) NULL DEFAULT NULL COMMENT '创建者',
  `created_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_by` bigint(20) NULL DEFAULT NULL COMMENT '更新者',
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间,有值表示已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_open_auth_quantity
-- ----------------------------

-- ----------------------------
-- Table structure for t_open_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `t_open_auth_rule`;
CREATE TABLE `t_open_auth_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '条件',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `menu_type` int(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型 0目录 1菜单 2按钮',
  `weigh` int(10) NOT NULL DEFAULT 0 COMMENT '权重',
  `status` int(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态',
  `always_show` int(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '显示状态',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `jump_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '跳转路由',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `is_frame` int(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否外链 1是 0否',
  `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属模块',
  `model_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模型ID',
  `is_cache` int(20) NULL DEFAULT 1 COMMENT '是否缓存(1 缓存, 2 不缓存)',
  `is_hide_child_menu` int(20) NULL DEFAULT 1 COMMENT '是否隐藏子菜单(1 隐藏, 2 不隐藏)',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除日期',
  `created_by` bigint(20) NULL DEFAULT NULL,
  `updated_by` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `weigh`(`weigh`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 9159164745528279041 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单节点表' ROW_FORMAT = Dynamic;


INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (9159164745528279040, 0, 'Product', '产品开发', 'icon-chanpinkaifa', '', '', 0, 1, 1, 1, '/product', '/product/product/index', 'PageView', 2, 'sys_admin', 0, 1, 2, '2022-05-12 15:35:25', '2024-01-10 17:26:36', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (8885045495148412928, 0, 'DataManage', '数据管理', 'icon-shujuguanli', '', '', 0, 3, 1, 1, '/dataManage', '/dataManage/overview/index', 'PageView', 2, 'sys_admin', 0, 1, 2, '2022-05-12 15:37:01', '2024-01-10 17:26:49', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (8737017658974765056, 0, 'Authority', '授权管理', 'icon-shouquanguanli', '', '', 0, 5, 1, 1, '/authority', '', 'PageView', 2, 'sys_admin', 0, 1, 2, '2022-05-12 15:37:36', '2024-01-10 17:26:56', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (8959096378445692928, 0, 'AppDevelop', 'APP开发', 'icon-APPkaifa', '', '', 0, 2, 1, 1, '/appDevelop', '/appDevelop/index', 'PageView', 2, 'sys_admin', 0, 1, 2, '2022-05-27 09:51:51', '2024-01-10 17:26:41', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (781674127982231552, 0, 'Marketing', '营销', 'icon-marketing', '', '', 0, 4, 1, 1, '/marketing', '/marketing/marketingService/helpCenter/index.vue', 'PageView', 2, 'sys_admin', 0, 1, 2, '2022-07-07 13:39:41', '2024-01-10 17:26:52', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (6890578711163273216, 2, 'guize22111', '', 'time', '', '', 0, 0, 1, 2, '/system/test11133', '/system/test22233', '', 0, 'sys_admin', 0, 1, 1, '2022-05-11 20:12:08', '2022-05-11 20:12:08', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (242, 238, 'Dictionaries', '字典管理', 'file-text', '', '', 1, 2, 1, 2, '/system/dictionaries/index', '', '/system/dictionaries/index', 0, 'sys_admin', 0, 1, 1, '2022-04-14 17:19:01', '2022-04-18 17:39:30', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (241, 238, 'Account', '个人信息', 'solution', '', '', 1, 1, 1, 2, '/system/account/index', '', '/system/account/index', 2, 'sys_admin', 0, 1, 1, '2022-04-14 17:18:07', '2022-04-21 19:26:30', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (255, 239, 'Role', '角色管理', 'team', '', '', 1, 6, 1, 2, '/authority/role/index', '', '/authority/role/index', 0, 'sys_admin', 0, 1, 1, '2022-04-15 10:59:12', '2022-04-18 17:38:41', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (256, 239, 'Department', '部门管理', 'cluster', '', '', 1, 5, 1, 2, '/authority/department/index', '', '/authority/department/index', 0, 'sys_admin', 0, 1, 1, '2022-04-15 10:59:57', '2022-04-18 17:39:12', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (257, 239, 'Post', '岗位管理', 'notification', '', '', 1, 8, 1, 2, '/authority/post/index', '', '/authority/post/index', 0, 'sys_admin', 0, 1, 1, '2022-04-15 11:01:00', '2022-04-18 17:38:37', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (258, 239, 'Developer', '开发者认证', 'key', '', '', 1, 9, 1, 2, '/authority/developer/index', '', '/authority/developer/index', 2, 'sys_admin', 0, 1, 1, '2022-04-15 11:01:33', '2022-04-25 14:10:47', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (244, 239, 'User', '用户管理', 'user', '', '', 1, 2, 1, 2, '/authority/user/index', '', '/authority/user/index', 2, 'sys_admin', 0, 1, 1, '2022-04-14 17:21:12', '2022-04-21 20:08:18', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (243, 239, 'CloudMenu', '云管平台菜单', 'bars', '', '', 1, 1, 1, 2, '/authority/cloud-menu/index', '', '/authority/cloudMenu/index', 0, 'sys_admin', 0, 1, 1, '2022-04-14 17:20:28', '2022-04-24 16:25:22', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (2403733430450880512, 239, 'xxx', 'testhu', 'right', '', '', 1, 333, 1, 2, 'xxx', '', 'xxx', 2, 'sys_admin', 0, 2, 2, '2022-04-25 14:47:01', '2022-04-25 14:52:50', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (251, 240, 'Device', '设备管理', 'tablet', '', '', 1, 6, 1, 2, '/product/device', '/product/device/index', 'BlankLayout', 0, 'sys_admin', 0, 2, 1, '2022-04-15 10:55:33', '2022-04-25 14:18:43', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (249, 240, 'Firmware', '固件管理', 'up-square', '', '', 1, 4, 1, 2, '/product/firmware', '/product/firmware/index', 'BlankLayout', 0, 'sys_admin', 0, 1, 1, '2022-04-15 10:54:25', '2022-04-25 14:20:23', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (248, 240, 'Classification', '产品分类管理', 'appstore', '', '', 1, 3, 1, 2, '/product/classification/index', '', '/product/classification/index', 0, 'sys_admin', 0, 1, 1, '2022-04-15 10:53:52', '2022-04-18 17:37:36', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (246, 240, 'StandardProduct', '产品类别', 'rocket', '', '', 1, 2, 1, 2, '/product/index', '/product/product/index', 'BlankLayout', 0, 'sys_admin', 0, 2, 1, '2022-04-15 09:48:57', '2022-04-25 13:43:13', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (245, 240, 'Chip', '芯片模组管理', 'deployment-unit', '', '', 1, 1, 1, 2, '/product/chip/index', '', '/product/chip/index', 0, 'sys_admin', 0, 1, 1, '2022-04-15 09:46:17', '2022-04-18 17:38:11', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1535601662729551872, 246, 'ProductDetails', '产品详情', 'dash', '', '', 1, 1, 1, 2, '/product/product/details/index', '', '/product/product/details/index', 2, 'sys_admin', 0, 2, 2, '2022-04-25 11:22:46', '2022-04-25 11:47:13', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (2084296393052028928, 246, 'ProductList', '产品列表', 'dash', '', '', 1, 1, 1, 2, '/product/product/index', '', '/product/product/index', 2, 'sys_admin', 0, 2, 2, '2022-04-25 11:24:57', '2022-04-25 11:26:16', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (4115308261576441856, 249, 'FirmwareList', '固件列表', 'dash', '', '', 1, 1, 1, 2, '/product/firmware/index', '', '/product/firmware/index', 2, 'sys_admin', 0, 2, 2, '2022-04-25 13:59:37', '2022-04-25 14:20:30', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (4897173437564223488, 249, 'FirmwareDetails', '固件详情', 'dash', '', '', 1, 2, 1, 2, '/product/firmware/details/index', '', '/product/firmware/details/index', 2, 'sys_admin', 0, 2, 2, '2022-04-25 14:02:43', '2022-04-25 14:02:43', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (5048246558223400960, 251, 'DeviceLog', '设备日志', 'dash', '', '', 1, 2, 1, 2, '/product/device/log/index', '', '/product/device/log/index', 2, 'sys_admin', 0, 2, 2, '2022-04-25 11:36:43', '2022-04-25 11:43:35', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (5209425033708863488, 251, 'DeviceDetails', '菜单详情', 'dash', '', '', 1, 1, 1, 2, '/product/device/details/index', '', '/product/device/details/index', 2, 'sys_admin', 0, 2, 2, '2022-04-25 11:37:22', '2022-04-25 11:43:05', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (7085924647471841280, 251, 'DeviceList', '设备列表', 'dash', '', '', 1, 1, 1, 2, '/product/device/index', '', '/product/device/index', 2, 'sys_admin', 0, 2, 2, '2022-04-25 11:44:49', '2022-04-25 14:19:23', NULL, NULL, NULL);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1178813131179065344, 781674127982231552, 'MarketingService', '营销服务', 'book', '', '', 1, 1, 1, 1, '/marketing/marketingService', '/marketing/marketingService/helpCenter/index.vue', 'PageView', 2, 'sys_admin', 0, 2, 2, '2022-07-07 13:41:16', '2023-03-27 11:49:55', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1831562684512763904, 781674127982231552, 'AfterSalesService', '售后服务', 'file-unknown', '', '', 1, 2, 1, 1, '/marketing/afterSalesService', '', 'PageView', 2, 'sys_admin', 0, 2, 2, '2022-07-07 13:43:51', '2022-07-07 13:46:42', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1576791311216705536, 1178813131179065344, 'HelpCenter', '帮助中心', 'file', '', '', 1, 1, 1, 1, '/marketing/marketingService/helpCenter/index.vue', '', '/marketing/marketingService/helpCenter/index.vue', 2, 'sys_admin', 0, 1, 2, '2022-07-07 13:42:51', '2024-01-10 17:28:31', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (4593347916386435072, 1178813131179065344, 'helpCenterDocManage', '帮助文档管理', 'hdd', '', '', 1, 2, 1, 2, '/marketing/marketingService/helpCenter/docManage/index', '', '/marketing/marketingService/helpCenter/docManage/index', 2, 'sys_admin', 0, 2, 2, '2022-07-11 09:32:23', '2022-07-11 09:32:23', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (2977052235772362752, 1178813131179065344, 'productDocManage', '产品帮助文档管理', 'copy', '', '', 1, 4, 1, 2, '/marketing/marketingService/productHelpCenter/docManage/index', '', '/marketing/marketingService/productHelpCenter/docManage/index', 2, 'sys_admin', 0, 2, 2, '2022-08-16 16:29:25', '2022-08-16 16:29:25', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1342078855191560192, 1178813131179065344, 'ProductHelpCenter', '产品帮助文档', 'copy', '', '', 1, 2, 1, 1, '/marketing/marketingService/productHelpCenter/index.vue', '', '/marketing/marketingService/productHelpCenter/index.vue', 2, 'sys_admin', 0, 1, 2, '2022-08-16 13:45:39', '2024-03-07 08:37:14', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (758281464944099328, 1178813131179065344, 'AppSplashScreenPushDetails', 'APP闪屏推送详情', 'credit-card', '', '', 1, 11, 1, 2, '/marketing/appSplashScreenPush/details/index', '', '/marketing/appSplashScreenPush/details/index', 2, 'sys_admin', 0, 2, 2, '2022-09-30 11:19:15', '2024-02-28 17:55:46', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (6214091536351723520, 1178813131179065344, 'AppSplashScreenPush', 'APP闪屏推送', 'code', '', '', 1, 3, 1, 1, '/marketing/appSplashScreenPush/index.vue', '', '/marketing/appSplashScreenPush/index.vue', 2, 'sys_admin', 0, 1, 2, '2022-09-30 10:57:35', '2024-03-07 14:20:27', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (2162740003589750784, 1831562684512763904, 'UserFeedback', 'APP用户反馈', 'file-unknown', '', '', 1, 1, 1, 1, '/marketing/afterSalesService/userFeedback/index.vue', '', '/marketing/afterSalesService/userFeedback/index.vue', 2, 'sys_admin', 0, 1, 2, '2022-07-07 13:45:10', '2024-03-07 08:37:41', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (2015813318220021760, 1831562684512763904, 'FeedbackDetails', '反馈详情', 'file-text', '', '', 1, 2, 1, 2, '/marketing/afterSalesService/userFeedback/details/index', '', '/marketing/afterSalesService/userFeedback/details/index', 2, 'sys_admin', 0, 2, 2, '2022-07-11 09:06:07', '2022-07-11 09:06:07', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (431674596966629376, 1831562684512763904, 'ProblemType', '问题类型管理', 'copy', '', '', 1, 2, 1, 1, '/marketing/afterSalesService/problemType/index.vue', '', '/marketing/afterSalesService/problemType/index.vue', 2, 'sys_admin', 0, 2, 2, '2022-10-20 10:14:27', '2022-10-20 10:14:27', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (2469472581170659328, 1891977961314091008, 'EmbedPanel', '应用面板设计', 'credit-card', '', '', 1, 2, 1, 2, '/panel/embed/screenDesign/index.vue', '', '/panel/embed/screenDesign/index.vue', 2, 'sys_admin', 0, 2, 2, '2023-06-26 16:59:53', '2023-08-10 16:06:50', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (2271806318794670080, 1891977961314091008, 'EmbedPanelList', '应用面板', 'credit-card', '', '', 1, 1, 1, 1, '/panel/embed/index.vue', '', '/panel/embed/index.vue', 2, 'sys_admin', 0, 2, 2, '2023-06-26 16:59:06', '2023-08-10 16:06:46', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (2928553496183668736, 3115948585393422336, 'firmwareOTA', '固件OTA', 'desktop', '', '', 1, 10, 1, 1, '/product/firmwareOTA/index', '', '/product/firmwareOTA/index', 2, 'sys_admin', 0, 1, 2, '2022-05-25 14:19:05', '2024-03-07 08:34:23', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (3317268708988125184, 3115948585393422336, 'deviceDebugging', '设备调试', 'tool', '', '', 1, 11, 1, 1, '/product/deviceDebugging/index', '', '/product/deviceDebugging/index', 2, 'sys_admin', 0, 1, 2, '2022-05-26 17:10:09', '2024-03-07 08:34:33', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (7740460241187143680, 3115948585393422336, 'Device', '设备管理', 'form', '', '', 1, 5, 1, 1, '/product/device/index', '', '/product/device/index', 2, 'sys_admin', 0, 1, 2, '2022-06-23 09:16:25', '2024-03-07 08:33:46', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (8155109377754169344, 3115948585393422336, 'DeviceDetails', '设备详情', 'copy', '', '', 1, 6, 1, 2, '/product/device/details/index', '', '/product/device/details/index', 2, 'sys_admin', 0, 2, 2, '2022-06-23 09:18:04', '2022-06-23 09:18:40', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (8279675191588454400, 3115948585393422336, 'DeviceLog', '设备日志', 'calendar', '', '', 1, 7, 1, 2, '/product/device/log/index', '', '/product/device/log/index', 2, 'sys_admin', 0, 2, 2, '2022-06-23 09:18:34', '2022-06-23 09:18:34', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (8482041016558190592, 3115948585393422336, 'Firmware', '固件管理', 'cloud-download', '', '', 1, 8, 1, 1, '/product/firmware/index', '', '/product/firmware/index', 2, 'sys_admin', 0, 1, 2, '2022-06-23 09:19:22', '2024-03-07 08:34:07', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (8667428119637819392, 3115948585393422336, 'FirmwareDetails', '固件详情', 'code', '', '', 1, 9, 1, 2, '/product/firmware/details/index', '', '/product/firmware/details/index', 2, 'sys_admin', 0, 2, 2, '2022-06-23 09:20:06', '2022-06-23 09:20:06', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (3269957101346193408, 3115948585393422336, 'DebugDetails', '真实设备调试', 'cloud', '', '', 1, 12, 1, 2, '/product/deviceDebugging/debugDetails/index', '', '/product/deviceDebugging/debugDetails/index', 2, 'sys_admin', 0, 2, 2, '2022-06-27 09:03:26', '2022-06-27 09:03:26', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1754747229723066368, 3115948585393422336, 'releaseLog', '固件OTA发布记录', 'file-text', '', '', 1, 13, 1, 2, '/product/firmwareOTA/record/index', '', '/product/firmwareOTA/record/index', 2, 'sys_admin', 0, 2, 2, '2022-08-17 15:37:16', '2022-08-17 15:37:16', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1982825917403004928, 3115948585393422336, 'DeviceMalfunction', '设备故障', 'frown', '', '', 1, 14, 1, 2, '/product/device/malfunction/index', '', '/product/device/malfunction/index', 2, 'sys_admin', 0, 2, 2, '2023-02-20 14:51:56', '2023-02-20 14:51:56', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1024434789655937024, 3115948585393422336, 'Manufacture', '生产管理', 'copy', '', '', 1, 4, 1, 1, '/product/manufacture/index', '', '/product/manufacture/index', 2, 'sys_admin', 0, 2, 2, '2023-12-08 17:35:15', '2024-01-10 01:54:52', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (2422285703049019392, 3232128938264133632, 'StudioPanel', '自定义控制面板设计', 'icon-panel', '', '', 1, 2, 1, 2, '/panel/studioPanel/screenDesign/index.vue', '', '/panel/studioPanel/screenDesign/index.vue', 2, 'sys_admin', 0, 2, 2, '2023-06-27 17:25:43', '2023-08-10 16:06:36', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (3833491004791357440, 3232128938264133632, 'StudioPanelList', '自定义控制面板', 'icon-panel', '', '', 1, 1, 1, 1, '/panel/studioPanel/index.vue', '', '/panel/studioPanel/index.vue', 2, 'sys_admin', 0, 2, 2, '2023-06-26 17:05:18', '2023-08-10 16:06:23', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1752185267978076160, 4100852806750142464, 'OEMAppCofiguration', 'APP开发配置', 'form', '', '', 1, 4, 1, 2, '/appDevelop/OEMApp/configuration/index', '', '/appDevelop/OEMApp/configuration/index', 2, 'sys_admin', 0, 2, 2, '2022-05-27 10:22:35', '2023-03-15 14:23:19', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (5299407709798301696, 4100852806750142464, 'OptionalConfig', '可选配置', 'file', '', '', 1, 3, 1, 1, '/appDevelop/optionalConfig/index.vue', '', '/appDevelop/optionalConfig/index.vue', 2, 'sys_admin', 0, 2, 2, '2022-07-06 15:58:13', '2023-03-15 14:23:06', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (6960539848083931136, 4100852806750142464, 'CustomizedAppVersion', '自定义APP版本管理', 'align-center', '', '', 1, 6, 1, 2, '/appDevelop/customizedApp/versionManage/index', '', '/appDevelop/customizedApp/versionManage/index', 2, 'sys_admin', 0, 2, 2, '2023-03-03 15:04:41', '2023-03-03 15:04:41', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (8177095095929962496, 4100852806750142464, 'CustomizedApp', '自定义APP开发', 'bank', '', '', 1, 2, 1, 1, '/appDevelop/customizedApp/index', '', '/appDevelop/customizedApp/index', 2, 'sys_admin', 0, 2, 2, '2023-03-03 11:29:36', '2023-03-03 11:29:36', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (870977444129964032, 4100852806750142464, 'CustomizedAppConfig', '自定义APP配置', 'align-left', '', '', 1, 5, 1, 2, '/appDevelop/customizedApp/configuration/index', '', '/appDevelop/customizedApp/configuration/index', 2, 'sys_admin', 0, 2, 2, '2023-03-03 14:40:29', '2023-03-03 14:40:29', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (258944583298088960, 4100852806750142464, 'OEMAppDebug', 'OEM  APP调试管理', 'paper-clip', '', '', 1, 8, 1, 1, '/appDevelop/OEMAppDebug/index.vue', '', '/appDevelop/OEMAppDebug/index.vue', 2, 'sys_admin', 0, 2, 2, '2024-03-13 15:10:52', '2024-03-13 15:10:52', '2024-03-13 15:11:14', 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (4838154040086790144, 4367313833849421824, 'RecommendScene', '推荐场景', 'copy', '', '', 1, 1, 1, 1, '/appDevelop/intelligentScene/index.vue', '', '/appDevelop/intelligentScene/index.vue', 2, 'sys_admin', 0, 1, 2, '2022-11-08 10:16:35', '2024-03-07 08:35:53', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (5088011157059502080, 4367313833849421824, 'RecommendSceneDetails', '推荐场景详情', 'copy', '', '', 1, 2, 1, 2, '/appDevelop/intelligentScene/details/index', '', '/appDevelop/intelligentScene/details/index', 2, 'sys_admin', 0, 2, 2, '2022-11-08 10:50:27', '2022-11-08 10:50:27', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (3232128938264133632, 5150040366262419456, 'AppPanel', 'App', 'icon-panel', '', '', 1, 1, 1, 1, '/panel', '/panel/studioPanel/index.vue', 'PageView', 2, 'sys_admin', 0, 2, 2, '2023-06-26 17:02:55', '2023-08-10 16:05:55', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1891977961314091008, 5150040366262419456, 'embed', '屏', 'cloud', '', '', 1, 2, 1, 1, '/panel/embed', '/panel/embed/index.vue', 'PageView', 2, 'sys_admin', 0, 2, 2, '2023-06-26 16:57:35', '2023-08-10 16:06:07', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (6829802911032049664, 6609772312344297472, 'ProductList', '产品开发', 'icon-chanpinkaifa', '', '', 1, 1, 1, 1, '/product/product/index', '', '/product/product/index', 2, 'sys_admin', 0, 1, 2, '2022-06-23 09:12:48', '2024-03-07 08:31:43', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (7079616222605508608, 6609772312344297472, 'CreateProduct', '创建产品', 'form', '', '', 1, 2, 1, 2, '/product/product/createProduct/index', '', '/product/product/createProduct/index', 2, 'sys_admin', 0, 2, 2, '2022-06-23 09:13:48', '2022-06-23 09:14:03', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (7312533820899491840, 6609772312344297472, 'Setting', '配置功能', 'form', '', '', 1, 3, 1, 2, '/product/product/setting/index', '', '/product/product/setting/index', 2, 'sys_admin', 0, 2, 2, '2022-06-23 09:14:43', '2022-06-23 09:14:43', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (7466223372586614784, 6609772312344297472, 'NetworkInfo', '配网信息', 'form', '', '', 1, 4, 1, 2, '/product/product/setting/basicConfig/networkInfo/index', '', '/product/product/setting/basicConfig/networkInfo/index', 2, 'sys_admin', 0, 2, 2, '2022-06-23 09:15:20', '2022-06-23 09:15:20', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (873599101126475776, 6609772312344297472, 'ConfigScheme', '配置语音平台方案', 'file-word', '', '', 1, 7, 1, 2, '/product/product/setting/voiceControl/configScheme/index', '', '/product/product/setting/voiceControl/configScheme/index', 2, 'sys_admin', 0, 2, 2, '2022-08-31 10:38:17', '2022-08-31 14:05:16', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1338271198261379072, 6609772312344297472, 'VerifyGuide', '验证语音平台功能', 'smile', '', '', 1, 6, 1, 2, '/product/product/setting/voiceControl/verifyGuide/index', '', '/product/product/setting/voiceControl/verifyGuide/index', 2, 'sys_admin', 0, 2, 2, '2022-08-31 10:29:30', '2022-08-31 14:05:13', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (1630572854741401600, 6609772312344297472, 'ReleaseRecord', '查看发布记录', 'hdd', '', '', 1, 5, 1, 2, '/product/product/setting/voiceControl/releaseRecord/index', '', '/product/product/setting/voiceControl/releaseRecord/index', 2, 'sys_admin', 0, 2, 2, '2022-08-31 10:28:20', '2022-08-31 14:05:05', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (5188143539115950080, 6609772312344297472, 'SceneLinkage', '场景联动', 'link', '', '', 1, 10, 1, 2, '/product/product/setting/basicConfig/sceneLinkage/index', '', '/product/product/setting/basicConfig/sceneLinkage/index', 2, 'sys_admin', 0, 2, 2, '2023-02-20 13:51:22', '2023-02-20 13:51:22', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (4569134046918901760, 6609772312344297472, 'RuleConfig', '规则配置', 'code', '', '', 1, 11, 1, 2, '/product/product/setting/basicConfig/ruleConfig/index', '', '/product/product/setting/basicConfig/ruleConfig/index', 2, 'sys_admin', 0, 2, 2, '2023-07-06 16:35:11', '2023-07-06 16:35:11', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (8661170261715746816, 7017623568535945216, 'ActivateData', '激活数据', 'heat-map', '', '', 1, 1, 1, 1, '/dataManage/deviceStatistics/activateData/index', '', '/dataManage/deviceStatistics/activateData/index', 2, 'sys_admin', 0, 2, 2, '2022-06-07 18:16:44', '2022-06-07 18:16:44', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (8862867757822017536, 7017623568535945216, 'FaultStatistics', '故障统计', 'stock', '', '', 1, 2, 1, 1, '/dataManage/deviceStatistics/faultStatistics/index', '', '/dataManage/deviceStatistics/faultStatistics/index', 2, 'sys_admin', 0, 2, 2, '2022-06-07 18:17:32', '2022-06-07 18:17:32', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (6882064572988096512, 7087187032707006464, 'appUsers', '用户管理', 'bar-chart', '', '', 1, 1, 1, 1, '/dataManage/appUsers/index', '', '/dataManage/appUsers/index', 2, 'sys_admin', 0, 1, 2, '2022-06-07 08:41:52', '2024-03-07 08:36:21', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (6664638759247642624, 7087187032707006464, 'UserDetails', '用户详情', 'edit', '', '', 1, 3, 1, 2, '/dataManage/appUsers/userDetails/index', '', '/dataManage/appUsers/userDetails/index', 2, 'sys_admin', 0, 2, 2, '2022-06-07 08:42:44', '2022-06-07 08:42:44', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (2276630482433179648, 7087187032707006464, 'UserAnalysis', '用户分析', 'dot-chart', '', '', 1, 2, 1, 1, '/dataManage/appUsers/userAnalysis/index', '', '/dataManage/appUsers/userAnalysis/index', 2, 'sys_admin', 0, 2, 2, '2022-06-07 17:33:17', '2022-06-07 17:33:17', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (7214225066630938624, 8234967918207664128, '多语言', '产品多语言', 'book', '', '', 1, 2, 1, 1, '/appDevelop/multilingual/index', '', '/appDevelop/multilingual/index', 2, 'sys_admin', 0, 2, 2, '2022-06-02 15:03:35', '2022-06-02 16:29:49', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (4285149633635581952, 8234967918207664128, 'appMultilingual', 'APP 多语言', 'copy', '', '', 1, 3, 1, 1, '/appDevelop/multilingual/appMultilingual/index', '', '/appDevelop/multilingual/appMultilingual/index', 2, 'sys_admin', 0, 2, 2, '2022-06-13 09:35:44', '2022-06-13 09:37:23', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (4938411809444036608, 8737017658974765056, 'Authority', '授权管理', 'user', '', '', 1, 1, 1, 1, '/authority/index', '', '/authority/index', 2, 'sys_admin', 0, 1, 2, '2022-05-12 15:52:42', '2024-03-07 08:36:45', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (4835238792686305280, 8737017658974765056, 'Authorization', '权限设置', 'user', '', '', 1, 2, 1, 2, '/authority/authorization/index', '', '/authority/authorization/index', 2, 'sys_admin', 0, 2, 2, '2022-05-12 15:53:07', '2022-05-12 15:53:07', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (4593671981483786240, 8737017658974765056, 'CustomRole', '自定义角色', 'border-right', '', '', 1, 3, 1, 2, '/authority/customRole/index', '', '/authority/customRole/index', 2, 'sys_admin', 0, 2, 2, '2022-05-12 15:54:04', '2022-05-12 15:54:04', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (7652015955438043136, 8885045495148412928, 'Overview', '概览', 'bar-chart', '', '', 1, 1, 1, 1, '/dataManage/overview/index', '', '/dataManage/overview/index', 2, 'sys_admin', 0, 2, 2, '2022-06-07 08:38:49', '2022-06-07 08:38:49', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (7087187032707006464, 8885045495148412928, 'App', 'APP', 'line-chart', '', '', 1, 2, 1, 1, '/dataManage/appUsers/', '/dataManage/appUser/index', 'PageView', 2, 'sys_admin', 0, 2, 2, '2022-06-07 08:41:03', '2022-06-07 08:46:39', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (7017623568535945216, 8885045495148412928, 'DeviceStatistics', '设备统计', 'line-chart', '', '', 1, 3, 1, 1, '/dataManage/deviceStatistics', '/dataManage/deviceStatistics/activateData/index', 'PageView', 2, 'sys_admin', 0, 2, 2, '2022-06-07 18:10:12', '2022-06-07 18:15:42', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (4100852806750142464, 8959096378445692928, 'OEMApp', 'APP', 'icon-APPkaifa', '', '', 1, 1, 1, 1, '/appDevelop/OEMApp', '/appDevelop/customizedApp/index', 'PageView', 2, 'sys_admin', 0, 1, 2, '2022-05-27 10:13:15', '2024-02-28 10:23:22', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (8234967918207664128, 8959096378445692928, '多语言', '多语言', 'tags', '', '', 1, 2, 1, 1, '/appDevelop/multilingual', '/appDevelop/multilingual/index', 'PageView', 2, 'sys_admin', 0, 2, 2, '2022-06-02 14:59:31', '2022-06-02 16:30:03', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (4367313833849421824, 8959096378445692928, 'IntelligentScene', '智能场景', 'file-ppt', '', '', 1, 1, 1, 1, '/appDevelop/intelligentScene', '/appDevelop/intelligentScene/index.vue', 'PageView', 2, 'sys_admin', 0, 2, 2, '2022-11-08 10:14:43', '2022-11-08 10:14:59', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (3115948585393422336, 9159164745528279040, 'Driver', '设备', 'appstore', '', '', 1, 2, 1, 1, '/driver', '', 'PageView', 2, 'sys_admin', 0, 2, 2, '2022-05-25 14:18:20', '2022-06-23 09:23:28', NULL, 0, 0);
INSERT INTO `iot_open_system`.`t_open_auth_rule` (`id`, `pid`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `status`, `always_show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `model_id`, `is_cache`, `is_hide_child_menu`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES (6609772312344297472, 9159164745528279040, 'ProductItem', '产品', 'icon-chanpinkaifa', '', '', 1, 1, 1, 1, '/productItem', '/product/product/index', 'PageView', 2, 'sys_admin', 0, 1, 2, '2022-06-23 09:11:55', '2024-01-10 17:27:46', NULL, 0, 0);

-- ----------------------------
-- Table structure for t_open_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `t_open_casbin_rule`;
CREATE TABLE `t_open_casbin_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v6` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v7` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_t_sys_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE,
  UNIQUE INDEX `idx_t_open_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1650006302 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_open_casbin_rule
-- ----------------------------
-- 默认iotcode的角色权限
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650005781, 'g', '5083013370315964416', '1145738928238526464', 'iotcode', '', '', '', '', '');

-- 默认管理员的菜单权限
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006371, 'p', '1145738928238526464', '1024434789655937024', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006320, 'p', '1145738928238526464', '1178813131179065344', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006382, 'p', '1145738928238526464', '1338271198261379072', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006346, 'p', '1145738928238526464', '1342078855191560192', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006306, 'p', '1145738928238526464', '1535601662729551872', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006321, 'p', '1145738928238526464', '1576791311216705536', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006375, 'p', '1145738928238526464', '1630572854741401600', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006367, 'p', '1145738928238526464', '1752185267978076160', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006406, 'p', '1145738928238526464', '1754747229723066368', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006343, 'p', '1145738928238526464', '1831562684512763904', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006354, 'p', '1145738928238526464', '1891977961314091008', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006409, 'p', '1145738928238526464', '1982825917403004928', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006344, 'p', '1145738928238526464', '2015813318220021760', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006307, 'p', '1145738928238526464', '2084296393052028928', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006322, 'p', '1145738928238526464', '2162740003589750784', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006328, 'p', '1145738928238526464', '2271806318794670080', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006340, 'p', '1145738928238526464', '2276630482433179648', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006410, 'p', '1145738928238526464', '2403733430450880512', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006304, 'p', '1145738928238526464', '241', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006331, 'p', '1145738928238526464', '242', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006352, 'p', '1145738928238526464', '2422285703049019392', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006303, 'p', '1145738928238526464', '243', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006330, 'p', '1145738928238526464', '244', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006305, 'p', '1145738928238526464', '245', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006329, 'p', '1145738928238526464', '246', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006353, 'p', '1145738928238526464', '2469472581170659328', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006356, 'p', '1145738928238526464', '248', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006366, 'p', '1145738928238526464', '249', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006379, 'p', '1145738928238526464', '251', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006377, 'p', '1145738928238526464', '255', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006372, 'p', '1145738928238526464', '256', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006388, 'p', '1145738928238526464', '257', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006392, 'p', '1145738928238526464', '258', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006396, 'p', '1145738928238526464', '2928553496183668736', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006370, 'p', '1145738928238526464', '2977052235772362752', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006335, 'p', '1145738928238526464', '3115948585393422336', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006327, 'p', '1145738928238526464', '3232128938264133632', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006403, 'p', '1145738928238526464', '3269957101346193408', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006399, 'p', '1145738928238526464', '3317268708988125184', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006326, 'p', '1145738928238526464', '3833491004791357440', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006402, 'p', '1145738928238526464', '3898960696095899648', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006313, 'p', '1145738928238526464', '4100852806750142464', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006310, 'p', '1145738928238526464', '4115308261576441856', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006362, 'p', '1145738928238526464', '4285149633635581952', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006347, 'p', '1145738928238526464', '431674596966629376', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006324, 'p', '1145738928238526464', '4367313833849421824', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006401, 'p', '1145738928238526464', '4569134046918901760', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006345, 'p', '1145738928238526464', '4593347916386435072', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006359, 'p', '1145738928238526464', '4593671981483786240', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006334, 'p', '1145738928238526464', '4835238792686305280', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006325, 'p', '1145738928238526464', '4838154040086790144', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006333, 'p', '1145738928238526464', '4897173437564223488', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006312, 'p', '1145738928238526464', '4938411809444036608', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006332, 'p', '1145738928238526464', '5048246558223400960', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006350, 'p', '1145738928238526464', '5088011157059502080', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006398, 'p', '1145738928238526464', '5188143539115950080', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006308, 'p', '1145738928238526464', '5209425033708863488', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006364, 'p', '1145738928238526464', '5299407709798301696', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006397, 'p', '1145738928238526464', '5500705610791813120', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006365, 'p', '1145738928238526464', '6214091536351723520', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006395, 'p', '1145738928238526464', '6590486867232063488', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006318, 'p', '1145738928238526464', '6609772312344297472', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006360, 'p', '1145738928238526464', '6664638759247642624', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006319, 'p', '1145738928238526464', '6829802911032049664', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006316, 'p', '1145738928238526464', '6882064572988096512', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006301, 'p', '1145738928238526464', '6890578711163273216', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006381, 'p', '1145738928238526464', '6960539848083931136', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006361, 'p', '1145738928238526464', '7017623568535945216', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006342, 'p', '1145738928238526464', '7079616222605508608', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006309, 'p', '1145738928238526464', '7085924647471841280', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006339, 'p', '1145738928238526464', '7087187032707006464', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006338, 'p', '1145738928238526464', '7214225066630938624', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006363, 'p', '1145738928238526464', '7312533820899491840', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006368, 'p', '1145738928238526464', '7466223372586614784', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006400, 'p', '1145738928238526464', '758281464944099328', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006315, 'p', '1145738928238526464', '7652015955438043136', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006374, 'p', '1145738928238526464', '7740460241187143680', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006369, 'p', '1145738928238526464', '781674127982231552', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006380, 'p', '1145738928238526464', '8155109377754169344', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006351, 'p', '1145738928238526464', '8177095095929962496', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006337, 'p', '1145738928238526464', '8234967918207664128', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006385, 'p', '1145738928238526464', '8279675191588454400', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006389, 'p', '1145738928238526464', '8482041016558190592', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006317, 'p', '1145738928238526464', '8661170261715746816', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006393, 'p', '1145738928238526464', '8667428119637819392', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006376, 'p', '1145738928238526464', '870977444129964032', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006386, 'p', '1145738928238526464', '873599101126475776', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006373, 'p', '1145738928238526464', '8737017658974765056', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006394, 'p', '1145738928238526464', '8772330656223887360', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006390, 'p', '1145738928238526464', '8832823208554954752', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006341, 'p', '1145738928238526464', '8862867757822017536', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006358, 'p', '1145738928238526464', '8885045495148412928', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006336, 'p', '1145738928238526464', '8959096378445692928', 'ALL', '', '', '', '', '');
INSERT INTO `iot_open_system`.`t_open_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1650006311, 'p', '1145738928238526464', '9159164745528279040', 'ALL', '', '', '', '', '');


-- ----------------------------
-- Table structure for t_open_company
-- ----------------------------
DROP TABLE IF EXISTS `t_open_company`;
CREATE TABLE `t_open_company`  (
  `id` bigint(20) NOT NULL COMMENT '唯一主键',
  `tenant_id` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户主键',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '企业名称',
  `nature` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '企业性质',
  `license_no` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '营业执照号码',
  `license` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '营业执照',
  `legal_person` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '法人姓名',
  `apply_person` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '申请人姓名',
  `idcard` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '身份证号',
  `idcard_front_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '身份证照片',
  `idcard_after_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `address` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系地址',
  `status` int(11) NULL DEFAULT NULL COMMENT '状态（=1 未提交 ,=2 认证中,   =3 已认证, =4 禁用',
  `account_type` int(10) NULL DEFAULT NULL COMMENT '账号类型 1 企业账号, 2 个人账号',
  `case_remak` varchar(1200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '审核备注(拒绝需要填写拒绝原因)',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户邮箱',
  `is_real_name` int(11) NULL DEFAULT NULL COMMENT '是否实名(1 已实名, 2 未实名)',
  `phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `request_auth_at` timestamp NULL DEFAULT NULL COMMENT '申请认证时间',
  `region` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '区域(中国/长沙)',
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '登录用户名(冗余字段.)',
  `created_by` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间标识',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_open_company
-- ----------------------------
INSERT INTO `t_open_company` VALUES (5083084987557642240, 'iotcode', 5083013370315964416, '某某科技有限公司', '2', '1689547865423156987', 'https://osspublic.aithinker.com/firmware/5ebe8248-d370-42a6-8905-e0af7cc6b801.png', 'zzz', 'zzz', '430482199004124069', 'https://osspublic.aithinker.com/firmware/c0785e37-27ee-449d-8f18-6ab061123c95.png', 'https://osspublic.aithinker.com/firmware/dce9d80e-7070-42da-a11e-583a1a48f9f0.png', '', 2, 1, '123456', '', 2, '13598756236', '2022-11-01 10:12:30', '1,7,247,3021', 'opensource@dev.com', 5083013370315964416, '2022-06-28 08:49:21', 1, '2023-05-23 14:29:40', NULL);

-- ----------------------------
-- Table structure for t_open_company_auth_logs
-- ----------------------------
DROP TABLE IF EXISTS `t_open_company_auth_logs`;
CREATE TABLE `t_open_company_auth_logs`  (
  `id` bigint(20) NOT NULL,
  `company_id` bigint(20) NULL DEFAULT NULL COMMENT '公司id',
  `auth_result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '审核结果',
  `auth_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作人',
  `auth_date` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '操作时间',
  `why` varchar(1200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '理由',
  `created_by` bigint(20) NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_by` bigint(20) NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_open_company_auth_logs
-- ----------------------------

-- ----------------------------
-- Table structure for t_open_company_connect
-- ----------------------------
DROP TABLE IF EXISTS `t_open_company_connect`;
CREATE TABLE `t_open_company_connect`  (
  `id` bigint(20) NOT NULL,
  `tenant_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '企业主键id',
  `account_type` int(1) NULL DEFAULT NULL COMMENT '账号类型（1 主账号 0 子账号）',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '姓名',
  `user_id` bigint(20) NULL DEFAULT NULL COMMENT '开发者用户编号',
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '账号',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '电话',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `job` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '职位',
  `created_by` bigint(20) NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` bigint(20) NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_open_config
-- ----------------------------
DROP TABLE IF EXISTS `t_open_config`;
CREATE TABLE `t_open_config`  (
  `config_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键值',
  `config_type` tinyint(1) NULL DEFAULT 0 COMMENT '系统内置（Y是 N否）',
  `create_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint(20) NULL DEFAULT NULL,
  `updated_by` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`config_id`) USING BTREE,
  UNIQUE INDEX `uni_config_key`(`config_key`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_open_login_log
-- ----------------------------
DROP TABLE IF EXISTS `t_open_login_log`;
CREATE TABLE `t_open_login_log`  (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `login_time` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 246 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统访问记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_open_login_log
-- ----------------------------

-- ----------------------------
-- Table structure for t_open_model_info
-- ----------------------------
DROP TABLE IF EXISTS `t_open_model_info`;
CREATE TABLE `t_open_model_info`  (
  `model_id` bigint(20) NOT NULL COMMENT '模型ID',
  `model_category_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模板分类id',
  `model_name` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模型标识',
  `model_title` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模型名称',
  `model_pk` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主键字段',
  `model_order` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '默认排序字段',
  `model_sort` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表单字段排序',
  `model_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '列表显示字段，为空显示全部',
  `model_edit` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '可编辑字段，为空则除主键外均可以编辑',
  `model_indexes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '索引字段',
  `search_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '高级搜索的字段',
  `create_time` bigint(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` bigint(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `model_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态',
  `model_engine` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'MyISAM' COMMENT '数据库引擎',
  `create_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`model_id`) USING BTREE,
  UNIQUE INDEX `name_uni`(`model_name`) USING BTREE COMMENT '模型名唯一'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文档模型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_open_model_info
-- ----------------------------

-- ----------------------------
-- Table structure for t_open_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `t_open_oper_log`;
CREATE TABLE `t_open_oper_log`  (
  `oper_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '模块标题',
  `business_type` int(2) NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求方式',
  `operator_type` int(1) NULL DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作地点',
  `oper_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求参数',
  `json_result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '返回参数',
  `status` int(1) NULL DEFAULT 0 COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '错误消息',
  `oper_time` timestamp NULL DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 16692 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_open_oper_log
-- ----------------------------

-- ----------------------------
-- Table structure for t_open_role
-- ----------------------------
DROP TABLE IF EXISTS `t_open_role`;
CREATE TABLE `t_open_role`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态;1: 正常; 2:禁用;',
  `list_order` int(11) NULL DEFAULT 0 COMMENT '排序',
  `is_default` int(11) NULL DEFAULT NULL COMMENT '是否默认角色(1 默认, 2 非默认)',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint(3) UNSIGNED NULL DEFAULT 3 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `is_admin` int(11) NULL DEFAULT NULL COMMENT '是否管理员角色(每一个空间都默认一个管理员角色,并且有所有菜单权限)[1 是, 2 否]',
  `created_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `created_by` bigint(20) NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` bigint(20) NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7889966926082965505 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

INSERT INTO `iot_open_system`.`t_open_role` (`id`, `tenant_id`, `status`, `list_order`, `is_default`, `name`, `remark`, `data_scope`, `is_admin`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_at`) VALUES (1145738928238526464, 'VJ1VgP', 1, 1, 1, '管理员', '', 1, 1, '2022-05-24 09:31:28', 7891575059146440704, '2022-05-24 09:31:28', 7891575059146440704, NULL);

-- ----------------------------
-- Table structure for t_open_user
-- ----------------------------
DROP TABLE IF EXISTS `t_open_user`;
CREATE TABLE `t_open_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '用户昵称',
  `birthday` int(11) NULL DEFAULT 0 COMMENT '生日',
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `user_salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密盐',
  `user_status` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `user_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '用户登录邮箱',
  `sex` tinyint(2) NULL DEFAULT 0 COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '用户头像',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '联系地址',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT ' 描述信息',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  `account_type` int(1) NULL DEFAULT NULL COMMENT '账号类型（=1 企业 =2 个人）',
  `account_origin` int(1) NULL DEFAULT NULL COMMENT '账号来源（=1 注册 =2 管理员创建）',
  `company_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `created_by` bigint(20) NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `updated_by` bigint(20) NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8980439143754596353 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_open_user
-- ----------------------------
INSERT INTO `t_open_user` (`id`, `user_name`, `mobile`, `user_nickname`, `birthday`, `user_password`, `user_salt`, `user_status`, `user_email`, `sex`, `avatar`, `remark`, `address`, `describe`, `last_login_ip`, `last_login_time`, `account_type`, `account_origin`, `company_name`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_at`) VALUES (5083013370315964416, 'opensource@dev.com', '18800000001', 'opensource@dev.com', 0, '4fb4866ca970a3d38d7bf14e642113f2', 'rzKLbEwJcv', 1, '', 0, '', '', '', '', '0.0.0.0', '0000-00-00 00:00:00', 1, NULL, '某某科技有限公司', '2024-02-26 00:00:00', 0, '2024-03-18 17:35:30', 0, NULL);

-- ----------------------------
-- Table structure for t_open_user_company
-- ----------------------------
DROP TABLE IF EXISTS `t_open_user_company`;
CREATE TABLE `t_open_user_company`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户主键id',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `company_id` bigint(20) NULL DEFAULT NULL COMMENT '企业主键id',
  `tenant_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户Id',
  `user_type` int(10) NULL DEFAULT NULL COMMENT '账号类型(1主账号 ,2子账号 )',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_by` bigint(20) NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_by` bigint(20) NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9219021856844120222 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_open_user_company
-- ----------------------------
INSERT INTO `t_open_user_company` VALUES (5083149572486627328, 5083013370315964416, 'opensource@dev.com', 5083084987557642240, 'iotcode', 1, '', 0, '2022-06-28 08:49:21', 0, '2022-06-28 08:49:21');

-- ----------------------------
-- Table structure for t_open_user_online
-- ----------------------------
DROP TABLE IF EXISTS `t_open_user_online`;
CREATE TABLE `t_open_user_online`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户标识',
  `token` varchar(600) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `ip` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `explorer` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_token`(`token`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 9091853428392361985 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户在线状态表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
