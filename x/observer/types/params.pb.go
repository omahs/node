// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: observer/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "github.com/zeta-chain/zetacore/common"
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

type Policy_Type int32

const (
	Policy_Type_stop_inbound_cctx    Policy_Type = 0
	Policy_Type_deploy_fungible_coin Policy_Type = 1
	Policy_Type_update_client_params Policy_Type = 2
)

var Policy_Type_name = map[int32]string{
	0: "stop_inbound_cctx",
	1: "deploy_fungible_coin",
	2: "update_client_params",
}

var Policy_Type_value = map[string]int32{
	"stop_inbound_cctx":    0,
	"deploy_fungible_coin": 1,
	"update_client_params": 2,
}

func (x Policy_Type) String() string {
	return proto.EnumName(Policy_Type_name, int32(x))
}

func (Policy_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{0}
}

type ClientParams struct {
	ConfirmationCount int64 `protobuf:"varint,1,opt,name=confirmation_count,json=confirmationCount,proto3" json:"confirmation_count,omitempty"`
	GasPriceTicker    int64 `protobuf:"varint,2,opt,name=gas_price_ticker,json=gasPriceTicker,proto3" json:"gas_price_ticker,omitempty"`
}

func (m *ClientParams) Reset()         { *m = ClientParams{} }
func (m *ClientParams) String() string { return proto.CompactTextString(m) }
func (*ClientParams) ProtoMessage()    {}
func (*ClientParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{0}
}
func (m *ClientParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientParams.Merge(m, src)
}
func (m *ClientParams) XXX_Size() int {
	return m.Size()
}
func (m *ClientParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientParams.DiscardUnknown(m)
}

var xxx_messageInfo_ClientParams proto.InternalMessageInfo

func (m *ClientParams) GetConfirmationCount() int64 {
	if m != nil {
		return m.ConfirmationCount
	}
	return 0
}

func (m *ClientParams) GetGasPriceTicker() int64 {
	if m != nil {
		return m.GasPriceTicker
	}
	return 0
}

type ObserverParams struct {
	Chain                 *common.Chain                          `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	BallotThreshold       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=ballot_threshold,json=ballotThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"ballot_threshold"`
	MinObserverDelegation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=min_observer_delegation,json=minObserverDelegation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_observer_delegation"`
	IsSupported           bool                                   `protobuf:"varint,5,opt,name=is_supported,json=isSupported,proto3" json:"is_supported,omitempty"`
	ClientParams          *ClientParams                          `protobuf:"bytes,6,opt,name=client_params,json=clientParams,proto3" json:"client_params,omitempty"`
}

func (m *ObserverParams) Reset()         { *m = ObserverParams{} }
func (m *ObserverParams) String() string { return proto.CompactTextString(m) }
func (*ObserverParams) ProtoMessage()    {}
func (*ObserverParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{1}
}
func (m *ObserverParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObserverParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObserverParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObserverParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObserverParams.Merge(m, src)
}
func (m *ObserverParams) XXX_Size() int {
	return m.Size()
}
func (m *ObserverParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ObserverParams.DiscardUnknown(m)
}

var xxx_messageInfo_ObserverParams proto.InternalMessageInfo

func (m *ObserverParams) GetChain() *common.Chain {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *ObserverParams) GetIsSupported() bool {
	if m != nil {
		return m.IsSupported
	}
	return false
}

func (m *ObserverParams) GetClientParams() *ClientParams {
	if m != nil {
		return m.ClientParams
	}
	return nil
}

