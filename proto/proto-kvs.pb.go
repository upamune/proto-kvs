// Code generated by protoc-gen-go.
// source: proto/proto-kvs.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/proto-kvs.proto

It has these top-level messages:
	GetRequest
	GetResponse
	SetRequest
	SetResponse
	DeleteRequest
	DeleteResponse
	ListRequest
	ListResponse
	WatchRequest
	Entry
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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

type WatchRequest struct {
	Prefix string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
}

func (m *WatchRequest) Reset()                    { *m = WatchRequest{} }
func (m *WatchRequest) String() string            { return proto1.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()               {}
func (*WatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *WatchRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type Entry struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto1.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Entry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
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
	proto1.RegisterType((*WatchRequest)(nil), "proto.WatchRequest")
	proto1.RegisterType((*Entry)(nil), "proto.Entry")
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
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Kvs_WatchClient, error)
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

func (c *kvsClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Kvs_WatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Kvs_serviceDesc.Streams[0], c.cc, "/proto.Kvs/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &kvsWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Kvs_WatchClient interface {
	Recv() (*Entry, error)
	grpc.ClientStream
}

type kvsWatchClient struct {
	grpc.ClientStream
}

func (x *kvsWatchClient) Recv() (*Entry, error) {
	m := new(Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Kvs service

type KvsServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Watch(*WatchRequest, Kvs_WatchServer) error
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

func _Kvs_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KvsServer).Watch(m, &kvsWatchServer{stream})
}

type Kvs_WatchServer interface {
	Send(*Entry) error
	grpc.ServerStream
}

type kvsWatchServer struct {
	grpc.ServerStream
}

func (x *kvsWatchServer) Send(m *Entry) error {
	return x.ServerStream.SendMsg(m)
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
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Kvs_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/proto-kvs.proto",
}

func init() { proto1.RegisterFile("proto/proto-kvs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x91, 0xdf, 0x8e, 0x9a, 0x40,
	0x14, 0x87, 0x83, 0x14, 0x93, 0x1e, 0xa4, 0xb5, 0x47, 0x69, 0x0c, 0x35, 0xa6, 0x9d, 0x36, 0x8d,
	0x31, 0xa9, 0xb4, 0xd6, 0x0b, 0xe3, 0x75, 0xab, 0x69, 0xea, 0x15, 0x5c, 0xf4, 0x9a, 0x6e, 0x66,
	0x5d, 0xa2, 0x61, 0x58, 0x18, 0xd9, 0x35, 0xc6, 0xbd, 0xf0, 0x15, 0xf6, 0xd1, 0xf6, 0x15, 0xf6,
	0x41, 0x36, 0x0c, 0x43, 0x00, 0xf7, 0x4f, 0x76, 0x6f, 0x60, 0xe6, 0x70, 0x7e, 0x1f, 0x73, 0xe6,
	0x03, 0x33, 0x8c, 0x18, 0x67, 0xb6, 0x78, 0x7e, 0x5b, 0x25, 0xf1, 0x50, 0xac, 0x50, 0x13, 0x2f,
	0xab, 0xbb, 0x64, 0x6c, 0xb9, 0xa6, 0xb6, 0x17, 0xfa, 0xb6, 0x17, 0x04, 0x8c, 0x7b, 0xdc, 0x67,
	0x81, 0x6c, 0x22, 0x3d, 0x80, 0x39, 0xe5, 0x0e, 0x3d, 0xdf, 0xd0, 0x98, 0x63, 0x13, 0xd4, 0x15,
	0xdd, 0x76, 0x94, 0x8f, 0x4a, 0xff, 0xb5, 0x93, 0x2e, 0xc9, 0x67, 0xd0, 0xc5, 0xf7, 0x38, 0x64,
	0x41, 0x4c, 0xb1, 0x0d, 0x5a, 0xe2, 0xad, 0x37, 0x54, 0xb6, 0x64, 0x1b, 0x32, 0x06, 0x70, 0x9f,
	0x80, 0x14, 0xa9, 0x5a, 0x39, 0x65, 0x80, 0xee, 0x16, 0x68, 0xf2, 0x09, 0x8c, 0x5f, 0x74, 0x4d,
	0x39, 0x7d, 0xfc, 0x30, 0x4d, 0x78, 0x93, 0xb7, 0xc8, 0x90, 0x01, 0xfa, 0xc2, 0x8f, 0xf3, 0x5f,
	0x93, 0x2b, 0x68, 0x64, 0x5b, 0x79, 0xdc, 0x31, 0x68, 0x31, 0x67, 0x51, 0x7a, 0x5c, 0xb5, 0xaf,
	0x8f, 0x7a, 0xd9, 0xd0, 0xc3, 0x72, 0xcf, 0xd0, 0x4d, 0x1b, 0x7e, 0x07, 0x3c, 0xda, 0x3a, 0x59,
	0xb3, 0x35, 0x01, 0x28, 0x8a, 0xcf, 0x1d, 0x67, 0x5a, 0x9b, 0x28, 0xe4, 0x2b, 0x34, 0xfe, 0x79,
	0xfc, 0xe4, 0x2c, 0x1f, 0xe1, 0x3d, 0xd4, 0xc3, 0x88, 0x9e, 0xfa, 0x97, 0x32, 0x2e, 0x77, 0xc4,
	0x06, 0xed, 0x45, 0xf0, 0xd1, 0x41, 0x05, 0xf5, 0x6f, 0x12, 0xe3, 0x1f, 0x50, 0xe7, 0x94, 0xe3,
	0x3b, 0x39, 0x48, 0xa1, 0xce, 0xc2, 0x72, 0x49, 0xde, 0xce, 0x87, 0xc3, 0xcd, 0xed, 0x75, 0xcd,
	0xc4, 0x96, 0x90, 0x9f, 0xfc, 0xb0, 0xc5, 0x7c, 0xf6, 0x6e, 0x45, 0xb7, 0x7b, 0x9c, 0x81, 0xea,
	0x96, 0x50, 0xee, 0x7d, 0x54, 0xd9, 0x4e, 0x47, 0xa0, 0x90, 0x18, 0x15, 0xd4, 0x54, 0x19, 0xe0,
	0x02, 0xea, 0x99, 0x14, 0x6c, 0xcb, 0x5c, 0x45, 0xa3, 0x65, 0x1e, 0x55, 0x25, 0xd0, 0x14, 0xc0,
	0xb7, 0x83, 0x2a, 0x10, 0x67, 0xf0, 0x2a, 0xb5, 0x83, 0x58, 0x51, 0x95, 0x91, 0x5a, 0x0f, 0xe8,
	0xcb, 0x39, 0x78, 0xc4, 0x71, 0x40, 0x13, 0x26, 0x30, 0x0f, 0x95, 0xbd, 0x58, 0x0d, 0x59, 0x14,
	0x12, 0xc8, 0x17, 0x81, 0xe8, 0x61, 0xb7, 0x7a, 0x4d, 0x17, 0x69, 0xc2, 0xde, 0x65, 0xca, 0xf6,
	0xdf, 0x95, 0xff, 0x75, 0x11, 0xfa, 0x79, 0x17, 0x00, 0x00, 0xff, 0xff, 0x50, 0xa9, 0xdf, 0x4c,
	0x70, 0x03, 0x00, 0x00,
}
