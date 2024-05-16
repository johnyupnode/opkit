// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: opkit/domain/tx.proto

package domain

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_UpdateParams_FullMethodName = "/opkit.domain.Msg/UpdateParams"
	Msg_CreateDomain_FullMethodName = "/opkit.domain.Msg/CreateDomain"
	Msg_UpdateDomain_FullMethodName = "/opkit.domain.Msg/UpdateDomain"
	Msg_DeleteDomain_FullMethodName = "/opkit.domain.Msg/DeleteDomain"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateDomain(ctx context.Context, in *MsgCreateDomain, opts ...grpc.CallOption) (*MsgCreateDomainResponse, error)
	UpdateDomain(ctx context.Context, in *MsgUpdateDomain, opts ...grpc.CallOption) (*MsgUpdateDomainResponse, error)
	DeleteDomain(ctx context.Context, in *MsgDeleteDomain, opts ...grpc.CallOption) (*MsgDeleteDomainResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateDomain(ctx context.Context, in *MsgCreateDomain, opts ...grpc.CallOption) (*MsgCreateDomainResponse, error) {
	out := new(MsgCreateDomainResponse)
	err := c.cc.Invoke(ctx, Msg_CreateDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateDomain(ctx context.Context, in *MsgUpdateDomain, opts ...grpc.CallOption) (*MsgUpdateDomainResponse, error) {
	out := new(MsgUpdateDomainResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteDomain(ctx context.Context, in *MsgDeleteDomain, opts ...grpc.CallOption) (*MsgDeleteDomainResponse, error) {
	out := new(MsgDeleteDomainResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateDomain(context.Context, *MsgCreateDomain) (*MsgCreateDomainResponse, error)
	UpdateDomain(context.Context, *MsgUpdateDomain) (*MsgUpdateDomainResponse, error)
	DeleteDomain(context.Context, *MsgDeleteDomain) (*MsgDeleteDomainResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateDomain(context.Context, *MsgCreateDomain) (*MsgCreateDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomain not implemented")
}
func (UnimplementedMsgServer) UpdateDomain(context.Context, *MsgUpdateDomain) (*MsgUpdateDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomain not implemented")
}
func (UnimplementedMsgServer) DeleteDomain(context.Context, *MsgDeleteDomain) (*MsgDeleteDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDomain not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateDomain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateDomain(ctx, req.(*MsgCreateDomain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateDomain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateDomain(ctx, req.(*MsgUpdateDomain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteDomain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteDomain(ctx, req.(*MsgDeleteDomain))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "opkit.domain.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateDomain",
			Handler:    _Msg_CreateDomain_Handler,
		},
		{
			MethodName: "UpdateDomain",
			Handler:    _Msg_UpdateDomain_Handler,
		},
		{
			MethodName: "DeleteDomain",
			Handler:    _Msg_DeleteDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opkit/domain/tx.proto",
}
