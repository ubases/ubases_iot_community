// Code generated by sgen.exe,2022-04-18 19:12:06. DO NOT EDIT.
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

type DemoGenTreeSvc struct {
	Ctx context.Context
}

// 创建DemoGenTree
func (s *DemoGenTreeSvc) CreateDemoGenTree(req *proto.DemoGenTree) (*proto.DemoGenTree, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TDemoGenTree
	do := t.WithContext(context.Background())
	dbObj := convert.DemoGenTree_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateDemoGenTree error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除DemoGenTree
func (s *DemoGenTreeSvc) DeleteDemoGenTree(req *proto.DemoGenTree) (*proto.DemoGenTree, error) {
	t := orm.Use(iotmodel.GetDB()).TDemoGenTree
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.ParentId != 0 { //整数
		do = do.Where(t.ParentId.Eq(req.ParentId))
	}
	if req.DemoName != "" { //字符串
		do = do.Where(t.DemoName.Eq(req.DemoName))
	}
	if req.DemoAge != 0 { //整数
		do = do.Where(t.DemoAge.Eq(req.DemoAge))
	}
	if req.Classes != "" { //字符串
		do = do.Where(t.Classes.Eq(req.Classes))
	}
	if req.TDemoGender != 0 { //整数
		do = do.Where(t.TDemoGender.Eq(req.TDemoGender))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	if req.DemoStatus != 0 { //整数
		do = do.Where(t.DemoStatus.Eq(req.DemoStatus))
	}
	if req.DemoCate != "" { //字符串
		do = do.Where(t.DemoCate.Eq(req.DemoCate))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteDemoGenTree error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除DemoGenTree
func (s *DemoGenTreeSvc) DeleteByIdDemoGenTree(req *proto.DemoGenTree) (*proto.DemoGenTree, error) {
	t := orm.Use(iotmodel.GetDB()).TDemoGenTree
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdDemoGenTree error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除DemoGenTree
func (s *DemoGenTreeSvc) DeleteByIdsDemoGenTree(req *proto.DemoGenTreeBatchDeleteRequest) (*proto.DemoGenTreeBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TDemoGenTree
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsDemoGenTree error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新DemoGenTree
func (s *DemoGenTreeSvc) UpdateDemoGenTree(req *proto.DemoGenTree) (*proto.DemoGenTree, error) {
	t := orm.Use(iotmodel.GetDB()).TDemoGenTree
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.ParentId != 0 { //整数
		updateField = append(updateField, t.ParentId)
	}
	if req.DemoName != "" { //字符串
		updateField = append(updateField, t.DemoName)
	}
	if req.DemoAge != 0 { //整数
		updateField = append(updateField, t.DemoAge)
	}
	if req.Classes != "" { //字符串
		updateField = append(updateField, t.Classes)
	}
	if req.TDemoGender != 0 { //整数
		updateField = append(updateField, t.TDemoGender)
	}
	if req.CreatedBy != 0 { //整数
		updateField = append(updateField, t.CreatedBy)
	}
	if req.UpdatedBy != 0 { //整数
		updateField = append(updateField, t.UpdatedBy)
	}
	if req.DemoStatus != 0 { //整数
		updateField = append(updateField, t.DemoStatus)
	}
	if req.DemoCate != "" { //字符串
		updateField = append(updateField, t.DemoCate)
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
		logger.Error("UpdateDemoGenTree error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.DemoGenTree_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateDemoGenTree error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段DemoGenTree
func (s *DemoGenTreeSvc) UpdateAllDemoGenTree(req *proto.DemoGenTree) (*proto.DemoGenTree, error) {
	t := orm.Use(iotmodel.GetDB()).TDemoGenTree
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.ParentId)
	updateField = append(updateField, t.DemoName)
	updateField = append(updateField, t.DemoAge)
	updateField = append(updateField, t.Classes)
	updateField = append(updateField, t.DemoBorn)
	updateField = append(updateField, t.TDemoGender)
	updateField = append(updateField, t.CreatedBy)
	updateField = append(updateField, t.UpdatedBy)
	updateField = append(updateField, t.DemoStatus)
	updateField = append(updateField, t.DemoCate)
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
		logger.Error("UpdateAllDemoGenTree error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.DemoGenTree_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllDemoGenTree error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *DemoGenTreeSvc) UpdateFieldsDemoGenTree(req *proto.DemoGenTreeUpdateFieldsRequest) (*proto.DemoGenTree, error) {
	t := orm.Use(iotmodel.GetDB()).TDemoGenTree
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
		logger.Error("UpdateFieldsDemoGenTree error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.DemoGenTree_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsDemoGenTree error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找DemoGenTree
func (s *DemoGenTreeSvc) FindDemoGenTree(req *proto.DemoGenTreeFilter) (*proto.DemoGenTree, error) {
	t := orm.Use(iotmodel.GetDB()).TDemoGenTree
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.ParentId != 0 { //整数
		do = do.Where(t.ParentId.Eq(req.ParentId))
	}
	if req.DemoName != "" { //字符串
		do = do.Where(t.DemoName.Like("%" + req.DemoName + "%"))
	}
	if req.DemoAge != 0 { //整数
		do = do.Where(t.DemoAge.Eq(req.DemoAge))
	}
	if req.Classes != "" { //字符串
		do = do.Where(t.Classes.Like("%" + req.Classes + "%"))
	}
	if req.TDemoGender != 0 { //整数
		do = do.Where(t.TDemoGender.Eq(req.TDemoGender))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	if req.DemoStatus != 0 { //整数
		do = do.Where(t.DemoStatus.Eq(req.DemoStatus))
	}
	if req.DemoCate != "" { //字符串
		do = do.Where(t.DemoCate.Like("%" + req.DemoCate + "%"))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindDemoGenTree error : %s", err.Error())
		return nil, err
	}
	res := convert.DemoGenTree_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找DemoGenTree
func (s *DemoGenTreeSvc) FindByIdDemoGenTree(req *proto.DemoGenTreeFilter) (*proto.DemoGenTree, error) {
	t := orm.Use(iotmodel.GetDB()).TDemoGenTree
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdDemoGenTree error : %s", err.Error())
		return nil, err
	}
	res := convert.DemoGenTree_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找DemoGenTree,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *DemoGenTreeSvc) GetListDemoGenTree(req *proto.DemoGenTreeListRequest) ([]*proto.DemoGenTree, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TDemoGenTree
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.ParentId != 0 { //整数
			do = do.Where(t.ParentId.Eq(query.ParentId))
		}
		if query.DemoName != "" { //字符串
			do = do.Where(t.DemoName.Like("%" + query.DemoName + "%"))
		}
		if query.DemoAge != 0 { //整数
			do = do.Where(t.DemoAge.Eq(query.DemoAge))
		}
		if query.Classes != "" { //字符串
			do = do.Where(t.Classes.Like("%" + query.Classes + "%"))
		}
		if query.TDemoGender != 0 { //整数
			do = do.Where(t.TDemoGender.Eq(query.TDemoGender))
		}
		if query.CreatedBy != 0 { //整数
			do = do.Where(t.CreatedBy.Eq(query.CreatedBy))
		}
		if query.UpdatedBy != 0 { //整数
			do = do.Where(t.UpdatedBy.Eq(query.UpdatedBy))
		}
		if query.DemoStatus != 0 { //整数
			do = do.Where(t.DemoStatus.Eq(query.DemoStatus))
		}
		if query.DemoCate != "" { //字符串
			do = do.Where(t.DemoCate.Like("%" + query.DemoCate + "%"))
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

	var list []*model.TDemoGenTree
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
		logger.Errorf("GetListDemoGenTree error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.DemoGenTree, len(list))
	for i, v := range list {
		result[i] = convert.DemoGenTree_db2pb(v)
	}
	return result, total, nil
}
