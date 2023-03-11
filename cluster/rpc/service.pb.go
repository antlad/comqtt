// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package rpc

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

type PublishRequest struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	ProtocolVersion      uint32   `protobuf:"varint,3,opt,name=protocolVersion,proto3" json:"protocolVersion,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *PublishRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *PublishRequest) GetProtocolVersion() uint32 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

func (m *PublishRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ConnectRequest struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ConnectRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type Response struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type ApplyRequest struct {
	Action               uint32   `protobuf:"varint,1,opt,name=action,proto3" json:"action,omitempty"`
	NodeId               string   `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Filter               []byte   `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyRequest) Reset()         { *m = ApplyRequest{} }
func (m *ApplyRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyRequest) ProtoMessage()    {}
func (*ApplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *ApplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRequest.Unmarshal(m, b)
}
func (m *ApplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRequest.Marshal(b, m, deterministic)
}
func (m *ApplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRequest.Merge(m, src)
}
func (m *ApplyRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyRequest.Size(m)
}
func (m *ApplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRequest proto.InternalMessageInfo

func (m *ApplyRequest) GetAction() uint32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *ApplyRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ApplyRequest) GetFilter() []byte {
	if m != nil {
		return m.Filter
	}
	return nil
}

type JoinRequest struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *JoinRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *JoinRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*PublishRequest)(nil), "PublishRequest")
	proto.RegisterType((*ConnectRequest)(nil), "ConnectRequest")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*ApplyRequest)(nil), "ApplyRequest")
	proto.RegisterType((*JoinRequest)(nil), "JoinRequest")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x24, 0xa1, 0x0a, 0xe9, 0x92, 0xb4, 0x92, 0x0f, 0x55, 0x54, 0x09, 0x51, 0x45, 0x42, 0xe4,
	0xd2, 0x54, 0x82, 0x2f, 0xe0, 0x71, 0x29, 0x12, 0xa8, 0xf2, 0xa1, 0x07, 0x6e, 0xa9, 0xe3, 0x52,
	0xab, 0xc6, 0x0e, 0xb6, 0x03, 0xca, 0x1f, 0xf0, 0x45, 0x7c, 0x1f, 0x8a, 0xe5, 0xa2, 0x84, 0x0b,
	0x07, 0x6e, 0x3b, 0x9b, 0xc9, 0xee, 0xcc, 0xac, 0x21, 0xd6, 0x54, 0xbd, 0x33, 0x42, 0xf3, 0x4a,
	0x49, 0x23, 0xd3, 0x4f, 0x0f, 0x46, 0xab, 0x7a, 0xc3, 0x99, 0xde, 0x61, 0xfa, 0x56, 0x53, 0x6d,
	0xd0, 0x04, 0x02, 0x21, 0x4b, 0xba, 0x2c, 0x13, 0x6f, 0xe6, 0x65, 0x43, 0xec, 0x10, 0x9a, 0x42,
	0x48, 0x38, 0xa3, 0xc2, 0x2c, 0xcb, 0xc4, 0xb7, 0x5f, 0x7e, 0x30, 0xca, 0x60, 0x6c, 0xe7, 0x11,
	0xc9, 0xd7, 0x54, 0x69, 0x26, 0x45, 0x72, 0x3c, 0xf3, 0xb2, 0x18, 0xff, 0x6e, 0xa3, 0x04, 0x4e,
	0xaa, 0xa2, 0xe1, 0xb2, 0x28, 0x93, 0xc1, 0xcc, 0xcb, 0x22, 0x7c, 0x80, 0xe9, 0x3d, 0x8c, 0xee,
	0xa4, 0x10, 0x94, 0x98, 0x7f, 0x28, 0x49, 0xa7, 0x10, 0x62, 0xaa, 0x2b, 0x29, 0x34, 0x45, 0x23,
	0xf0, 0xe5, 0xde, 0xfe, 0x1b, 0x62, 0x5f, 0xee, 0xd3, 0x35, 0x44, 0x37, 0x55, 0xc5, 0x9b, 0xce,
	0xfc, 0x82, 0x98, 0x56, 0xac, 0x67, 0xc5, 0x3a, 0xd4, 0xd9, 0xeb, 0xf7, 0xf6, 0x4e, 0x20, 0xd8,
	0x32, 0x6e, 0xa8, 0xb2, 0xe6, 0x22, 0xec, 0x50, 0xfa, 0x08, 0xa7, 0x0f, 0x92, 0x89, 0xbf, 0x64,
	0x23, 0x18, 0x14, 0x65, 0xa9, 0xdc, 0x50, 0x5b, 0xb7, 0xbd, 0x4a, 0x2a, 0xe3, 0xd2, 0xb2, 0xf5,
	0xd5, 0x97, 0x07, 0x01, 0xa6, 0xbc, 0x68, 0x34, 0x9a, 0x43, 0xec, 0xae, 0xb3, 0x2a, 0xc8, 0x9e,
	0x1a, 0x34, 0xce, 0xfb, 0xd7, 0x9a, 0x0e, 0xf3, 0x83, 0xdd, 0xf4, 0xa8, 0xa5, 0xbb, 0x08, 0x9f,
	0xa4, 0x61, 0xdb, 0x06, 0x8d, 0xf3, 0x7e, 0xa4, 0x7d, 0xfa, 0x25, 0x0c, 0x71, 0xb1, 0x35, 0x36,
	0x13, 0x14, 0xe7, 0xdd, 0x6c, 0xfa, 0xc4, 0x0b, 0x08, 0x5b, 0x62, 0x6b, 0x12, 0x45, 0x79, 0xc7,
	0x6b, 0x8f, 0x76, 0x7b, 0xfe, 0x7c, 0xf6, 0xc2, 0xcc, 0xae, 0xde, 0xe4, 0x44, 0xbe, 0x2e, 0x3e,
	0x98, 0x28, 0xe7, 0x64, 0x41, 0x78, 0xad, 0x0d, 0x55, 0x0b, 0x55, 0x91, 0x4d, 0x60, 0x5f, 0xc3,
	0xf5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x1d, 0xd6, 0x0c, 0x85, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RelaysClient is the client API for Relays service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RelaysClient interface {
	PublishPacket(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*Response, error)
	ConnectNotify(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*Response, error)
	RaftApply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*Response, error)
	RaftJoin(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*Response, error)
}

type relaysClient struct {
	cc *grpc.ClientConn
}

func NewRelaysClient(cc *grpc.ClientConn) RelaysClient {
	return &relaysClient{cc}
}

func (c *relaysClient) PublishPacket(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Relays/PublishPacket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relaysClient) ConnectNotify(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Relays/ConnectNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relaysClient) RaftApply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Relays/RaftApply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relaysClient) RaftJoin(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Relays/RaftJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelaysServer is the server API for Relays service.
type RelaysServer interface {
	PublishPacket(context.Context, *PublishRequest) (*Response, error)
	ConnectNotify(context.Context, *ConnectRequest) (*Response, error)
	RaftApply(context.Context, *ApplyRequest) (*Response, error)
	RaftJoin(context.Context, *JoinRequest) (*Response, error)
}

// UnimplementedRelaysServer can be embedded to have forward compatible implementations.
type UnimplementedRelaysServer struct {
}

func (*UnimplementedRelaysServer) PublishPacket(ctx context.Context, req *PublishRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishPacket not implemented")
}
func (*UnimplementedRelaysServer) ConnectNotify(ctx context.Context, req *ConnectRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectNotify not implemented")
}
func (*UnimplementedRelaysServer) RaftApply(ctx context.Context, req *ApplyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaftApply not implemented")
}
func (*UnimplementedRelaysServer) RaftJoin(ctx context.Context, req *JoinRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaftJoin not implemented")
}

func RegisterRelaysServer(s *grpc.Server, srv RelaysServer) {
	s.RegisterService(&_Relays_serviceDesc, srv)
}

func _Relays_PublishPacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelaysServer).PublishPacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Relays/PublishPacket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelaysServer).PublishPacket(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relays_ConnectNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelaysServer).ConnectNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Relays/ConnectNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelaysServer).ConnectNotify(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relays_RaftApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelaysServer).RaftApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Relays/RaftApply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelaysServer).RaftApply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relays_RaftJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelaysServer).RaftJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Relays/RaftJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelaysServer).RaftJoin(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Relays_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Relays",
	HandlerType: (*RelaysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishPacket",
			Handler:    _Relays_PublishPacket_Handler,
		},
		{
			MethodName: "ConnectNotify",
			Handler:    _Relays_ConnectNotify_Handler,
		},
		{
			MethodName: "RaftApply",
			Handler:    _Relays_RaftApply_Handler,
		},
		{
			MethodName: "RaftJoin",
			Handler:    _Relays_RaftJoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
