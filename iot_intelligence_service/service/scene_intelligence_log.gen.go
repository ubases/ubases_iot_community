// Code generated by sgen.exe,2022-05-20 13:36:03. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	"cloud_platform/iot_intelligence_service/convert"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_model/db_device/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type SceneIntelligenceLogSvc struct {
	Ctx context.Context
}

// 创建SceneIntelligenceLog
func (s *SceneIntelligenceLogSvc) CreateSceneIntelligenceLog(req *proto.SceneIntelligenceLog) (*proto.SceneIntelligenceLog, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceLog
	do := t.WithContext(context.Background())
	dbObj := convert.SceneIntelligenceLog_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateSceneIntelligenceLog error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除SceneIntelligenceLog
func (s *SceneIntelligenceLogSvc) DeleteSceneIntelligenceLog(req *proto.SceneIntelligenceLog) (*proto.SceneIntelligenceLog, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceLog
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.ObjectId != 0 { //整数
		do = do.Where(t.ObjectId.Eq(req.ObjectId))
	}
	if req.HomeId != 0 { //整数
		do = do.Where(t.HomeId.Eq(req.HomeId))
	}
	if req.UserId != 0 { //整数
		do = do.Where(t.UserId.Eq(req.UserId))
	}
	if req.Content != "" { //字符串
		do = do.Where(t.Content.Eq(req.Content))
	}
	if req.ResultId != 0 { //整数
		do = do.Where(t.ResultId.Eq(req.ResultId))
	}
	if req.IsSuccess != 0 { //整数
		do = do.Where(t.IsSuccess.Eq(req.IsSuccess))
	}
	if req.SceneType != 0 { //整数
		do = do.Where(t.SceneType.Eq(req.SceneType))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteSceneIntelligenceLog error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除SceneIntelligenceLog
func (s *SceneIntelligenceLogSvc) DeleteByIdSceneIntelligenceLog(req *proto.SceneIntelligenceLog) (*proto.SceneIntelligenceLog, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceLog
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdSceneIntelligenceLog error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除SceneIntelligenceLog
func (s *SceneIntelligenceLogSvc) DeleteByIdsSceneIntelligenceLog(req *proto.SceneIntelligenceLogBatchDeleteRequest) (*proto.SceneIntelligenceLogBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceLog
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsSceneIntelligenceLog error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新SceneIntelligenceLog
func (s *SceneIntelligenceLogSvc) UpdateSceneIntelligenceLog(req *proto.SceneIntelligenceLog) (*proto.SceneIntelligenceLog, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceLog
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.ObjectId != 0 { //整数
		updateField = append(updateField, t.ObjectId)
	}
	if req.HomeId != 0 { //整数
		updateField = append(updateField, t.HomeId)
	}
	if req.UserId != 0 { //整数
		updateField = append(updateField, t.UserId)
	}
	if req.Content != "" { //字符串
		updateField = append(updateField, t.Content)
	}
	if req.ResultId != 0 { //整数
		updateField = append(updateField, t.ResultId)
	}
	if req.IsSuccess != 0 { //整数
		updateField = append(updateField, t.IsSuccess)
	}
	if req.SceneType != 0 { //整数
		updateField = append(updateField, t.SceneType)
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
		logger.Error("UpdateSceneIntelligenceLog error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.SceneIntelligenceLog_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateSceneIntelligenceLog error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段SceneIntelligenceLog
func (s *SceneIntelligenceLogSvc) UpdateAllSceneIntelligenceLog(req *proto.SceneIntelligenceLog) (*proto.SceneIntelligenceLog, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceLog
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.ObjectId)
	updateField = append(updateField, t.HomeId)
	updateField = append(updateField, t.UserId)
	updateField = append(updateField, t.Content)
	updateField = append(updateField, t.ResultId)
	updateField = append(updateField, t.IsSuccess)
	updateField = append(updateField, t.SceneType)
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
		logger.Error("UpdateAllSceneIntelligenceLog error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SceneIntelligenceLog_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllSceneIntelligenceLog error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *SceneIntelligenceLogSvc) UpdateFieldsSceneIntelligenceLog(req *proto.SceneIntelligenceLogUpdateFieldsRequest) (*proto.SceneIntelligenceLog, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceLog
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsSceneIntelligenceLog error : missing updateField")
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
		logger.Error("UpdateFieldsSceneIntelligenceLog error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SceneIntelligenceLog_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsSceneIntelligenceLog error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找SceneIntelligenceLog
func (s *SceneIntelligenceLogSvc) FindSceneIntelligenceLog(req *proto.SceneIntelligenceLogFilter) (*proto.SceneIntelligenceLog, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceLog
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.ObjectId != 0 { //整数
		do = do.Where(t.ObjectId.Eq(req.ObjectId))
	}
	if req.HomeId != 0 { //整数
		do = do.Where(t.HomeId.Eq(req.HomeId))
	}
	if req.UserId != 0 { //整数
		do = do.Where(t.UserId.Eq(req.UserId))
	}
	if req.Content != "" { //字符串
		do = do.Where(t.Content.Eq(req.Content))
	}
	if req.ResultId != 0 { //整数
		do = do.Where(t.ResultId.Eq(req.ResultId))
	}
	if req.IsSuccess != 0 { //整数
		do = do.Where(t.IsSuccess.Eq(req.IsSuccess))
	}
	if req.SceneType != 0 { //整数
		do = do.Where(t.SceneType.Eq(req.SceneType))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindSceneIntelligenceLog error : %s", err.Error())
		return nil, err
	}
	res := convert.SceneIntelligenceLog_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找SceneIntelligenceLog
func (s *SceneIntelligenceLogSvc) FindByIdSceneIntelligenceLog(req *proto.SceneIntelligenceLogFilter) (*proto.SceneIntelligenceLog, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceLog
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdSceneIntelligenceLog error : %s", err.Error())
		return nil, err
	}
	res := convert.SceneIntelligenceLog_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找SceneIntelligenceLog,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *SceneIntelligenceLogSvc) GetListSceneIntelligenceLog(req *proto.SceneIntelligenceLogListRequest) ([]*proto.SceneIntelligenceLog, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	q := orm.Use(iotmodel.GetDB())
	t := q.TSceneIntelligenceLog
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {
		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.ObjectId != 0 { //整数
			do = do.Where(t.ObjectId.Eq(query.ObjectId))
		}
		if query.HomeId != 0 { //整数
			do = do.Where(t.HomeId.Eq(query.HomeId))
		}
		if query.UserId != 0 { //整数
			do = do.Where(t.UserId.Eq(query.UserId))
		}
		if query.Content != "" { //字符串
			do = do.Where(t.Content.Like("%" + query.Content + "%"))
		}
		if query.ResultId != 0 { //整数
			do = do.Where(t.ResultId.Eq(query.ResultId))
		}
		if query.IsSuccess != 0 { //整数
			do = do.Where(t.IsSuccess.Eq(query.IsSuccess))
		}
		if query.SceneType != 0 { //整数
			do = do.Where(t.SceneType.Eq(query.SceneType))
		}
	}
	orderCol, ok := t.GetFieldByName(req.OrderKey)
	if !ok {
		do = do.Order(t.CreatedAt.Desc())
	} else {
		if req.OrderDesc != "" {
			do = do.Order(orderCol.Desc())
		} else {
			do = do.Order(orderCol)
		}
	}
	var list []*model.TSceneIntelligenceLog
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
		logger.Errorf("GetListSceneIntelligenceLog error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.SceneIntelligenceLog, len(list))
	for i, v := range list {
		result[i] = convert.SceneIntelligenceLog_db2pb(v)
	}
	return result, total, nil
}
