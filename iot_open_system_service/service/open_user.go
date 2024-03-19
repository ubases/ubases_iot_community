// Code generated by sgen.exe,2022-04-27 10:55:26. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
	"context"
	"errors"
	"time"

	"gorm.io/gen"

	"go-micro.dev/v4/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gen/field"

	"cloud_platform/iot_common/iotutil"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_open_system/model"
	"cloud_platform/iot_model/db_open_system/orm"
	"cloud_platform/iot_open_system_service/convert"
	proto "cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4/metadata"
)

type OpenUserSvc struct {
	Ctx context.Context
}

func (s *OpenUserSvc) VerifyFiled(req *proto.OpenUser) error {
	if req.UserName == "" {
		return errors.New("phone or email not nil")
	}
	if req.UserPassword == "" {
		return errors.New("password not nil")
	}
	if req.UserSalt == "" {
		return errors.New("salt not nil")
	}
	return nil
}

// 设置公共属性[创建时间,创建人,修改时间,修改人]
func (s *OpenUserSvc) SetCommonFiled(req *proto.OpenUser, userId int64, opterType int) {
	if opterType == 1 {
		req.CreatedBy = userId
		req.CreatedAt = timestamppb.New(time.Now())
	} else if opterType == 2 {
		req.UpdatedBy = userId
		req.UpdatedAt = timestamppb.New(time.Now())
	}

}

