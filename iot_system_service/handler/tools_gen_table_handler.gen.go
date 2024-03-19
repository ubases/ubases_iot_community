// Code generated by sgen.exe,2022-04-18 19:12:11. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	proto "cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_system_service/service"
)

type ToolsGenTableHandler struct{}

// 创建
func (h *ToolsGenTableHandler) Create(ctx context.Context, req *proto.ToolsGenTable, resp *proto.Response) error {
	s := service.ToolsGenTableSvc{Ctx: ctx}
	_, err := s.CreateToolsGenTable(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *ToolsGenTableHandler) Delete(ctx context.Context, req *proto.ToolsGenTable, resp *proto.Response) error {
	s := service.ToolsGenTableSvc{Ctx: ctx}
	_, err := s.DeleteToolsGenTable(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *ToolsGenTableHandler) DeleteById(ctx context.Context, req *proto.ToolsGenTable, resp *proto.Response) error {
	s := service.ToolsGenTableSvc{Ctx: ctx}
	_, err := s.DeleteByIdToolsGenTable(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *ToolsGenTableHandler) DeleteByIds(ctx context.Context, req *proto.ToolsGenTableBatchDeleteRequest, resp *proto.Response) error {
	s := service.ToolsGenTableSvc{Ctx: ctx}
	_, err := s.DeleteByIdsToolsGenTable(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *ToolsGenTableHandler) Update(ctx context.Context, req *proto.ToolsGenTable, resp *proto.Response) error {
	s := service.ToolsGenTableSvc{Ctx: ctx}
	_, err := s.UpdateToolsGenTable(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *ToolsGenTableHandler) UpdateAll(ctx context.Context, req *proto.ToolsGenTable, resp *proto.Response) error {
	s := service.ToolsGenTableSvc{Ctx: ctx}
	_, err := s.UpdateAllToolsGenTable(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *ToolsGenTableHandler) UpdateFields(ctx context.Context, req *proto.ToolsGenTableUpdateFieldsRequest, resp *proto.Response) error {
	s := service.ToolsGenTableSvc{Ctx: ctx}
	_, err := s.UpdateFieldsToolsGenTable(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *ToolsGenTableHandler) Find(ctx context.Context, req *proto.ToolsGenTableFilter, resp *proto.ToolsGenTableResponse) error {
	s := service.ToolsGenTableSvc{Ctx: ctx}
	data, err := s.FindToolsGenTable(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *ToolsGenTableHandler) FindById(ctx context.Context, req *proto.ToolsGenTableFilter, resp *proto.ToolsGenTableResponse) error {
	s := service.ToolsGenTableSvc{Ctx: ctx}
	data, err := s.FindByIdToolsGenTable(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *ToolsGenTableHandler) Lists(ctx context.Context, req *proto.ToolsGenTableListRequest, resp *proto.ToolsGenTableResponse) error {
	s := service.ToolsGenTableSvc{Ctx: ctx}
	data, total, err := s.GetListToolsGenTable(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *ToolsGenTableHandler) SetResponse(resp *proto.ToolsGenTableResponse, data *proto.ToolsGenTable, err error) {
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

func (h *ToolsGenTableHandler) SetPageResponse(resp *proto.ToolsGenTableResponse, list []*proto.ToolsGenTable, total int64, err error) {
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
