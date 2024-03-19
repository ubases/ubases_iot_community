// Code generated by sgen.exe,2022-04-18 19:12:09. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	proto "cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_system_service/service"
)

type SysPostHandler struct{}

// 创建
func (h *SysPostHandler) Create(ctx context.Context, req *proto.SysPost, resp *proto.Response) error {
	s := service.SysPostSvc{Ctx: ctx}
	_, err := s.CreateSysPost(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *SysPostHandler) Delete(ctx context.Context, req *proto.SysPost, resp *proto.Response) error {
	s := service.SysPostSvc{Ctx: ctx}
	_, err := s.DeleteSysPost(req)
	SetResponse(resp, err)
	return nil

}

// 匹配ID删除
func (h *SysPostHandler) DeleteById(ctx context.Context, req *proto.SysPost, resp *proto.Response) error {
	s := service.SysPostSvc{Ctx: ctx}
	_, err := s.DeleteByIdSysPost(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *SysPostHandler) DeleteByIds(ctx context.Context, req *proto.SysPostBatchDeleteRequest, resp *proto.Response) error {
	s := service.SysPostSvc{Ctx: ctx}
	_, err := s.DeleteByIdsSysPost(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *SysPostHandler) Update(ctx context.Context, req *proto.SysPost, resp *proto.Response) error {
	s := service.SysPostSvc{Ctx: ctx}
	_, err := s.UpdateSysPost(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *SysPostHandler) UpdateAll(ctx context.Context, req *proto.SysPost, resp *proto.Response) error {
	s := service.SysPostSvc{Ctx: ctx}
	_, err := s.UpdateAllSysPost(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *SysPostHandler) UpdateFields(ctx context.Context, req *proto.SysPostUpdateFieldsRequest, resp *proto.Response) error {
	s := service.SysPostSvc{Ctx: ctx}
	_, err := s.UpdateFieldsSysPost(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *SysPostHandler) Find(ctx context.Context, req *proto.SysPostFilter, resp *proto.SysPostResponse) error {
	s := service.SysPostSvc{Ctx: ctx}
	data, err := s.FindSysPost(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *SysPostHandler) FindById(ctx context.Context, req *proto.SysPostFilter, resp *proto.SysPostResponse) error {
	s := service.SysPostSvc{Ctx: ctx}
	data, err := s.FindByIdSysPost(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *SysPostHandler) Lists(ctx context.Context, req *proto.SysPostListRequest, resp *proto.SysPostResponse) error {
	s := service.SysPostSvc{Ctx: ctx}
	data, total, err := s.GetListSysPost(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *SysPostHandler) SetResponse(resp *proto.SysPostResponse, data *proto.SysPost, err error) {
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

func (h *SysPostHandler) SetPageResponse(resp *proto.SysPostResponse, list []*proto.SysPost, total int64, err error) {
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
