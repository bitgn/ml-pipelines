// Code generated by protoc-gen-go.
// source: db.proto
// DO NOT EDIT!

/*
Package db is a generated protocol buffer package.

It is generated from these files:
	db.proto

It has these top-level messages:
	EntityRef
	ProjectData
	Expert
	DatasetData
	DatasetVersionData
	DatasetItemData
	Job
	JobRunData
	TenantStats
	AssetCache
	FieldSchema
	DatasetDescriptor
	AppVersionData
*/
package db

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import vo "mlp/catalog/vo"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Range int32

const (
	Range_Events             Range = 0
	Range_PROJECTS           Range = 1
	Range_DATASETS           Range = 2
	Range_JOBS               Range = 4
	Range_STATS              Range = 5
	Range_ASSET_CACHE        Range = 6
	Range_EXPERTS            Range = 7
	Range_DATASET_VERSIONS   Range = 8
	Range_DATASET_HEAD_ITEMS Range = 9
	Range_PROJECT_NAMES      Range = 10
	Range_JOB_RUNS           Range = 12
	// scoped to project
	Range_NAMES_SCOPED         Range = 11
	Range_IDX_PROJECT_DATASETS Range = 20
	Range_IDX_DATASET_VERSIONS Range = 21
	Range_VERSION              Range = 255
)

var Range_name = map[int32]string{
	0:   "Events",
	1:   "PROJECTS",
	2:   "DATASETS",
	4:   "JOBS",
	5:   "STATS",
	6:   "ASSET_CACHE",
	7:   "EXPERTS",
	8:   "DATASET_VERSIONS",
	9:   "DATASET_HEAD_ITEMS",
	10:  "PROJECT_NAMES",
	12:  "JOB_RUNS",
	11:  "NAMES_SCOPED",
	20:  "IDX_PROJECT_DATASETS",
	21:  "IDX_DATASET_VERSIONS",
	255: "VERSION",
}
var Range_value = map[string]int32{
	"Events":               0,
	"PROJECTS":             1,
	"DATASETS":             2,
	"JOBS":                 4,
	"STATS":                5,
	"ASSET_CACHE":          6,
	"EXPERTS":              7,
	"DATASET_VERSIONS":     8,
	"DATASET_HEAD_ITEMS":   9,
	"PROJECT_NAMES":        10,
	"JOB_RUNS":             12,
	"NAMES_SCOPED":         11,
	"IDX_PROJECT_DATASETS": 20,
	"IDX_DATASET_VERSIONS": 21,
	"VERSION":              255,
}