// RegisterUser
func (s *OpenUserSvc) RegisterUser(req *proto.OpenUserRegisterRequest) (int64, error) {
	//1.非空校验
	if req.Account == "" {
		return 0, errors.New("账号不能为空")
	}
	if req.Password == "" {
		return 0, errors.New("密码不能为空")
	}
	if req.CompanyName == "" && req.UserType == 1 {
		return 0, errors.New("公司名称不能为空")
	}
	q := orm.Use(iotmodel.GetDB())
	var userId int64
	var err error
	err = q.Transaction(func(tx *orm.Query) error {
		tOpenUser := q.TOpenUser
		do := tOpenUser.WithContext(context.Background())

		var userCheck gen.Condition
		//判断是否账号、手机还是邮箱
		mobile, email := "", ""
		if iotutil.IsPhone(req.Account) {
			mobile = req.Account
			userCheck = gen.Condition(do.Where(tOpenUser.UserName.Eq(req.Account)).Or(do.Where(tOpenUser.Mobile.Eq(req.Account))))
		} else if iotutil.IsEmail(req.Account) {
			email = req.Account
			userCheck = gen.Condition(do.Where(tOpenUser.UserName.Eq(req.Account)).Or(do.Where(tOpenUser.Mobile.Eq(req.Account))))
		} else {
			userCheck = gen.Condition(do.Where(tOpenUser.UserName.Eq(req.Account)))
		}

		if c, err := do.Where(userCheck).Count(); err != nil || c > 0 {
			if err != nil {
				return err
			}
			return errors.New("账号已存在")
		}

		//2.用户数据入库
		userId = iotutil.GetNextSeqInt64()
		UserSalt := iotutil.GetSecret(8)
		password := iotutil.Md5(req.Password + UserSalt)
		err = do.Create(&model.TOpenUser{
			Id:            userId,
			UserName:      req.Account,
			Mobile:        mobile,
			UserNickname:  req.Account, // 例如：150***8101
			UserPassword:  password,
			UserSalt:      UserSalt,
			UserStatus:    1,
			UserEmail:     email,
			Sex:           iotconst.SEX_UNKNOWN,
			LastLoginIp:   req.IP,
			AccountType:   req.UserType,
			CompanyName:   req.CompanyName,
			AccountOrigin: 1,
		})
		if err != nil {
			return err
		}
		tOpenCompany := q.TOpenCompany
		doOpenCompany := tOpenCompany.WithContext(context.Background())
		//3.组合company数据[组合基础数据入库]进行入库
		tenantId := iotutil.GetSecret(6) //租户Id

		companyName := req.CompanyName
		if req.UserType == 2 {
			companyName = ""
		}
		companyInfo := model.TOpenCompany{
			Id:          iotutil.GetNextSeqInt64(),
			TenantId:    tenantId,
			UserId:      userId, //todo 这个字段是否可以取消
			Name:        companyName,
			Status:      1,            //状态（=1 未提交 ,=2 认证中,   =3 已认证, =4 禁用
			AccountType: req.UserType, //账号类型 1 企业账号, 2 个人账号
			Email:       email,
			IsRealName:  2,
			CreatedBy:   userId,
			UserName:    req.Account,
		}
		err = doOpenCompany.Create(&companyInfo)
		if err != nil {
			return err
		}
		//4.建议用户与公司的关系
		tUserCompany := q.TOpenUserCompany
		doUserCompany := tUserCompany.WithContext(context.Background())
		userCompany := model.TOpenUserCompany{
			Id:        iotutil.GetNextSeqInt64(),
			UserId:    userId,
			TenantId:  tenantId,
			CompanyId: companyInfo.Id,
			UserType:  iotconst.OPEN_USER_MAIN_ACCOUNT,
			UserName:  req.Account, //这里用账号是正确的逻辑. 次字段只在提供空间列表的时候使用.
		}
		err = doUserCompany.Create(&userCompany)
		if err != nil {
			return err
		}

		//新增默认角色
		err = s.CreateAdminRoleAndMenuIds(iotutil.ToString(userId), tenantId)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (s *OpenUserSvc) AddUser(req *proto.DeveloperEntitys) (int64, error) {
	//1.非空校验
	if req.Account == "" {
		return 0, errors.New("账号不能为空")
	}
	if req.Password == "" {
		return 0, errors.New("密码不能为空")
	}
	// if req.CompanyName == "" {
	// 	return 0, errors.New("公司名称不能为空")
	// }
	q := orm.Use(iotmodel.GetDB())
	var userId int64
	var err error
	err = q.Transaction(func(tx *orm.Query) error {
		tOpenUser := q.TOpenUser
		do := tOpenUser.WithContext(context.Background())
		var userCheck gen.Condition
		//判断是否账号、手机还是邮箱
		email := ""
		if iotutil.IsPhone(req.Account) {
			//mobile = req.Account
			userCheck = gen.Condition(do.Where(tOpenUser.UserName.Eq(req.Account)).Or(do.Where(tOpenUser.Mobile.Eq(req.Account))))
		} else if iotutil.IsEmail(req.Account) {
			email = req.Account
			userCheck = gen.Condition(do.Where(tOpenUser.UserName.Eq(req.Account)).Or(do.Where(tOpenUser.UserEmail.Eq(req.Account))))
		} else {
			userCheck = gen.Condition(do.Where(tOpenUser.UserName.Eq(req.Account)))
		}

		if c, err := do.Where(userCheck).Count(); err != nil || c > 0 {
			if err != nil {
				return err
			}
			return errors.New("账号已存在")
		}

		//2.用户数据入库
		userId = iotutil.GetNextSeqInt64()
		UserSalt := iotutil.GetSecret(8)
		password := iotutil.Md5(req.Password + UserSalt)
		err = do.Create(&model.TOpenUser{
			Id:            userId,
			UserName:      req.Account,
			Mobile:        req.Phone,
			UserNickname:  req.Account, // 例如：150***8101
			UserPassword:  password,
			UserSalt:      UserSalt,
			UserStatus:    1,
			UserEmail:     req.Email,
			Address:       req.Address,
			Sex:           iotconst.SEX_UNKNOWN,
			LastLoginIp:   "",
			AccountType:   req.AccountType,
			CompanyName:   req.CompanyName,
			AccountOrigin: 2,
		})
		if err != nil {
			return err
		}
		tOpenCompany := q.TOpenCompany
		doOpenCompany := tOpenCompany.WithContext(context.Background())
		//3.组合company数据[组合基础数据入库]进行入库
		tenantId := iotutil.GetSecret(6) //租户Id
		companyInfo := model.TOpenCompany{
			Id:          iotutil.GetNextSeqInt64(),
			TenantId:    tenantId,
			UserId:      userId, //todo 这个字段是否可以取消
			Name:        req.CompanyName,
			Status:      1, //状态（=1 未提交 ,=2 认证中,   =3 已认证, =4 禁用
			AccountType: 2, //账号类型 1 企业账号, 2 个人账号
			Email:       email,
			IsRealName:  2,
			UserName:    req.Account,
			CreatedBy:   userId,
		}
		err = doOpenCompany.Create(&companyInfo)
		if err != nil {
			return err
		}
		//4.建议用户与公司的关系
		tUserCompany := q.TOpenUserCompany
		doUserCompany := tUserCompany.WithContext(context.Background())
		userCompany := model.TOpenUserCompany{
			Id:        iotutil.GetNextSeqInt64(),
			UserId:    userId,
			TenantId:  tenantId,
			CompanyId: companyInfo.Id,
			UserType:  iotconst.OPEN_USER_MAIN_ACCOUNT,
			UserName:  req.Account,
		}
		err = doUserCompany.Create(&userCompany)
		if err != nil {
			return err
		}
		//新增默认角色
		err = s.CreateAdminRoleAndMenuIds(iotutil.ToString(userId), tenantId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return userId, nil
}

// 给注册用户新增默认角色和菜单
func (s *OpenUserSvc) CreateAdminRoleAndMenuIds(userId string, tenandId string) error {
	var allRoleId int64
	role := OpenRoleSvc{}

	resRole, errRole := role.FindOpenRole(&proto.OpenRoleFilter{
		IsDefault: 1, //默认角色
		IsAdmin:   1, //管理员
	})
	if errRole != nil && errRole.Error() == ioterrs.ErrRecordNotFound {
		menu := OpenAuthRuleSvc{}
		menuIds, err := menu.GetAllMenuIds()
		if err != nil {
			return err
		}
		_, roleId, errRole := role.RoleAddByService(&proto.OpenRoleAddRequest{
			RoleName: "管理员",
			MenuIds:  menuIds,
		}, tenandId, userId, 1)
		if errRole != nil {
			return errRole
		}
		allRoleId = roleId
	} else if errRole != nil {
		return errRole
	} else {
		casbin := OpenCasbinRuleSvc{}
		rules, _, err := casbin.GetListOpenCasbinRule(&proto.OpenCasbinRuleListRequest{
			Query: &proto.OpenCasbinRule{
				V0: iotutil.ToString(resRole.Id),
			},
		})
		if err != nil {
			return err
		}
		if len(rules) == 0 {
			auth := OpenAuthRuleSvc{}
			menuids, err := auth.GetAllMenuIds()
			if err != nil {
				return err
			}
			//角色菜单权限
			//修改权限为批量新增方式
			var rules [][]string = make([][]string, 0)
			var idStr string = iotutil.ToString(resRole.Id)
			for _, mid := range menuids {
				var tmp []string = []string{idStr, mid, "ALL"}
				rules = append(rules, tmp)
			}
			_, err = Casbin_Enforcer.AddPolicies(rules)
			if err != nil {
				return err
			}
		}
		allRoleId = resRole.Id
	}
	//adminRoleId := "7889966926082965504" //TODO  后续改为从表里面查询出默认角色ID
	//role.SetUserRoleService(iotutil.ToString(roleId), userId, tenandId)
	role.SetUserRoleService(iotutil.ToString(allRoleId), userId, tenandId)

	return nil
}

// 创建OpenUser
func (s *OpenUserSvc) CreateOpenUser(req *proto.OpenUser) (*proto.OpenUser, error) {
	//1.非空校验
	errFiled := s.VerifyFiled(req)
	if errFiled != nil {
		return nil, errFiled
	}

	s.SetCommonFiled(req, req.CreatedBy, 1)

	//2.数据入库
	res, err := s.InsterOpenUser(req)
	if err != nil {
		return nil, err
	}

	//3.组合company数据[组合基础数据入库]进行入库
	var companyService = OpenCompanySvc{}
	companyId := iotutil.GetNextSeqInt64()
	_, errCompany := companyService.CreateOpenCompany(&proto.OpenCompany{
		Id:          companyId,
		UserId:      req.Id,
		Name:        req.UserNickname,
		Status:      1,
		AccountType: req.AccountType,
		Email:       req.UserEmail,
		IsRealName:  2,
		CreatedBy:   req.CreatedBy,
	})
	if errCompany != nil {
		//回滚用户
		s.DeleteByIdOpenUser(&proto.OpenUser{Id: req.Id})
		return nil, errCompany
	}
	//4.组合 user_company数据并且入库
	var userCompany = OpenUserCompanySvc{}
	userCompanyId := iotutil.GetNextSeqInt64()
	_, errUserCompany := userCompany.CreateOpenUserCompany(&proto.OpenUserCompany{
		Id:        userCompanyId,
		UserId:    req.Id,
		CompanyId: companyId,
		UserType:  2,
		CreatedBy: req.CreatedBy,
	})
	if errUserCompany != nil {
		//回滚公司
		companyService.DeleteByIdOpenCompany(&proto.OpenCompany{Id: companyId})
		//回滚用户
		s.DeleteByIdOpenUser(&proto.OpenUser{Id: req.Id})
		//抛出异常
		return nil, errUserCompany
	}

	return res, err
}

func (s *OpenUserSvc) InsterOpenUser(req *proto.OpenUser) (*proto.OpenUser, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())
	dbObj := convert.OpenUser_pb2db(req)
	err := do.Create(dbObj)
	if err != nil {
		logger.Errorf("InsterOpenUser error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据条件删除OpenUser
func (s *OpenUserSvc) DeleteOpenUser(req *proto.OpenUser) (*proto.OpenUser, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.UserName != "" { //字符串
		do = do.Where(t.UserName.Eq(req.UserName))
	}
	if req.Mobile != "" { //字符串
		do = do.Where(t.Mobile.Eq(req.Mobile))
	}
	if req.UserNickname != "" { //字符串
		do = do.Where(t.UserNickname.Eq(req.UserNickname))
	}
	if req.Birthday != 0 { //整数
		do = do.Where(t.Birthday.Eq(req.Birthday))
	}
	if req.UserPassword != "" { //字符串
		do = do.Where(t.UserPassword.Eq(req.UserPassword))
	}
	if req.UserSalt != "" { //字符串
		do = do.Where(t.UserSalt.Eq(req.UserSalt))
	}
	if req.UserStatus != 0 { //整数
		do = do.Where(t.UserStatus.Eq(req.UserStatus))
	}
	if req.UserEmail != "" { //字符串
		do = do.Where(t.UserEmail.Eq(req.UserEmail))
	}
	if req.Sex != 0 { //整数
		do = do.Where(t.Sex.Eq(req.Sex))
	}
	if req.Avatar != "" { //字符串
		do = do.Where(t.Avatar.Eq(req.Avatar))
	}
	if req.Remark != "" { //字符串
		do = do.Where(t.Remark.Eq(req.Remark))
	}
	if req.Address != "" { //字符串
		do = do.Where(t.Address.Eq(req.Address))
	}
	if req.Describe != "" { //字符串
		do = do.Where(t.Describe.Eq(req.Describe))
	}
	if req.LastLoginIp != "" { //字符串
		do = do.Where(t.LastLoginIp.Eq(req.LastLoginIp))
	}
	if req.AccountType != 0 { //整数
		do = do.Where(t.AccountType.Eq(req.AccountType))
	}
	if req.CompanyName != "" { //字符串
		do = do.Where(t.CompanyName.Eq(req.CompanyName))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteOpenUser error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键删除OpenUser
func (s *OpenUserSvc) DeleteByIdOpenUser(req *proto.OpenUser) (*proto.OpenUser, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())
	// fixme 请检查条件

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	_, err := do.Delete()
	if err != nil {
		logger.Errorf("DeleteByIdOpenUser error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// 根据数据库表主键批量删除OpenUser
func (s *OpenUserSvc) DeleteByIdsOpenUser(req *proto.OpenUserBatchDeleteRequest) (*proto.OpenUserBatchDeleteRequest, error) {
	var err error
	for _, k := range req.Keys {
		t := orm.Use(iotmodel.GetDB()).TOpenUser
		do := t.WithContext(context.Background())

		do = do.Where(t.Id.Eq(k.Id))

		_, err = do.Delete()
		if err != nil {
			logger.Errorf("DeleteByIdsOpenUser error : %s", err.Error())
			break
		}
	}
	return req, err
}

// 根据主键更新OpenUser
func (s *OpenUserSvc) UpdateOpenUser(req *proto.OpenUser) (*proto.OpenUser, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	if req.UserName != "" { //字符串
		updateField = append(updateField, t.UserName)
	}
	if req.Mobile != "" { //字符串
		updateField = append(updateField, t.Mobile)
	}
	if req.UserNickname != "" { //字符串
		updateField = append(updateField, t.UserNickname)
	}
	if req.Birthday != 0 { //整数
		updateField = append(updateField, t.Birthday)
	}
	if req.UserPassword != "" { //字符串
		updateField = append(updateField, t.UserPassword)
	}
	if req.UserSalt != "" { //字符串
		updateField = append(updateField, t.UserSalt)
	}
	if req.UserStatus != 0 { //整数
		updateField = append(updateField, t.UserStatus)
	}
	if req.UserEmail != "" { //字符串
		updateField = append(updateField, t.UserEmail)
	}
	if req.Sex != 0 { //整数
		updateField = append(updateField, t.Sex)
	}
	if req.Avatar != "" { //字符串
		updateField = append(updateField, t.Avatar)
	}
	if req.Remark != "" { //字符串
		updateField = append(updateField, t.Remark)
	}
	if req.Address != "" { //字符串
		updateField = append(updateField, t.Address)
	}
	if req.Describe != "" { //字符串
		updateField = append(updateField, t.Describe)
	}
	if req.LastLoginIp != "" { //字符串
		updateField = append(updateField, t.LastLoginIp)
	}
	if req.AccountType != 0 { //整数
		updateField = append(updateField, t.AccountType)
	}
	if req.CompanyName != "" { //字符串
		updateField = append(updateField, t.CompanyName)
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
		logger.Error("UpdateOpenUser error : Missing condition")
		return nil, errors.New("Missing condition")
	}

	dbObj := convert.OpenUser_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateOpenUser error : %s", err.Error())
		return nil, err
	}
	return req, err
}

// //根据主键更新所有字段OpenUser
func (s *OpenUserSvc) UpdateAllOpenUser(req *proto.OpenUser) (*proto.OpenUser, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数
	//要更新的字段,不包括主键
	var updateField []field.Expr

	updateField = append(updateField, t.UserName)
	updateField = append(updateField, t.Mobile)
	updateField = append(updateField, t.UserNickname)
	updateField = append(updateField, t.Birthday)
	updateField = append(updateField, t.UserPassword)
	updateField = append(updateField, t.UserSalt)
	updateField = append(updateField, t.UserStatus)
	updateField = append(updateField, t.UserEmail)
	updateField = append(updateField, t.Sex)
	updateField = append(updateField, t.Avatar)
	updateField = append(updateField, t.Remark)
	updateField = append(updateField, t.Address)
	updateField = append(updateField, t.Describe)
	updateField = append(updateField, t.LastLoginIp)
	updateField = append(updateField, t.LastLoginTime)
	updateField = append(updateField, t.AccountType)
	updateField = append(updateField, t.CompanyName)
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
		logger.Error("UpdateAllOpenUser error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OpenUser_pb2db(req)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateAllOpenUser error : %s", err.Error())
		return nil, err
	}
	return req, err
}

func (s *OpenUserSvc) UpdateFieldsOpenUser(req *proto.OpenUserUpdateFieldsRequest) (*proto.OpenUser, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())

	var updateField []field.Expr
	for _, v := range req.Fields {
		col, ok := t.GetFieldByName(v)
		if ok {
			updateField = append(updateField, col)
		}
	}
	if len(updateField) == 0 {
		err := errors.New("UpdateFieldsOpenUser error : missing updateField")
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
		logger.Error("UpdateFieldsOpenUser error : Missing condition")
		return nil, errors.New("Missing condition")
	}
	dbObj := convert.OpenUser_pb2db(req.Data)
	_, err := do.Updates(dbObj)
	if err != nil {
		logger.Errorf("UpdateFieldsOpenUser error : %s", err.Error())
		return nil, err
	}
	return req.Data, nil
}

// 根据非空条件查找OpenUser
func (s *OpenUserSvc) FindOpenUser(req *proto.OpenUserFilter) (*proto.OpenUser, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	if req.UserName != "" { //字符串
		do = do.Where(t.UserName.Eq(req.UserName))
	}
	if req.Mobile != "" { //字符串
		do = do.Where(t.Mobile.Eq(req.Mobile))
	}
	if req.UserNickname != "" { //字符串
		do = do.Where(t.UserNickname.Eq(req.UserNickname))
	}
	if req.Birthday != 0 { //整数
		do = do.Where(t.Birthday.Eq(req.Birthday))
	}
	if req.UserPassword != "" { //字符串
		do = do.Where(t.UserPassword.Eq(req.UserPassword))
	}
	if req.UserSalt != "" { //字符串
		do = do.Where(t.UserSalt.Eq(req.UserSalt))
	}
	if req.UserStatus != 0 { //整数
		do = do.Where(t.UserStatus.Eq(req.UserStatus))
	}
	if req.UserEmail != "" { //字符串
		do = do.Where(t.UserEmail.Eq(req.UserEmail))
	}
	if req.Sex != 0 { //整数
		do = do.Where(t.Sex.Eq(req.Sex))
	}
	if req.Avatar != "" { //字符串
		do = do.Where(t.Avatar.Eq(req.Avatar))
	}
	if req.Remark != "" { //字符串
		do = do.Where(t.Remark.Eq(req.Remark))
	}
	if req.Address != "" { //字符串
		do = do.Where(t.Address.Eq(req.Address))
	}
	if req.Describe != "" { //字符串
		do = do.Where(t.Describe.Eq(req.Describe))
	}
	if req.LastLoginIp != "" { //字符串
		do = do.Where(t.LastLoginIp.Eq(req.LastLoginIp))
	}
	if req.AccountType != 0 { //整数
		do = do.Where(t.AccountType.Eq(req.AccountType))
	}
	if req.CompanyName != "" { //字符串
		do = do.Where(t.CompanyName.Eq(req.CompanyName))
	}
	if req.CreatedBy != 0 { //整数
		do = do.Where(t.CreatedBy.Eq(req.CreatedBy))
	}
	if req.UpdatedBy != 0 { //整数
		do = do.Where(t.UpdatedBy.Eq(req.UpdatedBy))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindOpenUser error : %s", err.Error())
		return nil, err
	}
	res := convert.OpenUser_db2pb(dbObj)
	return res, err
}

// 根据数据库表主键查找OpenUser
func (s *OpenUserSvc) FindByIdOpenUser(req *proto.OpenUserFilter) (*proto.OpenUser, error) {
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())
	// fixme 请检查条件和校验参数

	if req.Id != 0 { //整数
		do = do.Where(t.Id.Eq(req.Id))
	}
	dbObj, err := do.First()
	if err != nil {
		logger.Errorf("FindByIdOpenUser error : %s", err.Error())
		return nil, err
	}
	res := convert.OpenUser_db2pb(dbObj)
	return res, err
}

// 根据分页条件查找OpenUser,请确保req.Query的结构字段与数据表model结构体字段保持一致，否则会有编译问题
func (s *OpenUserSvc) GetListOpenUser(req *proto.OpenUserListRequest) ([]*proto.OpenUser, int64, error) {
	// fixme 请检查条件和校验参数
	var err error
	t := orm.Use(iotmodel.GetDB()).TOpenUser
	do := t.WithContext(context.Background())
	query := req.Query
	if query != nil {

		if query.Id != 0 { //整数
			do = do.Where(t.Id.Eq(query.Id))
		}
		if query.UserName != "" { //字符串
			do = do.Where(t.UserName.Like("%" + query.UserName + "%"))
		}
		if query.Mobile != "" { //字符串
			do = do.Where(t.Mobile.Like("%" + query.Mobile + "%"))
		}
		if query.UserNickname != "" { //字符串
			do = do.Where(t.UserNickname.Like("%" + query.UserNickname + "%"))
		}
		if query.Birthday != 0 { //整数
			do = do.Where(t.Birthday.Eq(query.Birthday))
		}
		if query.UserPassword != "" { //字符串
			do = do.Where(t.UserPassword.Like("%" + query.UserPassword + "%"))
		}
		if query.UserSalt != "" { //字符串
			do = do.Where(t.UserSalt.Like("%" + query.UserSalt + "%"))
		}
		if query.UserStatus != 0 { //整数
			do = do.Where(t.UserStatus.Eq(query.UserStatus))
		}
		if query.UserEmail != "" { //字符串
			do = do.Where(t.UserEmail.Like("%" + query.UserEmail + "%"))
		}
		if query.Sex != 0 { //整数
			do = do.Where(t.Sex.Eq(query.Sex))
		}
		if query.Avatar != "" { //字符串
			do = do.Where(t.Avatar.Like("%" + query.Avatar + "%"))
		}
		if query.Remark != "" { //字符串
			do = do.Where(t.Remark.Like("%" + query.Remark + "%"))
		}
		if query.Address != "" { //字符串
			do = do.Where(t.Address.Like("%" + query.Address + "%"))
		}
		if query.Describe != "" { //字符串
			do = do.Where(t.Describe.Like("%" + query.Describe + "%"))
		}
		if query.LastLoginIp != "" { //字符串
			do = do.Where(t.LastLoginIp.Like("%" + query.LastLoginIp + "%"))
		}
		if query.AccountType != 0 { //整数
			do = do.Where(t.AccountType.Eq(query.AccountType))
		}
		if query.CompanyName != "" { //字符串
			do = do.Where(t.CompanyName.Like("%" + query.CompanyName + "%"))
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

	var list []*model.TOpenUser
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
		logger.Errorf("GetListOpenUser error : %s", err.Error())
		return nil, 0, err
	}
	if len(list) == 0 {
		return nil, total, nil
	}
	result := make([]*proto.OpenUser, len(list))
	for i, v := range list {
		result[i] = convert.OpenUser_db2pb(v)
	}
	return result, total, nil
}

func (s *OpenUserSvc) GetOpenUserRouters() ([]*proto.OpenMenuTree, error) {
	//用户id
	userid, _ := metadata.Get(s.Ctx, "userid")
	if userid == "" {
		return nil, errors.New("用户id获取失败.")
	}

	//用户空间
	tenantId, _ := metadata.Get(s.Ctx, "tenantid")
	if tenantId == "" {
		return nil, errors.New("公司空间获取失败.")
	}

	//根据用户id和用户空间 获取用户的角色id
	roleIds := Casbin_Enforcer.GetRolesForUserInDomain(userid, tenantId)
	if len(roleIds) == 0 {
		iotlogger.LogHelper.Errorf("用户在该空间下未设置角色, userId:%s, tenantId:%s", userid, tenantId)
		//return nil, errors.New("用户在该空间下未设置角色.")
		var nodata []*proto.OpenMenuTree

		return nodata, nil
	}
	roleId := roleIds[0]
	//根据用户的角色id 获取角色的菜单权限id
	var roleOpen = OpenRoleSvc{}

	roleModel, _ := roleOpen.FindByIdOpenRole(&proto.OpenRoleFilter{Id: iotutil.ToInt64(roleId)})
	var Menuids []string
	if roleModel.IsAdmin == 1 && roleModel.IsDefault == 1 {
		//TODO 后续考虑云管增加开发平台的默认角色配置
		//是管理员角色. 返回所有权限
		var authruleSvc = OpenAuthRuleSvc{}
		Menuids, _ = authruleSvc.GetAllMenuIds()
	} else {
		Menuids = roleOpen.GetRoleMenuIds(roleId)
	}

	if len(Menuids) == 0 {
		return nil, errors.New("用户角色未设置任何菜单权限")
	}
	//生成菜单树提供出去
	var authRule = OpenAuthRuleSvc{}
	Menuids = authRule.AddMeunIdChildByHide(Menuids)

	menuTree, errTree := authRule.GetUserMenuTree(Menuids)
	if errTree != nil {
		return nil, errTree
	}
	return menuTree, nil
}
