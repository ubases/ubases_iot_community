// Code generated by sgen.exe,2022-05-07 16:41:11. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_product_service/service"
	"context"
	"errors"
	"fmt"

	proto "cloud_platform/iot_proto/protos/protosService"
)

type OpmProductHandler struct{}

func (h *OpmProductHandler) FindByAllDetails(ctx context.Context, primarykey *proto.OpmProductPrimarykey, resp *proto.OpmProductAllDetails) error {
	s := service.OpmProductSvc{Ctx: ctx}
	res, err := s.FindProductAllDetaialById(primarykey)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.Product = res.Product
		resp.ThingModes = res.ThingModes
		resp.Module = res.Module
		resp.ControlPanel = res.ControlPanel
		resp.BaseProduct = res.BaseProduct
		resp.CustomFirmwares = res.CustomFirmwares
	}
	return nil
}

func (h *OpmProductHandler) ControlPanelsLists(ctx context.Context, req *proto.ControlPanelIdsRequest, resp *proto.PmControlPanelsVoResponse) error {
	s := service.OpmProductSvc{Ctx: ctx}
	ret, err := s.ControlPanelsLists(req)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.Data = ret.Data
	}
	return nil
}

func (h *OpmProductHandler) ModuleLists(ctx context.Context, req *proto.ModuleIdsRequest, resp *proto.PmModuleVoResponse) error {
	s := service.OpmProductSvc{Ctx: ctx}
	ret, err := s.ModuleLists(req)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.Data = ret.Data
	}
	return nil
}

// 创建
func (h *OpmProductHandler) Create(ctx context.Context, req *proto.OpmProduct, resp *proto.Response) error {
	s := service.OpmProductSvc{Ctx: ctx}
	ret, err := s.CreateOpmProduct(req, false)
	service.SetResponse(resp, err)
	if ret != nil && err == nil {
		resp.Data = ret.Id
		//service.GetJsPublisherMgr().PushData(&service.NatsPubData{
		//	Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
		//	Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, req.ProductKey, "name", req.Name, req.NameEn),
		//})

		iotnatsjs.GetJsClientPub().PushData(&iotnatsjs.NatsPubData{
			Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
			Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, req.ProductKey, "name", req.Name, req.NameEn),
		})
	}
	return nil
}

// 创建
func (h *OpmProductHandler) CreateDemoProduct(ctx context.Context, req *proto.CreateDemoProductRequest, resp *proto.CreateDemoProductResponse) error {
	pmSvc := TPmProductHandler{}
	baseInfo, err := pmSvc.GetBaseProductInfo(req.BaseProductId)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
		return nil
	}
	//用户Demo产品数据
	reqObj := &proto.OpmProduct{
		ImageUrl:         baseInfo.ImageUrl,
		PowerConsumeType: baseInfo.PowerConsumeType,
		NetworkType:      baseInfo.NetworkType,
		Name:             fmt.Sprintf("%v_demo", baseInfo.Name),
		Model:            iotutil.GetProductKeyRandomString(),
		ProductTypeId:    baseInfo.ProductTypeId,
		ProductId:        req.BaseProductId,
		WifiFlag:         baseInfo.WifiFlag,
		ControlPanelId:   req.ControlPanelId,
		IsDemoProduct:    1,
		TenantId:         req.TenantId,
		Status:           2,
	}
	//查询产品类型信息
	s := service.OpmProductSvc{Ctx: ctx}
	ret, err := s.CreateOpmProduct(reqObj, true)
	resp.Code = SUCCESS
	resp.Message = "success"

	if ret != nil && err == nil {
		resp.Data = ret
		//service.GetJsPublisherMgr().PushData(&service.NatsPubData{
		//	Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
		//	Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, reqObj.ProductKey, "name", reqObj.Name, reqObj.NameEn),
		//})
		iotnatsjs.GetJsClientPub().PushData(&iotnatsjs.NatsPubData{
			Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
			Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, reqObj.ProductKey, "name", reqObj.Name, reqObj.NameEn),
		})
	}
	return nil
}

