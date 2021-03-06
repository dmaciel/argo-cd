// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/version/version.proto

package version // import "github.com/argoproj/argo-cd/pkg/apiclient/version"

/*
	Version Service

	Version Service API returns the version of the API server.
*/

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionMessage) Reset()         { *m = VersionMessage{} }
func (m *VersionMessage) String() string { return proto.CompactTextString(m) }
func (*VersionMessage) ProtoMessage()    {}
func (*VersionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_version_ca42711144b8dc48, []int{0}
}
func (m *VersionMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VersionMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *VersionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionMessage.Merge(dst, src)
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

func init() {
	proto.RegisterType((*VersionMessage)(nil), "version.VersionMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VersionService service

type VersionServiceClient interface {
	// Version returns version information of the API server
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionMessage, error)
}

type versionServiceClient struct {
	cc *grpc.ClientConn
}

func NewVersionServiceClient(cc *grpc.ClientConn) VersionServiceClient {
	return &versionServiceClient{cc}
}

func (c *versionServiceClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionMessage, error) {
	out := new(VersionMessage)
	err := c.cc.Invoke(ctx, "/version.VersionService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VersionService service

type VersionServiceServer interface {
	// Version returns version information of the API server
	Version(context.Context, *empty.Empty) (*VersionMessage, error)
}

func RegisterVersionServiceServer(s *grpc.Server, srv VersionServiceServer) {
	s.RegisterService(&_VersionService_serviceDesc, srv)
}

func _VersionService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
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
		return srv.(VersionServiceServer).Version(ctx, req.(*empty.Empty))
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
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	if len(m.BuildDate) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.BuildDate)))
		i += copy(dAtA[i:], m.BuildDate)
	}
	if len(m.GitCommit) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.GitCommit)))
		i += copy(dAtA[i:], m.GitCommit)
	}
	if len(m.GitTag) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.GitTag)))
		i += copy(dAtA[i:], m.GitTag)
	}
	if len(m.GitTreeState) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.GitTreeState)))
		i += copy(dAtA[i:], m.GitTreeState)
	}
	if len(m.GoVersion) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.GoVersion)))
		i += copy(dAtA[i:], m.GoVersion)
	}
	if len(m.Compiler) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.Compiler)))
		i += copy(dAtA[i:], m.Compiler)
	}
	if len(m.Platform) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.Platform)))
		i += copy(dAtA[i:], m.Platform)
	}
	if len(m.KsonnetVersion) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.KsonnetVersion)))
		i += copy(dAtA[i:], m.KsonnetVersion)
	}
	if len(m.KustomizeVersion) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.KustomizeVersion)))
		i += copy(dAtA[i:], m.KustomizeVersion)
	}
	if len(m.HelmVersion) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.HelmVersion)))
		i += copy(dAtA[i:], m.HelmVersion)
	}
	if len(m.KubectlVersion) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintVersion(dAtA, i, uint64(len(m.KubectlVersion)))
		i += copy(dAtA[i:], m.KubectlVersion)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintVersion(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VersionMessage) Size() (n int) {
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
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVersion(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
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
			wire |= (uint64(b) & 0x7F) << shift
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KubectlVersion = string(dAtA[iNdEx:postIndex])
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
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthVersion
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowVersion
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
				next, err := skipVersion(dAtA[start:])
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
	ErrInvalidLengthVersion = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVersion   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("server/version/version.proto", fileDescriptor_version_ca42711144b8dc48)
}

var fileDescriptor_version_ca42711144b8dc48 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x4a, 0xe3, 0x40,
	0x18, 0xc5, 0x49, 0xbb, 0xdb, 0x3f, 0xd3, 0x52, 0x96, 0x61, 0xe9, 0x86, 0x6c, 0x29, 0xa5, 0x17,
	0xcb, 0xb2, 0xb0, 0x09, 0xbb, 0x3e, 0x80, 0xd0, 0x2a, 0x15, 0x8a, 0x50, 0xac, 0x78, 0xe1, 0xdd,
	0x24, 0xfd, 0x1a, 0x47, 0x33, 0xf9, 0xc2, 0x64, 0x52, 0xd0, 0x4b, 0x5f, 0x41, 0xf0, 0x99, 0xbc,
	0x14, 0x7c, 0x01, 0x29, 0x3e, 0x88, 0x64, 0x92, 0x89, 0x56, 0xaf, 0x3a, 0xe7, 0xfc, 0x4e, 0xcf,
	0x04, 0xce, 0x90, 0x41, 0x0a, 0x72, 0x03, 0xd2, 0xdb, 0x80, 0x4c, 0x39, 0xc6, 0xe6, 0xd7, 0x4d,
	0x24, 0x2a, 0xa4, 0xcd, 0x52, 0x3a, 0x83, 0x10, 0x31, 0x8c, 0xc0, 0x63, 0x09, 0xf7, 0x58, 0x1c,
	0xa3, 0x62, 0x8a, 0x63, 0x9c, 0x16, 0x31, 0xe7, 0x67, 0x49, 0xb5, 0xf2, 0xb3, 0xb5, 0x07, 0x22,
	0x51, 0xd7, 0x05, 0x1c, 0xdf, 0xd7, 0x49, 0xef, 0xac, 0xa8, 0x39, 0x86, 0x34, 0x65, 0x21, 0x50,
	0x9b, 0x34, 0x4b, 0xc7, 0xb6, 0x46, 0xd6, 0xef, 0xf6, 0x89, 0x91, 0x74, 0x40, 0xda, 0x93, 0x8c,
	0x47, 0xab, 0x03, 0xa6, 0xc0, 0xae, 0x69, 0xf6, 0x66, 0xe4, 0x74, 0xc6, 0xd5, 0x14, 0x85, 0xe0,
	0xca, 0xae, 0x17, 0xb4, 0x32, 0x68, 0x9f, 0x34, 0x66, 0x5c, 0x9d, 0xb2, 0xd0, 0xfe, 0xa2, 0x51,
	0xa9, 0xe8, 0x98, 0x74, 0xf3, 0x93, 0x04, 0x58, 0xaa, 0xbc, 0xf6, 0xab, 0xa6, 0x3b, 0x9e, 0x6e,
	0x46, 0xf3, 0x4d, 0x8d, 0xb2, 0xd9, 0x18, 0xd4, 0x21, 0xad, 0x29, 0x8a, 0x84, 0x47, 0x20, 0xed,
	0xa6, 0x86, 0x95, 0xce, 0xd9, 0x22, 0x62, 0x6a, 0x8d, 0x52, 0xd8, 0xad, 0x82, 0x19, 0x4d, 0x7f,
	0x91, 0xde, 0x3c, 0xc5, 0x38, 0x06, 0x65, 0xaa, 0xdb, 0x3a, 0xf1, 0xc1, 0xa5, 0x7f, 0xc8, 0xb7,
	0x79, 0x96, 0x2a, 0x14, 0xfc, 0x06, 0x4c, 0x92, 0xe8, 0xe4, 0x27, 0x9f, 0x8e, 0x48, 0xe7, 0x08,
	0x22, 0x61, 0x62, 0x1d, 0x1d, 0x7b, 0x6f, 0xe9, 0x5b, 0x33, 0x1f, 0x02, 0x15, 0x99, 0x50, 0xb7,
	0xbc, 0x75, 0xc7, 0xfd, 0xef, 0x57, 0xbb, 0x2c, 0x41, 0x6e, 0x78, 0x00, 0x74, 0x51, 0xed, 0x42,
	0xfb, 0x6e, 0xb1, 0xa9, 0x6b, 0x36, 0x75, 0x0f, 0xf3, 0x4d, 0x9d, 0x1f, 0xae, 0x79, 0x21, 0xbb,
	0x9b, 0x8e, 0xbf, 0xdf, 0x3e, 0xbd, 0xdc, 0xd5, 0x7a, 0xb4, 0xab, 0xdf, 0x48, 0x19, 0x9a, 0xec,
	0x3f, 0x6c, 0x87, 0xd6, 0xe3, 0x76, 0x68, 0x3d, 0x6f, 0x87, 0xd6, 0xf9, 0xbf, 0x90, 0xab, 0x8b,
	0xcc, 0x77, 0x03, 0x14, 0x1e, 0x93, 0x21, 0x26, 0x12, 0x2f, 0xf5, 0xe1, 0x6f, 0xb0, 0xf2, 0x92,
	0xab, 0x30, 0xff, 0x6b, 0x10, 0x71, 0x88, 0x95, 0x29, 0xf0, 0x1b, 0xfa, 0xfe, 0xbd, 0xd7, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0xc2, 0x44, 0x80, 0xa8, 0x02, 0x00, 0x00,
}
