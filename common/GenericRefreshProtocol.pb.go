// Code generated by protoc-gen-go.
// source: gopkg.in/freddierice/go-hproto.v1/common/GenericRefreshProtocol.proto
// DO NOT EDIT!

/*
Package hproto_common is a generated protocol buffer package.

It is generated from these files:
	gopkg.in/freddierice/go-hproto.v1/common/GenericRefreshProtocol.proto
	gopkg.in/freddierice/go-hproto.v1/common/GetUserMappingsProtocol.proto
	gopkg.in/freddierice/go-hproto.v1/common/HAServiceProtocol.proto
	gopkg.in/freddierice/go-hproto.v1/common/IpcConnectionContext.proto
	gopkg.in/freddierice/go-hproto.v1/common/ProtobufRpcEngine.proto
	gopkg.in/freddierice/go-hproto.v1/common/ProtocolInfo.proto
	gopkg.in/freddierice/go-hproto.v1/common/RefreshAuthorizationPolicyProtocol.proto
	gopkg.in/freddierice/go-hproto.v1/common/RefreshCallQueueProtocol.proto
	gopkg.in/freddierice/go-hproto.v1/common/RefreshUserMappingsProtocol.proto
	gopkg.in/freddierice/go-hproto.v1/common/RpcHeader.proto
	gopkg.in/freddierice/go-hproto.v1/common/Security.proto
	gopkg.in/freddierice/go-hproto.v1/common/TraceAdmin.proto
	gopkg.in/freddierice/go-hproto.v1/common/ZKFCProtocol.proto

It has these top-level messages:
	GenericRefreshRequestProto
	GenericRefreshResponseProto
	GenericRefreshResponseCollectionProto
	GetGroupsForUserRequestProto
	GetGroupsForUserResponseProto
	HAStateChangeRequestInfoProto
	MonitorHealthRequestProto
	MonitorHealthResponseProto
	TransitionToActiveRequestProto
	TransitionToActiveResponseProto
	TransitionToStandbyRequestProto
	TransitionToStandbyResponseProto
	GetServiceStatusRequestProto
	GetServiceStatusResponseProto
	UserInformationProto
	IpcConnectionContextProto
	RequestHeaderProto
	GetProtocolVersionsRequestProto
	ProtocolVersionProto
	GetProtocolVersionsResponseProto
	GetProtocolSignatureRequestProto
	GetProtocolSignatureResponseProto
	ProtocolSignatureProto
	RefreshServiceAclRequestProto
	RefreshServiceAclResponseProto
	RefreshCallQueueRequestProto
	RefreshCallQueueResponseProto
	RefreshUserToGroupsMappingsRequestProto
	RefreshUserToGroupsMappingsResponseProto
	RefreshSuperUserGroupsConfigurationRequestProto
	RefreshSuperUserGroupsConfigurationResponseProto
	RPCTraceInfoProto
	RPCCallerContextProto
	RpcRequestHeaderProto
	RpcResponseHeaderProto
	RpcSaslProto
	TokenProto
	CredentialsKVProto
	CredentialsProto
	GetDelegationTokenRequestProto
	GetDelegationTokenResponseProto
	RenewDelegationTokenRequestProto
	RenewDelegationTokenResponseProto
	CancelDelegationTokenRequestProto
	CancelDelegationTokenResponseProto
	ListSpanReceiversRequestProto
	SpanReceiverListInfo
	ListSpanReceiversResponseProto
	ConfigPair
	AddSpanReceiverRequestProto
	AddSpanReceiverResponseProto
	RemoveSpanReceiverRequestProto
	RemoveSpanReceiverResponseProto
	CedeActiveRequestProto
	CedeActiveResponseProto
	GracefulFailoverRequestProto
	GracefulFailoverResponseProto
*/
package hproto_common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// *
//  Refresh request.
type GenericRefreshRequestProto struct {
	Identifier       *string  `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Args             []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GenericRefreshRequestProto) Reset()                    { *m = GenericRefreshRequestProto{} }
func (m *GenericRefreshRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GenericRefreshRequestProto) ProtoMessage()               {}
func (*GenericRefreshRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GenericRefreshRequestProto) GetIdentifier() string {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return ""
}

func (m *GenericRefreshRequestProto) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

// *
// A single response from a refresh handler.
type GenericRefreshResponseProto struct {
	ExitStatus       *int32  `protobuf:"varint,1,opt,name=exitStatus" json:"exitStatus,omitempty"`
	UserMessage      *string `protobuf:"bytes,2,opt,name=userMessage" json:"userMessage,omitempty"`
	SenderName       *string `protobuf:"bytes,3,opt,name=senderName" json:"senderName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GenericRefreshResponseProto) Reset()                    { *m = GenericRefreshResponseProto{} }
func (m *GenericRefreshResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GenericRefreshResponseProto) ProtoMessage()               {}
func (*GenericRefreshResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GenericRefreshResponseProto) GetExitStatus() int32 {
	if m != nil && m.ExitStatus != nil {
		return *m.ExitStatus
	}
	return 0
}

func (m *GenericRefreshResponseProto) GetUserMessage() string {
	if m != nil && m.UserMessage != nil {
		return *m.UserMessage
	}
	return ""
}

func (m *GenericRefreshResponseProto) GetSenderName() string {
	if m != nil && m.SenderName != nil {
		return *m.SenderName
	}
	return ""
}

// *
// Collection of responses from zero or more handlers.
type GenericRefreshResponseCollectionProto struct {
	Responses        []*GenericRefreshResponseProto `protobuf:"bytes,1,rep,name=responses" json:"responses,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *GenericRefreshResponseCollectionProto) Reset()         { *m = GenericRefreshResponseCollectionProto{} }
func (m *GenericRefreshResponseCollectionProto) String() string { return proto.CompactTextString(m) }
func (*GenericRefreshResponseCollectionProto) ProtoMessage()    {}
func (*GenericRefreshResponseCollectionProto) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2}
}

func (m *GenericRefreshResponseCollectionProto) GetResponses() []*GenericRefreshResponseProto {
	if m != nil {
		return m.Responses
	}
	return nil
}

func init() {
	proto.RegisterType((*GenericRefreshRequestProto)(nil), "hproto.common.GenericRefreshRequestProto")
	proto.RegisterType((*GenericRefreshResponseProto)(nil), "hproto.common.GenericRefreshResponseProto")
	proto.RegisterType((*GenericRefreshResponseCollectionProto)(nil), "hproto.common.GenericRefreshResponseCollectionProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for GenericRefreshProtocolService service

type GenericRefreshProtocolServiceClient interface {
	Refresh(ctx context.Context, in *GenericRefreshRequestProto, opts ...grpc.CallOption) (*GenericRefreshResponseCollectionProto, error)
}

type genericRefreshProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewGenericRefreshProtocolServiceClient(cc *grpc.ClientConn) GenericRefreshProtocolServiceClient {
	return &genericRefreshProtocolServiceClient{cc}
}

func (c *genericRefreshProtocolServiceClient) Refresh(ctx context.Context, in *GenericRefreshRequestProto, opts ...grpc.CallOption) (*GenericRefreshResponseCollectionProto, error) {
	out := new(GenericRefreshResponseCollectionProto)
	err := grpc.Invoke(ctx, "/hproto.common.GenericRefreshProtocolService/refresh", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GenericRefreshProtocolService service

type GenericRefreshProtocolServiceServer interface {
	Refresh(context.Context, *GenericRefreshRequestProto) (*GenericRefreshResponseCollectionProto, error)
}

func RegisterGenericRefreshProtocolServiceServer(s *grpc.Server, srv GenericRefreshProtocolServiceServer) {
	s.RegisterService(&_GenericRefreshProtocolService_serviceDesc, srv)
}

func _GenericRefreshProtocolService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericRefreshRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericRefreshProtocolServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hproto.common.GenericRefreshProtocolService/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericRefreshProtocolServiceServer).Refresh(ctx, req.(*GenericRefreshRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _GenericRefreshProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hproto.common.GenericRefreshProtocolService",
	HandlerType: (*GenericRefreshProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "refresh",
			Handler:    _GenericRefreshProtocolService_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("gopkg.in/freddierice/go-hproto.v1/common/GenericRefreshProtocol.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x4a, 0x33, 0x31,
	0x10, 0x25, 0x5f, 0x3f, 0x91, 0x4e, 0xf1, 0x26, 0x57, 0xa5, 0x55, 0x29, 0x05, 0xa1, 0x0a, 0x66,
	0xb1, 0xf8, 0x02, 0x56, 0x44, 0x6f, 0x94, 0xb2, 0x7d, 0x82, 0xb0, 0x3b, 0xcd, 0x46, 0xdb, 0x4c,
	0x9a, 0xa4, 0xc5, 0x3b, 0x6f, 0x05, 0x5f, 0xc2, 0x47, 0x75, 0x7f, 0x8a, 0xfd, 0x61, 0x51, 0xaf,
	0x12, 0xce, 0x39, 0x73, 0xe6, 0xcc, 0x24, 0x70, 0xa7, 0xc8, 0xbe, 0x28, 0xa1, 0x4d, 0x34, 0x75,
	0x98, 0xa6, 0x1a, 0x9d, 0x4e, 0x30, 0x52, 0x74, 0x99, 0x59, 0x47, 0x81, 0xc4, 0xea, 0x2a, 0x4a,
	0x68, 0x3e, 0x27, 0x13, 0xdd, 0xa3, 0x29, 0xc8, 0x18, 0x73, 0xa1, 0xcf, 0xc6, 0x05, 0x9b, 0xd0,
	0x4c, 0x94, 0x32, 0x7e, 0xb4, 0x96, 0x57, 0xda, 0xfe, 0x18, 0x3a, 0xbb, 0xf2, 0x18, 0x17, 0x4b,
	0xf4, 0xa1, 0xac, 0xe2, 0xa7, 0x00, 0x3a, 0x45, 0x13, 0xf4, 0x34, 0x6f, 0xd7, 0x66, 0x3d, 0x36,
	0x68, 0xc6, 0x5b, 0x08, 0xe7, 0xf0, 0x5f, 0x3a, 0xe5, 0xdb, 0xff, 0x7a, 0x8d, 0x9c, 0x29, 0xef,
	0xfd, 0x37, 0xe8, 0xee, 0x3b, 0x7a, 0x4b, 0xc6, 0xe3, 0xb7, 0x25, 0xbe, 0xea, 0x30, 0x09, 0x32,
	0x2c, 0x7d, 0x69, 0x79, 0x10, 0x6f, 0x21, 0xbc, 0x07, 0xad, 0xa5, 0x47, 0xf7, 0x88, 0xde, 0x4b,
	0x85, 0xb9, 0x73, 0xd1, 0x73, 0x1b, 0x2a, 0x1c, 0x3c, 0x9a, 0x14, 0xdd, 0x93, 0x9c, 0x63, 0xbb,
	0x51, 0x85, 0xda, 0x20, 0xfd, 0x05, 0x9c, 0xd5, 0x07, 0xb8, 0xa5, 0xd9, 0x0c, 0x93, 0xa0, 0xc9,
	0x54, 0x51, 0x1e, 0xa0, 0xe9, 0xd6, 0x54, 0x91, 0xa4, 0x31, 0x68, 0x0d, 0x2f, 0xc4, 0xce, 0x7a,
	0xc4, 0x0f, 0x93, 0xc4, 0x9b, 0xe2, 0xe1, 0x07, 0x83, 0x93, 0xfa, 0xad, 0x4f, 0xd0, 0xad, 0xf2,
	0x77, 0xe2, 0xcf, 0x70, 0xe8, 0x2a, 0x86, 0x9f, 0xff, 0xd2, 0x63, 0xb3, 0xff, 0xce, 0xf5, 0x9f,
	0xe2, 0xec, 0xcd, 0x35, 0xba, 0x81, 0x2e, 0x39, 0x25, 0xa4, 0x95, 0x49, 0x86, 0x22, 0x93, 0x29,
	0x91, 0x15, 0xda, 0x26, 0xd5, 0x0f, 0x18, 0x1d, 0xd7, 0x27, 0x2d, 0x4f, 0xff, 0xce, 0xd8, 0x27,
	0x63, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x1c, 0xd9, 0xd8, 0x6d, 0x02, 0x00, 0x00,
}