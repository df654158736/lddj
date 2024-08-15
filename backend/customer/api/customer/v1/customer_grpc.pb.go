// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: api/customer/v1/customer.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Customer_CreateCustomer_FullMethodName = "/customer.v1.Customer/CreateCustomer"
	Customer_UpdateCustomer_FullMethodName = "/customer.v1.Customer/UpdateCustomer"
	Customer_DeleteCustomer_FullMethodName = "/customer.v1.Customer/DeleteCustomer"
	Customer_GetCustomer_FullMethodName    = "/customer.v1.Customer/GetCustomer"
	Customer_ListCustomer_FullMethodName   = "/customer.v1.Customer/ListCustomer"
	Customer_GetVerifyCode_FullMethodName  = "/customer.v1.Customer/GetVerifyCode"
)

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerClient interface {
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerReply, error)
	UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*UpdateCustomerReply, error)
	DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*DeleteCustomerReply, error)
	GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*GetCustomerReply, error)
	ListCustomer(ctx context.Context, in *ListCustomerRequest, opts ...grpc.CallOption) (*ListCustomerReply, error)
	GetVerifyCode(ctx context.Context, in *GetVerifyCodeRequest, opts ...grpc.CallOption) (*GetVerifyCodeReply, error)
}

type customerClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClient(cc grpc.ClientConnInterface) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCustomerReply)
	err := c.cc.Invoke(ctx, Customer_CreateCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*UpdateCustomerReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCustomerReply)
	err := c.cc.Invoke(ctx, Customer_UpdateCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*DeleteCustomerReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCustomerReply)
	err := c.cc.Invoke(ctx, Customer_DeleteCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*GetCustomerReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCustomerReply)
	err := c.cc.Invoke(ctx, Customer_GetCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) ListCustomer(ctx context.Context, in *ListCustomerRequest, opts ...grpc.CallOption) (*ListCustomerReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCustomerReply)
	err := c.cc.Invoke(ctx, Customer_ListCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) GetVerifyCode(ctx context.Context, in *GetVerifyCodeRequest, opts ...grpc.CallOption) (*GetVerifyCodeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVerifyCodeReply)
	err := c.cc.Invoke(ctx, Customer_GetVerifyCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
// All implementations must embed UnimplementedCustomerServer
// for forward compatibility.
type CustomerServer interface {
	CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerReply, error)
	UpdateCustomer(context.Context, *UpdateCustomerRequest) (*UpdateCustomerReply, error)
	DeleteCustomer(context.Context, *DeleteCustomerRequest) (*DeleteCustomerReply, error)
	GetCustomer(context.Context, *GetCustomerRequest) (*GetCustomerReply, error)
	ListCustomer(context.Context, *ListCustomerRequest) (*ListCustomerReply, error)
	GetVerifyCode(context.Context, *GetVerifyCodeRequest) (*GetVerifyCodeReply, error)
	mustEmbedUnimplementedCustomerServer()
}

// UnimplementedCustomerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCustomerServer struct{}

func (UnimplementedCustomerServer) CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerServer) UpdateCustomer(context.Context, *UpdateCustomerRequest) (*UpdateCustomerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedCustomerServer) DeleteCustomer(context.Context, *DeleteCustomerRequest) (*DeleteCustomerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
func (UnimplementedCustomerServer) GetCustomer(context.Context, *GetCustomerRequest) (*GetCustomerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (UnimplementedCustomerServer) ListCustomer(context.Context, *ListCustomerRequest) (*ListCustomerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCustomer not implemented")
}
func (UnimplementedCustomerServer) GetVerifyCode(context.Context, *GetVerifyCodeRequest) (*GetVerifyCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVerifyCode not implemented")
}
func (UnimplementedCustomerServer) mustEmbedUnimplementedCustomerServer() {}
func (UnimplementedCustomerServer) testEmbeddedByValue()                  {}

// UnsafeCustomerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServer will
// result in compilation errors.
type UnsafeCustomerServer interface {
	mustEmbedUnimplementedCustomerServer()
}

func RegisterCustomerServer(s grpc.ServiceRegistrar, srv CustomerServer) {
	// If the following call pancis, it indicates UnimplementedCustomerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Customer_ServiceDesc, srv)
}

func _Customer_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_CreateCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).CreateCustomer(ctx, req.(*CreateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_UpdateCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).UpdateCustomer(ctx, req.(*UpdateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_DeleteCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).DeleteCustomer(ctx, req.(*DeleteCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_GetCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetCustomer(ctx, req.(*GetCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_ListCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).ListCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_ListCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).ListCustomer(ctx, req.(*ListCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_GetVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_GetVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetVerifyCode(ctx, req.(*GetVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Customer_ServiceDesc is the grpc.ServiceDesc for Customer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Customer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer.v1.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _Customer_CreateCustomer_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _Customer_UpdateCustomer_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _Customer_DeleteCustomer_Handler,
		},
		{
			MethodName: "GetCustomer",
			Handler:    _Customer_GetCustomer_Handler,
		},
		{
			MethodName: "ListCustomer",
			Handler:    _Customer_ListCustomer_Handler,
		},
		{
			MethodName: "GetVerifyCode",
			Handler:    _Customer_GetVerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/customer/v1/customer.proto",
}
