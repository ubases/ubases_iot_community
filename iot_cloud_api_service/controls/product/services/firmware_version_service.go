package services

import (
	"cloud_platform/iot_cloud_api_service/controls/common/commonGlobal"
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotenums"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type PmFirmwareVersionService struct {
}

// 固件版本详细
func (s PmFirmwareVersionService) GetPmFirmwareVersionDetail(id string) (*entitys.PmFirmwareVersionEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientPmFirmwareVersionService.FindById(context.Background(), &protosService.PmFirmwareVersionFilter{Id: rid})
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
	return entitys.PmFirmwareVersion_pb2e(data), err
}

// QueryPmFirmwareVersionList 固件版本列表
func (s PmFirmwareVersionService) QueryPmFirmwareVersionList(filter entitys.PmFirmwareVersionQuery) ([]*entitys.PmFirmwareVersionEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	if filter.Query == nil || filter.Query.FirmwareId == 0 {
		return nil, 0, errors.New("firmwareId not found")
	}
	rep, err := rpc.ClientPmFirmwareVersionService.Lists(context.Background(), &protosService.PmFirmwareVersionListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		OrderKey:  filter.SortField,
		OrderDesc: filter.Sort,
		Query: &protosService.PmFirmwareVersion{
			FirmwareId: iotutil.ToInt64(filter.Query.FirmwareId),
			Status:     filter.Query.Status,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.PmFirmwareVersionEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.PmFirmwareVersion_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddPmFirmwareVersion 新增固件版本
func (s PmFirmwareVersionService) AddPmFirmwareVersion(req entitys.PmFirmwareVersionEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.PmFirmwareVersion_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Status = 2
	res, err := rpc.ClientPmFirmwareVersionService.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	if req.FilePath != "" && req.ZipFilePath != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTPmFirmwareVersion, iotutil.ToString(req.Id), req.FilePath, req.ZipFilePath)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改固件版本
func (s PmFirmwareVersionService) UpdatePmFirmwareVersion(req entitys.PmFirmwareVersionEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	//req.UpdatedAt = time.Now().Unix()
	res, err := rpc.ClientPmFirmwareVersionService.UpdateAll(context.Background(), entitys.PmFirmwareVersion_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	if req.FilePath != "" && req.ZipFilePath != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTPmFirmwareVersion, iotutil.ToString(req.Id), req.FilePath, req.ZipFilePath)
	}
	return iotutil.ToString(req.Id), err
}

// 删除固件版本
func (s PmFirmwareVersionService) DeletePmFirmwareVersion(req entitys.PmFirmwareVersionFilter) error {
	rep, err := rpc.ClientPmFirmwareVersionService.Delete(context.Background(), &protosService.PmFirmwareVersion{
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

// SetStatusPmFirmwareVersion 禁用/启用固件
func (s PmFirmwareVersionService) SetStatusPmFirmwareVersion(req entitys.PmFirmwareVersionFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	if req.Status == 0 {
		return errors.New("status not found")
	}

	//如果是禁用需要判断与模组的绑定关系
	if req.Status == 2 {
		rep, err := rpc.ClientPmModuleFirmwareVersionService.Lists(context.Background(), &protosService.PmModuleFirmwareVersionListRequest{
			Query: &protosService.PmModuleFirmwareVersion{
				VersionId: req.Id,
			},
		})
		if err != nil {
			return err
		}
		if rep.Code != 200 {
			return errors.New(rep.Message)
		}
		if len(rep.Data) > 0 {
			return errors.New("本版本已与模组关联，请先解除关联关系")
		}
	}

	rep, err := rpc.ClientPmFirmwareVersionService.UpdateFields(context.Background(), &protosService.PmFirmwareVersionUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.PmFirmwareVersion{
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

// QueryPmFirmwareChangeVersionList 固件版本列表
func (s PmFirmwareVersionService) QueryPmFirmwareChangeVersionList(moduleId, page, limit int64) ([]*entitys.PmFirmwareVersionEntitys, int64, error) {
	rep, err := rpc.ClientPmFirmwareVersionService.ModuleFirmwareVersionList(context.Background(),
		&protosService.ModuleFirmwareVersionRequest{
			Page:     page,
			PageSize: limit,
			ModuleId: moduleId,
		})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.PmFirmwareVersionEntitys{}
	for _, src := range rep.Data {
		entitysObj := &entitys.PmFirmwareVersionEntitys{
			Id:         iotutil.ToString(src.Id),
			FirmwareId: src.FirmwareId,
			Version:    src.Version,
			Desc:       src.Desc,
			Status:     src.Status,
			StatusName: iotenums.ToShelfName(src.Status),
			FilePath:   src.FilePath,
			FileName:   src.FileName,
			FileKey:    src.FileKey,
			FileSize:   src.FileSize,
			UpdatedAt:  src.UpdatedAt.AsTime().Unix(),
		}
		resultList = append(resultList, entitysObj)
	}
	return resultList, rep.Total, err
}
