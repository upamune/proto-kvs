// Code generated by protoc-gen-go.
// source: proto-kvs.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto-kvs.proto

It has these top-level messages:
	GetRequest
	GetResponse
	SetRequest
	SetResponse
	DeleteRequest
	DeleteResponse
	ListRequest
	ListResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResponse struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetRequest struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *SetRequest) Reset()                    { *m = SetRequest{} }
func (m *SetRequest) String() string            { return proto1.CompactTextString(m) }
func (*SetRequest) ProtoMessage()               {}
func (*SetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetResponse struct {
}

func (m *SetResponse) Reset()                    { *m = SetResponse{} }
func (m *SetResponse) String() string            { return proto1.CompactTextString(m) }
func (*SetResponse) ProtoMessage()               {}
func (*SetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DeleteRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto1.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ListResponse struct {
	Store map[string]string `protobuf:"bytes,1,rep,name=store" json:"store,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto1.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListResponse) GetStore() map[string]string {
	if m != nil {
		return m.Store
	}
	return nil
}

func init() {
	proto1.RegisterType((*GetRequest)(nil), "proto.GetRequest")
	proto1.RegisterType((*GetResponse)(nil), "proto.GetResponse")
	proto1.RegisterType((*SetRequest)(nil), "proto.SetRequest")
	proto1.RegisterType((*SetResponse)(nil), "proto.SetResponse")
	proto1.RegisterType((*DeleteRequest)(nil), "proto.DeleteRequest")
	proto1.RegisterType((*DeleteResponse)(nil), "proto.DeleteResponse")
	proto1.RegisterType((*ListRequest)(nil), "proto.ListRequest")
	proto1.RegisterType((*ListResponse)(nil), "proto.ListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Kvs service

type KvsClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type kvsClient struct {
	cc *grpc.ClientConn
}

func NewKvsClient(cc *grpc.ClientConn) KvsClient {
	return &kvsClient{cc}
}

func (c *kvsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/proto.Kvs/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvsClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := grpc.Invoke(ctx, "/proto.Kvs/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvsClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/proto.Kvs/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/proto.Kvs/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Kvs service

type KvsServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterKvsServer(s *grpc.Server, srv KvsServer) {
	s.RegisterService(&_Kvs_serviceDesc, srv)
}

func _Kvs_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Kvs/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kvs_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvsServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Kvs/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvsServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kvs_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Kvs/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvsServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kvs_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Kvs/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Kvs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Kvs",
	HandlerType: (*KvsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Kvs_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Kvs_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Kvs_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Kvs_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto-kvs.proto",
}

func init() { proto1.RegisterFile("proto-kvs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x18, 0xc4, 0x9b, 0x86, 0x54, 0xe2, 0x42, 0xa0, 0x7c, 0x14, 0x09, 0x65, 0xa8, 0xc0, 0x2c, 0x5d,
	0x88, 0x44, 0xa9, 0x44, 0xc5, 0x0c, 0xea, 0x00, 0x53, 0xfc, 0x04, 0x20, 0x7d, 0x03, 0x6a, 0xd5,
	0x94, 0xd8, 0x8d, 0xd4, 0x85, 0x77, 0xe4, 0x8d, 0x90, 0x63, 0x97, 0x38, 0xfc, 0x53, 0x27, 0xdb,
	0xe7, 0xdf, 0xd9, 0xdf, 0x1d, 0x8e, 0x56, 0x65, 0xa1, 0x8b, 0xab, 0x79, 0xa5, 0xb2, 0x7a, 0x47,
	0x51, 0xbd, 0x88, 0x21, 0x30, 0x63, 0x9d, 0xf3, 0xdb, 0x9a, 0x95, 0xa6, 0x3e, 0xc2, 0x39, 0x6f,
	0xce, 0x82, 0xf3, 0x60, 0xb4, 0x9f, 0x9b, 0xad, 0xb8, 0x44, 0x5c, 0xdf, 0xab, 0x55, 0xb1, 0x54,
	0x4c, 0x03, 0x44, 0xd5, 0xf3, 0x62, 0xcd, 0x0e, 0xb1, 0x07, 0x31, 0x01, 0xe4, 0x3f, 0x8f, 0x34,
	0xae, 0xae, 0xef, 0x4a, 0x10, 0xcb, 0xe6, 0x69, 0x71, 0x81, 0xe4, 0x9e, 0x17, 0xac, 0xf9, 0xef,
	0x61, 0xfa, 0x38, 0xdc, 0x22, 0xce, 0x94, 0x20, 0x7e, 0x7a, 0x55, 0xdb, 0xaf, 0xc5, 0x3b, 0x0e,
	0xec, 0xd1, 0x8d, 0x3b, 0x41, 0xa4, 0x74, 0x51, 0x9a, 0x71, 0xc3, 0x51, 0x3c, 0x1e, 0xda, 0xec,
	0x99, 0xcf, 0x64, 0xd2, 0x00, 0x0f, 0x4b, 0x5d, 0x6e, 0x72, 0x0b, 0xa7, 0x53, 0xa0, 0x11, 0x77,
	0x8d, 0x73, 0xd7, 0x9d, 0x06, 0xe3, 0x8f, 0x00, 0xe1, 0x63, 0xa5, 0x28, 0x43, 0x38, 0x63, 0x4d,
	0xc7, 0xee, 0xbf, 0xa6, 0xe1, 0x94, 0x7c, 0xc9, 0x85, 0xe8, 0x18, 0x5e, 0x7a, 0xbc, 0xfc, 0xc9,
	0xcb, 0x16, 0x7f, 0x8b, 0x9e, 0x2d, 0x82, 0x06, 0xee, 0xbe, 0x55, 0x5d, 0x7a, 0xfa, 0x4d, 0xfd,
	0x32, 0x5e, 0x63, 0xcf, 0x84, 0x27, 0x6a, 0x35, 0x61, 0x4d, 0x27, 0xbf, 0xb4, 0x23, 0x3a, 0x2f,
	0xbd, 0x5a, 0xbd, 0xf9, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xe0, 0x43, 0xe4, 0x42, 0x02, 0x00,
	0x00,
}
