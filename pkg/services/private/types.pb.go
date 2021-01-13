// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/services/private/types.proto

package private

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Health status of the storage node.
type HealthStatus int32

const (
	// Undefined status, default value.
	HealthStatus_STATUS_UNDEFINED HealthStatus = 0
	// Node is online.
	HealthStatus_ONLINE HealthStatus = 1
	// Node is offline.
	HealthStatus_OFFLINE HealthStatus = 2
)

var HealthStatus_name = map[int32]string{
	0: "STATUS_UNDEFINED",
	1: "ONLINE",
	2: "OFFLINE",
}

var HealthStatus_value = map[string]int32{
	"STATUS_UNDEFINED": 0,
	"ONLINE":           1,
	"OFFLINE":          2,
}

func (x HealthStatus) String() string {
	return proto.EnumName(HealthStatus_name, int32(x))
}

func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d50f4a50d05f699c, []int{0}
}

// Signature of some message.
type Signature struct {
	// Public key used for signing.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Binary signature.
	Sign                 []byte   `protobuf:"bytes,2,opt,name=sign,json=signature,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_d50f4a50d05f699c, []int{0}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
}
func (m *Signature) XXX_Size() int {
	return m.Size()
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Signature) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func init() {
	proto.RegisterEnum("private.HealthStatus", HealthStatus_name, HealthStatus_value)
	proto.RegisterType((*Signature)(nil), "private.Signature")
}

func init() { proto.RegisterFile("pkg/services/private/types.proto", fileDescriptor_d50f4a50d05f699c) }

var fileDescriptor_d50f4a50d05f699c = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0xc8, 0x4e, 0xd7,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2f, 0x28, 0xca, 0x2c, 0x4b, 0x2c, 0x49,
	0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x0a,
	0x2a, 0x99, 0x71, 0x71, 0x06, 0x67, 0xa6, 0xe7, 0x25, 0x96, 0x94, 0x16, 0xa5, 0x0a, 0x09, 0x70,
	0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0x98, 0x42, 0xe2, 0x5c,
	0x2c, 0xc5, 0x99, 0xe9, 0x79, 0x12, 0x4c, 0x60, 0x21, 0xce, 0x62, 0x98, 0x52, 0x2d, 0x5b, 0x2e,
	0x1e, 0x8f, 0xd4, 0xc4, 0x9c, 0x92, 0x8c, 0xe0, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x21, 0x11, 0x2e,
	0x81, 0xe0, 0x10, 0xc7, 0x90, 0xd0, 0xe0, 0xf8, 0x50, 0x3f, 0x17, 0x57, 0x37, 0x4f, 0x3f, 0x57,
	0x17, 0x01, 0x06, 0x21, 0x2e, 0x2e, 0x36, 0x7f, 0x3f, 0x1f, 0x4f, 0x3f, 0x57, 0x01, 0x46, 0x21,
	0x6e, 0x2e, 0x76, 0x7f, 0x37, 0x37, 0x30, 0x87, 0xc9, 0xc9, 0xed, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x24, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0xaf, 0xb8, 0x20, 0x39, 0x59, 0x37, 0x25,
	0xb5, 0x4c, 0x3f, 0x2f, 0x35, 0x3f, 0xad, 0x58, 0x37, 0x2f, 0x3f, 0x25, 0x55, 0x1f, 0x9b, 0x9f,
	0x92, 0xd8, 0xc0, 0xde, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x57, 0x15, 0x0c, 0xc5, 0xf2,
	0x00, 0x00, 0x00,
}

func (m *Signature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Signature) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Sign) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Sign)))
		i += copy(dAtA[i:], m.Sign)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Signature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Sign)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Signature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Signature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Signature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sign = append(m.Sign[:0], dAtA[iNdEx:postIndex]...)
			if m.Sign == nil {
				m.Sign = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTypes(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTypes
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)