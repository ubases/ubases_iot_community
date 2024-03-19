// Code generated by sgen.exe,2022-11-11 13:37:01. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"cloud_platform/iot_common/iotutil"
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

type SceneTemplateSvc struct {
	Ctx context.Context
}

// 创建SceneTemplate
func (s *SceneTemplateSvc) CreateSceneTemplate(req *proto.SceneTemplate) (*proto.SceneTemplate, error) {
	// fixme 请在这里校验参数
	var err error
	tenantId, err := CheckTenantId(s.Ctx)
	if err != nil {
		return nil, err
	}
	_, userId, err := CheckUserId(s.Ctx)
	if err != nil {
		return nil, err
	}
	q := orm.Use(iotmodel.GetDB())
	err = q.Transaction(func(tx *orm.Query) error {
		//场景模板新增
		do := tx.WithContext(context.Background()).TSceneTemplate
		dbObj := convert.SceneTemplate_pb2db(req)
		dbObj.Id = iotutil.GetNextSeqInt64()
		dbObj.CreatedBy = userId
		dbObj.TenantId = tenantId
		dbObj.Status = 2
		err := do.Create(dbObj)
		if err != nil {
			return err
		}
		//task
		err = s.createTemplateTask(tx, dbObj.Id, userId, req)
		if err != nil {
			return err
		}
		//conditions
		err = s.createTemplateCondition(tx, dbObj.Id, userId, req)
		if err != nil {
			return err
		}
		//appList
		err = s.createTemplateAppList(tx, dbObj.Id, userId, tenantId, req)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Errorf("CreateSceneTemplate error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 创建模板任务
func (s *SceneTemplateSvc) createTemplateTask(tx *orm.Query, id, userId int64, req *proto.SceneTemplate) error {
	//task
	dbTasksObj := make([]*model.TSceneTemplateTask, 0)
	for i, task := range req.Tasks {
		dbTasksObj = append(dbTasksObj, &model.TSceneTemplateTask{
			Id:              iotutil.GetNextSeqInt64(),
			SceneTemplateId: id,
			ProductId:       task.ProductId,
			ProductKey:      task.ProductKey,
			Functions:       task.Functions,
			Sort:            int32(i + 1),
			CreatedBy:       userId,
		})
	}
	doTask := tx.TSceneTemplateTask.WithContext(context.Background())
	err := doTask.Create(dbTasksObj...)
	if err != nil {
		return err
	}
	return err
}

// 创建模板条件
func (s *SceneTemplateSvc) createTemplateCondition(tx *orm.Query, id, userId int64, req *proto.SceneTemplate) error {
	dbConditionsObj := make([]*model.TSceneTemplateCondition, 0)
	for i, condition := range req.Conditions {
		dbConditionsObj = append(dbConditionsObj, &model.TSceneTemplateCondition{
			Id:              iotutil.GetNextSeqInt64(),
			SceneTemplateId: id,
			ConditionType:   condition.ConditionType,
			ProductId:       condition.ProductId,
			ProductKey:      condition.ProductKey,
			Functions:       condition.Functions,
			Sort:            int32(i + 1),
			CreatedBy:       userId,
			Desc:            condition.Desc,
			WeatherType:     condition.WeatherType,
			WeatherValue:    condition.WeatherValue,
			WeatherCompare:  condition.WeatherCompare,
			TimerWeeks:      condition.TimerWeeks,
			TimerValue:      condition.TimerValue,
			FuncKey:         condition.FuncKey,
			FuncCompare:     condition.FuncCompare,
			FuncValue:       condition.FuncValue,
			FuncDesc:        condition.FuncDesc,
			FuncIdentifier:  condition.FuncIdentifier,
		})
	}
	doCondition := tx.TSceneTemplateCondition.WithContext(context.Background())
	err := doCondition.Create(dbConditionsObj...)
	if err != nil {
		return err
	}
	return err
}

// 创建模板应用列表
func (s *SceneTemplateSvc) createTemplateAppList(tx *orm.Query, id, userId int64, tenantId string, req *proto.SceneTemplate) error {
	dbApps := make([]*model.TSceneTemplateAppRelation, 0)
	for _, app := range req.AppList {
		dbApps = append(dbApps, &model.TSceneTemplateAppRelation{
			Id:              iotutil.GetNextSeqInt64(),
			SceneTemplateId: id,
			AppId:           app.AppId,
			AppKey:          app.AppKey,
			AppName:         app.AppName,
			TenantId:        tenantId,
			CreatedBy:       userId,
		})
	}
	doApps := tx.TSceneTemplateAppRelation.WithContext(context.Background())
	err := doApps.Create(dbApps...)
	if err != nil {
		return err
	}
	return err
}

// 根据主键更新SceneTemplate
func (s *SceneTemplateSvc) UpdateSceneTemplate(req *proto.SceneTemplate) (*proto.SceneTemplate, error) {
	// fixme 请在这里校验参数
	var err error
	tenantId, err := CheckTenantId(s.Ctx)
	if err != nil {
		return nil, err
	}
	_, userId, err := CheckUserId(s.Ctx)
	if err != nil {
		return nil, err
	}
	q := orm.Use(iotmodel.GetDB())
	err = q.Transaction(func(tx *orm.Query) error {
		//场景模板修改
		do := tx.WithContext(context.Background()).TSceneTemplate
		dbObj := convert.SceneTemplate_pb2db(req)
		dbObj.UpdatedBy = userId
		dbObj.TenantId = tenantId
		_, err := do.Where(tx.TSceneTemplate.Id.Eq(req.Id)).Updates(dbObj)
		if err != nil {
			return err
		}
		//task
		err = s.updateTemplateTask(tx, dbObj.Id, userId, req)
		if err != nil {
			return err
		}
		//conditions
		err = s.updateTemplateCondition(tx, dbObj.Id, userId, req)
		if err != nil {
			return err
		}
		//appList
		err = s.updateTemplateAppList(tx, dbObj.Id, userId, tenantId, req)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Errorf("CreateSceneTemplate error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 修改模板任务
func (s *SceneTemplateSvc) updateTemplateTask(tx *orm.Query, id, userId int64, req *proto.SceneTemplate) error {
	//task
	dbTasksObj := make([]*model.TSceneTemplateTask, 0)
	for i, task := range req.Tasks {
		dbTask := &model.TSceneTemplateTask{
			Id:              task.Id,
			SceneTemplateId: id,
			ProductId:       task.ProductId,
			ProductKey:      task.ProductKey,
			Functions:       task.Functions,
			Sort:            int32(i + 1),
			CreatedBy:       userId,
		}
		if task.Id == 0 {
			dbTask.Id = iotutil.GetNextSeqInt64()
		}
		dbTasksObj = append(dbTasksObj, dbTask)
	}
	doTask := tx.TSceneTemplateTask.WithContext(context.Background())
	_, err := doTask.Omit(tx.TSceneTemplateTask.Id).Where(tx.TSceneTemplateTask.SceneTemplateId.Eq(id)).Delete()
	if err != nil {
		return err
	}
	err = doTask.Save(dbTasksObj...)
	if err != nil {
		return err
	}
	return err
}

// 修改模板条件
func (s *SceneTemplateSvc) updateTemplateCondition(tx *orm.Query, id, userId int64, req *proto.SceneTemplate) error {
	dbConditionsObj := make([]*model.TSceneTemplateCondition, 0)
	for i, condition := range req.Conditions {
		dbCondition := &model.TSceneTemplateCondition{
			Id:              condition.Id,
			SceneTemplateId: id,
			ConditionType:   condition.ConditionType,
			ProductId:       condition.ProductId,
			ProductKey:      condition.ProductKey,
			Functions:       condition.Functions,
			Sort:            int32(i + 1),
			CreatedBy:       userId,
			Desc:            condition.Desc,
			WeatherType:     condition.WeatherType,
			WeatherValue:    condition.WeatherValue,
			WeatherCompare:  condition.WeatherCompare,
			TimerWeeks:      condition.TimerWeeks,
			TimerValue:      condition.TimerValue,
			FuncKey:         condition.FuncKey,
			FuncCompare:     condition.FuncCompare,
			FuncValue:       condition.FuncValue,
			FuncDesc:        condition.FuncDesc,
			FuncIdentifier:  condition.FuncIdentifier,
		}
		if dbCondition.Id == 0 {
			dbCondition.Id = iotutil.GetNextSeqInt64()
		}
		dbConditionsObj = append(dbConditionsObj, dbCondition)
	}
	doCondition := tx.TSceneTemplateCondition.WithContext(context.Background())
	_, err := doCondition.Omit(tx.TSceneTemplateCondition.Id).Where(tx.TSceneTemplateCondition.SceneTemplateId.Eq(id)).Delete()
	if err != nil {
		return err
	}
	err = doCondition.Create(dbConditionsObj...)
	if err != nil {
		return err
	}
	return err
}

// 创模板应用列表
func (s *SceneTemplateSvc) updateTemplateAppList(tx *orm.Query, id, userId int64, tenantId string, req *proto.SceneTemplate) error {
	dbApps := make([]*model.TSceneTemplateAppRelation, 0)
	for _, app := range req.AppList {
		dbApp := &model.TSceneTemplateAppRelation{
			Id:              iotutil.GetNextSeqInt64(),
			SceneTemplateId: id,
			AppId:           app.AppId,
			AppKey:          app.AppKey,
			AppName:         app.AppName,
			TenantId:        tenantId,
			CreatedBy:       userId,
		}
		dbApps = append(dbApps, dbApp)
	}
	doApps := tx.TSceneTemplateAppRelation.WithContext(context.Background())
	_, err := doApps.Omit(tx.TSceneTemplateAppRelation.Id).Where(tx.TSceneTemplateAppRelation.SceneTemplateId.Eq(id)).Delete()
	if err != nil {
		return err
	}
	err = doApps.Create(dbApps...)
	if err != nil {
		return err
	}
	return err
}

// 根据条件删除SceneTemplate
func (s *SceneTemplateSvc) DeleteSceneTemplate(req *proto.SceneTemplate) (*proto.SceneTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneTemplate
	do := t.WithContext(context.Background())
	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.Type != 0 { //整数
		do = do.Where(t.Type.Eq(req.Type))
	}
	if req.Title != "" { //字符串
		do = do.Where(t.Title.Eq(req.Title))
	}
	if req.TitleEn != "" { //字符串
		do = do.Where(t.TitleEn.Eq(req.TitleEn))
	}
	if req.Desc != "" { //字符串
		do = do.Where(t.Desc.Eq(req.Desc))
	}
	if req.DescEn != "" { //字符串
		do = do.Where(t.DescEn.Eq(req.DescEn))
	}
	if req.SortNo != 0 { //整数
		do = do.Where(t.SortNo.Eq(req.SortNo))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.ConditionMode != 0 { //整数
		do = do.Where(t.ConditionMode.Eq(req.ConditionMode))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteSceneTemplate error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除SceneTemplate
func (s *SceneTemplateSvc) DeleteByIdSceneTemplate(req *proto.SceneTemplate) (*proto.SceneTemplate, error) {
	q := orm.Use(iotmodel.GetDB())
	err := q.Transaction(func(tx *orm.Query) error {
		t := tx.TSceneTemplate
		do := t.WithContext(context.Background())
		if req.Id == 0 {
			return errors.New("删除场景模板Id不能为空")
		}
		do = do.Where(t.Id.Eq(req.Id))
		_, err := do.Delete()
		if err != nil {
			return err
		}
		//task
		_, err = tx.TSceneTemplateTask.WithContext(context.Background()).Where(tx.TSceneTemplateTask.SceneTemplateId.Eq(req.Id)).Delete()
		if err != nil {
			return err
		}
		//condition
		_, err = tx.TSceneTemplateCondition.WithContext(context.Background()).Where(tx.TSceneTemplateCondition.SceneTemplateId.Eq(req.Id)).Delete()
		if err != nil {
			return err
		}
		//appList
		_, err = tx.TSceneTemplateAppRelation.WithContext(context.Background()).Where(tx.TSceneTemplateAppRelation.SceneTemplateId.Eq(req.Id)).Delete()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Errorf("DeleteByIdSceneTemplate error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除SceneTemplate
func (s *SceneTemplateSvc) DeleteByIdsSceneTemplate(req *proto.SceneTemplateBatchDeleteRequest) (*proto.SceneTemplateBatchDeleteRequest, error) {
	var err error
	deleteIds := []int64{}
	for _, k := range req.Keys {
		deleteIds = append(deleteIds, k.Id)
	}
	if len(deleteIds) == 0 {
		return nil, errors.New("请传入删除Id")
	}
	q := orm.Use(iotmodel.GetDB())
	err = q.Transaction(func(tx *orm.Query) error {
		t := tx.TSceneTemplate
		do := t.WithContext(context.Background())
		do = do.Where(t.Id.In(deleteIds...))
		_, err := do.Delete()
		if err != nil {
			return err
		}
		//task
		_, err = tx.TSceneTemplateTask.WithContext(context.Background()).Where(tx.TSceneTemplateTask.SceneTemplateId.In(deleteIds...)).Delete()
		if err != nil {
			return err
		}
		//condition
		_, err = tx.TSceneTemplateCondition.WithContext(context.Background()).Where(tx.TSceneTemplateCondition.SceneTemplateId.In(deleteIds...)).Delete()
		if err != nil {
			return err
		}
		//appList
		_, err = tx.TSceneTemplateAppRelation.WithContext(context.Background()).Where(tx.TSceneTemplateAppRelation.SceneTemplateId.In(deleteIds...)).Delete()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Errorf("DeleteByIdSceneTemplate error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段SceneTemplate
func (s *SceneTemplateSvc) UpdateAllSceneTemplate(req *proto.SceneTemplate) (*proto.SceneTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneTemplate
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.Type)
	updateField = append(updateField, t.Title)
	updateField = append(updateField, t.TitleEn)
	updateField = append(updateField, t.Desc)
	updateField = append(updateField, t.DescEn)
	updateField = append(updateField, t.SortNo)
	updateField = append(updateField, t.Status)
	updateField = append(updateField, t.ConditionMode)
	updateField = append(updateField, t.Icon)
	updateField = append(updateField, t.UseCount)
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
		logger.Error("UpdateAllSceneTemplate error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SceneTemplate_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllSceneTemplate error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *SceneTemplateSvc) UpdateFieldsSceneTemplate(req *proto.SceneTemplateUpdateFieldsRequest) (*proto.SceneTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneTemplate
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsSceneTemplate error : missing updateField")
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
		logger.Error("UpdateFieldsSceneTemplate error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.SceneTemplate_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsSceneTemplate error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找SceneTemplate
func (s *SceneTemplateSvc) FindSceneTemplate(req *proto.SceneTemplateFilter) (*proto.SceneTemplate, error) {
	t := orm.Use(iotmodel.GetDB()).TSceneTemplate
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.Type != 0 { //整数
		do = do.Where(t.Type.Eq(req.Type))
	}
	if req.Title != "" { //字符串
		do = do.Where(t.Title.Eq(req.Title))
	}
	if req.TitleEn != "" { //字符串
		do = do.Where(t.TitleEn.Eq(req.TitleEn))
	}
	if req.Desc != "" { //字符串
		do = do.Where(t.Desc.Eq(req.Desc))
	}
	if req.DescEn != "" { //字符串
		do = do.Where(t.DescEn.Eq(req.DescEn))
	}
	if req.SortNo != 0 { //整数
		do = do.Where(t.SortNo.Eq(req.SortNo))
	}
	if req.Status != 0 { //整数
		do = do.Where(t.Status.Eq(req.Status))
	}
	if req.ConditionMode != 0 { //整数
		do = do.Where(t.ConditionMode.Eq(req.ConditionMode))
	}
	if req.Icon != "" { //字符串
		do = do.Where(t.Icon.Eq(req.Icon))
	}
	if req.UseCount != 0 { //整数
		do = do.Where(t.UseCount.Eq(req.UseCount))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindSceneTemplate error : %s", err.Error())
		return nil, err
	}
	res := convert.SceneTemplate_db2pb(dbObj)
	taskMap, conditionMap, appMap := s.getSceneTemplateRelationDataMap(res.Id)
	res.Tasks = taskMap[res.Id]
	res.Conditions = conditionMap[res.Id]
	res.AppList = appMap[res.Id]
	return res, err
}

// 根据数据库表主键查找SceneTemplate
func (s *SceneTemplateSvc) FindByIdSceneTemplate(req *proto.SceneTemplateFilter) (*proto.SceneTemplate, error) {
	q := orm.Use(iotmodel.GetDB())
	t := q.TSceneTemplate
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdSceneTemplate error : %s", err.Error())
		return nil, err
	}
	res := convert.SceneTemplate_db2pb(dbObj)
	taskMap, conditionMap, appMap := s.getSceneTemplateRelationDataMap(res.Id)
	res.Tasks = taskMap[res.Id]
	res.Conditions = conditionMap[res.Id]
	res.AppList = appMap[res.Id]
	return res, err
}

// 根据分页条件查找SceneTemplate,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *SceneTemplateSvc) GetListSceneTemplate(req *proto.SceneTemplateListRequest) ([]*proto.SceneTemplate, int64, error) {
	q := orm.Use(iotmodel.GetDB())
	var err error
	t := q.TSceneTemplate
	tRelation := q.TSceneTemplateAppRelation
	doR := tRelation.WithContext(context.Background()).Group(tRelation.SceneTemplateId).
		Select(tRelation.SceneTemplateId,
			tRelation.AppKey.GroupConcat().As("app_key"),
			tRelation.AppName.GroupConcat().As("app_name"),
			tRelation.AppId.GroupConcat().As("app_id"))
	do := t.WithContext(context.Background()).LeftJoin(doR.As(tRelation.TableName()), tRelation.SceneTemplateId.EqCol(t.Id))
	query := req.Query
	if query != nil {
		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.Type != 0 { //整数
			do = do.Where(t.Type.Eq(query.Type))
		}
		if query.TenantId != "" {
			do = do.Where(t.TenantId.Eq(query.TenantId))
		}
		if query.Title != "" { //字符串
			do = do.Where(t.Title.Like("%" + query.Title + "%"))
		}
		if query.TitleEn != "" { //字符串
			do = do.Where(t.TitleEn.Like("%" + query.TitleEn + "%"))
		}
		if query.Desc != "" { //字符串
			do = do.Where(t.Desc.Like("%" + query.Desc + "%"))
		}
		if query.DescEn != "" { //字符串
			do = do.Where(t.DescEn.Like("%" + query.DescEn + "%"))
		}
		if query.SortNo != 0 { //整数
			do = do.Where(t.SortNo.Eq(query.SortNo))
		}
		if query.Status != 0 { //整数
			do = do.Where(t.Status.Eq(query.Status))
		}
		if query.ConditionMode != 0 { //整数
			do = do.Where(t.ConditionMode.Eq(query.ConditionMode))
		}
		if query.Icon != "" { //字符串
			do = do.Where(t.Icon.Like("%" + query.Icon + "%"))
		}
		if query.UseCount != 0 { //整数
			do = do.Where(t.UseCount.Eq(query.UseCount))
		}
		if query.AppList != nil && len(query.AppList) > 0 {
			appKey := []string{}
			for _, app := range query.AppList {
				if app.AppKey == "" {
					continue
				}
				appKey = append(appKey, app.AppKey)
			}
			if len(appKey) > 0 {
				orWhere := tRelation.WithContext(context.Background()).Where(tRelation.AppKey.Like("%" + appKey[0] + "%")).Or(t.IsSpecifyApp.Eq(1))
				do = do.Where(orWhere)
			}
		}
	}
	orderCol, ok := t.GetFieldByName(req.OrderKey)
	if !ok {
		do = do.Order(t.SortNo)
	} else {
		if req.OrderDesc != "" {
			do = do.Order(orderCol.Desc())
		} else {
			do = do.Order(orderCol)
		}
	}

	var list []*model.TSceneTemplate
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
		logger.Errorf("GetListSceneTemplate error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.SceneTemplate, len(list))
	templateIds := []int64{}
	for i, v := range list {
		result[i] = convert.SceneTemplate_db2pb(v)
		templateIds = append(templateIds, v.Id)
	}
	//task
	taskMap, conditionMap, appMap := s.getSceneTemplateRelationDataMap(templateIds...)
	for i, v := range result {
		result[i].Tasks = taskMap[v.Id]
		result[i].Conditions = conditionMap[v.Id]
		result[i].AppList = appMap[v.Id]
	}
	return result, total, nil
}

// 获取场景模板的其它关联数据
func (s SceneTemplateSvc) getSceneTemplateRelationDataMap(templateIds ...int64) (
	map[int64][]*proto.SceneTemplateTask,
	map[int64][]*proto.SceneTemplateCondition,
	map[int64][]*proto.SceneTemplateAppRelation,
) {
	q := orm.Use(iotmodel.GetDB())
	//task
	taskMap := map[int64][]*proto.SceneTemplateTask{}
	taskList, err := q.TSceneTemplateTask.WithContext(context.Background()).Where(q.TSceneTemplateTask.SceneTemplateId.In(templateIds...)).Find()
	if err == nil {
		for _, task := range taskList {
			taskMap[task.SceneTemplateId] = append(taskMap[task.SceneTemplateId], convert.SceneTemplateTask_db2pb(task))
		}
	}
	//condition
	conditionMap := map[int64][]*proto.SceneTemplateCondition{}
	conditionList, err := q.TSceneTemplateCondition.WithContext(context.Background()).Where(q.TSceneTemplateCondition.SceneTemplateId.In(templateIds...)).Find()
	if err == nil {
		for _, condition := range conditionList {
			conditionMap[condition.SceneTemplateId] = append(conditionMap[condition.SceneTemplateId], convert.SceneTemplateCondition_db2pb(condition))
		}
	}
	//appList
	appMap := map[int64][]*proto.SceneTemplateAppRelation{}
	appList, err := q.TSceneTemplateAppRelation.WithContext(context.Background()).Where(q.TSceneTemplateAppRelation.SceneTemplateId.In(templateIds...)).Find()
	if err == nil {
		for _, appRelation := range appList {
			appMap[appRelation.SceneTemplateId] = append(appMap[appRelation.SceneTemplateId], convert.SceneTemplateAppRelation_db2pb(appRelation))
		}
	}
	return taskMap, conditionMap, appMap
}
