// Code generated by protoc-gen-go.
// source: gopkg.in/freddierice/go-hproto.v1/yarn/yarn_server_nodemanager_recovery.proto
// DO NOT EDIT!

package hproto_yarn

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ContainerManagerApplicationProto struct {
	Id                          *ApplicationIdProto         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	User                        *string                     `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Credentials                 []byte                      `protobuf:"bytes,3,opt,name=credentials" json:"credentials,omitempty"`
	Acls                        []*ApplicationACLMapProto   `protobuf:"bytes,4,rep,name=acls" json:"acls,omitempty"`
	LogAggregationContext       *LogAggregationContextProto `protobuf:"bytes,5,opt,name=log_aggregation_context,json=logAggregationContext" json:"log_aggregation_context,omitempty"`
	AppLogAggregationInitedTime *int64                      `protobuf:"varint,6,opt,name=appLogAggregationInitedTime,def=-1" json:"appLogAggregationInitedTime,omitempty"`
	XXX_unrecognized            []byte                      `json:"-"`
}

func (m *ContainerManagerApplicationProto) Reset()         { *m = ContainerManagerApplicationProto{} }
func (m *ContainerManagerApplicationProto) String() string { return proto.CompactTextString(m) }
func (*ContainerManagerApplicationProto) ProtoMessage()    {}
func (*ContainerManagerApplicationProto) Descriptor() ([]byte, []int) {
	return fileDescriptor16, []int{0}
}

const Default_ContainerManagerApplicationProto_AppLogAggregationInitedTime int64 = -1

func (m *ContainerManagerApplicationProto) GetId() *ApplicationIdProto {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ContainerManagerApplicationProto) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *ContainerManagerApplicationProto) GetCredentials() []byte {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *ContainerManagerApplicationProto) GetAcls() []*ApplicationACLMapProto {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *ContainerManagerApplicationProto) GetLogAggregationContext() *LogAggregationContextProto {
	if m != nil {
		return m.LogAggregationContext
	}
	return nil
}

func (m *ContainerManagerApplicationProto) GetAppLogAggregationInitedTime() int64 {
	if m != nil && m.AppLogAggregationInitedTime != nil {
		return *m.AppLogAggregationInitedTime
	}
	return Default_ContainerManagerApplicationProto_AppLogAggregationInitedTime
}

type DeletionServiceDeleteTaskProto struct {
	Id               *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	User             *string  `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Subdir           *string  `protobuf:"bytes,3,opt,name=subdir" json:"subdir,omitempty"`
	DeletionTime     *int64   `protobuf:"varint,4,opt,name=deletionTime" json:"deletionTime,omitempty"`
	Basedirs         []string `protobuf:"bytes,5,rep,name=basedirs" json:"basedirs,omitempty"`
	SuccessorIds     []int32  `protobuf:"varint,6,rep,name=successorIds" json:"successorIds,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DeletionServiceDeleteTaskProto) Reset()                    { *m = DeletionServiceDeleteTaskProto{} }
func (m *DeletionServiceDeleteTaskProto) String() string            { return proto.CompactTextString(m) }
func (*DeletionServiceDeleteTaskProto) ProtoMessage()               {}
func (*DeletionServiceDeleteTaskProto) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

func (m *DeletionServiceDeleteTaskProto) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *DeletionServiceDeleteTaskProto) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *DeletionServiceDeleteTaskProto) GetSubdir() string {
	if m != nil && m.Subdir != nil {
		return *m.Subdir
	}
	return ""
}

func (m *DeletionServiceDeleteTaskProto) GetDeletionTime() int64 {
	if m != nil && m.DeletionTime != nil {
		return *m.DeletionTime
	}
	return 0
}

func (m *DeletionServiceDeleteTaskProto) GetBasedirs() []string {
	if m != nil {
		return m.Basedirs
	}
	return nil
}

func (m *DeletionServiceDeleteTaskProto) GetSuccessorIds() []int32 {
	if m != nil {
		return m.SuccessorIds
	}
	return nil
}

type LocalizedResourceProto struct {
	Resource         *LocalResourceProto `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	LocalPath        *string             `protobuf:"bytes,2,opt,name=localPath" json:"localPath,omitempty"`
	Size             *int64              `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *LocalizedResourceProto) Reset()                    { *m = LocalizedResourceProto{} }
func (m *LocalizedResourceProto) String() string            { return proto.CompactTextString(m) }
func (*LocalizedResourceProto) ProtoMessage()               {}
func (*LocalizedResourceProto) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

func (m *LocalizedResourceProto) GetResource() *LocalResourceProto {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *LocalizedResourceProto) GetLocalPath() string {
	if m != nil && m.LocalPath != nil {
		return *m.LocalPath
	}
	return ""
}

func (m *LocalizedResourceProto) GetSize() int64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

type LogDeleterProto struct {
	User             *string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	DeletionTime     *int64  `protobuf:"varint,2,opt,name=deletionTime" json:"deletionTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogDeleterProto) Reset()                    { *m = LogDeleterProto{} }
func (m *LogDeleterProto) String() string            { return proto.CompactTextString(m) }
func (*LogDeleterProto) ProtoMessage()               {}
func (*LogDeleterProto) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

func (m *LogDeleterProto) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *LogDeleterProto) GetDeletionTime() int64 {
	if m != nil && m.DeletionTime != nil {
		return *m.DeletionTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ContainerManagerApplicationProto)(nil), "hproto.yarn.ContainerManagerApplicationProto")
	proto.RegisterType((*DeletionServiceDeleteTaskProto)(nil), "hproto.yarn.DeletionServiceDeleteTaskProto")
	proto.RegisterType((*LocalizedResourceProto)(nil), "hproto.yarn.LocalizedResourceProto")
	proto.RegisterType((*LogDeleterProto)(nil), "hproto.yarn.LogDeleterProto")
}

