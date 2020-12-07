// Code generated by protoc-gen-go. DO NOT EDIT.
// source: welcome.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type WelcomeReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WelcomeReq) Reset()         { *m = WelcomeReq{} }
func (m *WelcomeReq) String() string { return proto.CompactTextString(m) }
func (*WelcomeReq) ProtoMessage()    {}
func (*WelcomeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f13780f0536c7504, []int{0}
}

func (m *WelcomeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WelcomeReq.Unmarshal(m, b)
}
func (m *WelcomeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WelcomeReq.Marshal(b, m, deterministic)
}
func (m *WelcomeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WelcomeReq.Merge(m, src)
}
func (m *WelcomeReq) XXX_Size() int {
	return xxx_messageInfo_WelcomeReq.Size(m)
}
func (m *WelcomeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_WelcomeReq.DiscardUnknown(m)
}

var xxx_messageInfo_WelcomeReq proto.InternalMessageInfo

type WelcomeResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WelcomeResp) Reset()         { *m = WelcomeResp{} }
func (m *WelcomeResp) String() string { return proto.CompactTextString(m) }
func (*WelcomeResp) ProtoMessage()    {}
func (*WelcomeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f13780f0536c7504, []int{1}
}

func (m *WelcomeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WelcomeResp.Unmarshal(m, b)
}
func (m *WelcomeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WelcomeResp.Marshal(b, m, deterministic)
}
func (m *WelcomeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WelcomeResp.Merge(m, src)
}
func (m *WelcomeResp) XXX_Size() int {
	return xxx_messageInfo_WelcomeResp.Size(m)
}
func (m *WelcomeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_WelcomeResp.DiscardUnknown(m)
}

var xxx_messageInfo_WelcomeResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*WelcomeReq)(nil), "main.WelcomeReq")
	proto.RegisterType((*WelcomeResp)(nil), "main.WelcomeResp")
}

func init() { proto.RegisterFile("welcome.proto", fileDescriptor_f13780f0536c7504) }

var fileDescriptor_f13780f0536c7504 = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4f, 0xcd, 0x49,
	0xce, 0xcf, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53,
	0xe2, 0xe1, 0xe2, 0x0a, 0x87, 0x08, 0x07, 0xa5, 0x16, 0x2a, 0xf1, 0x72, 0x71, 0xc3, 0x79, 0xc5,
	0x05, 0x46, 0x36, 0x5c, 0x5c, 0x8e, 0x05, 0x99, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42,
	0x7a, 0x5c, 0xec, 0x50, 0x49, 0x21, 0x01, 0x3d, 0x90, 0x66, 0x3d, 0x84, 0x4e, 0x29, 0x41, 0x34,
	0x91, 0xe2, 0x82, 0x24, 0x36, 0xb0, 0x3d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x63,
	0x8b, 0xbf, 0x78, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiServiceClient interface {
	Welcome(ctx context.Context, in *WelcomeReq, opts ...grpc.CallOption) (*WelcomeResp, error)
}

type apiServiceClient struct {
	cc *grpc.ClientConn
}

func NewApiServiceClient(cc *grpc.ClientConn) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) Welcome(ctx context.Context, in *WelcomeReq, opts ...grpc.CallOption) (*WelcomeResp, error) {
	out := new(WelcomeResp)
	err := c.cc.Invoke(ctx, "/main.ApiService/Welcome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
type ApiServiceServer interface {
	Welcome(context.Context, *WelcomeReq) (*WelcomeResp, error)
}

// UnimplementedApiServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (*UnimplementedApiServiceServer) Welcome(ctx context.Context, req *WelcomeReq) (*WelcomeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Welcome not implemented")
}

func RegisterApiServiceServer(s *grpc.Server, srv ApiServiceServer) {
	s.RegisterService(&_ApiService_serviceDesc, srv)
}

func _ApiService_Welcome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WelcomeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).Welcome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ApiService/Welcome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).Welcome(ctx, req.(*WelcomeReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Welcome",
			Handler:    _ApiService_Welcome_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "welcome.proto",
}