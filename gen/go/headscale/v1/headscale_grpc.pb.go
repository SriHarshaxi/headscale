// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: headscale/v1/headscale.proto

package v1

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

// HeadscaleServiceClient is the client API for HeadscaleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeadscaleServiceClient interface {
	// --- Namespace start ---
	GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error)
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error)
	RenameNamespace(ctx context.Context, in *RenameNamespaceRequest, opts ...grpc.CallOption) (*RenameNamespaceResponse, error)
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error)
	ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
	// --- PreAuthKeys start ---
	CreatePreAuthKey(ctx context.Context, in *CreatePreAuthKeyRequest, opts ...grpc.CallOption) (*CreatePreAuthKeyResponse, error)
	ExpirePreAuthKey(ctx context.Context, in *ExpirePreAuthKeyRequest, opts ...grpc.CallOption) (*ExpirePreAuthKeyResponse, error)
	ListPreAuthKeys(ctx context.Context, in *ListPreAuthKeysRequest, opts ...grpc.CallOption) (*ListPreAuthKeysResponse, error)
	// --- Machine start ---
	DebugCreateMachine(ctx context.Context, in *DebugCreateMachineRequest, opts ...grpc.CallOption) (*DebugCreateMachineResponse, error)
	GetMachine(ctx context.Context, in *GetMachineRequest, opts ...grpc.CallOption) (*GetMachineResponse, error)
	RegisterMachine(ctx context.Context, in *RegisterMachineRequest, opts ...grpc.CallOption) (*RegisterMachineResponse, error)
	DeleteMachine(ctx context.Context, in *DeleteMachineRequest, opts ...grpc.CallOption) (*DeleteMachineResponse, error)
	ExpireMachine(ctx context.Context, in *ExpireMachineRequest, opts ...grpc.CallOption) (*ExpireMachineResponse, error)
	ListMachines(ctx context.Context, in *ListMachinesRequest, opts ...grpc.CallOption) (*ListMachinesResponse, error)
	// --- Route start ---
	GetMachineRoute(ctx context.Context, in *GetMachineRouteRequest, opts ...grpc.CallOption) (*GetMachineRouteResponse, error)
	EnableMachineRoutes(ctx context.Context, in *EnableMachineRoutesRequest, opts ...grpc.CallOption) (*EnableMachineRoutesResponse, error)
	// --- ApiKeys start ---
	CreateApiKey(ctx context.Context, in *CreateApiKeyRequest, opts ...grpc.CallOption) (*CreateApiKeyResponse, error)
	ExpireApiKey(ctx context.Context, in *ExpireApiKeyRequest, opts ...grpc.CallOption) (*ExpireApiKeyResponse, error)
	ListApiKeys(ctx context.Context, in *ListApiKeysRequest, opts ...grpc.CallOption) (*ListApiKeysResponse, error)
}

type headscaleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeadscaleServiceClient(cc grpc.ClientConnInterface) HeadscaleServiceClient {
	return &headscaleServiceClient{cc}
}

