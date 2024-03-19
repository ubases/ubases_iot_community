package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"time"
)

type SysDeptService struct {
}

// 部门详细
func (s SysDeptService) GetSysDeptDetail(id string) (*entitys.SysDeptEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientSysDeptService.FindById(context.Background(), &protosService.SysDeptFilter{DeptId: rid})
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
	return entitys.SysDept_pb2e(data), err
}

// QuerySysDeptList 菜单列表
func (s SysDeptService) QuerySysDeptList(filter entitys.SysDeptQuery) ([]*entitys.SysDeptEntitys, int64, error) {
	rep1, err := rpc.ClientSysDeptService.Lists(context.Background(), &protosService.SysDeptListRequest{
		Query: &protosService.SysDept{
			DeptName: filter.DeptName,
			Status:   filter.Status,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep1.Code != 200 {
		return nil, 0, errors.New(rep1.Message)
	}
	var resultList = []*entitys.SysDeptEntitys{}
	for _, item := range rep1.Data {
		resultList = append(resultList, entitys.SysDept_pb2e(item))
	}
	return resultList, rep1.Total, err
}

// 获取所以启用的部门树
func (s SysDeptService) QuerySysDeptAllTree() []*entitys.SysDeptTreeRes {
	listdept, _, _ := s.QuerySysDeptList(entitys.SysDeptQuery{
		DeptName: "",
		Status:   "1",
	})
	ltree := s.GetDeptListTree(0, listdept)
	return ltree
}

// 递归设置部门树
func (s SysDeptService) GetDeptListTree(pid int64, list []*entitys.SysDeptEntitys) []*entitys.SysDeptTreeRes {
	tree := make([]*entitys.SysDeptTreeRes, 0, len(list))
	for _, v := range list {
		if v.ParentId == iotutil.ToString(pid) {
			t := &entitys.SysDeptTreeRes{
				SysDeptEntitys: v,
			}
			child := s.GetDeptListTree(iotutil.ToInt64(v.DeptId), list)
			if len(child) > 0 {
				t.Children = child
			}
			tree = append(tree, t)
		}
	}
	return tree
}

// AddSysDept 新增菜单
func (s SysDeptService) AddSysDept(req entitys.SysDeptEntitys) (string, error) {

	if req.OrderNum >= 10000000 {
		return "", errors.New("部门排序数值不可超过10000000")
	}

	if len(req.Leader) > 50 {
		return "", errors.New("部门责任人长度不可超过50")
	}

	req.DeptId = iotutil.ToString(iotutil.GetNextSeqInt64())
	req.CreatedAt = time.Now().Unix()
	res, err := rpc.ClientSysDeptService.Create(context.Background(), entitys.SysDept_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.DeptId), err
}

// 修改部门
func (s SysDeptService) UpdateSysDept(req entitys.SysDeptEntitys) (string, error) {
	req.UpdatedAt = time.Now().Unix()
	res, err := rpc.ClientSysDeptService.UpdateAll(context.Background(), entitys.SysDept_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.DeptId), err
}

// 删除部门
func (s SysDeptService) DeleteSysDept(req entitys.DeleteCommonQuery) error {
	var err error

	deptId := req.Ids[0]

	rep1, errSel := rpc.ClientSysDeptService.Lists(context.Background(), &protosService.SysDeptListRequest{
		Page:     1,
		PageSize: 100000000,
		Query:    &protosService.SysDept{},
	})
	if errSel != nil {
		return errSel
	}
	if rep1.Code != 200 {
		return errors.New(rep1.Message)
	}
	IsExistsParent := false
	isOneDept := false
	for _, v := range rep1.Data {
		if v.ParentId == iotutil.ToInt64(deptId) {
			IsExistsParent = true
			break
		}
		if v.DeptId == iotutil.ToInt64(deptId) && v.ParentId == 0 {
			isOneDept = true
			break
		}
	}
	if IsExistsParent {
		return errors.New("该部门存在下级部门. 不可删除")
	}
	if isOneDept {
		return errors.New("一级部门不可删除")
	}

	resUser, errUser := rpc.ClientSysUserService.Find(context.Background(), &protosService.SysUserFilter{
		DeptId: iotutil.ToInt64(deptId),
	})
	if errUser != nil && errUser.Error() != "record not found" {
		return errors.New("系统异常,请重试")
	}
	if resUser.Code != 200 && resUser.Message != "record not found" {
		return errors.New(resUser.Message)
	}
	if resUser.Data != nil && len(resUser.Data) > 0 {
		return errors.New("该部门下有关联账号,不可删除")
	}

	res, errDel := rpc.ClientSysDeptService.DeleteById(context.Background(), &protosService.SysDept{
		DeptId: iotutil.ToInt64(deptId),
	})
	if errDel != nil {
		err = errDel
	}
	if res.Code != 200 {
		err = errors.New(res.Message)
	}

	return err
}
