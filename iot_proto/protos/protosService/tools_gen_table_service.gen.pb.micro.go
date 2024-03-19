// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: tools_gen_table_service.gen.proto

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

// Api Endpoints for ToolsGenTableService service

func NewToolsGenTableServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "ToolsGenTableService.Create",
			Path:    []string{"/v1/toolsGenTable/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ToolsGenTableService.Delete",
			Path:    []string{"/v1/toolsGenTable/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ToolsGenTableService.DeleteById",
			Path:    []string{"/v1/toolsGenTable/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ToolsGenTableService.DeleteByIds",
			Path:    []string{"/v1/toolsGenTable/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ToolsGenTableService.Update",
			Path:    []string{"/v1/toolsGenTable/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ToolsGenTableService.UpdateAll",
			Path:    []string{"/v1/toolsGenTable/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ToolsGenTableService.UpdateFields",
			Path:    []string{"/v1/toolsGenTable/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ToolsGenTableService.FindById",
			Path:    []string{"/v1/toolsGenTable/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ToolsGenTableService.Find",
			Path:    []string{"/v1/toolsGenTable/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ToolsGenTableService.Lists",
			Path:    []string{"/v1/toolsGenTable/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for ToolsGenTableService service

type ToolsGenTableService interface {
	//创建
	Create(ctx context.Context, in *ToolsGenTable, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *ToolsGenTable, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *ToolsGenTable, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *ToolsGenTableBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *ToolsGenTable, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *ToolsGenTable, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *ToolsGenTableUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *ToolsGenTableFilter, opts ...client.CallOption) (*ToolsGenTableResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *ToolsGenTableFilter, opts ...client.CallOption) (*ToolsGenTableResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *ToolsGenTableListRequest, opts ...client.CallOption) (*ToolsGenTableResponse, error)
}

type toolsGenTableService struct {
	c    client.Client
	name string
}

func NewToolsGenTableService(name string, c client.Client) ToolsGenTableService {
	return &toolsGenTableService{
		c:    c,
		name: name,
	}
}

func (c *toolsGenTableService) Create(ctx context.Context, in *ToolsGenTable, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ToolsGenTableService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsGenTableService) Delete(ctx context.Context, in *ToolsGenTable, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ToolsGenTableService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsGenTableService) DeleteById(ctx context.Context, in *ToolsGenTable, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ToolsGenTableService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsGenTableService) DeleteByIds(ctx context.Context, in *ToolsGenTableBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ToolsGenTableService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsGenTableService) Update(ctx context.Context, in *ToolsGenTable, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ToolsGenTableService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsGenTableService) UpdateAll(ctx context.Context, in *ToolsGenTable, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ToolsGenTableService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsGenTableService) UpdateFields(ctx context.Context, in *ToolsGenTableUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ToolsGenTableService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsGenTableService) FindById(ctx context.Context, in *ToolsGenTableFilter, opts ...client.CallOption) (*ToolsGenTableResponse, error) {
	req := c.c.NewRequest(c.name, "ToolsGenTableService.FindById", in)
	out := new(ToolsGenTableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsGenTableService) Find(ctx context.Context, in *ToolsGenTableFilter, opts ...client.CallOption) (*ToolsGenTableResponse, error) {
	req := c.c.NewRequest(c.name, "ToolsGenTableService.Find", in)
	out := new(ToolsGenTableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsGenTableService) Lists(ctx context.Context, in *ToolsGenTableListRequest, opts ...client.CallOption) (*ToolsGenTableResponse, error) {
	req := c.c.NewRequest(c.name, "ToolsGenTableService.Lists", in)
	out := new(ToolsGenTableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ToolsGenTableService service

type ToolsGenTableServiceHandler interface {
	//创建
	Create(context.Context, *ToolsGenTable, *Response) error
	//匹配多条件删除
	Delete(context.Context, *ToolsGenTable, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *ToolsGenTable, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *ToolsGenTableBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *ToolsGenTable, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *ToolsGenTable, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *ToolsGenTableUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *ToolsGenTableFilter, *ToolsGenTableResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *ToolsGenTableFilter, *ToolsGenTableResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *ToolsGenTableListRequest, *ToolsGenTableResponse) error
}

func RegisterToolsGenTableServiceHandler(s server.Server, hdlr ToolsGenTableServiceHandler, opts ...server.HandlerOption) error {
	type toolsGenTableService interface {
		Create(ctx context.Context, in *ToolsGenTable, out *Response) error
		Delete(ctx context.Context, in *ToolsGenTable, out *Response) error
		DeleteById(ctx context.Context, in *ToolsGenTable, out *Response) error
		DeleteByIds(ctx context.Context, in *ToolsGenTableBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *ToolsGenTable, out *Response) error
		UpdateAll(ctx context.Context, in *ToolsGenTable, out *Response) error
		UpdateFields(ctx context.Context, in *ToolsGenTableUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *ToolsGenTableFilter, out *ToolsGenTableResponse) error
		Find(ctx context.Context, in *ToolsGenTableFilter, out *ToolsGenTableResponse) error
		Lists(ctx context.Context, in *ToolsGenTableListRequest, out *ToolsGenTableResponse) error
	}
	type ToolsGenTableService struct {
		toolsGenTableService
	}
	h := &toolsGenTableServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ToolsGenTableService.Create",
		Path:    []string{"/v1/toolsGenTable/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ToolsGenTableService.Delete",
		Path:    []string{"/v1/toolsGenTable/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ToolsGenTableService.DeleteById",
		Path:    []string{"/v1/toolsGenTable/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ToolsGenTableService.DeleteByIds",
		Path:    []string{"/v1/toolsGenTable/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ToolsGenTableService.Update",
		Path:    []string{"/v1/toolsGenTable/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ToolsGenTableService.UpdateAll",
		Path:    []string{"/v1/toolsGenTable/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ToolsGenTableService.UpdateFields",
		Path:    []string{"/v1/toolsGenTable/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ToolsGenTableService.FindById",
		Path:    []string{"/v1/toolsGenTable/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ToolsGenTableService.Find",
		Path:    []string{"/v1/toolsGenTable/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ToolsGenTableService.Lists",
		Path:    []string{"/v1/toolsGenTable/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&ToolsGenTableService{h}, opts...))
}

type toolsGenTableServiceHandler struct {
	ToolsGenTableServiceHandler
}

func (h *toolsGenTableServiceHandler) Create(ctx context.Context, in *ToolsGenTable, out *Response) error {
	return h.ToolsGenTableServiceHandler.Create(ctx, in, out)
}

func (h *toolsGenTableServiceHandler) Delete(ctx context.Context, in *ToolsGenTable, out *Response) error {
	return h.ToolsGenTableServiceHandler.Delete(ctx, in, out)
}

func (h *toolsGenTableServiceHandler) DeleteById(ctx context.Context, in *ToolsGenTable, out *Response) error {
	return h.ToolsGenTableServiceHandler.DeleteById(ctx, in, out)
}

func (h *toolsGenTableServiceHandler) DeleteByIds(ctx context.Context, in *ToolsGenTableBatchDeleteRequest, out *Response) error {
	return h.ToolsGenTableServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *toolsGenTableServiceHandler) Update(ctx context.Context, in *ToolsGenTable, out *Response) error {
	return h.ToolsGenTableServiceHandler.Update(ctx, in, out)
}

func (h *toolsGenTableServiceHandler) UpdateAll(ctx context.Context, in *ToolsGenTable, out *Response) error {
	return h.ToolsGenTableServiceHandler.UpdateAll(ctx, in, out)
}

func (h *toolsGenTableServiceHandler) UpdateFields(ctx context.Context, in *ToolsGenTableUpdateFieldsRequest, out *Response) error {
	return h.ToolsGenTableServiceHandler.UpdateFields(ctx, in, out)
}

func (h *toolsGenTableServiceHandler) FindById(ctx context.Context, in *ToolsGenTableFilter, out *ToolsGenTableResponse) error {
	return h.ToolsGenTableServiceHandler.FindById(ctx, in, out)
}

func (h *toolsGenTableServiceHandler) Find(ctx context.Context, in *ToolsGenTableFilter, out *ToolsGenTableResponse) error {
	return h.ToolsGenTableServiceHandler.Find(ctx, in, out)
}

func (h *toolsGenTableServiceHandler) Lists(ctx context.Context, in *ToolsGenTableListRequest, out *ToolsGenTableResponse) error {
	return h.ToolsGenTableServiceHandler.Lists(ctx, in, out)
}
