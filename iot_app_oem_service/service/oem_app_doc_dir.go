// Code generated by sgen.exe,2022-07-14 15:09:42. DO NOT EDIT.
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

type OemAppDocDirSvc struct {
	Ctx context.Context
}

// 创建OemAppDocDir
func (s *OemAppDocDirSvc) CreateOemAppDocDir(req *proto.OemAppDocDir) (*proto.OemAppDocDir, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
	do := t.WithContext(context.Background())
	dbObj := convert.OemAppDocDir_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateOemAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除OemAppDocDir
func (s *OemAppDocDirSvc) DeleteOemAppDocDir(req *proto.OemAppDocDir) (*proto.OemAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.DocId != 0 { //整数
		do = do.Where(t.DocId.Eq(req.DocId))
	}
	if req.ParentId != 0 { //整数
		do = do.Where(t.ParentId.Eq(req.ParentId))
	}
	if req.DirName != "" { //字符串
		do = do.Where(t.DirName.Eq(req.DirName))
	}
	if req.DirImg != "" { //字符串
		do = do.Where(t.DirImg.Eq(req.DirImg))
	}
	if req.Sort != 0 { //整数
		do = do.Where(t.Sort.Eq(req.Sort))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteOemAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除OemAppDocDir
func (s *OemAppDocDirSvc) DeleteByIdOemAppDocDir(req *proto.OemAppDocDir) (*proto.OemAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdOemAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除OemAppDocDir
func (s *OemAppDocDirSvc) DeleteByIdsOemAppDocDir(req *proto.OemAppDocDirBatchDeleteRequest) (*proto.OemAppDocDirBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsOemAppDocDir error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新OemAppDocDir
func (s *OemAppDocDirSvc) UpdateOemAppDocDir(req *proto.OemAppDocDir) (*proto.OemAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.DocId != 0 { //整数
		updateField = append(updateField, t.DocId)
	}
	if req.ParentId != 0 { //整数
		updateField = append(updateField, t.ParentId)
	}
	if req.DirName != "" { //字符串
		updateField = append(updateField, t.DirName)
	}
	if req.DirImg != "" { //字符串
		updateField = append(updateField, t.DirImg)
	}
	if req.Sort != 0 { //整数
		updateField = append(updateField, t.Sort)
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
		logger.Error("UpdateOemAppDocDir error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.OemAppDocDir_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateOemAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段OemAppDocDir
func (s *OemAppDocDirSvc) UpdateAllOemAppDocDir(req *proto.OemAppDocDir) (*proto.OemAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.DocId)
	updateField = append(updateField, t.ParentId)
	updateField = append(updateField, t.DirName)
	updateField = append(updateField, t.DirImg)
	updateField = append(updateField, t.Sort)
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
		logger.Error("UpdateAllOemAppDocDir error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OemAppDocDir_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllOemAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *OemAppDocDirSvc) UpdateFieldsOemAppDocDir(req *proto.OemAppDocDirUpdateFieldsRequest) (*proto.OemAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsOemAppDocDir error : missing updateField")
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
		logger.Error("UpdateFieldsOemAppDocDir error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OemAppDocDir_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsOemAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找OemAppDocDir
func (s *OemAppDocDirSvc) FindOemAppDocDir(req *proto.OemAppDocDirFilter) (*proto.OemAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.DocId != 0 { //整数
		do = do.Where(t.DocId.Eq(req.DocId))
	}
	if req.ParentId != 0 { //整数
		do = do.Where(t.ParentId.Eq(req.ParentId))
	}
	if req.DirName != "" { //字符串
		do = do.Where(t.DirName.Eq(req.DirName))
	}
	if req.DirImg != "" { //字符串
		do = do.Where(t.DirImg.Eq(req.DirImg))
	}
	if req.Sort != 0 { //整数
		do = do.Where(t.Sort.Eq(req.Sort))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindOemAppDocDir error : %s", err.Error())
		return nil, err
	}
	res := convert.OemAppDocDir_db2pb(dbObj)
	return res, err
}

// 返回目录id的下级目录id(包含目录id自己)
func (s *OemAppDocDirSvc) FindOemAppDocDirByIds(id int64) ([]int64, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
	do := t.WithContext(context.Background())
	do = do.Where(t.ParentId.Eq(id))
	dbObj, err := do.Find()
	var rs = make([]int64, 0)
	rs = append(rs, id)
	if err != nil {
		return rs, err
	}
	if dbObj != nil && len(dbObj) > 0 {
		for _, v := range dbObj {
			rs = append(rs, v.Id)
		}
	}
	return rs, nil
}

// 根据数据库表主键查找OemAppDocDir
func (s *OemAppDocDirSvc) FindByIdOemAppDocDir(req *proto.OemAppDocDirFilter) (*proto.OemAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdOemAppDocDir error : %s", err.Error())
		return nil, err
	}
	res := convert.OemAppDocDir_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找OemAppDocDir,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *OemAppDocDirSvc) GetListOemAppDocDir(req *proto.OemAppDocDirListRequest) ([]*proto.OemAppDocDir, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.DocId != 0 { //整数
			do = do.Where(t.DocId.Eq(query.DocId))
		}
		if query.ParentId > 0 { //整数
			do = do.Where(t.ParentId.Eq(query.ParentId))
		}
		if query.ParentId < 0 { //整数
			do = do.Where(t.ParentId.Eq(0))
		}
		if query.DirName != "" { //字符串
			do = do.Where(t.DirName.Like("%" + query.DirName + "%"))
		}
		if query.DirImg != "" { //字符串
			do = do.Where(t.DirImg.Like("%" + query.DirImg + "%"))
		}
		if query.Sort != 0 { //整数
			do = do.Where(t.Sort.Eq(query.Sort))
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

	var list []*model.TOemAppDocDir
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
		logger.Errorf("GetListOemAppDocDir error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.OemAppDocDir, len(list))
	for i, v := range list {
		result[i] = convert.OemAppDocDir_db2pb(v)
	}
	return result, total, nil
}

// 创建CreateOemAppDocDirBatch
func (s *OemAppDocDirSvc) CreateOemAppDocDirBatch(req *proto.OemAppDocDirBatchRequest) (*proto.OemAppDocDirBatchRequest, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TOemAppDocDir
	do := t.WithContext(context.Background())
	reqBatch := make([]*model.TOemAppDocDir, 0)
	for i := range req.OemAppDocDirs {
		reqBatch = append(reqBatch, convert.OemAppDocDir_pb2db(req.OemAppDocDirs[i]))
	}
	err := do.CreateInBatches(reqBatch, len(reqBatch))
	if err != nil {
		logger.Errorf("CreateOemAppDocDirBatch error : %s", err.Error())
		return nil, err
	}
	return req, err
}
