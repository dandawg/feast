// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/core/CoreService.proto

package core

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ApplyFeatureSetResponse_Status int32

const (
	// Latest feature set version is consistent with provided feature set
	ApplyFeatureSetResponse_NO_CHANGE ApplyFeatureSetResponse_Status = 0
	// New feature set or feature set version created
	ApplyFeatureSetResponse_CREATED ApplyFeatureSetResponse_Status = 1
	// Error occurred while trying to apply changes
	ApplyFeatureSetResponse_ERROR ApplyFeatureSetResponse_Status = 2
)

var ApplyFeatureSetResponse_Status_name = map[int32]string{
	0: "NO_CHANGE",
	1: "CREATED",
	2: "ERROR",
}

var ApplyFeatureSetResponse_Status_value = map[string]int32{
	"NO_CHANGE": 0,
	"CREATED":   1,
	"ERROR":     2,
}

func (x ApplyFeatureSetResponse_Status) String() string {
	return proto.EnumName(ApplyFeatureSetResponse_Status_name, int32(x))
}

func (ApplyFeatureSetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d9be266444105411, []int{5, 0}
}

// Retrieves details for all versions of a specific feature set
type GetFeatureSetsRequest struct {
	Filter               *GetFeatureSetsRequest_Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GetFeatureSetsRequest) Reset()         { *m = GetFeatureSetsRequest{} }
func (m *GetFeatureSetsRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeatureSetsRequest) ProtoMessage()    {}
func (*GetFeatureSetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9be266444105411, []int{0}
}

