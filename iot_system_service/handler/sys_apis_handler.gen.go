// Code generated by sgen.exe,2022-04-18 19:12:07. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	proto "cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_system_service/service"
)

type SysApisHandler struct{}

// 创建
func (h *SysApisHandler) Create(ctx context.Context, req *proto.SysApis, resp *proto.Response) error {
	s := service.SysApisSvc{Ctx: ctx}
	_, err := s.CreateSysApis(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *SysApisHandler) Delete(ctx context.Context, req *proto.SysApis, resp *proto.Response) error {
	s := service.SysApisSvc{Ctx: ctx}
	_, err := s.DeleteSysApis(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *SysApisHandler) DeleteById(ctx context.Context, req *proto.SysApis, resp *proto.Response) error {
	s := service.SysApisSvc{Ctx: ctx}
	_, err := s.DeleteByIdSysApis(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *SysApisHandler) DeleteByIds(ctx context.Context, req *proto.SysApisBatchDeleteRequest, resp *proto.Response) error {
	s := service.SysApisSvc{Ctx: ctx}
	_, err := s.DeleteByIdsSysApis(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *SysApisHandler) Update(ctx context.Context, req *proto.SysApis, resp *proto.Response) error {
	s := service.SysApisSvc{Ctx: ctx}
	_, err := s.UpdateSysApis(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *SysApisHandler) UpdateAll(ctx context.Context, req *proto.SysApis, resp *proto.Response) error {
	s := service.SysApisSvc{Ctx: ctx}
	_, err := s.UpdateAllSysApis(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *SysApisHandler) UpdateFields(ctx context.Context, req *proto.SysApisUpdateFieldsRequest, resp *proto.Response) error {
	s := service.SysApisSvc{Ctx: ctx}
	_, err := s.UpdateFieldsSysApis(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *SysApisHandler) Find(ctx context.Context, req *proto.SysApisFilter, resp *proto.SysApisResponse) error {
	s := service.SysApisSvc{Ctx: ctx}
	data, err := s.FindSysApis(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *SysApisHandler) FindById(ctx context.Context, req *proto.SysApisFilter, resp *proto.SysApisResponse) error {
	s := service.SysApisSvc{Ctx: ctx}
	data, err := s.FindByIdSysApis(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *SysApisHandler) Lists(ctx context.Context, req *proto.SysApisListRequest, resp *proto.SysApisResponse) error {
	s := service.SysApisSvc{Ctx: ctx}
	data, total, err := s.GetListSysApis(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *SysApisHandler) SetResponse(resp *proto.SysApisResponse, data *proto.SysApis, err error) {
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

func (h *SysApisHandler) SetPageResponse(resp *proto.SysApisResponse, list []*proto.SysApis, total int64, err error) {
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
