// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: api/verify_code/verify_code.proto

package s

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
	VerifyCode_CreateVerifyCode_FullMethodName = "/s.v1.VerifyCode/CreateVerifyCode"
	VerifyCode_UpdateVerifyCode_FullMethodName = "/s.v1.VerifyCode/UpdateVerifyCode"
	VerifyCode_DeleteVerifyCode_FullMethodName = "/s.v1.VerifyCode/DeleteVerifyCode"
	VerifyCode_GetVerifyCode_FullMethodName    = "/s.v1.VerifyCode/GetVerifyCode"
	VerifyCode_ListVerifyCode_FullMethodName   = "/s.v1.VerifyCode/ListVerifyCode"
)

// VerifyCodeClient is the client API for VerifyCode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerifyCodeClient interface {
	CreateVerifyCode(ctx context.Context, in *CreateVerifyCodeRequest, opts ...grpc.CallOption) (*CreateVerifyCodeReply, error)
	UpdateVerifyCode(ctx context.Context, in *UpdateVerifyCodeRequest, opts ...grpc.CallOption) (*UpdateVerifyCodeReply, error)
	DeleteVerifyCode(ctx context.Context, in *DeleteVerifyCodeRequest, opts ...grpc.CallOption) (*DeleteVerifyCodeReply, error)
	GetVerifyCode(ctx context.Context, in *GetVerifyCodeRequest, opts ...grpc.CallOption) (*GetVerifyCodeReply, error)
	ListVerifyCode(ctx context.Context, in *ListVerifyCodeRequest, opts ...grpc.CallOption) (*ListVerifyCodeReply, error)
}

type verifyCodeClient struct {
	cc grpc.ClientConnInterface
}

func NewVerifyCodeClient(cc grpc.ClientConnInterface) VerifyCodeClient {
	return &verifyCodeClient{cc}
}

func (c *verifyCodeClient) CreateVerifyCode(ctx context.Context, in *CreateVerifyCodeRequest, opts ...grpc.CallOption) (*CreateVerifyCodeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVerifyCodeReply)
	err := c.cc.Invoke(ctx, VerifyCode_CreateVerifyCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyCodeClient) UpdateVerifyCode(ctx context.Context, in *UpdateVerifyCodeRequest, opts ...grpc.CallOption) (*UpdateVerifyCodeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateVerifyCodeReply)
	err := c.cc.Invoke(ctx, VerifyCode_UpdateVerifyCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyCodeClient) DeleteVerifyCode(ctx context.Context, in *DeleteVerifyCodeRequest, opts ...grpc.CallOption) (*DeleteVerifyCodeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteVerifyCodeReply)
	err := c.cc.Invoke(ctx, VerifyCode_DeleteVerifyCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyCodeClient) GetVerifyCode(ctx context.Context, in *GetVerifyCodeRequest, opts ...grpc.CallOption) (*GetVerifyCodeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVerifyCodeReply)
	err := c.cc.Invoke(ctx, VerifyCode_GetVerifyCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyCodeClient) ListVerifyCode(ctx context.Context, in *ListVerifyCodeRequest, opts ...grpc.CallOption) (*ListVerifyCodeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVerifyCodeReply)
	err := c.cc.Invoke(ctx, VerifyCode_ListVerifyCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifyCodeServer is the server API for VerifyCode service.
// All implementations must embed UnimplementedVerifyCodeServer
// for forward compatibility.
type VerifyCodeServer interface {
	CreateVerifyCode(context.Context, *CreateVerifyCodeRequest) (*CreateVerifyCodeReply, error)
	UpdateVerifyCode(context.Context, *UpdateVerifyCodeRequest) (*UpdateVerifyCodeReply, error)
	DeleteVerifyCode(context.Context, *DeleteVerifyCodeRequest) (*DeleteVerifyCodeReply, error)
	GetVerifyCode(context.Context, *GetVerifyCodeRequest) (*GetVerifyCodeReply, error)
	ListVerifyCode(context.Context, *ListVerifyCodeRequest) (*ListVerifyCodeReply, error)
	mustEmbedUnimplementedVerifyCodeServer()
}

// UnimplementedVerifyCodeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVerifyCodeServer struct{}

func (UnimplementedVerifyCodeServer) CreateVerifyCode(context.Context, *CreateVerifyCodeRequest) (*CreateVerifyCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVerifyCode not implemented")
}
func (UnimplementedVerifyCodeServer) UpdateVerifyCode(context.Context, *UpdateVerifyCodeRequest) (*UpdateVerifyCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVerifyCode not implemented")
}
func (UnimplementedVerifyCodeServer) DeleteVerifyCode(context.Context, *DeleteVerifyCodeRequest) (*DeleteVerifyCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVerifyCode not implemented")
}
func (UnimplementedVerifyCodeServer) GetVerifyCode(context.Context, *GetVerifyCodeRequest) (*GetVerifyCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVerifyCode not implemented")
}
func (UnimplementedVerifyCodeServer) ListVerifyCode(context.Context, *ListVerifyCodeRequest) (*ListVerifyCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVerifyCode not implemented")
}
func (UnimplementedVerifyCodeServer) mustEmbedUnimplementedVerifyCodeServer() {}
func (UnimplementedVerifyCodeServer) testEmbeddedByValue()                    {}

// UnsafeVerifyCodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerifyCodeServer will
// result in compilation errors.
type UnsafeVerifyCodeServer interface {
	mustEmbedUnimplementedVerifyCodeServer()
}

func RegisterVerifyCodeServer(s grpc.ServiceRegistrar, srv VerifyCodeServer) {
	// If the following call pancis, it indicates UnimplementedVerifyCodeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VerifyCode_ServiceDesc, srv)
}

func _VerifyCode_CreateVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyCodeServer).CreateVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifyCode_CreateVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyCodeServer).CreateVerifyCode(ctx, req.(*CreateVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifyCode_UpdateVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyCodeServer).UpdateVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifyCode_UpdateVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyCodeServer).UpdateVerifyCode(ctx, req.(*UpdateVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifyCode_DeleteVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyCodeServer).DeleteVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifyCode_DeleteVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyCodeServer).DeleteVerifyCode(ctx, req.(*DeleteVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifyCode_GetVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyCodeServer).GetVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifyCode_GetVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyCodeServer).GetVerifyCode(ctx, req.(*GetVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifyCode_ListVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyCodeServer).ListVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifyCode_ListVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyCodeServer).ListVerifyCode(ctx, req.(*ListVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VerifyCode_ServiceDesc is the grpc.ServiceDesc for VerifyCode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VerifyCode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "s.v1.VerifyCode",
	HandlerType: (*VerifyCodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVerifyCode",
			Handler:    _VerifyCode_CreateVerifyCode_Handler,
		},
		{
			MethodName: "UpdateVerifyCode",
			Handler:    _VerifyCode_UpdateVerifyCode_Handler,
		},
		{
			MethodName: "DeleteVerifyCode",
			Handler:    _VerifyCode_DeleteVerifyCode_Handler,
		},
		{
			MethodName: "GetVerifyCode",
			Handler:    _VerifyCode_GetVerifyCode_Handler,
		},
		{
			MethodName: "ListVerifyCode",
			Handler:    _VerifyCode_ListVerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/verify_code/verify_code.proto",
}
