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
	"time"

	"github.com/mitchellh/mapstructure"
	"go-micro.dev/v4"
	"gorm.io/gen/field"
)

// The Register tPmThingModelServices handler.
func RegisterTPmThingModelServicesHandler(service micro.Service) error {
	err := protosService.RegisterTPmThingModelServicesHandler(service.Server(), new(TPmThingModelServicesHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterTPmThingModelServicesHandler发生错误:%s", err.Error())
	}
	return err
}

type TPmThingModelServicesHandler struct{}

// GetByIdTPmThingModelServices query struct by id
func (TPmThingModelServicesHandler) GetByIdTPmThingModelServices(ctx context.Context, request *protosService.TPmThingModelServicesFilterById, response *protosService.TPmThingModelServicesResponseObject) error {
	iotlogger.LogHelper.Info("Received Handler.GetByIdTPmThingModelServices request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelServices
	do := t.WithContext(context.Background())
	//创建查询条件
	info, err := do.Where(t.Id.Eq(int64(request.Id))).First()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetByIdTPmThingModelServices").Error(err)
		return err
	}
	response.Code = 200
	response.Data.Id = info.Id
	response.Data.ModelId = info.ModelId
	response.Data.ProductKey = info.ProductKey
	response.Data.Identifier = info.Identifier
	response.Data.ServiceName = info.ServiceName
	response.Data.InputParams = info.InputParams
	response.Data.OutputParams = info.OutputParams
	response.Data.Required = info.Required
	response.Data.CallType = info.CallType
	response.Data.Custom = info.Custom
	response.Data.Extension = info.Extension
	//response.Data.CreatedTime = info.CreatedTime
	response.Data.UpdatedBy = info.UpdatedBy
	//response.Data.UpdatedTime = info.UpdatedTime
	//response.Data.Deleted = info.Deleted
	response.Data.Desc = info.Desc
	response.Data.TriggerCond = info.TriggerCond
	response.Data.ExecCond = info.ExecCond

	// mapstructure.WeakDecode(info, &response.Data)
	return nil
}

// GetTPmThingModelServices query struct by struct
func (TPmThingModelServicesHandler) GetTPmThingModelServices(ctx context.Context, request *protosService.TPmThingModelServicesFilter, response *protosService.TPmThingModelServicesResponseObject) error {
	iotlogger.LogHelper.Info("Received Handler.GetTPmThingModelServices request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelServices
	do := t.WithContext(context.Background())
	// 判断参数进行查询
	if request.Id != 0 {
		do = do.Where(t.Id.Eq(int64(request.Id)))
	}

	info, err := do.First()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetTPmThingModelServices").Error(err)
		return err
	}
	response.Code = 200

	response.Data.Id = info.Id
	response.Data.ModelId = info.ModelId
	response.Data.ProductKey = info.ProductKey
	response.Data.Identifier = info.Identifier
	response.Data.ServiceName = info.ServiceName
	response.Data.InputParams = info.InputParams
	response.Data.OutputParams = info.OutputParams
	response.Data.Required = info.Required
	response.Data.CallType = info.CallType
	response.Data.Custom = info.Custom
	response.Data.Extension = info.Extension
	//response.Data.CreatedTime = info.CreatedTime
	response.Data.UpdatedBy = info.UpdatedBy
	//response.Data.UpdatedTime = info.UpdatedTime
	//response.Data.Deleted = info.Deleted
	response.Data.Desc = info.Desc
	response.Data.TriggerCond = info.TriggerCond
	response.Data.ExecCond = info.ExecCond

	//此方法效率较低，需要尽快找到替换方式
	//mapstructure.WeakDecode(info, &response.Data)
	return nil
}

// ListTPmThingModelServices query list by paging
func (TPmThingModelServicesHandler) ListTPmThingModelServices(ctx context.Context, request *protosService.TPmThingModelServicesFilterPage, response *protosService.TPmThingModelServicesResponseList) error {
	iotlogger.LogHelper.Info("Received Handler.ListTPmThingModelServices request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelServices
	do := t.WithContext(context.Background())
	// 判断参数进行查询
	if request.QueryObj != nil {
		if request.QueryObj.Id != 0 {
			do = do.Where(t.Id.Eq(request.QueryObj.Id))
		}
		if len(request.QueryObj.Ids) > 0 {
			do = do.Where(t.Id.In(request.QueryObj.Ids...))
		}
		if request.QueryObj.ModelId > 0 {
			do = do.Where(t.ModelId.Eq(request.QueryObj.ModelId))
		}
		if request.QueryObj.ProductKey != "" {
			do = do.Where(t.ProductKey.Eq(request.QueryObj.ProductKey))
		}
	}
	//if request.SearchKey != "" {
	//    do = do.Or(t.Name.Like(request.QueryObj.Remark), t.Remark.Like(request.QueryObj.Remark))
	//}
	var (
		list  []*model.TPmThingModelServices
		count int64
		err   error
	)
	hasPage := request.Page != 0 || request.Limit != 0
	if hasPage {
		offset := request.Limit * (request.Page - 1)
		list, count, err = do.Order(t.Dpid).FindByPage(int(offset), int(request.Limit))
	} else {
		list, err = do.Order(t.Dpid).Find()
	}
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "ListTPmThingModelServices").Error(err)
		return err
	}
	response.Code = 200
	response.Total = count
	response.List = make([]*protosService.TPmThingModelServicesRequest, len(list))
	for i, _ := range list {
		//此方法效率较低，需要尽快找到替换方式
		mapstructure.WeakDecode(list[i], &response.List[i])
	}
	return nil
}

