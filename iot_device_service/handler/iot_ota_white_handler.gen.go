// Code generated by sgen.exe,2022-04-21 14:24:41. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_device_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type IotOtaWhiteHandler struct{}

// 创建
func (h *IotOtaWhiteHandler) Create(ctx context.Context, req *proto.IotOtaWhite, resp *proto.Response) error {
	s := service.IotOtaWhiteSvc{Ctx: ctx}
	_, err := s.CreateIotOtaWhite(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *IotOtaWhiteHandler) Delete(ctx context.Context, req *proto.IotOtaWhite, resp *proto.Response) error {
	s := service.IotOtaWhiteSvc{Ctx: ctx}
	_, err := s.DeleteIotOtaWhite(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *IotOtaWhiteHandler) DeleteById(ctx context.Context, req *proto.IotOtaWhite, resp *proto.Response) error {
	s := service.IotOtaWhiteSvc{Ctx: ctx}
	_, err := s.DeleteByIdIotOtaWhite(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *IotOtaWhiteHandler) DeleteByIds(ctx context.Context, req *proto.IotOtaWhiteBatchDeleteRequest, resp *proto.Response) error {
	s := service.IotOtaWhiteSvc{Ctx: ctx}
	_, err := s.DeleteByIdsIotOtaWhite(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *IotOtaWhiteHandler) Update(ctx context.Context, req *proto.IotOtaWhite, resp *proto.Response) error {
	s := service.IotOtaWhiteSvc{Ctx: ctx}
	_, err := s.UpdateIotOtaWhite(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *IotOtaWhiteHandler) UpdateAll(ctx context.Context, req *proto.IotOtaWhite, resp *proto.Response) error {
	s := service.IotOtaWhiteSvc{Ctx: ctx}
	_, err := s.UpdateAllIotOtaWhite(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *IotOtaWhiteHandler) UpdateFields(ctx context.Context, req *proto.IotOtaWhiteUpdateFieldsRequest, resp *proto.Response) error {
	s := service.IotOtaWhiteSvc{Ctx: ctx}
	_, err := s.UpdateFieldsIotOtaWhite(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *IotOtaWhiteHandler) Find(ctx context.Context, req *proto.IotOtaWhiteFilter, resp *proto.IotOtaWhiteResponse) error {
	s := service.IotOtaWhiteSvc{Ctx: ctx}
	data, err := s.FindIotOtaWhite(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *IotOtaWhiteHandler) FindById(ctx context.Context, req *proto.IotOtaWhiteFilter, resp *proto.IotOtaWhiteResponse) error {
	s := service.IotOtaWhiteSvc{Ctx: ctx}
	data, err := s.FindByIdIotOtaWhite(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *IotOtaWhiteHandler) Lists(ctx context.Context, req *proto.IotOtaWhiteListRequest, resp *proto.IotOtaWhiteResponse) error {
	s := service.IotOtaWhiteSvc{Ctx: ctx}
	data, total, err := s.GetListIotOtaWhite(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *IotOtaWhiteHandler) SetResponse(resp *proto.IotOtaWhiteResponse, data *proto.IotOtaWhite, err error) {
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

func (h *IotOtaWhiteHandler) SetPageResponse(resp *proto.IotOtaWhiteResponse, list []*proto.IotOtaWhite, total int64, err error) {
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
