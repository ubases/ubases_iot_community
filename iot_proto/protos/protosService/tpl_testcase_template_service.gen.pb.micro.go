// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: tpl_testcase_template_service.gen.proto

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

// Api Endpoints for TplTestcaseTemplateService service

func NewTplTestcaseTemplateServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "TplTestcaseTemplateService.Create",
			Path:    []string{"/v1/tplTestcaseTemplate/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "TplTestcaseTemplateService.Delete",
			Path:    []string{"/v1/tplTestcaseTemplate/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "TplTestcaseTemplateService.DeleteById",
			Path:    []string{"/v1/tplTestcaseTemplate/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "TplTestcaseTemplateService.DeleteByIds",
			Path:    []string{"/v1/tplTestcaseTemplate/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "TplTestcaseTemplateService.Update",
			Path:    []string{"/v1/tplTestcaseTemplate/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "TplTestcaseTemplateService.UpdateAll",
			Path:    []string{"/v1/tplTestcaseTemplate/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "TplTestcaseTemplateService.UpdateFields",
			Path:    []string{"/v1/tplTestcaseTemplate/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "TplTestcaseTemplateService.FindById",
			Path:    []string{"/v1/tplTestcaseTemplate/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "TplTestcaseTemplateService.Find",
			Path:    []string{"/v1/tplTestcaseTemplate/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "TplTestcaseTemplateService.Lists",
			Path:    []string{"/v1/tplTestcaseTemplate/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for TplTestcaseTemplateService service

type TplTestcaseTemplateService interface {
	//创建
	Create(ctx context.Context, in *TplTestcaseTemplate, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *TplTestcaseTemplate, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *TplTestcaseTemplate, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *TplTestcaseTemplateBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *TplTestcaseTemplate, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *TplTestcaseTemplate, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *TplTestcaseTemplateUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *TplTestcaseTemplateFilter, opts ...client.CallOption) (*TplTestcaseTemplateResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *TplTestcaseTemplateFilter, opts ...client.CallOption) (*TplTestcaseTemplateResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *TplTestcaseTemplateListRequest, opts ...client.CallOption) (*TplTestcaseTemplateResponse, error)
}

type tplTestcaseTemplateService struct {
	c    client.Client
	name string
}

func NewTplTestcaseTemplateService(name string, c client.Client) TplTestcaseTemplateService {
	return &tplTestcaseTemplateService{
		c:    c,
		name: name,
	}
}

func (c *tplTestcaseTemplateService) Create(ctx context.Context, in *TplTestcaseTemplate, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TplTestcaseTemplateService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplTestcaseTemplateService) Delete(ctx context.Context, in *TplTestcaseTemplate, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TplTestcaseTemplateService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplTestcaseTemplateService) DeleteById(ctx context.Context, in *TplTestcaseTemplate, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TplTestcaseTemplateService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplTestcaseTemplateService) DeleteByIds(ctx context.Context, in *TplTestcaseTemplateBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TplTestcaseTemplateService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplTestcaseTemplateService) Update(ctx context.Context, in *TplTestcaseTemplate, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TplTestcaseTemplateService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplTestcaseTemplateService) UpdateAll(ctx context.Context, in *TplTestcaseTemplate, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TplTestcaseTemplateService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplTestcaseTemplateService) UpdateFields(ctx context.Context, in *TplTestcaseTemplateUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TplTestcaseTemplateService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplTestcaseTemplateService) FindById(ctx context.Context, in *TplTestcaseTemplateFilter, opts ...client.CallOption) (*TplTestcaseTemplateResponse, error) {
	req := c.c.NewRequest(c.name, "TplTestcaseTemplateService.FindById", in)
	out := new(TplTestcaseTemplateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplTestcaseTemplateService) Find(ctx context.Context, in *TplTestcaseTemplateFilter, opts ...client.CallOption) (*TplTestcaseTemplateResponse, error) {
	req := c.c.NewRequest(c.name, "TplTestcaseTemplateService.Find", in)
	out := new(TplTestcaseTemplateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplTestcaseTemplateService) Lists(ctx context.Context, in *TplTestcaseTemplateListRequest, opts ...client.CallOption) (*TplTestcaseTemplateResponse, error) {
	req := c.c.NewRequest(c.name, "TplTestcaseTemplateService.Lists", in)
	out := new(TplTestcaseTemplateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TplTestcaseTemplateService service

type TplTestcaseTemplateServiceHandler interface {
	//创建
	Create(context.Context, *TplTestcaseTemplate, *Response) error
	//匹配多条件删除
	Delete(context.Context, *TplTestcaseTemplate, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *TplTestcaseTemplate, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *TplTestcaseTemplateBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *TplTestcaseTemplate, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *TplTestcaseTemplate, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *TplTestcaseTemplateUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *TplTestcaseTemplateFilter, *TplTestcaseTemplateResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *TplTestcaseTemplateFilter, *TplTestcaseTemplateResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *TplTestcaseTemplateListRequest, *TplTestcaseTemplateResponse) error
}

func RegisterTplTestcaseTemplateServiceHandler(s server.Server, hdlr TplTestcaseTemplateServiceHandler, opts ...server.HandlerOption) error {
	type tplTestcaseTemplateService interface {
		Create(ctx context.Context, in *TplTestcaseTemplate, out *Response) error
		Delete(ctx context.Context, in *TplTestcaseTemplate, out *Response) error
		DeleteById(ctx context.Context, in *TplTestcaseTemplate, out *Response) error
		DeleteByIds(ctx context.Context, in *TplTestcaseTemplateBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *TplTestcaseTemplate, out *Response) error
		UpdateAll(ctx context.Context, in *TplTestcaseTemplate, out *Response) error
		UpdateFields(ctx context.Context, in *TplTestcaseTemplateUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *TplTestcaseTemplateFilter, out *TplTestcaseTemplateResponse) error
		Find(ctx context.Context, in *TplTestcaseTemplateFilter, out *TplTestcaseTemplateResponse) error
		Lists(ctx context.Context, in *TplTestcaseTemplateListRequest, out *TplTestcaseTemplateResponse) error
	}
	type TplTestcaseTemplateService struct {
		tplTestcaseTemplateService
	}
	h := &tplTestcaseTemplateServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "TplTestcaseTemplateService.Create",
		Path:    []string{"/v1/tplTestcaseTemplate/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "TplTestcaseTemplateService.Delete",
		Path:    []string{"/v1/tplTestcaseTemplate/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "TplTestcaseTemplateService.DeleteById",
		Path:    []string{"/v1/tplTestcaseTemplate/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "TplTestcaseTemplateService.DeleteByIds",
		Path:    []string{"/v1/tplTestcaseTemplate/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "TplTestcaseTemplateService.Update",
		Path:    []string{"/v1/tplTestcaseTemplate/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "TplTestcaseTemplateService.UpdateAll",
		Path:    []string{"/v1/tplTestcaseTemplate/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "TplTestcaseTemplateService.UpdateFields",
		Path:    []string{"/v1/tplTestcaseTemplate/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "TplTestcaseTemplateService.FindById",
		Path:    []string{"/v1/tplTestcaseTemplate/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "TplTestcaseTemplateService.Find",
		Path:    []string{"/v1/tplTestcaseTemplate/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "TplTestcaseTemplateService.Lists",
		Path:    []string{"/v1/tplTestcaseTemplate/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&TplTestcaseTemplateService{h}, opts...))
}

type tplTestcaseTemplateServiceHandler struct {
	TplTestcaseTemplateServiceHandler
}

func (h *tplTestcaseTemplateServiceHandler) Create(ctx context.Context, in *TplTestcaseTemplate, out *Response) error {
	return h.TplTestcaseTemplateServiceHandler.Create(ctx, in, out)
}

func (h *tplTestcaseTemplateServiceHandler) Delete(ctx context.Context, in *TplTestcaseTemplate, out *Response) error {
	return h.TplTestcaseTemplateServiceHandler.Delete(ctx, in, out)
}

func (h *tplTestcaseTemplateServiceHandler) DeleteById(ctx context.Context, in *TplTestcaseTemplate, out *Response) error {
	return h.TplTestcaseTemplateServiceHandler.DeleteById(ctx, in, out)
}

func (h *tplTestcaseTemplateServiceHandler) DeleteByIds(ctx context.Context, in *TplTestcaseTemplateBatchDeleteRequest, out *Response) error {
	return h.TplTestcaseTemplateServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *tplTestcaseTemplateServiceHandler) Update(ctx context.Context, in *TplTestcaseTemplate, out *Response) error {
	return h.TplTestcaseTemplateServiceHandler.Update(ctx, in, out)
}

func (h *tplTestcaseTemplateServiceHandler) UpdateAll(ctx context.Context, in *TplTestcaseTemplate, out *Response) error {
	return h.TplTestcaseTemplateServiceHandler.UpdateAll(ctx, in, out)
}

func (h *tplTestcaseTemplateServiceHandler) UpdateFields(ctx context.Context, in *TplTestcaseTemplateUpdateFieldsRequest, out *Response) error {
	return h.TplTestcaseTemplateServiceHandler.UpdateFields(ctx, in, out)
}

func (h *tplTestcaseTemplateServiceHandler) FindById(ctx context.Context, in *TplTestcaseTemplateFilter, out *TplTestcaseTemplateResponse) error {
	return h.TplTestcaseTemplateServiceHandler.FindById(ctx, in, out)
}

func (h *tplTestcaseTemplateServiceHandler) Find(ctx context.Context, in *TplTestcaseTemplateFilter, out *TplTestcaseTemplateResponse) error {
	return h.TplTestcaseTemplateServiceHandler.Find(ctx, in, out)
}

func (h *tplTestcaseTemplateServiceHandler) Lists(ctx context.Context, in *TplTestcaseTemplateListRequest, out *TplTestcaseTemplateResponse) error {
	return h.TplTestcaseTemplateServiceHandler.Lists(ctx, in, out)
}
