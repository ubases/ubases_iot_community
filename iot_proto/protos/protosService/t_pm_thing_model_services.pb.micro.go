// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: t_pm_thing_model_services.proto

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

// Api Endpoints for TPmThingModelServices service

func NewTPmThingModelServicesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TPmThingModelServices service

type TPmThingModelServicesService interface {
	UpdateTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, opts ...client.CallOption) (*TPmThingModelServicesResponse, error)
	DeleteTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, opts ...client.CallOption) (*TPmThingModelServicesResponse, error)
	GetByIdTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilterById, opts ...client.CallOption) (*TPmThingModelServicesResponseObject, error)
	GetTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilter, opts ...client.CallOption) (*TPmThingModelServicesResponseObject, error)
	ListTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilterPage, opts ...client.CallOption) (*TPmThingModelServicesResponseList, error)
	CreateTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, opts ...client.CallOption) (*TPmThingModelServicesResponse, error)
}

type tPmThingModelServicesService struct {
	c    client.Client
	name string
}

func NewTPmThingModelServicesService(name string, c client.Client) TPmThingModelServicesService {
	return &tPmThingModelServicesService{
		c:    c,
		name: name,
	}
}

func (c *tPmThingModelServicesService) UpdateTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, opts ...client.CallOption) (*TPmThingModelServicesResponse, error) {
	req := c.c.NewRequest(c.name, "TPmThingModelServices.UpdateTPmThingModelServices", in)
	out := new(TPmThingModelServicesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tPmThingModelServicesService) DeleteTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, opts ...client.CallOption) (*TPmThingModelServicesResponse, error) {
	req := c.c.NewRequest(c.name, "TPmThingModelServices.DeleteTPmThingModelServices", in)
	out := new(TPmThingModelServicesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tPmThingModelServicesService) GetByIdTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilterById, opts ...client.CallOption) (*TPmThingModelServicesResponseObject, error) {
	req := c.c.NewRequest(c.name, "TPmThingModelServices.GetByIdTPmThingModelServices", in)
	out := new(TPmThingModelServicesResponseObject)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tPmThingModelServicesService) GetTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilter, opts ...client.CallOption) (*TPmThingModelServicesResponseObject, error) {
	req := c.c.NewRequest(c.name, "TPmThingModelServices.GetTPmThingModelServices", in)
	out := new(TPmThingModelServicesResponseObject)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tPmThingModelServicesService) ListTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilterPage, opts ...client.CallOption) (*TPmThingModelServicesResponseList, error) {
	req := c.c.NewRequest(c.name, "TPmThingModelServices.ListTPmThingModelServices", in)
	out := new(TPmThingModelServicesResponseList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tPmThingModelServicesService) CreateTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, opts ...client.CallOption) (*TPmThingModelServicesResponse, error) {
	req := c.c.NewRequest(c.name, "TPmThingModelServices.CreateTPmThingModelServices", in)
	out := new(TPmThingModelServicesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TPmThingModelServices service

type TPmThingModelServicesHandler interface {
	UpdateTPmThingModelServices(context.Context, *TPmThingModelServicesRequest, *TPmThingModelServicesResponse) error
	DeleteTPmThingModelServices(context.Context, *TPmThingModelServicesRequest, *TPmThingModelServicesResponse) error
	GetByIdTPmThingModelServices(context.Context, *TPmThingModelServicesFilterById, *TPmThingModelServicesResponseObject) error
	GetTPmThingModelServices(context.Context, *TPmThingModelServicesFilter, *TPmThingModelServicesResponseObject) error
	ListTPmThingModelServices(context.Context, *TPmThingModelServicesFilterPage, *TPmThingModelServicesResponseList) error
	CreateTPmThingModelServices(context.Context, *TPmThingModelServicesRequest, *TPmThingModelServicesResponse) error
}

func RegisterTPmThingModelServicesHandler(s server.Server, hdlr TPmThingModelServicesHandler, opts ...server.HandlerOption) error {
	type tPmThingModelServices interface {
		UpdateTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, out *TPmThingModelServicesResponse) error
		DeleteTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, out *TPmThingModelServicesResponse) error
		GetByIdTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilterById, out *TPmThingModelServicesResponseObject) error
		GetTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilter, out *TPmThingModelServicesResponseObject) error
		ListTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilterPage, out *TPmThingModelServicesResponseList) error
		CreateTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, out *TPmThingModelServicesResponse) error
	}
	type TPmThingModelServices struct {
		tPmThingModelServices
	}
	h := &tPmThingModelServicesHandler{hdlr}
	return s.Handle(s.NewHandler(&TPmThingModelServices{h}, opts...))
}

type tPmThingModelServicesHandler struct {
	TPmThingModelServicesHandler
}

func (h *tPmThingModelServicesHandler) UpdateTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, out *TPmThingModelServicesResponse) error {
	return h.TPmThingModelServicesHandler.UpdateTPmThingModelServices(ctx, in, out)
}

func (h *tPmThingModelServicesHandler) DeleteTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, out *TPmThingModelServicesResponse) error {
	return h.TPmThingModelServicesHandler.DeleteTPmThingModelServices(ctx, in, out)
}

func (h *tPmThingModelServicesHandler) GetByIdTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilterById, out *TPmThingModelServicesResponseObject) error {
	return h.TPmThingModelServicesHandler.GetByIdTPmThingModelServices(ctx, in, out)
}

func (h *tPmThingModelServicesHandler) GetTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilter, out *TPmThingModelServicesResponseObject) error {
	return h.TPmThingModelServicesHandler.GetTPmThingModelServices(ctx, in, out)
}

func (h *tPmThingModelServicesHandler) ListTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesFilterPage, out *TPmThingModelServicesResponseList) error {
	return h.TPmThingModelServicesHandler.ListTPmThingModelServices(ctx, in, out)
}

func (h *tPmThingModelServicesHandler) CreateTPmThingModelServices(ctx context.Context, in *TPmThingModelServicesRequest, out *TPmThingModelServicesResponse) error {
	return h.TPmThingModelServicesHandler.CreateTPmThingModelServices(ctx, in, out)
}
