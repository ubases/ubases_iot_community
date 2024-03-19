// Code generated by sgen.exe,2022-06-02 11:15:11. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_app_oem_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OemAppBuildRecordHandler struct{}

// 创建
func (h *OemAppBuildRecordHandler) Create(ctx context.Context, req *proto.OemAppBuildRecord, resp *proto.Response) error {
	s := service.OemAppBuildRecordSvc{Ctx: ctx}
	_, err := s.CreateOemAppBuildRecord(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *OemAppBuildRecordHandler) Delete(ctx context.Context, req *proto.OemAppBuildRecord, resp *proto.Response) error {
	s := service.OemAppBuildRecordSvc{Ctx: ctx}
	_, err := s.DeleteOemAppBuildRecord(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *OemAppBuildRecordHandler) DeleteById(ctx context.Context, req *proto.OemAppBuildRecord, resp *proto.Response) error {
	s := service.OemAppBuildRecordSvc{Ctx: ctx}
	_, err := s.DeleteByIdOemAppBuildRecord(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *OemAppBuildRecordHandler) DeleteByIds(ctx context.Context, req *proto.OemAppBuildRecordBatchDeleteRequest, resp *proto.Response) error {
	s := service.OemAppBuildRecordSvc{Ctx: ctx}
	_, err := s.DeleteByIdsOemAppBuildRecord(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *OemAppBuildRecordHandler) Update(ctx context.Context, req *proto.OemAppBuildRecord, resp *proto.Response) error {
	s := service.OemAppBuildRecordSvc{Ctx: ctx}
	_, err := s.UpdateOemAppBuildRecord(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *OemAppBuildRecordHandler) UpdateAll(ctx context.Context, req *proto.OemAppBuildRecord, resp *proto.Response) error {
	s := service.OemAppBuildRecordSvc{Ctx: ctx}
	_, err := s.UpdateAllOemAppBuildRecord(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *OemAppBuildRecordHandler) UpdateFields(ctx context.Context, req *proto.OemAppBuildRecordUpdateFieldsRequest, resp *proto.Response) error {
	s := service.OemAppBuildRecordSvc{Ctx: ctx}
	_, err := s.UpdateFieldsOemAppBuildRecord(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *OemAppBuildRecordHandler) Find(ctx context.Context, req *proto.OemAppBuildRecordFilter, resp *proto.OemAppBuildRecordResponse) error {
	s := service.OemAppBuildRecordSvc{Ctx: ctx}
	data, err := s.FindOemAppBuildRecord(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *OemAppBuildRecordHandler) FindById(ctx context.Context, req *proto.OemAppBuildRecordFilter, resp *proto.OemAppBuildRecordResponse) error {
	s := service.OemAppBuildRecordSvc{Ctx: ctx}
	data, err := s.FindByIdOemAppBuildRecord(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *OemAppBuildRecordHandler) Lists(ctx context.Context, req *proto.OemAppBuildRecordListRequest, resp *proto.OemAppBuildRecordResponse) error {
	s := service.OemAppBuildRecordSvc{Ctx: ctx}
	data, total, err := s.GetListOemAppBuildRecord(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *OemAppBuildRecordHandler) SetResponse(resp *proto.OemAppBuildRecordResponse, data *proto.OemAppBuildRecord, err error) {
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

func (h *OemAppBuildRecordHandler) SetPageResponse(resp *proto.OemAppBuildRecordResponse, list []*proto.OemAppBuildRecord, total int64, err error) {
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
