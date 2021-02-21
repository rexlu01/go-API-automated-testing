// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/mweb.proto

package go_micro_srv_sendmweb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Mweb service

func NewMwebEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Mweb service

type MwebService interface {
	SendMweb(ctx context.Context, in *TryRequest, opts ...client.CallOption) (*TryRespons, error)
}

type mwebService struct {
	c    client.Client
	name string
}

func NewMwebService(name string, c client.Client) MwebService {
	return &mwebService{
		c:    c,
		name: name,
	}
}

func (c *mwebService) SendMweb(ctx context.Context, in *TryRequest, opts ...client.CallOption) (*TryRespons, error) {
	req := c.c.NewRequest(c.name, "Mweb.SendMweb", in)
	out := new(TryRespons)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mweb service

type MwebHandler interface {
	SendMweb(context.Context, *TryRequest, *TryRespons) error
}

func RegisterMwebHandler(s server.Server, hdlr MwebHandler, opts ...server.HandlerOption) error {
	type mweb interface {
		SendMweb(ctx context.Context, in *TryRequest, out *TryRespons) error
	}
	type Mweb struct {
		mweb
	}
	h := &mwebHandler{hdlr}
	return s.Handle(s.NewHandler(&Mweb{h}, opts...))
}

type mwebHandler struct {
	MwebHandler
}

func (h *mwebHandler) SendMweb(ctx context.Context, in *TryRequest, out *TryRespons) error {
	return h.MwebHandler.SendMweb(ctx, in, out)
}