type Admin_Policy struct {
	PolicyType Policy_Type `protobuf:"varint,1,opt,name=policy_type,json=policyType,proto3,enum=zetachain.zetacore.observer.Policy_Type" json:"policy_type,omitempty"`
	Address    string      `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Admin_Policy) Reset()         { *m = Admin_Policy{} }
func (m *Admin_Policy) String() string { return proto.CompactTextString(m) }
func (*Admin_Policy) ProtoMessage()    {}
func (*Admin_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{2}
}
func (m *Admin_Policy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Admin_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Admin_Policy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Admin_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Admin_Policy.Merge(m, src)
}
func (m *Admin_Policy) XXX_Size() int {
	return m.Size()
}
func (m *Admin_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Admin_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Admin_Policy proto.InternalMessageInfo

func (m *Admin_Policy) GetPolicyType() Policy_Type {
	if m != nil {
		return m.PolicyType
	}
	return Policy_Type_stop_inbound_cctx
}

func (m *Admin_Policy) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// Params defines the parameters for the module.
type Params struct {
	ObserverParams []*ObserverParams `protobuf:"bytes,1,rep,name=observer_params,json=observerParams,proto3" json:"observer_params,omitempty"`
	AdminPolicy    []*Admin_Policy   `protobuf:"bytes,2,rep,name=admin_policy,json=adminPolicy,proto3" json:"admin_policy,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4542fa62877488a1, []int{3}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetObserverParams() []*ObserverParams {
	if m != nil {
		return m.ObserverParams
	}
	return nil
}

func (m *Params) GetAdminPolicy() []*Admin_Policy {
	if m != nil {
		return m.AdminPolicy
	}
	return nil
}

func init() {
	proto.RegisterEnum("zetachain.zetacore.observer.Policy_Type", Policy_Type_name, Policy_Type_value)
	proto.RegisterType((*ClientParams)(nil), "zetachain.zetacore.observer.ClientParams")
	proto.RegisterType((*ObserverParams)(nil), "zetachain.zetacore.observer.ObserverParams")
	proto.RegisterType((*Admin_Policy)(nil), "zetachain.zetacore.observer.Admin_Policy")
	proto.RegisterType((*Params)(nil), "zetachain.zetacore.observer.Params")
}

func init() { proto.RegisterFile("observer/params.proto", fileDescriptor_4542fa62877488a1) }

