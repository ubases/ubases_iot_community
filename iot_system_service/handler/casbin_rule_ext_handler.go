package handler

import (
	"cloud_platform/iot_system_service/service"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"

	"go-micro.dev/v4"
)

// The Register tSysCasbinRule handler.
func RegisterCasbinRuleExtHandler(service micro.Service) error {
	err := protosService.RegisterCasbinExtServiceHandler(service.Server(), new(CasbinRuleExtHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterCasbinRuleExtHandler发生错误:%s", err.Error())
	}
	return err
}

type CasbinRuleExtHandler struct{}

func (h CasbinRuleExtHandler) GetPermissionsByRoleIds(ctx context.Context, req *protosService.PermissionsByRoleIdReq, resp *protosService.PermissionsByRoleIdResp) error {
	//v0 用户id,  v1  角色id
	groupPolicy := service.Casbin_Enforcer.GetFilteredPolicy(0, iotutil.ToString(req.RoleIds))
	var menuIds []int64
	if len(groupPolicy) > 0 {
		menuIds = make([]int64, len(groupPolicy))
		//得到角色id的切片
		for k, v := range groupPolicy {
			menuIds[k] = iotutil.ToInt64(v[1])
		}
	}
	resp.MenuIds = menuIds
	return nil
}

func (h CasbinRuleExtHandler) GetUserRole(ctx context.Context, req *protosService.UserRoleExtReq, resp *protosService.UserRoleExtResp) error {
	//v0 用户id,  v1  角色id
	groupPolicy := service.Casbin_Enforcer.GetFilteredGroupingPolicy(0, iotutil.ToString(req.UserId))
	var roleIds []string
	if len(groupPolicy) > 0 {
		roleIds = make([]string, len(groupPolicy))
		//得到角色id的切片
		for k, v := range groupPolicy {
			roleIds[k] = iotutil.ToString(v[1])
		}
	}
	resp.Ids = roleIds
	return nil
}

func (h CasbinRuleExtHandler) GetRoleUser(ctx context.Context, req *protosService.RoleUserExtReq, resp *protosService.UserRoleExtResp) error {
	//v0 用户id,  v1  角色id
	groupPolicy := service.Casbin_Enforcer.GetFilteredGroupingPolicy(0, iotutil.ToString(req.RoleId))
	var userIds []string
	if len(groupPolicy) > 0 {
		userIds = make([]string, len(groupPolicy))
		//得到角色id的切片
		for k, v := range groupPolicy {
			userIds[k] = iotutil.ToString(v[0])
		}
	}
	resp.Ids = userIds
	return nil
}

func (h CasbinRuleExtHandler) GetRoleUsers(roleId int64) []string {
	//v0 用户id,  v1  角色id
	groupPolicy := service.Casbin_Enforcer.GetFilteredGroupingPolicy(1, iotutil.ToString(roleId))
	var userIds []string
	if len(groupPolicy) > 0 {
		userIds = make([]string, len(groupPolicy))
		//得到角色id的切片
		for k, v := range groupPolicy {
			userIds[k] = iotutil.ToString(v[0])
		}
	}
	return userIds
}

// 用户角色添加
func (CasbinRuleExtHandler) AddUserRole(ctx context.Context, request *protosService.CasbinReq, response *protosService.CasbinResponse) error {

	//v0 用户id,  v1  角色id
	b, err := service.Casbin_Enforcer.AddGroupingPolicy(request.UserId, request.RoleId)
	if err != nil {
		return err
	}
	if !b {
		return errors.New("add user role error")
	}

	return nil
}

// 用户角色删除
func (CasbinRuleExtHandler) DeleteUserRole(ctx context.Context, request *protosService.CasbinReq, response *protosService.CasbinResponse) error {

	//v0 用户id
	service.Casbin_Enforcer.RemoveFilteredGroupingPolicy(0, request.UserId)
	// if err != nil {
	// 	return err
	// }
	// if !b {
	// 	return errors.New("delete user role error")
	// }
	return nil
}

// 用户角色删除
func (CasbinRuleExtHandler) DeleteAllRoleMenuById(ctx context.Context, request *protosService.CasbinReq, response *protosService.CasbinResponse) error {
	service.Casbin_Enforcer.RemoveFilteredPolicy(1, request.MenuId)
	return nil
}

// 角色菜单添加
func (CasbinRuleExtHandler) AddRoleMenu(ctx context.Context, request *protosService.CasbinReq, response *protosService.CasbinResponse) error {

	//v0 角色id,  v1 菜单id
	b, err := service.Casbin_Enforcer.AddPolicy(request.RoleId, request.MenuId, "ALL")
	if err != nil {
		return err
	}
	if !b {
		return errors.New("add role menuid error")
	}
	return nil
}

// 角色菜单删除
func (CasbinRuleExtHandler) DeleteRoleMenu(ctx context.Context, request *protosService.CasbinReq, response *protosService.CasbinResponse) error {
	//v0 角色id,  v1 菜单id
	b, err := service.Casbin_Enforcer.RemoveFilteredPolicy(0, request.RoleId)
	if err != nil {
		return err
	}
	if !b {
		return errors.New("delte role menuid error")
	}
	return nil
}

func (CasbinRuleExtHandler) GetRoleMenuList(ctx context.Context, request *protosService.CasbinRoleMenuReq, response *protosService.CasbinRoleMenuRes) error {
	menuids := service.Casbin_Enforcer.GetFilteredNamedPolicy("p", 0, request.RoleId)
	for _, v := range menuids {
		//gpSlice[k] = gconv.Int(v[1])
		response.MenuIds = append(response.MenuIds, int64(iotutil.ToInt(v[1])))
	}
	return nil

}
