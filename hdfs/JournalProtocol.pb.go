// Code generated by protoc-gen-go.
// source: gopkg.in/freddierice/go-hproto.v1/hdfs/JournalProtocol.proto
// DO NOT EDIT!

package hproto_hdfs

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

// *
// Journal information used by the journal receiver to identify a journal.
type JournalInfoProto struct {
	ClusterID        *string `protobuf:"bytes,1,req,name=clusterID" json:"clusterID,omitempty"`
	LayoutVersion    *uint32 `protobuf:"varint,2,opt,name=layoutVersion" json:"layoutVersion,omitempty"`
	NamespaceID      *uint32 `protobuf:"varint,3,opt,name=namespaceID" json:"namespaceID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *JournalInfoProto) Reset()                    { *m = JournalInfoProto{} }
func (m *JournalInfoProto) String() string            { return proto.CompactTextString(m) }
func (*JournalInfoProto) ProtoMessage()               {}
func (*JournalInfoProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *JournalInfoProto) GetClusterID() string {
	if m != nil && m.ClusterID != nil {
		return *m.ClusterID
	}
	return ""
}

func (m *JournalInfoProto) GetLayoutVersion() uint32 {
	if m != nil && m.LayoutVersion != nil {
		return *m.LayoutVersion
	}
	return 0
}

func (m *JournalInfoProto) GetNamespaceID() uint32 {
	if m != nil && m.NamespaceID != nil {
		return *m.NamespaceID
	}
	return 0
}

// *
// journalInfo - the information about the journal
// firstTxnId - the first txid in the journal records
// numTxns - Number of transactions in editlog
// records - bytes containing serialized journal records
// epoch - change to this represents change of journal writer
type JournalRequestProto struct {
	JournalInfo      *JournalInfoProto `protobuf:"bytes,1,req,name=journalInfo" json:"journalInfo,omitempty"`
	FirstTxnId       *uint64           `protobuf:"varint,2,req,name=firstTxnId" json:"firstTxnId,omitempty"`
	NumTxns          *uint32           `protobuf:"varint,3,req,name=numTxns" json:"numTxns,omitempty"`
	Records          []byte            `protobuf:"bytes,4,req,name=records" json:"records,omitempty"`
	Epoch            *uint64           `protobuf:"varint,5,req,name=epoch" json:"epoch,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *JournalRequestProto) Reset()                    { *m = JournalRequestProto{} }
func (m *JournalRequestProto) String() string            { return proto.CompactTextString(m) }
func (*JournalRequestProto) ProtoMessage()               {}
func (*JournalRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *JournalRequestProto) GetJournalInfo() *JournalInfoProto {
	if m != nil {
		return m.JournalInfo
	}
	return nil
}

func (m *JournalRequestProto) GetFirstTxnId() uint64 {
	if m != nil && m.FirstTxnId != nil {
		return *m.FirstTxnId
	}
	return 0
}

func (m *JournalRequestProto) GetNumTxns() uint32 {
	if m != nil && m.NumTxns != nil {
		return *m.NumTxns
	}
	return 0
}

func (m *JournalRequestProto) GetRecords() []byte {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *JournalRequestProto) GetEpoch() uint64 {
	if m != nil && m.Epoch != nil {
		return *m.Epoch
	}
	return 0
}

// *
// void response
type JournalResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *JournalResponseProto) Reset()                    { *m = JournalResponseProto{} }
func (m *JournalResponseProto) String() string            { return proto.CompactTextString(m) }
func (*JournalResponseProto) ProtoMessage()               {}
func (*JournalResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

// *
// journalInfo - the information about the journal
// txid - first txid in the new log
type StartLogSegmentRequestProto struct {
	JournalInfo      *JournalInfoProto `protobuf:"bytes,1,req,name=journalInfo" json:"journalInfo,omitempty"`
	Txid             *uint64           `protobuf:"varint,2,req,name=txid" json:"txid,omitempty"`
	Epoch            *uint64           `protobuf:"varint,3,req,name=epoch" json:"epoch,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *StartLogSegmentRequestProto) Reset()                    { *m = StartLogSegmentRequestProto{} }
func (m *StartLogSegmentRequestProto) String() string            { return proto.CompactTextString(m) }
func (*StartLogSegmentRequestProto) ProtoMessage()               {}
func (*StartLogSegmentRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *StartLogSegmentRequestProto) GetJournalInfo() *JournalInfoProto {
	if m != nil {
		return m.JournalInfo
	}
	return nil
}

func (m *StartLogSegmentRequestProto) GetTxid() uint64 {
	if m != nil && m.Txid != nil {
		return *m.Txid
	}
	return 0
}

func (m *StartLogSegmentRequestProto) GetEpoch() uint64 {
	if m != nil && m.Epoch != nil {
		return *m.Epoch
	}
	return 0
}

// *
// void response
type StartLogSegmentResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StartLogSegmentResponseProto) Reset()                    { *m = StartLogSegmentResponseProto{} }
func (m *StartLogSegmentResponseProto) String() string            { return proto.CompactTextString(m) }
func (*StartLogSegmentResponseProto) ProtoMessage()               {}
func (*StartLogSegmentResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

// *
// journalInfo - the information about the journal
// txid - first txid in the new log
type FenceRequestProto struct {
	JournalInfo      *JournalInfoProto `protobuf:"bytes,1,req,name=journalInfo" json:"journalInfo,omitempty"`
	Epoch            *uint64           `protobuf:"varint,2,req,name=epoch" json:"epoch,omitempty"`
	FencerInfo       *string           `protobuf:"bytes,3,opt,name=fencerInfo" json:"fencerInfo,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *FenceRequestProto) Reset()                    { *m = FenceRequestProto{} }
func (m *FenceRequestProto) String() string            { return proto.CompactTextString(m) }
func (*FenceRequestProto) ProtoMessage()               {}
func (*FenceRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *FenceRequestProto) GetJournalInfo() *JournalInfoProto {
	if m != nil {
		return m.JournalInfo
	}
	return nil
}

func (m *FenceRequestProto) GetEpoch() uint64 {
	if m != nil && m.Epoch != nil {
		return *m.Epoch
	}
	return 0
}

func (m *FenceRequestProto) GetFencerInfo() string {
	if m != nil && m.FencerInfo != nil {
		return *m.FencerInfo
	}
	return ""
}

// *
// previousEpoch - previous epoch if any or zero
// lastTransactionId - last valid transaction Id in the journal
// inSync - if all journal segments are available and in sync
type FenceResponseProto struct {
	PreviousEpoch     *uint64 `protobuf:"varint,1,opt,name=previousEpoch" json:"previousEpoch,omitempty"`
	LastTransactionId *uint64 `protobuf:"varint,2,opt,name=lastTransactionId" json:"lastTransactionId,omitempty"`
	InSync            *bool   `protobuf:"varint,3,opt,name=inSync" json:"inSync,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *FenceResponseProto) Reset()                    { *m = FenceResponseProto{} }
func (m *FenceResponseProto) String() string            { return proto.CompactTextString(m) }
func (*FenceResponseProto) ProtoMessage()               {}
func (*FenceResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *FenceResponseProto) GetPreviousEpoch() uint64 {
	if m != nil && m.PreviousEpoch != nil {
		return *m.PreviousEpoch
	}
	return 0
}

func (m *FenceResponseProto) GetLastTransactionId() uint64 {
	if m != nil && m.LastTransactionId != nil {
		return *m.LastTransactionId
	}
	return 0
}

func (m *FenceResponseProto) GetInSync() bool {
	if m != nil && m.InSync != nil {
		return *m.InSync
	}
	return false
}

func init() {
	proto.RegisterType((*JournalInfoProto)(nil), "hproto.hdfs.JournalInfoProto")
	proto.RegisterType((*JournalRequestProto)(nil), "hproto.hdfs.JournalRequestProto")
	proto.RegisterType((*JournalResponseProto)(nil), "hproto.hdfs.JournalResponseProto")
	proto.RegisterType((*StartLogSegmentRequestProto)(nil), "hproto.hdfs.StartLogSegmentRequestProto")
	proto.RegisterType((*StartLogSegmentResponseProto)(nil), "hproto.hdfs.StartLogSegmentResponseProto")
	proto.RegisterType((*FenceRequestProto)(nil), "hproto.hdfs.FenceRequestProto")
	proto.RegisterType((*FenceResponseProto)(nil), "hproto.hdfs.FenceResponseProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for JournalProtocolService service

type JournalProtocolServiceClient interface {
	// *
	// Request sent by active namenode to backup node via
	// EditLogBackupOutputStream to stream editlog records.
	Journal(ctx context.Context, in *JournalRequestProto, opts ...grpc.CallOption) (*JournalResponseProto, error)
	// *
	// Request sent by active namenode to backup node to notify
	// that the NameNode has rolled its edit logs and is now writing a
	// new log segment.
	StartLogSegment(ctx context.Context, in *StartLogSegmentRequestProto, opts ...grpc.CallOption) (*StartLogSegmentResponseProto, error)
	// *
	// Request to fence a journal receiver.
	Fence(ctx context.Context, in *FenceRequestProto, opts ...grpc.CallOption) (*FenceResponseProto, error)
}

type journalProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewJournalProtocolServiceClient(cc *grpc.ClientConn) JournalProtocolServiceClient {
	return &journalProtocolServiceClient{cc}
}

func (c *journalProtocolServiceClient) Journal(ctx context.Context, in *JournalRequestProto, opts ...grpc.CallOption) (*JournalResponseProto, error) {
	out := new(JournalResponseProto)
	err := grpc.Invoke(ctx, "/hproto.hdfs.JournalProtocolService/journal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalProtocolServiceClient) StartLogSegment(ctx context.Context, in *StartLogSegmentRequestProto, opts ...grpc.CallOption) (*StartLogSegmentResponseProto, error) {
	out := new(StartLogSegmentResponseProto)
	err := grpc.Invoke(ctx, "/hproto.hdfs.JournalProtocolService/startLogSegment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalProtocolServiceClient) Fence(ctx context.Context, in *FenceRequestProto, opts ...grpc.CallOption) (*FenceResponseProto, error) {
	out := new(FenceResponseProto)
	err := grpc.Invoke(ctx, "/hproto.hdfs.JournalProtocolService/fence", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JournalProtocolService service

type JournalProtocolServiceServer interface {
	// *
	// Request sent by active namenode to backup node via
	// EditLogBackupOutputStream to stream editlog records.
	Journal(context.Context, *JournalRequestProto) (*JournalResponseProto, error)
	// *
	// Request sent by active namenode to backup node to notify
	// that the NameNode has rolled its edit logs and is now writing a
	// new log segment.
	StartLogSegment(context.Context, *StartLogSegmentRequestProto) (*StartLogSegmentResponseProto, error)
	// *
	// Request to fence a journal receiver.
	Fence(context.Context, *FenceRequestProto) (*FenceResponseProto, error)
}

func RegisterJournalProtocolServiceServer(s *grpc.Server, srv JournalProtocolServiceServer) {
	s.RegisterService(&_JournalProtocolService_serviceDesc, srv)
}

func _JournalProtocolService_Journal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JournalRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalProtocolServiceServer).Journal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hproto.hdfs.JournalProtocolService/Journal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalProtocolServiceServer).Journal(ctx, req.(*JournalRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalProtocolService_StartLogSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartLogSegmentRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalProtocolServiceServer).StartLogSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hproto.hdfs.JournalProtocolService/StartLogSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalProtocolServiceServer).StartLogSegment(ctx, req.(*StartLogSegmentRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalProtocolService_Fence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FenceRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalProtocolServiceServer).Fence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hproto.hdfs.JournalProtocolService/Fence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalProtocolServiceServer).Fence(ctx, req.(*FenceRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _JournalProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hproto.hdfs.JournalProtocolService",
	HandlerType: (*JournalProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "journal",
			Handler:    _JournalProtocolService_Journal_Handler,
		},
		{
			MethodName: "startLogSegment",
			Handler:    _JournalProtocolService_StartLogSegment_Handler,
		},
		{
			MethodName: "fence",
			Handler:    _JournalProtocolService_Fence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor4,
}

func init() {
	proto.RegisterFile("gopkg.in/freddierice/go-hproto.v1/hdfs/JournalProtocol.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0xc0, 0xe5, 0xac, 0x63, 0xf4, 0x95, 0x09, 0x66, 0x46, 0x55, 0x95, 0x51, 0x4a, 0x04, 0x52,
	0x91, 0x20, 0xd5, 0x76, 0xe1, 0x82, 0x84, 0x34, 0x0d, 0xb4, 0x22, 0x84, 0x50, 0x8a, 0xb8, 0x47,
	0xc9, 0x4b, 0x1a, 0x48, 0xed, 0x60, 0x3b, 0x55, 0x77, 0xe3, 0x38, 0xf1, 0x11, 0x38, 0xf1, 0x51,
	0xf8, 0x68, 0x38, 0x6e, 0xba, 0x3a, 0x6d, 0x35, 0xed, 0xb0, 0x4b, 0x6a, 0xbf, 0xbf, 0xbf, 0xf7,
	0x9e, 0x5f, 0xe1, 0x6d, 0xc2, 0xf3, 0x1f, 0x89, 0x97, 0xb2, 0x61, 0x2c, 0x30, 0x8a, 0x52, 0x14,
	0x69, 0x88, 0xc3, 0x84, 0xbf, 0x9e, 0xe4, 0x82, 0x2b, 0xee, 0xcd, 0x8e, 0x87, 0x93, 0x28, 0x96,
	0xc3, 0x8f, 0xbc, 0x10, 0x2c, 0xc8, 0xbe, 0x94, 0xc2, 0x90, 0x67, 0x9e, 0xd1, 0xd2, 0x56, 0x65,
	0x55, 0x9a, 0x74, 0x8f, 0x6f, 0x18, 0xaa, 0xfc, 0x2c, 0xfc, 0xbb, 0x6f, 0x6e, 0xe8, 0x72, 0xae,
	0x3f, 0x63, 0x14, 0x33, 0x14, 0x0b, 0x47, 0x77, 0x0e, 0x0f, 0x2a, 0xa2, 0x11, 0x8b, 0xb9, 0xa1,
	0xa2, 0x47, 0xd0, 0x0c, 0xb3, 0x42, 0x2a, 0x14, 0xa3, 0xb3, 0x0e, 0xe9, 0x3b, 0x83, 0xa6, 0xbf,
	0x12, 0xd0, 0xe7, 0xb0, 0x9f, 0x05, 0x17, 0xbc, 0x50, 0xdf, 0x50, 0xc8, 0x94, 0xb3, 0x8e, 0xd3,
	0x27, 0x83, 0x7d, 0xbf, 0x2e, 0xa4, 0x7d, 0x68, 0xb1, 0x60, 0x8a, 0x32, 0x0f, 0x42, 0xd4, 0x51,
	0x76, 0x8c, 0x8d, 0x2d, 0x72, 0xff, 0x11, 0x78, 0x58, 0xa5, 0xf6, 0xf1, 0x67, 0x81, 0x52, 0x2d,
	0xb2, 0xbf, 0x83, 0xd6, 0xf7, 0x15, 0x91, 0xc9, 0xdf, 0x3a, 0x79, 0xe2, 0x59, 0x0d, 0xf2, 0xd6,
	0x89, 0x7d, 0xdb, 0x83, 0xf6, 0x00, 0xe2, 0x54, 0x48, 0xf5, 0x75, 0xce, 0x46, 0x91, 0xa6, 0x73,
	0x06, 0x0d, 0xdf, 0x92, 0xd0, 0x0e, 0xec, 0xb1, 0x62, 0xaa, 0xcf, 0x52, 0x63, 0x39, 0x1a, 0x6b,
	0x79, 0x2d, 0x35, 0x02, 0x43, 0x2e, 0x22, 0xd9, 0x69, 0x68, 0xcd, 0x3d, 0x7f, 0x79, 0xa5, 0x87,
	0xb0, 0x8b, 0x39, 0x0f, 0x27, 0x9d, 0x5d, 0x13, 0x6e, 0x71, 0x71, 0xdb, 0x70, 0x78, 0x55, 0x81,
	0xcc, 0x39, 0x93, 0x68, 0x70, 0xdc, 0x4b, 0x02, 0x8f, 0xc7, 0x2a, 0x10, 0xea, 0x13, 0x4f, 0xc6,
	0x98, 0x4c, 0x91, 0xa9, 0xdb, 0x2d, 0x91, 0x42, 0x43, 0xcd, 0xd3, 0x65, 0x71, 0xe6, 0xbc, 0x42,
	0xdc, 0xb1, 0x11, 0x7b, 0x70, 0xb4, 0x41, 0x62, 0xa3, 0xfe, 0x26, 0x70, 0xf0, 0x01, 0x59, 0x88,
	0xb7, 0x0b, 0x78, 0x05, 0xe3, 0x58, 0x30, 0x66, 0x32, 0x65, 0x2e, 0x61, 0xa2, 0x96, 0x6f, 0xa2,
	0xe9, 0x5b, 0x12, 0xf7, 0x17, 0x01, 0x5a, 0xc1, 0x58, 0x8c, 0xe5, 0x8b, 0xcb, 0x05, 0xce, 0x52,
	0x5e, 0xc8, 0xf7, 0x26, 0x28, 0xd1, 0x9e, 0x0d, 0xbf, 0x2e, 0xa4, 0xaf, 0xe0, 0x20, 0x0b, 0xf4,
	0x8c, 0x45, 0xc0, 0x64, 0x10, 0x2a, 0xfd, 0x08, 0xcd, 0xf4, 0x4b, 0xcb, 0x4d, 0x05, 0x6d, 0xc3,
	0x9d, 0x94, 0x8d, 0x2f, 0x58, 0x68, 0x30, 0xee, 0xfa, 0xd5, 0xed, 0xe4, 0x8f, 0x03, 0xed, 0xb5,
	0x15, 0x2d, 0xf7, 0x45, 0x6f, 0x13, 0xfd, 0x0c, 0x7b, 0x55, 0x89, 0xb4, 0xbf, 0xad, 0x15, 0x76,
	0x07, 0xbb, 0xcf, 0xb6, 0x5b, 0xd8, 0x65, 0xc5, 0x70, 0x5f, 0xd6, 0x47, 0x43, 0x07, 0x35, 0xaf,
	0x6b, 0x9e, 0x50, 0xf7, 0xe5, 0xf5, 0x96, 0x76, 0x9e, 0x73, 0xd8, 0x35, 0x3d, 0xa6, 0xbd, 0x9a,
	0xcf, 0xc6, 0xd4, 0xbb, 0x4f, 0xb7, 0xe9, 0xad, 0x48, 0xa7, 0x67, 0xf0, 0x82, 0x8b, 0xc4, 0x0b,
	0xf4, 0x06, 0x4f, 0xb0, 0x66, 0x9c, 0xd7, 0xfe, 0xce, 0x4e, 0x1f, 0xad, 0xb5, 0xd0, 0xfc, 0xca,
	0x4b, 0x42, 0xfe, 0x12, 0xf2, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x78, 0xf4, 0x4c, 0xca, 0x2a, 0x05,
	0x00, 0x00,
}