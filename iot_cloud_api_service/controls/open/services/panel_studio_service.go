package services

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type OpmPanelService struct {
	Ctx context.Context
}

func (s OpmPanelService) SetContext(ctx context.Context) OpmPanelService {
	s.Ctx = ctx
	return s
}

// 面板详细
func (s OpmPanelService) GetOpmPanelDetail(id string) (*entitys.OpmPanelEntitys, error) {
	rid, err := iotutil.ToInt64AndErr(id)
	if err != nil {
		return nil, err
	}
	req, err := rpc.ClientOpmPanelService.FindById(s.Ctx, &protosService.OpmPanelFilter{Id: rid})
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
	return entitys.OpmPanel_pb2e(data), err
}

// QueryOpmPanelList 面板列表
func (s OpmPanelService) QueryOpmPanelList(filter entitys.OpmPanelQuery) ([]*entitys.OpmPanelEntitys, int64, error) {
	var queryObj = filter.OpmPanelQuery_e2pb()
	rep, err := rpc.ClientOpmPanelService.Lists(s.Ctx, &protosService.OpmPanelListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		SearchKey: filter.SearchKey,
		OrderDesc: "desc",
		OrderKey:  "updated_at",
		Query:     queryObj,
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.OpmPanelEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmPanel_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddOpmPanel 新增面板
func (s OpmPanelService) AddOpmPanel(req entitys.OpmPanelEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.OpmPanel_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Status = 2 // 面板模组新增，默认为禁用
	res, err := rpc.ClientOpmPanelService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//services.SetDefaultTranslate(context.Background(), "t_pm_firmware", res.Data, "name", req.Name, req.NameEn)
	return iotutil.ToString(saveObj.Id), err
}

// 修改面板
func (s OpmPanelService) UpdateOpmPanel(req entitys.OpmPanelEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientOpmPanelService.UpdateAll(s.Ctx, entitys.OpmPanel_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//services.SetDefaultTranslate(context.Background(), "t_pm_firmware", req.Id, "name", req.Name, req.NameEn)
	return iotutil.ToString(req.Id), err
}

func (s OpmPanelService) UpdateOpmPanelStudio(req entitys.OpmPanelEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	req2pro := entitys.OpmPanel_e2pb(&req)
	zhLangs, err := cached.GetLangTranslate("zh")
	enLangs, err := cached.GetLangTranslate("en")
	req2pro.Zhs = zhLangs
	req2pro.Ens = enLangs
	res, err := rpc.ClientOpmPanelService.UpdateEditStudio(s.Ctx, req2pro)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除面板
func (s OpmPanelService) DeleteOpmPanel(req entitys.OpmPanelFilter) error {
	rep, err := rpc.ClientOpmPanelService.Delete(s.Ctx, &protosService.OpmPanel{
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

// SetStatusOpmPanel 禁用/启用面板
func (s OpmPanelService) SetStatusOpmPanel(req entitys.OpmPanelFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	if req.Status == 0 {
		return errors.New("status not found")
	}
	//面板信息
	panelInfo, err := s.GetOpmPanelDetail(iotutil.ToString(req.Id))
	if err != nil {
		return err
	}
	//如果是嵌入式面板，则直接编辑状态
	if panelInfo.PanelType == 2 {
		rep, err := rpc.ClientOpmPanelService.UpdateStatus(context.Background(), &protosService.OpmPanel{
			Id:     iotutil.ToInt64(req.Id),
			Status: req.Status,
		})
		if err != nil {
			return err
		}
		if rep.Code != 200 {
			return errors.New(rep.Message)
		}
	} else {
		//如果是执行发布，将发布状态设置为发布中
		//if req.Status == 3 {
		//	req.Status = 4 //发布中
		//}
		rep, err := rpc.ClientOpmPanelService.UpdateStatus(context.Background(), &protosService.OpmPanel{
			Id:     iotutil.ToInt64(req.Id),
			Status: req.Status,
		})
		if err != nil {
			return err
		}
		if rep.Code != 200 {
			return errors.New(rep.Message)
		}
	}
	return nil
}

// ClearProductPanelLang 清理产品的面板翻译
func ClearProductPanelLang(productId int64, productKey string) {
	defer iotutil.PanicHandler()
	if productKey != "" {
		cachedKey := persist.GetRedisKey(iotconst.APP_PRODUCT_PANEL_LANG, productKey)
		iotredis.GetClient().Del(context.Background(), cachedKey)
		return
	}
	proInfo, err := rpc.ClientOpmProductService.FindById(context.Background(), &protosService.OpmProductFilter{Id: productId})
	if err == nil {
		cachedKey := persist.GetRedisKey(iotconst.APP_PRODUCT_PANEL_LANG, proInfo.Data[0].ProductKey)
		iotredis.GetClient().Del(context.Background(), cachedKey)
	}
}
