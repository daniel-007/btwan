// Code generated by protoc-gen-go.
// source: btwanapi.proto
// DO NOT EDIT!

/*
Package btwan is a generated protocol buffer package.

It is generated from these files:
	btwanapi.proto

It has these top-level messages:
	Void
	InfoHash
	FileInfo
	MetadataInfo
	SearchReq
	SearchResp
	Event
*/
package btwan

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type InfoHash struct {
	Ih string `protobuf:"bytes,1,opt,name=ih" json:"ih,omitempty"`
	ID uint64 `protobuf:"varint,2,opt,name=ID" json:"ID,omitempty"`
}

func (m *InfoHash) Reset()                    { *m = InfoHash{} }
func (m *InfoHash) String() string            { return proto.CompactTextString(m) }
func (*InfoHash) ProtoMessage()               {}
func (*InfoHash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type FileInfo struct {
	Path   []string `protobuf:"bytes,1,rep,name=path" json:"path,omitempty"`
	Length uint64   `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
}

func (m *FileInfo) Reset()                    { *m = FileInfo{} }
func (m *FileInfo) String() string            { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()               {}
func (*FileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type MetadataInfo struct {
	ID          uint64      `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	InfoHash    string      `protobuf:"bytes,2,opt,name=infoHash" json:"infoHash,omitempty"`
	Name        string      `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Files       []*FileInfo `protobuf:"bytes,4,rep,name=files" json:"files,omitempty"`
	Length      uint64      `protobuf:"varint,5,opt,name=length" json:"length,omitempty"`
	CollectTime int64       `protobuf:"varint,6,opt,name=collectTime" json:"collectTime,omitempty"`
	IndexTime   int64       `protobuf:"varint,7,opt,name=indexTime" json:"indexTime,omitempty"`
	Degree      uint64      `protobuf:"varint,8,opt,name=degree" json:"degree,omitempty"`
	Reviews     uint64      `protobuf:"varint,9,opt,name=reviews" json:"reviews,omitempty"`
	Follows     uint64      `protobuf:"varint,10,opt,name=follows" json:"follows,omitempty"`
	Thumbs      uint64      `protobuf:"varint,11,opt,name=thumbs" json:"thumbs,omitempty"`
	Seeders     uint64      `protobuf:"varint,12,opt,name=seeders" json:"seeders,omitempty"`
	Downloaders uint64      `protobuf:"varint,13,opt,name=downloaders" json:"downloaders,omitempty"`
}

func (m *MetadataInfo) Reset()                    { *m = MetadataInfo{} }
func (m *MetadataInfo) String() string            { return proto.CompactTextString(m) }
func (*MetadataInfo) ProtoMessage()               {}
func (*MetadataInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MetadataInfo) GetFiles() []*FileInfo {
	if m != nil {
		return m.Files
	}
	return nil
}

type SearchReq struct {
	Q      string `protobuf:"bytes,1,opt,name=q" json:"q,omitempty"`
	Offset uint32 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	Limit  uint32 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *SearchReq) Reset()                    { *m = SearchReq{} }
func (m *SearchReq) String() string            { return proto.CompactTextString(m) }
func (*SearchReq) ProtoMessage()               {}
func (*SearchReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SearchResp struct {
	Request    *SearchReq      `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	TotalCount uint32          `protobuf:"varint,2,opt,name=totalCount" json:"totalCount,omitempty"`
	Count      uint32          `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Took       uint32          `protobuf:"varint,4,opt,name=took" json:"took,omitempty"`
	Metainfos  []*MetadataInfo `protobuf:"bytes,5,rep,name=metainfos" json:"metainfos,omitempty"`
}

func (m *SearchResp) Reset()                    { *m = SearchResp{} }
func (m *SearchResp) String() string            { return proto.CompactTextString(m) }
func (*SearchResp) ProtoMessage()               {}
func (*SearchResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SearchResp) GetRequest() *SearchReq {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *SearchResp) GetMetainfos() []*MetadataInfo {
	if m != nil {
		return m.Metainfos
	}
	return nil
}

type Event struct {
	Type       string            `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Attributes map[string]string `protobuf:"bytes,2,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Event) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterType((*Void)(nil), "btwan.Void")
	proto.RegisterType((*InfoHash)(nil), "btwan.InfoHash")
	proto.RegisterType((*FileInfo)(nil), "btwan.FileInfo")
	proto.RegisterType((*MetadataInfo)(nil), "btwan.MetadataInfo")
	proto.RegisterType((*SearchReq)(nil), "btwan.SearchReq")
	proto.RegisterType((*SearchResp)(nil), "btwan.SearchResp")
	proto.RegisterType((*Event)(nil), "btwan.Event")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for OwstoniService service

type OwstoniServiceClient interface {
	Send(ctx context.Context, opts ...grpc.CallOption) (OwstoniService_SendClient, error)
	Recv(ctx context.Context, in *Void, opts ...grpc.CallOption) (OwstoniService_RecvClient, error)
	SendInfoHash(ctx context.Context, in *InfoHash, opts ...grpc.CallOption) (*Void, error)
	GetMetadataInfo(ctx context.Context, in *InfoHash, opts ...grpc.CallOption) (*MetadataInfo, error)
	Index(ctx context.Context, in *MetadataInfo, opts ...grpc.CallOption) (*Void, error)
	Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchResp, error)
}

type owstoniServiceClient struct {
	cc *grpc.ClientConn
}

func NewOwstoniServiceClient(cc *grpc.ClientConn) OwstoniServiceClient {
	return &owstoniServiceClient{cc}
}

func (c *owstoniServiceClient) Send(ctx context.Context, opts ...grpc.CallOption) (OwstoniService_SendClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_OwstoniService_serviceDesc.Streams[0], c.cc, "/btwan.OwstoniService/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &owstoniServiceSendClient{stream}
	return x, nil
}

type OwstoniService_SendClient interface {
	Send(*Event) error
	CloseAndRecv() (*Void, error)
	grpc.ClientStream
}

type owstoniServiceSendClient struct {
	grpc.ClientStream
}

func (x *owstoniServiceSendClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *owstoniServiceSendClient) CloseAndRecv() (*Void, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Void)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *owstoniServiceClient) Recv(ctx context.Context, in *Void, opts ...grpc.CallOption) (OwstoniService_RecvClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_OwstoniService_serviceDesc.Streams[1], c.cc, "/btwan.OwstoniService/Recv", opts...)
	if err != nil {
		return nil, err
	}
	x := &owstoniServiceRecvClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OwstoniService_RecvClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type owstoniServiceRecvClient struct {
	grpc.ClientStream
}

func (x *owstoniServiceRecvClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *owstoniServiceClient) SendInfoHash(ctx context.Context, in *InfoHash, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/btwan.OwstoniService/SendInfoHash", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *owstoniServiceClient) GetMetadataInfo(ctx context.Context, in *InfoHash, opts ...grpc.CallOption) (*MetadataInfo, error) {
	out := new(MetadataInfo)
	err := grpc.Invoke(ctx, "/btwan.OwstoniService/GetMetadataInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *owstoniServiceClient) Index(ctx context.Context, in *MetadataInfo, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/btwan.OwstoniService/Index", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *owstoniServiceClient) Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchResp, error) {
	out := new(SearchResp)
	err := grpc.Invoke(ctx, "/btwan.OwstoniService/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OwstoniService service

type OwstoniServiceServer interface {
	Send(OwstoniService_SendServer) error
	Recv(*Void, OwstoniService_RecvServer) error
	SendInfoHash(context.Context, *InfoHash) (*Void, error)
	GetMetadataInfo(context.Context, *InfoHash) (*MetadataInfo, error)
	Index(context.Context, *MetadataInfo) (*Void, error)
	Search(context.Context, *SearchReq) (*SearchResp, error)
}

func RegisterOwstoniServiceServer(s *grpc.Server, srv OwstoniServiceServer) {
	s.RegisterService(&_OwstoniService_serviceDesc, srv)
}

func _OwstoniService_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OwstoniServiceServer).Send(&owstoniServiceSendServer{stream})
}

type OwstoniService_SendServer interface {
	SendAndClose(*Void) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type owstoniServiceSendServer struct {
	grpc.ServerStream
}

func (x *owstoniServiceSendServer) SendAndClose(m *Void) error {
	return x.ServerStream.SendMsg(m)
}

func (x *owstoniServiceSendServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OwstoniService_Recv_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OwstoniServiceServer).Recv(m, &owstoniServiceRecvServer{stream})
}

type OwstoniService_RecvServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type owstoniServiceRecvServer struct {
	grpc.ServerStream
}

func (x *owstoniServiceRecvServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _OwstoniService_SendInfoHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwstoniServiceServer).SendInfoHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btwan.OwstoniService/SendInfoHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwstoniServiceServer).SendInfoHash(ctx, req.(*InfoHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwstoniService_GetMetadataInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwstoniServiceServer).GetMetadataInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btwan.OwstoniService/GetMetadataInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwstoniServiceServer).GetMetadataInfo(ctx, req.(*InfoHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwstoniService_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetadataInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwstoniServiceServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btwan.OwstoniService/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwstoniServiceServer).Index(ctx, req.(*MetadataInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwstoniService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwstoniServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btwan.OwstoniService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwstoniServiceServer).Search(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _OwstoniService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "btwan.OwstoniService",
	HandlerType: (*OwstoniServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendInfoHash",
			Handler:    _OwstoniService_SendInfoHash_Handler,
		},
		{
			MethodName: "GetMetadataInfo",
			Handler:    _OwstoniService_GetMetadataInfo_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _OwstoniService_Index_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _OwstoniService_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _OwstoniService_Send_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Recv",
			Handler:       _OwstoniService_Recv_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("btwanapi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x8d, 0x9d, 0x9f, 0xc6, 0x93, 0xf4, 0xe7, 0xdb, 0x0f, 0xa1, 0x55, 0x54, 0xa1, 0xc8, 0x12,
	0x52, 0x54, 0xa4, 0x00, 0x45, 0x42, 0x80, 0xe0, 0x02, 0xd1, 0x52, 0x72, 0x81, 0x90, 0xb6, 0x88,
	0xfb, 0x4d, 0x3c, 0x69, 0x56, 0x75, 0xbc, 0x89, 0x77, 0x93, 0xd0, 0x37, 0xe1, 0x0d, 0x78, 0x02,
	0x1e, 0x8b, 0x77, 0x40, 0x3b, 0x5e, 0xa7, 0x6e, 0xda, 0xbb, 0x39, 0x73, 0x66, 0xce, 0x8c, 0x67,
	0x67, 0x0c, 0x07, 0x63, 0xbb, 0x91, 0x99, 0x5c, 0xa8, 0xe1, 0x22, 0xd7, 0x56, 0xb3, 0x26, 0xe1,
	0xb8, 0x05, 0x8d, 0x1f, 0x5a, 0x25, 0xf1, 0x09, 0xb4, 0x47, 0xd9, 0x54, 0x7f, 0x91, 0x66, 0xc6,
	0x0e, 0x20, 0x54, 0x33, 0x1e, 0xf4, 0x83, 0x41, 0x24, 0x42, 0x45, 0x78, 0x74, 0xc6, 0xc3, 0x7e,
	0x30, 0x68, 0x88, 0x70, 0x74, 0x16, 0xbf, 0x86, 0xf6, 0x67, 0x95, 0xa2, 0x8b, 0x67, 0x0c, 0x1a,
	0x0b, 0x69, 0x5d, 0x74, 0x7d, 0x10, 0x09, 0xb2, 0xd9, 0x63, 0x68, 0xa5, 0x98, 0x5d, 0xd9, 0x99,
	0xcf, 0xf1, 0x28, 0xfe, 0x1b, 0x42, 0xf7, 0x2b, 0x5a, 0x99, 0x48, 0x2b, 0x29, 0xb9, 0x10, 0x0e,
	0x4a, 0x61, 0xd6, 0x83, 0xb6, 0xf2, 0x4d, 0x50, 0x6a, 0x24, 0xb6, 0xd8, 0x15, 0xca, 0xe4, 0x1c,
	0x79, 0x9d, 0xfc, 0x64, 0xb3, 0xa7, 0xd0, 0x9c, 0xaa, 0x14, 0x0d, 0x6f, 0xf4, 0xeb, 0x83, 0xce,
	0xe9, 0xe1, 0x90, 0xbe, 0x69, 0x58, 0x36, 0x27, 0x0a, 0xb6, 0xd2, 0x4f, 0xb3, 0xda, 0x0f, 0xeb,
	0x43, 0x67, 0xa2, 0xd3, 0x14, 0x27, 0xf6, 0xbb, 0x9a, 0x23, 0x6f, 0xf5, 0x83, 0x41, 0x5d, 0x54,
	0x5d, 0xec, 0x18, 0x22, 0x95, 0x25, 0xf8, 0x93, 0xf8, 0x3d, 0xe2, 0x6f, 0x1d, 0x4e, 0x37, 0xc1,
	0xab, 0x1c, 0x91, 0xb7, 0x0b, 0xdd, 0x02, 0x31, 0x0e, 0x7b, 0x39, 0xae, 0x15, 0x6e, 0x0c, 0x8f,
	0x88, 0x28, 0xa1, 0x63, 0xa6, 0x3a, 0x4d, 0xf5, 0xc6, 0x70, 0x28, 0x18, 0x0f, 0x9d, 0x96, 0x9d,
	0xad, 0xe6, 0x63, 0xc3, 0x3b, 0x85, 0x56, 0x81, 0x5c, 0x86, 0x41, 0x4c, 0x30, 0x37, 0xbc, 0x5b,
	0x64, 0x78, 0xe8, 0xba, 0x4f, 0xf4, 0x26, 0x4b, 0xb5, 0x24, 0x76, 0x9f, 0xd8, 0xaa, 0x2b, 0xbe,
	0x80, 0xe8, 0x12, 0x65, 0x3e, 0x99, 0x09, 0x5c, 0xb2, 0x2e, 0x04, 0x4b, 0xff, 0xa6, 0xc1, 0xd2,
	0x95, 0xd3, 0xd3, 0xa9, 0x41, 0x4b, 0x73, 0xde, 0x17, 0x1e, 0xb1, 0x47, 0xd0, 0x4c, 0xd5, 0x5c,
	0x59, 0x1a, 0xf3, 0xbe, 0x28, 0x40, 0xfc, 0x27, 0x00, 0x28, 0x95, 0xcc, 0x82, 0x9d, 0xb8, 0xef,
	0x5b, 0xae, 0xd0, 0x58, 0x12, 0xec, 0x9c, 0x1e, 0xf9, 0xc1, 0x6f, 0xab, 0x89, 0x32, 0x80, 0x3d,
	0x01, 0xb0, 0xda, 0xca, 0xf4, 0x93, 0x5e, 0x65, 0x65, 0xb1, 0x8a, 0xc7, 0x15, 0x9c, 0x10, 0xe5,
	0x0b, 0x12, 0x70, 0x8f, 0x6d, 0xb5, 0xbe, 0xe6, 0x0d, 0x72, 0x92, 0xcd, 0x5e, 0x42, 0x34, 0x47,
	0x2b, 0xdd, 0x42, 0x18, 0xde, 0xa4, 0x07, 0xff, 0xdf, 0xd7, 0xad, 0x2e, 0x95, 0xb8, 0x8d, 0x8a,
	0x7f, 0x05, 0xd0, 0x3c, 0x5f, 0xa3, 0x17, 0xbc, 0x59, 0xa0, 0x1f, 0x00, 0xd9, 0xec, 0x3d, 0x80,
	0xb4, 0x36, 0x57, 0xe3, 0x95, 0x45, 0xc3, 0x43, 0x52, 0x3c, 0xf6, 0x8a, 0x94, 0x35, 0xfc, 0xb8,
	0xa5, 0xcf, 0x33, 0x9b, 0xdf, 0x88, 0x4a, 0x7c, 0xef, 0x03, 0x1c, 0xee, 0xd0, 0xec, 0x08, 0xea,
	0xd7, 0x78, 0xe3, 0x6b, 0x38, 0xd3, 0x7d, 0xdd, 0x5a, 0xa6, 0x2b, 0xf4, 0xdb, 0x5c, 0x80, 0x77,
	0xe1, 0x9b, 0xe0, 0xf4, 0x77, 0x08, 0x07, 0xdf, 0x36, 0xc6, 0xea, 0x4c, 0x5d, 0x62, 0xbe, 0x56,
	0x13, 0xb7, 0xcd, 0x8d, 0x4b, 0xcc, 0x12, 0xd6, 0xad, 0xf6, 0xd0, 0xeb, 0x78, 0x44, 0x57, 0x5a,
	0x1b, 0x04, 0x2e, 0x4c, 0xe0, 0x64, 0xcd, 0xaa, 0x44, 0xef, 0x4e, 0x4e, 0x5c, 0x7b, 0x11, 0xb0,
	0x21, 0x74, 0x9d, 0xda, 0xf6, 0xa8, 0xcb, 0xe3, 0x28, 0x1d, 0x3b, 0xc2, 0xec, 0x2d, 0x1c, 0x5e,
	0xa0, 0xbd, 0x73, 0x9e, 0xf7, 0x52, 0x1e, 0x9a, 0x77, 0x5c, 0x63, 0xcf, 0xa0, 0x39, 0x72, 0x47,
	0xc1, 0x1e, 0xe2, 0x77, 0xeb, 0x3c, 0x87, 0x56, 0xb1, 0x26, 0xec, 0xde, 0xd6, 0xf4, 0xfe, 0xdb,
	0xf1, 0x98, 0x45, 0x5c, 0x1b, 0xb7, 0xe8, 0x7f, 0xf5, 0xea, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd4, 0x9a, 0xf5, 0x75, 0xc1, 0x04, 0x00, 0x00,
}