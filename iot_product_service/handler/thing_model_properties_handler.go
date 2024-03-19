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

// The Register tPmThingModelProperties handler.
func RegisterTPmThingModelPropertiesHandler(service micro.Service) error {
	err := protosService.RegisterTPmThingModelPropertiesHandler(service.Server(), new(TPmThingModelPropertiesHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterTPmThingModelPropertiesHandler发生错误:%s", err.Error())
	}
	return err
}

type TPmThingModelPropertiesHandler struct{}

// UpdateTPmThingModelProperties update
func (TPmThingModelPropertiesHandler) UpdateTPmThingModelProperties(ctx context.Context, request *protosService.TPmThingModelPropertiesRequest, response *protosService.TPmThingModelPropertiesResponse) error {
	iotlogger.LogHelper.Info("Received Handler.UpdateTPmThingModelProperties request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelProperties
	do := t.WithContext(context.Background())

	// 赋值参数赋值
	var updateObj = model.TPmThingModelProperties{
		ModelId:       request.ModelId,
		ProductKey:    request.ProductKey,
		Identifier:    request.Identifier,
		DataType:      request.DataType,
		Name:          request.Name,
		RwFlag:        request.RwFlag,
		DataSpecs:     request.DataSpecs,
		DataSpecsList: request.DataSpecsList,
		Required:      request.Required,
		Custom:        request.Custom,
		Extension:     request.Extension,
		//CreatedTime:   request.CreatedTime,
		UpdatedBy: request.UpdatedBy,
		//UpdatedTime:   request.UpdatedTime,
		//Deleted:       request.Deleted,
		Desc:        request.Desc,
		TriggerCond: request.TriggerCond,
		ExecCond:    request.ExecCond,
	}
	updateObj.Id = request.Id
	var updateField = []field.Expr{
		t.ModelId, t.ProductKey, t.Identifier, t.DataType, t.Name, t.RwFlag, t.DataSpecs, t.DataSpecsList, t.Required, t.Custom, t.Extension, t.UpdatedBy,
	}
	_, err := do.Select(updateField...).Where(t.Id.Eq(int64(request.Id))).Updates(updateObj)
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "UpdateTPmThingModelProperties").Error(err)
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmThingModelPropertiesRequest{Id: updateObj.Id}
	return nil
}

func (s *TPmThingModelPropertiesHandler) Update(ctx context.Context, request *protosService.TPmThingModelPropertiesRequest, response *protosService.TPmThingModelPropertiesResponse) error {
	s.UpdateInfo(orm.Use(iotmodel.GetDB()), ctx, request, response)
	return nil
}

func (TPmThingModelPropertiesHandler) UpdateInfo(tx *orm.Query, ctx context.Context, request *protosService.TPmThingModelPropertiesRequest, response *protosService.TPmThingModelPropertiesResponse) error {
	iotlogger.LogHelper.Info("Received Handler.UpdateTPmThingModelProperties request")
	t := tx.TPmThingModelProperties
	do := t.WithContext(context.Background())

	// 赋值参数赋值
	var updateObj = model.TPmThingModelProperties{
		TriggerCond: request.TriggerCond,
		ExecCond:    request.ExecCond,
		Valid:       request.Valid,
	}
	updateObj.Id = request.Id
	var updateField = []field.Expr{t.TriggerCond, t.ExecCond, t.Valid}
	_, err := do.Select(updateField...).Where(t.Id.Eq(int64(request.Id))).Updates(updateObj)
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "UpdateTPmThingModelProperties").Error(err)
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmThingModelPropertiesRequest{Id: updateObj.Id}
	return nil
}

// DeleteTPmThingModelProperties delete
func (TPmThingModelPropertiesHandler) DeleteTPmThingModelProperties(ctx context.Context, request *protosService.TPmThingModelPropertiesRequest, response *protosService.TPmThingModelPropertiesResponse) error {
	iotlogger.LogHelper.Info("Received Handler.DeleteTPmThingModelProperties request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelProperties
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
		iotlogger.LogHelper.WithTag("func", "DeleteTPmThingModelProperties").Error(err)
		return err
	}
	response.Code = 200
	return nil
}

