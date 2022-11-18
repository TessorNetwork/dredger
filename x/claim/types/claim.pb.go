// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/claim/claim.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Action int32

const (
	ActionFree          Action = 0
	ActionLiquidStake   Action = 1
	ActionDelegateStake Action = 2
)

var Action_name = map[int32]string{
	0: "ActionFree",
	1: "ActionLiquidStake",
	2: "ActionDelegateStake",
}

var Action_value = map[string]int32{
	"ActionFree":          0,
	"ActionLiquidStake":   1,
	"ActionDelegateStake": 2,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4747d999b9dc0da, []int{0}
}

// A Claim Records is the metadata of claim data per address
type ClaimRecord struct {
	// airdrop identifier
	AirdropIdentifier string `protobuf:"bytes,1,opt,name=airdrop_identifier,json=airdropIdentifier,proto3" json:"airdrop_identifier,omitempty" yaml:"airdrop_identifier"`
	// address of claim user
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// weight that represent the portion from total allocation
	Weight github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight" yaml:"weight"`
	// true if action is completed
	// index of bool in array refers to action enum #
	ActionCompleted []bool `protobuf:"varint,4,rep,packed,name=action_completed,json=actionCompleted,proto3" json:"action_completed,omitempty" yaml:"action_completed"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4747d999b9dc0da, []int{0}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetAirdropIdentifier() string {
	if m != nil {
		return m.AirdropIdentifier
	}
	return ""
}

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetActionCompleted() []bool {
	if m != nil {
		return m.ActionCompleted
	}
	return nil
}

func init() {
	proto.RegisterEnum("stride.claim.Action", Action_name, Action_value)
	proto.RegisterType((*ClaimRecord)(nil), "stride.claim.ClaimRecord")
}

func init() { proto.RegisterFile("stride/claim/claim.proto", fileDescriptor_b4747d999b9dc0da) }

var fileDescriptor_b4747d999b9dc0da = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x6a, 0x14, 0x31,
	0x18, 0xc7, 0x67, 0xb6, 0x65, 0xd5, 0xa8, 0x75, 0x1b, 0x95, 0x8e, 0x2b, 0x66, 0xea, 0x1c, 0xa4,
	0x88, 0x9d, 0x50, 0x7a, 0xf3, 0x22, 0x6e, 0x4b, 0x51, 0x58, 0x3c, 0x4c, 0x0f, 0x82, 0x97, 0x92,
	0x49, 0x3e, 0xa7, 0x61, 0x77, 0x26, 0x6b, 0x92, 0x56, 0xfb, 0x06, 0x1e, 0x7d, 0x07, 0x6f, 0x3e,
	0x49, 0x8f, 0x3d, 0x8a, 0x87, 0x41, 0x76, 0xdf, 0x60, 0x9e, 0x40, 0x9a, 0x64, 0x17, 0xd1, 0x4b,
	0xf2, 0xe5, 0xff, 0xfb, 0xf3, 0x4f, 0x3e, 0xf2, 0xa1, 0xc4, 0x58, 0x2d, 0x05, 0x50, 0x3e, 0x65,
	0xb2, 0xf6, 0x6b, 0x3e, 0xd3, 0xca, 0x2a, 0x7c, 0xc7, 0x93, 0xdc, 0x69, 0xc3, 0x07, 0x95, 0xaa,
	0x94, 0x03, 0xf4, 0xba, 0xf2, 0x9e, 0x21, 0xe1, 0xca, 0xd4, 0xca, 0xd0, 0x92, 0x19, 0xa0, 0xe7,
	0x7b, 0x25, 0x58, 0xb6, 0x47, 0xb9, 0x92, 0x4d, 0xe0, 0x4f, 0x57, 0xbc, 0x99, 0xac, 0x78, 0x05,
	0x0d, 0x18, 0x69, 0xbc, 0x25, 0xfb, 0xd1, 0x43, 0xb7, 0x0f, 0xae, 0xaf, 0x28, 0x80, 0x2b, 0x2d,
	0xf0, 0x18, 0x61, 0x26, 0xb5, 0xd0, 0x6a, 0x76, 0x22, 0x05, 0x34, 0x56, 0x7e, 0x94, 0xa0, 0x93,
	0x78, 0x3b, 0xde, 0xb9, 0x35, 0x7a, 0xd2, 0xb5, 0xe9, 0xa3, 0x0b, 0x56, 0x4f, 0x5f, 0x66, 0xff,
	0x7b, 0xb2, 0x62, 0x33, 0x88, 0x6f, 0x57, 0x1a, 0x7e, 0x81, 0x6e, 0x30, 0x21, 0x34, 0x18, 0x93,
	0xf4, 0x5c, 0x04, 0xee, 0xda, 0x74, 0x23, 0x44, 0x78, 0x90, 0x15, 0x4b, 0x0b, 0x7e, 0x8f, 0xfa,
	0x9f, 0x41, 0x56, 0xa7, 0x36, 0x59, 0x73, 0xe6, 0x57, 0x97, 0x6d, 0x1a, 0xfd, 0x6a, 0xd3, 0x67,
	0x95, 0xb4, 0xa7, 0x67, 0x65, 0xce, 0x55, 0x4d, 0x43, 0x47, 0x7e, 0xdb, 0x35, 0x62, 0x42, 0xed,
	0xc5, 0x0c, 0x4c, 0x7e, 0x08, 0xbc, 0x6b, 0xd3, 0xbb, 0x3e, 0xda, 0xa7, 0x64, 0x45, 0x88, 0xc3,
	0x47, 0x68, 0xc0, 0xb8, 0x95, 0xaa, 0x39, 0xe1, 0xaa, 0x9e, 0x4d, 0xc1, 0x82, 0x48, 0xd6, 0xb7,
	0xd7, 0x76, 0x6e, 0x8e, 0x1e, 0x77, 0x6d, 0xba, 0x15, 0xde, 0xf3, 0x8f, 0x23, 0x2b, 0xee, 0x79,
	0xe9, 0x60, 0xa9, 0x3c, 0x7f, 0x87, 0xfa, 0xaf, 0x9d, 0x84, 0x37, 0x10, 0xf2, 0xd5, 0x91, 0x06,
	0x18, 0x44, 0xf8, 0x21, 0xda, 0xf4, 0xe7, 0xb1, 0xfc, 0x74, 0x26, 0xc5, 0xb1, 0x65, 0x13, 0x18,
	0xc4, 0x78, 0x0b, 0xdd, 0xf7, 0xf2, 0x21, 0x4c, 0xa1, 0x62, 0x16, 0x3c, 0xe8, 0x0d, 0xd7, 0xbf,
	0x7e, 0x27, 0xd1, 0xe8, 0xcd, 0xe5, 0x9c, 0xc4, 0x57, 0x73, 0x12, 0xff, 0x9e, 0x93, 0xf8, 0xdb,
	0x82, 0x44, 0x57, 0x0b, 0x12, 0xfd, 0x5c, 0x90, 0xe8, 0x43, 0xfe, 0x57, 0xcb, 0xc7, 0x6e, 0x10,
	0x76, 0xc7, 0xac, 0x34, 0x34, 0x8c, 0xcb, 0xf9, 0x3e, 0xfd, 0x12, 0x66, 0xc6, 0xb5, 0x5f, 0xf6,
	0xdd, 0x6f, 0xee, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x3b, 0x56, 0xf7, 0x50, 0x02, 0x00,
	0x00,
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionCompleted) > 0 {
		for iNdEx := len(m.ActionCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaim(dAtA, i, uint64(len(m.ActionCompleted)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaim(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AirdropIdentifier) > 0 {
		i -= len(m.AirdropIdentifier)
		copy(dAtA[i:], m.AirdropIdentifier)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.AirdropIdentifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AirdropIdentifier)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovClaim(uint64(l))
	if len(m.ActionCompleted) > 0 {
		n += 1 + sovClaim(uint64(len(m.ActionCompleted))) + len(m.ActionCompleted)*1
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AirdropIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
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
				m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
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
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionCompleted) == 0 {
					m.ActionCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
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
					m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)
