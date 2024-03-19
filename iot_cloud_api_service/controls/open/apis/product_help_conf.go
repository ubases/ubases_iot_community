package apis

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"time"

	goerrors "go-micro.dev/v4/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gin-gonic/gin"
)

var ProHelpConfController ProductHelpConfController

type ProductHelpConfController struct{}

func (ProductHelpConfController) EditProductHelpConf(c *gin.Context) {
	var conf entitys.ProductHelpConfEntitys
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind product help conf param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpConfParam, nil)
		return
	}
	if conf.ProductKey == "" {
		iotlogger.LogHelper.Helper.Error("product key is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudProductKeyEmpty, nil)
		return
	}
	tenantId := controls.GetTenantId(c)
	if tenantId == "" {
		iotlogger.LogHelper.Helper.Error("get tenantid is empty from token")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudTenantIdEmpty, nil)
		return
	}
	req, err := entitys.ProductHelpConf_e2pb(&conf)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("json marshal product help conf param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpConfParam, nil)
		return
	}
	req.UpdatedAt = timestamppb.New(time.Now())
	_, err = rpc.ClientProductHelpConfService.Update(controls.WithOpenUserContext(c), req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("rpc call product help conf server error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

func (ProductHelpConfController) SetProductHelpConf(c *gin.Context) {
	var conf entitys.ProductHelpConfEntitys
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind product help conf param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpConfParam, nil)
		return
	}
	req := &protosService.ProductHelpConfUpdateFieldsRequest{
		Fields: []string{"status", "updated_at"},
		Data: &protosService.ProductHelpConf{
			Id:        iotutil.ToInt64(conf.Id),
			Status:    conf.Status,
			UpdatedAt: timestamppb.New(time.Now()),
		},
	}
	_, err = rpc.ClientProductHelpConfService.UpdateFields(controls.WithOpenUserContext(c), req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("rpc call product help conf server error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

func (ProductHelpConfController) GetProductHelpConf(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotlogger.LogHelper.Helper.Error("query product help conf id param is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpConfParam, nil)
		return
	}
	req := &protosService.ProductHelpConfFilter{
		Id: iotutil.ToInt64(id),
	}
	resp, err := rpc.ClientProductHelpConfService.Find(controls.WithOpenUserContext(c), req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("rpc call product help conf server error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	data, err := entitys.ProductHelpConf_pb2e(resp.Data[0])
	if err != nil {
		iotlogger.LogHelper.Helper.Error("json unmarshal product help conf param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpConfParam, nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, data)
}

func (ProductHelpConfController) GetProductHelpConfList(c *gin.Context) {
	var conf entitys.ProductHelpConfQuery
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind product help conf param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpConfParam, nil)
		return
	}
	tenantId := controls.GetTenantId(c)
	if tenantId == "" {
		iotlogger.LogHelper.Helper.Error("get tenantid is empty from token")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudTenantIdEmpty, nil)
		return
	}
	req := &protosService.ProductHelpConfListRequest{
		Page:      int64(conf.Page),
		PageSize:  int64(conf.Limit),
		OrderKey:  "updated_at",
		OrderDesc: "desc",
		Query: &protosService.ProductHelpConf{
			TenantId: tenantId,
		},
	}
	if conf.Query != nil {
		req.Query.ProductName = conf.Query.ProductName
		if conf.Query.ProductTypeId != "" {
			req.Query.ProductTypeId = iotutil.ToInt64(conf.Query.ProductTypeId)
		}
		req.Query.Status = conf.Query.Status
	}
	resp, err := rpc.ClientProductHelpConfService.Lists(controls.WithOpenUserContext(c), req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("rpc call product help conf server error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	datas := []*entitys.ProductHelpConfEntitys{}
	for _, v := range resp.Data {
		data, err := entitys.ProductHelpConf_pb2e(v)
		if err != nil {
			iotlogger.LogHelper.Helper.Error("json unmarshal product help conf param error: ", err)
			ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpConfParam, nil)
			return
		}
		datas = append(datas, data)
	}
	ioterrs.ResponsePage(c, cached.RedisStore, ioterrs.Success, datas, resp.Total, int(conf.Page))
}
