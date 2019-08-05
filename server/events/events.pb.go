// Code generated by protoc-gen-go.
// source: events.proto
// DO NOT EDIT!

/*
Package events is a generated protocol buffer package.

It is generated from these files:
	events.proto

It has these top-level messages:
	ProjectCreated
	DatasetCreated
	DatasetUpdated
	JobAdded
	ExpertAdded
	DatasetMetadata
	Event
*/
package events

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FIELD_TYPES int32

const (
	FIELD_TYPES_FIELD_NONE             FIELD_TYPES = 0
	FIELD_TYPES_FIELD_RECORD_COUNT     FIELD_TYPES = 1
	FIELD_TYPES_FIELD_RAW_BYTES        FIELD_TYPES = 2
	FIELD_TYPES_FIELD_ZIP_BYTES        FIELD_TYPES = 3
	FIELD_TYPES_FIELD_SAMPLE           FIELD_TYPES = 4
	FIELD_TYPES_FIELD_UPDATE_TIMESTAMP FIELD_TYPES = 5
	FIELD_TYPES_FIELD_FILE_COUNT       FIELD_TYPES = 6
	FIELD_TYPES_FIELD_DATA_FORMAT      FIELD_TYPES = 7
	FIELD_TYPES_FIELD_DESCRIPTION      FIELD_TYPES = 8
	FIELD_TYPES_FIELD_LOCATION_ID      FIELD_TYPES = 9
	FIELD_TYPES_FIELD_LOCATION_URI     FIELD_TYPES = 10
)

var FIELD_TYPES_name = map[int32]string{
	0:  "FIELD_NONE",
	1:  "FIELD_RECORD_COUNT",
	2:  "FIELD_RAW_BYTES",
	3:  "FIELD_ZIP_BYTES",
	4:  "FIELD_SAMPLE",
	5:  "FIELD_UPDATE_TIMESTAMP",
	6:  "FIELD_FILE_COUNT",
	7:  "FIELD_DATA_FORMAT",
	8:  "FIELD_DESCRIPTION",
	9:  "FIELD_LOCATION_ID",
	10: "FIELD_LOCATION_URI",
}
var FIELD_TYPES_value = map[string]int32{
	"FIELD_NONE":             0,
	"FIELD_RECORD_COUNT":     1,
	"FIELD_RAW_BYTES":        2,
	"FIELD_ZIP_BYTES":        3,
	"FIELD_SAMPLE":           4,
	"FIELD_UPDATE_TIMESTAMP": 5,
	"FIELD_FILE_COUNT":       6,
	"FIELD_DATA_FORMAT":      7,
	"FIELD_DESCRIPTION":      8,
	"FIELD_LOCATION_ID":      9,
	"FIELD_LOCATION_URI":     10,
}

func (x FIELD_TYPES) String() string {
	return proto.EnumName(FIELD_TYPES_name, int32(x))
}
func (FIELD_TYPES) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Sample storage format. Doesn't necessarily match
// actual data format which could be very cryptic
type SAMPLE_KIND int32

const (
	SAMPLE_KIND_TEXT SAMPLE_KIND = 0
	SAMPLE_KIND_TSV  SAMPLE_KIND = 1
	SAMPLE_KIND_JSON SAMPLE_KIND = 2
)

var SAMPLE_KIND_name = map[int32]string{
	0: "TEXT",
	1: "TSV",
	2: "JSON",
}
var SAMPLE_KIND_value = map[string]int32{
	"TEXT": 0,
	"TSV":  1,
	"JSON": 2,
}

func (x SAMPLE_KIND) String() string {
	return proto.EnumName(SAMPLE_KIND_name, int32(x))
}
func (SAMPLE_KIND) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Event_Types int32

const (
	Event_None           Event_Types = 0
	Event_ProjectCreated Event_Types = 1
	Event_DatasetCreated Event_Types = 2
	Event_DatasetUpdated Event_Types = 3
	Event_ExpertAdded    Event_Types = 4
	Event_JobAdded       Event_Types = 5
)

var Event_Types_name = map[int32]string{
	0: "None",
	1: "ProjectCreated",
	2: "DatasetCreated",
	3: "DatasetUpdated",
	4: "ExpertAdded",
	5: "JobAdded",
}
var Event_Types_value = map[string]int32{
	"None":           0,
	"ProjectCreated": 1,
	"DatasetCreated": 2,
	"DatasetUpdated": 3,
	"ExpertAdded":    4,
	"JobAdded":       5,
}

