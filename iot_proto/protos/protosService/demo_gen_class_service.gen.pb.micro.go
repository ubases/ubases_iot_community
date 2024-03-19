// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: demo_gen_class_service.gen.proto

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

// Api Endpoints for DemoGenClassService service

func NewDemoGenClassServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "DemoGenClassService.Create",
			Path:    []string{"/v1/demoGenClass/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenClassService.Delete",
			Path:    []string{"/v1/demoGenClass/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenClassService.DeleteById",
			Path:    []string{"/v1/demoGenClass/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenClassService.DeleteByIds",
			Path:    []string{"/v1/demoGenClass/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenClassService.Update",
			Path:    []string{"/v1/demoGenClass/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenClassService.UpdateAll",
			Path:    []string{"/v1/demoGenClass/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenClassService.UpdateFields",
			Path:    []string{"/v1/demoGenClass/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenClassService.FindById",
			Path:    []string{"/v1/demoGenClass/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenClassService.Find",
			Path:    []string{"/v1/demoGenClass/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenClassService.Lists",
			Path:    []string{"/v1/demoGenClass/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for DemoGenClassService service

type DemoGenClassService interface {
	//创建
	Create(ctx context.Context, in *DemoGenClass, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *DemoGenClass, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *DemoGenClass, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *DemoGenClassBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *DemoGenClass, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *DemoGenClass, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *DemoGenClassUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *DemoGenClassFilter, opts ...client.CallOption) (*DemoGenClassResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *DemoGenClassFilter, opts ...client.CallOption) (*DemoGenClassResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *DemoGenClassListRequest, opts ...client.CallOption) (*DemoGenClassResponse, error)
}

type demoGenClassService struct {
	c    client.Client
	name string
}

func NewDemoGenClassService(name string, c client.Client) DemoGenClassService {
	return &demoGenClassService{
		c:    c,
		name: name,
	}
}

func (c *demoGenClassService) Create(ctx context.Context, in *DemoGenClass, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenClassService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenClassService) Delete(ctx context.Context, in *DemoGenClass, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenClassService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenClassService) DeleteById(ctx context.Context, in *DemoGenClass, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenClassService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenClassService) DeleteByIds(ctx context.Context, in *DemoGenClassBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenClassService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenClassService) Update(ctx context.Context, in *DemoGenClass, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenClassService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenClassService) UpdateAll(ctx context.Context, in *DemoGenClass, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenClassService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenClassService) UpdateFields(ctx context.Context, in *DemoGenClassUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenClassService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenClassService) FindById(ctx context.Context, in *DemoGenClassFilter, opts ...client.CallOption) (*DemoGenClassResponse, error) {
	req := c.c.NewRequest(c.name, "DemoGenClassService.FindById", in)
	out := new(DemoGenClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenClassService) Find(ctx context.Context, in *DemoGenClassFilter, opts ...client.CallOption) (*DemoGenClassResponse, error) {
	req := c.c.NewRequest(c.name, "DemoGenClassService.Find", in)
	out := new(DemoGenClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenClassService) Lists(ctx context.Context, in *DemoGenClassListRequest, opts ...client.CallOption) (*DemoGenClassResponse, error) {
	req := c.c.NewRequest(c.name, "DemoGenClassService.Lists", in)
	out := new(DemoGenClassResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DemoGenClassService service

type DemoGenClassServiceHandler interface {
	//创建
	Create(context.Context, *DemoGenClass, *Response) error
	//匹配多条件删除
	Delete(context.Context, *DemoGenClass, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *DemoGenClass, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *DemoGenClassBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *DemoGenClass, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *DemoGenClass, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *DemoGenClassUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *DemoGenClassFilter, *DemoGenClassResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *DemoGenClassFilter, *DemoGenClassResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *DemoGenClassListRequest, *DemoGenClassResponse) error
}

func RegisterDemoGenClassServiceHandler(s server.Server, hdlr DemoGenClassServiceHandler, opts ...server.HandlerOption) error {
	type demoGenClassService interface {
		Create(ctx context.Context, in *DemoGenClass, out *Response) error
		Delete(ctx context.Context, in *DemoGenClass, out *Response) error
		DeleteById(ctx context.Context, in *DemoGenClass, out *Response) error
		DeleteByIds(ctx context.Context, in *DemoGenClassBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *DemoGenClass, out *Response) error
		UpdateAll(ctx context.Context, in *DemoGenClass, out *Response) error
		UpdateFields(ctx context.Context, in *DemoGenClassUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *DemoGenClassFilter, out *DemoGenClassResponse) error
		Find(ctx context.Context, in *DemoGenClassFilter, out *DemoGenClassResponse) error
		Lists(ctx context.Context, in *DemoGenClassListRequest, out *DemoGenClassResponse) error
	}
	type DemoGenClassService struct {
		demoGenClassService
	}
	h := &demoGenClassServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenClassService.Create",
		Path:    []string{"/v1/demoGenClass/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenClassService.Delete",
		Path:    []string{"/v1/demoGenClass/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenClassService.DeleteById",
		Path:    []string{"/v1/demoGenClass/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenClassService.DeleteByIds",
		Path:    []string{"/v1/demoGenClass/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenClassService.Update",
		Path:    []string{"/v1/demoGenClass/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenClassService.UpdateAll",
		Path:    []string{"/v1/demoGenClass/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenClassService.UpdateFields",
		Path:    []string{"/v1/demoGenClass/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenClassService.FindById",
		Path:    []string{"/v1/demoGenClass/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenClassService.Find",
		Path:    []string{"/v1/demoGenClass/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenClassService.Lists",
		Path:    []string{"/v1/demoGenClass/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&DemoGenClassService{h}, opts...))
}

type demoGenClassServiceHandler struct {
	DemoGenClassServiceHandler
}

func (h *demoGenClassServiceHandler) Create(ctx context.Context, in *DemoGenClass, out *Response) error {
	return h.DemoGenClassServiceHandler.Create(ctx, in, out)
}

func (h *demoGenClassServiceHandler) Delete(ctx context.Context, in *DemoGenClass, out *Response) error {
	return h.DemoGenClassServiceHandler.Delete(ctx, in, out)
}

func (h *demoGenClassServiceHandler) DeleteById(ctx context.Context, in *DemoGenClass, out *Response) error {
	return h.DemoGenClassServiceHandler.DeleteById(ctx, in, out)
}

func (h *demoGenClassServiceHandler) DeleteByIds(ctx context.Context, in *DemoGenClassBatchDeleteRequest, out *Response) error {
	return h.DemoGenClassServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *demoGenClassServiceHandler) Update(ctx context.Context, in *DemoGenClass, out *Response) error {
	return h.DemoGenClassServiceHandler.Update(ctx, in, out)
}

func (h *demoGenClassServiceHandler) UpdateAll(ctx context.Context, in *DemoGenClass, out *Response) error {
	return h.DemoGenClassServiceHandler.UpdateAll(ctx, in, out)
}

func (h *demoGenClassServiceHandler) UpdateFields(ctx context.Context, in *DemoGenClassUpdateFieldsRequest, out *Response) error {
	return h.DemoGenClassServiceHandler.UpdateFields(ctx, in, out)
}

func (h *demoGenClassServiceHandler) FindById(ctx context.Context, in *DemoGenClassFilter, out *DemoGenClassResponse) error {
	return h.DemoGenClassServiceHandler.FindById(ctx, in, out)
}

func (h *demoGenClassServiceHandler) Find(ctx context.Context, in *DemoGenClassFilter, out *DemoGenClassResponse) error {
	return h.DemoGenClassServiceHandler.Find(ctx, in, out)
}

func (h *demoGenClassServiceHandler) Lists(ctx context.Context, in *DemoGenClassListRequest, out *DemoGenClassResponse) error {
	return h.DemoGenClassServiceHandler.Lists(ctx, in, out)
}
