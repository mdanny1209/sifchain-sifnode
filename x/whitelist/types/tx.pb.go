// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/whitelist/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgUpdateWhitelist struct {
	From  string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Units int64  `protobuf:"varint,3,opt,name=units,proto3" json:"units,omitempty"`
}

func (m *MsgUpdateWhitelist) Reset()         { *m = MsgUpdateWhitelist{} }
func (m *MsgUpdateWhitelist) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateWhitelist) ProtoMessage()    {}
func (*MsgUpdateWhitelist) Descriptor() ([]byte, []int) {
	return fileDescriptor_54b6ca421e01b902, []int{0}
}
func (m *MsgUpdateWhitelist) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateWhitelist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateWhitelist.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateWhitelist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateWhitelist.Merge(m, src)
}
func (m *MsgUpdateWhitelist) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateWhitelist) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateWhitelist.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateWhitelist proto.InternalMessageInfo

func (m *MsgUpdateWhitelist) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MsgUpdateWhitelist) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *MsgUpdateWhitelist) GetUnits() int64 {
	if m != nil {
		return m.Units
	}
	return 0
}

type MsgUpdateWhitelistResponse struct {
}

func (m *MsgUpdateWhitelistResponse) Reset()         { *m = MsgUpdateWhitelistResponse{} }
func (m *MsgUpdateWhitelistResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateWhitelistResponse) ProtoMessage()    {}
func (*MsgUpdateWhitelistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_54b6ca421e01b902, []int{1}
}
func (m *MsgUpdateWhitelistResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateWhitelistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateWhitelistResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateWhitelistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateWhitelistResponse.Merge(m, src)
}
func (m *MsgUpdateWhitelistResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateWhitelistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateWhitelistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateWhitelistResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateWhitelist)(nil), "sifnode.whitelist.v1.MsgUpdateWhitelist")
	proto.RegisterType((*MsgUpdateWhitelistResponse)(nil), "sifnode.whitelist.v1.MsgUpdateWhitelistResponse")
}

func init() { proto.RegisterFile("sifnode/whitelist/v1/tx.proto", fileDescriptor_54b6ca421e01b902) }

var fileDescriptor_54b6ca421e01b902 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xce, 0x4c, 0xcb,
	0xcb, 0x4f, 0x49, 0xd5, 0x2f, 0xcf, 0xc8, 0x2c, 0x49, 0xcd, 0xc9, 0x2c, 0x2e, 0xd1, 0x2f, 0x33,
	0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x81, 0x4a, 0xeb, 0xc1, 0xa5,
	0xf5, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x0a, 0xf4, 0x41, 0x2c, 0x88, 0x5a,
	0xa5, 0x10, 0x2e, 0x21, 0xdf, 0xe2, 0xf4, 0xd0, 0x82, 0x94, 0xc4, 0x92, 0xd4, 0x70, 0x98, 0x72,
	0x21, 0x21, 0x2e, 0x96, 0xb4, 0xa2, 0xfc, 0x5c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30,
	0x5b, 0x48, 0x84, 0x8b, 0x35, 0x25, 0x35, 0x2f, 0x3f, 0x57, 0x82, 0x09, 0x2c, 0x08, 0xe1, 0x80,
	0x44, 0x4b, 0xf3, 0x32, 0x4b, 0x8a, 0x25, 0x98, 0x15, 0x18, 0x35, 0x98, 0x83, 0x20, 0x1c, 0x25,
	0x19, 0x2e, 0x29, 0x4c, 0x53, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x8d, 0xca, 0xb8,
	0x98, 0x7d, 0x8b, 0xd3, 0x85, 0xf2, 0xb9, 0xf8, 0xd1, 0xed, 0xd5, 0xd0, 0xc3, 0xe6, 0x74, 0x3d,
	0x4c, 0xb3, 0xa4, 0x0c, 0x88, 0x55, 0x09, 0xb3, 0x55, 0x89, 0xc1, 0xc9, 0xfd, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0x83, 0x33, 0xd3, 0x92, 0x33, 0x12, 0x33, 0xf3, 0xf4, 0x61, 0x81, 0x5c, 0x81,
	0x14, 0xcc, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xb0, 0x33, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x2d, 0x43, 0xe2, 0xce, 0x88, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	UpdateWhitelist(ctx context.Context, in *MsgUpdateWhitelist, opts ...grpc.CallOption) (*MsgUpdateWhitelistResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateWhitelist(ctx context.Context, in *MsgUpdateWhitelist, opts ...grpc.CallOption) (*MsgUpdateWhitelistResponse, error) {
	out := new(MsgUpdateWhitelistResponse)
	err := c.cc.Invoke(ctx, "/sifnode.whitelist.v1.Msg/UpdateWhitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	UpdateWhitelist(context.Context, *MsgUpdateWhitelist) (*MsgUpdateWhitelistResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateWhitelist(ctx context.Context, req *MsgUpdateWhitelist) (*MsgUpdateWhitelistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWhitelist not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateWhitelist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sifnode.whitelist.v1.Msg/UpdateWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateWhitelist(ctx, req.(*MsgUpdateWhitelist))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sifnode.whitelist.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateWhitelist",
			Handler:    _Msg_UpdateWhitelist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sifnode/whitelist/v1/tx.proto",
}

func (m *MsgUpdateWhitelist) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateWhitelist) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateWhitelist) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Units != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Units))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateWhitelistResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateWhitelistResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateWhitelistResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateWhitelist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Units != 0 {
		n += 1 + sovTx(uint64(m.Units))
	}
	return n
}

func (m *MsgUpdateWhitelistResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateWhitelist) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateWhitelist: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateWhitelist: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Units", wireType)
			}
			m.Units = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Units |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateWhitelistResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateWhitelistResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateWhitelistResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
