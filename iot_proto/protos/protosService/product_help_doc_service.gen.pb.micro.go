// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: product_help_doc_service.gen.proto

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

// Api Endpoints for ProductHelpDocService service

func NewProductHelpDocServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "ProductHelpDocService.Create",
			Path:    []string{"/v1/productHelpDoc/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductHelpDocService.Delete",
			Path:    []string{"/v1/productHelpDoc/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductHelpDocService.DeleteById",
			Path:    []string{"/v1/productHelpDoc/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductHelpDocService.DeleteByIds",
			Path:    []string{"/v1/productHelpDoc/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductHelpDocService.Update",
			Path:    []string{"/v1/productHelpDoc/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductHelpDocService.UpdateAll",
			Path:    []string{"/v1/productHelpDoc/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductHelpDocService.UpdateFields",
			Path:    []string{"/v1/productHelpDoc/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductHelpDocService.FindById",
			Path:    []string{"/v1/productHelpDoc/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductHelpDocService.Find",
			Path:    []string{"/v1/productHelpDoc/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductHelpDocService.Lists",
			Path:    []string{"/v1/productHelpDoc/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for ProductHelpDocService service

type ProductHelpDocService interface {
	//创建
	Create(ctx context.Context, in *ProductHelpDoc, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *ProductHelpDoc, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *ProductHelpDoc, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *ProductHelpDocBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *ProductHelpDoc, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *ProductHelpDoc, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *ProductHelpDocUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *ProductHelpDocFilter, opts ...client.CallOption) (*ProductHelpDocResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *ProductHelpDocFilter, opts ...client.CallOption) (*ProductHelpDocResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *ProductHelpDocListRequest, opts ...client.CallOption) (*ProductHelpDocResponse, error)
}

type productHelpDocService struct {
	c    client.Client
	name string
}

func NewProductHelpDocService(name string, c client.Client) ProductHelpDocService {
	return &productHelpDocService{
		c:    c,
		name: name,
	}
}

func (c *productHelpDocService) Create(ctx context.Context, in *ProductHelpDoc, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductHelpDocService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHelpDocService) Delete(ctx context.Context, in *ProductHelpDoc, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductHelpDocService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHelpDocService) DeleteById(ctx context.Context, in *ProductHelpDoc, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductHelpDocService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHelpDocService) DeleteByIds(ctx context.Context, in *ProductHelpDocBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductHelpDocService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHelpDocService) Update(ctx context.Context, in *ProductHelpDoc, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductHelpDocService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHelpDocService) UpdateAll(ctx context.Context, in *ProductHelpDoc, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductHelpDocService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHelpDocService) UpdateFields(ctx context.Context, in *ProductHelpDocUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductHelpDocService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHelpDocService) FindById(ctx context.Context, in *ProductHelpDocFilter, opts ...client.CallOption) (*ProductHelpDocResponse, error) {
	req := c.c.NewRequest(c.name, "ProductHelpDocService.FindById", in)
	out := new(ProductHelpDocResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHelpDocService) Find(ctx context.Context, in *ProductHelpDocFilter, opts ...client.CallOption) (*ProductHelpDocResponse, error) {
	req := c.c.NewRequest(c.name, "ProductHelpDocService.Find", in)
	out := new(ProductHelpDocResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHelpDocService) Lists(ctx context.Context, in *ProductHelpDocListRequest, opts ...client.CallOption) (*ProductHelpDocResponse, error) {
	req := c.c.NewRequest(c.name, "ProductHelpDocService.Lists", in)
	out := new(ProductHelpDocResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductHelpDocService service

type ProductHelpDocServiceHandler interface {
	//创建
	Create(context.Context, *ProductHelpDoc, *Response) error
	//匹配多条件删除
	Delete(context.Context, *ProductHelpDoc, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *ProductHelpDoc, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *ProductHelpDocBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *ProductHelpDoc, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *ProductHelpDoc, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *ProductHelpDocUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *ProductHelpDocFilter, *ProductHelpDocResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *ProductHelpDocFilter, *ProductHelpDocResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *ProductHelpDocListRequest, *ProductHelpDocResponse) error
}

func RegisterProductHelpDocServiceHandler(s server.Server, hdlr ProductHelpDocServiceHandler, opts ...server.HandlerOption) error {
	type productHelpDocService interface {
		Create(ctx context.Context, in *ProductHelpDoc, out *Response) error
		Delete(ctx context.Context, in *ProductHelpDoc, out *Response) error
		DeleteById(ctx context.Context, in *ProductHelpDoc, out *Response) error
		DeleteByIds(ctx context.Context, in *ProductHelpDocBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *ProductHelpDoc, out *Response) error
		UpdateAll(ctx context.Context, in *ProductHelpDoc, out *Response) error
		UpdateFields(ctx context.Context, in *ProductHelpDocUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *ProductHelpDocFilter, out *ProductHelpDocResponse) error
		Find(ctx context.Context, in *ProductHelpDocFilter, out *ProductHelpDocResponse) error
		Lists(ctx context.Context, in *ProductHelpDocListRequest, out *ProductHelpDocResponse) error
	}
	type ProductHelpDocService struct {
		productHelpDocService
	}
	h := &productHelpDocServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductHelpDocService.Create",
		Path:    []string{"/v1/productHelpDoc/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductHelpDocService.Delete",
		Path:    []string{"/v1/productHelpDoc/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductHelpDocService.DeleteById",
		Path:    []string{"/v1/productHelpDoc/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductHelpDocService.DeleteByIds",
		Path:    []string{"/v1/productHelpDoc/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductHelpDocService.Update",
		Path:    []string{"/v1/productHelpDoc/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductHelpDocService.UpdateAll",
		Path:    []string{"/v1/productHelpDoc/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductHelpDocService.UpdateFields",
		Path:    []string{"/v1/productHelpDoc/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductHelpDocService.FindById",
		Path:    []string{"/v1/productHelpDoc/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductHelpDocService.Find",
		Path:    []string{"/v1/productHelpDoc/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductHelpDocService.Lists",
		Path:    []string{"/v1/productHelpDoc/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&ProductHelpDocService{h}, opts...))
}

type productHelpDocServiceHandler struct {
	ProductHelpDocServiceHandler
}

func (h *productHelpDocServiceHandler) Create(ctx context.Context, in *ProductHelpDoc, out *Response) error {
	return h.ProductHelpDocServiceHandler.Create(ctx, in, out)
}

func (h *productHelpDocServiceHandler) Delete(ctx context.Context, in *ProductHelpDoc, out *Response) error {
	return h.ProductHelpDocServiceHandler.Delete(ctx, in, out)
}

func (h *productHelpDocServiceHandler) DeleteById(ctx context.Context, in *ProductHelpDoc, out *Response) error {
	return h.ProductHelpDocServiceHandler.DeleteById(ctx, in, out)
}

func (h *productHelpDocServiceHandler) DeleteByIds(ctx context.Context, in *ProductHelpDocBatchDeleteRequest, out *Response) error {
	return h.ProductHelpDocServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *productHelpDocServiceHandler) Update(ctx context.Context, in *ProductHelpDoc, out *Response) error {
	return h.ProductHelpDocServiceHandler.Update(ctx, in, out)
}

func (h *productHelpDocServiceHandler) UpdateAll(ctx context.Context, in *ProductHelpDoc, out *Response) error {
	return h.ProductHelpDocServiceHandler.UpdateAll(ctx, in, out)
}

func (h *productHelpDocServiceHandler) UpdateFields(ctx context.Context, in *ProductHelpDocUpdateFieldsRequest, out *Response) error {
	return h.ProductHelpDocServiceHandler.UpdateFields(ctx, in, out)
}

func (h *productHelpDocServiceHandler) FindById(ctx context.Context, in *ProductHelpDocFilter, out *ProductHelpDocResponse) error {
	return h.ProductHelpDocServiceHandler.FindById(ctx, in, out)
}

func (h *productHelpDocServiceHandler) Find(ctx context.Context, in *ProductHelpDocFilter, out *ProductHelpDocResponse) error {
	return h.ProductHelpDocServiceHandler.Find(ctx, in, out)
}

func (h *productHelpDocServiceHandler) Lists(ctx context.Context, in *ProductHelpDocListRequest, out *ProductHelpDocResponse) error {
	return h.ProductHelpDocServiceHandler.Lists(ctx, in, out)
}
