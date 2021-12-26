// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package npool

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InternationalizationClient is the client API for Internationalization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternationalizationClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	AddLang(ctx context.Context, in *AddLangRequest, opts ...grpc.CallOption) (*AddLangResponse, error)
	UpdateLang(ctx context.Context, in *UpdateLangRequest, opts ...grpc.CallOption) (*UpdateLangResponse, error)
	GetLangs(ctx context.Context, in *GetLangsRequest, opts ...grpc.CallOption) (*GetLangsResponse, error)
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error)
	CreateMessages(ctx context.Context, in *CreateMessagesRequest, opts ...grpc.CallOption) (*CreateMessagesResponse, error)
	UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*UpdateMessageResponse, error)
	UpdateMessages(ctx context.Context, in *UpdateMessagesRequest, opts ...grpc.CallOption) (*UpdateMessagesResponse, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	GetMessageByMessageID(ctx context.Context, in *GetMessageByMessageIDRequest, opts ...grpc.CallOption) (*GetMessageByMessageIDResponse, error)
}

type internationalizationClient struct {
	cc grpc.ClientConnInterface
}

func NewInternationalizationClient(cc grpc.ClientConnInterface) InternationalizationClient {
	return &internationalizationClient{cc}
}

func (c *internationalizationClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/internationalization.v1.Internationalization/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internationalizationClient) AddLang(ctx context.Context, in *AddLangRequest, opts ...grpc.CallOption) (*AddLangResponse, error) {
	out := new(AddLangResponse)
	err := c.cc.Invoke(ctx, "/internationalization.v1.Internationalization/AddLang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internationalizationClient) UpdateLang(ctx context.Context, in *UpdateLangRequest, opts ...grpc.CallOption) (*UpdateLangResponse, error) {
	out := new(UpdateLangResponse)
	err := c.cc.Invoke(ctx, "/internationalization.v1.Internationalization/UpdateLang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internationalizationClient) GetLangs(ctx context.Context, in *GetLangsRequest, opts ...grpc.CallOption) (*GetLangsResponse, error) {
	out := new(GetLangsResponse)
	err := c.cc.Invoke(ctx, "/internationalization.v1.Internationalization/GetLangs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internationalizationClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error) {
	out := new(CreateMessageResponse)
	err := c.cc.Invoke(ctx, "/internationalization.v1.Internationalization/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internationalizationClient) CreateMessages(ctx context.Context, in *CreateMessagesRequest, opts ...grpc.CallOption) (*CreateMessagesResponse, error) {
	out := new(CreateMessagesResponse)
	err := c.cc.Invoke(ctx, "/internationalization.v1.Internationalization/CreateMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internationalizationClient) UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*UpdateMessageResponse, error) {
	out := new(UpdateMessageResponse)
	err := c.cc.Invoke(ctx, "/internationalization.v1.Internationalization/UpdateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internationalizationClient) UpdateMessages(ctx context.Context, in *UpdateMessagesRequest, opts ...grpc.CallOption) (*UpdateMessagesResponse, error) {
	out := new(UpdateMessagesResponse)
	err := c.cc.Invoke(ctx, "/internationalization.v1.Internationalization/UpdateMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internationalizationClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, "/internationalization.v1.Internationalization/GetMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internationalizationClient) GetMessageByMessageID(ctx context.Context, in *GetMessageByMessageIDRequest, opts ...grpc.CallOption) (*GetMessageByMessageIDResponse, error) {
	out := new(GetMessageByMessageIDResponse)
	err := c.cc.Invoke(ctx, "/internationalization.v1.Internationalization/GetMessageByMessageID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternationalizationServer is the server API for Internationalization service.
// All implementations must embed UnimplementedInternationalizationServer
// for forward compatibility
type InternationalizationServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	AddLang(context.Context, *AddLangRequest) (*AddLangResponse, error)
	UpdateLang(context.Context, *UpdateLangRequest) (*UpdateLangResponse, error)
	GetLangs(context.Context, *GetLangsRequest) (*GetLangsResponse, error)
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
	CreateMessages(context.Context, *CreateMessagesRequest) (*CreateMessagesResponse, error)
	UpdateMessage(context.Context, *UpdateMessageRequest) (*UpdateMessageResponse, error)
	UpdateMessages(context.Context, *UpdateMessagesRequest) (*UpdateMessagesResponse, error)
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	GetMessageByMessageID(context.Context, *GetMessageByMessageIDRequest) (*GetMessageByMessageIDResponse, error)
	mustEmbedUnimplementedInternationalizationServer()
}

// UnimplementedInternationalizationServer must be embedded to have forward compatible implementations.
type UnimplementedInternationalizationServer struct {
}

func (UnimplementedInternationalizationServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedInternationalizationServer) AddLang(context.Context, *AddLangRequest) (*AddLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLang not implemented")
}
func (UnimplementedInternationalizationServer) UpdateLang(context.Context, *UpdateLangRequest) (*UpdateLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLang not implemented")
}
func (UnimplementedInternationalizationServer) GetLangs(context.Context, *GetLangsRequest) (*GetLangsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLangs not implemented")
}
func (UnimplementedInternationalizationServer) CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedInternationalizationServer) CreateMessages(context.Context, *CreateMessagesRequest) (*CreateMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessages not implemented")
}
func (UnimplementedInternationalizationServer) UpdateMessage(context.Context, *UpdateMessageRequest) (*UpdateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedInternationalizationServer) UpdateMessages(context.Context, *UpdateMessagesRequest) (*UpdateMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessages not implemented")
}
func (UnimplementedInternationalizationServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedInternationalizationServer) GetMessageByMessageID(context.Context, *GetMessageByMessageIDRequest) (*GetMessageByMessageIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageByMessageID not implemented")
}
func (UnimplementedInternationalizationServer) mustEmbedUnimplementedInternationalizationServer() {}

// UnsafeInternationalizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternationalizationServer will
// result in compilation errors.
type UnsafeInternationalizationServer interface {
	mustEmbedUnimplementedInternationalizationServer()
}

func RegisterInternationalizationServer(s grpc.ServiceRegistrar, srv InternationalizationServer) {
	s.RegisterService(&Internationalization_ServiceDesc, srv)
}

func _Internationalization_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternationalizationServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internationalization.v1.Internationalization/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternationalizationServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internationalization_AddLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternationalizationServer).AddLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internationalization.v1.Internationalization/AddLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternationalizationServer).AddLang(ctx, req.(*AddLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internationalization_UpdateLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternationalizationServer).UpdateLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internationalization.v1.Internationalization/UpdateLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternationalizationServer).UpdateLang(ctx, req.(*UpdateLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internationalization_GetLangs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLangsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternationalizationServer).GetLangs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internationalization.v1.Internationalization/GetLangs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternationalizationServer).GetLangs(ctx, req.(*GetLangsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internationalization_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternationalizationServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internationalization.v1.Internationalization/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternationalizationServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internationalization_CreateMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternationalizationServer).CreateMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internationalization.v1.Internationalization/CreateMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternationalizationServer).CreateMessages(ctx, req.(*CreateMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internationalization_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternationalizationServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internationalization.v1.Internationalization/UpdateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternationalizationServer).UpdateMessage(ctx, req.(*UpdateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internationalization_UpdateMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternationalizationServer).UpdateMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internationalization.v1.Internationalization/UpdateMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternationalizationServer).UpdateMessages(ctx, req.(*UpdateMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internationalization_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternationalizationServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internationalization.v1.Internationalization/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternationalizationServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internationalization_GetMessageByMessageID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageByMessageIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternationalizationServer).GetMessageByMessageID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internationalization.v1.Internationalization/GetMessageByMessageID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternationalizationServer).GetMessageByMessageID(ctx, req.(*GetMessageByMessageIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Internationalization_ServiceDesc is the grpc.ServiceDesc for Internationalization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Internationalization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internationalization.v1.Internationalization",
	HandlerType: (*InternationalizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Internationalization_Version_Handler,
		},
		{
			MethodName: "AddLang",
			Handler:    _Internationalization_AddLang_Handler,
		},
		{
			MethodName: "UpdateLang",
			Handler:    _Internationalization_UpdateLang_Handler,
		},
		{
			MethodName: "GetLangs",
			Handler:    _Internationalization_GetLangs_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _Internationalization_CreateMessage_Handler,
		},
		{
			MethodName: "CreateMessages",
			Handler:    _Internationalization_CreateMessages_Handler,
		},
		{
			MethodName: "UpdateMessage",
			Handler:    _Internationalization_UpdateMessage_Handler,
		},
		{
			MethodName: "UpdateMessages",
			Handler:    _Internationalization_UpdateMessages_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _Internationalization_GetMessages_Handler,
		},
		{
			MethodName: "GetMessageByMessageID",
			Handler:    _Internationalization_GetMessageByMessageID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/internationalization.proto",
}
