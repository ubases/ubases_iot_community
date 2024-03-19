// Code generated by sgen.exe,2022-07-14 15:09:59. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_app_oem_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OemAppDocRelationHandler struct{}

// 创建
func (h *OemAppDocRelationHandler) Create(ctx context.Context, req *proto.OemAppDocRelation, resp *proto.Response) error {
	s := service.OemAppDocRelationSvc{Ctx: ctx}
	_, err := s.CreateOemAppDocRelation(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *OemAppDocRelationHandler) Delete(ctx context.Context, req *proto.OemAppDocRelation, resp *proto.Response) error {
	s := service.OemAppDocRelationSvc{Ctx: ctx}
	_, err := s.DeleteOemAppDocRelation(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *OemAppDocRelationHandler) DeleteById(ctx context.Context, req *proto.OemAppDocRelation, resp *proto.Response) error {
	s := service.OemAppDocRelationSvc{Ctx: ctx}
	_, err := s.DeleteByIdOemAppDocRelation(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *OemAppDocRelationHandler) DeleteByIds(ctx context.Context, req *proto.OemAppDocRelationBatchDeleteRequest, resp *proto.Response) error {
	s := service.OemAppDocRelationSvc{Ctx: ctx}
	_, err := s.DeleteByIdsOemAppDocRelation(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *OemAppDocRelationHandler) Update(ctx context.Context, req *proto.OemAppDocRelation, resp *proto.Response) error {
	s := service.OemAppDocRelationSvc{Ctx: ctx}
	_, err := s.UpdateOemAppDocRelation(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *OemAppDocRelationHandler) UpdateAll(ctx context.Context, req *proto.OemAppDocRelation, resp *proto.Response) error {
	s := service.OemAppDocRelationSvc{Ctx: ctx}
	_, err := s.UpdateAllOemAppDocRelation(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *OemAppDocRelationHandler) UpdateFields(ctx context.Context, req *proto.OemAppDocRelationUpdateFieldsRequest, resp *proto.Response) error {
	s := service.OemAppDocRelationSvc{Ctx: ctx}
	_, err := s.UpdateFieldsOemAppDocRelation(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *OemAppDocRelationHandler) Find(ctx context.Context, req *proto.OemAppDocRelationFilter, resp *proto.OemAppDocRelationResponse) error {
	s := service.OemAppDocRelationSvc{Ctx: ctx}
	data, err := s.FindOemAppDocRelation(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *OemAppDocRelationHandler) FindById(ctx context.Context, req *proto.OemAppDocRelationFilter, resp *proto.OemAppDocRelationResponse) error {
	s := service.OemAppDocRelationSvc{Ctx: ctx}
	data, err := s.FindByIdOemAppDocRelation(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *OemAppDocRelationHandler) Lists(ctx context.Context, req *proto.OemAppDocRelationListRequest, resp *proto.OemAppDocRelationResponse) error {
	s := service.OemAppDocRelationSvc{Ctx: ctx}
	data, total, err := s.GetListOemAppDocRelation(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *OemAppDocRelationHandler) SetResponse(resp *proto.OemAppDocRelationResponse, data *proto.OemAppDocRelation, err error) {
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

func (h *OemAppDocRelationHandler) SetPageResponse(resp *proto.OemAppDocRelationResponse, list []*proto.OemAppDocRelation, total int64, err error) {
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
