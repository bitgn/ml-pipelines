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
	ServiceData
	ServiceVersionData
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
	Range_Events              Range = 0
	Range_PROJECTS            Range = 1
	Range_DATASETS            Range = 2
	Range_JOBS                Range = 4
	Range_STATS               Range = 5
	Range_ASSET_CACHE         Range = 6
	Range_EXPERTS             Range = 7
	Range_DATASET_VERSIONS    Range = 8
	Range_DATASET_HEAD_ITEMS  Range = 9
	Range_SERVICES            Range = 13
	Range_SERVICE_VERSIONS    Range = 14
	Range_DATASET_BINARY_TREE Range = 22
	Range_PROJECT_NAMES       Range = 10
	Range_JOB_RUNS            Range = 12
	// scoped to project
	Range_NAMES_SCOPED         Range = 11
	Range_IDX_PROJECT_DATASETS Range = 20
	Range_IDX_DATASET_VERSIONS Range = 21
	// service UID/int32
	Range_IDX_SERVICE_VERSIONS Range = 23
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
	13:  "SERVICES",
	14:  "SERVICE_VERSIONS",
	22:  "DATASET_BINARY_TREE",
	10:  "PROJECT_NAMES",
	12:  "JOB_RUNS",
	11:  "NAMES_SCOPED",
	20:  "IDX_PROJECT_DATASETS",
	21:  "IDX_DATASET_VERSIONS",
	23:  "IDX_SERVICE_VERSIONS",
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
	"SERVICES":             13,
	"SERVICE_VERSIONS":     14,
	"DATASET_BINARY_TREE":  22,
	"PROJECT_NAMES":        10,
	"JOB_RUNS":             12,
	"NAMES_SCOPED":         11,
	"IDX_PROJECT_DATASETS": 20,
	"IDX_DATASET_VERSIONS": 21,
	"IDX_SERVICE_VERSIONS": 23,
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
	ServiceCount int32  `protobuf:"varint,8,opt,name=service_count,json=serviceCount" json:"service_count,omitempty"`
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
	HeadVersion     []byte            `protobuf:"bytes,20,opt,name=head_version,json=headVersion,proto3" json:"head_version,omitempty"`
	VersionCount    int32             `protobuf:"varint,22,opt,name=version_count,json=versionCount" json:"version_count,omitempty"`
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
	VersionNum   int32                  `protobuf:"varint,11,opt,name=version_num,json=versionNum" json:"version_num,omitempty"`
	Timestamp    int64                  `protobuf:"varint,12,opt,name=timestamp" json:"timestamp,omitempty"`
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