func (x Range) String() string {
	return proto.EnumName(Range_name, int32(x))
}
func (Range) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EntityRef struct {
	Kind vo.ENTITY `protobuf:"varint,1,opt,name=kind,enum=ENTITY" json:"kind,omitempty"`
	Uid  []byte    `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (m *EntityRef) Reset()                    { *m = EntityRef{} }
func (m *EntityRef) String() string            { return proto.CompactTextString(m) }
func (*EntityRef) ProtoMessage()               {}
func (*EntityRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ProjectData struct {
	Title        string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Uid          []byte `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	DatasetCount int32  `protobuf:"varint,4,opt,name=dataset_count,json=datasetCount" json:"dataset_count,omitempty"`
	StorageBytes int64  `protobuf:"varint,5,opt,name=storage_bytes,json=storageBytes" json:"storage_bytes,omitempty"`
	JobCount     int32  `protobuf:"varint,6,opt,name=job_count,json=jobCount" json:"job_count,omitempty"`
	Description  string `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
}

func (m *ProjectData) Reset()                    { *m = ProjectData{} }
func (m *ProjectData) String() string            { return proto.CompactTextString(m) }
func (*ProjectData) ProtoMessage()               {}
func (*ProjectData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Expert struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Uid      []byte `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	FullName string `protobuf:"bytes,3,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
}

func (m *Expert) Reset()                    { *m = Expert{} }
func (m *Expert) String() string            { return proto.CompactTextString(m) }
func (*Expert) ProtoMessage()               {}
func (*Expert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DatasetData struct {
	Uid             []byte            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name            string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Title           string            `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	ProjectUid      []byte            `protobuf:"bytes,4,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	RecordCount     int64             `protobuf:"varint,5,opt,name=record_count,json=recordCount" json:"record_count,omitempty"`
	FileCount       int64             `protobuf:"varint,6,opt,name=file_count,json=fileCount" json:"file_count,omitempty"`
	StorageBytes    int64             `protobuf:"varint,7,opt,name=storage_bytes,json=storageBytes" json:"storage_bytes,omitempty"`
	Sample          *vo.DatasetSample `protobuf:"bytes,8,opt,name=sample" json:"sample,omitempty"`
	UpdateTimestamp int64             `protobuf:"varint,9,opt,name=update_timestamp,json=updateTimestamp" json:"update_timestamp,omitempty"`
	DataFormat      string            `protobuf:"bytes,10,opt,name=data_format,json=dataFormat" json:"data_format,omitempty"`
	Description     string            `protobuf:"bytes,11,opt,name=description" json:"description,omitempty"`
	LocationId      string            `protobuf:"bytes,12,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	LocationUri     string            `protobuf:"bytes,13,opt,name=location_uri,json=locationUri" json:"location_uri,omitempty"`
	Experts         [][]byte          `protobuf:"bytes,14,rep,name=experts,proto3" json:"experts,omitempty"`
	ProjectName     string            `protobuf:"bytes,17,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
}

func (m *DatasetData) Reset()                    { *m = DatasetData{} }
func (m *DatasetData) String() string            { return proto.CompactTextString(m) }
func (*DatasetData) ProtoMessage()               {}
func (*DatasetData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DatasetData) GetSample() *vo.DatasetSample {
	if m != nil {
		return m.Sample
	}
	return nil
}

type DatasetVersionData struct {
	Uid       []byte `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ParentUid []byte `protobuf:"bytes,2,opt,name=parent_uid,json=parentUid,proto3" json:"parent_uid,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	// total
	RecordCount  int64                  `protobuf:"varint,4,opt,name=record_count,json=recordCount" json:"record_count,omitempty"`
	ItemCount    int64                  `protobuf:"varint,5,opt,name=item_count,json=itemCount" json:"item_count,omitempty"`
	StorageBytes int64                  `protobuf:"varint,6,opt,name=storage_bytes,json=storageBytes" json:"storage_bytes,omitempty"`
	Inputs       []*vo.DatasetVerInput  `protobuf:"bytes,7,rep,name=inputs" json:"inputs,omitempty"`
	Items        []*vo.DatasetItem      `protobuf:"bytes,8,rep,name=items" json:"items,omitempty"`
	DatasetUid   []byte                 `protobuf:"bytes,9,opt,name=dataset_uid,json=datasetUid,proto3" json:"dataset_uid,omitempty"`
	Outputs      []*vo.DatasetVerOutput `protobuf:"bytes,10,rep,name=outputs" json:"outputs,omitempty"`
}

func (m *DatasetVersionData) Reset()                    { *m = DatasetVersionData{} }
func (m *DatasetVersionData) String() string            { return proto.CompactTextString(m) }
func (*DatasetVersionData) ProtoMessage()               {}
func (*DatasetVersionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DatasetVersionData) GetInputs() []*vo.DatasetVerInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *DatasetVersionData) GetItems() []*vo.DatasetItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *DatasetVersionData) GetOutputs() []*vo.DatasetVerOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type DatasetItemData struct {
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	StorageBytes int64  `protobuf:"varint,3,opt,name=storage_bytes,json=storageBytes" json:"storage_bytes,omitempty"`
	RecordsCount int64  `protobuf:"varint,4,opt,name=records_count,json=recordsCount" json:"records_count,omitempty"`
}

func (m *DatasetItemData) Reset()                    { *m = DatasetItemData{} }
func (m *DatasetItemData) String() string            { return proto.CompactTextString(m) }
func (*DatasetItemData) ProtoMessage()               {}
func (*DatasetItemData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Job struct {
	Uid        []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ProjectUid []byte   `protobuf:"bytes,2,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	Name       string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Title      string   `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Experts    [][]byte `protobuf:"bytes,7,rep,name=experts,proto3" json:"experts,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type JobRunData struct {
	Uid       []byte             `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	JobUid    []byte             `protobuf:"bytes,2,opt,name=job_uid,json=jobUid,proto3" json:"job_uid,omitempty"`
	Title     string             `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Inputs    []*vo.JobRunInput  `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	Outputs   []*vo.JobRunOutput `protobuf:"bytes,5,rep,name=outputs" json:"outputs,omitempty"`
	Timestamp int64              `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *JobRunData) Reset()                    { *m = JobRunData{} }
func (m *JobRunData) String() string            { return proto.CompactTextString(m) }
func (*JobRunData) ProtoMessage()               {}
func (*JobRunData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *JobRunData) GetInputs() []*vo.JobRunInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *JobRunData) GetOutputs() []*vo.JobRunOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type TenantStats struct {
	ProjectCount int32 `protobuf:"varint,1,opt,name=project_count,json=projectCount" json:"project_count,omitempty"`
	DatasetCount int32 `protobuf:"varint,2,opt,name=dataset_count,json=datasetCount" json:"dataset_count,omitempty"`
	JobCount     int32 `protobuf:"varint,3,opt,name=job_count,json=jobCount" json:"job_count,omitempty"`
	ReportCount  int32 `protobuf:"varint,4,opt,name=report_count,json=reportCount" json:"report_count,omitempty"`
	ExpertCount  int32 `protobuf:"varint,5,opt,name=expert_count,json=expertCount" json:"expert_count,omitempty"`
	ModelCount   int32 `protobuf:"varint,6,opt,name=model_count,json=modelCount" json:"model_count,omitempty"`
}

