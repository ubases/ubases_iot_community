// Code generated by sgen.exe,2022-04-20 13:52:29. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_product_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type PmFirmwareVersionHandler struct{}

func (h *PmFirmwareVersionHandler) ModuleFirmwareVersionList(ctx context.Context, request *proto.ModuleFirmwareVersionRequest, response *proto.PmFirmwareVersionResponse) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	data, total, err := s.ModuleFirmwareVersionList(request)
	h.SetPageResponse(response, data, total, err)
	return nil
}

// 创建
func (h *PmFirmwareVersionHandler) Create(ctx context.Context, req *proto.PmFirmwareVersion, resp *proto.Response) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	_, err := s.CreatePmFirmwareVersion(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *PmFirmwareVersionHandler) Delete(ctx context.Context, req *proto.PmFirmwareVersion, resp *proto.Response) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	_, err := s.DeletePmFirmwareVersion(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *PmFirmwareVersionHandler) DeleteById(ctx context.Context, req *proto.PmFirmwareVersion, resp *proto.Response) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	_, err := s.DeleteByIdPmFirmwareVersion(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *PmFirmwareVersionHandler) DeleteByIds(ctx context.Context, req *proto.PmFirmwareVersionBatchDeleteRequest, resp *proto.Response) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	_, err := s.DeleteByIdsPmFirmwareVersion(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *PmFirmwareVersionHandler) Update(ctx context.Context, req *proto.PmFirmwareVersion, resp *proto.Response) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	_, err := s.UpdatePmFirmwareVersion(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *PmFirmwareVersionHandler) UpdateAll(ctx context.Context, req *proto.PmFirmwareVersion, resp *proto.Response) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	_, err := s.UpdateAllPmFirmwareVersion(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *PmFirmwareVersionHandler) UpdateFields(ctx context.Context, req *proto.PmFirmwareVersionUpdateFieldsRequest, resp *proto.Response) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	_, err := s.UpdateFieldsPmFirmwareVersion(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *PmFirmwareVersionHandler) Find(ctx context.Context, req *proto.PmFirmwareVersionFilter, resp *proto.PmFirmwareVersionResponse) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	data, err := s.FindPmFirmwareVersion(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *PmFirmwareVersionHandler) FindById(ctx context.Context, req *proto.PmFirmwareVersionFilter, resp *proto.PmFirmwareVersionResponse) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	data, err := s.FindByIdPmFirmwareVersion(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *PmFirmwareVersionHandler) Lists(ctx context.Context, req *proto.PmFirmwareVersionListRequest, resp *proto.PmFirmwareVersionResponse) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	data, total, err := s.GetListPmFirmwareVersion(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *PmFirmwareVersionHandler) UpdateStatusByFirmware(ctx context.Context, req *proto.PmFirmwareVersionFilter, response *proto.Response) error {
	s := service.PmFirmwareVersionSvc{Ctx: ctx}
	err := s.UpdateStatusByFirmware(req)
	SetResponse(response, err)
	return nil
}

func (h *PmFirmwareVersionHandler) SetResponse(resp *proto.PmFirmwareVersionResponse, data *proto.PmFirmwareVersion, err error) {
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

func (h *PmFirmwareVersionHandler) SetPageResponse(resp *proto.PmFirmwareVersionResponse, list []*proto.PmFirmwareVersion, total int64, err error) {
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
