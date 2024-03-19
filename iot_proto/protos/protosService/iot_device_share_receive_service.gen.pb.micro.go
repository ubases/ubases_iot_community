// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: iot_device_share_receive_service.gen.proto

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

// Api Endpoints for IotDeviceShareReceiveService service

func NewIotDeviceShareReceiveServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "IotDeviceShareReceiveService.Create",
			Path:    []string{"/v1/iotDeviceShareReceive/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotDeviceShareReceiveService.Delete",
			Path:    []string{"/v1/iotDeviceShareReceive/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotDeviceShareReceiveService.DeleteById",
			Path:    []string{"/v1/iotDeviceShareReceive/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotDeviceShareReceiveService.DeleteByIds",
			Path:    []string{"/v1/iotDeviceShareReceive/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotDeviceShareReceiveService.Update",
			Path:    []string{"/v1/iotDeviceShareReceive/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotDeviceShareReceiveService.UpdateAll",
			Path:    []string{"/v1/iotDeviceShareReceive/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotDeviceShareReceiveService.UpdateFields",
			Path:    []string{"/v1/iotDeviceShareReceive/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotDeviceShareReceiveService.FindById",
			Path:    []string{"/v1/iotDeviceShareReceive/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotDeviceShareReceiveService.Find",
			Path:    []string{"/v1/iotDeviceShareReceive/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotDeviceShareReceiveService.Lists",
			Path:    []string{"/v1/iotDeviceShareReceive/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for IotDeviceShareReceiveService service

type IotDeviceShareReceiveService interface {
	//创建
	Create(ctx context.Context, in *IotDeviceShareReceive, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *IotDeviceShareReceive, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *IotDeviceShareReceive, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *IotDeviceShareReceiveBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *IotDeviceShareReceive, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *IotDeviceShareReceive, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *IotDeviceShareReceiveUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *IotDeviceShareReceiveFilter, opts ...client.CallOption) (*IotDeviceShareReceiveResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *IotDeviceShareReceiveFilter, opts ...client.CallOption) (*IotDeviceShareReceiveResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *IotDeviceShareReceiveListRequest, opts ...client.CallOption) (*IotDeviceShareReceiveResponse, error)
}

type iotDeviceShareReceiveService struct {
	c    client.Client
	name string
}

func NewIotDeviceShareReceiveService(name string, c client.Client) IotDeviceShareReceiveService {
	return &iotDeviceShareReceiveService{
		c:    c,
		name: name,
	}
}

func (c *iotDeviceShareReceiveService) Create(ctx context.Context, in *IotDeviceShareReceive, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotDeviceShareReceiveService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotDeviceShareReceiveService) Delete(ctx context.Context, in *IotDeviceShareReceive, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotDeviceShareReceiveService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotDeviceShareReceiveService) DeleteById(ctx context.Context, in *IotDeviceShareReceive, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotDeviceShareReceiveService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotDeviceShareReceiveService) DeleteByIds(ctx context.Context, in *IotDeviceShareReceiveBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotDeviceShareReceiveService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotDeviceShareReceiveService) Update(ctx context.Context, in *IotDeviceShareReceive, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotDeviceShareReceiveService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotDeviceShareReceiveService) UpdateAll(ctx context.Context, in *IotDeviceShareReceive, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotDeviceShareReceiveService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotDeviceShareReceiveService) UpdateFields(ctx context.Context, in *IotDeviceShareReceiveUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotDeviceShareReceiveService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotDeviceShareReceiveService) FindById(ctx context.Context, in *IotDeviceShareReceiveFilter, opts ...client.CallOption) (*IotDeviceShareReceiveResponse, error) {
	req := c.c.NewRequest(c.name, "IotDeviceShareReceiveService.FindById", in)
	out := new(IotDeviceShareReceiveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotDeviceShareReceiveService) Find(ctx context.Context, in *IotDeviceShareReceiveFilter, opts ...client.CallOption) (*IotDeviceShareReceiveResponse, error) {
	req := c.c.NewRequest(c.name, "IotDeviceShareReceiveService.Find", in)
	out := new(IotDeviceShareReceiveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotDeviceShareReceiveService) Lists(ctx context.Context, in *IotDeviceShareReceiveListRequest, opts ...client.CallOption) (*IotDeviceShareReceiveResponse, error) {
	req := c.c.NewRequest(c.name, "IotDeviceShareReceiveService.Lists", in)
	out := new(IotDeviceShareReceiveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IotDeviceShareReceiveService service

type IotDeviceShareReceiveServiceHandler interface {
	//创建
	Create(context.Context, *IotDeviceShareReceive, *Response) error
	//匹配多条件删除
	Delete(context.Context, *IotDeviceShareReceive, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *IotDeviceShareReceive, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *IotDeviceShareReceiveBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *IotDeviceShareReceive, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *IotDeviceShareReceive, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *IotDeviceShareReceiveUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *IotDeviceShareReceiveFilter, *IotDeviceShareReceiveResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *IotDeviceShareReceiveFilter, *IotDeviceShareReceiveResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *IotDeviceShareReceiveListRequest, *IotDeviceShareReceiveResponse) error
}

func RegisterIotDeviceShareReceiveServiceHandler(s server.Server, hdlr IotDeviceShareReceiveServiceHandler, opts ...server.HandlerOption) error {
	type iotDeviceShareReceiveService interface {
		Create(ctx context.Context, in *IotDeviceShareReceive, out *Response) error
		Delete(ctx context.Context, in *IotDeviceShareReceive, out *Response) error
		DeleteById(ctx context.Context, in *IotDeviceShareReceive, out *Response) error
		DeleteByIds(ctx context.Context, in *IotDeviceShareReceiveBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *IotDeviceShareReceive, out *Response) error
		UpdateAll(ctx context.Context, in *IotDeviceShareReceive, out *Response) error
		UpdateFields(ctx context.Context, in *IotDeviceShareReceiveUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *IotDeviceShareReceiveFilter, out *IotDeviceShareReceiveResponse) error
		Find(ctx context.Context, in *IotDeviceShareReceiveFilter, out *IotDeviceShareReceiveResponse) error
		Lists(ctx context.Context, in *IotDeviceShareReceiveListRequest, out *IotDeviceShareReceiveResponse) error
	}
	type IotDeviceShareReceiveService struct {
		iotDeviceShareReceiveService
	}
	h := &iotDeviceShareReceiveServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotDeviceShareReceiveService.Create",
		Path:    []string{"/v1/iotDeviceShareReceive/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotDeviceShareReceiveService.Delete",
		Path:    []string{"/v1/iotDeviceShareReceive/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotDeviceShareReceiveService.DeleteById",
		Path:    []string{"/v1/iotDeviceShareReceive/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotDeviceShareReceiveService.DeleteByIds",
		Path:    []string{"/v1/iotDeviceShareReceive/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotDeviceShareReceiveService.Update",
		Path:    []string{"/v1/iotDeviceShareReceive/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotDeviceShareReceiveService.UpdateAll",
		Path:    []string{"/v1/iotDeviceShareReceive/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotDeviceShareReceiveService.UpdateFields",
		Path:    []string{"/v1/iotDeviceShareReceive/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotDeviceShareReceiveService.FindById",
		Path:    []string{"/v1/iotDeviceShareReceive/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotDeviceShareReceiveService.Find",
		Path:    []string{"/v1/iotDeviceShareReceive/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotDeviceShareReceiveService.Lists",
		Path:    []string{"/v1/iotDeviceShareReceive/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&IotDeviceShareReceiveService{h}, opts...))
}

type iotDeviceShareReceiveServiceHandler struct {
	IotDeviceShareReceiveServiceHandler
}

func (h *iotDeviceShareReceiveServiceHandler) Create(ctx context.Context, in *IotDeviceShareReceive, out *Response) error {
	return h.IotDeviceShareReceiveServiceHandler.Create(ctx, in, out)
}

func (h *iotDeviceShareReceiveServiceHandler) Delete(ctx context.Context, in *IotDeviceShareReceive, out *Response) error {
	return h.IotDeviceShareReceiveServiceHandler.Delete(ctx, in, out)
}

func (h *iotDeviceShareReceiveServiceHandler) DeleteById(ctx context.Context, in *IotDeviceShareReceive, out *Response) error {
	return h.IotDeviceShareReceiveServiceHandler.DeleteById(ctx, in, out)
}

func (h *iotDeviceShareReceiveServiceHandler) DeleteByIds(ctx context.Context, in *IotDeviceShareReceiveBatchDeleteRequest, out *Response) error {
	return h.IotDeviceShareReceiveServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *iotDeviceShareReceiveServiceHandler) Update(ctx context.Context, in *IotDeviceShareReceive, out *Response) error {
	return h.IotDeviceShareReceiveServiceHandler.Update(ctx, in, out)
}

func (h *iotDeviceShareReceiveServiceHandler) UpdateAll(ctx context.Context, in *IotDeviceShareReceive, out *Response) error {
	return h.IotDeviceShareReceiveServiceHandler.UpdateAll(ctx, in, out)
}

func (h *iotDeviceShareReceiveServiceHandler) UpdateFields(ctx context.Context, in *IotDeviceShareReceiveUpdateFieldsRequest, out *Response) error {
	return h.IotDeviceShareReceiveServiceHandler.UpdateFields(ctx, in, out)
}

func (h *iotDeviceShareReceiveServiceHandler) FindById(ctx context.Context, in *IotDeviceShareReceiveFilter, out *IotDeviceShareReceiveResponse) error {
	return h.IotDeviceShareReceiveServiceHandler.FindById(ctx, in, out)
}

func (h *iotDeviceShareReceiveServiceHandler) Find(ctx context.Context, in *IotDeviceShareReceiveFilter, out *IotDeviceShareReceiveResponse) error {
	return h.IotDeviceShareReceiveServiceHandler.Find(ctx, in, out)
}

func (h *iotDeviceShareReceiveServiceHandler) Lists(ctx context.Context, in *IotDeviceShareReceiveListRequest, out *IotDeviceShareReceiveResponse) error {
	return h.IotDeviceShareReceiveServiceHandler.Lists(ctx, in, out)
}
