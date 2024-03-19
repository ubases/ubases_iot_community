package ioterrs

// 业务级错误
// 数据库错误以ErrDB开头001-500范围
// 其他业务错误501-999范围
// iot_device_service业务错误定义
var (
	// 数据库错误以ErrDB开头001-500范围
	ErrDBCountDownCreate int32 = 323001
	ErrDBCountDownUpdate int32 = 323002
	ErrDBCountDownDelete int32 = 323003
	ErrDBCountDownGet    int32 = 323004
	ErrDBCountDownList   int32 = 323005
	ErrDBTimerCreate     int32 = 323006
	ErrDBTimerUpdate     int32 = 323007
	ErrDBTimerDelete     int32 = 323008
	ErrDBTimerGet        int32 = 323009
	ErrDBTimerList       int32 = 323010
	ErrDBDeviceCreate    int32 = 323011
	ErrDBDeviceUpdate    int32 = 323012
	ErrDBDeviceDelete    int32 = 323013
	ErrDBDeviceGet       int32 = 323014
	ErrDBDeviceList      int32 = 323015
	// 其他业务错误501-999范围
	ErrDeviceNotActive         int32 = 323501
	ErrDevProductKeyEmpty      int32 = 323502
	ErrDevJsonMarshal          int32 = 323503
	ErrCountDownNotExist       int32 = 323504
	ErrCountDownAlreadyStarted int32 = 323505
	ErrCountDownAlreadyStopped int32 = 323506
	ErrTimerAlreadyStarted     int32 = 323507
	ErrTimerAlreadyStopted     int32 = 323508
	ErrDevJsonUnMarshal        int32 = 323509
	ErrDevTimerFuncKeyIsEmpty  int32 = 323510
)

// iot-document-service业务错误定义
var (
	ErrDBProductHelpConfCreate int32 = 333001
	ErrDBProductHelpConfUpdate int32 = 333002
	ErrDBProductHelpConfDelete int32 = 333003
	ErrDBProductHelpConfGet    int32 = 333004
	ErrDBProductHelpConfList   int32 = 333005
	ErrDBProductHelpDocCreate  int32 = 333006
	ErrDBProductHelpDocUpdate  int32 = 333007
	ErrDBProductHelpDocDelete  int32 = 333008
	ErrDBProductHelpDocGet     int32 = 333009
	ErrDBProductHelpDocList    int32 = 333010
)

// iot_app_oem_service业务错误定义
var (
	ErrDBFlashScreenCreate            int32 = 355001
	ErrDBFlashScreenUpdate            int32 = 355002
	ErrDBFlashScreenDelete            int32 = 355003
	ErrDBFlashScreenGet               int32 = 355004
	ErrDBFlashScreenList              int32 = 355005
	ErrDBFlashScreenUserCreate        int32 = 355006
	ErrDBFlashScreenUserUpdate        int32 = 355007
	ErrDBFlashScreenUserDelete        int32 = 355008
	ErrDBFlashScreenUserGet           int32 = 355009
	ErrDBFlashScreenUserList          int32 = 355010
	ErrDBOemAppVersionRecordCreate    int32 = 355011
	ErrDBOemAppVersionRecordUpdate    int32 = 355012
	ErrDBOemAppVersionRecordDelete    int32 = 355013
	ErrDBOemAppVersionRecordGet       int32 = 355014
	ErrDBOemAppVersionRecordList      int32 = 355015
	ErrDBCustomAppVersionRecordCreate int32 = 355016
	ErrDBCustomAppVersionRecordUpdate int32 = 355017
	ErrDBCustomAppVersionRecordDelete int32 = 355018
	ErrDBCustomAppVersionRecordGet    int32 = 355019
	ErrDBCustomAppVersionRecordList   int32 = 355020
)

// iot_job_service业务错误定义
var (
	// 数据库错误以ErrDB开头001-500范围
	ErrDBJobCreate int32 = 356001
	ErrDBJobUpdate int32 = 356002
	ErrDBJobDelete int32 = 356003
	ErrDBJobGet    int32 = 356004
	ErrDBJobList   int32 = 356005
	// 其他业务错误501-999范围
	ErrJobTaskCreate         int32 = 356501
	ErrJobTaskUpdate         int32 = 356502
	ErrJobTaskDelete         int32 = 356503
	ErrJobTaskGet            int32 = 356504
	ErrJobTaskList           int32 = 356505
	ErrJobTaskAlreadyStarted int32 = 356506
	ErrJobTaskAlreadyStopped int32 = 356507
)

// iot_product_service业务错误定义
var (
	// 数据库错误以ErrDB开头001-500范围
	ErrDBMaterialCreate             int32 = 322001
	ErrDBMaterialUpdate             int32 = 322002
	ErrDBMaterialDelete             int32 = 322003
	ErrDBMaterialGet                int32 = 322004
	ErrDBMaterialList               int32 = 322005
	ErrDBMaterialRelCreate          int32 = 322006
	ErrDBMaterialRelUpdate          int32 = 322007
	ErrDBMaterialRelDelete          int32 = 322008
	ErrDBMaterialRelGet             int32 = 322009
	ErrDBMaterialRelList            int32 = 322010
	ErrDBMaterialTypeCreate         int32 = 322011
	ErrDBMaterialTypeUpdate         int32 = 322012
	ErrDBMaterialTypeDelete         int32 = 322013
	ErrDBMaterialTypeGet            int32 = 322014
	ErrDBMaterialTypeList           int32 = 322015
	ErrDBMaterialLanguageCreate     int32 = 322016
	ErrDBMaterialLanguageUpdate     int32 = 322017
	ErrDBMaterialLanguageDelete     int32 = 322018
	ErrDBMaterialLanguageGet        int32 = 322019
	ErrDBMaterialLanguageList       int32 = 322020
	ErrDBMaterialTypeLanguageCreate int32 = 322021
	ErrDBMaterialTypeLanguageUpdate int32 = 322022
	ErrDBMaterialTypeLanguageDelete int32 = 322023
	ErrDBMaterialTypeLanguageGet    int32 = 322024
	ErrDBMaterialTypeLanguageList   int32 = 322025
	ErrDBProductManualCreate        int32 = 322026
	ErrDBProductManualUpdate        int32 = 322027
	ErrDBProductManualDelete        int32 = 322028
	ErrDBProductManualGet           int32 = 322029
	ErrDBProductManualList          int32 = 322030
)

// iot_log_service业务错误定义
var (
	// 数据库错误以ErrDB开头001-500范围
	ErrDBAppLogUserList    int32 = 329001
	ErrDBAppLogRecordsList int32 = 329002
)
