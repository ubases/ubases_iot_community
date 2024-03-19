package handler

import (
	"cloud_platform/iot_open_system_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
)

type DeveloperHandler struct{}

func (h *DeveloperHandler) Add(ctx context.Context, req *proto.DeveloperEntitys, resp *proto.Response) error {
	s := service.DeveloperSvc{Ctx: ctx}
	_, err := s.Add(req)
	SetResponse(resp, err)
	return nil
}

func (h *DeveloperHandler) Edit(ctx context.Context, req *proto.DeveloperEntitys, resp *proto.Response) error {
	s := service.DeveloperSvc{Ctx: ctx}
	_, err := s.Edit(req)
	SetResponse(resp, err)
	return nil
}

func (h *DeveloperHandler) Detail(ctx context.Context, req *proto.DeveloperFilterReq, resp *proto.DeveloperEntitys) error {
	s := service.DeveloperSvc{Ctx: ctx}
	ret, err := s.Detail(req)
	if err == nil {
		*resp = *ret
	}
	return nil
}

func (h *DeveloperHandler) Delete(ctx context.Context, req *proto.DeveloperFilterReq, resp *proto.Response) error {
	s := service.DeveloperSvc{Ctx: ctx}
	_, err := s.Delete(req)
	SetResponse(resp, err)
	return nil
}

func (h *DeveloperHandler) SetStatus(ctx context.Context, req *proto.DeveloperStatusReq, resp *proto.Response) error {
	s := service.DeveloperSvc{Ctx: ctx}
	_, err := s.SetStatus(req)
	SetResponse(resp, err)
	return nil
}

func (h *DeveloperHandler) List(ctx context.Context, req *proto.DeveloperListRequest, resp *proto.DeveloperListResponse) error {
	s := service.DeveloperSvc{Ctx: ctx}
	ret, total, err := s.List(req)
	if err == nil {
		resp.Data = ret
		resp.Code = 200
		resp.Total = total
	} else {
		resp.Data = ret
		resp.Code = 0
		resp.Message = err.Error()
		resp.Total = 0
	}
	return nil
}

func (h *DeveloperHandler) BasicList(ctx context.Context, req *proto.DeveloperListRequest, resp *proto.DeveloperListResponse) error {
	s := service.DeveloperSvc{Ctx: ctx}
	ret, total, err := s.BasicList(req)
	if err == nil {
		resp.Data = ret
		resp.Code = 200
		resp.Total = total
	} else {
		resp.Data = ret
		resp.Code = 0
		resp.Message = err.Error()
		resp.Total = 0
	}
	return nil
}

func (h *DeveloperHandler) ResetPassword(ctx context.Context, req *proto.DeveloperResetPasswordReq, resp *proto.Response) error {
	s := service.DeveloperSvc{Ctx: ctx}
	_, err := s.ResetPassword(req)
	SetResponse(resp, err)
	return nil
}
