// Code generated by sgen.exe,2022-04-27 10:55:25. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_open_system/model"
	"cloud_platform/iot_model/db_open_system/orm"
	"cloud_platform/iot_open_system_service/convert"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OpenConfigSvc struct {
	Ctx context.Context
}

// 创建OpenConfig
func (s *OpenConfigSvc) CreateOpenConfig(req *proto.OpenConfig) (*proto.OpenConfig, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TOpenConfig
	do := t.WithContext(context.Background())
	dbObj := convert.OpenConfig_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateOpenConfig error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除OpenConfig
func (s *OpenConfigSvc) DeleteOpenConfig(req *proto.OpenConfig) (*proto.OpenConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenConfig
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.ConfigId != 0 { //整数
		do = do.Where(t.ConfigId.Eq(req.ConfigId))
	}
	if req.ConfigName != "" { //字符串
		do = do.Where(t.ConfigName.Eq(req.ConfigName))
	}
	if req.ConfigKey != "" { //字符串
		do = do.Where(t.ConfigKey.Eq(req.ConfigKey))
	}
	if req.ConfigValue != "" { //字符串
		do = do.Where(t.ConfigValue.Eq(req.ConfigValue))
	}
	if req.ConfigType != 0 { //整数
		do = do.Where(t.ConfigType.Eq(req.ConfigType))
	}
	if req.CreateBy != 0 { //整数
		do = do.Where(t.CreateBy.Eq(req.CreateBy))
	}
	if req.UpdateBy != 0 { //整数
		do = do.Where(t.UpdateBy.Eq(req.UpdateBy))
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
		logger.Errorf("DeleteOpenConfig error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除OpenConfig
func (s *OpenConfigSvc) DeleteByIdOpenConfig(req *proto.OpenConfig) (*proto.OpenConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenConfig
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.ConfigId != 0 { //整数
		do = do.Where(t.ConfigId.Eq(req.ConfigId))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdOpenConfig error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除OpenConfig
func (s *OpenConfigSvc) DeleteByIdsOpenConfig(req *proto.OpenConfigBatchDeleteRequest) (*proto.OpenConfigBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TOpenConfig
		do := t.WithContext(context.Background())

		do = do.Where(t.ConfigId.Eq(k.ConfigId))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsOpenConfig error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新OpenConfig
func (s *OpenConfigSvc) UpdateOpenConfig(req *proto.OpenConfig) (*proto.OpenConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenConfig
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.ConfigName != "" { //字符串
		updateField = append(updateField, t.ConfigName)
	}
	if req.ConfigKey != "" { //字符串
		updateField = append(updateField, t.ConfigKey)
	}
	if req.ConfigValue != "" { //字符串
		updateField = append(updateField, t.ConfigValue)
	}
	if req.ConfigType != 0 { //整数
		updateField = append(updateField, t.ConfigType)
	}
	if req.CreateBy != 0 { //整数
		updateField = append(updateField, t.CreateBy)
	}
	if req.UpdateBy != 0 { //整数
		updateField = append(updateField, t.UpdateBy)
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

	if req.ConfigId != 0 { //整数
		do = do.Where(t.ConfigId.Eq(req.ConfigId))
		HasPrimaryKey = true
	}

	if !HasPrimaryKey {
		logger.Error("UpdateOpenConfig error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.OpenConfig_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateOpenConfig error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段OpenConfig
func (s *OpenConfigSvc) UpdateAllOpenConfig(req *proto.OpenConfig) (*proto.OpenConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenConfig
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.ConfigName)
	updateField = append(updateField, t.ConfigKey)
	updateField = append(updateField, t.ConfigValue)
	updateField = append(updateField, t.ConfigType)
	updateField = append(updateField, t.CreateBy)
	updateField = append(updateField, t.UpdateBy)
	updateField = append(updateField, t.Remark)
	updateField = append(updateField, t.CreatedBy)
	updateField = append(updateField, t.UpdatedBy)
	if len(updateField) > 0 {
		do = do.Select(updateField...)
	}
	//主键条件
	HasPrimaryKey := false
	if req.ConfigId != 0 { //整数
		do = do.Where(t.ConfigId.Eq(req.ConfigId))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateAllOpenConfig error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OpenConfig_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllOpenConfig error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *OpenConfigSvc) UpdateFieldsOpenConfig(req *proto.OpenConfigUpdateFieldsRequest) (*proto.OpenConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenConfig
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsOpenConfig error : missing updateField")
		logger.Error(err)
		return nil, err
	}
	do = do.Select(updateField...)

	//主键条件
	HasPrimaryKey := false
	if req.Data.ConfigId != 0 { //整数
		do = do.Where(t.ConfigId.Eq(req.Data.ConfigId))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateFieldsOpenConfig error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OpenConfig_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsOpenConfig error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找OpenConfig
func (s *OpenConfigSvc) FindOpenConfig(req *proto.OpenConfigFilter) (*proto.OpenConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenConfig
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.ConfigId != 0 { //整数
		do = do.Where(t.ConfigId.Eq(req.ConfigId))
	}
	if req.ConfigName != "" { //字符串
		do = do.Where(t.ConfigName.Eq(req.ConfigName))
	}
	if req.ConfigKey != "" { //字符串
		do = do.Where(t.ConfigKey.Eq(req.ConfigKey))
	}
	if req.ConfigValue != "" { //字符串
		do = do.Where(t.ConfigValue.Eq(req.ConfigValue))
	}
	if req.ConfigType != 0 { //整数
		do = do.Where(t.ConfigType.Eq(req.ConfigType))
	}
	if req.CreateBy != 0 { //整数
		do = do.Where(t.CreateBy.Eq(req.CreateBy))
	}
	if req.UpdateBy != 0 { //整数
		do = do.Where(t.UpdateBy.Eq(req.UpdateBy))
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
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindOpenConfig error : %s", err.Error())
		return nil, err
	}
	res := convert.OpenConfig_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找OpenConfig
func (s *OpenConfigSvc) FindByIdOpenConfig(req *proto.OpenConfigFilter) (*proto.OpenConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenConfig
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.ConfigId != 0 { //整数
		do = do.Where(t.ConfigId.Eq(req.ConfigId))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdOpenConfig error : %s", err.Error())
		return nil, err
	}
	res := convert.OpenConfig_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找OpenConfig,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *OpenConfigSvc) GetListOpenConfig(req *proto.OpenConfigListRequest) ([]*proto.OpenConfig, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TOpenConfig
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.ConfigId != 0 { //整数
			do = do.Where(t.ConfigId.Eq(query.ConfigId))
		}
		if query.ConfigName != "" { //字符串
			do = do.Where(t.ConfigName.Like("%" + query.ConfigName + "%"))
		}
		if query.ConfigKey != "" { //字符串
			do = do.Where(t.ConfigKey.Like("%" + query.ConfigKey + "%"))
		}
		if query.ConfigValue != "" { //字符串
			do = do.Where(t.ConfigValue.Like("%" + query.ConfigValue + "%"))
		}
		if query.ConfigType != 0 { //整数
			do = do.Where(t.ConfigType.Eq(query.ConfigType))
		}
		if query.CreateBy != 0 { //整数
			do = do.Where(t.CreateBy.Eq(query.CreateBy))
		}
		if query.UpdateBy != 0 { //整数
			do = do.Where(t.UpdateBy.Eq(query.UpdateBy))
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
	orderCol, ok := t.GetFieldByName(req.OrderKey)
	if !ok {
		orderCol = t.ConfigId
	}
	if req.OrderDesc != "" {
		do = do.Order(orderCol.Desc())
	} else {
		do = do.Order(orderCol)
	}

	var list []*model.TOpenConfig
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
		logger.Errorf("GetListOpenConfig error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.OpenConfig, len(list))
	for i, v := range list {
		result[i] = convert.OpenConfig_db2pb(v)
	}
	return result, total, nil
}
