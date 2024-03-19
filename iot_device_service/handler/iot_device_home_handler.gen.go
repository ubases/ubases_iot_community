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

type IotDeviceHomeHandler struct{}

func (h *IotDeviceHomeHandler) QueryDeviceAreas(ctx context.Context, req *proto.DeviceAreaRequest, resp *proto.DeviceAreaResponse) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	res, err := s.QueryDeviceAreas(req)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.Areas = res
	}
	return nil
}

// 创建
func (h *IotDeviceHomeHandler) Create(ctx context.Context, req *proto.IotDeviceHome, resp *proto.Response) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	_, err := s.CreateIotDeviceHome(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *IotDeviceHomeHandler) Delete(ctx context.Context, req *proto.IotDeviceHome, resp *proto.Response) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	_, err := s.DeleteIotDeviceHome(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *IotDeviceHomeHandler) DeleteById(ctx context.Context, req *proto.IotDeviceHome, resp *proto.Response) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	_, err := s.DeleteByIdIotDeviceHome(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *IotDeviceHomeHandler) DeleteByIds(ctx context.Context, req *proto.IotDeviceHomeBatchDeleteRequest, resp *proto.Response) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	_, err := s.DeleteByIdsIotDeviceHome(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *IotDeviceHomeHandler) Update(ctx context.Context, req *proto.IotDeviceHome, resp *proto.Response) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	_, err := s.UpdateIotDeviceHome(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *IotDeviceHomeHandler) UpdateAll(ctx context.Context, req *proto.IotDeviceHome, resp *proto.Response) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	_, err := s.UpdateAllIotDeviceHome(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *IotDeviceHomeHandler) UpdateFields(ctx context.Context, req *proto.IotDeviceHomeUpdateFieldsRequest, resp *proto.Response) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	_, err := s.UpdateFieldsIotDeviceHome(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *IotDeviceHomeHandler) Find(ctx context.Context, req *proto.IotDeviceHomeFilter, resp *proto.IotDeviceHomeResponse) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	data, err := s.FindIotDeviceHome(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *IotDeviceHomeHandler) FindById(ctx context.Context, req *proto.IotDeviceHomeFilter, resp *proto.IotDeviceHomeResponse) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	data, err := s.FindByIdIotDeviceHome(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *IotDeviceHomeHandler) Lists(ctx context.Context, req *proto.IotDeviceHomeListRequest, resp *proto.IotDeviceHomeResponse) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	data, total, err := s.GetListIotDeviceHome(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *IotDeviceHomeHandler) SetResponse(resp *proto.IotDeviceHomeResponse, data *proto.IotDeviceHome, err error) {
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

func (h *IotDeviceHomeHandler) SetPageResponse(resp *proto.IotDeviceHomeResponse, list []*proto.IotDeviceHome, total int64, err error) {
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

// 家庭房间设备数量
func (h *IotDeviceHomeHandler) DevCount(ctx context.Context, req *proto.IotDeviceHomeDevCount, resp *proto.DevCountResponse) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	data, err := s.GetDevCountByHomeRoom(req)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.Keys = data
	}
	return err
}

func (h *IotDeviceHomeHandler) UserDev(ctx context.Context, req *proto.IotUserHomeDev, resp *proto.IotDeviceHomeResponse) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	//data,err := s.UserDevCount(req)
	data, total, err := s.UserDevCount(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *IotDeviceHomeHandler) UserDevList(ctx context.Context, req *proto.IotDeviceHomeHomeId, resp *proto.DevListResponse) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	data, err := s.HomeDevList(req.HomeId, req.HomeIds)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.DevList = data
	}
	return err
}

func (h *IotDeviceHomeHandler) HomeDevCount(ctx context.Context, req *proto.IotDeviceHomeHomeId, resp *proto.IotHomeDevCountResponse) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	data, err := s.HomeDevCount(req)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.DevCount = data.DevCount
		resp.DevCounts = data.DevCounts
	}
	return err
}

func (h *IotDeviceHomeHandler) HomeDevList(ctx context.Context, req *proto.IotDeviceHomeHomeId, resp *proto.DevListResponse) error {
	iotlogger.LogHelper.Infof("进入设备服务HomeDevList")
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	data, err := s.HomeDevList(req.HomeId, req.HomeIds)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
		iotlogger.LogHelper.Errorf("进入设备服务HomeDevList出现异常, %s", err.Error())
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.DevList = data
	}
	return err
}

func (h *IotDeviceHomeHandler) RemoveDev(ctx context.Context, request *proto.RemoveDevRequest, response *proto.RemoveDevResponse) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	err := s.RemoveDev(request)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
	}
	return err
}

func (h *IotDeviceHomeHandler) UpdateDeviceInfo(ctx context.Context, req *proto.IotDeviceHome, resp *proto.Response) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	_, err := s.UpdateDeviceInfo(req)
	SetResponse(resp, err)
	return nil
}

func (h *IotDeviceHomeHandler) SetDevSort(ctx context.Context, req *proto.SetDevSortRequest, resp *proto.Response) error {
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	_, err := s.SetDevSort(req)
	SetResponse(resp, err)
	return nil
}

func (h *IotDeviceHomeHandler) HomeDevListExcludeVirtualDevices(ctx context.Context, req *proto.IotDeviceHomeHomeId, resp *proto.DevListResponse) error {
	iotlogger.LogHelper.Infof("进入设备服务HomeDevList")
	s := service.IotDeviceHomeSvc{Ctx: ctx}
	data, err := s.HomeDevListExcludeVirtualDevices(req.HomeId, req.HomeIds)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
		iotlogger.LogHelper.Errorf("进入设备服务HomeDevList出现异常, %s", err.Error())
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.DevList = data
	}
	return err
}
