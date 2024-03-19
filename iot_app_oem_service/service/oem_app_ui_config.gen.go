// Code generated by sgen.exe,2022-06-02 11:15:12. DO NOT EDIT.
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

type OemAppUiConfigSvc struct {
	Ctx context.Context
}

// 创建OemAppUiConfig
func (s *OemAppUiConfigSvc) CreateOemAppUiConfig(req *proto.OemAppUiConfig) (*proto.OemAppUiConfig, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TOemAppUiConfig
	do := t.WithContext(context.Background())
	dbObj := convert.OemAppUiConfig_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateOemAppUiConfig error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除OemAppUiConfig
func (s *OemAppUiConfigSvc) DeleteOemAppUiConfig(req *proto.OemAppUiConfig) (*proto.OemAppUiConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppUiConfig
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
	if req.IconUrl != "" { //字符串
		do = do.Where(t.IconUrl.Eq(req.IconUrl))
	}
	if req.IosLaunchScreen != "" { //字符串
		do = do.Where(t.IosLaunchScreen.Eq(req.IosLaunchScreen))
	}
	if req.AndroidLaunchScreen != "" { //字符串
		do = do.Where(t.AndroidLaunchScreen.Eq(req.AndroidLaunchScreen))
	}
	if req.ThemeColors != "" { //字符串
		do = do.Where(t.ThemeColors.Eq(req.ThemeColors))
	}
	if req.BottomMenu != "" { //字符串
		do = do.Where(t.BottomMenu.Eq(req.BottomMenu))
	}
	if req.Personalize != "" { //字符串
		do = do.Where(t.Personalize.Eq(req.Personalize))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteOemAppUiConfig error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除OemAppUiConfig
func (s *OemAppUiConfigSvc) DeleteByIdOemAppUiConfig(req *proto.OemAppUiConfig) (*proto.OemAppUiConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppUiConfig
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdOemAppUiConfig error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除OemAppUiConfig
func (s *OemAppUiConfigSvc) DeleteByIdsOemAppUiConfig(req *proto.OemAppUiConfigBatchDeleteRequest) (*proto.OemAppUiConfigBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TOemAppUiConfig
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsOemAppUiConfig error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新OemAppUiConfig
func (s *OemAppUiConfigSvc) UpdateOemAppUiConfig(req *proto.OemAppUiConfig) (*proto.OemAppUiConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppUiConfig
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
	if req.IconUrl != "" { //字符串
		updateField = append(updateField, t.IconUrl)
	}
	if req.IosLaunchScreen != "" { //字符串
		updateField = append(updateField, t.IosLaunchScreen)
	}
	if req.AndroidLaunchScreen != "" { //字符串
		updateField = append(updateField, t.AndroidLaunchScreen)
	}
	if req.ThemeColors != "" { //字符串
		updateField = append(updateField, t.ThemeColors)
	}
	if req.BottomMenu != "" { //字符串
		updateField = append(updateField, t.BottomMenu)
	}
	if req.Personalize != "" { //字符串
		updateField = append(updateField, t.Personalize)
	}
	if req.Room != "" { //字符串
		updateField = append(updateField, t.Room)
	}
	if req.RoomIcons != "" {
		updateField = append(updateField, t.RoomIcons)
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
		logger.Error("UpdateOemAppUiConfig error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.OemAppUiConfig_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateOemAppUiConfig error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段OemAppUiConfig
func (s *OemAppUiConfigSvc) UpdateAllOemAppUiConfig(req *proto.OemAppUiConfig) (*proto.OemAppUiConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppUiConfig
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.AppId)
	updateField = append(updateField, t.Version)
	updateField = append(updateField, t.IconUrl)
	updateField = append(updateField, t.IosLaunchScreen)
	updateField = append(updateField, t.AndroidLaunchScreen)
	updateField = append(updateField, t.ThemeColors)
	updateField = append(updateField, t.BottomMenu)
	updateField = append(updateField, t.Personalize)
	updateField = append(updateField, t.Room)
	updateField = append(updateField, t.RoomIcons)

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
		logger.Error("UpdateAllOemAppUiConfig error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OemAppUiConfig_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllOemAppUiConfig error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *OemAppUiConfigSvc) UpdateFieldsOemAppUiConfig(req *proto.OemAppUiConfigUpdateFieldsRequest) (*proto.OemAppUiConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppUiConfig
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsOemAppUiConfig error : missing updateField")
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
		logger.Error("UpdateFieldsOemAppUiConfig error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OemAppUiConfig_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsOemAppUiConfig error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找OemAppUiConfig
func (s *OemAppUiConfigSvc) FindOemAppUiConfig(req *proto.OemAppUiConfigFilter) (*proto.OemAppUiConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppUiConfig
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
	if req.IconUrl != "" { //字符串
		do = do.Where(t.IconUrl.Eq(req.IconUrl))
	}
	if req.IosLaunchScreen != "" { //字符串
		do = do.Where(t.IosLaunchScreen.Eq(req.IosLaunchScreen))
	}
	if req.AndroidLaunchScreen != "" { //字符串
		do = do.Where(t.AndroidLaunchScreen.Eq(req.AndroidLaunchScreen))
	}
	if req.ThemeColors != "" { //字符串
		do = do.Where(t.ThemeColors.Eq(req.ThemeColors))
	}
	if req.BottomMenu != "" { //字符串
		do = do.Where(t.BottomMenu.Eq(req.BottomMenu))
	}
	if req.Personalize != "" { //字符串
		do = do.Where(t.Personalize.Eq(req.Personalize))
	}
	if req.Room != "" {
		do = do.Where(t.Room.Eq(req.Room))
	}

	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindOemAppUiConfig error : %s", err.Error())
		return nil, err
	}
	res := convert.OemAppUiConfig_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找OemAppUiConfig
func (s *OemAppUiConfigSvc) FindByIdOemAppUiConfig(req *proto.OemAppUiConfigFilter) (*proto.OemAppUiConfig, error) {
	t := orm.Use(iotmodel.GetDB()).TOemAppUiConfig
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdOemAppUiConfig error : %s", err.Error())
		return nil, err
	}
	res := convert.OemAppUiConfig_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找OemAppUiConfig,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *OemAppUiConfigSvc) GetListOemAppUiConfig(req *proto.OemAppUiConfigListRequest) ([]*proto.OemAppUiConfig, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TOemAppUiConfig
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
		if query.IconUrl != "" { //字符串
			do = do.Where(t.IconUrl.Like("%" + query.IconUrl + "%"))
		}
		if query.IosLaunchScreen != "" { //字符串
			do = do.Where(t.IosLaunchScreen.Like("%" + query.IosLaunchScreen + "%"))
		}
		if query.AndroidLaunchScreen != "" { //字符串
			do = do.Where(t.AndroidLaunchScreen.Like("%" + query.AndroidLaunchScreen + "%"))
		}
		if query.ThemeColors != "" { //字符串
			do = do.Where(t.ThemeColors.Like("%" + query.ThemeColors + "%"))
		}
		if query.BottomMenu != "" { //字符串
			do = do.Where(t.BottomMenu.Like("%" + query.BottomMenu + "%"))
		}
		if query.Personalize != "" { //字符串
			do = do.Where(t.Personalize.Like("%" + query.Personalize + "%"))
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

	var list []*model.TOemAppUiConfig
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
		logger.Errorf("GetListOemAppUiConfig error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.OemAppUiConfig, len(list))
	for i, v := range list {
		result[i] = convert.OemAppUiConfig_db2pb(v)
	}
	return result, total, nil
}
