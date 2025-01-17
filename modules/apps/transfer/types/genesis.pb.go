// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/transfer/v1/genesis.proto

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

// GenesisState defines the ibc-transfer genesis state
type GenesisState struct {
	PortId      string `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty" yaml:"port_id"`
	DenomTraces Traces `protobuf:"bytes,2,rep,name=denom_traces,json=denomTraces,proto3,castrepeated=Traces" json:"denom_traces" yaml:"denom_traces"`
	Params      Params `protobuf:"bytes,3,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4f788affd5bea89, []int{0}
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

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetDenomTraces() Traces {
	if m != nil {
		return m.DenomTraces
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.applications.transfer.v1.GenesisState")
}

func init() {
	proto.RegisterFile("ibc/applications/transfer/v1/genesis.proto", fileDescriptor_a4f788affd5bea89)
}

var fileDescriptor_a4f788affd5bea89 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x1b, 0x27, 0x15, 0xbb, 0xe1, 0xa1, 0x7a, 0x18, 0x43, 0xd2, 0x51, 0x14, 0x8a, 0xc3,
	0x84, 0x6d, 0x78, 0xf1, 0x58, 0x04, 0xf1, 0xa4, 0x4c, 0x4f, 0x5e, 0x46, 0x9a, 0xc6, 0x1a, 0x68,
	0x9b, 0x90, 0x64, 0x83, 0xbd, 0x85, 0xcf, 0xe1, 0x93, 0xec, 0xb8, 0xa3, 0xa7, 0x29, 0xdb, 0x1b,
	0xcc, 0x17, 0x90, 0x74, 0x73, 0xec, 0xb4, 0xdb, 0x47, 0xf2, 0xfb, 0xbe, 0xef, 0xcf, 0xe7, 0x5d,
	0xf1, 0x84, 0x62, 0x22, 0x65, 0xce, 0x29, 0x31, 0x5c, 0x94, 0x1a, 0x1b, 0x45, 0x4a, 0xfd, 0xc6,
	0x14, 0x1e, 0x77, 0x71, 0xc6, 0x4a, 0xa6, 0xb9, 0x46, 0x52, 0x09, 0x23, 0xfc, 0x73, 0x9e, 0x50,
	0xb4, 0xcb, 0xa2, 0x7f, 0x16, 0x8d, 0xbb, 0xad, 0xce, 0xde, 0xa4, 0x2d, 0x59, 0x45, 0xb5, 0xce,
	0x32, 0x91, 0x89, 0x4a, 0x62, 0xab, 0xd6, 0xaf, 0xe1, 0x2f, 0xf0, 0x1a, 0xf7, 0xeb, 0xca, 0x67,
	0x43, 0x0c, 0xf3, 0x3b, 0xde, 0x91, 0x14, 0xca, 0x0c, 0x79, 0xda, 0x04, 0x6d, 0x10, 0x1d, 0xc7,
	0xfe, 0x6a, 0x1e, 0x9c, 0x4c, 0x48, 0x91, 0xdf, 0x86, 0x9b, 0x8f, 0x70, 0xe0, 0x5a, 0xf5, 0x90,
	0xfa, 0xca, 0x6b, 0xa4, 0xac, 0x14, 0xc5, 0xd0, 0x28, 0x42, 0x99, 0x6e, 0x1e, 0xb4, 0x6b, 0x51,
	0xbd, 0x17, 0xa1, 0x7d, 0x57, 0xa3, 0x3b, 0xeb, 0x78, 0xb1, 0x86, 0xf8, 0x72, 0x3a, 0x0f, 0x9c,
	0xd5, 0x3c, 0x38, 0x5d, 0xe7, 0xef, 0x66, 0x85, 0x9f, 0xdf, 0x81, 0x5b, 0x51, 0x7a, 0x50, 0x4f,
	0xb7, 0x16, 0xed, 0xc7, 0x9e, 0x2b, 0x89, 0x22, 0x85, 0x6e, 0xd6, 0xda, 0x20, 0xaa, 0xf7, 0x2e,
	0xf6, 0xb7, 0x3d, 0x55, 0x6c, 0x7c, 0x68, 0x9b, 0x06, 0x1b, 0x67, 0xfc, 0x38, 0x5d, 0x40, 0x30,
	0x5b, 0x40, 0xf0, 0xb3, 0x80, 0xe0, 0x63, 0x09, 0x9d, 0xd9, 0x12, 0x3a, 0x5f, 0x4b, 0xe8, 0xbc,
	0xde, 0x64, 0xdc, 0xbc, 0x8f, 0x12, 0x44, 0x45, 0x81, 0x73, 0x5e, 0x32, 0xcc, 0x13, 0x7a, 0x9d,
	0x09, 0x3c, 0xee, 0xe3, 0x42, 0xa4, 0xa3, 0x9c, 0x69, 0x3b, 0xf8, 0xce, 0xd0, 0x66, 0x22, 0x99,
	0x4e, 0xdc, 0x6a, 0xcd, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x51, 0x32, 0x2e, 0xdc,
	0x01, 0x00, 0x00,
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.DenomTraces) > 0 {
		for iNdEx := len(m.DenomTraces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomTraces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0xa
	}
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
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.DenomTraces) > 0 {
		for _, e := range m.DenomTraces {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomTraces", wireType)
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
			m.DenomTraces = append(m.DenomTraces, DenomTrace{})
			if err := m.DenomTraces[len(m.DenomTraces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
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
