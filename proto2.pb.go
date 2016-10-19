// Code generated by protoc-gen-go.
// source: proto2.proto
// DO NOT EDIT!

/*
Package prototest is a generated protocol buffer package.

It is generated from these files:
	proto2.proto
	proto3.proto

It has these top-level messages:
	Small2
	Inner2
	Small3
	Inner3
	Group1
	Group2
*/
package prototest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import unsafe "unsafe"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = unsafe.Pointer(nil)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Small2_Enum int32

const (
	Small2_EA Small2_Enum = 0
	Small2_EB Small2_Enum = 1
	Small2_EC Small2_Enum = 2
)

var Small2_Enum_name = map[int32]string{
	0: "EA",
	1: "EB",
	2: "EC",
}
var Small2_Enum_value = map[string]int32{
	"EA": 0,
	"EB": 1,
	"EC": 2,
}

func (x Small2_Enum) Enum() *Small2_Enum {
	p := new(Small2_Enum)
	*p = x
	return p
}
func (x Small2_Enum) String() string {
	return proto.EnumName(Small2_Enum_name, int32(x))
}
func (x *Small2_Enum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Small2_Enum_value, data, "Small2_Enum")
	if err != nil {
		return err
	}
	*x = Small2_Enum(value)
	return nil
}
func (Small2_Enum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Small2 struct {
	A *int32           `protobuf:"varint,1,req,name=a" json:"a,omitempty"`
	B *uint32          `protobuf:"fixed32,2,opt,name=b" json:"b,omitempty"`
	C *string          `protobuf:"bytes,3,opt,name=c" json:"c,omitempty"`
	D []float64        `protobuf:"fixed64,4,rep,name=d" json:"d,omitempty"`
	E *Inner2          `protobuf:"bytes,5,opt,name=e" json:"e,omitempty"`
	F []*Inner2        `protobuf:"bytes,6,rep,name=f" json:"f,omitempty"`
	G *Small2_G        `protobuf:"group,8,opt,name=G,json=g" json:"g,omitempty"`
	H []*Small2_H      `protobuf:"group,10,rep,name=H,json=h" json:"h,omitempty"`
	I *Small2_Enum     `protobuf:"varint,12,opt,name=i,enum=prototest.Small2_Enum" json:"i,omitempty"`
	J map[int32]string `protobuf:"bytes,13,rep,name=j" json:"j,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to Oneof:
	//	*Small2_X
	//	*Small2_Y
	Oneof            isSmall2_Oneof `protobuf_oneof:"Oneof"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Small2) Reset()                    { *m = Small2{} }
func (m *Small2) String() string            { return proto.CompactTextString(m) }
func (*Small2) ProtoMessage()               {}
func (*Small2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isSmall2_Oneof interface {
	isSmall2_Oneof()
}

type Small2_X struct {
	X int64 `protobuf:"varint,14,opt,name=x,oneof"`
}
type Small2_Y struct {
	Y string `protobuf:"bytes,15,opt,name=y,oneof"`
}

func (*Small2_X) isSmall2_Oneof() {}
func (*Small2_Y) isSmall2_Oneof() {}

func (m *Small2) GetOneof() isSmall2_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (m *Small2) GetA() int32 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func (m *Small2) GetB() uint32 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

func (m *Small2) GetC() string {
	if m != nil && m.C != nil {
		return *m.C
	}
	return ""
}

func (m *Small2) GetD() []float64 {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *Small2) GetE() *Inner2 {
	if m != nil {
		return m.E
	}
	return nil
}

func (m *Small2) GetF() []*Inner2 {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *Small2) GetG() *Small2_G {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *Small2) GetH() []*Small2_H {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *Small2) GetI() Small2_Enum {
	if m != nil && m.I != nil {
		return *m.I
	}
	return Small2_EA
}

func (m *Small2) GetJ() map[int32]string {
	if m != nil {
		return m.J
	}
	return nil
}

func (m *Small2) GetX() int64 {
	if x, ok := m.GetOneof().(*Small2_X); ok {
		return x.X
	}
	return 0
}

func (m *Small2) GetY() string {
	if x, ok := m.GetOneof().(*Small2_Y); ok {
		return x.Y
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Small2) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Small2_OneofMarshaler, _Small2_OneofUnmarshaler, _Small2_OneofSizer, []interface{}{
		(*Small2_X)(nil),
		(*Small2_Y)(nil),
	}
}

func _Small2_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Small2)
	// Oneof
	switch x := m.Oneof.(type) {
	case *Small2_X:
		b.EncodeVarint(14<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.X))
	case *Small2_Y:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Y)
	case nil:
	default:
		return fmt.Errorf("Small2.Oneof has unexpected type %T", x)
	}
	return nil
}

