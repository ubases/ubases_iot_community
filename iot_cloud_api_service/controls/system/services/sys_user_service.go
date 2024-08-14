package services

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
	micorerrors "go-micro.dev/v4/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SysUserService struct {
}

// GetUserDetail 查询用户详情
func (s SysUserService) GetUserDetail(id string) (*entitys.SysUserDetailRoleDeptRes, error) {
	rid := iotutil.ToUint64(id)
	rep, err := rpc.ClientSysUserService.FindById(context.Background(), &protosService.SysUserFilter{Id: int64(rid)})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	if rep.Data == nil || len(rep.Data) == 0 {
		return nil, errors.New("record not found")
	}

	var userData = rep.Data[0]
	userInfo := entitys.SysUser_pb2e(userData)
	resUserInfo, _ := s.GetUserDetailRoleDept(userInfo)
	return resUserInfo, err
}

// QueryLoginUserInfo 查询用户详情
func (s SysUserService) QueryLoginUserInfo(id string) (*entitys.SysLoginUserInfoRes, error) {
	rid := iotutil.ToUint64(id)
	rep, err := rpc.ClientSysUserService.FindById(context.Background(), &protosService.SysUserFilter{Id: int64(rid)})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	if rep.Data == nil || len(rep.Data) == 0 {
		return nil, errors.New("record not found")
	}

	var userData = rep.Data[0]
	userInfo := entitys.SysUser_pb2e(userData)
	resUserInfo, err := s.GetLoginUserInfo(userInfo)
	return resUserInfo, err
}

// QueryUserList 查询用户列表
func (s SysUserService) GetUserProfile(userId string) (*entitys.SysUserRoleDeptRes, error) {
	rep, err := rpc.ClientSysUserService.FindById(context.Background(), &protosService.SysUserFilter{
		Id: iotutil.ToInt64(userId),
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	if rep.Data == nil || len(rep.Data) == 0 {
		return nil, errors.New("record not found")
	}

	var resultList = []*entitys.SysUserEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.SysUser_pb2e(item))
	}
	resList, err := s.GetUsersRoleDept(resultList)
	if len(resList) > 0 {
		return resList[0], err
	}
	return nil, errors.New("not found user")
}

// QueryUserList 查询用户列表
func (s SysUserService) QueryUserList(filter entitys.QueryUser) ([]*entitys.SysUserRoleDeptRes, int64, error) {
	if filter.DeptId == "" {
		filter.DeptId = "0"
	}
	if filter.PostId == "" {
		filter.PostId = "0"
	}
	rep, err := rpc.ClientSysUserService.Lists(context.Background(), &protosService.SysUserListRequest{
		Page:      filter.Page,
		PageSize:  filter.Limit,
		SearchKey: filter.SearchKey,
		OrderKey:  iotutil.Camel2Case(filter.SortField),
		OrderDesc: filter.Sort,
		BeginTime: filter.BeginTime,
		EndTime:   filter.EndTime,
		Query: &protosService.SysUser{
			UserStatus: int32(filter.Status),
			PhoneNum:   filter.Phonenumber,
			DeptId:     iotutil.ToInt64(filter.DeptId),
			PostId:     iotutil.ToInt64(filter.PostId),
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}

	var resultList = []*entitys.SysUserEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.SysUser_pb2e(item))
	}
	resList, err := s.GetUsersRoleDept(resultList)

	//根据岗位ID 查询
	if filter.PostId != "" && filter.PostId != "0" {
		var filterResList []*entitys.SysUserRoleDeptRes
		for _, v := range resList {
			if v.PostIds != nil && len(v.PostIds) > 0 {
				//是否有岗位条件
				if iotutil.ArraysExistsString(v.PostIds, filter.PostId) {
					filterResList = append(filterResList, v)
				}
			}
		}
		return filterResList, iotutil.ToInt64(len(filterResList)), nil
	}
	return resList, rep.Total, err
}

