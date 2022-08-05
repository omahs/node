// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mirror/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgDeployERC20Mirror struct {
	Creator                  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	HomeChain                string `protobuf:"bytes,2,opt,name=homeChain,proto3" json:"homeChain,omitempty"`
	HomeERC20ContractAddress string `protobuf:"bytes,3,opt,name=homeERC20ContractAddress,proto3" json:"homeERC20ContractAddress,omitempty"`
	Name                     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Symbol                   string `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimals                 uint32 `protobuf:"varint,6,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (m *MsgDeployERC20Mirror) Reset()         { *m = MsgDeployERC20Mirror{} }
func (m *MsgDeployERC20Mirror) String() string { return proto.CompactTextString(m) }
func (*MsgDeployERC20Mirror) ProtoMessage()    {}
func (*MsgDeployERC20Mirror) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b5ebd4422c633b3, []int{0}
}
func (m *MsgDeployERC20Mirror) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeployERC20Mirror) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeployERC20Mirror.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeployERC20Mirror) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeployERC20Mirror.Merge(m, src)
}
func (m *MsgDeployERC20Mirror) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeployERC20Mirror) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeployERC20Mirror.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeployERC20Mirror proto.InternalMessageInfo

func (m *MsgDeployERC20Mirror) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeployERC20Mirror) GetHomeChain() string {
	if m != nil {
		return m.HomeChain
	}
	return ""
}

func (m *MsgDeployERC20Mirror) GetHomeERC20ContractAddress() string {
	if m != nil {
		return m.HomeERC20ContractAddress
	}
	return ""
}

func (m *MsgDeployERC20Mirror) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgDeployERC20Mirror) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MsgDeployERC20Mirror) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

type MsgDeployERC20MirrorResponse struct {
}

func (m *MsgDeployERC20MirrorResponse) Reset()         { *m = MsgDeployERC20MirrorResponse{} }
func (m *MsgDeployERC20MirrorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeployERC20MirrorResponse) ProtoMessage()    {}
func (*MsgDeployERC20MirrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b5ebd4422c633b3, []int{1}
}
func (m *MsgDeployERC20MirrorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeployERC20MirrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeployERC20MirrorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeployERC20MirrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeployERC20MirrorResponse.Merge(m, src)
}
func (m *MsgDeployERC20MirrorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeployERC20MirrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeployERC20MirrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeployERC20MirrorResponse proto.InternalMessageInfo

type MsgDepoistERC20 struct {
	Creator                  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	HomeERC20ContractAddress string `protobuf:"bytes,2,opt,name=homeERC20ContractAddress,proto3" json:"homeERC20ContractAddress,omitempty"`
	RecipientAddress         string `protobuf:"bytes,3,opt,name=recipientAddress,proto3" json:"recipientAddress,omitempty"`
	Amount                   string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *MsgDepoistERC20) Reset()         { *m = MsgDepoistERC20{} }
func (m *MsgDepoistERC20) String() string { return proto.CompactTextString(m) }
func (*MsgDepoistERC20) ProtoMessage()    {}
func (*MsgDepoistERC20) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b5ebd4422c633b3, []int{2}
}
func (m *MsgDepoistERC20) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepoistERC20) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepoistERC20.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepoistERC20) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepoistERC20.Merge(m, src)
}
func (m *MsgDepoistERC20) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepoistERC20) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepoistERC20.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepoistERC20 proto.InternalMessageInfo

func (m *MsgDepoistERC20) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDepoistERC20) GetHomeERC20ContractAddress() string {
	if m != nil {
		return m.HomeERC20ContractAddress
	}
	return ""
}

func (m *MsgDepoistERC20) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *MsgDepoistERC20) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type MsgDepoistERC20Response struct {
}

func (m *MsgDepoistERC20Response) Reset()         { *m = MsgDepoistERC20Response{} }
func (m *MsgDepoistERC20Response) String() string { return proto.CompactTextString(m) }
func (*MsgDepoistERC20Response) ProtoMessage()    {}
func (*MsgDepoistERC20Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b5ebd4422c633b3, []int{3}
}
func (m *MsgDepoistERC20Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepoistERC20Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepoistERC20Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepoistERC20Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepoistERC20Response.Merge(m, src)
}
func (m *MsgDepoistERC20Response) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepoistERC20Response) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepoistERC20Response.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepoistERC20Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDeployERC20Mirror)(nil), "zetachain.zetacore.mirror.MsgDeployERC20Mirror")
	proto.RegisterType((*MsgDeployERC20MirrorResponse)(nil), "zetachain.zetacore.mirror.MsgDeployERC20MirrorResponse")
	proto.RegisterType((*MsgDepoistERC20)(nil), "zetachain.zetacore.mirror.MsgDepoistERC20")
	proto.RegisterType((*MsgDepoistERC20Response)(nil), "zetachain.zetacore.mirror.MsgDepoistERC20Response")
}

