// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	CreateProjectRequest
	CreateProjectResponse
	OkResponse
	ScenarioRequest
	ScenarioResponse
	KillRequest
	PingRequest
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import events "mlp/catalog/events"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateProjectRequest struct {
}

func (m *CreateProjectRequest) Reset()                    { *m = CreateProjectRequest{} }
func (m *CreateProjectRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateProjectRequest) ProtoMessage()               {}
func (*CreateProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateProjectResponse struct {
}

func (m *CreateProjectResponse) Reset()                    { *m = CreateProjectResponse{} }
func (m *CreateProjectResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateProjectResponse) ProtoMessage()               {}
func (*CreateProjectResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type OkResponse struct {
}

func (m *OkResponse) Reset()                    { *m = OkResponse{} }
func (m *OkResponse) String() string            { return proto.CompactTextString(m) }
func (*OkResponse) ProtoMessage()               {}
func (*OkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ScenarioRequest struct {
	Name      string          `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Events    []*events.Event `protobuf:"bytes,2,rep,name=Events" json:"Events,omitempty"`
	Timestamp int64           `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *ScenarioRequest) Reset()                    { *m = ScenarioRequest{} }
func (m *ScenarioRequest) String() string            { return proto.CompactTextString(m) }
func (*ScenarioRequest) ProtoMessage()               {}
func (*ScenarioRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ScenarioRequest) GetEvents() []*events.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type ScenarioResponse struct {
}

func (m *ScenarioResponse) Reset()                    { *m = ScenarioResponse{} }
func (m *ScenarioResponse) String() string            { return proto.CompactTextString(m) }
func (*ScenarioResponse) ProtoMessage()               {}
func (*ScenarioResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type KillRequest struct {
}

func (m *KillRequest) Reset()                    { *m = KillRequest{} }
func (m *KillRequest) String() string            { return proto.CompactTextString(m) }
func (*KillRequest) ProtoMessage()               {}
func (*KillRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*CreateProjectRequest)(nil), "CreateProjectRequest")
	proto.RegisterType((*CreateProjectResponse)(nil), "CreateProjectResponse")
	proto.RegisterType((*OkResponse)(nil), "OkResponse")
	proto.RegisterType((*ScenarioRequest)(nil), "ScenarioRequest")
	proto.RegisterType((*ScenarioResponse)(nil), "ScenarioResponse")
	proto.RegisterType((*KillRequest)(nil), "KillRequest")
	proto.RegisterType((*PingRequest)(nil), "PingRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Catalog service

type CatalogClient interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
}

type catalogClient struct {
	cc *grpc.ClientConn
}

func NewCatalogClient(cc *grpc.ClientConn) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	out := new(CreateProjectResponse)
	err := grpc.Invoke(ctx, "/Catalog/CreateProject", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Catalog service

type CatalogServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
}

func RegisterCatalogServer(s *grpc.Server, srv CatalogServer) {
	s.RegisterService(&_Catalog_serviceDesc, srv)
}

func _Catalog_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Catalog/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Catalog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _Catalog_CreateProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

// Client API for Test service

type TestClient interface {
	// Setup a given state in the database
	Setup(ctx context.Context, in *ScenarioRequest, opts ...grpc.CallOption) (*OkResponse, error)
	Kill(ctx context.Context, in *KillRequest, opts ...grpc.CallOption) (*OkResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*OkResponse, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) Setup(ctx context.Context, in *ScenarioRequest, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := grpc.Invoke(ctx, "/Test/Setup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Kill(ctx context.Context, in *KillRequest, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := grpc.Invoke(ctx, "/Test/Kill", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := grpc.Invoke(ctx, "/Test/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Test service

type TestServer interface {
	// Setup a given state in the database
	Setup(context.Context, *ScenarioRequest) (*OkResponse, error)
	Kill(context.Context, *KillRequest) (*OkResponse, error)
	Ping(context.Context, *PingRequest) (*OkResponse, error)
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScenarioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Setup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test/Setup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Setup(ctx, req.(*ScenarioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Kill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Kill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test/Kill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Kill(ctx, req.(*KillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Setup",
			Handler:    _Test_Setup_Handler,
		},
		{
			MethodName: "Kill",
			Handler:    _Test_Kill_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Test_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xd1, 0x4a, 0x84, 0x40,
	0x14, 0x86, 0xd7, 0x34, 0xc3, 0xa3, 0xd2, 0x72, 0x68, 0x37, 0x91, 0x08, 0xb1, 0x1b, 0xaf, 0xe6,
	0xc2, 0x5e, 0x20, 0x58, 0xba, 0x5a, 0xa8, 0xc5, 0xed, 0x05, 0x26, 0x39, 0x2c, 0x53, 0xea, 0x4c,
	0xce, 0x6c, 0x17, 0x3d, 0x7d, 0x68, 0xd6, 0x9a, 0x78, 0x37, 0xff, 0xe1, 0x0c, 0xff, 0xf7, 0xcd,
	0x80, 0xc7, 0x95, 0x60, 0xaa, 0x95, 0x46, 0xc6, 0x01, 0x7d, 0x52, 0x63, 0xf4, 0x4f, 0x4a, 0xd7,
	0x70, 0xb5, 0x69, 0x89, 0x1b, 0xda, 0xb5, 0xf2, 0x8d, 0x4a, 0x53, 0xd0, 0xc7, 0x91, 0xb4, 0x49,
	0xaf, 0x61, 0x35, 0x99, 0x6b, 0x25, 0x1b, 0x4d, 0x69, 0x00, 0xf0, 0xfc, 0xfe, 0x97, 0x4a, 0xb8,
	0xdc, 0x97, 0xd4, 0xf0, 0x56, 0xc8, 0xe1, 0x26, 0x22, 0x38, 0x4f, 0xbc, 0xa6, 0xc8, 0x4a, 0xac,
	0xcc, 0x2b, 0xfa, 0x33, 0xde, 0x82, 0xfb, 0xd8, 0xb7, 0x46, 0x67, 0x89, 0x9d, 0xf9, 0xb9, 0xcb,
	0xfa, 0x58, 0x0c, 0x53, 0xbc, 0x01, 0xcf, 0x88, 0x9a, 0xb4, 0xe1, 0xb5, 0x8a, 0xec, 0xc4, 0xca,
	0xec, 0xe2, 0x34, 0x48, 0x11, 0x96, 0xa7, 0x92, 0xa1, 0x38, 0x04, 0x7f, 0x2b, 0xaa, 0xea, 0x17,
	0x37, 0x04, 0x7f, 0x27, 0x9a, 0xc3, 0x10, 0xf3, 0x2d, 0x5c, 0x6c, 0xb8, 0xe1, 0x95, 0x3c, 0xe0,
	0x03, 0x84, 0xff, 0x44, 0x70, 0xc5, 0xe6, 0x84, 0xe3, 0x35, 0x9b, 0xf7, 0x5d, 0xe4, 0x5f, 0xe0,
	0xbc, 0x74, 0x62, 0x19, 0x9c, 0xef, 0xc9, 0x1c, 0x15, 0x2e, 0xd9, 0xc4, 0x39, 0xf6, 0xd9, 0xe8,
	0x4d, 0x16, 0x78, 0x07, 0x4e, 0x07, 0x87, 0x01, 0x1b, 0x31, 0xce, 0x2c, 0x75, 0xc8, 0x18, 0xb0,
	0x11, 0xf9, 0x64, 0xe9, 0xd5, 0xed, 0x7f, 0xe9, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x35, 0xbc,
	0x9c, 0x3c, 0xc0, 0x01, 0x00, 0x00,
}