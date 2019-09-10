// Code generated by protoc-gen-go.
// source: db.proto
// DO NOT EDIT!

/*
Package db is a generated protocol buffer package.

It is generated from these files:
	db.proto

It has these top-level messages:
	SystemLink
	EntityRef
	ProjectData
	Expert
	DatasetData
	DatasetVersionData
	SystemData
	SystemVersionData
	DatasetItemData
	Job
	JobRunData
	JobRunOutput
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
	Range_SYSTEMS             Range = 13
	Range_SYSTEM_VERSIONS     Range = 14
	Range_SYSTEM_LINKS        Range = 30
	Range_DATASET_BINARY_TREE Range = 22
	Range_PROJECT_NAMES       Range = 10
	Range_JOB_RUNS            Range = 12
	// scoped to project
	Range_NAMES_SCOPED         Range = 11
	Range_IDX_PROJECT_DATASETS Range = 20
	Range_IDX_DATASET_VERSIONS Range = 21
	// service UID/int32
	Range_IDX_SYSTEM_VERSIONS Range = 23
	Range_IDX_PROJECT_SYSTEMS Range = 24
	Range_IDX_PROJECT_JOBS    Range = 25
	Range_VERSION             Range = 255
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
	13:  "SYSTEMS",
	14:  "SYSTEM_VERSIONS",
	30:  "SYSTEM_LINKS",
	22:  "DATASET_BINARY_TREE",
	10:  "PROJECT_NAMES",
	12:  "JOB_RUNS",
	11:  "NAMES_SCOPED",
	20:  "IDX_PROJECT_DATASETS",
	21:  "IDX_DATASET_VERSIONS",
	23:  "IDX_SYSTEM_VERSIONS",
	24:  "IDX_PROJECT_SYSTEMS",
	25:  "IDX_PROJECT_JOBS",
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
	"SYSTEMS":              13,
	"SYSTEM_VERSIONS":      14,
	"SYSTEM_LINKS":         30,
	"DATASET_BINARY_TREE":  22,
	"PROJECT_NAMES":        10,
	"JOB_RUNS":             12,
	"NAMES_SCOPED":         11,
	"IDX_PROJECT_DATASETS": 20,
	"IDX_DATASET_VERSIONS": 21,
	"IDX_SYSTEM_VERSIONS":  23,
	"IDX_PROJECT_SYSTEMS":  24,
	"IDX_PROJECT_JOBS":     25,
	"VERSION":              255,
}

func (x Range) String() string {
	return proto.EnumName(Range_name, int32(x))
}
func (Range) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SystemLink_Type int32

const (
	SystemLink_None             SystemLink_Type = 0
	SystemLink_Input_SystemVer  SystemLink_Type = 1
	SystemLink_Input_JobRun     SystemLink_Type = 2
	SystemLink_Output_SystemVer SystemLink_Type = 3
	SystemLink_Output_JobRun    SystemLink_Type = 4
)

var SystemLink_Type_name = map[int32]string{
	0: "None",
	1: "Input_SystemVer",
	2: "Input_JobRun",
	3: "Output_SystemVer",
	4: "Output_JobRun",
}
var SystemLink_Type_value = map[string]int32{
	"None":             0,
	"Input_SystemVer":  1,
	"Input_JobRun":     2,
	"Output_SystemVer": 3,
	"Output_JobRun":    4,
}

func (x SystemLink_Type) String() string {
	return proto.EnumName(SystemLink_Type_name, int32(x))
}
func (SystemLink_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type JobRunOutput_Type int32

const (
	JobRunOutput_None       JobRunOutput_Type = 0
	JobRunOutput_DatasetVer JobRunOutput_Type = 1
	JobRunOutput_SystemVer  JobRunOutput_Type = 2
)

var JobRunOutput_Type_name = map[int32]string{
	0: "None",
	1: "DatasetVer",
	2: "SystemVer",
}
var JobRunOutput_Type_value = map[string]int32{
	"None":       0,
	"DatasetVer": 1,
	"SystemVer":  2,
}

func (x JobRunOutput_Type) String() string {
	return proto.EnumName(JobRunOutput_Type_name, int32(x))
}
func (JobRunOutput_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{11, 0} }

// serviceID | Entity Container  | Exact Reference
type SystemLink struct {
	InstanceUid  []byte          `protobuf:"bytes,1,opt,name=instance_uid,json=instanceUid,proto3" json:"instance_uid,omitempty"`
	ContainerUid []byte          `protobuf:"bytes,2,opt,name=container_uid,json=containerUid,proto3" json:"container_uid,omitempty"`
	Type         SystemLink_Type `protobuf:"varint,3,opt,name=type,enum=SystemLink_Type" json:"type,omitempty"`
}

func (m *SystemLink) Reset()                    { *m = SystemLink{} }
func (m *SystemLink) String() string            { return proto.CompactTextString(m) }
func (*SystemLink) ProtoMessage()               {}
func (*SystemLink) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	SystemCount  int32  `protobuf:"varint,8,opt,name=system_count,json=systemCount" json:"system_count,omitempty"`
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
	ProjectName     string            `protobuf:"bytes,17,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
	HeadVersion     []byte            `protobuf:"bytes,20,opt,name=head_version,json=headVersion,proto3" json:"head_version,omitempty"`
	VersionCount    int32             `protobuf:"varint,22,opt,name=version_count,json=versionCount" json:"version_count,omitempty"`
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
func (*DatasetVersionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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

type SystemData struct {
	Uid              []byte        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name             string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Title            string        `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	ProjectUid       []byte        `protobuf:"bytes,4,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	Description      string        `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	LocationId       string        `protobuf:"bytes,6,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	LocationUri      string        `protobuf:"bytes,7,opt,name=location_uri,json=locationUri" json:"location_uri,omitempty"`
	Experts          [][]byte      `protobuf:"bytes,14,rep,name=experts,proto3" json:"experts,omitempty"`
	ProjectName      string        `protobuf:"bytes,8,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
	VersionNum       int32         `protobuf:"varint,9,opt,name=version_num,json=versionNum" json:"version_num,omitempty"`
	VersionUid       []byte        `protobuf:"bytes,10,opt,name=version_uid,json=versionUid,proto3" json:"version_uid,omitempty"`
	VersionTimestamp int64         `protobuf:"varint,11,opt,name=version_timestamp,json=versionTimestamp" json:"version_timestamp,omitempty"`
	Kind             vo.SystemKind `protobuf:"varint,12,opt,name=kind,enum=SystemKind" json:"kind,omitempty"`
}

func (m *SystemData) Reset()                    { *m = SystemData{} }
func (m *SystemData) String() string            { return proto.CompactTextString(m) }
func (*SystemData) ProtoMessage()               {}
func (*SystemData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type SystemVersionData struct {
	Uid        []byte                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Title      string                    `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	VersionNum int32                     `protobuf:"varint,3,opt,name=version_num,json=versionNum" json:"version_num,omitempty"`
	SystemUid  []byte                    `protobuf:"bytes,4,opt,name=system_uid,json=systemUid,proto3" json:"system_uid,omitempty"`
	Timestamp  int64                     `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	Inputs     []*vo.SystemVersionInput  `protobuf:"bytes,10,rep,name=inputs" json:"inputs,omitempty"`
	Outputs    []*vo.SystemVersionOutput `protobuf:"bytes,11,rep,name=outputs" json:"outputs,omitempty"`
}

func (m *SystemVersionData) Reset()                    { *m = SystemVersionData{} }
func (m *SystemVersionData) String() string            { return proto.CompactTextString(m) }
func (*SystemVersionData) ProtoMessage()               {}
func (*SystemVersionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SystemVersionData) GetInputs() []*vo.SystemVersionInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *SystemVersionData) GetOutputs() []*vo.SystemVersionOutput {
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
func (*DatasetItemData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type Job struct {
	Uid            []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ProjectUid     []byte   `protobuf:"bytes,2,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	Name           string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Title          string   `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Experts        [][]byte `protobuf:"bytes,7,rep,name=experts,proto3" json:"experts,omitempty"`
	RunNum         int32    `protobuf:"varint,8,opt,name=run_num,json=runNum" json:"run_num,omitempty"`
	RunUid         []byte   `protobuf:"bytes,10,opt,name=run_uid,json=runUid,proto3" json:"run_uid,omitempty"`
	ProjectName    string   `protobuf:"bytes,9,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
	LastSuccessUid []byte   `protobuf:"bytes,11,opt,name=last_success_uid,json=lastSuccessUid,proto3" json:"last_success_uid,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type JobRunData struct {
	Uid             []byte            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	JobUid          []byte            `protobuf:"bytes,2,opt,name=job_uid,json=jobUid,proto3" json:"job_uid,omitempty"`
	Title           string            `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Inputs          []*vo.JobRunInput `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	Outputs         []*JobRunOutput   `protobuf:"bytes,5,rep,name=outputs" json:"outputs,omitempty"`
	StartTimestamp  int64             `protobuf:"varint,6,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	UpdateTimestamp int64             `protobuf:"varint,9,opt,name=update_timestamp,json=updateTimestamp" json:"update_timestamp,omitempty"`
	Status          vo.JOB_STATUS     `protobuf:"varint,7,opt,name=status,enum=JOB_STATUS" json:"status,omitempty"`
	RunNum          int32             `protobuf:"varint,8,opt,name=run_num,json=runNum" json:"run_num,omitempty"`
}

func (m *JobRunData) Reset()                    { *m = JobRunData{} }
func (m *JobRunData) String() string            { return proto.CompactTextString(m) }
func (*JobRunData) ProtoMessage()               {}
func (*JobRunData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *JobRunData) GetInputs() []*vo.JobRunInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *JobRunData) GetOutputs() []*JobRunOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type JobRunOutput struct {
	// dataset
	Uid  []byte            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type JobRunOutput_Type `protobuf:"varint,2,opt,name=type,enum=JobRunOutput_Type" json:"type,omitempty"`
}

func (m *JobRunOutput) Reset()                    { *m = JobRunOutput{} }
func (m *JobRunOutput) String() string            { return proto.CompactTextString(m) }
func (*JobRunOutput) ProtoMessage()               {}
func (*JobRunOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type TenantStats struct {
	ProjectCount int32 `protobuf:"varint,1,opt,name=project_count,json=projectCount" json:"project_count,omitempty"`
	DatasetCount int32 `protobuf:"varint,2,opt,name=dataset_count,json=datasetCount" json:"dataset_count,omitempty"`
	JobCount     int32 `protobuf:"varint,3,opt,name=job_count,json=jobCount" json:"job_count,omitempty"`
	ReportCount  int32 `protobuf:"varint,4,opt,name=report_count,json=reportCount" json:"report_count,omitempty"`
	ExpertCount  int32 `protobuf:"varint,5,opt,name=expert_count,json=expertCount" json:"expert_count,omitempty"`
	ModelCount   int32 `protobuf:"varint,6,opt,name=model_count,json=modelCount" json:"model_count,omitempty"`
	SystemCount  int32 `protobuf:"varint,7,opt,name=system_count,json=systemCount" json:"system_count,omitempty"`
}

func (m *TenantStats) Reset()                    { *m = TenantStats{} }
func (m *TenantStats) String() string            { return proto.CompactTextString(m) }
func (*TenantStats) ProtoMessage()               {}
func (*TenantStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type AssetCache struct {
	Digest []byte `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	Body   []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *AssetCache) Reset()                    { *m = AssetCache{} }
func (m *AssetCache) String() string            { return proto.CompactTextString(m) }
func (*AssetCache) ProtoMessage()               {}
func (*AssetCache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type FieldSchema struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *FieldSchema) Reset()                    { *m = FieldSchema{} }
func (m *FieldSchema) String() string            { return proto.CompactTextString(m) }
func (*FieldSchema) ProtoMessage()               {}
func (*FieldSchema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type DatasetDescriptor struct {
	Type         string         `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	RecordCount  int64          `protobuf:"varint,2,opt,name=record_count,json=recordCount" json:"record_count,omitempty"`
	Schema       []*FieldSchema `protobuf:"bytes,3,rep,name=schema" json:"schema,omitempty"`
	SampleString string         `protobuf:"bytes,4,opt,name=sample_string,json=sampleString" json:"sample_string,omitempty"`
}

func (m *DatasetDescriptor) Reset()                    { *m = DatasetDescriptor{} }
func (m *DatasetDescriptor) String() string            { return proto.CompactTextString(m) }
func (*DatasetDescriptor) ProtoMessage()               {}
func (*DatasetDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

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
func (*AppVersionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func init() {
	proto.RegisterType((*SystemLink)(nil), "SystemLink")
	proto.RegisterType((*EntityRef)(nil), "EntityRef")
	proto.RegisterType((*ProjectData)(nil), "ProjectData")
	proto.RegisterType((*Expert)(nil), "Expert")
	proto.RegisterType((*DatasetData)(nil), "DatasetData")
	proto.RegisterType((*DatasetVersionData)(nil), "DatasetVersionData")
	proto.RegisterType((*SystemData)(nil), "SystemData")
	proto.RegisterType((*SystemVersionData)(nil), "SystemVersionData")
	proto.RegisterType((*DatasetItemData)(nil), "DatasetItemData")
	proto.RegisterType((*Job)(nil), "Job")
	proto.RegisterType((*JobRunData)(nil), "JobRunData")
	proto.RegisterType((*JobRunOutput)(nil), "JobRunOutput")
	proto.RegisterType((*TenantStats)(nil), "TenantStats")
	proto.RegisterType((*AssetCache)(nil), "AssetCache")
	proto.RegisterType((*FieldSchema)(nil), "FieldSchema")
	proto.RegisterType((*DatasetDescriptor)(nil), "DatasetDescriptor")
	proto.RegisterType((*AppVersionData)(nil), "AppVersionData")
	proto.RegisterEnum("Range", Range_name, Range_value)
	proto.RegisterEnum("SystemLink_Type", SystemLink_Type_name, SystemLink_Type_value)
	proto.RegisterEnum("JobRunOutput_Type", JobRunOutput_Type_name, JobRunOutput_Type_value)
}

func init() { proto.RegisterFile("db.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x57, 0xdd, 0x6e, 0xdb, 0x56,
	0x12, 0x0e, 0x45, 0xfd, 0x71, 0x48, 0xcb, 0xf4, 0x89, 0xd7, 0xe6, 0x6e, 0x36, 0x1b, 0x87, 0x09,
	0x12, 0x6d, 0x53, 0xa8, 0x80, 0x8b, 0x02, 0x45, 0x2f, 0x0a, 0xc8, 0x36, 0x83, 0xc8, 0x49, 0x64,
	0x83, 0x94, 0x83, 0xe4, 0x8a, 0xa5, 0xc4, 0x63, 0x87, 0x8e, 0x44, 0x0a, 0xe4, 0x61, 0x50, 0x3d,
	0x48, 0x6f, 0x8a, 0xf6, 0x15, 0x8a, 0x3e, 0x50, 0x5f, 0xa3, 0xbd, 0x29, 0xd0, 0x16, 0xe7, 0x87,
	0x7f, 0xa2, 0xdc, 0xa0, 0x28, 0x7a, 0x47, 0x7e, 0x33, 0x67, 0x78, 0x66, 0xe6, 0x9b, 0x1f, 0x42,
	0xd7, 0x9f, 0x0e, 0x96, 0x71, 0x44, 0xa2, 0xff, 0x74, 0xdf, 0x47, 0xfc, 0xc9, 0xfc, 0x49, 0x02,
	0x70, 0x56, 0x09, 0xc1, 0x8b, 0x17, 0x41, 0xf8, 0x0e, 0xdd, 0x07, 0x2d, 0x08, 0x13, 0xe2, 0x85,
	0x33, 0xec, 0xa6, 0x81, 0x6f, 0x48, 0x07, 0x52, 0x5f, 0xb3, 0xd5, 0x0c, 0xbb, 0x08, 0x7c, 0xf4,
	0x00, 0xb6, 0x66, 0x51, 0x48, 0xbc, 0x20, 0xc4, 0x31, 0xd3, 0x69, 0x30, 0x1d, 0x2d, 0x07, 0xa9,
	0xd2, 0x43, 0x68, 0x92, 0xd5, 0x12, 0x1b, 0xf2, 0x81, 0xd4, 0xef, 0x1d, 0xea, 0x83, 0xe2, 0x13,
	0x83, 0xc9, 0x6a, 0x89, 0x6d, 0x26, 0x35, 0xbf, 0x82, 0x26, 0x7d, 0x43, 0x5d, 0x68, 0x8e, 0xa3,
	0x10, 0xeb, 0xb7, 0xd0, 0x6d, 0xd8, 0x1e, 0x85, 0xcb, 0x94, 0xb8, 0xfc, 0xc0, 0x2b, 0x1c, 0xeb,
	0x12, 0xd2, 0x41, 0xe3, 0xe0, 0x69, 0x34, 0xb5, 0xd3, 0x50, 0x6f, 0xa0, 0x5d, 0xd0, 0xcf, 0x52,
	0x52, 0xd5, 0x93, 0xd1, 0x0e, 0x6c, 0x09, 0x54, 0x28, 0x36, 0xcd, 0x2f, 0x40, 0xb1, 0x42, 0x12,
	0x90, 0x95, 0x8d, 0x2f, 0xd1, 0x1d, 0x68, 0xbe, 0x0b, 0x42, 0xee, 0x54, 0xef, 0xb0, 0x33, 0xb0,
	0xc6, 0x93, 0xd1, 0xe4, 0x8d, 0xcd, 0x40, 0xa4, 0x83, 0x5c, 0x38, 0x43, 0x1f, 0xcd, 0x9f, 0x25,
	0x50, 0xcf, 0xe3, 0xe8, 0x1a, 0xcf, 0xc8, 0x89, 0x47, 0x3c, 0xb4, 0x0b, 0x2d, 0x12, 0x90, 0x39,
	0x66, 0xe7, 0x15, 0x9b, 0xbf, 0x20, 0x04, 0xcd, 0xd0, 0x5b, 0x60, 0x76, 0x50, 0xb1, 0xd9, 0x73,
	0x66, 0x4b, 0xce, 0x6d, 0xd1, 0xa0, 0xf9, 0x1e, 0xf1, 0x12, 0x4c, 0xdc, 0x59, 0x94, 0x86, 0xc4,
	0x68, 0x1e, 0x48, 0xfd, 0x96, 0xad, 0x09, 0xf0, 0x98, 0x62, 0x54, 0x29, 0x21, 0x51, 0xec, 0x5d,
	0x61, 0x77, 0xba, 0x22, 0x38, 0x31, 0x5a, 0x07, 0x52, 0x5f, 0xb6, 0x35, 0x01, 0x1e, 0x51, 0x0c,
	0xdd, 0x01, 0xe5, 0x3a, 0x9a, 0x0a, 0x2b, 0x6d, 0x66, 0xa5, 0x7b, 0x1d, 0x4d, 0xb9, 0x85, 0x03,
	0x50, 0x7d, 0x9c, 0xcc, 0xe2, 0x60, 0x49, 0x82, 0x28, 0x34, 0x3a, 0xec, 0x4e, 0x65, 0x88, 0x26,
	0x38, 0x61, 0x21, 0x13, 0x16, 0xba, 0xcc, 0x82, 0xca, 0x31, 0x66, 0xc4, 0xf4, 0xa0, 0x6d, 0x7d,
	0xbd, 0xc4, 0x31, 0xc9, 0x7d, 0x93, 0xea, 0xbe, 0x15, 0x71, 0xa2, 0x37, 0xba, 0x4c, 0xe7, 0x73,
	0x97, 0xa9, 0xca, 0x4c, 0xb5, 0x4b, 0x81, 0x31, 0x55, 0xdf, 0x85, 0x16, 0x5e, 0x78, 0xc1, 0x9c,
	0x39, 0xac, 0xd8, 0xfc, 0xc5, 0xfc, 0xb1, 0x09, 0xea, 0x09, 0x77, 0x9d, 0x85, 0x56, 0x18, 0x95,
	0x0a, 0xa3, 0x9b, 0xc2, 0x9a, 0x27, 0x40, 0x2e, 0x27, 0xe0, 0x1e, 0xa8, 0x4b, 0x9e, 0x25, 0xc6,
	0xc6, 0x26, 0xb3, 0x01, 0x02, 0xa2, 0x5c, 0xbc, 0x0f, 0x5a, 0x8c, 0x67, 0x51, 0xec, 0x0b, 0x97,
	0x79, 0x54, 0x55, 0x8e, 0xf1, 0xb8, 0xdd, 0x05, 0xb8, 0x0c, 0xe6, 0xb8, 0x14, 0x55, 0xd9, 0x56,
	0x28, 0x72, 0x43, 0x62, 0x3a, 0x1b, 0x12, 0xf3, 0x08, 0xda, 0x89, 0xb7, 0x58, 0xce, 0x31, 0x8b,
	0xa9, 0x7a, 0xd8, 0x1b, 0x08, 0x0f, 0x1d, 0x86, 0xda, 0x42, 0x8a, 0xfe, 0x0f, 0x7a, 0xba, 0xf4,
	0x3d, 0x82, 0x5d, 0x12, 0x2c, 0x70, 0x42, 0xbc, 0xc5, 0xd2, 0x50, 0x98, 0xbd, 0x6d, 0x8e, 0x4f,
	0x32, 0x98, 0xba, 0x46, 0x09, 0xe2, 0x5e, 0x46, 0xf1, 0xc2, 0x23, 0x06, 0x30, 0xb7, 0x81, 0x42,
	0x4f, 0x19, 0xb2, 0x9e, 0x6f, 0xb5, 0x9e, 0xef, 0x7b, 0xa0, 0xce, 0xa3, 0x99, 0x47, 0x9f, 0xdd,
	0xc0, 0x37, 0x34, 0x6e, 0x22, 0x83, 0x46, 0x2c, 0x3a, 0xb9, 0x42, 0x1a, 0x07, 0xc6, 0x16, 0xb7,
	0x91, 0x61, 0x17, 0x71, 0x80, 0x0c, 0xe8, 0x60, 0x46, 0x88, 0xc4, 0xe8, 0x1d, 0xc8, 0x7d, 0xcd,
	0xce, 0x5e, 0xe9, 0xe1, 0x2c, 0xf6, 0x2c, 0x5b, 0x3b, 0xfc, 0xb0, 0xc0, 0x18, 0x01, 0xee, 0x83,
	0xf6, 0x16, 0x7b, 0xbe, 0xfb, 0x1e, 0xc7, 0x09, 0xbd, 0xe3, 0x2e, 0xef, 0x28, 0x14, 0x7b, 0xc5,
	0x21, 0x1a, 0x5e, 0x21, 0x15, 0x09, 0xd8, 0xe3, 0xc5, 0x21, 0x40, 0xce, 0xca, 0xef, 0x65, 0x40,
	0x22, 0xa0, 0xe2, 0xdc, 0x0d, 0xcc, 0xb9, 0x0b, 0xb0, 0xf4, 0x62, 0x1c, 0x92, 0x52, 0x73, 0x52,
	0x38, 0x42, 0xd9, 0xb0, 0x99, 0x44, 0xeb, 0x1c, 0x69, 0x6e, 0xe4, 0x48, 0x50, 0xd4, 0x0d, 0x27,
	0x91, 0x12, 0x64, 0x55, 0x53, 0xe7, 0x48, 0x7b, 0x03, 0x47, 0xfa, 0xd0, 0x0e, 0x68, 0x27, 0xa3,
	0x0c, 0x92, 0xfb, 0xea, 0xa1, 0x3e, 0x28, 0x5c, 0x62, 0x2d, 0xce, 0x16, 0x72, 0x64, 0x42, 0x8b,
	0xda, 0x4e, 0x8c, 0x2e, 0x53, 0xd4, 0x32, 0xc5, 0x11, 0xc1, 0x0b, 0x9b, 0x8b, 0x32, 0x7a, 0xd0,
	0xa6, 0x42, 0x5d, 0x55, 0x38, 0xf3, 0x05, 0x44, 0x7d, 0x7d, 0x02, 0x9d, 0x88, 0x35, 0xc4, 0xc4,
	0x00, 0x66, 0x66, 0xa7, 0xf4, 0x3d, 0xde, 0x2a, 0xed, 0x4c, 0x83, 0x5a, 0xcb, 0xb2, 0x10, 0xa6,
	0x0b, 0xc6, 0xa5, 0x96, 0x0d, 0x02, 0x1a, 0xa7, 0x0b, 0xf4, 0x5f, 0x50, 0x0a, 0xc6, 0x6a, 0xdc,
	0xff, 0x1c, 0x30, 0xbf, 0x93, 0xb3, 0x41, 0xf2, 0xcf, 0x57, 0xf4, 0x1a, 0xed, 0x5b, 0x1f, 0xa4,
	0x7d, 0xfb, 0x83, 0xb4, 0xef, 0xfc, 0x1d, 0xda, 0x77, 0xeb, 0xb4, 0x5f, 0x8b, 0xa6, 0x52, 0x8b,
	0x66, 0x49, 0x81, 0x3a, 0x09, 0xdc, 0x49, 0x01, 0xf1, 0xe4, 0xed, 0x64, 0x0a, 0x45, 0xd8, 0x55,
	0x16, 0x76, 0x5d, 0x08, 0xca, 0x9d, 0x82, 0x8f, 0x36, 0x8d, 0x8d, 0x36, 0x55, 0xcc, 0xdb, 0xe7,
	0x41, 0xe8, 0xf3, 0xf1, 0x66, 0xfe, 0x22, 0xc1, 0x4e, 0x3e, 0x2b, 0xff, 0xa4, 0x7a, 0xf2, 0x8c,
	0x34, 0xd6, 0x32, 0x52, 0xf6, 0x46, 0xae, 0x79, 0x73, 0x17, 0x40, 0x8c, 0x95, 0x22, 0x63, 0x0a,
	0x47, 0xa8, 0x2f, 0x15, 0xea, 0xb4, 0xd6, 0xa8, 0x83, 0x9e, 0xe4, 0x55, 0xc1, 0x59, 0x7a, 0x7b,
	0x50, 0xb9, 0x69, 0xb5, 0x30, 0x06, 0x05, 0xa7, 0x55, 0xa6, 0xbd, 0x5b, 0xd5, 0x5e, 0xa3, 0xb5,
	0x19, 0xc1, 0x76, 0xa9, 0x74, 0x98, 0xd7, 0x9b, 0x98, 0x58, 0x2b, 0x5f, 0x79, 0x43, 0xf9, 0x3e,
	0x80, 0x2d, 0xde, 0x11, 0x92, 0x4a, 0x9b, 0x10, 0xad, 0x23, 0xe1, 0x8d, 0xea, 0x57, 0x09, 0xe4,
	0xd3, 0x68, 0xba, 0x21, 0xb6, 0x6b, 0xbc, 0x6e, 0xd4, 0x78, 0x9d, 0x5d, 0x4c, 0xde, 0x54, 0x22,
	0xcd, 0x72, 0x42, 0x4a, 0xdc, 0xec, 0x54, 0xb9, 0xb9, 0x0f, 0x9d, 0x38, 0xe5, 0x69, 0xe2, 0xb3,
	0xbd, 0x1d, 0xa7, 0x2c, 0x45, 0x42, 0x50, 0x90, 0x8d, 0x0a, 0xc4, 0x7c, 0xac, 0xb0, 0x59, 0xa9,
	0xb3, 0xb9, 0x0f, 0xfa, 0xdc, 0x4b, 0x88, 0x9b, 0xa4, 0xb3, 0x19, 0x4e, 0x12, 0x66, 0x44, 0x65,
	0x46, 0x7a, 0x14, 0x77, 0x38, 0x7c, 0x11, 0xf8, 0xe6, 0x0f, 0x0d, 0x00, 0xbe, 0x7d, 0xdd, 0x40,
	0xb0, 0x7d, 0xe8, 0xd0, 0xfd, 0xa5, 0x08, 0x40, 0xfb, 0x3a, 0x9a, 0xde, 0xdc, 0x98, 0x1f, 0xe6,
	0xdc, 0x68, 0x8a, 0x46, 0xc8, 0xad, 0x57, 0x49, 0xf1, 0xb8, 0x20, 0x45, 0x8b, 0xa9, 0x6d, 0x09,
	0xb5, 0xf5, 0x26, 0xf7, 0x18, 0xb6, 0x13, 0xe2, 0xc5, 0xa4, 0x54, 0x52, 0xbc, 0x4f, 0xf7, 0x18,
	0x5c, 0x14, 0xd4, 0x5f, 0x98, 0xd2, 0x0f, 0xa0, 0x9d, 0x10, 0x8f, 0xa4, 0x7c, 0x2d, 0xa0, 0xd5,
	0x77, 0x7a, 0x76, 0xe4, 0x3a, 0x93, 0xe1, 0xe4, 0xc2, 0xb1, 0x85, 0xe8, 0xc6, 0xb4, 0x98, 0x2b,
	0xd0, 0xca, 0x57, 0xdd, 0x10, 0xb1, 0x47, 0x62, 0x97, 0x6e, 0x30, 0xeb, 0xa8, 0xe2, 0x59, 0x79,
	0x9b, 0xfe, 0xa4, 0xb6, 0x4d, 0xf7, 0x00, 0x8a, 0x7e, 0xaf, 0x4b, 0x68, 0x0b, 0x94, 0x62, 0x5f,
	0x6e, 0x98, 0xbf, 0x49, 0xa0, 0x4e, 0x70, 0xe8, 0x85, 0xc4, 0x21, 0x1e, 0x61, 0xf4, 0xce, 0x88,
	0xc0, 0xe9, 0x2d, 0xf1, 0x39, 0x2c, 0xc0, 0x7c, 0xce, 0x55, 0x37, 0xd9, 0xc6, 0x86, 0x4d, 0xb6,
	0xb2, 0xa4, 0xca, 0x6b, 0x4b, 0x2a, 0x9b, 0xb5, 0xcb, 0x28, 0xae, 0xae, 0xc2, 0x2a, 0xc7, 0x72,
	0x15, 0xce, 0xe7, 0xd2, 0xb4, 0x6d, 0xd9, 0x2a, 0xc7, 0xb8, 0xca, 0x3d, 0x50, 0x17, 0x91, 0x8f,
	0xe7, 0x95, 0x4d, 0x18, 0x18, 0x94, 0xdb, 0xa8, 0x6c, 0xba, 0x9d, 0xfa, 0xa6, 0xfb, 0x39, 0xc0,
	0x30, 0xa1, 0x97, 0xf6, 0x66, 0x6f, 0x31, 0xda, 0x83, 0xb6, 0x1f, 0x5c, 0xe1, 0x84, 0x88, 0xe0,
	0x8b, 0x37, 0x5a, 0x95, 0xd3, 0xc8, 0x5f, 0x09, 0xba, 0xb2, 0x67, 0xf3, 0x33, 0x50, 0x9f, 0x06,
	0x78, 0xee, 0x3b, 0xb3, 0xb7, 0x78, 0xe1, 0x6d, 0x5c, 0x94, 0x51, 0x29, 0x6d, 0x8a, 0x48, 0xd1,
	0x37, 0x12, 0xec, 0x64, 0x7b, 0xaf, 0x98, 0x56, 0x51, 0x9c, 0x6b, 0x4a, 0x85, 0x66, 0x6d, 0x21,
	0x69, 0xd4, 0x17, 0x92, 0x87, 0xd0, 0x4e, 0xd8, 0xe7, 0x0d, 0x59, 0x94, 0x46, 0xe9, 0x4a, 0xb6,
	0x90, 0xb1, 0xc6, 0xc6, 0x16, 0x4f, 0x37, 0x21, 0x71, 0x10, 0x5e, 0x89, 0x3e, 0xa2, 0x71, 0xd0,
	0x61, 0x98, 0xf9, 0x25, 0xf4, 0x86, 0xcb, 0x65, 0x79, 0x32, 0x7c, 0x0c, 0x3b, 0xe2, 0xdf, 0x27,
	0x88, 0x42, 0x21, 0x10, 0x17, 0xac, 0x0b, 0x3e, 0xfa, 0x56, 0x86, 0x96, 0xed, 0x85, 0x57, 0x18,
	0x01, 0xb4, 0xad, 0xf7, 0x38, 0x24, 0x89, 0x7e, 0x0b, 0x69, 0xd0, 0x3d, 0xb7, 0xcf, 0x4e, 0xad,
	0xe3, 0x89, 0xa3, 0x4b, 0xf4, 0xed, 0x64, 0x38, 0x19, 0x3a, 0xd6, 0xc4, 0xd1, 0x1b, 0x94, 0xa4,
	0xa7, 0x67, 0x47, 0x8e, 0xde, 0x44, 0x0a, 0xb4, 0x68, 0xad, 0x38, 0x7a, 0x0b, 0x6d, 0x83, 0x3a,
	0x74, 0x1c, 0x6b, 0xe2, 0x1e, 0x0f, 0x8f, 0x9f, 0x59, 0x7a, 0x1b, 0xa9, 0xd0, 0xb1, 0x5e, 0x9f,
	0x5b, 0xf6, 0xc4, 0xd1, 0x3b, 0xf4, 0xa7, 0x4f, 0x18, 0x70, 0x5f, 0x59, 0xb6, 0x33, 0x3a, 0x1b,
	0x3b, 0x7a, 0x17, 0xed, 0x01, 0xca, 0xd0, 0x67, 0xd6, 0xf0, 0xc4, 0x1d, 0x4d, 0xac, 0x97, 0x8e,
	0xae, 0xd0, 0xa3, 0xce, 0x1b, 0x87, 0xbd, 0x6c, 0xd1, 0xdf, 0x4a, 0xfe, 0x52, 0x9c, 0xec, 0xd1,
	0xdf, 0x4a, 0x01, 0xbe, 0x18, 0x8d, 0x9f, 0x3b, 0xfa, 0xff, 0xd0, 0x3e, 0xdc, 0xce, 0x6c, 0x1d,
	0x8d, 0xc6, 0x43, 0xfb, 0x8d, 0x3b, 0xb1, 0x2d, 0x4b, 0xdf, 0xa3, 0x7f, 0x96, 0xc2, 0x13, 0x77,
	0x3c, 0x7c, 0x69, 0x39, 0x3a, 0x50, 0x77, 0x68, 0x99, 0xdb, 0x17, 0x63, 0x47, 0xd7, 0xa8, 0x2d,
	0x26, 0x70, 0x9d, 0xe3, 0xb3, 0x73, 0xeb, 0x44, 0x57, 0x91, 0x01, 0xbb, 0xa3, 0x93, 0xd7, 0x6e,
	0x76, 0x2c, 0x77, 0x7d, 0x37, 0x93, 0xd4, 0x7c, 0xf9, 0x17, 0xfd, 0x3e, 0x95, 0xac, 0x5f, 0x75,
	0x3f, 0x13, 0x64, 0xc6, 0x32, 0xc7, 0x0c, 0x1a, 0x93, 0xb2, 0x80, 0x85, 0xf4, 0xdf, 0x48, 0x83,
	0x8e, 0x38, 0xac, 0xff, 0x2e, 0x4d, 0xdb, 0xec, 0x4f, 0xff, 0xd3, 0x3f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xb5, 0x8e, 0xda, 0x44, 0xff, 0x0f, 0x00, 0x00,
}