func (m *TenantStats) Reset()                    { *m = TenantStats{} }
func (m *TenantStats) String() string            { return proto.CompactTextString(m) }
func (*TenantStats) ProtoMessage()               {}
func (*TenantStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type AssetCache struct {
	Digest []byte `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	Body   []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *AssetCache) Reset()                    { *m = AssetCache{} }
func (m *AssetCache) String() string            { return proto.CompactTextString(m) }
func (*AssetCache) ProtoMessage()               {}
func (*AssetCache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type FieldSchema struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *FieldSchema) Reset()                    { *m = FieldSchema{} }
func (m *FieldSchema) String() string            { return proto.CompactTextString(m) }
func (*FieldSchema) ProtoMessage()               {}
func (*FieldSchema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type DatasetDescriptor struct {
	Type         string         `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	RecordCount  int64          `protobuf:"varint,2,opt,name=record_count,json=recordCount" json:"record_count,omitempty"`
	Schema       []*FieldSchema `protobuf:"bytes,3,rep,name=schema" json:"schema,omitempty"`
	SampleString string         `protobuf:"bytes,4,opt,name=sample_string,json=sampleString" json:"sample_string,omitempty"`
}

func (m *DatasetDescriptor) Reset()                    { *m = DatasetDescriptor{} }
func (m *DatasetDescriptor) String() string            { return proto.CompactTextString(m) }
func (*DatasetDescriptor) ProtoMessage()               {}
func (*DatasetDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DatasetDescriptor) GetSchema() []*FieldSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

type AppVersionData struct {
	ProjectionVersion string `protobuf:"bytes,1,opt,name=ProjectionVersion" json:"ProjectionVersion,omitempty"`
}

func (m *AppVersionData) Reset()                    { *m = AppVersionData{} }
func (m *AppVersionData) String() string            { return proto.CompactTextString(m) }
func (*AppVersionData) ProtoMessage()               {}
func (*AppVersionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*EntityRef)(nil), "EntityRef")
	proto.RegisterType((*ProjectData)(nil), "ProjectData")
	proto.RegisterType((*Expert)(nil), "Expert")
	proto.RegisterType((*DatasetData)(nil), "DatasetData")
	proto.RegisterType((*DatasetVersionData)(nil), "DatasetVersionData")
	proto.RegisterType((*DatasetItemData)(nil), "DatasetItemData")
	proto.RegisterType((*Job)(nil), "Job")
	proto.RegisterType((*JobRunData)(nil), "JobRunData")
	proto.RegisterType((*TenantStats)(nil), "TenantStats")
	proto.RegisterType((*AssetCache)(nil), "AssetCache")
	proto.RegisterType((*FieldSchema)(nil), "FieldSchema")
	proto.RegisterType((*DatasetDescriptor)(nil), "DatasetDescriptor")
	proto.RegisterType((*AppVersionData)(nil), "AppVersionData")
	proto.RegisterEnum("Range", Range_name, Range_value)
}

