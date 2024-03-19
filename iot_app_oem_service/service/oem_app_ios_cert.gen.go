// Code generated by sgen.exe,2022-06-02 11:15:11. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	"cloud_platform/iot_app_oem_service/convert"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_app_oem/model"
	"cloud_platform/iot_model/db_app_oem/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OemAppIosCertSvc struct {
	Ctx context.Context
}

// 创建OemAppIosCert
func (s *OemAppIosCertSvc) CreateOemAppIosCert(req *proto.OemAppIosCert) (*proto.OemAppIosCert, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TOemAppIosCert
	do := t.WithContext(context.Background())
	dbObj := convert.OemAppIosCert_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateOemAppIosCert error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除OemAppIosCert
func (s *OemAppIosCertSvc) DeleteOemAppIosCert(req *proto.OemAppIosCert) (*proto.OemAppIosCert, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppIosCert
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.AppId != 0 { //整数
		do = do.Where(t.AppId.Eq(req.AppId))
	}
	if req.Version != "" { //字符串
		do = do.Where(t.Version.Eq(req.Version))
	}
	if req.DistProvision != "" { //字符串
		do = do.Where(t.DistProvision.Eq(req.DistProvision))
	}
	if req.DistCert != "" { //字符串
		do = do.Where(t.DistCert.Eq(req.DistCert))
	}
	if req.DistCertSecret != "" { //字符串
		do = do.Where(t.DistCertSecret.Eq(req.DistCertSecret))
	}
	if req.DistCertOfficial != "" { //字符串
		do = do.Where(t.DistCertOfficial.Eq(req.DistCertOfficial))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteOemAppIosCert error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除OemAppIosCert
func (s *OemAppIosCertSvc) DeleteByIdOemAppIosCert(req *proto.OemAppIosCert) (*proto.OemAppIosCert, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppIosCert
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdOemAppIosCert error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除OemAppIosCert
func (s *OemAppIosCertSvc) DeleteByIdsOemAppIosCert(req *proto.OemAppIosCertBatchDeleteRequest) (*proto.OemAppIosCertBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TOemAppIosCert
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsOemAppIosCert error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新OemAppIosCert
func (s *OemAppIosCertSvc) UpdateOemAppIosCert(req *proto.OemAppIosCert) (*proto.OemAppIosCert, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppIosCert
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.AppId != 0 { //整数
		updateField = append(updateField, t.AppId)
	}
	if req.Version != "" { //字符串
		updateField = append(updateField, t.Version)
	}
	if req.DistProvision != "" { //字符串
		updateField = append(updateField, t.DistProvision)
	}
	if req.DistCert != "" { //字符串
		updateField = append(updateField, t.DistCert)
	}
	if req.DistCertSecret != "" { //字符串
		updateField = append(updateField, t.DistCertSecret)
	}
	if req.DistCertOfficial != "" { //字符串
		updateField = append(updateField, t.DistCertOfficial)
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
		logger.Error("UpdateOemAppIosCert error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.OemAppIosCert_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateOemAppIosCert error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段OemAppIosCert
func (s *OemAppIosCertSvc) UpdateAllOemAppIosCert(req *proto.OemAppIosCert) (*proto.OemAppIosCert, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppIosCert
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.AppId)
	updateField = append(updateField, t.Version)
	updateField = append(updateField, t.DistProvision)
	updateField = append(updateField, t.DistCert)
	updateField = append(updateField, t.DistCertSecret)
	updateField = append(updateField, t.DistCertOfficial)

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
		logger.Error("UpdateAllOemAppIosCert error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OemAppIosCert_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllOemAppIosCert error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *OemAppIosCertSvc) UpdateFieldsOemAppIosCert(req *proto.OemAppIosCertUpdateFieldsRequest) (*proto.OemAppIosCert, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppIosCert
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsOemAppIosCert error : missing updateField")
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
		logger.Error("UpdateFieldsOemAppIosCert error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OemAppIosCert_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsOemAppIosCert error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找OemAppIosCert
func (s *OemAppIosCertSvc) FindOemAppIosCert(req *proto.OemAppIosCertFilter) (*proto.OemAppIosCert, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppIosCert
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.AppId != 0 { //整数
		do = do.Where(t.AppId.Eq(req.AppId))
	}
	if req.Version != "" { //字符串
		do = do.Where(t.Version.Eq(req.Version))
	}
	if req.DistProvision != "" { //字符串
		do = do.Where(t.DistProvision.Eq(req.DistProvision))
	}
	if req.DistCert != "" { //字符串
		do = do.Where(t.DistCert.Eq(req.DistCert))
	}
	if req.DistCertSecret != "" { //字符串
		do = do.Where(t.DistCertSecret.Eq(req.DistCertSecret))
	}
	if req.DistCertOfficial != "" { //字符串
		do = do.Where(t.DistCertOfficial.Eq(req.DistCertOfficial))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindOemAppIosCert error : %s", err.Error())
		return nil, err
	}
	res := convert.OemAppIosCert_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找OemAppIosCert
func (s *OemAppIosCertSvc) FindByIdOemAppIosCert(req *proto.OemAppIosCertFilter) (*proto.OemAppIosCert, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppIosCert
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdOemAppIosCert error : %s", err.Error())
		return nil, err
	}
	res := convert.OemAppIosCert_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找OemAppIosCert,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *OemAppIosCertSvc) GetListOemAppIosCert(req *proto.OemAppIosCertListRequest) ([]*proto.OemAppIosCert, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TOemAppIosCert
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.AppId != 0 { //整数
			do = do.Where(t.AppId.Eq(query.AppId))
		}
		if query.Version != "" { //字符串
			do = do.Where(t.Version.Like("%" + query.Version + "%"))
		}
		if query.DistProvision != "" { //字符串
			do = do.Where(t.DistProvision.Like("%" + query.DistProvision + "%"))
		}
		if query.DistCert != "" { //字符串
			do = do.Where(t.DistCert.Like("%" + query.DistCert + "%"))
		}
		if query.DistCertSecret != "" { //字符串
			do = do.Where(t.DistCertSecret.Like("%" + query.DistCertSecret + "%"))
		}
		if query.DistCertOfficial != "" { //字符串
			do = do.Where(t.DistCertOfficial.Like("%" + query.DistCertOfficial + "%"))
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

	var list []*model.TOemAppIosCert
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
		logger.Errorf("GetListOemAppIosCert error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.OemAppIosCert, len(list))
	for i, v := range list {
		result[i] = convert.OemAppIosCert_db2pb(v)
	}
	return result, total, nil
}
