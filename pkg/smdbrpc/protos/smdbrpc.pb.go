// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: smdbrpc/protos/smdbrpc.proto

package execinfrapb

/*
	package cockroach.sql.distsqlrun;
*/

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type HLCTimestamp struct {
	Walltime *int64 `protobuf:"varint,1,opt,name=walltime" json:"walltime,omitempty"`
}

func (m *HLCTimestamp) Reset()         { *m = HLCTimestamp{} }
func (m *HLCTimestamp) String() string { return proto.CompactTextString(m) }
func (*HLCTimestamp) ProtoMessage()    {}
func (*HLCTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_smdbrpc_3ce012c60fc25c9d, []int{0}
}
func (m *HLCTimestamp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HLCTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *HLCTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HLCTimestamp.Merge(dst, src)
}
func (m *HLCTimestamp) XXX_Size() int {
	return m.Size()
}
func (m *HLCTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_HLCTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_HLCTimestamp proto.InternalMessageInfo

type KVPair struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KVPair) Reset()         { *m = KVPair{} }
func (m *KVPair) String() string { return proto.CompactTextString(m) }
func (*KVPair) ProtoMessage()    {}
func (*KVPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_smdbrpc_3ce012c60fc25c9d, []int{1}
}
func (m *KVPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KVPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *KVPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVPair.Merge(dst, src)
}
func (m *KVPair) XXX_Size() int {
	return m.Size()
}
func (m *KVPair) XXX_DiscardUnknown() {
	xxx_messageInfo_KVPair.DiscardUnknown(m)
}

var xxx_messageInfo_KVPair proto.InternalMessageInfo

type HotshardRequest struct {
	Hlctimestamp *HLCTimestamp `protobuf:"bytes,1,opt,name=hlctimestamp" json:"hlctimestamp,omitempty"`
	WriteKeyset  []*KVPair     `protobuf:"bytes,2,rep,name=write_keyset,json=writeKeyset" json:"write_keyset,omitempty"`
	ReadKeyset   [][]byte      `protobuf:"bytes,3,rep,name=read_keyset,json=readKeyset" json:"read_keyset,omitempty"`
}

func (m *HotshardRequest) Reset()         { *m = HotshardRequest{} }
func (m *HotshardRequest) String() string { return proto.CompactTextString(m) }
func (*HotshardRequest) ProtoMessage()    {}
func (*HotshardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_smdbrpc_3ce012c60fc25c9d, []int{2}
}
func (m *HotshardRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HotshardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *HotshardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotshardRequest.Merge(dst, src)
}
func (m *HotshardRequest) XXX_Size() int {
	return m.Size()
}
func (m *HotshardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HotshardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HotshardRequest proto.InternalMessageInfo

// The response message containing the greetings
type HotshardReply struct {
	IsCommitted  *bool     `protobuf:"varint,1,opt,name=is_committed,json=isCommitted" json:"is_committed,omitempty"`
	ReadValueset []*KVPair `protobuf:"bytes,2,rep,name=read_valueset,json=readValueset" json:"read_valueset,omitempty"`
}

func (m *HotshardReply) Reset()         { *m = HotshardReply{} }
func (m *HotshardReply) String() string { return proto.CompactTextString(m) }
func (*HotshardReply) ProtoMessage()    {}
func (*HotshardReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_smdbrpc_3ce012c60fc25c9d, []int{3}
}
func (m *HotshardReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HotshardReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *HotshardReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotshardReply.Merge(dst, src)
}
func (m *HotshardReply) XXX_Size() int {
	return m.Size()
}
func (m *HotshardReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HotshardReply.DiscardUnknown(m)
}

var xxx_messageInfo_HotshardReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HLCTimestamp)(nil), "smdbrpc.HLCTimestamp")
	proto.RegisterType((*KVPair)(nil), "smdbrpc.KVPair")
	proto.RegisterType((*HotshardRequest)(nil), "smdbrpc.HotshardRequest")
	proto.RegisterType((*HotshardReply)(nil), "smdbrpc.HotshardReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HotshardGatewayClient is the client API for HotshardGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HotshardGatewayClient interface {
	// Sends a greeting
	ContactHotshard(ctx context.Context, in *HotshardRequest, opts ...grpc.CallOption) (*HotshardReply, error)
}

type hotshardGatewayClient struct {
	cc *grpc.ClientConn
}

func NewHotshardGatewayClient(cc *grpc.ClientConn) HotshardGatewayClient {
	return &hotshardGatewayClient{cc}
}

