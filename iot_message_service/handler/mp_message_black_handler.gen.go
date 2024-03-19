// Code generated by sgen.exe,2022-05-22 07:25:57. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_message_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type MpMessageBlackHandler struct{}

// 创建
func (h *MpMessageBlackHandler) Create(ctx context.Context, req *proto.MpMessageBlack, resp *proto.Response) error {
	s := service.MpMessageBlackSvc{Ctx: ctx}
	_, err := s.CreateMpMessageBlack(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *MpMessageBlackHandler) Delete(ctx context.Context, req *proto.MpMessageBlack, resp *proto.Response) error {
	s := service.MpMessageBlackSvc{Ctx: ctx}
	_, err := s.DeleteMpMessageBlack(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *MpMessageBlackHandler) DeleteById(ctx context.Context, req *proto.MpMessageBlack, resp *proto.Response) error {
	s := service.MpMessageBlackSvc{Ctx: ctx}
	_, err := s.DeleteByIdMpMessageBlack(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *MpMessageBlackHandler) DeleteByIds(ctx context.Context, req *proto.MpMessageBlackBatchDeleteRequest, resp *proto.Response) error {
	s := service.MpMessageBlackSvc{Ctx: ctx}
	_, err := s.DeleteByIdsMpMessageBlack(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *MpMessageBlackHandler) Update(ctx context.Context, req *proto.MpMessageBlack, resp *proto.Response) error {
	s := service.MpMessageBlackSvc{Ctx: ctx}
	_, err := s.UpdateMpMessageBlack(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *MpMessageBlackHandler) UpdateAll(ctx context.Context, req *proto.MpMessageBlack, resp *proto.Response) error {
	s := service.MpMessageBlackSvc{Ctx: ctx}
	_, err := s.UpdateAllMpMessageBlack(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *MpMessageBlackHandler) UpdateFields(ctx context.Context, req *proto.MpMessageBlackUpdateFieldsRequest, resp *proto.Response) error {
	s := service.MpMessageBlackSvc{Ctx: ctx}
	_, err := s.UpdateFieldsMpMessageBlack(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *MpMessageBlackHandler) Find(ctx context.Context, req *proto.MpMessageBlackFilter, resp *proto.MpMessageBlackResponse) error {
	s := service.MpMessageBlackSvc{Ctx: ctx}
	data, err := s.FindMpMessageBlack(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *MpMessageBlackHandler) FindById(ctx context.Context, req *proto.MpMessageBlackFilter, resp *proto.MpMessageBlackResponse) error {
	s := service.MpMessageBlackSvc{Ctx: ctx}
	data, err := s.FindByIdMpMessageBlack(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *MpMessageBlackHandler) Lists(ctx context.Context, req *proto.MpMessageBlackListRequest, resp *proto.MpMessageBlackResponse) error {
	s := service.MpMessageBlackSvc{Ctx: ctx}
	data, total, err := s.GetListMpMessageBlack(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *MpMessageBlackHandler) SetResponse(resp *proto.MpMessageBlackResponse, data *proto.MpMessageBlack, err error) {
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

func (h *MpMessageBlackHandler) SetPageResponse(resp *proto.MpMessageBlackResponse, list []*proto.MpMessageBlack, total int64, err error) {
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