func (x Event_Types) String() string {
	return proto.EnumName(Event_Types_name, int32(x))
}
func (Event_Types) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

type ProjectCreated struct {
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ProjectCreated) Reset()                    { *m = ProjectCreated{} }
func (m *ProjectCreated) String() string            { return proto.CompactTextString(m) }
func (*ProjectCreated) ProtoMessage()               {}
func (*ProjectCreated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DatasetCreated struct {
	DatasetId string           `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
	Name      string           `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ProjectId string           `protobuf:"bytes,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Metadata  *DatasetMetadata `protobuf:"bytes,4,opt,name=metadata" json:"metadata,omitempty"`
	Experts   []string         `protobuf:"bytes,5,rep,name=experts" json:"experts,omitempty"`
}

func (m *DatasetCreated) Reset()                    { *m = DatasetCreated{} }
func (m *DatasetCreated) String() string            { return proto.CompactTextString(m) }
func (*DatasetCreated) ProtoMessage()               {}
func (*DatasetCreated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DatasetCreated) GetMetadata() *DatasetMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type DatasetUpdated struct {
	DatasetId string           `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
	ProjectId string           `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Metadata  *DatasetMetadata `protobuf:"bytes,4,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *DatasetUpdated) Reset()                    { *m = DatasetUpdated{} }
func (m *DatasetUpdated) String() string            { return proto.CompactTextString(m) }
func (*DatasetUpdated) ProtoMessage()               {}
func (*DatasetUpdated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DatasetUpdated) GetMetadata() *DatasetMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type JobAdded struct {
	JobId     string   `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	JobName   string   `protobuf:"bytes,2,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	Inputs    []string `protobuf:"bytes,3,rep,name=inputs" json:"inputs,omitempty"`
	Outputs   []string `protobuf:"bytes,4,rep,name=outputs" json:"outputs,omitempty"`
	ProjectId string   `protobuf:"bytes,5,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Experts   []string `protobuf:"bytes,6,rep,name=experts" json:"experts,omitempty"`
}

func (m *JobAdded) Reset()                    { *m = JobAdded{} }
func (m *JobAdded) String() string            { return proto.CompactTextString(m) }
func (*JobAdded) ProtoMessage()               {}
func (*JobAdded) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ExpertAdded struct {
	ExpertId   string `protobuf:"bytes,1,opt,name=expert_id,json=expertId" json:"expert_id,omitempty"`
	ExpertName string `protobuf:"bytes,2,opt,name=expert_name,json=expertName" json:"expert_name,omitempty"`
}

func (m *ExpertAdded) Reset()                    { *m = ExpertAdded{} }
func (m *ExpertAdded) String() string            { return proto.CompactTextString(m) }
func (*ExpertAdded) ProtoMessage()               {}
func (*ExpertAdded) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// if a field kind is set, then the corresponding property carries
// a new value (which could be null)
type DatasetMetadata struct {
	SetFields       []FIELD_TYPES `protobuf:"varint,1,rep,packed,name=set_fields,json=setFields,enum=FIELD_TYPES" json:"set_fields,omitempty"`
	DelFields       []FIELD_TYPES `protobuf:"varint,2,rep,packed,name=del_fields,json=delFields,enum=FIELD_TYPES" json:"del_fields,omitempty"`
	RecordCount     int64         `protobuf:"varint,3,opt,name=record_count,json=recordCount" json:"record_count,omitempty"`
	FileCount       int64         `protobuf:"varint,4,opt,name=file_count,json=fileCount" json:"file_count,omitempty"`
	RawBytes        int64         `protobuf:"varint,5,opt,name=raw_bytes,json=rawBytes" json:"raw_bytes,omitempty"`
	ZipBytes        int64         `protobuf:"varint,6,opt,name=zip_bytes,json=zipBytes" json:"zip_bytes,omitempty"`
	SampleBody      []byte        `protobuf:"bytes,7,opt,name=sample_body,json=sampleBody,proto3" json:"sample_body,omitempty"`
	SampleKind      SAMPLE_KIND   `protobuf:"varint,8,opt,name=sample_kind,json=sampleKind,enum=SAMPLE_KIND" json:"sample_kind,omitempty"`
	UpdateTimestamp int64         `protobuf:"varint,9,opt,name=update_timestamp,json=updateTimestamp" json:"update_timestamp,omitempty"`
	DataFormat      string        `protobuf:"bytes,10,opt,name=data_format,json=dataFormat" json:"data_format,omitempty"`
	Description     string        `protobuf:"bytes,11,opt,name=description" json:"description,omitempty"`
	LocationId      string        `protobuf:"bytes,12,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	LocationUri     string        `protobuf:"bytes,13,opt,name=location_uri,json=locationUri" json:"location_uri,omitempty"`
}

func (m *DatasetMetadata) Reset()                    { *m = DatasetMetadata{} }
func (m *DatasetMetadata) String() string            { return proto.CompactTextString(m) }
func (*DatasetMetadata) ProtoMessage()               {}
func (*DatasetMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Event struct {
	Body []byte      `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
	Type Event_Types `protobuf:"varint,2,opt,name=Type,enum=Event_Types" json:"Type,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*ProjectCreated)(nil), "ProjectCreated")
	proto.RegisterType((*DatasetCreated)(nil), "DatasetCreated")
	proto.RegisterType((*DatasetUpdated)(nil), "DatasetUpdated")
	proto.RegisterType((*JobAdded)(nil), "JobAdded")
	proto.RegisterType((*ExpertAdded)(nil), "ExpertAdded")
	proto.RegisterType((*DatasetMetadata)(nil), "DatasetMetadata")
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterEnum("FIELD_TYPES", FIELD_TYPES_name, FIELD_TYPES_value)
	proto.RegisterEnum("SAMPLE_KIND", SAMPLE_KIND_name, SAMPLE_KIND_value)
	proto.RegisterEnum("Event_Types", Event_Types_name, Event_Types_value)
}

func init() { proto.RegisterFile("events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 809 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x95, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0xeb, 0x38, 0x5f, 0x3e, 0x0e, 0xed, 0x30, 0xb0, 0x95, 0xd9, 0x15, 0xaa, 0xc9, 0x55,
	0x28, 0xd0, 0x95, 0x96, 0x27, 0x70, 0x13, 0x47, 0xf2, 0xb6, 0x71, 0x22, 0xc7, 0x01, 0x96, 0x1b,
	0xcb, 0xce, 0x4c, 0x57, 0x2e, 0x8e, 0xc7, 0xb2, 0x27, 0x2c, 0x5d, 0x71, 0xc1, 0xdb, 0x20, 0xc1,
	0x93, 0xf1, 0x16, 0x68, 0x66, 0x9c, 0xc4, 0x89, 0x7a, 0x81, 0xc4, 0xdd, 0x9c, 0xdf, 0xf9, 0xeb,
	0x7c, 0xf9, 0xcc, 0x18, 0x06, 0xf4, 0x57, 0x9a, 0xf3, 0xea, 0xa6, 0x28, 0x19, 0x67, 0x2f, 0xaf,
	0xde, 0x33, 0xf6, 0x3e, 0xa3, 0xaf, 0xa5, 0x95, 0x6c, 0x1f, 0x5e, 0xf3, 0x74, 0x43, 0x2b, 0x1e,
	0x6f, 0x0a, 0x25, 0x18, 0x8e, 0xe1, 0x7c, 0x51, 0xb2, 0x47, 0xba, 0xe6, 0xe3, 0x92, 0xc6, 0x9c,
	0x12, 0xfc, 0x25, 0x40, 0xa1, 0x48, 0x94, 0x12, 0x4b, 0xb3, 0xb5, 0x91, 0x11, 0x18, 0x35, 0xf1,
	0x08, 0xc6, 0xd0, 0xce, 0xe3, 0x0d, 0xb5, 0x5a, 0xd2, 0x21, 0xcf, 0xc3, 0xbf, 0x34, 0x38, 0x9f,
	0xc4, 0x3c, 0xae, 0x68, 0x33, 0x0a, 0x51, 0xa4, 0x11, 0xa5, 0x26, 0xcf, 0x47, 0x39, 0x49, 0xac,
	0x9f, 0x26, 0xfe, 0x16, 0xfa, 0x1b, 0xca, 0x63, 0x11, 0xc3, 0x6a, 0xdb, 0xda, 0xc8, 0x7c, 0x83,
	0x6e, 0xea, 0xa4, 0xb3, 0x9a, 0x07, 0x7b, 0x05, 0xb6, 0xa0, 0x47, 0x7f, 0x2b, 0x68, 0xc9, 0x2b,
	0xab, 0x63, 0xeb, 0x23, 0x23, 0xd8, 0x99, 0xc3, 0xdf, 0xf7, 0xb5, 0xae, 0x0a, 0xf2, 0x5f, 0x6a,
	0x3d, 0xae, 0xab, 0xf5, 0xbf, 0xea, 0x1a, 0xfe, 0xa9, 0x41, 0xff, 0x2d, 0x4b, 0x1c, 0x42, 0x28,
	0xc1, 0x2f, 0xa0, 0xfb, 0xc8, 0x92, 0x43, 0xd2, 0xce, 0x23, 0x4b, 0x3c, 0x82, 0xbf, 0x80, 0xbe,
	0xc0, 0x8d, 0x01, 0xf5, 0x1e, 0x59, 0xe2, 0x8b, 0x19, 0x5d, 0x42, 0x37, 0xcd, 0x8b, 0x2d, 0xaf,
	0x2c, 0x5d, 0x76, 0x55, 0x5b, 0xa2, 0x5d, 0xb6, 0xe5, 0xd2, 0xd1, 0x56, 0xed, 0xd6, 0xe6, 0x49,
	0xf5, 0x9d, 0xd3, 0xea, 0x1b, 0x73, 0xea, 0x1e, 0xcf, 0xe9, 0x0e, 0x4c, 0x57, 0x1e, 0x55, 0xad,
	0xaf, 0xc0, 0x50, 0x9e, 0x43, 0xb9, 0x7d, 0x05, 0x3c, 0x82, 0xaf, 0xc0, 0xac, 0x9d, 0x8d, 0xa2,
	0x41, 0x21, 0x51, 0xf7, 0xf0, 0x1f, 0x1d, 0x2e, 0x4e, 0x86, 0x82, 0xbf, 0x01, 0x10, 0x23, 0x7f,
	0x48, 0x69, 0x46, 0x2a, 0x4b, 0xb3, 0xf5, 0xd1, 0xf9, 0x9b, 0xc1, 0xcd, 0xd4, 0x73, 0xef, 0x27,
	0x51, 0xf8, 0x6e, 0xe1, 0x2e, 0x03, 0xa3, 0xa2, 0x7c, 0x2a, 0xdd, 0x42, 0x4c, 0x68, 0xb6, 0x13,
	0xb7, 0x9e, 0x13, 0x13, 0x9a, 0xd5, 0xe2, 0xaf, 0x60, 0x50, 0xd2, 0x35, 0x2b, 0x49, 0xb4, 0x66,
	0xdb, 0x9c, 0xcb, 0x5d, 0xd2, 0x03, 0x53, 0xb1, 0xb1, 0x40, 0x62, 0x2c, 0x0f, 0x69, 0x46, 0x6b,
	0x41, 0x5b, 0x0a, 0x0c, 0x41, 0x94, 0xfb, 0x15, 0x18, 0x65, 0xfc, 0x21, 0x4a, 0x9e, 0x38, 0xad,
	0xe4, 0xd0, 0xf4, 0xa0, 0x5f, 0xc6, 0x1f, 0x6e, 0x85, 0x2d, 0x9c, 0x1f, 0xd3, 0xa2, 0x76, 0x76,
	0x95, 0xf3, 0x63, 0x5a, 0x28, 0xe7, 0x15, 0x98, 0x55, 0xbc, 0x29, 0x32, 0x1a, 0x25, 0x8c, 0x3c,
	0x59, 0x3d, 0x5b, 0x1b, 0x0d, 0x02, 0x50, 0xe8, 0x96, 0x91, 0x27, 0xfc, 0xdd, 0x5e, 0xf0, 0x4b,
	0x9a, 0x13, 0xab, 0x6f, 0x6b, 0xb2, 0x95, 0xa5, 0x33, 0x5b, 0xdc, 0xbb, 0xd1, 0x9d, 0xe7, 0x4f,
	0x76, 0xf2, 0xbb, 0x34, 0x27, 0xf8, 0x6b, 0x40, 0x5b, 0xb9, 0xa7, 0xd1, 0xfe, 0xea, 0x5a, 0x86,
	0xcc, 0x79, 0xa1, 0x78, 0xb8, 0xc3, 0x22, 0xb5, 0x18, 0x6c, 0xf4, 0xc0, 0xca, 0x4d, 0xcc, 0x2d,
	0x50, 0x5f, 0x41, 0xa0, 0xa9, 0x24, 0xd8, 0x06, 0x93, 0xd0, 0x6a, 0x5d, 0xa6, 0x05, 0x4f, 0x59,
	0x6e, 0x99, 0x52, 0xd0, 0x44, 0x22, 0x44, 0xc6, 0xd6, 0xb1, 0x38, 0x8b, 0xef, 0x3c, 0x50, 0x21,
	0x76, 0xc8, 0x23, 0x62, 0xb4, 0x7b, 0xc1, 0xb6, 0x4c, 0xad, 0x4f, 0x54, 0x8c, 0x1d, 0x5b, 0x95,
	0xe9, 0xf0, 0x6f, 0x0d, 0x3a, 0xae, 0x78, 0x84, 0xc4, 0x2d, 0x17, 0x2d, 0xcb, 0x75, 0x19, 0x04,
	0xf2, 0x8c, 0x6d, 0x68, 0x87, 0x4f, 0x85, 0xda, 0x11, 0xd1, 0xb7, 0x54, 0xde, 0x08, 0x54, 0x05,
	0xd2, 0x33, 0xcc, 0xa0, 0x23, 0x4d, 0xdc, 0x87, 0xb6, 0xcf, 0x72, 0x8a, 0xce, 0x30, 0x3e, 0x7d,
	0xa5, 0x90, 0x26, 0xd8, 0xf1, 0x9b, 0x83, 0x5a, 0x0d, 0x56, 0xdf, 0x6d, 0xa4, 0xe3, 0x8b, 0xa3,
	0x3d, 0x46, 0x6d, 0x3c, 0x38, 0xdc, 0x40, 0xd4, 0xb9, 0xfe, 0xa3, 0x05, 0x66, 0x63, 0x8d, 0xf0,
	0x39, 0x80, 0x32, 0xfd, 0xb9, 0xef, 0xa2, 0x33, 0x7c, 0x09, 0x58, 0xd9, 0x81, 0x3b, 0x9e, 0x07,
	0x93, 0x68, 0x3c, 0x5f, 0xf9, 0x21, 0xd2, 0xf0, 0x67, 0x70, 0x51, 0x73, 0xe7, 0xc7, 0xe8, 0xf6,
	0x5d, 0xe8, 0x2e, 0x51, 0xeb, 0x00, 0x7f, 0xf6, 0x16, 0x35, 0xd4, 0x31, 0x82, 0x81, 0x82, 0xea,
	0x13, 0xa3, 0x36, 0x7e, 0x09, 0x97, 0x8a, 0xac, 0x16, 0x13, 0x27, 0x74, 0xa3, 0xd0, 0x9b, 0xb9,
	0xcb, 0xd0, 0x99, 0x2d, 0x50, 0x07, 0x7f, 0x0e, 0x48, 0xf9, 0xa6, 0xde, 0xbd, 0x5b, 0x67, 0xeb,
	0xe2, 0x17, 0xf0, 0xa9, 0xa2, 0x13, 0x27, 0x74, 0xa2, 0xe9, 0x3c, 0x98, 0x39, 0x21, 0xea, 0x35,
	0xb0, 0xbb, 0x1c, 0x07, 0xde, 0x22, 0xf4, 0xe6, 0x3e, 0xea, 0x1f, 0xf0, 0xfd, 0x7c, 0xec, 0x08,
	0x16, 0x79, 0x13, 0x64, 0x1c, 0x5a, 0xd9, 0xe3, 0x55, 0xe0, 0x21, 0xb8, 0xbe, 0x06, 0xb3, 0xb1,
	0x7d, 0x62, 0xec, 0xa1, 0xfb, 0x53, 0x88, 0xce, 0x70, 0x0f, 0xf4, 0x70, 0xf9, 0x03, 0xd2, 0x04,
	0x7a, 0xbb, 0x9c, 0xfb, 0xa8, 0x95, 0x74, 0xe5, 0x6f, 0xe3, 0xfb, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x48, 0x83, 0x14, 0xdd, 0x67, 0x06, 0x00, 0x00,
}
