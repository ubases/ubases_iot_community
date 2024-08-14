package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/service/email"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_message/model"
	"cloud_platform/iot_model/db_message/orm"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"strings"
)

type EmailSvc struct {
	Ctx context.Context
}

func (s *EmailSvc) SendEmailUserCode(request *protosService.SendEmailUserCodeRequest) (*protosService.SendEmailResponse, error) {
	ret := protosService.SendEmailResponse{Status: false}
	dbObj, err := s.GetEmailTpl(iotconst.NotifierBusinesses(request.TplType), request.Lang)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	templateInput := CodeInput{UserName: request.UserName, Code: request.Code}
	err = s.SendEmail(request.Email, dbObj.TplSubject, dbObj.TplContent, templateInput, dbObj, request.Lang, request.TenantId, request.AppKey)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	ret.Status = true
	return &ret, nil
}

func (s *EmailSvc) SendEmailUserLoggedIn(request *protosService.SendEmailUserLoggedInRequest) (*protosService.SendEmailResponse, error) {
	ret := protosService.SendEmailResponse{Status: false}
	dbObj, err := s.GetEmailTpl(iotconst.NB_LOGGEDIN, request.Lang)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	templateInput := LoggedInInput{UserName: request.UserName, IP: request.Ip}
	err = s.SendEmail(request.Email, dbObj.TplSubject, dbObj.TplContent, templateInput, dbObj, request.Lang, request.TenantId, request.AppKey)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	ret.Status = true
	return &ret, nil
}

func (s *EmailSvc) SendEmailUserRegister(request *protosService.SendEmailUserRegisterRequest) (*protosService.SendEmailResponse, error) {
	ret := protosService.SendEmailResponse{Status: false}
	dbObj, err := s.GetEmailTpl(iotconst.NB_REGISTER, request.Lang)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	templateInput := RegisterInput{UserName: request.UserName}
	err = s.SendEmail(request.Email, dbObj.TplSubject, dbObj.TplContent, templateInput, dbObj, request.Lang, request.TenantId, request.AppKey)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return &ret, err
	}
	ret.Status = true
	return &ret, nil
}

func (s *EmailSvc) GetEmailTpl(tplType iotconst.NotifierBusinesses, langNew string) (*model.TMsNoticeTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TMsNoticeTemplate
	do := t.WithContext(context.Background())
	do = do.Where(t.TplType.Eq(int32(tplType)), t.Method.Eq(2), t.Lang.Eq(langNew))
	dbObj, err := do.First()
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return nil, err
	}
	return dbObj, nil
}

func (s *EmailSvc) SendEmail(to, subject, body string, data interface{}, tempObj *model.TMsNoticeTemplate, lang, tenantId, appKey string) error {
	msg := email.SendEmailInput{To: strings.TrimSpace(to), Subject: subject, Body: body}
	if err := msg.GenerateBodyFromContent(body, data); err != nil {
		iotlogger.LogHelper.Error(err)
		return err
	}
	msg.RecordId = iotutil.GetNextSeqInt64()
	err := MsNoticerecordSvc.SaveNoticeRecord(msg.RecordId, 1, tempObj.TplName, tempObj.TplContent,
		lang, tenantId, appKey, tempObj, to)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return err
	}
	if err := EmailMgr.Send(&msg); err != nil {
		return err
	}
	return nil
}
