// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: logger.proto

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

// Client API for Logger service

type LoggerService interface {
	Push(ctx context.Context, in *LoggerPushRequest, opts ...client.CallOption) (*LoggerPushResponse, error)
}

type loggerService struct {
	c    client.Client
	name string
}

func NewLoggerService(name string, c client.Client) LoggerService {
	return &loggerService{
		c:    c,
		name: name,
	}
}

func (c *loggerService) Push(ctx context.Context, in *LoggerPushRequest, opts ...client.CallOption) (*LoggerPushResponse, error) {
	req := c.c.NewRequest(c.name, "Logger.Push", in)
	out := new(LoggerPushResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Logger service

type LoggerHandler interface {
	Push(context.Context, *LoggerPushRequest, *LoggerPushResponse) error
}

func RegisterLoggerHandler(s server.Server, hdlr LoggerHandler, opts ...server.HandlerOption) error {
	type logger interface {
		Push(ctx context.Context, in *LoggerPushRequest, out *LoggerPushResponse) error
	}
	type Logger struct {
		logger
	}
	h := &loggerHandler{hdlr}
	return s.Handle(s.NewHandler(&Logger{h}, opts...))
}

type loggerHandler struct {
	LoggerHandler
}

func (h *loggerHandler) Push(ctx context.Context, in *LoggerPushRequest, out *LoggerPushResponse) error {
	return h.LoggerHandler.Push(ctx, in, out)
}
