DROP DATABASE IF EXISTS iot_log;
CREATE DATABASE iot_log;
USE iot_log;

CREATE TABLE IF NOT EXISTS iot_log.iot_device_log_DomeDev
(
    `id` String,
    `product_key` String,
    `device_id` String,
    `ver` String,
    `ns` String,
    `name` String,
    `gid` String,
    `from` String,
    `onlineStatus` String,
    `startTs` DateTime64(3),
    `endTs` Nullable(DateTime64(3)),
    `controls` String,
    `light_brightness` Nullable(Int64),
    `mist_level2` Nullable(Int32),
    `charging_state` Nullable(Bool),
    `light_mode` Nullable(Int32),
    `mist_countdown` Nullable(Int32),
    `essential_oil_id` Nullable(String),
    `play_switch` Nullable(Int32),
    `play_volume` Nullable(Int64),
    `play_source` Nullable(Int32),
    `mist_level3` Nullable(Int32),
    `play_music` Nullable(Int32),
    `light_switch` Nullable(Bool),
    `remain_electricity` Nullable(Int64),
    `classic_bluetooth_switch` Nullable(Bool),
    `powerstate` Nullable(Bool),
    `mist_countdown_remain` Nullable(Int64),
    `light_colour` Nullable(String),
    `updated_at` Nullable(DateTime64(3)),
    `created_at` DateTime64(3),
    `code` Nullable(Int32)
)
ENGINE = MergeTree
ORDER BY tuple()
SETTINGS index_granularity = 8192;

CREATE TABLE IF NOT EXISTS iot_log.iot_device_log_PKEVvgUr
(
    `id` String,
    `product_key` String,
    `device_id` String,
    `ver` String,
    `ns` String,
    `name` String,
    `gid` String,
    `from` String,
    `onlineStatus` String,
    `startTs` DateTime64(3),
    `endTs` Nullable(DateTime64(3)),
    `controls` String,
    `light_brightness` Nullable(Int64),
    `light_colour` Nullable(String),
    `play_switch` Nullable(Int32),
    `mist_level2` Nullable(Int32),
    `lymc_hide` Nullable(String),
    `remain_electricity` Nullable(Int64),
    `essential_oil_id` Nullable(String),
    `light_switch` Nullable(Bool),
    `powerstate` Nullable(Bool),
    `classic_bluetooth_switch` Nullable(Bool),
    `charging_state` Nullable(Bool),
    `play_music` Nullable(Int32),
    `mist_level3` Nullable(Int32),
    `play_source` Nullable(Int32),
    `mist_level` Nullable(Int32),
    `mist_countdown` Nullable(Int32),
    `mist_countdown_remain` Nullable(Int64),
    `play_volume` Nullable(Int64),
    `light_mode` Nullable(Int32),
    `updated_at` Nullable(DateTime64(3)),
    `created_at` DateTime64(3),
    `code` Nullable(Int32)
)
ENGINE = MergeTree
ORDER BY tuple()
SETTINGS index_granularity = 8192;

CREATE TABLE IF NOT EXISTS iot_log.iot_device_log_PKiXITWV
(
    `id` String,
    `product_key` String,
    `device_id` String,
    `ver` String,
    `ns` String,
    `name` String,
    `gid` String,
    `from` String,
    `onlineStatus` String,
    `startTs` DateTime64(3),
    `endTs` Nullable(DateTime64(3)),
    `controls` String,
    `light_mode` Nullable(Int32),
    `mist_countdown_remain` Nullable(Int64),
    `play_source` Nullable(Int32),
    `lymc_hide` Nullable(String),
    `light_colour` Nullable(String),
    `play_volume` Nullable(Int64),
    `charging_state` Nullable(Bool),
    `classic_bluetooth_switch` Nullable(Bool),
    `light_brightness` Nullable(Int64),
    `play_switch` Nullable(Int32),
    `mist_level3` Nullable(Int32),
    `mist_level2` Nullable(Int32),
    `mist_level` Nullable(Int32),
    `remain_electricity` Nullable(Int64),
    `powerstate` Nullable(Bool),
    `light_switch` Nullable(Bool),
    `play_music` Nullable(Int32),
    `essential_oil_id` Nullable(String),
    `mist_countdown` Nullable(Int32),
    `updated_at` Nullable(DateTime64(3)),
    `created_at` DateTime64(3),
    `code` Nullable(Int32)
)
ENGINE = MergeTree
ORDER BY tuple()
SETTINGS index_granularity = 8192;

