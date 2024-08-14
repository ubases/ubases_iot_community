package services

import (
	"cloud_platform/iot_cloud_api_service/controls/common/commonGlobal"
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type PmModuleService struct {
}

// 模组芯片详细
func (s PmModuleService) GetPmModuleDetail(id string) (*entitys.PmModuleEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientModuleService.FindById(context.Background(), &protosService.PmModuleFilter{Id: rid})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}
	if len(req.Data) == 0 {
		return nil, errors.New("not found")
	}
	var data = req.Data[0]
	return entitys.PmModule_pb2e(data), err
}

// QueryPmModuleList 模组芯片列表
func (s PmModuleService) QueryPmModuleList(filter entitys.PmModuleQuery) ([]*entitys.PmModuleEntitys, int64, error) {

	firmwareType := ""
	if filter.Query.FirmwareType != 0 {
		firmwareType = iotutil.ToString(filter.Query.FirmwareType)
	}

	var newFilter = &protosService.PmModule{
		ModuleName:   filter.Query.ModuleName,
		ModuleNameEn: filter.Query.ModuleNameEn,
		FirmwareType: firmwareType,
		FirmwareFlag: filter.Query.FirmwareFlag,
		Status:       filter.Query.Status,
	}
	if filter.Query.FirmwareId != "" {
		newFilter.FirmwareId = iotutil.ToInt64(filter.Query.FirmwareId)
	}

	rep, err := rpc.ClientModuleService.Lists(context.Background(), &protosService.PmModuleListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query:    newFilter,
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}

	var resultList = []*entitys.PmModuleEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.PmModule_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddPmModule 新增模组芯片
func (s PmModuleService) AddPmModule(req entitys.PmModuleEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.PmModule_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Status = 2 // 模组新增，默认为禁用
	//saveObj.CreatedAt = timestamppb.Now()
	//saveObj.UpdatedAt = timestamppb.Now()
	res, err := rpc.ClientModuleService.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//services.SetDefaultTranslate(context.Background(), "t_pm_module", saveObj.Id, "name", req.ModuleName, req.ModuleNameEn)
	if req.ImgUrl != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTPmModule, iotutil.ToString(saveObj.Id), req.ImgUrl)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改模组芯片
func (s PmModuleService) UpdatePmModule(req entitys.PmModuleEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	//req.UpdatedAt = time.Now().Unix()
	res, err := rpc.ClientModuleService.UpdateAll(context.Background(), entitys.PmModule_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//services.SetDefaultTranslate(context.Background(), "t_pm_module", req.Id, "name", req.ModuleName, req.ModuleNameEn)
	if req.ImgUrl != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTPmModule, iotutil.ToString(req.Id), req.ImgUrl)
	}
	return iotutil.ToString(req.Id), err
}

// 修改部分模组芯片数据，按需修改
func (s PmModuleService) UpdatePartPmModule(req entitys.PmModuleEntitys) (string, error) {
	if req.Id == "" {
		return "", errors.New("模组芯片Id不能为空")
	}
	res, err := rpc.ClientModuleService.Update(context.Background(), entitys.PmModule_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//services.SetDefaultTranslate(context.Background(), "t_pm_module", req.Id, "name", req.ModuleName, req.ModuleNameEn)
	return iotutil.ToString(req.Id), err
}

// 删除模组芯片
func (s PmModuleService) DeletePmModule(req entitys.PmModuleFilter) error {
	rep, err := rpc.ClientModuleService.Delete(context.Background(), &protosService.PmModule{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}

// SetStatusPmFirmware 禁用/启用固件
func (s PmModuleService) SetStatusPmModule(req entitys.PmModuleFilter) error {
	if req.Id == "" {
		return errors.New("id not found")
	}
	if req.Status == 0 {
		return errors.New("status not found")
	}
	rep, err := rpc.ClientModuleService.UpdateFields(context.Background(), &protosService.PmModuleUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.PmModule{
			Id:     iotutil.ToInt64(req.Id),
			Status: req.Status,
		},
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}
