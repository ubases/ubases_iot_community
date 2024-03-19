// Code generated by sgen.exe,2022-04-21 13:46:12. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	"cloud_platform/iot_device_service/convert"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_model/db_device/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type IotOtaPkgSvc struct {
	Ctx context.Context
}

// 创建IotOtaPkg
func (s *IotOtaPkgSvc) CreateIotOtaPkg(req *proto.IotOtaPkg) (*proto.IotOtaPkg, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TIotOtaPkg
	do := t.WithContext(context.Background())
	dbObj := convert.IotOtaPkg_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateIotOtaPkg error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除IotOtaPkg
func (s *IotOtaPkgSvc) DeleteIotOtaPkg(req *proto.IotOtaPkg) (*proto.IotOtaPkg, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaPkg
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.Type != 0 { //整数
		do = do.Where(t.Type.Eq(req.Type))
	}
	if req.Name != "" { //字符串
		do = do.Where(t.Name.Eq(req.Name))
	}
	if req.Version != "" { //字符串
		do = do.Where(t.Version.Eq(req.Version))
	}
	if req.UpgradeMode != 0 { //整数
		do = do.Where(t.UpgradeMode.Eq(req.UpgradeMode))
	}
	if req.Url != "" { //字符串
		do = do.Where(t.Url.Eq(req.Url))
	}
	if req.KeyVersionFlag != 0 { //整数
		do = do.Where(t.KeyVersionFlag.Eq(req.KeyVersionFlag))
	}
	if req.SystemType != 0 { //整数
		do = do.Where(t.SystemType.Eq(req.SystemType))
	}
	if req.MinimumEcuRequired != "" { //字符串
		do = do.Where(t.MinimumEcuRequired.Eq(req.MinimumEcuRequired))
	}
	if req.MinimumMcuRequired != "" { //字符串
		do = do.Where(t.MinimumMcuRequired.Eq(req.MinimumMcuRequired))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteIotOtaPkg error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除IotOtaPkg
func (s *IotOtaPkgSvc) DeleteByIdIotOtaPkg(req *proto.IotOtaPkg) (*proto.IotOtaPkg, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaPkg
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdIotOtaPkg error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除IotOtaPkg
func (s *IotOtaPkgSvc) DeleteByIdsIotOtaPkg(req *proto.IotOtaPkgBatchDeleteRequest) (*proto.IotOtaPkgBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TIotOtaPkg
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsIotOtaPkg error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新IotOtaPkg
func (s *IotOtaPkgSvc) UpdateIotOtaPkg(req *proto.IotOtaPkg) (*proto.IotOtaPkg, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaPkg
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.Type != 0 { //整数
		updateField = append(updateField, t.Type)
	}
	if req.Name != "" { //字符串
		updateField = append(updateField, t.Name)
	}
	if req.Version != "" { //字符串
		updateField = append(updateField, t.Version)
	}
	if req.UpgradeMode != 0 { //整数
		updateField = append(updateField, t.UpgradeMode)
	}
	if req.Url != "" { //字符串
		updateField = append(updateField, t.Url)
	}
	if req.KeyVersionFlag != 0 { //整数
		updateField = append(updateField, t.KeyVersionFlag)
	}
	if req.SystemType != 0 { //整数
		updateField = append(updateField, t.SystemType)
	}
	if req.MinimumEcuRequired != "" { //字符串
		updateField = append(updateField, t.MinimumEcuRequired)
	}
	if req.MinimumMcuRequired != "" { //字符串
		updateField = append(updateField, t.MinimumMcuRequired)
	}
	if req.Status != 0 { //整数
		updateField = append(updateField, t.Status)
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

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
		HasPrimaryKey = true
	}

	if !HasPrimaryKey {
		logger.Error("UpdateIotOtaPkg error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.IotOtaPkg_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateIotOtaPkg error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段IotOtaPkg
func (s *IotOtaPkgSvc) UpdateAllIotOtaPkg(req *proto.IotOtaPkg) (*proto.IotOtaPkg, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaPkg
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.Type)
	updateField = append(updateField, t.Name)
	updateField = append(updateField, t.Version)
	updateField = append(updateField, t.UpgradeMode)
	updateField = append(updateField, t.Url)
	updateField = append(updateField, t.KeyVersionFlag)
	updateField = append(updateField, t.SystemType)
	updateField = append(updateField, t.MinimumEcuRequired)
	updateField = append(updateField, t.MinimumMcuRequired)
	updateField = append(updateField, t.Status)
	updateField = append(updateField, t.UploadTime)
	updateField = append(updateField, t.CreatedBy)
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
		logger.Error("UpdateAllIotOtaPkg error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.IotOtaPkg_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllIotOtaPkg error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *IotOtaPkgSvc) UpdateFieldsIotOtaPkg(req *proto.IotOtaPkgUpdateFieldsRequest) (*proto.IotOtaPkg, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaPkg
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsIotOtaPkg error : missing updateField")
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
		logger.Error("UpdateFieldsIotOtaPkg error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.IotOtaPkg_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsIotOtaPkg error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找IotOtaPkg
func (s *IotOtaPkgSvc) FindIotOtaPkg(req *proto.IotOtaPkgFilter) (*proto.IotOtaPkg, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaPkg
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.Type != 0 { //整数
		do = do.Where(t.Type.Eq(req.Type))
	}
	if req.Name != "" { //字符串
		do = do.Where(t.Name.Like("%" + req.Name + "%"))
	}
	if req.Version != "" { //字符串
		do = do.Where(t.Version.Like("%" + req.Version + "%"))
	}
	if req.UpgradeMode != 0 { //整数
		do = do.Where(t.UpgradeMode.Eq(req.UpgradeMode))
	}
	if req.Url != "" { //字符串
		do = do.Where(t.Url.Like("%" + req.Url + "%"))
	}
	if req.KeyVersionFlag != 0 { //整数
		do = do.Where(t.KeyVersionFlag.Eq(req.KeyVersionFlag))
	}
	if req.SystemType != 0 { //整数
		do = do.Where(t.SystemType.Eq(req.SystemType))
	}
	if req.MinimumEcuRequired != "" { //字符串
		do = do.Where(t.MinimumEcuRequired.Like("%" + req.MinimumEcuRequired + "%"))
	}
	if req.MinimumMcuRequired != "" { //字符串
		do = do.Where(t.MinimumMcuRequired.Like("%" + req.MinimumMcuRequired + "%"))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindIotOtaPkg error : %s", err.Error())
		return nil, err
	}
	res := convert.IotOtaPkg_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找IotOtaPkg
func (s *IotOtaPkgSvc) FindByIdIotOtaPkg(req *proto.IotOtaPkgFilter) (*proto.IotOtaPkg, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaPkg
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdIotOtaPkg error : %s", err.Error())
		return nil, err
	}
	res := convert.IotOtaPkg_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找IotOtaPkg,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *IotOtaPkgSvc) GetListIotOtaPkg(req *proto.IotOtaPkgListRequest) ([]*proto.IotOtaPkg, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TIotOtaPkg
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.Type != 0 { //整数
			do = do.Where(t.Type.Eq(query.Type))
		}
		if query.Name != "" { //字符串
			do = do.Where(t.Name.Like("%" + query.Name + "%"))
		}
		if query.Version != "" { //字符串
			do = do.Where(t.Version.Like("%" + query.Version + "%"))
		}
		if query.UpgradeMode != 0 { //整数
			do = do.Where(t.UpgradeMode.Eq(query.UpgradeMode))
		}
		if query.Url != "" { //字符串
			do = do.Where(t.Url.Like("%" + query.Url + "%"))
		}
		if query.KeyVersionFlag != 0 { //整数
			do = do.Where(t.KeyVersionFlag.Eq(query.KeyVersionFlag))
		}
		if query.SystemType != 0 { //整数
			do = do.Where(t.SystemType.Eq(query.SystemType))
		}
		if query.MinimumEcuRequired != "" { //字符串
			do = do.Where(t.MinimumEcuRequired.Like("%" + query.MinimumEcuRequired + "%"))
		}
		if query.MinimumMcuRequired != "" { //字符串
			do = do.Where(t.MinimumMcuRequired.Like("%" + query.MinimumMcuRequired + "%"))
		}
		if query.Status != 0 { //整数
			do = do.Where(t.Status.Eq(query.Status))
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
		orderCol = t.Id
	}
	if req.OrderDesc != "" {
		do = do.Order(orderCol.Desc())
	} else {
		do = do.Order(orderCol)
	}

	var list []*model.TIotOtaPkg
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
		logger.Errorf("GetListIotOtaPkg error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.IotOtaPkg, len(list))
	for i, v := range list {
		result[i] = convert.IotOtaPkg_db2pb(v)
	}
	return result, total, nil
}
