// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import/message.proto

package _import

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//  import request
type ImportRequest struct {
	Quantity             int64    `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportRequest) Reset()         { *m = ImportRequest{} }
func (m *ImportRequest) String() string { return proto.CompactTextString(m) }
func (*ImportRequest) ProtoMessage()    {}
func (*ImportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_841f828cfc241ae5, []int{0}
}
func (m *ImportRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ImportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ImportRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ImportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportRequest.Merge(m, src)
}
func (m *ImportRequest) XXX_Size() int {
	return m.Size()
}
func (m *ImportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImportRequest proto.InternalMessageInfo

func (m *ImportRequest) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

//  import response{
type ImportResponse struct {
	NumberOfCreated      int64    `protobuf:"varint,1,opt,name=number_of_created,json=numberOfCreated,proto3" json:"number_of_created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportResponse) Reset()         { *m = ImportResponse{} }
func (m *ImportResponse) String() string { return proto.CompactTextString(m) }
func (*ImportResponse) ProtoMessage()    {}
func (*ImportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_841f828cfc241ae5, []int{1}
}
func (m *ImportResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ImportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ImportResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ImportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportResponse.Merge(m, src)
}
func (m *ImportResponse) XXX_Size() int {
	return m.Size()
}
func (m *ImportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImportResponse proto.InternalMessageInfo

func (m *ImportResponse) GetNumberOfCreated() int64 {
	if m != nil {
		return m.NumberOfCreated
	}
	return 0
}

func init() {
	proto.RegisterType((*ImportRequest)(nil), "import.ImportRequest")
	proto.RegisterType((*ImportResponse)(nil), "import.ImportResponse")
}

func init() { proto.RegisterFile("import/message.proto", fileDescriptor_841f828cfc241ae5) }

var fileDescriptor_841f828cfc241ae5 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0xd1, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x83, 0x88, 0x2a, 0x69, 0x73, 0xf1, 0x7a, 0x82, 0x59, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x52, 0x5c, 0x1c, 0x85, 0xa5, 0x89, 0x79, 0x25, 0x99, 0x25, 0x95, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x70, 0xbe, 0x92, 0x0d, 0x17, 0x1f, 0x4c, 0x71, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x16, 0x97, 0x60, 0x5e, 0x69, 0x6e, 0x52, 0x6a, 0x51, 0x7c, 0x7e, 0x5a,
	0x7c, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x6a, 0x0a, 0x54, 0x1b, 0x3f, 0x44, 0xc2, 0x3f, 0xcd, 0x19,
	0x22, 0xec, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31,
	0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0x76, 0x8b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x82,
	0x1c, 0x86, 0x46, 0xa3, 0x00, 0x00, 0x00,
}

func (m *ImportRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImportRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ImportRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Quantity != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Quantity))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ImportResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImportResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ImportResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.NumberOfCreated != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.NumberOfCreated))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ImportRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Quantity != 0 {
		n += 1 + sovMessage(uint64(m.Quantity))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ImportResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumberOfCreated != 0 {
		n += 1 + sovMessage(uint64(m.NumberOfCreated))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ImportRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: ImportRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImportRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			m.Quantity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quantity |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *ImportResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: ImportResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImportResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberOfCreated", wireType)
			}
			m.NumberOfCreated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumberOfCreated |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
