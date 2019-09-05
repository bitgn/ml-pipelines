// Code generated by protoc-gen-go.
// source: events.proto
// DO NOT EDIT!

/*
Package events is a generated protocol buffer package.

It is generated from these files:
	events.proto

It has these top-level messages:
	ProjectCreated
	JobAdded
	JobRunStarted
	ServiceCreated
	DatasetCreated
	DatasetUpdated
	DatasetVersionAdded
	ExpertAdded
	Event
*/
package events

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

type Type int32

const (
	Type_None                      Type = 0
	Type_Event_ProjectCreated      Type = 1
	Type_Event_DatasetCreated      Type = 2
	Type_Event_DatasetUpdated      Type = 3
	Type_Event_ExpertAdded         Type = 4
	Type_Event_JobAdded            Type = 5
	Type_Event_ServiceCreated      Type = 6
	Type_Event_DatasetVersionAdded Type = 7
)

var Type_name = map[int32]string{
	0: "None",
	1: "Event_ProjectCreated",
	2: "Event_DatasetCreated",
	3: "Event_DatasetUpdated",
	4: "Event_ExpertAdded",
	5: "Event_JobAdded",
	6: "Event_ServiceCreated",
	7: "Event_DatasetVersionAdded",
}
var Type_value = map[string]int32{
	"None":                      0,
	"Event_ProjectCreated":      1,
	"Event_DatasetCreated":      2,
	"Event_DatasetUpdated":      3,
	"Event_ExpertAdded":         4,
	"Event_JobAdded":            5,
	"Event_ServiceCreated":      6,
	"Event_DatasetVersionAdded": 7,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ProjectCreated struct {
	Uid  []byte                   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name string                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Meta *vo.ProjectMetadataDelta `protobuf:"bytes,3,opt,name=meta" json:"meta,omitempty"`
}

func (m *ProjectCreated) Reset()                    { *m = ProjectCreated{} }
func (m *ProjectCreated) String() string            { return proto.CompactTextString(m) }
func (*ProjectCreated) ProtoMessage()               {}
func (*ProjectCreated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProjectCreated) GetMeta() *vo.ProjectMetadataDelta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type JobAdded struct {
	Uid         []byte               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ProjectUid  []byte               `protobuf:"bytes,2,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	Name        string               `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Meta        *vo.JobMetadataDelta `protobuf:"bytes,4,opt,name=meta" json:"meta,omitempty"`
	ProjectName string               `protobuf:"bytes,5,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
}

func (m *JobAdded) Reset()                    { *m = JobAdded{} }
func (m *JobAdded) String() string            { return proto.CompactTextString(m) }
func (*JobAdded) ProtoMessage()               {}
func (*JobAdded) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JobAdded) GetMeta() *vo.JobMetadataDelta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type JobRunStarted struct {
	JobUid    []byte            `protobuf:"bytes,1,opt,name=job_uid,json=jobUid,proto3" json:"job_uid,omitempty"`
	Uid       []byte            `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Title     string            `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Inputs    []*vo.JobRunInput `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	Timestamp int64             `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *JobRunStarted) Reset()                    { *m = JobRunStarted{} }
func (m *JobRunStarted) String() string            { return proto.CompactTextString(m) }
func (*JobRunStarted) ProtoMessage()               {}
func (*JobRunStarted) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JobRunStarted) GetInputs() []*vo.JobRunInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type ServiceCreated struct {
	Uid         []byte                   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ProjectUid  []byte                   `protobuf:"bytes,2,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	Name        string                   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Meta        *vo.ServiceMetadataDelta `protobuf:"bytes,4,opt,name=meta" json:"meta,omitempty"`
	ProjectName string                   `protobuf:"bytes,5,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
}

func (m *ServiceCreated) Reset()                    { *m = ServiceCreated{} }
func (m *ServiceCreated) String() string            { return proto.CompactTextString(m) }
func (*ServiceCreated) ProtoMessage()               {}
func (*ServiceCreated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ServiceCreated) GetMeta() *vo.ServiceMetadataDelta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type DatasetCreated struct {
	ProjectUid  []byte                   `protobuf:"bytes,1,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	Uid         []byte                   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name        string                   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Meta        *vo.DatasetMetadataDelta `protobuf:"bytes,4,opt,name=meta" json:"meta,omitempty"`
	ProjectName string                   `protobuf:"bytes,5,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
}

func (m *DatasetCreated) Reset()                    { *m = DatasetCreated{} }
func (m *DatasetCreated) String() string            { return proto.CompactTextString(m) }
func (*DatasetCreated) ProtoMessage()               {}
func (*DatasetCreated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DatasetCreated) GetMeta() *vo.DatasetMetadataDelta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type DatasetUpdated struct {
	Uid        []byte                   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ProjectUid []byte                   `protobuf:"bytes,2,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	Meta       *vo.DatasetMetadataDelta `protobuf:"bytes,3,opt,name=meta" json:"meta,omitempty"`
}

func (m *DatasetUpdated) Reset()                    { *m = DatasetUpdated{} }
func (m *DatasetUpdated) String() string            { return proto.CompactTextString(m) }
func (*DatasetUpdated) ProtoMessage()               {}
func (*DatasetUpdated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DatasetUpdated) GetMeta() *vo.DatasetMetadataDelta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type DatasetVersionAdded struct {
	Uid         []byte                `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ProjectUid  []byte                `protobuf:"bytes,2,opt,name=project_uid,json=projectUid,proto3" json:"project_uid,omitempty"`
	ParentUid   []byte                `protobuf:"bytes,3,opt,name=parent_uid,json=parentUid,proto3" json:"parent_uid,omitempty"`
	ProjectName string                `protobuf:"bytes,4,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
	Title       string                `protobuf:"bytes,5,opt,name=title" json:"title,omitempty"`
	Timestamp   int64                 `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Items       []*vo.DatasetItem     `protobuf:"bytes,7,rep,name=items" json:"items,omitempty"`
	Remove      []*vo.DatasetItem     `protobuf:"bytes,10,rep,name=remove" json:"remove,omitempty"`
	Inputs      []*vo.DatasetVerInput `protobuf:"bytes,8,rep,name=inputs" json:"inputs,omitempty"`
	DatasetUid  []byte                `protobuf:"bytes,9,opt,name=dataset_uid,json=datasetUid,proto3" json:"dataset_uid,omitempty"`
}

func (m *DatasetVersionAdded) Reset()                    { *m = DatasetVersionAdded{} }
func (m *DatasetVersionAdded) String() string            { return proto.CompactTextString(m) }
func (*DatasetVersionAdded) ProtoMessage()               {}
func (*DatasetVersionAdded) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DatasetVersionAdded) GetItems() []*vo.DatasetItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *DatasetVersionAdded) GetRemove() []*vo.DatasetItem {
	if m != nil {
		return m.Remove
	}
	return nil
}

func (m *DatasetVersionAdded) GetInputs() []*vo.DatasetVerInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type ExpertAdded struct {
	Uid  []byte                  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name string                  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Meta *vo.ExpertMetadataDelta `protobuf:"bytes,3,opt,name=meta" json:"meta,omitempty"`
}

func (m *ExpertAdded) Reset()                    { *m = ExpertAdded{} }
func (m *ExpertAdded) String() string            { return proto.CompactTextString(m) }
func (*ExpertAdded) ProtoMessage()               {}
func (*ExpertAdded) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ExpertAdded) GetMeta() *vo.ExpertMetadataDelta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type Event struct {
	Body []byte `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
	Type Type   `protobuf:"varint,2,opt,name=Type,enum=Type" json:"Type,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*ProjectCreated)(nil), "ProjectCreated")
	proto.RegisterType((*JobAdded)(nil), "JobAdded")
	proto.RegisterType((*JobRunStarted)(nil), "JobRunStarted")
	proto.RegisterType((*ServiceCreated)(nil), "ServiceCreated")
	proto.RegisterType((*DatasetCreated)(nil), "DatasetCreated")
	proto.RegisterType((*DatasetUpdated)(nil), "DatasetUpdated")
	proto.RegisterType((*DatasetVersionAdded)(nil), "DatasetVersionAdded")
	proto.RegisterType((*ExpertAdded)(nil), "ExpertAdded")
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterEnum("Type", Type_name, Type_value)
}

func init() { proto.RegisterFile("events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xc5, 0xb1, 0x9d, 0xc7, 0x4d, 0x88, 0xdc, 0x21, 0x11, 0x2e, 0xa2, 0x6a, 0xb0, 0x40, 0x72,
	0x59, 0x78, 0x11, 0x24, 0xf6, 0x40, 0xbb, 0x68, 0x24, 0x2a, 0xe4, 0x12, 0x76, 0x28, 0x1a, 0xd7,
	0x77, 0xe1, 0xa8, 0xf6, 0x58, 0xf6, 0x24, 0xa2, 0xff, 0xc1, 0x8e, 0x35, 0x3f, 0xc2, 0x2f, 0xf0,
	0x43, 0x68, 0x1e, 0x89, 0xe3, 0x3c, 0x50, 0x95, 0xdd, 0xf8, 0x9e, 0xc9, 0xb9, 0x67, 0xce, 0x9c,
	0xb9, 0x81, 0x1e, 0x2e, 0x31, 0xe3, 0x65, 0x90, 0x17, 0x8c, 0xb3, 0x17, 0xed, 0x25, 0x53, 0x2b,
	0x8f, 0x42, 0xff, 0x4b, 0xc1, 0xe6, 0x78, 0xc7, 0x3f, 0x15, 0x48, 0x39, 0xc6, 0xc4, 0x01, 0x73,
	0x91, 0xc4, 0xae, 0x31, 0x32, 0xfc, 0x5e, 0x28, 0x96, 0x84, 0x80, 0x95, 0xd1, 0x14, 0xdd, 0xc6,
	0xc8, 0xf0, 0x3b, 0xa1, 0x5c, 0x93, 0x0b, 0xb0, 0x52, 0xe4, 0xd4, 0x35, 0x47, 0x86, 0xdf, 0x1d,
	0x0f, 0x03, 0x4d, 0xf2, 0x19, 0x39, 0x8d, 0x29, 0xa7, 0x97, 0x78, 0xcf, 0x69, 0x28, 0xb7, 0x78,
	0xbf, 0x0c, 0x68, 0x4f, 0x58, 0xf4, 0x21, 0x8e, 0xf7, 0xb2, 0x9f, 0x43, 0x37, 0x57, 0x3f, 0x9e,
	0x09, 0xa4, 0x21, 0x11, 0xd0, 0xa5, 0xe9, 0x46, 0x7b, 0x73, 0xa3, 0xfd, 0x1b, 0xdd, 0xde, 0x92,
	0xed, 0x4f, 0x82, 0x09, 0x8b, 0xf6, 0xb4, 0x26, 0xaf, 0xa0, 0xb7, 0xe2, 0x96, 0x14, 0xb6, 0xa4,
	0x58, 0xf5, 0xbb, 0xa1, 0x29, 0x7a, 0x3f, 0x0d, 0x78, 0x3a, 0x61, 0x51, 0xb8, 0xc8, 0x6e, 0x39,
	0x2d, 0x84, 0x01, 0xcf, 0xa1, 0x35, 0x67, 0xd1, 0xac, 0x92, 0xd9, 0x9c, 0xb3, 0x48, 0x08, 0xd1,
	0xda, 0x1b, 0x95, 0xf6, 0x01, 0xd8, 0x3c, 0xe1, 0xf7, 0x2b, 0x6d, 0xea, 0x83, 0xbc, 0x86, 0x66,
	0x92, 0xe5, 0x0b, 0x5e, 0xba, 0xd6, 0xc8, 0xf4, 0xbb, 0xe3, 0x5e, 0xa0, 0x1a, 0x5c, 0x8b, 0x62,
	0xa8, 0x31, 0xf2, 0x12, 0x3a, 0x3c, 0x49, 0xb1, 0xe4, 0x34, 0xcd, 0xdd, 0xe6, 0xc8, 0xf0, 0xcd,
	0xb0, 0x2a, 0x78, 0xbf, 0x0d, 0xe8, 0xdf, 0x62, 0xb1, 0x4c, 0xee, 0xf0, 0xf0, 0xc5, 0x1c, 0x65,
	0xdd, 0x45, 0xcd, 0xba, 0x61, 0xa0, 0xbb, 0x1c, 0x69, 0x9f, 0xd0, 0x79, 0x49, 0x39, 0x2d, 0x71,
	0x1d, 0xa0, 0x2d, 0x55, 0xc6, 0x8e, 0xaa, 0x5d, 0x1f, 0x1f, 0xa3, 0x53, 0x77, 0x39, 0x52, 0x67,
	0xb6, 0x96, 0x39, 0xcd, 0xe3, 0x63, 0xed, 0xdc, 0x0e, 0xfd, 0x61, 0x49, 0xde, 0xdf, 0x06, 0x3c,
	0xd3, 0xf0, 0x37, 0x2c, 0xca, 0x84, 0x65, 0x47, 0xe7, 0xff, 0x0c, 0x20, 0xa7, 0x05, 0x66, 0x0a,
	0x37, 0x25, 0xde, 0x51, 0x15, 0x01, 0x6f, 0x1f, 0xde, 0xda, 0x39, 0x7c, 0x15, 0x53, 0x7b, 0x33,
	0xa6, 0xff, 0x0d, 0x20, 0xf1, 0xc0, 0x4e, 0x38, 0xa6, 0xa5, 0xdb, 0xd2, 0x19, 0xd6, 0xa7, 0xb9,
	0xe6, 0x98, 0x86, 0x0a, 0x12, 0x41, 0x2f, 0x30, 0x65, 0x4b, 0x74, 0x61, 0xcf, 0x26, 0x8d, 0x11,
	0x7f, 0xfd, 0x1c, 0xda, 0x72, 0x97, 0x13, 0x54, 0xc6, 0xd4, 0x9f, 0xc4, 0x39, 0x74, 0x63, 0x05,
	0xc9, 0xa3, 0x76, 0x94, 0x15, 0xba, 0x34, 0x4d, 0x62, 0xef, 0x3b, 0x74, 0xaf, 0x7e, 0xe4, 0x58,
	0xf0, 0x43, 0x66, 0xee, 0x1b, 0x55, 0x7e, 0xed, 0xd6, 0x06, 0x81, 0x62, 0xd8, 0x77, 0x69, 0xef,
	0xc1, 0xbe, 0x12, 0x63, 0x52, 0xd0, 0x7c, 0x64, 0xf1, 0x83, 0x66, 0x96, 0x6b, 0x72, 0x0a, 0xd6,
	0xd7, 0x87, 0x5c, 0x51, 0xf7, 0xc7, 0x76, 0x20, 0x3e, 0x42, 0x59, 0x7a, 0xfb, 0xc7, 0x50, 0x18,
	0x69, 0x83, 0x75, 0xc3, 0x32, 0x74, 0x9e, 0x10, 0x17, 0x06, 0x92, 0x6a, 0x56, 0x9f, 0xae, 0x8e,
	0x51, 0x21, 0xf5, 0x67, 0xe3, 0x34, 0x76, 0x10, 0x9d, 0x54, 0xc7, 0x24, 0x43, 0x38, 0x51, 0xc8,
	0xc6, 0xe9, 0x1d, 0x8b, 0x10, 0xe8, 0xab, 0xf2, 0x6a, 0xbc, 0x3a, 0x76, 0x45, 0x52, 0x9f, 0x1e,
	0x4e, 0x93, 0x9c, 0xc1, 0x69, 0x8d, 0x7e, 0x33, 0x97, 0x4e, 0x2b, 0x6a, 0xca, 0x3f, 0x84, 0x77,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x96, 0xac, 0xc0, 0x2a, 0x06, 0x00, 0x00,
}
