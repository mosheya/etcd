// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lease.proto

package leasepb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	etcdserverpb "github.com/mosheya/etcd/etcdserver/etcdserverpb"
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

type Lease struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TTL                  int64    `protobuf:"varint,2,opt,name=TTL,proto3" json:"TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lease) Reset()         { *m = Lease{} }
func (m *Lease) String() string { return proto.CompactTextString(m) }
func (*Lease) ProtoMessage()    {}
func (*Lease) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dd57e402472b33a, []int{0}
}
func (m *Lease) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lease.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lease.Merge(m, src)
}
func (m *Lease) XXX_Size() int {
	return m.Size()
}
func (m *Lease) XXX_DiscardUnknown() {
	xxx_messageInfo_Lease.DiscardUnknown(m)
}

var xxx_messageInfo_Lease proto.InternalMessageInfo

type LeaseInternalRequest struct {
	LeaseTimeToLiveRequest *etcdserverpb.LeaseTimeToLiveRequest `protobuf:"bytes,1,opt,name=LeaseTimeToLiveRequest,proto3" json:"LeaseTimeToLiveRequest,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                             `json:"-"`
	XXX_unrecognized       []byte                               `json:"-"`
	XXX_sizecache          int32                                `json:"-"`
}

func (m *LeaseInternalRequest) Reset()         { *m = LeaseInternalRequest{} }
func (m *LeaseInternalRequest) String() string { return proto.CompactTextString(m) }
func (*LeaseInternalRequest) ProtoMessage()    {}
func (*LeaseInternalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dd57e402472b33a, []int{1}
}
func (m *LeaseInternalRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseInternalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseInternalRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaseInternalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseInternalRequest.Merge(m, src)
}
func (m *LeaseInternalRequest) XXX_Size() int {
	return m.Size()
}
func (m *LeaseInternalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseInternalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseInternalRequest proto.InternalMessageInfo

type LeaseInternalResponse struct {
	LeaseTimeToLiveResponse *etcdserverpb.LeaseTimeToLiveResponse `protobuf:"bytes,1,opt,name=LeaseTimeToLiveResponse,proto3" json:"LeaseTimeToLiveResponse,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                              `json:"-"`
	XXX_unrecognized        []byte                                `json:"-"`
	XXX_sizecache           int32                                 `json:"-"`
}

func (m *LeaseInternalResponse) Reset()         { *m = LeaseInternalResponse{} }
func (m *LeaseInternalResponse) String() string { return proto.CompactTextString(m) }
func (*LeaseInternalResponse) ProtoMessage()    {}
func (*LeaseInternalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dd57e402472b33a, []int{2}
}
func (m *LeaseInternalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseInternalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseInternalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaseInternalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseInternalResponse.Merge(m, src)
}
func (m *LeaseInternalResponse) XXX_Size() int {
	return m.Size()
}
func (m *LeaseInternalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseInternalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseInternalResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Lease)(nil), "leasepb.Lease")
	proto.RegisterType((*LeaseInternalRequest)(nil), "leasepb.LeaseInternalRequest")
	proto.RegisterType((*LeaseInternalResponse)(nil), "leasepb.LeaseInternalResponse")
}

func init() { proto.RegisterFile("lease.proto", fileDescriptor_3dd57e402472b33a) }

var fileDescriptor_3dd57e402472b33a = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x49, 0x4d, 0x2c,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0x73, 0x0a, 0x92, 0xa4, 0x44, 0xd2,
	0xf3, 0xd3, 0xf3, 0xc1, 0x62, 0xfa, 0x20, 0x16, 0x44, 0x5a, 0x4a, 0x2d, 0xb5, 0x24, 0x39, 0x45,
	0x1f, 0x44, 0x14, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0x21, 0x31, 0x0b, 0x92, 0xf4, 0x8b, 0x0a, 0x92,
	0x21, 0xea, 0x94, 0x34, 0xb9, 0x58, 0x7d, 0x40, 0x06, 0x09, 0xf1, 0x71, 0x31, 0x79, 0xba, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x31, 0x79, 0xba, 0x08, 0x09, 0x70, 0x31, 0x87, 0x84, 0xf8,
	0x48, 0x30, 0x81, 0x05, 0x40, 0x4c, 0xa5, 0x12, 0x2e, 0x11, 0xb0, 0x52, 0xcf, 0xbc, 0x92, 0xd4,
	0xa2, 0xbc, 0xc4, 0x9c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0xa1, 0x18, 0x2e, 0x31, 0xb0,
	0x78, 0x48, 0x66, 0x6e, 0x6a, 0x48, 0xbe, 0x4f, 0x66, 0x59, 0x2a, 0x54, 0x06, 0x6c, 0x1a, 0xb7,
	0x91, 0x8a, 0x1e, 0xb2, 0xdd, 0x7a, 0xd8, 0xd5, 0x06, 0xe1, 0x30, 0x43, 0xa9, 0x82, 0x4b, 0x14,
	0xcd, 0xd6, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa1, 0x78, 0x2e, 0x71, 0x0c, 0x2d, 0x10, 0x29,
	0xa8, 0xbd, 0xaa, 0x04, 0xec, 0x85, 0x28, 0x0e, 0xc2, 0x65, 0x8a, 0x93, 0xc4, 0x89, 0x87, 0x72,
	0x0c, 0x17, 0x1e, 0xca, 0x31, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0xc3, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0xf2, 0x42, 0xe0, 0x91, 0x01, 0x00, 0x00,
}

func (m *Lease) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lease) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lease) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TTL != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.TTL))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LeaseInternalRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseInternalRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseInternalRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LeaseTimeToLiveRequest != nil {
		{
			size, err := m.LeaseTimeToLiveRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLease(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LeaseInternalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseInternalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseInternalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LeaseTimeToLiveResponse != nil {
		{
			size, err := m.LeaseTimeToLiveResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLease(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLease(dAtA []byte, offset int, v uint64) int {
	offset -= sovLease(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lease) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLease(uint64(m.ID))
	}
	if m.TTL != 0 {
		n += 1 + sovLease(uint64(m.TTL))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseInternalRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LeaseTimeToLiveRequest != nil {
		l = m.LeaseTimeToLiveRequest.Size()
		n += 1 + l + sovLease(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseInternalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LeaseTimeToLiveResponse != nil {
		l = m.LeaseTimeToLiveResponse.Size()
		n += 1 + l + sovLease(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLease(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLease(x uint64) (n int) {
	return sovLease(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lease) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
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
			return fmt.Errorf("proto: Lease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLease
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLease
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
func (m *LeaseInternalRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
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
			return fmt.Errorf("proto: LeaseInternalRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseInternalRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseTimeToLiveRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
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
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeaseTimeToLiveRequest == nil {
				m.LeaseTimeToLiveRequest = &etcdserverpb.LeaseTimeToLiveRequest{}
			}
			if err := m.LeaseTimeToLiveRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLease
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLease
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
func (m *LeaseInternalResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
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
			return fmt.Errorf("proto: LeaseInternalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseInternalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseTimeToLiveResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
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
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeaseTimeToLiveResponse == nil {
				m.LeaseTimeToLiveResponse = &etcdserverpb.LeaseTimeToLiveResponse{}
			}
			if err := m.LeaseTimeToLiveResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLease
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLease
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
func skipLease(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLease
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
					return 0, ErrIntOverflowLease
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
					return 0, ErrIntOverflowLease
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
				return 0, ErrInvalidLengthLease
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLease
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLease
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
				next, err := skipLease(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLease
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
	ErrInvalidLengthLease = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLease   = fmt.Errorf("proto: integer overflow")
)
