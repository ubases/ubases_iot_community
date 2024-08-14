package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_message/model"
	"cloud_platform/iot_model/db_message/orm"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
)

type SmsSvc struct {
	Ctx context.Context
}

func (s *SmsSvc) SendCode(request *protosService.SendSMSCodeRequest) (*protosService.SendSMSResponse, error) {
	ret := protosService.SendSMSResponse{Status: false}
	dbObj, err := s.GetSMSTpl(request.Lang, iotconst.NB_CODE)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	err = SmdMgr.SendSMS(CodeInput{UserName: request.UserName, Code: request.Code}, dbObj.TplCode, dbObj,
		request.Lang, request.TenantId, request.AppKey, request.PhoneNumber)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	ret.Status = true
	return &ret, nil
}

func (s *SmsSvc) SendSMSVerifyCode(request *protosService.SendSMSVerifyCodeRequest) (*protosService.SendSMSResponse, error) {
	ret := protosService.SendSMSResponse{Status: false}
	dbObj, err := s.GetSMSTplByType(request.Lang, request.PhoneType, iotconst.NotifierBusinesses(request.TplType))
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}

	err = SmdMgr.SendSMS(CodeInput{UserName: request.UserName, Code: request.Code, PhoneType: iotutil.ToString(request.PhoneType), Lang: request.Lang, Template: dbObj.TplContent}, dbObj.TplCode, dbObj,
		request.Lang, request.TenantId, request.AppKey, request.PhoneNumber)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	ret.Status = true
	return &ret, nil
}

func (s *SmsSvc) SendLoggedIn(request *protosService.SendSMSLoggedInRequest) (*protosService.SendSMSResponse, error) {
	ret := protosService.SendSMSResponse{Status: false}
	dbObj, err := s.GetSMSTpl(request.Lang, iotconst.NB_LOGGEDIN)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	err = SmdMgr.SendSMS(LoggedInInput{UserName: request.UserName, IP: request.Ip}, dbObj.TplCode, dbObj,
		request.Lang, request.TenantId, request.AppKey, request.PhoneNumber)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	ret.Status = true
	return &ret, nil
}

func (s *SmsSvc) SendRegister(request *protosService.SendSMSRegisterRequest) (*protosService.SendSMSResponse, error) {
	ret := protosService.SendSMSResponse{Status: false}
	dbObj, err := s.GetSMSTpl(request.Lang, iotconst.NB_REGISTER)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	err = SmdMgr.SendSMS(RegisterInput{UserName: request.UserName}, dbObj.TplCode, dbObj,
		request.Lang, request.TenantId, request.AppKey, request.PhoneNumber)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	ret.Status = true
	return &ret, nil
}

func (s *SmsSvc) GetSMSTplByType(lang string, method int32, tplType iotconst.NotifierBusinesses) (*model.TMsNoticeTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TMsNoticeTemplate
	do := t.WithContext(context.Background())
	//update by hogan 短信验证码翻译，需要读取对应语言的短信模板, 通知模板增加短信类型（短信、短信（英文））
	do = do.Where(t.TplType.Eq(int32(tplType)), t.Method.Eq(method), t.Lang.Eq(lang))
	dbObj, err := do.First()
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return nil, err
	}
	return dbObj, nil
}

func (s *SmsSvc) GetSMSTpl(lang string, tplType iotconst.NotifierBusinesses) (*model.TMsNoticeTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TMsNoticeTemplate
	do := t.WithContext(context.Background())
	//update by hogan 短信验证码翻译，需要读取对应语言的短信模板, 通知模板增加短信类型（短信、短信（英文））
	var method int32 = 3
	if lang == "zh" {
		method = 1
	}
	do = do.Where(t.TplType.Eq(int32(tplType)), t.Method.Eq(method))
	dbObj, err := do.First()
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return nil, err
	}
	return dbObj, nil
}