type ServiceData struct {
	Uid         []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Title       string   `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	ProjectUid  []byte   `protobuf:"bytes,4,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	Description string   `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	LocationId  string   `protobuf:"bytes,6,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	LocationUri string   `protobuf:"bytes,7,opt,name=location_uri,json=locationUri" json:"location_uri,omitempty"`
	Experts     [][]byte `protobuf:"bytes,14,rep,name=experts,proto3" json:"experts,omitempty"`
	ProjectName string   `protobuf:"bytes,8,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
	VersionNum  int32    `protobuf:"varint,9,opt,name=version_num,json=versionNum" json:"version_num,omitempty"`
	VersionUid  []byte   `protobuf:"bytes,10,opt,name=version_uid,json=versionUid,proto3" json:"version_uid,omitempty"`
}

func (m *ServiceData) Reset()                    { *m = ServiceData{} }
func (m *ServiceData) String() string            { return proto.CompactTextString(m) }
func (*ServiceData) ProtoMessage()               {}
func (*ServiceData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ServiceVersionData struct {
	Uid        []byte                     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Title      string                     `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	VersionNum int32                      `protobuf:"varint,3,opt,name=version_num,json=versionNum" json:"version_num,omitempty"`
	ServiceUid []byte                     `protobuf:"bytes,4,opt,name=service_uid,json=serviceUid,proto3" json:"service_uid,omitempty"`
	Timestamp  int64                      `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	Inputs     []*vo.ServiceVersionInput  `protobuf:"bytes,10,rep,name=inputs" json:"inputs,omitempty"`
	Outputs    []*vo.ServiceVersionOutput `protobuf:"bytes,11,rep,name=outputs" json:"outputs,omitempty"`
}

func (m *ServiceVersionData) Reset()                    { *m = ServiceVersionData{} }
func (m *ServiceVersionData) String() string            { return proto.CompactTextString(m) }
func (*ServiceVersionData) ProtoMessage()               {}
func (*ServiceVersionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServiceVersionData) GetInputs() []*vo.ServiceVersionInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *ServiceVersionData) GetOutputs() []*vo.ServiceVersionOutput {
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
func (*DatasetItemData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Job struct {
	Uid        []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ProjectUid []byte   `protobuf:"bytes,2,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	Name       string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Title      string   `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Experts    [][]byte `protobuf:"bytes,7,rep,name=experts,proto3" json:"experts,omitempty"`
	RunCount   int32    `protobuf:"varint,8,opt,name=run_count,json=runCount" json:"run_count,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type JobRunData struct {
	Uid       []byte             `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	JobUid    []byte             `protobuf:"bytes,2,opt,name=job_uid,json=jobUid,proto3" json:"job_uid,omitempty"`
	Title     string             `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Inputs    []*vo.JobRunInput  `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	Outputs   []*vo.JobRunOutput `protobuf:"bytes,5,rep,name=outputs" json:"outputs,omitempty"`
	Timestamp int64              `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Status    vo.JOB_STATUS      `protobuf:"varint,7,opt,name=status,enum=JOB_STATUS" json:"status,omitempty"`
	RunNum    int32              `protobuf:"varint,8,opt,name=run_num,json=runNum" json:"run_num,omitempty"`
}

func (m *JobRunData) Reset()                    { *m = JobRunData{} }
func (m *JobRunData) String() string            { return proto.CompactTextString(m) }
func (*JobRunData) ProtoMessage()               {}
func (*JobRunData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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
	ServiceCount int32 `protobuf:"varint,7,opt,name=service_count,json=serviceCount" json:"service_count,omitempty"`
}

func (m *TenantStats) Reset()                    { *m = TenantStats{} }
func (m *TenantStats) String() string            { return proto.CompactTextString(m) }
func (*TenantStats) ProtoMessage()               {}
func (*TenantStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type AssetCache struct {
	Digest []byte `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	Body   []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *AssetCache) Reset()                    { *m = AssetCache{} }
