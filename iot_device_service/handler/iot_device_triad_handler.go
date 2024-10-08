// Code generated by sgen.exe,2022-04-21 14:24:40. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"context"

	"cloud_platform/iot_device_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type IotDeviceTriadHandler struct{}

func (h *IotDeviceTriadHandler) SetTestDeviceTriad(ctx context.Context, request *proto.SetTestTriadRequest, response *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	err := s.SetTestDeviceTriad(request)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
	}
	return nil
}

func (h *IotDeviceTriadHandler) BindTestAccount(ctx context.Context, request *proto.BindTestAccountRequest, response *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	err := s.BindTestAccount(request)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
	}
	return nil
}

func (h *IotDeviceTriadHandler) GetDeviceTriadCountByTenantId(ctx context.Context, request *proto.IotDeviceTriadFilter, response *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	total, err := s.GetDeviceTriadCountByTenantId(request)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
		response.Data = total
	}
	return nil
}

func (h *IotDeviceTriadHandler) GeneratorDeviceTriad(ctx context.Context, request *proto.IotDeviceTriadGenerateRequest, response *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	err := s.GeneratorDeviceTriad(request)
	SetResponse(response, err)
	return nil
}

func (h *IotDeviceTriadHandler) CreateAndBindDeviceTriad(ctx context.Context, request *proto.IotDeviceTriadGenerateRequest, response *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	err := s.CreateAndBindDeviceTriad(request)
	SetResponse(response, err)
	return nil
}

// 创建
func (h *IotDeviceTriadHandler) Create(ctx context.Context, req *proto.IotDeviceTriad, resp *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	_, err := s.CreateIotDeviceTriad(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *IotDeviceTriadHandler) Delete(ctx context.Context, req *proto.IotDeviceTriad, resp *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	_, err := s.DeleteIotDeviceTriad(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *IotDeviceTriadHandler) DeleteById(ctx context.Context, req *proto.IotDeviceTriad, resp *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	_, err := s.DeleteByIdIotDeviceTriad(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *IotDeviceTriadHandler) DeleteByIds(ctx context.Context, req *proto.IotDeviceTriadBatchDeleteRequest, resp *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	_, err := s.DeleteByIdsIotDeviceTriad(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *IotDeviceTriadHandler) Update(ctx context.Context, req *proto.IotDeviceTriad, resp *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	_, err := s.UpdateIotDeviceTriad(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *IotDeviceTriadHandler) UpdateAll(ctx context.Context, req *proto.IotDeviceTriad, resp *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	_, err := s.UpdateAllIotDeviceTriad(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *IotDeviceTriadHandler) UpdateFields(ctx context.Context, req *proto.IotDeviceTriadUpdateFieldsRequest, resp *proto.Response) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	_, err := s.UpdateFieldsIotDeviceTriad(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *IotDeviceTriadHandler) Find(ctx context.Context, req *proto.IotDeviceTriadFilter, resp *proto.IotDeviceTriadResponse) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	data, err := s.FindIotDeviceTriad(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *IotDeviceTriadHandler) FindById(ctx context.Context, req *proto.IotDeviceTriadFilter, resp *proto.IotDeviceTriadResponse) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	data, err := s.FindByIdIotDeviceTriad(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *IotDeviceTriadHandler) Lists(ctx context.Context, req *proto.IotDeviceTriadListRequest, resp *proto.IotDeviceTriadResponse) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	data, total, err := s.GetListIotDeviceTriad(req)
	if err != nil {
		iotlogger.LogHelper.Errorf("进入设备服务Lists出现异常, %s", err.Error())
	}
	h.SetPageResponse(resp, data, total, err)
	return nil
}

// 设备导出数据统计
func (h *IotDeviceTriadHandler) SetExportCount(ctx context.Context, req *proto.IotDeviceTriadListRequest, resp *proto.IotDeviceTriadResponse) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	err := s.SetExportCount(req)
	h.SetResponse(resp, nil, err)
	return nil
}

// 设备导出数据统计
func (h *IotDeviceTriadHandler) GetDeviceTriadCount(ctx context.Context, req *proto.IotDeviceTriadCountRequest, resp *proto.IotDeviceTriadCountResponse) error {
	s := service.IotDeviceTriadSvc{Ctx: ctx}
	res, err := s.GetDeviceTriadCount(req)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.Data = res
	}
	return nil
}


func (h *IotDeviceTriadHandler) SetResponse(resp *proto.IotDeviceTriadResponse, data *proto.IotDeviceTriad, err error) {
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

func (h *IotDeviceTriadHandler) SetPageResponse(resp *proto.IotDeviceTriadResponse, list []*proto.IotDeviceTriad, total int64, err error) {
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
