// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: product_fault_type_service.gen.proto

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

// Api Endpoints for ProductFaultTypeService service

func NewProductFaultTypeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "ProductFaultTypeService.Create",
			Path:    []string{"/v1/productFaultType/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductFaultTypeService.Delete",
			Path:    []string{"/v1/productFaultType/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductFaultTypeService.DeleteById",
			Path:    []string{"/v1/productFaultType/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductFaultTypeService.DeleteByIds",
			Path:    []string{"/v1/productFaultType/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductFaultTypeService.Update",
			Path:    []string{"/v1/productFaultType/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductFaultTypeService.UpdateAll",
			Path:    []string{"/v1/productFaultType/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductFaultTypeService.UpdateFields",
			Path:    []string{"/v1/productFaultType/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductFaultTypeService.FindById",
			Path:    []string{"/v1/productFaultType/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductFaultTypeService.Find",
			Path:    []string{"/v1/productFaultType/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "ProductFaultTypeService.Lists",
			Path:    []string{"/v1/productFaultType/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for ProductFaultTypeService service

type ProductFaultTypeService interface {
	//创建
	Create(ctx context.Context, in *ProductFaultType, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *ProductFaultType, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *ProductFaultType, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *ProductFaultTypeBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *ProductFaultType, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *ProductFaultType, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *ProductFaultTypeUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *ProductFaultTypeFilter, opts ...client.CallOption) (*ProductFaultTypeResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *ProductFaultTypeFilter, opts ...client.CallOption) (*ProductFaultTypeResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *ProductFaultTypeListRequest, opts ...client.CallOption) (*ProductFaultTypeResponse, error)
}

type productFaultTypeService struct {
	c    client.Client
	name string
}

func NewProductFaultTypeService(name string, c client.Client) ProductFaultTypeService {
	return &productFaultTypeService{
		c:    c,
		name: name,
	}
}

func (c *productFaultTypeService) Create(ctx context.Context, in *ProductFaultType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductFaultTypeService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFaultTypeService) Delete(ctx context.Context, in *ProductFaultType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductFaultTypeService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFaultTypeService) DeleteById(ctx context.Context, in *ProductFaultType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductFaultTypeService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFaultTypeService) DeleteByIds(ctx context.Context, in *ProductFaultTypeBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductFaultTypeService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFaultTypeService) Update(ctx context.Context, in *ProductFaultType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductFaultTypeService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFaultTypeService) UpdateAll(ctx context.Context, in *ProductFaultType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductFaultTypeService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFaultTypeService) UpdateFields(ctx context.Context, in *ProductFaultTypeUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductFaultTypeService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFaultTypeService) FindById(ctx context.Context, in *ProductFaultTypeFilter, opts ...client.CallOption) (*ProductFaultTypeResponse, error) {
	req := c.c.NewRequest(c.name, "ProductFaultTypeService.FindById", in)
	out := new(ProductFaultTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFaultTypeService) Find(ctx context.Context, in *ProductFaultTypeFilter, opts ...client.CallOption) (*ProductFaultTypeResponse, error) {
	req := c.c.NewRequest(c.name, "ProductFaultTypeService.Find", in)
	out := new(ProductFaultTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFaultTypeService) Lists(ctx context.Context, in *ProductFaultTypeListRequest, opts ...client.CallOption) (*ProductFaultTypeResponse, error) {
	req := c.c.NewRequest(c.name, "ProductFaultTypeService.Lists", in)
	out := new(ProductFaultTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductFaultTypeService service

type ProductFaultTypeServiceHandler interface {
	//创建
	Create(context.Context, *ProductFaultType, *Response) error
	//匹配多条件删除
	Delete(context.Context, *ProductFaultType, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *ProductFaultType, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *ProductFaultTypeBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *ProductFaultType, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *ProductFaultType, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *ProductFaultTypeUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *ProductFaultTypeFilter, *ProductFaultTypeResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *ProductFaultTypeFilter, *ProductFaultTypeResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *ProductFaultTypeListRequest, *ProductFaultTypeResponse) error
}

func RegisterProductFaultTypeServiceHandler(s server.Server, hdlr ProductFaultTypeServiceHandler, opts ...server.HandlerOption) error {
	type productFaultTypeService interface {
		Create(ctx context.Context, in *ProductFaultType, out *Response) error
		Delete(ctx context.Context, in *ProductFaultType, out *Response) error
		DeleteById(ctx context.Context, in *ProductFaultType, out *Response) error
		DeleteByIds(ctx context.Context, in *ProductFaultTypeBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *ProductFaultType, out *Response) error
		UpdateAll(ctx context.Context, in *ProductFaultType, out *Response) error
		UpdateFields(ctx context.Context, in *ProductFaultTypeUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *ProductFaultTypeFilter, out *ProductFaultTypeResponse) error
		Find(ctx context.Context, in *ProductFaultTypeFilter, out *ProductFaultTypeResponse) error
		Lists(ctx context.Context, in *ProductFaultTypeListRequest, out *ProductFaultTypeResponse) error
	}
	type ProductFaultTypeService struct {
		productFaultTypeService
	}
	h := &productFaultTypeServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductFaultTypeService.Create",
		Path:    []string{"/v1/productFaultType/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductFaultTypeService.Delete",
		Path:    []string{"/v1/productFaultType/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductFaultTypeService.DeleteById",
		Path:    []string{"/v1/productFaultType/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductFaultTypeService.DeleteByIds",
		Path:    []string{"/v1/productFaultType/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductFaultTypeService.Update",
		Path:    []string{"/v1/productFaultType/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductFaultTypeService.UpdateAll",
		Path:    []string{"/v1/productFaultType/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductFaultTypeService.UpdateFields",
		Path:    []string{"/v1/productFaultType/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductFaultTypeService.FindById",
		Path:    []string{"/v1/productFaultType/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductFaultTypeService.Find",
		Path:    []string{"/v1/productFaultType/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ProductFaultTypeService.Lists",
		Path:    []string{"/v1/productFaultType/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&ProductFaultTypeService{h}, opts...))
}

type productFaultTypeServiceHandler struct {
	ProductFaultTypeServiceHandler
}

func (h *productFaultTypeServiceHandler) Create(ctx context.Context, in *ProductFaultType, out *Response) error {
	return h.ProductFaultTypeServiceHandler.Create(ctx, in, out)
}

func (h *productFaultTypeServiceHandler) Delete(ctx context.Context, in *ProductFaultType, out *Response) error {
	return h.ProductFaultTypeServiceHandler.Delete(ctx, in, out)
}

func (h *productFaultTypeServiceHandler) DeleteById(ctx context.Context, in *ProductFaultType, out *Response) error {
	return h.ProductFaultTypeServiceHandler.DeleteById(ctx, in, out)
}

func (h *productFaultTypeServiceHandler) DeleteByIds(ctx context.Context, in *ProductFaultTypeBatchDeleteRequest, out *Response) error {
	return h.ProductFaultTypeServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *productFaultTypeServiceHandler) Update(ctx context.Context, in *ProductFaultType, out *Response) error {
	return h.ProductFaultTypeServiceHandler.Update(ctx, in, out)
}

func (h *productFaultTypeServiceHandler) UpdateAll(ctx context.Context, in *ProductFaultType, out *Response) error {
	return h.ProductFaultTypeServiceHandler.UpdateAll(ctx, in, out)
}

func (h *productFaultTypeServiceHandler) UpdateFields(ctx context.Context, in *ProductFaultTypeUpdateFieldsRequest, out *Response) error {
	return h.ProductFaultTypeServiceHandler.UpdateFields(ctx, in, out)
}

func (h *productFaultTypeServiceHandler) FindById(ctx context.Context, in *ProductFaultTypeFilter, out *ProductFaultTypeResponse) error {
	return h.ProductFaultTypeServiceHandler.FindById(ctx, in, out)
}

func (h *productFaultTypeServiceHandler) Find(ctx context.Context, in *ProductFaultTypeFilter, out *ProductFaultTypeResponse) error {
	return h.ProductFaultTypeServiceHandler.Find(ctx, in, out)
}

func (h *productFaultTypeServiceHandler) Lists(ctx context.Context, in *ProductFaultTypeListRequest, out *ProductFaultTypeResponse) error {
	return h.ProductFaultTypeServiceHandler.Lists(ctx, in, out)
}