func init() {
	proto.RegisterFile("gopkg.in/freddierice/go-hproto.v1/yarn/yarn_server_nodemanager_recovery.proto", fileDescriptor16)
}

var fileDescriptor16 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0x92, 0x6e, 0xda, 0xdc, 0x09, 0x24, 0x4b, 0x8c, 0x68, 0x4c, 0x50, 0x65, 0x0f, 0xf0,
	0xb2, 0x54, 0xe3, 0x05, 0x04, 0x4f, 0x5d, 0xf7, 0x12, 0xa9, 0x45, 0x93, 0xb7, 0x17, 0x9e, 0x22,
	0xcf, 0xbe, 0xb8, 0xd6, 0xb2, 0x38, 0xb2, 0xd3, 0x89, 0xf1, 0x03, 0xf0, 0x19, 0x7c, 0x0a, 0xbf,
	0xc3, 0x5f, 0x70, 0xe3, 0x84, 0xb6, 0x81, 0x0d, 0x89, 0x97, 0xea, 0xfa, 0xfa, 0x9c, 0x7b, 0xce,
	0x3d, 0xa9, 0xc9, 0x5c, 0x99, 0xea, 0x5a, 0xa5, 0xba, 0x1c, 0x7f, 0xb2, 0x20, 0xa5, 0x06, 0xab,
	0x05, 0x8c, 0x95, 0x39, 0x5e, 0x54, 0xd6, 0xd4, 0x26, 0xbd, 0x3d, 0x19, 0xdf, 0x71, 0x5b, 0xfa,
	0x9f, 0xdc, 0x81, 0xbd, 0x05, 0x9b, 0x97, 0x46, 0xc2, 0x0d, 0x2f, 0xb9, 0xc2, 0xda, 0x82, 0x30,
	0xd8, 0xbb, 0x4b, 0x3d, 0x9c, 0x0e, 0x3b, 0x5a, 0x03, 0x3f, 0x78, 0xfb, 0x3f, 0xb3, 0x7d, 0xcb,
	0xb5, 0x63, 0x92, 0x9f, 0x21, 0x19, 0x4d, 0x4d, 0x59, 0x73, 0x5d, 0x82, 0x9d, 0xb7, 0x52, 0x93,
	0xaa, 0x2a, 0xb4, 0xe0, 0xb5, 0x36, 0xe5, 0xb9, 0xd7, 0x1a, 0x93, 0x50, 0xcb, 0x38, 0x18, 0x05,
	0xaf, 0x86, 0xaf, 0x5f, 0xa4, 0x1b, 0xc2, 0xe9, 0x06, 0x34, 0x93, 0x1e, 0xcc, 0x10, 0x4a, 0x29,
	0x19, 0x2c, 0x71, 0x85, 0x38, 0x44, 0xca, 0x2e, 0xf3, 0x35, 0x1d, 0x91, 0xa1, 0x40, 0x73, 0x50,
	0xd6, 0x9a, 0x17, 0x2e, 0x8e, 0xf0, 0x6a, 0x8f, 0x6d, 0xb6, 0xe8, 0x1b, 0x32, 0xe0, 0x02, 0xaf,
	0x06, 0xa3, 0x08, 0x85, 0x8e, 0x1e, 0x12, 0x9a, 0x4c, 0x67, 0x73, 0x5e, 0xb5, 0x62, 0x9e, 0x40,
	0x73, 0xf2, 0xb4, 0x30, 0x2a, 0xe7, 0x4a, 0x59, 0x50, 0x1e, 0x93, 0x0b, 0xdc, 0x09, 0x3e, 0xd7,
	0xf1, 0x96, 0x37, 0xfd, 0xb2, 0x37, 0x6b, 0x66, 0xd4, 0x64, 0x0d, 0x9d, 0xb6, 0xc8, 0x76, 0xde,
	0x93, 0xe2, 0xbe, 0x3b, 0x7a, 0x46, 0x9e, 0xf1, 0xaa, 0xea, 0xf3, 0xb2, 0x52, 0xd7, 0x20, 0x2f,
	0xf5, 0x0d, 0xc4, 0xdb, 0x28, 0x12, 0xbd, 0x0b, 0x8f, 0x4f, 0xd8, 0xbf, 0x60, 0xc9, 0x8f, 0x80,
	0x3c, 0x3f, 0x83, 0x02, 0x9a, 0xf6, 0x05, 0x7e, 0x60, 0xfc, 0x46, 0xfe, 0x08, 0x97, 0xdc, 0x5d,
	0xb7, 0x49, 0x3f, 0x5a, 0x25, 0xbd, 0xf5, 0x60, 0x90, 0xfb, 0x64, 0xdb, 0x2d, 0xaf, 0xa4, 0xb6,
	0x3e, 0xc3, 0x5d, 0xd6, 0x9d, 0x68, 0x42, 0xf6, 0x64, 0x37, 0xdd, 0xbb, 0x1a, 0x34, 0xae, 0x58,
	0xaf, 0x47, 0x0f, 0xc8, 0xce, 0x15, 0x77, 0x80, 0x70, 0x87, 0xd1, 0x44, 0xc8, 0x5e, 0x9d, 0x1b,
	0xbe, 0x5b, 0x0a, 0x01, 0xce, 0x19, 0x9b, 0x49, 0x87, 0x5b, 0x45, 0xe8, 0xa2, 0xd7, 0x4b, 0xbe,
	0x06, 0x64, 0x7f, 0x66, 0x04, 0x2f, 0xf4, 0x17, 0x90, 0x0c, 0x9c, 0x59, 0x5a, 0x01, 0xad, 0xf5,
	0xf7, 0x64, 0xc7, 0x76, 0x8d, 0x7b, 0xff, 0x2a, 0x9e, 0xd6, 0xa3, 0xb0, 0x15, 0x81, 0x1e, 0x92,
	0xdd, 0xa2, 0xb9, 0x3f, 0xe7, 0xf5, 0xa2, 0x5b, 0x76, 0xdd, 0x68, 0x52, 0x70, 0x28, 0xe8, 0xf7,
	0x8d, 0x98, 0xaf, 0x93, 0x8c, 0x3c, 0xc6, 0xa0, 0xdb, 0xfc, 0x6c, 0xeb, 0xe0, 0x77, 0x58, 0xc1,
	0x46, 0x58, 0x7f, 0x86, 0x12, 0xfe, 0x1d, 0xca, 0x69, 0x46, 0x0e, 0x8d, 0x55, 0x29, 0xaf, 0xb8,
	0x58, 0x40, 0xcf, 0xb3, 0x2f, 0x4f, 0x8f, 0x3e, 0x62, 0x7d, 0xe1, 0x5f, 0xe4, 0x87, 0xf5, 0x83,
	0x64, 0xdd, 0x7b, 0xf4, 0xda, 0xee, 0x5b, 0x10, 0x7c, 0x0f, 0x82, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x0e, 0xb1, 0x6b, 0x3a, 0xe5, 0x03, 0x00, 0x00,
}