// Code generated by sgen.exe,2022-05-31 13:46:36. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_basic_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type LangResourcesHandler struct{}

// 创建
func (h *LangResourcesHandler) Create(ctx context.Context, req *proto.LangResources, resp *proto.Response) error {
	s := service.LangResourcesSvc{Ctx: ctx}
	_, err := s.CreateLangResources(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *LangResourcesHandler) Delete(ctx context.Context, req *proto.LangResources, resp *proto.Response) error {
	s := service.LangResourcesSvc{Ctx: ctx}
	_, err := s.DeleteLangResources(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *LangResourcesHandler) DeleteById(ctx context.Context, req *proto.LangResources, resp *proto.Response) error {
	s := service.LangResourcesSvc{Ctx: ctx}
	_, err := s.DeleteByIdLangResources(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *LangResourcesHandler) DeleteByIds(ctx context.Context, req *proto.LangResourcesBatchDeleteRequest, resp *proto.Response) error {
	s := service.LangResourcesSvc{Ctx: ctx}
	_, err := s.DeleteByIdsLangResources(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *LangResourcesHandler) Update(ctx context.Context, req *proto.LangResources, resp *proto.Response) error {
	s := service.LangResourcesSvc{Ctx: ctx}
	_, err := s.UpdateLangResources(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *LangResourcesHandler) UpdateAll(ctx context.Context, req *proto.LangResources, resp *proto.Response) error {
	s := service.LangResourcesSvc{Ctx: ctx}
	_, err := s.UpdateAllLangResources(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *LangResourcesHandler) UpdateFields(ctx context.Context, req *proto.LangResourcesUpdateFieldsRequest, resp *proto.Response) error {
	s := service.LangResourcesSvc{Ctx: ctx}
	_, err := s.UpdateFieldsLangResources(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *LangResourcesHandler) Find(ctx context.Context, req *proto.LangResourcesFilter, resp *proto.LangResourcesResponse) error {
	s := service.LangResourcesSvc{Ctx: ctx}
	data, err := s.FindLangResources(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *LangResourcesHandler) FindById(ctx context.Context, req *proto.LangResourcesFilter, resp *proto.LangResourcesResponse) error {
	s := service.LangResourcesSvc{Ctx: ctx}
	data, err := s.FindByIdLangResources(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *LangResourcesHandler) Lists(ctx context.Context, req *proto.LangResourcesListRequest, resp *proto.LangResourcesResponse) error {
	s := service.LangResourcesSvc{Ctx: ctx}
	data, total, err := s.GetListLangResources(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *LangResourcesHandler) SetResponse(resp *proto.LangResourcesResponse, data *proto.LangResources, err error) {
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

func (h *LangResourcesHandler) SetPageResponse(resp *proto.LangResourcesResponse, list []*proto.LangResources, total int64, err error) {
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
