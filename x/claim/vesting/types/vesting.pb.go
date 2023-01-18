// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dredger/vesting/vesting.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
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

// BaseVestingAccount implements the VestingAccount interface. It contains all
// the necessary fields needed for any vesting account implementation.
type BaseVestingAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty"`
	OriginalVesting    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=original_vesting,json=originalVesting,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"original_vesting" yaml:"original_vesting"`
	DelegatedFree      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=delegated_free,json=delegatedFree,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"delegated_free" yaml:"delegated_free"`
	DelegatedVesting   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=delegated_vesting,json=delegatedVesting,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"delegated_vesting" yaml:"delegated_vesting"`
	EndTime            int64                                    `protobuf:"varint,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" yaml:"end_time"`
}

func (m *BaseVestingAccount) Reset()      { *m = BaseVestingAccount{} }
func (*BaseVestingAccount) ProtoMessage() {}
func (*BaseVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f0278a453c26b3, []int{0}
}
func (m *BaseVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseVestingAccount.Merge(m, src)
}
func (m *BaseVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *BaseVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_BaseVestingAccount proto.InternalMessageInfo

// Period defines a length of time and amount of coins that will vest.
type Period struct {
	StartTime  int64                                    `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Length     int64                                    `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Amount     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	ActionType int32                                    `protobuf:"varint,4,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
}

func (m *Period) Reset()      { *m = Period{} }
func (*Period) ProtoMessage() {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f0278a453c26b3, []int{1}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Period) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Period) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Period) GetActionType() int32 {
	if m != nil {
		return m.ActionType
	}
	return 0
}

// StridePeriodicVestingAccount implements the VestingAccount interface. It
// periodically vests by unlocking coins during each specified period.
type StridePeriodicVestingAccount struct {
	*BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	VestingPeriods      []Period `protobuf:"bytes,3,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods" yaml:"vesting_periods"`
}

func (m *StridePeriodicVestingAccount) Reset()      { *m = StridePeriodicVestingAccount{} }
func (*StridePeriodicVestingAccount) ProtoMessage() {}
func (*StridePeriodicVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f0278a453c26b3, []int{2}
}
func (m *StridePeriodicVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StridePeriodicVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StridePeriodicVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StridePeriodicVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StridePeriodicVestingAccount.Merge(m, src)
}
func (m *StridePeriodicVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *StridePeriodicVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_StridePeriodicVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_StridePeriodicVestingAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseVestingAccount)(nil), "dredger.vesting.BaseVestingAccount")
	proto.RegisterType((*Period)(nil), "dredger.vesting.Period")
	proto.RegisterType((*StridePeriodicVestingAccount)(nil), "dredger.vesting.StridePeriodicVestingAccount")
}

func init() { proto.RegisterFile("dredger/vesting/vesting.proto", fileDescriptor_41f0278a453c26b3) }