// GetUserDetailRoleDept 获取用户详细信息角色 部门信息
func (s *SysUserService) GetUserDetailRoleDept(userInfo *entitys.SysUserEntitys) (*entitys.SysUserDetailRoleDeptRes, error) {

	resultUser := &entitys.SysUserDetailRoleDeptRes{
		CheckedRoleIds: []string{},
		CheckedPosts:   []string{},
		UserInfo:       userInfo,
		RoleList:       []*entitys.SysRoleEntitys{},
		Posts:          []*entitys.SysPostEntitys{},
	}

	roleService := SysRoleService{}
	allRoles, _, err := roleService.QuerySysRoleList(entitys.SysRoleQuery{
		Page:  1,
		Limit: 1000,
	})
	if err != nil {
		g.Log().Error(err)
		return resultUser, err
	}

	roles, err := s.GetAdminRole(iotutil.ToInt64(userInfo.Id), allRoles)
	if err != nil {
		return resultUser, err
	}
	for _, r := range roles {
		resultUser.CheckedRoleIds = append(resultUser.CheckedRoleIds, r.Id)
		resultUser.RoleList = append(resultUser.RoleList, r)
	}
	postIds, posts, err := s.GetUserPostInfo(iotutil.ToInt64(userInfo.Id))
	if err != nil {
		return resultUser, err
	}
	resultUser.CheckedPosts = postIds
	resultUser.Posts = posts
	return resultUser, nil
}

// GetLoginUserInfo 获取用户详细信息角色 部门信息
func (s *SysUserService) GetLoginUserInfo(userInfo *entitys.SysUserEntitys) (*entitys.SysLoginUserInfoRes, error) {
	roleService := SysRoleService{}
	allRoles, _, err := roleService.QuerySysRoleList(entitys.SysRoleQuery{
		Page:  1,
		Limit: 1000,
	})
	if err != nil {
		g.Log().Error(err)
		return nil, err
	}
	roles, err := s.GetAdminRole(iotutil.ToInt64(userInfo.Id), allRoles)
	if err != nil {
		return nil, err
	}
	resultUser := &entitys.SysLoginUserInfoRes{
		Permissions: []string{},
		Roles:       []string{},
		User:        userInfo,
	}
	roleIds := []int64{}
	for _, r := range roles {
		roleIds = append(roleIds, iotutil.ToInt64(r.Id))
		resultUser.Roles = append(resultUser.Roles, r.Name)
	}
	//if iotconst.AccountType(userInfo.AccountType) == iotconst.PLATFROM_USER { }
	//如果是当前用户为管理ID
	//resultUser.Permissions = []string{"*/*/*"}
	if userInfo.IsAdmin == 1 {
		resultUser.Permissions = []string{"*/*/*"}
	} else {
		permissions, err := s.GetPermissions(roleIds)
		if err != nil {
			return nil, err
		}
		resultUser.Permissions = permissions
	}
	return resultUser, nil
}

// GetUsersRoleDept 获取多个用户角色 部门信息
func (s *SysUserService) GetUsersRoleDept(userList []*entitys.SysUserEntitys) ([]*entitys.SysUserRoleDeptRes, error) {
	roleService := SysRoleService{}
	deptService := SysDeptService{}
	allRoles, _, err := roleService.QuerySysRoleList(entitys.SysRoleQuery{
		Page:  1,
		Limit: 1000,
	})
	if err != nil {
		g.Log().Error(err)
		return nil, err
	}
	depts, _, err := deptService.QuerySysDeptList(entitys.SysDeptQuery{
		//Page:  1,
		//Limit: 1000,
	})

	if err != nil {
		g.Log().Error(err)
		return nil, err
	}
	users := make([]*entitys.SysUserRoleDeptRes, len(userList))
	for k, u := range userList {
		var dept *entitys.SysDeptEntitys
		users[k] = &entitys.SysUserRoleDeptRes{
			SysUserEntitys: u,
		}
		for _, d := range depts {
			if iotutil.ToString(u.DeptId) == d.DeptId {
				dept = d
			}
		}
		users[k].Dept = dept
		roles, err := s.GetAdminRole(iotutil.ToInt64(u.Id), allRoles)
		if err != nil {
			return nil, err
		}
		for _, r := range roles {
			users[k].RoleInfo = append(users[k].RoleInfo, &struct {
				RoleId int64  `json:"roleId"`
				Name   string `json:"name"`
			}{RoleId: iotutil.ToInt64(r.Id), Name: r.Name})
		}

		postIds, posts, err := s.GetUserPostInfo(iotutil.ToInt64(u.Id))
		if err != nil && err.Error() != "record not found" {
			return nil, err
		}
		for _, p := range posts {
			users[k].Post = append(users[k].Post, &struct {
				PostId   string `json:"postId"`
				PostName string `json:"postName"`
			}{PostId: p.PostId, PostName: p.PostName})
			//把岗位id加入集合. 方便前端使用岗位id 进行搜索
			users[k].PostIds = postIds
		}

	}
	return users, nil
}

