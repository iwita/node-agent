// Code generated by protoc-gen-go. DO NOT EDIT.
// source: info.proto

package info

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

type PodType struct {
	PodName              string   `protobuf:"bytes,1,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PodType) Reset()         { *m = PodType{} }
func (m *PodType) String() string { return proto.CompactTextString(m) }
func (*PodType) ProtoMessage()    {}
func (*PodType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{0}
}

func (m *PodType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodType.Unmarshal(m, b)
}
func (m *PodType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodType.Marshal(b, m, deterministic)
}
func (m *PodType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodType.Merge(m, src)
}
func (m *PodType) XXX_Size() int {
	return xxx_messageInfo_PodType.Size(m)
}
func (m *PodType) XXX_DiscardUnknown() {
	xxx_messageInfo_PodType.DiscardUnknown(m)
}

var xxx_messageInfo_PodType proto.InternalMessageInfo

func (m *PodType) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

type SocketType struct {
	SocketId             int32    `protobuf:"varint,1,opt,name=socket_id,json=socketId,proto3" json:"socket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketType) Reset()         { *m = SocketType{} }
func (m *SocketType) String() string { return proto.CompactTextString(m) }
func (*SocketType) ProtoMessage()    {}
func (*SocketType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{1}
}

func (m *SocketType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketType.Unmarshal(m, b)
}
func (m *SocketType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketType.Marshal(b, m, deterministic)
}
func (m *SocketType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketType.Merge(m, src)
}
func (m *SocketType) XXX_Size() int {
	return xxx_messageInfo_SocketType.Size(m)
}
func (m *SocketType) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketType.DiscardUnknown(m)
}

var xxx_messageInfo_SocketType proto.InternalMessageInfo

func (m *SocketType) GetSocketId() int32 {
	if m != nil {
		return m.SocketId
	}
	return 0
}

