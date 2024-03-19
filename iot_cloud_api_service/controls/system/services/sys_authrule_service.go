package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"

	"go-micro.dev/v4/logger"
)

type SysAuthRuleService struct {
}

func (s SysAuthRuleService) QueryUserRouters(userid int64) ([]*entitys.SysAuthRuleTreeRes, error) {
	//2.获取用户所有角色
	res, err := rpc.ClientCasbinExtService.GetUserRole(context.Background(), &protosService.UserRoleExtReq{
		UserId: iotutil.ToString(userid),
	})

	if err != nil {
		return nil, err
	}
	if res == nil || res.Ids == nil {
		return nil, errors.New("该用户未设置角色,请联系管理员.")
	}

	iSSuperAdmin := false
	resUser, errUser := rpc.ClientSysUserService.FindById(context.Background(), &protosService.SysUserFilter{Id: userid})
	if errUser != nil || len(resUser.Data) == 0 {
		return nil, errors.New("用户不存在")
	}
	if resUser.Data[0].UserName == "admin" {
		iSSuperAdmin = true
	}

	//存放角色所有的菜单.
	var MenuIds = make([]int64, 0)
	if len(res.Ids) > 0 {
		//3.获取所有角色的菜单
		for _, roleId := range res.Ids {
			resMenuIds, _ := rpc.ClientCasbinExtService.GetRoleMenuList(context.Background(), &protosService.CasbinRoleMenuReq{
				RoleId: roleId,
			})
			if len(resMenuIds.MenuIds) > 0 {
				for _, v := range resMenuIds.MenuIds {
					//避免重复加入
					if iotutil.ArraysExistsInt64(MenuIds, v) == false {
						MenuIds = append(MenuIds, v)
					}
				}
			}
		}
	}
	//获取所有菜单
	rep1, err := rpc.ClientSysAuthRuleService.Lists(context.Background(), &protosService.SysAuthRuleListRequest{
		Page:     1,
		PageSize: 10000000,
		Query: &protosService.SysAuthRule{
			Title:  "",
			Status: 1,
		},
	})
	if err != nil {
		return nil, err
	}
	//4.过滤用户菜单

	var resultList = []*entitys.SysAuthRuleEntitys{}
	for _, item := range rep1.Data {
		//如果是超级管理员,则有所有权限菜单.
		if iotutil.ArraysExistsInt64(MenuIds, item.Id) == false && iSSuperAdmin {
			MenuIds = append(MenuIds, item.Id)
		}
		resultList = append(resultList, entitys.SysAuthRule_pb2e(item))
	}
	//菜单设置tree

	list := s.GetMenuListTreeFilterMenuList(0, resultList, MenuIds)
	return list, nil

}

// 菜单详情
func (s SysAuthRuleService) GetAuthRuleDetail(id string) (*entitys.SysAuthRuleEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientSysAuthRuleService.FindById(context.Background(), &protosService.SysAuthRuleFilter{Id: rid})
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
	return entitys.SysAuthRule_pb2e(data), err
}

