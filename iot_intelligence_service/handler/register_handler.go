package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {
	err := protosService.RegisterSceneIntelligenceLogServiceHandler(s.Server(), new(SceneIntelligenceLogHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSceneIntelligenceLogServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSceneIntelligenceConditionServiceHandler(s.Server(), new(SceneIntelligenceConditionHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSceneIntelligenceConditionServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSceneIntelligenceResultServiceHandler(s.Server(), new(SceneIntelligenceResultHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSceneIntelligenceResultServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSceneIntelligenceResultTaskServiceHandler(s.Server(), new(SceneIntelligenceResultTaskHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSceneIntelligenceLogServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSceneIntelligenceServiceHandler(s.Server(), new(SceneIntelligenceHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSceneIntelligenceServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSceneIntelligenceTaskServiceHandler(s.Server(), new(SceneIntelligenceTaskHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSceneIntelligenceTaskServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterSceneTemplateServiceHandler(s.Server(), new(SceneTemplateHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSceneTemplateServiceHandler 错误:%s", err.Error())
		return err
	}

	//不需要注册
	//scene_template_app_relation_handler.gen.go
	//scene_template_task_handler.gen.go
	//scene_template_condition_handler.gen.go

	return nil
}
