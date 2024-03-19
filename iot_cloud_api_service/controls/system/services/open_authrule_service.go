package services

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type OpenAuthRuleService struct {
}

// 菜单详情
func (s OpenAuthRuleService) GetAuthRuleDetail(id string) (*entitys.OpenAuthRuleEntitys, error) {
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientOpenAuthRuleService.FindById(context.Background(), &protosService.OpenAuthRuleFilter{Id: rid})
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
	return entitys.OpenAuthRule_pb2e(data), err
}

// QueryAuthRuleList 菜单列表
func (s OpenAuthRuleService) QueryAuthRuleList(filter entitys.OpenAuthRuleQuery) ([]*entitys.OpenAuthRuleTreeRes, error) {
	rep1, err := rpc.ClientOpenAuthRuleService.Lists(context.Background(), &protosService.OpenAuthRuleListRequest{
		Page:     1,
		PageSize: 10000000,
		Query: &protosService.OpenAuthRule{
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
		return nil, errors.New("未查到记录")
	}

	if len(rep1.Data) > 0 {
		if filter.Status > 0 || filter.Title != "" {
			list := make([]*entitys.OpenAuthRuleTreeRes, 0, len(rep1.Data))
			for _, menu := range rep1.Data {
				menu2 := entitys.OpenAuthRule_pb2e(menu)
				list = append(list, &entitys.OpenAuthRuleTreeRes{
					OpenAuthRuleEntitys: menu2,
				})
			}
			return list, nil

		} else {
			var resultList = []*entitys.OpenAuthRuleEntitys{}
			for _, item := range rep1.Data {
				resultList = append(resultList, entitys.OpenAuthRule_pb2e(item))
			}
			//菜单设置tree
			list := s.GetMenuListTree(0, resultList)
			return list, nil
		}
	}
	return nil, err
}

// 获取所有菜单,不做树.
func (s OpenAuthRuleService) GetMenuList(filter entitys.OpenAuthRuleQuery) ([]*entitys.OpenAuthRuleEntitys, error) {
	rep1, err := rpc.ClientOpenAuthRuleService.Lists(context.Background(), &protosService.OpenAuthRuleListRequest{
		Page:     1,
		PageSize: 10000000,
		Query: &protosService.OpenAuthRule{
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

	list := make([]*entitys.OpenAuthRuleEntitys, 0, len(rep1.Data))
	for _, menu := range rep1.Data {
		menu2 := entitys.OpenAuthRule_pb2e(menu)
		list = append(list, menu2)
	}
	return list, err
}

func (s OpenAuthRuleService) GetMenuListTree(pid int64, list []*entitys.OpenAuthRuleEntitys) []*entitys.OpenAuthRuleTreeRes {
	tree := make([]*entitys.OpenAuthRuleTreeRes, 0, len(list))
	for _, menu := range list {
		if menu.Pid == iotutil.ToString(pid) {
			t := &entitys.OpenAuthRuleTreeRes{
				OpenAuthRuleEntitys: menu,
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

// AddAuthRule 新增菜单
func (s OpenAuthRuleService) AddAuthRule(req entitys.OpenAuthRuleEntitys) (string, error) {
	req.Id = iotutil.ToString(iotutil.GetNextSeqInt64())
	//req.CreatedAt = time.Now()
	res, err := rpc.ClientOpenAuthRuleService.Create(context.Background(), entitys.OpenAuthRule_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 修改菜单
func (s OpenAuthRuleService) UpdateAuthRule(req entitys.OpenAuthRuleEntitys) (string, error) {
	//req.UpdatedAt = time.Now()
	res, err := rpc.ClientOpenAuthRuleService.UpdateAll(context.Background(), entitys.OpenAuthRule_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(req.Id), err
}

// 删除菜单
func (s OpenAuthRuleService) DeleteAuthRule(req entitys.DeleteCommonQuery) error {
	var err error
	for _, id := range req.Ids {
		res, errDel := rpc.ClientOpenAuthRuleService.Delete(context.Background(), &protosService.OpenAuthRule{
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