var fileDescriptor_4542fa62877488a1 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xb7, 0xd3, 0x3f, 0xd0, 0x73, 0x9a, 0xa6, 0x47, 0xab, 0x5a, 0x45, 0x72, 0x43, 0x90, 0x90,
	0x01, 0xc5, 0x96, 0xc2, 0xc6, 0x46, 0xd3, 0xa5, 0x12, 0x82, 0xc8, 0x64, 0x81, 0x81, 0x93, 0x7d,
	0xbe, 0x38, 0xa7, 0xda, 0x77, 0x96, 0xef, 0x8c, 0x1a, 0x3e, 0x05, 0x23, 0x23, 0x03, 0x03, 0x33,
	0x9f, 0xa2, 0x63, 0x47, 0xc4, 0x50, 0xa1, 0x64, 0xe4, 0x4b, 0x20, 0x9f, 0xed, 0xc8, 0x59, 0x32,
	0x30, 0xdd, 0xbb, 0xf7, 0xee, 0xfd, 0xde, 0xfb, 0xbd, 0xdf, 0x3b, 0x70, 0xcc, 0x03, 0x41, 0xb2,
	0x4f, 0x24, 0x73, 0x53, 0x3f, 0xf3, 0x13, 0xe1, 0xa4, 0x19, 0x97, 0x1c, 0x3e, 0xfc, 0x4c, 0xa4,
	0x8f, 0x67, 0x3e, 0x65, 0x8e, 0xb2, 0x78, 0x46, 0x9c, 0xfa, 0xe5, 0xe9, 0x51, 0xc4, 0x23, 0xae,
	0xde, 0xb9, 0x85, 0x55, 0xa6, 0x9c, 0x9e, 0xac, 0x90, 0x6a, 0xa3, 0x0a, 0x3c, 0xc0, 0x3c, 0x49,
	0x38, 0x73, 0xcb, 0xa3, 0x74, 0xf6, 0x23, 0xd0, 0x1e, 0xc5, 0x94, 0x30, 0x39, 0x56, 0x65, 0xe1,
	0x00, 0x40, 0xcc, 0xd9, 0x94, 0x66, 0x89, 0x2f, 0x29, 0x67, 0x08, 0xf3, 0x9c, 0x49, 0x53, 0xef,
	0xe9, 0xf6, 0x96, 0x77, 0xd8, 0x8c, 0x8c, 0x8a, 0x00, 0xb4, 0x41, 0x37, 0xf2, 0x05, 0x4a, 0x33,
	0x8a, 0x09, 0x92, 0x14, 0x5f, 0x91, 0xcc, 0x6c, 0xa9, 0xc7, 0x9d, 0xc8, 0x17, 0xe3, 0xc2, 0x3d,
	0x51, 0xde, 0xfe, 0xdf, 0x16, 0xe8, 0xbc, 0xad, 0x1a, 0xaa, 0x6a, 0x3d, 0x06, 0x3b, 0x8a, 0x9a,
	0x82, 0x37, 0x86, 0xfb, 0x4e, 0xd5, 0xd9, 0xa8, 0x70, 0x7a, 0x65, 0x0c, 0xbe, 0x07, 0xdd, 0xc0,
	0x8f, 0x63, 0x2e, 0x91, 0x9c, 0x65, 0x44, 0xcc, 0x78, 0x1c, 0x9a, 0x5b, 0x3d, 0xdd, 0xde, 0x3b,
	0x77, 0x6e, 0xee, 0xce, 0xb4, 0xdf, 0x77, 0x67, 0x4f, 0x22, 0x2a, 0x67, 0x79, 0x50, 0x64, 0xbb,
	0x98, 0x8b, 0x84, 0x8b, 0xea, 0x18, 0x88, 0xf0, 0xca, 0x95, 0xf3, 0x94, 0x08, 0xe7, 0x82, 0x60,
	0xef, 0xa0, 0xc4, 0x99, 0xd4, 0x30, 0x70, 0x0a, 0x4e, 0x12, 0xca, 0x50, 0x3d, 0x26, 0x14, 0x92,
	0x98, 0x44, 0x8a, 0x9c, 0xb9, 0xfd, 0x5f, 0x15, 0x8e, 0x13, 0xca, 0x6a, 0x8e, 0x17, 0x2b, 0x30,
	0xf8, 0x08, 0xb4, 0xa9, 0x40, 0x22, 0x4f, 0x53, 0x9e, 0x49, 0x12, 0x9a, 0x3b, 0x3d, 0xdd, 0xbe,
	0xef, 0x19, 0x54, 0xbc, 0xab, 0x5d, 0xf0, 0x0d, 0xd8, 0xc7, 0x4a, 0x06, 0x54, 0xca, 0x6f, 0xee,
	0xaa, 0x91, 0x3c, 0x75, 0x36, 0xe8, 0xef, 0x34, 0x85, 0xf3, 0xda, 0xb8, 0x71, 0xeb, 0x0b, 0xd0,
	0x7e, 0x15, 0x16, 0xe4, 0xc6, 0x3c, 0xa6, 0x78, 0x0e, 0x2f, 0x81, 0x91, 0x2a, 0x0b, 0x15, 0xdd,
	0xaa, 0x81, 0x77, 0x86, 0xf6, 0x46, 0xf4, 0x32, 0x13, 0x4d, 0xe6, 0x29, 0xf1, 0x40, 0x99, 0x5c,
	0xd8, 0xd0, 0x04, 0xf7, 0xfc, 0x30, 0xcc, 0x88, 0x10, 0x4a, 0xe9, 0x3d, 0xaf, 0xbe, 0xf6, 0x7f,
	0xea, 0x60, 0xb7, 0x92, 0x76, 0x02, 0x0e, 0x56, 0x63, 0xad, 0x18, 0xe9, 0xbd, 0x2d, 0xdb, 0x18,
	0x3e, 0xdf, 0x58, 0x73, 0x7d, 0x41, 0xbc, 0x0e, 0x5f, 0x5f, 0x98, 0xd7, 0xa0, 0xed, 0x2b, 0x56,
	0x65, 0x3b, 0x66, 0x4b, 0x41, 0x6e, 0x1e, 0x52, 0x73, 0x0c, 0x9e, 0xa1, 0xd2, 0xcb, 0xcb, 0xcb,
	0xed, 0xaf, 0xdf, 0xce, 0xb4, 0x67, 0x1f, 0x81, 0xd1, 0x60, 0x0a, 0x8f, 0xc1, 0xa1, 0x90, 0x3c,
	0x45, 0x94, 0x05, 0x3c, 0x67, 0x21, 0xc2, 0x58, 0x5e, 0x77, 0x35, 0x68, 0x82, 0xa3, 0x90, 0xa4,
	0x31, 0x9f, 0xa3, 0x69, 0xce, 0x22, 0x1a, 0xc4, 0x04, 0x61, 0x4e, 0x59, 0x57, 0x2f, 0x22, 0x79,
	0x1a, 0xfa, 0x92, 0xa0, 0x35, 0x01, 0xbb, 0xad, 0xd3, 0xed, 0x1f, 0xdf, 0x2d, 0xfd, 0xfc, 0xf2,
	0x66, 0x61, 0xe9, 0xb7, 0x0b, 0x4b, 0xff, 0xb3, 0xb0, 0xf4, 0x2f, 0x4b, 0x4b, 0xbb, 0x5d, 0x5a,
	0xda, 0xaf, 0xa5, 0xa5, 0x7d, 0x70, 0x1b, 0x5b, 0x55, 0xf4, 0x3d, 0x50, 0x14, 0xdc, 0x9a, 0x82,
	0x7b, 0xbd, 0xfa, 0xc0, 0xe5, 0x8a, 0x05, 0xbb, 0xea, 0xcb, 0xbe, 0xf8, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x75, 0x8c, 0xa3, 0x5e, 0x2c, 0x04, 0x00, 0x00,
}

