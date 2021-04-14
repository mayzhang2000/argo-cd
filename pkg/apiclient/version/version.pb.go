// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/version/version.proto

// Version Service
//
// Version Service API returns the version of the API server.

package version

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// VersionMessage represents version of the Argo CD API server
type VersionMessage struct {
	Version              string   `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	BuildDate            string   `protobuf:"bytes,2,opt,name=BuildDate,proto3" json:"BuildDate,omitempty"`
	GitCommit            string   `protobuf:"bytes,3,opt,name=GitCommit,proto3" json:"GitCommit,omitempty"`
	GitTag               string   `protobuf:"bytes,4,opt,name=GitTag,proto3" json:"GitTag,omitempty"`
	GitTreeState         string   `protobuf:"bytes,5,opt,name=GitTreeState,proto3" json:"GitTreeState,omitempty"`
	GoVersion            string   `protobuf:"bytes,6,opt,name=GoVersion,proto3" json:"GoVersion,omitempty"`
	Compiler             string   `protobuf:"bytes,7,opt,name=Compiler,proto3" json:"Compiler,omitempty"`
	Platform             string   `protobuf:"bytes,8,opt,name=Platform,proto3" json:"Platform,omitempty"`
	KsonnetVersion       string   `protobuf:"bytes,9,opt,name=KsonnetVersion,proto3" json:"KsonnetVersion,omitempty"`
	KustomizeVersion     string   `protobuf:"bytes,10,opt,name=KustomizeVersion,proto3" json:"KustomizeVersion,omitempty"`
	HelmVersion          string   `protobuf:"bytes,11,opt,name=HelmVersion,proto3" json:"HelmVersion,omitempty"`
	KubectlVersion       string   `protobuf:"bytes,12,opt,name=KubectlVersion,proto3" json:"KubectlVersion,omitempty"`
	JsonnetVersion       string   `protobuf:"bytes,13,opt,name=JsonnetVersion,proto3" json:"JsonnetVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionMessage) Reset()         { *m = VersionMessage{} }
func (m *VersionMessage) String() string { return proto.CompactTextString(m) }
func (*VersionMessage) ProtoMessage()    {}
func (*VersionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8be80977d07a4107, []int{0}
}
func (m *VersionMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VersionMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VersionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionMessage.Merge(m, src)
}
func (m *VersionMessage) XXX_Size() int {
	return m.Size()
}
func (m *VersionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_VersionMessage proto.InternalMessageInfo

func (m *VersionMessage) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionMessage) GetBuildDate() string {
	if m != nil {
		return m.BuildDate
	}
	return ""
}

func (m *VersionMessage) GetGitCommit() string {
	if m != nil {
		return m.GitCommit
	}
	return ""
}

func (m *VersionMessage) GetGitTag() string {
	if m != nil {
		return m.GitTag
	}
	return ""
}

func (m *VersionMessage) GetGitTreeState() string {
	if m != nil {
		return m.GitTreeState
	}
	return ""
}

func (m *VersionMessage) GetGoVersion() string {
	if m != nil {
		return m.GoVersion
	}
	return ""
}

func (m *VersionMessage) GetCompiler() string {
	if m != nil {
		return m.Compiler
	}
	return ""
}

func (m *VersionMessage) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *VersionMessage) GetKsonnetVersion() string {
	if m != nil {
		return m.KsonnetVersion
	}
	return ""
}

func (m *VersionMessage) GetKustomizeVersion() string {
	if m != nil {
		return m.KustomizeVersion
	}
	return ""
}

func (m *VersionMessage) GetHelmVersion() string {
	if m != nil {
		return m.HelmVersion
	}
	return ""
}

func (m *VersionMessage) GetKubectlVersion() string {
	if m != nil {
		return m.KubectlVersion
	}
	return ""
}

func (m *VersionMessage) GetJsonnetVersion() string {
	if m != nil {
		return m.JsonnetVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionMessage)(nil), "version.VersionMessage")
}

func init() { proto.RegisterFile("server/version/version.proto", fileDescriptor_8be80977d07a4107) }

