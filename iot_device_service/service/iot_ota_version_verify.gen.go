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

type IotOtaVersionVerifySvc struct {
	Ctx context.Context
}

// 创建IotOtaVersionVerify
func (s *IotOtaVersionVerifySvc) CreateIotOtaVersionVerify(req *proto.IotOtaVersionVerify) (*proto.IotOtaVersionVerify, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TIotOtaVersionVerify
	do := t.WithContext(context.Background())
	dbObj := convert.IotOtaVersionVerify_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateIotOtaVersionVerify error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除IotOtaVersionVerify
func (s *IotOtaVersionVerifySvc) DeleteIotOtaVersionVerify(req *proto.IotOtaVersionVerify) (*proto.IotOtaVersionVerify, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaVersionVerify
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.VersionId != 0 { //整数
		do = do.Where(t.VersionId.Eq(req.VersionId))
	}
	if req.DeviceVersion != "" { //字符串
		do = do.Where(t.DeviceVersion.Eq(req.DeviceVersion))
	}
	if req.Did != "" { //字符串
		do = do.Where(t.Did.Eq(req.Did))
	}
	if req.DeviceId != 0 { //整数
		do = do.Where(t.DeviceId.Eq(req.DeviceId))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.DeviceLog != "" { //字符串
		do = do.Where(t.DeviceLog.Eq(req.DeviceLog))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteIotOtaVersionVerify error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除IotOtaVersionVerify
func (s *IotOtaVersionVerifySvc) DeleteByIdIotOtaVersionVerify(req *proto.IotOtaVersionVerify) (*proto.IotOtaVersionVerify, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaVersionVerify
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdIotOtaVersionVerify error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除IotOtaVersionVerify
func (s *IotOtaVersionVerifySvc) DeleteByIdsIotOtaVersionVerify(req *proto.IotOtaVersionVerifyBatchDeleteRequest) (*proto.IotOtaVersionVerifyBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TIotOtaVersionVerify
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsIotOtaVersionVerify error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新IotOtaVersionVerify
func (s *IotOtaVersionVerifySvc) UpdateIotOtaVersionVerify(req *proto.IotOtaVersionVerify) (*proto.IotOtaVersionVerify, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaVersionVerify
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.VersionId != 0 { //整数
		updateField = append(updateField, t.VersionId)
	}
	if req.DeviceVersion != "" { //字符串
		updateField = append(updateField, t.DeviceVersion)
	}
	if req.Did != "" { //字符串
		updateField = append(updateField, t.Did)
	}
	if req.DeviceId != 0 { //整数
		updateField = append(updateField, t.DeviceId)
	}
	if req.Status != 0 { //整数
		updateField = append(updateField, t.Status)
	}
	if req.DeviceLog != "" { //字符串
		updateField = append(updateField, t.DeviceLog)
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
		logger.Error("UpdateIotOtaVersionVerify error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.IotOtaVersionVerify_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateIotOtaVersionVerify error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段IotOtaVersionVerify
func (s *IotOtaVersionVerifySvc) UpdateAllIotOtaVersionVerify(req *proto.IotOtaVersionVerify) (*proto.IotOtaVersionVerify, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaVersionVerify
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.VersionId)
	updateField = append(updateField, t.DeviceVersion)
	updateField = append(updateField, t.Did)
	updateField = append(updateField, t.DeviceId)
	updateField = append(updateField, t.Status)
	updateField = append(updateField, t.DeviceLog)
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
		logger.Error("UpdateAllIotOtaVersionVerify error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.IotOtaVersionVerify_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllIotOtaVersionVerify error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *IotOtaVersionVerifySvc) UpdateFieldsIotOtaVersionVerify(req *proto.IotOtaVersionVerifyUpdateFieldsRequest) (*proto.IotOtaVersionVerify, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaVersionVerify
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsIotOtaVersionVerify error : missing updateField")
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
		logger.Error("UpdateFieldsIotOtaVersionVerify error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.IotOtaVersionVerify_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsIotOtaVersionVerify error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找IotOtaVersionVerify
func (s *IotOtaVersionVerifySvc) FindIotOtaVersionVerify(req *proto.IotOtaVersionVerifyFilter) (*proto.IotOtaVersionVerify, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaVersionVerify
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.VersionId != 0 { //整数
		do = do.Where(t.VersionId.Eq(req.VersionId))
	}
	if req.DeviceVersion != "" { //字符串
		do = do.Where(t.DeviceVersion.Like("%" + req.DeviceVersion + "%"))
	}
	if req.Did != "" { //字符串
		do = do.Where(t.Did.Like("%" + req.Did + "%"))
	}
	if req.DeviceId != 0 { //整数
		do = do.Where(t.DeviceId.Eq(req.DeviceId))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.DeviceLog != "" { //字符串
		do = do.Where(t.DeviceLog.Like("%" + req.DeviceLog + "%"))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindIotOtaVersionVerify error : %s", err.Error())
		return nil, err
	}
	res := convert.IotOtaVersionVerify_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找IotOtaVersionVerify
func (s *IotOtaVersionVerifySvc) FindByIdIotOtaVersionVerify(req *proto.IotOtaVersionVerifyFilter) (*proto.IotOtaVersionVerify, error) {
	t := orm.Use(iotmodel.GetDB()).TIotOtaVersionVerify
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdIotOtaVersionVerify error : %s", err.Error())
		return nil, err
	}
	res := convert.IotOtaVersionVerify_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找IotOtaVersionVerify,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *IotOtaVersionVerifySvc) GetListIotOtaVersionVerify(req *proto.IotOtaVersionVerifyListRequest) ([]*proto.IotOtaVersionVerify, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TIotOtaVersionVerify
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.VersionId != 0 { //整数
			do = do.Where(t.VersionId.Eq(query.VersionId))
		}
		if query.DeviceVersion != "" { //字符串
			do = do.Where(t.DeviceVersion.Like("%" + query.DeviceVersion + "%"))
		}
		if query.Did != "" { //字符串
			do = do.Where(t.Did.Like("%" + query.Did + "%"))
		}
		if query.DeviceId != 0 { //整数
			do = do.Where(t.DeviceId.Eq(query.DeviceId))
		}
		if query.Status != 0 { //整数
			do = do.Where(t.Status.Eq(query.Status))
		}
		if query.DeviceLog != "" { //字符串
			do = do.Where(t.DeviceLog.Like("%" + query.DeviceLog + "%"))
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

	var list []*model.TIotOtaVersionVerify
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
		logger.Errorf("GetListIotOtaVersionVerify error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.IotOtaVersionVerify, len(list))
	for i, v := range list {
		result[i] = convert.IotOtaVersionVerify_db2pb(v)
	}
	return result, total, nil
}