func (m *ClientParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasPriceTicker != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.GasPriceTicker))
		i--
		dAtA[i] = 0x10
	}
	if m.ConfirmationCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ConfirmationCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ObserverParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObserverParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObserverParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClientParams != nil {
		{
			size, err := m.ClientParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.IsSupported {
		i--
		if m.IsSupported {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.MinObserverDelegation.Size()
		i -= size
		if _, err := m.MinObserverDelegation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.BallotThreshold.Size()
		i -= size
		if _, err := m.BallotThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Chain != nil {
		{
			size, err := m.Chain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Admin_Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Admin_Policy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Admin_Policy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.PolicyType != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PolicyType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AdminPolicy) > 0 {
		for iNdEx := len(m.AdminPolicy) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AdminPolicy[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ObserverParams) > 0 {
		for iNdEx := len(m.ObserverParams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ObserverParams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConfirmationCount != 0 {
		n += 1 + sovParams(uint64(m.ConfirmationCount))
	}
	if m.GasPriceTicker != 0 {
		n += 1 + sovParams(uint64(m.GasPriceTicker))
	}
	return n
}

func (m *ObserverParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Chain != nil {
		l = m.Chain.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.BallotThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinObserverDelegation.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.IsSupported {
		n += 2
	}
	if m.ClientParams != nil {
		l = m.ClientParams.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Admin_Policy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PolicyType != 0 {
		n += 1 + sovParams(uint64(m.PolicyType))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ObserverParams) > 0 {
		for _, e := range m.ObserverParams {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.AdminPolicy) > 0 {
		for _, e := range m.AdminPolicy {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: ClientParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmationCount", wireType)
			}
			m.ConfirmationCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConfirmationCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPriceTicker", wireType)
			}
			m.GasPriceTicker = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPriceTicker |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *ObserverParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: ObserverParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObserverParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Chain == nil {
				m.Chain = &common.Chain{}
			}
			if err := m.Chain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BallotThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinObserverDelegation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinObserverDelegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSupported", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsSupported = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClientParams == nil {
				m.ClientParams = &ClientParams{}
			}
			if err := m.ClientParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Admin_Policy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Admin_Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Admin_Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PolicyType", wireType)
			}
			m.PolicyType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PolicyType |= Policy_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObserverParams = append(m.ObserverParams, &ObserverParams{})
			if err := m.ObserverParams[len(m.ObserverParams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminPolicy = append(m.AdminPolicy, &Admin_Policy{})
			if err := m.AdminPolicy[len(m.AdminPolicy)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