// GetAdminRole 通过用户获取用户的角色信息
func (s SysUserService) GetAdminRole(userid int64, allRoleList []*entitys.SysRoleEntitys) (roles []*entitys.SysRoleEntitys, err error) {
	res, err := rpc.ClientCasbinExtService.GetUserRole(context.Background(), &protosService.UserRoleExtReq{
		UserId: iotutil.ToString(userid),
	})
	if err != nil {
		return nil, err
	}

	roleIds := res.Ids
	roles = make([]*entitys.SysRoleEntitys, 0, len(allRoleList))
	for _, v := range allRoleList {
		for _, id := range roleIds {
			if id == iotutil.ToString(v.Id) {
				roles = append(roles, v)
			}
		}
		if len(roles) == len(roleIds) {
			break
		}
	}
	return
}

// GetUserPostInfo 通过用户Id获取用户的岗位信息
func (s SysUserService) GetUserPostInfo(userid int64) (postIds []string, posts []*entitys.SysPostEntitys, err error) {
	res, err := rpc.ClientSysUserPostService.Lists(context.Background(), &protosService.SysUserPostListRequest{
		Query: &protosService.SysUserPost{
			UserId: userid,
		},
	})
	if err != nil {
		return nil, nil, err
	}
	if res.Code != 200 {
		return nil, nil, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) == 0 {
		return nil, nil, errors.New("record not found")
	}

	postIds = make([]string, 0, len(res.Data))
	postIdMaps := map[int64]int64{}
	for _, v := range res.Data {
		if _, ok := postIdMaps[v.PostId]; !ok {
			postIdMaps[v.PostId] = 1
		}
	}
	for k, _ := range postIdMaps {
		postIds = append(postIds, iotutil.ToString(k))
	}

	postService := SysPostService{}
	allPosts, _, err := postService.QuerySysPostList(entitys.SysPostQuery{})
	if err != nil {
		g.Log().Error(err)
		return nil, nil, err
	}

	posts = make([]*entitys.SysPostEntitys, 0)
	for _, post := range allPosts {
		postid := iotutil.ToInt64(post.PostId)
		if _, ok := postIdMaps[postid]; ok {
			posts = append(posts, post)
		}
	}
	return
}

// AddUser 新增用户
func (s SysUserService) AddUser(req entitys.UserCreateReq) (string, error) {
	//添加用户
	userid := iotutil.GetNextSeqInt64()
	req.Id = iotutil.ToString(userid)
	//生成密码盐值
	req.UserSalt = iotutil.GetSecret(10)
	req.UserPassword = iotutil.Md5(req.UserPassword + req.UserSalt)

	res, err := rpc.ClientSysUserService.Create(context.Background(), entitys.SysUserAdd_e2pb(&req))

	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	//保存用户角色
	err = s.SaveUserRole(userid, req.RoleIds, false)
	if err != nil {
		return "", err
	}

	//保存用户岗位
	err = s.SaveUserPost(userid, req.PostIds, false)
	if err != nil {
		return "", err
	}

	return iotutil.ToString(userid), nil
}

