// Code generated by protoc-gen-gogo.
// source: kythe/proto/filecontext.proto
// DO NOT EDIT!

/*
	Package filecontext_proto is a generated protocol buffer package.

	It is generated from these files:
		kythe/proto/filecontext.proto

	It has these top-level messages:
		ContextDependentVersion
*/
package filecontext_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// ContextDependentVersion columns and rows define a table that relates input
// contexts (keyed by a single source context per row) to tuples of (byte
// offset * linked context).  When a file F being processed in context C refers
// to another file F' at offset O (perhaps because F has an #include directive
// at O) the context in which F' should be processed is the linked context
// derived from this table.
type ContextDependentVersion struct {
	Row []*ContextDependentVersion_Row `protobuf:"bytes,1,rep,name=row" json:"row,omitempty"`
}

func (m *ContextDependentVersion) Reset()         { *m = ContextDependentVersion{} }
func (m *ContextDependentVersion) String() string { return proto.CompactTextString(m) }
func (*ContextDependentVersion) ProtoMessage()    {}
func (*ContextDependentVersion) Descriptor() ([]byte, []int) {
	return fileDescriptorFilecontext, []int{0}
}

func (m *ContextDependentVersion) GetRow() []*ContextDependentVersion_Row {
	if m != nil {
		return m.Row
	}
	return nil
}

type ContextDependentVersion_Column struct {
	// The byte offset into the file resource.
	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// The signature for the resulting context.
	LinkedContext string `protobuf:"bytes,2,opt,name=linked_context,json=linkedContext,proto3" json:"linked_context,omitempty"`
}

func (m *ContextDependentVersion_Column) Reset()         { *m = ContextDependentVersion_Column{} }
func (m *ContextDependentVersion_Column) String() string { return proto.CompactTextString(m) }
func (*ContextDependentVersion_Column) ProtoMessage()    {}
func (*ContextDependentVersion_Column) Descriptor() ([]byte, []int) {
	return fileDescriptorFilecontext, []int{0, 0}
}

// See ContextDependentVersionColumn for details.  It is valid for a Row to
// have no columns. In this case, the associated file was seen to exist in
// some context C, but did not refer to any other files while in that
// context.
type ContextDependentVersion_Row struct {
	// The context to be applied to all columns.
	SourceContext string `protobuf:"bytes,1,opt,name=source_context,json=sourceContext,proto3" json:"source_context,omitempty"`
	// A map from byte offsets to linked contexts.
	Column []*ContextDependentVersion_Column `protobuf:"bytes,2,rep,name=column" json:"column,omitempty"`
	// If true, this version should always be processed regardless of any
	// claiming.
	AlwaysProcess bool `protobuf:"varint,3,opt,name=always_process,json=alwaysProcess,proto3" json:"always_process,omitempty"`
}

func (m *ContextDependentVersion_Row) Reset()         { *m = ContextDependentVersion_Row{} }
func (m *ContextDependentVersion_Row) String() string { return proto.CompactTextString(m) }
func (*ContextDependentVersion_Row) ProtoMessage()    {}
func (*ContextDependentVersion_Row) Descriptor() ([]byte, []int) {
	return fileDescriptorFilecontext, []int{0, 1}
}

func (m *ContextDependentVersion_Row) GetColumn() []*ContextDependentVersion_Column {
	if m != nil {
		return m.Column
	}
	return nil
}

func init() {
	proto.RegisterType((*ContextDependentVersion)(nil), "kythe.proto.ContextDependentVersion")
	proto.RegisterType((*ContextDependentVersion_Column)(nil), "kythe.proto.ContextDependentVersion.Column")
	proto.RegisterType((*ContextDependentVersion_Row)(nil), "kythe.proto.ContextDependentVersion.Row")
}
func (m *ContextDependentVersion) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ContextDependentVersion) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Row) > 0 {
		for _, msg := range m.Row {
			data[i] = 0xa
			i++
			i = encodeVarintFilecontext(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ContextDependentVersion_Column) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ContextDependentVersion_Column) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Offset != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintFilecontext(data, i, uint64(m.Offset))
	}
	if len(m.LinkedContext) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintFilecontext(data, i, uint64(len(m.LinkedContext)))
		i += copy(data[i:], m.LinkedContext)
	}
	return i, nil
}

