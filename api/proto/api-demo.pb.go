// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api-demo.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type StringMessage struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMessage) Reset()         { *m = StringMessage{} }
func (m *StringMessage) String() string { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()    {}
func (*StringMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_441f1be6025a891e, []int{0}
}

func (m *StringMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMessage.Unmarshal(m, b)
}
func (m *StringMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMessage.Marshal(b, m, deterministic)
}
func (m *StringMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMessage.Merge(m, src)
}
func (m *StringMessage) XXX_Size() int {
	return xxx_messageInfo_StringMessage.Size(m)
}
func (m *StringMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StringMessage proto.InternalMessageInfo

func (m *StringMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "api.StringMessage")
}

func init() { proto.RegisterFile("api-demo.proto", fileDescriptor_441f1be6025a891e) }

var fileDescriptor_441f1be6025a891e = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2c, 0xc8, 0xd4,
	0x4d, 0x49, 0xcd, 0xcd, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94,
	0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb,
	0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x51, 0x52, 0xe5, 0xe2, 0x0d, 0x2e, 0x29,
	0xca, 0xcc, 0x4b, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x4b,
	0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x8c, 0xa6, 0x30, 0x72,
	0x71, 0x07, 0xa5, 0x16, 0x97, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x39, 0x73, 0x31,
	0xbb, 0xa7, 0x96, 0x08, 0x09, 0xe9, 0x25, 0x16, 0x64, 0xea, 0xa1, 0x18, 0x20, 0x85, 0x45, 0x4c,
	0x49, 0xa4, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x7c, 0x42, 0x3c, 0xfa, 0xe9, 0xa9, 0x25, 0xfa, 0xd5,
	0x60, 0x33, 0x6b, 0x85, 0x1c, 0xb8, 0x58, 0x02, 0xf2, 0x8b, 0x89, 0x37, 0x45, 0x00, 0x6c, 0x0a,
	0x97, 0x12, 0xab, 0x7e, 0x41, 0x7e, 0x71, 0x89, 0x15, 0xa3, 0x56, 0x12, 0x1b, 0xd8, 0x13, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x08, 0x8b, 0x44, 0xf9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RestServiceClient is the client API for RestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RestServiceClient interface {
	Get(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
	Post(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type restServiceClient struct {
	cc *grpc.ClientConn
}

func NewRestServiceClient(cc *grpc.ClientConn) RestServiceClient {
	return &restServiceClient{cc}
}

func (c *restServiceClient) Get(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/api.RestService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restServiceClient) Post(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/api.RestService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestServiceServer is the server API for RestService service.
type RestServiceServer interface {
	Get(context.Context, *StringMessage) (*StringMessage, error)
	Post(context.Context, *StringMessage) (*StringMessage, error)
}

// UnimplementedRestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRestServiceServer struct {
}

func (*UnimplementedRestServiceServer) Get(ctx context.Context, req *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRestServiceServer) Post(ctx context.Context, req *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}

func RegisterRestServiceServer(s *grpc.Server, srv RestServiceServer) {
	s.RegisterService(&_RestService_serviceDesc, srv)
}

func _RestService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RestService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestServiceServer).Get(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RestService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestServiceServer).Post(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _RestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.RestService",
	HandlerType: (*RestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RestService_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _RestService_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api-demo.proto",
}