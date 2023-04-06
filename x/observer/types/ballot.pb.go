// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetacore/observer/ballot.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type VoteType int32

const (
	VoteType_SuccessObservation VoteType = 0
	VoteType_FailureObservation VoteType = 1
	VoteType_NotYetVoted        VoteType = 2
)

var VoteType_name = map[int32]string{
	0: "SuccessObservation",
	1: "FailureObservation",
	2: "NotYetVoted",
}

var VoteType_value = map[string]int32{
	"SuccessObservation": 0,
	"FailureObservation": 1,
	"NotYetVoted":        2,
}

func (x VoteType) String() string {
	return proto.EnumName(VoteType_name, int32(x))
}

func (VoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_934fc1e2777762f1, []int{0}
}

type BallotStatus int32

const (
	BallotStatus_BallotFinalized_SuccessObservation BallotStatus = 0
	BallotStatus_BallotFinalized_FailureObservation BallotStatus = 1
	BallotStatus_BallotInProgress                   BallotStatus = 2
)

var BallotStatus_name = map[int32]string{
	0: "BallotFinalized_SuccessObservation",
	1: "BallotFinalized_FailureObservation",
	2: "BallotInProgress",
}

var BallotStatus_value = map[string]int32{
	"BallotFinalized_SuccessObservation": 0,
	"BallotFinalized_FailureObservation": 1,
	"BallotInProgress":                   2,
}

func (x BallotStatus) String() string {
	return proto.EnumName(BallotStatus_name, int32(x))
}

func (BallotStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_934fc1e2777762f1, []int{1}
}

type Ballot struct {
	Index            string                                 `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	BallotIdentifier string                                 `protobuf:"bytes,2,opt,name=ballot_identifier,json=ballotIdentifier,proto3" json:"ballot_identifier,omitempty"`
	VoterList        []string                               `protobuf:"bytes,3,rep,name=voter_list,json=voterList,proto3" json:"voter_list,omitempty"`
	Votes            []VoteType                             `protobuf:"varint,4,rep,packed,name=votes,proto3,enum=zetacore.observer.VoteType" json:"votes,omitempty"`
	ObservationType  ObservationType                        `protobuf:"varint,5,opt,name=observation_type,json=observationType,proto3,enum=zetacore.observer.ObservationType" json:"observation_type,omitempty"`
	BallotThreshold  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=BallotThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"BallotThreshold"`
	BallotStatus     BallotStatus                           `protobuf:"varint,7,opt,name=ballot_status,json=ballotStatus,proto3,enum=zetacore.observer.BallotStatus" json:"ballot_status,omitempty"`
}

func (m *Ballot) Reset()         { *m = Ballot{} }
func (m *Ballot) String() string { return proto.CompactTextString(m) }
func (*Ballot) ProtoMessage()    {}
func (*Ballot) Descriptor() ([]byte, []int) {
	return fileDescriptor_934fc1e2777762f1, []int{0}
}
func (m *Ballot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ballot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ballot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ballot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ballot.Merge(m, src)
}
func (m *Ballot) XXX_Size() int {
	return m.Size()
}
func (m *Ballot) XXX_DiscardUnknown() {
	xxx_messageInfo_Ballot.DiscardUnknown(m)
}

var xxx_messageInfo_Ballot proto.InternalMessageInfo

func (m *Ballot) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Ballot) GetBallotIdentifier() string {
	if m != nil {
		return m.BallotIdentifier
	}
	return ""
}

func (m *Ballot) GetVoterList() []string {
	if m != nil {
		return m.VoterList
	}
	return nil
}

func (m *Ballot) GetVotes() []VoteType {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *Ballot) GetObservationType() ObservationType {
	if m != nil {
		return m.ObservationType
	}
	return ObservationType_EmptyObserverType
}

func (m *Ballot) GetBallotStatus() BallotStatus {
	if m != nil {
		return m.BallotStatus
	}
	return BallotStatus_BallotFinalized_SuccessObservation
}

func init() {
	proto.RegisterEnum("zetacore.observer.VoteType", VoteType_name, VoteType_value)
	proto.RegisterEnum("zetacore.observer.BallotStatus", BallotStatus_name, BallotStatus_value)
	proto.RegisterType((*Ballot)(nil), "zetacore.observer.Ballot")
}

func init() { proto.RegisterFile("zetacore/observer/ballot.proto", fileDescriptor_934fc1e2777762f1) }

