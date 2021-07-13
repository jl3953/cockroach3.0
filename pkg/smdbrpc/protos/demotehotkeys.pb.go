// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: smdbrpc/protos/demotehotkeys.proto

package execinfrapb

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

type KVVersion struct {
	Key       *string       `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value     []byte        `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Timestamp *HLCTimestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Hotness   *int64        `protobuf:"varint,4,opt,name=hotness" json:"hotness,omitempty"`
}

func (m *KVVersion) Reset()         { *m = KVVersion{} }
func (m *KVVersion) String() string { return proto.CompactTextString(m) }
func (*KVVersion) ProtoMessage()    {}
func (*KVVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_demotehotkeys_280d6bcacc7102eb, []int{0}
}
func (m *KVVersion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KVVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *KVVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVVersion.Merge(dst, src)
}
func (m *KVVersion) XXX_Size() int {
	return m.Size()
}
func (m *KVVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_KVVersion.DiscardUnknown(m)
}

var xxx_messageInfo_KVVersion proto.InternalMessageInfo

type KVDemotionStatus struct {
	Key                   *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	IsSuccessfullyDemoted *bool   `protobuf:"varint,2,opt,name=is_successfully_demoted,json=isSuccessfullyDemoted" json:"is_successfully_demoted,omitempty"`
}

func (m *KVDemotionStatus) Reset()         { *m = KVDemotionStatus{} }
func (m *KVDemotionStatus) String() string { return proto.CompactTextString(m) }
func (*KVDemotionStatus) ProtoMessage()    {}
func (*KVDemotionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_demotehotkeys_280d6bcacc7102eb, []int{1}
}
func (m *KVDemotionStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KVDemotionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *KVDemotionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVDemotionStatus.Merge(dst, src)
}
func (m *KVDemotionStatus) XXX_Size() int {
	return m.Size()
}
func (m *KVDemotionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_KVDemotionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_KVDemotionStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*KVVersion)(nil), "smdbrpc.KVVersion")
	proto.RegisterType((*KVDemotionStatus)(nil), "smdbrpc.KVDemotionStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoteHotkeysGatewayClient is the client API for DemoteHotkeysGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoteHotkeysGatewayClient interface {
	DemoteHotkeys(ctx context.Context, opts ...grpc.CallOption) (DemoteHotkeysGateway_DemoteHotkeysClient, error)
}

type demoteHotkeysGatewayClient struct {
	cc *grpc.ClientConn
}

func NewDemoteHotkeysGatewayClient(cc *grpc.ClientConn) DemoteHotkeysGatewayClient {
	return &demoteHotkeysGatewayClient{cc}
}

func (c *demoteHotkeysGatewayClient) DemoteHotkeys(ctx context.Context, opts ...grpc.CallOption) (DemoteHotkeysGateway_DemoteHotkeysClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DemoteHotkeysGateway_serviceDesc.Streams[0], "/smdbrpc.DemoteHotkeysGateway/DemoteHotkeys", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoteHotkeysGatewayDemoteHotkeysClient{stream}
	return x, nil
}

type DemoteHotkeysGateway_DemoteHotkeysClient interface {
	Send(*KVVersion) error
	Recv() (*KVDemotionStatus, error)
	grpc.ClientStream
}

type demoteHotkeysGatewayDemoteHotkeysClient struct {
	grpc.ClientStream
}

func (x *demoteHotkeysGatewayDemoteHotkeysClient) Send(m *KVVersion) error {
	return x.ClientStream.SendMsg(m)
}

func (x *demoteHotkeysGatewayDemoteHotkeysClient) Recv() (*KVDemotionStatus, error) {
	m := new(KVDemotionStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DemoteHotkeysGatewayServer is the server API for DemoteHotkeysGateway service.
type DemoteHotkeysGatewayServer interface {
	DemoteHotkeys(DemoteHotkeysGateway_DemoteHotkeysServer) error
}

func RegisterDemoteHotkeysGatewayServer(s *grpc.Server, srv DemoteHotkeysGatewayServer) {
	s.RegisterService(&_DemoteHotkeysGateway_serviceDesc, srv)
}

func _DemoteHotkeysGateway_DemoteHotkeys_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoteHotkeysGatewayServer).DemoteHotkeys(&demoteHotkeysGatewayDemoteHotkeysServer{stream})
}

