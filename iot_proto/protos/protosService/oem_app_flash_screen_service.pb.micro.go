// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: oem_app_flash_screen_service.proto

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

// Api Endpoints for OemAppFlashScreenService service

func NewOemAppFlashScreenServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "OemAppFlashScreenService.Create",
			Path:    []string{"/v1/oemAppFlashScreen/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppFlashScreenService.Delete",
			Path:    []string{"/v1/oemAppFlashScreen/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppFlashScreenService.DeleteById",
			Path:    []string{"/v1/oemAppFlashScreen/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppFlashScreenService.DeleteByIds",
			Path:    []string{"/v1/oemAppFlashScreen/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppFlashScreenService.Update",
			Path:    []string{"/v1/oemAppFlashScreen/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppFlashScreenService.UpdateAll",
			Path:    []string{"/v1/oemAppFlashScreen/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppFlashScreenService.UpdateFields",
			Path:    []string{"/v1/oemAppFlashScreen/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppFlashScreenService.FindById",
			Path:    []string{"/v1/oemAppFlashScreen/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppFlashScreenService.Find",
			Path:    []string{"/v1/oemAppFlashScreen/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppFlashScreenService.Lists",
			Path:    []string{"/v1/oemAppFlashScreen/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "OemAppFlashScreenService.GetFlashScreen",
			Path:    []string{"/v1/oemAppFlashScreen/getFlashScreen"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for OemAppFlashScreenService service

type OemAppFlashScreenService interface {
	//创建
	Create(ctx context.Context, in *OemAppFlashScreen, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *OemAppFlashScreen, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *OemAppFlashScreen, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *OemAppFlashScreenBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *OemAppFlashScreen, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *OemAppFlashScreen, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *OemAppFlashScreenUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *OemAppFlashScreenFilter, opts ...client.CallOption) (*OemAppFlashScreenResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *OemAppFlashScreenFilter, opts ...client.CallOption) (*OemAppFlashScreenResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *OemAppFlashScreenListRequest, opts ...client.CallOption) (*OemAppFlashScreenResponse, error)
	//多条件查找，返回单条数据
	GetFlashScreen(ctx context.Context, in *OemAppFlashScreenRequest, opts ...client.CallOption) (*OemAppFlashScreenResponse, error)
}

type oemAppFlashScreenService struct {
	c    client.Client
	name string
}

func NewOemAppFlashScreenService(name string, c client.Client) OemAppFlashScreenService {
	return &oemAppFlashScreenService{
		c:    c,
		name: name,
	}
}

func (c *oemAppFlashScreenService) Create(ctx context.Context, in *OemAppFlashScreen, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppFlashScreenService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppFlashScreenService) Delete(ctx context.Context, in *OemAppFlashScreen, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppFlashScreenService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppFlashScreenService) DeleteById(ctx context.Context, in *OemAppFlashScreen, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppFlashScreenService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppFlashScreenService) DeleteByIds(ctx context.Context, in *OemAppFlashScreenBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppFlashScreenService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppFlashScreenService) Update(ctx context.Context, in *OemAppFlashScreen, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppFlashScreenService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppFlashScreenService) UpdateAll(ctx context.Context, in *OemAppFlashScreen, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppFlashScreenService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppFlashScreenService) UpdateFields(ctx context.Context, in *OemAppFlashScreenUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OemAppFlashScreenService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppFlashScreenService) FindById(ctx context.Context, in *OemAppFlashScreenFilter, opts ...client.CallOption) (*OemAppFlashScreenResponse, error) {
	req := c.c.NewRequest(c.name, "OemAppFlashScreenService.FindById", in)
	out := new(OemAppFlashScreenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppFlashScreenService) Find(ctx context.Context, in *OemAppFlashScreenFilter, opts ...client.CallOption) (*OemAppFlashScreenResponse, error) {
	req := c.c.NewRequest(c.name, "OemAppFlashScreenService.Find", in)
	out := new(OemAppFlashScreenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppFlashScreenService) Lists(ctx context.Context, in *OemAppFlashScreenListRequest, opts ...client.CallOption) (*OemAppFlashScreenResponse, error) {
	req := c.c.NewRequest(c.name, "OemAppFlashScreenService.Lists", in)
	out := new(OemAppFlashScreenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oemAppFlashScreenService) GetFlashScreen(ctx context.Context, in *OemAppFlashScreenRequest, opts ...client.CallOption) (*OemAppFlashScreenResponse, error) {
	req := c.c.NewRequest(c.name, "OemAppFlashScreenService.GetFlashScreen", in)
	out := new(OemAppFlashScreenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OemAppFlashScreenService service

type OemAppFlashScreenServiceHandler interface {
	//创建
	Create(context.Context, *OemAppFlashScreen, *Response) error
	//匹配多条件删除
	Delete(context.Context, *OemAppFlashScreen, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *OemAppFlashScreen, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *OemAppFlashScreenBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *OemAppFlashScreen, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *OemAppFlashScreen, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *OemAppFlashScreenUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *OemAppFlashScreenFilter, *OemAppFlashScreenResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *OemAppFlashScreenFilter, *OemAppFlashScreenResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *OemAppFlashScreenListRequest, *OemAppFlashScreenResponse) error
	//多条件查找，返回单条数据
	GetFlashScreen(context.Context, *OemAppFlashScreenRequest, *OemAppFlashScreenResponse) error
}

func RegisterOemAppFlashScreenServiceHandler(s server.Server, hdlr OemAppFlashScreenServiceHandler, opts ...server.HandlerOption) error {
	type oemAppFlashScreenService interface {
		Create(ctx context.Context, in *OemAppFlashScreen, out *Response) error
		Delete(ctx context.Context, in *OemAppFlashScreen, out *Response) error
		DeleteById(ctx context.Context, in *OemAppFlashScreen, out *Response) error
		DeleteByIds(ctx context.Context, in *OemAppFlashScreenBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *OemAppFlashScreen, out *Response) error
		UpdateAll(ctx context.Context, in *OemAppFlashScreen, out *Response) error
		UpdateFields(ctx context.Context, in *OemAppFlashScreenUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *OemAppFlashScreenFilter, out *OemAppFlashScreenResponse) error
		Find(ctx context.Context, in *OemAppFlashScreenFilter, out *OemAppFlashScreenResponse) error
		Lists(ctx context.Context, in *OemAppFlashScreenListRequest, out *OemAppFlashScreenResponse) error
		GetFlashScreen(ctx context.Context, in *OemAppFlashScreenRequest, out *OemAppFlashScreenResponse) error
	}
	type OemAppFlashScreenService struct {
		oemAppFlashScreenService
	}
	h := &oemAppFlashScreenServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppFlashScreenService.Create",
		Path:    []string{"/v1/oemAppFlashScreen/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppFlashScreenService.Delete",
		Path:    []string{"/v1/oemAppFlashScreen/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppFlashScreenService.DeleteById",
		Path:    []string{"/v1/oemAppFlashScreen/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppFlashScreenService.DeleteByIds",
		Path:    []string{"/v1/oemAppFlashScreen/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppFlashScreenService.Update",
		Path:    []string{"/v1/oemAppFlashScreen/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppFlashScreenService.UpdateAll",
		Path:    []string{"/v1/oemAppFlashScreen/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppFlashScreenService.UpdateFields",
		Path:    []string{"/v1/oemAppFlashScreen/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppFlashScreenService.FindById",
		Path:    []string{"/v1/oemAppFlashScreen/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppFlashScreenService.Find",
		Path:    []string{"/v1/oemAppFlashScreen/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppFlashScreenService.Lists",
		Path:    []string{"/v1/oemAppFlashScreen/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "OemAppFlashScreenService.GetFlashScreen",
		Path:    []string{"/v1/oemAppFlashScreen/getFlashScreen"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&OemAppFlashScreenService{h}, opts...))
}

type oemAppFlashScreenServiceHandler struct {
	OemAppFlashScreenServiceHandler
}

func (h *oemAppFlashScreenServiceHandler) Create(ctx context.Context, in *OemAppFlashScreen, out *Response) error {
	return h.OemAppFlashScreenServiceHandler.Create(ctx, in, out)
}

func (h *oemAppFlashScreenServiceHandler) Delete(ctx context.Context, in *OemAppFlashScreen, out *Response) error {
	return h.OemAppFlashScreenServiceHandler.Delete(ctx, in, out)
}

func (h *oemAppFlashScreenServiceHandler) DeleteById(ctx context.Context, in *OemAppFlashScreen, out *Response) error {
	return h.OemAppFlashScreenServiceHandler.DeleteById(ctx, in, out)
}

func (h *oemAppFlashScreenServiceHandler) DeleteByIds(ctx context.Context, in *OemAppFlashScreenBatchDeleteRequest, out *Response) error {
	return h.OemAppFlashScreenServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *oemAppFlashScreenServiceHandler) Update(ctx context.Context, in *OemAppFlashScreen, out *Response) error {
	return h.OemAppFlashScreenServiceHandler.Update(ctx, in, out)
}

func (h *oemAppFlashScreenServiceHandler) UpdateAll(ctx context.Context, in *OemAppFlashScreen, out *Response) error {
	return h.OemAppFlashScreenServiceHandler.UpdateAll(ctx, in, out)
}

func (h *oemAppFlashScreenServiceHandler) UpdateFields(ctx context.Context, in *OemAppFlashScreenUpdateFieldsRequest, out *Response) error {
	return h.OemAppFlashScreenServiceHandler.UpdateFields(ctx, in, out)
}

func (h *oemAppFlashScreenServiceHandler) FindById(ctx context.Context, in *OemAppFlashScreenFilter, out *OemAppFlashScreenResponse) error {
	return h.OemAppFlashScreenServiceHandler.FindById(ctx, in, out)
}

func (h *oemAppFlashScreenServiceHandler) Find(ctx context.Context, in *OemAppFlashScreenFilter, out *OemAppFlashScreenResponse) error {
	return h.OemAppFlashScreenServiceHandler.Find(ctx, in, out)
}

func (h *oemAppFlashScreenServiceHandler) Lists(ctx context.Context, in *OemAppFlashScreenListRequest, out *OemAppFlashScreenResponse) error {
	return h.OemAppFlashScreenServiceHandler.Lists(ctx, in, out)
}

func (h *oemAppFlashScreenServiceHandler) GetFlashScreen(ctx context.Context, in *OemAppFlashScreenRequest, out *OemAppFlashScreenResponse) error {
	return h.OemAppFlashScreenServiceHandler.GetFlashScreen(ctx, in, out)
}
