// Code generated by protoc-gen-go. DO NOT EDIT.
// source: news.proto

package domain

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetNewsResponse struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNewsResponse) Reset()         { *m = GetNewsResponse{} }
func (m *GetNewsResponse) String() string { return proto.CompactTextString(m) }
func (*GetNewsResponse) ProtoMessage()    {}
func (*GetNewsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0382e93bed6d84, []int{0}
}

func (m *GetNewsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNewsResponse.Unmarshal(m, b)
}
func (m *GetNewsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNewsResponse.Marshal(b, m, deterministic)
}
func (m *GetNewsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNewsResponse.Merge(m, src)
}
func (m *GetNewsResponse) XXX_Size() int {
	return xxx_messageInfo_GetNewsResponse.Size(m)
}
func (m *GetNewsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNewsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNewsResponse proto.InternalMessageInfo

func (m *GetNewsResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *GetNewsResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GetNewsResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type GetNewsRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNewsRequest) Reset()         { *m = GetNewsRequest{} }
func (m *GetNewsRequest) String() string { return proto.CompactTextString(m) }
func (*GetNewsRequest) ProtoMessage()    {}
func (*GetNewsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0382e93bed6d84, []int{1}
}

func (m *GetNewsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNewsRequest.Unmarshal(m, b)
}
func (m *GetNewsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNewsRequest.Marshal(b, m, deterministic)
}
func (m *GetNewsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNewsRequest.Merge(m, src)
}
func (m *GetNewsRequest) XXX_Size() int {
	return xxx_messageInfo_GetNewsRequest.Size(m)
}
func (m *GetNewsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNewsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNewsRequest proto.InternalMessageInfo

func (m *GetNewsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*GetNewsResponse)(nil), "domain.GetNewsResponse")
	proto.RegisterType((*GetNewsRequest)(nil), "domain.GetNewsRequest")
}

func init() { proto.RegisterFile("news.proto", fileDescriptor_2c0382e93bed6d84) }

var fileDescriptor_2c0382e93bed6d84 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4b, 0x2d, 0x2f,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x53, 0x0a,
	0xe7, 0xe2, 0x77, 0x4f, 0x2d, 0xf1, 0x4b, 0x2d, 0x2f, 0x0e, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b,
	0x4e, 0x15, 0x12, 0xe0, 0x62, 0x2e, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0x31, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25, 0x98, 0xc0, 0x62, 0x10, 0x8e,
	0x90, 0x0c, 0x17, 0x67, 0x49, 0x66, 0x6e, 0x6a, 0x71, 0x49, 0x62, 0x6e, 0x81, 0x04, 0xb3, 0x02,
	0xa3, 0x06, 0x73, 0x10, 0x42, 0x40, 0x49, 0x81, 0x8b, 0x0f, 0x6e, 0x70, 0x61, 0x69, 0x6a, 0x71,
	0x89, 0x10, 0x1f, 0x17, 0x13, 0xdc, 0x58, 0xa6, 0xcc, 0x14, 0x23, 0x4f, 0x2e, 0x6e, 0x90, 0x74,
	0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x15, 0x17, 0x3b, 0x54, 0x83, 0x90, 0x98, 0x1e,
	0xc4, 0x75, 0x7a, 0xa8, 0x26, 0x48, 0x89, 0x63, 0x88, 0x43, 0x9c, 0x9c, 0xc4, 0x06, 0xf6, 0x94,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xea, 0x38, 0x17, 0x5c, 0xe2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NewsServiceClient is the client API for NewsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NewsServiceClient interface {
	GetNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*GetNewsResponse, error)
}

type newsServiceClient struct {
	cc *grpc.ClientConn
}

func NewNewsServiceClient(cc *grpc.ClientConn) NewsServiceClient {
	return &newsServiceClient{cc}
}

func (c *newsServiceClient) GetNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*GetNewsResponse, error) {
	out := new(GetNewsResponse)
	err := c.cc.Invoke(ctx, "/domain.NewsService/GetNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsServiceServer is the server API for NewsService service.
type NewsServiceServer interface {
	GetNews(context.Context, *GetNewsRequest) (*GetNewsResponse, error)
}

// UnimplementedNewsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNewsServiceServer struct {
}

func (*UnimplementedNewsServiceServer) GetNews(ctx context.Context, req *GetNewsRequest) (*GetNewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNews not implemented")
}

func RegisterNewsServiceServer(s *grpc.Server, srv NewsServiceServer) {
	s.RegisterService(&_NewsService_serviceDesc, srv)
}

func _NewsService_GetNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.NewsService/GetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetNews(ctx, req.(*GetNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NewsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "domain.NewsService",
	HandlerType: (*NewsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNews",
			Handler:    _NewsService_GetNews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news.proto",
}
