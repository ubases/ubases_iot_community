package handler

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_model/db_product/orm"
	"cloud_platform/iot_product_service/service"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/mitchellh/mapstructure"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"

	"go-micro.dev/v4"
)

// 以字母开头,允许3~32长度
var regVariable = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{1,31}$`)

// The Register tPmProductType handler.
func RegisterTPmProductTypeHandler(service micro.Service) error {
	err := protosService.RegisterTPmProductTypeHandler(service.Server(), new(TPmProductTypeHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterTPmProductTypeHandler发生错误:%s", err.Error())
	}
	return err
}

type TPmProductTypeUpdate struct {
	Id         int64     `json:"id"`         //主键（雪花算法19位）
	Name       string    `json:"name"`       //分类名称
	NameEn     string    `json:"nameEn"`     //分类名称（英文）
	Identifier string    `json:"identifier"` //属性的标识符。可包含大小写英文字母、数字、下划线（_），长度不超过50个字符。
	Sort       int32     `json:"sort"`       //排序
	ParentId   int64     `json:"parentId"`   //父ID
	Desc       string    `json:"desc"`       //描述
	CreatedBy  int64     `json:"createdBy"`  //创建人
	CreatedAt  time.Time `json:"createdAt"`  //创建时间
	UpdatedBy  int64     `json:"updatedBy"`  //修改人
	UpdatedAt  time.Time `json:"updatedAt"`  //修改时间
	DeletedAt  time.Time `json:"deletedAt"`  //删除的标识 0-正常 1-删除
}

type TPmProductTypeHandler struct{}

// ListTPmProductType query list by paging
func (TPmProductTypeHandler) ListTPmProductType(ctx context.Context, request *protosService.TPmProductTypeFilterPage, response *protosService.TPmProductTypeResponseList) error {
	iotlogger.LogHelper.Info("Received Handler.ListTPmProductType request")
	t := orm.Use(iotmodel.GetDB()).TPmProductType
	tp := orm.Use(iotmodel.GetDB()).TPmProduct
	do := t.WithContext(ctx).Select(t.ALL, tp.ProductTypeId).
		LeftJoin(tp, tp.ProductTypeIdPath.FindInSetWithInt64Col(t.Id), tp.DeletedAt.IsNull())

	// 判断参数进行查询
	if request.QueryObj != nil {
		if request.QueryObj.Id != 0 {
			do = do.Where(t.Id.Eq(request.QueryObj.Id))
		}
	}
	if !iotutil.IsEmpty(request.SearchKey) {
		if !isDigit(request.SearchKey) {
			do = do.Where(t.WithContext(context.Background()).
				Where(t.Name.Like("%" + request.SearchKey + "%")).
				Or(t.NameEn.Like("%" + request.SearchKey + "%")))
		} else {
			do = do.Where(t.WithContext(context.Background()).Where(t.Id.Eq(iotutil.ToInt64(request.SearchKey))).
				Or(t.Name.Like("%" + request.SearchKey + "%")).
				Or(t.NameEn.Like("%" + request.SearchKey + "%")))
		}
	}
	do1 := gen.Table(do.As("p")).Select(
		field.NewString("p", "*"),
		field.NewInt64("p", "product_type_id").Count().As("count")).
		Group(field.NewInt64("p", "id")).Order(field.NewTime("p", "updated_at").Desc())
	var (
		list  []*TPmProductType
		count int64
		err   error
	)
	err = do1.Scan(&list)
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "ListTPmProductType").Error(err)
		return err
	}
	response.Code = 200
	response.Total = count
	response.List = make([]*protosService.TPmProductTypeRequest, len(list))
	for i, _ := range response.List {
		mapstructure.WeakDecode(list[i], &response.List[i])
		response.List[i].CreatedTime = list[i].CreatedAt.Format("2006-01-02 15:04:05")
		response.List[i].UpdatedTime = list[i].UpdatedAt.Format("2006-01-02 15:04:05")
	}
	return nil
}

// CreateTPmProductType create
func (s *TPmProductTypeHandler) CreateTPmProductType(ctx context.Context, request *protosService.TPmProductTypeRequest, response *protosService.TPmProductTypeResponse) error {
	iotlogger.LogHelper.Info("Received Handler.CreateTPmProductType request")
	var setError = func(msg string) error {
		response.Msg = msg
		return nil
	}

	//参数判断 request
	if iotutil.IsEmpty(request.Id) {
		return setError("id 主键（雪花算法19位） is null")
	}
	if iotutil.IsEmpty(request.Name) {
		return setError("name 分类名称 is null")
	}
	if iotutil.IsEmpty(request.NameEn) {
		return setError("name_en 分类名称（英文） is null")
	}
	if iotutil.IsEmpty(request.Identifier) {
		return setError("identifier 属性的标识符。可包含大小写英文字母、数字、下划线（_），长度不超过50个字符。 is null")
	}
	if iotutil.IsEmpty(request.Sort) {
		return setError("sort 排序 is null")
	}
	if iotutil.IsEmpty(request.ParentId) {
		return setError("parent_id 父ID is null")
	}
	if iotutil.IsEmpty(request.CreatedBy) {
		return setError("created_by 创建人 is null")
	}
	if iotutil.IsEmpty(request.CreatedTime) {
		return setError("created_time 创建时间 is null")
	}
	if iotutil.IsEmpty(request.UpdatedBy) {
		return setError("updated_by 修改人 is null")
	}
	if iotutil.IsEmpty(request.UpdatedTime) {
		return setError("updated_at 修改时间 is null")
	}
	if iotutil.IsEmpty(request.Deleted) {
		return setError("deleted 删除的标识 0-正常 1-删除 is null")
	}

	//一级分类不允许导入物模型
	if request.ParentId == 0 && len(request.ModelsItems) > 0 {
		response.Msg = "一级分类不允许导入物模型"
		response.Code = 400
		return nil
	}

	//分类名称去重
	isExists, err := s.existsTypeByName(request.Name, request.NameEn, 0)
	if err != nil {
		response.Msg = err.Error()
		return nil
	}
	if isExists {
		response.Msg = "产品分类名称已存在"
		return nil
	}
	var saveObj = model.TPmProductType{}
	mapstructure.WeakDecode(request, &saveObj)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.CreatedAt = time.Now()
	saveObj.UpdatedAt = time.Now()

	var modelItems []*model.TPmThingModelItem
	for _, v := range request.ModelsItems {
		modelItemsObj := model.TPmThingModelItem{
			Id:            v.Id,
			ProductTypeId: saveObj.Id,
			Dpid:          v.Dpid,
			Identifier:    strings.TrimSpace(v.Identifier),
			Name:          v.Name,
			RwFlag:        v.RwFlag,
			DataType:      v.DataType,
			Properties:    v.Properties,
			Mark:          v.Mark,
			Required:      v.Required,
		}
		modelItems = append(modelItems, &modelItemsObj)
	}

	var thingModelObj = model.TPmThingModel{
		Id:            iotutil.GetNextSeqInt64(),
		ProductTypeId: saveObj.Id,
		Standard:      1,
		Version:       "V1.0.0",
		Description:   "",
		CreatedBy:     request.CreatedBy,
	}

	modelpropertieslist, err := toTPmThingModelProperties(thingModelObj.Id, modelItems)
	if err != nil {
		response.Msg = err.Error()
		return err
	}

	q := orm.Use(iotmodel.GetDB())
	err = q.Transaction(func(tx *orm.Query) error {
		//保存分类
		if errx := tx.TPmProductType.WithContext(ctx).Create(&saveObj); errx != nil {
			return errx
		}
		//保存分类物模型原始表
		if errx := tx.TPmThingModelItem.WithContext(ctx).CreateInBatches(modelItems, len(modelItems)); errx != nil {
			return errx
		}
		//保存分类物模型
		if errx := tx.TPmThingModel.WithContext(ctx).Create(&thingModelObj); errx != nil {
			return errx
		}
		//保存分类物模型数据属性
		if errx := tx.TPmThingModelProperties.WithContext(ctx).CreateInBatches(modelpropertieslist, len(modelpropertieslist)); errx != nil {
			return errx
		}
		return nil
	})
	//新增入库
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "CreateTPmProductType").Error(err)
		return nil
	}

	response.Code = 200
	response.Data = &protosService.TPmProductTypeRequest{Id: saveObj.Id}

	service.GetJsPublisherMgr().PushData(&service.NatsPubData{
		Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
		Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_T_PM_PRODUCT_TYPE, saveObj.Id, "name", saveObj.Name, saveObj.NameEn),
	})
	return nil
}

// UpdateTPmProductType update
func (s *TPmProductTypeHandler) UpdateTPmProductType(ctx context.Context, filter *protosService.TPmProductTypeRequest, response *protosService.TPmProductTypeResponse) error {
	iotlogger.LogHelper.Info("Received Handler.UpdateTPmProductType request")

	//一级分类不允许导入物模型
	if filter.ParentId == 0 && len(filter.ModelsItems) > 0 {
		response.Msg = "一级分类不允许导入物模型"
		response.Code = 400
		return nil
	}

	//分类名称去重
	isExists, err := s.existsTypeByName(filter.Name, filter.NameEn, filter.Id)
	if err != nil {
		return err
	}
	if isExists {
		return errors.New("产品分类名称已存在")
	}
	// 赋值参数赋值
	var updateObj = model.TPmProductType{}
	updateObj.UpdatedAt = time.Now()
	mapstructure.WeakDecode(filter, &updateObj)

	var modelItems []*model.TPmThingModelItem
	for _, v := range filter.ModelsItems {
		modelItemsObj := model.TPmThingModelItem{
			Id:            iotutil.GetNextSeqInt64(), // v.Id,
			ProductTypeId: filter.Id,
			Dpid:          v.Dpid,
			Identifier:    strings.TrimSpace(v.Identifier),
			Name:          v.Name,
			RwFlag:        v.RwFlag,
			DataType:      v.DataType,
			Properties:    v.Properties,
			Mark:          v.Mark,
			Required:      v.Required,
		}
		modelItems = append(modelItems, &modelItemsObj)
	}

	var modelid int64
	t := orm.Use(iotmodel.GetDB()).TPmThingModel
	do := t.WithContext(context.Background())
	err = do.Select(t.Id).Where(t.ProductTypeId.Eq(filter.Id), t.ProductKey.Eq("")).Scan(&modelid)
	if err != nil {
		return err
	}
	var thingModelObj *model.TPmThingModel = nil
	if modelid == 0 {
		thingModelObj = &model.TPmThingModel{
			Id:            iotutil.GetNextSeqInt64(),
			ProductTypeId: filter.Id,
			Standard:      1,
			Version:       "V1.0.0",
			Description:   "",
			CreatedBy:     filter.CreatedBy,
		}
		modelid = thingModelObj.Id
	}

	modelpropertieslist, err := toTPmThingModelProperties(modelid, modelItems)
	if err != nil {
		response.Msg = err.Error()
		return err
	}

	q := orm.Use(iotmodel.GetDB())
	err = q.Transaction(func(tx *orm.Query) error {
		//更新分类
		tptype := tx.TPmProductType
		res, errx := tptype.WithContext(ctx).Select(tptype.Name, tptype.NameEn, tptype.Sort, tptype.Identifier, tptype.Desc, tptype.ParentId,
			tptype.Standard, tptype.ImgSize, tptype.ImgPath, tptype.ImgFullPath, tptype.ImgName, tptype.ImgKey).Where(tptype.Id.Eq(filter.Id)).Updates(updateObj)
		if errx != nil {
			return errx
		}
		if res.RowsAffected == 0 {
			return errors.New("记录不存在")
		}

		//删除分类物模型原始表记录
		ttmi := tx.TPmThingModelItem
		_, errx = ttmi.WithContext(ctx).Where(ttmi.ProductTypeId.Eq(filter.Id)).Delete()
		if errx != nil {
			return errx
		}
		//保存分类物模型原始表
		errx = tx.TPmThingModelItem.WithContext(ctx).CreateInBatches(modelItems, len(modelItems))
		if errx != nil {
			return errx
		}

		//若需要保存分类物模型，则保存
		if thingModelObj != nil {
			errx = tx.TPmThingModel.WithContext(ctx).Create(thingModelObj)
			if errx != nil {
				return errx
			}
		}

		ttmp := tx.TPmThingModelProperties
		_, errx = ttmp.WithContext(ctx).Unscoped().Where(ttmp.ModelId.Eq(modelid)).Delete()
		if errx != nil {
			return errx
		}

		//保存分类物模型数据属性
		errx = tx.TPmThingModelProperties.WithContext(ctx).CreateInBatches(modelpropertieslist, len(modelpropertieslist))
		if errx != nil {
			return errx
		}

		//更新产品数据中的分类层级数据
		tP := tx.TPmProduct
		tx.TPmProduct.WithContext(ctx).Where(tP.ProductTypeIdPath.Like("%"+iotutil.ToString(filter.Id))).
			Update(tP.ProductTypeIdPath, fmt.Sprintf("%v,%v", filter.ParentId, filter.Id))
		return nil
	})

	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "UpdateTIotDeviceTriad").Error(err)
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmProductTypeRequest{Id: updateObj.Id}

	service.GetJsPublisherMgr().PushData(&service.NatsPubData{
		Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
		Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_T_PM_PRODUCT_TYPE, updateObj.Id, "name", updateObj.Name, updateObj.NameEn),
	})
	return nil
}

// DeleteTPmProductType delete
func (s TPmProductTypeHandler) DeleteTPmProductType(ctx context.Context, request *protosService.TPmProductTypeRequest, response *protosService.TPmProductTypeResponse) error {
	iotlogger.LogHelper.Info("Received Handler.DeleteTPmProductType request")
	t := orm.Use(iotmodel.GetDB()).TPmProductType
	do := t.WithContext(context.Background())

	//判断分类是否已经被使用
	exists, err := s.existsProducts(request.Id)
	if err != nil {
		response.Msg = err.Error()
		return nil
	}
	if exists {
		response.Msg = "产品分类已使用，无法删除！"
		return nil
	}

	_, err = do.Where(t.Id.Eq(int64(request.Id))).Delete()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "DeleteTPmProductType").Error(err)
		return err
	}
	response.Code = 200
	return nil
}

// GetByIdTPmProductType query struct by id
func (TPmProductTypeHandler) GetByIdTPmProductType(ctx context.Context, request *protosService.TPmProductTypeFilterById, response *protosService.TPmProductTypeResponseObject) error {
	iotlogger.LogHelper.Info("Received Handler.GetByIdTPmProductType request")
	t := orm.Use(iotmodel.GetDB()).TPmProductType
	do := t.WithContext(context.Background())
	//创建查询条件
	info, err := do.Where(t.Id.Eq(int64(request.Id))).First()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetByIdTPmProductType").Error(err)
		return err
	}

	var list []*model.TPmThingModelItem
	ttmi := orm.Use(iotmodel.GetDB()).TPmThingModelItem
	do1 := ttmi.WithContext(context.Background())
	do1 = do1.Where(ttmi.ProductTypeId.Eq(request.Id))
	list, err = do1.Find()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetByIdTPmProductType").Error(err)
		return err
	}
	response.Code = 200
	mapstructure.WeakDecode(info, &response.Data)
	for _, v := range list {
		modelitem := protosService.ModelsItem{
			Id:            v.Id,
			ProductTypeId: v.ProductTypeId,
			Dpid:          v.Dpid,
			Identifier:    v.Identifier,
			Name:          v.Name,
			RwFlag:        v.RwFlag,
			DataType:      v.DataType,
			Properties:    v.Properties,
			Mark:          v.Mark,
			Required:      v.Required,
		}
		response.Data.ModelsItems = append(response.Data.ModelsItems, &modelitem)
	}
	return nil
}

// GetTPmProductType query struct by struct
func (TPmProductTypeHandler) GetTPmProductType(ctx context.Context, request *protosService.TPmProductTypeFilter, response *protosService.TPmProductTypeResponseObject) error {
	iotlogger.LogHelper.Info("Received Handler.GetTPmProductType request")
	t := orm.Use(iotmodel.GetDB()).TPmProductType
	do := t.WithContext(context.Background())
	// 判断参数进行查询
	if request.Id != 0 {
		do = do.Where(t.Id.Eq(int64(request.Id)))
	}

	info, err := do.First()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetTPmProductType").Error(err)
		return err
	}
	response.Code = 200
	mapstructure.WeakDecode(info, &response.Data)
	return nil
}

// private method
func isDigit(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}

// 新增和修改的时候判断分类名称是否重复
func (s *TPmProductTypeHandler) existsTypeByName(name string, nameEn string, id int64) (bool, error) {
	t := orm.Use(iotmodel.GetDB()).TPmProductType
	do := t.WithContext(context.Background())
	do = do.Where(do.Where(t.Name.Eq(name)).Or(do.Where(t.NameEn.Eq(nameEn))))

	//编辑的时候验证名称是否重复
	if id != 0 {
		do = do.Where(t.Id.Neq(id))
	}
	count, err := do.Count()
	if err != nil {
		return true, err
	}

	return count > 0, err
}

// GetAppProductType query list by paging
func (TPmProductTypeHandler) GetAppProductType(ctx context.Context, request *protosService.TPmProductTypeFilterPage, response *protosService.TPmProductTypeResponseList) error {
	iotlogger.LogHelper.Info("Received Handler.GetAppProductType request")
	t := orm.Use(iotmodel.GetDB()).TPmProductType
	do := t.WithContext(context.Background())
	//关联子查询
	sub := t.As("sub")
	do.Select(t.ALL, sub.Name.As("ParentName")).Join(sub, sub.Id.EqCol(sub.ParentId))

	do = do.Where(t.ParentId.Eq(0))
	// 判断参数进行查询
	if request.QueryObj != nil {
		if request.QueryObj.Id != 0 {
			do = do.Where(t.Id.Eq(request.QueryObj.Id))
		}
	}
	if !iotutil.IsEmpty(request.SearchKey) {
		if !isDigit(request.SearchKey) {
			do = do.Where(t.Name.Like("%" + request.SearchKey + "%"))
		} else {
			do = do.Where(t.Id.Eq(iotutil.ToInt64(request.SearchKey))).Or(t.Name.Like("%" + request.SearchKey + "%"))
		}
	}
	//排序（update_at倒序）
	orderCol, ok := t.GetFieldByName("updated_at")
	if ok {
		do.Order(orderCol.Desc())
	}

	var (
		list  []*model.TPmProductType
		count int64
		err   error
	)
	list, err = do.Find()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetAppProductType").Error(err)
		return err
	}
	response.Code = 200
	response.Total = count
	response.List = make([]*protosService.TPmProductTypeRequest, len(list))
	for i, _ := range response.List {
		mapstructure.WeakDecode(list[i], &response.List[i])
		response.List[i].CreatedTime = list[i].CreatedAt.Format("2006-01-02 15:04:05")
		response.List[i].UpdatedTime = list[i].UpdatedAt.Format("2006-01-02 15:04:05")
	}
	return nil
}

// 新增和修改的时候判断分类名称是否重复
func (s *TPmProductTypeHandler) existsProducts(productTypeId int64) (bool, error) {
	t := orm.Use(iotmodel.GetDB()).TPmProduct
	do := t.WithContext(context.Background())
	do = do.Where(do.Where(t.ProductTypeId.Eq(productTypeId)))
	count, err := do.Count()
	if err != nil {
		return true, err
	}

	return count > 0, err
}

// Enum 例如：0-暂停;1-播放
func getEnumDescList(marks string) []string {
	var enumDesclist []string = make([]string, 0)
	if strings.Index(marks, ";") != -1 {
		enumDesclist = strings.Split(strings.Trim(marks, " "), ";") //例如：0-暂停;1-播放
	} else if strings.Index(marks, "；") != -1 {
		enumDesclist = strings.Split(strings.Trim(marks, " "), "；") //例如：0-暂停；1-播放
	} else if strings.Index(marks, "，") != -1 {
		enumDesclist = strings.Split(strings.Trim(marks, " "), "，") //例如：0-呼吸，1-渐变，2-追光
	} else if strings.Index(marks, ",") != -1 {
		enumDesclist = strings.Split(strings.Trim(marks, " "), ",") //例如：0-暂停,1-播放
	}
	return enumDesclist
}

// Bool 例如：false-关;true-开
func getDescMap(marks string) map[string]string {
	enumMap := make(map[string]string)
	var enumDesclist []string = make([]string, 0)
	if strings.Index(marks, ";") != -1 {
		enumDesclist = strings.Split(strings.Trim(marks, " "), ";") //例如：0-暂停;1-播放
	} else if strings.Index(marks, "；") != -1 {
		enumDesclist = strings.Split(strings.Trim(marks, " "), "；") //例如：0-暂停；1-播放
	} else if strings.Index(marks, "，") != -1 {
		enumDesclist = strings.Split(strings.Trim(marks, " "), "，") //例如：0-呼吸，1-渐变，2-追光
	} else if strings.Index(marks, ",") != -1 {
		enumDesclist = strings.Split(strings.Trim(marks, " "), ",") //例如：0-暂停,1-播放
	}
	for i, v := range enumDesclist {
		key := iotutil.ToString(i)
		val := ""
		ldesc := strings.Split(v, "-")
		if len(ldesc) > 1 {
			key = ldesc[0]
			val = ldesc[1]
		} else {
			//兼容
			val = ldesc[0]
		}
		enumMap[key] = val
	}
	return enumMap
}

func toTPmThingModelProperties(modelid int64, modelItems []*model.TPmThingModelItem) ([]*model.TPmThingModelProperties, error) {
	var retList []*model.TPmThingModelProperties
	for _, v := range modelItems {
		p := model.TPmThingModelProperties{
			Id:         iotutil.GetNextSeqInt64(),
			ModelId:    modelid,
			Identifier: v.Identifier,
			DataType:   v.DataType,
			Name:       v.Name,
			RwFlag:     v.RwFlag,
			Custom:     0,
			Desc:       v.Mark,
			Dpid:       v.Dpid,
		}
		//必填，字典配置required_flag
		if v.Required == 1 {
			p.Required = 1
		} else {
			p.Required = 2
		}
		if v.RwFlag != "READ_WRITE" && v.RwFlag != "READ" && v.RwFlag != "WRITE" {
			return nil, fmt.Errorf("rwFlag error.(Identifier=%s,rwFlag=%s)", v.Identifier, v.RwFlag)
		}
		if !regVariable.MatchString(v.Identifier) {
			return nil, fmt.Errorf("identifier错误,要求以字母开头,可包含字母、数字、下划线,长度2~32的字符串.(identifier=%s)", v.Identifier)
		}

		if p.DataType == "ENUM" {
			properties := strings.TrimSpace(v.Properties)
			marks := strings.TrimSpace(v.Mark)
			if strings.Index(marks, ":") != -1 {
				marks = strings.Split(marks, ":")[1]
			}
			var list []EnumDataSpaces
			if properties != "" {
				enumlist := strings.Split(strings.Trim(properties, " "), ";") //例如：PAUSE;PLAY
				enumDescMap := getDescMap(marks)

				for k, enumVal := range enumlist {
					l := strings.Split(enumVal, ":")
					var vName string
					var vVal int
					var vDesc string
					if len(l) > 1 {
						vName = l[0]
						vVal, _ = iotutil.ToIntErr(l[1])
					} else {
						//兼容
						vName = l[0]
						vVal = k
					}

					if v, ok := enumDescMap[iotutil.ToString(vVal)]; ok {
						vDesc = v
					} else {
						vDesc = vName
					}

					if regVariable.MatchString(vName) {
						ds := EnumDataSpaces{Custom: 0, DataType: p.DataType, Name: vName, Value: int64(vVal), Desc: vDesc}
						list = append(list, ds)
					} else {
						return nil, fmt.Errorf("enum properties error.(Identifier=%s,name=%s)", v.Identifier, vName)
					}
				}
			}
			if len(list) > 0 {
				buf, err := json.Marshal(list)
				if err == nil {
					p.DataSpecsList = string(buf)
				}
			} else {
				p.DataSpecsList = "[]"
			}
		} else if p.DataType == "BOOL" {
			var list []EnumDataSpaces
			properties := strings.TrimSpace(v.Properties)
			if properties != "" {
				enumlist := strings.Split(strings.Trim(properties, " "), ";")
				if len(enumlist) != 2 || enumlist[0] != "false" || enumlist[1] != "true" {
					return nil, fmt.Errorf("BOOL properties error.(Identifier=%s,properties=%s)", v.Identifier, v.Properties)
				}
				marks := strings.TrimSpace(v.Mark)
				boolDescMap := getDescMap(marks)
				for k, vv := range enumlist {
					var vDesc string
					if v, ok := boolDescMap[iotutil.ToString(vv)]; ok {
						vDesc = v
					} else {
						vDesc = vv
					}
					ds := EnumDataSpaces{Custom: 0, DataType: p.DataType, Name: vv, Value: int64(k), Desc: vDesc}
					list = append(list, ds)
				}
			} else {
				return nil, fmt.Errorf("BOOL properties error.(Identifier=%s,properties=%s)", v.Identifier, v.Properties)
			}
			buf, _ := json.Marshal(list)
			p.DataSpecsList = string(buf)
		} else if p.DataType == "INT" || p.DataType == "FLOAT" || p.DataType == "DOUBLE" {
			properties := strings.Trim(v.Properties, " ")
			valuedataspecs := ValueDataSpecs{DataType: p.DataType, Custom: 0}
			if properties != "" {
				vallist := strings.Split(properties, ";")
				for _, kv := range vallist {
					vals := strings.Split(kv, ":")
					if len(vals) == 2 {
						vals1 := strings.Trim(vals[1], " ")
						switch strings.Trim(vals[0], " ") { // range:-20-50;step:1;unit:℃
						case "range":
							minmax := strings.Split(vals1, "-")
							if len(minmax) == 2 { //10-15
								valuedataspecs.Min = minmax[0]
								valuedataspecs.Max = minmax[1]
							} else if len(minmax) == 3 { //-10-5
								valuedataspecs.Min = "-" + minmax[1]
								valuedataspecs.Max = minmax[2]
							} else if len(minmax) == 4 { //-10--5
								valuedataspecs.Min = "-" + minmax[1]
								valuedataspecs.Max = "-" + minmax[3]
							}
						case "step":
							valuedataspecs.Step = vals1
						case "unit":
							valuedataspecs.Unit = vals1
						case "multiple":
							valuedataspecs.Multiple = vals1
						}
					}
				}
				//检查数据是否合法
				if err := validate(valuedataspecs, p, v); err != nil {
					return nil, err
				}
			} else {
				return nil, fmt.Errorf("%s properties error.(Identifier=%s,properties=%s)", p.DataType, v.Identifier, v.Properties)
			}
			buf, _ := json.Marshal(valuedataspecs)
			p.DataSpecs = string(buf)
		} else if p.DataType == "TEXT" || p.DataType == "DATE" || p.DataType == "JSON" {
			valuedataspecs := StringDataSpecs{DataType: p.DataType, Custom: 0}
			buf, _ := json.Marshal(valuedataspecs)
			p.DataSpecs = string(buf)
		} else {
			return nil, fmt.Errorf("That dataType %s is not supported", p.DataType)
		}
		retList = append(retList, &p)
	}
	err := ValidateThingModelProperties(retList)
	if err != nil {
		return nil, err
	}
	return retList, nil
}

func validate(valuedataspecs ValueDataSpecs, p model.TPmThingModelProperties, v *model.TPmThingModelItem) error {
	var err error
	var fmin, fmax, fstep float64
	if valuedataspecs.Min != "" {
		fmin, err = strconv.ParseFloat(valuedataspecs.Min, 64)
		if err != nil {
			return fmt.Errorf("%s properties error.(Identifier=%s,properties=%s)", p.DataType, v.Identifier, v.Properties)
		}
	}
	if valuedataspecs.Max != "" {
		fmax, err = strconv.ParseFloat(valuedataspecs.Max, 64)
		if err != nil {
			return fmt.Errorf("%s properties error.(Identifier=%s,properties=%s)", p.DataType, v.Identifier, v.Properties)
		}
	}
	if valuedataspecs.Step != "" {
		fstep, err = strconv.ParseFloat(valuedataspecs.Step, 64)
		if err != nil {
			return fmt.Errorf("%s properties error.(Identifier=%s,properties=%s)", p.DataType, v.Identifier, v.Properties)
		}
	}
	if fmin > fmax {
		return fmt.Errorf("%s properties error.(Identifier=%s,properties=%s)", p.DataType, v.Identifier, v.Properties)
	}
	if p.DataType == "INT" && strings.Contains(valuedataspecs.Step, ".") {
		return fmt.Errorf("%s properties error.(Identifier=%s,properties=%s)", p.DataType, v.Identifier, v.Properties)
	}
	if fmax-fmin < fstep {
		return fmt.Errorf("%s properties error.(Identifier=%s,properties=%s)", p.DataType, v.Identifier, v.Properties)
	}
	return nil
}

func ValidateThingModelProperties(list []*model.TPmThingModelProperties) error {
	mapDpid := make(map[int32]int)
	mapIdentifier := make(map[string]int)
	mapName := make(map[string]int)
	for _, v := range list {
		if _, ok := mapDpid[v.Dpid]; ok {
			mapDpid[v.Dpid] = mapDpid[v.Dpid] + 1
		} else {
			mapDpid[v.Dpid] = 1
		}
		if _, ok := mapIdentifier[v.Identifier]; ok {
			mapIdentifier[v.Identifier] = mapIdentifier[v.Identifier] + 1
		} else {
			mapIdentifier[v.Identifier] = 1
		}
		if _, ok := mapName[v.Name]; ok {
			mapName[v.Name] = mapName[v.Name] + 1
		} else {
			mapName[v.Name] = 1
		}
	}
	if len(list) != len(mapDpid) {
		return fmt.Errorf("dpid dup error.")
	}
	if len(list) != len(mapIdentifier) {
		return fmt.Errorf("identifier dup error.")
	}
	if len(list) != len(mapName) {
		return fmt.Errorf("name dup error.")
	}
	return nil
}

type ValueDataSpecs struct {
	DataType     string `json:"dataType,omitempty"`     // 取值为INT、FLOAT或DOUBLE。
	Max          string `json:"max,omitempty"`          // 最大值。取值为INT、FLOAT或DOUBLE，必须与dataType设置一致
	Min          string `json:"min,omitempty"`          // 最小值。取值为INT、FLOAT或DOUBLE，必须与dataType设置一致
	Step         string `json:"step,omitempty"`         // 步长，数据每次变化的增量。取值为INT、FLOAT或DOUBLE，必须与dataType设置一致
	Precise      string `json:"precise,omitempty"`      // 精度。当dataType取值为FLOAT或DOUBLE时，可传入的参数
	DefaultValue string `json:"defaultValue,omitempty"` // 传入此参数，可存入一个默认值
	Unit         string `json:"unit,omitempty"`         // 单位的符号
	UnitName     string `json:"unitName,omitempty"`     // 单位的名称
	Custom       int32  `json:"custom,omitempty"`       // 是否是自定义功能。1：是 0：否
	Multiple     string `json:"multiple,omitempty"`     //倍数
}

// 字符型DataSpecs（TEXT & DATE & JSON）
type StringDataSpecs struct {
	DataType     string `json:"dataType,omitempty"`     // 取值为DATE或TEXT。
	Length       int32  `json:"length,omitempty"`       // 数据长度，取值不能超过2048，单位：字节。dataType取值为TEXT时，需传入该参数。
	DefaultValue string `json:"defaultValue,omitempty"` // 传入此参数，可存入一个默认值。
	Custom       int32  `json:"custom,omitempty"`       // 是否是自定义功能。1：是 0：否
}

// 枚举型DataSpaceList（ENUM）
type EnumDataSpaces struct {
	DataType string `json:"dataType"` // 数据类型
	Name     string `json:"name"`     // 名称
	Value    int64  `json:"value"`    // 数值
	Custom   int32  `json:"custom"`   // 是否是自定义功能。1：是 0：否
	Desc     string `json:"desc"`     // 名称
}

type TPmProductType struct {
	Id          int64          `gorm:"column:id;primaryKey" json:"id"`                        // 主键（雪花算法19位）
	ParentId    int64          `gorm:"column:parent_id;not null" json:"parentId"`             // 父ID
	Name        string         `gorm:"column:name;not null" json:"name"`                      // 产品分类名称
	NameEn      string         `gorm:"column:name_en;not null" json:"nameEn"`                 // 产品分类名称（英文）
	Identifier  string         `gorm:"column:identifier;not null" json:"identifier"`          // 属性的标识符。可包含大小写英文字母、数字、下划线（_），长度不超过50个字符。
	Sort        int32          `gorm:"column:sort;not null" json:"sort"`                      // 排序
	Standard    int32          `gorm:"column:standard;not null" json:"standard"`              // 是否标准物模型[0:否 1:是]
	Desc        string         `gorm:"column:desc" json:"desc"`                               // 描述
	ImgSize     int64          `gorm:"column:img_size" json:"imgSize"`                        // 图片大小，单位B
	Count       int64          `gorm:"column:count" json:"count"`                             // 关联产品类型数量
	ImgPath     string         `gorm:"column:img_path" json:"imgPath"`                        // 图片路径
	ImgFullPath string         `gorm:"column:img_full_path" json:"imgFullPath"`               // 图片完整路径
	ImgName     string         `gorm:"column:img_name" json:"imgName"`                        // 图片名称
	ImgKey      string         `gorm:"column:img_key" json:"imgKey"`                          // 图片MD5
	CreatedBy   int64          `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
	UpdatedBy   int64          `gorm:"column:updated_by;not null;default:0" json:"updatedBy"` // 修改人
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`                    // 修改时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                    // 删除时间
}