func (m *ContextDependentVersion_Row) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ContextDependentVersion_Row) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SourceContext) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintFilecontext(data, i, uint64(len(m.SourceContext)))
		i += copy(data[i:], m.SourceContext)
	}
	if len(m.Column) > 0 {
		for _, msg := range m.Column {
			data[i] = 0x12
			i++
			i = encodeVarintFilecontext(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.AlwaysProcess {
		data[i] = 0x18
		i++
		if m.AlwaysProcess {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Filecontext(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Filecontext(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFilecontext(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ContextDependentVersion) Size() (n int) {
	var l int
	_ = l
	if len(m.Row) > 0 {
		for _, e := range m.Row {
			l = e.Size()
			n += 1 + l + sovFilecontext(uint64(l))
		}
	}
	return n
}

func (m *ContextDependentVersion_Column) Size() (n int) {
	var l int
	_ = l
	if m.Offset != 0 {
		n += 1 + sovFilecontext(uint64(m.Offset))
	}
	l = len(m.LinkedContext)
	if l > 0 {
		n += 1 + l + sovFilecontext(uint64(l))
	}
	return n
}

func (m *ContextDependentVersion_Row) Size() (n int) {
	var l int
	_ = l
	l = len(m.SourceContext)
	if l > 0 {
		n += 1 + l + sovFilecontext(uint64(l))
	}
	if len(m.Column) > 0 {
		for _, e := range m.Column {
			l = e.Size()
			n += 1 + l + sovFilecontext(uint64(l))
		}
	}
	if m.AlwaysProcess {
		n += 2
	}
	return n
}

func sovFilecontext(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFilecontext(x uint64) (n int) {
	return sovFilecontext(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContextDependentVersion) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilecontext
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContextDependentVersion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContextDependentVersion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Row", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilecontext
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFilecontext
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Row = append(m.Row, &ContextDependentVersion_Row{})
			if err := m.Row[len(m.Row)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFilecontext(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFilecontext
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
func (m *ContextDependentVersion_Column) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilecontext
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Column: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Column: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilecontext
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Offset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LinkedContext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilecontext
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFilecontext
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LinkedContext = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFilecontext(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFilecontext
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
func (m *ContextDependentVersion_Row) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilecontext
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Row: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Row: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceContext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilecontext
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFilecontext
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceContext = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Column", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilecontext
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFilecontext
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Column = append(m.Column, &ContextDependentVersion_Column{})
			if err := m.Column[len(m.Column)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlwaysProcess", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilecontext
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AlwaysProcess = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFilecontext(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFilecontext
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
func skipFilecontext(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFilecontext
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowFilecontext
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowFilecontext
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFilecontext
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFilecontext
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipFilecontext(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthFilecontext = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFilecontext   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorFilecontext = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x8e, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x7d, 0x09, 0x06, 0x9d, 0x52, 0x17, 0x59, 0x68, 0x28, 0x18, 0x83, 0x20, 0x04, 0x84,
	0x29, 0xea, 0xce, 0xa5, 0x11, 0xdc, 0xca, 0x2c, 0xdc, 0x96, 0x9a, 0xbc, 0xd4, 0xd0, 0xe9, 0xbc,
	0x90, 0x99, 0x1a, 0x7b, 0x93, 0x5e, 0xc5, 0x1b, 0xb8, 0xf4, 0x08, 0x12, 0x2f, 0x22, 0xc9, 0x0c,
	0xd2, 0x8d, 0xe0, 0x6a, 0x98, 0xef, 0xfd, 0x3f, 0xff, 0xc7, 0x4e, 0x97, 0x1b, 0xf3, 0x82, 0xd3,
	0xba, 0x21, 0x43, 0xd3, 0xb2, 0x92, 0x98, 0x93, 0x32, 0xf8, 0x66, 0xf8, 0x40, 0xc2, 0xd1, 0x70,
	0xb6, 0x9f, 0xf3, 0x77, 0x8f, 0x9d, 0x64, 0xf6, 0x7c, 0x8f, 0x35, 0xaa, 0x02, 0x95, 0x79, 0xc2,
	0x46, 0x57, 0xa4, 0xc2, 0x5b, 0xe6, 0x37, 0xd4, 0x46, 0x90, 0xf8, 0xe9, 0xe8, 0x3a, 0xe5, 0x3b,
	0x35, 0xfe, 0x47, 0x85, 0x0b, 0x6a, 0x45, 0x5f, 0x9a, 0x3c, 0xb0, 0x20, 0x23, 0xb9, 0x5e, 0xa9,
	0xf0, 0x98, 0x05, 0x54, 0x96, 0x1a, 0x4d, 0x04, 0x09, 0xa4, 0xfb, 0xc2, 0xfd, 0xc2, 0x0b, 0x76,
	0x24, 0x2b, 0xb5, 0xc4, 0x62, 0xe6, 0xf4, 0x22, 0x2f, 0x81, 0xf4, 0x50, 0x8c, 0x2d, 0x75, 0x0b,
	0x93, 0x2d, 0x30, 0x5f, 0x50, 0xdb, 0xc7, 0x35, 0xad, 0x9b, 0x1c, 0x7f, 0xe3, 0x60, 0xe3, 0x96,
	0xba, 0x78, 0x98, 0xb1, 0x20, 0x1f, 0x76, 0x23, 0x6f, 0xd0, 0xbe, 0xfc, 0x97, 0xb6, 0x55, 0x15,
	0xae, 0xda, 0x6f, 0xcd, 0x65, 0x3b, 0xdf, 0xe8, 0x59, 0xdd, 0x50, 0x8e, 0x5a, 0x47, 0x7e, 0x02,
	0xe9, 0x81, 0x18, 0x5b, 0xfa, 0x68, 0xe1, 0xdd, 0xd5, 0x47, 0x17, 0xc3, 0x67, 0x17, 0xc3, 0x57,
	0x17, 0xc3, 0xf6, 0x3b, 0xde, 0x63, 0x67, 0x39, 0xad, 0xf8, 0x82, 0x68, 0x21, 0x91, 0x17, 0xf8,
	0x6a, 0x88, 0xa4, 0xde, 0x15, 0x78, 0x0e, 0x86, 0xe7, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb5,
	0xb1, 0xc6, 0xb2, 0xa3, 0x01, 0x00, 0x00,
}