var fileDescriptor_934fc1e2777762f1 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xb5, 0x9b, 0x26, 0x90, 0xa1, 0x34, 0xee, 0x28, 0x42, 0x56, 0x10, 0x13, 0x2b, 0x8b, 0x2a,
	0x2a, 0xaa, 0x2d, 0xca, 0x1f, 0x44, 0x55, 0xa5, 0x48, 0x3c, 0xdd, 0x0a, 0x01, 0x9b, 0xc8, 0x8f,
	0x4b, 0x32, 0xc2, 0xf5, 0x8d, 0x66, 0x26, 0x55, 0x9b, 0x6f, 0x60, 0xc1, 0x47, 0xb0, 0xe0, 0x53,
	0xba, 0xec, 0x12, 0xb1, 0xa8, 0x50, 0xf2, 0x23, 0xc8, 0x33, 0x75, 0x6b, 0xa8, 0x57, 0x33, 0xf7,
	0x9e, 0x33, 0xe7, 0xcc, 0x3d, 0xba, 0x84, 0x2d, 0x41, 0x45, 0x09, 0x0a, 0x08, 0x30, 0x96, 0x20,
	0xce, 0x40, 0x04, 0x71, 0x94, 0x65, 0xa8, 0xfc, 0xb9, 0x40, 0x85, 0x74, 0xa7, 0xc4, 0xfd, 0x12,
	0xef, 0x75, 0xa7, 0x38, 0x45, 0x8d, 0x06, 0xc5, 0xcd, 0x10, 0x7b, 0xde, 0x7d, 0xa1, 0xf2, 0x62,
	0x18, 0x83, 0x6f, 0x0d, 0xd2, 0x1a, 0x69, 0x6d, 0xda, 0x25, 0x4d, 0x9e, 0xa7, 0x70, 0xee, 0xda,
	0x9e, 0x3d, 0x6c, 0x87, 0xa6, 0xa0, 0xcf, 0xc9, 0x8e, 0xf1, 0x9e, 0xf0, 0x14, 0x72, 0xc5, 0xbf,
	0x70, 0x10, 0xee, 0x86, 0x66, 0x38, 0x06, 0x18, 0xdf, 0xf6, 0xe9, 0x33, 0x42, 0xce, 0x50, 0x81,
	0x98, 0x64, 0x5c, 0x2a, 0xb7, 0xe1, 0x35, 0x86, 0xed, 0xb0, 0xad, 0x3b, 0xaf, 0xb8, 0x54, 0xf4,
	0x05, 0x69, 0x16, 0x85, 0x74, 0x37, 0xbd, 0xc6, 0x70, 0xfb, 0xe0, 0xa9, 0x7f, 0x6f, 0x0e, 0xff,
	0x03, 0x2a, 0x38, 0xb9, 0x98, 0x43, 0x68, 0x98, 0xf4, 0x35, 0x71, 0x0c, 0x16, 0x29, 0x8e, 0xf9,
	0x44, 0x5d, 0xcc, 0xc1, 0x6d, 0x7a, 0xf6, 0x70, 0xfb, 0x60, 0x50, 0xf3, 0xfa, 0xed, 0x1d, 0x55,
	0x8b, 0x74, 0xf0, 0xdf, 0x06, 0xfd, 0x48, 0x3a, 0x66, 0xda, 0x93, 0x99, 0x00, 0x39, 0xc3, 0x2c,
	0x75, 0x5b, 0xc5, 0x2c, 0x23, 0xff, 0xf2, 0xba, 0x6f, 0xfd, 0xbe, 0xee, 0xef, 0x4e, 0xb9, 0x9a,
	0x2d, 0x62, 0x3f, 0xc1, 0xd3, 0x20, 0x41, 0x79, 0x8a, 0xf2, 0xe6, 0xd8, 0x97, 0xe9, 0xd7, 0xa0,
	0xb0, 0x97, 0xfe, 0x21, 0x24, 0xe1, 0xff, 0x32, 0xf4, 0x90, 0x3c, 0xbe, 0xc9, 0x49, 0xaa, 0x48,
	0x2d, 0xa4, 0xfb, 0x40, 0xff, 0xb2, 0x5f, 0xf3, 0x4b, 0xf3, 0xf4, 0x58, 0xd3, 0xc2, 0xad, 0xb8,
	0x52, 0xed, 0xbd, 0x27, 0x0f, 0xcb, 0x04, 0xe8, 0x13, 0x42, 0x8f, 0x17, 0x49, 0x02, 0x52, 0x56,
	0xc6, 0x72, 0xac, 0xa2, 0x7f, 0x14, 0xf1, 0x6c, 0x21, 0xa0, 0xda, 0xb7, 0x69, 0x87, 0x3c, 0x7a,
	0x83, 0xea, 0x13, 0xa8, 0x42, 0x21, 0x75, 0x36, 0x7a, 0x9b, 0x3f, 0x7f, 0x30, 0x7b, 0x6f, 0x49,
	0xb6, 0xaa, 0x86, 0x74, 0x97, 0x0c, 0x4c, 0x7d, 0xc4, 0xf3, 0x28, 0xe3, 0x4b, 0x48, 0x27, 0xb5,
	0x36, 0x35, 0xbc, 0x5a, 0xdb, 0x2e, 0x71, 0x0c, 0x6f, 0x9c, 0xbf, 0x13, 0x38, 0x15, 0x20, 0x65,
	0xe9, 0x3d, 0x1a, 0x5f, 0xae, 0x98, 0x7d, 0xb5, 0x62, 0xf6, 0x9f, 0x15, 0xb3, 0xbf, 0xaf, 0x99,
	0x75, 0xb5, 0x66, 0xd6, 0xaf, 0x35, 0xb3, 0x3e, 0x07, 0x95, 0x9c, 0x8b, 0x84, 0xf6, 0x93, 0x59,
	0xc4, 0xf3, 0xe0, 0x76, 0x5f, 0xcf, 0xef, 0x36, 0x56, 0x87, 0x1e, 0xb7, 0xf4, 0xbe, 0xbe, 0xfc,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x94, 0x6c, 0x58, 0xd7, 0x1c, 0x03, 0x00, 0x00,
}

