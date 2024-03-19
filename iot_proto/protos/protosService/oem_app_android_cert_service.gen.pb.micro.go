// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: oem_app_android_cert_service.gen.proto

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

// Api Endpoints for OemAppAndroidCertService service

func NewOemAppAndroidCertServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "OemAppAndroidCertService.Create",
			Path:    []string{"/v1/oemAppAndroidCert/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppAndroidCertService.Delete",
			Path:    []string{"/v1/oemAppAndroidCert/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppAndroidCertService.DeleteById",
			Path:    []string{"/v1/oemAppAndroidCert/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppAndroidCertService.DeleteByIds",
			Path:    []string{"/v1/oemAppAndroidCert/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppAndroidCertService.Update",
			Path:    []string{"/v1/oemAppAndroidCert/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppAndroidCertService.UpdateAll",
			Path:    []string{"/v1/oemAppAndroidCert/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppAndroidCertService.UpdateFields",
			Path:    []string{"/v1/oemAppAndroidCert/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppAndroidCertService.FindById",
			Path:    []string{"/v1/oemAppAndroidCert/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppAndroidCertService.Find",
			Path:    []string{"/v1/oemAppAndroidCert/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppAndroidCertService.Lists",
			Path:    []string{"/v1/oemAppAndroidCert/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppAndroidCertService.Regenerate",
			Path:    []string{"/v1/oemAppAndroidCert/regenerate"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for OemAppAndroidCertService service

type OemAppAndroidCertService interface {
	//创建
	Create(ctx context.Context, in *OemAppAndroidCert, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *OemAppAndroidCert, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *OemAppAndroidCert, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *OemAppAndroidCertBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *OemAppAndroidCert, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *OemAppAndroidCert, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *OemAppAndroidCertUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *OemAppAndroidCertFilter, opts ...client.CallOption) (*OemAppAndroidCertResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *OemAppAndroidCertFilter, opts ...client.CallOption) (*OemAppAndroidCertResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *OemAppAndroidCertListRequest, opts ...client.CallOption) (*OemAppAndroidCertResponse, error)
	//重新生成keystore文件
	Regenerate(ctx context.Context, in *OemAppAndroidCertFilter, opts ...client.CallOption) (*OemAppAndroidCertResponse, error)
}

type oemAppAndroidCertService struct {
	c    client.Client
	name string
}

func NewOemAppAndroidCertService(name string, c client.Client) OemAppAndroidCertService {
	return &oemAppAndroidCertService{
		c:    c,
		name: name,
	}
}

func (c *oemAppAndroidCertService) Create(ctx context.Context, in *OemAppAndroidCert, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppAndroidCertService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppAndroidCertService) Delete(ctx context.Context, in *OemAppAndroidCert, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppAndroidCertService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppAndroidCertService) DeleteById(ctx context.Context, in *OemAppAndroidCert, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppAndroidCertService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppAndroidCertService) DeleteByIds(ctx context.Context, in *OemAppAndroidCertBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppAndroidCertService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppAndroidCertService) Update(ctx context.Context, in *OemAppAndroidCert, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppAndroidCertService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppAndroidCertService) UpdateAll(ctx context.Context, in *OemAppAndroidCert, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppAndroidCertService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppAndroidCertService) UpdateFields(ctx context.Context, in *OemAppAndroidCertUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppAndroidCertService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppAndroidCertService) FindById(ctx context.Context, in *OemAppAndroidCertFilter, opts ...client.CallOption) (*OemAppAndroidCertResponse, error) {
	req := c.c.NewRequest(c.name, "OemAppAndroidCertService.FindById", in)
	out := new(OemAppAndroidCertResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppAndroidCertService) Find(ctx context.Context, in *OemAppAndroidCertFilter, opts ...client.CallOption) (*OemAppAndroidCertResponse, error) {
	req := c.c.NewRequest(c.name, "OemAppAndroidCertService.Find", in)
	out := new(OemAppAndroidCertResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppAndroidCertService) Lists(ctx context.Context, in *OemAppAndroidCertListRequest, opts ...client.CallOption) (*OemAppAndroidCertResponse, error) {
	req := c.c.NewRequest(c.name, "OemAppAndroidCertService.Lists", in)
	out := new(OemAppAndroidCertResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppAndroidCertService) Regenerate(ctx context.Context, in *OemAppAndroidCertFilter, opts ...client.CallOption) (*OemAppAndroidCertResponse, error) {
	req := c.c.NewRequest(c.name, "OemAppAndroidCertService.Regenerate", in)
	out := new(OemAppAndroidCertResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OemAppAndroidCertService service

type OemAppAndroidCertServiceHandler interface {
	//创建
	Create(context.Context, *OemAppAndroidCert, *Response) error
	//匹配多条件删除
	Delete(context.Context, *OemAppAndroidCert, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *OemAppAndroidCert, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *OemAppAndroidCertBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *OemAppAndroidCert, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *OemAppAndroidCert, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *OemAppAndroidCertUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *OemAppAndroidCertFilter, *OemAppAndroidCertResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *OemAppAndroidCertFilter, *OemAppAndroidCertResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *OemAppAndroidCertListRequest, *OemAppAndroidCertResponse) error
	//重新生成keystore文件
	Regenerate(context.Context, *OemAppAndroidCertFilter, *OemAppAndroidCertResponse) error
}

func RegisterOemAppAndroidCertServiceHandler(s server.Server, hdlr OemAppAndroidCertServiceHandler, opts ...server.HandlerOption) error {
	type oemAppAndroidCertService interface {
		Create(ctx context.Context, in *OemAppAndroidCert, out *Response) error
		Delete(ctx context.Context, in *OemAppAndroidCert, out *Response) error
		DeleteById(ctx context.Context, in *OemAppAndroidCert, out *Response) error
		DeleteByIds(ctx context.Context, in *OemAppAndroidCertBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *OemAppAndroidCert, out *Response) error
		UpdateAll(ctx context.Context, in *OemAppAndroidCert, out *Response) error
		UpdateFields(ctx context.Context, in *OemAppAndroidCertUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *OemAppAndroidCertFilter, out *OemAppAndroidCertResponse) error
		Find(ctx context.Context, in *OemAppAndroidCertFilter, out *OemAppAndroidCertResponse) error
		Lists(ctx context.Context, in *OemAppAndroidCertListRequest, out *OemAppAndroidCertResponse) error
		Regenerate(ctx context.Context, in *OemAppAndroidCertFilter, out *OemAppAndroidCertResponse) error
	}
	type OemAppAndroidCertService struct {
		oemAppAndroidCertService
	}
	h := &oemAppAndroidCertServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppAndroidCertService.Create",
		Path:    []string{"/v1/oemAppAndroidCert/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppAndroidCertService.Delete",
		Path:    []string{"/v1/oemAppAndroidCert/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppAndroidCertService.DeleteById",
		Path:    []string{"/v1/oemAppAndroidCert/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppAndroidCertService.DeleteByIds",
		Path:    []string{"/v1/oemAppAndroidCert/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppAndroidCertService.Update",
		Path:    []string{"/v1/oemAppAndroidCert/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppAndroidCertService.UpdateAll",
		Path:    []string{"/v1/oemAppAndroidCert/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppAndroidCertService.UpdateFields",
		Path:    []string{"/v1/oemAppAndroidCert/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppAndroidCertService.FindById",
		Path:    []string{"/v1/oemAppAndroidCert/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppAndroidCertService.Find",
		Path:    []string{"/v1/oemAppAndroidCert/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppAndroidCertService.Lists",
		Path:    []string{"/v1/oemAppAndroidCert/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppAndroidCertService.Regenerate",
		Path:    []string{"/v1/oemAppAndroidCert/regenerate"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&OemAppAndroidCertService{h}, opts...))
}

type oemAppAndroidCertServiceHandler struct {
	OemAppAndroidCertServiceHandler
}

func (h *oemAppAndroidCertServiceHandler) Create(ctx context.Context, in *OemAppAndroidCert, out *Response) error {
	return h.OemAppAndroidCertServiceHandler.Create(ctx, in, out)
}

func (h *oemAppAndroidCertServiceHandler) Delete(ctx context.Context, in *OemAppAndroidCert, out *Response) error {
	return h.OemAppAndroidCertServiceHandler.Delete(ctx, in, out)
}

func (h *oemAppAndroidCertServiceHandler) DeleteById(ctx context.Context, in *OemAppAndroidCert, out *Response) error {
	return h.OemAppAndroidCertServiceHandler.DeleteById(ctx, in, out)
}

func (h *oemAppAndroidCertServiceHandler) DeleteByIds(ctx context.Context, in *OemAppAndroidCertBatchDeleteRequest, out *Response) error {
	return h.OemAppAndroidCertServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *oemAppAndroidCertServiceHandler) Update(ctx context.Context, in *OemAppAndroidCert, out *Response) error {
	return h.OemAppAndroidCertServiceHandler.Update(ctx, in, out)
}

func (h *oemAppAndroidCertServiceHandler) UpdateAll(ctx context.Context, in *OemAppAndroidCert, out *Response) error {
	return h.OemAppAndroidCertServiceHandler.UpdateAll(ctx, in, out)
}

func (h *oemAppAndroidCertServiceHandler) UpdateFields(ctx context.Context, in *OemAppAndroidCertUpdateFieldsRequest, out *Response) error {
	return h.OemAppAndroidCertServiceHandler.UpdateFields(ctx, in, out)
}

func (h *oemAppAndroidCertServiceHandler) FindById(ctx context.Context, in *OemAppAndroidCertFilter, out *OemAppAndroidCertResponse) error {
	return h.OemAppAndroidCertServiceHandler.FindById(ctx, in, out)
}

func (h *oemAppAndroidCertServiceHandler) Find(ctx context.Context, in *OemAppAndroidCertFilter, out *OemAppAndroidCertResponse) error {
	return h.OemAppAndroidCertServiceHandler.Find(ctx, in, out)
}

func (h *oemAppAndroidCertServiceHandler) Lists(ctx context.Context, in *OemAppAndroidCertListRequest, out *OemAppAndroidCertResponse) error {
	return h.OemAppAndroidCertServiceHandler.Lists(ctx, in, out)
}

func (h *oemAppAndroidCertServiceHandler) Regenerate(ctx context.Context, in *OemAppAndroidCertFilter, out *OemAppAndroidCertResponse) error {
	return h.OemAppAndroidCertServiceHandler.Regenerate(ctx, in, out)
}
