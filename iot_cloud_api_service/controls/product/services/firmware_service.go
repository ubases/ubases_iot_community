package services

import (
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type PmFirmwareService struct {
}

// 固件详细
func (s PmFirmwareService) GetPmFirmwareDetail(id string) (*entitys.PmFirmwareEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientPmFirmwareService.FindById(context.Background(), &protosService.PmFirmwareFilter{Id: rid})
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
	return entitys.PmFirmware_pb2e(data), err
}

// QueryPmFirmwareList 固件列表
func (s PmFirmwareService) QueryPmFirmwareList(filter entitys.PmFirmwareQuery) ([]*entitys.PmFirmwareEntitys, int64, error) {
	var queryObj = new(protosService.PmFirmware)
	if filter.Query != nil {
		queryObj.Status = filter.Query.Status
		if filter.Query.Name != "" {
			queryObj.Name = filter.Query.Name
		}
		if filter.Query.NameEn != "" {
			queryObj.NameEn = filter.Query.NameEn
		}
		if filter.Query.Flag != "" {
			queryObj.Flag = filter.Query.Flag
		}
		if filter.Query.Type != 0 {
			queryObj.Type = iotutil.ToString(filter.Query.Type)
		}
		if filter.Query.ModuleIds != nil {
			queryObj.ModuleIds = filter.Query.ModuleIds
		}
	}
	rep, err := rpc.ClientPmFirmwareService.Lists(context.Background(), &protosService.PmFirmwareListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		Query:     queryObj,
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.PmFirmwareEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.PmFirmware_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddPmFirmware 新增固件
func (s PmFirmwareService) AddPmFirmware(req entitys.PmFirmwareEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.PmFirmware_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Status = 2 // 固件模组新增，默认为禁用
	res, err := rpc.ClientPmFirmwareService.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//services.SetDefaultTranslate(context.Background(), "t_pm_firmware", res.Data, "name", req.Name, req.NameEn)
	return iotutil.ToString(saveObj.Id), err
}

// 修改固件
func (s PmFirmwareService) UpdatePmFirmware(req entitys.PmFirmwareEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientPmFirmwareService.UpdateAll(context.Background(), entitys.PmFirmware_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//services.SetDefaultTranslate(context.Background(), "t_pm_firmware", req.Id, "name", req.Name, req.NameEn)
	return iotutil.ToString(req.Id), err
}

// 删除固件
func (s PmFirmwareService) DeletePmFirmware(req entitys.PmFirmwareFilter) error {
	rep, err := rpc.ClientPmFirmwareService.Delete(context.Background(), &protosService.PmFirmware{
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
func (s PmFirmwareService) SetStatusPmFirmware(req entitys.PmFirmwareFilter) error {
	if req.Id == "" {
		return errors.New("id not found")
	}
	if req.Status == 0 {
		return errors.New("status not found")
	}
	//如果是禁用则需检查固件模组绑定关系；
	if req.Status == 2 {
		rep, err := rpc.ClientModuleService.Find(context.Background(), &protosService.PmModuleFilter{
			FirmwareId: iotutil.ToInt64(req.Id),
			Status:     1,
		})
		if err != nil {
			return err
		}
		if rep.Code != 200 && rep.Message != ioterrs.ErrRecordNotFound {
			return errors.New(rep.Message)
		}
		if len(rep.Data) > 0 {
			return errors.New("固件已绑定模组已启用, 无法进行禁用操作")
		}
	}
	rep, err := rpc.ClientPmFirmwareService.UpdateStatus(context.Background(), &protosService.PmFirmware{
		Id:     iotutil.ToInt64(req.Id),
		Status: req.Status,
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}
