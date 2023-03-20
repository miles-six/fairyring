// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/fairblock/encrypted_tx.proto

package types

import (
	fmt "fmt"
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

type EncryptedTx struct {
	TargetHeight uint64 `protobuf:"varint,1,opt,name=targetHeight,proto3" json:"targetHeight,omitempty"`
	Index        uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Data         string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Creator      string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *EncryptedTx) Reset()         { *m = EncryptedTx{} }
func (m *EncryptedTx) String() string { return proto.CompactTextString(m) }
func (*EncryptedTx) ProtoMessage()    {}
func (*EncryptedTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1183a5fa212eacc, []int{0}
}
func (m *EncryptedTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncryptedTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EncryptedTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EncryptedTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedTx.Merge(m, src)
}
func (m *EncryptedTx) XXX_Size() int {
	return m.Size()
}
func (m *EncryptedTx) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedTx.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedTx proto.InternalMessageInfo

func (m *EncryptedTx) GetTargetHeight() uint64 {
	if m != nil {
		return m.TargetHeight
	}
	return 0
}

func (m *EncryptedTx) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *EncryptedTx) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *EncryptedTx) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type EncryptedTxArray struct {
	EncryptedTx []EncryptedTx `protobuf:"bytes,1,rep,name=encryptedTx,proto3" json:"encryptedTx"`
}

func (m *EncryptedTxArray) Reset()         { *m = EncryptedTxArray{} }
func (m *EncryptedTxArray) String() string { return proto.CompactTextString(m) }
func (*EncryptedTxArray) ProtoMessage()    {}
func (*EncryptedTxArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1183a5fa212eacc, []int{1}
}
func (m *EncryptedTxArray) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncryptedTxArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EncryptedTxArray.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EncryptedTxArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedTxArray.Merge(m, src)
}
func (m *EncryptedTxArray) XXX_Size() int {
	return m.Size()
}
func (m *EncryptedTxArray) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedTxArray.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedTxArray proto.InternalMessageInfo

func (m *EncryptedTxArray) GetEncryptedTx() []EncryptedTx {
	if m != nil {
		return m.EncryptedTx
	}
	return nil
}

func init() {
	proto.RegisterType((*EncryptedTx)(nil), "fairyring.fairblock.EncryptedTx")
	proto.RegisterType((*EncryptedTxArray)(nil), "fairyring.fairblock.EncryptedTxArray")
}

func init() {
	proto.RegisterFile("fairyring/fairblock/encrypted_tx.proto", fileDescriptor_e1183a5fa212eacc)
}

var fileDescriptor_e1183a5fa212eacc = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0x4b, 0xcc, 0x2c,
	0xaa, 0x2c, 0xca, 0xcc, 0x4b, 0xd7, 0x07, 0xb1, 0x92, 0x72, 0xf2, 0x93, 0xb3, 0xf5, 0x53, 0xf3,
	0x92, 0x8b, 0x2a, 0x0b, 0x4a, 0x52, 0x53, 0xe2, 0x4b, 0x2a, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x84, 0xe1, 0xea, 0xf4, 0xe0, 0xea, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xf2, 0xfa,
	0x20, 0x16, 0x44, 0xa9, 0x52, 0x29, 0x17, 0xb7, 0x2b, 0xcc, 0x80, 0x90, 0x0a, 0x21, 0x25, 0x2e,
	0x9e, 0x92, 0xc4, 0xa2, 0xf4, 0xd4, 0x12, 0x8f, 0xd4, 0xcc, 0xf4, 0x8c, 0x12, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x96, 0x20, 0x14, 0x31, 0x21, 0x11, 0x2e, 0xd6, 0xcc, 0xbc, 0x94, 0xd4, 0x0a, 0x09,
	0x26, 0xb0, 0x24, 0x84, 0x23, 0x24, 0xc4, 0xc5, 0x92, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xac, 0xc0,
	0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x49, 0x70, 0xb1, 0x27, 0x17, 0xa5, 0x26, 0x96, 0xe4, 0x17,
	0x49, 0xb0, 0x80, 0x85, 0x61, 0x5c, 0xa5, 0x18, 0x2e, 0x01, 0x24, 0x6b, 0x1d, 0x8b, 0x8a, 0x12,
	0x2b, 0x85, 0x3c, 0xb8, 0xb8, 0x53, 0x11, 0x62, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x0a,
	0x7a, 0x58, 0xfc, 0xa2, 0x87, 0xa4, 0xd7, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0x64, 0xad,
	0x4e, 0xa6, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x8d, 0x08, 0xc1,
	0x0a, 0xa4, 0x30, 0x2c, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x07, 0x89, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x11, 0x67, 0xf1, 0x9d, 0x67, 0x01, 0x00, 0x00,
}

func (m *EncryptedTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncryptedTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EncryptedTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEncryptedTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintEncryptedTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Index != 0 {
		i = encodeVarintEncryptedTx(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if m.TargetHeight != 0 {
		i = encodeVarintEncryptedTx(dAtA, i, uint64(m.TargetHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EncryptedTxArray) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncryptedTxArray) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EncryptedTxArray) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptedTx) > 0 {
		for iNdEx := len(m.EncryptedTx) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EncryptedTx[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEncryptedTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintEncryptedTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovEncryptedTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EncryptedTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TargetHeight != 0 {
		n += 1 + sovEncryptedTx(uint64(m.TargetHeight))
	}
	if m.Index != 0 {
		n += 1 + sovEncryptedTx(uint64(m.Index))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovEncryptedTx(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEncryptedTx(uint64(l))
	}
	return n
}

func (m *EncryptedTxArray) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.EncryptedTx) > 0 {
		for _, e := range m.EncryptedTx {
			l = e.Size()
			n += 1 + l + sovEncryptedTx(uint64(l))
		}
	}
	return n
}

func sovEncryptedTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEncryptedTx(x uint64) (n int) {
	return sovEncryptedTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EncryptedTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEncryptedTx
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
			return fmt.Errorf("proto: EncryptedTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncryptedTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetHeight", wireType)
			}
			m.TargetHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncryptedTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncryptedTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncryptedTx
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
				return ErrInvalidLengthEncryptedTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEncryptedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncryptedTx
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
				return ErrInvalidLengthEncryptedTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEncryptedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEncryptedTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEncryptedTx
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
func (m *EncryptedTxArray) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEncryptedTx
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
			return fmt.Errorf("proto: EncryptedTxArray: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncryptedTxArray: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncryptedTx
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
				return ErrInvalidLengthEncryptedTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEncryptedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedTx = append(m.EncryptedTx, EncryptedTx{})
			if err := m.EncryptedTx[len(m.EncryptedTx)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEncryptedTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEncryptedTx
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
func skipEncryptedTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEncryptedTx
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
					return 0, ErrIntOverflowEncryptedTx
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
					return 0, ErrIntOverflowEncryptedTx
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
				return 0, ErrInvalidLengthEncryptedTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEncryptedTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEncryptedTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEncryptedTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEncryptedTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEncryptedTx = fmt.Errorf("proto: unexpected end of group")
)