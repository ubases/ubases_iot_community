// Code generated by sgen.exe,2022-04-18 19:12:09. DO NOT EDIT.
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

type SysModelInfoSvc struct {
	Ctx context.Context
}

// 创建SysModelInfo
func (s *SysModelInfoSvc) CreateSysModelInfo(req *proto.SysModelInfo) (*proto.SysModelInfo, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TSysModelInfo
	do := t.WithContext(context.Background())
	dbObj := convert.SysModelInfo_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateSysModelInfo error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除SysModelInfo
func (s *SysModelInfoSvc) DeleteSysModelInfo(req *proto.SysModelInfo) (*proto.SysModelInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TSysModelInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.ModelId != 0 { //整数
		do = do.Where(t.ModelId.Eq(req.ModelId))
	}
	if req.ModelCategoryId != 0 { //整数
		do = do.Where(t.ModelCategoryId.Eq(req.ModelCategoryId))
	}
	if req.ModelName != "" { //字符串
		do = do.Where(t.ModelName.Eq(req.ModelName))
	}
	if req.ModelTitle != "" { //字符串
		do = do.Where(t.ModelTitle.Eq(req.ModelTitle))
	}
	if req.ModelPk != "" { //字符串
		do = do.Where(t.ModelPk.Eq(req.ModelPk))
	}
	if req.ModelOrder != "" { //字符串
		do = do.Where(t.ModelOrder.Eq(req.ModelOrder))
	}
	if req.ModelSort != "" { //字符串
		do = do.Where(t.ModelSort.Eq(req.ModelSort))
	}
	if req.ModelList != "" { //字符串
		do = do.Where(t.ModelList.Eq(req.ModelList))
	}
	if req.ModelEdit != "" { //字符串
		do = do.Where(t.ModelEdit.Eq(req.ModelEdit))
	}
	if req.ModelIndexes != "" { //字符串
		do = do.Where(t.ModelIndexes.Eq(req.ModelIndexes))
	}
	if req.SearchList != "" { //字符串
		do = do.Where(t.SearchList.Eq(req.SearchList))
	}
	if req.CreateTime != 0 { //整数
		do = do.Where(t.CreateTime.Eq(req.CreateTime))
	}
	if req.UpdateTime != 0 { //整数
		do = do.Where(t.UpdateTime.Eq(req.UpdateTime))
	}
	if req.ModelStatus != 0 { //整数
		do = do.Where(t.ModelStatus.Eq(req.ModelStatus))
	}
	if req.ModelEngine != "" { //字符串
		do = do.Where(t.ModelEngine.Eq(req.ModelEngine))
	}
	if req.CreateBy != 0 { //整数
		do = do.Where(t.CreateBy.Eq(req.CreateBy))
	}
	if req.UpdateBy != 0 { //整数
		do = do.Where(t.UpdateBy.Eq(req.UpdateBy))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteSysModelInfo error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除SysModelInfo
func (s *SysModelInfoSvc) DeleteByIdSysModelInfo(req *proto.SysModelInfo) (*proto.SysModelInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TSysModelInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.ModelId != 0 { //整数
		do = do.Where(t.ModelId.Eq(req.ModelId))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdSysModelInfo error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除SysModelInfo
func (s *SysModelInfoSvc) DeleteByIdsSysModelInfo(req *proto.SysModelInfoBatchDeleteRequest) (*proto.SysModelInfoBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TSysModelInfo
		do := t.WithContext(context.Background())

		do = do.Where(t.ModelId.Eq(k.ModelId))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsSysModelInfo error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新SysModelInfo
func (s *SysModelInfoSvc) UpdateSysModelInfo(req *proto.SysModelInfo) (*proto.SysModelInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TSysModelInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.ModelCategoryId != 0 { //整数
		updateField = append(updateField, t.ModelCategoryId)
	}
	if req.ModelName != "" { //字符串
		updateField = append(updateField, t.ModelName)
	}
	if req.ModelTitle != "" { //字符串
		updateField = append(updateField, t.ModelTitle)
	}
	if req.ModelPk != "" { //字符串
		updateField = append(updateField, t.ModelPk)
	}
	if req.ModelOrder != "" { //字符串
		updateField = append(updateField, t.ModelOrder)
	}
	if req.ModelSort != "" { //字符串
		updateField = append(updateField, t.ModelSort)
	}
	if req.ModelList != "" { //字符串
		updateField = append(updateField, t.ModelList)
	}
	if req.ModelEdit != "" { //字符串
		updateField = append(updateField, t.ModelEdit)
	}
	if req.ModelIndexes != "" { //字符串
		updateField = append(updateField, t.ModelIndexes)
	}
	if req.SearchList != "" { //字符串
		updateField = append(updateField, t.SearchList)
	}
	if req.CreateTime != 0 { //整数
		updateField = append(updateField, t.CreateTime)
	}
	if req.UpdateTime != 0 { //整数
		updateField = append(updateField, t.UpdateTime)
	}
	if req.ModelStatus != 0 { //整数
		updateField = append(updateField, t.ModelStatus)
	}
	if req.ModelEngine != "" { //字符串
		updateField = append(updateField, t.ModelEngine)
	}
	if req.CreateBy != 0 { //整数
		updateField = append(updateField, t.CreateBy)
	}
	if req.UpdateBy != 0 { //整数
		updateField = append(updateField, t.UpdateBy)
	}
	if len(updateField) > 0 {
		do = do.Select(updateField...)
	}
	//主键条件
	HasPrimaryKey := false

	if req.ModelId != 0 { //整数
		do = do.Where(t.ModelId.Eq(req.ModelId))
		HasPrimaryKey = true
	}

	if !HasPrimaryKey {
		logger.Error("UpdateSysModelInfo error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.SysModelInfo_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateSysModelInfo error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段SysModelInfo
func (s *SysModelInfoSvc) UpdateAllSysModelInfo(req *proto.SysModelInfo) (*proto.SysModelInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TSysModelInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.ModelCategoryId)
	updateField = append(updateField, t.ModelName)
	updateField = append(updateField, t.ModelTitle)
	updateField = append(updateField, t.ModelPk)
	updateField = append(updateField, t.ModelOrder)
	updateField = append(updateField, t.ModelSort)
	updateField = append(updateField, t.ModelList)
	updateField = append(updateField, t.ModelEdit)
	updateField = append(updateField, t.ModelIndexes)
	updateField = append(updateField, t.SearchList)
	updateField = append(updateField, t.CreateTime)
	updateField = append(updateField, t.UpdateTime)
	updateField = append(updateField, t.ModelStatus)
	updateField = append(updateField, t.ModelEngine)
	updateField = append(updateField, t.CreateBy)
	updateField = append(updateField, t.UpdateBy)
	if len(updateField) > 0 {
		do = do.Select(updateField...)
	}
	//主键条件
	HasPrimaryKey := false
	if req.ModelId != 0 { //整数
		do = do.Where(t.ModelId.Eq(req.ModelId))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateAllSysModelInfo error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SysModelInfo_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllSysModelInfo error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *SysModelInfoSvc) UpdateFieldsSysModelInfo(req *proto.SysModelInfoUpdateFieldsRequest) (*proto.SysModelInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TSysModelInfo
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
	if req.Data.ModelId != 0 { //整数
		do = do.Where(t.ModelId.Eq(req.Data.ModelId))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateFieldsSysModelInfo error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SysModelInfo_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsSysModelInfo error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找SysModelInfo
func (s *SysModelInfoSvc) FindSysModelInfo(req *proto.SysModelInfoFilter) (*proto.SysModelInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TSysModelInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.ModelId != 0 { //整数
		do = do.Where(t.ModelId.Eq(req.ModelId))
	}
	if req.ModelCategoryId != 0 { //整数
		do = do.Where(t.ModelCategoryId.Eq(req.ModelCategoryId))
	}
	if req.ModelName != "" { //字符串
		do = do.Where(t.ModelName.Like("%" + req.ModelName + "%"))
	}
	if req.ModelTitle != "" { //字符串
		do = do.Where(t.ModelTitle.Like("%" + req.ModelTitle + "%"))
	}
	if req.ModelPk != "" { //字符串
		do = do.Where(t.ModelPk.Like("%" + req.ModelPk + "%"))
	}
	if req.ModelOrder != "" { //字符串
		do = do.Where(t.ModelOrder.Like("%" + req.ModelOrder + "%"))
	}
	if req.ModelSort != "" { //字符串
		do = do.Where(t.ModelSort.Like("%" + req.ModelSort + "%"))
	}
	if req.ModelList != "" { //字符串
		do = do.Where(t.ModelList.Like("%" + req.ModelList + "%"))
	}
	if req.ModelEdit != "" { //字符串
		do = do.Where(t.ModelEdit.Like("%" + req.ModelEdit + "%"))
	}
	if req.ModelIndexes != "" { //字符串
		do = do.Where(t.ModelIndexes.Like("%" + req.ModelIndexes + "%"))
	}
	if req.SearchList != "" { //字符串
		do = do.Where(t.SearchList.Like("%" + req.SearchList + "%"))
	}
	if req.CreateTime != 0 { //整数
		do = do.Where(t.CreateTime.Eq(req.CreateTime))
	}
	if req.UpdateTime != 0 { //整数
		do = do.Where(t.UpdateTime.Eq(req.UpdateTime))
	}
	if req.ModelStatus != 0 { //整数
		do = do.Where(t.ModelStatus.Eq(req.ModelStatus))
	}
	if req.ModelEngine != "" { //字符串
		do = do.Where(t.ModelEngine.Like("%" + req.ModelEngine + "%"))
	}
	if req.CreateBy != 0 { //整数
		do = do.Where(t.CreateBy.Eq(req.CreateBy))
	}
	if req.UpdateBy != 0 { //整数
		do = do.Where(t.UpdateBy.Eq(req.UpdateBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindSysModelInfo error : %s", err.Error())
		return nil, err
	}
	res := convert.SysModelInfo_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找SysModelInfo
func (s *SysModelInfoSvc) FindByIdSysModelInfo(req *proto.SysModelInfoFilter) (*proto.SysModelInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TSysModelInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.ModelId != 0 { //整数
		do = do.Where(t.ModelId.Eq(req.ModelId))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdSysModelInfo error : %s", err.Error())
		return nil, err
	}
	res := convert.SysModelInfo_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找SysModelInfo,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *SysModelInfoSvc) GetListSysModelInfo(req *proto.SysModelInfoListRequest) ([]*proto.SysModelInfo, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TSysModelInfo
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.ModelId != 0 { //整数
			do = do.Where(t.ModelId.Eq(query.ModelId))
		}
		if query.ModelCategoryId != 0 { //整数
			do = do.Where(t.ModelCategoryId.Eq(query.ModelCategoryId))
		}
		if query.ModelName != "" { //字符串
			do = do.Where(t.ModelName.Like("%" + query.ModelName + "%"))
		}
		if query.ModelTitle != "" { //字符串
			do = do.Where(t.ModelTitle.Like("%" + query.ModelTitle + "%"))
		}
		if query.ModelPk != "" { //字符串
			do = do.Where(t.ModelPk.Like("%" + query.ModelPk + "%"))
		}
		if query.ModelOrder != "" { //字符串
			do = do.Where(t.ModelOrder.Like("%" + query.ModelOrder + "%"))
		}
		if query.ModelSort != "" { //字符串
			do = do.Where(t.ModelSort.Like("%" + query.ModelSort + "%"))
		}
		if query.ModelList != "" { //字符串
			do = do.Where(t.ModelList.Like("%" + query.ModelList + "%"))
		}
		if query.ModelEdit != "" { //字符串
			do = do.Where(t.ModelEdit.Like("%" + query.ModelEdit + "%"))
		}
		if query.ModelIndexes != "" { //字符串
			do = do.Where(t.ModelIndexes.Like("%" + query.ModelIndexes + "%"))
		}
		if query.SearchList != "" { //字符串
			do = do.Where(t.SearchList.Like("%" + query.SearchList + "%"))
		}
		if query.CreateTime != 0 { //整数
			do = do.Where(t.CreateTime.Eq(query.CreateTime))
		}
		if query.UpdateTime != 0 { //整数
			do = do.Where(t.UpdateTime.Eq(query.UpdateTime))
		}
		if query.ModelStatus != 0 { //整数
			do = do.Where(t.ModelStatus.Eq(query.ModelStatus))
		}
		if query.ModelEngine != "" { //字符串
			do = do.Where(t.ModelEngine.Like("%" + query.ModelEngine + "%"))
		}
		if query.CreateBy != 0 { //整数
			do = do.Where(t.CreateBy.Eq(query.CreateBy))
		}
		if query.UpdateBy != 0 { //整数
			do = do.Where(t.UpdateBy.Eq(query.UpdateBy))
		}
	}
	orderCol, ok := t.GetFieldByName(req.OrderKey)
	if !ok {
		orderCol = t.ModelId
	}
	if req.OrderDesc != "" {
		do = do.Order(orderCol.Desc())
	} else {
		do = do.Order(orderCol)
	}

	var list []*model.TSysModelInfo
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
		logger.Errorf("GetListSysModelInfo error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.SysModelInfo, len(list))
	for i, v := range list {
		result[i] = convert.SysModelInfo_db2pb(v)
	}
	return result, total, nil
}