var fileDescriptor_41f0278a453c26b3 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x41, 0x6f, 0xd3, 0x3e,
	0x1c, 0x8d, 0xd7, 0xae, 0xff, 0xfd, 0x5d, 0x58, 0x47, 0x18, 0x25, 0x4c, 0x23, 0xa9, 0x72, 0xea,
	0x65, 0x09, 0x1b, 0x88, 0x43, 0x6f, 0x04, 0x09, 0x69, 0x62, 0x87, 0x29, 0x4c, 0x1c, 0x76, 0x89,
	0x9c, 0xc4, 0xa4, 0x16, 0x4d, 0x5c, 0xc5, 0xee, 0x44, 0xbf, 0x01, 0x12, 0x17, 0x90, 0x38, 0x70,
	0xdc, 0x99, 0x4f, 0x32, 0x89, 0x4b, 0xc5, 0x89, 0x53, 0x41, 0xad, 0xf8, 0x02, 0xfb, 0x04, 0x28,
	0xb6, 0xd3, 0xb2, 0xec, 0x50, 0x8d, 0x53, 0x62, 0xff, 0xfc, 0x9e, 0xdf, 0xef, 0xd9, 0xcf, 0x70,
	0x97, 0xf1, 0x9c, 0xc4, 0xd8, 0x3d, 0xc3, 0x8c, 0x93, 0x2c, 0x29, 0xbf, 0xce, 0x30, 0xa7, 0x9c,
	0xea, 0x9b, 0xb2, 0xea, 0xa8, 0xd9, 0x9d, 0xed, 0x84, 0x26, 0x54, 0x94, 0xdc, 0xe2, 0x4f, 0xae,
	0xda, 0x31, 0x23, 0xca, 0x52, 0xca, 0xdc, 0x10, 0x31, 0xec, 0x9e, 0xed, 0x87, 0x98, 0xa3, 0x7d,
	0x37, 0xa2, 0x24, 0xab, 0xd4, 0xd1, 0x88, 0xf7, 0x17, 0xf5, 0x62, 0x20, 0xeb, 0xf6, 0xf7, 0x3a,
	0xd4, 0x3d, 0xc4, 0xf0, 0x6b, 0xb9, 0xcb, 0xb3, 0x28, 0xa2, 0xa3, 0x8c, 0xeb, 0x87, 0xf0, 0x56,
	0xc1, 0x18, 0x20, 0x39, 0x36, 0x40, 0x07, 0x74, 0x9b, 0x07, 0x1d, 0x47, 0xb2, 0x39, 0x82, 0x40,
	0xb1, 0x39, 0x05, 0x5c, 0xe1, 0xbc, 0xfa, 0x64, 0x6a, 0x01, 0xbf, 0x19, 0x2e, 0xa7, 0xf4, 0x4f,
	0x00, 0x6e, 0xd1, 0x9c, 0x24, 0x24, 0x43, 0x83, 0x40, 0x35, 0x63, 0xac, 0x75, 0x6a, 0xdd, 0xe6,
	0xc1, 0x83, 0x92, 0xaf, 0x58, 0xbf, 0xe0, 0x7b, 0x4e, 0x49, 0xe6, 0xbd, 0xbc, 0x98, 0x5a, 0xda,
	0xe5, 0xd4, 0xba, 0x3f, 0x46, 0xe9, 0xa0, 0x67, 0x57, 0x09, 0xec, 0xaf, 0x3f, 0xad, 0x6e, 0x42,
	0x78, 0x7f, 0x14, 0x3a, 0x11, 0x4d, 0x5d, 0xd5, 0xa5, 0xfc, 0xec, 0xb1, 0xf8, 0xad, 0xcb, 0xc7,
	0x43, 0xcc, 0x04, 0x17, 0xf3, 0x5b, 0x25, 0x5c, 0x75, 0xa9, 0x7f, 0x00, 0x70, 0x33, 0xc6, 0x03,
	0x9c, 0x20, 0x8e, 0xe3, 0xe0, 0x4d, 0x8e, 0xb1, 0x51, 0x5b, 0xa5, 0xe8, 0x50, 0x29, 0xba, 0x27,
	0x15, 0x5d, 0x85, 0xdf, 0x4c, 0xcf, 0xed, 0x05, 0xf8, 0x45, 0x8e, 0xb1, 0xfe, 0x19, 0xc0, 0x3b,
	0x4b, 0xba, 0xd2, 0xa2, 0xfa, 0x2a, 0x41, 0x47, 0x4a, 0x90, 0x51, 0x15, 0xf4, 0x4f, 0x1e, 0x6d,
	0x2d, 0xf0, 0xa5, 0x49, 0x0e, 0xdc, 0xc0, 0x59, 0x1c, 0x70, 0x92, 0x62, 0x63, 0xbd, 0x03, 0xba,
	0x35, 0xef, 0xee, 0xe5, 0xd4, 0x6a, 0xc9, 0xdd, 0xca, 0x8a, 0xed, 0xff, 0x87, 0xb3, 0xf8, 0x84,
	0xa4, 0xb8, 0xb7, 0xf1, 0xfe, 0xdc, 0xd2, 0xbe, 0x9c, 0x5b, 0x9a, 0xfd, 0x0d, 0xc0, 0xc6, 0x31,
	0xce, 0x09, 0x8d, 0xf5, 0x87, 0x10, 0x32, 0x8e, 0x72, 0x2e, 0x69, 0x8a, 0x6b, 0x54, 0xf3, 0xff,
	0x17, 0x33, 0x05, 0x46, 0x6f, 0xc3, 0xc6, 0x00, 0x67, 0x09, 0xef, 0x1b, 0x6b, 0xa2, 0xa4, 0x46,
	0x7a, 0x04, 0x1b, 0x28, 0x15, 0x37, 0x6f, 0xe5, 0xb9, 0x3c, 0x2a, 0x6c, 0xb8, 0x51, 0xab, 0x8a,
	0x5a, 0xb7, 0x60, 0x13, 0x45, 0x9c, 0xd0, 0x2c, 0x28, 0xaa, 0x46, 0xbd, 0x03, 0xba, 0xeb, 0x3e,
	0x94, 0x53, 0x27, 0xe3, 0x21, 0xee, 0xd5, 0x45, 0x37, 0xbf, 0x01, 0xdc, 0x7d, 0x25, 0xb2, 0x28,
	0x7b, 0x22, 0x51, 0x25, 0x2c, 0xa7, 0x70, 0x5b, 0x84, 0x45, 0xf9, 0x5e, 0x09, 0x8d, 0xed, 0x5c,
	0x0d, 0xb2, 0x73, 0x3d, 0x6e, 0x2a, 0x36, 0x7a, 0x78, 0x3d, 0x88, 0x01, 0x6c, 0x95, 0xb4, 0x43,
	0xb1, 0x3b, 0x53, 0x8e, 0xb4, 0xab, 0xb4, 0x52, 0x9c, 0x67, 0xaa, 0x5b, 0xd1, 0x96, 0xe7, 0x54,
	0x01, 0xdb, 0xfe, 0xa6, 0x9a, 0x91, 0xcb, 0xd9, 0xf2, 0xd4, 0xbc, 0xe3, 0x8b, 0x99, 0x09, 0x26,
	0x33, 0x13, 0xfc, 0x9a, 0x99, 0xe0, 0xe3, 0xdc, 0xd4, 0x26, 0x73, 0x53, 0xfb, 0x31, 0x37, 0xb5,
	0xd3, 0xa7, 0x7f, 0x59, 0x2b, 0x9d, 0xd8, 0x3b, 0x42, 0x21, 0x73, 0xcb, 0xf7, 0xeb, 0x89, 0xfb,
	0xce, 0x8d, 0x06, 0x88, 0xa4, 0x8b, 0xa7, 0x4c, 0xd8, 0x1d, 0x36, 0xc4, 0x1b, 0xf3, 0xf8, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x19, 0x40, 0xa0, 0xe9, 0x04, 0x00, 0x00,
}