func init() { proto.RegisterFile("mirror/tx.proto", fileDescriptor_8b5ebd4422c633b3) }

var fileDescriptor_8b5ebd4422c633b3 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x4e, 0x02, 0x31,
	0x10, 0xa6, 0x80, 0x28, 0x13, 0x0d, 0xda, 0x18, 0x2d, 0x84, 0x6c, 0xc8, 0x9e, 0x08, 0x09, 0xbb,
	0x06, 0x0f, 0x26, 0xde, 0x14, 0x8d, 0x27, 0x2e, 0x7b, 0xf4, 0xb6, 0x2c, 0x0d, 0x6c, 0x42, 0xdb,
	0x4d, 0x5b, 0x12, 0x30, 0xf1, 0x1d, 0x7c, 0x0f, 0x5f, 0xc4, 0x23, 0xf1, 0xe4, 0xd1, 0xc0, 0x2b,
	0xf8, 0x00, 0x66, 0x0b, 0x8b, 0x91, 0x5f, 0xbd, 0xcd, 0x37, 0xf3, 0xf5, 0xeb, 0x7c, 0x33, 0x19,
	0x28, 0xb0, 0x50, 0x4a, 0x21, 0x5d, 0x3d, 0x74, 0x22, 0x29, 0xb4, 0xc0, 0xc5, 0x27, 0xaa, 0xfd,
	0xa0, 0xe7, 0x87, 0xdc, 0x31, 0x91, 0x90, 0xd4, 0x99, 0x71, 0xec, 0x77, 0x04, 0xa7, 0x2d, 0xd5,
	0xbd, 0xa3, 0x51, 0x5f, 0x8c, 0xee, 0xbd, 0x66, 0xe3, 0xa2, 0x65, 0x0a, 0x98, 0xc0, 0x7e, 0x20,
	0xa9, 0xaf, 0x85, 0x24, 0xa8, 0x82, 0xaa, 0x79, 0x2f, 0x81, 0xb8, 0x0c, 0xf9, 0x9e, 0x60, 0xb4,
	0x19, 0xeb, 0x91, 0xb4, 0xa9, 0xfd, 0x24, 0xf0, 0x35, 0x90, 0x18, 0x18, 0xa9, 0xa6, 0xe0, 0x5a,
	0xfa, 0x81, 0xbe, 0xe9, 0x74, 0x24, 0x55, 0x8a, 0x64, 0x0c, 0x79, 0x63, 0x1d, 0x63, 0xc8, 0x72,
	0x9f, 0x51, 0x92, 0x35, 0x3c, 0x13, 0xe3, 0x33, 0xc8, 0xa9, 0x11, 0x6b, 0x8b, 0x3e, 0xd9, 0x33,
	0xd9, 0x39, 0xc2, 0x25, 0x38, 0xe8, 0xd0, 0x20, 0x64, 0x7e, 0x5f, 0x91, 0x5c, 0x05, 0x55, 0x8f,
	0xbc, 0x05, 0xb6, 0x2d, 0x28, 0xaf, 0xf3, 0xe4, 0x51, 0x15, 0x09, 0xae, 0xa8, 0xfd, 0x8a, 0xa0,
	0x30, 0x23, 0x88, 0x50, 0x69, 0xc3, 0xd8, 0xe2, 0x77, 0x9b, 0xa3, 0xf4, 0x0e, 0x47, 0x35, 0x38,
	0x96, 0x34, 0x08, 0xa3, 0x90, 0xf2, 0xa5, 0x29, 0xac, 0xe4, 0x63, 0xa7, 0x3e, 0x13, 0x03, 0xae,
	0xe7, 0xfe, 0xe7, 0xc8, 0x2e, 0xc2, 0xf9, 0x52, 0xb3, 0x89, 0x91, 0xc6, 0x17, 0x82, 0x4c, 0x4b,
	0x75, 0xf1, 0x33, 0x9c, 0xac, 0x6e, 0xd0, 0x75, 0x36, 0xae, 0xdd, 0x59, 0x37, 0x9e, 0xd2, 0xd5,
	0x3f, 0x1f, 0x24, 0x6d, 0x60, 0x0e, 0x87, 0xbf, 0x66, 0x59, 0xdb, 0x29, 0xb4, 0xe0, 0x96, 0x1a,
	0x7f, 0xe7, 0x26, 0xff, 0xdd, 0x3e, 0xbc, 0x4d, 0x2c, 0x34, 0x9e, 0x58, 0xe8, 0x73, 0x62, 0xa1,
	0x97, 0xa9, 0x95, 0x1a, 0x4f, 0xad, 0xd4, 0xc7, 0xd4, 0x4a, 0x3d, 0xd6, 0xbb, 0xa1, 0xee, 0x0d,
	0xda, 0x4e, 0x20, 0x98, 0x1b, 0xab, 0xd5, 0x8d, 0xb0, 0x9b, 0x08, 0xbb, 0x43, 0x37, 0xb9, 0x8d,
	0x51, 0x44, 0x55, 0x3b, 0x67, 0xee, 0xe3, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x74, 0xc6, 0x6f,
	0xb8, 0x32, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	DeployERC20Mirror(ctx context.Context, in *MsgDeployERC20Mirror, opts ...grpc.CallOption) (*MsgDeployERC20MirrorResponse, error)
	DepoistERC20(ctx context.Context, in *MsgDepoistERC20, opts ...grpc.CallOption) (*MsgDepoistERC20Response, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DeployERC20Mirror(ctx context.Context, in *MsgDeployERC20Mirror, opts ...grpc.CallOption) (*MsgDeployERC20MirrorResponse, error) {
	out := new(MsgDeployERC20MirrorResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.mirror.Msg/DeployERC20Mirror", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DepoistERC20(ctx context.Context, in *MsgDepoistERC20, opts ...grpc.CallOption) (*MsgDepoistERC20Response, error) {
	out := new(MsgDepoistERC20Response)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.mirror.Msg/DepoistERC20", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	DeployERC20Mirror(context.Context, *MsgDeployERC20Mirror) (*MsgDeployERC20MirrorResponse, error)
	DepoistERC20(context.Context, *MsgDepoistERC20) (*MsgDepoistERC20Response, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) DeployERC20Mirror(ctx context.Context, req *MsgDeployERC20Mirror) (*MsgDeployERC20MirrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployERC20Mirror not implemented")
}
func (*UnimplementedMsgServer) DepoistERC20(ctx context.Context, req *MsgDepoistERC20) (*MsgDepoistERC20Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepoistERC20 not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_DeployERC20Mirror_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeployERC20Mirror)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeployERC20Mirror(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.mirror.Msg/DeployERC20Mirror",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeployERC20Mirror(ctx, req.(*MsgDeployERC20Mirror))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DepoistERC20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDepoistERC20)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DepoistERC20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.mirror.Msg/DepoistERC20",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DepoistERC20(ctx, req.(*MsgDepoistERC20))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zetachain.zetacore.mirror.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeployERC20Mirror",
			Handler:    _Msg_DeployERC20Mirror_Handler,
		},
		{
			MethodName: "DepoistERC20",
			Handler:    _Msg_DepoistERC20_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mirror/tx.proto",
}

func (m *MsgDeployERC20Mirror) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeployERC20Mirror) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeployERC20Mirror) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimals != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.HomeERC20ContractAddress) > 0 {
		i -= len(m.HomeERC20ContractAddress)
		copy(dAtA[i:], m.HomeERC20ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.HomeERC20ContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.HomeChain) > 0 {
		i -= len(m.HomeChain)
		copy(dAtA[i:], m.HomeChain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.HomeChain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeployERC20MirrorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeployERC20MirrorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeployERC20MirrorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDepoistERC20) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepoistERC20) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepoistERC20) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RecipientAddress) > 0 {
		i -= len(m.RecipientAddress)
		copy(dAtA[i:], m.RecipientAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RecipientAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.HomeERC20ContractAddress) > 0 {
		i -= len(m.HomeERC20ContractAddress)
		copy(dAtA[i:], m.HomeERC20ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.HomeERC20ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDepoistERC20Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepoistERC20Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepoistERC20Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgDeployERC20Mirror) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.HomeChain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.HomeERC20ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovTx(uint64(m.Decimals))
	}
	return n
}

func (m *MsgDeployERC20MirrorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDepoistERC20) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.HomeERC20ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RecipientAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDepoistERC20Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgDeployERC20Mirror) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeployERC20Mirror: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeployERC20Mirror: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomeChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HomeChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomeERC20ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HomeERC20ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeployERC20MirrorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeployERC20MirrorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeployERC20MirrorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDepoistERC20) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDepoistERC20: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepoistERC20: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomeERC20ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HomeERC20ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDepoistERC20Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDepoistERC20Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepoistERC20Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