func (m *AssetCache) String() string            { return proto.CompactTextString(m) }
func (*AssetCache) ProtoMessage()               {}
func (*AssetCache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type FieldSchema struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *FieldSchema) Reset()                    { *m = FieldSchema{} }
func (m *FieldSchema) String() string            { return proto.CompactTextString(m) }
func (*FieldSchema) ProtoMessage()               {}
func (*FieldSchema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type DatasetDescriptor struct {
	Type         string         `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	RecordCount  int64          `protobuf:"varint,2,opt,name=record_count,json=recordCount" json:"record_count,omitempty"`
	Schema       []*FieldSchema `protobuf:"bytes,3,rep,name=schema" json:"schema,omitempty"`
	SampleString string         `protobuf:"bytes,4,opt,name=sample_string,json=sampleString" json:"sample_string,omitempty"`
}

func (m *DatasetDescriptor) Reset()                    { *m = DatasetDescriptor{} }
func (m *DatasetDescriptor) String() string            { return proto.CompactTextString(m) }
func (*DatasetDescriptor) ProtoMessage()               {}
func (*DatasetDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

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
func (*AppVersionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func init() {
	proto.RegisterType((*EntityRef)(nil), "EntityRef")
	proto.RegisterType((*ProjectData)(nil), "ProjectData")
	proto.RegisterType((*Expert)(nil), "Expert")
	proto.RegisterType((*DatasetData)(nil), "DatasetData")
	proto.RegisterType((*DatasetVersionData)(nil), "DatasetVersionData")
	proto.RegisterType((*ServiceData)(nil), "ServiceData")
	proto.RegisterType((*ServiceVersionData)(nil), "ServiceVersionData")
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
	// 1345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x57, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x2e, 0x45, 0x89, 0x12, 0x87, 0xb4, 0x43, 0x6f, 0x1d, 0x87, 0x68, 0x1a, 0xc4, 0x51, 0x82,
	0x56, 0x6d, 0x03, 0x17, 0x70, 0x51, 0xa0, 0xe8, 0xa1, 0x80, 0x6c, 0x33, 0x88, 0x0c, 0x44, 0x36,
	0x96, 0x74, 0x90, 0x9c, 0x08, 0x4a, 0x5c, 0x3b, 0x74, 0x25, 0x92, 0x20, 0x97, 0x46, 0x7d, 0xeb,
	0x4b, 0xf4, 0xd0, 0x43, 0xdf, 0xa1, 0x6f, 0xd2, 0x17, 0xe8, 0x7b, 0xb4, 0x40, 0x0f, 0x2d, 0xf6,
	0x87, 0x14, 0x29, 0xd1, 0xcd, 0xa1, 0xe8, 0x8d, 0xfb, 0xcd, 0xec, 0x68, 0x66, 0xbe, 0x6f, 0x77,
	0x47, 0x30, 0x08, 0x67, 0x07, 0x69, 0x96, 0xd0, 0xe4, 0xa3, 0xc1, 0x4d, 0x22, 0xbe, 0x86, 0xdf,
	0x82, 0xee, 0xc4, 0x34, 0xa2, 0xb7, 0x98, 0x5c, 0xa2, 0x87, 0xd0, 0xfd, 0x3e, 0x8a, 0x43, 0x5b,
	0xd9, 0x57, 0x46, 0xdb, 0x87, 0xfd, 0x03, 0x67, 0xea, 0x4d, 0xbc, 0xb7, 0x98, 0x83, 0xc8, 0x02,
	0xb5, 0x88, 0x42, 0xbb, 0xb3, 0xaf, 0x8c, 0x4c, 0xcc, 0x3e, 0x87, 0x7f, 0x28, 0x60, 0x9c, 0x67,
	0xc9, 0x35, 0x99, 0xd3, 0x93, 0x80, 0x06, 0x68, 0x17, 0x7a, 0x34, 0xa2, 0x0b, 0xc2, 0xf7, 0xeb,
	0x58, 0x2c, 0x10, 0x82, 0x6e, 0x1c, 0x2c, 0x09, 0xdf, 0xa8, 0x63, 0xfe, 0x5d, 0xc6, 0x52, 0xab,
	0x58, 0xe8, 0x29, 0x6c, 0x85, 0x01, 0x0d, 0x72, 0x42, 0xfd, 0x79, 0x52, 0xc4, 0xd4, 0xee, 0xee,
	0x2b, 0xa3, 0x1e, 0x36, 0x25, 0x78, 0xcc, 0x30, 0xe6, 0x94, 0xd3, 0x24, 0x0b, 0xae, 0x88, 0x3f,
	0xbb, 0xa5, 0x24, 0xb7, 0x7b, 0xfb, 0xca, 0x48, 0xc5, 0xa6, 0x04, 0x8f, 0x18, 0x86, 0x1e, 0x82,
	0x7e, 0x9d, 0xcc, 0x64, 0x14, 0x8d, 0x47, 0x19, 0x5c, 0x27, 0x33, 0x11, 0x61, 0x1f, 0x8c, 0x90,
	0xe4, 0xf3, 0x2c, 0x4a, 0x69, 0x94, 0xc4, 0x76, 0x9f, 0xe7, 0x54, 0x87, 0xf8, 0x6f, 0x90, 0xec,
	0x26, 0x9a, 0x13, 0x19, 0x62, 0x20, 0x12, 0x91, 0x20, 0x0f, 0x33, 0x0c, 0x40, 0x73, 0x7e, 0x48,
	0x49, 0x46, 0xab, 0xea, 0x94, 0xcd, 0xea, 0x56, 0x9d, 0x62, 0x39, 0x5d, 0x16, 0x8b, 0x85, 0xcf,
	0x5d, 0x55, 0xee, 0x3a, 0x60, 0xc0, 0x94, 0xb9, 0xef, 0x42, 0x8f, 0x2c, 0x83, 0x68, 0xc1, 0x4b,
	0xd6, 0xb1, 0x58, 0x0c, 0x7f, 0xed, 0x82, 0x71, 0x22, 0x8a, 0xe7, 0xcd, 0x95, 0x41, 0x95, 0x55,
	0xd0, 0xb6, 0xc6, 0x56, 0x14, 0xa8, 0x75, 0x0a, 0x1e, 0x83, 0x91, 0x0a, 0x9e, 0x7c, 0x16, 0xa3,
	0xcb, 0x63, 0x80, 0x84, 0x2e, 0xa2, 0x10, 0x3d, 0x01, 0x33, 0x23, 0xf3, 0x24, 0x0b, 0x65, 0xcd,
	0xa2, 0xaf, 0x86, 0xc0, 0x44, 0xe7, 0x1e, 0x01, 0x5c, 0x46, 0x0b, 0x52, 0xeb, 0xab, 0x8a, 0x75,
	0x86, 0xdc, 0x41, 0x4d, 0xbf, 0x85, 0x9a, 0x4f, 0x40, 0xcb, 0x83, 0x65, 0xba, 0x20, 0xbc, 0xa9,
	0xc6, 0xe1, 0xf6, 0x81, 0xac, 0xd0, 0xe5, 0x28, 0x96, 0x56, 0xf4, 0x19, 0x58, 0x45, 0x1a, 0x06,
	0x94, 0xf8, 0x34, 0x5a, 0x92, 0x9c, 0x06, 0xcb, 0xd4, 0xd6, 0x79, 0xbc, 0x7b, 0x02, 0xf7, 0x4a,
	0x98, 0x95, 0xc6, 0x24, 0xe2, 0x5f, 0x26, 0xd9, 0x32, 0xa0, 0x36, 0xf0, 0xb2, 0x81, 0x41, 0x2f,
	0x38, 0xb2, 0xce, 0xb8, 0xb1, 0xc9, 0xf8, 0x63, 0x30, 0x16, 0xc9, 0x3c, 0x60, 0xdf, 0x7e, 0x14,
	0xda, 0xa6, 0x08, 0x51, 0x42, 0x13, 0xde, 0x9d, 0xca, 0xa1, 0xc8, 0x22, 0x7b, 0x4b, 0xc4, 0x28,
	0xb1, 0x8b, 0x2c, 0x42, 0x36, 0xf4, 0x09, 0x17, 0x44, 0x6e, 0x6f, 0xef, 0xab, 0x23, 0x13, 0x97,
	0x4b, 0xb6, 0xb9, 0xec, 0x3d, 0x67, 0x6b, 0x47, 0x6c, 0x96, 0x18, 0x17, 0xc0, 0x13, 0x30, 0xdf,
	0x91, 0x20, 0xf4, 0x6f, 0x48, 0x96, 0xb3, 0x1c, 0x77, 0x39, 0x3f, 0x06, 0xc3, 0x5e, 0x0b, 0x88,
	0xb5, 0x57, 0x5a, 0x25, 0x01, 0x7b, 0x42, 0x95, 0x12, 0x14, 0xaa, 0xfc, 0x45, 0x05, 0x24, 0x1b,
	0x2a, 0xf7, 0xdd, 0xa1, 0x9c, 0x47, 0x00, 0x69, 0x90, 0x91, 0x58, 0xc8, 0x41, 0xe8, 0x54, 0x17,
	0x08, 0x53, 0x43, 0xbb, 0x88, 0xd6, 0x35, 0xd2, 0x6d, 0xd5, 0x48, 0x44, 0xc9, 0xb2, 0x21, 0x22,
	0x9d, 0x21, 0x77, 0x68, 0x44, 0x6b, 0xd1, 0xc8, 0x08, 0xb4, 0x28, 0x4e, 0x0b, 0xca, 0x14, 0xa4,
	0x8e, 0x8c, 0x43, 0xeb, 0x60, 0x55, 0xd2, 0x84, 0x19, 0xb0, 0xb4, 0xa3, 0x21, 0xf4, 0x58, 0xec,
	0xdc, 0x1e, 0x70, 0x47, 0xb3, 0x74, 0x9c, 0x50, 0xb2, 0xc4, 0xc2, 0x54, 0xca, 0x83, 0x5d, 0x2b,
	0xac, 0x54, 0x5d, 0x28, 0x5f, 0x42, 0xac, 0xd6, 0x2f, 0xa0, 0x9f, 0x14, 0x94, 0xff, 0x1e, 0xf0,
	0x30, 0x3b, 0xb5, 0xdf, 0x3b, 0xe3, 0x16, 0x5c, 0x7a, 0xb0, 0x68, 0x25, 0x0b, 0x71, 0xb1, 0xe4,
	0x5a, 0xea, 0x61, 0x90, 0xd0, 0xb4, 0x58, 0xa2, 0x8f, 0x41, 0x5f, 0x29, 0xd6, 0x14, 0xf5, 0x57,
	0xc0, 0xf0, 0xb7, 0x0e, 0x18, 0xae, 0xb8, 0x46, 0xfe, 0xff, 0x23, 0xbd, 0xa6, 0xfb, 0xde, 0x7b,
	0x75, 0xaf, 0xbd, 0x57, 0xf7, 0xfd, 0xff, 0xa2, 0xfb, 0xc1, 0xa6, 0xee, 0xd7, 0xda, 0xa9, 0x6f,
	0xb4, 0xb3, 0xe6, 0xc0, 0x8a, 0x04, 0x51, 0xa4, 0x84, 0x2e, 0xa2, 0x70, 0xf8, 0x97, 0x02, 0x48,
	0x76, 0xf4, 0xdf, 0x15, 0x5f, 0x35, 0xb1, 0xb3, 0xd6, 0xc4, 0x7a, 0x02, 0x6a, 0x5b, 0x02, 0xe5,
	0x63, 0x50, 0xeb, 0xb2, 0x84, 0x58, 0x97, 0x1b, 0x84, 0xf7, 0xd6, 0x08, 0x47, 0xcf, 0x2b, 0x2d,
	0x0b, 0x6d, 0xed, 0x1e, 0x34, 0x93, 0x6d, 0xea, 0xf9, 0xcb, 0x95, 0x14, 0x0d, 0xee, 0x7e, 0x7f,
	0xcd, 0x7d, 0x4d, 0x8e, 0xc3, 0x04, 0xee, 0xd5, 0x24, 0xcf, 0x2b, 0x6f, 0x13, 0xd0, 0xc6, 0xb1,
	0x53, 0x5b, 0x8e, 0xdd, 0x53, 0xd8, 0x12, 0x27, 0x39, 0x6f, 0x1c, 0x6f, 0x79, 0xe4, 0x73, 0x71,
	0xc1, 0xfc, 0xac, 0x80, 0x7a, 0x9a, 0xcc, 0x5a, 0xfa, 0xbb, 0x26, 0xc7, 0xce, 0x86, 0x1c, 0xcb,
	0xc4, 0xd4, 0x36, 0x65, 0x77, 0xeb, 0xa4, 0xd4, 0x24, 0xd5, 0x6f, 0x4a, 0xea, 0x21, 0xe8, 0x59,
	0x11, 0x37, 0x9e, 0xe5, 0x41, 0x56, 0xc8, 0xcb, 0xef, 0x4f, 0x05, 0xe0, 0x34, 0x99, 0xe1, 0xe2,
	0x2e, 0x09, 0x3c, 0x80, 0x3e, 0x9b, 0x0b, 0x56, 0xe9, 0x69, 0xd7, 0xc9, 0xec, 0xee, 0xeb, 0xee,
	0x59, 0xc5, 0x5d, 0x57, 0x5e, 0x2f, 0x22, 0x7a, 0x93, 0xb3, 0x4f, 0x57, 0x9c, 0xf5, 0xb8, 0xdb,
	0x96, 0x74, 0x5b, 0xbf, 0x3a, 0x1a, 0x42, 0xd1, 0xd6, 0x85, 0xf2, 0x14, 0xb4, 0x9c, 0x06, 0xb4,
	0x10, 0xcf, 0xe6, 0xf6, 0xa1, 0x71, 0x70, 0x7a, 0x76, 0xe4, 0xbb, 0xde, 0xd8, 0xbb, 0x70, 0xb1,
	0x34, 0xb1, 0x02, 0x58, 0xf9, 0x4c, 0xa9, 0xa2, 0x78, 0x2d, 0x2b, 0x98, 0x4a, 0x87, 0x3f, 0x76,
	0xc0, 0xf0, 0x48, 0x1c, 0xc4, 0xd4, 0xa5, 0x01, 0xe5, 0x5c, 0x96, 0x64, 0x88, 0x5e, 0x29, 0xe2,
	0xb1, 0x90, 0x60, 0x75, 0x19, 0x37, 0x07, 0xae, 0x4e, 0xcb, 0xc0, 0xd5, 0x98, 0xa5, 0xd4, 0xb5,
	0x59, 0x8a, 0x3f, 0x08, 0x69, 0x92, 0x35, 0x27, 0x36, 0x43, 0x60, 0x95, 0x8b, 0x20, 0xaf, 0xf6,
	0x24, 0xf4, 0xb0, 0x21, 0x30, 0xe1, 0xf2, 0x18, 0x8c, 0x65, 0x12, 0x92, 0x45, 0x63, 0x60, 0x03,
	0x0e, 0xad, 0x5e, 0x8d, 0xc6, 0x40, 0xd6, 0x6f, 0x19, 0xc8, 0xbe, 0x01, 0x18, 0xe7, 0x2c, 0xed,
	0x60, 0xfe, 0x8e, 0xa0, 0x3d, 0xd0, 0xc2, 0xe8, 0x8a, 0xe4, 0x54, 0xf2, 0x2f, 0x57, 0x4c, 0x84,
	0xb3, 0x24, 0xbc, 0x95, 0xfc, 0xf3, 0xef, 0xe1, 0xd7, 0x60, 0xbc, 0x88, 0xc8, 0x22, 0x74, 0xe7,
	0xef, 0xc8, 0x32, 0x68, 0x9d, 0xe7, 0x10, 0x74, 0xe9, 0x6d, 0x5a, 0x1d, 0x2a, 0xf6, 0x3d, 0xfc,
	0x49, 0x81, 0x9d, 0x72, 0x3c, 0x93, 0x77, 0x6a, 0x92, 0x55, 0x9e, 0xca, 0xca, 0x73, 0xe3, 0xdd,
	0xec, 0x6c, 0xbe, 0x9b, 0xcf, 0x40, 0xcb, 0xf9, 0xcf, 0xdb, 0xaa, 0xd4, 0x5a, 0x2d, 0x25, 0x2c,
	0x6d, 0xbc, 0x11, 0x7c, 0x3e, 0xf2, 0x73, 0x9a, 0x45, 0xf1, 0x95, 0x3c, 0x36, 0xa6, 0x00, 0x5d,
	0x8e, 0x0d, 0xbf, 0x83, 0xed, 0x71, 0x9a, 0xd6, 0x2f, 0xc3, 0xe7, 0xb0, 0x23, 0x87, 0xf4, 0x28,
	0x89, 0xa5, 0x41, 0x26, 0xb8, 0x69, 0xf8, 0xfc, 0xf7, 0x0e, 0xf4, 0x70, 0x10, 0x5f, 0x11, 0x04,
	0xa0, 0x39, 0x37, 0x24, 0xa6, 0xb9, 0xf5, 0x01, 0x32, 0x61, 0x70, 0x8e, 0xcf, 0x4e, 0x9d, 0x63,
	0xcf, 0xb5, 0x14, 0xb6, 0x3a, 0x19, 0x7b, 0x63, 0xd7, 0xf1, 0x5c, 0xab, 0x83, 0x06, 0xd0, 0x3d,
	0x3d, 0x3b, 0x72, 0xad, 0x2e, 0xd2, 0xa1, 0xc7, 0x24, 0xeb, 0x5a, 0x3d, 0x74, 0x0f, 0x8c, 0xb1,
	0xeb, 0x3a, 0x9e, 0x7f, 0x3c, 0x3e, 0x7e, 0xe9, 0x58, 0x1a, 0x32, 0xa0, 0xef, 0xbc, 0x39, 0x77,
	0xb0, 0xe7, 0x5a, 0x7d, 0xb4, 0x0b, 0x96, 0x0c, 0xe0, 0xbf, 0x76, 0xb0, 0x3b, 0x39, 0x9b, 0xba,
	0xd6, 0x00, 0xed, 0x01, 0x2a, 0xd1, 0x97, 0xce, 0xf8, 0xc4, 0x9f, 0x78, 0xce, 0x2b, 0xd7, 0xd2,
	0xd9, 0xcf, 0xb9, 0x0e, 0x7e, 0x3d, 0x39, 0x76, 0x5c, 0x6b, 0x8b, 0xed, 0x95, 0xab, 0xd5, 0xde,
	0x6d, 0xf4, 0x00, 0x3e, 0x2c, 0xf7, 0x1e, 0x4d, 0xa6, 0x63, 0xfc, 0xd6, 0xf7, 0xb0, 0xe3, 0x58,
	0x7b, 0x68, 0x07, 0xb6, 0x64, 0xe6, 0xfe, 0x74, 0xfc, 0xca, 0x71, 0x2d, 0x60, 0xf1, 0xd8, 0xe9,
	0xc2, 0x17, 0x53, 0xd7, 0x32, 0x91, 0x05, 0x26, 0x37, 0xf8, 0xee, 0xf1, 0xd9, 0xb9, 0x73, 0x62,
	0x19, 0xc8, 0x86, 0xdd, 0xc9, 0xc9, 0x1b, 0xbf, 0xdc, 0x56, 0x95, 0xba, 0x5b, 0x5a, 0x36, 0x72,
	0xbf, 0x5f, 0x5a, 0x36, 0x32, 0x7b, 0x80, 0x4c, 0xe8, 0xcb, 0x95, 0xf5, 0xb7, 0x32, 0xd3, 0xf8,
	0xbf, 0xae, 0xaf, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x81, 0xe4, 0x91, 0x78, 0x8b, 0x0d, 0x00,
	0x00,
}