func (m *Ballot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ballot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ballot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BallotStatus != 0 {
		i = encodeVarintBallot(dAtA, i, uint64(m.BallotStatus))
		i--
		dAtA[i] = 0x38
	}
	{
		size := m.BallotThreshold.Size()
		i -= size
		if _, err := m.BallotThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBallot(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.ObservationType != 0 {
		i = encodeVarintBallot(dAtA, i, uint64(m.ObservationType))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Votes) > 0 {
		dAtA2 := make([]byte, len(m.Votes)*10)
		var j1 int
		for _, num := range m.Votes {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintBallot(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.VoterList) > 0 {
		for iNdEx := len(m.VoterList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.VoterList[iNdEx])
			copy(dAtA[i:], m.VoterList[iNdEx])
			i = encodeVarintBallot(dAtA, i, uint64(len(m.VoterList[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.BallotIdentifier) > 0 {
		i -= len(m.BallotIdentifier)
		copy(dAtA[i:], m.BallotIdentifier)
		i = encodeVarintBallot(dAtA, i, uint64(len(m.BallotIdentifier)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintBallot(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBallot(dAtA []byte, offset int, v uint64) int {
	offset -= sovBallot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Ballot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovBallot(uint64(l))
	}
	l = len(m.BallotIdentifier)
	if l > 0 {
		n += 1 + l + sovBallot(uint64(l))
	}
	if len(m.VoterList) > 0 {
		for _, s := range m.VoterList {
			l = len(s)
			n += 1 + l + sovBallot(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		l = 0
		for _, e := range m.Votes {
			l += sovBallot(uint64(e))
		}
		n += 1 + sovBallot(uint64(l)) + l
	}
	if m.ObservationType != 0 {
		n += 1 + sovBallot(uint64(m.ObservationType))
	}
	l = m.BallotThreshold.Size()
	n += 1 + l + sovBallot(uint64(l))
	if m.BallotStatus != 0 {
		n += 1 + sovBallot(uint64(m.BallotStatus))
	}
	return n
}

func sovBallot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBallot(x uint64) (n int) {
	return sovBallot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ballot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBallot
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
			return fmt.Errorf("proto: Ballot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ballot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
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
				return ErrInvalidLengthBallot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBallot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
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
				return ErrInvalidLengthBallot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBallot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BallotIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoterList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
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
				return ErrInvalidLengthBallot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBallot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoterList = append(m.VoterList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v VoteType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBallot
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= VoteType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Votes = append(m.Votes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBallot
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBallot
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthBallot
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Votes) == 0 {
					m.Votes = make([]VoteType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v VoteType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBallot
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= VoteType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Votes = append(m.Votes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObservationType", wireType)
			}
			m.ObservationType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ObservationType |= ObservationType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
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
				return ErrInvalidLengthBallot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBallot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BallotThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BallotStatus", wireType)
			}
			m.BallotStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBallot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BallotStatus |= BallotStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBallot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBallot
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
func skipBallot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBallot
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
					return 0, ErrIntOverflowBallot
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
					return 0, ErrIntOverflowBallot
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
				return 0, ErrInvalidLengthBallot
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBallot
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBallot
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBallot        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBallot          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBallot = fmt.Errorf("proto: unexpected end of group")
)
