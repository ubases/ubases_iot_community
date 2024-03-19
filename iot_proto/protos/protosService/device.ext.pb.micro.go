// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: device.ext.proto

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

// Api Endpoints for DeviceInfoExtService service

func NewDeviceInfoExtServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DeviceInfoExtService service

type DeviceInfoExtService interface {
	GetDeviceInfo(ctx context.Context, in *DeviceInfoFilterReq, opts ...client.CallOption) (*DeviceInfoRsp, error)
	GetDeviceInfoList(ctx context.Context, in *DeviceInfoListFilterReq, opts ...client.CallOption) (*DeviceInfoRsp, error)
	SetDeviceActive(ctx context.Context, in *DeviceActiveReq, opts ...client.CallOption) (*DeviceActiveRsp, error)
}

type deviceInfoExtService struct {
	c    client.Client
	name string
}

func NewDeviceInfoExtService(name string, c client.Client) DeviceInfoExtService {
	return &deviceInfoExtService{
		c:    c,
		name: name,
	}
}

func (c *deviceInfoExtService) GetDeviceInfo(ctx context.Context, in *DeviceInfoFilterReq, opts ...client.CallOption) (*DeviceInfoRsp, error) {
	req := c.c.NewRequest(c.name, "DeviceInfoExtService.GetDeviceInfo", in)
	out := new(DeviceInfoRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceInfoExtService) GetDeviceInfoList(ctx context.Context, in *DeviceInfoListFilterReq, opts ...client.CallOption) (*DeviceInfoRsp, error) {
	req := c.c.NewRequest(c.name, "DeviceInfoExtService.GetDeviceInfoList", in)
	out := new(DeviceInfoRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceInfoExtService) SetDeviceActive(ctx context.Context, in *DeviceActiveReq, opts ...client.CallOption) (*DeviceActiveRsp, error) {
	req := c.c.NewRequest(c.name, "DeviceInfoExtService.SetDeviceActive", in)
	out := new(DeviceActiveRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceInfoExtService service

type DeviceInfoExtServiceHandler interface {
	GetDeviceInfo(context.Context, *DeviceInfoFilterReq, *DeviceInfoRsp) error
	GetDeviceInfoList(context.Context, *DeviceInfoListFilterReq, *DeviceInfoRsp) error
	SetDeviceActive(context.Context, *DeviceActiveReq, *DeviceActiveRsp) error
}

func RegisterDeviceInfoExtServiceHandler(s server.Server, hdlr DeviceInfoExtServiceHandler, opts ...server.HandlerOption) error {
	type deviceInfoExtService interface {
		GetDeviceInfo(ctx context.Context, in *DeviceInfoFilterReq, out *DeviceInfoRsp) error
		GetDeviceInfoList(ctx context.Context, in *DeviceInfoListFilterReq, out *DeviceInfoRsp) error
		SetDeviceActive(ctx context.Context, in *DeviceActiveReq, out *DeviceActiveRsp) error
	}
	type DeviceInfoExtService struct {
		deviceInfoExtService
	}
	h := &deviceInfoExtServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DeviceInfoExtService{h}, opts...))
}

type deviceInfoExtServiceHandler struct {
	DeviceInfoExtServiceHandler
}

func (h *deviceInfoExtServiceHandler) GetDeviceInfo(ctx context.Context, in *DeviceInfoFilterReq, out *DeviceInfoRsp) error {
	return h.DeviceInfoExtServiceHandler.GetDeviceInfo(ctx, in, out)
}

func (h *deviceInfoExtServiceHandler) GetDeviceInfoList(ctx context.Context, in *DeviceInfoListFilterReq, out *DeviceInfoRsp) error {
	return h.DeviceInfoExtServiceHandler.GetDeviceInfoList(ctx, in, out)
}

func (h *deviceInfoExtServiceHandler) SetDeviceActive(ctx context.Context, in *DeviceActiveReq, out *DeviceActiveRsp) error {
	return h.DeviceInfoExtServiceHandler.SetDeviceActive(ctx, in, out)
}
