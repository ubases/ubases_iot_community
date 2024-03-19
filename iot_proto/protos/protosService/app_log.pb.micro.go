// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: app_log.proto

package protosService

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Api Endpoints for AppLogService service

func NewAppLogServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AppLogService service

type AppLogService interface {
	CreateAppLogUser(ctx context.Context, in *AppLogUser, opts ...client.CallOption) (*AppLogCommonResponse, error)
	UpdateAppLogUser(ctx context.Context, in *AppLogUser, opts ...client.CallOption) (*AppLogCommonResponse, error)
	DeleteAppLogUser(ctx context.Context, in *AppLogUser, opts ...client.CallOption) (*AppLogCommonResponse, error)
	GetAppLogUser(ctx context.Context, in *AppLogUser, opts ...client.CallOption) (*AppLogUser, error)
	GetAppLogUserList(ctx context.Context, in *AppLogUserListReq, opts ...client.CallOption) (*AppLogUserListResp, error)
	GetAppLogRecordsList(ctx context.Context, in *AppLogRecordsListReq, opts ...client.CallOption) (*AppLogRecordsListResp, error)
}

type appLogService struct {
	c    client.Client
	name string
}

func NewAppLogService(name string, c client.Client) AppLogService {
	return &appLogService{
		c:    c,
		name: name,
	}
}

func (c *appLogService) CreateAppLogUser(ctx context.Context, in *AppLogUser, opts ...client.CallOption) (*AppLogCommonResponse, error) {
	req := c.c.NewRequest(c.name, "AppLogService.CreateAppLogUser", in)
	out := new(AppLogCommonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appLogService) UpdateAppLogUser(ctx context.Context, in *AppLogUser, opts ...client.CallOption) (*AppLogCommonResponse, error) {
	req := c.c.NewRequest(c.name, "AppLogService.UpdateAppLogUser", in)
	out := new(AppLogCommonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appLogService) DeleteAppLogUser(ctx context.Context, in *AppLogUser, opts ...client.CallOption) (*AppLogCommonResponse, error) {
	req := c.c.NewRequest(c.name, "AppLogService.DeleteAppLogUser", in)
	out := new(AppLogCommonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appLogService) GetAppLogUser(ctx context.Context, in *AppLogUser, opts ...client.CallOption) (*AppLogUser, error) {
	req := c.c.NewRequest(c.name, "AppLogService.GetAppLogUser", in)
	out := new(AppLogUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appLogService) GetAppLogUserList(ctx context.Context, in *AppLogUserListReq, opts ...client.CallOption) (*AppLogUserListResp, error) {
	req := c.c.NewRequest(c.name, "AppLogService.GetAppLogUserList", in)
	out := new(AppLogUserListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appLogService) GetAppLogRecordsList(ctx context.Context, in *AppLogRecordsListReq, opts ...client.CallOption) (*AppLogRecordsListResp, error) {
	req := c.c.NewRequest(c.name, "AppLogService.GetAppLogRecordsList", in)
	out := new(AppLogRecordsListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppLogService service

type AppLogServiceHandler interface {
	CreateAppLogUser(context.Context, *AppLogUser, *AppLogCommonResponse) error
	UpdateAppLogUser(context.Context, *AppLogUser, *AppLogCommonResponse) error
	DeleteAppLogUser(context.Context, *AppLogUser, *AppLogCommonResponse) error
	GetAppLogUser(context.Context, *AppLogUser, *AppLogUser) error
	GetAppLogUserList(context.Context, *AppLogUserListReq, *AppLogUserListResp) error
	GetAppLogRecordsList(context.Context, *AppLogRecordsListReq, *AppLogRecordsListResp) error
}

func RegisterAppLogServiceHandler(s server.Server, hdlr AppLogServiceHandler, opts ...server.HandlerOption) error {
	type appLogService interface {
		CreateAppLogUser(ctx context.Context, in *AppLogUser, out *AppLogCommonResponse) error
		UpdateAppLogUser(ctx context.Context, in *AppLogUser, out *AppLogCommonResponse) error
		DeleteAppLogUser(ctx context.Context, in *AppLogUser, out *AppLogCommonResponse) error
		GetAppLogUser(ctx context.Context, in *AppLogUser, out *AppLogUser) error
		GetAppLogUserList(ctx context.Context, in *AppLogUserListReq, out *AppLogUserListResp) error
		GetAppLogRecordsList(ctx context.Context, in *AppLogRecordsListReq, out *AppLogRecordsListResp) error
	}
	type AppLogService struct {
		appLogService
	}
	h := &appLogServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AppLogService{h}, opts...))
}

type appLogServiceHandler struct {
	AppLogServiceHandler
}

func (h *appLogServiceHandler) CreateAppLogUser(ctx context.Context, in *AppLogUser, out *AppLogCommonResponse) error {
	return h.AppLogServiceHandler.CreateAppLogUser(ctx, in, out)
}

func (h *appLogServiceHandler) UpdateAppLogUser(ctx context.Context, in *AppLogUser, out *AppLogCommonResponse) error {
	return h.AppLogServiceHandler.UpdateAppLogUser(ctx, in, out)
}

func (h *appLogServiceHandler) DeleteAppLogUser(ctx context.Context, in *AppLogUser, out *AppLogCommonResponse) error {
	return h.AppLogServiceHandler.DeleteAppLogUser(ctx, in, out)
}

func (h *appLogServiceHandler) GetAppLogUser(ctx context.Context, in *AppLogUser, out *AppLogUser) error {
	return h.AppLogServiceHandler.GetAppLogUser(ctx, in, out)
}

func (h *appLogServiceHandler) GetAppLogUserList(ctx context.Context, in *AppLogUserListReq, out *AppLogUserListResp) error {
	return h.AppLogServiceHandler.GetAppLogUserList(ctx, in, out)
}

func (h *appLogServiceHandler) GetAppLogRecordsList(ctx context.Context, in *AppLogRecordsListReq, out *AppLogRecordsListResp) error {
	return h.AppLogServiceHandler.GetAppLogRecordsList(ctx, in, out)
}