func (c *hotshardGatewayClient) ContactHotshard(ctx context.Context, in *HotshardRequest, opts ...grpc.CallOption) (*HotshardReply, error) {
	out := new(HotshardReply)
	err := c.cc.Invoke(ctx, "/smdbrpc.HotshardGateway/ContactHotshard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotshardGatewayServer is the server API for HotshardGateway service.
type HotshardGatewayServer interface {
	// Sends a greeting
	ContactHotshard(context.Context, *HotshardRequest) (*HotshardReply, error)
}

func RegisterHotshardGatewayServer(s *grpc.Server, srv HotshardGatewayServer) {
	s.RegisterService(&_HotshardGateway_serviceDesc, srv)
}

func _HotshardGateway_ContactHotshard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotshardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotshardGatewayServer).ContactHotshard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smdbrpc.HotshardGateway/ContactHotshard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotshardGatewayServer).ContactHotshard(ctx, req.(*HotshardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HotshardGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smdbrpc.HotshardGateway",
	HandlerType: (*HotshardGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContactHotshard",
			Handler:    _HotshardGateway_ContactHotshard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smdbrpc/protos/smdbrpc.proto",
}

func (m *HLCTimestamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HLCTimestamp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Walltime != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSmdbrpc(dAtA, i, uint64(*m.Walltime))
	}
	return i, nil
}

func (m *KVPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KVPair) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSmdbrpc(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if m.Value != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSmdbrpc(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *HotshardRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HotshardRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Hlctimestamp != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSmdbrpc(dAtA, i, uint64(m.Hlctimestamp.Size()))
		n1, err := m.Hlctimestamp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.WriteKeyset) > 0 {
		for _, msg := range m.WriteKeyset {
			dAtA[i] = 0x12
			i++
			i = encodeVarintSmdbrpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ReadKeyset) > 0 {
		for _, b := range m.ReadKeyset {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintSmdbrpc(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func (m *HotshardReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HotshardReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.IsCommitted != nil {
		dAtA[i] = 0x8
		i++
		if *m.IsCommitted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.ReadValueset) > 0 {
		for _, msg := range m.ReadValueset {
			dAtA[i] = 0x12
			i++
			i = encodeVarintSmdbrpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintSmdbrpc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HLCTimestamp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Walltime != nil {
		n += 1 + sovSmdbrpc(uint64(*m.Walltime))
	}
	return n
}

func (m *KVPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = len(m.Key)
		n += 1 + l + sovSmdbrpc(uint64(l))
	}
	if m.Value != nil {
		l = len(m.Value)
		n += 1 + l + sovSmdbrpc(uint64(l))
	}
	return n
}

func (m *HotshardRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Hlctimestamp != nil {
		l = m.Hlctimestamp.Size()
		n += 1 + l + sovSmdbrpc(uint64(l))
	}
	if len(m.WriteKeyset) > 0 {
		for _, e := range m.WriteKeyset {
			l = e.Size()
			n += 1 + l + sovSmdbrpc(uint64(l))
		}
	}
	if len(m.ReadKeyset) > 0 {
		for _, b := range m.ReadKeyset {
			l = len(b)
			n += 1 + l + sovSmdbrpc(uint64(l))
		}
	}
	return n
}

func (m *HotshardReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsCommitted != nil {
		n += 2
	}
	if len(m.ReadValueset) > 0 {
		for _, e := range m.ReadValueset {
			l = e.Size()
			n += 1 + l + sovSmdbrpc(uint64(l))
		}
	}
	return n
}

func sovSmdbrpc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSmdbrpc(x uint64) (n int) {
	return sovSmdbrpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HLCTimestamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSmdbrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HLCTimestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HLCTimestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Walltime", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmdbrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Walltime = &v
		default:
			iNdEx = preIndex
			skippy, err := skipSmdbrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSmdbrpc
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
func (m *KVPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSmdbrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KVPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KVPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmdbrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmdbrpc
			}
			postIndex := iNdEx + byteLen
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
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmdbrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmdbrpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSmdbrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSmdbrpc
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
func (m *HotshardRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSmdbrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HotshardRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HotshardRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hlctimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmdbrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSmdbrpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Hlctimestamp == nil {
				m.Hlctimestamp = &HLCTimestamp{}
			}
			if err := m.Hlctimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WriteKeyset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmdbrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSmdbrpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WriteKeyset = append(m.WriteKeyset, &KVPair{})
			if err := m.WriteKeyset[len(m.WriteKeyset)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadKeyset", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmdbrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmdbrpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReadKeyset = append(m.ReadKeyset, make([]byte, postIndex-iNdEx))
			copy(m.ReadKeyset[len(m.ReadKeyset)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSmdbrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSmdbrpc
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
func (m *HotshardReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSmdbrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HotshardReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HotshardReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsCommitted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmdbrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.IsCommitted = &b
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadValueset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmdbrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSmdbrpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReadValueset = append(m.ReadValueset, &KVPair{})
			if err := m.ReadValueset[len(m.ReadValueset)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSmdbrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSmdbrpc
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
func skipSmdbrpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSmdbrpc
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
					return 0, ErrIntOverflowSmdbrpc
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
					return 0, ErrIntOverflowSmdbrpc
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSmdbrpc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSmdbrpc
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
				next, err := skipSmdbrpc(dAtA[start:])
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
	ErrInvalidLengthSmdbrpc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSmdbrpc   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("smdbrpc/protos/smdbrpc.proto", fileDescriptor_smdbrpc_3ce012c60fc25c9d)
}

var fileDescriptor_smdbrpc_3ce012c60fc25c9d = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x6b, 0xea, 0x40,
	0x10, 0xc7, 0x13, 0xc3, 0x7b, 0x4f, 0x26, 0x11, 0x1f, 0xcb, 0x7b, 0x25, 0x48, 0xd9, 0xda, 0x9c,
	0xa4, 0x50, 0x2d, 0xa1, 0x97, 0x5e, 0x2b, 0xa5, 0x82, 0x3d, 0x94, 0x50, 0xa4, 0xf4, 0x22, 0x6b,
	0x32, 0xc5, 0xc5, 0xc4, 0xa4, 0xd9, 0xb5, 0x36, 0xdf, 0xa2, 0x5f, 0xa1, 0xdf, 0xc6, 0xa3, 0x47,
	0x8f, 0x6d, 0xfc, 0x22, 0xc5, 0xd5, 0x68, 0x85, 0xd2, 0x5b, 0xfe, 0x33, 0xbf, 0x99, 0xfc, 0x76,
	0xe0, 0x50, 0x44, 0xc1, 0x20, 0x4d, 0xfc, 0x56, 0x92, 0xc6, 0x32, 0x16, 0xad, 0x4d, 0x6c, 0xaa,
	0x48, 0xfe, 0x6c, 0xa2, 0x73, 0x02, 0x56, 0xe7, 0xa6, 0x7d, 0xc7, 0x23, 0x14, 0x92, 0x45, 0x09,
	0xa9, 0x41, 0x79, 0xca, 0xc2, 0x50, 0xf2, 0x08, 0x6d, 0xbd, 0xae, 0x37, 0x0c, 0x6f, 0x9b, 0x9d,
	0x33, 0xf8, 0xdd, 0xed, 0xdd, 0x32, 0x9e, 0x92, 0xbf, 0x60, 0x8c, 0x30, 0x53, 0x80, 0xe5, 0xad,
	0x3e, 0xc9, 0x3f, 0xf8, 0xf5, 0xcc, 0xc2, 0x09, 0xda, 0x25, 0x55, 0x5b, 0x07, 0xe7, 0x4d, 0x87,
	0x6a, 0x27, 0x96, 0x62, 0xc8, 0xd2, 0xc0, 0xc3, 0xa7, 0x09, 0x0a, 0x49, 0x2e, 0xc0, 0x1a, 0x86,
	0xbe, 0x2c, 0xfe, 0xa8, 0x96, 0x98, 0xee, 0xff, 0x66, 0x21, 0xf8, 0x55, 0xc7, 0xdb, 0x43, 0x89,
	0x0b, 0xd6, 0x34, 0xe5, 0x12, 0xfb, 0x23, 0xcc, 0x04, 0x4a, 0xbb, 0x54, 0x37, 0x1a, 0xa6, 0x5b,
	0xdd, 0x8e, 0xae, 0xed, 0x3c, 0x53, 0x41, 0x5d, 0xc5, 0x90, 0x23, 0x30, 0x53, 0x64, 0x41, 0x31,
	0x62, 0xd4, 0x8d, 0x86, 0xe5, 0xc1, 0xaa, 0xb4, 0x06, 0x9c, 0x21, 0x54, 0x76, 0x8a, 0x49, 0x98,
	0x91, 0x63, 0xb0, 0xb8, 0xe8, 0xfb, 0x71, 0x14, 0x71, 0x29, 0x31, 0x50, 0x82, 0x65, 0xcf, 0xe4,
	0xa2, 0x5d, 0x94, 0xc8, 0x39, 0x54, 0xd4, 0x52, 0xf5, 0xca, 0x1f, 0x4c, 0xac, 0x15, 0xd5, 0xdb,
	0x40, 0xee, 0xfd, 0xee, 0x18, 0xd7, 0x4c, 0xe2, 0x94, 0x65, 0xe4, 0x0a, 0xaa, 0xed, 0x78, 0x2c,
	0x99, 0x2f, 0x8b, 0x0e, 0xb1, 0x77, 0x97, 0xd8, 0xbf, 0x5c, 0xed, 0xe0, 0x9b, 0x4e, 0x12, 0x66,
	0x8e, 0x76, 0x79, 0x3a, 0xfb, 0xa0, 0xda, 0x2c, 0xa7, 0xfa, 0x3c, 0xa7, 0xfa, 0x22, 0xa7, 0xfa,
	0x7b, 0x4e, 0xf5, 0xd7, 0x25, 0xd5, 0xe6, 0x4b, 0xaa, 0x2d, 0x96, 0x54, 0x7b, 0x30, 0xf1, 0x05,
	0x7d, 0x3e, 0x7e, 0x4c, 0x59, 0x32, 0xf8, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x7a, 0x68, 0xaa,
	0x1c, 0x02, 0x00, 0x00,
}