func (m *GetFeatureSetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeatureSetsRequest.Unmarshal(m, b)
}
func (m *GetFeatureSetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeatureSetsRequest.Marshal(b, m, deterministic)
}
func (m *GetFeatureSetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeatureSetsRequest.Merge(m, src)
}
func (m *GetFeatureSetsRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeatureSetsRequest.Size(m)
}
func (m *GetFeatureSetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeatureSetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeatureSetsRequest proto.InternalMessageInfo

func (m *GetFeatureSetsRequest) GetFilter() *GetFeatureSetsRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type GetFeatureSetsRequest_Filter struct {
	FeatureSetName       string   `protobuf:"bytes,1,opt,name=featureSetName,proto3" json:"featureSetName,omitempty"`
	FeatureSetVersion    string   `protobuf:"bytes,2,opt,name=featureSetVersion,proto3" json:"featureSetVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeatureSetsRequest_Filter) Reset()         { *m = GetFeatureSetsRequest_Filter{} }
func (m *GetFeatureSetsRequest_Filter) String() string { return proto.CompactTextString(m) }
func (*GetFeatureSetsRequest_Filter) ProtoMessage()    {}
func (*GetFeatureSetsRequest_Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9be266444105411, []int{0, 0}
}

func (m *GetFeatureSetsRequest_Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeatureSetsRequest_Filter.Unmarshal(m, b)
}
func (m *GetFeatureSetsRequest_Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeatureSetsRequest_Filter.Marshal(b, m, deterministic)
}
func (m *GetFeatureSetsRequest_Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeatureSetsRequest_Filter.Merge(m, src)
}
func (m *GetFeatureSetsRequest_Filter) XXX_Size() int {
	return xxx_messageInfo_GetFeatureSetsRequest_Filter.Size(m)
}
func (m *GetFeatureSetsRequest_Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeatureSetsRequest_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeatureSetsRequest_Filter proto.InternalMessageInfo

func (m *GetFeatureSetsRequest_Filter) GetFeatureSetName() string {
	if m != nil {
		return m.FeatureSetName
	}
	return ""
}

func (m *GetFeatureSetsRequest_Filter) GetFeatureSetVersion() string {
	if m != nil {
		return m.FeatureSetVersion
	}
	return ""
}

type GetFeatureSetsResponse struct {
	FeatureSets          []*FeatureSetSpec `protobuf:"bytes,1,rep,name=featureSets,proto3" json:"featureSets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetFeatureSetsResponse) Reset()         { *m = GetFeatureSetsResponse{} }
func (m *GetFeatureSetsResponse) String() string { return proto.CompactTextString(m) }
func (*GetFeatureSetsResponse) ProtoMessage()    {}
func (*GetFeatureSetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9be266444105411, []int{1}
}

func (m *GetFeatureSetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeatureSetsResponse.Unmarshal(m, b)
}
func (m *GetFeatureSetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeatureSetsResponse.Marshal(b, m, deterministic)
}
func (m *GetFeatureSetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeatureSetsResponse.Merge(m, src)
}
func (m *GetFeatureSetsResponse) XXX_Size() int {
	return xxx_messageInfo_GetFeatureSetsResponse.Size(m)
}
func (m *GetFeatureSetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeatureSetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeatureSetsResponse proto.InternalMessageInfo

func (m *GetFeatureSetsResponse) GetFeatureSets() []*FeatureSetSpec {
	if m != nil {
		return m.FeatureSets
	}
	return nil
}

type GetStoresRequest struct {
	Filter               *GetStoresRequest_Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetStoresRequest) Reset()         { *m = GetStoresRequest{} }
func (m *GetStoresRequest) String() string { return proto.CompactTextString(m) }
func (*GetStoresRequest) ProtoMessage()    {}
func (*GetStoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9be266444105411, []int{2}
}

func (m *GetStoresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoresRequest.Unmarshal(m, b)
}
func (m *GetStoresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoresRequest.Marshal(b, m, deterministic)
}
func (m *GetStoresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoresRequest.Merge(m, src)
}
func (m *GetStoresRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoresRequest.Size(m)
}
func (m *GetStoresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoresRequest proto.InternalMessageInfo

func (m *GetStoresRequest) GetFilter() *GetStoresRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type GetStoresRequest_Filter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStoresRequest_Filter) Reset()         { *m = GetStoresRequest_Filter{} }
func (m *GetStoresRequest_Filter) String() string { return proto.CompactTextString(m) }
func (*GetStoresRequest_Filter) ProtoMessage()    {}
func (*GetStoresRequest_Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9be266444105411, []int{2, 0}
}

func (m *GetStoresRequest_Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoresRequest_Filter.Unmarshal(m, b)
}
func (m *GetStoresRequest_Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoresRequest_Filter.Marshal(b, m, deterministic)
}
func (m *GetStoresRequest_Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoresRequest_Filter.Merge(m, src)
}
func (m *GetStoresRequest_Filter) XXX_Size() int {
	return xxx_messageInfo_GetStoresRequest_Filter.Size(m)
}
func (m *GetStoresRequest_Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoresRequest_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoresRequest_Filter proto.InternalMessageInfo

func (m *GetStoresRequest_Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetStoresResponse struct {
	Store                []*Store `protobuf:"bytes,1,rep,name=store,proto3" json:"store,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStoresResponse) Reset()         { *m = GetStoresResponse{} }
func (m *GetStoresResponse) String() string { return proto.CompactTextString(m) }
func (*GetStoresResponse) ProtoMessage()    {}
func (*GetStoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9be266444105411, []int{3}
}

func (m *GetStoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoresResponse.Unmarshal(m, b)
}
func (m *GetStoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoresResponse.Marshal(b, m, deterministic)
}
func (m *GetStoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoresResponse.Merge(m, src)
}
func (m *GetStoresResponse) XXX_Size() int {
	return xxx_messageInfo_GetStoresResponse.Size(m)
}
func (m *GetStoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoresResponse proto.InternalMessageInfo

func (m *GetStoresResponse) GetStore() []*Store {
	if m != nil {
		return m.Store
	}
	return nil
}

type ApplyFeatureSetRequest struct {
	// Feature set version and source will be ignored
	FeatureSet           *FeatureSetSpec `protobuf:"bytes,1,opt,name=featureSet,proto3" json:"featureSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ApplyFeatureSetRequest) Reset()         { *m = ApplyFeatureSetRequest{} }
func (m *ApplyFeatureSetRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyFeatureSetRequest) ProtoMessage()    {}
func (*ApplyFeatureSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9be266444105411, []int{4}
}

func (m *ApplyFeatureSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyFeatureSetRequest.Unmarshal(m, b)
}
func (m *ApplyFeatureSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyFeatureSetRequest.Marshal(b, m, deterministic)
}
func (m *ApplyFeatureSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyFeatureSetRequest.Merge(m, src)
}
func (m *ApplyFeatureSetRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyFeatureSetRequest.Size(m)
}
func (m *ApplyFeatureSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyFeatureSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyFeatureSetRequest proto.InternalMessageInfo

func (m *ApplyFeatureSetRequest) GetFeatureSet() *FeatureSetSpec {
	if m != nil {
		return m.FeatureSet
	}
	return nil
}

type ApplyFeatureSetResponse struct {
	// Feature set response has been enriched with version and source information
	FeatureSet           *FeatureSetSpec                `protobuf:"bytes,1,opt,name=featureSet,proto3" json:"featureSet,omitempty"`
	Status               ApplyFeatureSetResponse_Status `protobuf:"varint,2,opt,name=status,proto3,enum=feast.core.ApplyFeatureSetResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ApplyFeatureSetResponse) Reset()         { *m = ApplyFeatureSetResponse{} }
func (m *ApplyFeatureSetResponse) String() string { return proto.CompactTextString(m) }
func (*ApplyFeatureSetResponse) ProtoMessage()    {}
func (*ApplyFeatureSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9be266444105411, []int{5}
}

func (m *ApplyFeatureSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyFeatureSetResponse.Unmarshal(m, b)
}
func (m *ApplyFeatureSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyFeatureSetResponse.Marshal(b, m, deterministic)
}
func (m *ApplyFeatureSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyFeatureSetResponse.Merge(m, src)
}
func (m *ApplyFeatureSetResponse) XXX_Size() int {
	return xxx_messageInfo_ApplyFeatureSetResponse.Size(m)
}
func (m *ApplyFeatureSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyFeatureSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyFeatureSetResponse proto.InternalMessageInfo

func (m *ApplyFeatureSetResponse) GetFeatureSet() *FeatureSetSpec {
	if m != nil {
		return m.FeatureSet
	}
	return nil
}

func (m *ApplyFeatureSetResponse) GetStatus() ApplyFeatureSetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ApplyFeatureSetResponse_NO_CHANGE
}

type GetFeastCoreVersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeastCoreVersionResponse) Reset()         { *m = GetFeastCoreVersionResponse{} }
func (m *GetFeastCoreVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetFeastCoreVersionResponse) ProtoMessage()    {}
func (*GetFeastCoreVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9be266444105411, []int{6}
}

func (m *GetFeastCoreVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeastCoreVersionResponse.Unmarshal(m, b)
}
func (m *GetFeastCoreVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeastCoreVersionResponse.Marshal(b, m, deterministic)
}
func (m *GetFeastCoreVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeastCoreVersionResponse.Merge(m, src)
}
func (m *GetFeastCoreVersionResponse) XXX_Size() int {
	return xxx_messageInfo_GetFeastCoreVersionResponse.Size(m)
}
func (m *GetFeastCoreVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeastCoreVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeastCoreVersionResponse proto.InternalMessageInfo

func (m *GetFeastCoreVersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterEnum("feast.core.ApplyFeatureSetResponse_Status", ApplyFeatureSetResponse_Status_name, ApplyFeatureSetResponse_Status_value)
	proto.RegisterType((*GetFeatureSetsRequest)(nil), "feast.core.GetFeatureSetsRequest")
	proto.RegisterType((*GetFeatureSetsRequest_Filter)(nil), "feast.core.GetFeatureSetsRequest.Filter")
	proto.RegisterType((*GetFeatureSetsResponse)(nil), "feast.core.GetFeatureSetsResponse")
	proto.RegisterType((*GetStoresRequest)(nil), "feast.core.GetStoresRequest")
	proto.RegisterType((*GetStoresRequest_Filter)(nil), "feast.core.GetStoresRequest.Filter")
	proto.RegisterType((*GetStoresResponse)(nil), "feast.core.GetStoresResponse")
	proto.RegisterType((*ApplyFeatureSetRequest)(nil), "feast.core.ApplyFeatureSetRequest")
	proto.RegisterType((*ApplyFeatureSetResponse)(nil), "feast.core.ApplyFeatureSetResponse")
	proto.RegisterType((*GetFeastCoreVersionResponse)(nil), "feast.core.GetFeastCoreVersionResponse")
}

func init() { proto.RegisterFile("feast/core/CoreService.proto", fileDescriptor_d9be266444105411) }

var fileDescriptor_d9be266444105411 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xe1, 0x6e, 0x12, 0x41,
	0x10, 0xf6, 0xd0, 0xd2, 0x30, 0x44, 0x84, 0x35, 0x22, 0x39, 0x30, 0xa9, 0xd7, 0xc4, 0x12, 0x63,
	0x6e, 0x13, 0x8c, 0x31, 0xd1, 0xfe, 0x10, 0x90, 0xd2, 0x5f, 0xd4, 0x2c, 0x15, 0x93, 0xfe, 0xd0,
	0x00, 0x0e, 0x27, 0x0a, 0xec, 0x79, 0xbb, 0x34, 0xe9, 0xa3, 0xf9, 0x0a, 0x3e, 0x87, 0x0f, 0x62,
	0x6e, 0x77, 0xe1, 0xb6, 0x07, 0x45, 0x93, 0xfe, 0x83, 0xf9, 0xbe, 0xf9, 0x6e, 0xe6, 0x9b, 0x99,
	0x85, 0xda, 0x04, 0x87, 0x42, 0xd2, 0x31, 0x8f, 0x90, 0xb6, 0x79, 0x84, 0x7d, 0x8c, 0x2e, 0xa7,
	0x63, 0xf4, 0xc3, 0x88, 0x4b, 0x4e, 0x40, 0xa1, 0x7e, 0x8c, 0xba, 0x55, 0x8b, 0x79, 0x82, 0x43,
	0xb9, 0x8c, 0xc9, 0x52, 0x13, 0xdd, 0xb2, 0x05, 0xf6, 0x25, 0x8f, 0x8c, 0x80, 0x5b, 0x0d, 0x38,
	0x0f, 0x66, 0x48, 0xd5, 0xbf, 0xd1, 0x72, 0x42, 0x71, 0x1e, 0xca, 0x2b, 0x0d, 0x7a, 0xbf, 0x1c,
	0x78, 0xd4, 0x45, 0x99, 0x88, 0x09, 0x86, 0x3f, 0x97, 0x28, 0x24, 0x79, 0x07, 0xd9, 0xc9, 0x74,
	0x26, 0x31, 0xaa, 0x38, 0x07, 0x4e, 0x3d, 0xdf, 0xa8, 0xfb, 0x49, 0x21, 0xfe, 0xd6, 0x14, 0xff,
	0x44, 0xf1, 0x99, 0xc9, 0x73, 0x3f, 0x43, 0x56, 0x47, 0xc8, 0x33, 0x28, 0x4c, 0xd6, 0xf4, 0xde,
	0x70, 0x8e, 0x4a, 0x33, 0xc7, 0x52, 0x51, 0xf2, 0x02, 0x4a, 0x49, 0x64, 0x80, 0x91, 0x98, 0xf2,
	0x45, 0x25, 0xa3, 0xa8, 0x9b, 0x80, 0x37, 0x80, 0x72, 0xba, 0x0e, 0x11, 0xf2, 0x85, 0x40, 0x72,
	0x0c, 0xf9, 0x84, 0x2e, 0x2a, 0xce, 0xc1, 0xdd, 0x7a, 0xbe, 0xe1, 0xda, 0x0d, 0x24, 0x59, 0xfd,
	0x10, 0xc7, 0xcc, 0xa6, 0x7b, 0x73, 0x28, 0x76, 0x51, 0x2a, 0x0b, 0xd7, 0x6e, 0xbc, 0x4d, 0xb9,
	0x71, 0x98, 0x72, 0xe3, 0x1a, 0x3b, 0x6d, 0x44, 0x6d, 0x6d, 0x04, 0x81, 0x7b, 0x8b, 0xa4, 0x7d,
	0xf5, 0xdb, 0x3b, 0x86, 0x92, 0x25, 0x60, 0x3a, 0x38, 0x82, 0x3d, 0x11, 0x47, 0x4c, 0xed, 0x25,
	0xfb, 0x73, 0x8a, 0xca, 0x34, 0xee, 0x9d, 0x43, 0xb9, 0x19, 0x86, 0xb3, 0xab, 0xa4, 0xa1, 0x55,
	0xc9, 0x6f, 0x00, 0x92, 0xae, 0x4c, 0xd9, 0xbb, 0x3c, 0xb0, 0xd8, 0xde, 0x6f, 0x07, 0x1e, 0x6f,
	0xc8, 0x9a, 0xd2, 0x6e, 0xa1, 0x4b, 0x5a, 0x90, 0x15, 0x72, 0x28, 0x97, 0x42, 0x4d, 0xb5, 0xd0,
	0x78, 0x6e, 0xe7, 0xdd, 0xf0, 0x41, 0xbf, 0xaf, 0x32, 0x98, 0xc9, 0xf4, 0x28, 0x64, 0x75, 0x84,
	0xdc, 0x87, 0x5c, 0xef, 0xec, 0x4b, 0xfb, 0xb4, 0xd9, 0xeb, 0x76, 0x8a, 0x77, 0x48, 0x1e, 0xf6,
	0xdb, 0xac, 0xd3, 0x3c, 0xef, 0xbc, 0x2f, 0x3a, 0x24, 0x07, 0x7b, 0x1d, 0xc6, 0xce, 0x58, 0x31,
	0xe3, 0xbd, 0x86, 0xaa, 0xde, 0x13, 0x21, 0xe3, 0xf3, 0x32, 0xeb, 0xb3, 0xee, 0xa7, 0x02, 0xfb,
	0x97, 0x66, 0xd5, 0xf4, 0x58, 0x56, 0x7f, 0x1b, 0x7f, 0x32, 0x90, 0xb7, 0x0e, 0x92, 0x0c, 0xe0,
	0xe1, 0x16, 0x21, 0x52, 0xf6, 0xf5, 0x85, 0xf9, 0xab, 0x0b, 0xf3, 0x3b, 0xf1, 0x85, 0xb9, 0x47,
	0x9b, 0x17, 0xb3, 0xbd, 0x82, 0x4f, 0x50, 0xb8, 0xbe, 0xc8, 0xe4, 0xe9, 0x3f, 0x8f, 0xcd, 0xf5,
	0x76, 0x51, 0x8c, 0xf0, 0x29, 0xe4, 0xd6, 0xab, 0x45, 0x6a, 0xbb, 0x56, 0xd6, 0x7d, 0x72, 0x03,
	0x6a, 0x94, 0x2e, 0xe0, 0x41, 0x6a, 0x3c, 0xc4, 0xdb, 0x39, 0x3b, 0xad, 0x7a, 0xf8, 0x1f, 0xf3,
	0x6d, 0x7d, 0x04, 0xeb, 0x8d, 0x6b, 0x15, 0x2d, 0xc7, 0x3f, 0xc4, 0x86, 0x5e, 0xbc, 0x0a, 0xa6,
	0xf2, 0xdb, 0x72, 0xe4, 0x8f, 0xf9, 0x9c, 0x06, 0xfc, 0x3b, 0xfe, 0xa0, 0xfa, 0xa5, 0x53, 0x76,
	0x0b, 0x1a, 0xe0, 0x02, 0xa3, 0xa1, 0xc4, 0xaf, 0x34, 0xe0, 0x34, 0x79, 0x03, 0x47, 0x59, 0x85,
	0xbf, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x3f, 0x4e, 0x5b, 0x5f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoreServiceClient is the client API for CoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreServiceClient interface {
	// Retrieve version information about this Feast deployment
	GetFeastCoreVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetFeastCoreVersionResponse, error)
	// Retrieve feature set details given a filter.
	// Returns all featureSets matching that filter.
	GetFeatureSets(ctx context.Context, in *GetFeatureSetsRequest, opts ...grpc.CallOption) (*GetFeatureSetsResponse, error)
	// Retrieve store details given a filter.
	// Returns all stores matching that filter.
	GetStores(ctx context.Context, in *GetStoresRequest, opts ...grpc.CallOption) (*GetStoresResponse, error)
	// Idempotent creation of feature set. Will not create a new feature set if schema does not change
	ApplyFeatureSet(ctx context.Context, in *ApplyFeatureSetRequest, opts ...grpc.CallOption) (*ApplyFeatureSetResponse, error)
}

type coreServiceClient struct {
	cc *grpc.ClientConn
}

func NewCoreServiceClient(cc *grpc.ClientConn) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) GetFeastCoreVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetFeastCoreVersionResponse, error) {
	out := new(GetFeastCoreVersionResponse)
	err := c.cc.Invoke(ctx, "/feast.core.CoreService/GetFeastCoreVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetFeatureSets(ctx context.Context, in *GetFeatureSetsRequest, opts ...grpc.CallOption) (*GetFeatureSetsResponse, error) {
	out := new(GetFeatureSetsResponse)
	err := c.cc.Invoke(ctx, "/feast.core.CoreService/GetFeatureSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) GetStores(ctx context.Context, in *GetStoresRequest, opts ...grpc.CallOption) (*GetStoresResponse, error) {
	out := new(GetStoresResponse)
	err := c.cc.Invoke(ctx, "/feast.core.CoreService/GetStores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) ApplyFeatureSet(ctx context.Context, in *ApplyFeatureSetRequest, opts ...grpc.CallOption) (*ApplyFeatureSetResponse, error) {
	out := new(ApplyFeatureSetResponse)
	err := c.cc.Invoke(ctx, "/feast.core.CoreService/ApplyFeatureSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServiceServer is the server API for CoreService service.
type CoreServiceServer interface {
	// Retrieve version information about this Feast deployment
	GetFeastCoreVersion(context.Context, *empty.Empty) (*GetFeastCoreVersionResponse, error)
	// Retrieve feature set details given a filter.
	// Returns all featureSets matching that filter.
	GetFeatureSets(context.Context, *GetFeatureSetsRequest) (*GetFeatureSetsResponse, error)
	// Retrieve store details given a filter.
	// Returns all stores matching that filter.
	GetStores(context.Context, *GetStoresRequest) (*GetStoresResponse, error)
	// Idempotent creation of feature set. Will not create a new feature set if schema does not change
	ApplyFeatureSet(context.Context, *ApplyFeatureSetRequest) (*ApplyFeatureSetResponse, error)
}

// UnimplementedCoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCoreServiceServer struct {
}

func (*UnimplementedCoreServiceServer) GetFeastCoreVersion(ctx context.Context, req *empty.Empty) (*GetFeastCoreVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeastCoreVersion not implemented")
}
func (*UnimplementedCoreServiceServer) GetFeatureSets(ctx context.Context, req *GetFeatureSetsRequest) (*GetFeatureSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeatureSets not implemented")
}
func (*UnimplementedCoreServiceServer) GetStores(ctx context.Context, req *GetStoresRequest) (*GetStoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStores not implemented")
}
func (*UnimplementedCoreServiceServer) ApplyFeatureSet(ctx context.Context, req *ApplyFeatureSetRequest) (*ApplyFeatureSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyFeatureSet not implemented")
}

func RegisterCoreServiceServer(s *grpc.Server, srv CoreServiceServer) {
	s.RegisterService(&_CoreService_serviceDesc, srv)
}

func _CoreService_GetFeastCoreVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetFeastCoreVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.core.CoreService/GetFeastCoreVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetFeastCoreVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetFeatureSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeatureSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetFeatureSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.core.CoreService/GetFeatureSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetFeatureSets(ctx, req.(*GetFeatureSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_GetStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).GetStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.core.CoreService/GetStores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).GetStores(ctx, req.(*GetStoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_ApplyFeatureSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyFeatureSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).ApplyFeatureSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.core.CoreService/ApplyFeatureSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).ApplyFeatureSet(ctx, req.(*ApplyFeatureSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feast.core.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeastCoreVersion",
			Handler:    _CoreService_GetFeastCoreVersion_Handler,
		},
		{
			MethodName: "GetFeatureSets",
			Handler:    _CoreService_GetFeatureSets_Handler,
		},
		{
			MethodName: "GetStores",
			Handler:    _CoreService_GetStores_Handler,
		},
		{
			MethodName: "ApplyFeatureSet",
			Handler:    _CoreService_ApplyFeatureSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feast/core/CoreService.proto",
}
