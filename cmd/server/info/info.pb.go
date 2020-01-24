// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cmd/server/info/info.proto

package info

import (
	context "context"
	fmt "fmt"
	_ "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "k8s.io/api/core/v1"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
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

type GetInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b893c3550e6e6df, []int{0}
}
func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(m, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

type InfoResponse struct {
	ManagedNamespace     string   `protobuf:"bytes,1,opt,name=managedNamespace,proto3" json:"managedNamespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b893c3550e6e6df, []int{1}
}
func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

func (m *InfoResponse) GetManagedNamespace() string {
	if m != nil {
		return m.ManagedNamespace
	}
	return ""
}

func init() {
	proto.RegisterType((*GetInfoRequest)(nil), "info.GetInfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "info.InfoResponse")
}

func init() { proto.RegisterFile("cmd/server/info/info.proto", fileDescriptor_8b893c3550e6e6df) }

var fileDescriptor_8b893c3550e6e6df = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0x3b, 0x31,
	0x10, 0xc6, 0xd9, 0x3f, 0x7f, 0x14, 0xd7, 0x52, 0xca, 0xd2, 0x83, 0x2c, 0x52, 0x64, 0x4f, 0xc5,
	0xc3, 0x0e, 0x55, 0x11, 0xf1, 0xa8, 0x07, 0xe9, 0xc5, 0x43, 0xbd, 0x88, 0xb7, 0xe9, 0x76, 0x9a,
	0xc6, 0x36, 0x99, 0x98, 0xa4, 0x5b, 0xbc, 0xfa, 0x0a, 0xbe, 0x94, 0x47, 0xc1, 0x17, 0x90, 0xe2,
	0x83, 0xc8, 0xa6, 0x0b, 0x55, 0xab, 0x97, 0x30, 0xf9, 0x3e, 0xbe, 0xf9, 0x25, 0x33, 0x71, 0x5a,
	0xa8, 0x11, 0x38, 0xb2, 0x25, 0x59, 0x90, 0x7a, 0xcc, 0xe1, 0xc8, 0x8d, 0x65, 0xcf, 0xc9, 0xff,
	0xaa, 0x4e, 0xdb, 0x82, 0x05, 0x07, 0x01, 0xaa, 0x6a, 0xe5, 0xa5, 0xfb, 0x82, 0x59, 0xcc, 0x08,
	0xd0, 0x48, 0x40, 0xad, 0xd9, 0xa3, 0x97, 0xac, 0x5d, 0xed, 0x9e, 0x4c, 0xcf, 0x5c, 0x2e, 0xb9,
	0x72, 0x15, 0x16, 0x13, 0xa9, 0xc9, 0x3e, 0x82, 0x99, 0x8a, 0x4a, 0x70, 0xa0, 0xc8, 0x23, 0x94,
	0x3d, 0x10, 0xa4, 0xc9, 0xa2, 0xa7, 0x51, 0x9d, 0xba, 0x14, 0xd2, 0x4f, 0xe6, 0xc3, 0xbc, 0x60,
	0x05, 0x68, 0x03, 0xf4, 0x3e, 0x14, 0xeb, 0xe8, 0x82, 0xed, 0x74, 0x3c, 0xe3, 0x05, 0x94, 0x3d,
	0x9c, 0x99, 0x09, 0x6e, 0x36, 0xc9, 0xd6, 0x68, 0x28, 0xd8, 0xd2, 0x2f, 0xa0, 0xac, 0x15, 0x37,
	0xaf, 0xc8, 0xf7, 0xf5, 0x98, 0x07, 0xf4, 0x30, 0x27, 0xe7, 0xb3, 0xf3, 0xb8, 0xb1, 0xba, 0x3a,
	0xc3, 0xda, 0x51, 0x72, 0x18, 0xb7, 0x14, 0x6a, 0x14, 0x34, 0xba, 0x46, 0x45, 0xce, 0x60, 0x41,
	0x7b, 0xd1, 0x41, 0xd4, 0xdd, 0x19, 0x6c, 0xe8, 0x47, 0xb7, 0xf1, 0x6e, 0x95, 0xbd, 0x21, 0x5b,
	0xca, 0x82, 0x92, 0x7e, 0xbc, 0x5d, 0x37, 0x4f, 0xda, 0x79, 0x98, 0xe6, 0x77, 0x56, 0x9a, 0xac,
	0xd4, 0xaf, 0xbc, 0xac, 0xfd, 0xf4, 0xf6, 0xf1, 0xfc, 0xaf, 0x99, 0x34, 0xc2, 0xbb, 0xcb, 0x5e,
	0x58, 0xc3, 0xc5, 0xe9, 0xcb, 0xb2, 0x13, 0xbd, 0x2e, 0x3b, 0xd1, 0xfb, 0xb2, 0x13, 0xdd, 0x75,
	0xff, 0x1c, 0xcf, 0x8f, 0x1d, 0x0e, 0xb7, 0xc2, 0x37, 0x8f, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x05, 0xce, 0x15, 0x70, 0xdd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InfoServiceClient is the client API for InfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfoServiceClient interface {
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type infoServiceClient struct {
	cc *grpc.ClientConn
}

func NewInfoServiceClient(cc *grpc.ClientConn) InfoServiceClient {
	return &infoServiceClient{cc}
}

func (c *infoServiceClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/info.InfoService/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServiceServer is the server API for InfoService service.
type InfoServiceServer interface {
	GetInfo(context.Context, *GetInfoRequest) (*InfoResponse, error)
}

// UnimplementedInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInfoServiceServer struct {
}

func (*UnimplementedInfoServiceServer) GetInfo(ctx context.Context, req *GetInfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}

func RegisterInfoServiceServer(s *grpc.Server, srv InfoServiceServer) {
	s.RegisterService(&_InfoService_serviceDesc, srv)
}

func _InfoService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.InfoService/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "info.InfoService",
	HandlerType: (*InfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _InfoService_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/server/info/info.proto",
}

func (m *GetInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *InfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ManagedNamespace) > 0 {
		i -= len(m.ManagedNamespace)
		copy(dAtA[i:], m.ManagedNamespace)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.ManagedNamespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ManagedNamespace)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInfo(x uint64) (n int) {
	return sovInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: GetInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInfo
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
func (m *InfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: InfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManagedNamespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ManagedNamespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInfo
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
func skipInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
				return 0, ErrInvalidLengthInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInfo = fmt.Errorf("proto: unexpected end of group")
)