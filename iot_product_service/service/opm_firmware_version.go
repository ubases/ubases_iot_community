// Code generated by sgen.exe,2022-04-29 15:04:30. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"cloud_platform/iot_product_service/convert"
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_model/db_product/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OpmFirmwareVersionSvc struct {
	Ctx context.Context
}

// 创建OpmFirmwareVersion
func (s *OpmFirmwareVersionSvc) CreateOpmFirmwareVersion(req *proto.OpmFirmwareVersion) (*proto.OpmFirmwareVersion, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TOpmFirmwareVersion
	do := t.WithContext(context.Background())
	dbObj := convert.OpmFirmwareVersion_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateOpmFirmwareVersion error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除OpmFirmwareVersion
func (s *OpmFirmwareVersionSvc) DeleteOpmFirmwareVersion(req *proto.OpmFirmwareVersion) (*proto.OpmFirmwareVersion, error) {
	t := orm.Use(iotmodel.GetDB()).TOpmFirmwareVersion
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.FirmwareId != 0 { //整数
		do = do.Where(t.FirmwareId.Eq(req.FirmwareId))
	}
	if req.Version != "" { //字符串
		do = do.Where(t.Version.Eq(req.Version))
	}
	if req.Desc != "" { //字符串
		do = do.Where(t.Desc.Eq(req.Desc))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteOpmFirmwareVersion error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除OpmFirmwareVersion
func (s *OpmFirmwareVersionSvc) DeleteByIdOpmFirmwareVersion(req *proto.OpmFirmwareVersion) (*proto.OpmFirmwareVersion, error) {
	t := orm.Use(iotmodel.GetDB()).TOpmFirmwareVersion
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdOpmFirmwareVersion error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除OpmFirmwareVersion
func (s *OpmFirmwareVersionSvc) DeleteByIdsOpmFirmwareVersion(req *proto.OpmFirmwareVersionBatchDeleteRequest) (*proto.OpmFirmwareVersionBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TOpmFirmwareVersion
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsOpmFirmwareVersion error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新OpmFirmwareVersion
func (s *OpmFirmwareVersionSvc) UpdateOpmFirmwareVersion(req *proto.OpmFirmwareVersion) (*proto.OpmFirmwareVersion, error) {
	t := orm.Use(iotmodel.GetDB()).TOpmFirmwareVersion
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.FirmwareId != 0 { //整数
		updateField = append(updateField, t.FirmwareId)
	}
	if req.Version != "" { //字符串
		updateField = append(updateField, t.Version)
	}
	if req.Desc != "" { //字符串
		updateField = append(updateField, t.Desc)
	}
	if req.Status != 0 { //整数
		updateField = append(updateField, t.Status)
	}
	if req.UpgradeFileName != "" { //字符串
		updateField = append(updateField, t.UpgradeFileName)
	}
	if req.UpgradeFilePath != "" { //字符串
		updateField = append(updateField, t.UpgradeFilePath)
	}
	if req.UpgradeFileKey != "" { //字符串
		updateField = append(updateField, t.UpgradeFileKey)
	}
	if req.UpgradeFileSize != 0 { //整数
		updateField = append(updateField, t.UpgradeFileSize)
	}
	if req.IsMust != 0 { //整数
		updateField = append(updateField, t.IsMust)
	}
	if req.UpgradeMode != 0 { //整数
		updateField = append(updateField, t.UpgradeMode)
	}
	if req.ProdFileName != "" { //字符串
		updateField = append(updateField, t.ProdFileName)
	}
	if req.ProdFilePath != "" { //字符串
		updateField = append(updateField, t.ProdFilePath)
	}
	if req.ProdFileKey != "" { //字符串
		updateField = append(updateField, t.ProdFileKey)
	}
	if req.ProdFileSize != 0 { //整数
		updateField = append(updateField, t.ProdFileSize)
	}

	if req.UpdatedBy != 0 { //整数
		updateField = append(updateField, t.UpdatedBy)
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
		logger.Error("UpdateOpmFirmwareVersion error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.OpmFirmwareVersion_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateOpmFirmwareVersion error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段OpmFirmwareVersion
func (s *OpmFirmwareVersionSvc) UpdateAllOpmFirmwareVersion(req *proto.OpmFirmwareVersion) (*proto.OpmFirmwareVersion, error) {
	t := orm.Use(iotmodel.GetDB()).TOpmFirmwareVersion
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.FirmwareId)
	updateField = append(updateField, t.Version)
	updateField = append(updateField, t.Desc)
	updateField = append(updateField, t.Status)
	updateField = append(updateField, t.UpgradeMode)
	updateField = append(updateField, t.IsMust)
	updateField = append(updateField, t.UpgradeFileName)
	updateField = append(updateField, t.UpgradeFilePath)
	updateField = append(updateField, t.UpgradeFileKey)
	updateField = append(updateField, t.UpgradeFileSize)
	updateField = append(updateField, t.ProdFileName)
	updateField = append(updateField, t.ProdFilePath)
	updateField = append(updateField, t.ProdFileKey)
	updateField = append(updateField, t.ProdFileSize)
	updateField = append(updateField, t.UpdatedBy)
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
		logger.Error("UpdateAllOpmFirmwareVersion error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OpmFirmwareVersion_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllOpmFirmwareVersion error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *OpmFirmwareVersionSvc) UpdateFieldsOpmFirmwareVersion(req *proto.OpmFirmwareVersionUpdateFieldsRequest) (*proto.OpmFirmwareVersion, error) {
	t := orm.Use(iotmodel.GetDB()).TOpmFirmwareVersion
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsOpmFirmwareVersion error : missing updateField")
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
		logger.Error("UpdateFieldsOpmFirmwareVersion error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OpmFirmwareVersion_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsOpmFirmwareVersion error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找OpmFirmwareVersion
func (s *OpmFirmwareVersionSvc) FindOpmFirmwareVersion(req *proto.OpmFirmwareVersionFilter) (*proto.OpmFirmwareVersion, error) {
	t := orm.Use(iotmodel.GetDB()).TOpmFirmwareVersion
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.FirmwareId != 0 { //整数
		do = do.Where(t.FirmwareId.Eq(req.FirmwareId))
	}
	if req.Version != "" { //字符串
		do = do.Where(t.Version.Eq(req.Version))
	}
	if req.Desc != "" { //字符串
		do = do.Where(t.Desc.Eq(req.Desc))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindOpmFirmwareVersion error : %s", err.Error())
		return nil, err
	}
	res := convert.OpmFirmwareVersion_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找OpmFirmwareVersion
func (s *OpmFirmwareVersionSvc) FindByIdOpmFirmwareVersion(req *proto.OpmFirmwareVersionFilter) (*proto.OpmFirmwareVersion, error) {
	t := orm.Use(iotmodel.GetDB()).TOpmFirmwareVersion
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdOpmFirmwareVersion error : %s", err.Error())
		return nil, err
	}
	res := convert.OpmFirmwareVersion_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找OpmFirmwareVersion,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *OpmFirmwareVersionSvc) GetListOpmFirmwareVersion(req *proto.OpmFirmwareVersionListRequest) ([]*proto.OpmFirmwareVersion, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TOpmFirmwareVersion
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.FirmwareId != 0 { //整数
			do = do.Where(t.FirmwareId.Eq(query.FirmwareId))
		}
		if query.Version != "" { //字符串
			do = do.Where(t.Version.Like("%" + query.Version + "%"))
		}
		if query.Desc != "" { //字符串
			do = do.Where(t.Desc.Like("%" + query.Desc + "%"))
		}
		if query.Status >= 0 { //整数
			do = do.Where(t.Status.Eq(query.Status))
		}
		if query.UpdatedBy != 0 { //整数
			do = do.Where(t.UpdatedBy.Eq(query.UpdatedBy))
		}
	}
	orderCol, ok := t.GetFieldByName(req.OrderKey)
	if !ok {
		do = do.Order(field.Func.VersionOrder(t.Version))
	} else {
		if req.OrderDesc != "" {
			do = do.Order(orderCol.Desc())
		} else {
			do = do.Order(orderCol)
		}
	}
	var list []*model.TOpmFirmwareVersion
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
		logger.Errorf("GetListOpmFirmwareVersion error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.OpmFirmwareVersion, len(list))
	for i, v := range list {
		result[i] = convert.OpmFirmwareVersion_db2pb(v)
	}
	return result, total, nil
}
