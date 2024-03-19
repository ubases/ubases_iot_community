// fixme 本文件是demo，展示service实现；正式服务可以删除或者替换成自己的Service
// 代码是自动生成的，演示数据库增删改查
// 文件中的proto包，就是google proto3文件生成的go文件所在的包名，可手动修改为自己生成的go文件包名

package service

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	"cloud_platform/iot_demo_service/convert"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_config/model"
	"cloud_platform/iot_model/db_config/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type ConfigDictTypeSvc struct {
	Ctx context.Context
}

// 创建ConfigDictType
func (s *ConfigDictTypeSvc) CreateConfigDictType(req *proto.ConfigDictType) (*proto.ConfigDictType, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TConfigDictType
	do := t.WithContext(context.Background())

	//判断重复
	isExists, err := s.existsTypeByName(req.DictName, req.DictType, 0)
	if err != nil {
		return nil, err
	}
	if isExists {
		return nil, errors.New("字典名称或者编码已存在")
	}

	dbObj := convert.ConfigDictType_pb2db(req)
	err = do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateConfigDictType error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除ConfigDictType
func (s *ConfigDictTypeSvc) DeleteConfigDictType(req *proto.ConfigDictType) (*proto.ConfigDictType, error) {
	t := orm.Use(iotmodel.GetDB()).TConfigDictType
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.DictId != 0 { //整数
		do = do.Where(t.DictId.Eq(req.DictId))
	}
	if req.DictName != "" { //字符串
		do = do.Where(t.DictName.Eq(req.DictName))
	}
	if req.DictType != "" { //字符串
		do = do.Where(t.DictType.Eq(req.DictType))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.IsSystem != 0 { //整数
		do = do.Where(t.IsSystem.Eq(req.IsSystem))
	}
	if req.Remark != "" { //字符串
		do = do.Where(t.Remark.Eq(req.Remark))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteConfigDictType error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除ConfigDictType
func (s *ConfigDictTypeSvc) DeleteByIdConfigDictType(req *proto.ConfigDictType) (*proto.ConfigDictType, error) {
	t := orm.Use(iotmodel.GetDB()).TConfigDictType
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.DictId != 0 { //整数
		do = do.Where(t.DictId.Eq(req.DictId))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdConfigDictType error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除ConfigDictType
func (s *ConfigDictTypeSvc) DeleteByIdsConfigDictType(req *proto.ConfigDictTypeBatchDeleteRequest) (*proto.ConfigDictTypeBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TConfigDictType
		do := t.WithContext(context.Background())

		do = do.Where(t.DictId.Eq(k.DictId))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsConfigDictType error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新ConfigDictType
func (s *ConfigDictTypeSvc) UpdateConfigDictType(req *proto.ConfigDictType) (*proto.ConfigDictType, error) {
	t := orm.Use(iotmodel.GetDB()).TConfigDictType
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.DictName != "" { //字符串
		updateField = append(updateField, t.DictName)
	}
	if req.DictType != "" { //字符串
		updateField = append(updateField, t.DictType)
	}
	if req.Status != 0 { //整数
		updateField = append(updateField, t.Status)
	}
	if req.ValueType != 0 { //整数
		updateField = append(updateField, t.ValueType)
	}
	if req.IsSystem != 0 { //整数
		updateField = append(updateField, t.IsSystem)
	}
	if req.Remark != "" { //字符串
		updateField = append(updateField, t.Remark)
	}
	if req.CreatedBy != 0 { //整数
		updateField = append(updateField, t.CreatedBy)
	}
	if req.UpdatedBy != 0 { //整数
		updateField = append(updateField, t.UpdatedBy)
	}
	if len(updateField) > 0 {
		do = do.Select(updateField...)
	}
	//主键条件
	HasPrimaryKey := false

	if req.DictId != 0 { //整数
		do = do.Where(t.DictId.Eq(req.DictId))
		HasPrimaryKey = true
	}

	if !HasPrimaryKey {
		logger.Error("UpdateConfigDictType error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.ConfigDictType_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateConfigDictType error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段ConfigDictType
func (s *ConfigDictTypeSvc) UpdateAllConfigDictType(req *proto.ConfigDictType) (*proto.ConfigDictType, error) {
	t := orm.Use(iotmodel.GetDB()).TConfigDictType
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	//判断重复
	isExists, err := s.existsTypeByName(req.DictName, req.DictType, req.DictId)
	if err != nil {
		return nil, err
	}
	if isExists {
		return nil, errors.New("字典名称或者编码已存在")
	}
	updateField = append(updateField, t.DictName)
	updateField = append(updateField, t.DictType)
	updateField = append(updateField, t.Status)
	updateField = append(updateField, t.ValueType)
	updateField = append(updateField, t.IsSystem)
	updateField = append(updateField, t.Remark)
	updateField = append(updateField, t.CreatedBy)
	updateField = append(updateField, t.UpdatedBy)
	if len(updateField) > 0 {
		do = do.Select(updateField...)
	}
	//主键条件
	HasPrimaryKey := false
	if req.DictId != 0 { //整数
		do = do.Where(t.DictId.Eq(req.DictId))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateAllConfigDictType error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.ConfigDictType_pb2db(req)
	_, err = do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllConfigDictType error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *ConfigDictTypeSvc) UpdateFieldsConfigDictType(req *proto.ConfigDictTypeUpdateFieldsRequest) (*proto.ConfigDictType, error) {
	t := orm.Use(iotmodel.GetDB()).TConfigDictType
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	//主键条件
	HasPrimaryKey := false
	if req.Data.DictId != 0 { //整数
		do = do.Where(t.DictId.Eq(req.Data.DictId))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateFieldsConfigDictType error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.ConfigDictType_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsConfigDictType error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找ConfigDictType
func (s *ConfigDictTypeSvc) FindConfigDictType(req *proto.ConfigDictTypeFilter) (*proto.ConfigDictType, error) {
	t := orm.Use(iotmodel.GetDB()).TConfigDictType
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.DictId != 0 { //整数
		do = do.Where(t.DictId.Eq(req.DictId))
	}
	if req.DictName != "" { //字符串
		do = do.Where(t.DictName.Like("%" + req.DictName + "%"))
	}
	if req.DictType != "" { //字符串
		do = do.Where(t.DictType.Like("%" + req.DictType + "%"))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.ValueType != 0 { //整数
		do = do.Where(t.ValueType.Eq(req.ValueType))
	}
	if req.IsSystem != 0 { //整数
		do = do.Where(t.IsSystem.Eq(req.IsSystem))
	}
	if req.Remark != "" { //字符串
		do = do.Where(t.Remark.Like("%" + req.Remark + "%"))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindConfigDictType error : %s", err.Error())
		return nil, err
	}
	res := convert.ConfigDictType_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找ConfigDictType
func (s *ConfigDictTypeSvc) FindByIdConfigDictType(req *proto.ConfigDictTypeFilter) (*proto.ConfigDictType, error) {
	t := orm.Use(iotmodel.GetDB()).TConfigDictType
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.DictId != 0 { //整数
		do = do.Where(t.DictId.Eq(req.DictId))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdConfigDictType error : %s", err.Error())
		return nil, err
	}
	res := convert.ConfigDictType_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找ConfigDictType,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *ConfigDictTypeSvc) GetListConfigDictType(req *proto.ConfigDictTypeListRequest) ([]*proto.ConfigDictType, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TConfigDictType
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.DictId != 0 { //整数
			do = do.Where(t.DictId.Eq(query.DictId))
		}
		if query.DictName != "" { //字符串
			do = do.Where(t.DictName.Like("%" + query.DictName + "%"))
		}
		if query.DictType != "" { //字符串
			do = do.Where(t.DictType.Like("%" + query.DictType + "%"))
		}
		if query.Status != 0 { //整数
			do = do.Where(t.Status.Eq(query.Status))
		}
		if query.IsSystem != 0 { //整数
			do = do.Where(t.IsSystem.Eq(query.IsSystem))
		}
		if query.ValueType != 0 { //整数
			do = do.Where(t.ValueType.Eq(query.ValueType))
		}
		if query.Remark != "" { //字符串
			do = do.Where(t.Remark.Like("%" + query.Remark + "%"))
		}
		if query.CreatedBy != 0 { //整数
			do = do.Where(t.CreatedBy.Eq(query.CreatedBy))
		}
		if query.UpdatedBy != 0 { //整数
			do = do.Where(t.UpdatedBy.Eq(query.UpdatedBy))
		}
	}
	orderCol, ok := t.GetFieldByName("created_at")
	if !ok {
		orderCol = t.DictId
	}
	do = do.Order(orderCol.Desc())

	var list []*model.TConfigDictType
	var total int64
	if req.PageSize > 0 {
		limit := req.PageSize
		if req.Page == 0 {
			req.Page = 1
		}
		offset := req.PageSize * (req.Page - 1)
		list, total, err = do.FindByPage(int(offset), int(limit))
	} else {
		list, err = do.Find()
		total = int64(len(list))
	}
	if err != nil {
		logger.Errorf("GetListConfigDictType error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.ConfigDictType, len(list))
	for i, v := range list {
		result[i] = convert.ConfigDictType_db2pb(v)
	}
	return result, total, nil
}

// 新增和修改的时候判断分类名称是否重复
func (s *ConfigDictTypeSvc) existsTypeByName(dictNamee string, dictType string, id int64) (bool, error) {
	t := orm.Use(iotmodel.GetDB()).TConfigDictType
	do := t.WithContext(context.Background())
	do = do.Where(do.Where(t.DictName.Eq(dictNamee)).Or(do.Where(t.DictType.Eq(dictType))))

	//编辑的时候验证名称是否重复
	if id != 0 {
		do = do.Where(t.DictId.Neq(id))
	}
	count, err := do.Count()
	if err != nil {
		return true, err
	}

	return count > 0, err
}
