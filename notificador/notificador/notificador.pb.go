// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notificador.proto

package notificador

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

type CorreoRequest struct {
	Destino              string   `protobuf:"bytes,1,opt,name=destino,proto3" json:"destino,omitempty"`
	Asunto               string   `protobuf:"bytes,2,opt,name=asunto,proto3" json:"asunto,omitempty"`
	Mensaje              string   `protobuf:"bytes,3,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CorreoRequest) Reset()         { *m = CorreoRequest{} }
func (m *CorreoRequest) String() string { return proto.CompactTextString(m) }
func (*CorreoRequest) ProtoMessage()    {}
func (*CorreoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead5d8a612625130, []int{0}
}

func (m *CorreoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CorreoRequest.Unmarshal(m, b)
}
func (m *CorreoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CorreoRequest.Marshal(b, m, deterministic)
}
func (m *CorreoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CorreoRequest.Merge(m, src)
}
func (m *CorreoRequest) XXX_Size() int {
	return xxx_messageInfo_CorreoRequest.Size(m)
}
func (m *CorreoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CorreoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CorreoRequest proto.InternalMessageInfo

func (m *CorreoRequest) GetDestino() string {
	if m != nil {
		return m.Destino
	}
	return ""
}

func (m *CorreoRequest) GetAsunto() string {
	if m != nil {
		return m.Asunto
	}
	return ""
}

func (m *CorreoRequest) GetMensaje() string {
	if m != nil {
		return m.Mensaje
	}
	return ""
}

type CorreoResponse struct {
	Enviado              bool     `protobuf:"varint,1,opt,name=enviado,proto3" json:"enviado,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CorreoResponse) Reset()         { *m = CorreoResponse{} }
func (m *CorreoResponse) String() string { return proto.CompactTextString(m) }
func (*CorreoResponse) ProtoMessage()    {}
func (*CorreoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ead5d8a612625130, []int{1}
}

func (m *CorreoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CorreoResponse.Unmarshal(m, b)
}
func (m *CorreoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CorreoResponse.Marshal(b, m, deterministic)
}
func (m *CorreoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CorreoResponse.Merge(m, src)
}
func (m *CorreoResponse) XXX_Size() int {
	return xxx_messageInfo_CorreoResponse.Size(m)
}
func (m *CorreoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CorreoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CorreoResponse proto.InternalMessageInfo

func (m *CorreoResponse) GetEnviado() bool {
	if m != nil {
		return m.Enviado
	}
	return false
}

func init() {
	proto.RegisterType((*CorreoRequest)(nil), "notificador.CorreoRequest")
	proto.RegisterType((*CorreoResponse)(nil), "notificador.CorreoResponse")
}

func init() { proto.RegisterFile("notificador.proto", fileDescriptor_ead5d8a612625130) }

var fileDescriptor_ead5d8a612625130 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xcb, 0x2f, 0xc9,
	0x4c, 0xcb, 0x4c, 0x4e, 0x4c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x52, 0x8a, 0xe6, 0xe2, 0x75, 0xce, 0x2f, 0x2a, 0x4a, 0xcd, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0x49, 0x2d, 0x2e, 0xc9, 0xcc, 0xcb, 0x97, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xc4, 0xb8, 0xd8, 0x12, 0x8b, 0x4b, 0xf3, 0x4a, 0xf2,
	0x25, 0x98, 0xc0, 0x12, 0x50, 0x1e, 0x48, 0x47, 0x6e, 0x6a, 0x5e, 0x71, 0x62, 0x56, 0xaa, 0x04,
	0x33, 0x44, 0x07, 0x94, 0xab, 0xa4, 0xc5, 0xc5, 0x07, 0x33, 0xbc, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x15, 0xa4, 0x36, 0x35, 0xaf, 0x2c, 0x33, 0x31, 0x05, 0x62, 0x3a, 0x47, 0x10, 0x8c, 0x6b, 0x14,
	0xc1, 0xc5, 0xed, 0x87, 0x70, 0x97, 0x90, 0x27, 0x17, 0x8f, 0x2b, 0x48, 0xa6, 0x08, 0x62, 0x80,
	0x90, 0x94, 0x1e, 0xb2, 0x47, 0x50, 0x9c, 0x2c, 0x25, 0x8d, 0x55, 0x0e, 0x62, 0xa3, 0x12, 0x43,
	0x12, 0x1b, 0xd8, 0xdb, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0x2e, 0x36, 0x60, 0x0b,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificadorClient is the client API for Notificador service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificadorClient interface {
	EnviarCorreo(ctx context.Context, in *CorreoRequest, opts ...grpc.CallOption) (*CorreoResponse, error)
}

type notificadorClient struct {
	cc *grpc.ClientConn
}

func NewNotificadorClient(cc *grpc.ClientConn) NotificadorClient {
	return &notificadorClient{cc}
}

func (c *notificadorClient) EnviarCorreo(ctx context.Context, in *CorreoRequest, opts ...grpc.CallOption) (*CorreoResponse, error) {
	out := new(CorreoResponse)
	err := c.cc.Invoke(ctx, "/notificador.Notificador/EnviarCorreo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificadorServer is the server API for Notificador service.
type NotificadorServer interface {
	EnviarCorreo(context.Context, *CorreoRequest) (*CorreoResponse, error)
}

// UnimplementedNotificadorServer can be embedded to have forward compatible implementations.
type UnimplementedNotificadorServer struct {
}

func (*UnimplementedNotificadorServer) EnviarCorreo(ctx context.Context, req *CorreoRequest) (*CorreoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarCorreo not implemented")
}

func RegisterNotificadorServer(s *grpc.Server, srv NotificadorServer) {
	s.RegisterService(&_Notificador_serviceDesc, srv)
}

func _Notificador_EnviarCorreo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CorreoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificadorServer).EnviarCorreo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notificador.Notificador/EnviarCorreo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificadorServer).EnviarCorreo(ctx, req.(*CorreoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notificador_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notificador.Notificador",
	HandlerType: (*NotificadorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarCorreo",
			Handler:    _Notificador_EnviarCorreo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notificador.proto",
}
