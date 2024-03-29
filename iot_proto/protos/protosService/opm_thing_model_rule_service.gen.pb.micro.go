// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: opm_thing_model_rule_service.gen.proto

package protosService

import (
	fmt "fmt"
	math "math"

	proto "google.golang.org/protobuf/proto"

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

// Api Endpoints for OpmThingModelRuleService service

func NewOpmThingModelRuleServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "OpmThingModelRuleService.Create",
			Path:    []string{"/v1/opmThingModelRule/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OpmThingModelRuleService.Delete",
			Path:    []string{"/v1/opmThingModelRule/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OpmThingModelRuleService.DeleteById",
			Path:    []string{"/v1/opmThingModelRule/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OpmThingModelRuleService.DeleteByIds",
			Path:    []string{"/v1/opmThingModelRule/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OpmThingModelRuleService.Update",
			Path:    []string{"/v1/opmThingModelRule/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OpmThingModelRuleService.UpdateAll",
			Path:    []string{"/v1/opmThingModelRule/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OpmThingModelRuleService.UpdateFields",
			Path:    []string{"/v1/opmThingModelRule/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OpmThingModelRuleService.FindById",
			Path:    []string{"/v1/opmThingModelRule/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OpmThingModelRuleService.Find",
			Path:    []string{"/v1/opmThingModelRule/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OpmThingModelRuleService.Lists",
			Path:    []string{"/v1/opmThingModelRule/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OpmThingModelRuleService.UpdateStatus",
			Path:    []string{"/v1/opmThingModelRule/updateStatus"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for OpmThingModelRuleService service

type OpmThingModelRuleService interface {
	//创建
	Create(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *OpmThingModelRuleBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *OpmThingModelRuleUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *OpmThingModelRuleFilter, opts ...client.CallOption) (*OpmThingModelRuleResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *OpmThingModelRuleFilter, opts ...client.CallOption) (*OpmThingModelRuleResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *OpmThingModelRuleListRequest, opts ...client.CallOption) (*OpmThingModelRuleResponse, error)
	//修改状态
	UpdateStatus(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error)
}

type opmThingModelRuleService struct {
	c    client.Client
	name string
}

func NewOpmThingModelRuleService(name string, c client.Client) OpmThingModelRuleService {
	return &opmThingModelRuleService{
		c:    c,
		name: name,
	}
}

func (c *opmThingModelRuleService) Create(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpmThingModelRuleService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opmThingModelRuleService) Delete(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpmThingModelRuleService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opmThingModelRuleService) DeleteById(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpmThingModelRuleService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opmThingModelRuleService) DeleteByIds(ctx context.Context, in *OpmThingModelRuleBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpmThingModelRuleService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opmThingModelRuleService) Update(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpmThingModelRuleService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opmThingModelRuleService) UpdateAll(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpmThingModelRuleService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opmThingModelRuleService) UpdateFields(ctx context.Context, in *OpmThingModelRuleUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpmThingModelRuleService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opmThingModelRuleService) FindById(ctx context.Context, in *OpmThingModelRuleFilter, opts ...client.CallOption) (*OpmThingModelRuleResponse, error) {
	req := c.c.NewRequest(c.name, "OpmThingModelRuleService.FindById", in)
	out := new(OpmThingModelRuleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opmThingModelRuleService) Find(ctx context.Context, in *OpmThingModelRuleFilter, opts ...client.CallOption) (*OpmThingModelRuleResponse, error) {
	req := c.c.NewRequest(c.name, "OpmThingModelRuleService.Find", in)
	out := new(OpmThingModelRuleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opmThingModelRuleService) Lists(ctx context.Context, in *OpmThingModelRuleListRequest, opts ...client.CallOption) (*OpmThingModelRuleResponse, error) {
	req := c.c.NewRequest(c.name, "OpmThingModelRuleService.Lists", in)
	out := new(OpmThingModelRuleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opmThingModelRuleService) UpdateStatus(ctx context.Context, in *OpmThingModelRule, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpmThingModelRuleService.UpdateStatus", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OpmThingModelRuleService service

type OpmThingModelRuleServiceHandler interface {
	//创建
	Create(context.Context, *OpmThingModelRule, *Response) error
	//匹配多条件删除
	Delete(context.Context, *OpmThingModelRule, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *OpmThingModelRule, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *OpmThingModelRuleBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *OpmThingModelRule, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *OpmThingModelRule, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *OpmThingModelRuleUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *OpmThingModelRuleFilter, *OpmThingModelRuleResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *OpmThingModelRuleFilter, *OpmThingModelRuleResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *OpmThingModelRuleListRequest, *OpmThingModelRuleResponse) error
	//修改状态
	UpdateStatus(context.Context, *OpmThingModelRule, *Response) error
}

func RegisterOpmThingModelRuleServiceHandler(s server.Server, hdlr OpmThingModelRuleServiceHandler, opts ...server.HandlerOption) error {
	type opmThingModelRuleService interface {
		Create(ctx context.Context, in *OpmThingModelRule, out *Response) error
		Delete(ctx context.Context, in *OpmThingModelRule, out *Response) error
		DeleteById(ctx context.Context, in *OpmThingModelRule, out *Response) error
		DeleteByIds(ctx context.Context, in *OpmThingModelRuleBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *OpmThingModelRule, out *Response) error
		UpdateAll(ctx context.Context, in *OpmThingModelRule, out *Response) error
		UpdateFields(ctx context.Context, in *OpmThingModelRuleUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *OpmThingModelRuleFilter, out *OpmThingModelRuleResponse) error
		Find(ctx context.Context, in *OpmThingModelRuleFilter, out *OpmThingModelRuleResponse) error
		Lists(ctx context.Context, in *OpmThingModelRuleListRequest, out *OpmThingModelRuleResponse) error
		UpdateStatus(ctx context.Context, in *OpmThingModelRule, out *Response) error
	}
	type OpmThingModelRuleService struct {
		opmThingModelRuleService
	}
	h := &opmThingModelRuleServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OpmThingModelRuleService.Create",
		Path:    []string{"/v1/opmThingModelRule/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OpmThingModelRuleService.Delete",
		Path:    []string{"/v1/opmThingModelRule/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OpmThingModelRuleService.DeleteById",
		Path:    []string{"/v1/opmThingModelRule/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OpmThingModelRuleService.DeleteByIds",
		Path:    []string{"/v1/opmThingModelRule/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OpmThingModelRuleService.Update",
		Path:    []string{"/v1/opmThingModelRule/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OpmThingModelRuleService.UpdateAll",
		Path:    []string{"/v1/opmThingModelRule/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OpmThingModelRuleService.UpdateFields",
		Path:    []string{"/v1/opmThingModelRule/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OpmThingModelRuleService.FindById",
		Path:    []string{"/v1/opmThingModelRule/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OpmThingModelRuleService.Find",
		Path:    []string{"/v1/opmThingModelRule/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OpmThingModelRuleService.Lists",
		Path:    []string{"/v1/opmThingModelRule/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OpmThingModelRuleService.UpdateStatus",
		Path:    []string{"/v1/opmThingModelRule/updateStatus"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&OpmThingModelRuleService{h}, opts...))
}

type opmThingModelRuleServiceHandler struct {
	OpmThingModelRuleServiceHandler
}

func (h *opmThingModelRuleServiceHandler) Create(ctx context.Context, in *OpmThingModelRule, out *Response) error {
	return h.OpmThingModelRuleServiceHandler.Create(ctx, in, out)
}

func (h *opmThingModelRuleServiceHandler) Delete(ctx context.Context, in *OpmThingModelRule, out *Response) error {
	return h.OpmThingModelRuleServiceHandler.Delete(ctx, in, out)
}

func (h *opmThingModelRuleServiceHandler) DeleteById(ctx context.Context, in *OpmThingModelRule, out *Response) error {
	return h.OpmThingModelRuleServiceHandler.DeleteById(ctx, in, out)
}

func (h *opmThingModelRuleServiceHandler) DeleteByIds(ctx context.Context, in *OpmThingModelRuleBatchDeleteRequest, out *Response) error {
	return h.OpmThingModelRuleServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *opmThingModelRuleServiceHandler) Update(ctx context.Context, in *OpmThingModelRule, out *Response) error {
	return h.OpmThingModelRuleServiceHandler.Update(ctx, in, out)
}

func (h *opmThingModelRuleServiceHandler) UpdateAll(ctx context.Context, in *OpmThingModelRule, out *Response) error {
	return h.OpmThingModelRuleServiceHandler.UpdateAll(ctx, in, out)
}

func (h *opmThingModelRuleServiceHandler) UpdateFields(ctx context.Context, in *OpmThingModelRuleUpdateFieldsRequest, out *Response) error {
	return h.OpmThingModelRuleServiceHandler.UpdateFields(ctx, in, out)
}

func (h *opmThingModelRuleServiceHandler) FindById(ctx context.Context, in *OpmThingModelRuleFilter, out *OpmThingModelRuleResponse) error {
	return h.OpmThingModelRuleServiceHandler.FindById(ctx, in, out)
}

func (h *opmThingModelRuleServiceHandler) Find(ctx context.Context, in *OpmThingModelRuleFilter, out *OpmThingModelRuleResponse) error {
	return h.OpmThingModelRuleServiceHandler.Find(ctx, in, out)
}

func (h *opmThingModelRuleServiceHandler) Lists(ctx context.Context, in *OpmThingModelRuleListRequest, out *OpmThingModelRuleResponse) error {
	return h.OpmThingModelRuleServiceHandler.Lists(ctx, in, out)
}

func (h *opmThingModelRuleServiceHandler) UpdateStatus(ctx context.Context, in *OpmThingModelRule, out *Response) error {
	return h.OpmThingModelRuleServiceHandler.UpdateStatus(ctx, in, out)
}
