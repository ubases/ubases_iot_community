package services

import (
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpenRoleService struct {
	Ctx context.Context
}

func (s OpenRoleService) SetContext(ctx context.Context) OpenRoleService {
	s.Ctx = ctx
	return s
}

func (s OpenRoleService) AddRole(req *entitys.OpenRoleAddReq) (string, error) {
	res, err := rpc.ClientOpenRoleService.RoleAdd(s.Ctx, &protosService.OpenRoleAddRequest{
		RoleName: req.RoleName,
		MenuIds:  req.MenuIds,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", nil
}

func (s OpenRoleService) EditRole(req *entitys.OpenRoleEditReq) (string, error) {

	res, err := rpc.ClientOpenRoleService.RoleEdit(s.Ctx, &protosService.OpenRoleEditRequest{
		Id:       req.Id,
		RoleName: req.RoleName,
		MenuIds:  req.MenuIds,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", nil
}
func (s OpenRoleService) DeleteRole(id string) (string, error) {
	res, err := rpc.ClientOpenRoleService.DeleteById(s.Ctx, &protosService.OpenRole{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", nil
}

func (s OpenRoleService) RoleDefaultList() (*protosService.OpenRoleResponse, error) {
	res, err := rpc.ClientOpenRoleService.Lists(s.Ctx, &protosService.OpenRoleListRequest{
		Query: &protosService.OpenRole{
			Status:    1,
			IsDefault: 1,
		},
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	return res, nil
}

func (s OpenRoleService) RoleList(tenantId string) (*[]entitys.OpenRoleListRes, error) {
	res, err := rpc.ClientOpenRoleService.Lists(s.Ctx, &protosService.OpenRoleListRequest{
		Query: &protosService.OpenRole{
			TenantId:  tenantId,
			Status:    1,
			IsDefault: 2,
		},
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}

	var reslist []entitys.OpenRoleListRes
	resDefault, errDefault := s.RoleDefaultList()
	if err != nil {
		return nil, errDefault
	}

	//默认的角色
	for _, r := range resDefault.Data {
		reslist = append(reslist, *entitys.OpenRole_pb2e(r))
	}

	//自定义角色
	for _, v := range res.Data {
		reslist = append(reslist, *entitys.OpenRole_pb2e(v))
	}
	return &reslist, nil
}

func (s OpenRoleService) RoleDetail(id string) (*entitys.OpenRoleDetailRes, error) {
	res, err := rpc.ClientOpenRoleService.RoleDetail(s.Ctx, &protosService.OpenRolePrimarykey{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}

	resRole := entitys.OpenRoleDetailRes{
		MenuIds:   res.Menuids,
		Name:      res.Role.Name,
		Status:    res.Role.Status,
		ListOrder: res.Role.ListOrder,
		IsDefault: res.Role.IsDefault,
	}
	return &resRole, nil
}

// 给用户分配角色
func (s OpenRoleService) RoleSetUser(req *entitys.OpenRoleSetUserReq) (string, error) {
	res, err := rpc.ClientOpenRoleService.RoleSetUser(s.Ctx, &protosService.OpenRoleSetUserRequest{
		UserName: req.UserName,
		RoleId:   req.RoleId,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return "success", nil
}