type DemoteHotkeysGateway_DemoteHotkeysServer interface {
	Send(*KVDemotionStatus) error
	Recv() (*KVVersion, error)
	grpc.ServerStream
}

type demoteHotkeysGatewayDemoteHotkeysServer struct {
	grpc.ServerStream
}

func (x *demoteHotkeysGatewayDemoteHotkeysServer) Send(m *KVDemotionStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *demoteHotkeysGatewayDemoteHotkeysServer) Recv() (*KVVersion, error) {
	m := new(KVVersion)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DemoteHotkeysGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smdbrpc.DemoteHotkeysGateway",
	HandlerType: (*DemoteHotkeysGatewayServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DemoteHotkeys",
			Handler:       _DemoteHotkeysGateway_DemoteHotkeys_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "smdbrpc/protos/demotehotkeys.proto",
}

func (m *KVVersion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KVVersion) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDemotehotkeys(dAtA, i, uint64(len(*m.Key)))
		i += copy(dAtA[i:], *m.Key)
	}
	if m.Value != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDemotehotkeys(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.Timestamp != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDemotehotkeys(dAtA, i, uint64(m.Timestamp.Size()))
		n1, err := m.Timestamp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Hotness != nil {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDemotehotkeys(dAtA, i, uint64(*m.Hotness))
	}
	return i, nil
}

func (m *KVDemotionStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KVDemotionStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDemotehotkeys(dAtA, i, uint64(len(*m.Key)))
		i += copy(dAtA[i:], *m.Key)
	}
	if m.IsSuccessfullyDemoted != nil {
		dAtA[i] = 0x10
		i++
		if *m.IsSuccessfullyDemoted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintDemotehotkeys(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *KVVersion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = len(*m.Key)
		n += 1 + l + sovDemotehotkeys(uint64(l))
	}
	if m.Value != nil {
		l = len(m.Value)
		n += 1 + l + sovDemotehotkeys(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovDemotehotkeys(uint64(l))
	}
	if m.Hotness != nil {
		n += 1 + sovDemotehotkeys(uint64(*m.Hotness))
	}
	return n
}

func (m *KVDemotionStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = len(*m.Key)
		n += 1 + l + sovDemotehotkeys(uint64(l))
	}
	if m.IsSuccessfullyDemoted != nil {
		n += 2
	}
	return n
}