type Info struct {
	Socket               *SocketType `protobuf:"bytes,1,opt,name=socket,proto3" json:"socket,omitempty"`
	Pod                  *PodType    `protobuf:"bytes,2,opt,name=pod,proto3" json:"pod,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{2}
}

func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetSocket() *SocketType {
	if m != nil {
		return m.Socket
	}
	return nil
}

func (m *Info) GetPod() *PodType {
	if m != nil {
		return m.Pod
	}
	return nil
}

type AgentMessage struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentMessage) Reset()         { *m = AgentMessage{} }
func (m *AgentMessage) String() string { return proto.CompactTextString(m) }
func (*AgentMessage) ProtoMessage()    {}
func (*AgentMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{3}
}

func (m *AgentMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentMessage.Unmarshal(m, b)
}
func (m *AgentMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentMessage.Marshal(b, m, deterministic)
}
func (m *AgentMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentMessage.Merge(m, src)
}
func (m *AgentMessage) XXX_Size() int {
	return xxx_messageInfo_AgentMessage.Size(m)
}
func (m *AgentMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AgentMessage proto.InternalMessageInfo

func (m *AgentMessage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*PodType)(nil), "info.PodType")
	proto.RegisterType((*SocketType)(nil), "info.SocketType")
	proto.RegisterType((*Info)(nil), "info.Info")
	proto.RegisterType((*AgentMessage)(nil), "info.AgentMessage")
}

func init() { proto.RegisterFile("info.proto", fileDescriptor_f140d5b28dddb141) }

var fileDescriptor_f140d5b28dddb141 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xad, 0x5d, 0xfb, 0x31, 0x7e, 0x20, 0x73, 0xaa, 0x7a, 0x50, 0x82, 0x60, 0xbd, 0xf4,
	0x50, 0x4f, 0x1e, 0x7b, 0xec, 0x41, 0xa9, 0xd1, 0x7b, 0x49, 0x9d, 0x69, 0x09, 0x65, 0x33, 0x61,
	0x13, 0x84, 0xfd, 0xef, 0x65, 0x27, 0x0b, 0x7a, 0x7b, 0x79, 0xef, 0xc7, 0xcb, 0x63, 0x00, 0x7c,
	0xd8, 0xcb, 0x22, 0x36, 0x92, 0x05, 0xab, 0x4e, 0x9b, 0x47, 0x18, 0x6f, 0x84, 0xbe, 0xda, 0xc8,
	0x78, 0x03, 0x93, 0x28, 0xb4, 0x0d, 0xae, 0xe6, 0xd9, 0xe0, 0x61, 0x30, 0x9f, 0xda, 0x71, 0x14,
	0x7a, 0x77, 0x35, 0x9b, 0x67, 0x80, 0x4f, 0xf9, 0x3e, 0x72, 0x56, 0xf0, 0x0e, 0xa6, 0x49, 0x5f,
	0x5b, 0x4f, 0x4a, 0x9e, 0xd9, 0x49, 0x31, 0xd6, 0x64, 0x3e, 0xa0, 0x5a, 0x87, 0xbd, 0xe0, 0x1c,
	0x46, 0xc5, 0x53, 0xe2, 0x7c, 0x79, 0xbd, 0xd0, 0xbf, 0xff, 0x6a, 0x6c, 0x9f, 0xe3, 0x3d, 0x0c,
	0xa3, 0xd0, 0xec, 0x54, 0xb1, 0xcb, 0x82, 0xf5, 0x9b, 0x6c, 0x97, 0x18, 0x03, 0x17, 0xab, 0x03,
	0x87, 0xfc, 0xc6, 0x29, 0xb9, 0x03, 0x23, 0x42, 0xb5, 0x13, 0x6a, 0xfb, 0x91, 0xaa, 0x97, 0xaf,
	0x70, 0x95, 0x5d, 0x3a, 0x26, 0xce, 0x2b, 0xfa, 0xf1, 0x49, 0x1a, 0x7c, 0x82, 0xe1, 0xc6, 0x07,
	0x84, 0x52, 0xd8, 0x6d, 0xba, 0xc5, 0xa2, 0xff, 0x97, 0x99, 0x93, 0xdd, 0x48, 0xef, 0xf1, 0xf2,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x05, 0x2c, 0x5b, 0xcb, 0x1d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TasksetAdvisorClient is the client API for TasksetAdvisor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TasksetAdvisorClient interface {
	Pin(ctx context.Context, in *Info, opts ...grpc.CallOption) (*AgentMessage, error)
}

type tasksetAdvisorClient struct {
	cc *grpc.ClientConn
}

func NewTasksetAdvisorClient(cc *grpc.ClientConn) TasksetAdvisorClient {
	return &tasksetAdvisorClient{cc}
}

func (c *tasksetAdvisorClient) Pin(ctx context.Context, in *Info, opts ...grpc.CallOption) (*AgentMessage, error) {
	out := new(AgentMessage)
	err := c.cc.Invoke(ctx, "/info.tasksetAdvisor/Pin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksetAdvisorServer is the server API for TasksetAdvisor service.
type TasksetAdvisorServer interface {
	Pin(context.Context, *Info) (*AgentMessage, error)
}

// UnimplementedTasksetAdvisorServer can be embedded to have forward compatible implementations.
type UnimplementedTasksetAdvisorServer struct {
}

func (*UnimplementedTasksetAdvisorServer) Pin(ctx context.Context, req *Info) (*AgentMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pin not implemented")
}

func RegisterTasksetAdvisorServer(s *grpc.Server, srv TasksetAdvisorServer) {
	s.RegisterService(&_TasksetAdvisor_serviceDesc, srv)
}

func _TasksetAdvisor_Pin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Info)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksetAdvisorServer).Pin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.tasksetAdvisor/Pin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksetAdvisorServer).Pin(ctx, req.(*Info))
	}
	return interceptor(ctx, in, info, handler)
}

var _TasksetAdvisor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "info.tasksetAdvisor",
	HandlerType: (*TasksetAdvisorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pin",
			Handler:    _TasksetAdvisor_Pin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info.proto",
}