func (s SysUserService) SaveUserPost(userid int64, postIds []string, isDelete bool) error {
	var err error
	//1.删除原有岗位
	if isDelete {
		s.DeleteUserPost(userid)
	}
	//保存用户岗位
	if len(postIds) > 0 {
		for _, postId := range postIds {
			_, err = rpc.ClientSysUserPostService.Create(context.Background(), &protosService.SysUserPost{
				UserId: userid,
				PostId: iotutil.ToInt64(postId),
			})
			if err != nil {
				iotlogger.LogHelper.Error("rpc create userpost error--->" + err.Error())
				break
			}
		}
	}
	return err
}

// DeleteUserPost 删除用户岗位
func (s SysUserService) DeleteUserPost(userid int64) error {
	res, err := rpc.ClientSysUserPostService.Delete(context.Background(), &protosService.SysUserPost{
		UserId: userid,
	})
	if err != nil {
		iotlogger.LogHelper.Error("rpc delete userpost error--->" + err.Error())
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}

	return nil
}

// SaveUserRole 保存用户角色
func (s SysUserService) SaveUserRole(userid int64, roleids []string, isDelete bool) error {
	//是否删除原来用户角色[新增用false,修改用true]
	if isDelete {
		s.DeleteUserRole(userid)
	}

	if len(roleids) > 0 {
		for _, id := range roleids {
			rpc.ClientCasbinExtService.AddUserRole(context.Background(), &protosService.CasbinReq{
				UserId: iotutil.ToString(userid),
				RoleId: id,
			})
		}
	}
	return nil

}

func (s *SysUserService) GetPermissions(roleIds []int64) ([]string, error) {
	res, err := rpc.ClientCasbinExtService.GetPermissionsByRoleIds(context.Background(), &protosService.PermissionsByRoleIdReq{
		RoleIds: roleIds,
	})
	if err != nil {
		return nil, err
	}

	menuIds := map[int64]int64{}
	for _, id := range res.MenuIds {
		menuIds[id] = id
	}
	//获取所有开启的按钮
	allButtons, err := s.GetIsButtonStatusList()
	userButtons := make([]string, 0, len(allButtons))
	for _, button := range allButtons {
		if _, ok := menuIds[iotutil.ToInt64(button.Id)]; button.Condition == "nocheck" || ok {
			userButtons = append(userButtons, button.Name)
		}
	}
	return userButtons, nil
}

func (s *SysUserService) GetIsButtonStatusList() ([]*entitys.SysAuthRuleEntitys, error) {
	list, err := s.GetMenuList()
	if err != nil {
		return nil, err
	}
	var gList = make([]*entitys.SysAuthRuleEntitys, 0, len(list))
	for _, v := range list {
		if v.MenuType == 2 && v.Status == 1 {
			gList = append(gList, v)
		}
	}
	return gList, nil
}

// GetMenuList 获取菜单列表
func (s *SysUserService) GetMenuList() ([]*entitys.SysAuthRuleEntitys, error) {
	return SysAuthRuleService{}.GetMenuList(entitys.SysAuthRuleQuery{
		Title:  "",
		Status: 0,
	})
}

// DeleteUserRole 删除用户角色
func (s SysUserService) DeleteUserRole(userid int64) error {
	res, err := rpc.ClientCasbinExtService.DeleteUserRole(context.Background(), &protosService.CasbinReq{
		UserId: iotutil.ToString(userid),
	})
	if res.Code != 200 {
		return errors.New(res.Message)
	}

	return err
}

// UpdateUser 修改用户
func (s SysUserService) UpdateUser(req entitys.UserCreateReq) (string, error) {
	userid := iotutil.ToInt64(req.Id)

	res, err := rpc.ClientSysUserService.UpdateFields(context.Background(), &protosService.SysUserUpdateFieldsRequest{
		Fields: []string{"dept_id", "user_nickname", "phone_num", "user_email", "sex", "user_status", "remark", "address"},
		Data:   entitys.SysUserAdd_e2pb(&req),
	})

	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	//保存用户角色
	err = s.SaveUserRole(userid, req.RoleIds, true)
	if err != nil {
		return "", err
	}

	//保存用户岗位
	err = s.SaveUserPost(userid, req.PostIds, true)
	if err != nil {
		return "", err
	}

	//如果用户是禁用，则删除用户所有token
	//清除用户token
	//如果为禁用用户
	if req.UserStatus == 2 {
		controls.ClearTokenByUserId(userid)
	}
	return iotutil.ToString(userid), err
}

