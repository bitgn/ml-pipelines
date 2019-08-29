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
	ServiceCreated
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
func (*DatasetCreated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
func (*DatasetUpdated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DatasetUpdated) GetMeta() *vo.DatasetMetadataDelta {
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
func (*JobAdded) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *JobAdded) GetMeta() *vo.JobMetadataDelta {
	if m != nil {
		return m.Meta
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
func (*ServiceCreated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ServiceCreated) GetMeta() *vo.ServiceMetadataDelta {
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
	Inputs      []*vo.DatasetVerInput `protobuf:"bytes,8,rep,name=inputs" json:"inputs,omitempty"`
	DatasetUid  []byte                `protobuf:"bytes,9,opt,name=dataset_uid,json=datasetUid,proto3" json:"dataset_uid,omitempty"`
}

func (m *DatasetVersionAdded) Reset()                    { *m = DatasetVersionAdded{} }
func (m *DatasetVersionAdded) String() string            { return proto.CompactTextString(m) }
func (*DatasetVersionAdded) ProtoMessage()               {}
func (*DatasetVersionAdded) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DatasetVersionAdded) GetItems() []*vo.DatasetItem {
	if m != nil {
		return m.Items
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
func (*ExpertAdded) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*ProjectCreated)(nil), "ProjectCreated")
	proto.RegisterType((*DatasetCreated)(nil), "DatasetCreated")
	proto.RegisterType((*DatasetUpdated)(nil), "DatasetUpdated")
	proto.RegisterType((*JobAdded)(nil), "JobAdded")
	proto.RegisterType((*ServiceCreated)(nil), "ServiceCreated")
	proto.RegisterType((*DatasetVersionAdded)(nil), "DatasetVersionAdded")
	proto.RegisterType((*ExpertAdded)(nil), "ExpertAdded")
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterEnum("Type", Type_name, Type_value)
}

func init() { proto.RegisterFile("events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x5f, 0x69, 0x32, 0x89, 0xa2, 0xed, 0x92, 0x4a, 0x2e, 0xa2, 0xc2, 0x58, 0x42, 0x32,
	0x1c, 0x7c, 0x08, 0x12, 0x77, 0xa0, 0x3d, 0xb4, 0x12, 0x15, 0x32, 0x84, 0x1b, 0xaa, 0x36, 0xf5,
	0x1c, 0x16, 0xd5, 0x5e, 0xcb, 0xde, 0x46, 0xf4, 0xb7, 0x70, 0xe6, 0xc8, 0x9f, 0xe0, 0x97, 0x55,
	0xde, 0x5d, 0xc7, 0x75, 0xe2, 0x5c, 0x7c, 0x1b, 0xcf, 0x5b, 0xbf, 0x79, 0xf3, 0x76, 0x76, 0x60,
	0x86, 0x1b, 0xcc, 0x65, 0x15, 0x17, 0xa5, 0x90, 0xe2, 0xc5, 0x78, 0x23, 0x74, 0x14, 0x32, 0x98,
	0x7f, 0x2d, 0xc5, 0x2f, 0xbc, 0x95, 0x9f, 0x4b, 0x64, 0x12, 0x53, 0x4a, 0xc0, 0xb9, 0xe7, 0xa9,
	0x6f, 0x05, 0x56, 0x34, 0x4b, 0xea, 0x90, 0x52, 0x70, 0x73, 0x96, 0xa1, 0x6f, 0x07, 0x56, 0x34,
	0x49, 0x54, 0x4c, 0xdf, 0x82, 0x9b, 0xa1, 0x64, 0xbe, 0x13, 0x58, 0xd1, 0x74, 0x79, 0x12, 0x1b,
	0x92, 0x2f, 0x28, 0x59, 0xca, 0x24, 0x3b, 0xc7, 0x3b, 0xc9, 0x12, 0x75, 0x24, 0xfc, 0x6b, 0xc1,
	0xfc, 0x9c, 0x49, 0x56, 0xe1, 0xb6, 0xc6, 0x2b, 0x98, 0x16, 0xfa, 0x87, 0x9b, 0xb6, 0x16, 0x98,
	0xd4, 0x8a, 0x6f, 0x45, 0xd8, 0xfb, 0x22, 0x9c, 0x1e, 0x11, 0xae, 0x11, 0x61, 0xaa, 0xf4, 0x88,
	0xa0, 0xaf, 0x61, 0xd6, 0x54, 0x54, 0x34, 0x9e, 0xa2, 0x69, 0x54, 0x5c, 0xb3, 0x0c, 0xc3, 0x7c,
	0x2b, 0x73, 0x55, 0xa4, 0x07, 0xac, 0xd8, 0x11, 0x6e, 0xef, 0x09, 0xdf, 0xf5, 0xe5, 0xb0, 0xa4,
	0xf0, 0x8f, 0x05, 0xe3, 0x2b, 0xb1, 0xfe, 0x98, 0xa6, 0xc3, 0x4a, 0xf5, 0x39, 0xf2, 0xa6, 0xe3,
	0xc8, 0x71, 0x7c, 0x25, 0xd6, 0x03, 0xdd, 0xa8, 0x6f, 0xed, 0x1b, 0x96, 0x1b, 0x7e, 0x8b, 0x87,
	0x27, 0x63, 0x90, 0xc6, 0xdd, 0x5b, 0x33, 0x55, 0x06, 0xea, 0xfc, 0x67, 0xc3, 0x73, 0x63, 0xf2,
	0x0f, 0x2c, 0x2b, 0x2e, 0xf2, 0xc1, 0x86, 0x9e, 0x01, 0x14, 0xac, 0xc4, 0x5c, 0xe3, 0x8e, 0xc2,
	0x27, 0x3a, 0x53, 0xc3, 0xbb, 0x62, 0xdc, 0x3d, 0x31, 0x74, 0x01, 0x9e, 0xe4, 0xf2, 0xae, 0x11,
	0xaa, 0x3f, 0xe8, 0x4b, 0x98, 0x48, 0x9e, 0x61, 0x25, 0x59, 0x56, 0xf8, 0xa3, 0xc0, 0x8a, 0x9c,
	0xa4, 0x4d, 0xd0, 0x10, 0x3c, 0x2e, 0x31, 0xab, 0xfc, 0xa3, 0xc0, 0x89, 0xa6, 0xcb, 0x59, 0x33,
	0x32, 0x97, 0x12, 0xb3, 0x44, 0x43, 0x34, 0x82, 0x11, 0xcf, 0x8b, 0x7b, 0x59, 0xf9, 0x63, 0x75,
	0x88, 0xc4, 0x6d, 0xcb, 0x97, 0x35, 0x90, 0x18, 0xbc, 0x6e, 0x32, 0xd5, 0x90, 0x6a, 0x62, 0xa2,
	0x9b, 0x34, 0xa9, 0x15, 0x4f, 0xc3, 0x9f, 0x30, 0xbd, 0xf8, 0x5d, 0x60, 0x29, 0x0f, 0xd9, 0xd4,
	0xf7, 0xda, 0xa3, 0xce, 0x54, 0x2f, 0x62, 0xcd, 0xd0, 0x37, 0xd4, 0x1f, 0xc0, 0xbb, 0xa8, 0x37,
	0x4d, 0x4d, 0xf3, 0x49, 0xa4, 0x0f, 0x86, 0x59, 0xc5, 0xf4, 0x14, 0xdc, 0xef, 0x0f, 0x85, 0xa6,
	0x9e, 0x2f, 0xbd, 0xb8, 0xfe, 0x48, 0x54, 0xea, 0xdd, 0x7f, 0x4b, 0x63, 0x74, 0x0c, 0xee, 0xb5,
	0xc8, 0x91, 0x3c, 0xa3, 0x3e, 0x2c, 0x14, 0xd5, 0x4d, 0x77, 0x41, 0x11, 0xab, 0x45, 0xba, 0x6b,
	0x85, 0xd8, 0x7b, 0x88, 0x79, 0xc9, 0xc4, 0xa1, 0x27, 0x70, 0xac, 0x91, 0x27, 0xdd, 0x13, 0x97,
	0x52, 0x98, 0xeb, 0x74, 0xf3, 0x12, 0x89, 0xd7, 0x92, 0x74, 0xe7, 0x9f, 0x8c, 0xe8, 0x19, 0x9c,
	0x76, 0xe8, 0x9f, 0x4e, 0x1c, 0x39, 0x5a, 0x8f, 0xd4, 0x4e, 0x7d, 0xff, 0x18, 0x00, 0x00, 0xff,
	0xff, 0x39, 0x76, 0x54, 0x23, 0x6d, 0x05, 0x00, 0x00,
}
