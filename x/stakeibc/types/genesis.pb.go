// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dredger/stakeibc/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the stakeibc module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	// list of zones that are registered by the protocol
	IcaAccount    *ICAAccount `protobuf:"bytes,4,opt,name=ica_account,json=icaAccount,proto3" json:"ica_account,omitempty"`
	HostZoneList  []HostZone  `protobuf:"bytes,5,rep,name=host_zone_list,json=hostZoneList,proto3" json:"host_zone_list"`
	HostZoneCount uint64      `protobuf:"varint,6,opt,name=host_zone_count,json=hostZoneCount,proto3" json:"host_zone_count,omitempty"`
	// stores a map from hostZone base denom to hostZone
	DenomToHostZone  map[string]string `protobuf:"bytes,9,rep,name=denom_to_host_zone,json=denomToHostZone,proto3" json:"denom_to_host_zone,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EpochTrackerList []EpochTracker    `protobuf:"bytes,10,rep,name=epoch_tracker_list,json=epochTrackerList,proto3" json:"epoch_tracker_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea81129ed6fb77a, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetIcaAccount() *ICAAccount {
	if m != nil {
		return m.IcaAccount
	}
	return nil
}

func (m *GenesisState) GetHostZoneList() []HostZone {
	if m != nil {
		return m.HostZoneList
	}
	return nil
}

func (m *GenesisState) GetHostZoneCount() uint64 {
	if m != nil {
		return m.HostZoneCount
	}
	return 0
}

func (m *GenesisState) GetDenomToHostZone() map[string]string {
	if m != nil {
		return m.DenomToHostZone
	}
	return nil
}

func (m *GenesisState) GetEpochTrackerList() []EpochTracker {
	if m != nil {
		return m.EpochTrackerList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "dredger.stakeibc.GenesisState")
	proto.RegisterMapType((map[string]string)(nil), "dredger.stakeibc.GenesisState.DenomToHostZoneEntry")
}

func init() { proto.RegisterFile("dredger/stakeibc/genesis.proto", fileDescriptor_dea81129ed6fb77a) }