func _Small2_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Small2)
	switch tag {
	case 14: // Oneof.x
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Oneof = &Small2_X{int64(x)}
		return true, err
	case 15: // Oneof.y
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Oneof = &Small2_Y{x}
		return true, err
	default:
		return false, nil
	}
}

func _Small2_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Small2)
	// Oneof
	switch x := m.Oneof.(type) {
	case *Small2_X:
		n += proto.SizeVarint(14<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.X))
	case *Small2_Y:
		n += proto.SizeVarint(15<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Y)))
		n += len(x.Y)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func (m *Small2) MergeFullCustom(b []byte) error {
	for len(b) > 0 {
		x, n := proto.DecodeVarint(b)
		if n == 0 {
			return proto.ErrInternalBadWireType
		}
		b = b[n:]
		switch x >> 3 {
		case 1:
			x, n = proto.DecodeVarint(b)
			if n == 0 {
				return proto.ErrInternalBadWireType
			}
			b = b[n:]
			v := int32(x)
			m.A = &v
		case 2:
			if len(b) < 4 {
				return proto.ErrInternalBadWireType
			}
			v := uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
			b = b[4:]
			m.B = &v
		case 3:
			x, n = proto.DecodeVarint(b)
			if n == 0 {
				return proto.ErrInternalBadWireType
			}
			b = b[n:]
			if uint64(len(b)) < x {
				return proto.ErrInternalBadWireType
			}
			v := string(b[:x])
			b = b[x:]
			m.C = &v
		case 4:
			if len(b) < 8 {
				return proto.ErrInternalBadWireType
			}
			v := math.Float64frombits(uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56)
			b = b[8:]
			m.D = append(m.D, v)
		case 5:
			x, n = proto.DecodeVarint(b)
			if n == 0 {
				return proto.ErrInternalBadWireType
			}
			b = b[n:]
			if uint64(len(b)) < x {
				return proto.ErrInternalBadWireType
			}
			var v Inner2
			if err := v.MergeFullCustom(b[:x]); err != nil {
				return err
			}
			b = b[x:]
			m.E = &v
		case 6:
			x, n = proto.DecodeVarint(b)
			if n == 0 {
				return proto.ErrInternalBadWireType
			}
			b = b[n:]
			if uint64(len(b)) < x {
				return proto.ErrInternalBadWireType
			}
			var v Inner2
			if err := v.MergeFullCustom(b[:x]); err != nil {
				return err
			}
			b = b[x:]
			m.F = append(m.F, &v)
		case 8:
			i, j := proto.FindEndGroup(b)
			if i < 0 {
				return proto.ErrInternalBadWireType
			}
			var v Small2_G
			if err := v.MergeFullCustom(b[:i]); err != nil {
				return err
			}
			b = b[j:]
			m.G = &v
		case 10:
			i, j := proto.FindEndGroup(b)
			if i < 0 {
				return proto.ErrInternalBadWireType
			}
			var v Small2_H
			if err := v.MergeFullCustom(b[:i]); err != nil {
				return err
			}
			b = b[j:]
			m.H = append(m.H, &v)
		case 12:
			x, n = proto.DecodeVarint(b)
			if n == 0 {
				return proto.ErrInternalBadWireType
			}
			b = b[n:]
			v := Small2_Enum(x)
			m.I = &v
		case 13:
			x, n = proto.DecodeVarint(b)
			if n == 0 {
				return proto.ErrInternalBadWireType
			}
			b = b[n:]
			if uint64(len(b)) < x {
				return proto.ErrInternalBadWireType
			}
			b = b[x:]
		case 14:
			x, n = proto.DecodeVarint(b)
			if n == 0 {
				return proto.ErrInternalBadWireType
			}
			b = b[n:]
			v := int64(x)
			w := Small2_X{X: v}
			m.Oneof = &w
		case 15:
			x, n = proto.DecodeVarint(b)
			if n == 0 {
				return proto.ErrInternalBadWireType
			}
			b = b[n:]
			if uint64(len(b)) < x {
				return proto.ErrInternalBadWireType
			}
			v := string(b[:x])
			b = b[x:]
			w := Small2_Y{Y: v}
			m.Oneof = &w
		default:
			return proto.ErrInternalBadWireType
		}
	}
	return nil
}

