// Code generated by sgen.exe,2022-04-21 12:44:21. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_model/db_product/orm"
	"cloud_platform/iot_product_service/convert"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type PmProductPanelRelationSvc struct {
	Ctx context.Context
}

// 创建PmProductPanelRelation
func (s *PmProductPanelRelationSvc) CreatePmProductPanelRelation(tx *orm.Query, req *proto.PmProductPanelRelation) (*proto.PmProductPanelRelation, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TPmProductPanelRelation
	do := t.WithContext(context.Background())
	dbObj := convert.PmProductPanelRelation_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreatePmProductPanelRelation error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 批量创建PmProductPanelRelation
func (s *PmProductPanelRelationSvc) BatchCreatePmProductPanelRelation(tx *orm.Query, req []*proto.PmProductPanelRelation) ([]*proto.PmProductPanelRelation, error) {
	// fixme 请在这里校验参数
	t := tx.TPmProductPanelRelation
	do := t.WithContext(context.Background())
	var dbObjs = make([]*model.TPmProductPanelRelation, len(req))
	for i, relation := range req {
		dbObjs[i] = convert.PmProductPanelRelation_pb2db(relation)
	}
	err := do.CreateInBatches(dbObjs, len(req))
	if err != nil {
		logger.Errorf("BatchCreatePmProductPanelRelation error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除PmProductPanelRelation
func (s *PmProductPanelRelationSvc) DeletePmProductPanelRelation(tx *orm.Query, req *proto.PmProductPanelRelation) (*proto.PmProductPanelRelation, error) {
	t := tx.TPmProductPanelRelation
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.ProductId != 0 { //整数
		do = do.Where(t.ProductId.Eq(req.ProductId))
	}
	if req.ControlPanelId != 0 { //整数
		do = do.Where(t.ControlPanelId.Eq(req.ControlPanelId))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeletePmProductPanelRelation error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除PmProductPanelRelation
func (s *PmProductPanelRelationSvc) DeleteByIdPmProductPanelRelation(req *proto.PmProductPanelRelation) (*proto.PmProductPanelRelation, error) {
	t := orm.Use(iotmodel.GetDB()).TPmProductPanelRelation
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdPmProductPanelRelation error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除PmProductPanelRelation
func (s *PmProductPanelRelationSvc) DeleteByIdsPmProductPanelRelation(req *proto.PmProductPanelRelationBatchDeleteRequest) (*proto.PmProductPanelRelationBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TPmProductPanelRelation
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsPmProductPanelRelation error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新PmProductPanelRelation
func (s *PmProductPanelRelationSvc) UpdatePmProductPanelRelation(req *proto.PmProductPanelRelation) (*proto.PmProductPanelRelation, error) {
	t := orm.Use(iotmodel.GetDB()).TPmProductPanelRelation
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.ProductId != 0 { //整数
		updateField = append(updateField, t.ProductId)
	}
	if req.ControlPanelId != 0 { //整数
		updateField = append(updateField, t.ControlPanelId)
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
		logger.Error("UpdatePmProductPanelRelation error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.PmProductPanelRelation_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdatePmProductPanelRelation error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段PmProductPanelRelation
func (s *PmProductPanelRelationSvc) UpdateAllPmProductPanelRelation(req *proto.PmProductPanelRelation) (*proto.PmProductPanelRelation, error) {
	t := orm.Use(iotmodel.GetDB()).TPmProductPanelRelation
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.ProductId)
	updateField = append(updateField, t.ControlPanelId)
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
		logger.Error("UpdateAllPmProductPanelRelation error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.PmProductPanelRelation_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllPmProductPanelRelation error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *PmProductPanelRelationSvc) UpdateFieldsPmProductPanelRelation(req *proto.PmProductPanelRelationUpdateFieldsRequest) (*proto.PmProductPanelRelation, error) {
	t := orm.Use(iotmodel.GetDB()).TPmProductPanelRelation
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsPmProductPanelRelation error : missing updateField")
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
		logger.Error("UpdateFieldsPmProductPanelRelation error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.PmProductPanelRelation_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsPmProductPanelRelation error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找PmProductPanelRelation
func (s *PmProductPanelRelationSvc) FindPmProductPanelRelation(req *proto.PmProductPanelRelationFilter) (*proto.PmProductPanelRelation, error) {
	t := orm.Use(iotmodel.GetDB()).TPmProductPanelRelation
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.ProductId != 0 { //整数
		do = do.Where(t.ProductId.Eq(req.ProductId))
	}
	if req.ControlPanelId != 0 { //整数
		do = do.Where(t.ControlPanelId.Eq(req.ControlPanelId))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindPmProductPanelRelation error : %s", err.Error())
		return nil, err
	}
	res := convert.PmProductPanelRelation_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找PmProductPanelRelation
func (s *PmProductPanelRelationSvc) FindByIdPmProductPanelRelation(req *proto.PmProductPanelRelationFilter) (*proto.PmProductPanelRelation, error) {
	t := orm.Use(iotmodel.GetDB()).TPmProductPanelRelation
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdPmProductPanelRelation error : %s", err.Error())
		return nil, err
	}
	res := convert.PmProductPanelRelation_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找PmProductPanelRelation,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *PmProductPanelRelationSvc) GetListPmProductPanelRelation(req *proto.PmProductPanelRelationListRequest) ([]*proto.PmProductPanelRelation, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TPmProductPanelRelation
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.ProductId != 0 { //整数
			do = do.Where(t.ProductId.Eq(query.ProductId))
		}
		if query.ControlPanelId != 0 { //整数
			do = do.Where(t.ControlPanelId.Eq(query.ControlPanelId))
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

	var list []*model.TPmProductPanelRelation
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
		logger.Errorf("GetListPmProductPanelRelation error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.PmProductPanelRelation, len(list))
	for i, v := range list {
		result[i] = convert.PmProductPanelRelation_db2pb(v)
	}
	return result, total, nil
}

// 根据分页条件查找PmProductPanelRelation,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *PmProductPanelRelationSvc) GetListPmProductPanelRelationAndPanelStatus(req *proto.PmProductPanelRelationListRequest) ([]*proto.PmProductPanelRelation, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB())
	tPmProductPanelRelation := t.TPmProductPanelRelation
	tPmControlPanels := t.TPmControlPanels
	do := tPmProductPanelRelation.WithContext(context.Background()).LeftJoin(tPmControlPanels, tPmProductPanelRelation.ControlPanelId.EqCol(tPmControlPanels.Id))
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(tPmProductPanelRelation.Id.Eq(query.Id))
		}
		if query.ProductId != 0 { //整数
			do = do.Where(tPmProductPanelRelation.ProductId.Eq(query.ProductId))
		}
		if query.ControlPanelId != 0 { //整数
			do = do.Where(tPmProductPanelRelation.ControlPanelId.Eq(query.ControlPanelId))
		}
	}
	orderCol, ok := tPmProductPanelRelation.GetFieldByName(req.OrderKey)
	if !ok {
		orderCol = tPmProductPanelRelation.Id
	}
	if req.OrderDesc != "" {
		do = do.Order(orderCol.Desc())
	} else {
		do = do.Order(orderCol)
	}
	do = do.Where(tPmControlPanels.Status.Eq(1), tPmControlPanels.DeletedAt.IsNull())
	var list []*model.TPmProductPanelRelation
	var total int64
	err = do.Select(tPmProductPanelRelation.ALL).Scan(&list)
	if err != nil {
		logger.Errorf("GetListPmProductPanelRelation error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.PmProductPanelRelation, len(list))
	for i, v := range list {
		result[i] = convert.PmProductPanelRelation_db2pb(v)
	}
	return result, total, nil
}