CREATE TABLE IF NOT EXISTS iot_log.iot_device_log_PKNg1ZmK
(
    `id` String,
    `product_key` String,
    `device_id` String,
    `ver` String,
    `ns` String,
    `name` String,
    `gid` String,
    `from` String,
    `onlineStatus` String,
    `startTs` DateTime64(3),
    `endTs` Nullable(DateTime64(3)),
    `controls` String,
    `dff` Nullable(Bool),
    `play_music` Nullable(Int32),
    `charging_state` Nullable(Bool),
    `light_mode` Nullable(Int32),
    `classic_bluetooth_switch` Nullable(Bool),
    `remain_electricity` Nullable(Int64),
    `powerstate` Nullable(Bool),
    `light_colour` Nullable(String),
    `essential_oil_id` Nullable(String),
    `light_switch` Nullable(Bool),
    `play_switch` Nullable(Int32),
    `mist_level3` Nullable(Int32),
    `lymc_hide` Nullable(String),
    `mist_countdown` Nullable(Int32),
    `light_brightness` Nullable(Int64),
    `dd` Nullable(Bool),
    `mist_level` Nullable(Int32),
    `mist_countdown_remain` Nullable(Int64),
    `play_volume` Nullable(Int64),
    `mist_level2` Nullable(Int32),
    `updated_at` Nullable(DateTime64(3)),
    `created_at` DateTime64(3),
    `code` Nullable(Int32)
)
ENGINE = MergeTree
ORDER BY tuple()
SETTINGS index_granularity = 8192;

CREATE TABLE IF NOT EXISTS iot_log.iot_device_operation_fail_log
(
    `id` Int64,
    `device_id` String,
    `type` Int32 DEFAULT 0,
    `content` String,
    `user_account` String,
    `user_id` String,
    `tenant_id` String,
    `app_key` String,
    `created_at` Int64
)ENGINE = Log;

CREATE TABLE IF NOT EXISTS iot_log.iot_device_operation_fail_log_ex
(
    `id` Int64,
    `device_id` String,
    `type` Int32 DEFAULT 0,
    `content` String,
    `user_account` String,
    `user_id` String,
    `tenant_id` String,
    `app_key` String,
    `created_at` DateTime,
    `product_key` String,
    `code` Int32,
    `timezone` String,
    `region` String,
    `lang` String,
    `os` String,
    `model` String,
    `version` String,
    `desc` String,
    `upload_from` String,
    `upload_method` String
)ENGINE = Log;

CREATE TABLE IF NOT EXISTS iot_log.nginx_access
(
    `RemoteAddr` String,
    `TimeIso8601` DateTime,
    `RequestUri` String,
    `Status` UInt16,
    `RequestTime` Float64,
    `UpstreamHeaderTime` Float64,
    `UpstreamConnectTime` Float64,
    `UpstreamResponseTime` Float64,
    `BytesSent` UInt64,
    `BodyBytesSent` UInt64,
    `HttpReferer` String,
    `HttpUserAgent` String,
    `Connection` String,
    `RequestMethod` String,
    `ServerProtocol` String,
    `RequestLength` UInt64
) ENGINE = Log;

CREATE TABLE IF NOT EXISTS iot_log.t_iot_log_app_records
(
    `id` Int64 COMMENT '编号',
    `account` String COMMENT '用户账号,手机号或者邮箱',
    `tenant_id` String COMMENT '租户id',
    `app_key` String COMMENT 'appKey',
    `log_type` String COMMENT '日志类型',
    `event_name` String COMMENT '事件名称',
    `details` Map(String, String) COMMENT '详情',
    `created_at` DateTime DEFAULT now() COMMENT '创建时间',
    `region_server_id` Int64 COMMENT '区域服务编号'
)
ENGINE = MergeTree
PARTITION BY toYYYYMMDD(created_at)
ORDER BY (account,tenant_id,app_key)
SETTINGS index_granularity = 8192;

CREATE TABLE IF NOT EXISTS iot_log.t_iot_log_app_user
(
    `id` Int64 COMMENT '编号',
    `account` String COMMENT '用户账号,手机号或者邮箱',
    `tenant_id` String COMMENT '租户id',
    `app_key` String COMMENT 'appKey',
    `app_name` String COMMENT 'APP名称',
    `region` String COMMENT '服务区',
    `login_time` DateTime COMMENT '最后登录时间',
    `created_at` DateTime DEFAULT now() COMMENT '创建时间',
    `region_server_id` Int64 COMMENT '区域服务编号'
)
ENGINE = MergeTree
PARTITION BY toYYYYMMDD(created_at)
ORDER BY (account,tenant_id,app_key)
SETTINGS index_granularity = 8192;