func init() { proto.RegisterFile("db.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1084 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x2e, 0x45, 0x89, 0x12, 0x87, 0xb4, 0x4d, 0x2f, 0xdc, 0x94, 0x68, 0x1a, 0x44, 0x65, 0x82,
	0x56, 0xfd, 0x81, 0x0f, 0x2e, 0x0a, 0x14, 0x3d, 0x14, 0x90, 0x2d, 0x06, 0x91, 0x81, 0x58, 0xc6,
	0x92, 0x0e, 0xd2, 0x13, 0x41, 0x89, 0x6b, 0x87, 0xae, 0xc4, 0x25, 0xc8, 0x55, 0x50, 0x03, 0xbd,
	0xf6, 0x11, 0xfa, 0x32, 0x7d, 0x8c, 0x5e, 0x0b, 0xf4, 0xdc, 0xb7, 0x68, 0xb1, 0x3f, 0xa4, 0x28,
	0x8b, 0xba, 0xed, 0x7e, 0x33, 0x9c, 0x9d, 0xf9, 0xbe, 0xd9, 0x59, 0xc2, 0x20, 0x99, 0x9f, 0xe6,
	0x05, 0x65, 0xf4, 0xd3, 0xc1, 0x07, 0x2a, 0x57, 0xde, 0x8f, 0x60, 0xfa, 0x19, 0x4b, 0xd9, 0x03,
	0x26, 0xb7, 0xe8, 0x29, 0x74, 0x7f, 0x49, 0xb3, 0xc4, 0xd5, 0x86, 0xda, 0xe8, 0xf0, 0xac, 0x7f,
	0xea, 0x5f, 0x85, 0xd3, 0xf0, 0x67, 0x2c, 0x40, 0xe4, 0x80, 0xbe, 0x4e, 0x13, 0xb7, 0x33, 0xd4,
	0x46, 0x36, 0xe6, 0x4b, 0xef, 0x2f, 0x0d, 0xac, 0xeb, 0x82, 0xde, 0x93, 0x05, 0x9b, 0xc4, 0x2c,
	0x46, 0x27, 0xd0, 0x63, 0x29, 0x5b, 0x12, 0xf1, 0xbd, 0x89, 0xe5, 0x06, 0x21, 0xe8, 0x66, 0xf1,
	0x8a, 0x88, 0x0f, 0x4d, 0x2c, 0xd6, 0x55, 0x2c, 0xbd, 0x8e, 0x85, 0x5e, 0xc0, 0x41, 0x12, 0xb3,
	0xb8, 0x24, 0x2c, 0x5a, 0xd0, 0x75, 0xc6, 0xdc, 0xee, 0x50, 0x1b, 0xf5, 0xb0, 0xad, 0xc0, 0x0b,
	0x8e, 0x71, 0xa7, 0x92, 0xd1, 0x22, 0xbe, 0x23, 0xd1, 0xfc, 0x81, 0x91, 0xd2, 0xed, 0x0d, 0xb5,
	0x91, 0x8e, 0x6d, 0x05, 0x9e, 0x73, 0x0c, 0x3d, 0x05, 0xf3, 0x9e, 0xce, 0x55, 0x14, 0x43, 0x44,
	0x19, 0xdc, 0xd3, 0xb9, 0x8c, 0x30, 0x04, 0x2b, 0x21, 0xe5, 0xa2, 0x48, 0x73, 0x96, 0xd2, 0xcc,
	0xed, 0x8b, 0x9c, 0x9a, 0x90, 0x17, 0x83, 0xe1, 0xff, 0x9a, 0x93, 0x82, 0xd5, 0x89, 0x6b, 0xbb,
	0x89, 0x6f, 0x48, 0xe0, 0xc7, 0xdd, 0xae, 0x97, 0xcb, 0x48, 0xb8, 0xea, 0xc2, 0x75, 0xc0, 0x81,
	0x2b, 0xee, 0x7e, 0x02, 0x3d, 0xb2, 0x8a, 0xd3, 0xa5, 0xa8, 0xc6, 0xc4, 0x72, 0xe3, 0xfd, 0xab,
	0x83, 0x35, 0x91, 0x75, 0x09, 0xde, 0x54, 0x50, 0x6d, 0x13, 0xb4, 0x8d, 0xb3, 0x9a, 0x5d, 0xbd,
	0xc9, 0xee, 0x73, 0xb0, 0x72, 0x29, 0x41, 0xc4, 0x63, 0x74, 0x45, 0x0c, 0x50, 0xd0, 0x4d, 0x9a,
	0xa0, 0xcf, 0xc1, 0x2e, 0xc8, 0x82, 0x16, 0x89, 0x62, 0x44, 0x52, 0x66, 0x49, 0x4c, 0x92, 0xf2,
	0x0c, 0xe0, 0x36, 0x5d, 0x92, 0x06, 0x65, 0x3a, 0x36, 0x39, 0xb2, 0x87, 0xf5, 0x7e, 0x0b, 0xeb,
	0x5f, 0x80, 0x51, 0xc6, 0xab, 0x7c, 0x49, 0xdc, 0xc1, 0x50, 0x1b, 0x59, 0x67, 0x87, 0xa7, 0xaa,
	0xc2, 0x40, 0xa0, 0x58, 0x59, 0xd1, 0x57, 0xe0, 0xac, 0xf3, 0x24, 0x66, 0x24, 0x62, 0xe9, 0x8a,
	0x94, 0x2c, 0x5e, 0xe5, 0xae, 0x29, 0xe2, 0x1d, 0x49, 0x3c, 0xac, 0x60, 0x5e, 0x1a, 0x57, 0x3f,
	0xba, 0xa5, 0xc5, 0x2a, 0x66, 0x2e, 0x88, 0xb2, 0x81, 0x43, 0xaf, 0x04, 0xf2, 0x58, 0x4c, 0x6b,
	0x47, 0x4c, 0x1e, 0x62, 0x49, 0x17, 0x31, 0x5f, 0x47, 0x69, 0xe2, 0xda, 0x32, 0x44, 0x05, 0x4d,
	0x05, 0x3b, 0xb5, 0xc3, 0xba, 0x48, 0xdd, 0x03, 0x19, 0xa3, 0xc2, 0x6e, 0x8a, 0x14, 0xb9, 0xd0,
	0x27, 0xa2, 0x21, 0x4a, 0xf7, 0x70, 0xa8, 0x8f, 0x6c, 0x5c, 0x6d, 0xf9, 0xc7, 0x15, 0xf7, 0x42,
	0xad, 0x63, 0xf9, 0xb1, 0xc2, 0x78, 0x03, 0x78, 0xff, 0x74, 0x00, 0x29, 0x22, 0xde, 0x92, 0xa2,
	0x4c, 0x69, 0xb6, 0x47, 0xf1, 0x67, 0x00, 0x79, 0x5c, 0x90, 0x4c, 0xca, 0x28, 0xfb, 0xcb, 0x94,
	0x08, 0x57, 0xb1, 0x5d, 0xfc, 0xc7, 0xda, 0x76, 0x5b, 0xb5, 0x4d, 0x19, 0x59, 0x6d, 0x89, 0x6f,
	0x72, 0x64, 0x8f, 0xb6, 0x46, 0x8b, 0xb6, 0x23, 0x30, 0xd2, 0x2c, 0x5f, 0x33, 0xae, 0xbc, 0x3e,
	0xb2, 0xce, 0x9c, 0xd3, 0x4d, 0x49, 0x53, 0x6e, 0xc0, 0xca, 0x8e, 0x3c, 0xe8, 0xf1, 0xd8, 0xa5,
	0x3b, 0x10, 0x8e, 0x76, 0xe5, 0x38, 0x65, 0x64, 0x85, 0xa5, 0xa9, 0x92, 0x95, 0xdf, 0x74, 0x5e,
	0xaa, 0x29, 0x3b, 0x56, 0x41, 0xbc, 0xd6, 0x6f, 0xa0, 0x4f, 0xd7, 0x4c, 0x9c, 0x07, 0x22, 0xcc,
	0x71, 0xe3, 0xbc, 0x99, 0xb0, 0xe0, 0xca, 0xc3, 0xa3, 0x70, 0xd4, 0x38, 0x43, 0x90, 0xdb, 0x76,
	0x79, 0x76, 0xea, 0xd4, 0x5b, 0xea, 0x7c, 0x01, 0x07, 0x92, 0xba, 0x72, 0x8b, 0x4f, 0xc5, 0x71,
	0x29, 0x18, 0xf3, 0x7e, 0x03, 0xfd, 0x92, 0xce, 0x5b, 0x14, 0x7c, 0x74, 0x13, 0x3b, 0x3b, 0x37,
	0xb1, 0xca, 0x4b, 0x6f, 0xbb, 0xd4, 0xdd, 0xa6, 0xae, 0x8d, 0x96, 0xeb, 0x6f, 0xb5, 0x9c, 0xf7,
	0xa7, 0x06, 0x70, 0x49, 0xe7, 0x78, 0xbd, 0xaf, 0x8f, 0x3e, 0x81, 0x3e, 0x9f, 0x7e, 0x9b, 0x0c,
	0x8c, 0x7b, 0x3a, 0xdf, 0xdf, 0x41, 0x2f, 0x6b, 0x69, 0xbb, 0x4a, 0x31, 0x19, 0x7d, 0x5b, 0xd6,
	0x2f, 0x37, 0x8a, 0xf4, 0x84, 0xdb, 0x81, 0x72, 0x7b, 0xa4, 0x06, 0xfa, 0x0c, 0xcc, 0xcd, 0xb5,
	0x56, 0x83, 0xa4, 0x06, 0xbc, 0xbf, 0x35, 0xb0, 0x42, 0x92, 0xc5, 0x19, 0x0b, 0x58, 0xcc, 0x04,
	0xdf, 0x15, 0x63, 0x92, 0x6f, 0x4d, 0xce, 0x7c, 0x05, 0xd6, 0x1d, 0xba, 0xfd, 0x30, 0x74, 0x5a,
	0x1e, 0x86, 0xad, 0x99, 0xaf, 0x3f, 0x9a, 0xf9, 0xe2, 0x96, 0xe4, 0xb4, 0xd8, 0x7e, 0x59, 0x2c,
	0x89, 0xd5, 0x2e, 0x92, 0xe1, 0xc6, 0x3d, 0xe9, 0x61, 0x4b, 0x62, 0xd2, 0xe5, 0x39, 0x58, 0x2b,
	0x9a, 0x90, 0xe5, 0xd6, 0xc3, 0x02, 0x02, 0x92, 0x8d, 0xf1, 0x03, 0xc0, 0xb8, 0xe4, 0x19, 0xc5,
	0x8b, 0xf7, 0x04, 0x3d, 0x01, 0x23, 0x49, 0xef, 0x48, 0xc9, 0x94, 0x38, 0x6a, 0xc7, 0x9b, 0x60,
	0x4e, 0x93, 0x07, 0x25, 0x8e, 0x58, 0x7b, 0xdf, 0x83, 0xf5, 0x2a, 0x25, 0xcb, 0x24, 0x58, 0xbc,
	0x27, 0xab, 0xb8, 0xf5, 0xdd, 0x41, 0xd0, 0x65, 0x0f, 0x79, 0xdd, 0xd3, 0x7c, 0xed, 0xfd, 0xa1,
	0xc1, 0x71, 0xf5, 0x8c, 0xa8, 0x99, 0x47, 0x8b, 0xda, 0x53, 0xdb, 0x78, 0xee, 0xcc, 0x89, 0xce,
	0xee, 0x9c, 0x78, 0x09, 0x46, 0x29, 0x8e, 0x77, 0x75, 0xd5, 0x08, 0x8d, 0x94, 0xb0, 0xb2, 0x89,
	0x6b, 0x24, 0xe6, 0x78, 0x54, 0xb2, 0x22, 0xcd, 0xee, 0x54, 0xdb, 0xda, 0x12, 0x0c, 0x04, 0xe6,
	0xfd, 0x04, 0x87, 0xe3, 0x3c, 0x6f, 0x8e, 0xbb, 0x6f, 0xe1, 0x58, 0xfd, 0x27, 0xa4, 0x34, 0x53,
	0x06, 0x95, 0xe0, 0xae, 0xe1, 0xeb, 0xdf, 0x3b, 0xd0, 0xc3, 0x71, 0x76, 0x47, 0x10, 0x80, 0xe1,
	0x7f, 0x20, 0x19, 0x2b, 0x9d, 0x8f, 0x90, 0x0d, 0x83, 0x6b, 0x3c, 0xbb, 0xf4, 0x2f, 0xc2, 0xc0,
	0xd1, 0xf8, 0x6e, 0x32, 0x0e, 0xc7, 0x81, 0x1f, 0x06, 0x4e, 0x07, 0x0d, 0xa0, 0x7b, 0x39, 0x3b,
	0x0f, 0x9c, 0x2e, 0x32, 0xa1, 0x17, 0x84, 0xe3, 0x30, 0x70, 0x7a, 0xe8, 0x08, 0xac, 0x71, 0x10,
	0xf8, 0x61, 0x74, 0x31, 0xbe, 0x78, 0xed, 0x3b, 0x06, 0xb2, 0xa0, 0xef, 0xbf, 0xbb, 0xf6, 0x71,
	0x18, 0x38, 0x7d, 0x74, 0x02, 0x8e, 0x0a, 0x10, 0xbd, 0xf5, 0x71, 0x30, 0x9d, 0x5d, 0x05, 0xce,
	0x00, 0x3d, 0x01, 0x54, 0xa1, 0xaf, 0xfd, 0xf1, 0x24, 0x9a, 0x86, 0xfe, 0x9b, 0xc0, 0x31, 0xd1,
	0x31, 0x1c, 0xa8, 0xc3, 0xa3, 0xab, 0xf1, 0x1b, 0x3f, 0x70, 0x80, 0x67, 0x70, 0x39, 0x3b, 0x8f,
	0xf0, 0xcd, 0x55, 0xe0, 0xd8, 0xc8, 0x01, 0x5b, 0x18, 0xa2, 0xe0, 0x62, 0x76, 0xed, 0x4f, 0x1c,
	0x0b, 0xb9, 0x70, 0x32, 0x9d, 0xbc, 0x8b, 0xaa, 0xcf, 0xea, 0x6c, 0x4f, 0x2a, 0xcb, 0xce, 0xf1,
	0x1f, 0x23, 0x1b, 0xfa, 0x6a, 0xe7, 0xfc, 0xa7, 0xcd, 0x0d, 0xf1, 0x87, 0xf6, 0xdd, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x5e, 0xba, 0xf4, 0xc7, 0xb7, 0x09, 0x00, 0x00,
}
