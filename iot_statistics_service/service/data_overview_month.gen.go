// Code generated by sgen.exe,2022-06-17 09:58:13. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"cloud_platform/iot_statistics_service/config"
	"context"
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_statistics/model"
	"cloud_platform/iot_model/db_statistics/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_statistics_service/convert"
)

type DataOverviewMonthSvc struct {
	Ctx context.Context
}

// 创建DataOverviewMonth
func (s *DataOverviewMonthSvc) CreateDataOverviewMonth(req *proto.DataOverviewMonth) (*proto.DataOverviewMonth, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TDataOverviewMonth
	do := t.WithContext(context.Background())
	dbObj := convert.DataOverviewMonth_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateDataOverviewMonth error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除DataOverviewMonth
func (s *DataOverviewMonthSvc) DeleteDataOverviewMonth(req *proto.DataOverviewMonth) (*proto.DataOverviewMonth, error) {
	t := orm.Use(iotmodel.GetDB()).TDataOverviewMonth
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.TenantId != "" { //字符串
		do = do.Where(t.TenantId.Eq(req.TenantId))
	}
	if req.DeviceActiveSum != 0 { //整数
		do = do.Where(t.DeviceActiveSum.Eq(req.DeviceActiveSum))
	}
	if req.DeviceFaultSum != 0 { //整数
		do = do.Where(t.DeviceFaultSum.Eq(req.DeviceFaultSum))
	}
	if req.DeveloperRegisterSum != 0 { //整数
		do = do.Where(t.DeveloperRegisterSum.Eq(req.DeveloperRegisterSum))
	}
	if req.UserRegisterSum != 0 { //整数
		do = do.Where(t.UserRegisterSum.Eq(req.UserRegisterSum))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteDataOverviewMonth error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除DataOverviewMonth
func (s *DataOverviewMonthSvc) DeleteByIdDataOverviewMonth(req *proto.DataOverviewMonth) (*proto.DataOverviewMonth, error) {
	t := orm.Use(iotmodel.GetDB()).TDataOverviewMonth
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.TenantId != "" { //字符串
		do = do.Where(t.TenantId.Eq(req.TenantId))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdDataOverviewMonth error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除DataOverviewMonth
func (s *DataOverviewMonthSvc) DeleteByIdsDataOverviewMonth(req *proto.DataOverviewMonthBatchDeleteRequest) (*proto.DataOverviewMonthBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TDataOverviewMonth
		do := t.WithContext(context.Background())

		do = do.Where(t.DataTime.Eq(k.DataTime.AsTime()))

		do = do.Where(t.TenantId.Eq(k.TenantId))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsDataOverviewMonth error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新DataOverviewMonth
func (s *DataOverviewMonthSvc) UpdateDataOverviewMonth(req *proto.DataOverviewMonth) (*proto.DataOverviewMonth, error) {
	t := orm.Use(iotmodel.GetDB()).TDataOverviewMonth
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.DeviceActiveSum != 0 { //整数
		updateField = append(updateField, t.DeviceActiveSum)
	}
	if req.DeviceFaultSum != 0 { //整数
		updateField = append(updateField, t.DeviceFaultSum)
	}
	if req.DeveloperRegisterSum != 0 { //整数
		updateField = append(updateField, t.DeveloperRegisterSum)
	}
	if req.UserRegisterSum != 0 { //整数
		updateField = append(updateField, t.UserRegisterSum)
	}
	if len(updateField) > 0 {
		do = do.Select(updateField...)
	}
	//主键条件
	HasPrimaryKey := false

	if req.TenantId != "" { //字符串
		do = do.Where(t.TenantId.Eq(req.TenantId))
		HasPrimaryKey = true
	}

	if !HasPrimaryKey {
		logger.Error("UpdateDataOverviewMonth error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.DataOverviewMonth_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateDataOverviewMonth error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段DataOverviewMonth
func (s *DataOverviewMonthSvc) UpdateAllDataOverviewMonth(req *proto.DataOverviewMonth) (*proto.DataOverviewMonth, error) {
	t := orm.Use(iotmodel.GetDB()).TDataOverviewMonth
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.DeviceActiveSum)
	updateField = append(updateField, t.DeviceFaultSum)
	updateField = append(updateField, t.DeveloperRegisterSum)
	updateField = append(updateField, t.UserRegisterSum)
	if len(updateField) > 0 {
		do = do.Select(updateField...)
	}
	//主键条件
	HasPrimaryKey := false
	if req.TenantId != "" { //字符串
		do = do.Where(t.TenantId.Eq(req.TenantId))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateAllDataOverviewMonth error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.DataOverviewMonth_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllDataOverviewMonth error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *DataOverviewMonthSvc) UpdateFieldsDataOverviewMonth(req *proto.DataOverviewMonthUpdateFieldsRequest) (*proto.DataOverviewMonth, error) {
	t := orm.Use(iotmodel.GetDB()).TDataOverviewMonth
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsDataOverviewMonth error : missing updateField")
		logger.Error(err)
		return nil, err
	}
	do = do.Select(updateField...)

	//主键条件
	HasPrimaryKey := false
	if req.Data.TenantId != "" { //字符串
		do = do.Where(t.TenantId.Eq(req.Data.TenantId))
		HasPrimaryKey = true
	}
	if !HasPrimaryKey {
		logger.Error("UpdateFieldsDataOverviewMonth error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.DataOverviewMonth_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsDataOverviewMonth error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找DataOverviewMonth
func (s *DataOverviewMonthSvc) FindDataOverviewMonth(req *proto.DataOverviewMonthFilter) (*proto.DataOverviewMonth, error) {
	t := orm.Use(iotmodel.GetDB()).TDataOverviewMonth
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.TenantId != "" { //字符串
		do = do.Where(t.TenantId.Eq(req.TenantId))
	}
	if req.DeviceActiveSum != 0 { //整数
		do = do.Where(t.DeviceActiveSum.Eq(req.DeviceActiveSum))
	}
	if req.DeviceFaultSum != 0 { //整数
		do = do.Where(t.DeviceFaultSum.Eq(req.DeviceFaultSum))
	}
	if req.DeveloperRegisterSum != 0 { //整数
		do = do.Where(t.DeveloperRegisterSum.Eq(req.DeveloperRegisterSum))
	}
	if req.UserRegisterSum != 0 { //整数
		do = do.Where(t.UserRegisterSum.Eq(req.UserRegisterSum))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindDataOverviewMonth error : %s", err.Error())
		return nil, err
	}
	res := convert.DataOverviewMonth_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找DataOverviewMonth
func (s *DataOverviewMonthSvc) FindByIdDataOverviewMonth(req *proto.DataOverviewMonthFilter) (*proto.DataOverviewMonth, error) {
	t := orm.Use(iotmodel.GetDB()).TDataOverviewMonth
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.TenantId != "" { //字符串
		do = do.Where(t.TenantId.Eq(req.TenantId))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdDataOverviewMonth error : %s", err.Error())
		return nil, err
	}
	res := convert.DataOverviewMonth_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找DataOverviewMonth,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *DataOverviewMonthSvc) GetListDataOverviewMonth(req *proto.DataOverviewMonthListRequest) ([]*proto.DataOverviewMonth, int64, error) {
	// fixme 请检查条件和校验参数
	db, ok := config.DBMap["iot_statistics"]
	if !ok {
		return nil, 0, errors.New("数据库未初始化")
	}
	var err error
	t := orm.Use(db).TDataOverviewMonth
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {
		do = do.Where(t.TenantId.Eq(query.TenantId))
		if query.StartTime != nil {
			do = do.Where(t.DataTime.Gte(query.StartTime.AsTime()))
		}
		if query.EndTime != nil {
			do = do.Where(t.DataTime.Lte(query.EndTime.AsTime()))
		} else {
			do = do.Where(t.DataTime.Lte(time.Now()))
		}
	} else {
		return nil, 0, errors.New("缺查询条件")
	}
	orderCol, ok := t.GetFieldByName(req.OrderKey)
	if !ok {
		orderCol = t.DataTime
		orderCol = t.TenantId
	}
	if req.OrderDesc != "" {
		do = do.Order(orderCol.Desc())
	} else {
		do = do.Order(orderCol)
	}

	var list []*model.TDataOverviewMonth
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
		logger.Errorf("GetListDataOverviewMonth error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.DataOverviewMonth, 0, len(list)+1)
	for _, v := range list {
		result = append(result, convert.DataOverviewMonth_db2pb(v))
	}

	if query.TenantId == "" {
		if totalObj, err := s.GetPMDataOverviewTotal(); err != nil {
			return nil, 0, errors.New("查询总计失败")
		} else {
			result = append(result, totalObj)
		}
	} else {
		if totalObj, err := s.GetOpenDataOverviewTotal(query.TenantId); err != nil {
			return nil, 0, errors.New("查询总计失败")
		} else {
			result = append(result, totalObj)
		}
	}
	return result, total, nil
}

func (s *DataOverviewMonthSvc) GetOpenDataOverviewTotal(tenantId string) (*proto.DataOverviewMonth, error) {
	db, ok := config.DBMap["iot_statistics"]
	if !ok {
		return nil, errors.New("数据库未初始化")
	}
	var err error
	var scanObj TotalStruct2
	t := orm.Use(db).TDeviceDataSum
	err = t.WithContext(context.Background()).Select(t.ActiveSum.Sum().IfNull(0).As("device_active_sum"),
		t.FaultSum.Sum().IfNull(0).As("device_fault_sum")).Where(t.TenantId.Eq(tenantId)).Scan(&scanObj)
	if err != nil {
		return nil, err
	}
	var totalObj proto.DataOverviewMonth
	totalObj.DataTime = timestamppb.New(time.Time{})
	totalObj.DeviceActiveSum = scanObj.DeviceActiveSum
	totalObj.DeviceFaultSum = scanObj.DeviceFaultSum

	var total int64
	t2 := orm.Use(db).TAppUserSum
	err = t2.WithContext(context.Background()).Select(t2.RegisterSum.Sum().IfNull(0).As("total")).Where(t2.TenantId.Eq(tenantId)).Scan(&total)
	if err != nil {
		return nil, err
	}
	totalObj.UserRegisterSum = total
	return &totalObj, nil
}

func (s *DataOverviewMonthSvc) GetPMDataOverviewTotal() (*proto.DataOverviewMonth, error) {
	db, ok := config.DBMap["iot_statistics"]
	if !ok {
		return nil, errors.New("数据库未初始化")
	}
	var err error
	var scanObj TotalStruct2
	t := orm.Use(db).TDeviceDataSum
	err = t.WithContext(context.Background()).Select(t.ActiveSum.Sum().IfNull(0).As("device_active_sum"),
		t.FaultSum.Sum().As("device_fault_sum")).Scan(&scanObj)
	if err != nil {
		return nil, err
	}
	var totalObj proto.DataOverviewMonth
	totalObj.DataTime = timestamppb.New(time.Time{})
	totalObj.DeviceActiveSum = scanObj.DeviceActiveSum
	totalObj.DeviceFaultSum = scanObj.DeviceFaultSum

	var total int64
	t2 := orm.Use(db).TAppUserSum
	err = t2.WithContext(context.Background()).Select(t2.RegisterSum.Sum().IfNull(0).As("total")).Scan(&total)
	if err != nil {
		return nil, err
	}
	totalObj.UserRegisterSum = total

	//云管平台有开发者统计
	total = 0
	t3 := orm.Use(db).TDeveloperSum
	err = t3.WithContext(context.Background()).Select(t3.DeveloperSum.IfNull(0).As("total")).Scan(&total)
	if err != nil {
		return nil, err
	}
	totalObj.DeveloperRegisterSum = total
	return &totalObj, nil
}

type TotalStruct struct {
	DeviceActiveSum      int64
	DeviceFaultSum       int64
	DeveloperRegisterSum int64
	UserRegisterSum      int64
}

type TotalStruct2 struct {
	DeviceActiveSum int64
	DeviceFaultSum  int64
}
