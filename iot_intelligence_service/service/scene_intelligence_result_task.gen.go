// Code generated by sgen.exe,2022-05-20 13:36:04. DO NOT EDIT.
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

type SceneIntelligenceResultTaskSvc struct {
	Ctx context.Context
}

// 创建SceneIntelligenceResultTask
func (s *SceneIntelligenceResultTaskSvc) CreateSceneIntelligenceResultTask(req *proto.SceneIntelligenceResultTask) (*proto.SceneIntelligenceResultTask, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceResultTask
	do := t.WithContext(context.Background())
	dbObj := convert.SceneIntelligenceResultTask_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateSceneIntelligenceResultTask error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除SceneIntelligenceResultTask
func (s *SceneIntelligenceResultTaskSvc) DeleteSceneIntelligenceResultTask(req *proto.SceneIntelligenceResultTask) (*proto.SceneIntelligenceResultTask, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceResultTask
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.IntelligenceId != 0 { //整数
		do = do.Where(t.IntelligenceId.Eq(req.IntelligenceId))
	}
	if req.ResultId != 0 { //整数
		do = do.Where(t.ResultId.Eq(req.ResultId))
	}
	if req.IsSuccess != 0 { //整数
		do = do.Where(t.IsSuccess.Eq(req.IsSuccess))
	}
	if req.ResultMsg != "" { //字符串
		do = do.Where(t.ResultMsg.Eq(req.ResultMsg))
	}
	if req.TaskId != 0 { //整数
		do = do.Where(t.TaskId.Eq(req.TaskId))
	}
	if req.TaskImg != "" { //字符串
		do = do.Where(t.TaskImg.Eq(req.TaskImg))
	}
	if req.TaskDesc != "" { //字符串
		do = do.Where(t.TaskDesc.Eq(req.TaskDesc))
	}
	if req.TaskType != 0 { //整数
		do = do.Where(t.TaskType.Eq(req.TaskType))
	}
	if req.ObjectId != "" { //字符串
		do = do.Where(t.ObjectId.Eq(req.ObjectId))
	}
	if req.ObjectDesc != "" { //字符串
		do = do.Where(t.ObjectDesc.Eq(req.ObjectDesc))
	}
	if req.FuncKey != "" { //字符串
		do = do.Where(t.FuncKey.Eq(req.FuncKey))
	}
	if req.FuncDesc != "" { //字符串
		do = do.Where(t.FuncDesc.Eq(req.FuncDesc))
	}
	if req.FuncValue != "" { //字符串
		do = do.Where(t.FuncValue.Eq(req.FuncValue))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteSceneIntelligenceResultTask error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除SceneIntelligenceResultTask
func (s *SceneIntelligenceResultTaskSvc) DeleteByIdSceneIntelligenceResultTask(req *proto.SceneIntelligenceResultTask) (*proto.SceneIntelligenceResultTask, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceResultTask
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdSceneIntelligenceResultTask error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除SceneIntelligenceResultTask
func (s *SceneIntelligenceResultTaskSvc) DeleteByIdsSceneIntelligenceResultTask(req *proto.SceneIntelligenceResultTaskBatchDeleteRequest) (*proto.SceneIntelligenceResultTaskBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceResultTask
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsSceneIntelligenceResultTask error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新SceneIntelligenceResultTask
func (s *SceneIntelligenceResultTaskSvc) UpdateSceneIntelligenceResultTask(req *proto.SceneIntelligenceResultTask) (*proto.SceneIntelligenceResultTask, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceResultTask
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.IntelligenceId != 0 { //整数
		updateField = append(updateField, t.IntelligenceId)
	}
	if req.ResultId != 0 { //整数
		updateField = append(updateField, t.ResultId)
	}
	if req.IsSuccess != 0 { //整数
		updateField = append(updateField, t.IsSuccess)
	}
	if req.ResultMsg != "" { //字符串
		updateField = append(updateField, t.ResultMsg)
	}
	if req.TaskId != 0 { //整数
		updateField = append(updateField, t.TaskId)
	}
	if req.TaskImg != "" { //字符串
		updateField = append(updateField, t.TaskImg)
	}
	if req.TaskDesc != "" { //字符串
		updateField = append(updateField, t.TaskDesc)
	}
	if req.TaskType != 0 { //整数
		updateField = append(updateField, t.TaskType)
	}
	if req.ObjectId != "" { //字符串
		updateField = append(updateField, t.ObjectId)
	}
	if req.ObjectDesc != "" { //字符串
		updateField = append(updateField, t.ObjectDesc)
	}
	if req.FuncKey != "" { //字符串
		updateField = append(updateField, t.FuncKey)
	}
	if req.FuncDesc != "" { //字符串
		updateField = append(updateField, t.FuncDesc)
	}
	if req.FuncValue != "" { //字符串
		updateField = append(updateField, t.FuncValue)
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
		logger.Error("UpdateSceneIntelligenceResultTask error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.SceneIntelligenceResultTask_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateSceneIntelligenceResultTask error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段SceneIntelligenceResultTask
func (s *SceneIntelligenceResultTaskSvc) UpdateAllSceneIntelligenceResultTask(req *proto.SceneIntelligenceResultTask) (*proto.SceneIntelligenceResultTask, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceResultTask
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.StartTime)
	updateField = append(updateField, t.EndTime)
	updateField = append(updateField, t.IntelligenceId)
	updateField = append(updateField, t.ResultId)
	updateField = append(updateField, t.IsSuccess)
	updateField = append(updateField, t.ResultMsg)
	updateField = append(updateField, t.TaskId)
	updateField = append(updateField, t.TaskImg)
	updateField = append(updateField, t.TaskDesc)
	updateField = append(updateField, t.TaskType)
	updateField = append(updateField, t.ObjectId)
	updateField = append(updateField, t.ObjectDesc)
	updateField = append(updateField, t.FuncKey)
	updateField = append(updateField, t.FuncDesc)
	updateField = append(updateField, t.FuncValue)
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
		logger.Error("UpdateAllSceneIntelligenceResultTask error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SceneIntelligenceResultTask_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllSceneIntelligenceResultTask error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *SceneIntelligenceResultTaskSvc) UpdateFieldsSceneIntelligenceResultTask(req *proto.SceneIntelligenceResultTaskUpdateFieldsRequest) (*proto.SceneIntelligenceResultTask, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceResultTask
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsSceneIntelligenceResultTask error : missing updateField")
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
		logger.Error("UpdateFieldsSceneIntelligenceResultTask error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SceneIntelligenceResultTask_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsSceneIntelligenceResultTask error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找SceneIntelligenceResultTask
func (s *SceneIntelligenceResultTaskSvc) FindSceneIntelligenceResultTask(req *proto.SceneIntelligenceResultTaskFilter) (*proto.SceneIntelligenceResultTask, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceResultTask
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.IntelligenceId != 0 { //整数
		do = do.Where(t.IntelligenceId.Eq(req.IntelligenceId))
	}
	if req.ResultId != 0 { //整数
		do = do.Where(t.ResultId.Eq(req.ResultId))
	}
	if req.IsSuccess != 0 { //整数
		do = do.Where(t.IsSuccess.Eq(req.IsSuccess))
	}
	if req.ResultMsg != "" { //字符串
		do = do.Where(t.ResultMsg.Eq(req.ResultMsg))
	}
	if req.TaskId != 0 { //整数
		do = do.Where(t.TaskId.Eq(req.TaskId))
	}
	if req.TaskImg != "" { //字符串
		do = do.Where(t.TaskImg.Eq(req.TaskImg))
	}
	if req.TaskDesc != "" { //字符串
		do = do.Where(t.TaskDesc.Eq(req.TaskDesc))
	}
	if req.TaskType != 0 { //整数
		do = do.Where(t.TaskType.Eq(req.TaskType))
	}
	if req.ObjectId != "" { //字符串
		do = do.Where(t.ObjectId.Eq(req.ObjectId))
	}
	if req.ObjectDesc != "" { //字符串
		do = do.Where(t.ObjectDesc.Eq(req.ObjectDesc))
	}
	if req.FuncKey != "" { //字符串
		do = do.Where(t.FuncKey.Eq(req.FuncKey))
	}
	if req.FuncDesc != "" { //字符串
		do = do.Where(t.FuncDesc.Eq(req.FuncDesc))
	}
	if req.FuncValue != "" { //字符串
		do = do.Where(t.FuncValue.Eq(req.FuncValue))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindSceneIntelligenceResultTask error : %s", err.Error())
		return nil, err
	}
	res := convert.SceneIntelligenceResultTask_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找SceneIntelligenceResultTask
func (s *SceneIntelligenceResultTaskSvc) FindByIdSceneIntelligenceResultTask(req *proto.SceneIntelligenceResultTaskFilter) (*proto.SceneIntelligenceResultTask, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceResultTask
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdSceneIntelligenceResultTask error : %s", err.Error())
		return nil, err
	}
	res := convert.SceneIntelligenceResultTask_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找SceneIntelligenceResultTask,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *SceneIntelligenceResultTaskSvc) GetListSceneIntelligenceResultTask(req *proto.SceneIntelligenceResultTaskListRequest) ([]*proto.SceneIntelligenceResultTask, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TSceneIntelligenceResultTask
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.IntelligenceId != 0 { //整数
			do = do.Where(t.IntelligenceId.Eq(query.IntelligenceId))
		}
		if query.ResultId != 0 { //整数
			do = do.Where(t.ResultId.Eq(query.ResultId))
		}
		if query.IsSuccess != 0 { //整数
			do = do.Where(t.IsSuccess.Eq(query.IsSuccess))
		}
		if query.ResultMsg != "" { //字符串
			do = do.Where(t.ResultMsg.Like("%" + query.ResultMsg + "%"))
		}
		if query.TaskId != 0 { //整数
			do = do.Where(t.TaskId.Eq(query.TaskId))
		}
		if query.TaskImg != "" { //字符串
			do = do.Where(t.TaskImg.Like("%" + query.TaskImg + "%"))
		}
		if query.TaskDesc != "" { //字符串
			do = do.Where(t.TaskDesc.Like("%" + query.TaskDesc + "%"))
		}
		if query.TaskType != 0 { //整数
			do = do.Where(t.TaskType.Eq(query.TaskType))
		}
		if query.ObjectId != "" { //字符串
			do = do.Where(t.ObjectId.Like("%" + query.ObjectId + "%"))
		}
		if query.ObjectDesc != "" { //字符串
			do = do.Where(t.ObjectDesc.Like("%" + query.ObjectDesc + "%"))
		}
		if query.FuncKey != "" { //字符串
			do = do.Where(t.FuncKey.Like("%" + query.FuncKey + "%"))
		}
		if query.FuncDesc != "" { //字符串
			do = do.Where(t.FuncDesc.Like("%" + query.FuncDesc + "%"))
		}
		if query.FuncValue != "" { //字符串
			do = do.Where(t.FuncValue.Like("%" + query.FuncValue + "%"))
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

	var list []*model.TSceneIntelligenceResultTask
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
		logger.Errorf("GetListSceneIntelligenceResultTask error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.SceneIntelligenceResultTask, len(list))
	for i, v := range list {
		result[i] = convert.SceneIntelligenceResultTask_db2pb(v)
	}
	return result, total, nil
}