// GetByIdTPmThingModelProperties query struct by id
func (TPmThingModelPropertiesHandler) GetByIdTPmThingModelProperties(ctx context.Context, request *protosService.TPmThingModelPropertiesFilterById, response *protosService.TPmThingModelPropertiesResponseObject) error {
	iotlogger.LogHelper.Info("Received Handler.GetByIdTPmThingModelProperties request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelProperties
	do := t.WithContext(context.Background())
	//创建查询条件
	info, err := do.Where(t.Id.Eq(int64(request.Id))).First()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetByIdTPmThingModelProperties").Error(err)
		return err
	}
	response.Code = 200
	response.Data.Id = info.Id
	response.Data.ModelId = info.ModelId
	response.Data.ProductKey = info.ProductKey
	response.Data.Identifier = info.Identifier
	response.Data.DataType = info.DataType
	response.Data.Name = info.Name
	response.Data.RwFlag = info.RwFlag
	response.Data.DataSpecs = info.DataSpecs
	response.Data.DataSpecsList = info.DataSpecsList
	response.Data.Required = info.Required
	response.Data.Custom = info.Custom
	response.Data.Extension = info.Extension
	//response.Data.CreatedTime = info.CreatedTime
	response.Data.UpdatedBy = info.UpdatedBy
	//response.Data.UpdatedTime = info.UpdatedTisme
	//response.Data.Deleted = info.Deleted
	response.Data.Desc = info.Desc
	response.Data.TriggerCond = info.TriggerCond
	response.Data.ExecCond = info.ExecCond

	// mapstructure.WeakDecode(info, &response.Data)
	return nil
}