func (c *headscaleServiceClient) GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error) {
	out := new(GetNamespaceResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/GetNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error) {
	out := new(CreateNamespaceResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/CreateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) RenameNamespace(ctx context.Context, in *RenameNamespaceRequest, opts ...grpc.CallOption) (*RenameNamespaceResponse, error) {
	out := new(RenameNamespaceResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/RenameNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error) {
	out := new(DeleteNamespaceResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/DeleteNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/ListNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) CreatePreAuthKey(ctx context.Context, in *CreatePreAuthKeyRequest, opts ...grpc.CallOption) (*CreatePreAuthKeyResponse, error) {
	out := new(CreatePreAuthKeyResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/CreatePreAuthKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) ExpirePreAuthKey(ctx context.Context, in *ExpirePreAuthKeyRequest, opts ...grpc.CallOption) (*ExpirePreAuthKeyResponse, error) {
	out := new(ExpirePreAuthKeyResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/ExpirePreAuthKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) ListPreAuthKeys(ctx context.Context, in *ListPreAuthKeysRequest, opts ...grpc.CallOption) (*ListPreAuthKeysResponse, error) {
	out := new(ListPreAuthKeysResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/ListPreAuthKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) DebugCreateMachine(ctx context.Context, in *DebugCreateMachineRequest, opts ...grpc.CallOption) (*DebugCreateMachineResponse, error) {
	out := new(DebugCreateMachineResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/DebugCreateMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) GetMachine(ctx context.Context, in *GetMachineRequest, opts ...grpc.CallOption) (*GetMachineResponse, error) {
	out := new(GetMachineResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/GetMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) RegisterMachine(ctx context.Context, in *RegisterMachineRequest, opts ...grpc.CallOption) (*RegisterMachineResponse, error) {
	out := new(RegisterMachineResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/RegisterMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) DeleteMachine(ctx context.Context, in *DeleteMachineRequest, opts ...grpc.CallOption) (*DeleteMachineResponse, error) {
	out := new(DeleteMachineResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/DeleteMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) ExpireMachine(ctx context.Context, in *ExpireMachineRequest, opts ...grpc.CallOption) (*ExpireMachineResponse, error) {
	out := new(ExpireMachineResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/ExpireMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) ListMachines(ctx context.Context, in *ListMachinesRequest, opts ...grpc.CallOption) (*ListMachinesResponse, error) {
	out := new(ListMachinesResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/ListMachines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) GetMachineRoute(ctx context.Context, in *GetMachineRouteRequest, opts ...grpc.CallOption) (*GetMachineRouteResponse, error) {
	out := new(GetMachineRouteResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/GetMachineRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) EnableMachineRoutes(ctx context.Context, in *EnableMachineRoutesRequest, opts ...grpc.CallOption) (*EnableMachineRoutesResponse, error) {
	out := new(EnableMachineRoutesResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/EnableMachineRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) CreateApiKey(ctx context.Context, in *CreateApiKeyRequest, opts ...grpc.CallOption) (*CreateApiKeyResponse, error) {
	out := new(CreateApiKeyResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/CreateApiKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) ExpireApiKey(ctx context.Context, in *ExpireApiKeyRequest, opts ...grpc.CallOption) (*ExpireApiKeyResponse, error) {
	out := new(ExpireApiKeyResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/ExpireApiKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headscaleServiceClient) ListApiKeys(ctx context.Context, in *ListApiKeysRequest, opts ...grpc.CallOption) (*ListApiKeysResponse, error) {
	out := new(ListApiKeysResponse)
	err := c.cc.Invoke(ctx, "/headscale.v1.HeadscaleService/ListApiKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeadscaleServiceServer is the server API for HeadscaleService service.
// All implementations must embed UnimplementedHeadscaleServiceServer
// for forward compatibility
type HeadscaleServiceServer interface {
	// --- Namespace start ---
	GetNamespace(context.Context, *GetNamespaceRequest) (*GetNamespaceResponse, error)
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error)
	RenameNamespace(context.Context, *RenameNamespaceRequest) (*RenameNamespaceResponse, error)
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error)
	ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
	// --- PreAuthKeys start ---
	CreatePreAuthKey(context.Context, *CreatePreAuthKeyRequest) (*CreatePreAuthKeyResponse, error)
	ExpirePreAuthKey(context.Context, *ExpirePreAuthKeyRequest) (*ExpirePreAuthKeyResponse, error)
	ListPreAuthKeys(context.Context, *ListPreAuthKeysRequest) (*ListPreAuthKeysResponse, error)
	// --- Machine start ---
	DebugCreateMachine(context.Context, *DebugCreateMachineRequest) (*DebugCreateMachineResponse, error)
	GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error)
	RegisterMachine(context.Context, *RegisterMachineRequest) (*RegisterMachineResponse, error)
	DeleteMachine(context.Context, *DeleteMachineRequest) (*DeleteMachineResponse, error)
	ExpireMachine(context.Context, *ExpireMachineRequest) (*ExpireMachineResponse, error)
	ListMachines(context.Context, *ListMachinesRequest) (*ListMachinesResponse, error)
	// --- Route start ---
	GetMachineRoute(context.Context, *GetMachineRouteRequest) (*GetMachineRouteResponse, error)
	EnableMachineRoutes(context.Context, *EnableMachineRoutesRequest) (*EnableMachineRoutesResponse, error)
	// --- ApiKeys start ---
	CreateApiKey(context.Context, *CreateApiKeyRequest) (*CreateApiKeyResponse, error)
	ExpireApiKey(context.Context, *ExpireApiKeyRequest) (*ExpireApiKeyResponse, error)
	ListApiKeys(context.Context, *ListApiKeysRequest) (*ListApiKeysResponse, error)
	mustEmbedUnimplementedHeadscaleServiceServer()
}

// UnimplementedHeadscaleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHeadscaleServiceServer struct {
}

func (UnimplementedHeadscaleServiceServer) GetNamespace(context.Context, *GetNamespaceRequest) (*GetNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespace not implemented")
}
func (UnimplementedHeadscaleServiceServer) CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamespace not implemented")
}
func (UnimplementedHeadscaleServiceServer) RenameNamespace(context.Context, *RenameNamespaceRequest) (*RenameNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameNamespace not implemented")
}
func (UnimplementedHeadscaleServiceServer) DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedHeadscaleServiceServer) ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespaces not implemented")
}
func (UnimplementedHeadscaleServiceServer) CreatePreAuthKey(context.Context, *CreatePreAuthKeyRequest) (*CreatePreAuthKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePreAuthKey not implemented")
}
func (UnimplementedHeadscaleServiceServer) ExpirePreAuthKey(context.Context, *ExpirePreAuthKeyRequest) (*ExpirePreAuthKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpirePreAuthKey not implemented")
}
func (UnimplementedHeadscaleServiceServer) ListPreAuthKeys(context.Context, *ListPreAuthKeysRequest) (*ListPreAuthKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPreAuthKeys not implemented")
}
func (UnimplementedHeadscaleServiceServer) DebugCreateMachine(context.Context, *DebugCreateMachineRequest) (*DebugCreateMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DebugCreateMachine not implemented")
}
func (UnimplementedHeadscaleServiceServer) GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachine not implemented")
}
func (UnimplementedHeadscaleServiceServer) RegisterMachine(context.Context, *RegisterMachineRequest) (*RegisterMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMachine not implemented")
}
func (UnimplementedHeadscaleServiceServer) DeleteMachine(context.Context, *DeleteMachineRequest) (*DeleteMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMachine not implemented")
}
func (UnimplementedHeadscaleServiceServer) ExpireMachine(context.Context, *ExpireMachineRequest) (*ExpireMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpireMachine not implemented")
}
func (UnimplementedHeadscaleServiceServer) ListMachines(context.Context, *ListMachinesRequest) (*ListMachinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMachines not implemented")
}
func (UnimplementedHeadscaleServiceServer) GetMachineRoute(context.Context, *GetMachineRouteRequest) (*GetMachineRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachineRoute not implemented")
}
func (UnimplementedHeadscaleServiceServer) EnableMachineRoutes(context.Context, *EnableMachineRoutesRequest) (*EnableMachineRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableMachineRoutes not implemented")
}
func (UnimplementedHeadscaleServiceServer) CreateApiKey(context.Context, *CreateApiKeyRequest) (*CreateApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApiKey not implemented")
}
func (UnimplementedHeadscaleServiceServer) ExpireApiKey(context.Context, *ExpireApiKeyRequest) (*ExpireApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpireApiKey not implemented")
}
func (UnimplementedHeadscaleServiceServer) ListApiKeys(context.Context, *ListApiKeysRequest) (*ListApiKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApiKeys not implemented")
}
func (UnimplementedHeadscaleServiceServer) mustEmbedUnimplementedHeadscaleServiceServer() {}

// UnsafeHeadscaleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeadscaleServiceServer will
// result in compilation errors.
type UnsafeHeadscaleServiceServer interface {
	mustEmbedUnimplementedHeadscaleServiceServer()
}

func RegisterHeadscaleServiceServer(s grpc.ServiceRegistrar, srv HeadscaleServiceServer) {
	s.RegisterService(&HeadscaleService_ServiceDesc, srv)
}

func _HeadscaleService_GetNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).GetNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/GetNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).GetNamespace(ctx, req.(*GetNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/CreateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).CreateNamespace(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_RenameNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).RenameNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/RenameNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).RenameNamespace(ctx, req.(*RenameNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/DeleteNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).DeleteNamespace(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/ListNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).ListNamespaces(ctx, req.(*ListNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_CreatePreAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePreAuthKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).CreatePreAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/CreatePreAuthKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).CreatePreAuthKey(ctx, req.(*CreatePreAuthKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_ExpirePreAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpirePreAuthKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).ExpirePreAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/ExpirePreAuthKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).ExpirePreAuthKey(ctx, req.(*ExpirePreAuthKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_ListPreAuthKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPreAuthKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).ListPreAuthKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/ListPreAuthKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).ListPreAuthKeys(ctx, req.(*ListPreAuthKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_DebugCreateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebugCreateMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).DebugCreateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/DebugCreateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).DebugCreateMachine(ctx, req.(*DebugCreateMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_GetMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).GetMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/GetMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).GetMachine(ctx, req.(*GetMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_RegisterMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).RegisterMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/RegisterMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).RegisterMachine(ctx, req.(*RegisterMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_DeleteMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).DeleteMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/DeleteMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).DeleteMachine(ctx, req.(*DeleteMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_ExpireMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpireMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).ExpireMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/ExpireMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).ExpireMachine(ctx, req.(*ExpireMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_ListMachines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMachinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).ListMachines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/ListMachines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).ListMachines(ctx, req.(*ListMachinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_GetMachineRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMachineRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).GetMachineRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/GetMachineRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).GetMachineRoute(ctx, req.(*GetMachineRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_EnableMachineRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableMachineRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).EnableMachineRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/EnableMachineRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).EnableMachineRoutes(ctx, req.(*EnableMachineRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_CreateApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).CreateApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/CreateApiKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).CreateApiKey(ctx, req.(*CreateApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_ExpireApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpireApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).ExpireApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/ExpireApiKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).ExpireApiKey(ctx, req.(*ExpireApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadscaleService_ListApiKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApiKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadscaleServiceServer).ListApiKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/headscale.v1.HeadscaleService/ListApiKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadscaleServiceServer).ListApiKeys(ctx, req.(*ListApiKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HeadscaleService_ServiceDesc is the grpc.ServiceDesc for HeadscaleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeadscaleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "headscale.v1.HeadscaleService",
	HandlerType: (*HeadscaleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNamespace",
			Handler:    _HeadscaleService_GetNamespace_Handler,
		},
		{
			MethodName: "CreateNamespace",
			Handler:    _HeadscaleService_CreateNamespace_Handler,
		},
		{
			MethodName: "RenameNamespace",
			Handler:    _HeadscaleService_RenameNamespace_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _HeadscaleService_DeleteNamespace_Handler,
		},
		{
			MethodName: "ListNamespaces",
			Handler:    _HeadscaleService_ListNamespaces_Handler,
		},
		{
			MethodName: "CreatePreAuthKey",
			Handler:    _HeadscaleService_CreatePreAuthKey_Handler,
		},
		{
			MethodName: "ExpirePreAuthKey",
			Handler:    _HeadscaleService_ExpirePreAuthKey_Handler,
		},
		{
			MethodName: "ListPreAuthKeys",
			Handler:    _HeadscaleService_ListPreAuthKeys_Handler,
		},
		{
			MethodName: "DebugCreateMachine",
			Handler:    _HeadscaleService_DebugCreateMachine_Handler,
		},
		{
			MethodName: "GetMachine",
			Handler:    _HeadscaleService_GetMachine_Handler,
		},
		{
			MethodName: "RegisterMachine",
			Handler:    _HeadscaleService_RegisterMachine_Handler,
		},
		{
			MethodName: "DeleteMachine",
			Handler:    _HeadscaleService_DeleteMachine_Handler,
		},
		{
			MethodName: "ExpireMachine",
			Handler:    _HeadscaleService_ExpireMachine_Handler,
		},
		{
			MethodName: "ListMachines",
			Handler:    _HeadscaleService_ListMachines_Handler,
		},
		{
			MethodName: "GetMachineRoute",
			Handler:    _HeadscaleService_GetMachineRoute_Handler,
		},
		{
			MethodName: "EnableMachineRoutes",
			Handler:    _HeadscaleService_EnableMachineRoutes_Handler,
		},
		{
			MethodName: "CreateApiKey",
			Handler:    _HeadscaleService_CreateApiKey_Handler,
		},
		{
			MethodName: "ExpireApiKey",
			Handler:    _HeadscaleService_ExpireApiKey_Handler,
		},
		{
			MethodName: "ListApiKeys",
			Handler:    _HeadscaleService_ListApiKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "headscale/v1/headscale.proto",
}
