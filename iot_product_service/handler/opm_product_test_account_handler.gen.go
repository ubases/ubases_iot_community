// Code generated by sgen,2024-03-11 16:22:42. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
    "context"
    proto "cloud_platform/iot_proto/protos/protosService"
    "cloud_platform/iot_product_service/service"
)

type OpmProductTestAccountHandler struct{}

//创建
func (h *OpmProductTestAccountHandler) Create(ctx context.Context, req *proto.OpmProductTestAccount,resp *proto.Response)  error {
    s := service.OpmProductTestAccountSvc{Ctx: ctx}
	_, err := s.CreateOpmProductTestAccount(req)
    SetResponse(resp, err)
	return nil
}

//匹配多条件删除
func (h *OpmProductTestAccountHandler) Delete(ctx context.Context, req *proto.OpmProductTestAccount,resp *proto.Response) error {
    s := service.OpmProductTestAccountSvc{Ctx: ctx}
	_, err := s.DeleteOpmProductTestAccount(req)
    SetResponse(resp, err)
	return nil
}

//匹配ID删除
func (h *OpmProductTestAccountHandler) DeleteById(ctx context.Context, req *proto.OpmProductTestAccount,resp *proto.Response) error {
    s := service.OpmProductTestAccountSvc{Ctx: ctx}
	_, err := s.DeleteByIdOpmProductTestAccount(req)
    SetResponse(resp, err)
	return nil
}

//根据主键批量删除
func (h *OpmProductTestAccountHandler) DeleteByIds(ctx context.Context, req *proto.OpmProductTestAccountBatchDeleteRequest,resp *proto.Response) error {
    s := service.OpmProductTestAccountSvc{Ctx: ctx}
	_, err := s.DeleteByIdsOpmProductTestAccount(req)
    SetResponse(resp, err)
	return nil
}

//更新
func (h *OpmProductTestAccountHandler) Update(ctx context.Context, req *proto.OpmProductTestAccount,resp *proto.Response) error {
    s := service.OpmProductTestAccountSvc{Ctx: ctx}
	_, err := s.UpdateOpmProductTestAccount(req)
    SetResponse(resp, err)
	return nil
}

//根据主键更新所有字段
func (h *OpmProductTestAccountHandler) UpdateAll(ctx context.Context, req *proto.OpmProductTestAccount,resp *proto.Response) error {
    s := service.OpmProductTestAccountSvc{Ctx: ctx}
	_, err := s.UpdateAllOpmProductTestAccount(req)
    SetResponse(resp, err)
	return nil
}

//根据主键更新指定列
func (h *OpmProductTestAccountHandler) UpdateFields(ctx context.Context, req *proto.OpmProductTestAccountUpdateFieldsRequest,resp *proto.Response) error {
    s := service.OpmProductTestAccountSvc{Ctx: ctx}
	_, err := s.UpdateFieldsOpmProductTestAccount(req)
    SetResponse(resp, err)
	return nil
}

//多条件查找，返回单条数据
func (h *OpmProductTestAccountHandler) Find(ctx context.Context, req *proto.OpmProductTestAccountFilter,resp *proto.OpmProductTestAccountResponse)  error {
    s := service.OpmProductTestAccountSvc{Ctx: ctx}
	data, err := s.FindOpmProductTestAccount(req)
    h.SetResponse(resp, data, err)
	return nil
}

//根据ID查找，返回单条数据
func (h *OpmProductTestAccountHandler) FindById(ctx context.Context, req *proto.OpmProductTestAccountFilter,resp *proto.OpmProductTestAccountResponse) error {
    s := service.OpmProductTestAccountSvc{Ctx: ctx}
	data, err := s.FindByIdOpmProductTestAccount(req)
	h.SetResponse(resp, data, err)
	return nil
}

//查找，支持分页，可返回多条数据
func (h *OpmProductTestAccountHandler) Lists(ctx context.Context, req *proto.OpmProductTestAccountListRequest,resp *proto.OpmProductTestAccountResponse) error {
    s := service.OpmProductTestAccountSvc{Ctx: ctx}
    data, total, err := s.GetListOpmProductTestAccount(req)
    h.SetPageResponse(resp, data, total, err)
	return nil
}


func (h *OpmProductTestAccountHandler) SetResponse(resp *proto.OpmProductTestAccountResponse, data *proto.OpmProductTestAccount, err error) {
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

func (h *OpmProductTestAccountHandler) SetPageResponse(resp *proto.OpmProductTestAccountResponse, list []*proto.OpmProductTestAccount, total int64, err error) {
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
