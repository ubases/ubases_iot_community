// Code generated by sgen.exe,2022-04-18 21:26:14. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	_const "cloud_platform/iot_app_api_service/controls/user/const"
	"cloud_platform/iot_app_user_service/cached"
	"cloud_platform/iot_app_user_service/rpc/rpcClient"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"go-micro.dev/v4/logger"
	"gorm.io/gen/field"

	"cloud_platform/iot_app_user_service/convert"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_app/model"
	"cloud_platform/iot_model/db_app/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type UcHomeSvc struct {
	Ctx context.Context
}

// 创建UcHome
func (s *UcHomeSvc) CreateUcHome(req *proto.UcHome) (*proto.UcHome, error) {
	// fixme 请在这里校验参数
	t := orm.Use(iotmodel.GetDB()).TUcHome
	do := t.WithContext(context.Background())
	dbObj := convert.UcHome_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("CreateUcHome error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除UcHome
func (s *UcHomeSvc) DeleteUcHome(req *proto.UcHome) (*proto.UcHome, error) {
	t := orm.Use(iotmodel.GetDB()).TUcHome
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.Name != "" { //字符串
		do = do.Where(t.Name.Eq(req.Name))
	}
	if req.IconUrl != "" { //字符串
		do = do.Where(t.IconUrl.Eq(req.IconUrl))
	}
	if req.Sid != 0 { //整数
		do = do.Where(t.Sid.Eq(req.Sid))
	}
	if req.Country != "" { //字符串
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
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteUcHome error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除UcHome
func (s *UcHomeSvc) DeleteByIdUcHome(req *proto.UcHome) (*proto.UcHome, error) {
	t := orm.Use(iotmodel.GetDB()).TUcHome
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdUcHome error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除UcHome
func (s *UcHomeSvc) DeleteByIdsUcHome(req *proto.UcHomeBatchDeleteRequest) (*proto.UcHomeBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TUcHome
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsUcHome error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新UcHome
func (s *UcHomeSvc) UpdateUcHome(req *proto.UcHome) (*proto.UcHome, error) {
	t := orm.Use(iotmodel.GetDB()).TUcHome
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.Name != "" { //字符串
		updateField = append(updateField, t.Name)
	}
	if req.IconUrl != "" { //字符串
		updateField = append(updateField, t.IconUrl)
	}
	if req.Sid != 0 { //整数
		updateField = append(updateField, t.Sid)
	}
	if req.Country != "" { //字符串
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
	if req.Address != "" { //字符串
		updateField = append(updateField, t.Address)
	}
	if req.Lng != 0 { //整数
		updateField = append(updateField, t.Lng)
	}
	if req.Lat != 0 { //整数
		updateField = append(updateField, t.Lat)
	}
	if req.CoordType != 0 { //字符串
		updateField = append(updateField, t.CoordType)
	}
	if req.CreatedBy != 0 { //整数
		updateField = append(updateField, t.CreatedBy)
	}
	if req.UpdatedBy != 0 { //整数
		updateField = append(updateField, t.UpdatedBy)
	}
	if req.CountryId != 0 { //整数
		updateField = append(updateField, t.CountryId)
		updateField = append(updateField, t.ProvinceId)
		updateField = append(updateField, t.CityId)
	}
	if req.ProvinceId != 0 { //整数
		updateField = append(updateField, t.ProvinceId)
	}
	if req.CityId != 0 { //整数
		updateField = append(updateField, t.CityId)
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
		logger.Error("UpdateUcHome error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.UcHome_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateUcHome error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段UcHome
func (s *UcHomeSvc) UpdateAllUcHome(req *proto.UcHome) (*proto.UcHome, error) {
	t := orm.Use(iotmodel.GetDB()).TUcHome
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.Name)
	updateField = append(updateField, t.IconUrl)
	updateField = append(updateField, t.Sid)
	updateField = append(updateField, t.Lat)
	updateField = append(updateField, t.Lng)
	updateField = append(updateField, t.Country)
	updateField = append(updateField, t.Province)
	updateField = append(updateField, t.City)
	updateField = append(updateField, t.District)
	updateField = append(updateField, t.CreatedBy)
	updateField = append(updateField, t.UpdatedBy)
	updateField = append(updateField, t.CountryId)
	updateField = append(updateField, t.ProvinceId)
	updateField = append(updateField, t.CityId)
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
		logger.Error("UpdateAllUcHome error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.UcHome_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllUcHome error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *UcHomeSvc) UpdateFieldsUcHome(req *proto.UcHomeUpdateFieldsRequest) (*proto.UcHome, error) {
	t := orm.Use(iotmodel.GetDB()).TUcHome
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFields error : missing updateField")
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
		logger.Error("UpdateFieldsUcHome error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.UcHome_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsUcHome error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找UcHome
func (s *UcHomeSvc) FindUcHome(req *proto.UcHomeFilter) (*proto.UcHome, error) {
	t := orm.Use(iotmodel.GetDB()).TUcHome
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.Name != "" { //字符串
		do = do.Where(t.Name.Like("%" + req.Name + "%"))
	}
	if req.IconUrl != "" { //字符串
		do = do.Where(t.IconUrl.Like("%" + req.IconUrl + "%"))
	}
	if req.Sid != 0 { //整数
		do = do.Where(t.Sid.Eq(req.Sid))
	}
	if req.Country != "" { //字符串
		do = do.Where(t.Country.Like("%" + req.Country + "%"))
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
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindUcHome error : %s", err.Error())
		return nil, err
	}
	res := convert.UcHome_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找UcHome
func (s *UcHomeSvc) FindByIdUcHome(req *proto.UcHomeFilter) (*proto.UcHome, error) {
	t := orm.Use(iotmodel.GetDB()).TUcHome
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdUcHome error : %s", err.Error())
		return nil, err
	}
	res := convert.UcHome_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找UcHome,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *UcHomeSvc) GetListUcHome(req *proto.UcHomeListRequest) ([]*proto.UcHome, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TUcHome
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.Name != "" { //字符串
			do = do.Where(t.Name.Like("%" + query.Name + "%"))
		}
		if query.IconUrl != "" { //字符串
			do = do.Where(t.IconUrl.Like("%" + query.IconUrl + "%"))
		}
		if query.Sid != 0 { //整数
			do = do.Where(t.Sid.Eq(query.Sid))
		}
		if query.Country != "" { //字符串
			do = do.Where(t.Country.Like("%" + query.Country + "%"))
		}
		if query.Province != "" { //字符串
			do = do.Where(t.Province.Like("%" + query.Province + "%"))
		}
		if query.City != "" { //字符串
			do = do.Where(t.City.Like("%" + query.City + "%"))
		}
		if query.District != "" { //字符串
			do = do.Where(t.District.Like("%" + query.District + "%"))
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
		do = do.Order(t.CreatedAt.Desc())
	} else {
		if req.OrderDesc != "" {
			do = do.Order(orderCol.Desc())
		} else {
			do = do.Order(orderCol)
		}
	}

	var list []*model.TUcHome
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
		logger.Errorf("GetListUcHome error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.UcHome, len(list))
	for i, v := range list {
		result[i] = convert.UcHome_db2pb(v)
	}
	return result, total, nil
}

// 添加家庭
func (s *UcHomeSvc) AddHome(req *proto.UcHomeReq) error {
	var (
		lang, _     = CheckLang(s.Ctx)
		appKey, _   = CheckAppKey(s.Ctx)
		tenantId, _ = CheckTenantId(s.Ctx)
	)
	tran := orm.Use(iotmodel.GetDB())
	tran.Transaction(func(tx *orm.Query) error {
		//todo 最大家庭数量为100个待处理
		//fixme 请检查条件和校验参数
		dbObj := convert.UcHomeReq_pb2db(req)
		err := tx.TUcHome.WithContext(context.Background()).Create(dbObj)
		if err != nil {
			logger.Errorf("TUcHome create failed, error:%s", err.Error())
			return err
		}

		tUcHomeUser := &model.TUcHomeUser{
			Id:        iotutil.GetNextSeqInt64(),
			HomeId:    dbObj.Id,
			UserId:    req.UserId,
			RoleType:  iotutil.ToInt32(_const.RoleSuperAdministrator),
			BindTime:  timestamppb.Now().AsTime(),
			Shared:    iotutil.ToInt32(_const.NormalHome),
			CreatedAt: timestamppb.Now().AsTime(),
		}
		err = tx.TUcHomeUser.WithContext(context.Background()).Create(tUcHomeUser)
		if err != nil {
			logger.Errorf("TUcHomeUser create failed, error:%s", err.Error())
			return err
		}

		rooms := make([]*model.TUcHomeRoom, 0)
		defaultRooms := UcUserSvc{Ctx: s.Ctx}.GetDefaultRooms(lang, tenantId, appKey)
		for i, v := range req.RoomList { //添加多个房间
			var roomTemplateId int64 = 0
			for _, room := range defaultRooms {
				if room.DictLabel == v.Name && room.Listimg == v.Icon {
					roomTemplateId = room.DictCode
				}
			}
			rooms = append(rooms, &model.TUcHomeRoom{
				Id:             iotutil.GetNextSeqInt64(),
				HomeId:         dbObj.Id,
				RoomName:       v.Name,
				IconUrl:        v.Icon,
				Sort:           int32(i) + 1,
				RoomTemplateId: roomTemplateId,
				CreatedBy:      req.UserId,
			})
		}
		err = tx.TUcHomeRoom.WithContext(context.Background()).Create(rooms...)
		if err != nil {
			logger.Errorf("TUcHomeRoom create failed, error:%s", err.Error())
			return err
		}
		return nil
	})
	return nil
}

// 成员加入家庭
func (s *UcHomeSvc) JoinHome(req *proto.UcHomeDetailRequest) error {
	tUcHomeUserTable := orm.Use(iotmodel.GetDB()).TUcHomeUser
	tUcHomeUserDo := tUcHomeUserTable.WithContext(context.Background())
	tUcHomeUser := &model.TUcHomeUser{
		Id:        iotutil.GetNextSeqInt64(),
		HomeId:    req.HomeId,
		UserId:    req.UserId,
		RoleType:  iotutil.ToInt32(_const.RoleMember),
		BindTime:  timestamppb.Now().AsTime(),
		Shared:    iotutil.ToInt32(_const.ShareHome),
		CreatedAt: timestamppb.Now().AsTime(),
		UpdatedAt: timestamppb.Now().AsTime(),
		DeletedAt: gorm.DeletedAt{},
	}

	err := tUcHomeUserDo.Create(tUcHomeUser)
	if err != nil {
		logger.Errorf("JoinHome failed, error:%s", err.Error())
		return err
	}
	return nil
}

// 家庭成员角色设置
func (s *UcHomeSvc) SetRole(req *proto.UcHomeDetailRequest) error {
	tUcHomeUserTable := orm.Use(iotmodel.GetDB()).TUcHomeUser
	tUcHomeUserDo := tUcHomeUserTable.WithContext(context.Background())

	//要更新的字段,不包括主键
	var updateField []field.Expr
	updateField = append(updateField, tUcHomeUserTable.RoleType)
	tUcHomeUserDo = tUcHomeUserDo.Select(updateField...)
	tUcHomeUserDo = tUcHomeUserDo.Where(tUcHomeUserTable.HomeId.Eq(req.HomeId), tUcHomeUserTable.UserId.Eq(req.ThirdUserId))
	dbObj := convert.UcHomeUserReq_pb2db(req)
	_, err := tUcHomeUserDo.Updates(dbObj)
	if err != nil {
		logger.Errorf("SetRole error : %s", err.Error())
		return err
	}
	return nil
}

// 移除成员
func (s *UcHomeSvc) RemoveMembers(req *proto.UcHomeDetailRequest) error {
	tUcHomeUserTable := orm.Use(iotmodel.GetDB()).TUcHomeUser
	tUcHomeUserDo := tUcHomeUserTable.WithContext(context.Background())
	tUcHomeUserDo = tUcHomeUserDo.Where(tUcHomeUserTable.HomeId.Eq(req.HomeId), tUcHomeUserTable.UserId.Eq(req.ThirdUserId))
	_, err := tUcHomeUserDo.Delete()
	if err != nil {
		logger.Errorf("DeleteUcHome error : %s", err.Error())
		return err
	}
	return nil
}

// 转移家庭所有权
func (s *UcHomeSvc) TransferOwnership(req *proto.UcHomeDetailRequest) error {
	tran := orm.Use(iotmodel.GetDB())
	tran.Transaction(func(tx *orm.Query) error {
		tUcHomeUserTable := orm.Use(iotmodel.GetDB()).TUcHomeUser
		tUcHomeUserDo := tUcHomeUserTable.WithContext(context.Background())
		//要更新的字段,不包括主键
		var updateField []field.Expr
		updateField = append(updateField, tUcHomeUserTable.RoleType)
		tUcHomeUserDo = tUcHomeUserDo.Select(updateField...)
		tUcHomeUserDo = tUcHomeUserDo.Where(tUcHomeUserTable.HomeId.Eq(req.HomeId), tUcHomeUserTable.UserId.Eq(req.ThirdUserId))
		dbObj := model.TUcHomeUser{
			RoleType: iotutil.ToInt32(_const.RoleSuperAdministrator),
		}
		_, err := tUcHomeUserDo.Updates(dbObj)
		if err != nil {
			logger.Errorf("TransferOwnership error : %s", err.Error())
			return err
		}

		newTUcHomeUserDo := tUcHomeUserTable.WithContext(context.Background())
		//要更新的字段,不包括主键
		var newUpdateField []field.Expr
		newUpdateField = append(newUpdateField, tUcHomeUserTable.RoleType)
		newTUcHomeUserDo = newTUcHomeUserDo.Select(newUpdateField...)
		newTUcHomeUserDo = newTUcHomeUserDo.Where(tUcHomeUserTable.HomeId.Eq(req.HomeId), tUcHomeUserTable.UserId.Eq(req.UserId))
		newDbObj := model.TUcHomeUser{
			RoleType: iotutil.ToInt32(_const.RoleAdministrator),
		}
		_, err = newTUcHomeUserDo.Updates(newDbObj)
		if err != nil {
			logger.Errorf("TransferOwnership error : %s", err.Error())
			return err
		}
		return nil
	})
	return nil
}

type TempHomeInfo struct {
	*model.TUcHomeUser
	Photo    string `gorm:"column:photo" json:"photo"`        // 头像URL
	NickName string `gorm:"column:nick_name" json:"nickName"` // 昵称
}

// 家庭详情（unloadSet设置不需要加载的数据内容，默认全部加载）
func (s *UcHomeSvc) HomeDetail(req *proto.UcHomeDetailRequest) (*proto.UcHomeDetail, error) {
	if req.HomeId == 0 {
		return nil, errors.New("HomeId is empty")
	}
	var q = orm.Use(iotmodel.GetDB())
	tTUcHome := q.TUcHome
	tTUcHomeUser := q.TUcHomeUser
	tTUcUser := q.TUcUser
	tTUcHomeRoom := q.TUcHomeRoom

	//获取家庭信息
	tUcHome, err := tTUcHome.WithContext(context.Background()).Where(tTUcHome.Id.Eq(req.HomeId)).Find()
	if err != nil {
		logger.Errorf("HomeDetail query homeInfo error")
		return nil, err
	}
	ucHomeDetail := proto.UcHomeDetail{}
	ucHomeDetail.Data = convert.UcHome_db2pb(tUcHome[0])

	unloadSet := req.UnloadSet
	//设置默认值，默认全部为false，表示全部加载
	if unloadSet == nil {
		unloadSet = &proto.UnLoadDataSet{}
	}
	//获取家庭用户列表
	if !unloadSet.UnLoadUsers {
		var homeUserList []*TempHomeInfo
		err = tTUcHomeUser.WithContext(context.Background()).
			LeftJoin(tTUcUser, tTUcHomeUser.UserId.EqCol(tTUcUser.Id)).
			Where(tTUcHomeUser.HomeId.Eq(req.HomeId), tTUcHomeUser.DeletedAt.IsNull(), tTUcUser.DeletedAt.IsNull()).
			Order(tTUcHomeUser.CreatedAt).
			Select(tTUcHomeUser.ALL, tTUcUser.NickName, tTUcUser.Photo).
			Scan(&homeUserList)

		if err != nil {
			logger.Errorf("HomeDetail users error: %s", err.Error())
			return &ucHomeDetail, err
		}
		if len(homeUserList) == 0 {
			logger.Errorf("HomeDetail users is 0")
			return &ucHomeDetail, err
		}
		ucHomeDetail.UserList = make([]*proto.UserHome, 0)
		for _, v := range homeUserList {
			userHome := proto.UserHome{
				Uid:      iotutil.ToString(v.UserId),
				Role:     v.RoleType,
				Photo:    v.Photo,
				NickName: v.NickName,
			}
			ucHomeDetail.UserList = append(ucHomeDetail.UserList, &userHome)
		}
	}

	//获取家庭房间列表
	if !unloadSet.UnLoadRooms {
		homeRoomList, err := tTUcHomeRoom.WithContext(context.Background()).
			Where(tTUcHomeRoom.HomeId.Eq(req.HomeId)).
			Order(tTUcHomeRoom.Sort).Find()
		if err != nil {
			logger.Errorf("HomeDetail rooms error : %s", err.Error())
			return &ucHomeDetail, err
		}
		ucHomeDetail.RoomList = make([]*proto.RoomHome, 0)
		//判断是否沿用了默认房间，如果沿用则需要进行翻译，否则不需要翻译
		//GetDefaultRooms
		for _, v := range homeRoomList {
			homeRoom := proto.RoomHome{
				RoomId:         iotutil.ToString(v.Id),
				Name:           v.RoomName,
				Sort:           v.Sort,
				Icon:           v.IconUrl,
				RoomTemplateId: v.RoomTemplateId,
			}
			ucHomeDetail.RoomList = append(ucHomeDetail.RoomList, &homeRoom)
		}
	}

	//加载家庭用户数据
	if !unloadSet.UnLoadDevices {
		//获取家庭设备列表
		devListProtoResp, err := rpcClient.IotDeviceHomeService.HomeDevList(context.Background(), &proto.IotDeviceHomeHomeId{
			HomeId: req.HomeId,
		})
		if err != nil {
			iotlogger.LogHelper.Errorf("调用IotDeviceHomeService.HomeDevList异常，%s", err.Error())
			return &ucHomeDetail, err
		}
		if devListProtoResp.Code != 200 {
			iotlogger.LogHelper.Errorf("调用IotDeviceHomeService.HomeDevList异常，- %s", devListProtoResp.Message)
			return &ucHomeDetail, err
		}
		if devListProtoResp != nil && devListProtoResp.Code == 200 {
			result := make([]*proto.DeviceInfo, len(devListProtoResp.DevList))
			productIds := []int64{}
			for _, v := range devListProtoResp.DevList {
				productIds = append(productIds, v.ProductId)
			}
			productInfoList := &proto.OpmProductResponse{}
			if len(productIds) > 0 {
				productInfoList, err = rpcClient.OpmProductService.ListsByProductIds(context.Background(), &proto.ListsByProductIdsRequest{
					ProductIds: productIds,
				})
			}

			for i, v := range devListProtoResp.DevList {
				var productPic, roomName, roomIcon, productKey string
				var roomTemplateId int64
				var roomSort string
				deviceInfo := proto.DeviceInfo{}
				for _, homeRoom := range ucHomeDetail.RoomList {
					if v.RoomId == homeRoom.RoomId {
						roomName = homeRoom.Name
						roomTemplateId = homeRoom.RoomTemplateId
						roomIcon = homeRoom.Icon
						roomSort = iotutil.ToString(homeRoom.Sort)
						break
					}
				}
				if productInfoList != nil && len(productInfoList.Data) > 0 {
					for _, productInfo := range productInfoList.Data {
						if v.ProductId == productInfo.Id {
							productPic = productInfo.ImageUrl
							productKey = productInfo.ProductKey
							break
						}
					}
				}
				v.RoomName = roomName
				v.RoomIcon = roomIcon
				v.RoomSort = roomSort
				v.ProductPic = productPic
				v.ProductKey = productKey
				devInfo := convert.IotDeviceInfo_pb2db(v)
				devInfo.RoomTemplateId = roomTemplateId
				deviceInfo.Data = devInfo
				result[i] = &deviceInfo
			}
			ucHomeDetail.DeviceList = result //设备列表
		}
	}
	return &ucHomeDetail, nil
}

// 只有一个家庭时，删除此家庭，新建一个默认家庭
func (s *UcHomeSvc) CreateNewHome(req *proto.UcHomeDetailRequest) error {
	lang, _ := CheckLang(s.Ctx)
	//geo, err := iotutil.Geoip(req.Ip, config.Global.IpService.QueryUrl, config.Global.IpService.AppCode) //根据ip获取位置信息
	geo, err := Geoip(req.Ip)
	if err != nil {
		logger.Errorf("get address by ip[%s], error:%s", req.Ip, err.Error())
	}
	homeId := iotutil.GetNextSeqInt64()
	//查询默认房间信息
	dictRet, err := rpcClient.DictDataService.Lists(context.Background(), &proto.ConfigDictDataListRequest{Query: &proto.ConfigDictData{
		DictType: iotconst.Dict_type_default_rooms,
	}})
	if err != nil {
		return err
	}
	if dictRet.Code != 200 {
		return errors.New("默认房间信息获取失败")
	}

	q := orm.Use(iotmodel.GetDB())
	q.Transaction(func(tx *orm.Query) error {
		tUcHomeDo := tx.TUcHome.WithContext(context.Background())
		tUcHome := &model.TUcHome{
			Id:        homeId,
			Name:      getAppDefaultHomeName(lang),
			IconUrl:   "",
			Sid:       0,
			Lat:       geo.Lat,
			Lng:       geo.Lng,
			Country:   geo.Country,
			Province:  geo.Province,
			City:      geo.City,
			District:  geo.District,
			CreatedBy: req.ThirdUserId,
		}

		err := tUcHomeDo.Create(tUcHome)
		if err != nil {
			logger.Errorf("TUcHome create failed, error:%s", err.Error())
			return err
		}

		tUcHomeUserDo := tx.TUcHomeUser.WithContext(context.Background())
		tUcHomeUser := &model.TUcHomeUser{
			Id:        iotutil.GetNextSeqInt64(),
			HomeId:    homeId,
			UserId:    req.ThirdUserId,
			RoleType:  1,
			BindTime:  time.Now(),
			Shared:    1,
			CreatedBy: req.ThirdUserId,
		}

		err = tUcHomeUserDo.Create(tUcHomeUser)
		if err != nil {
			logger.Errorf("TUcHomeUser create failed, error:%s", err.Error())
			return err
		}

		rooms := make([]*model.TUcHomeRoom, 0)
		for _, dict := range dictRet.Data {
			rooms = append(rooms, &model.TUcHomeRoom{
				Id:        iotutil.GetNextSeqInt64(),
				HomeId:    homeId,
				RoomName:  dict.DictLabel,
				IconUrl:   dict.Listimg,
				Sort:      iotutil.ToInt32(dict.DictValue),
				CreatedBy: req.ThirdUserId,
			})
		}
		err = tx.TUcHomeRoom.WithContext(context.Background()).Create(rooms...)
		if err != nil {
			logger.Errorf("tUcHomeRoom create failed, error:%s", err.Error())
			return err
		}
		return nil
	})
	return nil
}

func (s *UcHomeSvc) ChangeUserDefaultHomeId(req *proto.UcHomeUserListRequest, ucHomeUserList []*proto.UcHomeUser) error {
	var q = orm.Use(iotmodel.GetDB())
	param := req.Query
	tTUcUser := q.TUcUser

	//查询默认家庭id为当前家庭id的用户信息
	do := tTUcUser.WithContext(context.Background())
	do = do.Where(tTUcUser.Id.Eq(param.UserId), tTUcUser.DefaultHomeId.Eq(iotutil.ToString(param.HomeId)))
	var list []*struct {
		*model.TUcUser
	}
	err := do.Select(tTUcUser.ALL).Scan(&list)
	if err == nil && len(list) == 1 {
		userId := list[0].Id
		var changeHomeId int64
		for _, homeUserObj := range ucHomeUserList {
			if homeUserObj.HomeId != param.HomeId {
				changeHomeId = homeUserObj.HomeId
				break
			}
		}
		if changeHomeId == 0 {
			return errors.New("没有找到用户的其他家庭信息")
		}
		tTUcUserDo := tTUcUser.WithContext(context.Background())
		// fixme 请检查条件和校验参数
		//要更新的字段,不包括主键
		var updateField []field.Expr
		updateField = append(updateField, tTUcUser.DefaultHomeId)
		tTUcUserDo = tTUcUserDo.Select(updateField...)
		tTUcUserDo = tTUcUserDo.Where(tTUcUser.Id.Eq(userId))
		//自动修改默认家庭id
		dbObj := convert.UcUser_pb2db(&proto.UcUser{DefaultHomeId: iotutil.ToString(changeHomeId)})
		_, err := tTUcUserDo.Updates(dbObj)
		if err != nil {
			logger.Errorf("UpdateUcUser error : %s", err.Error())
			return err
		}
	}
	return nil
}

func clearHomeCached(data iotstruct.DeviceRedisUpdate) error {
	iotlogger.LogHelper.Infof("清理缓存，参数：%v", iotutil.ToString(data))
	// 删除家庭详情缓存persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(userId)),
	keys := make([]string, 0)
	//家庭房间
	for _, l := range iotconst.APP_SUPPORT_LANGUAGE {
		keys = append(keys, persist.GetRedisKey(iotconst.APP_HOME_ROOM_LIST_DATA, l, data.HomeId))
	}
	ctx := context.Background()
	svc := UcHomeUserSvc{Ctx: context.Background()}
	list, _, err := svc.GetListUcHomeUser(&proto.UcHomeUserListRequest{Query: &proto.UcHomeUser{HomeId: iotutil.ToInt64(data.HomeId)}})
	if err != nil {
		return err
	}
	for i := range list {
		keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(list[i].UserId)))
		keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, data.HomeId, iotutil.ToString(list[i].UserId)))
	}
	pipe := cached.RedisStore.Pipeline()
	pipe.Del(ctx, keys...)
	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}
	return nil
}
