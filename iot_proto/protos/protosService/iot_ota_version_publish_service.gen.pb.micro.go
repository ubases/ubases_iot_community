// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: iot_ota_version_publish_service.gen.proto

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

// Api Endpoints for IotOtaVersionPublishService service

func NewIotOtaVersionPublishServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "IotOtaVersionPublishService.Create",
			Path:    []string{"/v1/iotOtaVersionPublish/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaVersionPublishService.Delete",
			Path:    []string{"/v1/iotOtaVersionPublish/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaVersionPublishService.DeleteById",
			Path:    []string{"/v1/iotOtaVersionPublish/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaVersionPublishService.DeleteByIds",
			Path:    []string{"/v1/iotOtaVersionPublish/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaVersionPublishService.Update",
			Path:    []string{"/v1/iotOtaVersionPublish/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaVersionPublishService.UpdateAll",
			Path:    []string{"/v1/iotOtaVersionPublish/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaVersionPublishService.UpdateFields",
			Path:    []string{"/v1/iotOtaVersionPublish/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaVersionPublishService.FindById",
			Path:    []string{"/v1/iotOtaVersionPublish/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaVersionPublishService.Find",
			Path:    []string{"/v1/iotOtaVersionPublish/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotOtaVersionPublishService.Lists",
			Path:    []string{"/v1/iotOtaVersionPublish/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for IotOtaVersionPublishService service

type IotOtaVersionPublishService interface {
	//创建
	Create(ctx context.Context, in *IotOtaVersionPublish, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *IotOtaVersionPublish, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *IotOtaVersionPublish, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *IotOtaVersionPublishBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *IotOtaVersionPublish, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *IotOtaVersionPublish, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *IotOtaVersionPublishUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *IotOtaVersionPublishFilter, opts ...client.CallOption) (*IotOtaVersionPublishResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *IotOtaVersionPublishFilter, opts ...client.CallOption) (*IotOtaVersionPublishResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *IotOtaVersionPublishListRequest, opts ...client.CallOption) (*IotOtaVersionPublishResponse, error)
}

type iotOtaVersionPublishService struct {
	c    client.Client
	name string
}

func NewIotOtaVersionPublishService(name string, c client.Client) IotOtaVersionPublishService {
	return &iotOtaVersionPublishService{
		c:    c,
		name: name,
	}
}

func (c *iotOtaVersionPublishService) Create(ctx context.Context, in *IotOtaVersionPublish, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaVersionPublishService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaVersionPublishService) Delete(ctx context.Context, in *IotOtaVersionPublish, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaVersionPublishService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaVersionPublishService) DeleteById(ctx context.Context, in *IotOtaVersionPublish, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaVersionPublishService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaVersionPublishService) DeleteByIds(ctx context.Context, in *IotOtaVersionPublishBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaVersionPublishService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaVersionPublishService) Update(ctx context.Context, in *IotOtaVersionPublish, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaVersionPublishService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaVersionPublishService) UpdateAll(ctx context.Context, in *IotOtaVersionPublish, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaVersionPublishService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaVersionPublishService) UpdateFields(ctx context.Context, in *IotOtaVersionPublishUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotOtaVersionPublishService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaVersionPublishService) FindById(ctx context.Context, in *IotOtaVersionPublishFilter, opts ...client.CallOption) (*IotOtaVersionPublishResponse, error) {
	req := c.c.NewRequest(c.name, "IotOtaVersionPublishService.FindById", in)
	out := new(IotOtaVersionPublishResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaVersionPublishService) Find(ctx context.Context, in *IotOtaVersionPublishFilter, opts ...client.CallOption) (*IotOtaVersionPublishResponse, error) {
	req := c.c.NewRequest(c.name, "IotOtaVersionPublishService.Find", in)
	out := new(IotOtaVersionPublishResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotOtaVersionPublishService) Lists(ctx context.Context, in *IotOtaVersionPublishListRequest, opts ...client.CallOption) (*IotOtaVersionPublishResponse, error) {
	req := c.c.NewRequest(c.name, "IotOtaVersionPublishService.Lists", in)
	out := new(IotOtaVersionPublishResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IotOtaVersionPublishService service

type IotOtaVersionPublishServiceHandler interface {
	//创建
	Create(context.Context, *IotOtaVersionPublish, *Response) error
	//匹配多条件删除
	Delete(context.Context, *IotOtaVersionPublish, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *IotOtaVersionPublish, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *IotOtaVersionPublishBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *IotOtaVersionPublish, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *IotOtaVersionPublish, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *IotOtaVersionPublishUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *IotOtaVersionPublishFilter, *IotOtaVersionPublishResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *IotOtaVersionPublishFilter, *IotOtaVersionPublishResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *IotOtaVersionPublishListRequest, *IotOtaVersionPublishResponse) error
}

func RegisterIotOtaVersionPublishServiceHandler(s server.Server, hdlr IotOtaVersionPublishServiceHandler, opts ...server.HandlerOption) error {
	type iotOtaVersionPublishService interface {
		Create(ctx context.Context, in *IotOtaVersionPublish, out *Response) error
		Delete(ctx context.Context, in *IotOtaVersionPublish, out *Response) error
		DeleteById(ctx context.Context, in *IotOtaVersionPublish, out *Response) error
		DeleteByIds(ctx context.Context, in *IotOtaVersionPublishBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *IotOtaVersionPublish, out *Response) error
		UpdateAll(ctx context.Context, in *IotOtaVersionPublish, out *Response) error
		UpdateFields(ctx context.Context, in *IotOtaVersionPublishUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *IotOtaVersionPublishFilter, out *IotOtaVersionPublishResponse) error
		Find(ctx context.Context, in *IotOtaVersionPublishFilter, out *IotOtaVersionPublishResponse) error
		Lists(ctx context.Context, in *IotOtaVersionPublishListRequest, out *IotOtaVersionPublishResponse) error
	}
	type IotOtaVersionPublishService struct {
		iotOtaVersionPublishService
	}
	h := &iotOtaVersionPublishServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaVersionPublishService.Create",
		Path:    []string{"/v1/iotOtaVersionPublish/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaVersionPublishService.Delete",
		Path:    []string{"/v1/iotOtaVersionPublish/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaVersionPublishService.DeleteById",
		Path:    []string{"/v1/iotOtaVersionPublish/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaVersionPublishService.DeleteByIds",
		Path:    []string{"/v1/iotOtaVersionPublish/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaVersionPublishService.Update",
		Path:    []string{"/v1/iotOtaVersionPublish/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaVersionPublishService.UpdateAll",
		Path:    []string{"/v1/iotOtaVersionPublish/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaVersionPublishService.UpdateFields",
		Path:    []string{"/v1/iotOtaVersionPublish/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaVersionPublishService.FindById",
		Path:    []string{"/v1/iotOtaVersionPublish/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaVersionPublishService.Find",
		Path:    []string{"/v1/iotOtaVersionPublish/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotOtaVersionPublishService.Lists",
		Path:    []string{"/v1/iotOtaVersionPublish/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&IotOtaVersionPublishService{h}, opts...))
}

type iotOtaVersionPublishServiceHandler struct {
	IotOtaVersionPublishServiceHandler
}

func (h *iotOtaVersionPublishServiceHandler) Create(ctx context.Context, in *IotOtaVersionPublish, out *Response) error {
	return h.IotOtaVersionPublishServiceHandler.Create(ctx, in, out)
}

func (h *iotOtaVersionPublishServiceHandler) Delete(ctx context.Context, in *IotOtaVersionPublish, out *Response) error {
	return h.IotOtaVersionPublishServiceHandler.Delete(ctx, in, out)
}

func (h *iotOtaVersionPublishServiceHandler) DeleteById(ctx context.Context, in *IotOtaVersionPublish, out *Response) error {
	return h.IotOtaVersionPublishServiceHandler.DeleteById(ctx, in, out)
}

func (h *iotOtaVersionPublishServiceHandler) DeleteByIds(ctx context.Context, in *IotOtaVersionPublishBatchDeleteRequest, out *Response) error {
	return h.IotOtaVersionPublishServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *iotOtaVersionPublishServiceHandler) Update(ctx context.Context, in *IotOtaVersionPublish, out *Response) error {
	return h.IotOtaVersionPublishServiceHandler.Update(ctx, in, out)
}

func (h *iotOtaVersionPublishServiceHandler) UpdateAll(ctx context.Context, in *IotOtaVersionPublish, out *Response) error {
	return h.IotOtaVersionPublishServiceHandler.UpdateAll(ctx, in, out)
}

func (h *iotOtaVersionPublishServiceHandler) UpdateFields(ctx context.Context, in *IotOtaVersionPublishUpdateFieldsRequest, out *Response) error {
	return h.IotOtaVersionPublishServiceHandler.UpdateFields(ctx, in, out)
}

func (h *iotOtaVersionPublishServiceHandler) FindById(ctx context.Context, in *IotOtaVersionPublishFilter, out *IotOtaVersionPublishResponse) error {
	return h.IotOtaVersionPublishServiceHandler.FindById(ctx, in, out)
}

func (h *iotOtaVersionPublishServiceHandler) Find(ctx context.Context, in *IotOtaVersionPublishFilter, out *IotOtaVersionPublishResponse) error {
	return h.IotOtaVersionPublishServiceHandler.Find(ctx, in, out)
}

func (h *iotOtaVersionPublishServiceHandler) Lists(ctx context.Context, in *IotOtaVersionPublishListRequest, out *IotOtaVersionPublishResponse) error {
	return h.IotOtaVersionPublishServiceHandler.Lists(ctx, in, out)
}
