package services

import (
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"go-micro.dev/v4/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
)

type OpmOtaPkgService struct {
	Ctx context.Context
}

func (s *OpmOtaPkgService) SetContext(ctx context.Context) *OpmOtaPkgService {
	s.Ctx = ctx
	return s
}

// 固件Ota详细
func (s OpmOtaPkgService) GetOpmOtaPkgDetail(id string) (*entitys.OpmOtaPkgEntitys, error) {
	if id == "" {
		return nil, errors.New("id not found")
	}
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientOtaPkgService.FindById(s.Ctx, &protosService.OpmOtaPkgFilter{Id: rid})
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
	return entitys.OpmOtaPkg_pb2e(data), err
}

// QueryOpmOtaPkgList 固件Ota列表
func (s OpmOtaPkgService) QueryOpmOtaPkgList(filter entitys.OpmOtaPkgQuery) ([]*entitys.OpmOtaPkgEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	rep, err := rpc.ClientOtaPkgService.Lists(s.Ctx, &protosService.OpmOtaPkgListRequest{
		Page:      filter.Page,
		PageSize:  filter.Limit,
		SearchKey: filter.SearchKey,
		Query:     entitys.OpmOtaPkgFilter_e2pb(filter.Query),
		OrderDesc: "desc",
		OrderKey:  "updated_at",
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.OpmOtaPkgEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmOtaPkg_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddOpmOtaPkg 新增固件Ota
func (s OpmOtaPkgService) AddOpmOtaPkg(req entitys.OpmOtaPkgEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.OpmOtaPkg_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Status = iotconst.STATUS_RELEASE_PENDING

	proRes, err := rpc.ClientOpmProductService.FindById(s.Ctx, &protosService.OpmProductFilter{Id: req.ProductId})
	if err != nil {
		return "", err
	}
	if proRes.Code != 200 {
		return "", errors.New(proRes.Message)
	}
	if len(proRes.Data) == 0 {
		return "", errors.New("产品Id不存在")
	}
	saveObj.ProductKey = proRes.Data[0].ProductKey

	res, err := rpc.ClientOtaPkgService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// 修改固件Ota
func (s OpmOtaPkgService) UpdateOpmOtaPkg(req entitys.OpmOtaPkgEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientOtaPkgService.Update(s.Ctx, entitys.OpmOtaPkg_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除固件Ota
func (s OpmOtaPkgService) DeleteOpmOtaPkg(req entitys.OpmOtaPkgFilter) error {
	id, _ := iotutil.ToInt64AndErr(req.Id)
	if id == 0 {
		return errors.New("id not found")
	}
	rep, err := rpc.ClientOtaPkgService.Delete(s.Ctx, &protosService.OpmOtaPkg{
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

// SetStatusOpmOtaPkg 禁用/启用固件Ota
func (s OpmOtaPkgService) SetStatusOpmOtaPkg(req entitys.OpmOtaPkgFilter) error {
	id, _ := iotutil.ToInt64AndErr(req.Id)
	if id == 0 {
		return errors.New("固件OTA包编号不能为空")
	}
	if req.Status == 0 {
		return errors.New("状态不能为空")
	}
	rep, err := rpc.ClientOtaPkgService.UpdateFields(s.Ctx, &protosService.OpmOtaPkgUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.OpmOtaPkg{
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

func (s OpmOtaPkgService) checkDeviceCount(req entitys.OtaReleaseRequest) error {
	rep, err := rpc.ClientIotDeviceInfoServer.QueryCount(s.Ctx, &protosService.IotDeviceInfoListRequest{
		Query: &protosService.IotDeviceInfo{
			ProductKey: req.ProductKey,
			ProductId:  req.ProductID,
		},
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return err
	}
	if int64(req.GrayScale) >= rep.Data.ActiveTotal {
		return errors.New("灰度设备总数不得大于当前账号拥有设备数量")
	}
	return nil
}

// 检查设备Id是否有效
func (s *OpmOtaPkgService) checkDeviceIds(req entitys.OtaReleaseRequest) ([]string, error) {
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	if tenantId == "" {
		return nil, errors.New("未获取到当前用户的租户Id")
	}
	rep, err := rpc.ClientIotDeviceServer.Lists(s.Ctx, &protosService.IotDeviceTriadListRequest{
		Query: &protosService.IotDeviceTriad{
			DeviceIds: req.DeviceIds,
			//ProductKey:       req.ProductKey, //产品Key
			//ProductId:        req.ProductID,
			TenantId:         tenantId,
			IsQueryTriadData: true,
			Status:           -1,
			IsTest:           -1,
			UseType:          -1,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, err
	}

	dbDevs := make([]string, 0)
	noInDevIds := make([]string, 0)
	for _, d := range rep.Data {
		//不在平台的设备Id
		if d.ProductId != req.ProductID {
			noInDevIds = append(noInDevIds, d.Did)
		}
		dbDevs = append(dbDevs, d.Did)
	}

	//不存在的设备Id
	missingDevIds := iotutil.FindMissingElements(req.DeviceIds, dbDevs)

	//设备ID不存在平台设备库的：提示设备ID不存在
	//非这个产品的设备ID：提示设备ID不属于该产品
	var errMsg []string = make([]string, 0)
	if len(missingDevIds) > 0 {
		errMsg = append(errMsg, "以下设备Id不存在\r\n（"+strings.Join(missingDevIds, "、")+"）")
	}
	if len(noInDevIds) > 0 {
		errMsg = append(errMsg, "以下备ID不属于该产品\r\n（"+strings.Join(noInDevIds, "、")+"）")
	}
	if len(errMsg) > 0 {
		return iotutil.Union(missingDevIds, noInDevIds), errors.New(strings.Join(errMsg, "\r\n"))
	}
	return missingDevIds, nil
}

// OtaPublish 固件Ota发布
func (s *OpmOtaPkgService) OtaPublish(req entitys.OtaReleaseRequest) ([]string, error) {
	mdids := make([]string, 0)
	if err := req.CheckParams(); err != nil {
		return mdids, err

	}
	if req.ReleaseMode == 2 {
		switch req.GraySetting {
		case 1:
			if req.GraySetting == 0 {
				return mdids, errors.New("灰度设置不能空")
			}
			if req.GrayScale == 0 {
				return mdids, errors.New("灰度比例不能为空")
			}
			if req.GrayScale > 100 || req.GrayScale < 0 {
				return mdids, errors.New("灰度比例值必须为[1-100]")
			}
		case 2:
			if req.GraySetting == 0 {
				return mdids, errors.New("灰度设置不能空")
			}
			if req.GrayScale == 0 {
				return mdids, errors.New("灰度比例不能为空")
			}
			err := s.checkDeviceCount(req)
			if err != nil {
				return mdids, err
			}
		case 3:
			missingDevIds, err := s.checkDeviceIds(req)
			if err != nil {
				return missingDevIds, err
			}
		default:
			return mdids, errors.New("不存在的灰度设置")
		}
	}

	//TODO 临时方案，之后从前端传入productKey
	prodSvc := OpmProductService{Ctx: s.Ctx}
	productInfo, err := prodSvc.GetOpmProductDetail(iotutil.ToString(req.ProductID))
	if err != nil {
		return mdids, err
	}

	otaPkg, err := rpc.ClientOtaPkgService.FindById(s.Ctx, &protosService.OpmOtaPkgFilter{Id: req.OtaPkgId})
	if err != nil {
		return mdids, err
	}
	if otaPkg.Code != 200 {
		return mdids, errors.New(otaPkg.Message)
	}
	otaPkgInfo := otaPkg.Data[0]
	rep, err := rpc.ClientOtaPkgService.SetPublish(s.Ctx, &protosService.SetOtaPublishRequest{
		OtaPkgId: req.OtaPkgId,
		Publish: &protosService.OpmOtaPublish{
			IsGray:               req.ReleaseMode,
			GrayScale:            req.GrayScale,
			GrayType:             req.GraySetting,
			PkgId:                req.OtaPkgId,
			Status:               iotconst.STATUS_RELEASE, //状态[1:已发布,2:待发布,3:已暂停]
			PublishAt:            timestamppb.Now(),
			Version:              otaPkgInfo.Version,
			VersionId:            otaPkgInfo.VersionId,
			TenantId:             otaPkgInfo.TenantId,
			IsAuto:               otaPkgInfo.IsAuto,
			AutoStartAt:          otaPkgInfo.AutoStartAt,
			AutoEndAt:            otaPkgInfo.AutoEndAt,
			SpecifiedVersionMode: otaPkgInfo.SpecifiedVersionMode,
			SpecifiedVersion:     otaPkgInfo.SpecifiedVersion,
			SpecifiedAreaMode:    otaPkgInfo.SpecifiedAreaMode,
			SpecifiedArea:        otaPkgInfo.SpecifiedArea,
			UpgradeMode:          otaPkgInfo.UpgradeMode,
			UpgradeTimeMode:      otaPkgInfo.UpgradeTimeMode,
			UpgradeDesc:          otaPkgInfo.UpgradeDesc,
			UpgradeDescEn:        otaPkgInfo.UpgradeDescEn,
			ProductKey:           productInfo.ProductKey,
			DeviceIds:            req.DeviceIds,
		},
	})
	if err != nil {
		return mdids, err
	}
	if rep.Code != 200 {
		return mdids, errors.New(rep.Message)
	}
	return mdids, nil
}

// OtaPublishStop 暂停固件Ota
func (s OpmOtaPkgService) OtaPublishStop(req entitys.OpmOtaPublishFilter) error {
	if req.Id == 0 {
		return errors.New("发布编号不能为空")
	}
	rep, err := rpc.ClientOtaPkgService.SetPublish(s.Ctx, &protosService.SetOtaPublishRequest{
		Publish: &protosService.OpmOtaPublish{
			Id:        req.Id,
			TenantId:  req.TenantId,
			Status:    iotconst.STATUS_RELEASE_STOP, //状态[1:已发布,2:待发布,3:已暂停]
			PublishAt: timestamppb.Now(),
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

// OtaRecoveryPublish 恢复Ota
func (s OpmOtaPkgService) OtaRecoveryPublish(req entitys.OpmOtaPublishFilter) error {
	if req.Id == 0 {
		return errors.New("发布编号不能为空")
	}
	rep, err := rpc.ClientOtaPublishService.UpdateFields(s.Ctx, &protosService.OpmOtaPublishUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.OpmOtaPublish{
			Id:     iotutil.ToInt64(req.Id),
			Status: iotconst.STATUS_RELEASE,
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

// QueryOpmOtaPublishList 固件Ota发布
func (s OpmOtaPkgService) QueryOpmOtaPublishList(filter entitys.OpmOtaPublishQuery) ([]*entitys.OpmOtaPublishEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	rep, err := rpc.ClientOtaPublishService.Lists(s.Ctx, &protosService.OpmOtaPublishListRequest{
		Page:     filter.Page,
		PageSize: filter.Limit,
		Query:    entitys.OpmOtaPublishFilter_e2pb(filter.Query),
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.OpmOtaPublishEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmOtaPublish_pb2e(item))
	}
	return resultList, rep.Total, err
}

// QueryOtaResultList 查询OTA结果列表
func (s OpmOtaPkgService) QueryOtaResultList(deviceId, version string, isGray int32, publishId int64, status, result int32, area string, page, limit int64) (map[string]interface{}, error) {
	var queryStatus int32 = 0
	var queryStatusList []int32 = make([]int32, 0)
	if status == 1 {
		//未升级
		queryStatus = 1
	}
	if result != 0 {
		switch result {
		case 1: //下载中
			queryStatus = 2
		case 2: //安装中
			queryStatus = 3
		case 3: //成功
			queryStatus = 4
		case 4: //失败
			queryStatus = 5
		}
	} else {
		if status == 2 {
			queryStatusList = append(queryStatusList, 2, 3, 4, 5)
		}
	}
	rep, err := rpc.ClientOtaUpgradeRecordService.Lists(s.Ctx, &protosService.IotOtaUpgradeRecordListRequest{
		Page:     page,
		PageSize: limit,
		Query: &protosService.IotOtaUpgradeRecord{
			PublishId:  publishId,
			DeviceId:   deviceId,
			Version:    version,
			IsGray:     isGray,
			Status:     queryStatus,
			StatusList: queryStatusList,
			Area:       area,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	var resultList = []*entitys.IotOtaUpgradeRecordEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.IotOtaUpgradeRecord_pb2e(item))
	}
	return map[string]interface{}{
		"deviceList":   resultList,
		"total":        rep.Total,
		"successTotal": rep.SuccessTotal,
		"failTotal":    rep.FailTotal,
		"page":         page,
	}, nil
}

// QueryOtaVersions 待升级的固件版本
func (s OpmOtaPkgService) QueryOtaVersions(productId, firmwareId int64) ([]string, error) {
	//var resVersions = []string{
	//	"1.0.1",
	//	"1.0.2",
	//	"1.0.3",
	//	"1.0.4",
	//}
	otaVersion, err := rpc.ClientOtaPkgService.GetProductOtaVersion(context.Background(), &protosService.ProductOtaVersionRequest{
		ProductId:  productId,
		FirmwareId: firmwareId,
	})
	if err != nil {
		return nil, err
	}
	if otaVersion.Code != 200 {
		return nil, errors.New(otaVersion.Message)
	}
	return otaVersion.Versions, nil
}

// QueryOtaAreas 获取可升级的区域
func (s OpmOtaPkgService) QueryOtaAreas(productId, firmwareId int64) ([]*iotstruct.DropdownItem, error) {
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	//var resAreas = []map[string]string{
	//	{
	//		"code": "changsha",
	//		"name": "长沙市",
	//	},
	//	{
	//		"code": "shenzhen",
	//		"name": "深圳市",
	//	},
	//}
	otaVersion, err := rpc.IotDeviceHomeService.QueryDeviceAreas(context.Background(), &protosService.DeviceAreaRequest{
		ProductId: productId,
		TenantId:  tenantId,
	})
	if err != nil {
		return nil, err
	}
	if otaVersion.Code != 200 {
		return nil, errors.New(otaVersion.Message)
	}

	res := make([]*iotstruct.DropdownItem, 0)
	for k, v := range otaVersion.Areas {
		if k == "" {
			continue
		}
		res = append(res, &iotstruct.DropdownItem{Code: k, Count: v, Name: k})
	}
	return res, nil
}
