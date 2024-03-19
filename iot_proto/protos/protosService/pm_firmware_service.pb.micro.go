// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pm_firmware_service.proto

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

// Api Endpoints for PmFirmwareService service

func NewPmFirmwareServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "PmFirmwareService.Create",
			Path:    []string{"/v1/pmFirmware/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmFirmwareService.Delete",
			Path:    []string{"/v1/pmFirmware/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmFirmwareService.DeleteById",
			Path:    []string{"/v1/pmFirmware/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmFirmwareService.DeleteByIds",
			Path:    []string{"/v1/pmFirmware/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmFirmwareService.Update",
			Path:    []string{"/v1/pmFirmware/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmFirmwareService.UpdateAll",
			Path:    []string{"/v1/pmFirmware/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmFirmwareService.UpdateFields",
			Path:    []string{"/v1/pmFirmware/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmFirmwareService.FindById",
			Path:    []string{"/v1/pmFirmware/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmFirmwareService.Find",
			Path:    []string{"/v1/pmFirmware/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmFirmwareService.Lists",
			Path:    []string{"/v1/pmFirmware/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "PmFirmwareService.UpdateStatus",
			Path:    []string{"/v1/pmFirmware/UpdateStatus"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for PmFirmwareService service

type PmFirmwareService interface {
	//创建
	Create(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *PmFirmwareBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *PmFirmwareUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *PmFirmwareFilter, opts ...client.CallOption) (*PmFirmwareResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *PmFirmwareFilter, opts ...client.CallOption) (*PmFirmwareResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *PmFirmwareListRequest, opts ...client.CallOption) (*PmFirmwareResponse, error)
	//修改状态
	UpdateStatus(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error)
}

type pmFirmwareService struct {
	c    client.Client
	name string
}

func NewPmFirmwareService(name string, c client.Client) PmFirmwareService {
	return &pmFirmwareService{
		c:    c,
		name: name,
	}
}

func (c *pmFirmwareService) Create(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmFirmwareService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmFirmwareService) Delete(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmFirmwareService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmFirmwareService) DeleteById(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmFirmwareService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmFirmwareService) DeleteByIds(ctx context.Context, in *PmFirmwareBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmFirmwareService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmFirmwareService) Update(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmFirmwareService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmFirmwareService) UpdateAll(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmFirmwareService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmFirmwareService) UpdateFields(ctx context.Context, in *PmFirmwareUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmFirmwareService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmFirmwareService) FindById(ctx context.Context, in *PmFirmwareFilter, opts ...client.CallOption) (*PmFirmwareResponse, error) {
	req := c.c.NewRequest(c.name, "PmFirmwareService.FindById", in)
	out := new(PmFirmwareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmFirmwareService) Find(ctx context.Context, in *PmFirmwareFilter, opts ...client.CallOption) (*PmFirmwareResponse, error) {
	req := c.c.NewRequest(c.name, "PmFirmwareService.Find", in)
	out := new(PmFirmwareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmFirmwareService) Lists(ctx context.Context, in *PmFirmwareListRequest, opts ...client.CallOption) (*PmFirmwareResponse, error) {
	req := c.c.NewRequest(c.name, "PmFirmwareService.Lists", in)
	out := new(PmFirmwareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmFirmwareService) UpdateStatus(ctx context.Context, in *PmFirmware, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PmFirmwareService.UpdateStatus", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PmFirmwareService service

type PmFirmwareServiceHandler interface {
	//创建
	Create(context.Context, *PmFirmware, *Response) error
	//匹配多条件删除
	Delete(context.Context, *PmFirmware, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *PmFirmware, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *PmFirmwareBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *PmFirmware, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *PmFirmware, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *PmFirmwareUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *PmFirmwareFilter, *PmFirmwareResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *PmFirmwareFilter, *PmFirmwareResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *PmFirmwareListRequest, *PmFirmwareResponse) error
	//修改状态
	UpdateStatus(context.Context, *PmFirmware, *Response) error
}

func RegisterPmFirmwareServiceHandler(s server.Server, hdlr PmFirmwareServiceHandler, opts ...server.HandlerOption) error {
	type pmFirmwareService interface {
		Create(ctx context.Context, in *PmFirmware, out *Response) error
		Delete(ctx context.Context, in *PmFirmware, out *Response) error
		DeleteById(ctx context.Context, in *PmFirmware, out *Response) error
		DeleteByIds(ctx context.Context, in *PmFirmwareBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *PmFirmware, out *Response) error
		UpdateAll(ctx context.Context, in *PmFirmware, out *Response) error
		UpdateFields(ctx context.Context, in *PmFirmwareUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *PmFirmwareFilter, out *PmFirmwareResponse) error
		Find(ctx context.Context, in *PmFirmwareFilter, out *PmFirmwareResponse) error
		Lists(ctx context.Context, in *PmFirmwareListRequest, out *PmFirmwareResponse) error
		UpdateStatus(ctx context.Context, in *PmFirmware, out *Response) error
	}
	type PmFirmwareService struct {
		pmFirmwareService
	}
	h := &pmFirmwareServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmFirmwareService.Create",
		Path:    []string{"/v1/pmFirmware/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmFirmwareService.Delete",
		Path:    []string{"/v1/pmFirmware/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmFirmwareService.DeleteById",
		Path:    []string{"/v1/pmFirmware/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmFirmwareService.DeleteByIds",
		Path:    []string{"/v1/pmFirmware/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmFirmwareService.Update",
		Path:    []string{"/v1/pmFirmware/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmFirmwareService.UpdateAll",
		Path:    []string{"/v1/pmFirmware/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmFirmwareService.UpdateFields",
		Path:    []string{"/v1/pmFirmware/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmFirmwareService.FindById",
		Path:    []string{"/v1/pmFirmware/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmFirmwareService.Find",
		Path:    []string{"/v1/pmFirmware/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmFirmwareService.Lists",
		Path:    []string{"/v1/pmFirmware/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PmFirmwareService.UpdateStatus",
		Path:    []string{"/v1/pmFirmware/UpdateStatus"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&PmFirmwareService{h}, opts...))
}

type pmFirmwareServiceHandler struct {
	PmFirmwareServiceHandler
}

func (h *pmFirmwareServiceHandler) Create(ctx context.Context, in *PmFirmware, out *Response) error {
	return h.PmFirmwareServiceHandler.Create(ctx, in, out)
}

func (h *pmFirmwareServiceHandler) Delete(ctx context.Context, in *PmFirmware, out *Response) error {
	return h.PmFirmwareServiceHandler.Delete(ctx, in, out)
}

func (h *pmFirmwareServiceHandler) DeleteById(ctx context.Context, in *PmFirmware, out *Response) error {
	return h.PmFirmwareServiceHandler.DeleteById(ctx, in, out)
}

func (h *pmFirmwareServiceHandler) DeleteByIds(ctx context.Context, in *PmFirmwareBatchDeleteRequest, out *Response) error {
	return h.PmFirmwareServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *pmFirmwareServiceHandler) Update(ctx context.Context, in *PmFirmware, out *Response) error {
	return h.PmFirmwareServiceHandler.Update(ctx, in, out)
}

func (h *pmFirmwareServiceHandler) UpdateAll(ctx context.Context, in *PmFirmware, out *Response) error {
	return h.PmFirmwareServiceHandler.UpdateAll(ctx, in, out)
}

func (h *pmFirmwareServiceHandler) UpdateFields(ctx context.Context, in *PmFirmwareUpdateFieldsRequest, out *Response) error {
	return h.PmFirmwareServiceHandler.UpdateFields(ctx, in, out)
}

func (h *pmFirmwareServiceHandler) FindById(ctx context.Context, in *PmFirmwareFilter, out *PmFirmwareResponse) error {
	return h.PmFirmwareServiceHandler.FindById(ctx, in, out)
}

func (h *pmFirmwareServiceHandler) Find(ctx context.Context, in *PmFirmwareFilter, out *PmFirmwareResponse) error {
	return h.PmFirmwareServiceHandler.Find(ctx, in, out)
}

func (h *pmFirmwareServiceHandler) Lists(ctx context.Context, in *PmFirmwareListRequest, out *PmFirmwareResponse) error {
	return h.PmFirmwareServiceHandler.Lists(ctx, in, out)
}

func (h *pmFirmwareServiceHandler) UpdateStatus(ctx context.Context, in *PmFirmware, out *Response) error {
	return h.PmFirmwareServiceHandler.UpdateStatus(ctx, in, out)
}
