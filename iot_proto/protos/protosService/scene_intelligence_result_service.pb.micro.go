// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: scene_intelligence_result_service.proto

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

// Api Endpoints for SceneIntelligenceResultService service

func NewSceneIntelligenceResultServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "SceneIntelligenceResultService.Create",
			Path:    []string{"/v1/sceneIntelligenceResult/create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SceneIntelligenceResultService.Delete",
			Path:    []string{"/v1/sceneIntelligenceResult/delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SceneIntelligenceResultService.DeleteById",
			Path:    []string{"/v1/sceneIntelligenceResult/deleteById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SceneIntelligenceResultService.DeleteByIds",
			Path:    []string{"/v1/sceneIntelligenceResult/DeleteByIds"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SceneIntelligenceResultService.Update",
			Path:    []string{"/v1/sceneIntelligenceResult/update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SceneIntelligenceResultService.UpdateAll",
			Path:    []string{"/v1/sceneIntelligenceResult/updateAll"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SceneIntelligenceResultService.UpdateFields",
			Path:    []string{"/v1/sceneIntelligenceResult/updateFields"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SceneIntelligenceResultService.FindById",
			Path:    []string{"/v1/sceneIntelligenceResult/findById"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SceneIntelligenceResultService.Find",
			Path:    []string{"/v1/sceneIntelligenceResult/find"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "SceneIntelligenceResultService.Lists",
			Path:    []string{"/v1/sceneIntelligenceResult/lists"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for SceneIntelligenceResultService service

type SceneIntelligenceResultService interface {
	//创建
	Create(ctx context.Context, in *SceneIntelligenceResult, opts ...client.CallOption) (*Response, error)
	//匹配多条件删除
	Delete(ctx context.Context, in *SceneIntelligenceResult, opts ...client.CallOption) (*Response, error)
	//匹配主键删除,一般是Id
	DeleteById(ctx context.Context, in *SceneIntelligenceResult, opts ...client.CallOption) (*Response, error)
	//匹配主键批量删除
	DeleteByIds(ctx context.Context, in *SceneIntelligenceResultBatchDeleteRequest, opts ...client.CallOption) (*Response, error)
	//根据主键更新非空字段
	Update(ctx context.Context, in *SceneIntelligenceResult, opts ...client.CallOption) (*Response, error)
	//根据主键更新所有字段
	UpdateAll(ctx context.Context, in *SceneIntelligenceResult, opts ...client.CallOption) (*Response, error)
	//根据主键更新指定列
	UpdateFields(ctx context.Context, in *SceneIntelligenceResultUpdateFieldsRequest, opts ...client.CallOption) (*Response, error)
	//根据主键查找,一般是Id,返回单条数据
	FindById(ctx context.Context, in *SceneIntelligenceResultFilter, opts ...client.CallOption) (*SceneIntelligenceResultResponse, error)
	//多条件查找，返回单条数据
	Find(ctx context.Context, in *SceneIntelligenceResultFilter, opts ...client.CallOption) (*SceneIntelligenceResultResponse, error)
	//查找，支持分页，可返回多条数据
	Lists(ctx context.Context, in *SceneIntelligenceResultListRequest, opts ...client.CallOption) (*SceneIntelligenceResultResponse, error)
}

type sceneIntelligenceResultService struct {
	c    client.Client
	name string
}

func NewSceneIntelligenceResultService(name string, c client.Client) SceneIntelligenceResultService {
	return &sceneIntelligenceResultService{
		c:    c,
		name: name,
	}
}

func (c *sceneIntelligenceResultService) Create(ctx context.Context, in *SceneIntelligenceResult, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SceneIntelligenceResultService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneIntelligenceResultService) Delete(ctx context.Context, in *SceneIntelligenceResult, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SceneIntelligenceResultService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneIntelligenceResultService) DeleteById(ctx context.Context, in *SceneIntelligenceResult, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SceneIntelligenceResultService.DeleteById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneIntelligenceResultService) DeleteByIds(ctx context.Context, in *SceneIntelligenceResultBatchDeleteRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SceneIntelligenceResultService.DeleteByIds", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneIntelligenceResultService) Update(ctx context.Context, in *SceneIntelligenceResult, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SceneIntelligenceResultService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneIntelligenceResultService) UpdateAll(ctx context.Context, in *SceneIntelligenceResult, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SceneIntelligenceResultService.UpdateAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneIntelligenceResultService) UpdateFields(ctx context.Context, in *SceneIntelligenceResultUpdateFieldsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SceneIntelligenceResultService.UpdateFields", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneIntelligenceResultService) FindById(ctx context.Context, in *SceneIntelligenceResultFilter, opts ...client.CallOption) (*SceneIntelligenceResultResponse, error) {
	req := c.c.NewRequest(c.name, "SceneIntelligenceResultService.FindById", in)
	out := new(SceneIntelligenceResultResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneIntelligenceResultService) Find(ctx context.Context, in *SceneIntelligenceResultFilter, opts ...client.CallOption) (*SceneIntelligenceResultResponse, error) {
	req := c.c.NewRequest(c.name, "SceneIntelligenceResultService.Find", in)
	out := new(SceneIntelligenceResultResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneIntelligenceResultService) Lists(ctx context.Context, in *SceneIntelligenceResultListRequest, opts ...client.CallOption) (*SceneIntelligenceResultResponse, error) {
	req := c.c.NewRequest(c.name, "SceneIntelligenceResultService.Lists", in)
	out := new(SceneIntelligenceResultResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SceneIntelligenceResultService service

type SceneIntelligenceResultServiceHandler interface {
	//创建
	Create(context.Context, *SceneIntelligenceResult, *Response) error
	//匹配多条件删除
	Delete(context.Context, *SceneIntelligenceResult, *Response) error
	//匹配主键删除,一般是Id
	DeleteById(context.Context, *SceneIntelligenceResult, *Response) error
	//匹配主键批量删除
	DeleteByIds(context.Context, *SceneIntelligenceResultBatchDeleteRequest, *Response) error
	//根据主键更新非空字段
	Update(context.Context, *SceneIntelligenceResult, *Response) error
	//根据主键更新所有字段
	UpdateAll(context.Context, *SceneIntelligenceResult, *Response) error
	//根据主键更新指定列
	UpdateFields(context.Context, *SceneIntelligenceResultUpdateFieldsRequest, *Response) error
	//根据主键查找,一般是Id,返回单条数据
	FindById(context.Context, *SceneIntelligenceResultFilter, *SceneIntelligenceResultResponse) error
	//多条件查找，返回单条数据
	Find(context.Context, *SceneIntelligenceResultFilter, *SceneIntelligenceResultResponse) error
	//查找，支持分页，可返回多条数据
	Lists(context.Context, *SceneIntelligenceResultListRequest, *SceneIntelligenceResultResponse) error
}

func RegisterSceneIntelligenceResultServiceHandler(s server.Server, hdlr SceneIntelligenceResultServiceHandler, opts ...server.HandlerOption) error {
	type sceneIntelligenceResultService interface {
		Create(ctx context.Context, in *SceneIntelligenceResult, out *Response) error
		Delete(ctx context.Context, in *SceneIntelligenceResult, out *Response) error
		DeleteById(ctx context.Context, in *SceneIntelligenceResult, out *Response) error
		DeleteByIds(ctx context.Context, in *SceneIntelligenceResultBatchDeleteRequest, out *Response) error
		Update(ctx context.Context, in *SceneIntelligenceResult, out *Response) error
		UpdateAll(ctx context.Context, in *SceneIntelligenceResult, out *Response) error
		UpdateFields(ctx context.Context, in *SceneIntelligenceResultUpdateFieldsRequest, out *Response) error
		FindById(ctx context.Context, in *SceneIntelligenceResultFilter, out *SceneIntelligenceResultResponse) error
		Find(ctx context.Context, in *SceneIntelligenceResultFilter, out *SceneIntelligenceResultResponse) error
		Lists(ctx context.Context, in *SceneIntelligenceResultListRequest, out *SceneIntelligenceResultResponse) error
	}
	type SceneIntelligenceResultService struct {
		sceneIntelligenceResultService
	}
	h := &sceneIntelligenceResultServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SceneIntelligenceResultService.Create",
		Path:    []string{"/v1/sceneIntelligenceResult/create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SceneIntelligenceResultService.Delete",
		Path:    []string{"/v1/sceneIntelligenceResult/delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SceneIntelligenceResultService.DeleteById",
		Path:    []string{"/v1/sceneIntelligenceResult/deleteById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SceneIntelligenceResultService.DeleteByIds",
		Path:    []string{"/v1/sceneIntelligenceResult/DeleteByIds"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SceneIntelligenceResultService.Update",
		Path:    []string{"/v1/sceneIntelligenceResult/update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SceneIntelligenceResultService.UpdateAll",
		Path:    []string{"/v1/sceneIntelligenceResult/updateAll"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SceneIntelligenceResultService.UpdateFields",
		Path:    []string{"/v1/sceneIntelligenceResult/updateFields"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SceneIntelligenceResultService.FindById",
		Path:    []string{"/v1/sceneIntelligenceResult/findById"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SceneIntelligenceResultService.Find",
		Path:    []string{"/v1/sceneIntelligenceResult/find"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "SceneIntelligenceResultService.Lists",
		Path:    []string{"/v1/sceneIntelligenceResult/lists"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&SceneIntelligenceResultService{h}, opts...))
}

type sceneIntelligenceResultServiceHandler struct {
	SceneIntelligenceResultServiceHandler
}

func (h *sceneIntelligenceResultServiceHandler) Create(ctx context.Context, in *SceneIntelligenceResult, out *Response) error {
	return h.SceneIntelligenceResultServiceHandler.Create(ctx, in, out)
}

func (h *sceneIntelligenceResultServiceHandler) Delete(ctx context.Context, in *SceneIntelligenceResult, out *Response) error {
	return h.SceneIntelligenceResultServiceHandler.Delete(ctx, in, out)
}

func (h *sceneIntelligenceResultServiceHandler) DeleteById(ctx context.Context, in *SceneIntelligenceResult, out *Response) error {
	return h.SceneIntelligenceResultServiceHandler.DeleteById(ctx, in, out)
}

func (h *sceneIntelligenceResultServiceHandler) DeleteByIds(ctx context.Context, in *SceneIntelligenceResultBatchDeleteRequest, out *Response) error {
	return h.SceneIntelligenceResultServiceHandler.DeleteByIds(ctx, in, out)
}

func (h *sceneIntelligenceResultServiceHandler) Update(ctx context.Context, in *SceneIntelligenceResult, out *Response) error {
	return h.SceneIntelligenceResultServiceHandler.Update(ctx, in, out)
}

func (h *sceneIntelligenceResultServiceHandler) UpdateAll(ctx context.Context, in *SceneIntelligenceResult, out *Response) error {
	return h.SceneIntelligenceResultServiceHandler.UpdateAll(ctx, in, out)
}

func (h *sceneIntelligenceResultServiceHandler) UpdateFields(ctx context.Context, in *SceneIntelligenceResultUpdateFieldsRequest, out *Response) error {
	return h.SceneIntelligenceResultServiceHandler.UpdateFields(ctx, in, out)
}

func (h *sceneIntelligenceResultServiceHandler) FindById(ctx context.Context, in *SceneIntelligenceResultFilter, out *SceneIntelligenceResultResponse) error {
	return h.SceneIntelligenceResultServiceHandler.FindById(ctx, in, out)
}

func (h *sceneIntelligenceResultServiceHandler) Find(ctx context.Context, in *SceneIntelligenceResultFilter, out *SceneIntelligenceResultResponse) error {
	return h.SceneIntelligenceResultServiceHandler.Find(ctx, in, out)
}

func (h *sceneIntelligenceResultServiceHandler) Lists(ctx context.Context, in *SceneIntelligenceResultListRequest, out *SceneIntelligenceResultResponse) error {
	return h.SceneIntelligenceResultServiceHandler.Lists(ctx, in, out)
}
