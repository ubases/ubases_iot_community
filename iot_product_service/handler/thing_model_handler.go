package handler

import (
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_model/db_product/orm"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"

	"github.com/mitchellh/mapstructure"
	"go-micro.dev/v4"
	"gorm.io/gen/field"
)

// The Register tPmThingModel handler.
func RegisterTPmThingModelHandler(service micro.Service) error {
	err := protosService.RegisterTPmThingModelHandler(service.Server(), new(TPmThingModelHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterTPmThingModelHandler发生错误:%s", err.Error())
	}
	return err
}

type TPmThingModelHandler struct{}

// ListTPmThingModel query list by paging
func (TPmThingModelHandler) ListTPmThingModel(ctx context.Context, request *protosService.TPmThingModelFilterPage, response *protosService.TPmThingModelResponseList) error {
	iotlogger.LogHelper.Info("Received Handler.ListTPmThingModel request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModel
	do := t.WithContext(context.Background())
	// 判断参数进行查询
	if request.QueryObj != nil {
		if request.QueryObj.Id != 0 {
			do = do.Where(t.Id.Eq(int64(request.QueryObj.Id)))
		}
	}
	//if request.SearchKey != "" {
	//    do = do.Or(t.Name.Like(request.QueryObj.Remark), t.Remark.Like(request.QueryObj.Remark))
	//}
	var (
		list  []*model.TPmThingModel
		count int64
		err   error
	)
	hasPage := request.Page != 0 || request.Limit != 0
	if hasPage {
		offset := request.Limit * (request.Page - 1)
		list, count, err = do.FindByPage(int(offset), int(request.Limit))
	} else {
		list, err = do.Find()
	}
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "ListTPmThingModel").Error(err)
		return err
	}
	response.Code = 200
	response.Total = count
	response.List = make([]*protosService.TPmThingModelRequest, len(list))
	for i, _ := range list {
		//此方法效率较低，需要尽快找到替换方式
		mapstructure.WeakDecode(list[i], &response.List[i])
	}
	return nil
}

// CreateTPmThingModel create
func (s *TPmThingModelHandler) CreateTPmThingModel(ctx context.Context, request *protosService.TPmThingModelRequest, response *protosService.TPmThingModelResponse) error {
	saveObj, err := s.SaveTPmThingModel(orm.Use(iotmodel.GetDB()), ctx, request)
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "CreateTPmThingModel").Error(err)
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmThingModelRequest{Id: saveObj.Id}
	return nil
}

// SaveTPmThingModel create
func (TPmThingModelHandler) SaveTPmThingModel(tx *orm.Query, ctx context.Context, request *protosService.TPmThingModelRequest) (*model.TPmThingModel, error) {
	iotlogger.LogHelper.Info("Received Handler.CreateTPmThingModel request")
	t := tx.TPmThingModel
	do := t.WithContext(context.Background())
	// 判断参数

	//参数判断 request
	if iotutil.IsEmpty(request.Id) {
		return nil, errors.New("id 主键（雪花算法19位） is null")
	}
	if iotutil.IsEmpty(request.ProductTypeId) {
		return nil, errors.New("product_type_id 产品类型ID is null")
	}
	if iotutil.IsEmpty(request.Standard) {
		return nil, errors.New("standard 是否标准品类 0-否 1-是 is null")
	}
	if iotutil.IsEmpty(request.Version) {
		return nil, errors.New("version 物模型版本号 is null")
	}
	if iotutil.IsEmpty(request.CreatedBy) {
		return nil, errors.New("created_by 创建人 is null")
	}
	if iotutil.IsEmpty(request.CreatedTime) {
		return nil, errors.New("created_time 创建时间 is null")
	}
	if iotutil.IsEmpty(request.UpdatedBy) {
		return nil, errors.New("updated_by 修改人 is null")
	}
	if iotutil.IsEmpty(request.UpdatedTime) {
		return nil, errors.New("updated_at 修改时间 is null")
	}

	// 赋值参数赋值
	var saveObj = model.TPmThingModel{
		ProductKey:    request.ProductKey,
		ProductTypeId: request.ProductTypeId,
		Standard:      request.Standard,
		Version:       request.Version,
		Description:   request.Description,
		CreatedBy:     request.CreatedBy,
	}
	saveObj.Id = iotutil.GetNextSeqInt64()
	err := do.Create(&saveObj)
	if err != nil {
		iotlogger.LogHelper.WithTag("func", "CreateTPmThingModel").Error(err)
		return nil, err
	}
	return &saveObj, nil
}