var XXX_Unpack_Small2 = proto.UnpackMessageInfo{
	Make: func() unsafe.Pointer {
		return unsafe.Pointer(new(Small2))
	},
	Dense: []proto.UnpackFieldInfo{
		1: {
			Offset: unsafe.Offsetof(Small2{}.A),
			Unpack: proto.UnpackInt32_2,
		},
		2: {
			Offset: unsafe.Offsetof(Small2{}.B),
			Unpack: proto.UnpackFixed32_2,
		},
		3: {
			Offset: unsafe.Offsetof(Small2{}.C),
			Unpack: proto.UnpackString_2,
		},
		4: {
			Offset: unsafe.Offsetof(Small2{}.D),
			Unpack: proto.UnpackDouble_R,
		},
		5: {
			Offset: unsafe.Offsetof(Small2{}.E),
			Sub:    &XXX_Unpack_Inner2,
			Unpack: proto.UnpackMessage,
		},
		6: {
			Offset: unsafe.Offsetof(Small2{}.F),
			Sub:    &XXX_Unpack_Inner2,
			Unpack: proto.UnpackMessage_R,
		},
		8: {
			Offset: unsafe.Offsetof(Small2{}.G),
			Sub:    &XXX_Unpack_Small2_G,
			Unpack: proto.UnpackGroup,
		},
		10: {
			Offset: unsafe.Offsetof(Small2{}.H),
			Sub:    &XXX_Unpack_Small2_H,
			Unpack: proto.UnpackGroup_R,
		},
		12: {
			Offset: unsafe.Offsetof(Small2{}.I),
			Unpack: proto.UnpackEnum_2,
		},
	},
	Sparse:             nil,
	UnrecognizedOffset: unsafe.Offsetof(Small2{}.XXX_unrecognized),
}

func (m *Small2) MergeTableDriven(b []byte) error {
	return proto.UnmarshalMsg(b, unsafe.Pointer(m), &XXX_Unpack_Small2)
}

type Small2_G struct {
	X                *float32 `protobuf:"fixed32,9,opt,name=x" json:"x,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Small2_G) Reset()                    { *m = Small2_G{} }
func (m *Small2_G) String() string            { return proto.CompactTextString(m) }
func (*Small2_G) ProtoMessage()               {}
func (*Small2_G) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Small2_G) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Small2_G) MergeFullCustom(b []byte) error {
	for len(b) > 0 {
		x, n := proto.DecodeVarint(b)
		if n == 0 {
			return proto.ErrInternalBadWireType
		}
		b = b[n:]
		switch x >> 3 {
		case 9:
			if len(b) < 4 {
				return proto.ErrInternalBadWireType
			}
			v := math.Float32frombits(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24)
			b = b[4:]
			m.X = &v
		default:
			return proto.ErrInternalBadWireType
		}
	}
	return nil
}

var XXX_Unpack_Small2_G = proto.UnpackMessageInfo{
	Make: func() unsafe.Pointer {
		return unsafe.Pointer(new(Small2_G))
	},
	Dense: []proto.UnpackFieldInfo{
		9: {
			Offset: unsafe.Offsetof(Small2_G{}.X),
			Unpack: proto.UnpackFloat_2,
		},
	},
	Sparse:             nil,
	UnrecognizedOffset: unsafe.Offsetof(Small2_G{}.XXX_unrecognized),
}

func (m *Small2_G) MergeTableDriven(b []byte) error {
	return proto.UnmarshalMsg(b, unsafe.Pointer(m), &XXX_Unpack_Small2_G)
}

type Small2_H struct {
	X                *int32 `protobuf:"zigzag32,11,opt,name=x" json:"x,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Small2_H) Reset()                    { *m = Small2_H{} }
func (m *Small2_H) String() string            { return proto.CompactTextString(m) }
func (*Small2_H) ProtoMessage()               {}
func (*Small2_H) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *Small2_H) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Small2_H) MergeFullCustom(b []byte) error {
	for len(b) > 0 {
		x, n := proto.DecodeVarint(b)
		if n == 0 {
			return proto.ErrInternalBadWireType
		}
		b = b[n:]
		switch x >> 3 {
		case 11:
			x, n = proto.DecodeVarint(b)
			if n == 0 {
				return proto.ErrInternalBadWireType
			}
			b = b[n:]
			v := int32(x>>1) ^ int32(x)<<31>>31
			m.X = &v
		default:
			return proto.ErrInternalBadWireType
		}
	}
	return nil
}

