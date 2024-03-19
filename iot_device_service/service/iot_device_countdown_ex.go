package service

import (
	"cloud_platform/iot_common/ioterrs"
	"context"
	"errors"
	goerrors "go-micro.dev/v4/errors"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	"cloud_platform/iot_device_service/convert"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type IotDeviceCountdownSvcEx struct {
	Ctx context.Context
}

func (s *IotDeviceCountdownSvcEx) StartIotDeviceCountdownJob(req *proto.IotDeviceCountdownJobReq) error {
	infoReq := &proto.IotDeviceCountdownFilter{
		Id: req.Id,
	}
	jobInfo, err := s.FindByIdIotDeviceCountdown(infoReq)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBDeviceGet)
	}
	if jobInfo.Enabled == 1 {
		return goerrors.New("", "倒计时任务已启动", ioterrs.ErrCountDownAlreadyStarted)
	}
	// 调用job服务rpc接口
	reqJob := &proto.JobReq{
		Id: req.Id,
	}
	jobSvc := IotJobSvc{Ctx: context.Background()}
	err = jobSvc.StartJob(s.Ctx, reqJob)
	if err != nil {
		return err
	}
	repJob := &proto.IotDeviceCountdown{
		Id:      req.Id,
		Enabled: 1,
	}
	_, err = s.UpdateIotDeviceCountdown(repJob)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBCountDownUpdate)
	}
	return nil
}

func (s *IotDeviceCountdownSvcEx) StopIotDeviceCountdownJob(req *proto.IotDeviceCountdownJobReq) error {
	infoReq := &proto.IotDeviceCountdownFilter{
		Id: req.Id,
	}
	jobInfo, err := s.FindByIdIotDeviceCountdown(infoReq)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBDeviceGet)
	}
	if jobInfo.Enabled == 2 {
		return goerrors.New("", "倒计时任务已停止", ioterrs.ErrCountDownAlreadyStopped)
	}
	// 调用job服务rpc接口
	reqJob := &proto.JobReq{
		Id: req.Id,
	}
	jobSvc := IotJobSvc{Ctx: context.Background()}
	err = jobSvc.StopJob(s.Ctx, reqJob)
	if err != nil {
		return err
	}
	repJob := &proto.IotDeviceCountdown{
		Id:      req.Id,
		Enabled: 2,
	}
	_, err = s.UpdateIotDeviceCountdown(repJob)
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrDBCountDownUpdate)
	}
	return nil
}

// UpdateIotDeviceCountdown 根据主键更新IotDeviceCountdown
func (s *IotDeviceCountdownSvcEx) UpdateIotDeviceCountdown(req *proto.IotDeviceCountdown) (*proto.IotDeviceCountdown, error) {
	t := orm.Use(iotmodel.GetDB()).TIotDeviceCountdown
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.Id != 0 { //整数
		updateField = append(updateField, t.Id)
	}
	if req.Hour != "" { //字符串
		updateField = append(updateField, t.Hour)
	}
	if req.Minute != "" { //字符串
		updateField = append(updateField, t.Minute)
	}
	if req.Remark != "" { //字符串
		updateField = append(updateField, t.Remark)
	}
	if req.FuncKey != "" { //字符串
		updateField = append(updateField, t.FuncKey)
	}
	if req.FuncValue != "" { //字符串
		updateField = append(updateField, t.FuncValue)
	}
	if req.UserId != 0 { //整数
		updateField = append(updateField, t.UserId)
	}
	if req.DeviceId != "" { //字符串
		updateField = append(updateField, t.DeviceId)
	}
	if req.Enabled != 0 { //整数
		updateField = append(updateField, t.Enabled)
	}
	if req.TotalSecond != 0 { //整数
		updateField = append(updateField, t.TotalSecond)
	}
	if req.TaskId != "" { //字符串
		updateField = append(updateField, t.TaskId)
	}
	if req.Cron != "" { //字符串
		updateField = append(updateField, t.Cron)
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
		logger.Error("UpdateIotDeviceCountdown error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.IotDeviceCountdown_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateIotDeviceCountdown error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// FindByIdIotDeviceCountdown 根据数据库表主键查找IotDeviceCountdown
func (s *IotDeviceCountdownSvcEx) FindByIdIotDeviceCountdown(req *proto.IotDeviceCountdownFilter) (*proto.IotDeviceCountdown, error) {
	t := orm.Use(iotmodel.GetDB()).TIotDeviceCountdown
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdIotDeviceCountdown error : %s", err.Error())
		return nil, err
	}
	res := convert.IotDeviceCountdown_db2pb(dbObj)
	return res, err
}