// UpdateTPmThingModel update
func (TPmThingModelHandler) UpdateTPmThingModel(ctx context.Context, request *protosService.TPmThingModelRequest, response *protosService.TPmThingModelResponse) error {
	iotlogger.LogHelper.Info("Received Handler.UpdateTPmThingModel request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModel
	do := t.WithContext(context.Background())

	// 赋值参数赋值
	var updateObj = model.TPmThingModel{
		ProductKey:    request.ProductKey,
		ProductTypeId: request.ProductTypeId,
		Standard:      request.Standard,
		Version:       request.Version,
		Description:   request.Description,
		//CreatedTime:   request.CreatedTime,
		UpdatedBy: request.UpdatedBy,
		//UpdatedTime:   request.UpdatedTime,
		//Deleted:       request.Deleted,
	}
	updateObj.Id = request.Id
	var updateField = []field.Expr{
		t.ProductKey, t.ProductTypeId, t.Standard, t.Version, t.Description, t.UpdatedBy,
	}
	_, err := do.Select(updateField...).Where(t.Id.Eq(int64(request.Id))).Updates(updateObj)
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "UpdateTPmThingModel").Error(err)
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmThingModelRequest{Id: updateObj.Id}
	return nil
}

// DeleteTPmThingModel delete
func (TPmThingModelHandler) DeleteTPmThingModel(ctx context.Context, request *protosService.TPmThingModelRequest, response *protosService.TPmThingModelResponse) error {
	iotlogger.LogHelper.Info("Received Handler.DeleteTPmThingModel request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModel
	do := t.WithContext(context.Background())
	// 检查数据是否存在
	_, err := do.Where(t.Id.Eq(int64(request.Id))).Delete()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "DeleteTPmThingModel").Error(err)
		return err
	}
	response.Code = 200
	return nil
}

// GetByIdTPmThingModel query struct by id
func (TPmThingModelHandler) GetByIdTPmThingModel(ctx context.Context, request *protosService.TPmThingModelFilterById, response *protosService.TPmThingModelResponseObject) error {
	iotlogger.LogHelper.Info("Received Handler.GetByIdTPmThingModel request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModel
	do := t.WithContext(context.Background())
	//创建查询条件
	info, err := do.Where(t.Id.Eq(int64(request.Id))).First()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetByIdTPmThingModel").Error(err)
		return err
	}
	response.Code = 200
	response.Data.Id = info.Id
	response.Data.ProductKey = info.ProductKey
	response.Data.ProductTypeId = info.ProductTypeId
	response.Data.Standard = info.Standard
	response.Data.Version = info.Version
	response.Data.Description = info.Description
	//response.Data.CreatedTime = info.CreatedTime
	response.Data.UpdatedBy = info.UpdatedBy
	//response.Data.UpdatedTime = info.UpdatedTime
	//response.Data.Deleted = info.Deleted

	// mapstructure.WeakDecode(info, &response.Data)
	return nil
}

// GetTPmThingModel query struct by struct
func (TPmThingModelHandler) GetTPmThingModel(ctx context.Context, request *protosService.TPmThingModelFilter, response *protosService.TPmThingModelResponseObject) error {
	iotlogger.LogHelper.Info("Received Handler.GetTPmThingModel request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModel
	do := t.WithContext(context.Background())
	// 判断参数进行查询
	if request.Id != 0 {
		do = do.Where(t.Id.Eq(request.Id))
	}
	if request.ProductTypeId != 0 {
		do = do.Where(t.ProductTypeId.Eq(request.ProductTypeId))
	}
	if request.Standard >= 0 {
		do = do.Where(t.Standard.Eq(request.Standard))
	}
	if !iotutil.IsEmpty(request.Version) {
		do = do.Where(t.Version.Eq(request.Version))
	}
	if !iotutil.IsEmpty(request.ProductKey) {
		do = do.Where(t.ProductKey.Eq(request.ProductKey))
	}

	info, err := do.First()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetTPmThingModel").Error(err)
		return err
	}
	response.Code = 200

	//此方法效率较低，需要尽快找到替换方式
	mapstructure.WeakDecode(info, &response.Data)
	return nil
}
