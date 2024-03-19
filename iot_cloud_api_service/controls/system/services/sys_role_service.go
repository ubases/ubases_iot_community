package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type SysRoleService struct {
}

// 角色详情
func (s SysRoleService) GetSysRoleDetail(id string) (*entitys.SysRoleDetailRes, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientSysRoleService.FindById(context.Background(), &protosService.SysRoleFilter{Id: rid})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}
	if len(req.Data) == 0 {
		return nil, errors.New("not found")
	}
	var data = req.Data[0]

	var menu = SysAuthRuleService{}
	menulist, errMenu := menu.QueryRoleAuthRuleList(entitys.SysAuthRuleQuery{
		Title:  "",
		Status: 1,
	})
	if errMenu != nil {
		return nil, errMenu
	}

	var res = entitys.SysRoleDetailRes{}
	res.MenuList = menulist

	res.Role = *entitys.SysRole_pb2e(data)

	menuids, errMenuids := rpc.ClientCasbinExtService.GetRoleMenuList(context.Background(), &protosService.CasbinRoleMenuReq{
		RoleId: id,
	})
	if errMenuids != nil {
		return nil, errMenuids
	}

	var strMenuids = make([]string, 0)
	for _, v := range menuids.MenuIds {
		strMenuids = append(strMenuids, iotutil.ToString(v))
	}

	res.CheckedRules = strMenuids

	return &res, err
}

// 获取角色数据权限
func (s SysRoleService) RoleDeptTreeSelect(id string) (*entitys.SysRoleDeptTreeSelectRes, error) {
	roledept, err := rpc.ClientSysRoleDeptService.Lists(context.Background(), &protosService.SysRoleDeptListRequest{
		Page:     1,
		PageSize: 100000000,
		Query: &protosService.SysRoleDept{
			RoleId: iotutil.ToInt64(id),
		},
	})
	if err != nil {
		return nil, err
	}
	if roledept.Code != 200 {
		return nil, errors.New(roledept.Message)
	}
	if roledept.Data == nil || len(roledept.Data) == 0 {
		return nil, errors.New("record not found")
	}

	//获取角色关联的部门ids
	list := make([]string, 0, len(roledept.Data))
	for _, v := range roledept.Data {
		list = append(list, iotutil.ToString(v.DeptId))
	}
	//获取部门树
	var dept = SysDeptService{}
	treelist := dept.QuerySysDeptAllTree()

	var res = entitys.SysRoleDeptTreeSelectRes{}
	res.DeptIds = list
	res.DeptTree = treelist

	return &res, nil
}

// QuerySysRoleList 角色列表
func (s SysRoleService) QuerySysRoleList(filter entitys.SysRoleQuery) ([]*entitys.SysRoleEntitys, int64, error) {
	rep, err := rpc.ClientSysRoleService.Lists(context.Background(), &protosService.SysRoleListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query: &protosService.SysRole{
			Name:   filter.RoleName,
			Status: filter.Status,
		},
	})
	if err != nil && err.Error() != "record not found" {

		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	if rep.Data == nil || len(rep.Data) == 0 {
		var notdata = []*entitys.SysRoleEntitys{}
		return notdata, 0, nil
	}
	var resultList = []*entitys.SysRoleEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.SysRole_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddSysRole 新增角色
func (s SysRoleService) AddSysRole(req entitys.SysRoleAddEntitys) (string, error) {
	req.Id = iotutil.ToString(iotutil.GetNextSeqInt64())
	res, err := rpc.ClientSysRoleService.Create(context.Background(), entitys.SysRoleAdd_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	//角色菜单授权
	if len(req.MenuIds) > 0 {
		for _, v := range req.MenuIds {
			rpc.ClientCasbinExtService.AddRoleMenu(context.Background(), &protosService.CasbinReq{
				RoleId: iotutil.ToString(req.Id),
				MenuId: iotutil.ToString(v),
			})
		}
	}

	if len(req.DeptIds) > 0 {
		for _, v := range req.DeptIds {
			rpc.ClientSysRoleDeptService.Create(context.Background(), &protosService.SysRoleDept{
				RoleId: iotutil.ToInt64(req.Id),
				DeptId: iotutil.ToInt64(v),
			})
		}
	}

	return iotutil.ToString(req.Id), err
}

// 修改角色
func (s SysRoleService) UpdateSysRole(req entitys.SysRoleAddEntitys) (string, error) {
	res, err := rpc.ClientSysRoleService.UpdateAll(context.Background(), entitys.SysRoleAdd_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	//角色菜单授权
	if len(req.MenuIds) > 0 {
		rpc.ClientCasbinExtService.DeleteRoleMenu(context.Background(), &protosService.CasbinReq{
			RoleId: iotutil.ToString(req.Id),
		})
		for _, v := range req.MenuIds {
			rpc.ClientCasbinExtService.AddRoleMenu(context.Background(), &protosService.CasbinReq{
				RoleId: iotutil.ToString(req.Id),
				MenuId: v,
			})
		}
	}

	return iotutil.ToString(req.Id), err
}

// 修改角色数据权限
func (s SysRoleService) RoleDataScope(req entitys.SysRoleAddEntitys) (string, error) {

	//m := entitys.SysRoleAdd_e2pb(&req)
	res, err := rpc.ClientSysRoleService.UpdateFields(context.Background(), &protosService.SysRoleUpdateFieldsRequest{
		Fields: []string{"data_scope"},
		Data: &protosService.SysRole{
			DataScope: req.DataScope,
			Id:        iotutil.ToInt64(req.Id),
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	if req.DataScope == 2 {
		//删除所有部门
		rpc.ClientSysRoleDeptService.Delete(context.Background(), &protosService.SysRoleDept{
			RoleId: iotutil.ToInt64(req.Id),
		})
		//重新设置部门
		if len(req.DeptIds) > 0 {
			for _, v := range req.DeptIds {
				rpc.ClientSysRoleDeptService.Create(context.Background(), &protosService.SysRoleDept{
					RoleId: iotutil.ToInt64(req.Id),
					DeptId: iotutil.ToInt64(v),
				})
			}
		}
	}

	return iotutil.ToString(req.Id), err
}

// 修改角色状态
func (s SysRoleService) RoleStatusUpdate(req entitys.SysRoleStatusReq) (string, error) {
	res, err := rpc.ClientSysRoleService.UpdateFields(context.Background(), &protosService.SysRoleUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.SysRole{
			Status: req.Status,
			Id:     iotutil.ToInt64(req.RoleId),
		},
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.RoleId), err
}

// 删除角色
func (s SysRoleService) DeleteSysRole(ctx context.Context, req entitys.DeleteCommonQuery) error {
	var err error
	for _, id := range req.Ids {
		res, errDel := rpc.ClientSysRoleService.Delete(ctx, &protosService.SysRole{
			Id: iotutil.ToInt64(id),
		})
		if errDel != nil {
			err = errDel
			break
		}
		if res.Code != 200 {
			err = errors.New(res.Message)
			break
		}
	}
	return err
}
