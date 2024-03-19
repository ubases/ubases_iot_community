// Code generated by sgen.exe,2022-04-21 13:46:11. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_device_service/cached"
	"cloud_platform/iot_device_service/config"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	"cloud_platform/iot_device_service/convert"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_model/db_device/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type IotDeviceInfoSvc struct {
	Ctx context.Context
}

func (s *IotDeviceInfoSvc) ActiveDevice(req *proto.ActiveDeviceRequest) error {
	//var resObj DeviceInfoReport
	//err := json.Unmarshal([]byte(req.ReportMsg), &resObj)
	//if err != nil {
	//	iotlogger.LogHelper.Error("数据转换失败," + err.Error())
	//	return err
	//}
	//return ActiveDevice(resObj)
	return nil
}

func (s *IotDeviceInfoSvc) SetOnlineStatus(did string, onlineStatus int32) error {
	tx := orm.Use(iotmodel.GetDB())
	do := tx.TIotDeviceInfo.WithContext(context.Background())
	_, err := do.Where(tx.TIotDeviceInfo.Did.Eq(did)).Update(tx.TIotDeviceInfo.OnlineStatus, onlineStatus)
	if err != nil {
		logger.Errorf("SetUse delete old error : %s", err.Error())
		return err
	}
	if err != nil {
		logger.Errorf("SetUse error : %s", err.Error())
		return err
	}
	return nil
}

func (s *IotDeviceInfoSvc) TranCreate(tx *orm.Query, req *model.TIotDeviceInfo) error {
	//查询之前的激活时间
	//do1 := tx.TIotDeviceInfo.WithContext(context.Background())
	//var tt time.Time
	//_ = do1.Unscoped().Select(tx.TIotDeviceInfo.ActivatedTime.Min().As("tt")).Where(tx.TIotDeviceInfo.Did.Eq(req.Did)).Scan(&tt)
	//if !tt.IsZero() {
	//	req.ActivatedTime = tt
	//}

	do := tx.TIotDeviceInfo.WithContext(context.Background())
	//TODO 先删除原来的设备信息
	_, err := do.Where(tx.TIotDeviceInfo.Did.Eq(req.Did)).Delete()
	if err != nil {
		logger.Errorf("TransalteCreate delete old error : %s", err.Error())
		return err
	}
	err = do.Create(req)
	if err != nil {
		logger.Errorf("TransalteCreate error : %s", err.Error())
		return err
	}
	return nil
}

