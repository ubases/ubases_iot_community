package service

import (
	iotmodel "cloud_platform/iot_model"
	appBuildOrm "cloud_platform/iot_model/db_app_oem/orm"
	deviceOrm "cloud_platform/iot_model/db_device/orm"
	openSystemOrm "cloud_platform/iot_model/db_open_system/orm"
	"cloud_platform/iot_model/db_statistics/model"
	statisticsOrm "cloud_platform/iot_model/db_statistics/orm"
	"cloud_platform/iot_statistics_service/config"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"sort"
	"strconv"
	"time"

	"gorm.io/gorm"

	"go-micro.dev/v4/logger"

	"gorm.io/gen/field"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type StatisticsSvc struct {
	Ctx context.Context
}

func (s *StatisticsSvc) GetDeveloperList(request *proto.DeveloperStatListRequest) (*proto.DeveloperStatListResponse, error) {
	list, err := GetDeveloperList(request)
	if err != nil {
		return nil, err
	}
	develperObjList, err := GetListPmDevelopData(list)
	if err != nil {
		return nil, err
	}
	rsp := proto.DeveloperStatListResponse{
		Code:    200,
		Message: "success",
		Total:   int64(len(list)),
	}
	for _, v0 := range list {
		obj := proto.DeveloperStat{UserId: iotutil.ToString(v0.Id), UserName: v0.UserName, RegisterTime: timestamppb.New(v0.CreatedAt), Online: 2, LoginAddr: "-"}
		for _, v1 := range develperObjList {
			if v0.TenantId == v1.TenantId {
				obj.Quantity = v1.DeviceSum
				obj.ActiveDeviceTotal = v1.DeviceActiveSum
				obj.AppTotal = v1.AppSum
				break
			}
		}
		rsp.Data = append(rsp.Data, &obj)
	}

	return &rsp, nil
}

func (s *StatisticsSvc) GetDeveloperDetail(request *proto.DeveloperDetailFilter) (*proto.DeveloperDetailResponse, error) {
	info, err := GetDeveloperDetail(request)
	if err != nil {
		return nil, err
	}
	rsp := proto.DeveloperDetailResponse{
		UserId:      iotutil.ToString(info.Id),
		UserName:    info.UserName,
		Account:     info.UserName,
		CompanyName: info.CompanyName,
		RoleName:    info.RoleName,
		//ActiveDeviceTotal: developData.DeviceActiveSum,
		//AppTotal:          developData.AppSum,
	}

	developData, err := GetDevelopActiveDeviceSum(info.TenantId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if developData != nil {
		rsp.ActiveDeviceTotal = developData.DeviceActiveSum
		rsp.AppTotal = developData.AppSum
	}
	appList, err := GetDevelopAppList(info.TenantId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	for _, v := range appList {
		obj := proto.AppStatistics{
			AppId:     iotutil.ToString(v.Id),
			AppKey:    v.AppKey,
			AppName:   v.Name,
			DevStatus: iotutil.ToString(v.Status),
			Version:   v.Version,
			VerTotal:  int64(v.VerTotal),
		}
		rsp.AppList = append(rsp.AppList, &obj)
	}
	return &rsp, nil
}

func (s *StatisticsSvc) GetDeveloperStatistics(null *proto.NULL) (*proto.DeveloperStatisticsResponse, error) {
	openSystem, ok := config.DBMap["iot_open_system"]
	if !ok {
		return nil, errors.New("iot_open_system数据库未初始化")
	}
	//SELECT COUNT(*) FROM t_open_user
	var total int64
	topenUser := openSystemOrm.Use(openSystem).TOpenUser
	err := topenUser.WithContext(context.Background()).Select(topenUser.Id.Count().IfNull(0).As("total")).Scan(&total)
	if err != nil {
		return nil, err
	}
	ret := &proto.DeveloperStatisticsResponse{
		Total:       total,
		OnlineTotal: 0,
	}

	//fixme 在线数据未保存到数据库
	//todo 因为token的原因，判断是否在线需要考虑具体方案

	return ret, nil
}

func (s *StatisticsSvc) GetAppDataDetail(request *proto.AppDataDetailFilter) (*proto.AppDataDetailResponse, error) {
	appId := iotutil.ToInt64(request.AppId)
	if appId == 0 {
		return nil, errors.New("appId error.")
	}
	t := statisticsOrm.Use(iotmodel.GetDB()).TPmAppData
	dbObj, err := t.WithContext(context.Background()).Where(t.AppId.Eq(appId)).First()
	if err != nil {
		logger.Errorf("FindPmAppData error : %s", err.Error())
		return nil, err
	}

	//SELECT `version`,COUNT(*), MAX(updated_at) FROM t_oem_app_build_record
	//WHERE app_id = 181369044765605888 GROUP BY `version`
	appBuild, ok := config.DBMap["iot_app_build"]
	if !ok {
		return nil, errors.New("iot_app_build数据库未初始化")
	}
	//查找开发app数量
	var appList []AppVersionRecord
	tOemAppBuild := appBuildOrm.Use(appBuild).TOemAppBuildRecord
	err = tOemAppBuild.WithContext(context.Background()).Select(tOemAppBuild.Version,
		tOemAppBuild.Version.Count().IfNull(0).As("total"), tOemAppBuild.UpdatedAt.Max().As("updated_at")).
		Where(tOemAppBuild.AppId.Eq(appId)).Group(tOemAppBuild.Version).Scan(&appList)
	if err != nil {
		return nil, err
	}

	mapVersionStatus := make(map[string]int)
	var appVersionStatusList []AppVersionStatusRecord
	toemappversionrecord := appBuildOrm.Use(appBuild).TOemAppVersionRecord
	err = toemappversionrecord.WithContext(context.Background()).Select(toemappversionrecord.Version, toemappversionrecord.Status).
		Where(toemappversionrecord.AppId.Eq(appId)).Scan(&appVersionStatusList)
	//if err != nil {
	//	return nil, err
	//}
	for _, v := range appVersionStatusList {
		mapVersionStatus[v.Version] = v.Status
	}

	resp := proto.AppDataDetailResponse{
		Account:           dbObj.DevAccount,
		AcitveUserTotal:   dbObj.ActiveUserSum,
		AppName:           dbObj.AppName,
		AppType:           "oem",
		RegisterUserTotal: dbObj.RegisterUserSum,
	}
	resp.VersionList = make([]*proto.AppVersion, 0, len(appList))
	for _, v := range appList {
		devStatus, _ := mapVersionStatus[v.Version]
		if devStatus == 0 {
			devStatus = 1
		}
		obj := proto.AppVersion{
			AppVersion:  v.Version,
			BuildNumber: v.Total,
			DevStatus:   int64(devStatus),
			LastOptTime: v.UpdatedAt.Unix(),
			LastOptUser: dbObj.DevAccount,
		}
		resp.VersionList = append(resp.VersionList, &obj)
	}

	sort.Slice(resp.VersionList, func(i, j int) bool {
		if ret, err1 := iotutil.VerCompare(resp.VersionList[i].AppVersion, resp.VersionList[j].AppVersion); err1 == nil {
			if ret == -1 {
				return false
			}
		}
		return true
	})

	return &resp, nil
}

func GetDeveloperList(req *proto.DeveloperStatListRequest) ([]DeveloperInfo, error) {
	openSystem, ok := config.DBMap["iot_open_system"]
	if !ok {
		return nil, errors.New("iot_open_system数据库未初始化")
	}
	//SELECT u.id,u.user_name,c.tenant_id,u.created_at
	//FROM t_open_user u ,t_open_company c WHERE u.id = c.user_id
	var developerInfoList []DeveloperInfo
	topenUser := openSystemOrm.Use(openSystem).TOpenUser
	topenCompany := openSystemOrm.Use(openSystem).TOpenCompany
	do := topenUser.WithContext(context.Background()).Select(topenUser.Id, topenUser.UserName, topenUser.CreatedAt, topenCompany.TenantId)
	do = do.Join(topenCompany, topenUser.Id.EqCol(topenCompany.UserId))
	if req.Query != nil {
		if req.Query.UserName != "" {
			do = do.Where(topenUser.UserName.Like("%" + req.Query.UserName + "%"))
		}
		if req.Query.StartTime != nil && !req.Query.StartTime.AsTime().IsZero() {
			do = do.Where(topenUser.CreatedAt.Gte(req.Query.StartTime.AsTime()))
		}
		if req.Query.EndTime != nil && !req.Query.EndTime.AsTime().IsZero() {
			do = do.Where(topenUser.CreatedAt.Lte(req.Query.EndTime.AsTime()))
		}
	}
	err := do.Order(topenUser.CreatedAt).Scan(&developerInfoList)
	return developerInfoList, err
}

type DeveloperInfo struct {
	Id        int64
	UserName  string
	CreatedAt time.Time
	TenantId  string
}

func GetListPmDevelopData(infolist []DeveloperInfo) ([]*model.TPmDevelopData, error) {
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return nil, errors.New("iot_statistics数据库未初始化")
	}
	tenantIdList := make([]string, 0, len(infolist))
	for _, v := range infolist {
		tenantIdList = append(tenantIdList, v.TenantId)
	}
	tdeveloper := statisticsOrm.Use(statDB).TPmDevelopData
	do := tdeveloper.WithContext(context.Background())
	do = do.Where(tdeveloper.TenantId.In(tenantIdList...))
	return do.Find()
}

type DeveloperDetailInfo struct {
	Id          int64
	UserName    string
	CompanyName string
	TenantId    string
	RoleName    string
}

func GetDeveloperDetail(request *proto.DeveloperDetailFilter) (*DeveloperDetailInfo, error) {
	openSystem, ok := config.DBMap["iot_open_system"]
	if !ok {
		return nil, errors.New("iot_open_system数据库未初始化")
	}
	userId, err := strconv.ParseInt(request.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	//-- 主账户信息
	//SELECT u.id,u.user_name,c.name,c.tenant_id FROM t_open_user u,t_open_company  c
	//WHERE u.id = 112344250921877504 AND u.id = c.user_id
	var info DeveloperDetailInfo
	topenUser := openSystemOrm.Use(openSystem).TOpenUser
	topenCompany := openSystemOrm.Use(openSystem).TOpenCompany
	do := topenUser.WithContext(context.Background()).Select(topenUser.Id, topenUser.UserName, topenCompany.Name.As(
		"company_name"), topenCompany.TenantId)
	do = do.Join(topenCompany, topenUser.Id.EqCol(topenCompany.UserId))
	do = do.Where(topenUser.Id.Eq(userId))
	err = do.Scan(&info)
	if err != nil {
		return nil, err
	}
	//查询角色
	//SELECT r.name FROM t_open_role r JOIN t_open_casbin_rule  c ON r.id = c.v1
	//WHERE c.ptype = 'g' AND c.v0 = '1348267746932129792' AND c.v2= 'LhchpV'
	var roleName string
	topenrole := openSystemOrm.Use(openSystem).TOpenRole
	topencasbinrule := openSystemOrm.Use(openSystem).TOpenCasbinRule
	err = topenrole.WithContext(context.Background()).Select(topenrole.Name.As("role_name")).Join(topencasbinrule,
		topenrole.Id.EqCol(topencasbinrule.V1)).Where(topencasbinrule.Ptype.Eq("g"),
		topencasbinrule.V0.Eq(request.UserId), topencasbinrule.V2.Eq(info.TenantId)).Scan(&roleName)
	//info := DeveloperDetailInfo{infoA, RoleName}
	info.RoleName = roleName
	return &info, nil
}

func GetDevelopActiveDeviceSum(tenandId string) (*model.TPmDevelopData, error) {
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return nil, errors.New("iot_statistics数据库未初始化")
	}
	//-- 已激活数量
	//SELECT device_active_sum FROM t_pm_develop_data WHERE tenant_id ='ioqp4r'
	tdeveloper := statisticsOrm.Use(statDB).TPmDevelopData
	do := tdeveloper.WithContext(context.Background())
	do = do.Where(tdeveloper.TenantId.Eq(tenandId))
	return do.First()
}

func GetDevelopAppList(tenandId string) ([]AppInfoList, error) {
	//-- 开发者app列表
	//SELECT a.id,a.app_key,a.`name`,'oem',a.status,a.version,IFNULL(r.verTotal,0) AS verTotal FROM t_oem_app  a
	//LEFT JOIN (SELECT app_id,COUNT(distinct version) AS verTotal FROM t_oem_app_build_record GROUP BY app_id)  r
	//ON (a.id = r.app_id)
	//WHERE a.tenant_id = 'ioqp4r'
	appBuild, ok := config.DBMap["iot_app_build"]
	if !ok {
		return nil, errors.New("iot_app_build数据库未初始化")
	}
	//查找开发app数量
	var appList []AppInfoList
	tOemApp := appBuildOrm.Use(appBuild).TOemApp
	tOemAppBuild := appBuildOrm.Use(appBuild).TOemAppBuildRecord
	r := tOemAppBuild.As("r")
	subQuery := tOemAppBuild.WithContext(context.Background()).Select(tOemAppBuild.AppId,
		tOemAppBuild.Version.Distinct().Count().IfNull(0).As("ver_total")).Group(tOemAppBuild.AppId)
	do := tOemApp.WithContext(context.Background()).Select(tOemApp.Id, tOemApp.AppKey, tOemApp.Name, tOemApp.Status,
		tOemApp.Version, field.NewInt32("r", "ver_total").IfNull(0).As("ver_total")).LeftJoin(subQuery.As("r"), tOemApp.Id.EqCol(r.AppId))
	err := do.Where(tOemApp.TenantId.Eq(tenandId)).Scan(&appList)
	if err != nil {
		return nil, err
	}
	return appList, nil
}

func (s *StatisticsSvc) GetDeviceTotalStatistics(null *proto.NULL) (*proto.DeviceStatisticsResponse, error) {
	deviceDB, ok := config.DBMap["iot_device"]
	if !ok {
		return nil, errors.New("iot_device数据库未初始化")
	}
	ret := &proto.DeviceStatisticsResponse{ActiveTotal: 0, OnlineTotal: 0}
	var total int64
	t := deviceOrm.Use(deviceDB).TIotDeviceInfo
	err := t.WithContext(context.Background()).Select(t.Id.Count().IfNull(0).As("total")).Where(t.UseType.Eq(0)).Scan(&total)
	if err == nil {
		ret.ActiveTotal = total
	}
	err = t.WithContext(context.Background()).Select(t.Id.Count().IfNull(0).As("total")).Where(t.UseType.Eq(0), t.OnlineStatus.Eq(1)).Scan(&total)
	if err == nil {
		ret.OnlineTotal = total
	}
	return ret, nil
}

type AppInfoList struct {
	Id       int64
	AppKey   string
	Name     string
	Status   int
	Version  string
	VerTotal int
}

type AppVersionRecord struct {
	Version   string
	Total     int64
	UpdatedAt time.Time
}

type AppVersionStatusRecord struct {
	Version string
	Status  int
}
