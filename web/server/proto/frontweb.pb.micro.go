// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/frontweb.proto

package go_micro_srv_frontweb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/api"
	server "github.com/micro/go-micro/server"
	client "github.com/micro/go-micro/v2/client"
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

// Api Endpoints for Fweb service

func NewFwebEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Fweb service

type FwebService interface {
	MakeWeb(ctx context.Context, in *FrontRequest, opts ...client.CallOption) (*FrontRespons, error)
}

type fwebService struct {
	c    client.Client
	name string
}

func NewFwebService(name string, c client.Client) FwebService {
	return &fwebService{
		c:    c,
		name: name,
	}
}

func (c *fwebService) MakeWeb(ctx context.Context, in *FrontRequest, opts ...client.CallOption) (*FrontRespons, error) {
	req := c.c.NewRequest(c.name, "Fweb.MakeWeb", in)
	out := new(FrontRespons)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Fweb service

type FwebHandler interface {
	MakeWeb(context.Context, *FrontRequest, *FrontRespons) error
}

func RegisterFwebHandler(s server.Server, hdlr FwebHandler, opts ...server.HandlerOption) error {
	type fweb interface {
		MakeWeb(ctx context.Context, in *FrontRequest, out *FrontRespons) error
	}
	type Fweb struct {
		fweb
	}
	h := &fwebHandler{hdlr}
	return s.Handle(s.NewHandler(&Fweb{h}, opts...))
}

type fwebHandler struct {
	FwebHandler
}

func (h *fwebHandler) MakeWeb(ctx context.Context, in *FrontRequest, out *FrontRespons) error {
	return h.FwebHandler.MakeWeb(ctx, in, out)
}
