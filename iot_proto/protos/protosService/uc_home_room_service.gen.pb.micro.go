// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: uc_home_room_service.gen.proto

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

// Api Endpoints for UcHomeRoomService service

func NewUcHomeRoomServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "UcHomeRoomService.Create",
			Path:    []string{"/v1/ucHomeRoom/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "UcHomeRoomService.Delete",
			Path:    []string{"/v1/ucHomeRoom/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "UcHomeRoomService.DeleteById",
			Path:    []string{"/v1/ucHomeRoom/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "UcHomeRoomService.DeleteByIds",
			Path:    []string{"/v1/ucHomeRoom/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "UcHomeRoomService.Update",
			Path:    []string{"/v1/ucHomeRoom/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "UcHomeRoomService.UpdateAll",
			Path:    []string{"/v1/ucHomeRoom/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "UcHomeRoomService.UpdateFields",
			Path:    []string{"/v1/ucHomeRoom/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "UcHomeRoomService.FindById",
			Path:    []string{"/v1/ucHomeRoom/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "UcHomeRoomService.Find",
			Path:    []string{"/v1/ucHomeRoom/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "UcHomeRoomService.Lists",
			Path:    []string{"/v1/ucHomeRoom/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "UcHomeRoomService.SetSort",
			Path:    []string{"/v1/ucHomeRoom/setSort"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "UcHomeRoomService.FindByIds",
			Path:    []string{"/v1/ucHomeRoom/FindByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for UcHomeRoomService service

type UcHomeRoomService interface {
	//创建
	Create(ctx context.Context, in *UcHomeRoom, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *UcHomeRoom, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *UcHomeRoom, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *UcHomeRoomBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *UcHomeRoom, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *UcHomeRoom, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *UcHomeRoomUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *UcHomeRoomFilter, opts ...client.CallOption) (*UcHomeRoomResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *UcHomeRoomFilter, opts ...client.CallOption) (*UcHomeRoomResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *UcHomeRoomListRequest, opts ...client.CallOption) (*UcHomeRoomResponse, error)
	//家庭房间排序
	SetSort(ctx context.Context, in *UcHomeRoomSortRequest, opts ...client.CallOption) (*UcHomeRoomResponse, error)
	//根据ids查找
	FindByIds(ctx context.Context, in *UcHomeRoomFilter, opts ...client.CallOption) (*UcHomeRoomResponse, error)
}

type ucHomeRoomService struct {
	c    client.Client
	name string
}

func NewUcHomeRoomService(name string, c client.Client) UcHomeRoomService {
	return &ucHomeRoomService{
		c:    c,
		name: name,
	}
}

func (c *ucHomeRoomService) Create(ctx context.Context, in *UcHomeRoom, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ucHomeRoomService) Delete(ctx context.Context, in *UcHomeRoom, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ucHomeRoomService) DeleteById(ctx context.Context, in *UcHomeRoom, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ucHomeRoomService) DeleteByIds(ctx context.Context, in *UcHomeRoomBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ucHomeRoomService) Update(ctx context.Context, in *UcHomeRoom, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ucHomeRoomService) UpdateAll(ctx context.Context, in *UcHomeRoom, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ucHomeRoomService) UpdateFields(ctx context.Context, in *UcHomeRoomUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ucHomeRoomService) FindById(ctx context.Context, in *UcHomeRoomFilter, opts ...client.CallOption) (*UcHomeRoomResponse, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.FindById", in)
	out := new(UcHomeRoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ucHomeRoomService) Find(ctx context.Context, in *UcHomeRoomFilter, opts ...client.CallOption) (*UcHomeRoomResponse, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.Find", in)
	out := new(UcHomeRoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ucHomeRoomService) Lists(ctx context.Context, in *UcHomeRoomListRequest, opts ...client.CallOption) (*UcHomeRoomResponse, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.Lists", in)
	out := new(UcHomeRoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ucHomeRoomService) SetSort(ctx context.Context, in *UcHomeRoomSortRequest, opts ...client.CallOption) (*UcHomeRoomResponse, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.SetSort", in)
	out := new(UcHomeRoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ucHomeRoomService) FindByIds(ctx context.Context, in *UcHomeRoomFilter, opts ...client.CallOption) (*UcHomeRoomResponse, error) {
	req := c.c.NewRequest(c.name, "UcHomeRoomService.FindByIds", in)
	out := new(UcHomeRoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UcHomeRoomService service

type UcHomeRoomServiceHandler interface {
	//创建
	Create(context.Context, *UcHomeRoom, *Response) error
	//匹配多条件删除
	Delete(context.Context, *UcHomeRoom, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *UcHomeRoom, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *UcHomeRoomBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *UcHomeRoom, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *UcHomeRoom, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *UcHomeRoomUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *UcHomeRoomFilter, *UcHomeRoomResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *UcHomeRoomFilter, *UcHomeRoomResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *UcHomeRoomListRequest, *UcHomeRoomResponse) error
	//家庭房间排序
	SetSort(context.Context, *UcHomeRoomSortRequest, *UcHomeRoomResponse) error
	//根据ids查找
	FindByIds(context.Context, *UcHomeRoomFilter, *UcHomeRoomResponse) error
}

func RegisterUcHomeRoomServiceHandler(s server.Server, hdlr UcHomeRoomServiceHandler, opts ...server.HandlerOption) error {
	type ucHomeRoomService interface {
		Create(ctx context.Context, in *UcHomeRoom, out *Response) error
		Delete(ctx context.Context, in *UcHomeRoom, out *Response) error
		DeleteById(ctx context.Context, in *UcHomeRoom, out *Response) error
		DeleteByIds(ctx context.Context, in *UcHomeRoomBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *UcHomeRoom, out *Response) error
		UpdateAll(ctx context.Context, in *UcHomeRoom, out *Response) error
		UpdateFields(ctx context.Context, in *UcHomeRoomUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *UcHomeRoomFilter, out *UcHomeRoomResponse) error
		Find(ctx context.Context, in *UcHomeRoomFilter, out *UcHomeRoomResponse) error
		Lists(ctx context.Context, in *UcHomeRoomListRequest, out *UcHomeRoomResponse) error
		SetSort(ctx context.Context, in *UcHomeRoomSortRequest, out *UcHomeRoomResponse) error
		FindByIds(ctx context.Context, in *UcHomeRoomFilter, out *UcHomeRoomResponse) error
	}
	type UcHomeRoomService struct {
		ucHomeRoomService
	}
	h := &ucHomeRoomServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.Create",
		Path:    []string{"/v1/ucHomeRoom/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.Delete",
		Path:    []string{"/v1/ucHomeRoom/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.DeleteById",
		Path:    []string{"/v1/ucHomeRoom/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.DeleteByIds",
		Path:    []string{"/v1/ucHomeRoom/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.Update",
		Path:    []string{"/v1/ucHomeRoom/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.UpdateAll",
		Path:    []string{"/v1/ucHomeRoom/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.UpdateFields",
		Path:    []string{"/v1/ucHomeRoom/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.FindById",
		Path:    []string{"/v1/ucHomeRoom/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.Find",
		Path:    []string{"/v1/ucHomeRoom/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.Lists",
		Path:    []string{"/v1/ucHomeRoom/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.SetSort",
		Path:    []string{"/v1/ucHomeRoom/setSort"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UcHomeRoomService.FindByIds",
		Path:    []string{"/v1/ucHomeRoom/FindByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&UcHomeRoomService{h}, opts...))
}

type ucHomeRoomServiceHandler struct {
	UcHomeRoomServiceHandler
}

func (h *ucHomeRoomServiceHandler) Create(ctx context.Context, in *UcHomeRoom, out *Response) error {
	return h.UcHomeRoomServiceHandler.Create(ctx, in, out)
}

func (h *ucHomeRoomServiceHandler) Delete(ctx context.Context, in *UcHomeRoom, out *Response) error {
	return h.UcHomeRoomServiceHandler.Delete(ctx, in, out)
}

func (h *ucHomeRoomServiceHandler) DeleteById(ctx context.Context, in *UcHomeRoom, out *Response) error {
	return h.UcHomeRoomServiceHandler.DeleteById(ctx, in, out)
}

func (h *ucHomeRoomServiceHandler) DeleteByIds(ctx context.Context, in *UcHomeRoomBatchDeleteRequest, out *Response) error {
	return h.UcHomeRoomServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *ucHomeRoomServiceHandler) Update(ctx context.Context, in *UcHomeRoom, out *Response) error {
	return h.UcHomeRoomServiceHandler.Update(ctx, in, out)
}

func (h *ucHomeRoomServiceHandler) UpdateAll(ctx context.Context, in *UcHomeRoom, out *Response) error {
	return h.UcHomeRoomServiceHandler.UpdateAll(ctx, in, out)
}

func (h *ucHomeRoomServiceHandler) UpdateFields(ctx context.Context, in *UcHomeRoomUpdateFieldsRequest, out *Response) error {
	return h.UcHomeRoomServiceHandler.UpdateFields(ctx, in, out)
}

func (h *ucHomeRoomServiceHandler) FindById(ctx context.Context, in *UcHomeRoomFilter, out *UcHomeRoomResponse) error {
	return h.UcHomeRoomServiceHandler.FindById(ctx, in, out)
}

func (h *ucHomeRoomServiceHandler) Find(ctx context.Context, in *UcHomeRoomFilter, out *UcHomeRoomResponse) error {
	return h.UcHomeRoomServiceHandler.Find(ctx, in, out)
}

func (h *ucHomeRoomServiceHandler) Lists(ctx context.Context, in *UcHomeRoomListRequest, out *UcHomeRoomResponse) error {
	return h.UcHomeRoomServiceHandler.Lists(ctx, in, out)
}

func (h *ucHomeRoomServiceHandler) SetSort(ctx context.Context, in *UcHomeRoomSortRequest, out *UcHomeRoomResponse) error {
	return h.UcHomeRoomServiceHandler.SetSort(ctx, in, out)
}

func (h *ucHomeRoomServiceHandler) FindByIds(ctx context.Context, in *UcHomeRoomFilter, out *UcHomeRoomResponse) error {
	return h.UcHomeRoomServiceHandler.FindByIds(ctx, in, out)
}
