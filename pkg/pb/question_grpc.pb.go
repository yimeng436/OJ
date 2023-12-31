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
	QuestionService_GetById_FullMethodName            = "/QuestionService/GetById"
	QuestionService_ValidQuestion_FullMethodName      = "/QuestionService/ValidQuestion"
	QuestionService_ListQuestionPage_FullMethodName   = "/QuestionService/ListQuestionPage"
	QuestionService_AddQuestion_FullMethodName        = "/QuestionService/AddQuestion"
	QuestionService_GetQuestionVoById_FullMethodName  = "/QuestionService/GetQuestionVoById"
	QuestionService_GetQuestionById_FullMethodName    = "/QuestionService/GetQuestionById"
	QuestionService_DeleteQuestionById_FullMethodName = "/QuestionService/DeleteQuestionById"
	QuestionService_UpdateQuestion_FullMethodName     = "/QuestionService/UpdateQuestion"
	QuestionService_GetQuestionVo_FullMethodName      = "/QuestionService/GetQuestionVo"
	QuestionService_GetQuestionTotal_FullMethodName   = "/QuestionService/GetQuestionTotal"
	QuestionService_ListQuestionVoPage_FullMethodName = "/QuestionService/ListQuestionVoPage"
)

// QuestionServiceClient is the client API for QuestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionServiceClient interface {
	GetById(ctx context.Context, in *QuestionIdRequest, opts ...grpc.CallOption) (*QuestionInfo, error)
	ValidQuestion(ctx context.Context, in *ValidQuestionRequest, opts ...grpc.CallOption) (*Empty, error)
	ListQuestionPage(ctx context.Context, in *GetQuestionPageRequest, opts ...grpc.CallOption) (*GetQuestionPageResponse, error)
	AddQuestion(ctx context.Context, in *QuestionAddRequest, opts ...grpc.CallOption) (*BoolResponse, error)
	GetQuestionVoById(ctx context.Context, in *QuestionIdRequest, opts ...grpc.CallOption) (*QuestionVo, error)
	GetQuestionById(ctx context.Context, in *QuestionIdRequest, opts ...grpc.CallOption) (*QuestionInfo, error)
	DeleteQuestionById(ctx context.Context, in *QuestionIdRequest, opts ...grpc.CallOption) (*BoolResponse, error)
	UpdateQuestion(ctx context.Context, in *QuestionInfo, opts ...grpc.CallOption) (*BoolResponse, error)
	GetQuestionVo(ctx context.Context, in *QuestionInfo, opts ...grpc.CallOption) (*QuestionVo, error)
	GetQuestionTotal(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TotalResponse, error)
	ListQuestionVoPage(ctx context.Context, in *GetQuestionPageRequest, opts ...grpc.CallOption) (*ListQuestionPageVoResponse, error)
}

type questionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionServiceClient(cc grpc.ClientConnInterface) QuestionServiceClient {
	return &questionServiceClient{cc}
}

func (c *questionServiceClient) GetById(ctx context.Context, in *QuestionIdRequest, opts ...grpc.CallOption) (*QuestionInfo, error) {
	out := new(QuestionInfo)
	err := c.cc.Invoke(ctx, QuestionService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) ValidQuestion(ctx context.Context, in *ValidQuestionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, QuestionService_ValidQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) ListQuestionPage(ctx context.Context, in *GetQuestionPageRequest, opts ...grpc.CallOption) (*GetQuestionPageResponse, error) {
	out := new(GetQuestionPageResponse)
	err := c.cc.Invoke(ctx, QuestionService_ListQuestionPage_FullMethodName, in, out, opts...)
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

func (c *questionServiceClient) GetQuestionTotal(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TotalResponse, error) {
	out := new(TotalResponse)
	err := c.cc.Invoke(ctx, QuestionService_GetQuestionTotal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) ListQuestionVoPage(ctx context.Context, in *GetQuestionPageRequest, opts ...grpc.CallOption) (*ListQuestionPageVoResponse, error) {
	out := new(ListQuestionPageVoResponse)
	err := c.cc.Invoke(ctx, QuestionService_ListQuestionVoPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionServiceServer is the server API for QuestionService service.
// All implementations must embed UnimplementedQuestionServiceServer
// for forward compatibility
type QuestionServiceServer interface {
	GetById(context.Context, *QuestionIdRequest) (*QuestionInfo, error)
	ValidQuestion(context.Context, *ValidQuestionRequest) (*Empty, error)
	ListQuestionPage(context.Context, *GetQuestionPageRequest) (*GetQuestionPageResponse, error)
	AddQuestion(context.Context, *QuestionAddRequest) (*BoolResponse, error)
	GetQuestionVoById(context.Context, *QuestionIdRequest) (*QuestionVo, error)
	GetQuestionById(context.Context, *QuestionIdRequest) (*QuestionInfo, error)
	DeleteQuestionById(context.Context, *QuestionIdRequest) (*BoolResponse, error)
	UpdateQuestion(context.Context, *QuestionInfo) (*BoolResponse, error)
	GetQuestionVo(context.Context, *QuestionInfo) (*QuestionVo, error)
	GetQuestionTotal(context.Context, *Empty) (*TotalResponse, error)
	ListQuestionVoPage(context.Context, *GetQuestionPageRequest) (*ListQuestionPageVoResponse, error)
	mustEmbedUnimplementedQuestionServiceServer()
}

// UnimplementedQuestionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuestionServiceServer struct {
}

func (UnimplementedQuestionServiceServer) GetById(context.Context, *QuestionIdRequest) (*QuestionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedQuestionServiceServer) ValidQuestion(context.Context, *ValidQuestionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) ListQuestionPage(context.Context, *GetQuestionPageRequest) (*GetQuestionPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuestionPage not implemented")
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
func (UnimplementedQuestionServiceServer) GetQuestionTotal(context.Context, *Empty) (*TotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionTotal not implemented")
}
func (UnimplementedQuestionServiceServer) ListQuestionVoPage(context.Context, *GetQuestionPageRequest) (*ListQuestionPageVoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuestionVoPage not implemented")
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

func _QuestionService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetById(ctx, req.(*QuestionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
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

func _QuestionService_ListQuestionPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).ListQuestionPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_ListQuestionPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).ListQuestionPage(ctx, req.(*GetQuestionPageRequest))
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

func _QuestionService_GetQuestionTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetQuestionTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetQuestionTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetQuestionTotal(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_ListQuestionVoPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).ListQuestionVoPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_ListQuestionVoPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).ListQuestionVoPage(ctx, req.(*GetQuestionPageRequest))
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
			MethodName: "GetById",
			Handler:    _QuestionService_GetById_Handler,
		},
		{
			MethodName: "ValidQuestion",
			Handler:    _QuestionService_ValidQuestion_Handler,
		},
		{
			MethodName: "ListQuestionPage",
			Handler:    _QuestionService_ListQuestionPage_Handler,
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
		{
			MethodName: "GetQuestionTotal",
			Handler:    _QuestionService_GetQuestionTotal_Handler,
		},
		{
			MethodName: "ListQuestionVoPage",
			Handler:    _QuestionService_ListQuestionVoPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "question.proto",
}
