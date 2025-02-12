// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/transfer/v1/authz.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Allocation defines the spend limit for a particular port and channel
type Allocation struct {
	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
	// spend limitation on the channel
	SpendLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=spend_limit,json=spendLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"spend_limit"`
	// allow list of receivers, an empty allow list permits any receiver address
	AllowList []string `protobuf:"bytes,4,rep,name=allow_list,json=allowList,proto3" json:"allow_list,omitempty"`
	// allow list of memo strings, an empty list prohibits all memo strings;
	// a list only with "*" permits any memo string
	AllowedPacketData []string `protobuf:"bytes,5,rep,name=allowed_packet_data,json=allowedPacketData,proto3" json:"allowed_packet_data,omitempty"`
}

func (m *Allocation) Reset()         { *m = Allocation{} }
func (m *Allocation) String() string { return proto.CompactTextString(m) }
func (*Allocation) ProtoMessage()    {}
func (*Allocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1a28b55d17325aa, []int{0}
}
func (m *Allocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Allocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Allocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Allocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Allocation.Merge(m, src)
}
func (m *Allocation) XXX_Size() int {
	return m.Size()
}
func (m *Allocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Allocation.DiscardUnknown(m)
}

var xxx_messageInfo_Allocation proto.InternalMessageInfo

func (m *Allocation) GetSourcePort() string {
	if m != nil {
		return m.SourcePort
	}
	return ""
}

func (m *Allocation) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *Allocation) GetSpendLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.SpendLimit
	}
	return nil
}

func (m *Allocation) GetAllowList() []string {
	if m != nil {
		return m.AllowList
	}
	return nil
}

func (m *Allocation) GetAllowedPacketData() []string {
	if m != nil {
		return m.AllowedPacketData
	}
	return nil
}

// TransferAuthorization allows the grantee to spend up to spend_limit coins from
// the granter's account for ibc transfer on a specific channel
type TransferAuthorization struct {
	// port and channel amounts
	Allocations []Allocation `protobuf:"bytes,1,rep,name=allocations,proto3" json:"allocations"`
}

func (m *TransferAuthorization) Reset()         { *m = TransferAuthorization{} }
func (m *TransferAuthorization) String() string { return proto.CompactTextString(m) }
func (*TransferAuthorization) ProtoMessage()    {}
func (*TransferAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1a28b55d17325aa, []int{1}
}
func (m *TransferAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransferAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransferAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferAuthorization.Merge(m, src)
}
func (m *TransferAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *TransferAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_TransferAuthorization proto.InternalMessageInfo

func (m *TransferAuthorization) GetAllocations() []Allocation {
	if m != nil {
		return m.Allocations
	}
	return nil
}

func init() {
	proto.RegisterType((*Allocation)(nil), "ibc.applications.transfer.v1.Allocation")
	proto.RegisterType((*TransferAuthorization)(nil), "ibc.applications.transfer.v1.TransferAuthorization")
}

func init() {
	proto.RegisterFile("ibc/applications/transfer/v1/authz.proto", fileDescriptor_b1a28b55d17325aa)
}

var fileDescriptor_b1a28b55d17325aa = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x9b, 0x75, 0x20, 0xd5, 0x15, 0x48, 0x04, 0x90, 0xb2, 0x09, 0xd2, 0xaa, 0x12, 0x28,
	0x12, 0xaa, 0x4d, 0xe1, 0x80, 0xe0, 0xb6, 0x8e, 0xe3, 0x0e, 0xa5, 0xe2, 0xc4, 0x25, 0x72, 0x1c,
	0xd3, 0x5a, 0x73, 0xf2, 0x46, 0xf1, 0x9b, 0x20, 0xf6, 0x29, 0xd8, 0xd7, 0xe0, 0xcc, 0x87, 0x98,
	0x38, 0xed, 0xc8, 0x09, 0x50, 0xfb, 0x45, 0x50, 0x6c, 0x6f, 0x14, 0x21, 0xed, 0x94, 0xf8, 0x79,
	0x1f, 0xbf, 0x7f, 0x7e, 0x7e, 0x49, 0xa2, 0x32, 0xc1, 0x78, 0x55, 0x69, 0x25, 0x38, 0x2a, 0x28,
	0x0d, 0xc3, 0x9a, 0x97, 0xe6, 0xa3, 0xac, 0x59, 0x3b, 0x63, 0xbc, 0xc1, 0xf5, 0x19, 0xad, 0x6a,
	0x40, 0x08, 0x1f, 0xa9, 0x4c, 0xd0, 0x5d, 0x27, 0xbd, 0x72, 0xd2, 0x76, 0x76, 0x78, 0x20, 0xc0,
	0x14, 0x60, 0x52, 0xeb, 0x65, 0xee, 0xe0, 0x2e, 0x1e, 0x3e, 0x58, 0xc1, 0x0a, 0x9c, 0xde, 0xfd,
	0x79, 0x35, 0x76, 0x1e, 0x96, 0x71, 0x23, 0x59, 0x3b, 0xcb, 0x24, 0xf2, 0x19, 0x13, 0xa0, 0x4a,
	0x1f, 0x7f, 0x76, 0x63, 0x63, 0xd7, 0xa5, 0xad, 0x79, 0x72, 0xbe, 0x47, 0xc8, 0x91, 0xd6, 0xe0,
	0xac, 0xe1, 0x88, 0x0c, 0x0d, 0x34, 0xb5, 0x90, 0x69, 0x05, 0x35, 0x46, 0xc1, 0x38, 0x48, 0x06,
	0x4b, 0xe2, 0xa4, 0x05, 0xd4, 0x18, 0x3e, 0x21, 0x77, 0xbd, 0x41, 0xac, 0x79, 0x59, 0x4a, 0x1d,
	0xed, 0x59, 0xcf, 0x1d, 0xa7, 0x1e, 0x3b, 0x31, 0xd4, 0x64, 0x68, 0x2a, 0x59, 0xe6, 0xa9, 0x56,
	0x85, 0xc2, 0xa8, 0x3f, 0xee, 0x27, 0xc3, 0x17, 0x07, 0xd4, 0x4f, 0xd7, 0x75, 0x4e, 0x7d, 0xe7,
	0xf4, 0x18, 0x54, 0x39, 0x7f, 0x7e, 0xf1, 0x73, 0xd4, 0xfb, 0xfa, 0x6b, 0x94, 0xac, 0x14, 0xae,
	0x9b, 0x8c, 0x0a, 0x28, 0x3c, 0x0a, 0xff, 0x99, 0x9a, 0xfc, 0x94, 0xe1, 0xe7, 0x4a, 0x1a, 0x7b,
	0xc1, 0x2c, 0x89, 0xcd, 0x7f, 0xd2, 0xa5, 0x0f, 0x1f, 0x13, 0xc2, 0xb5, 0x86, 0x4f, 0xa9, 0x56,
	0x06, 0xa3, 0xfd, 0x71, 0x3f, 0x19, 0x2c, 0x07, 0x56, 0x39, 0x51, 0x06, 0x43, 0x4a, 0xee, 0xdb,
	0x83, 0xcc, 0xd3, 0x8a, 0x8b, 0x53, 0x89, 0x69, 0xce, 0x91, 0x47, 0xb7, 0xac, 0xef, 0x9e, 0x0f,
	0x2d, 0x6c, 0xe4, 0x2d, 0x47, 0x3e, 0x39, 0x0f, 0xc8, 0xc3, 0xf7, 0x1e, 0xd3, 0x51, 0x83, 0x6b,
	0xa8, 0xd5, 0x99, 0xc3, 0xb3, 0x20, 0x43, 0x7e, 0x0d, 0xcb, 0x44, 0x81, 0x1d, 0x2b, 0xa1, 0x37,
	0xbd, 0x2f, 0xfd, 0x4b, 0x77, 0xbe, 0xdf, 0x4d, 0xb9, 0xdc, 0x4d, 0xf1, 0xe6, 0xe9, 0xf7, 0x6f,
	0xd3, 0x89, 0xc7, 0xe2, 0x76, 0xe6, 0x8a, 0xcb, 0x3f, 0x95, 0xe7, 0xef, 0x2e, 0x36, 0x71, 0x70,
	0xb9, 0x89, 0x83, 0xdf, 0x9b, 0x38, 0xf8, 0xb2, 0x8d, 0x7b, 0x97, 0xdb, 0xb8, 0xf7, 0x63, 0x1b,
	0xf7, 0x3e, 0xbc, 0xfa, 0x1f, 0x99, 0xca, 0xc4, 0x74, 0x05, 0xac, 0x7d, 0xcd, 0x0a, 0xc8, 0x1b,
	0x2d, 0x4d, 0xb7, 0x0e, 0x3b, 0x6b, 0x60, 0x39, 0x66, 0xb7, 0xed, 0x06, 0xbc, 0xfc, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x89, 0xe3, 0xe3, 0x15, 0xc9, 0x02, 0x00, 0x00,
}

func (m *Allocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Allocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Allocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowedPacketData) > 0 {
		for iNdEx := len(m.AllowedPacketData) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedPacketData[iNdEx])
			copy(dAtA[i:], m.AllowedPacketData[iNdEx])
			i = encodeVarintAuthz(dAtA, i, uint64(len(m.AllowedPacketData[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.AllowList) > 0 {
		for iNdEx := len(m.AllowList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowList[iNdEx])
			copy(dAtA[i:], m.AllowList[iNdEx])
			i = encodeVarintAuthz(dAtA, i, uint64(len(m.AllowList[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.SpendLimit) > 0 {
		for iNdEx := len(m.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthz(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TransferAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransferAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Allocations) > 0 {
		for iNdEx := len(m.Allocations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Allocations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthz(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Allocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	if len(m.SpendLimit) > 0 {
		for _, e := range m.SpendLimit {
			l = e.Size()
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	if len(m.AllowList) > 0 {
		for _, s := range m.AllowList {
			l = len(s)
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	if len(m.AllowedPacketData) > 0 {
		for _, s := range m.AllowedPacketData {
			l = len(s)
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func (m *TransferAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Allocations) > 0 {
		for _, e := range m.Allocations {
			l = e.Size()
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Allocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: Allocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Allocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendLimit = append(m.SpendLimit, types.Coin{})
			if err := m.SpendLimit[len(m.SpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowList = append(m.AllowList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedPacketData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedPacketData = append(m.AllowedPacketData, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *TransferAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: TransferAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Allocations = append(m.Allocations, Allocation{})
			if err := m.Allocations[len(m.Allocations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
