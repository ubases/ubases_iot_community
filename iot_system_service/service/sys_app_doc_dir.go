// Code generated by sgen.exe,2022-11-15 11:05:52. DO NOT EDIT.
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

type SysAppDocDirSvc struct {
	Ctx context.Context
}

// 创建SysAppDocDir
func (s *SysAppDocDirSvc) CreateSysAppDocDir(req *proto.SysAppDocDir) (*proto.SysAppDocDir, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TSysAppDocDir
	do := t.WithContext(context.Background())
	dbObj := convert.SysAppDocDir_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateSysAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除SysAppDocDir
func (s *SysAppDocDirSvc) DeleteSysAppDocDir(req *proto.SysAppDocDir) (*proto.SysAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TSysAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
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
	if req.HelpId != 0 { //整数
		do = do.Where(t.HelpId.Eq(req.HelpId))
	}
	if req.Sort != 0 { //整数
		do = do.Where(t.Sort.Eq(req.Sort))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteSysAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除SysAppDocDir
func (s *SysAppDocDirSvc) DeleteByIdSysAppDocDir(req *proto.SysAppDocDir) (*proto.SysAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TSysAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdSysAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除SysAppDocDir
func (s *SysAppDocDirSvc) DeleteByIdsSysAppDocDir(req *proto.SysAppDocDirBatchDeleteRequest) (*proto.SysAppDocDirBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TSysAppDocDir
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsSysAppDocDir error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新SysAppDocDir
func (s *SysAppDocDirSvc) UpdateSysAppDocDir(req *proto.SysAppDocDir) (*proto.SysAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TSysAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.ParentId != 0 { //整数
		updateField = append(updateField, t.ParentId)
	}
	if req.DirName != "" { //字符串
		updateField = append(updateField, t.DirName)
	}
	if req.DirImg != "" { //字符串
		updateField = append(updateField, t.DirImg)
	}
	if req.HelpId != 0 { //整数
		updateField = append(updateField, t.HelpId)
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
		logger.Error("UpdateSysAppDocDir error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.SysAppDocDir_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateSysAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段SysAppDocDir
func (s *SysAppDocDirSvc) UpdateAllSysAppDocDir(req *proto.SysAppDocDir) (*proto.SysAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TSysAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.ParentId)
	updateField = append(updateField, t.DirName)
	updateField = append(updateField, t.DirImg)
	updateField = append(updateField, t.HelpId)
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
		logger.Error("UpdateAllSysAppDocDir error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SysAppDocDir_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllSysAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *SysAppDocDirSvc) UpdateFieldsSysAppDocDir(req *proto.SysAppDocDirUpdateFieldsRequest) (*proto.SysAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TSysAppDocDir
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsSysAppDocDir error : missing updateField")
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
		logger.Error("UpdateFieldsSysAppDocDir error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SysAppDocDir_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsSysAppDocDir error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找SysAppDocDir
func (s *SysAppDocDirSvc) FindSysAppDocDir(req *proto.SysAppDocDirFilter) (*proto.SysAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TSysAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
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
	if req.HelpId != 0 { //整数
		do = do.Where(t.HelpId.Eq(req.HelpId))
	}
	if req.Sort != 0 { //整数
		do = do.Where(t.Sort.Eq(req.Sort))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindSysAppDocDir error : %s", err.Error())
		return nil, err
	}
	res := convert.SysAppDocDir_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找SysAppDocDir
func (s *SysAppDocDirSvc) FindByIdSysAppDocDir(req *proto.SysAppDocDirFilter) (*proto.SysAppDocDir, error) {
	t := orm.Use(iotmodel.GetDB()).TSysAppDocDir
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdSysAppDocDir error : %s", err.Error())
		return nil, err
	}
	res := convert.SysAppDocDir_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找SysAppDocDir,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *SysAppDocDirSvc) GetListSysAppDocDir(req *proto.SysAppDocDirListRequest) ([]*proto.SysAppDocDir, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	q := orm.Use(iotmodel.GetDB())
	t := q.TSysAppDocDir
	tHelp := q.TSysAppHelpCenter
	do := t.WithContext(context.Background()).Join(tHelp, tHelp.Id.EqCol(t.HelpId))
	query := req.Query
	if query != nil {
		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
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
		if query.HelpId != 0 { //整数
			do = do.Where(t.HelpId.Eq(query.HelpId))
		}
		//关联APP模板进行查询
		if query.AppTemplateType != 0 {
			do = do.Where(tHelp.TemplateType.Eq(query.AppTemplateType))
		}
		if query.AppTemplateVersion != "" {
			do = do.Where(tHelp.Version.Eq(query.AppTemplateVersion))
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

	var list []*model.TSysAppDocDir
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
		logger.Errorf("GetListSysAppDocDir error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.SysAppDocDir, len(list))
	for i, v := range list {
		result[i] = convert.SysAppDocDir_db2pb(v)
	}
	return result, total, nil
}

// 返回目录id的下级目录id(包含目录id自己)
func (s *SysAppDocDirSvc) FindOemAppDocDirByIds(id int64) ([]int64, error) {
	t := orm.Use(iotmodel.GetDB()).TSysAppDocDir
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

// 创建CreateSysAppDocDirBatch
func (s *SysAppDocDirSvc) CreateSysAppDocDirBatch(req *proto.SysAppDocDirBatchRequest) (*proto.SysAppDocDirBatchRequest, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TSysAppDocDir
	do := t.WithContext(context.Background())
	reqBatch := make([]*model.TSysAppDocDir, 0)
	for i := range req.SysAppDocDirs {
		reqBatch = append(reqBatch, convert.SysAppDocDir_pb2db(req.SysAppDocDirs[i]))
	}
	err := do.CreateInBatches(reqBatch, len(reqBatch))
	if err != nil {
		logger.Errorf("CreateSysAppDocDirBatch error : %s", err.Error())
		return nil, err
	}
	return req, err
}