func sovDemotehotkeys(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDemotehotkeys(x uint64) (n int) {
	return sovDemotehotkeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KVVersion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemotehotkeys
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
			return fmt.Errorf("proto: KVVersion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KVVersion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemotehotkeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDemotehotkeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Key = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemotehotkeys
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
				return ErrInvalidLengthDemotehotkeys
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemotehotkeys
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
				return ErrInvalidLengthDemotehotkeys
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &HLCTimestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hotness", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemotehotkeys
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
			m.Hotness = &v
		default:
			iNdEx = preIndex
			skippy, err := skipDemotehotkeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDemotehotkeys
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
func (m *KVDemotionStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemotehotkeys
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
			return fmt.Errorf("proto: KVDemotionStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KVDemotionStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemotehotkeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDemotehotkeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Key = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSuccessfullyDemoted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemotehotkeys
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
			m.IsSuccessfullyDemoted = &b
		default:
			iNdEx = preIndex
			skippy, err := skipDemotehotkeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDemotehotkeys
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
func skipDemotehotkeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDemotehotkeys
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
					return 0, ErrIntOverflowDemotehotkeys
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
					return 0, ErrIntOverflowDemotehotkeys
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
				return 0, ErrInvalidLengthDemotehotkeys
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDemotehotkeys
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
				next, err := skipDemotehotkeys(dAtA[start:])
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
	ErrInvalidLengthDemotehotkeys = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDemotehotkeys   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("smdbrpc/protos/demotehotkeys.proto", fileDescriptor_demotehotkeys_280d6bcacc7102eb)
}

var fileDescriptor_demotehotkeys_280d6bcacc7102eb = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xcd, 0x4a, 0x02, 0x51,
	0x14, 0x9e, 0x93, 0x85, 0x79, 0x2d, 0x90, 0x8b, 0xd2, 0x24, 0x71, 0x19, 0x66, 0x35, 0x9b, 0x34,
	0x0c, 0x7a, 0x80, 0x12, 0x12, 0x6c, 0x35, 0x86, 0x8b, 0x10, 0x64, 0x1c, 0x8f, 0x38, 0xe8, 0xcc,
	0x1d, 0xe6, 0xdc, 0xa9, 0x66, 0xdb, 0x13, 0xf4, 0x58, 0x2e, 0x5d, 0xba, 0xac, 0xf1, 0x45, 0x42,
	0x47, 0x4d, 0xa3, 0xdd, 0xf9, 0x7e, 0x38, 0x7c, 0xdf, 0xc7, 0x4c, 0xf2, 0x87, 0x83, 0x28, 0x74,
	0xeb, 0x61, 0x24, 0x95, 0xa4, 0xfa, 0x10, 0x7d, 0xa9, 0x70, 0x2c, 0xd5, 0x04, 0x13, 0xaa, 0xad,
	0x49, 0x9e, 0xdf, 0x78, 0xaa, 0x57, 0x7f, 0xcc, 0x1b, 0x98, 0xd9, 0xcc, 0x0f, 0x60, 0x85, 0x76,
	0xb7, 0x8b, 0x11, 0x79, 0x32, 0xe0, 0x25, 0x96, 0x9b, 0x60, 0xa2, 0x83, 0x01, 0x56, 0xc1, 0x5e,
	0x9d, 0xbc, 0xcc, 0x4e, 0x5e, 0x9d, 0x69, 0x8c, 0xfa, 0x91, 0x01, 0xd6, 0x99, 0x9d, 0x01, 0x7e,
	0xcb, 0x0a, 0xca, 0xf3, 0x91, 0x94, 0xe3, 0x87, 0x7a, 0xce, 0x00, 0xab, 0xd8, 0xa8, 0xd4, 0xb6,
	0x8f, 0x5b, 0x4f, 0x0f, 0xcf, 0x5b, 0xd1, 0xfe, 0xf5, 0x71, 0x9d, 0xe5, 0xc7, 0x52, 0x05, 0x48,
	0xa4, 0x1f, 0x1b, 0x60, 0xe5, 0xec, 0x2d, 0x34, 0x7b, 0xac, 0xd4, 0xee, 0x36, 0x57, 0x25, 0x3c,
	0x19, 0x74, 0x94, 0xa3, 0x62, 0xfa, 0x27, 0xca, 0x1d, 0xbb, 0xf0, 0xa8, 0x4f, 0xb1, 0xeb, 0x22,
	0xd1, 0x28, 0x9e, 0x4e, 0x93, 0x7e, 0x56, 0x7c, 0xb8, 0x0e, 0x77, 0x6a, 0x57, 0x3c, 0xea, 0xec,
	0xa9, 0xcd, 0x4c, 0x6c, 0xf4, 0x58, 0x39, 0x3b, 0x5b, 0xd9, 0x40, 0x8f, 0x8e, 0xc2, 0x37, 0x27,
	0xe1, 0x4d, 0x76, 0x7e, 0xc0, 0x73, 0xbe, 0xab, 0xb0, 0x5b, 0xa4, 0x7a, 0xb9, 0xc7, 0x1d, 0x26,
	0x34, 0x35, 0x0b, 0x6e, 0xe0, 0xfe, 0x7a, 0xf6, 0x2d, 0xb4, 0x59, 0x2a, 0x60, 0x9e, 0x0a, 0x58,
	0xa4, 0x02, 0xbe, 0x52, 0x01, 0x9f, 0x4b, 0xa1, 0xcd, 0x97, 0x42, 0x5b, 0x2c, 0x85, 0xf6, 0x52,
	0xc4, 0x77, 0x74, 0xbd, 0x60, 0x14, 0x39, 0xe1, 0xe0, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x0b,
	0x41, 0xe1, 0xbb, 0x01, 0x00, 0x00,
}
