DROP DATABASE IF EXISTS iot_device;
CREATE DATABASE iot_device;
USE iot_device;

CREATE TABLE IF NOT EXISTS iot_device.t_iot_ota_upgrade_record
(
    `id` Int64,
    `version` String,
    `publish_id` Int64,
    `product_key` String,
    `device_id` String,
    `original_version` String,
    `status` Int32,
    `is_gray` Int32,
    `tenant_id` String,
    `device_secret` String,
    `created_at` DateTime64(3),
    `updated_at` DateTime64(3),
    `ota_state` String,
    `ota_code` Int32,
    `ota_progress` Int32,
    `area` String,
    `fw_ver` String,
    `is_auto_upgrade` Bool
)
ENGINE = MergeTree
ORDER BY tuple()
SETTINGS index_granularity = 8192;
