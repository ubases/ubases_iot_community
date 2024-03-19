// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sys_region_server_service.gen.proto

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

// Api Endpoints for SysRegionServerService service

func NewSysRegionServerServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "SysRegionServerService.Create",
			Path:    []string{"/v1/sysRegionServer/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SysRegionServerService.Delete",
			Path:    []string{"/v1/sysRegionServer/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SysRegionServerService.DeleteById",
			Path:    []string{"/v1/sysRegionServer/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SysRegionServerService.DeleteByIds",
			Path:    []string{"/v1/sysRegionServer/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SysRegionServerService.Update",
			Path:    []string{"/v1/sysRegionServer/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SysRegionServerService.UpdateAll",
			Path:    []string{"/v1/sysRegionServer/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SysRegionServerService.UpdateFields",
			Path:    []string{"/v1/sysRegionServer/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SysRegionServerService.FindById",
			Path:    []string{"/v1/sysRegionServer/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SysRegionServerService.Find",
			Path:    []string{"/v1/sysRegionServer/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SysRegionServerService.Lists",
			Path:    []string{"/v1/sysRegionServer/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for SysRegionServerService service

type SysRegionServerService interface {
	//创建
	Create(ctx context.Context, in *SysRegionServer, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *SysRegionServer, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *SysRegionServer, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *SysRegionServerBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *SysRegionServer, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *SysRegionServer, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *SysRegionServerUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *SysRegionServerFilter, opts ...client.CallOption) (*SysRegionServerResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *SysRegionServerFilter, opts ...client.CallOption) (*SysRegionServerResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *SysRegionServerListRequest, opts ...client.CallOption) (*SysRegionServerResponse, error)
}

type sysRegionServerService struct {
	c    client.Client
	name string
}

func NewSysRegionServerService(name string, c client.Client) SysRegionServerService {
	return &sysRegionServerService{
		c:    c,
		name: name,
	}
}

func (c *sysRegionServerService) Create(ctx context.Context, in *SysRegionServer, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SysRegionServerService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRegionServerService) Delete(ctx context.Context, in *SysRegionServer, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SysRegionServerService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRegionServerService) DeleteById(ctx context.Context, in *SysRegionServer, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SysRegionServerService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRegionServerService) DeleteByIds(ctx context.Context, in *SysRegionServerBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SysRegionServerService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRegionServerService) Update(ctx context.Context, in *SysRegionServer, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SysRegionServerService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRegionServerService) UpdateAll(ctx context.Context, in *SysRegionServer, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SysRegionServerService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRegionServerService) UpdateFields(ctx context.Context, in *SysRegionServerUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SysRegionServerService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRegionServerService) FindById(ctx context.Context, in *SysRegionServerFilter, opts ...client.CallOption) (*SysRegionServerResponse, error) {
	req := c.c.NewRequest(c.name, "SysRegionServerService.FindById", in)
	out := new(SysRegionServerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRegionServerService) Find(ctx context.Context, in *SysRegionServerFilter, opts ...client.CallOption) (*SysRegionServerResponse, error) {
	req := c.c.NewRequest(c.name, "SysRegionServerService.Find", in)
	out := new(SysRegionServerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRegionServerService) Lists(ctx context.Context, in *SysRegionServerListRequest, opts ...client.CallOption) (*SysRegionServerResponse, error) {
	req := c.c.NewRequest(c.name, "SysRegionServerService.Lists", in)
	out := new(SysRegionServerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SysRegionServerService service

type SysRegionServerServiceHandler interface {
	//创建
	Create(context.Context, *SysRegionServer, *Response) error
	//匹配多条件删除
	Delete(context.Context, *SysRegionServer, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *SysRegionServer, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *SysRegionServerBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *SysRegionServer, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *SysRegionServer, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *SysRegionServerUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *SysRegionServerFilter, *SysRegionServerResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *SysRegionServerFilter, *SysRegionServerResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *SysRegionServerListRequest, *SysRegionServerResponse) error
}

func RegisterSysRegionServerServiceHandler(s server.Server, hdlr SysRegionServerServiceHandler, opts ...server.HandlerOption) error {
	type sysRegionServerService interface {
		Create(ctx context.Context, in *SysRegionServer, out *Response) error
		Delete(ctx context.Context, in *SysRegionServer, out *Response) error
		DeleteById(ctx context.Context, in *SysRegionServer, out *Response) error
		DeleteByIds(ctx context.Context, in *SysRegionServerBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *SysRegionServer, out *Response) error
		UpdateAll(ctx context.Context, in *SysRegionServer, out *Response) error
		UpdateFields(ctx context.Context, in *SysRegionServerUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *SysRegionServerFilter, out *SysRegionServerResponse) error
		Find(ctx context.Context, in *SysRegionServerFilter, out *SysRegionServerResponse) error
		Lists(ctx context.Context, in *SysRegionServerListRequest, out *SysRegionServerResponse) error
	}
	type SysRegionServerService struct {
		sysRegionServerService
	}
	h := &sysRegionServerServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SysRegionServerService.Create",
		Path:    []string{"/v1/sysRegionServer/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SysRegionServerService.Delete",
		Path:    []string{"/v1/sysRegionServer/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SysRegionServerService.DeleteById",
		Path:    []string{"/v1/sysRegionServer/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SysRegionServerService.DeleteByIds",
		Path:    []string{"/v1/sysRegionServer/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SysRegionServerService.Update",
		Path:    []string{"/v1/sysRegionServer/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SysRegionServerService.UpdateAll",
		Path:    []string{"/v1/sysRegionServer/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SysRegionServerService.UpdateFields",
		Path:    []string{"/v1/sysRegionServer/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SysRegionServerService.FindById",
		Path:    []string{"/v1/sysRegionServer/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SysRegionServerService.Find",
		Path:    []string{"/v1/sysRegionServer/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SysRegionServerService.Lists",
		Path:    []string{"/v1/sysRegionServer/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&SysRegionServerService{h}, opts...))
}

type sysRegionServerServiceHandler struct {
	SysRegionServerServiceHandler
}

func (h *sysRegionServerServiceHandler) Create(ctx context.Context, in *SysRegionServer, out *Response) error {
	return h.SysRegionServerServiceHandler.Create(ctx, in, out)
}

func (h *sysRegionServerServiceHandler) Delete(ctx context.Context, in *SysRegionServer, out *Response) error {
	return h.SysRegionServerServiceHandler.Delete(ctx, in, out)
}

func (h *sysRegionServerServiceHandler) DeleteById(ctx context.Context, in *SysRegionServer, out *Response) error {
	return h.SysRegionServerServiceHandler.DeleteById(ctx, in, out)
}

func (h *sysRegionServerServiceHandler) DeleteByIds(ctx context.Context, in *SysRegionServerBatchDeleteRequest, out *Response) error {
	return h.SysRegionServerServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *sysRegionServerServiceHandler) Update(ctx context.Context, in *SysRegionServer, out *Response) error {
	return h.SysRegionServerServiceHandler.Update(ctx, in, out)
}

func (h *sysRegionServerServiceHandler) UpdateAll(ctx context.Context, in *SysRegionServer, out *Response) error {
	return h.SysRegionServerServiceHandler.UpdateAll(ctx, in, out)
}

func (h *sysRegionServerServiceHandler) UpdateFields(ctx context.Context, in *SysRegionServerUpdateFieldsRequest, out *Response) error {
	return h.SysRegionServerServiceHandler.UpdateFields(ctx, in, out)
}

func (h *sysRegionServerServiceHandler) FindById(ctx context.Context, in *SysRegionServerFilter, out *SysRegionServerResponse) error {
	return h.SysRegionServerServiceHandler.FindById(ctx, in, out)
}

func (h *sysRegionServerServiceHandler) Find(ctx context.Context, in *SysRegionServerFilter, out *SysRegionServerResponse) error {
	return h.SysRegionServerServiceHandler.Find(ctx, in, out)
}

func (h *sysRegionServerServiceHandler) Lists(ctx context.Context, in *SysRegionServerListRequest, out *SysRegionServerResponse) error {
	return h.SysRegionServerServiceHandler.Lists(ctx, in, out)
}