var fileDescriptor_dea81129ed6fb77a = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0x35, 0x2b, 0xab, 0x5b, 0x68, 0x65, 0x55, 0x5a, 0x28, 0x2c, 0x2b, 0x20, 0xa1,
	0x5e, 0x48, 0x44, 0x01, 0x09, 0x21, 0x2e, 0xeb, 0xa8, 0x60, 0x63, 0x07, 0xc8, 0x76, 0xda, 0xc5,
	0x72, 0x1d, 0xab, 0xb5, 0xba, 0xc6, 0x51, 0xfc, 0x3a, 0x51, 0xbe, 0x04, 0x7c, 0xac, 0x1d, 0x77,
	0xe4, 0x84, 0x50, 0xfb, 0x45, 0x50, 0x6c, 0x77, 0xeb, 0xd2, 0xdd, 0xec, 0xbc, 0x9f, 0xff, 0xff,
	0xff, 0xcb, 0x7b, 0x68, 0x4f, 0x41, 0x26, 0x62, 0x1e, 0x2a, 0xa0, 0x13, 0x2e, 0x86, 0x2c, 0x1c,
	0xf1, 0x84, 0x2b, 0xa1, 0x82, 0x34, 0x93, 0x20, 0x71, 0xc3, 0x94, 0x83, 0x55, 0xb9, 0xdd, 0x1a,
	0xc9, 0x91, 0xd4, 0xb5, 0x30, 0x3f, 0x19, 0xac, 0xfd, 0xb4, 0xa8, 0x92, 0xd2, 0x8c, 0x4e, 0xad,
	0x48, 0xfb, 0x59, 0xb1, 0x2a, 0x18, 0x25, 0x94, 0x31, 0x39, 0x4b, 0xc0, 0x22, 0xfb, 0x45, 0x64,
	0x2c, 0x15, 0x90, 0x9f, 0x32, 0xe1, 0x16, 0x78, 0x51, 0x04, 0x78, 0x2a, 0xd9, 0x98, 0x40, 0x46,
	0xd9, 0x84, 0x67, 0x06, 0x7a, 0xfe, 0xcb, 0x45, 0xf5, 0xcf, 0x26, 0xff, 0x29, 0x50, 0xe0, 0xf8,
	0x1d, 0xaa, 0x98, 0x24, 0x9e, 0xd3, 0x71, 0xba, 0xb5, 0xde, 0x6e, 0x50, 0xe8, 0x27, 0xf8, 0xa6,
	0xcb, 0x7d, 0xf7, 0xea, 0xef, 0x7e, 0x29, 0xb2, 0x30, 0xde, 0x45, 0x0f, 0x52, 0x99, 0x01, 0x11,
	0xb1, 0xb7, 0xd5, 0x71, 0xba, 0xd5, 0xa8, 0x92, 0x5f, 0x8f, 0x62, 0xfc, 0x11, 0xd5, 0xd6, 0xb2,
	0x7b, 0xae, 0x16, 0x7d, 0xb2, 0x21, 0x7a, 0x74, 0x78, 0x70, 0x60, 0x90, 0x08, 0x09, 0x46, 0xed,
	0x19, 0x0f, 0xd0, 0xa3, 0x9b, 0xb6, 0xc8, 0x85, 0x50, 0xe0, 0x6d, 0x77, 0xca, 0xdd, 0x5a, 0xef,
	0xf1, 0x86, 0xc0, 0x17, 0xa9, 0xe0, 0x5c, 0x26, 0xdc, 0xe6, 0xaa, 0x8f, 0xed, 0xfd, 0x44, 0x28,
	0xc0, 0x2f, 0x51, 0xe3, 0x56, 0xc6, 0x04, 0xa9, 0x74, 0x9c, 0xae, 0x1b, 0x3d, 0x5c, 0x61, 0x87,
	0xda, 0x8e, 0x20, 0x1c, 0xf3, 0x44, 0x4e, 0x09, 0x48, 0x72, 0xf3, 0xc0, 0xab, 0x6a, 0xcb, 0xde,
	0x86, 0xe5, 0xfa, 0x7f, 0x0b, 0x3e, 0xe5, 0xef, 0xce, 0xe4, 0x2a, 0xc6, 0x20, 0x81, 0x6c, 0x1e,
	0x35, 0xe2, 0xbb, 0x5f, 0xf1, 0x77, 0x84, 0xef, 0x4c, 0xc1, 0xf4, 0x84, 0xb4, 0xc1, 0xde, 0x86,
	0xc1, 0x20, 0x47, 0xcf, 0x0c, 0x69, 0xfb, 0x6a, 0xf2, 0xb5, 0x6f, 0x79, 0x6f, 0xed, 0x3e, 0x6a,
	0xdd, 0xe7, 0x8d, 0x9b, 0xa8, 0x3c, 0xe1, 0x73, 0x3d, 0xc5, 0x6a, 0x94, 0x1f, 0x71, 0x0b, 0x6d,
	0x5f, 0xd2, 0x8b, 0x19, 0xb7, 0x13, 0x32, 0x97, 0x0f, 0x5b, 0xef, 0x9d, 0x63, 0x77, 0xa7, 0xdc,
	0x74, 0x8f, 0xdd, 0x9d, 0x5a, 0xb3, 0xde, 0xff, 0x7a, 0xb5, 0xf0, 0x9d, 0xeb, 0x85, 0xef, 0xfc,
	0x5b, 0xf8, 0xce, 0xef, 0xa5, 0x5f, 0xba, 0x5e, 0xfa, 0xa5, 0x3f, 0x4b, 0xbf, 0x74, 0xfe, 0x7a,
	0x24, 0x60, 0x3c, 0x1b, 0x06, 0x4c, 0x4e, 0xc3, 0x53, 0x1d, 0xf5, 0xd5, 0x09, 0x1d, 0xaa, 0xd0,
	0xee, 0xd9, 0xe5, 0xdb, 0xf0, 0xc7, 0xed, 0xb2, 0xc1, 0x3c, 0xe5, 0x6a, 0x58, 0xd1, 0x5b, 0xf6,
	0xe6, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x35, 0xe5, 0x2f, 0x34, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EpochTrackerList) > 0 {
		for iNdEx := len(m.EpochTrackerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EpochTrackerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.DenomToHostZone) > 0 {
		for k := range m.DenomToHostZone {
			v := m.DenomToHostZone[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGenesis(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.HostZoneCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.HostZoneCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.HostZoneList) > 0 {
		for iNdEx := len(m.HostZoneList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HostZoneList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.IcaAccount != nil {
		{
			size, err := m.IcaAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.IcaAccount != nil {
		l = m.IcaAccount.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.HostZoneList) > 0 {
		for _, e := range m.HostZoneList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.HostZoneCount != 0 {
		n += 1 + sovGenesis(uint64(m.HostZoneCount))
	}
	if len(m.DenomToHostZone) > 0 {
		for k, v := range m.DenomToHostZone {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + 1 + len(v) + sovGenesis(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	if len(m.EpochTrackerList) > 0 {
		for _, e := range m.EpochTrackerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcaAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IcaAccount == nil {
				m.IcaAccount = &ICAAccount{}
			}
			if err := m.IcaAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZoneList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostZoneList = append(m.HostZoneList, HostZone{})
			if err := m.HostZoneList[len(m.HostZoneList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZoneCount", wireType)
			}
			m.HostZoneCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HostZoneCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomToHostZone", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DenomToHostZone == nil {
				m.DenomToHostZone = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.DenomToHostZone[mapkey] = mapvalue
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochTrackerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochTrackerList = append(m.EpochTrackerList, EpochTracker{})
			if err := m.EpochTrackerList[len(m.EpochTrackerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
