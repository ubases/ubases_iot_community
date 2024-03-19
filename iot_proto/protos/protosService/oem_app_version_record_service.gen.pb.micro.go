// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: oem_app_version_record_service.gen.proto

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

// Api Endpoints for OemAppVersionRecordService service

func NewOemAppVersionRecordServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "OemAppVersionRecordService.Create",
			Path:    []string{"/v1/oemAppVersionRecord/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppVersionRecordService.Delete",
			Path:    []string{"/v1/oemAppVersionRecord/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppVersionRecordService.DeleteById",
			Path:    []string{"/v1/oemAppVersionRecord/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppVersionRecordService.DeleteByIds",
			Path:    []string{"/v1/oemAppVersionRecord/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppVersionRecordService.Update",
			Path:    []string{"/v1/oemAppVersionRecord/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppVersionRecordService.UpdateAll",
			Path:    []string{"/v1/oemAppVersionRecord/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppVersionRecordService.UpdateFields",
			Path:    []string{"/v1/oemAppVersionRecord/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppVersionRecordService.FindById",
			Path:    []string{"/v1/oemAppVersionRecord/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppVersionRecordService.Find",
			Path:    []string{"/v1/oemAppVersionRecord/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppVersionRecordService.Lists",
			Path:    []string{"/v1/oemAppVersionRecord/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for OemAppVersionRecordService service

type OemAppVersionRecordService interface {
	//创建
	Create(ctx context.Context, in *OemAppVersionRecord, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *OemAppVersionRecord, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *OemAppVersionRecord, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *OemAppVersionRecordBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *OemAppVersionRecord, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *OemAppVersionRecord, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *OemAppVersionRecordUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *OemAppVersionRecordFilter, opts ...client.CallOption) (*OemAppVersionRecordResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *OemAppVersionRecordFilter, opts ...client.CallOption) (*OemAppVersionRecordResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *OemAppVersionRecordListRequest, opts ...client.CallOption) (*OemAppVersionRecordResponse, error)
}

type oemAppVersionRecordService struct {
	c    client.Client
	name string
}

func NewOemAppVersionRecordService(name string, c client.Client) OemAppVersionRecordService {
	return &oemAppVersionRecordService{
		c:    c,
		name: name,
	}
}

func (c *oemAppVersionRecordService) Create(ctx context.Context, in *OemAppVersionRecord, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppVersionRecordService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppVersionRecordService) Delete(ctx context.Context, in *OemAppVersionRecord, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppVersionRecordService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppVersionRecordService) DeleteById(ctx context.Context, in *OemAppVersionRecord, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppVersionRecordService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppVersionRecordService) DeleteByIds(ctx context.Context, in *OemAppVersionRecordBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppVersionRecordService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppVersionRecordService) Update(ctx context.Context, in *OemAppVersionRecord, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppVersionRecordService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppVersionRecordService) UpdateAll(ctx context.Context, in *OemAppVersionRecord, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppVersionRecordService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppVersionRecordService) UpdateFields(ctx context.Context, in *OemAppVersionRecordUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppVersionRecordService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppVersionRecordService) FindById(ctx context.Context, in *OemAppVersionRecordFilter, opts ...client.CallOption) (*OemAppVersionRecordResponse, error) {
	req := c.c.NewRequest(c.name, "OemAppVersionRecordService.FindById", in)
	out := new(OemAppVersionRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppVersionRecordService) Find(ctx context.Context, in *OemAppVersionRecordFilter, opts ...client.CallOption) (*OemAppVersionRecordResponse, error) {
	req := c.c.NewRequest(c.name, "OemAppVersionRecordService.Find", in)
	out := new(OemAppVersionRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppVersionRecordService) Lists(ctx context.Context, in *OemAppVersionRecordListRequest, opts ...client.CallOption) (*OemAppVersionRecordResponse, error) {
	req := c.c.NewRequest(c.name, "OemAppVersionRecordService.Lists", in)
	out := new(OemAppVersionRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OemAppVersionRecordService service

type OemAppVersionRecordServiceHandler interface {
	//创建
	Create(context.Context, *OemAppVersionRecord, *Response) error
	//匹配多条件删除
	Delete(context.Context, *OemAppVersionRecord, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *OemAppVersionRecord, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *OemAppVersionRecordBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *OemAppVersionRecord, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *OemAppVersionRecord, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *OemAppVersionRecordUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *OemAppVersionRecordFilter, *OemAppVersionRecordResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *OemAppVersionRecordFilter, *OemAppVersionRecordResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *OemAppVersionRecordListRequest, *OemAppVersionRecordResponse) error
}

func RegisterOemAppVersionRecordServiceHandler(s server.Server, hdlr OemAppVersionRecordServiceHandler, opts ...server.HandlerOption) error {
	type oemAppVersionRecordService interface {
		Create(ctx context.Context, in *OemAppVersionRecord, out *Response) error
		Delete(ctx context.Context, in *OemAppVersionRecord, out *Response) error
		DeleteById(ctx context.Context, in *OemAppVersionRecord, out *Response) error
		DeleteByIds(ctx context.Context, in *OemAppVersionRecordBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *OemAppVersionRecord, out *Response) error
		UpdateAll(ctx context.Context, in *OemAppVersionRecord, out *Response) error
		UpdateFields(ctx context.Context, in *OemAppVersionRecordUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *OemAppVersionRecordFilter, out *OemAppVersionRecordResponse) error
		Find(ctx context.Context, in *OemAppVersionRecordFilter, out *OemAppVersionRecordResponse) error
		Lists(ctx context.Context, in *OemAppVersionRecordListRequest, out *OemAppVersionRecordResponse) error
	}
	type OemAppVersionRecordService struct {
		oemAppVersionRecordService
	}
	h := &oemAppVersionRecordServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppVersionRecordService.Create",
		Path:    []string{"/v1/oemAppVersionRecord/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppVersionRecordService.Delete",
		Path:    []string{"/v1/oemAppVersionRecord/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppVersionRecordService.DeleteById",
		Path:    []string{"/v1/oemAppVersionRecord/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppVersionRecordService.DeleteByIds",
		Path:    []string{"/v1/oemAppVersionRecord/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppVersionRecordService.Update",
		Path:    []string{"/v1/oemAppVersionRecord/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppVersionRecordService.UpdateAll",
		Path:    []string{"/v1/oemAppVersionRecord/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppVersionRecordService.UpdateFields",
		Path:    []string{"/v1/oemAppVersionRecord/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppVersionRecordService.FindById",
		Path:    []string{"/v1/oemAppVersionRecord/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppVersionRecordService.Find",
		Path:    []string{"/v1/oemAppVersionRecord/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppVersionRecordService.Lists",
		Path:    []string{"/v1/oemAppVersionRecord/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&OemAppVersionRecordService{h}, opts...))
}

type oemAppVersionRecordServiceHandler struct {
	OemAppVersionRecordServiceHandler
}

func (h *oemAppVersionRecordServiceHandler) Create(ctx context.Context, in *OemAppVersionRecord, out *Response) error {
	return h.OemAppVersionRecordServiceHandler.Create(ctx, in, out)
}

func (h *oemAppVersionRecordServiceHandler) Delete(ctx context.Context, in *OemAppVersionRecord, out *Response) error {
	return h.OemAppVersionRecordServiceHandler.Delete(ctx, in, out)
}

func (h *oemAppVersionRecordServiceHandler) DeleteById(ctx context.Context, in *OemAppVersionRecord, out *Response) error {
	return h.OemAppVersionRecordServiceHandler.DeleteById(ctx, in, out)
}

func (h *oemAppVersionRecordServiceHandler) DeleteByIds(ctx context.Context, in *OemAppVersionRecordBatchDeleteRequest, out *Response) error {
	return h.OemAppVersionRecordServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *oemAppVersionRecordServiceHandler) Update(ctx context.Context, in *OemAppVersionRecord, out *Response) error {
	return h.OemAppVersionRecordServiceHandler.Update(ctx, in, out)
}

func (h *oemAppVersionRecordServiceHandler) UpdateAll(ctx context.Context, in *OemAppVersionRecord, out *Response) error {
	return h.OemAppVersionRecordServiceHandler.UpdateAll(ctx, in, out)
}

func (h *oemAppVersionRecordServiceHandler) UpdateFields(ctx context.Context, in *OemAppVersionRecordUpdateFieldsRequest, out *Response) error {
	return h.OemAppVersionRecordServiceHandler.UpdateFields(ctx, in, out)
}

func (h *oemAppVersionRecordServiceHandler) FindById(ctx context.Context, in *OemAppVersionRecordFilter, out *OemAppVersionRecordResponse) error {
	return h.OemAppVersionRecordServiceHandler.FindById(ctx, in, out)
}

func (h *oemAppVersionRecordServiceHandler) Find(ctx context.Context, in *OemAppVersionRecordFilter, out *OemAppVersionRecordResponse) error {
	return h.OemAppVersionRecordServiceHandler.Find(ctx, in, out)
}

func (h *oemAppVersionRecordServiceHandler) Lists(ctx context.Context, in *OemAppVersionRecordListRequest, out *OemAppVersionRecordResponse) error {
	return h.OemAppVersionRecordServiceHandler.Lists(ctx, in, out)
}
