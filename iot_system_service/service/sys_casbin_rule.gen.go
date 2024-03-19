// Code generated by sgen.exe,2022-04-18 19:12:08. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_system/model"
	"cloud_platform/iot_model/db_system/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_system_service/convert"
)

type SysCasbinRuleSvc struct {
	Ctx context.Context
}

// 创建SysCasbinRule
func (s *SysCasbinRuleSvc) CreateSysCasbinRule(req *proto.SysCasbinRule) (*proto.SysCasbinRule, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TSysCasbinRule
	do := t.WithContext(context.Background())
	dbObj := convert.SysCasbinRule_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateSysCasbinRule error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除SysCasbinRule
func (s *SysCasbinRuleSvc) DeleteSysCasbinRule(req *proto.SysCasbinRule) (*proto.SysCasbinRule, error) {
	t := orm.Use(iotmodel.GetDB()).TSysCasbinRule
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.Ptype != "" { //字符串
		do = do.Where(t.Ptype.Eq(req.Ptype))
	}
	if req.V0 != "" { //字符串
		do = do.Where(t.V0.Eq(req.V0))
	}
	if req.V1 != "" { //字符串
		do = do.Where(t.V1.Eq(req.V1))
	}
	if req.V2 != "" { //字符串
		do = do.Where(t.V2.Eq(req.V2))
	}
	if req.V3 != "" { //字符串
		do = do.Where(t.V3.Eq(req.V3))
	}
	if req.V4 != "" { //字符串
		do = do.Where(t.V4.Eq(req.V4))
	}
	if req.V5 != "" { //字符串
		do = do.Where(t.V5.Eq(req.V5))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteSysCasbinRule error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除SysCasbinRule
func (s *SysCasbinRuleSvc) DeleteByIdSysCasbinRule(req *proto.SysCasbinRule) (*proto.SysCasbinRule, error) {
	t := orm.Use(iotmodel.GetDB()).TSysCasbinRule
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdSysCasbinRule error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除SysCasbinRule
func (s *SysCasbinRuleSvc) DeleteByIdsSysCasbinRule(req *proto.SysCasbinRuleBatchDeleteRequest) (*proto.SysCasbinRuleBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TSysCasbinRule
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsSysCasbinRule error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新SysCasbinRule
func (s *SysCasbinRuleSvc) UpdateSysCasbinRule(req *proto.SysCasbinRule) (*proto.SysCasbinRule, error) {
	t := orm.Use(iotmodel.GetDB()).TSysCasbinRule
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.Ptype != "" { //字符串
		updateField = append(updateField, t.Ptype)
	}
	if req.V0 != "" { //字符串
		updateField = append(updateField, t.V0)
	}
	if req.V1 != "" { //字符串
		updateField = append(updateField, t.V1)
	}
	if req.V2 != "" { //字符串
		updateField = append(updateField, t.V2)
	}
	if req.V3 != "" { //字符串
		updateField = append(updateField, t.V3)
	}
	if req.V4 != "" { //字符串
		updateField = append(updateField, t.V4)
	}
	if req.V5 != "" { //字符串
		updateField = append(updateField, t.V5)
	}
	if len(updateField) > 0 {
		do = do.Select(updateField...)
	}
	//主键条件
	HasPrimaryKey := false

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
		HasPrimaryKey = true
	}

	if !HasPrimaryKey {
		logger.Error("UpdateSysCasbinRule error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.SysCasbinRule_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateSysCasbinRule error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段SysCasbinRule
func (s *SysCasbinRuleSvc) UpdateAllSysCasbinRule(req *proto.SysCasbinRule) (*proto.SysCasbinRule, error) {
	t := orm.Use(iotmodel.GetDB()).TSysCasbinRule
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.Ptype)
	updateField = append(updateField, t.V0)
	updateField = append(updateField, t.V1)
	updateField = append(updateField, t.V2)
	updateField = append(updateField, t.V3)
	updateField = append(updateField, t.V4)
	updateField = append(updateField, t.V5)
	if len(updateField) > 0 {
		do = do.Select(updateField...)
	}
	//主键条件
	HasPrimaryKey := false
	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateAllSysCasbinRule error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SysCasbinRule_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllSysCasbinRule error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *SysCasbinRuleSvc) UpdateFieldsSysCasbinRule(req *proto.SysCasbinRuleUpdateFieldsRequest) (*proto.SysCasbinRule, error) {
	t := orm.Use(iotmodel.GetDB()).TSysCasbinRule
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFields error : missing updateField")
		logger.Error(err)
		return nil, err
	}
	do = do.Select(updateField...)
	//主键条件
	HasPrimaryKey := false
	if req.Data.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Data.Id))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateFieldsSysCasbinRule error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SysCasbinRule_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsSysCasbinRule error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找SysCasbinRule
func (s *SysCasbinRuleSvc) FindSysCasbinRule(req *proto.SysCasbinRuleFilter) (*proto.SysCasbinRule, error) {
	t := orm.Use(iotmodel.GetDB()).TSysCasbinRule
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.Ptype != "" { //字符串
		do = do.Where(t.Ptype.Like("%" + req.Ptype + "%"))
	}
	if req.V0 != "" { //字符串
		do = do.Where(t.V0.Like("%" + req.V0 + "%"))
	}
	if req.V1 != "" { //字符串
		do = do.Where(t.V1.Like("%" + req.V1 + "%"))
	}
	if req.V2 != "" { //字符串
		do = do.Where(t.V2.Like("%" + req.V2 + "%"))
	}
	if req.V3 != "" { //字符串
		do = do.Where(t.V3.Like("%" + req.V3 + "%"))
	}
	if req.V4 != "" { //字符串
		do = do.Where(t.V4.Like("%" + req.V4 + "%"))
	}
	if req.V5 != "" { //字符串
		do = do.Where(t.V5.Like("%" + req.V5 + "%"))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindSysCasbinRule error : %s", err.Error())
		return nil, err
	}
	res := convert.SysCasbinRule_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找SysCasbinRule
func (s *SysCasbinRuleSvc) FindByIdSysCasbinRule(req *proto.SysCasbinRuleFilter) (*proto.SysCasbinRule, error) {
	t := orm.Use(iotmodel.GetDB()).TSysCasbinRule
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdSysCasbinRule error : %s", err.Error())
		return nil, err
	}
	res := convert.SysCasbinRule_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找SysCasbinRule,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *SysCasbinRuleSvc) GetListSysCasbinRule(req *proto.SysCasbinRuleListRequest) ([]*proto.SysCasbinRule, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TSysCasbinRule
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.Ptype != "" { //字符串
			do = do.Where(t.Ptype.Like("%" + query.Ptype + "%"))
		}
		if query.V0 != "" { //字符串
			do = do.Where(t.V0.Like("%" + query.V0 + "%"))
		}
		if query.V1 != "" { //字符串
			do = do.Where(t.V1.Like("%" + query.V1 + "%"))
		}
		if query.V2 != "" { //字符串
			do = do.Where(t.V2.Like("%" + query.V2 + "%"))
		}
		if query.V3 != "" { //字符串
			do = do.Where(t.V3.Like("%" + query.V3 + "%"))
		}
		if query.V4 != "" { //字符串
			do = do.Where(t.V4.Like("%" + query.V4 + "%"))
		}
		if query.V5 != "" { //字符串
			do = do.Where(t.V5.Like("%" + query.V5 + "%"))
		}
	}
	orderCol, ok := t.GetFieldByName(req.OrderKey)
	if !ok {
		orderCol = t.Id
	}
	if req.OrderDesc != "" {
		do = do.Order(orderCol.Desc())
	} else {
		do = do.Order(orderCol)
	}

	var list []*model.TSysCasbinRule
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
		logger.Errorf("GetListSysCasbinRule error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.SysCasbinRule, len(list))
	for i, v := range list {
		result[i] = convert.SysCasbinRule_db2pb(v)
	}
	return result, total, nil
}
