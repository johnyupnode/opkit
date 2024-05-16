// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opkit/domain/domain.proto

package types

import (
	fmt "fmt"
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

type Domain struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Domain    string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Owner     string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Txhash    string `protobuf:"bytes,5,opt,name=txhash,proto3" json:"txhash,omitempty"`
	Creator   string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Domain) Reset()         { *m = Domain{} }
func (m *Domain) String() string { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()    {}
func (*Domain) Descriptor() ([]byte, []int) {
	return fileDescriptor_855f959cfe6a0c3b, []int{0}
}
func (m *Domain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Domain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Domain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Domain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Domain.Merge(m, src)
}
func (m *Domain) XXX_Size() int {
	return m.Size()
}
func (m *Domain) XXX_DiscardUnknown() {
	xxx_messageInfo_Domain.DiscardUnknown(m)
}

var xxx_messageInfo_Domain proto.InternalMessageInfo

func (m *Domain) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Domain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Domain) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Domain) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Domain) GetTxhash() string {
	if m != nil {
		return m.Txhash
	}
	return ""
}

func (m *Domain) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Domain)(nil), "opkit.domain.Domain")
}

func init() { proto.RegisterFile("opkit/domain/domain.proto", fileDescriptor_855f959cfe6a0c3b) }

var fileDescriptor_855f959cfe6a0c3b = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x2f, 0xc8, 0xce,
	0x2c, 0xd1, 0x4f, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x83, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x3c, 0x60, 0x29, 0x3d, 0x88, 0x98, 0xd2, 0x34, 0x46, 0x2e, 0x36, 0x17, 0x30, 0x53, 0x88,
	0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x25, 0x88, 0x29, 0x33, 0x45, 0x48,
	0x8c, 0x8b, 0x0d, 0xa2, 0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xca, 0x13, 0x12, 0xe1,
	0x62, 0xcd, 0x2f, 0xcf, 0x4b, 0x2d, 0x92, 0x60, 0x06, 0x0b, 0x43, 0x38, 0x42, 0x32, 0x5c, 0x9c,
	0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x12, 0x2c, 0x60, 0x19, 0x84, 0x00, 0xc8,
	0xac, 0x92, 0x8a, 0x8c, 0xc4, 0xe2, 0x0c, 0x09, 0x56, 0x88, 0x59, 0x10, 0x9e, 0x90, 0x04, 0x17,
	0x7b, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x7e, 0x91, 0x04, 0x1b, 0x58, 0x02, 0xc6, 0x75, 0xd2, 0x3b,
	0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63,
	0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x11, 0x88, 0xdf, 0x2a, 0x60, 0xbe,
	0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xfb, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x4d, 0x79, 0x25, 0x6c, 0xfa, 0x00, 0x00, 0x00,
}

func (m *Domain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Domain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Domain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDomain(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Txhash) > 0 {
		i -= len(m.Txhash)
		copy(dAtA[i:], m.Txhash)
		i = encodeVarintDomain(dAtA, i, uint64(len(m.Txhash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Timestamp) > 0 {
		i -= len(m.Timestamp)
		copy(dAtA[i:], m.Timestamp)
		i = encodeVarintDomain(dAtA, i, uint64(len(m.Timestamp)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintDomain(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintDomain(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintDomain(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDomain(dAtA []byte, offset int, v uint64) int {
	offset -= sovDomain(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Domain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovDomain(uint64(m.Id))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovDomain(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDomain(uint64(l))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovDomain(uint64(l))
	}
	l = len(m.Txhash)
	if l > 0 {
		n += 1 + l + sovDomain(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDomain(uint64(l))
	}
	return n
}

func sovDomain(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDomain(x uint64) (n int) {
	return sovDomain(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Domain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDomain
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
			return fmt.Errorf("proto: Domain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Domain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomain
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
				return ErrInvalidLengthDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomain
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
				return ErrInvalidLengthDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomain
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
				return ErrInvalidLengthDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txhash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomain
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
				return ErrInvalidLengthDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txhash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomain
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
				return ErrInvalidLengthDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDomain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDomain
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
func skipDomain(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDomain
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
					return 0, ErrIntOverflowDomain
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
					return 0, ErrIntOverflowDomain
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
				return 0, ErrInvalidLengthDomain
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDomain
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDomain
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDomain        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDomain          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDomain = fmt.Errorf("proto: unexpected end of group")
)