func (m *BaseVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DelegatedVesting) > 0 {
		for iNdEx := len(m.DelegatedVesting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatedVesting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.DelegatedFree) > 0 {
		for iNdEx := len(m.DelegatedFree) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatedFree[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OriginalVesting) > 0 {
		for iNdEx := len(m.OriginalVesting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OriginalVesting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Period) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ActionType != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.ActionType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Length != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x10
	}
	if m.StartTime != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StridePeriodicVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StridePeriodicVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StridePeriodicVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if len(m.OriginalVesting) > 0 {
		for _, e := range m.OriginalVesting {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.DelegatedFree) > 0 {
		for _, e := range m.DelegatedFree {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.DelegatedVesting) > 0 {
		for _, e := range m.DelegatedVesting {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if m.EndTime != 0 {
		n += 1 + sovVesting(uint64(m.EndTime))
	}
	return n
}

func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartTime != 0 {
		n += 1 + sovVesting(uint64(m.StartTime))
	}
	if m.Length != 0 {
		n += 1 + sovVesting(uint64(m.Length))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if m.ActionType != 0 {
		n += 1 + sovVesting(uint64(m.ActionType))
	}
	return n
}

func (m *StridePeriodicVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func sovVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVesting(x uint64) (n int) {
	return sovVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: BaseVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalVesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalVesting = append(m.OriginalVesting, types1.Coin{})
			if err := m.OriginalVesting[len(m.OriginalVesting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatedFree", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatedFree = append(m.DelegatedFree, types1.Coin{})
			if err := m.DelegatedFree[len(m.DelegatedFree)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatedVesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatedVesting = append(m.DelegatedVesting, types1.Coin{})
			if err := m.DelegatedVesting[len(m.DelegatedVesting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types1.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionType", wireType)
			}
			m.ActionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActionType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func (m *StridePeriodicVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: StridePeriodicVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StridePeriodicVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, Period{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func skipVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
				return 0, ErrInvalidLengthVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVesting = fmt.Errorf("proto: unexpected end of group")
)