// CreateTPmThingModelServices create
func (TPmThingModelServicesHandler) CreateTPmThingModelServices(ctx context.Context, request *protosService.TPmThingModelServicesRequest, response *protosService.TPmThingModelServicesResponse) error {
	iotlogger.LogHelper.Info("Received Handler.CreateTPmThingModelServices request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelServices
	do := t.WithContext(context.Background())
	// 判断参数

	//参数判断 request
	if iotutil.IsEmpty(request.Id) {
		return errors.New("id 主键（雪花算法19位） is null")
	}
	if iotutil.IsEmpty(request.ModelId) {
		return errors.New("model_id 物模型ID(t_pm_thing_model.id) is null")
	}
	if iotutil.IsEmpty(request.ProductKey) {
		return errors.New("product_key 产品唯一标识 is null")
	}
	if iotutil.IsEmpty(request.CreateTs) {
		return errors.New("create_ts 功能创建的时间戳，默认长度是13位。可手动传入也可由系统生成。功能定义会根据该时间由小到大进行排序。 is null")
	}
	if iotutil.IsEmpty(request.Identifier) {
		return errors.New("identifier 属性的标识符。可包含大小写英文字母、数字、下划线（_），长度不超过50个字符。 is null")
	}
	if iotutil.IsEmpty(request.ServiceName) {
		return errors.New("service_name 服务名称。可包含中文、大小写英文字母、数字、短划线（-）、下划线（_）和半角句号（.），且必须以中文、英文字母或数字开头，长度不超过30个字符，一个中文计为一个字符。 is null")
	}
	if iotutil.IsEmpty(request.Required) {
		return errors.New("required 是否是标准品类的必选服务。1：是, 0：否 is null")
	}
	if iotutil.IsEmpty(request.CallType) {
		return errors.New("call_type 服务的调用方式。1：异步调用, 0：同步调用 is null")
	}
	if iotutil.IsEmpty(request.Custom) {
		return errors.New("custom 是否是自定义功能。1：是, 0：否 is null")
	}
	if iotutil.IsEmpty(request.Extension) {
		return errors.New("extension 扩展字段 is null")
	}
	if iotutil.IsEmpty(request.CreatedBy) {
		return errors.New("created_by 创建人 is null")
	}
	if iotutil.IsEmpty(request.CreatedTime) {
		return errors.New("created_time 创建时间 is null")
	}
	if iotutil.IsEmpty(request.UpdatedBy) {
		return errors.New("updated_by 修改人 is null")
	}
	if iotutil.IsEmpty(request.UpdatedTime) {
		return errors.New("updated_at 修改时间 is null")
	}

	// 赋值参数赋值
	var saveObj = model.TPmThingModelServices{
		ModelId:      request.ModelId,
		ProductKey:   request.ProductKey,
		Identifier:   request.Identifier,
		ServiceName:  request.ServiceName,
		InputParams:  request.InputParams,
		OutputParams: request.OutputParams,
		Required:     request.Required,
		CallType:     request.CallType,
		Custom:       request.Custom,
		Extension:    request.Extension,
		CreatedBy:    request.CreatedBy,
		//CreatedTime:  request.CreatedTime,
		//UpdatedTime:  request.UpdatedTime,
		//Deleted:      request.Deleted,
		Desc:        request.Desc,
		TriggerCond: request.TriggerCond,
		ExecCond:    request.ExecCond,
	}
	saveObj.Id = iotutil.GetNextSeqInt64()
	err := do.Create(&saveObj)
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "CreateTPmThingModelServices").Error(err)
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmThingModelServicesRequest{Id: saveObj.Id}
	return nil
}

