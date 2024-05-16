// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: opkit/domain/query.proto

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
	Query_Params_FullMethodName             = "/opkit.domain.Query/Params"
	Query_Domain_FullMethodName             = "/opkit.domain.Query/Domain"
	Query_DomainAll_FullMethodName          = "/opkit.domain.Query/DomainAll"
	Query_ListDomainOpkit_FullMethodName    = "/opkit.domain.Query/ListDomainOpkit"
	Query_ListDomainEvm_FullMethodName      = "/opkit.domain.Query/ListDomainEvm"
	Query_ListDomainByString_FullMethodName = "/opkit.domain.Query/ListDomainByString"
	Query_Info_FullMethodName               = "/opkit.domain.Query/Info"
	Query_CheckOwner_FullMethodName         = "/opkit.domain.Query/CheckOwner"
	Query_Reward_FullMethodName             = "/opkit.domain.Query/Reward"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Domain items.
	Domain(ctx context.Context, in *QueryGetDomainRequest, opts ...grpc.CallOption) (*QueryGetDomainResponse, error)
	DomainAll(ctx context.Context, in *QueryAllDomainRequest, opts ...grpc.CallOption) (*QueryAllDomainResponse, error)
	// Queries a list of ListDomainOpkit items.
	ListDomainOpkit(ctx context.Context, in *QueryListDomainOpkitRequest, opts ...grpc.CallOption) (*QueryListDomainOpkitResponse, error)
	// Queries a list of ListDomainEvm items.
	ListDomainEvm(ctx context.Context, in *QueryListDomainEvmRequest, opts ...grpc.CallOption) (*QueryListDomainEvmResponse, error)
	// Queries a list of ListDomainByString items.
	ListDomainByString(ctx context.Context, in *QueryListDomainByStringRequest, opts ...grpc.CallOption) (*QueryListDomainByStringResponse, error)
	// Queries a list of Info items.
	Info(ctx context.Context, in *QueryInfoRequest, opts ...grpc.CallOption) (*QueryInfoResponse, error)
	// Queries a list of CheckOwner items.
	CheckOwner(ctx context.Context, in *QueryCheckOwnerRequest, opts ...grpc.CallOption) (*QueryCheckOwnerResponse, error)
	// Queries a list of Reward items.
	Reward(ctx context.Context, in *QueryRewardRequest, opts ...grpc.CallOption) (*QueryRewardResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Domain(ctx context.Context, in *QueryGetDomainRequest, opts ...grpc.CallOption) (*QueryGetDomainResponse, error) {
	out := new(QueryGetDomainResponse)
	err := c.cc.Invoke(ctx, Query_Domain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DomainAll(ctx context.Context, in *QueryAllDomainRequest, opts ...grpc.CallOption) (*QueryAllDomainResponse, error) {
	out := new(QueryAllDomainResponse)
	err := c.cc.Invoke(ctx, Query_DomainAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListDomainOpkit(ctx context.Context, in *QueryListDomainOpkitRequest, opts ...grpc.CallOption) (*QueryListDomainOpkitResponse, error) {
	out := new(QueryListDomainOpkitResponse)
	err := c.cc.Invoke(ctx, Query_ListDomainOpkit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListDomainEvm(ctx context.Context, in *QueryListDomainEvmRequest, opts ...grpc.CallOption) (*QueryListDomainEvmResponse, error) {
	out := new(QueryListDomainEvmResponse)
	err := c.cc.Invoke(ctx, Query_ListDomainEvm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListDomainByString(ctx context.Context, in *QueryListDomainByStringRequest, opts ...grpc.CallOption) (*QueryListDomainByStringResponse, error) {
	out := new(QueryListDomainByStringResponse)
	err := c.cc.Invoke(ctx, Query_ListDomainByString_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Info(ctx context.Context, in *QueryInfoRequest, opts ...grpc.CallOption) (*QueryInfoResponse, error) {
	out := new(QueryInfoResponse)
	err := c.cc.Invoke(ctx, Query_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CheckOwner(ctx context.Context, in *QueryCheckOwnerRequest, opts ...grpc.CallOption) (*QueryCheckOwnerResponse, error) {
	out := new(QueryCheckOwnerResponse)
	err := c.cc.Invoke(ctx, Query_CheckOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Reward(ctx context.Context, in *QueryRewardRequest, opts ...grpc.CallOption) (*QueryRewardResponse, error) {
	out := new(QueryRewardResponse)
	err := c.cc.Invoke(ctx, Query_Reward_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Domain items.
	Domain(context.Context, *QueryGetDomainRequest) (*QueryGetDomainResponse, error)
	DomainAll(context.Context, *QueryAllDomainRequest) (*QueryAllDomainResponse, error)
	// Queries a list of ListDomainOpkit items.
	ListDomainOpkit(context.Context, *QueryListDomainOpkitRequest) (*QueryListDomainOpkitResponse, error)
	// Queries a list of ListDomainEvm items.
	ListDomainEvm(context.Context, *QueryListDomainEvmRequest) (*QueryListDomainEvmResponse, error)
	// Queries a list of ListDomainByString items.
	ListDomainByString(context.Context, *QueryListDomainByStringRequest) (*QueryListDomainByStringResponse, error)
	// Queries a list of Info items.
	Info(context.Context, *QueryInfoRequest) (*QueryInfoResponse, error)
	// Queries a list of CheckOwner items.
	CheckOwner(context.Context, *QueryCheckOwnerRequest) (*QueryCheckOwnerResponse, error)
	// Queries a list of Reward items.
	Reward(context.Context, *QueryRewardRequest) (*QueryRewardResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Domain(context.Context, *QueryGetDomainRequest) (*QueryGetDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Domain not implemented")
}
func (UnimplementedQueryServer) DomainAll(context.Context, *QueryAllDomainRequest) (*QueryAllDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DomainAll not implemented")
}
func (UnimplementedQueryServer) ListDomainOpkit(context.Context, *QueryListDomainOpkitRequest) (*QueryListDomainOpkitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomainOpkit not implemented")
}
func (UnimplementedQueryServer) ListDomainEvm(context.Context, *QueryListDomainEvmRequest) (*QueryListDomainEvmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomainEvm not implemented")
}
func (UnimplementedQueryServer) ListDomainByString(context.Context, *QueryListDomainByStringRequest) (*QueryListDomainByStringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomainByString not implemented")
}
func (UnimplementedQueryServer) Info(context.Context, *QueryInfoRequest) (*QueryInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedQueryServer) CheckOwner(context.Context, *QueryCheckOwnerRequest) (*QueryCheckOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOwner not implemented")
}
func (UnimplementedQueryServer) Reward(context.Context, *QueryRewardRequest) (*QueryRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reward not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Domain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Domain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Domain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Domain(ctx, req.(*QueryGetDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DomainAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DomainAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DomainAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DomainAll(ctx, req.(*QueryAllDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListDomainOpkit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListDomainOpkitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListDomainOpkit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListDomainOpkit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListDomainOpkit(ctx, req.(*QueryListDomainOpkitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListDomainEvm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListDomainEvmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListDomainEvm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListDomainEvm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListDomainEvm(ctx, req.(*QueryListDomainEvmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListDomainByString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListDomainByStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListDomainByString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListDomainByString_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListDomainByString(ctx, req.(*QueryListDomainByStringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Info(ctx, req.(*QueryInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CheckOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCheckOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CheckOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CheckOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CheckOwner(ctx, req.(*QueryCheckOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Reward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Reward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Reward_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Reward(ctx, req.(*QueryRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "opkit.domain.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Domain",
			Handler:    _Query_Domain_Handler,
		},
		{
			MethodName: "DomainAll",
			Handler:    _Query_DomainAll_Handler,
		},
		{
			MethodName: "ListDomainOpkit",
			Handler:    _Query_ListDomainOpkit_Handler,
		},
		{
			MethodName: "ListDomainEvm",
			Handler:    _Query_ListDomainEvm_Handler,
		},
		{
			MethodName: "ListDomainByString",
			Handler:    _Query_ListDomainByString_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Query_Info_Handler,
		},
		{
			MethodName: "CheckOwner",
			Handler:    _Query_CheckOwner_Handler,
		},
		{
			MethodName: "Reward",
			Handler:    _Query_Reward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opkit/domain/query.proto",
}
