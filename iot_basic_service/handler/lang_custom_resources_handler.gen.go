// Code generated by sgen.exe,2022-05-31 13:46:36. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_basic_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type LangCustomResourcesHandler struct{}

func (h *LangCustomResourcesHandler) CreateOpRecord(ctx context.Context, record *proto.CustomerResourceRecord, response *proto.Response) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	err := s.CreateRecord(record)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
	}
	return nil
}

// 创建
func (h *LangCustomResourcesHandler) ResourceUseRecord(ctx context.Context, req *proto.ResourceOperationRecordRequest, resp *proto.ResourceOperationRecordResponse) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	i, e, err := s.ResourceUseRecord(req)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.ImportCount = i
		resp.ExportCount = e
	}
	return nil
}

// 创建
func (h *LangCustomResourcesHandler) ImportCreate(ctx context.Context, req *proto.ImportLangCustomResource, resp *proto.Response) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	err := s.ImportLangCustomResources(req)
	SetResponse(resp, err)
	return nil
}

// 批量新增资源 BatchSaveCustomResources，通过BelongId和BelongType
func (h *LangCustomResourcesHandler) BatchSaveCustomResources(ctx context.Context, req *proto.BatchCustomResourcesRequest, resp *proto.Response) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	err := s.BatchSaveCustomResources(req)
	SetResponse(resp, err)
	return nil
}

// 创建
func (h *LangCustomResourcesHandler) Create(ctx context.Context, req *proto.LangCustomResources, resp *proto.Response) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	_, err := s.CreateLangCustomResources(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *LangCustomResourcesHandler) Delete(ctx context.Context, req *proto.LangCustomResources, resp *proto.Response) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	_, err := s.DeleteLangCustomResources(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *LangCustomResourcesHandler) DeleteById(ctx context.Context, req *proto.LangCustomResources, resp *proto.Response) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	_, err := s.DeleteByIdLangCustomResources(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *LangCustomResourcesHandler) DeleteByIds(ctx context.Context, req *proto.LangCustomResourcesBatchDeleteRequest, resp *proto.Response) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	_, err := s.DeleteByIdsLangCustomResources(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *LangCustomResourcesHandler) Update(ctx context.Context, req *proto.LangCustomResources, resp *proto.Response) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	_, err := s.UpdateLangCustomResources(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *LangCustomResourcesHandler) UpdateAll(ctx context.Context, req *proto.LangCustomResources, resp *proto.Response) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	_, err := s.UpdateAllLangCustomResources(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *LangCustomResourcesHandler) UpdateFields(ctx context.Context, req *proto.LangCustomResourcesUpdateFieldsRequest, resp *proto.Response) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	_, err := s.UpdateFieldsLangCustomResources(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *LangCustomResourcesHandler) Find(ctx context.Context, req *proto.LangCustomResourcesFilter, resp *proto.LangCustomResourcesResponse) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	data, err := s.FindLangCustomResources(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *LangCustomResourcesHandler) FindById(ctx context.Context, req *proto.LangCustomResourcesFilter, resp *proto.LangCustomResourcesResponse) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	data, err := s.FindByIdLangCustomResources(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *LangCustomResourcesHandler) Lists(ctx context.Context, req *proto.LangCustomResourcesListRequest, resp *proto.LangCustomResourcesResponse) error {
	s := service.LangCustomResourcesSvc{Ctx: ctx}
	data, total, err := s.GetListLangCustomResources(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *LangCustomResourcesHandler) SetResponse(resp *proto.LangCustomResourcesResponse, data *proto.LangCustomResources, err error) {
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

func (h *LangCustomResourcesHandler) SetPageResponse(resp *proto.LangCustomResourcesResponse, list []*proto.LangCustomResources, total int64, err error) {
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
