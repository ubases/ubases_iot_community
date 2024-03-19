// Code generated by sgen.exe,2022-05-22 07:25:57. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_message_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type MpMessageUserOutHandler struct{}

func (h *MpMessageUserOutHandler) GroupLists(ctx context.Context, request *proto.MpMessageUserOutListRequest, response *proto.MpMessageUserOutGroupResponse) error {
	s := service.MpMessageUserOutSvc{Ctx: ctx}
	res, total, err := s.GetGroupListMpMessageUserOut(request)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
		response.Total = total
		response.Data = res
	}
	return nil
}

// 创建
func (h *MpMessageUserOutHandler) Create(ctx context.Context, req *proto.MpMessageUserOut, resp *proto.Response) error {
	s := service.MpMessageUserOutSvc{Ctx: ctx}
	_, err := s.CreateMpMessageUserOut(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *MpMessageUserOutHandler) Delete(ctx context.Context, req *proto.MpMessageUserOut, resp *proto.Response) error {
	s := service.MpMessageUserOutSvc{Ctx: ctx}
	_, err := s.DeleteMpMessageUserOut(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *MpMessageUserOutHandler) DeleteById(ctx context.Context, req *proto.MpMessageUserOut, resp *proto.Response) error {
	s := service.MpMessageUserOutSvc{Ctx: ctx}
	_, err := s.DeleteByIdMpMessageUserOut(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *MpMessageUserOutHandler) DeleteByIds(ctx context.Context, req *proto.MpMessageUserOutBatchDeleteRequest, resp *proto.Response) error {
	s := service.MpMessageUserOutSvc{Ctx: ctx}
	_, err := s.DeleteByIdsMpMessageUserOut(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *MpMessageUserOutHandler) Update(ctx context.Context, req *proto.MpMessageUserOut, resp *proto.Response) error {
	s := service.MpMessageUserOutSvc{Ctx: ctx}
	_, err := s.UpdateMpMessageUserOut(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *MpMessageUserOutHandler) UpdateAll(ctx context.Context, req *proto.MpMessageUserOut, resp *proto.Response) error {
	s := service.MpMessageUserOutSvc{Ctx: ctx}
	_, err := s.UpdateAllMpMessageUserOut(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *MpMessageUserOutHandler) UpdateFields(ctx context.Context, req *proto.MpMessageUserOutUpdateFieldsRequest, resp *proto.Response) error {
	s := service.MpMessageUserOutSvc{Ctx: ctx}
	_, err := s.UpdateFieldsMpMessageUserOut(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *MpMessageUserOutHandler) Find(ctx context.Context, req *proto.MpMessageUserOutFilter, resp *proto.MpMessageUserOutResponse) error {
	s := service.MpMessageUserOutSvc{Ctx: ctx}
	data, err := s.FindMpMessageUserOut(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *MpMessageUserOutHandler) FindById(ctx context.Context, req *proto.MpMessageUserOutFilter, resp *proto.MpMessageUserOutResponse) error {
	s := service.MpMessageUserOutSvc{Ctx: ctx}
	data, err := s.FindByIdMpMessageUserOut(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *MpMessageUserOutHandler) Lists(ctx context.Context, req *proto.MpMessageUserOutListRequest, resp *proto.MpMessageUserOutResponse) error {
	s := service.MpMessageUserOutSvc{Ctx: ctx}
	data, total, err := s.GetListMpMessageUserOut(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *MpMessageUserOutHandler) SetResponse(resp *proto.MpMessageUserOutResponse, data *proto.MpMessageUserOut, err error) {
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

func (h *MpMessageUserOutHandler) SetPageResponse(resp *proto.MpMessageUserOutResponse, list []*proto.MpMessageUserOut, total int64, err error) {
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
