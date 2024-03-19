// Code generated by sgen.exe,2022-06-02 11:15:12. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_app_oem_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OemAppPushCertHandler struct{}

// 创建
func (h *OemAppPushCertHandler) Create(ctx context.Context, req *proto.OemAppPushCert, resp *proto.Response) error {
	s := service.OemAppPushCertSvc{Ctx: ctx}
	_, err := s.CreateOemAppPushCert(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *OemAppPushCertHandler) Delete(ctx context.Context, req *proto.OemAppPushCert, resp *proto.Response) error {
	s := service.OemAppPushCertSvc{Ctx: ctx}
	_, err := s.DeleteOemAppPushCert(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *OemAppPushCertHandler) DeleteById(ctx context.Context, req *proto.OemAppPushCert, resp *proto.Response) error {
	s := service.OemAppPushCertSvc{Ctx: ctx}
	_, err := s.DeleteByIdOemAppPushCert(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *OemAppPushCertHandler) DeleteByIds(ctx context.Context, req *proto.OemAppPushCertBatchDeleteRequest, resp *proto.Response) error {
	s := service.OemAppPushCertSvc{Ctx: ctx}
	_, err := s.DeleteByIdsOemAppPushCert(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *OemAppPushCertHandler) Update(ctx context.Context, req *proto.OemAppPushCert, resp *proto.Response) error {
	s := service.OemAppPushCertSvc{Ctx: ctx}
	_, err := s.UpdateOemAppPushCert(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *OemAppPushCertHandler) UpdateAll(ctx context.Context, req *proto.OemAppPushCert, resp *proto.Response) error {
	s := service.OemAppPushCertSvc{Ctx: ctx}
	_, err := s.UpdateAllOemAppPushCert(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *OemAppPushCertHandler) UpdateFields(ctx context.Context, req *proto.OemAppPushCertUpdateFieldsRequest, resp *proto.Response) error {
	s := service.OemAppPushCertSvc{Ctx: ctx}
	_, err := s.UpdateFieldsOemAppPushCert(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *OemAppPushCertHandler) Find(ctx context.Context, req *proto.OemAppPushCertFilter, resp *proto.OemAppPushCertResponse) error {
	s := service.OemAppPushCertSvc{Ctx: ctx}
	data, err := s.FindOemAppPushCert(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *OemAppPushCertHandler) FindById(ctx context.Context, req *proto.OemAppPushCertFilter, resp *proto.OemAppPushCertResponse) error {
	s := service.OemAppPushCertSvc{Ctx: ctx}
	data, err := s.FindByIdOemAppPushCert(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *OemAppPushCertHandler) Lists(ctx context.Context, req *proto.OemAppPushCertListRequest, resp *proto.OemAppPushCertResponse) error {
	s := service.OemAppPushCertSvc{Ctx: ctx}
	data, total, err := s.GetListOemAppPushCert(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *OemAppPushCertHandler) SetResponse(resp *proto.OemAppPushCertResponse, data *proto.OemAppPushCert, err error) {
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

func (h *OemAppPushCertHandler) SetPageResponse(resp *proto.OemAppPushCertResponse, list []*proto.OemAppPushCert, total int64, err error) {
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
