// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/types/FeatureRow.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type FeatureRow struct {
	Features             []*Feature             `protobuf:"bytes,2,rep,name=features,proto3" json:"features,omitempty"`
	EventTimestamp       *timestamp.Timestamp   `protobuf:"bytes,3,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
	FeatureSet           *FeatureRow_FeatureSet `protobuf:"bytes,4,opt,name=featureSet,proto3" json:"featureSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FeatureRow) Reset()         { *m = FeatureRow{} }
func (m *FeatureRow) String() string { return proto.CompactTextString(m) }
func (*FeatureRow) ProtoMessage()    {}
func (*FeatureRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbbea9c89787d1c7, []int{0}
}

func (m *FeatureRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureRow.Unmarshal(m, b)
}
func (m *FeatureRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureRow.Marshal(b, m, deterministic)
}
func (m *FeatureRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureRow.Merge(m, src)
}
func (m *FeatureRow) XXX_Size() int {
	return xxx_messageInfo_FeatureRow.Size(m)
}
func (m *FeatureRow) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureRow.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureRow proto.InternalMessageInfo

func (m *FeatureRow) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *FeatureRow) GetEventTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.EventTimestamp
	}
	return nil
}

func (m *FeatureRow) GetFeatureSet() *FeatureRow_FeatureSet {
	if m != nil {
		return m.FeatureSet
	}
	return nil
}

type FeatureRow_FeatureSet struct {
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Version              int32    `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeatureRow_FeatureSet) Reset()         { *m = FeatureRow_FeatureSet{} }
func (m *FeatureRow_FeatureSet) String() string { return proto.CompactTextString(m) }
func (*FeatureRow_FeatureSet) ProtoMessage()    {}
func (*FeatureRow_FeatureSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbbea9c89787d1c7, []int{0, 0}
}

func (m *FeatureRow_FeatureSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureRow_FeatureSet.Unmarshal(m, b)
}
func (m *FeatureRow_FeatureSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureRow_FeatureSet.Marshal(b, m, deterministic)
}
func (m *FeatureRow_FeatureSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureRow_FeatureSet.Merge(m, src)
}
func (m *FeatureRow_FeatureSet) XXX_Size() int {
	return xxx_messageInfo_FeatureRow_FeatureSet.Size(m)
}
func (m *FeatureRow_FeatureSet) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureRow_FeatureSet.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureRow_FeatureSet proto.InternalMessageInfo

func (m *FeatureRow_FeatureSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeatureRow_FeatureSet) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*FeatureRow)(nil), "feast.types.FeatureRow")
	proto.RegisterType((*FeatureRow_FeatureSet)(nil), "feast.types.FeatureRow.FeatureSet")
}

func init() { proto.RegisterFile("feast/types/FeatureRow.proto", fileDescriptor_fbbea9c89787d1c7) }

var fileDescriptor_fbbea9c89787d1c7 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x3d, 0x4f, 0x84, 0x40,
	0x10, 0x0d, 0xf7, 0xa5, 0x0e, 0x89, 0x26, 0x1b, 0x8b, 0x95, 0x98, 0x48, 0xae, 0xa2, 0xda, 0x35,
	0x67, 0x62, 0x61, 0x49, 0x61, 0x6d, 0xc0, 0x58, 0xd8, 0x81, 0x0e, 0x2b, 0x2a, 0x0c, 0x61, 0x87,
	0xbb, 0xd8, 0xfa, 0xcb, 0x8d, 0x8b, 0x70, 0xc4, 0x5c, 0x37, 0x93, 0xf7, 0xe6, 0xbd, 0x37, 0x0f,
	0x2e, 0x0b, 0xcc, 0x2c, 0x6b, 0xfe, 0x6a, 0xd0, 0xea, 0x7b, 0xcc, 0xb8, 0x6b, 0x31, 0xa1, 0x9d,
	0x6a, 0x5a, 0x62, 0x12, 0xbe, 0x43, 0x95, 0x43, 0x83, 0x2b, 0x43, 0x64, 0x3e, 0x51, 0x3b, 0x28,
	0xef, 0x0a, 0xcd, 0x65, 0x85, 0x96, 0xb3, 0xaa, 0xe9, 0xd9, 0xc1, 0xc5, 0x01, 0xad, 0x1e, 0x5a,
	0x7f, 0xcf, 0x00, 0xf6, 0xea, 0xe2, 0x1a, 0x8e, 0x8b, 0x7e, 0xb3, 0x72, 0x16, 0xce, 0x23, 0x7f,
	0x73, 0xae, 0x26, 0x56, 0x6a, 0xa0, 0x8e, 0x2c, 0x11, 0xc3, 0x29, 0x6e, 0xb1, 0xe6, 0xc7, 0xc1,
	0x53, 0xce, 0x43, 0x2f, 0xf2, 0x37, 0x81, 0xea, 0x53, 0xa9, 0x21, 0x95, 0x1a, 0x19, 0xc9, 0xbf,
	0x0b, 0x11, 0x03, 0xfc, 0xe9, 0xa5, 0xc8, 0x72, 0xe1, 0xee, 0xd7, 0x07, 0x7d, 0x69, 0x37, 0x8c,
	0x29, 0x72, 0x32, 0xb9, 0x0a, 0xee, 0xc6, 0x3f, 0x52, 0x64, 0x21, 0x60, 0x51, 0x67, 0x15, 0xca,
	0x65, 0xe8, 0x45, 0x27, 0x89, 0x9b, 0x85, 0x84, 0xa3, 0x2d, 0xb6, 0xb6, 0xa4, 0x5a, 0xae, 0x42,
	0x2f, 0x5a, 0x26, 0xc3, 0x1a, 0x3f, 0xc1, 0xb4, 0xcf, 0xf8, 0x6c, 0xef, 0xf6, 0xf0, 0x1b, 0xfe,
	0xf9, 0xd6, 0x94, 0xfc, 0xd6, 0xe5, 0xea, 0x85, 0x2a, 0x6d, 0xe8, 0x1d, 0x3f, 0x74, 0x5f, 0xa8,
	0x7b, 0xcd, 0x6a, 0x83, 0x35, 0xb6, 0x19, 0xe3, 0xab, 0x36, 0xa4, 0x27, 0x55, 0xe7, 0x2b, 0x47,
	0xb8, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x17, 0x1b, 0x50, 0xcc, 0x01, 0x00, 0x00,
}
