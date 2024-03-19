// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pm_thing_model_events_service.gen.proto

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

// Api Endpoints for PmThingModelEventsService service

func NewPmThingModelEventsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "PmThingModelEventsService.Create",
			Path:    []string{"/v1/pmThingModelEvents/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmThingModelEventsService.Delete",
			Path:    []string{"/v1/pmThingModelEvents/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmThingModelEventsService.DeleteById",
			Path:    []string{"/v1/pmThingModelEvents/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmThingModelEventsService.DeleteByIds",
			Path:    []string{"/v1/pmThingModelEvents/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmThingModelEventsService.Update",
			Path:    []string{"/v1/pmThingModelEvents/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmThingModelEventsService.UpdateAll",
			Path:    []string{"/v1/pmThingModelEvents/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmThingModelEventsService.UpdateFields",
			Path:    []string{"/v1/pmThingModelEvents/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmThingModelEventsService.FindById",
			Path:    []string{"/v1/pmThingModelEvents/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmThingModelEventsService.Find",
			Path:    []string{"/v1/pmThingModelEvents/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmThingModelEventsService.Lists",
			Path:    []string{"/v1/pmThingModelEvents/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for PmThingModelEventsService service

type PmThingModelEventsService interface {
	//创建
	Create(ctx context.Context, in *PmThingModelEvents, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *PmThingModelEvents, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *PmThingModelEvents, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *PmThingModelEventsBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *PmThingModelEvents, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *PmThingModelEvents, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *PmThingModelEventsUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *PmThingModelEventsFilter, opts ...client.CallOption) (*PmThingModelEventsResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *PmThingModelEventsFilter, opts ...client.CallOption) (*PmThingModelEventsResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *PmThingModelEventsListRequest, opts ...client.CallOption) (*PmThingModelEventsResponse, error)
}

type pmThingModelEventsService struct {
	c    client.Client
	name string
}

func NewPmThingModelEventsService(name string, c client.Client) PmThingModelEventsService {
	return &pmThingModelEventsService{
		c:    c,
		name: name,
	}
}

func (c *pmThingModelEventsService) Create(ctx context.Context, in *PmThingModelEvents, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmThingModelEventsService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmThingModelEventsService) Delete(ctx context.Context, in *PmThingModelEvents, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmThingModelEventsService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmThingModelEventsService) DeleteById(ctx context.Context, in *PmThingModelEvents, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmThingModelEventsService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmThingModelEventsService) DeleteByIds(ctx context.Context, in *PmThingModelEventsBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmThingModelEventsService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmThingModelEventsService) Update(ctx context.Context, in *PmThingModelEvents, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmThingModelEventsService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmThingModelEventsService) UpdateAll(ctx context.Context, in *PmThingModelEvents, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmThingModelEventsService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmThingModelEventsService) UpdateFields(ctx context.Context, in *PmThingModelEventsUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmThingModelEventsService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmThingModelEventsService) FindById(ctx context.Context, in *PmThingModelEventsFilter, opts ...client.CallOption) (*PmThingModelEventsResponse, error) {
	req := c.c.NewRequest(c.name, "PmThingModelEventsService.FindById", in)
	out := new(PmThingModelEventsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmThingModelEventsService) Find(ctx context.Context, in *PmThingModelEventsFilter, opts ...client.CallOption) (*PmThingModelEventsResponse, error) {
	req := c.c.NewRequest(c.name, "PmThingModelEventsService.Find", in)
	out := new(PmThingModelEventsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmThingModelEventsService) Lists(ctx context.Context, in *PmThingModelEventsListRequest, opts ...client.CallOption) (*PmThingModelEventsResponse, error) {
	req := c.c.NewRequest(c.name, "PmThingModelEventsService.Lists", in)
	out := new(PmThingModelEventsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PmThingModelEventsService service

type PmThingModelEventsServiceHandler interface {
	//创建
	Create(context.Context, *PmThingModelEvents, *Response) error
	//匹配多条件删除
	Delete(context.Context, *PmThingModelEvents, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *PmThingModelEvents, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *PmThingModelEventsBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *PmThingModelEvents, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *PmThingModelEvents, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *PmThingModelEventsUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *PmThingModelEventsFilter, *PmThingModelEventsResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *PmThingModelEventsFilter, *PmThingModelEventsResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *PmThingModelEventsListRequest, *PmThingModelEventsResponse) error
}

func RegisterPmThingModelEventsServiceHandler(s server.Server, hdlr PmThingModelEventsServiceHandler, opts ...server.HandlerOption) error {
	type pmThingModelEventsService interface {
		Create(ctx context.Context, in *PmThingModelEvents, out *Response) error
		Delete(ctx context.Context, in *PmThingModelEvents, out *Response) error
		DeleteById(ctx context.Context, in *PmThingModelEvents, out *Response) error
		DeleteByIds(ctx context.Context, in *PmThingModelEventsBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *PmThingModelEvents, out *Response) error
		UpdateAll(ctx context.Context, in *PmThingModelEvents, out *Response) error
		UpdateFields(ctx context.Context, in *PmThingModelEventsUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *PmThingModelEventsFilter, out *PmThingModelEventsResponse) error
		Find(ctx context.Context, in *PmThingModelEventsFilter, out *PmThingModelEventsResponse) error
		Lists(ctx context.Context, in *PmThingModelEventsListRequest, out *PmThingModelEventsResponse) error
	}
	type PmThingModelEventsService struct {
		pmThingModelEventsService
	}
	h := &pmThingModelEventsServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmThingModelEventsService.Create",
		Path:    []string{"/v1/pmThingModelEvents/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmThingModelEventsService.Delete",
		Path:    []string{"/v1/pmThingModelEvents/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmThingModelEventsService.DeleteById",
		Path:    []string{"/v1/pmThingModelEvents/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmThingModelEventsService.DeleteByIds",
		Path:    []string{"/v1/pmThingModelEvents/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmThingModelEventsService.Update",
		Path:    []string{"/v1/pmThingModelEvents/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmThingModelEventsService.UpdateAll",
		Path:    []string{"/v1/pmThingModelEvents/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmThingModelEventsService.UpdateFields",
		Path:    []string{"/v1/pmThingModelEvents/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmThingModelEventsService.FindById",
		Path:    []string{"/v1/pmThingModelEvents/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmThingModelEventsService.Find",
		Path:    []string{"/v1/pmThingModelEvents/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmThingModelEventsService.Lists",
		Path:    []string{"/v1/pmThingModelEvents/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&PmThingModelEventsService{h}, opts...))
}

type pmThingModelEventsServiceHandler struct {
	PmThingModelEventsServiceHandler
}

func (h *pmThingModelEventsServiceHandler) Create(ctx context.Context, in *PmThingModelEvents, out *Response) error {
	return h.PmThingModelEventsServiceHandler.Create(ctx, in, out)
}

func (h *pmThingModelEventsServiceHandler) Delete(ctx context.Context, in *PmThingModelEvents, out *Response) error {
	return h.PmThingModelEventsServiceHandler.Delete(ctx, in, out)
}

func (h *pmThingModelEventsServiceHandler) DeleteById(ctx context.Context, in *PmThingModelEvents, out *Response) error {
	return h.PmThingModelEventsServiceHandler.DeleteById(ctx, in, out)
}

func (h *pmThingModelEventsServiceHandler) DeleteByIds(ctx context.Context, in *PmThingModelEventsBatchDeleteRequest, out *Response) error {
	return h.PmThingModelEventsServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *pmThingModelEventsServiceHandler) Update(ctx context.Context, in *PmThingModelEvents, out *Response) error {
	return h.PmThingModelEventsServiceHandler.Update(ctx, in, out)
}

func (h *pmThingModelEventsServiceHandler) UpdateAll(ctx context.Context, in *PmThingModelEvents, out *Response) error {
	return h.PmThingModelEventsServiceHandler.UpdateAll(ctx, in, out)
}

func (h *pmThingModelEventsServiceHandler) UpdateFields(ctx context.Context, in *PmThingModelEventsUpdateFieldsRequest, out *Response) error {
	return h.PmThingModelEventsServiceHandler.UpdateFields(ctx, in, out)
}

func (h *pmThingModelEventsServiceHandler) FindById(ctx context.Context, in *PmThingModelEventsFilter, out *PmThingModelEventsResponse) error {
	return h.PmThingModelEventsServiceHandler.FindById(ctx, in, out)
}

func (h *pmThingModelEventsServiceHandler) Find(ctx context.Context, in *PmThingModelEventsFilter, out *PmThingModelEventsResponse) error {
	return h.PmThingModelEventsServiceHandler.Find(ctx, in, out)
}

func (h *pmThingModelEventsServiceHandler) Lists(ctx context.Context, in *PmThingModelEventsListRequest, out *PmThingModelEventsResponse) error {
	return h.PmThingModelEventsServiceHandler.Lists(ctx, in, out)
}
