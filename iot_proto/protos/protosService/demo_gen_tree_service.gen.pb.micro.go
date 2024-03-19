// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: demo_gen_tree_service.gen.proto

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

// Api Endpoints for DemoGenTreeService service

func NewDemoGenTreeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "DemoGenTreeService.Create",
			Path:    []string{"/v1/demoGenTree/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenTreeService.Delete",
			Path:    []string{"/v1/demoGenTree/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenTreeService.DeleteById",
			Path:    []string{"/v1/demoGenTree/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenTreeService.DeleteByIds",
			Path:    []string{"/v1/demoGenTree/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenTreeService.Update",
			Path:    []string{"/v1/demoGenTree/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenTreeService.UpdateAll",
			Path:    []string{"/v1/demoGenTree/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenTreeService.UpdateFields",
			Path:    []string{"/v1/demoGenTree/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenTreeService.FindById",
			Path:    []string{"/v1/demoGenTree/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenTreeService.Find",
			Path:    []string{"/v1/demoGenTree/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DemoGenTreeService.Lists",
			Path:    []string{"/v1/demoGenTree/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for DemoGenTreeService service

type DemoGenTreeService interface {
	//创建
	Create(ctx context.Context, in *DemoGenTree, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *DemoGenTree, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *DemoGenTree, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *DemoGenTreeBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *DemoGenTree, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *DemoGenTree, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *DemoGenTreeUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *DemoGenTreeFilter, opts ...client.CallOption) (*DemoGenTreeResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *DemoGenTreeFilter, opts ...client.CallOption) (*DemoGenTreeResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *DemoGenTreeListRequest, opts ...client.CallOption) (*DemoGenTreeResponse, error)
}

type demoGenTreeService struct {
	c    client.Client
	name string
}

func NewDemoGenTreeService(name string, c client.Client) DemoGenTreeService {
	return &demoGenTreeService{
		c:    c,
		name: name,
	}
}

func (c *demoGenTreeService) Create(ctx context.Context, in *DemoGenTree, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenTreeService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenTreeService) Delete(ctx context.Context, in *DemoGenTree, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenTreeService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenTreeService) DeleteById(ctx context.Context, in *DemoGenTree, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenTreeService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenTreeService) DeleteByIds(ctx context.Context, in *DemoGenTreeBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenTreeService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenTreeService) Update(ctx context.Context, in *DemoGenTree, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenTreeService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenTreeService) UpdateAll(ctx context.Context, in *DemoGenTree, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenTreeService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenTreeService) UpdateFields(ctx context.Context, in *DemoGenTreeUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DemoGenTreeService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenTreeService) FindById(ctx context.Context, in *DemoGenTreeFilter, opts ...client.CallOption) (*DemoGenTreeResponse, error) {
	req := c.c.NewRequest(c.name, "DemoGenTreeService.FindById", in)
	out := new(DemoGenTreeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenTreeService) Find(ctx context.Context, in *DemoGenTreeFilter, opts ...client.CallOption) (*DemoGenTreeResponse, error) {
	req := c.c.NewRequest(c.name, "DemoGenTreeService.Find", in)
	out := new(DemoGenTreeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGenTreeService) Lists(ctx context.Context, in *DemoGenTreeListRequest, opts ...client.CallOption) (*DemoGenTreeResponse, error) {
	req := c.c.NewRequest(c.name, "DemoGenTreeService.Lists", in)
	out := new(DemoGenTreeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DemoGenTreeService service

type DemoGenTreeServiceHandler interface {
	//创建
	Create(context.Context, *DemoGenTree, *Response) error
	//匹配多条件删除
	Delete(context.Context, *DemoGenTree, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *DemoGenTree, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *DemoGenTreeBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *DemoGenTree, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *DemoGenTree, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *DemoGenTreeUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *DemoGenTreeFilter, *DemoGenTreeResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *DemoGenTreeFilter, *DemoGenTreeResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *DemoGenTreeListRequest, *DemoGenTreeResponse) error
}

func RegisterDemoGenTreeServiceHandler(s server.Server, hdlr DemoGenTreeServiceHandler, opts ...server.HandlerOption) error {
	type demoGenTreeService interface {
		Create(ctx context.Context, in *DemoGenTree, out *Response) error
		Delete(ctx context.Context, in *DemoGenTree, out *Response) error
		DeleteById(ctx context.Context, in *DemoGenTree, out *Response) error
		DeleteByIds(ctx context.Context, in *DemoGenTreeBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *DemoGenTree, out *Response) error
		UpdateAll(ctx context.Context, in *DemoGenTree, out *Response) error
		UpdateFields(ctx context.Context, in *DemoGenTreeUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *DemoGenTreeFilter, out *DemoGenTreeResponse) error
		Find(ctx context.Context, in *DemoGenTreeFilter, out *DemoGenTreeResponse) error
		Lists(ctx context.Context, in *DemoGenTreeListRequest, out *DemoGenTreeResponse) error
	}
	type DemoGenTreeService struct {
		demoGenTreeService
	}
	h := &demoGenTreeServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenTreeService.Create",
		Path:    []string{"/v1/demoGenTree/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenTreeService.Delete",
		Path:    []string{"/v1/demoGenTree/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenTreeService.DeleteById",
		Path:    []string{"/v1/demoGenTree/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenTreeService.DeleteByIds",
		Path:    []string{"/v1/demoGenTree/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenTreeService.Update",
		Path:    []string{"/v1/demoGenTree/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenTreeService.UpdateAll",
		Path:    []string{"/v1/demoGenTree/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenTreeService.UpdateFields",
		Path:    []string{"/v1/demoGenTree/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenTreeService.FindById",
		Path:    []string{"/v1/demoGenTree/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenTreeService.Find",
		Path:    []string{"/v1/demoGenTree/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DemoGenTreeService.Lists",
		Path:    []string{"/v1/demoGenTree/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&DemoGenTreeService{h}, opts...))
}

type demoGenTreeServiceHandler struct {
	DemoGenTreeServiceHandler
}

func (h *demoGenTreeServiceHandler) Create(ctx context.Context, in *DemoGenTree, out *Response) error {
	return h.DemoGenTreeServiceHandler.Create(ctx, in, out)
}

func (h *demoGenTreeServiceHandler) Delete(ctx context.Context, in *DemoGenTree, out *Response) error {
	return h.DemoGenTreeServiceHandler.Delete(ctx, in, out)
}

func (h *demoGenTreeServiceHandler) DeleteById(ctx context.Context, in *DemoGenTree, out *Response) error {
	return h.DemoGenTreeServiceHandler.DeleteById(ctx, in, out)
}

func (h *demoGenTreeServiceHandler) DeleteByIds(ctx context.Context, in *DemoGenTreeBatchDeleteRequest, out *Response) error {
	return h.DemoGenTreeServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *demoGenTreeServiceHandler) Update(ctx context.Context, in *DemoGenTree, out *Response) error {
	return h.DemoGenTreeServiceHandler.Update(ctx, in, out)
}

func (h *demoGenTreeServiceHandler) UpdateAll(ctx context.Context, in *DemoGenTree, out *Response) error {
	return h.DemoGenTreeServiceHandler.UpdateAll(ctx, in, out)
}

func (h *demoGenTreeServiceHandler) UpdateFields(ctx context.Context, in *DemoGenTreeUpdateFieldsRequest, out *Response) error {
	return h.DemoGenTreeServiceHandler.UpdateFields(ctx, in, out)
}

func (h *demoGenTreeServiceHandler) FindById(ctx context.Context, in *DemoGenTreeFilter, out *DemoGenTreeResponse) error {
	return h.DemoGenTreeServiceHandler.FindById(ctx, in, out)
}

func (h *demoGenTreeServiceHandler) Find(ctx context.Context, in *DemoGenTreeFilter, out *DemoGenTreeResponse) error {
	return h.DemoGenTreeServiceHandler.Find(ctx, in, out)
}

func (h *demoGenTreeServiceHandler) Lists(ctx context.Context, in *DemoGenTreeListRequest, out *DemoGenTreeResponse) error {
	return h.DemoGenTreeServiceHandler.Lists(ctx, in, out)
}