// 匹配多条件删除
func (h *OpmProductHandler) Delete(ctx context.Context, req *proto.OpmProduct, resp *proto.Response) error {
	s := service.OpmProductSvc{Ctx: ctx}
	_, err := s.DeleteOpmProduct(req)
	service.SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *OpmProductHandler) DeleteById(ctx context.Context, req *proto.OpmProduct, resp *proto.Response) error {
	s := service.OpmProductSvc{Ctx: ctx}
	_, err := s.DeleteByIdOpmProduct(req)
	service.SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *OpmProductHandler) DeleteByIds(ctx context.Context, req *proto.OpmProductBatchDeleteRequest, resp *proto.Response) error {
	s := service.OpmProductSvc{Ctx: ctx}
	_, err := s.DeleteByIdsOpmProduct(req)
	service.SetResponse(resp, err)
	return nil
}

// 更新
func (h *OpmProductHandler) Update(ctx context.Context, req *proto.OpmProduct, resp *proto.Response) error {
	s := service.OpmProductSvc{Ctx: ctx}
	_, err := s.UpdateOpmProduct(req)
	service.SetResponse(resp, err)
	//if err == nil && req.Name != "" {
	//	data, err := s.FindByIdOpmProduct(&proto.OpmProductFilter{Id: req.Id})
	//	if err == nil {
	//		service.GetJsPublisherMgr().PushData(&service.NatsPubData{
	//			Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
	//			Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, data.ProductKey, "name", req.Name, req.NameEn),
	//		})
	//	}
	//}
	return nil
}

// 根据主键更新所有字段
func (h *OpmProductHandler) UpdateAll(ctx context.Context, req *proto.OpmProduct, resp *proto.Response) error {
	s := service.OpmProductSvc{Ctx: ctx}
	_, err := s.UpdateAllOpmProduct(req)
	service.SetResponse(resp, err)
	//if err == nil && req.Name != "" {
	//	data, err := s.FindByIdOpmProduct(&proto.OpmProductFilter{Id: req.Id})
	//	if err == nil {
	//		service.GetJsPublisherMgr().PushData(&service.NatsPubData{
	//			Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
	//			Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, data.ProductKey, "name", req.Name, req.NameEn),
	//		})
	//	}
	//}
	return nil
}

// 根据主键更新指定列
func (h *OpmProductHandler) UpdateFields(ctx context.Context, req *proto.OpmProductUpdateFieldsRequest, resp *proto.Response) error {
	s := service.OpmProductSvc{Ctx: ctx}
	_, err := s.UpdateFieldsOpmProduct(req)
	service.SetResponse(resp, err)
	//if err == nil {
	//	if req.Data.Name != "" {
	//		data, err := s.FindByIdOpmProduct(&proto.OpmProductFilter{Id: req.Data.Id})
	//		if err == nil {
	//			service.GetJsPublisherMgr().PushData(&service.NatsPubData{
	//				Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
	//				Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, data.ProductKey, "name", req.Data.Name, req.Data.NameEn),
	//			})
	//		}
	//	}
	//}
	return nil
}

