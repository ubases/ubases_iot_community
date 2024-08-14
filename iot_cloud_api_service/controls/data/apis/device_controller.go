package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/data/entitys"
	"cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gin-gonic/gin"
)

var devicedatacontroller DeviceDataController

type DeviceDataController struct{} //部门操作控制器

func (DeviceDataController) getFaultList(c *gin.Context) {
	var req entitys.IotDeviceFaultQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	rpcReq := protosService.IotDeviceFaultListRequest{
		Page:      int64(req.Page),
		PageSize:  int64(req.Limit),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
		SearchKey: req.SearchKey,
	}
	if req.Query != nil {
		if req.Query.Developer != 0 {
			company, err := rpc.ClientOpenCompanyService.Find(context.Background(), &protosService.OpenCompanyFilter{UserId: req.Query.Developer})
			if err != nil {
				iotgin.ResErrCli(c, err)
				return
			}
			req.Query.TenantId = company.Data[0].TenantId
		}
		productId, _ := iotutil.ToInt64AndErr(req.Query.ProductId)
		faultCode := ""
		if req.Query.FaultCode >= 0 {
			faultCode = strconv.Itoa(int(req.Query.FaultCode))
		}
		baseProductId, _ := iotutil.ToInt64AndErr(req.Query.BaseProductId)
		rpcReq.Query = &protosService.IotDeviceFaultListFilter{
			//DeviceId:  int64(deviceId),
			DeviceKey:     req.Query.Did,
			ProductId:     productId,
			TenantId:      req.Query.TenantId,
			BaseProductId: baseProductId,
			FaultCode:     faultCode,
			LastDay:       req.Query.LastDay,
			StartTime:     timestamppb.New(time.Unix(req.Query.StartTime, 0)),
			EndTime:       timestamppb.New(time.Unix(req.Query.EndTime, 0)),
		}
	}
	//获取产品开发者信息
	openDev := services.OpenDevService{}
	openDevMap, err := openDev.GetOpenDevMap(0)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//获取故障数据
	grpcRes, err := rpc.ClientIotDeviceFault.Lists(context.Background(), &rpcReq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if grpcRes.Code != 200 {
		iotgin.ResErrCli(c, errors.New(grpcRes.Message))
		return
	}
	list := make([]*entitys.IotDeviceFaultEntitys, 0, len(grpcRes.Data))
	for _, v := range grpcRes.Data {
		row := entitys.IotDeviceFault_pb2e(v)
		if v.TenantId != "" {
			if devel, ok := openDevMap[v.TenantId]; ok {
				row.Developer = devel[0].UserName
			}
		}
		list = append(list, row)
	}
	iotgin.ResPageSuccess(c, list, grpcRes.Total, int(req.Page))
}

func (DeviceDataController) GetDeviceTotalStatistics(c *gin.Context) {
	grpcRes, err := rpc.ClientStatisticsService.GetDeviceTotalStatistics(c, nil)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if grpcRes == nil {
		iotgin.ResErrCli(c, errors.New("error."))
	}
	iotgin.ResSuccess(c, gin.H{"devActiveTotal": grpcRes.GetActiveTotal(), "devOnlineTotal": grpcRes.GetOnlineTotal()})
}

func (DeviceDataController) getFailLogList(c *gin.Context) {
	var req entitys.FailLogQueryEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//必须要指定设备key
	if req.Query.Did == "" {
		iotgin.ResBadRequest(c, "请指定设备编号(deviceKey)")
		return
	}
	rpcReq := protosService.DeviceOperationFailLogListRequest{
		Page:      int64(req.Page),
		PageSize:  int64(req.Limit),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
		SearchKey: req.SearchKey,
		Query: &protosService.DeviceOperationFailLogQueryObj{
			Did:          req.Query.Did,
			Code:         req.Query.Code,
			ProductKey:   req.Query.ProductKey,
			LastDay:      req.Query.LastDay,
			EndTime:      req.Query.EndTime,
			StartTime:    req.Query.StartTime,
			UploadFrom:   req.Query.UploadFrom,
			UploadMethod: req.Query.UploadMethod,
		},
	}
	grpcRes, err := rpc.ClientIotDeviceLogServer.FailLogLists(context.Background(), &rpcReq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if grpcRes.Code != 200 {
		iotgin.ResErrCli(c, errors.New(grpcRes.Message))
		return
	}
	list := make([]*entitys.DeviceOperationFailLogListResponseObj, 0, len(grpcRes.Data))
	for _, v := range grpcRes.Data {
		list = append(list, entitys.DeviceOperationFailLogListResponseObj_pb2e(v))
	}
	iotgin.ResPageSuccess(c, list, grpcRes.Total, int(req.Page))
}
