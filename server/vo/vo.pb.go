// Code generated by protoc-gen-go.
// source: vo.proto
// DO NOT EDIT!

/*
Package vo is a generated protocol buffer package.

It is generated from these files:
	vo.proto

It has these top-level messages:
	DatasetMetadataDelta
	Sample
*/
package vo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ACTIVITY_LEVEL int32

const (
	ACTIVITY_LEVEL_ACTIVITY_VERBOSE ACTIVITY_LEVEL = 0
	ACTIVITY_LEVEL_ACTIVITY_INFO    ACTIVITY_LEVEL = 1
	ACTIVITY_LEVEL_ACTIVITY_PROBLEM ACTIVITY_LEVEL = 2
)

var ACTIVITY_LEVEL_name = map[int32]string{
	0: "ACTIVITY_VERBOSE",
	1: "ACTIVITY_INFO",
	2: "ACTIVITY_PROBLEM",
}
var ACTIVITY_LEVEL_value = map[string]int32{
	"ACTIVITY_VERBOSE": 0,
	"ACTIVITY_INFO":    1,
	"ACTIVITY_PROBLEM": 2,
}

func (x ACTIVITY_LEVEL) String() string {
	return proto.EnumName(ACTIVITY_LEVEL_name, int32(x))
}
func (ACTIVITY_LEVEL) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Sample storage format. Doesn't necessarily match
// actual data format which could be very cryptic
type Sample_FORMAT int32

const (
	Sample_NONE Sample_FORMAT = 0
	Sample_TEXT Sample_FORMAT = 1
	Sample_TSV  Sample_FORMAT = 2
	Sample_JSON Sample_FORMAT = 3
)

var Sample_FORMAT_name = map[int32]string{
	0: "NONE",
	1: "TEXT",
	2: "TSV",
	3: "JSON",
}
var Sample_FORMAT_value = map[string]int32{
	"NONE": 0,
	"TEXT": 1,
	"TSV":  2,
	"JSON": 3,
}

func (x Sample_FORMAT) String() string {
	return proto.EnumName(Sample_FORMAT_name, int32(x))
}
func (Sample_FORMAT) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// if a field kind is set, then the corresponding property carries
// a new value (which could be null)
type DatasetMetadataDelta struct {
	Sample         *Sample `protobuf:"bytes,1,opt,name=sample" json:"sample,omitempty"`
	SampleSet      bool    `protobuf:"varint,2,opt,name=sample_set,json=sampleSet" json:"sample_set,omitempty"`
	Description    string  `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	DescriptionSet bool    `protobuf:"varint,6,opt,name=description_set,json=descriptionSet" json:"description_set,omitempty"`
	Summary        string  `protobuf:"bytes,7,opt,name=summary" json:"summary,omitempty"`
	SummarySet     bool    `protobuf:"varint,8,opt,name=summary_set,json=summarySet" json:"summary_set,omitempty"`
}

func (m *DatasetMetadataDelta) Reset()                    { *m = DatasetMetadataDelta{} }
func (m *DatasetMetadataDelta) String() string            { return proto.CompactTextString(m) }
func (*DatasetMetadataDelta) ProtoMessage()               {}
func (*DatasetMetadataDelta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DatasetMetadataDelta) GetSample() *Sample {
	if m != nil {
		return m.Sample
	}
	return nil
}

type Sample struct {
	Format Sample_FORMAT `protobuf:"varint,1,opt,name=format,enum=Sample_FORMAT" json:"format,omitempty"`
	Body   string        `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (m *Sample) String() string            { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*DatasetMetadataDelta)(nil), "DatasetMetadataDelta")
	proto.RegisterType((*Sample)(nil), "Sample")
	proto.RegisterEnum("ACTIVITY_LEVEL", ACTIVITY_LEVEL_name, ACTIVITY_LEVEL_value)
	proto.RegisterEnum("Sample_FORMAT", Sample_FORMAT_name, Sample_FORMAT_value)
}

