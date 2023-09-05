// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: question.proto

package pb

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
	QuestionService_ValidQuestion_FullMethodName      = "/QuestionService/ValidQuestion"
	QuestionService_GetQuestionVoPage_FullMethodName  = "/QuestionService/GetQuestionVoPage"
	QuestionService_AddQuestion_FullMethodName        = "/QuestionService/AddQuestion"
	QuestionService_GetQuestionVoById_FullMethodName  = "/QuestionService/GetQuestionVoById"
	QuestionService_GetQuestionById_FullMethodName    = "/QuestionService/GetQuestionById"
	QuestionService_DeleteQuestionById_FullMethodName = "/QuestionService/DeleteQuestionById"
	QuestionService_UpdateQuestion_FullMethodName     = "/QuestionService/UpdateQuestion"
	QuestionService_GetQuestionVo_FullMethodName      = "/QuestionService/GetQuestionVo"
)

// QuestionServiceClient is the client API for QuestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionServiceClient interface {
	ValidQuestion(ctx context.Context, in *ValidQuestionRequest, opts ...grpc.CallOption) (*Empty, error)
	GetQuestionVoPage(ctx context.Context, in *GetQuestionVoPageRequest, opts ...grpc.CallOption) (*GetQuestionPageVoResponse, error)
	AddQuestion(ctx context.Context, in *QuestionAddRequest, opts ...grpc.CallOption) (*BoolResponse, error)
	GetQuestionVoById(ctx context.Context, in *QuestionIdRequest, opts ...grpc.CallOption) (*QuestionVo, error)
	GetQuestionById(ctx context.Context, in *QuestionIdRequest, opts ...grpc.CallOption) (*QuestionInfo, error)
	DeleteQuestionById(ctx context.Context, in *QuestionIdRequest, opts ...grpc.CallOption) (*BoolResponse, error)
	UpdateQuestion(ctx context.Context, in *QuestionInfo, opts ...grpc.CallOption) (*BoolResponse, error)
	GetQuestionVo(ctx context.Context, in *QuestionInfo, opts ...grpc.CallOption) (*QuestionVo, error)
}

type questionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionServiceClient(cc grpc.ClientConnInterface) QuestionServiceClient {
	return &questionServiceClient{cc}
}