// UpdateUser 修改用户
func (s SysUserService) UpdateUserCenter(req entitys.UserCenterEditReq, userid int64) (string, error) {
	var user = protosService.SysUser{
		PhoneNum:  req.PhoneNum,
		UserEmail: req.Email,
		Avatar:    req.Avatar,
		Id:        userid,
	}

	res, err := rpc.ClientSysUserService.UpdateFields(context.Background(), &protosService.SysUserUpdateFieldsRequest{
		Fields: []string{"phone_num", "user_email", "avatar"},
		Data:   &user,
	})

	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(userid), err
}

// 重置用户密码
func (s SysUserService) UpdateResetPassword(req entitys.UserResetPasswordReq) error {

	user, errUser := rpc.ClientSysUserService.FindById(context.Background(), &protosService.SysUserFilter{
		Id: req.UserId,
	})
	if errUser != nil {
		return errUser
	}
	if user.Code != 200 {
		return errors.New(user.Message)
	}
	if len(user.Data) <= 0 {
		return errors.New("userid not exists")
	}
	salt := user.Data[0].UserSalt
	//password := iotutil.Md5(req.Password + salt)
	newPassword := iotutil.Md5(config.Global.Service.DefaultPassword)
	password := iotutil.Md5(newPassword + salt)
	res, err := rpc.ClientSysUserService.UpdateFields(context.Background(), &protosService.SysUserUpdateFieldsRequest{
		Fields: []string{"user_password", "updated_at"},
		Data:   &protosService.SysUser{Id: req.UserId, UserPassword: password, UpdatedAt: timestamppb.Now(), LastLoginTime: timestamppb.Now()},
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	//清除用户token
	controls.ClearTokenByUserId(req.UserId)
	return err
}

// 修改用户密码
func (s SysUserService) UpdatePassword(userId string, req entitys.UserUpdatePasswordReq) error {
	//todo 验证用户密码是否正确
	userid := iotutil.ToInt64(userId)
	//获取原用户信息
	user, errUser := rpc.ClientSysUserService.FindById(context.Background(), &protosService.SysUserFilter{
		Id: userid,
	})
	if errUser != nil {
		return errUser
	}
	if user.Code != 200 {
		return errors.New(user.Message)
	}
	if len(user.Data) <= 0 {
		return errors.New("账号不存在")
	}
	salt := user.Data[0].UserSalt
	old := iotutil.Md5(req.OldPassword + salt)
	if old != user.Data[0].UserPassword {
		return errors.New("旧密码不正确")
	}
	password := iotutil.Md5(req.NewPassword + salt)

	res, err := rpc.ClientSysUserService.UpdateFields(context.Background(), &protosService.SysUserUpdateFieldsRequest{
		Fields: []string{"user_password", "updated_at"},
		Data:   &protosService.SysUser{Id: userid, UserPassword: password, UpdatedAt: timestamppb.Now(), LastLoginTime: timestamppb.Now()},
	})

	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}

	//清除用户token
	controls.ClearTokenByUserId(userid)
	return err
}

// DeleteUser 删除用户
func (s SysUserService) DeleteUser(req entitys.DeleteCommonQuery) error {
	var err error
	if len(req.Ids) > 0 {
		for _, item := range req.Ids {
			id := iotutil.ToInt64(item)
			res, errUser := rpc.ClientSysUserService.DeleteById(context.Background(), &protosService.SysUser{
				Id: id,
			})
			if errUser != nil {
				iotlogger.LogHelper.Error("rpc delete userpost error--->" + errUser.Error())
				err = errUser
				break
			}
			if res.Code != 200 {
				err = errors.New(res.Message)
				break
			}

			//删除岗位
			errPost := s.DeleteUserPost(id)
			if errPost != nil {
				iotlogger.LogHelper.Error("rpc delete userpost error--->" + errPost.Error())
				err = errPost
				break
			}
			//删除角色
			errRole := s.DeleteUserRole(id)
			if errRole != nil {
				iotlogger.LogHelper.Error("rpc delete userrole error--->" + errRole.Error())
				err = errPost
				break
			}
			//清除用户token
			controls.ClearTokenByUserId(id)
		}

	}
	return err
}

// UserLogin 用户登录
func (s SysUserService) UserLogin(req entitys.UserLogin) (string, string, int64, error) {
	resp, err := rpc.ClientCloudAuthService.PasswordLogin(context.Background(), &protosService.PasswordLoginRequest{
		Channel:      req.Channel,
		LoginName:    req.Username,
		Password:     req.Password,
		VerifyCode:   req.Verifycode,
		ClientIp:     req.ClientIp,
		Explorer:     req.Explorer,
		Os:           req.Os,
		PlatformType: string(iotconst.PLATFORMTYPE_CLOUD),
	})
	if err != nil {
		return "", "", 0, errors.New(micorerrors.FromError(err).Detail)
	}
	if resp == nil || resp.UserInfo == nil {
		return "", "", 0, errors.New("内部服务错误")
	}
	userinfo := CloudUserInfo_pb2e(resp.GetUserInfo())
	expires := time.Unix(resp.ExpiresAt, 0).Sub(time.Now())
	err = cached.RedisStore.Set(resp.Token, *userinfo, expires)
	if err != nil {
		iotlogger.LogHelper.Errorf("UserLogin,缓存token失败:%s", err.Error())
	}
	controls.CacheTokenByUserId(userinfo.UserID, resp.Token, resp.GetExpiresAt())
	//for debug
	//var userInfo controls.UserInfo
	//err = cached.RedisStore.Get(resp.Token, &userInfo)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for debug
	return resp.GetToken(), resp.GetRefreshToken(), resp.GetExpiresAt(), nil
}

// UserLogout 用户登录
func (s SysUserService) UserLogout(token string) error {
	//删除redis即可
	cached.RedisStore.Delete(token)
	//删除在线信息
	rpc.ClientSysUserOnlineService.Delete(context.Background(), &protosService.SysUserOnline{Token: token})
	return nil
}

func CloudUserInfo_pb2e(src *protosService.CloudUserInfo) *controls.UserInfo {
	if src == nil {
		return nil
	}
	uiObj := controls.UserInfo{
		UserID:   src.UserId,
		Nickname: src.NickName,
		Avatar:   src.Avatar,
		DeptId:   src.DeptId,
		RoleIds:  src.RoleIds,
		PostIds:  src.PostIds,
	}
	return &uiObj
}

// UserLogin 用户登录
func (s SysUserService) RefreshToken(req entitys.RefreshToken) (string, string, int64, error) {
	resp, err := rpc.ClientCloudAuthService.RefreshToken(context.Background(), &protosService.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return "", "", 0, err
	}
	if resp == nil || (resp.GetValid() && resp.GetData() == nil) {
		return "", "", 0, errors.New("RefreshToken:内部服务错误")
	}
	if !resp.GetValid() {
		return "", "", 0, errors.New("RefreshToken:refresh token 已失效")
	}
	userinfo := CloudUserInfo_pb2e(resp.GetData().GetUserInfo())
	expires := time.Unix(resp.GetData().GetExpiresAt(), 0).Sub(time.Now())
	err = cached.RedisStore.Set(resp.GetData().GetToken(), *userinfo, expires)
	if err != nil {
		iotlogger.LogHelper.Errorf("RefreshToken:缓存token失败:%s", err.Error())
	}
	controls.CacheTokenByUserId(userinfo.UserID, resp.GetData().GetToken(), resp.GetData().GetExpiresAt())
	return resp.GetData().GetToken(), resp.GetData().GetRefreshToken(), resp.GetData().GetExpiresAt(), nil
}

func (s SysUserService) ForgetPassword(req entitys.UserResetPasswordNoTokenReq) (string, error) {
	//获取原用户信息
	user, errUser := rpc.ClientSysUserService.Find(context.Background(), &protosService.SysUserFilter{
		UserName:   req.Account,
		UserStatus: 1,
	})
	if errUser != nil {
		return "", errUser
	}
	if user.Code != 200 {
		if user.Message == "record not found" {
			return "", errors.New("用户名不存在")
		}
		return "", errors.New(user.Message)
	}
	if len(user.Data) <= 0 {
		return "", errors.New("用户名不存在")
	}
	isVer, verMsg := s.VerificationCode(req.Account, req.VerificationCode)
	if !isVer {
		return "", errors.New(verMsg)
	}

	salt := user.Data[0].UserSalt
	//password := iotutil.Md5(req.Password + salt)
	password := iotutil.Md5(req.Password + salt)
	res, err := rpc.ClientSysUserService.UpdateFields(context.Background(), &protosService.SysUserUpdateFieldsRequest{
		Fields: []string{"user_password", "updated_at"},
		Data:   &protosService.SysUser{Id: user.Data[0].Id, UserPassword: password, UpdatedAt: timestamppb.Now()},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", nil
}

// 发送验证码(会验证用户名是否存在) 忘记密码和登录的时候使用.
func (s SysUserService) SendVerificationCodeForExists(userName, tenantId, lang string, codeType int32) (string, error) {
	res, err := rpc.ClientSysUserService.Find(context.Background(), &protosService.SysUserFilter{
		UserName:   userName,
		UserStatus: 1,
	})
	if err != nil {
		return "", err
	}

	if res.Code != 200 {
		if res.Message == "record not found" {
			return "", errors.New("用户名不存在")
		}
		return "", errors.New(res.Message)
	}
	if len(res.Data) == 0 {
		return "", errors.New("用户名不存在")
	}
	return s.SendVerificationCode(userName, tenantId, lang, codeType)
}

// 发送验证码
func (s SysUserService) SendVerificationCode(userName, tenantId, lang string, codeType int32) (string, error) {
	index := strings.Index(userName, "@")
	code := iotutil.GetRandomNumber(6) //租户Id
	if index > 0 {
		res, err := rpc.ClientEmailService.SendEmailUserCode(context.Background(), &protosService.SendEmailUserCodeRequest{
			Email:    userName,
			UserName: userName,
			Code:     code,
			Lang:     lang,
			TplType:  codeType,
			TenantId: tenantId,
		})
		if err != nil {
			return "", err
		}
		if !res.Status {
			return "error", errors.New("验证码获取错误")
		}
	} else {
		res, err := rpc.ClientSmsService.SendSMSVerifyCode(context.Background(), &protosService.SendSMSVerifyCodeRequest{
			PhoneNumber: userName,
			UserName:    userName,
			Code:        code,
			Lang:        lang,
			TplType:     codeType,
			PhoneType:   1,
			TenantId:    tenantId,
		})
		if err != nil {
			return "", err
		}
		if !res.Status {
			return "error", errors.New("验证码获取错误")
		}

	}
	//验证码10分钟失效.
	expires := time.Minute * 10
	//避免开发平台和云管的数据冲突了.
	key := "sys_" + userName
	cached.RedisStore.Set(key, code, expires)
	return "ok", nil
}

// 验证验证码
func (s SysUserService) VerificationCode(userName string, code string) (bool, string) {
	//TODO 方便测试
	//if code == "999999" {
	//	return true, ""
	//}
	var codeRedis string
	//避免开发平台和云管的数据冲突了.
	key := "sys_" + userName
	cached.RedisStore.Get(key, &codeRedis)
	if codeRedis == "" {
		return false, "验证码已失效"
	}

	if codeRedis != code {
		return false, "验证码错误"
	}
	return true, ""
}
