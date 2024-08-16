// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v3.12.4
// source: api/customer/v1/customer.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCustomerGetVerifyCode = "/customer.v1.Customer/GetVerifyCode"
const OperationCustomerLogin = "/customer.v1.Customer/Login"

type CustomerHTTPServer interface {
	GetVerifyCode(context.Context, *GetVerifyCodeRequest) (*GetVerifyCodeReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterCustomerHTTPServer(s *http.Server, srv CustomerHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/customer/get-verify-code", _Customer_GetVerifyCode0_HTTP_Handler(srv))
	r.POST("/api/v1/customer/login", _Customer_Login0_HTTP_Handler(srv))
}

func _Customer_GetVerifyCode0_HTTP_Handler(srv CustomerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVerifyCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCustomerGetVerifyCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVerifyCode(ctx, req.(*GetVerifyCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetVerifyCodeReply)
		return ctx.Result(200, reply)
	}
}

func _Customer_Login0_HTTP_Handler(srv CustomerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCustomerLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

type CustomerHTTPClient interface {
	GetVerifyCode(ctx context.Context, req *GetVerifyCodeRequest, opts ...http.CallOption) (rsp *GetVerifyCodeReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
}

type CustomerHTTPClientImpl struct {
	cc *http.Client
}

func NewCustomerHTTPClient(client *http.Client) CustomerHTTPClient {
	return &CustomerHTTPClientImpl{client}
}

func (c *CustomerHTTPClientImpl) GetVerifyCode(ctx context.Context, in *GetVerifyCodeRequest, opts ...http.CallOption) (*GetVerifyCodeReply, error) {
	var out GetVerifyCodeReply
	pattern := "/api/v1/customer/get-verify-code"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCustomerGetVerifyCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CustomerHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/api/v1/customer/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCustomerLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