var fileDescriptor_8be80977d07a4107 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcd, 0x6a, 0xdb, 0x40,
	0x14, 0x85, 0x91, 0xdd, 0xfa, 0x67, 0xec, 0x9a, 0x32, 0x14, 0x57, 0xa8, 0xc6, 0x18, 0x2f, 0x4a,
	0x29, 0x54, 0x02, 0xb7, 0x4f, 0x60, 0xb7, 0xb8, 0xd4, 0x14, 0x4c, 0x5d, 0xba, 0xe8, 0x6e, 0x24,
	0x5f, 0xab, 0xd3, 0x68, 0x74, 0xc5, 0x68, 0x24, 0x48, 0x96, 0x79, 0x85, 0xbc, 0x4f, 0xd6, 0x59,
	0x06, 0xf2, 0x02, 0xc1, 0xe4, 0x41, 0x82, 0x46, 0x1a, 0x25, 0x4a, 0x56, 0x9a, 0x7b, 0xce, 0xa7,
	0x33, 0x17, 0xce, 0x90, 0x49, 0x0a, 0x32, 0x07, 0xe9, 0xe5, 0x20, 0x53, 0x8e, 0xb1, 0xf9, 0xba,
	0x89, 0x44, 0x85, 0xb4, 0x5b, 0x8d, 0xce, 0x24, 0x44, 0x0c, 0x23, 0xf0, 0x58, 0xc2, 0x3d, 0x16,
	0xc7, 0xa8, 0x98, 0xe2, 0x18, 0xa7, 0x25, 0xe6, 0xbc, 0xab, 0x5c, 0x3d, 0xf9, 0xd9, 0xc1, 0x03,
	0x91, 0xa8, 0xd3, 0xd2, 0x9c, 0x5f, 0xb6, 0xc9, 0xe8, 0x4f, 0x19, 0xf3, 0x13, 0xd2, 0x94, 0x85,
	0x40, 0x6d, 0xd2, 0xad, 0x14, 0xdb, 0x9a, 0x59, 0x1f, 0xfa, 0xbf, 0xcc, 0x48, 0x27, 0xa4, 0xbf,
	0xcc, 0x78, 0xb4, 0xff, 0xca, 0x14, 0xd8, 0x2d, 0xed, 0x3d, 0x08, 0x85, 0xbb, 0xe6, 0x6a, 0x85,
	0x42, 0x70, 0x65, 0xb7, 0x4b, 0xb7, 0x16, 0xe8, 0x98, 0x74, 0xd6, 0x5c, 0xfd, 0x66, 0xa1, 0xfd,
	0x42, 0x5b, 0xd5, 0x44, 0xe7, 0x64, 0x58, 0x9c, 0x24, 0xc0, 0x4e, 0x15, 0xb1, 0x2f, 0xb5, 0xdb,
	0xd0, 0x74, 0x32, 0x9a, 0x9d, 0x3a, 0x55, 0xb2, 0x11, 0xa8, 0x43, 0x7a, 0x2b, 0x14, 0x09, 0x8f,
	0x40, 0xda, 0x5d, 0x6d, 0xd6, 0x73, 0xe1, 0x6d, 0x23, 0xa6, 0x0e, 0x28, 0x85, 0xdd, 0x2b, 0x3d,
	0x33, 0xd3, 0xf7, 0x64, 0xb4, 0x49, 0x31, 0x8e, 0x41, 0x99, 0xe8, 0xbe, 0x26, 0x9e, 0xa8, 0xf4,
	0x23, 0x79, 0xbd, 0xc9, 0x52, 0x85, 0x82, 0x9f, 0x81, 0x21, 0x89, 0x26, 0x9f, 0xe9, 0x74, 0x46,
	0x06, 0xdf, 0x21, 0x12, 0x06, 0x1b, 0x68, 0xec, 0xb1, 0xa4, 0x6f, 0xcd, 0x7c, 0x08, 0x54, 0x64,
	0xa0, 0x61, 0x75, 0x6b, 0x43, 0x2d, 0xb8, 0x1f, 0xcd, 0xed, 0x5e, 0x95, 0x5c, 0x53, 0x5d, 0xf8,
	0x75, 0x7f, 0x3b, 0x90, 0x39, 0x0f, 0x80, 0x6e, 0xeb, 0xfe, 0xe8, 0xd8, 0x2d, 0xbb, 0x77, 0x4d,
	0xf7, 0xee, 0xb7, 0xa2, 0x7b, 0xe7, 0xad, 0x6b, 0x5e, 0x52, 0xb3, 0xfb, 0xf9, 0x9b, 0xf3, 0x9b,
	0xbb, 0x8b, 0xd6, 0x88, 0x0e, 0xf5, 0x5b, 0xaa, 0xa0, 0xe5, 0xf2, 0xea, 0x38, 0xb5, 0xae, 0x8f,
	0x53, 0xeb, 0xf6, 0x38, 0xb5, 0xfe, 0x7e, 0x09, 0xb9, 0xfa, 0x97, 0xf9, 0x6e, 0x80, 0xc2, 0x63,
	0x32, 0xc4, 0x44, 0xe2, 0x7f, 0x7d, 0xf8, 0x14, 0xec, 0xbd, 0x7c, 0xe1, 0x25, 0x27, 0x61, 0xf1,
	0x77, 0x10, 0x71, 0x88, 0x95, 0xc9, 0xf0, 0x3b, 0x7a, 0x85, 0xcf, 0xf7, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x48, 0x07, 0x04, 0xf4, 0xd3, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VersionServiceClient is the client API for VersionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VersionServiceClient interface {
	// Version returns version information of the API server
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionMessage, error)
}

