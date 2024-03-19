// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: device_active_hour_service.proto

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

// Api Endpoints for DeviceActiveHourService service

func NewDeviceActiveHourServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "DeviceActiveHourService.Create",
			Path:    []string{"/v1/deviceActiveHour/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DeviceActiveHourService.Delete",
			Path:    []string{"/v1/deviceActiveHour/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DeviceActiveHourService.DeleteById",
			Path:    []string{"/v1/deviceActiveHour/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DeviceActiveHourService.DeleteByIds",
			Path:    []string{"/v1/deviceActiveHour/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DeviceActiveHourService.Update",
			Path:    []string{"/v1/deviceActiveHour/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DeviceActiveHourService.UpdateAll",
			Path:    []string{"/v1/deviceActiveHour/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DeviceActiveHourService.UpdateFields",
			Path:    []string{"/v1/deviceActiveHour/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DeviceActiveHourService.FindById",
			Path:    []string{"/v1/deviceActiveHour/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DeviceActiveHourService.Find",
			Path:    []string{"/v1/deviceActiveHour/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "DeviceActiveHourService.Lists",
			Path:    []string{"/v1/deviceActiveHour/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for DeviceActiveHourService service

type DeviceActiveHourService interface {
	//创建
	Create(ctx context.Context, in *DeviceActiveHour, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *DeviceActiveHour, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *DeviceActiveHour, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *DeviceActiveHourBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *DeviceActiveHour, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *DeviceActiveHour, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *DeviceActiveHourUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *DeviceActiveHourFilter, opts ...client.CallOption) (*DeviceActiveHourResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *DeviceActiveHourFilter, opts ...client.CallOption) (*DeviceActiveHourResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *DeviceActiveHourListRequest, opts ...client.CallOption) (*DeviceActiveHourResponse, error)
}

type deviceActiveHourService struct {
	c    client.Client
	name string
}

func NewDeviceActiveHourService(name string, c client.Client) DeviceActiveHourService {
	return &deviceActiveHourService{
		c:    c,
		name: name,
	}
}

func (c *deviceActiveHourService) Create(ctx context.Context, in *DeviceActiveHour, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeviceActiveHourService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceActiveHourService) Delete(ctx context.Context, in *DeviceActiveHour, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeviceActiveHourService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceActiveHourService) DeleteById(ctx context.Context, in *DeviceActiveHour, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeviceActiveHourService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceActiveHourService) DeleteByIds(ctx context.Context, in *DeviceActiveHourBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeviceActiveHourService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceActiveHourService) Update(ctx context.Context, in *DeviceActiveHour, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeviceActiveHourService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceActiveHourService) UpdateAll(ctx context.Context, in *DeviceActiveHour, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeviceActiveHourService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceActiveHourService) UpdateFields(ctx context.Context, in *DeviceActiveHourUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "DeviceActiveHourService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceActiveHourService) FindById(ctx context.Context, in *DeviceActiveHourFilter, opts ...client.CallOption) (*DeviceActiveHourResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceActiveHourService.FindById", in)
	out := new(DeviceActiveHourResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceActiveHourService) Find(ctx context.Context, in *DeviceActiveHourFilter, opts ...client.CallOption) (*DeviceActiveHourResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceActiveHourService.Find", in)
	out := new(DeviceActiveHourResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceActiveHourService) Lists(ctx context.Context, in *DeviceActiveHourListRequest, opts ...client.CallOption) (*DeviceActiveHourResponse, error) {
	req := c.c.NewRequest(c.name, "DeviceActiveHourService.Lists", in)
	out := new(DeviceActiveHourResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceActiveHourService service

type DeviceActiveHourServiceHandler interface {
	//创建
	Create(context.Context, *DeviceActiveHour, *Response) error
	//匹配多条件删除
	Delete(context.Context, *DeviceActiveHour, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *DeviceActiveHour, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *DeviceActiveHourBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *DeviceActiveHour, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *DeviceActiveHour, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *DeviceActiveHourUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *DeviceActiveHourFilter, *DeviceActiveHourResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *DeviceActiveHourFilter, *DeviceActiveHourResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *DeviceActiveHourListRequest, *DeviceActiveHourResponse) error
}

func RegisterDeviceActiveHourServiceHandler(s server.Server, hdlr DeviceActiveHourServiceHandler, opts ...server.HandlerOption) error {
	type deviceActiveHourService interface {
		Create(ctx context.Context, in *DeviceActiveHour, out *Response) error
		Delete(ctx context.Context, in *DeviceActiveHour, out *Response) error
		DeleteById(ctx context.Context, in *DeviceActiveHour, out *Response) error
		DeleteByIds(ctx context.Context, in *DeviceActiveHourBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *DeviceActiveHour, out *Response) error
		UpdateAll(ctx context.Context, in *DeviceActiveHour, out *Response) error
		UpdateFields(ctx context.Context, in *DeviceActiveHourUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *DeviceActiveHourFilter, out *DeviceActiveHourResponse) error
		Find(ctx context.Context, in *DeviceActiveHourFilter, out *DeviceActiveHourResponse) error
		Lists(ctx context.Context, in *DeviceActiveHourListRequest, out *DeviceActiveHourResponse) error
	}
	type DeviceActiveHourService struct {
		deviceActiveHourService
	}
	h := &deviceActiveHourServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DeviceActiveHourService.Create",
		Path:    []string{"/v1/deviceActiveHour/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DeviceActiveHourService.Delete",
		Path:    []string{"/v1/deviceActiveHour/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DeviceActiveHourService.DeleteById",
		Path:    []string{"/v1/deviceActiveHour/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DeviceActiveHourService.DeleteByIds",
		Path:    []string{"/v1/deviceActiveHour/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DeviceActiveHourService.Update",
		Path:    []string{"/v1/deviceActiveHour/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DeviceActiveHourService.UpdateAll",
		Path:    []string{"/v1/deviceActiveHour/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DeviceActiveHourService.UpdateFields",
		Path:    []string{"/v1/deviceActiveHour/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DeviceActiveHourService.FindById",
		Path:    []string{"/v1/deviceActiveHour/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DeviceActiveHourService.Find",
		Path:    []string{"/v1/deviceActiveHour/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DeviceActiveHourService.Lists",
		Path:    []string{"/v1/deviceActiveHour/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&DeviceActiveHourService{h}, opts...))
}

type deviceActiveHourServiceHandler struct {
	DeviceActiveHourServiceHandler
}

func (h *deviceActiveHourServiceHandler) Create(ctx context.Context, in *DeviceActiveHour, out *Response) error {
	return h.DeviceActiveHourServiceHandler.Create(ctx, in, out)
}

func (h *deviceActiveHourServiceHandler) Delete(ctx context.Context, in *DeviceActiveHour, out *Response) error {
	return h.DeviceActiveHourServiceHandler.Delete(ctx, in, out)
}

func (h *deviceActiveHourServiceHandler) DeleteById(ctx context.Context, in *DeviceActiveHour, out *Response) error {
	return h.DeviceActiveHourServiceHandler.DeleteById(ctx, in, out)
}

func (h *deviceActiveHourServiceHandler) DeleteByIds(ctx context.Context, in *DeviceActiveHourBatchDeleteRequest, out *Response) error {
	return h.DeviceActiveHourServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *deviceActiveHourServiceHandler) Update(ctx context.Context, in *DeviceActiveHour, out *Response) error {
	return h.DeviceActiveHourServiceHandler.Update(ctx, in, out)
}

func (h *deviceActiveHourServiceHandler) UpdateAll(ctx context.Context, in *DeviceActiveHour, out *Response) error {
	return h.DeviceActiveHourServiceHandler.UpdateAll(ctx, in, out)
}

func (h *deviceActiveHourServiceHandler) UpdateFields(ctx context.Context, in *DeviceActiveHourUpdateFieldsRequest, out *Response) error {
	return h.DeviceActiveHourServiceHandler.UpdateFields(ctx, in, out)
}

func (h *deviceActiveHourServiceHandler) FindById(ctx context.Context, in *DeviceActiveHourFilter, out *DeviceActiveHourResponse) error {
	return h.DeviceActiveHourServiceHandler.FindById(ctx, in, out)
}

func (h *deviceActiveHourServiceHandler) Find(ctx context.Context, in *DeviceActiveHourFilter, out *DeviceActiveHourResponse) error {
	return h.DeviceActiveHourServiceHandler.Find(ctx, in, out)
}

func (h *deviceActiveHourServiceHandler) Lists(ctx context.Context, in *DeviceActiveHourListRequest, out *DeviceActiveHourResponse) error {
	return h.DeviceActiveHourServiceHandler.Lists(ctx, in, out)
}
