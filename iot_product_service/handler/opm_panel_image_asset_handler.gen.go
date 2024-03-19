// Code generated by sgen,2023-09-26 13:54:18. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_product_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OpmPanelImageAssetHandler struct{}

//创建
func (h *OpmPanelImageAssetHandler) Create(ctx context.Context, req *proto.OpmPanelImageAsset, resp *proto.Response) error {
	s := service.OpmPanelImageAssetSvc{Ctx: ctx}
	rsp, err := s.CreateOpmPanelImageAsset(req)
	SetResponse(resp, err)
	if rsp != nil {
		resp.Data = rsp.Id
	}
	return nil
}

//匹配多条件删除
func (h *OpmPanelImageAssetHandler) Delete(ctx context.Context, req *proto.OpmPanelImageAsset, resp *proto.Response) error {
	s := service.OpmPanelImageAssetSvc{Ctx: ctx}
	_, err := s.DeleteOpmPanelImageAsset(req)
	SetResponse(resp, err)
	return nil
}

//匹配ID删除
func (h *OpmPanelImageAssetHandler) DeleteById(ctx context.Context, req *proto.OpmPanelImageAsset, resp *proto.Response) error {
	s := service.OpmPanelImageAssetSvc{Ctx: ctx}
	_, err := s.DeleteByIdOpmPanelImageAsset(req)
	SetResponse(resp, err)
	return nil
}

//根据主键批量删除
func (h *OpmPanelImageAssetHandler) DeleteByIds(ctx context.Context, req *proto.OpmPanelImageAssetBatchDeleteRequest, resp *proto.Response) error {
	s := service.OpmPanelImageAssetSvc{Ctx: ctx}
	_, err := s.DeleteByIdsOpmPanelImageAsset(req)
	SetResponse(resp, err)
	return nil
}

//更新
func (h *OpmPanelImageAssetHandler) Update(ctx context.Context, req *proto.OpmPanelImageAsset, resp *proto.Response) error {
	s := service.OpmPanelImageAssetSvc{Ctx: ctx}
	_, err := s.UpdateOpmPanelImageAsset(req)
	SetResponse(resp, err)
	return nil
}

//根据主键更新所有字段
func (h *OpmPanelImageAssetHandler) UpdateAll(ctx context.Context, req *proto.OpmPanelImageAsset, resp *proto.Response) error {
	s := service.OpmPanelImageAssetSvc{Ctx: ctx}
	_, err := s.UpdateAllOpmPanelImageAsset(req)
	SetResponse(resp, err)
	return nil
}

//根据主键更新指定列
func (h *OpmPanelImageAssetHandler) UpdateFields(ctx context.Context, req *proto.OpmPanelImageAssetUpdateFieldsRequest, resp *proto.Response) error {
	s := service.OpmPanelImageAssetSvc{Ctx: ctx}
	_, err := s.UpdateFieldsOpmPanelImageAsset(req)
	SetResponse(resp, err)
	return nil
}

//多条件查找，返回单条数据
func (h *OpmPanelImageAssetHandler) Find(ctx context.Context, req *proto.OpmPanelImageAssetFilter, resp *proto.OpmPanelImageAssetResponse) error {
	s := service.OpmPanelImageAssetSvc{Ctx: ctx}
	data, err := s.FindOpmPanelImageAsset(req)
	h.SetResponse(resp, data, err)
	return nil
}

//根据ID查找，返回单条数据
func (h *OpmPanelImageAssetHandler) FindById(ctx context.Context, req *proto.OpmPanelImageAssetFilter, resp *proto.OpmPanelImageAssetResponse) error {
	s := service.OpmPanelImageAssetSvc{Ctx: ctx}
	data, err := s.FindByIdOpmPanelImageAsset(req)
	h.SetResponse(resp, data, err)
	return nil
}

//查找，支持分页，可返回多条数据
func (h *OpmPanelImageAssetHandler) Lists(ctx context.Context, req *proto.OpmPanelImageAssetListRequest, resp *proto.OpmPanelImageAssetResponse) error {
	s := service.OpmPanelImageAssetSvc{Ctx: ctx}
	data, total, err := s.GetListOpmPanelImageAsset(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *OpmPanelImageAssetHandler) SetResponse(resp *proto.OpmPanelImageAssetResponse, data *proto.OpmPanelImageAsset, err error) {
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
		resp.Total = 0
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		if data != nil {
			resp.Total = 1
			resp.Data = append(resp.Data, data)
		}
	}
}

func (h *OpmPanelImageAssetHandler) SetPageResponse(resp *proto.OpmPanelImageAssetResponse, list []*proto.OpmPanelImageAsset, total int64, err error) {
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.Total = total
		resp.Data = list
	}
}
