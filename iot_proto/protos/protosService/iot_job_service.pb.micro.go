// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: iot_job_service.proto

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

// Api Endpoints for IotJobService service

func NewIotJobServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "IotJobService.Create",
			Path:    []string{"/v1/iotJob/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotJobService.Delete",
			Path:    []string{"/v1/iotJob/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotJobService.DeleteById",
			Path:    []string{"/v1/iotJob/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotJobService.DeleteByIds",
			Path:    []string{"/v1/iotJob/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotJobService.Update",
			Path:    []string{"/v1/iotJob/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotJobService.UpdateAll",
			Path:    []string{"/v1/iotJob/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotJobService.UpdateFields",
			Path:    []string{"/v1/iotJob/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotJobService.FindById",
			Path:    []string{"/v1/iotJob/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotJobService.Find",
			Path:    []string{"/v1/iotJob/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotJobService.Lists",
			Path:    []string{"/v1/iotJob/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotJobService.StartJob",
			Path:    []string{"/v1/iotJob/startJob"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "IotJobService.StopJob",
			Path:    []string{"/v1/iotJob/stopJob"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for IotJobService service

type IotJobService interface {
	//创建
	Create(ctx context.Context, in *IotJob, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *IotJob, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *IotJob, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *IotJobBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *IotJob, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *IotJob, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *IotJobUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *IotJobFilter, opts ...client.CallOption) (*IotJobResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *IotJobFilter, opts ...client.CallOption) (*IotJobResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *IotJobListRequest, opts ...client.CallOption) (*IotJobResponse, error)
	// 开始job任务
	StartJob(ctx context.Context, in *JobReq, opts ...client.CallOption) (*IotJobResponse, error)
	// 关闭job任务
	StopJob(ctx context.Context, in *JobReq, opts ...client.CallOption) (*IotJobResponse, error)
}

type iotJobService struct {
	c    client.Client
	name string
}

func NewIotJobService(name string, c client.Client) IotJobService {
	return &iotJobService{
		c:    c,
		name: name,
	}
}

func (c *iotJobService) Create(ctx context.Context, in *IotJob, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotJobService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotJobService) Delete(ctx context.Context, in *IotJob, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotJobService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotJobService) DeleteById(ctx context.Context, in *IotJob, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotJobService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotJobService) DeleteByIds(ctx context.Context, in *IotJobBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotJobService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotJobService) Update(ctx context.Context, in *IotJob, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotJobService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotJobService) UpdateAll(ctx context.Context, in *IotJob, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotJobService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotJobService) UpdateFields(ctx context.Context, in *IotJobUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "IotJobService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotJobService) FindById(ctx context.Context, in *IotJobFilter, opts ...client.CallOption) (*IotJobResponse, error) {
	req := c.c.NewRequest(c.name, "IotJobService.FindById", in)
	out := new(IotJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotJobService) Find(ctx context.Context, in *IotJobFilter, opts ...client.CallOption) (*IotJobResponse, error) {
	req := c.c.NewRequest(c.name, "IotJobService.Find", in)
	out := new(IotJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotJobService) Lists(ctx context.Context, in *IotJobListRequest, opts ...client.CallOption) (*IotJobResponse, error) {
	req := c.c.NewRequest(c.name, "IotJobService.Lists", in)
	out := new(IotJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotJobService) StartJob(ctx context.Context, in *JobReq, opts ...client.CallOption) (*IotJobResponse, error) {
	req := c.c.NewRequest(c.name, "IotJobService.StartJob", in)
	out := new(IotJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotJobService) StopJob(ctx context.Context, in *JobReq, opts ...client.CallOption) (*IotJobResponse, error) {
	req := c.c.NewRequest(c.name, "IotJobService.StopJob", in)
	out := new(IotJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IotJobService service

type IotJobServiceHandler interface {
	//创建
	Create(context.Context, *IotJob, *Response) error
	//匹配多条件删除
	Delete(context.Context, *IotJob, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *IotJob, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *IotJobBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *IotJob, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *IotJob, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *IotJobUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *IotJobFilter, *IotJobResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *IotJobFilter, *IotJobResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *IotJobListRequest, *IotJobResponse) error
	// 开始job任务
	StartJob(context.Context, *JobReq, *IotJobResponse) error
	// 关闭job任务
	StopJob(context.Context, *JobReq, *IotJobResponse) error
}

func RegisterIotJobServiceHandler(s server.Server, hdlr IotJobServiceHandler, opts ...server.HandlerOption) error {
	type iotJobService interface {
		Create(ctx context.Context, in *IotJob, out *Response) error
		Delete(ctx context.Context, in *IotJob, out *Response) error
		DeleteById(ctx context.Context, in *IotJob, out *Response) error
		DeleteByIds(ctx context.Context, in *IotJobBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *IotJob, out *Response) error
		UpdateAll(ctx context.Context, in *IotJob, out *Response) error
		UpdateFields(ctx context.Context, in *IotJobUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *IotJobFilter, out *IotJobResponse) error
		Find(ctx context.Context, in *IotJobFilter, out *IotJobResponse) error
		Lists(ctx context.Context, in *IotJobListRequest, out *IotJobResponse) error
		StartJob(ctx context.Context, in *JobReq, out *IotJobResponse) error
		StopJob(ctx context.Context, in *JobReq, out *IotJobResponse) error
	}
	type IotJobService struct {
		iotJobService
	}
	h := &iotJobServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.Create",
		Path:    []string{"/v1/iotJob/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.Delete",
		Path:    []string{"/v1/iotJob/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.DeleteById",
		Path:    []string{"/v1/iotJob/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.DeleteByIds",
		Path:    []string{"/v1/iotJob/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.Update",
		Path:    []string{"/v1/iotJob/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.UpdateAll",
		Path:    []string{"/v1/iotJob/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.UpdateFields",
		Path:    []string{"/v1/iotJob/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.FindById",
		Path:    []string{"/v1/iotJob/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.Find",
		Path:    []string{"/v1/iotJob/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.Lists",
		Path:    []string{"/v1/iotJob/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.StartJob",
		Path:    []string{"/v1/iotJob/startJob"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IotJobService.StopJob",
		Path:    []string{"/v1/iotJob/stopJob"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&IotJobService{h}, opts...))
}

type iotJobServiceHandler struct {
	IotJobServiceHandler
}

func (h *iotJobServiceHandler) Create(ctx context.Context, in *IotJob, out *Response) error {
	return h.IotJobServiceHandler.Create(ctx, in, out)
}

func (h *iotJobServiceHandler) Delete(ctx context.Context, in *IotJob, out *Response) error {
	return h.IotJobServiceHandler.Delete(ctx, in, out)
}

func (h *iotJobServiceHandler) DeleteById(ctx context.Context, in *IotJob, out *Response) error {
	return h.IotJobServiceHandler.DeleteById(ctx, in, out)
}

func (h *iotJobServiceHandler) DeleteByIds(ctx context.Context, in *IotJobBatchDeleteRequest, out *Response) error {
	return h.IotJobServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *iotJobServiceHandler) Update(ctx context.Context, in *IotJob, out *Response) error {
	return h.IotJobServiceHandler.Update(ctx, in, out)
}

func (h *iotJobServiceHandler) UpdateAll(ctx context.Context, in *IotJob, out *Response) error {
	return h.IotJobServiceHandler.UpdateAll(ctx, in, out)
}

func (h *iotJobServiceHandler) UpdateFields(ctx context.Context, in *IotJobUpdateFieldsRequest, out *Response) error {
	return h.IotJobServiceHandler.UpdateFields(ctx, in, out)
}

func (h *iotJobServiceHandler) FindById(ctx context.Context, in *IotJobFilter, out *IotJobResponse) error {
	return h.IotJobServiceHandler.FindById(ctx, in, out)
}

func (h *iotJobServiceHandler) Find(ctx context.Context, in *IotJobFilter, out *IotJobResponse) error {
	return h.IotJobServiceHandler.Find(ctx, in, out)
}

func (h *iotJobServiceHandler) Lists(ctx context.Context, in *IotJobListRequest, out *IotJobResponse) error {
	return h.IotJobServiceHandler.Lists(ctx, in, out)
}

func (h *iotJobServiceHandler) StartJob(ctx context.Context, in *JobReq, out *IotJobResponse) error {
	return h.IotJobServiceHandler.StartJob(ctx, in, out)
}

func (h *iotJobServiceHandler) StopJob(ctx context.Context, in *JobReq, out *IotJobResponse) error {
	return h.IotJobServiceHandler.StopJob(ctx, in, out)
}
