package ioterrs

// 内部服务错误
var (
	// iot_device_service内部服务错误定义
	ErrDeviceServiceUnavailable int32 = 423001
)

// iot_log_service内部服务错误定义
var (
	ErrLogServiceUnavailable int32 = 429001
)

// iot_job_service内部服务错误定义
var (
	ErrJobServiceUnavailable int32 = 456001
)
