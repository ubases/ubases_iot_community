// Code generated by sgen.exe,2022-07-25 09:29:04. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	proto "cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_system_service/service"
)

type SysAppEntrySetingHandler struct{}

// 创建
func (h *SysAppEntrySetingHandler) Create(ctx context.Context, req *proto.SysAppEntrySeting, resp *proto.Response) error {
	s := service.SysAppEntrySetingSvc{Ctx: ctx}
	_, err := s.CreateSysAppEntrySeting(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *SysAppEntrySetingHandler) Delete(ctx context.Context, req *proto.SysAppEntrySeting, resp *proto.Response) error {
	s := service.SysAppEntrySetingSvc{Ctx: ctx}
	_, err := s.DeleteSysAppEntrySeting(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *SysAppEntrySetingHandler) DeleteById(ctx context.Context, req *proto.SysAppEntrySeting, resp *proto.Response) error {
	s := service.SysAppEntrySetingSvc{Ctx: ctx}
	_, err := s.DeleteByIdSysAppEntrySeting(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *SysAppEntrySetingHandler) DeleteByIds(ctx context.Context, req *proto.SysAppEntrySetingBatchDeleteRequest, resp *proto.Response) error {
	s := service.SysAppEntrySetingSvc{Ctx: ctx}
	_, err := s.DeleteByIdsSysAppEntrySeting(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *SysAppEntrySetingHandler) Update(ctx context.Context, req *proto.SysAppEntrySeting, resp *proto.Response) error {
	s := service.SysAppEntrySetingSvc{Ctx: ctx}
	_, err := s.UpdateSysAppEntrySeting(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *SysAppEntrySetingHandler) UpdateAll(ctx context.Context, req *proto.SysAppEntrySeting, resp *proto.Response) error {
	s := service.SysAppEntrySetingSvc{Ctx: ctx}
	_, err := s.UpdateAllSysAppEntrySeting(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *SysAppEntrySetingHandler) UpdateFields(ctx context.Context, req *proto.SysAppEntrySetingUpdateFieldsRequest, resp *proto.Response) error {
	s := service.SysAppEntrySetingSvc{Ctx: ctx}
	_, err := s.UpdateFieldsSysAppEntrySeting(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *SysAppEntrySetingHandler) Find(ctx context.Context, req *proto.SysAppEntrySetingFilter, resp *proto.SysAppEntrySetingResponse) error {
	s := service.SysAppEntrySetingSvc{Ctx: ctx}
	data, err := s.FindSysAppEntrySeting(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *SysAppEntrySetingHandler) FindById(ctx context.Context, req *proto.SysAppEntrySetingFilter, resp *proto.SysAppEntrySetingResponse) error {
	s := service.SysAppEntrySetingSvc{Ctx: ctx}
	data, err := s.FindByIdSysAppEntrySeting(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *SysAppEntrySetingHandler) Lists(ctx context.Context, req *proto.SysAppEntrySetingListRequest, resp *proto.SysAppEntrySetingResponse) error {
	s := service.SysAppEntrySetingSvc{Ctx: ctx}
	data, total, err := s.GetListSysAppEntrySeting(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *SysAppEntrySetingHandler) SetResponse(resp *proto.SysAppEntrySetingResponse, data *proto.SysAppEntrySeting, err error) {
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

func (h *SysAppEntrySetingHandler) SetPageResponse(resp *proto.SysAppEntrySetingResponse, list []*proto.SysAppEntrySeting, total int64, err error) {
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

// 创建
func (h *SysAppEntrySetingHandler) CreateBatch(ctx context.Context, req *proto.SysAppEntrySetingBatchRequest, resp *proto.Response) error {
	s := service.SysAppEntrySetingSvc{Ctx: ctx}
	_, err := s.CreateSysAppEntrySetingBatch(req)
	SetResponse(resp, err)
	return nil
}