// QueryAuthRuleList 菜单列表
func (s SysAuthRuleService) QueryAuthRuleList(filter entitys.SysAuthRuleQuery) ([]*entitys.SysAuthRuleTreeRes, error) {
	rep1, err := rpc.ClientSysAuthRuleService.Lists(context.Background(), &protosService.SysAuthRuleListRequest{
		Page:     1,
		PageSize: 10000000,
		Query: &protosService.SysAuthRule{
			Title:  filter.Title,
			Status: int32(filter.Status),
		},
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	if rep1.Code != 200 {
		return nil, errors.New(rep1.Message)
	}
	if rep1.Data == nil || len(rep1.Data) == 0 {
		return nil, errors.New("未查到菜单数据.")
	}

	if len(rep1.Data) > 0 {
		if filter.Status > 0 || filter.Title != "" {
			list := make([]*entitys.SysAuthRuleTreeRes, 0, len(rep1.Data))
			for _, menu := range rep1.Data {
				menu2 := entitys.SysAuthRule_pb2e(menu)
				list = append(list, &entitys.SysAuthRuleTreeRes{
					SysAuthRuleEntitys: menu2,
				})
			}
			return list, nil

		} else {
			var resultList = []*entitys.SysAuthRuleEntitys{}
			for _, item := range rep1.Data {
				resultList = append(resultList, entitys.SysAuthRule_pb2e(item))
			}
			//菜单设置tree
			list := s.GetMenuListTree(0, resultList)
			return list, nil
		}
	}
	return nil, err
}

// 角色详情获取菜单列表使用
func (s SysAuthRuleService) QueryRoleAuthRuleList(filter entitys.SysAuthRuleQuery) ([]*entitys.SysAuthRuleTreeRes, error) {
	rep1, err := rpc.ClientSysAuthRuleService.Lists(context.Background(), &protosService.SysAuthRuleListRequest{
		Page:     1,
		PageSize: 10000000,
		Query: &protosService.SysAuthRule{
			Title:  filter.Title,
			Status: int32(filter.Status),
		},
	})
	if err != nil {
		return nil, err
	}

	if rep1.Code != 200 {
		return nil, errors.New(rep1.Message)
	}
	if rep1.Data == nil || len(rep1.Data) == 0 {
		return nil, errors.New("record not found")
	}

	if len(rep1.Data) > 0 {

		var resultList = []*entitys.SysAuthRuleEntitys{}
		for _, item := range rep1.Data {
			resultList = append(resultList, entitys.SysAuthRule_pb2e(item))
		}
		//菜单设置tree
		list := s.GetMenuListTree(0, resultList)
		return list, nil

	}
	return nil, err
}

// 获取所有菜单,不做树.
func (s SysAuthRuleService) GetMenuList(filter entitys.SysAuthRuleQuery) ([]*entitys.SysAuthRuleEntitys, error) {
	rep1, err := rpc.ClientSysAuthRuleService.Lists(context.Background(), &protosService.SysAuthRuleListRequest{
		Page:     1,
		PageSize: 10000000,
		Query: &protosService.SysAuthRule{
			Title:  filter.Title,
			Status: int32(filter.Status),
		},
	})
	if err != nil {
		return nil, err
	}
	if rep1.Code != 200 {
		return nil, errors.New(rep1.Message)
	}
	if rep1.Data == nil || len(rep1.Data) == 0 {
		return nil, errors.New("record not found")
	}

	list := make([]*entitys.SysAuthRuleEntitys, 0, len(rep1.Data))
	for _, menu := range rep1.Data {
		menu2 := entitys.SysAuthRule_pb2e(menu)
		list = append(list, menu2)
	}
	return list, err
}

func (s SysAuthRuleService) GetMenuListTree(pid int64, list []*entitys.SysAuthRuleEntitys) []*entitys.SysAuthRuleTreeRes {
	tree := make([]*entitys.SysAuthRuleTreeRes, 0, len(list))
	for _, menu := range list {
		if menu.Pid == iotutil.ToString(pid) {
			t := &entitys.SysAuthRuleTreeRes{
				SysAuthRuleEntitys: menu,
			}
			child := s.GetMenuListTree(iotutil.ToInt64(menu.Id), list)
			if child != nil {
				t.Children = child
			}
			tree = append(tree, t)
		}
	}
	return tree
}

// 只有Menuids有的菜单. 才加入菜单树
func (s SysAuthRuleService) GetMenuListTreeFilterMenuList(pid int64, list []*entitys.SysAuthRuleEntitys, MenuIds []int64) []*entitys.SysAuthRuleTreeRes {
	tree := make([]*entitys.SysAuthRuleTreeRes, 0, len(list))
	for _, menu := range list {
		if menu.Pid == iotutil.ToString(pid) {
			t := &entitys.SysAuthRuleTreeRes{
				SysAuthRuleEntitys: menu,
			}
			child := s.GetMenuListTreeFilterMenuList(iotutil.ToInt64(menu.Id), list, MenuIds)
			if child != nil {
				t.Children = child
			}
			//有权限的,才加入菜单树
			if iotutil.ArraysExistsInt64(MenuIds, iotutil.ToInt64(menu.Id)) {
				tree = append(tree, t)
			}

		}
	}
	return tree
}

// AddAuthRule 新增菜单
func (s SysAuthRuleService) AddAuthRule(req entitys.SysAuthRuleEntitys) (string, error) {
	req.Id = iotutil.ToString(iotutil.GetNextSeqInt64())
	//req.CreatedAt = time.Now()
	res, err := rpc.ClientSysAuthRuleService.Create(context.Background(), entitys.SysAuthRule_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 修改菜单
func (s SysAuthRuleService) UpdateAuthRule(req entitys.SysAuthRuleEntitys) (string, error) {
	//req.UpdatedAt = time.Now()
	res, err := rpc.ClientSysAuthRuleService.UpdateAll(context.Background(), entitys.SysAuthRule_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除菜单
func (s SysAuthRuleService) DeleteAuthRule(req entitys.DeleteCommonQuery) error {
	var err error
	for _, id := range req.Ids {

		//把所有角色关联的菜单id都删掉
		_, errRoleDel := rpc.ClientCasbinExtService.DeleteAllRoleMenuById(context.Background(), &protosService.CasbinReq{
			MenuId: id,
		})
		if errRoleDel != nil {
			logger.Error("delete menuid error")
		}

		res, errDel := rpc.ClientSysAuthRuleService.Delete(context.Background(), &protosService.SysAuthRule{
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