func init() { proto.RegisterFile("vo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0x51, 0x6b, 0xc2, 0x30,
	0x14, 0x85, 0x17, 0x75, 0x6d, 0xbd, 0xb2, 0x2e, 0x0b, 0x3e, 0xf4, 0x65, 0x58, 0x7c, 0xd8, 0x64,
	0x0f, 0x0e, 0xdc, 0x2f, 0xd0, 0x19, 0xc1, 0xa1, 0xcd, 0x48, 0x42, 0xd9, 0x9e, 0x24, 0xce, 0x0c,
	0x04, 0xbb, 0x48, 0x1b, 0x07, 0xfe, 0xd7, 0xfd, 0x98, 0xd1, 0xb4, 0x13, 0xf7, 0x76, 0xce, 0x97,
	0x7b, 0x0f, 0x39, 0x5c, 0x08, 0xbe, 0xcd, 0x70, 0x9f, 0x1b, 0x6b, 0xfa, 0x3f, 0x08, 0xba, 0x53,
	0x65, 0x55, 0xa1, 0xed, 0x52, 0x5b, 0xb5, 0x51, 0x56, 0x4d, 0xf5, 0xce, 0x2a, 0xd2, 0x03, 0xaf,
	0x50, 0xd9, 0x7e, 0xa7, 0x23, 0x14, 0xa3, 0x41, 0x67, 0xe4, 0x0f, 0x85, 0xb3, 0xbc, 0xc6, 0xe4,
	0x16, 0xa0, 0x52, 0xab, 0x42, 0xdb, 0xa8, 0x11, 0xa3, 0x41, 0xc0, 0xdb, 0x15, 0x11, 0xda, 0x92,
	0x18, 0x3a, 0x1b, 0x5d, 0x7c, 0xe4, 0xdb, 0xbd, 0xdd, 0x9a, 0xaf, 0xe8, 0x32, 0x46, 0x83, 0x36,
	0x3f, 0x47, 0xe4, 0x1e, 0xae, 0xcf, 0xac, 0x4b, 0xf1, 0x5c, 0x4a, 0x78, 0x86, 0xcb, 0xa8, 0x08,
	0xfc, 0xe2, 0x90, 0x65, 0x2a, 0x3f, 0x46, 0xbe, 0x8b, 0xf9, 0xb3, 0xa4, 0x07, 0x9d, 0x5a, 0xba,
	0xf5, 0xc0, 0xad, 0x43, 0x8d, 0x84, 0xb6, 0xfd, 0x03, 0x78, 0xd5, 0xb7, 0xc9, 0x1d, 0x78, 0x9f,
	0x26, 0xcf, 0x94, 0x75, 0x7d, 0xc2, 0x51, 0x58, 0xf7, 0x19, 0xce, 0x18, 0x5f, 0x8e, 0x25, 0xaf,
	0x5f, 0x09, 0x81, 0xd6, 0xda, 0x6c, 0x8e, 0xae, 0x50, 0x9b, 0x3b, 0xdd, 0x7f, 0x04, 0xaf, 0x9a,
	0x22, 0x01, 0xb4, 0x12, 0x96, 0x50, 0x7c, 0x51, 0x2a, 0x49, 0xdf, 0x24, 0x46, 0xc4, 0x87, 0xa6,
	0x14, 0x29, 0x6e, 0x94, 0xe8, 0x45, 0xb0, 0x04, 0x37, 0x1f, 0x18, 0x84, 0xe3, 0x67, 0x39, 0x4f,
	0xe7, 0xf2, 0x7d, 0xb5, 0xa0, 0x29, 0x5d, 0x90, 0x2e, 0xe0, 0x13, 0x49, 0x29, 0x9f, 0x30, 0x51,
	0x86, 0xdc, 0xc0, 0xd5, 0x89, 0xce, 0x93, 0x19, 0xc3, 0xe8, 0xdf, 0xe0, 0x2b, 0x67, 0x93, 0x05,
	0x5d, 0xe2, 0xc6, 0xda, 0x73, 0xd7, 0x7a, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xac, 0x52, 0xfb,
	0x48, 0xb9, 0x01, 0x00, 0x00,
}
