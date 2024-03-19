// Code generated by sgen,2023-07-10 11:45:48. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_product_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OpmThingModelRuleHandler struct{}

// 创建
func (h *OpmThingModelRuleHandler) UpdateStatus(ctx context.Context, req *proto.OpmThingModelRule, resp *proto.Response) error {
	s := service.OpmThingModelRuleSvc{Ctx: ctx}
	err := s.UpdateStatus(req)
	SetResponse(resp, err)
	return nil
}

// 创建
func (h *OpmThingModelRuleHandler) Create(ctx context.Context, req *proto.OpmThingModelRule, resp *proto.Response) error {
	s := service.OpmThingModelRuleSvc{Ctx: ctx}
	_, err := s.CreateOpmThingModelRule(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *OpmThingModelRuleHandler) Delete(ctx context.Context, req *proto.OpmThingModelRule, resp *proto.Response) error {
	s := service.OpmThingModelRuleSvc{Ctx: ctx}
	_, err := s.DeleteOpmThingModelRule(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *OpmThingModelRuleHandler) DeleteById(ctx context.Context, req *proto.OpmThingModelRule, resp *proto.Response) error {
	s := service.OpmThingModelRuleSvc{Ctx: ctx}
	_, err := s.DeleteByIdOpmThingModelRule(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *OpmThingModelRuleHandler) DeleteByIds(ctx context.Context, req *proto.OpmThingModelRuleBatchDeleteRequest, resp *proto.Response) error {
	s := service.OpmThingModelRuleSvc{Ctx: ctx}
	_, err := s.DeleteByIdsOpmThingModelRule(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *OpmThingModelRuleHandler) Update(ctx context.Context, req *proto.OpmThingModelRule, resp *proto.Response) error {
	s := service.OpmThingModelRuleSvc{Ctx: ctx}
	_, err := s.UpdateOpmThingModelRule(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *OpmThingModelRuleHandler) UpdateAll(ctx context.Context, req *proto.OpmThingModelRule, resp *proto.Response) error {
	s := service.OpmThingModelRuleSvc{Ctx: ctx}
	_, err := s.UpdateAllOpmThingModelRule(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *OpmThingModelRuleHandler) UpdateFields(ctx context.Context, req *proto.OpmThingModelRuleUpdateFieldsRequest, resp *proto.Response) error {
	s := service.OpmThingModelRuleSvc{Ctx: ctx}
	_, err := s.UpdateFieldsOpmThingModelRule(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *OpmThingModelRuleHandler) Find(ctx context.Context, req *proto.OpmThingModelRuleFilter, resp *proto.OpmThingModelRuleResponse) error {
	s := service.OpmThingModelRuleSvc{Ctx: ctx}
	data, err := s.FindOpmThingModelRule(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *OpmThingModelRuleHandler) FindById(ctx context.Context, req *proto.OpmThingModelRuleFilter, resp *proto.OpmThingModelRuleResponse) error {
	s := service.OpmThingModelRuleSvc{Ctx: ctx}
	data, err := s.FindByIdOpmThingModelRule(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *OpmThingModelRuleHandler) Lists(ctx context.Context, req *proto.OpmThingModelRuleListRequest, resp *proto.OpmThingModelRuleResponse) error {
	s := service.OpmThingModelRuleSvc{Ctx: ctx}
	data, total, err := s.GetListOpmThingModelRule(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *OpmThingModelRuleHandler) SetResponse(resp *proto.OpmThingModelRuleResponse, data *proto.OpmThingModelRule, err error) {
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
		resp.Total = 0
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		if data != nil {
			resp.Total = 1
			resp.Data = append(resp.Data, data)
		}
	}
}

func (h *OpmThingModelRuleHandler) SetPageResponse(resp *proto.OpmThingModelRuleResponse, list []*proto.OpmThingModelRule, total int64, err error) {
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.Total = total
		resp.Data = list
	}
}