var XXX_Unpack_Small2_H = proto.UnpackMessageInfo{
	Make: func() unsafe.Pointer {
		return unsafe.Pointer(new(Small2_H))
	},
	Dense: []proto.UnpackFieldInfo{
		11: {
			Offset: unsafe.Offsetof(Small2_H{}.X),
			Unpack: proto.UnpackSint32_2,
		},
	},
	Sparse:             nil,
	UnrecognizedOffset: unsafe.Offsetof(Small2_H{}.XXX_unrecognized),
}

func (m *Small2_H) MergeTableDriven(b []byte) error {
	return proto.UnmarshalMsg(b, unsafe.Pointer(m), &XXX_Unpack_Small2_H)
}

type Inner2 struct {
	A                *int64 `protobuf:"varint,1,req,name=a" json:"a,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Inner2) Reset()                    { *m = Inner2{} }
func (m *Inner2) String() string            { return proto.CompactTextString(m) }
func (*Inner2) ProtoMessage()               {}
func (*Inner2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Inner2) GetA() int64 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func (m *Inner2) MergeFullCustom(b []byte) error {
	for len(b) > 0 {
		x, n := proto.DecodeVarint(b)
		if n == 0 {
			return proto.ErrInternalBadWireType
		}
		b = b[n:]
		switch x >> 3 {
		case 1:
			x, n = proto.DecodeVarint(b)
			if n == 0 {
				return proto.ErrInternalBadWireType
			}
			b = b[n:]
			v := int64(x)
			m.A = &v
		default:
			return proto.ErrInternalBadWireType
		}
	}
	return nil
}

var XXX_Unpack_Inner2 = proto.UnpackMessageInfo{
	Make: func() unsafe.Pointer {
		return unsafe.Pointer(new(Inner2))
	},
	Dense: []proto.UnpackFieldInfo{
		1: {
			Offset: unsafe.Offsetof(Inner2{}.A),
			Unpack: proto.UnpackInt64_2,
		},
	},
	Sparse:             nil,
	UnrecognizedOffset: unsafe.Offsetof(Inner2{}.XXX_unrecognized),
}

func (m *Inner2) MergeTableDriven(b []byte) error {
	return proto.UnmarshalMsg(b, unsafe.Pointer(m), &XXX_Unpack_Inner2)
}
func init() {
	proto.RegisterType((*Small2)(nil), "prototest.Small2")
	proto.RegisterType((*Small2_G)(nil), "prototest.Small2.G")
	proto.RegisterType((*Small2_H)(nil), "prototest.Small2.H")
	proto.RegisterType((*Inner2)(nil), "prototest.Inner2")
	proto.RegisterEnum("prototest.Small2_Enum", Small2_Enum_name, Small2_Enum_value)
}

func init() { proto.RegisterFile("proto2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x39, 0x1d, 0x5a, 0xbe, 0x5e, 0xf8, 0xb0, 0x8c, 0x86, 0xdc, 0xb8, 0xd0, 0x91, 0x18,
	0x33, 0xab, 0x2e, 0x1a, 0x17, 0xc6, 0x9d, 0x18, 0x02, 0xba, 0x31, 0x19, 0x9f, 0xa0, 0x40, 0xf9,
	0x27, 0x14, 0x83, 0xc5, 0xb4, 0xcf, 0xe6, 0xcb, 0x99, 0x69, 0x13, 0x35, 0x41, 0x57, 0x67, 0x7e,
	0x37, 0xbf, 0xdc, 0xdc, 0x9c, 0xa1, 0xd6, 0xeb, 0x6e, 0x9b, 0x6d, 0xa3, 0xb0, 0x0c, 0xe9, 0x97,
	0x91, 0x25, 0x6f, 0x59, 0xef, 0x43, 0x90, 0xf7, 0xbc, 0x89, 0xd7, 0xeb, 0x48, 0xb6, 0x08, 0x31,
	0x43, 0x39, 0xda, 0x35, 0x88, 0x2d, 0x8d, 0xd9, 0x51, 0xd0, 0x0d, 0x83, 0xb1, 0xa5, 0x09, 0x0b,
	0x05, 0xed, 0x1b, 0x4c, 0x2c, 0x4d, 0xb9, 0xae, 0x84, 0x86, 0xc1, 0x54, 0x9e, 0x13, 0x12, 0x76,
	0x15, 0x74, 0x33, 0xea, 0x84, 0x5f, 0x9b, 0xc3, 0x87, 0x34, 0x4d, 0x76, 0x91, 0x41, 0x62, 0x85,
	0x19, 0x7b, 0x4a, 0xfc, 0x21, 0xcc, 0xe4, 0x05, 0x61, 0xce, 0xff, 0x14, 0x34, 0x45, 0xc7, 0x3f,
	0x84, 0xea, 0xae, 0x70, 0x68, 0x30, 0xb7, 0xca, 0x82, 0x49, 0x89, 0xdf, 0x95, 0x91, 0xc1, 0x42,
	0x5e, 0x12, 0x96, 0xdc, 0x52, 0xd0, 0xed, 0xa8, 0x7b, 0xa8, 0x0c, 0xd2, 0xfd, 0xc6, 0x60, 0x29,
	0xaf, 0x08, 0x2b, 0xfe, 0x5f, 0x1e, 0xc3, 0x87, 0xd6, 0xe3, 0x20, 0xcd, 0x76, 0x85, 0xc1, 0x4a,
	0xb6, 0x09, 0x39, 0xb7, 0x15, 0xb4, 0x18, 0xd5, 0x0c, 0x72, 0xcb, 0x05, 0x1f, 0xd9, 0x06, 0x2c,
	0x17, 0xa7, 0x1d, 0xc2, 0xd0, 0x16, 0x91, 0xb3, 0xaf, 0xa0, 0x1d, 0x83, 0xdc, 0x8e, 0x46, 0xd5,
	0xa8, 0xa9, 0xa0, 0x3b, 0x76, 0x74, 0x4d, 0x5e, 0xb5, 0x52, 0x06, 0x24, 0x5e, 0x92, 0x82, 0xa1,
	0xa0, 0x5d, 0x63, 0x9f, 0xf2, 0x84, 0xdc, 0xf7, 0x78, 0xbd, 0x4f, 0xca, 0x96, 0x7d, 0x53, 0xc1,
	0xad, 0x73, 0x83, 0xde, 0x19, 0xd5, 0xed, 0xb9, 0xd2, 0x23, 0x67, 0x70, 0x17, 0xd4, 0xca, 0xec,
	0x07, 0x28, 0xf3, 0x3e, 0x70, 0xfa, 0x0d, 0x72, 0x9f, 0xd2, 0x64, 0x3b, 0xeb, 0x75, 0xc9, 0xab,
	0x5a, 0xfc, 0xfe, 0x3c, 0x61, 0x10, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x1f, 0x59, 0xc2,
	0xef, 0x01, 0x00, 0x00,
}
