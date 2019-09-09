// Code generated by protoc-gen-go.
// source: vo.proto
// DO NOT EDIT!

/*
Package vo is a generated protocol buffer package.

It is generated from these files:
	vo.proto

It has these top-level messages:
	ProjectMetadataDelta
	SystemMetadataDelta
	ExpertMetadataDelta
	Relation
	JobMetadataDelta
	SystemVersionInput
	SystemVersionOutput
	JobRunInput
	JobOutput
	DatasetVerInput
	DatasetVerOutput
	DatasetMetadataDelta
	DatasetItem
	DatasetSample
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

type ENTITY int32

const (
	ENTITY_NONE            ENTITY = 0
	ENTITY_DATASET         ENTITY = 1
	ENTITY_JOB             ENTITY = 2
	ENTITY_SYSTEM          ENTITY = 3
	ENTITY_MODEL           ENTITY = 4
	ENTITY_PROJECT         ENTITY = 5
	ENTITY_DATASET_VERSION ENTITY = 6
	ENTITY_JOB_RUN         ENTITY = 7
	ENTITY_SYSTEM_VERSION  ENTITY = 9
)

var ENTITY_name = map[int32]string{
	0: "NONE",
	1: "DATASET",
	2: "JOB",
	3: "SYSTEM",
	4: "MODEL",
	5: "PROJECT",
	6: "DATASET_VERSION",
	7: "JOB_RUN",
	9: "SYSTEM_VERSION",
}
var ENTITY_value = map[string]int32{
	"NONE":            0,
	"DATASET":         1,
	"JOB":             2,
	"SYSTEM":          3,
	"MODEL":           4,
	"PROJECT":         5,
	"DATASET_VERSION": 6,
	"JOB_RUN":         7,
	"SYSTEM_VERSION":  9,
}

func (x ENTITY) String() string {
	return proto.EnumName(ENTITY_name, int32(x))
}
func (ENTITY) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type JOB_STATUS int32

const (
	JOB_STATUS_NEVER   JOB_STATUS = 0
	JOB_STATUS_SUCCESS JOB_STATUS = 1
	JOB_STATUS_FAIL    JOB_STATUS = 2
	JOB_STATUS_RUNNING JOB_STATUS = 3
)

var JOB_STATUS_name = map[int32]string{
	0: "NEVER",
	1: "SUCCESS",
	2: "FAIL",
	3: "RUNNING",
}
var JOB_STATUS_value = map[string]int32{
	"NEVER":   0,
	"SUCCESS": 1,
	"FAIL":    2,
	"RUNNING": 3,
}

func (x JOB_STATUS) String() string {
	return proto.EnumName(JOB_STATUS_name, int32(x))
}
func (JOB_STATUS) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SystemVersionInput_Type int32

const (
	SystemVersionInput_None   SystemVersionInput_Type = 0
	SystemVersionInput_System SystemVersionInput_Type = 1
	SystemVersionInput_JobRun SystemVersionInput_Type = 2
)

var SystemVersionInput_Type_name = map[int32]string{
	0: "None",
	1: "System",
	2: "JobRun",
}
var SystemVersionInput_Type_value = map[string]int32{
	"None":   0,
	"System": 1,
	"JobRun": 2,
}

func (x SystemVersionInput_Type) String() string {
	return proto.EnumName(SystemVersionInput_Type_name, int32(x))
}
func (SystemVersionInput_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type SystemVersionOutput_Type int32

const (
	SystemVersionOutput_None   SystemVersionOutput_Type = 0
	SystemVersionOutput_System SystemVersionOutput_Type = 1
)

var SystemVersionOutput_Type_name = map[int32]string{
	0: "None",
	1: "System",
}
var SystemVersionOutput_Type_value = map[string]int32{
	"None":   0,
	"System": 1,
}

func (x SystemVersionOutput_Type) String() string {
	return proto.EnumName(SystemVersionOutput_Type_name, int32(x))
}
func (SystemVersionOutput_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

type JobRunInput_Type int32

const (
	JobRunInput_None       JobRunInput_Type = 0
	JobRunInput_DatasetVer JobRunInput_Type = 1
	JobRunInput_System     JobRunInput_Type = 2
)

var JobRunInput_Type_name = map[int32]string{
	0: "None",
	1: "DatasetVer",
	2: "System",
}
var JobRunInput_Type_value = map[string]int32{
	"None":       0,
	"DatasetVer": 1,
	"System":     2,
}

func (x JobRunInput_Type) String() string {
	return proto.EnumName(JobRunInput_Type_name, int32(x))
}
func (JobRunInput_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type JobOutput_Type int32

const (
	JobOutput_None    JobOutput_Type = 0
	JobOutput_Dataset JobOutput_Type = 1
)

var JobOutput_Type_name = map[int32]string{
	0: "None",
	1: "Dataset",
}
var JobOutput_Type_value = map[string]int32{
	"None":    0,
	"Dataset": 1,
}

func (x JobOutput_Type) String() string {
	return proto.EnumName(JobOutput_Type_name, int32(x))
}
func (JobOutput_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

type DatasetVerInput_TYPE int32

const (
	DatasetVerInput_NONE    DatasetVerInput_TYPE = 0
	DatasetVerInput_JOB_RUN DatasetVerInput_TYPE = 1
)

var DatasetVerInput_TYPE_name = map[int32]string{
	0: "NONE",
	1: "JOB_RUN",
}
var DatasetVerInput_TYPE_value = map[string]int32{
	"NONE":    0,
	"JOB_RUN": 1,
}

func (x DatasetVerInput_TYPE) String() string {
	return proto.EnumName(DatasetVerInput_TYPE_name, int32(x))
}
func (DatasetVerInput_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

type DatasetVerOutput_TYPE int32

const (
	DatasetVerOutput_NONE    DatasetVerOutput_TYPE = 0
	DatasetVerOutput_JOB_RUN DatasetVerOutput_TYPE = 1
)

var DatasetVerOutput_TYPE_name = map[int32]string{
	0: "NONE",
	1: "JOB_RUN",
}
var DatasetVerOutput_TYPE_value = map[string]int32{
	"NONE":    0,
	"JOB_RUN": 1,
}

func (x DatasetVerOutput_TYPE) String() string {
	return proto.EnumName(DatasetVerOutput_TYPE_name, int32(x))
}
func (DatasetVerOutput_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{10, 0} }

// Sample storage format. Doesn't necessarily match
// actual data format which could be very cryptic
type DatasetSample_FORMAT int32

const (
	DatasetSample_TEXT DatasetSample_FORMAT = 0
	DatasetSample_TSV  DatasetSample_FORMAT = 1
	DatasetSample_JSON DatasetSample_FORMAT = 2
)

var DatasetSample_FORMAT_name = map[int32]string{
	0: "TEXT",
	1: "TSV",
	2: "JSON",
}
var DatasetSample_FORMAT_value = map[string]int32{
	"TEXT": 0,
	"TSV":  1,
	"JSON": 2,
}

func (x DatasetSample_FORMAT) String() string {
	return proto.EnumName(DatasetSample_FORMAT_name, int32(x))
}
func (DatasetSample_FORMAT) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{13, 0} }

type ProjectMetadataDelta struct {
	Title          string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	TitleSet       bool   `protobuf:"varint,2,opt,name=title_set,json=titleSet" json:"title_set,omitempty"`
	Description    string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	DescriptionSet bool   `protobuf:"varint,4,opt,name=description_set,json=descriptionSet" json:"description_set,omitempty"`
}

func (m *ProjectMetadataDelta) Reset()                    { *m = ProjectMetadataDelta{} }
func (m *ProjectMetadataDelta) String() string            { return proto.CompactTextString(m) }
func (*ProjectMetadataDelta) ProtoMessage()               {}
func (*ProjectMetadataDelta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SystemMetadataDelta struct {
	Title          string   `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	TitleSet       bool     `protobuf:"varint,2,opt,name=title_set,json=titleSet" json:"title_set,omitempty"`
	LocationId     string   `protobuf:"bytes,3,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	LocationIdSet  bool     `protobuf:"varint,4,opt,name=location_id_set,json=locationIdSet" json:"location_id_set,omitempty"`
	LocationUri    string   `protobuf:"bytes,5,opt,name=location_uri,json=locationUri" json:"location_uri,omitempty"`
	LocationUriSet bool     `protobuf:"varint,6,opt,name=location_uri_set,json=locationUriSet" json:"location_uri_set,omitempty"`
	Description    string   `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	DescriptionSet bool     `protobuf:"varint,8,opt,name=description_set,json=descriptionSet" json:"description_set,omitempty"`
	Experts        [][]byte `protobuf:"bytes,9,rep,name=experts,proto3" json:"experts,omitempty"`
	ExpertsSet     bool     `protobuf:"varint,10,opt,name=experts_set,json=expertsSet" json:"experts_set,omitempty"`
}

func (m *SystemMetadataDelta) Reset()                    { *m = SystemMetadataDelta{} }
func (m *SystemMetadataDelta) String() string            { return proto.CompactTextString(m) }
func (*SystemMetadataDelta) ProtoMessage()               {}
func (*SystemMetadataDelta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ExpertMetadataDelta struct {
	FullName    string `protobuf:"bytes,1,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	FullNameSet bool   `protobuf:"varint,2,opt,name=full_name_set,json=fullNameSet" json:"full_name_set,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	EmailSet    bool   `protobuf:"varint,4,opt,name=email_set,json=emailSet" json:"email_set,omitempty"`
}

func (m *ExpertMetadataDelta) Reset()                    { *m = ExpertMetadataDelta{} }
func (m *ExpertMetadataDelta) String() string            { return proto.CompactTextString(m) }
func (*ExpertMetadataDelta) ProtoMessage()               {}
func (*ExpertMetadataDelta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Relation struct {
	TargetId   []byte `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	TargetType ENTITY `protobuf:"varint,2,opt,name=target_type,json=targetType,enum=ENTITY" json:"target_type,omitempty"`
}

func (m *Relation) Reset()                    { *m = Relation{} }
func (m *Relation) String() string            { return proto.CompactTextString(m) }
func (*Relation) ProtoMessage()               {}
func (*Relation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type JobMetadataDelta struct {
	Title    string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	TitleSet bool   `protobuf:"varint,2,opt,name=title_set,json=titleSet" json:"title_set,omitempty"`
}

func (m *JobMetadataDelta) Reset()                    { *m = JobMetadataDelta{} }
func (m *JobMetadataDelta) String() string            { return proto.CompactTextString(m) }
func (*JobMetadataDelta) ProtoMessage()               {}
func (*JobMetadataDelta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SystemVersionInput struct {
	Uid  []byte                  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type SystemVersionInput_Type `protobuf:"varint,2,opt,name=type,enum=SystemVersionInput_Type" json:"type,omitempty"`
}

func (m *SystemVersionInput) Reset()                    { *m = SystemVersionInput{} }
func (m *SystemVersionInput) String() string            { return proto.CompactTextString(m) }
func (*SystemVersionInput) ProtoMessage()               {}
func (*SystemVersionInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type SystemVersionOutput struct {
	Uid  []byte                   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type SystemVersionOutput_Type `protobuf:"varint,2,opt,name=type,enum=SystemVersionOutput_Type" json:"type,omitempty"`
}

func (m *SystemVersionOutput) Reset()                    { *m = SystemVersionOutput{} }
func (m *SystemVersionOutput) String() string            { return proto.CompactTextString(m) }
func (*SystemVersionOutput) ProtoMessage()               {}
func (*SystemVersionOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type JobRunInput struct {
	// dataset
	Uid  []byte           `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type JobRunInput_Type `protobuf:"varint,2,opt,name=type,enum=JobRunInput_Type" json:"type,omitempty"`
}

func (m *JobRunInput) Reset()                    { *m = JobRunInput{} }
func (m *JobRunInput) String() string            { return proto.CompactTextString(m) }
func (*JobRunInput) ProtoMessage()               {}
func (*JobRunInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type JobOutput struct {
	// dataset
	TargetId []byte         `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	Type     JobOutput_Type `protobuf:"varint,2,opt,name=type,enum=JobOutput_Type" json:"type,omitempty"`
}

func (m *JobOutput) Reset()                    { *m = JobOutput{} }
func (m *JobOutput) String() string            { return proto.CompactTextString(m) }
func (*JobOutput) ProtoMessage()               {}
func (*JobOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type DatasetVerInput struct {
	Uid  []byte               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type DatasetVerInput_TYPE `protobuf:"varint,2,opt,name=type,enum=DatasetVerInput_TYPE" json:"type,omitempty"`
}

func (m *DatasetVerInput) Reset()                    { *m = DatasetVerInput{} }
func (m *DatasetVerInput) String() string            { return proto.CompactTextString(m) }
func (*DatasetVerInput) ProtoMessage()               {}
func (*DatasetVerInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type DatasetVerOutput struct {
	Uid  []byte                `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type DatasetVerOutput_TYPE `protobuf:"varint,2,opt,name=type,enum=DatasetVerOutput_TYPE" json:"type,omitempty"`
}

func (m *DatasetVerOutput) Reset()                    { *m = DatasetVerOutput{} }
func (m *DatasetVerOutput) String() string            { return proto.CompactTextString(m) }
func (*DatasetVerOutput) ProtoMessage()               {}
func (*DatasetVerOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// if a field kind is set, then the corresponding property carries
// a new value (which could be null)
type DatasetMetadataDelta struct {
	Sample         *DatasetSample `protobuf:"bytes,7,opt,name=sample" json:"sample,omitempty"`
	SampleSet      bool           `protobuf:"varint,8,opt,name=sample_set,json=sampleSet" json:"sample_set,omitempty"`
	DataFormat     string         `protobuf:"bytes,11,opt,name=data_format,json=dataFormat" json:"data_format,omitempty"`
	DataFormatSet  bool           `protobuf:"varint,12,opt,name=data_format_set,json=dataFormatSet" json:"data_format_set,omitempty"`
	Description    string         `protobuf:"bytes,13,opt,name=description" json:"description,omitempty"`
	DescriptionSet bool           `protobuf:"varint,14,opt,name=description_set,json=descriptionSet" json:"description_set,omitempty"`
	LocationId     string         `protobuf:"bytes,15,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	LocationIdSet  bool           `protobuf:"varint,16,opt,name=location_id_set,json=locationIdSet" json:"location_id_set,omitempty"`
	LocationUri    string         `protobuf:"bytes,17,opt,name=location_uri,json=locationUri" json:"location_uri,omitempty"`
	LocationUriSet bool           `protobuf:"varint,18,opt,name=location_uri_set,json=locationUriSet" json:"location_uri_set,omitempty"`
	Experts        [][]byte       `protobuf:"bytes,19,rep,name=experts,proto3" json:"experts,omitempty"`
	ExpertsSet     bool           `protobuf:"varint,20,opt,name=experts_set,json=expertsSet" json:"experts_set,omitempty"`
	Title          string         `protobuf:"bytes,21,opt,name=title" json:"title,omitempty"`
	TitleSet       bool           `protobuf:"varint,22,opt,name=title_set,json=titleSet" json:"title_set,omitempty"`
}

func (m *DatasetMetadataDelta) Reset()                    { *m = DatasetMetadataDelta{} }
func (m *DatasetMetadataDelta) String() string            { return proto.CompactTextString(m) }
func (*DatasetMetadataDelta) ProtoMessage()               {}
func (*DatasetMetadataDelta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DatasetMetadataDelta) GetSample() *DatasetSample {
	if m != nil {
		return m.Sample
	}
	return nil
}

type DatasetItem struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	StorageBytes int64  `protobuf:"varint,2,opt,name=storage_bytes,json=storageBytes" json:"storage_bytes,omitempty"`
	Records      int64  `protobuf:"varint,3,opt,name=records" json:"records,omitempty"`
	Uid          []byte `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (m *DatasetItem) Reset()                    { *m = DatasetItem{} }
func (m *DatasetItem) String() string            { return proto.CompactTextString(m) }
func (*DatasetItem) ProtoMessage()               {}
func (*DatasetItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type DatasetSample struct {
	Format DatasetSample_FORMAT `protobuf:"varint,1,opt,name=format,enum=DatasetSample_FORMAT" json:"format,omitempty"`
	Body   string               `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *DatasetSample) Reset()                    { *m = DatasetSample{} }
func (m *DatasetSample) String() string            { return proto.CompactTextString(m) }
func (*DatasetSample) ProtoMessage()               {}
func (*DatasetSample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*ProjectMetadataDelta)(nil), "ProjectMetadataDelta")
	proto.RegisterType((*SystemMetadataDelta)(nil), "SystemMetadataDelta")
	proto.RegisterType((*ExpertMetadataDelta)(nil), "ExpertMetadataDelta")
	proto.RegisterType((*Relation)(nil), "Relation")
	proto.RegisterType((*JobMetadataDelta)(nil), "JobMetadataDelta")
	proto.RegisterType((*SystemVersionInput)(nil), "SystemVersionInput")
	proto.RegisterType((*SystemVersionOutput)(nil), "SystemVersionOutput")
	proto.RegisterType((*JobRunInput)(nil), "JobRunInput")
	proto.RegisterType((*JobOutput)(nil), "JobOutput")
	proto.RegisterType((*DatasetVerInput)(nil), "DatasetVerInput")
	proto.RegisterType((*DatasetVerOutput)(nil), "DatasetVerOutput")
	proto.RegisterType((*DatasetMetadataDelta)(nil), "DatasetMetadataDelta")
	proto.RegisterType((*DatasetItem)(nil), "DatasetItem")
	proto.RegisterType((*DatasetSample)(nil), "DatasetSample")
	proto.RegisterEnum("ENTITY", ENTITY_name, ENTITY_value)
	proto.RegisterEnum("JOB_STATUS", JOB_STATUS_name, JOB_STATUS_value)
	proto.RegisterEnum("SystemVersionInput_Type", SystemVersionInput_Type_name, SystemVersionInput_Type_value)
	proto.RegisterEnum("SystemVersionOutput_Type", SystemVersionOutput_Type_name, SystemVersionOutput_Type_value)
	proto.RegisterEnum("JobRunInput_Type", JobRunInput_Type_name, JobRunInput_Type_value)
	proto.RegisterEnum("JobOutput_Type", JobOutput_Type_name, JobOutput_Type_value)
	proto.RegisterEnum("DatasetVerInput_TYPE", DatasetVerInput_TYPE_name, DatasetVerInput_TYPE_value)
	proto.RegisterEnum("DatasetVerOutput_TYPE", DatasetVerOutput_TYPE_name, DatasetVerOutput_TYPE_value)
	proto.RegisterEnum("DatasetSample_FORMAT", DatasetSample_FORMAT_name, DatasetSample_FORMAT_value)
}

func init() { proto.RegisterFile("vo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 925 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0xcf, 0x6f, 0xdb, 0x36,
	0x14, 0xae, 0xfc, 0x53, 0x7a, 0xf2, 0x0f, 0x96, 0x71, 0x0b, 0x0f, 0x5d, 0xb1, 0x4c, 0x45, 0x3b,
	0x2f, 0x68, 0x75, 0xe8, 0x8e, 0x3b, 0x39, 0x89, 0x32, 0xd8, 0x68, 0xa4, 0x8c, 0x92, 0x8d, 0xe5,
	0x64, 0xc8, 0x31, 0x5b, 0x78, 0x90, 0x2d, 0x43, 0xa2, 0x87, 0xf9, 0xb2, 0xc3, 0x6e, 0xbb, 0xef,
	0x4f, 0xdc, 0xbf, 0xb0, 0xfb, 0x40, 0x52, 0xb2, 0x29, 0xc7, 0x09, 0x3c, 0xe4, 0x46, 0x7e, 0x7c,
	0xef, 0x7d, 0x1f, 0x9f, 0x1e, 0x3f, 0x08, 0xf4, 0xdf, 0x62, 0x7b, 0x95, 0xc4, 0x2c, 0xb6, 0xfe,
	0xd6, 0xa0, 0x73, 0x93, 0xc4, 0xbf, 0xd2, 0x3b, 0x76, 0x4d, 0x59, 0x38, 0x0b, 0x59, 0x78, 0x49,
	0x23, 0x16, 0xe2, 0x0e, 0x54, 0xd9, 0x9c, 0x45, 0xb4, 0xab, 0x9d, 0x6a, 0x3d, 0x83, 0xc8, 0x0d,
	0x7e, 0x05, 0x86, 0x58, 0x4c, 0x52, 0xca, 0xba, 0xa5, 0x53, 0xad, 0xa7, 0x13, 0x5d, 0x00, 0x3e,
	0x65, 0xf8, 0x14, 0xcc, 0x19, 0x4d, 0xef, 0x92, 0xf9, 0x8a, 0xcd, 0xe3, 0x65, 0xb7, 0x2c, 0x12,
	0x55, 0x08, 0x7f, 0x07, 0x6d, 0x65, 0x2b, 0x8a, 0x54, 0x44, 0x91, 0x96, 0x02, 0xfb, 0x94, 0x59,
	0xff, 0x94, 0xe0, 0xc4, 0xdf, 0xa4, 0x8c, 0x2e, 0x9e, 0xac, 0xea, 0x1b, 0x30, 0xa3, 0xf8, 0x2e,
	0x14, 0x84, 0xf3, 0x59, 0xa6, 0x0a, 0x72, 0x68, 0x30, 0xc3, 0xef, 0xa0, 0xad, 0x04, 0x28, 0xa2,
	0x9a, 0xbb, 0x20, 0x5e, 0xe8, 0x5b, 0x68, 0x6c, 0xe3, 0xd6, 0xc9, 0xbc, 0x5b, 0x95, 0xf7, 0xcb,
	0xb1, 0x51, 0x32, 0xc7, 0x3d, 0x40, 0x6a, 0x88, 0xa8, 0x55, 0x93, 0x17, 0x54, 0xc2, 0x0e, 0xf4,
	0xaa, 0x7e, 0x54, 0xaf, 0xf4, 0x43, 0xbd, 0xc2, 0x5d, 0xa8, 0xd3, 0xdf, 0x57, 0x34, 0x61, 0x69,
	0xd7, 0x38, 0x2d, 0xf7, 0x1a, 0x24, 0xdf, 0xf2, 0xab, 0x67, 0x4b, 0x91, 0x0e, 0x22, 0x1d, 0x32,
	0x88, 0xb7, 0xf9, 0x2f, 0x0d, 0x4e, 0x1c, 0xb1, 0x2d, 0xb6, 0xf9, 0x15, 0x18, 0x9f, 0xd7, 0x51,
	0x34, 0x59, 0x86, 0x8b, 0xbc, 0xd5, 0x3a, 0x07, 0xdc, 0x70, 0x41, 0xb1, 0x05, 0xcd, 0xed, 0xa1,
	0xd2, 0x71, 0x33, 0x0f, 0xe0, 0x9a, 0x3a, 0x50, 0xa5, 0x8b, 0x70, 0x1e, 0x65, 0xed, 0x96, 0x1b,
	0x5e, 0x56, 0x2c, 0x94, 0x1e, 0xeb, 0x02, 0xe0, 0x5a, 0x7e, 0x06, 0x9d, 0xd0, 0x48, 0xf4, 0x48,
	0x7c, 0xd0, 0x30, 0xf9, 0x42, 0x19, 0xff, 0x62, 0x9c, 0xbf, 0x41, 0x74, 0x09, 0x0c, 0x66, 0xb8,
	0x07, 0x66, 0x76, 0xc8, 0x36, 0x2b, 0x2a, 0xd8, 0x5b, 0x1f, 0xeb, 0xb6, 0xe3, 0x06, 0x83, 0xe0,
	0x96, 0x80, 0x3c, 0x0b, 0x36, 0x2b, 0x6a, 0x39, 0x80, 0x86, 0xf1, 0xf4, 0xa9, 0x13, 0x64, 0xfd,
	0x01, 0x58, 0xce, 0xe2, 0x98, 0x26, 0x29, 0x1f, 0x87, 0xe5, 0x6a, 0xcd, 0x30, 0x82, 0xf2, 0x7a,
	0xab, 0x8e, 0x2f, 0xf1, 0x7b, 0xa8, 0x28, 0x8a, 0xba, 0xf6, 0xfd, 0x24, 0x9b, 0xcb, 0x22, 0x22,
	0xca, 0xea, 0x41, 0x85, 0xef, 0xb0, 0x0e, 0x15, 0x37, 0x5e, 0x52, 0xf4, 0x0c, 0x03, 0xd4, 0x64,
	0x0a, 0xd2, 0xf8, 0x7a, 0x18, 0x4f, 0xc9, 0x7a, 0x89, 0x4a, 0x16, 0xcb, 0xdf, 0x42, 0x56, 0xca,
	0x5b, 0xb3, 0xc3, 0x02, 0x3e, 0x14, 0x04, 0x7c, 0x65, 0x1f, 0xc8, 0x52, 0x15, 0x7c, 0xfd, 0x98,
	0x02, 0x8b, 0x81, 0x29, 0x15, 0x3c, 0x74, 0xdd, 0xb7, 0x05, 0xb6, 0xe7, 0xb6, 0x12, 0xad, 0xb2,
	0xbc, 0xbf, 0xc7, 0xd2, 0x02, 0xb8, 0x0c, 0x59, 0x98, 0x52, 0x36, 0xa6, 0x89, 0xbc, 0x6b, 0xc6,
	0x5a, 0xb2, 0x22, 0x30, 0x86, 0xf1, 0x34, 0xbb, 0xe1, 0xa3, 0x63, 0xf0, 0xa6, 0x40, 0xdf, 0xb6,
	0xb7, 0x69, 0x2a, 0xf9, 0xeb, 0x7b, 0xe4, 0x26, 0xd4, 0x33, 0x72, 0xa4, 0x59, 0x0b, 0x68, 0xef,
	0x94, 0x3c, 0x74, 0xcf, 0xef, 0x0b, 0x44, 0x2f, 0xec, 0xbd, 0x0c, 0x3b, 0xb8, 0xbd, 0x71, 0x14,
	0xba, 0xdb, 0x1b, 0x47, 0xd0, 0x79, 0xae, 0x23, 0xe9, 0x86, 0xde, 0xf9, 0x84, 0x8c, 0x5c, 0xa4,
	0x59, 0x31, 0xa0, 0x5d, 0xf2, 0x83, 0x5f, 0xf1, 0xac, 0xc0, 0xf7, 0xd2, 0xde, 0x4f, 0xf9, 0x1f,
	0x84, 0xff, 0x96, 0xa1, 0x93, 0xa5, 0x17, 0x5f, 0xc1, 0x3b, 0xa8, 0xa5, 0xe1, 0x62, 0x15, 0x51,
	0xe1, 0x3c, 0xe6, 0xc7, 0x56, 0xce, 0xe2, 0x0b, 0x94, 0x64, 0xa7, 0xf8, 0x35, 0x80, 0x5c, 0x29,
	0xfe, 0x63, 0x48, 0x24, 0xf3, 0x56, 0x5e, 0x73, 0xf2, 0x39, 0x4e, 0x16, 0x21, 0xeb, 0x9a, 0xd2,
	0x5b, 0x39, 0x74, 0x25, 0x10, 0xee, 0xad, 0x4a, 0x80, 0x28, 0xd2, 0x90, 0xde, 0xba, 0x0b, 0x3a,
	0x60, 0x87, 0xcd, 0xa3, 0xec, 0xb0, 0x75, 0xd0, 0x0e, 0xf7, 0xfc, 0xbe, 0x7d, 0x8c, 0xdf, 0xa3,
	0x63, 0xfc, 0xfe, 0xf9, 0x71, 0x7e, 0x8f, 0x0f, 0xfa, 0xbd, 0x62, 0xd2, 0x27, 0x8f, 0x9a, 0x74,
	0x67, 0xdf, 0xa4, 0x77, 0x8e, 0xf5, 0xe2, 0x41, 0xc7, 0x7a, 0xb9, 0xe7, 0x58, 0x09, 0x98, 0xd9,
	0xf7, 0x1c, 0x30, 0xba, 0xc0, 0x18, 0x2a, 0x8a, 0x93, 0x8b, 0x35, 0x7e, 0x03, 0xcd, 0x94, 0xc5,
	0x49, 0xf8, 0x85, 0x4e, 0xa6, 0x1b, 0x46, 0x53, 0x31, 0x6e, 0x65, 0xd2, 0xc8, 0xc0, 0x73, 0x8e,
	0x71, 0xd5, 0x09, 0xbd, 0x8b, 0x93, 0x59, 0x2a, 0x8c, 0xbc, 0x4c, 0xf2, 0x6d, 0x3e, 0xb6, 0x95,
	0xed, 0xd8, 0x5a, 0x1b, 0x68, 0x16, 0x66, 0x08, 0x7f, 0x80, 0x5a, 0x36, 0x17, 0x5a, 0xf1, 0xe5,
	0xc8, 0x73, 0xfb, 0xca, 0x23, 0xd7, 0xfd, 0x80, 0x64, 0x41, 0x5c, 0xe4, 0x34, 0x9e, 0x6d, 0x84,
	0x0e, 0x83, 0x88, 0xb5, 0xf5, 0x16, 0x6a, 0x32, 0x8a, 0x0f, 0x78, 0xe0, 0xfc, 0x12, 0xa0, 0x67,
	0xb8, 0x0e, 0xe5, 0xc0, 0x1f, 0x23, 0x8d, 0x43, 0x43, 0xdf, 0x73, 0x51, 0xe9, 0xec, 0x4f, 0x0d,
	0x6a, 0xd2, 0xfe, 0x8b, 0x0f, 0xe1, 0xb2, 0x1f, 0xf4, 0x7d, 0x27, 0x40, 0x1a, 0x4f, 0x1a, 0x7a,
	0xe7, 0xa8, 0x24, 0xbc, 0xe6, 0xd6, 0x0f, 0x9c, 0x6b, 0x54, 0xc6, 0x06, 0x54, 0xaf, 0xbd, 0x4b,
	0xe7, 0x13, 0xaa, 0xf0, 0xe0, 0x1b, 0xe2, 0x0d, 0x9d, 0x8b, 0x00, 0x55, 0xf1, 0x09, 0xb4, 0xb3,
	0xcc, 0xc9, 0xd8, 0x21, 0xfe, 0xc0, 0x73, 0x51, 0x4d, 0x7d, 0x57, 0x75, 0x8c, 0xa1, 0x25, 0xab,
	0x6c, 0x03, 0x8c, 0xb3, 0x1f, 0x01, 0x78, 0x80, 0x1f, 0xf4, 0x83, 0x91, 0xcf, 0x6b, 0xbb, 0xce,
	0xd8, 0x21, 0x52, 0x88, 0x3f, 0xba, 0xb8, 0x70, 0x7c, 0x5f, 0x8a, 0xbe, 0xea, 0x0f, 0x3e, 0xa1,
	0x12, 0x87, 0xc9, 0xc8, 0x75, 0x07, 0xee, 0x4f, 0xa8, 0x3c, 0xad, 0x89, 0xbf, 0xb1, 0x1f, 0xfe,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x41, 0x48, 0x62, 0x99, 0x09, 0x00, 0x00,
}