// 多条件查找，返回单条数据
func (h *OpmProductHandler) Find(ctx context.Context, req *proto.OpmProductFilter, resp *proto.OpmProductResponse) error {
	s := service.OpmProductSvc{Ctx: ctx}
	data, err := s.FindOpmProduct(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *OpmProductHandler) FindById(ctx context.Context, req *proto.OpmProductFilter, resp *proto.OpmProductResponse) error {
	s := service.OpmProductSvc{Ctx: ctx}
	data, err := s.FindByIdOpmProduct(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *OpmProductHandler) Lists(ctx context.Context, req *proto.OpmProductListRequest, resp *proto.OpmProductResponse) error {
	s := service.OpmProductSvc{Ctx: ctx}
	data, total, err := s.GetListOpmProduct(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *OpmProductHandler) SetResponse(resp *proto.OpmProductResponse, data *proto.OpmProduct, err error) {
	if err != nil {
		resp.Code = service.ERROR
		resp.Message = err.Error()
		resp.Total = 0
	} else {
		resp.Code = service.SUCCESS
		resp.Message = "success"
		if data != nil {
			resp.Total = 1
			resp.Data = append(resp.Data, data)
		}
	}
}

func (h *OpmProductHandler) SetPageResponse(resp *proto.OpmProductResponse, list []*proto.OpmProduct, total int64, err error) {
	if err != nil {
		resp.Code = service.ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = service.SUCCESS
		resp.Message = "success"
		resp.Total = total
		resp.Data = list
	}
}

// App查找，支持分页，可返回多条数据
func (h *OpmProductHandler) AppLists(ctx context.Context, req *proto.AppOpmProductListRequest, resp *proto.OpmProductResponse) error {
	tenantId, err := service.CheckTenantId(ctx)
	if err != nil {
		return errors.New("tenantId is null")
	}

	s := service.OpmProductSvc{Ctx: ctx}
	data, total, err := s.GetListAppOpmProduct(req, tenantId)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

// App,获取控制面板url和md5
func (h *OpmProductHandler) ControlPanelsUrlAndMd5(ctx context.Context, request *proto.ControlPanelsUrlAndMd5Request, response *proto.ControlPanelsUrlAndMd5Response) error {
	s := service.OpmProductSvc{Ctx: ctx}
	result, err := s.ControlPanelsUrlAndMd5(request)
	if err != nil {
		response.Code = service.ERROR
		response.Message = err.Error()
	} else {
		response.Code = service.SUCCESS
		response.Message = "success"
		response.Url = result.Url
		response.ControlpageMd5 = result.ControlpageMd5
	}
	return nil
}

func (h *OpmProductHandler) ListsByProductIds(ctx context.Context, req *proto.ListsByProductIdsRequest, resp *proto.OpmProductResponse) error {
	s := service.OpmProductSvc{Ctx: ctx}
	data, total, err := s.ListsByProductIds(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

// 重新设置产品物模型
func (h *OpmProductHandler) ResetOpmProductThingsModel(ctx context.Context, req *proto.OpmProduct, resp *proto.Response) error {
	s := service.OpmProductSvc{Ctx: ctx}
	err := s.ResetOpmProductThingsModel(req)
	service.SetResponse(resp, err)
	return nil
}

func (h *OpmProductHandler) PanelListsByProductIds(ctx context.Context, req *proto.ListsByProductIdsRequest, resp *proto.OpmProductResponse) error {
	s := service.OpmProductSvc{Ctx: ctx}
	data, total, err := s.PanelListsByProductIds(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *OpmProductHandler) MergeProductThingsModel(ctx context.Context, req *proto.ListsByProductIdsRequest, resp *proto.OpmThingModelByProductResponse) error {
	s := service.OpmProductSvc{Ctx: ctx}
	productIds := iotutil.Int64ArrayToString(req.ProductIds)
	data, err := s.MergeProductThingsModel(productIds, req.ProductKeys)
	if err != nil {
		resp.Code = service.ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = service.SUCCESS
		resp.Data = &proto.OpmThingModelAllList{
			Properties: data,
		}
	}
	return nil
}

//获取指定产品的面板信息
func (h *OpmProductHandler) GetProductPanelInfo(ctx context.Context, req *proto.ListsByProductIdsRequest, resp *proto.ProductPanelInfoResponse) error {
	s := service.OpmProductSvc{Ctx: ctx}
	data, err := s.PanelInfosByProducts(req)
	if err != nil {
		resp.Code = service.ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = service.SUCCESS
		resp.Message = "success"
		resp.Data = data
	}
	return nil
}