// 创建IotDeviceInfo
func (s *IotDeviceInfoSvc) CreateIotDeviceInfo(req *proto.IotDeviceInfo) (*proto.IotDeviceInfo, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
	do := t.WithContext(context.Background())
	dbObj := convert.IotDeviceInfo_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateIotDeviceInfo error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除IotDeviceInfo
func (s *IotDeviceInfoSvc) DeleteIotDeviceInfo(req *proto.IotDeviceInfo) (*proto.IotDeviceInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.Did != "" { //字符串
		do = do.Where(t.Did.Eq(req.Did))
	}
	if req.ProductId != 0 { //整数
		do = do.Where(t.ProductId.Eq(req.ProductId))
	}
	if req.OnlineStatus != 0 { //整数
		do = do.Where(t.OnlineStatus.Eq(req.OnlineStatus))
	}
	if req.DeviceName != "" { //字符串
		do = do.Where(t.DeviceName.Eq(req.DeviceName))
	}
	if req.DeviceNature != "" { //字符串
		do = do.Where(t.DeviceNature.Eq(req.DeviceNature))
	}
	if req.Sn != "" { //字符串
		do = do.Where(t.Sn.Eq(req.Sn))
	}
	if req.BatchId != 0 { //整数
		do = do.Where(t.BatchId.Eq(req.BatchId))
	}
	if req.GroupId != 0 { //整数
		do = do.Where(t.GroupId.Eq(req.GroupId))
	}
	if req.DeviceModel != "" { //字符串
		do = do.Where(t.DeviceModel.Eq(req.DeviceModel))
	}
	if req.UserName != "" { //字符串
		do = do.Where(t.UserName.Eq(req.UserName))
	}
	if req.Passward != "" { //字符串
		do = do.Where(t.Passward.Eq(req.Passward))
	}
	if req.Salt != "" { //字符串
		do = do.Where(t.Salt.Eq(req.Salt))
	}
	if req.DeviceSecretHttp != "" { //字符串
		do = do.Where(t.DeviceSecretHttp.Eq(req.DeviceSecretHttp))
	}
	if req.DeviceSecretMqtt != "" { //字符串
		do = do.Where(t.DeviceSecretMqtt.Eq(req.DeviceSecretMqtt))
	}
	if req.IpAddress != "" { //字符串
		do = do.Where(t.IpAddress.Eq(req.IpAddress))
	}
	if req.Lat != 0 { //整数
		do = do.Where(t.Lat.Eq(req.Lat))
	}
	if req.Lng != 0 { //整数
		do = do.Where(t.Lng.Eq(req.Lng))
	}
	if req.Country != "" { //整数
		do = do.Where(t.Country.Eq(req.Country))
	}
	if req.Province != "" { //字符串
		do = do.Where(t.Province.Eq(req.Province))
	}
	if req.City != "" { //字符串
		do = do.Where(t.City.Eq(req.City))
	}
	if req.District != "" { //字符串
		do = do.Where(t.District.Eq(req.District))
	}
	//if req.ActivatedTime != "" { //字符串
	//	do = do.Where(t.ActivatedTime.Eq(req.ActivatedTime))
	//}
	if req.MacAddress != "" { //字符串
		do = do.Where(t.MacAddress.Eq(req.MacAddress))
	}
	if req.DeviceVersion != "" { //字符串
		do = do.Where(t.DeviceVersion.Eq(req.DeviceVersion))
	}
	if req.ActiveStatus != "" { //字符串
		do = do.Where(t.ActiveStatus.Eq(req.ActiveStatus))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteIotDeviceInfo error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除IotDeviceInfo
func (s *IotDeviceInfoSvc) DeleteByIdIotDeviceInfo(req *proto.IotDeviceInfo) (*proto.IotDeviceInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdIotDeviceInfo error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除IotDeviceInfo
func (s *IotDeviceInfoSvc) DeleteByIdsIotDeviceInfo(req *proto.IotDeviceInfoBatchDeleteRequest) (*proto.IotDeviceInfoBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsIotDeviceInfo error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新IotDeviceInfo
func (s *IotDeviceInfoSvc) UpdateIotDeviceInfo(req *proto.IotDeviceInfo) (*proto.IotDeviceInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.Did != "" { //字符串
		updateField = append(updateField, t.Did)
	}
	if req.ProductId != 0 { //整数
		updateField = append(updateField, t.ProductId)
	}
	if req.OnlineStatus != 0 { //整数
		updateField = append(updateField, t.OnlineStatus)
	}
	if req.DeviceName != "" { //字符串
		updateField = append(updateField, t.DeviceName)
	}
	if req.DeviceNature != "" { //字符串
		updateField = append(updateField, t.DeviceNature)
	}
	if req.Sn != "" { //字符串
		updateField = append(updateField, t.Sn)
	}
	if req.BatchId != 0 { //整数
		updateField = append(updateField, t.BatchId)
	}
	if req.GroupId != 0 { //整数
		updateField = append(updateField, t.GroupId)
	}
	if req.DeviceModel != "" { //字符串
		updateField = append(updateField, t.DeviceModel)
	}
	if req.UserName != "" { //字符串
		updateField = append(updateField, t.UserName)
	}
	if req.Passward != "" { //字符串
		updateField = append(updateField, t.Passward)
	}
	if req.Salt != "" { //字符串
		updateField = append(updateField, t.Salt)
	}
	if req.DeviceSecretHttp != "" { //字符串
		updateField = append(updateField, t.DeviceSecretHttp)
	}
	if req.DeviceSecretMqtt != "" { //字符串
		updateField = append(updateField, t.DeviceSecretMqtt)
	}
	if req.IpAddress != "" { //字符串
		updateField = append(updateField, t.IpAddress)
	}
	if req.Lat != 0 { //整数
		updateField = append(updateField, t.Lat)
	}
	if req.Lng != 0 { //整数
		updateField = append(updateField, t.Lng)
	}
	if req.Country != "" { //整数
		updateField = append(updateField, t.Country)
	}
	if req.Province != "" { //字符串
		updateField = append(updateField, t.Province)
	}
	if req.City != "" { //字符串
		updateField = append(updateField, t.City)
	}
	if req.District != "" { //字符串
		updateField = append(updateField, t.District)
	}
	//if req.ActivatedTime != "" { //字符串
	//	updateField = append(updateField, t.ActivatedTime)
	//}
	if req.MacAddress != "" { //字符串
		updateField = append(updateField, t.MacAddress)
	}
	if req.DeviceVersion != "" { //字符串
		updateField = append(updateField, t.DeviceVersion)
	}
	if req.ActiveStatus != "" { //字符串
		updateField = append(updateField, t.ActiveStatus)
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
		logger.Error("UpdateIotDeviceInfo error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.IotDeviceInfo_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateIotDeviceInfo error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段IotDeviceInfo
func (s *IotDeviceInfoSvc) UpdateAllIotDeviceInfo(req *proto.IotDeviceInfo) (*proto.IotDeviceInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.Did)
	updateField = append(updateField, t.ProductId)
	updateField = append(updateField, t.OnlineStatus)
	updateField = append(updateField, t.DeviceName)
	updateField = append(updateField, t.DeviceNature)
	updateField = append(updateField, t.Sn)
	updateField = append(updateField, t.BatchId)
	updateField = append(updateField, t.GroupId)
	updateField = append(updateField, t.DeviceModel)
	updateField = append(updateField, t.UserName)
	updateField = append(updateField, t.Passward)
	updateField = append(updateField, t.Salt)
	updateField = append(updateField, t.DeviceSecretHttp)
	updateField = append(updateField, t.DeviceSecretMqtt)
	updateField = append(updateField, t.IpAddress)
	updateField = append(updateField, t.Lat)
	updateField = append(updateField, t.Lng)
	updateField = append(updateField, t.Country)
	updateField = append(updateField, t.Province)
	updateField = append(updateField, t.City)
	updateField = append(updateField, t.District)
	updateField = append(updateField, t.ActivatedTime)
	updateField = append(updateField, t.MacAddress)
	updateField = append(updateField, t.DeviceVersion)
	updateField = append(updateField, t.ActiveStatus)
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
		logger.Error("UpdateAllIotDeviceInfo error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.IotDeviceInfo_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllIotDeviceInfo error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *IotDeviceInfoSvc) UpdateFieldsIotDeviceInfo(req *proto.IotDeviceInfoUpdateFieldsRequest) (*proto.IotDeviceInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsIotDeviceInfo error : missing updateField")
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
		logger.Error("UpdateFieldsIotDeviceInfo error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.IotDeviceInfo_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsIotDeviceInfo error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找IotDeviceInfo
func (s *IotDeviceInfoSvc) FindIotDeviceInfo(req *proto.IotDeviceInfoFilter) (*proto.IotDeviceInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.Did != "" { //字符串
		do = do.Where(t.Did.Like("%" + req.Did + "%"))
	}
	if req.ProductId != 0 { //整数
		do = do.Where(t.ProductId.Eq(req.ProductId))
	}
	if req.OnlineStatus != 0 { //整数
		do = do.Where(t.OnlineStatus.Eq(req.OnlineStatus))
	}
	if req.DeviceName != "" { //字符串
		do = do.Where(t.DeviceName.Like("%" + req.DeviceName + "%"))
	}
	if req.DeviceNature != "" { //字符串
		do = do.Where(t.DeviceNature.Like("%" + req.DeviceNature + "%"))
	}
	if req.Sn != "" { //字符串
		do = do.Where(t.Sn.Like("%" + req.Sn + "%"))
	}
	if req.BatchId != 0 { //整数
		do = do.Where(t.BatchId.Eq(req.BatchId))
	}
	if req.GroupId != 0 { //整数
		do = do.Where(t.GroupId.Eq(req.GroupId))
	}
	if req.DeviceModel != "" { //字符串
		do = do.Where(t.DeviceModel.Like("%" + req.DeviceModel + "%"))
	}
	if req.UserName != "" { //字符串
		do = do.Where(t.UserName.Like("%" + req.UserName + "%"))
	}
	if req.Passward != "" { //字符串
		do = do.Where(t.Passward.Like("%" + req.Passward + "%"))
	}
	if req.Salt != "" { //字符串
		do = do.Where(t.Salt.Like("%" + req.Salt + "%"))
	}
	if req.DeviceSecretHttp != "" { //字符串
		do = do.Where(t.DeviceSecretHttp.Like("%" + req.DeviceSecretHttp + "%"))
	}
	if req.DeviceSecretMqtt != "" { //字符串
		do = do.Where(t.DeviceSecretMqtt.Like("%" + req.DeviceSecretMqtt + "%"))
	}
	if req.IpAddress != "" { //字符串
		do = do.Where(t.IpAddress.Like("%" + req.IpAddress + "%"))
	}
	if req.Lat != 0 { //整数
		do = do.Where(t.Lat.Eq(req.Lat))
	}
	if req.Lng != 0 { //整数
		do = do.Where(t.Lng.Eq(req.Lng))
	}
	if req.Country != "" { //整数
		do = do.Where(t.Country.Eq(req.Country))
	}
	if req.Province != "" { //字符串
		do = do.Where(t.Province.Like("%" + req.Province + "%"))
	}
	if req.City != "" { //字符串
		do = do.Where(t.City.Like("%" + req.City + "%"))
	}
	if req.District != "" { //字符串
		do = do.Where(t.District.Like("%" + req.District + "%"))
	}
	//if req.ActivatedTime != "" { //字符串
	//	do = do.Where(t.ActivatedTime.Like("%" + req.ActivatedTime + "%"))
	//}
	if req.MacAddress != "" { //字符串
		do = do.Where(t.MacAddress.Like("%" + req.MacAddress + "%"))
	}
	if req.DeviceVersion != "" { //字符串
		do = do.Where(t.DeviceVersion.Like("%" + req.DeviceVersion + "%"))
	}
	if req.ActiveStatus != "" { //字符串
		do = do.Where(t.ActiveStatus.Like("%" + req.ActiveStatus + "%"))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindIotDeviceInfo error : %s", err.Error())
		return nil, err
	}
	res := convert.IotDeviceInfo_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找IotDeviceInfo
func (s *IotDeviceInfoSvc) FindByIdIotDeviceInfo(req *proto.IotDeviceInfoFilter) (*proto.IotDeviceInfo, error) {
	t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdIotDeviceInfo error : %s", err.Error())
		return nil, err
	}
	res := convert.IotDeviceInfo_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找IotDeviceInfo,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *IotDeviceInfoSvc) GetListIotDeviceInfo(req *proto.IotDeviceInfoListRequest) ([]*proto.IotDeviceInfo, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	tenantId, err := CheckTenantId(s.Ctx)
	if err != nil {
		return nil, 0, err
	}
	//userId, err := CheckUserId(s.Ctx)
	//if err != nil {
	//	return nil, 0, err
	//}

	t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
	do := t.WithContext(context.Background())
	query := req.Query

	//var SearchKeyid int = 0
	//if req.SearchKey != "" {
	//	SearchKeyid, _ = strconv.Atoi(req.SearchKey)
	//}
	//fixme 临时处理方案，主要需要通过用户的产品进行过滤 （District暂时用作tenantId）
	if tenantId != "" { //整数
		do = do.Where(t.District.Eq(tenantId))
	}
	if query != nil {
		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}

		if query.Did != "" { //字符串
			do = do.Where(t.Did.Eq(query.Did))
		}
		//else {
		//	if req.SearchKey != "" {
		//		do = do.Where(t.Did.Like("%" + req.SearchKey + "%"))
		//	}
		//}
		if query.ProductId != 0 { //整数
			do = do.Where(t.ProductId.Eq(query.ProductId))
		}
		if query.OnlineStatus >= 0 { //整数
			do = do.Where(t.OnlineStatus.Eq(query.OnlineStatus))
		}
		if query.DeviceName != "" { //字符串
			SearchKeyid, _ := strconv.Atoi(query.DeviceName)
			if SearchKeyid == 0 {
				do = do.Where(do.Where(t.DeviceName.Like("%" + query.DeviceName + "%")).Or(t.Did.Like("%" + query.DeviceName + "%")))
			} else {
				do = do.Where(do.Where(t.DeviceName.Like("%" + query.DeviceName + "%")).Or(t.Did.Like("%" + query.DeviceName + "%")).Or(t.ProductId.Eq(int64(SearchKeyid))))
			}
		}
		if query.DeviceNature != "" { //字符串
			do = do.Where(t.DeviceNature.Eq(query.DeviceNature))
		}
		if query.Sn != "" { //字符串
			do = do.Where(t.Sn.Like("%" + query.Sn + "%"))
		}
		if query.BatchId != 0 { //整数
			do = do.Where(t.BatchId.Eq(query.BatchId))
		}
		if query.GroupId != 0 { //整数
			do = do.Where(t.GroupId.Eq(query.GroupId))
		}
		if query.DeviceModel != "" { //字符串
			do = do.Where(t.DeviceModel.Like("%" + query.DeviceModel + "%"))
		}
		if query.UserName != "" { //字符串
			do = do.Where(t.UserName.Like("%" + query.UserName + "%"))
		}
		if query.DeviceSecretHttp != "" { //字符串
			do = do.Where(t.DeviceSecretHttp.Like("%" + query.DeviceSecretHttp + "%"))
		}
		if query.DeviceSecretMqtt != "" { //字符串
			do = do.Where(t.DeviceSecretMqtt.Like("%" + query.DeviceSecretMqtt + "%"))
		}
		if query.IpAddress != "" { //字符串
			do = do.Where(t.IpAddress.Like("%" + query.IpAddress + "%"))
		}
		if query.Lat != 0 { //整数
			do = do.Where(t.Lat.Eq(query.Lat))
		}
		if query.Lng != 0 { //整数
			do = do.Where(t.Lng.Eq(query.Lng))
		}
		if query.Country != "" { //整数
			do = do.Where(t.Country.Eq(query.Country))
		}
		if query.Province != "" { //字符串
			do = do.Where(t.Province.Eq(query.Province))
		}
		if query.City != "" { //字符串
			do = do.Where(t.City.Eq(query.City))
		}
		if query.District != "" { //字符串
			do = do.Where(t.District.Eq(query.District))
		}
		//if query.ActivatedTime != "" { //字符串
		//	do = do.Where(t.ActivatedTime.Like("%" + query.ActivatedTime + "%"))
		//}
		if query.MacAddress != "" { //字符串
			do = do.Where(t.MacAddress.Eq(query.MacAddress))
		}
		if query.DeviceVersion != "" { //字符串
			do = do.Where(t.DeviceVersion.Eq(query.DeviceVersion))
		}
		if query.ActiveStatus != "" { //字符串
			do = do.Where(t.ActiveStatus.Eq(query.ActiveStatus))
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

	var list []*model.TIotDeviceInfo
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
		logger.Errorf("GetListIotDeviceInfo error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.IotDeviceInfo, len(list))
	for i, v := range list {
		result[i] = convert.IotDeviceInfo_db2pb(v)
	}
	return result, total, nil
}

// 设备统计信息
func (s *IotDeviceInfoSvc) GetCountIotDeviceInfoCount(req *proto.IotDeviceInfoListRequest) (*proto.IotDeviceInfoCount, error) {
	//var err error
	tenantId, _ := CheckTenantId(s.Ctx)
	//if err != nil {
	//	return nil, err
	//}
	q := orm.Use(iotmodel.GetDB())
	t := q.TIotDeviceInfo
	tTriad := q.TIotDeviceTriad

	var totalCount int64 = 0
	var userCount int64 = 0
	var onlineCount int64 = 0

	//有租户Id代表开发平台调用接口查询
	if tenantId != "" {
		totalCount, _ = tTriad.WithContext(context.Background()).Where(tTriad.TenantId.Eq(tenantId), tTriad.UseType.Eq(0), tTriad.FirstActiveTime.IsNotNull(), tTriad.FirstActiveTime.Neq(time.Time{})).Count()
		userCount, _ = tTriad.WithContext(context.Background()).LeftJoin(t, tTriad.Did.EqCol(t.Did)).Where(tTriad.TenantId.Eq(tenantId), tTriad.UseType.Eq(0), tTriad.Status.Eq(1), t.DeletedAt.IsNull(), t.Did.IsNotNull(), tTriad.FirstActiveTime.IsNotNull(), tTriad.FirstActiveTime.Neq(time.Time{})).Count()
		onlineCount, _ = tTriad.WithContext(context.Background()).LeftJoin(t, tTriad.Did.EqCol(t.Did)).Where(tTriad.TenantId.Eq(tenantId), tTriad.UseType.Eq(0), tTriad.Status.Eq(1), t.OnlineStatus.Eq(1), t.DeletedAt.IsNull(), t.Did.IsNotNull(), tTriad.FirstActiveTime.IsNotNull(), tTriad.FirstActiveTime.Neq(time.Time{})).Count()
	} else {
		if req.Query.DeveloperTenantIds != nil && len(req.Query.DeveloperTenantIds) > 0 {
			var tenantIds = req.Query.DeveloperTenantIds
			totalCount, _ = tTriad.WithContext(context.Background()).Where(tTriad.TenantId.In(tenantIds...), tTriad.UseType.Eq(0), tTriad.FirstActiveTime.IsNotNull(), tTriad.FirstActiveTime.Neq(time.Time{})).Count()
			userCount, _ = tTriad.WithContext(context.Background()).LeftJoin(t, tTriad.Did.EqCol(t.Did)).Where(tTriad.TenantId.In(tenantIds...), tTriad.UseType.Eq(0), tTriad.Status.Eq(1), t.DeletedAt.IsNull(), t.Did.IsNotNull(), tTriad.FirstActiveTime.IsNotNull(), tTriad.FirstActiveTime.Neq(time.Time{})).Count()
			onlineCount, _ = tTriad.WithContext(context.Background()).LeftJoin(t, tTriad.Did.EqCol(t.Did)).Where(tTriad.TenantId.In(tenantIds...), tTriad.UseType.Eq(0), tTriad.Status.Eq(1), t.OnlineStatus.Eq(1), t.DeletedAt.IsNull(), t.Did.IsNotNull(), tTriad.FirstActiveTime.IsNotNull(), tTriad.FirstActiveTime.Neq(time.Time{})).Count()
		} else {
			totalCount, _ = tTriad.WithContext(context.Background()).Where(tTriad.UseType.Eq(0), tTriad.FirstActiveTime.IsNotNull(), tTriad.FirstActiveTime.Neq(time.Time{}), tTriad.FirstActiveTime.IsNotNull(), tTriad.FirstActiveTime.Neq(time.Time{})).Count()
			userCount, _ = tTriad.WithContext(context.Background()).LeftJoin(t, tTriad.Did.EqCol(t.Did)).Where(tTriad.Status.Eq(1), tTriad.UseType.Eq(0), t.DeletedAt.IsNull(), t.Did.IsNotNull(), tTriad.FirstActiveTime.IsNotNull(), tTriad.FirstActiveTime.Neq(time.Time{})).Count()
			onlineCount, _ = tTriad.WithContext(context.Background()).LeftJoin(t, tTriad.Did.EqCol(t.Did)).Where(tTriad.Status.Eq(1), t.OnlineStatus.Eq(1), tTriad.UseType.Eq(0), t.DeletedAt.IsNull(), t.Did.IsNotNull(), tTriad.FirstActiveTime.IsNotNull(), tTriad.FirstActiveTime.Neq(time.Time{})).Count()
		}
	}

	return &proto.IotDeviceInfoCount{
		ActiveTotal: userCount,
		DeviceTotal: totalCount,
		OnlineTotal: onlineCount,
	}, nil
}

// 设备详细信息
func (s *IotDeviceInfoSvc) GetCountIotDeviceInfoDetails(req *proto.IotDeviceInfoFilter) (*proto.IotDeviceInfoDetails, error) {
	//t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
	//do := t.WithContext(context.Background())
	//// fixme 请检查条件和校验参数

	q := orm.Use(iotmodel.GetDB())
	t := q.TIotDeviceTriad
	tInfo := q.TIotDeviceInfo
	//dbObj, err := t.WithContext(context.Background()).Where(t.Did.Eq(req.Did)).First()
	//if err != nil {
	//	logger.Errorf("FindByIdIotDeviceInfo error : %s", err.Error())
	//	return nil, err
	//}

	devInfoObj, err := tInfo.WithContext(context.Background()).Where(tInfo.Did.Eq(req.Did)).First()
	if err != nil {
		devInfoObj = &model.TIotDeviceInfo{}
	}

	// 查询设备绑定给那个家庭
	//tHomeDevice := orm.Use(iotmodel.GetDB()).TIotDeviceHome
	//doHomeDevice := tHomeDevice.WithContext(context.Background())
	//deviceHome, _ := doHomeDevice.Where(t.Did.Eq(dbObj.Did)).First()
	//activeUser := ""
	//if deviceHome != nil {
	//	activeUser = ""  //active user
	//}

	// 查询设备是那个开发者的
	tDeviceTriad := orm.Use(iotmodel.GetDB()).TIotDeviceTriad
	doDeviceTriad := tDeviceTriad.WithContext(context.Background())
	deviceTriad, err := doDeviceTriad.Where(t.Did.Eq(req.Did)).First()
	if err != nil {
		logger.Errorf("GetCountIotDeviceInfoDetails redis error : %s", err.Error())
		return nil, err
	}
	belongUser := iotutil.ToString(deviceTriad.UserId)
	// 查询设备属性信息
	var deviceStatus map[string]interface{}
	cached.RedisStore.Get(cached.DEVICE_STATUS+req.Did, &deviceStatus)
	//iotconst.HKEY_DEV_DATA_PREFIX+id
	newDeviceStatus, redisErr := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+req.Did).Result()
	if redisErr != nil {
		logger.Errorf("GetCountIotDeviceInfoDetails redis error : %s", redisErr.Error())
		//return
	}
	deviceVersion := devInfoObj.DeviceVersion
	mcuVer := ""
	//优先redis缓存中fwVer版本号
	if newDeviceStatus != nil {
		if val, ok := newDeviceStatus["fwVer"]; ok && val != "" {
			deviceVersion = iotutil.ToString(val)
		}
		if val, ok := newDeviceStatus["mcuVer"]; ok && val != "" {
			mcuVer = iotutil.ToString(val)
		}
	}

	activeInfo := &proto.IotDeviceInfoActiveInfo{
		ActiveStatus:   devInfoObj.ActiveStatus,
		ActiveUser:     devInfoObj.ActiveUserName,
		ActiveApp:      devInfoObj.AppKey,
		ActiveLoc:      fmt.Sprintf("[%.6f, %.6f]", devInfoObj.Lat, devInfoObj.Lng),
		ActiveChannel:  "",              //devInfoObj.ActiveChannel
		ActiveTimeZone: "Asia/Shanghai", //TODO 考虑获取方法
		BelogOpenUser:  belongUser,
	}
	if devInfoObj.Id != 0 {
		activeInfo.ActivatedTime = timestamppb.New(devInfoObj.ActivatedTime)
		activeInfo.LastActivatedTime = timestamppb.New(devInfoObj.LastActivatedTime)
		activeInfo.UpdateTime = timestamppb.New(devInfoObj.UpdatedAt)
	}
	activeInfo.OnlineStatus = toOnlineStatusName(devInfoObj.OnlineStatus)
	activeInfo.ActiveStatus = toActiveStatusName(devInfoObj.ActiveStatus)
	//统计数据
	return &proto.IotDeviceInfoDetails{
		ActiveInfo: activeInfo,
		DeviceInfo: &proto.IotDeviceInfoBasicInfo{
			Id:                 deviceTriad.Id,
			Did:                deviceTriad.Did,
			DeviceName:         devInfoObj.DeviceName,
			ProductId:          deviceTriad.ProductId,
			ProductKey:         deviceTriad.ProductKey,
			TenantId:           deviceTriad.TenantId,
			ProductName:        "",
			FirmwallKey:        "",
			FirmwallVersion:    deviceVersion, //?
			McuFirmwallKey:     "",
			McuFirmwallVersion: mcuVer,
			ModuleSN:           "",
			ModuleVersion:      devInfoObj.ModuleVersion,
			DeviceSN:           deviceTriad.SerialNumber,
		},
		DeviceStatus: newDeviceStatus,
	}, nil
}

func toOnlineStatusName(onlineStatus int32) string {
	if onlineStatus == 1 {
		return "在线"
	}
	return "离线"
}

func toActiveStatusName(activeStatus string) string {
	if activeStatus == "1" {
		return "已激活"
	}
	return "未激活"
}

// 通用设备信息
func (s *IotDeviceInfoSvc) CurrentDeviceInfo(req *proto.CurrentDeviceInfoFilter) (*proto.CurrentDeviceInfoResponse, error) {
	var q = orm.Use(iotmodel.GetDB())
	tIotDeviceHome := q.TIotDeviceHome
	tIotDeviceInfo := q.TIotDeviceInfo
	do := tIotDeviceInfo.WithContext(context.Background()).LeftJoin(tIotDeviceHome, tIotDeviceHome.DeviceId.EqCol(tIotDeviceInfo.Did), tIotDeviceHome.DeletedAt.IsNull())
	do = do.Where(tIotDeviceInfo.Did.Eq(req.DevId), tIotDeviceHome.DeletedAt.IsNull())
	var deviceInfoData *convert.TIotDeviceInfoData
	err := do.Select(tIotDeviceInfo.ALL, tIotDeviceHome.RoomId, tIotDeviceHome.CustomName, tIotDeviceHome.Secrtkey, tIotDeviceHome.HomeId).Scan(&deviceInfoData)
	if err != nil {
		logger.Errorf("GetListIotDeviceHome error : %s", err.Error())
		return nil, err
	}
	if deviceInfoData == nil {
		logger.Errorf("GetListIotDeviceHome error")
		return nil, errors.New("GetListIotDeviceHome error")
	}
	devInfo := convert.IotDevInfo_db2pbNew(deviceInfoData)
	// 查询设备属性信息
	var deviceStatus map[string]interface{}
	cached.RedisStore.Get(cached.DEVICE_STATUS+req.DevId, &deviceStatus)
	DeviceStatus, redisErr := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+req.DevId).Result()
	if redisErr != nil {
		logger.Errorf("FindByIdIotDeviceInfo redis error : %s", redisErr.Error())
	}
	newDeviceStatus := iotutil.MapStringToInterface(DeviceStatus)
	newDeviceStatusResult, _ := json.Marshal(newDeviceStatus)

	var MqttServer string
	if len(config.Global.AppMQTT.Addrs) > 0 {
		MqttServer = config.Global.AppMQTT.Addrs[0]
	}

	accessToken, _ := iotutil.AES_CBC_EncryptBase64([]byte(iotutil.ToString(req.DevId)), []byte(req.DevSecret))
	deviceInfo := proto.CurrentIotDeviceInfo{
		AccessToken:   accessToken,
		Batch:         iotutil.ToString(devInfo.BatchId),
		DevId:         iotutil.ToString(req.DevId),
		HomeId:        iotutil.ToString(deviceInfoData.HomeId),
		ProductId:     devInfo.ProductId,
		Name:          devInfo.DeviceName,
		RoomId:        deviceInfoData.RoomId,
		RoomName:      devInfo.RoomName,
		Secretkey:     devInfo.Salt,
		MqttServer:    MqttServer,
		Ssid:          "",
		State:         devInfo.OnlineStatus,
		Switch:        0,
		DeviceStatus:  newDeviceStatusResult,
		DeviceVersion: devInfo.DeviceVersion,
	}

	//优先redis缓存中fwVer版本号
	if newDeviceStatusResult != nil {
		if val, ok := newDeviceStatus["fwVer"]; ok && val != "" {
			deviceInfo.DeviceVersion = iotutil.ToString(val)
		}
		if val, ok := newDeviceStatus["mcuVer"]; ok && val != "" {
			deviceInfo.DeviceMcuVersion = iotutil.ToString(val)
		}
	}

	//获取用户对属性值的自定义显示，不需要返回错误
	deviceInfo.FuncDescMap, _ = s.GetFunctionSetMap(req.DevId)
	CurrentDeviceInfoResponse := proto.CurrentDeviceInfoResponse{
		Data: &deviceInfo,
	}
	return &CurrentDeviceInfoResponse, nil
}

func (s *IotDeviceInfoSvc) DeviceInfoListByDevIds(req *proto.DeviceInfoListByDevIdsFilter) ([]*proto.IotDeviceInfo, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TIotDeviceInfo
	do := t.WithContext(context.Background())

	if len(req.DevIds) > 0 { //字符串
		do = do.Where(t.Did.In(req.DevIds...))
	}

	var list []*model.TIotDeviceInfo
	list, err = do.Find()
	if err != nil {
		logger.Errorf("GetListIotDeviceInfo error : %s", err.Error())
		return nil, 0, err
	}
	result := make([]*proto.IotDeviceInfo, len(list))
	for i, v := range list {
		result[i] = convert.IotDeviceInfo_db2pb(v)
	}
	if len(result) > 0 {
		result[0].FuncDescMap, _ = s.GetFunctionSetMap(result[0].Did)
	}
	return result, 0, nil
}

func (s *IotDeviceInfoSvc) GetDeviceFunctionSetList(req *proto.IotDeviceFunctionSet) ([]*proto.IotDeviceFunctionSet, int64, error) {
	if req.DeviceId == "" {
		return nil, 0, errors.New("设备Id不能为空")
	}
	var err error
	t := orm.Use(iotmodel.GetDB()).TIotDeviceFunctionSet
	do := t.WithContext(context.Background())

	if len(req.DeviceId) > 0 { //字符串
		do = do.Where(t.DeviceId.Eq(req.DeviceId))
	}

	var list []*model.TIotDeviceFunctionSet
	list, err = do.Find()
	if err != nil {
		logger.Errorf("GetListIotDeviceInfo error : %s", err.Error())
		return nil, 0, err
	}
	result := make([]*proto.IotDeviceFunctionSet, len(list))
	for i, v := range list {
		result[i] = convert.IotDeviceFunctionSet_db2pb(v)
	}
	return result, 0, nil
}

func (s *IotDeviceInfoSvc) SaveDeviceFunctionSet(req *proto.IotDeviceFunctionSet) error {
	if req.DeviceId == "" {
		return errors.New("设备Id不能为空")
	}
	t := orm.Use(iotmodel.GetDB()).TIotDeviceFunctionSet
	do := t.WithContext(context.Background())
	saveObj := convert.IotDeviceFunctionSet_pb2db(req)
	saveObj.Id = iotutil.GetNextSeqInt64()

	if req.CustomType == iotconst.FUNCTION_CUSTOM_PROPERTY_VALUE_SET {
		obj, err := do.Where(t.DeviceId.Eq(req.DeviceId), t.FuncKey.Eq(req.FuncKey), t.FuncValue.Eq(req.FuncValue), t.CustomType.Eq(req.CustomType)).Find()
		if err != nil {
			return err
		}
		if len(obj) > 0 {
			saveObj.Id = obj[0].Id
		}
	} else if req.CustomType == iotconst.FUNCTION_CUSTOM_PROPERTY_SET {
		obj, err := do.Where(t.DeviceId.Eq(req.DeviceId), t.FuncKey.Eq(req.FuncKey), t.CustomType.Eq(req.CustomType)).Find()
		if err != nil {
			return err
		}
		if len(obj) > 0 {
			saveObj.Id = obj[0].Id
		}
	} else {
		return errors.New("CustomType不正确")
	}
	err := do.Save(saveObj)
	if err != nil {
		return err
	}
	return nil
}

func (s *IotDeviceInfoSvc) GetFunctionSetMap(deviceId string) (map[string]string, error) {
	list, _, err := s.GetDeviceFunctionSetList(&proto.IotDeviceFunctionSet{DeviceId: deviceId})
	if err != nil {
		return nil, err
	}
	resMap := make(map[string]string)
	for _, l := range list {
		if l.CustomType == iotconst.FUNCTION_CUSTOM_PROPERTY_SET {
			resMap[fmt.Sprintf("%s_%s", l.ProductKey, l.FuncKey)] = l.CustomDesc
		} else if l.CustomType == iotconst.FUNCTION_CUSTOM_PROPERTY_VALUE_SET {
			resMap[fmt.Sprintf("%s_%s_%v", l.ProductKey, l.FuncKey, l.FuncValue)] = l.CustomDesc
		}
	}
	return resMap, nil
}

func (s *IotDeviceInfoSvc) SaveDeviceFunctionBatchSet(req *proto.IotDeviceFunctionBatchSet) error {
	if req.DeviceId == "" {
		return errors.New("设备Id不能为空")
	}
	q := orm.Use(iotmodel.GetDB())

	err := q.Transaction(func(tx *orm.Query) error {
		t := tx.TIotDeviceFunctionSet
		do := t.WithContext(context.Background())
		saveList := make([]*model.TIotDeviceFunctionSet, 0)
		if req.CustomType == iotconst.FUNCTION_CUSTOM_PROPERTY_VALUE_SET {
			objs, err := do.Where(t.DeviceId.Eq(req.DeviceId), t.FuncKey.Eq(req.FuncKey), t.CustomType.Eq(req.CustomType)).Find()
			if err != nil {
				return err
			}
			saveIds := make([]int64, 0)
			for _, data := range req.Datas {
				has := false
				for i, obj := range objs {
					if obj.FuncValue == data.FuncValue && obj.FuncKey == req.FuncKey {
						//历史更新
						objs[i].CustomDesc = data.CustomDesc
						newObj := convert.IotDeviceFunctionSet_db2NewDb(objs[i])
						newObj.FuncValue = data.FuncValue
						newObj.CustomDesc = data.CustomDesc
						saveIds = append(saveIds, newObj.Id)
						saveList = append(saveList, newObj)
						has = true
						break
					}
				}
				//TODO objs数据库有，但是request里面没有的数据，是否需要考虑删除
				if !has {
					//不存在的新增
					newObj := convert.IotDeviceFunctionBatchSet_pb2db(req)
					newObj.FuncValue = data.FuncValue
					newObj.CustomDesc = data.CustomDesc
					newObj.Id = iotutil.GetNextSeqInt64()
					saveIds = append(saveIds, newObj.Id)
					saveList = append(saveList, newObj)
				}
			}
			_, err = do.Where(t.DeviceId.Eq(req.DeviceId), t.FuncKey.Eq(req.FuncKey), t.CustomType.Eq(req.CustomType), t.Id.NotIn(saveIds...)).Delete()
			if err != nil {
				return err
			}
		} else if req.CustomType == iotconst.FUNCTION_CUSTOM_PROPERTY_SET {
			saveObj := convert.IotDeviceFunctionBatchSet_pb2db(req)
			saveObj.Id = iotutil.GetNextSeqInt64()
			saveObj.FuncValue = req.Datas[0].FuncValue
			saveObj.CustomDesc = req.Datas[0].CustomDesc
			obj, err := do.Where(t.DeviceId.Eq(req.DeviceId), t.FuncKey.Eq(req.FuncKey), t.CustomType.Eq(req.CustomType)).Find()
			if err != nil {
				return err
			}
			if len(obj) > 0 {
				saveObj.Id = obj[0].Id
			}
			saveList = append(saveList, saveObj)
		} else {
			return errors.New("CustomType不正确")
		}
		return do.Save(saveList...)
	})
	if err != nil {
		return err
	}
	return nil
}