// UpdateTPmThingModelServices update
func (TPmThingModelServicesHandler) UpdateTPmThingModelServices(ctx context.Context, request *protosService.TPmThingModelServicesRequest, response *protosService.TPmThingModelServicesResponse) error {
	iotlogger.LogHelper.Info("Received Handler.UpdateTPmThingModelServices request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelServices
	do := t.WithContext(context.Background())

	// 赋值参数赋值
	var updateObj = model.TPmThingModelServices{
		ModelId:      request.ModelId,
		ProductKey:   request.ProductKey,
		Identifier:   request.Identifier,
		ServiceName:  request.ServiceName,
		InputParams:  request.InputParams,
		OutputParams: request.OutputParams,
		Required:     request.Required,
		CallType:     request.CallType,
		Custom:       request.Custom,
		Extension:    request.Extension,
		//CreatedTime:  request.CreatedTime,
		UpdatedBy: request.UpdatedBy,
		//UpdatedTime:  request.UpdatedTime,
		//Deleted:      request.Deleted,
		Desc:        request.Desc,
		TriggerCond: request.TriggerCond,
		ExecCond:    request.ExecCond,
	}
	updateObj.Id = request.Id
	var updateField = []field.Expr{
		t.ModelId, t.ProductKey, t.Identifier, t.ServiceName, t.InputParams, t.OutputParams, t.Required, t.CallType, t.Custom, t.Extension, t.UpdatedBy,
	}
	_, err := do.Select(updateField...).Where(t.Id.Eq(int64(request.Id))).Updates(updateObj)
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "UpdateTPmThingModelServices").Error(err)
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmThingModelServicesRequest{Id: updateObj.Id}
	return nil
}

func (s *TPmThingModelServicesHandler) Update(ctx context.Context, request *protosService.TPmThingModelServicesRequest, response *protosService.TPmThingModelServicesResponse) error {
	s.UpdateInfo(orm.Use(iotmodel.GetDB()), ctx, request, response)
	return nil
}

func (TPmThingModelServicesHandler) UpdateInfo(tx *orm.Query, ctx context.Context, request *protosService.TPmThingModelServicesRequest, response *protosService.TPmThingModelServicesResponse) error {
	iotlogger.LogHelper.Info("Received Handler.UpdateTPmThingModelServices request")
	t := tx.TPmThingModelServices
	do := t.WithContext(context.Background())

	// 赋值参数赋值
	var updateObj = model.TPmThingModelServices{
		TriggerCond: request.TriggerCond,
		ExecCond:    request.ExecCond,
		Valid:       request.Valid,
	}
	updateObj.Id = request.Id
	var updateField = []field.Expr{
		t.TriggerCond, t.ExecCond, t.Valid,
	}
	_, err := do.Select(updateField...).Where(t.Id.Eq(int64(request.Id))).Updates(updateObj)
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "UpdateTPmThingModelServices").Error(err)
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmThingModelServicesRequest{Id: updateObj.Id}
	return nil
}

// DeleteTPmThingModelServices delete
func (TPmThingModelServicesHandler) DeleteTPmThingModelServices(ctx context.Context, request *protosService.TPmThingModelServicesRequest, response *protosService.TPmThingModelServicesResponse) error {
	iotlogger.LogHelper.Info("Received Handler.DeleteTPmThingModelServices request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelServices
	do := t.WithContext(context.Background())

	// 判断参数进行查询
	var (
		err  error
		flag = false
	)
	if request.ModelId > 0 {
		do = do.Where(t.ModelId.Eq(request.ModelId))
		flag = true
	}
	if !iotutil.IsEmpty(request.ProductKey) {
		do = do.Where(t.ProductKey.Eq(request.ProductKey))
		flag = true
	}
	if request.Id > 0 {
		do = do.Where(t.Id.Eq(request.Id))
		flag = true
	}
	if flag {
		_, err = do.Unscoped().Delete()
	} else {
		_, err = do.Where(t.Id.Eq(request.Id)).Unscoped().Delete()
	}
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "DeleteTPmThingModelServices").Error(err)
		return err
	}
	response.Code = 200
	return nil
}

