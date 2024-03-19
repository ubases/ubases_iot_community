// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cms_document_service.gen.proto

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

// Api Endpoints for CmsDocumentService service

func NewCmsDocumentServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "CmsDocumentService.Create",
			Path:    []string{"/v1/cmsDocument/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "CmsDocumentService.Delete",
			Path:    []string{"/v1/cmsDocument/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "CmsDocumentService.DeleteById",
			Path:    []string{"/v1/cmsDocument/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "CmsDocumentService.DeleteByIds",
			Path:    []string{"/v1/cmsDocument/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "CmsDocumentService.Update",
			Path:    []string{"/v1/cmsDocument/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "CmsDocumentService.UpdateAll",
			Path:    []string{"/v1/cmsDocument/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "CmsDocumentService.UpdateFields",
			Path:    []string{"/v1/cmsDocument/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "CmsDocumentService.FindById",
			Path:    []string{"/v1/cmsDocument/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "CmsDocumentService.Find",
			Path:    []string{"/v1/cmsDocument/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "CmsDocumentService.Lists",
			Path:    []string{"/v1/cmsDocument/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for CmsDocumentService service

type CmsDocumentService interface {
	//创建
	Create(ctx context.Context, in *CmsDocument, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *CmsDocument, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *CmsDocument, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *CmsDocumentBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *CmsDocument, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *CmsDocument, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *CmsDocumentUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *CmsDocumentFilter, opts ...client.CallOption) (*CmsDocumentResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *CmsDocumentFilter, opts ...client.CallOption) (*CmsDocumentResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *CmsDocumentListRequest, opts ...client.CallOption) (*CmsDocumentResponse, error)
}

type cmsDocumentService struct {
	c    client.Client
	name string
}

func NewCmsDocumentService(name string, c client.Client) CmsDocumentService {
	return &cmsDocumentService{
		c:    c,
		name: name,
	}
}

func (c *cmsDocumentService) Create(ctx context.Context, in *CmsDocument, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CmsDocumentService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsDocumentService) Delete(ctx context.Context, in *CmsDocument, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CmsDocumentService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsDocumentService) DeleteById(ctx context.Context, in *CmsDocument, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CmsDocumentService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsDocumentService) DeleteByIds(ctx context.Context, in *CmsDocumentBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CmsDocumentService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsDocumentService) Update(ctx context.Context, in *CmsDocument, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CmsDocumentService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsDocumentService) UpdateAll(ctx context.Context, in *CmsDocument, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CmsDocumentService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsDocumentService) UpdateFields(ctx context.Context, in *CmsDocumentUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CmsDocumentService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsDocumentService) FindById(ctx context.Context, in *CmsDocumentFilter, opts ...client.CallOption) (*CmsDocumentResponse, error) {
	req := c.c.NewRequest(c.name, "CmsDocumentService.FindById", in)
	out := new(CmsDocumentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsDocumentService) Find(ctx context.Context, in *CmsDocumentFilter, opts ...client.CallOption) (*CmsDocumentResponse, error) {
	req := c.c.NewRequest(c.name, "CmsDocumentService.Find", in)
	out := new(CmsDocumentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsDocumentService) Lists(ctx context.Context, in *CmsDocumentListRequest, opts ...client.CallOption) (*CmsDocumentResponse, error) {
	req := c.c.NewRequest(c.name, "CmsDocumentService.Lists", in)
	out := new(CmsDocumentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CmsDocumentService service

type CmsDocumentServiceHandler interface {
	//创建
	Create(context.Context, *CmsDocument, *Response) error
	//匹配多条件删除
	Delete(context.Context, *CmsDocument, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *CmsDocument, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *CmsDocumentBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *CmsDocument, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *CmsDocument, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *CmsDocumentUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *CmsDocumentFilter, *CmsDocumentResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *CmsDocumentFilter, *CmsDocumentResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *CmsDocumentListRequest, *CmsDocumentResponse) error
}

func RegisterCmsDocumentServiceHandler(s server.Server, hdlr CmsDocumentServiceHandler, opts ...server.HandlerOption) error {
	type cmsDocumentService interface {
		Create(ctx context.Context, in *CmsDocument, out *Response) error
		Delete(ctx context.Context, in *CmsDocument, out *Response) error
		DeleteById(ctx context.Context, in *CmsDocument, out *Response) error
		DeleteByIds(ctx context.Context, in *CmsDocumentBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *CmsDocument, out *Response) error
		UpdateAll(ctx context.Context, in *CmsDocument, out *Response) error
		UpdateFields(ctx context.Context, in *CmsDocumentUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *CmsDocumentFilter, out *CmsDocumentResponse) error
		Find(ctx context.Context, in *CmsDocumentFilter, out *CmsDocumentResponse) error
		Lists(ctx context.Context, in *CmsDocumentListRequest, out *CmsDocumentResponse) error
	}
	type CmsDocumentService struct {
		cmsDocumentService
	}
	h := &cmsDocumentServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "CmsDocumentService.Create",
		Path:    []string{"/v1/cmsDocument/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "CmsDocumentService.Delete",
		Path:    []string{"/v1/cmsDocument/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "CmsDocumentService.DeleteById",
		Path:    []string{"/v1/cmsDocument/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "CmsDocumentService.DeleteByIds",
		Path:    []string{"/v1/cmsDocument/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "CmsDocumentService.Update",
		Path:    []string{"/v1/cmsDocument/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "CmsDocumentService.UpdateAll",
		Path:    []string{"/v1/cmsDocument/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "CmsDocumentService.UpdateFields",
		Path:    []string{"/v1/cmsDocument/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "CmsDocumentService.FindById",
		Path:    []string{"/v1/cmsDocument/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "CmsDocumentService.Find",
		Path:    []string{"/v1/cmsDocument/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "CmsDocumentService.Lists",
		Path:    []string{"/v1/cmsDocument/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&CmsDocumentService{h}, opts...))
}

type cmsDocumentServiceHandler struct {
	CmsDocumentServiceHandler
}

func (h *cmsDocumentServiceHandler) Create(ctx context.Context, in *CmsDocument, out *Response) error {
	return h.CmsDocumentServiceHandler.Create(ctx, in, out)
}

func (h *cmsDocumentServiceHandler) Delete(ctx context.Context, in *CmsDocument, out *Response) error {
	return h.CmsDocumentServiceHandler.Delete(ctx, in, out)
}

func (h *cmsDocumentServiceHandler) DeleteById(ctx context.Context, in *CmsDocument, out *Response) error {
	return h.CmsDocumentServiceHandler.DeleteById(ctx, in, out)
}

func (h *cmsDocumentServiceHandler) DeleteByIds(ctx context.Context, in *CmsDocumentBatchDeleteRequest, out *Response) error {
	return h.CmsDocumentServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *cmsDocumentServiceHandler) Update(ctx context.Context, in *CmsDocument, out *Response) error {
	return h.CmsDocumentServiceHandler.Update(ctx, in, out)
}

func (h *cmsDocumentServiceHandler) UpdateAll(ctx context.Context, in *CmsDocument, out *Response) error {
	return h.CmsDocumentServiceHandler.UpdateAll(ctx, in, out)
}

func (h *cmsDocumentServiceHandler) UpdateFields(ctx context.Context, in *CmsDocumentUpdateFieldsRequest, out *Response) error {
	return h.CmsDocumentServiceHandler.UpdateFields(ctx, in, out)
}

func (h *cmsDocumentServiceHandler) FindById(ctx context.Context, in *CmsDocumentFilter, out *CmsDocumentResponse) error {
	return h.CmsDocumentServiceHandler.FindById(ctx, in, out)
}

func (h *cmsDocumentServiceHandler) Find(ctx context.Context, in *CmsDocumentFilter, out *CmsDocumentResponse) error {
	return h.CmsDocumentServiceHandler.Find(ctx, in, out)
}

func (h *cmsDocumentServiceHandler) Lists(ctx context.Context, in *CmsDocumentListRequest, out *CmsDocumentResponse) error {
	return h.CmsDocumentServiceHandler.Lists(ctx, in, out)
}
