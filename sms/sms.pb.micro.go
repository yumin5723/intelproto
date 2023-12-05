// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sms.proto

package sms

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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

// Api Endpoints for Sms service

func NewSmsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Sms service

type SmsService interface {
	SendSms(ctx context.Context, in *SmsRequest, opts ...client.CallOption) (*Response, error)
	ValidSms(ctx context.Context, in *ValidSmsRequest, opts ...client.CallOption) (*ValidResponse, error)
}

type smsService struct {
	c    client.Client
	name string
}

func NewSmsService(name string, c client.Client) SmsService {
	return &smsService{
		c:    c,
		name: name,
	}
}

func (c *smsService) SendSms(ctx context.Context, in *SmsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Sms.SendSms", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsService) ValidSms(ctx context.Context, in *ValidSmsRequest, opts ...client.CallOption) (*ValidResponse, error) {
	req := c.c.NewRequest(c.name, "Sms.ValidSms", in)
	out := new(ValidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sms service

type SmsHandler interface {
	SendSms(context.Context, *SmsRequest, *Response) error
	ValidSms(context.Context, *ValidSmsRequest, *ValidResponse) error
}

func RegisterSmsHandler(s server.Server, hdlr SmsHandler, opts ...server.HandlerOption) error {
	type sms interface {
		SendSms(ctx context.Context, in *SmsRequest, out *Response) error
		ValidSms(ctx context.Context, in *ValidSmsRequest, out *ValidResponse) error
	}
	type Sms struct {
		sms
	}
	h := &smsHandler{hdlr}
	return s.Handle(s.NewHandler(&Sms{h}, opts...))
}

type smsHandler struct {
	SmsHandler
}

func (h *smsHandler) SendSms(ctx context.Context, in *SmsRequest, out *Response) error {
	return h.SmsHandler.SendSms(ctx, in, out)
}

func (h *smsHandler) ValidSms(ctx context.Context, in *ValidSmsRequest, out *ValidResponse) error {
	return h.SmsHandler.ValidSms(ctx, in, out)
}
