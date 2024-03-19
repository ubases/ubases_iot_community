// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: developer.proto

package protosService

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for DeveloperService service

func NewDeveloperServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DeveloperService service

type DeveloperService interface {
	//添加开发者
	Add(ctx context.Context, in *DeveloperEntitys, opts ...client.CallOption) (*Response, error)
	//修改开发者
	Edit(ctx context.Context, in *DeveloperEntitys, opts ...client.CallOption) (*Response, error)
	//查看详情
	Detail(ctx context.Context, in *DeveloperFilterReq, opts ...client.CallOption) (*DeveloperEntitys, error)
	//删除开发者
	Delete(ctx context.Context, in *DeveloperFilterReq, opts ...client.CallOption) (*Response, error)
	//开发者启用禁用
	SetStatus(ctx context.Context, in *DeveloperStatusReq, opts ...client.CallOption) (*Response, error)
	//查询开发者列表
	List(ctx context.Context, in *DeveloperListRequest, opts ...client.CallOption) (*DeveloperListResponse, error)
	//查询开发者列表（不包括角色和授权数量）
	BasicList(ctx context.Context, in *DeveloperListRequest, opts ...client.CallOption) (*DeveloperListResponse, error)
	//查询开发者公司列表
	//rpc ListCompany() returns () {}
	//重置密码
	ResetPassword(ctx context.Context, in *DeveloperResetPasswordReq, opts ...client.CallOption) (*Response, error)
}

type developerService struct {
	c    client.Client
	name string
}

func NewDeveloperService(name string, c client.Client) DeveloperService {
	return &developerService{
		c:    c,
		name: name,
	}
}

func (c *developerService) Add(ctx context.Context, in *DeveloperEntitys, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeveloperService.Add", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerService) Edit(ctx context.Context, in *DeveloperEntitys, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeveloperService.Edit", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerService) Detail(ctx context.Context, in *DeveloperFilterReq, opts ...client.CallOption) (*DeveloperEntitys, error) {
	req := c.c.NewRequest(c.name, "DeveloperService.Detail", in)
	out := new(DeveloperEntitys)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerService) Delete(ctx context.Context, in *DeveloperFilterReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeveloperService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerService) SetStatus(ctx context.Context, in *DeveloperStatusReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeveloperService.SetStatus", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerService) List(ctx context.Context, in *DeveloperListRequest, opts ...client.CallOption) (*DeveloperListResponse, error) {
	req := c.c.NewRequest(c.name, "DeveloperService.List", in)
	out := new(DeveloperListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerService) BasicList(ctx context.Context, in *DeveloperListRequest, opts ...client.CallOption) (*DeveloperListResponse, error) {
	req := c.c.NewRequest(c.name, "DeveloperService.BasicList", in)
	out := new(DeveloperListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerService) ResetPassword(ctx context.Context, in *DeveloperResetPasswordReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeveloperService.ResetPassword", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeveloperService service

type DeveloperServiceHandler interface {
	//添加开发者
	Add(context.Context, *DeveloperEntitys, *Response) error
	//修改开发者
	Edit(context.Context, *DeveloperEntitys, *Response) error
	//查看详情
	Detail(context.Context, *DeveloperFilterReq, *DeveloperEntitys) error
	//删除开发者
	Delete(context.Context, *DeveloperFilterReq, *Response) error
	//开发者启用禁用
	SetStatus(context.Context, *DeveloperStatusReq, *Response) error
	//查询开发者列表
	List(context.Context, *DeveloperListRequest, *DeveloperListResponse) error
	//查询开发者列表（不包括角色和授权数量）
	BasicList(context.Context, *DeveloperListRequest, *DeveloperListResponse) error
	//查询开发者公司列表
	//rpc ListCompany() returns () {}
	//重置密码
	ResetPassword(context.Context, *DeveloperResetPasswordReq, *Response) error
}

func RegisterDeveloperServiceHandler(s server.Server, hdlr DeveloperServiceHandler, opts ...server.HandlerOption) error {
	type developerService interface {
		Add(ctx context.Context, in *DeveloperEntitys, out *Response) error
		Edit(ctx context.Context, in *DeveloperEntitys, out *Response) error
		Detail(ctx context.Context, in *DeveloperFilterReq, out *DeveloperEntitys) error
		Delete(ctx context.Context, in *DeveloperFilterReq, out *Response) error
		SetStatus(ctx context.Context, in *DeveloperStatusReq, out *Response) error
		List(ctx context.Context, in *DeveloperListRequest, out *DeveloperListResponse) error
		BasicList(ctx context.Context, in *DeveloperListRequest, out *DeveloperListResponse) error
		ResetPassword(ctx context.Context, in *DeveloperResetPasswordReq, out *Response) error
	}
	type DeveloperService struct {
		developerService
	}
	h := &developerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DeveloperService{h}, opts...))
}

type developerServiceHandler struct {
	DeveloperServiceHandler
}

func (h *developerServiceHandler) Add(ctx context.Context, in *DeveloperEntitys, out *Response) error {
	return h.DeveloperServiceHandler.Add(ctx, in, out)
}

func (h *developerServiceHandler) Edit(ctx context.Context, in *DeveloperEntitys, out *Response) error {
	return h.DeveloperServiceHandler.Edit(ctx, in, out)
}

func (h *developerServiceHandler) Detail(ctx context.Context, in *DeveloperFilterReq, out *DeveloperEntitys) error {
	return h.DeveloperServiceHandler.Detail(ctx, in, out)
}

func (h *developerServiceHandler) Delete(ctx context.Context, in *DeveloperFilterReq, out *Response) error {
	return h.DeveloperServiceHandler.Delete(ctx, in, out)
}

func (h *developerServiceHandler) SetStatus(ctx context.Context, in *DeveloperStatusReq, out *Response) error {
	return h.DeveloperServiceHandler.SetStatus(ctx, in, out)
}

func (h *developerServiceHandler) List(ctx context.Context, in *DeveloperListRequest, out *DeveloperListResponse) error {
	return h.DeveloperServiceHandler.List(ctx, in, out)
}

func (h *developerServiceHandler) BasicList(ctx context.Context, in *DeveloperListRequest, out *DeveloperListResponse) error {
	return h.DeveloperServiceHandler.BasicList(ctx, in, out)
}

func (h *developerServiceHandler) ResetPassword(ctx context.Context, in *DeveloperResetPasswordReq, out *Response) error {
	return h.DeveloperServiceHandler.ResetPassword(ctx, in, out)
}
