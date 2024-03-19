// Code generated by sgen,2023-08-12 17:29:59. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_message_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type AppPushTokenHandler struct{}

// 创建
func (h *AppPushTokenHandler) Create(ctx context.Context, req *proto.AppPushToken, resp *proto.Response) error {
	s := service.AppPushTokenSvc{Ctx: ctx}
	_, err := s.CreateAppPushToken(req)
	SetResponse(resp, err)
	return nil
}

// 匹配多条件删除
func (h *AppPushTokenHandler) Delete(ctx context.Context, req *proto.AppPushToken, resp *proto.Response) error {
	s := service.AppPushTokenSvc{Ctx: ctx}
	_, err := s.DeleteAppPushToken(req)
	SetResponse(resp, err)
	return nil
}

// 匹配ID删除
func (h *AppPushTokenHandler) DeleteById(ctx context.Context, req *proto.AppPushToken, resp *proto.Response) error {
	s := service.AppPushTokenSvc{Ctx: ctx}
	_, err := s.DeleteByIdAppPushToken(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键批量删除
func (h *AppPushTokenHandler) DeleteByIds(ctx context.Context, req *proto.AppPushTokenBatchDeleteRequest, resp *proto.Response) error {
	s := service.AppPushTokenSvc{Ctx: ctx}
	_, err := s.DeleteByIdsAppPushToken(req)
	SetResponse(resp, err)
	return nil
}

// 更新
func (h *AppPushTokenHandler) Update(ctx context.Context, req *proto.AppPushToken, resp *proto.Response) error {
	s := service.AppPushTokenSvc{Ctx: ctx}
	_, err := s.UpdateAppPushToken(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新所有字段
func (h *AppPushTokenHandler) UpdateAll(ctx context.Context, req *proto.AppPushToken, resp *proto.Response) error {
	s := service.AppPushTokenSvc{Ctx: ctx}
	_, err := s.UpdateAllAppPushToken(req)
	SetResponse(resp, err)
	return nil
}

// 根据主键更新指定列
func (h *AppPushTokenHandler) UpdateFields(ctx context.Context, req *proto.AppPushTokenUpdateFieldsRequest, resp *proto.Response) error {
	s := service.AppPushTokenSvc{Ctx: ctx}
	_, err := s.UpdateFieldsAppPushToken(req)
	SetResponse(resp, err)
	return nil
}

// 多条件查找，返回单条数据
func (h *AppPushTokenHandler) Find(ctx context.Context, req *proto.AppPushTokenFilter, resp *proto.AppPushTokenResponse) error {
	s := service.AppPushTokenSvc{Ctx: ctx}
	data, err := s.FindAppPushToken(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *AppPushTokenHandler) FindById(ctx context.Context, req *proto.AppPushTokenFilter, resp *proto.AppPushTokenResponse) error {
	s := service.AppPushTokenSvc{Ctx: ctx}
	data, err := s.FindByIdAppPushToken(req)
	h.SetResponse(resp, data, err)
	return nil
}

// 查找，支持分页，可返回多条数据
func (h *AppPushTokenHandler) Lists(ctx context.Context, req *proto.AppPushTokenListRequest, resp *proto.AppPushTokenResponse) error {
	s := service.AppPushTokenSvc{Ctx: ctx}
	data, total, err := s.GetListAppPushToken(req)
	h.SetPageResponse(resp, data, total, err)
	return nil
}

func (h *AppPushTokenHandler) SetResponse(resp *proto.AppPushTokenResponse, data *proto.AppPushToken, err error) {
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

func (h *AppPushTokenHandler) SetPageResponse(resp *proto.AppPushTokenResponse, list []*proto.AppPushToken, total int64, err error) {
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
