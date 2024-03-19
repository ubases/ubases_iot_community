// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: iot_ota_white_group_service.gen.proto

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

// Api Endpoints for IotOtaWhiteGroupService service

func NewIotOtaWhiteGroupServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "IotOtaWhiteGroupService.Create",
			Path:    []string{"/v1/iotOtaWhiteGroup/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaWhiteGroupService.Delete",
			Path:    []string{"/v1/iotOtaWhiteGroup/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaWhiteGroupService.DeleteById",
			Path:    []string{"/v1/iotOtaWhiteGroup/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaWhiteGroupService.DeleteByIds",
			Path:    []string{"/v1/iotOtaWhiteGroup/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaWhiteGroupService.Update",
			Path:    []string{"/v1/iotOtaWhiteGroup/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaWhiteGroupService.UpdateAll",
			Path:    []string{"/v1/iotOtaWhiteGroup/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaWhiteGroupService.UpdateFields",
			Path:    []string{"/v1/iotOtaWhiteGroup/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaWhiteGroupService.FindById",
			Path:    []string{"/v1/iotOtaWhiteGroup/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaWhiteGroupService.Find",
			Path:    []string{"/v1/iotOtaWhiteGroup/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaWhiteGroupService.Lists",
			Path:    []string{"/v1/iotOtaWhiteGroup/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for IotOtaWhiteGroupService service

type IotOtaWhiteGroupService interface {
	//创建
	Create(ctx context.Context, in *IotOtaWhiteGroup, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *IotOtaWhiteGroup, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *IotOtaWhiteGroup, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *IotOtaWhiteGroupBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *IotOtaWhiteGroup, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *IotOtaWhiteGroup, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *IotOtaWhiteGroupUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *IotOtaWhiteGroupFilter, opts ...client.CallOption) (*IotOtaWhiteGroupResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *IotOtaWhiteGroupFilter, opts ...client.CallOption) (*IotOtaWhiteGroupResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *IotOtaWhiteGroupListRequest, opts ...client.CallOption) (*IotOtaWhiteGroupResponse, error)
}

type iotOtaWhiteGroupService struct {
	c    client.Client
	name string
}

func NewIotOtaWhiteGroupService(name string, c client.Client) IotOtaWhiteGroupService {
	return &iotOtaWhiteGroupService{
		c:    c,
		name: name,
	}
}

func (c *iotOtaWhiteGroupService) Create(ctx context.Context, in *IotOtaWhiteGroup, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaWhiteGroupService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaWhiteGroupService) Delete(ctx context.Context, in *IotOtaWhiteGroup, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaWhiteGroupService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaWhiteGroupService) DeleteById(ctx context.Context, in *IotOtaWhiteGroup, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaWhiteGroupService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaWhiteGroupService) DeleteByIds(ctx context.Context, in *IotOtaWhiteGroupBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaWhiteGroupService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaWhiteGroupService) Update(ctx context.Context, in *IotOtaWhiteGroup, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaWhiteGroupService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaWhiteGroupService) UpdateAll(ctx context.Context, in *IotOtaWhiteGroup, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaWhiteGroupService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaWhiteGroupService) UpdateFields(ctx context.Context, in *IotOtaWhiteGroupUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaWhiteGroupService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaWhiteGroupService) FindById(ctx context.Context, in *IotOtaWhiteGroupFilter, opts ...client.CallOption) (*IotOtaWhiteGroupResponse, error) {
	req := c.c.NewRequest(c.name, "IotOtaWhiteGroupService.FindById", in)
	out := new(IotOtaWhiteGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaWhiteGroupService) Find(ctx context.Context, in *IotOtaWhiteGroupFilter, opts ...client.CallOption) (*IotOtaWhiteGroupResponse, error) {
	req := c.c.NewRequest(c.name, "IotOtaWhiteGroupService.Find", in)
	out := new(IotOtaWhiteGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaWhiteGroupService) Lists(ctx context.Context, in *IotOtaWhiteGroupListRequest, opts ...client.CallOption) (*IotOtaWhiteGroupResponse, error) {
	req := c.c.NewRequest(c.name, "IotOtaWhiteGroupService.Lists", in)
	out := new(IotOtaWhiteGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IotOtaWhiteGroupService service

type IotOtaWhiteGroupServiceHandler interface {
	//创建
	Create(context.Context, *IotOtaWhiteGroup, *Response) error
	//匹配多条件删除
	Delete(context.Context, *IotOtaWhiteGroup, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *IotOtaWhiteGroup, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *IotOtaWhiteGroupBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *IotOtaWhiteGroup, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *IotOtaWhiteGroup, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *IotOtaWhiteGroupUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *IotOtaWhiteGroupFilter, *IotOtaWhiteGroupResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *IotOtaWhiteGroupFilter, *IotOtaWhiteGroupResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *IotOtaWhiteGroupListRequest, *IotOtaWhiteGroupResponse) error
}

func RegisterIotOtaWhiteGroupServiceHandler(s server.Server, hdlr IotOtaWhiteGroupServiceHandler, opts ...server.HandlerOption) error {
	type iotOtaWhiteGroupService interface {
		Create(ctx context.Context, in *IotOtaWhiteGroup, out *Response) error
		Delete(ctx context.Context, in *IotOtaWhiteGroup, out *Response) error
		DeleteById(ctx context.Context, in *IotOtaWhiteGroup, out *Response) error
		DeleteByIds(ctx context.Context, in *IotOtaWhiteGroupBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *IotOtaWhiteGroup, out *Response) error
		UpdateAll(ctx context.Context, in *IotOtaWhiteGroup, out *Response) error
		UpdateFields(ctx context.Context, in *IotOtaWhiteGroupUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *IotOtaWhiteGroupFilter, out *IotOtaWhiteGroupResponse) error
		Find(ctx context.Context, in *IotOtaWhiteGroupFilter, out *IotOtaWhiteGroupResponse) error
		Lists(ctx context.Context, in *IotOtaWhiteGroupListRequest, out *IotOtaWhiteGroupResponse) error
	}
	type IotOtaWhiteGroupService struct {
		iotOtaWhiteGroupService
	}
	h := &iotOtaWhiteGroupServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaWhiteGroupService.Create",
		Path:    []string{"/v1/iotOtaWhiteGroup/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaWhiteGroupService.Delete",
		Path:    []string{"/v1/iotOtaWhiteGroup/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaWhiteGroupService.DeleteById",
		Path:    []string{"/v1/iotOtaWhiteGroup/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaWhiteGroupService.DeleteByIds",
		Path:    []string{"/v1/iotOtaWhiteGroup/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaWhiteGroupService.Update",
		Path:    []string{"/v1/iotOtaWhiteGroup/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaWhiteGroupService.UpdateAll",
		Path:    []string{"/v1/iotOtaWhiteGroup/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaWhiteGroupService.UpdateFields",
		Path:    []string{"/v1/iotOtaWhiteGroup/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaWhiteGroupService.FindById",
		Path:    []string{"/v1/iotOtaWhiteGroup/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaWhiteGroupService.Find",
		Path:    []string{"/v1/iotOtaWhiteGroup/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaWhiteGroupService.Lists",
		Path:    []string{"/v1/iotOtaWhiteGroup/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&IotOtaWhiteGroupService{h}, opts...))
}

type iotOtaWhiteGroupServiceHandler struct {
	IotOtaWhiteGroupServiceHandler
}

func (h *iotOtaWhiteGroupServiceHandler) Create(ctx context.Context, in *IotOtaWhiteGroup, out *Response) error {
	return h.IotOtaWhiteGroupServiceHandler.Create(ctx, in, out)
}

func (h *iotOtaWhiteGroupServiceHandler) Delete(ctx context.Context, in *IotOtaWhiteGroup, out *Response) error {
	return h.IotOtaWhiteGroupServiceHandler.Delete(ctx, in, out)
}

func (h *iotOtaWhiteGroupServiceHandler) DeleteById(ctx context.Context, in *IotOtaWhiteGroup, out *Response) error {
	return h.IotOtaWhiteGroupServiceHandler.DeleteById(ctx, in, out)
}

func (h *iotOtaWhiteGroupServiceHandler) DeleteByIds(ctx context.Context, in *IotOtaWhiteGroupBatchDeleteRequest, out *Response) error {
	return h.IotOtaWhiteGroupServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *iotOtaWhiteGroupServiceHandler) Update(ctx context.Context, in *IotOtaWhiteGroup, out *Response) error {
	return h.IotOtaWhiteGroupServiceHandler.Update(ctx, in, out)
}

func (h *iotOtaWhiteGroupServiceHandler) UpdateAll(ctx context.Context, in *IotOtaWhiteGroup, out *Response) error {
	return h.IotOtaWhiteGroupServiceHandler.UpdateAll(ctx, in, out)
}

func (h *iotOtaWhiteGroupServiceHandler) UpdateFields(ctx context.Context, in *IotOtaWhiteGroupUpdateFieldsRequest, out *Response) error {
	return h.IotOtaWhiteGroupServiceHandler.UpdateFields(ctx, in, out)
}

func (h *iotOtaWhiteGroupServiceHandler) FindById(ctx context.Context, in *IotOtaWhiteGroupFilter, out *IotOtaWhiteGroupResponse) error {
	return h.IotOtaWhiteGroupServiceHandler.FindById(ctx, in, out)
}

func (h *iotOtaWhiteGroupServiceHandler) Find(ctx context.Context, in *IotOtaWhiteGroupFilter, out *IotOtaWhiteGroupResponse) error {
	return h.IotOtaWhiteGroupServiceHandler.Find(ctx, in, out)
}

func (h *iotOtaWhiteGroupServiceHandler) Lists(ctx context.Context, in *IotOtaWhiteGroupListRequest, out *IotOtaWhiteGroupResponse) error {
	return h.IotOtaWhiteGroupServiceHandler.Lists(ctx, in, out)
}
