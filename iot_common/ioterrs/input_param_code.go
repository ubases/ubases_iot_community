package ioterrs

// 输入参数错误

// iot_cloud_api_service输入参数错误定义
var (
	ErrAppLogUserListReqParam    int32 = 221001
	ErrAppLogRecordsListReqParam int32 = 221002
	ErrCloudQueryParamIsNil      int32 = 221003
	ErrCloudTenantIdEmpty        int32 = 221004
	ErrCloudProductKeyEmpty      int32 = 221005
	ErrProductHelpConfParam      int32 = 221006
	ErrProductHelpDocParam       int32 = 221007
	ErrShouldBindJSON            int32 = 221008
	ErrCreateFlashScreen         int32 = 221009
	ErrCloudRequestParamIsEmpty  int32 = 221010
	ErrFlashScreenAlreadyExist   int32 = 221011
	ErrFlashScreenAlreadyExpired int32 = 221012
	ErrCloudRequestParam         int32 = 221013
	ErrCloudVersionTooLow        int32 = 221014
	ErrCloudGenOrUploadPlist     int32 = 221015
	ErrMaterialTypeAlreadyExist  int32 = 221016
)

// iot-document-service输入参数错误定义
var (
// ErrAppLogUserListReqParam    int32 = 233001
// ErrAppLogRecordsListReqParam int32 = 221002
)

// iot_app_api_service输入参数错误定义
var (
	ErrAppCountDownAddParam int32 = 251001
	ErrAppDeviceIdEmpty     int32 = 251002
	ErrAppFuncKeyEmpty      int32 = 251003
	ErrAppCountDownNotExist int32 = 251004
	ErrAppTenantIdEmpty     int32 = 251005
	ErrAppKeyEmpty          int32 = 251006
	ErrAppQueryParamIsNil   int32 = 251007
	ErrAppRequestParam      int32 = 251008
	ErrInterAppointment     int32 = 251009
	ErrUserTimeToLocalTime  int32 = 251010
)