// CreateTPmThingModelServicesBatch batch create
func (TPmThingModelServicesHandler) CreateTPmThingModelServicesBatch(ctx context.Context, requests []*protosService.TPmThingModelServicesRequest, response *protosService.TPmThingModelServicesResponse) error {
	iotlogger.LogHelper.Info("Received Handler.CreateTPmThingModelServicesBatch request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelServices
	do := t.WithContext(context.Background())

	//参数判断 requests
	var (
		saveObjs = make([]*model.TPmThingModelServices, len(requests))
		ids      = make([]int64, len(requests))
	)
	for index, request := range requests {
		if iotutil.IsEmpty(request.Id) {
			return errors.New("id 主键（雪花算法19位） is null")
		}
		if iotutil.IsEmpty(request.ModelId) {
			return errors.New("model_id 物模型ID(t_pm_thing_model.id) is null")
		}
		if iotutil.IsEmpty(request.ProductKey) {
			return errors.New("product_key 产品唯一标识 is null")
		}
		if iotutil.IsEmpty(request.CreateTs) {
			return errors.New("create_ts 功能创建的时间戳，默认长度是13位。可手动传入也可由系统生成。功能定义会根据该时间由小到大进行排序。 is null")
		}
		if iotutil.IsEmpty(request.Identifier) {
			return errors.New("identifier 属性的标识符。可包含大小写英文字母、数字、下划线（_），长度不超过50个字符。 is null")
		}
		if iotutil.IsEmpty(request.ServiceName) {
			return errors.New("service_name 服务名称。可包含中文、大小写英文字母、数字、短划线（-）、下划线（_）和半角句号（.），且必须以中文、英文字母或数字开头，长度不超过30个字符，一个中文计为一个字符。 is null")
		}
		if iotutil.IsEmpty(request.Required) {
			return errors.New("required 是否是标准品类的必选服务。1：是, 0：否 is null")
		}
		if iotutil.IsEmpty(request.CallType) {
			return errors.New("call_type 服务的调用方式。1：异步调用, 0：同步调用 is null")
		}
		if iotutil.IsEmpty(request.Custom) {
			return errors.New("custom 是否是自定义功能。1：是, 0：否 is null")
		}
		if iotutil.IsEmpty(request.CreatedBy) {
			return errors.New("created_by 创建人 is null")
		}
		if iotutil.IsEmpty(request.CreatedTime) {
			return errors.New("created_time 创建时间 is null")
		}
		if iotutil.IsEmpty(request.UpdatedBy) {
			return errors.New("updated_by 修改人 is null")
		}
		if iotutil.IsEmpty(request.UpdatedTime) {
			return errors.New("updated_at 修改时间 is null")
		}

		// 赋值参数赋值
		var saveObj = &model.TPmThingModelServices{
			ModelId:      request.ModelId,
			ProductKey:   request.ProductKey,
			Identifier:   request.Identifier,
			ServiceName:  request.ServiceName,
			InputParams:  request.InputParams,
			OutputParams: request.OutputParams,
			Required:     request.Required,
			CallType:     request.CallType,
			Custom:       request.Custom,
			Extension:    request.Extension,
			CreatedBy:    request.CreatedBy,
			//CreatedTime:  request.CreatedTime,
			//UpdatedTime:  request.UpdatedTime,
			//Deleted:      request.Deleted,
			Desc:        request.Desc,
			TriggerCond: request.TriggerCond,
			ExecCond:    request.ExecCond,
		}
		saveObj.Id = iotutil.GetNextSeqInt64()
		saveObjs[index] = saveObj
		ids[index] = saveObj.Id
	}

	err := do.CreateInBatches(saveObjs, len(saveObjs))
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "CreateTPmThingModelServicesBatch").Error(err)
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmThingModelServicesRequest{Ids: ids}
	return nil
}

func UpdateInvalidForModelServices(tx *orm.Query, productModelId int64, productKey string) error {
	if productModelId == 0 || productKey == "" {
		return errors.New("缺条件")
	}
	t := tx.TPmThingModelServices
	_, err := t.WithContext(context.Background()).Where(t.ProductKey.Eq(productKey), t.ModelId.Eq(productModelId)).Update(t.Valid, 0)
	return err
}

func CopyTypeToProductForModelServices(tx *orm.Query, typeid int64, productModelId int64, productKey string, mapInfos map[int64]Cond) error {
	ctx := context.Background()
	tm := tx.TPmThingModel
	t := tx.TPmThingModelServices
	list, err := t.WithContext(ctx).Join(tm, tm.Id.EqCol(t.ModelId), tm.ProductTypeId.Eq(typeid), tm.Standard.Eq(1)).Find()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}
	for _, v := range list {
		if info, ok := mapInfos[v.Id]; ok {
			v.TriggerCond = info.TriggerCond
			v.ExecCond = info.ExecCond
			v.Valid = 1
		} else {
			v.Valid = 0
			v.TriggerCond = 0
			v.ExecCond = 0
		}
		v.StdId = v.Id
		v.Id = iotutil.GetNextSeqInt64()
		v.ModelId = productModelId
		v.ProductKey = productKey
		v.CreatedAt = time.Now()
		v.UpdatedAt = v.CreatedAt
	}
	do := t.WithContext(context.Background())
	err = do.CreateInBatches(list, len(list))
	if err != nil {
		return err
	}
	return nil
}
