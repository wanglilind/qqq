// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/proto/identity.proto

package identity

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

// IdentityServiceClient is the client API for IdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityServiceClient interface {
	// 注册新用户身份
	RegisterIdentity(ctx context.Context, in *RegisterIdentityRequest, opts ...grpc.CallOption) (*RegisterIdentityResponse, error)
	// 验证用户身份
	VerifyIdentity(ctx context.Context, in *VerifyIdentityRequest, opts ...grpc.CallOption) (*VerifyIdentityResponse, error)
	// 获取身份状态
	GetIdentityStatus(ctx context.Context, in *GetIdentityStatusRequest, opts ...grpc.CallOption) (*GetIdentityStatusResponse, error)
	// 更新生物特征数据
	UpdateBiometricData(ctx context.Context, in *UpdateBiometricDataRequest, opts ...grpc.CallOption) (*UpdateBiometricDataResponse, error)
}

type identityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityServiceClient(cc grpc.ClientConnInterface) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) RegisterIdentity(ctx context.Context, in *RegisterIdentityRequest, opts ...grpc.CallOption) (*RegisterIdentityResponse, error) {
	out := new(RegisterIdentityResponse)
	err := c.cc.Invoke(ctx, "/identity.IdentityService/RegisterIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) VerifyIdentity(ctx context.Context, in *VerifyIdentityRequest, opts ...grpc.CallOption) (*VerifyIdentityResponse, error) {
	out := new(VerifyIdentityResponse)
	err := c.cc.Invoke(ctx, "/identity.IdentityService/VerifyIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) GetIdentityStatus(ctx context.Context, in *GetIdentityStatusRequest, opts ...grpc.CallOption) (*GetIdentityStatusResponse, error) {
	out := new(GetIdentityStatusResponse)
	err := c.cc.Invoke(ctx, "/identity.IdentityService/GetIdentityStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) UpdateBiometricData(ctx context.Context, in *UpdateBiometricDataRequest, opts ...grpc.CallOption) (*UpdateBiometricDataResponse, error) {
	out := new(UpdateBiometricDataResponse)
	err := c.cc.Invoke(ctx, "/identity.IdentityService/UpdateBiometricData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServiceServer is the server API for IdentityService service.
// All implementations must embed UnimplementedIdentityServiceServer
// for forward compatibility
type IdentityServiceServer interface {
	// 注册新用户身份
	RegisterIdentity(context.Context, *RegisterIdentityRequest) (*RegisterIdentityResponse, error)
	// 验证用户身份
	VerifyIdentity(context.Context, *VerifyIdentityRequest) (*VerifyIdentityResponse, error)
	// 获取身份状态
	GetIdentityStatus(context.Context, *GetIdentityStatusRequest) (*GetIdentityStatusResponse, error)
	// 更新生物特征数据
	UpdateBiometricData(context.Context, *UpdateBiometricDataRequest) (*UpdateBiometricDataResponse, error)
	mustEmbedUnimplementedIdentityServiceServer()
}

// UnimplementedIdentityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityServiceServer struct {
}

func (UnimplementedIdentityServiceServer) RegisterIdentity(context.Context, *RegisterIdentityRequest) (*RegisterIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterIdentity not implemented")
}
func (UnimplementedIdentityServiceServer) VerifyIdentity(context.Context, *VerifyIdentityRequest) (*VerifyIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyIdentity not implemented")
}
func (UnimplementedIdentityServiceServer) GetIdentityStatus(context.Context, *GetIdentityStatusRequest) (*GetIdentityStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentityStatus not implemented")
}
func (UnimplementedIdentityServiceServer) UpdateBiometricData(context.Context, *UpdateBiometricDataRequest) (*UpdateBiometricDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBiometricData not implemented")
}
func (UnimplementedIdentityServiceServer) mustEmbedUnimplementedIdentityServiceServer() {}

// UnsafeIdentityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServiceServer will
// result in compilation errors.
type UnsafeIdentityServiceServer interface {
	mustEmbedUnimplementedIdentityServiceServer()
}

func RegisterIdentityServiceServer(s grpc.ServiceRegistrar, srv IdentityServiceServer) {
	s.RegisterService(&IdentityService_ServiceDesc, srv)
}

func _IdentityService_RegisterIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).RegisterIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.IdentityService/RegisterIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).RegisterIdentity(ctx, req.(*RegisterIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_VerifyIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).VerifyIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.IdentityService/VerifyIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).VerifyIdentity(ctx, req.(*VerifyIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_GetIdentityStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdentityStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).GetIdentityStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.IdentityService/GetIdentityStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).GetIdentityStatus(ctx, req.(*GetIdentityStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_UpdateBiometricData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBiometricDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).UpdateBiometricData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.IdentityService/UpdateBiometricData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).UpdateBiometricData(ctx, req.(*UpdateBiometricDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityService_ServiceDesc is the grpc.ServiceDesc for IdentityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "identity.IdentityService",
	HandlerType: (*IdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterIdentity",
			Handler:    _IdentityService_RegisterIdentity_Handler,
		},
		{
			MethodName: "VerifyIdentity",
			Handler:    _IdentityService_VerifyIdentity_Handler,
		},
		{
			MethodName: "GetIdentityStatus",
			Handler:    _IdentityService_GetIdentityStatus_Handler,
		},
		{
			MethodName: "UpdateBiometricData",
			Handler:    _IdentityService_UpdateBiometricData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/identity.proto",
}