func (c *questionServiceClient) ValidQuestion(ctx context.Context, in *ValidQuestionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, QuestionService_ValidQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) GetQuestionVoPage(ctx context.Context, in *GetQuestionVoPageRequest, opts ...grpc.CallOption) (*GetQuestionPageVoResponse, error) {
	out := new(GetQuestionPageVoResponse)
	err := c.cc.Invoke(ctx, QuestionService_GetQuestionVoPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) AddQuestion(ctx context.Context, in *QuestionAddRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, QuestionService_AddQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) GetQuestionVoById(ctx context.Context, in *QuestionIdRequest, opts ...grpc.CallOption) (*QuestionVo, error) {
	out := new(QuestionVo)
	err := c.cc.Invoke(ctx, QuestionService_GetQuestionVoById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) GetQuestionById(ctx context.Context, in *QuestionIdRequest, opts ...grpc.CallOption) (*QuestionInfo, error) {
	out := new(QuestionInfo)
	err := c.cc.Invoke(ctx, QuestionService_GetQuestionById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) DeleteQuestionById(ctx context.Context, in *QuestionIdRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, QuestionService_DeleteQuestionById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) UpdateQuestion(ctx context.Context, in *QuestionInfo, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, QuestionService_UpdateQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) GetQuestionVo(ctx context.Context, in *QuestionInfo, opts ...grpc.CallOption) (*QuestionVo, error) {
	out := new(QuestionVo)
	err := c.cc.Invoke(ctx, QuestionService_GetQuestionVo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionServiceServer is the server API for QuestionService service.
// All implementations must embed UnimplementedQuestionServiceServer
// for forward compatibility
type QuestionServiceServer interface {
	ValidQuestion(context.Context, *ValidQuestionRequest) (*Empty, error)
	GetQuestionVoPage(context.Context, *GetQuestionVoPageRequest) (*GetQuestionPageVoResponse, error)
	AddQuestion(context.Context, *QuestionAddRequest) (*BoolResponse, error)
	GetQuestionVoById(context.Context, *QuestionIdRequest) (*QuestionVo, error)
	GetQuestionById(context.Context, *QuestionIdRequest) (*QuestionInfo, error)
	DeleteQuestionById(context.Context, *QuestionIdRequest) (*BoolResponse, error)
	UpdateQuestion(context.Context, *QuestionInfo) (*BoolResponse, error)
	GetQuestionVo(context.Context, *QuestionInfo) (*QuestionVo, error)
	mustEmbedUnimplementedQuestionServiceServer()
}

// UnimplementedQuestionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuestionServiceServer struct {
}

func (UnimplementedQuestionServiceServer) ValidQuestion(context.Context, *ValidQuestionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) GetQuestionVoPage(context.Context, *GetQuestionVoPageRequest) (*GetQuestionPageVoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionVoPage not implemented")
}
func (UnimplementedQuestionServiceServer) AddQuestion(context.Context, *QuestionAddRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) GetQuestionVoById(context.Context, *QuestionIdRequest) (*QuestionVo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionVoById not implemented")
}
func (UnimplementedQuestionServiceServer) GetQuestionById(context.Context, *QuestionIdRequest) (*QuestionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionById not implemented")
}
func (UnimplementedQuestionServiceServer) DeleteQuestionById(context.Context, *QuestionIdRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuestionById not implemented")
}
func (UnimplementedQuestionServiceServer) UpdateQuestion(context.Context, *QuestionInfo) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) GetQuestionVo(context.Context, *QuestionInfo) (*QuestionVo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionVo not implemented")
}
func (UnimplementedQuestionServiceServer) mustEmbedUnimplementedQuestionServiceServer() {}

// UnsafeQuestionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionServiceServer will
// result in compilation errors.
type UnsafeQuestionServiceServer interface {
	mustEmbedUnimplementedQuestionServiceServer()
}

func RegisterQuestionServiceServer(s grpc.ServiceRegistrar, srv QuestionServiceServer) {
	s.RegisterService(&QuestionService_ServiceDesc, srv)
}

func _QuestionService_ValidQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).ValidQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_ValidQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).ValidQuestion(ctx, req.(*ValidQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_GetQuestionVoPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionVoPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetQuestionVoPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetQuestionVoPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetQuestionVoPage(ctx, req.(*GetQuestionVoPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_AddQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).AddQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_AddQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).AddQuestion(ctx, req.(*QuestionAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_GetQuestionVoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetQuestionVoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetQuestionVoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetQuestionVoById(ctx, req.(*QuestionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_GetQuestionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetQuestionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetQuestionById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetQuestionById(ctx, req.(*QuestionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_DeleteQuestionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).DeleteQuestionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_DeleteQuestionById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).DeleteQuestionById(ctx, req.(*QuestionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_UpdateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).UpdateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_UpdateQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).UpdateQuestion(ctx, req.(*QuestionInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_GetQuestionVo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetQuestionVo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetQuestionVo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetQuestionVo(ctx, req.(*QuestionInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// QuestionService_ServiceDesc is the grpc.ServiceDesc for QuestionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuestionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "QuestionService",
	HandlerType: (*QuestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidQuestion",
			Handler:    _QuestionService_ValidQuestion_Handler,
		},
		{
			MethodName: "GetQuestionVoPage",
			Handler:    _QuestionService_GetQuestionVoPage_Handler,
		},
		{
			MethodName: "AddQuestion",
			Handler:    _QuestionService_AddQuestion_Handler,
		},
		{
			MethodName: "GetQuestionVoById",
			Handler:    _QuestionService_GetQuestionVoById_Handler,
		},
		{
			MethodName: "GetQuestionById",
			Handler:    _QuestionService_GetQuestionById_Handler,
		},
		{
			MethodName: "DeleteQuestionById",
			Handler:    _QuestionService_DeleteQuestionById_Handler,
		},
		{
			MethodName: "UpdateQuestion",
			Handler:    _QuestionService_UpdateQuestion_Handler,
		},
		{
			MethodName: "GetQuestionVo",
			Handler:    _QuestionService_GetQuestionVo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "question.proto",
}
