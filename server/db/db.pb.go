// Code generated by protoc-gen-go.
// source: db.proto
// DO NOT EDIT!

/*
Package db is a generated protocol buffer package.

It is generated from these files:
	db.proto

It has these top-level messages:
	DatasetVersionData
	EntityRef
	ProjectData
	Expert
	DatasetData
	DatasetItemData
	Job
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
	// scoped to project
	Range_NAMES_SCOPED         Range = 11
	Range_IDX_PROJECT_DATASETS Range = 20
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
	11:  "NAMES_SCOPED",
	20:  "IDX_PROJECT_DATASETS",
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
	"NAMES_SCOPED":         11,
	"IDX_PROJECT_DATASETS": 20,
	"VERSION":              255,
}

func (x Range) String() string {
	return proto.EnumName(Range_name, int32(x))
}
func (Range) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DatasetVersionData struct {
	Uid       []byte `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ParentUid []byte `protobuf:"bytes,2,opt,name=parent_uid,json=parentUid,proto3" json:"parent_uid,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	// total
	RecordCount  int64                 `protobuf:"varint,4,opt,name=record_count,json=recordCount" json:"record_count,omitempty"`
	ItemCount    int64                 `protobuf:"varint,5,opt,name=item_count,json=itemCount" json:"item_count,omitempty"`
	StorageBytes int64                 `protobuf:"varint,6,opt,name=storage_bytes,json=storageBytes" json:"storage_bytes,omitempty"`
	Inputs       []*vo.DatasetVerInput `protobuf:"bytes,7,rep,name=inputs" json:"inputs,omitempty"`
	Items        []*vo.DatasetItem     `protobuf:"bytes,8,rep,name=items" json:"items,omitempty"`
	DatasetUid   []byte                `protobuf:"bytes,9,opt,name=dataset_uid,json=datasetUid,proto3" json:"dataset_uid,omitempty"`
}

func (m *DatasetVersionData) Reset()                    { *m = DatasetVersionData{} }
func (m *DatasetVersionData) String() string            { return proto.CompactTextString(m) }
func (*DatasetVersionData) ProtoMessage()               {}
func (*DatasetVersionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

type EntityRef struct {
	Kind vo.ENTITY `protobuf:"varint,1,opt,name=kind,enum=ENTITY" json:"kind,omitempty"`
	Uid  []byte    `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (m *EntityRef) Reset()                    { *m = EntityRef{} }
func (m *EntityRef) String() string            { return proto.CompactTextString(m) }
func (*EntityRef) ProtoMessage()               {}
func (*EntityRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
func (*ProjectData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Expert struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Uid      []byte `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	FullName string `protobuf:"bytes,3,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
}

func (m *Expert) Reset()                    { *m = Expert{} }
func (m *Expert) String() string            { return proto.CompactTextString(m) }
func (*Expert) ProtoMessage()               {}
func (*Expert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
	UpstreamJobs    [][]byte          `protobuf:"bytes,15,rep,name=upstream_jobs,json=upstreamJobs,proto3" json:"upstream_jobs,omitempty"`
	DownstreamJobs  [][]byte          `protobuf:"bytes,16,rep,name=downstream_jobs,json=downstreamJobs,proto3" json:"downstream_jobs,omitempty"`
	ProjectName     string            `protobuf:"bytes,17,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
}

func (m *DatasetData) Reset()                    { *m = DatasetData{} }
func (m *DatasetData) String() string            { return proto.CompactTextString(m) }
func (*DatasetData) ProtoMessage()               {}
func (*DatasetData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DatasetData) GetSample() *vo.DatasetSample {
	if m != nil {
		return m.Sample
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
	Uid        []byte          `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ProjectUid []byte          `protobuf:"bytes,2,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	Name       string          `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Title      string          `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Inputs     []*vo.JobInput  `protobuf:"bytes,5,rep,name=inputs" json:"inputs,omitempty"`
	Outputs    []*vo.JobOutput `protobuf:"bytes,6,rep,name=outputs" json:"outputs,omitempty"`
	Experts    [][]byte        `protobuf:"bytes,7,rep,name=experts,proto3" json:"experts,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Job) GetInputs() []*vo.JobInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Job) GetOutputs() []*vo.JobOutput {
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
func (*TenantStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AssetCache struct {
	Digest []byte `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	Body   []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *AssetCache) Reset()                    { *m = AssetCache{} }
func (m *AssetCache) String() string            { return proto.CompactTextString(m) }
func (*AssetCache) ProtoMessage()               {}
func (*AssetCache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type FieldSchema struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *FieldSchema) Reset()                    { *m = FieldSchema{} }
func (m *FieldSchema) String() string            { return proto.CompactTextString(m) }
func (*FieldSchema) ProtoMessage()               {}
func (*FieldSchema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type DatasetDescriptor struct {
	Type         string         `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	RecordCount  int64          `protobuf:"varint,2,opt,name=record_count,json=recordCount" json:"record_count,omitempty"`
	Schema       []*FieldSchema `protobuf:"bytes,3,rep,name=schema" json:"schema,omitempty"`
	SampleString string         `protobuf:"bytes,4,opt,name=sample_string,json=sampleString" json:"sample_string,omitempty"`
}

func (m *DatasetDescriptor) Reset()                    { *m = DatasetDescriptor{} }
func (m *DatasetDescriptor) String() string            { return proto.CompactTextString(m) }
func (*DatasetDescriptor) ProtoMessage()               {}
func (*DatasetDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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
func (*AppVersionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*DatasetVersionData)(nil), "DatasetVersionData")
	proto.RegisterType((*EntityRef)(nil), "EntityRef")
	proto.RegisterType((*ProjectData)(nil), "ProjectData")
	proto.RegisterType((*Expert)(nil), "Expert")
	proto.RegisterType((*DatasetData)(nil), "DatasetData")
	proto.RegisterType((*DatasetItemData)(nil), "DatasetItemData")
	proto.RegisterType((*Job)(nil), "Job")
	proto.RegisterType((*TenantStats)(nil), "TenantStats")
	proto.RegisterType((*AssetCache)(nil), "AssetCache")
	proto.RegisterType((*FieldSchema)(nil), "FieldSchema")
	proto.RegisterType((*DatasetDescriptor)(nil), "DatasetDescriptor")
	proto.RegisterType((*AppVersionData)(nil), "AppVersionData")
	proto.RegisterEnum("Range", Range_name, Range_value)
}

func init() { proto.RegisterFile("db.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1057 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x71, 0xec, 0xc4, 0xc7, 0x6e, 0xeb, 0x8e, 0xaa, 0x95, 0xc5, 0x0a, 0x6d, 0xd6, 0x5d,
	0x41, 0x40, 0xa8, 0x17, 0x45, 0x48, 0x88, 0x0b, 0xa4, 0x6c, 0xeb, 0xd5, 0xa6, 0xd2, 0xb6, 0xd5,
	0xd8, 0x5d, 0x2d, 0x57, 0x96, 0x13, 0x4f, 0xbb, 0x53, 0x62, 0x8f, 0x65, 0x4f, 0x16, 0xfa, 0x20,
	0x3c, 0x08, 0x37, 0xbc, 0x01, 0x2f, 0xc0, 0x2d, 0xe2, 0x59, 0x40, 0xf3, 0xe3, 0xfc, 0x34, 0xee,
	0xdd, 0xcc, 0x77, 0xbe, 0x9c, 0x39, 0x73, 0xbe, 0x6f, 0x8e, 0x03, 0xc3, 0x7c, 0x76, 0x52, 0xd5,
	0x8c, 0xb3, 0xcf, 0x87, 0x9f, 0x98, 0x5a, 0x85, 0x7f, 0xf4, 0x00, 0x9d, 0x67, 0x3c, 0x6b, 0x08,
	0x7f, 0x4f, 0xea, 0x86, 0xb2, 0x52, 0xec, 0x90, 0x0f, 0xe6, 0x92, 0xe6, 0x81, 0x31, 0x32, 0xc6,
	0x1e, 0x16, 0x4b, 0xf4, 0x05, 0x40, 0x95, 0xd5, 0xa4, 0xe4, 0xa9, 0x08, 0xf4, 0x64, 0xc0, 0x51,
	0xc8, 0x0d, 0xcd, 0xd1, 0x11, 0x58, 0x9c, 0xf2, 0x05, 0x09, 0xcc, 0x91, 0x31, 0x76, 0xb0, 0xda,
	0xa0, 0x97, 0xe0, 0xd5, 0x64, 0xce, 0xea, 0x3c, 0x9d, 0xb3, 0x65, 0xc9, 0x83, 0xfe, 0xc8, 0x18,
	0x9b, 0xd8, 0x55, 0xd8, 0x99, 0x80, 0x44, 0x5e, 0xca, 0x49, 0xa1, 0x09, 0x96, 0x24, 0x38, 0x02,
	0x51, 0xe1, 0x63, 0xd8, 0x6b, 0x38, 0xab, 0xb3, 0x3b, 0x92, 0xce, 0x1e, 0x38, 0x69, 0x02, 0x5b,
	0x32, 0x3c, 0x0d, 0xbe, 0x16, 0x18, 0x1a, 0x83, 0x4d, 0xcb, 0x6a, 0xc9, 0x9b, 0x60, 0x30, 0x32,
	0xc7, 0xee, 0xa9, 0x7f, 0xb2, 0xbe, 0xd2, 0x54, 0x04, 0xb0, 0x8e, 0xa3, 0x10, 0x2c, 0x91, 0xbb,
	0x09, 0x86, 0x92, 0xe8, 0xb5, 0xc4, 0x29, 0x27, 0x05, 0x56, 0x21, 0xf4, 0x02, 0xdc, 0x5c, 0xa1,
	0xf2, 0xaa, 0x8e, 0xbc, 0x2a, 0x68, 0xe8, 0x86, 0xe6, 0xe1, 0x8f, 0xe0, 0x44, 0x25, 0xa7, 0xfc,
	0x01, 0x93, 0x5b, 0xf4, 0x1c, 0xfa, 0xbf, 0xd0, 0x52, 0xb5, 0x6a, 0xff, 0x74, 0x70, 0x12, 0x5d,
	0x26, 0xd3, 0xe4, 0x67, 0x2c, 0xc1, 0xb6, 0x8d, 0xbd, 0x55, 0x1b, 0xc3, 0xbf, 0x0d, 0x70, 0xaf,
	0x6b, 0x76, 0x4f, 0xe6, 0x5c, 0x36, 0x7a, 0xd5, 0x37, 0x63, 0xb3, 0x6f, 0x08, 0xfa, 0x65, 0x56,
	0x10, 0xf9, 0x43, 0x07, 0xcb, 0x75, 0x9b, 0xcb, 0x5c, 0x4b, 0x72, 0x0c, 0x7b, 0x6d, 0xa1, 0xeb,
	0xf6, 0x5a, 0xd8, 0xd3, 0xe0, 0x13, 0x0d, 0xb4, 0x3a, 0x1a, 0xf8, 0x1c, 0x9c, 0x7b, 0x36, 0xd3,
	0x59, 0x6c, 0x99, 0x65, 0x78, 0xcf, 0x66, 0x2a, 0xc3, 0x08, 0xdc, 0x9c, 0x34, 0xf3, 0x9a, 0x56,
	0x9c, 0xb2, 0x32, 0x18, 0xc8, 0x9a, 0x36, 0xa1, 0x30, 0x03, 0x3b, 0xfa, 0xad, 0x22, 0x35, 0x5f,
	0x15, 0x6e, 0xec, 0x16, 0xbe, 0x6e, 0x82, 0x38, 0xee, 0x76, 0xb9, 0x58, 0xa4, 0x92, 0xaa, 0x0c,
	0x33, 0x14, 0xc0, 0xa5, 0xa0, 0x1f, 0x81, 0x45, 0x8a, 0x8c, 0x2e, 0xe4, 0x6d, 0x1c, 0xac, 0x36,
	0xe1, 0x9f, 0x7d, 0x70, 0xb5, 0x56, 0x4f, 0x18, 0xb4, 0xab, 0x67, 0xdd, 0xae, 0x7c, 0x01, 0x6e,
	0xa5, 0x24, 0x90, 0x02, 0xf7, 0x95, 0xc0, 0x1a, 0x12, 0x66, 0x7e, 0x6c, 0x5b, 0xab, 0xd3, 0xb6,
	0xb7, 0x74, 0x41, 0x36, 0x5a, 0x66, 0x62, 0x47, 0x20, 0x4f, 0x74, 0x7d, 0xd0, 0xd1, 0xf5, 0x2f,
	0xc1, 0x6e, 0xb2, 0xa2, 0x5a, 0x90, 0x60, 0x38, 0x32, 0xc6, 0xee, 0xe9, 0x7e, 0xeb, 0xc6, 0x58,
	0xa2, 0x58, 0x47, 0xd1, 0xd7, 0xe0, 0x2f, 0xab, 0x3c, 0xe3, 0x24, 0xe5, 0xb4, 0x20, 0x0d, 0xcf,
	0x8a, 0x4a, 0xba, 0xd2, 0xc4, 0x07, 0x0a, 0x4f, 0x5a, 0xb8, 0xf5, 0x6e, 0x7a, 0xcb, 0xea, 0x22,
	0xe3, 0x01, 0xc8, 0x6b, 0x4b, 0xef, 0xbe, 0x91, 0xc8, 0x63, 0x31, 0xdd, 0x1d, 0x31, 0x45, 0x8a,
	0x05, 0x9b, 0x67, 0x62, 0x9d, 0xd2, 0x3c, 0xf0, 0x54, 0x8a, 0x16, 0x9a, 0xca, 0xee, 0xac, 0x08,
	0xcb, 0x9a, 0x06, 0x7b, 0x2a, 0x47, 0x8b, 0xdd, 0xd4, 0x14, 0x05, 0x30, 0x20, 0xd2, 0x10, 0x4d,
	0xb0, 0x3f, 0x32, 0xc7, 0x1e, 0x6e, 0xb7, 0xa2, 0x31, 0xcb, 0xaa, 0xe1, 0x35, 0xc9, 0x8a, 0xf4,
	0x9e, 0xcd, 0x9a, 0xe0, 0x40, 0xc6, 0xbd, 0x16, 0xbc, 0x60, 0xb3, 0x06, 0x7d, 0x05, 0x07, 0x39,
	0xfb, 0xb5, 0xdc, 0xa4, 0xf9, 0x92, 0xb6, 0xbf, 0x86, 0x25, 0xf1, 0x25, 0x78, 0xad, 0x92, 0x52,
	0xfb, 0x43, 0x55, 0x8a, 0xc6, 0x84, 0x9d, 0x42, 0x06, 0x07, 0x1b, 0x6f, 0x5c, 0x7a, 0xa7, 0xcb,
	0x29, 0x3b, 0x82, 0x99, 0x1d, 0x82, 0x1d, 0xc3, 0x9e, 0xf2, 0x40, 0xb3, 0x35, 0xcf, 0xb4, 0x59,
	0x1a, 0x29, 0x7d, 0xf8, 0x97, 0x01, 0xe6, 0x05, 0x9b, 0x75, 0x38, 0xf4, 0x91, 0xef, 0x7a, 0x3b,
	0xbe, 0x6b, 0x0b, 0x33, 0xbb, 0x2c, 0xdc, 0xdf, 0x1e, 0xac, 0xed, 0xc4, 0xb3, 0xe4, 0x20, 0x73,
	0x4e, 0x2e, 0xd8, 0x6c, 0x7b, 0xd4, 0xbd, 0x82, 0x01, 0x5b, 0x72, 0xc9, 0xb1, 0x25, 0x07, 0x04,
	0xe7, 0x4a, 0x42, 0xb8, 0x0d, 0x6d, 0x2a, 0x35, 0xd8, 0x52, 0x2a, 0xfc, 0xc7, 0x00, 0x37, 0x21,
	0x65, 0x56, 0xf2, 0x98, 0x67, 0x4a, 0xb9, 0xb6, 0x7a, 0x75, 0x79, 0x43, 0x4d, 0x1b, 0x0d, 0xae,
	0x7c, 0xbf, 0x3d, 0x92, 0x7a, 0x1d, 0x23, 0x69, 0x6b, 0xda, 0x98, 0x8f, 0xa6, 0x8d, 0x7c, 0x7b,
	0x15, 0xab, 0xb7, 0x67, 0x9a, 0xab, 0xb0, 0x15, 0x45, 0x15, 0xb9, 0xf1, 0x3c, 0x2d, 0xec, 0x2a,
	0x4c, 0x51, 0x5e, 0x80, 0x5b, 0xb0, 0x9c, 0x2c, 0xb6, 0x46, 0x1a, 0x48, 0x48, 0xa9, 0xf4, 0x03,
	0xc0, 0xa4, 0x11, 0x15, 0x65, 0xf3, 0x8f, 0x04, 0x3d, 0x03, 0x3b, 0xa7, 0x77, 0xa4, 0xe1, 0x5a,
	0x2e, 0xbd, 0x13, 0x82, 0xcc, 0x58, 0xfe, 0xa0, 0xa5, 0x92, 0xeb, 0xf0, 0x7b, 0x70, 0xdf, 0x50,
	0xb2, 0xc8, 0xe3, 0xf9, 0x47, 0x52, 0x64, 0x9d, 0x13, 0x0f, 0x41, 0x9f, 0x3f, 0x54, 0x2b, 0x83,
	0x89, 0x75, 0xf8, 0xbb, 0x01, 0x87, 0xed, 0x00, 0xd3, 0xaf, 0x8d, 0xd5, 0x2b, 0xa6, 0xb1, 0x66,
	0xee, 0x4c, 0x9f, 0xde, 0xee, 0xf4, 0x79, 0x05, 0x76, 0x23, 0x8f, 0x0f, 0x4c, 0xfd, 0x1d, 0xdb,
	0x28, 0x09, 0xeb, 0x98, 0xf4, 0xb4, 0x9c, 0x20, 0x69, 0xc3, 0x6b, 0x5a, 0xde, 0x69, 0x0b, 0x79,
	0x0a, 0x8c, 0x25, 0x16, 0xfe, 0x04, 0xfb, 0x93, 0xaa, 0xda, 0xfc, 0xf6, 0x7f, 0x0b, 0x87, 0xfa,
	0x0b, 0x45, 0x59, 0xa9, 0x03, 0xba, 0xc0, 0xdd, 0xc0, 0x37, 0xff, 0x1a, 0x60, 0xe1, 0xac, 0xbc,
	0x23, 0x08, 0xc0, 0x8e, 0x3e, 0x91, 0x92, 0x37, 0xfe, 0x67, 0xc8, 0x83, 0xe1, 0x35, 0xbe, 0xba,
	0x88, 0xce, 0x92, 0xd8, 0x37, 0xc4, 0xee, 0x7c, 0x92, 0x4c, 0xe2, 0x28, 0x89, 0xfd, 0x1e, 0x1a,
	0x42, 0xff, 0xe2, 0xea, 0x75, 0xec, 0xf7, 0x91, 0x03, 0x56, 0x9c, 0x4c, 0x92, 0xd8, 0xb7, 0xd0,
	0x01, 0xb8, 0x93, 0x38, 0x8e, 0x92, 0xf4, 0x6c, 0x72, 0xf6, 0x36, 0xf2, 0x6d, 0xe4, 0xc2, 0x20,
	0xfa, 0x70, 0x1d, 0xe1, 0x24, 0xf6, 0x07, 0xe8, 0x08, 0x7c, 0x9d, 0x20, 0x7d, 0x1f, 0xe1, 0x78,
	0x7a, 0x75, 0x19, 0xfb, 0x43, 0xf4, 0x0c, 0x50, 0x8b, 0xbe, 0x8d, 0x26, 0xe7, 0xe9, 0x34, 0x89,
	0xde, 0xc5, 0xbe, 0x83, 0x0e, 0x61, 0x4f, 0x1f, 0x9e, 0x5e, 0x4e, 0xde, 0x45, 0xb1, 0x0f, 0xc8,
	0x07, 0x4f, 0x2e, 0xd3, 0xf8, 0xec, 0xea, 0x3a, 0x3a, 0xf7, 0x5d, 0x14, 0xc0, 0xd1, 0xf4, 0xfc,
	0x43, 0xda, 0x12, 0x57, 0xf5, 0x1d, 0x21, 0x0f, 0x06, 0xfa, 0x10, 0xff, 0x3f, 0x63, 0x66, 0xcb,
	0xff, 0x49, 0xdf, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x75, 0x7e, 0x30, 0x3d, 0x09, 0x00,
	0x00,
}
