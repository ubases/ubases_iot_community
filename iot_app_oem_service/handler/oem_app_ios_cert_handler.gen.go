// Code generated by sgen.exe,2022-06-02 11:15:11. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_app_oem_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OemAppIosCertHandler struct{}

// 创建
func (h *OemAppIosCertHandler) Create(ctx context.Context, req *proto.OemAppIosCert, resp *proto.Response) error {
	s := service.OemAppIosCertSvc{Ctx: ctx}
	_, err := s.CreateOemAppIosCert(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *OemAppIosCertHandler) Delete(ctx context.Context, req *proto.OemAppIosCert, resp *proto.Response) error {
	s := service.OemAppIosCertSvc{Ctx: ctx}
	_, err := s.DeleteOemAppIosCert(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *OemAppIosCertHandler) DeleteById(ctx context.Context, req *proto.OemAppIosCert, resp *proto.Response) error {
	s := service.OemAppIosCertSvc{Ctx: ctx}
	_, err := s.DeleteByIdOemAppIosCert(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *OemAppIosCertHandler) DeleteByIds(ctx context.Context, req *proto.OemAppIosCertBatchDeleteRequest, resp *proto.Response) error {
	s := service.OemAppIosCertSvc{Ctx: ctx}
	_, err := s.DeleteByIdsOemAppIosCert(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *OemAppIosCertHandler) Update(ctx context.Context, req *proto.OemAppIosCert, resp *proto.Response) error {
	s := service.OemAppIosCertSvc{Ctx: ctx}
	_, err := s.UpdateOemAppIosCert(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *OemAppIosCertHandler) UpdateAll(ctx context.Context, req *proto.OemAppIosCert, resp *proto.Response) error {
	s := service.OemAppIosCertSvc{Ctx: ctx}
	_, err := s.UpdateAllOemAppIosCert(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *OemAppIosCertHandler) UpdateFields(ctx context.Context, req *proto.OemAppIosCertUpdateFieldsRequest, resp *proto.Response) error {
	s := service.OemAppIosCertSvc{Ctx: ctx}
	_, err := s.UpdateFieldsOemAppIosCert(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *OemAppIosCertHandler) Find(ctx context.Context, req *proto.OemAppIosCertFilter, resp *proto.OemAppIosCertResponse) error {
	s := service.OemAppIosCertSvc{Ctx: ctx}
	data, err := s.FindOemAppIosCert(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *OemAppIosCertHandler) FindById(ctx context.Context, req *proto.OemAppIosCertFilter, resp *proto.OemAppIosCertResponse) error {
	s := service.OemAppIosCertSvc{Ctx: ctx}
	data, err := s.FindByIdOemAppIosCert(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *OemAppIosCertHandler) Lists(ctx context.Context, req *proto.OemAppIosCertListRequest, resp *proto.OemAppIosCertResponse) error {
	s := service.OemAppIosCertSvc{Ctx: ctx}
	data, total, err := s.GetListOemAppIosCert(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *OemAppIosCertHandler) SetResponse(resp *proto.OemAppIosCertResponse, data *proto.OemAppIosCert, err error) {
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

func (h *OemAppIosCertHandler) SetPageResponse(resp *proto.OemAppIosCertResponse, list []*proto.OemAppIosCert, total int64, err error) {
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
