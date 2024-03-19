// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: voice.proto

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

// Api Endpoints for VoiceService service

func NewVoiceServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for VoiceService service

type VoiceService interface {
	CreateClientInfo(ctx context.Context, in *ClientInfoReq, opts ...client.CallOption) (*ClientInfoResp, error)
}

type voiceService struct {
	c    client.Client
	name string
}

func NewVoiceService(name string, c client.Client) VoiceService {
	return &voiceService{
		c:    c,
		name: name,
	}
}

func (c *voiceService) CreateClientInfo(ctx context.Context, in *ClientInfoReq, opts ...client.CallOption) (*ClientInfoResp, error) {
	req := c.c.NewRequest(c.name, "VoiceService.CreateClientInfo", in)
	out := new(ClientInfoResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VoiceService service

type VoiceServiceHandler interface {
	CreateClientInfo(context.Context, *ClientInfoReq, *ClientInfoResp) error
}

func RegisterVoiceServiceHandler(s server.Server, hdlr VoiceServiceHandler, opts ...server.HandlerOption) error {
	type voiceService interface {
		CreateClientInfo(ctx context.Context, in *ClientInfoReq, out *ClientInfoResp) error
	}
	type VoiceService struct {
		voiceService
	}
	h := &voiceServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&VoiceService{h}, opts...))
}

type voiceServiceHandler struct {
	VoiceServiceHandler
}

func (h *voiceServiceHandler) CreateClientInfo(ctx context.Context, in *ClientInfoReq, out *ClientInfoResp) error {
	return h.VoiceServiceHandler.CreateClientInfo(ctx, in, out)
}
