// Code generated by sgen.exe,2022-04-21 19:06:38. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_device_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type IotDeviceLogHandler struct{}

func (h *IotDeviceLogHandler) FailLogLists(ctx context.Context, request *proto.DeviceOperationFailLogListRequest, response *proto.DeviceOperationFailLogListResponse) error {
	s := service.IotDeviceLogProductSvc{Ctx: ctx}
	data, total, err := s.FailLogLists(request)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
		response.Total = total
		response.Data = data
	}
	return nil
}

func (h *IotDeviceLogHandler) OperationFailLogReport(ctx context.Context, request *proto.OperationFailLogRequest, response *proto.Response) error {
	s := service.IotDeviceLogProductSvc{Ctx: ctx}
	err := s.SaveOperationFailLog(request)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
	}
	return nil
}

func (h *IotDeviceLogHandler) DeviceOperationFailLogReport(ctx context.Context, request *proto.DeviceOperationFailLogRequest, response *proto.Response) error {
	s := service.IotDeviceLogProductSvc{Ctx: ctx}
	err := s.SaveDeviceOperationFailLog(request)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
	}
	return nil
}

func (h *IotDeviceLogHandler) ProductEventLogReport(ctx context.Context, request *proto.ProductLogRequest, response *proto.ProductLogResponse) error {
	s := service.IotDeviceLogProductSvc{Ctx: ctx}
	list, err := s.QueryDeviceLogs(request)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
		response.List = list
	}
	return nil
}

func (h *IotDeviceLogHandler) ProductReportLogRecord(ctx context.Context, request *proto.ProductLogRequest, response *proto.ProductLogRecordResponse) error {
	s := service.IotDeviceLogProductSvc{Ctx: ctx}
	list, err := s.QueryDeviceLogRecord(request)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
		response.Records = list
	}
	return nil
}

func (h *IotDeviceLogHandler) CreateProductLogTable(ctx context.Context, req *proto.CreateProductLogTableResponse, resp *proto.Response) error {
	s := service.IotDeviceLogProductSvc{Ctx: ctx}
	err := s.CreateLogTable(req)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
	}
	return nil
}

func (h *IotDeviceLogHandler) Export(ctx context.Context, request *proto.IotDeviceLogListRequest, response *proto.IotDeviceLogDExportResponse) error {
	//TODO implement me
	panic("implement me")
}

func (h *IotDeviceLogHandler) Count(ctx context.Context, req *proto.IotDeviceLogListRequest, resp *proto.IotDeviceLogResponse) error {
	s := service.IotDeviceLogSvc{Ctx: ctx}
	data, total, err := s.GetListIotDeviceLog(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

// 创建
func (h *IotDeviceLogHandler) Create(ctx context.Context, req *proto.IotDeviceLog, resp *proto.Response) error {
	s := service.IotDeviceLogSvc{Ctx: ctx}
	_, err := s.CreateIotDeviceLog(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *IotDeviceLogHandler) Delete(ctx context.Context, req *proto.IotDeviceLog, resp *proto.Response) error {
	s := service.IotDeviceLogSvc{Ctx: ctx}
	_, err := s.DeleteIotDeviceLog(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *IotDeviceLogHandler) DeleteById(ctx context.Context, req *proto.IotDeviceLog, resp *proto.Response) error {
	s := service.IotDeviceLogSvc{Ctx: ctx}
	_, err := s.DeleteByIdIotDeviceLog(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *IotDeviceLogHandler) DeleteByIds(ctx context.Context, req *proto.IotDeviceLogBatchDeleteRequest, resp *proto.Response) error {
	s := service.IotDeviceLogSvc{Ctx: ctx}
	_, err := s.DeleteByIdsIotDeviceLog(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *IotDeviceLogHandler) Update(ctx context.Context, req *proto.IotDeviceLog, resp *proto.Response) error {
	s := service.IotDeviceLogSvc{Ctx: ctx}
	_, err := s.UpdateIotDeviceLog(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *IotDeviceLogHandler) UpdateAll(ctx context.Context, req *proto.IotDeviceLog, resp *proto.Response) error {
	s := service.IotDeviceLogSvc{Ctx: ctx}
	_, err := s.UpdateAllIotDeviceLog(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *IotDeviceLogHandler) UpdateFields(ctx context.Context, req *proto.IotDeviceLogUpdateFieldsRequest, resp *proto.Response) error {
	s := service.IotDeviceLogSvc{Ctx: ctx}
	_, err := s.UpdateFieldsIotDeviceLog(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *IotDeviceLogHandler) Find(ctx context.Context, req *proto.IotDeviceLogFilter, resp *proto.IotDeviceLogResponse) error {
	s := service.IotDeviceLogSvc{Ctx: ctx}
	data, err := s.FindIotDeviceLog(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *IotDeviceLogHandler) FindById(ctx context.Context, req *proto.IotDeviceLogFilter, resp *proto.IotDeviceLogResponse) error {
	s := service.IotDeviceLogSvc{Ctx: ctx}
	data, err := s.FindByIdIotDeviceLog(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *IotDeviceLogHandler) Lists(ctx context.Context, req *proto.IotDeviceLogListRequest, resp *proto.IotDeviceLogResponse) error {
	s := service.IotDeviceLogSvc{Ctx: ctx}
	data, total, err := s.GetListIotDeviceLog(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *IotDeviceLogHandler) SetResponse(resp *proto.IotDeviceLogResponse, data *proto.IotDeviceLog, err error) {
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

func (h *IotDeviceLogHandler) SetPageResponse(resp *proto.IotDeviceLogResponse, list []*proto.IotDeviceLog, total int64, err error) {
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

// 清空设备日志
func (h *IotDeviceLogHandler) ClearDeviceLogs(ctx context.Context, request *proto.ProductLogRequest, response *proto.Response) error {
	s := service.IotDeviceLogProductSvc{Ctx: ctx}
	err := s.ClearDeviceLogs(request)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
	}
	return nil
}
