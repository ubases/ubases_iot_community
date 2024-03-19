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

var ProHelpDocController ProductHelpDocController

type ProductHelpDocController struct{}

func (ProductHelpDocController) AddProductHelpDoc(c *gin.Context) {
	var conf entitys.ProductHelpDocEntitys
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind product help doc param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpDocParam, nil)
		return
	}
	if conf.ProductKey == "" {
		iotlogger.LogHelper.Helper.Error("product help doc productkey is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudProductKeyEmpty, nil)
		return
	}
	conf.Id = iotutil.ToString(iotutil.GetNextSeqInt64())
	conf.RelationId = iotutil.ToString(iotutil.GetNextSeqInt64())
	req := entitys.ProductHelpDoc_e2pb(&conf)
	tenantId := controls.GetTenantId(c)
	if tenantId == "" {
		iotlogger.LogHelper.Helper.Error("get tenantid is empty from token")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudTenantIdEmpty, nil)
		return
	}
	req.TenantId = tenantId
	req.CreatedAt = timestamppb.New(time.Now())
	req.UpdatedAt = timestamppb.New(time.Now())
	_, err = rpc.ClientProductHelpDocService.Create(controls.WithOpenUserContext(c), req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("create product help doc error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	// 创建文档后，需要同步更新产品配置的更新时间
	resp, err := rpc.ClientProductHelpConfService.Find(controls.WithOpenUserContext(c), &protosService.ProductHelpConfFilter{
		TenantId:   tenantId,
		ProductKey: conf.ProductKey,
	})
	if err != nil {
		iotlogger.LogHelper.Helper.Error("rpc call product help conf server error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	reqConf := &protosService.ProductHelpConf{
		Id:        resp.Data[0].Id,
		CreatedAt: resp.Data[0].CreatedAt,
		UpdatedAt: timestamppb.New(time.Now()),
	}
	_, err = rpc.ClientProductHelpConfService.Update(controls.WithOpenUserContext(c), reqConf)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("rpc call product help conf server error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

func (ProductHelpDocController) EditProductHelpDoc(c *gin.Context) {
	var conf entitys.ProductHelpDocEntitys
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind product help doc param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpDocParam, nil)
		return
	}
	if conf.ProductKey == "" {
		iotlogger.LogHelper.Helper.Error("product help doc productkey is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudProductKeyEmpty, nil)
		return
	}
	tenantId := controls.GetTenantId(c)
	if tenantId == "" {
		iotlogger.LogHelper.Helper.Error("get tenantid is empty from token")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudTenantIdEmpty, nil)
		return
	}
	// 判断是更新现有语种文档，还是新增语种文档
	reqFind := &protosService.ProductHelpDocFilter{
		RelationId: iotutil.ToInt64(conf.RelationId),
		Lang:       conf.Lang,
	}
	respFind, err := rpc.ClientProductHelpDocService.Find(controls.WithOpenUserContext(c), reqFind)
	if err != nil && goerrors.FromError(err).GetDetail() != "record not found" {
		iotlogger.LogHelper.Helper.Error("find product help doc error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	if respFind != nil && respFind.Data != nil && len(respFind.Data) != 0 {
		reqUpdate := &protosService.ProductHelpDocUpdateFieldsRequest{
			Fields: []string{"title", "content", "updated_at"},
			Data: &protosService.ProductHelpDoc{
				Id:        iotutil.ToInt64(conf.Id),
				Title:     conf.Title,
				Content:   conf.Content,
				UpdatedAt: timestamppb.New(time.Now()),
			},
		}
		_, err := rpc.ClientProductHelpDocService.UpdateFields(controls.WithOpenUserContext(c), reqUpdate)
		if err != nil {
			iotlogger.LogHelper.Helper.Error("update product help doc error: ", err)
			ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
			return
		}
	} else {
		conf.Id = iotutil.ToString(iotutil.GetNextSeqInt64())
		req := entitys.ProductHelpDoc_e2pb(&conf)
		req.TenantId = tenantId
		req.CreatedAt = timestamppb.New(time.Now())
		req.UpdatedAt = timestamppb.New(time.Now())
		_, err := rpc.ClientProductHelpDocService.Create(controls.WithOpenUserContext(c), req)
		if err != nil {
			iotlogger.LogHelper.Helper.Error("create product help doc error: ", err)
			ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
			return
		}
	}
	// 创建文档后，需要同步更新产品配置的更新时间
	resp, err := rpc.ClientProductHelpConfService.Find(controls.WithOpenUserContext(c), &protosService.ProductHelpConfFilter{
		TenantId:   tenantId,
		ProductKey: conf.ProductKey,
	})
	if err != nil {
		iotlogger.LogHelper.Helper.Error("rpc call product help conf server error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	reqConf := &protosService.ProductHelpConf{
		Id:        resp.Data[0].Id,
		CreatedAt: resp.Data[0].CreatedAt,
		UpdatedAt: timestamppb.New(time.Now()),
	}
	_, err = rpc.ClientProductHelpConfService.Update(controls.WithOpenUserContext(c), reqConf)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("rpc call product help conf server error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

func (ProductHelpDocController) SetProductHelpDoc(c *gin.Context) {
	var conf entitys.ProductHelpDocEntitys
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind product help doc param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpDocParam, nil)
		return
	}
	req := &protosService.ProductHelpDocUpdateFieldsRequest{
		Fields: []string{"sort_id", "status", "updated_at"},
		Data: &protosService.ProductHelpDoc{
			RelationId: iotutil.ToInt64(conf.RelationId),
			SortId:     conf.SortId,
			Status:     conf.Status,
			UpdatedAt:  timestamppb.New(time.Now()),
		},
	}
	_, err = rpc.ClientProductHelpDocService.UpdateFields(controls.WithOpenUserContext(c), req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("set product help doc error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

func (ProductHelpDocController) DeleteProductHelpDoc(c *gin.Context) {
	var conf entitys.ProductHelpDocEntitys
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind product help doc param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpDocParam, nil)
		return
	}
	req := &protosService.ProductHelpDoc{
		RelationId: iotutil.ToInt64(conf.RelationId),
	}
	_, err = rpc.ClientProductHelpDocService.Delete(controls.WithOpenUserContext(c), req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("delete product help doc error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

func (ProductHelpDocController) GetProductHelpDoc(c *gin.Context) {
	relationId := c.Query("relationId")
	if relationId == "" {
		iotlogger.LogHelper.Helper.Error("query product help conf relation id param is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpConfParam, nil)
		return
	}
	lang := c.Query("lang")
	if lang == "" {
		iotlogger.LogHelper.Helper.Error("query product help conf lang param is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpConfParam, nil)
		return
	}
	req := &protosService.ProductHelpDocListRequest{
		Query: &protosService.ProductHelpDoc{
			RelationId: iotutil.ToInt64(relationId),
		},
		OrderKey: "sort_id",
	}
	resp, err := rpc.ClientProductHelpDocService.Lists(controls.WithOpenUserContext(c), req)
	if err != nil && goerrors.FromError(err).GetDetail() != "record not found" {
		iotlogger.LogHelper.Helper.Error("get product help doc error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	if resp != nil && resp.Data != nil && len(resp.Data) != 0 {
		for i := range resp.Data {
			if resp.Data[i].Lang == lang {
				ioterrs.Response(c, cached.RedisStore, ioterrs.Success, entitys.ProductHelpDoc_pb2e(resp.Data[i]))
				return
			}
		}
		data := &entitys.ProductHelpDocEmpty{
			TenantId:   resp.Data[0].TenantId,
			ProductKey: resp.Data[0].ProductKey,
			Lang:       lang,
			Title:      "",
			Content:    "",
			RelationId: relationId,
			SortId:     resp.Data[0].SortId,
			Status:     resp.Data[0].Status,
			CreatedAt:  resp.Data[0].CreatedAt.AsTime().Unix(),
			UpdatedAt:  resp.Data[0].UpdatedAt.AsTime().Unix(),
		}
		ioterrs.Response(c, cached.RedisStore, ioterrs.Success, data)
		return
	}
	data := &entitys.ProductHelpDocEmpty{
		RelationId: relationId,
		Lang:       lang,
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, data)
}

func (ProductHelpDocController) GetProductHelpDocList(c *gin.Context) {
	var conf entitys.ProductHelpDocQuery
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind product help doc param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrProductHelpDocParam, nil)
		return
	}
	if conf.Query == nil {
		iotlogger.LogHelper.Helper.Error("product help doc param is nil")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudQueryParamIsNil, nil)
		return
	}
	if conf.Query.ProductKey == "" {
		iotlogger.LogHelper.Helper.Error("product help doc productkey is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudProductKeyEmpty, nil)
		return
	}
	tenantId := controls.GetTenantId(c)
	if tenantId == "" {
		iotlogger.LogHelper.Helper.Error("get tenantid is empty from token")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudTenantIdEmpty, nil)
		return
	}
	req := &protosService.ProductHelpDocListRequest{
		Page:     int64(conf.Page),
		PageSize: int64(conf.Limit),
		Query: &protosService.ProductHelpDoc{
			TenantId:   tenantId,
			ProductKey: conf.Query.ProductKey,
			Lang:       conf.Query.Lang,
			Title:      conf.Query.Title,
			Status:     conf.Query.Status,
		},
		OrderKey: "sort_id",
	}
	resp, err := rpc.ClientProductHelpDocService.Lists(controls.WithOpenUserContext(c), req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("get product help doc list error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	datas := []*entitys.ProductHelpDocEntitys{}
	for _, v := range resp.Data {
		data := entitys.ProductHelpDoc_pb2e(v)
		datas = append(datas, data)
	}
	ioterrs.ResponsePage(c, cached.RedisStore, ioterrs.Success, datas, resp.Total, int(conf.Page))
}