// GetTPmThingModelProperties query struct by struct
func (TPmThingModelPropertiesHandler) GetTPmThingModelProperties(ctx context.Context, request *protosService.TPmThingModelPropertiesFilter, response *protosService.TPmThingModelPropertiesResponseObject) error {
	iotlogger.LogHelper.Info("Received Handler.GetTPmThingModelProperties request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelProperties
	do := t.WithContext(context.Background())
	// 判断参数进行查询
	if request.Id != 0 {
		do = do.Where(t.Id.Eq(int64(request.Id)))
	}

	info, err := do.First()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetTPmThingModelProperties").Error(err)
		return err
	}
	response.Code = 200

	response.Data.Id = info.Id
	response.Data.ModelId = info.ModelId
	response.Data.ProductKey = info.ProductKey
	response.Data.Identifier = info.Identifier
	response.Data.DataType = info.DataType
	response.Data.Name = info.Name
	response.Data.RwFlag = info.RwFlag
	response.Data.DataSpecs = info.DataSpecs
	response.Data.DataSpecsList = info.DataSpecsList
	response.Data.Required = info.Required
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

// ListTPmThingModelProperties query list by paging
func (TPmThingModelPropertiesHandler) ListTPmThingModelProperties(ctx context.Context, request *protosService.TPmThingModelPropertiesFilterPage, response *protosService.TPmThingModelPropertiesResponseList) error {
	iotlogger.LogHelper.Info("Received Handler.ListTPmThingModelProperties request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelProperties
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
		list  []*model.TPmThingModelProperties
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
		iotlogger.LogHelper.WithTag("func", "ListTPmThingModelProperties").Error(err)
		return err
	}
	response.Code = 200
	response.Total = int64(count)
	response.List = make([]*protosService.TPmThingModelPropertiesRequest, len(list))
	for i, _ := range list {
		//此方法效率较低，需要尽快找到替换方式
		mapstructure.WeakDecode(list[i], &response.List[i])
	}
	return nil
}

// CreateTPmThingModelProperties create
func (TPmThingModelPropertiesHandler) CreateTPmThingModelProperties(ctx context.Context, request *protosService.TPmThingModelPropertiesRequest, response *protosService.TPmThingModelPropertiesResponse) error {
	iotlogger.LogHelper.Info("Received Handler.CreateTPmThingModelProperties request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelProperties
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
	if iotutil.IsEmpty(request.DataType) {
		return errors.New("data_type 属性值的数据类型。可选值：ARRAY、STRUCT、INT、FLOAT、DOUBLE、TEXT、DATE、JSON、ENUM、BOOL。不同数据类型，可传入的参数不同。 is null")
	}
	if iotutil.IsEmpty(request.Name) {
		return errors.New("name 属性名称。可包含中文、大小写英文字母、数字、短划线（-）、下划线（_）和半角句号（.），且必须以中文、英文字母或数字开头，长度不超过30个字符，一个中文计为一个字符。 is null")
	}
	if iotutil.IsEmpty(request.RwFlag) {
		return errors.New("rw_flag 在云端可以对该属性进行的操作类型。READ_WRITE：读写。READ_ONLY：只读。WRITE_ONLY：只写。 is null")
	}
	if iotutil.IsEmpty(request.Required) {
		return errors.New("required 是否是标准品类的必选属性。1：是, 0：否 is null")
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
	if iotutil.IsEmpty(request.Deleted) {
		return errors.New("deleted 删除的标识 0-正常 1-删除 is null")
	}

	// 赋值参数赋值
	var saveObj = model.TPmThingModelProperties{
		ModelId:       request.ModelId,
		ProductKey:    request.ProductKey,
		Identifier:    request.Identifier,
		DataType:      request.DataType,
		Name:          request.Name,
		RwFlag:        request.RwFlag,
		DataSpecs:     request.DataSpecs,
		DataSpecsList: request.DataSpecsList,
		Required:      request.Required,
		Custom:        request.Custom,
		Extension:     request.Extension,
		CreatedBy:     request.CreatedBy,
		//CreatedTime:   request.CreatedTime,
		//UpdatedTime:   request.UpdatedTime,
		//Deleted:       request.Deleted,
		Desc:        request.Desc,
		TriggerCond: request.TriggerCond,
		ExecCond:    request.ExecCond,
	}
	saveObj.Id = iotutil.GetNextSeqInt64()
	err := do.Create(&saveObj)
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "CreateTPmThingModelProperties").Error(err)
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmThingModelPropertiesRequest{Id: saveObj.Id}
	return nil
}

// CreateTPmThingModelPropertiesBatch batch create
func (TPmThingModelPropertiesHandler) CreateTPmThingModelPropertiesBatch(ctx context.Context, requests []*protosService.TPmThingModelPropertiesRequest, response *protosService.TPmThingModelPropertiesResponse) error {
	iotlogger.LogHelper.Info("Received Handler.CreateTPmThingModelPropertiesBatch request")
	t := orm.Use(iotmodel.GetDB()).TPmThingModelProperties
	do := t.WithContext(context.Background())

	//参数判断 requests
	var (
		saveObjs = make([]*model.TPmThingModelProperties, len(requests))
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
		if iotutil.IsEmpty(request.DataType) {
			return errors.New("data_type 属性值的数据类型。可选值：ARRAY、STRUCT、INT、FLOAT、DOUBLE、TEXT、DATE、JSON、ENUM、BOOL。不同数据类型，可传入的参数不同。 is null")
		}
		if iotutil.IsEmpty(request.Name) {
			return errors.New("name 属性名称。可包含中文、大小写英文字母、数字、短划线（-）、下划线（_）和半角句号（.），且必须以中文、英文字母或数字开头，长度不超过30个字符，一个中文计为一个字符。 is null")
		}
		if iotutil.IsEmpty(request.RwFlag) {
			return errors.New("rw_flag 在云端可以对该属性进行的操作类型。READ_WRITE：读写。READ_ONLY：只读。WRITE_ONLY：只写。 is null")
		}
		if iotutil.IsEmpty(request.Required) {
			return errors.New("required 是否是标准品类的必选属性。1：是, 0：否 is null")
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
		if iotutil.IsEmpty(request.Deleted) {
			return errors.New("deleted 删除的标识 0-正常 1-删除 is null")
		}
		// 赋值参数赋值
		var saveObj = &model.TPmThingModelProperties{
			ModelId:       request.ModelId,
			ProductKey:    request.ProductKey,
			Identifier:    request.Identifier,
			DataType:      request.DataType,
			Name:          request.Name,
			RwFlag:        request.RwFlag,
			DataSpecs:     request.DataSpecs,
			DataSpecsList: request.DataSpecsList,
			Required:      request.Required,
			Custom:        request.Custom,
			Extension:     request.Extension,
			CreatedBy:     request.CreatedBy,
			//CreatedTime:   request.CreatedTime,
			//UpdatedTime:   request.UpdatedTime,
			//Deleted:       request.Deleted,
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
		iotlogger.LogHelper.WithTag("func", "CreateTPmThingModelPropertiesBatch").Error(err)
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmThingModelPropertiesRequest{Ids: ids}
	return nil
}

func UpdateInvalidForModelProperties(tx *orm.Query, productModelId int64, productKey string) error {
	if productModelId == 0 || productKey == "" {
		return errors.New("缺条件")
	}
	t := tx.TPmThingModelProperties
	_, err := t.WithContext(context.Background()).Where(t.ProductKey.Eq(productKey), t.ModelId.Eq(productModelId)).Update(t.Valid, 0)
	return err
}

func CopyTypeToProductForModelProperties(tx *orm.Query, typeid int64, productModelId int64, productKey string, mapInfos map[int64]Cond) error {
	ctx := context.Background()
	tm := tx.TPmThingModel
	t := tx.TPmThingModelProperties
	//list, err := t.WithContext(ctx).Where(t.WithContext(ctx).Columns(t.ModelId).In(tm.WithContext(ctx).Select(tm.Id).
	//	Where(tm.ProductTypeId.Eq(typeid), tm.Standard.Eq(1)))).Find()
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
