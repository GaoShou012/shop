// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: jwt.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Jwt service

type JwtService interface {
	Encode(ctx context.Context, in *JwtEncodeRequest, opts ...client.CallOption) (*JwtEncodeResponse, error)
	Decode(ctx context.Context, in *JwtDecodeRequest, opts ...client.CallOption) (*JwtDecodeResponse, error)
}

type jwtService struct {
	c    client.Client
	name string
}

func NewJwtService(name string, c client.Client) JwtService {
	return &jwtService{
		c:    c,
		name: name,
	}
}

func (c *jwtService) Encode(ctx context.Context, in *JwtEncodeRequest, opts ...client.CallOption) (*JwtEncodeResponse, error) {
	req := c.c.NewRequest(c.name, "Jwt.Encode", in)
	out := new(JwtEncodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwtService) Decode(ctx context.Context, in *JwtDecodeRequest, opts ...client.CallOption) (*JwtDecodeResponse, error) {
	req := c.c.NewRequest(c.name, "Jwt.Decode", in)
	out := new(JwtDecodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Jwt service

type JwtHandler interface {
	Encode(context.Context, *JwtEncodeRequest, *JwtEncodeResponse) error
	Decode(context.Context, *JwtDecodeRequest, *JwtDecodeResponse) error
}

func RegisterJwtHandler(s server.Server, hdlr JwtHandler, opts ...server.HandlerOption) error {
	type jwt interface {
		Encode(ctx context.Context, in *JwtEncodeRequest, out *JwtEncodeResponse) error
		Decode(ctx context.Context, in *JwtDecodeRequest, out *JwtDecodeResponse) error
	}
	type Jwt struct {
		jwt
	}
	h := &jwtHandler{hdlr}
	return s.Handle(s.NewHandler(&Jwt{h}, opts...))
}

type jwtHandler struct {
	JwtHandler
}

func (h *jwtHandler) Encode(ctx context.Context, in *JwtEncodeRequest, out *JwtEncodeResponse) error {
	return h.JwtHandler.Encode(ctx, in, out)
}

func (h *jwtHandler) Decode(ctx context.Context, in *JwtDecodeRequest, out *JwtDecodeResponse) error {
	return h.JwtHandler.Decode(ctx, in, out)
}