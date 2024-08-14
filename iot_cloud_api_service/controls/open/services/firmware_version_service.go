package services

import (
	"cloud_platform/iot_cloud_api_service/controls/common/commonGlobal"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpmFirmwareVersionService struct {
	Ctx context.Context
}

func (s OpmFirmwareVersionService) SetContext(ctx context.Context) OpmFirmwareVersionService {
	s.Ctx = ctx
	return s
}

// GetOpmFirmwareVersionDetail 固件版本详细
func (s OpmFirmwareVersionService) GetOpmFirmwareVersionDetail(id string) (*entitys.OpmFirmwareVersionEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientFirmwareVersionService.FindById(s.Ctx, &protosService.OpmFirmwareVersionFilter{Id: rid})
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
	return entitys.OpmFirmwareVersion_pb2e(data), err
}

// QueryOpmFirmwareVersionList 固件版本列表
func (s OpmFirmwareVersionService) QueryOpmFirmwareVersionList(filter entitys.OpmFirmwareVersionQuery) ([]*entitys.OpmFirmwareVersionEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	if filter.Query == nil || filter.Query.FirmwareId == 0 {
		return nil, 0, errors.New("固件Id不能为空")
	}
	rep, err := rpc.ClientFirmwareVersionService.Lists(s.Ctx, &protosService.OpmFirmwareVersionListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query: &protosService.OpmFirmwareVersion{
			Status:     *filter.Query.Status,
			FirmwareId: iotutil.ToInt64(filter.Query.FirmwareId),
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.OpmFirmwareVersionEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmFirmwareVersion_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddOpmFirmwareVersion 新增固件版本
func (s OpmFirmwareVersionService) AddOpmFirmwareVersion(req entitys.OpmFirmwareVersionEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.OpmFirmwareVersion_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Status = 2

	//版本号检查
	fvRes, err := rpc.ClientFirmwareVersionService.Lists(s.Ctx, &protosService.OpmFirmwareVersionListRequest{
		Page:     1,
		PageSize: 1,
		Query: &protosService.OpmFirmwareVersion{
			FirmwareId: iotutil.ToInt64(req.FirmwareId),
			Status:     -1,
		},
	})
	if err != nil {
		return "", err
	}
	if len(fvRes.Data) > 0 {
		if r, _ := iotutil.VerCompare(req.Version, fvRes.Data[0].Version); r != 1 {
			return "", errors.New("上传的版本不能小于或等于最新的版本")
		}
	}

	res, err := rpc.ClientFirmwareVersionService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//面板产品图片
	if req.ProdFilePath != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTOpmFirmwareVersion+"_prod", iotutil.ToString(req.Id), req.ProdFilePath)
		commonGlobal.SetAttachmentStatus(model.TableNameTOpmFirmwareVersion+"_upgrade", iotutil.ToString(req.Id), req.UpgradeFilePath)
	}
	return iotutil.ToString(saveObj.Id), err
}

// UpdateOpmFirmwareVersion 修改固件版本
func (s OpmFirmwareVersionService) UpdateOpmFirmwareVersion(req entitys.OpmFirmwareVersionEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientFirmwareVersionService.Update(s.Ctx, entitys.OpmFirmwareVersion_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	if req.ProdFilePath != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTOpmFirmwareVersion+"_prod", iotutil.ToString(req.Id), req.ProdFilePath)
		commonGlobal.SetAttachmentStatus(model.TableNameTOpmFirmwareVersion+"_upgrade", iotutil.ToString(req.Id), req.UpgradeFilePath)
	}
	return iotutil.ToString(req.Id), err
}

// DeleteOpmFirmwareVersion 删除固件版本
func (s OpmFirmwareVersionService) DeleteOpmFirmwareVersion(req entitys.OpmFirmwareVersionFilter) error {
	rep, err := rpc.ClientFirmwareVersionService.Delete(s.Ctx, &protosService.OpmFirmwareVersion{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}

// SetStatusOpmFirmwareVersion 禁用/启用固件
func (s OpmFirmwareVersionService) SetStatusOpmFirmwareVersion(req entitys.OpmFirmwareVersionFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	if req.Status == nil {
		return errors.New("status not found")
	}
	rep, err := rpc.ClientFirmwareVersionService.UpdateFields(s.Ctx, &protosService.OpmFirmwareVersionUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.OpmFirmwareVersion{
			Id:     iotutil.ToInt64(req.Id),
			Status: *req.Status,
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

// OnShelfOpmFirmwareVersion 上架
func (s OpmFirmwareVersionService) OnShelfOpmFirmwareVersion(req entitys.OpmFirmwareVersionFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	if req.Status == nil {
		return errors.New("status not found")
	}
	rep, err := rpc.ClientFirmwareVersionService.Update(s.Ctx, &protosService.OpmFirmwareVersion{
		Id:              req.Id,
		Status:          1, //上架
		IsMust:          req.IsMust,
		UpgradeMode:     req.UpgradeMode,
		UpgradeFileName: req.UpgradeFileName,
		UpgradeFilePath: req.UpgradeFilePath,
		UpgradeFileSize: iotutil.ToInt32(req.UpgradeFileSize),
		UpgradeFileKey:  req.UpgradeFileKey,
		ProdFilePath:    req.ProdFilePath,
		ProdFileSize:    iotutil.ToInt32(req.ProdFileSize),
		ProdFileKey:     req.ProdFileKey,
		ProdFileName:    req.ProdFileName,
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}