type versionServiceClient struct {
	cc *grpc.ClientConn
}

func NewVersionServiceClient(cc *grpc.ClientConn) VersionServiceClient {
	return &versionServiceClient{cc}
}

func (c *versionServiceClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionMessage, error) {
	out := new(VersionMessage)
	err := c.cc.Invoke(ctx, "/version.VersionService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionServiceServer is the server API for VersionService service.
type VersionServiceServer interface {
	// Version returns version information of the API server
	Version(context.Context, *emptypb.Empty) (*VersionMessage, error)
}

// UnimplementedVersionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVersionServiceServer struct {
}

func (*UnimplementedVersionServiceServer) Version(ctx context.Context, req *emptypb.Empty) (*VersionMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}

func RegisterVersionServiceServer(s *grpc.Server, srv VersionServiceServer) {
	s.RegisterService(&_VersionService_serviceDesc, srv)
}

func _VersionService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.VersionService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _VersionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "version.VersionService",
	HandlerType: (*VersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _VersionService_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/version/version.proto",
}

func (m *VersionMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VersionMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.JsonnetVersion) > 0 {
		i -= len(m.JsonnetVersion)
		copy(dAtA[i:], m.JsonnetVersion)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.JsonnetVersion)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.KubectlVersion) > 0 {
		i -= len(m.KubectlVersion)
		copy(dAtA[i:], m.KubectlVersion)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.KubectlVersion)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.HelmVersion) > 0 {
		i -= len(m.HelmVersion)
		copy(dAtA[i:], m.HelmVersion)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.HelmVersion)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.KustomizeVersion) > 0 {
		i -= len(m.KustomizeVersion)
		copy(dAtA[i:], m.KustomizeVersion)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.KustomizeVersion)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.KsonnetVersion) > 0 {
		i -= len(m.KsonnetVersion)
		copy(dAtA[i:], m.KsonnetVersion)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.KsonnetVersion)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Platform) > 0 {
		i -= len(m.Platform)
		copy(dAtA[i:], m.Platform)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.Platform)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Compiler) > 0 {
		i -= len(m.Compiler)
		copy(dAtA[i:], m.Compiler)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.Compiler)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.GoVersion) > 0 {
		i -= len(m.GoVersion)
		copy(dAtA[i:], m.GoVersion)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.GoVersion)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.GitTreeState) > 0 {
		i -= len(m.GitTreeState)
		copy(dAtA[i:], m.GitTreeState)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.GitTreeState)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.GitTag) > 0 {
		i -= len(m.GitTag)
		copy(dAtA[i:], m.GitTag)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.GitTag)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.GitCommit) > 0 {
		i -= len(m.GitCommit)
		copy(dAtA[i:], m.GitCommit)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.GitCommit)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BuildDate) > 0 {
		i -= len(m.BuildDate)
		copy(dAtA[i:], m.BuildDate)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.BuildDate)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVersion(dAtA []byte, offset int, v uint64) int {
	offset -= sovVersion(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VersionMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.BuildDate)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.GitCommit)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.GitTag)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.GitTreeState)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.GoVersion)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.Compiler)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.Platform)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.KsonnetVersion)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.KustomizeVersion)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.HelmVersion)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.KubectlVersion)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	l = len(m.JsonnetVersion)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVersion(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVersion(x uint64) (n int) {
	return sovVersion(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VersionMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVersion
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
			return fmt.Errorf("proto: VersionMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuildDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuildDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitCommit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitCommit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitTag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitTag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitTreeState", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitTreeState = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compiler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Compiler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platform", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Platform = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KsonnetVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KsonnetVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KustomizeVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KustomizeVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HelmVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HelmVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KubectlVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KubectlVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JsonnetVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JsonnetVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVersion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVersion
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVersion
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
func skipVersion(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVersion
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
					return 0, ErrIntOverflowVersion
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
					return 0, ErrIntOverflowVersion
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
				return 0, ErrInvalidLengthVersion
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVersion
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVersion
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVersion        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVersion          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVersion = fmt.Errorf("proto: unexpected end of group")
)
