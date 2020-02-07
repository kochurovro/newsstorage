// Code generated by protoc-gen-go. DO NOT EDIT.
// source: news.proto

package spec

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
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
	proto.RegisterType((*GetNewsResponse)(nil), "web.GetNewsResponse")
	proto.RegisterType((*GetNewsRequest)(nil), "web.GetNewsRequest")
}

func init() { proto.RegisterFile("news.proto", fileDescriptor_2c0382e93bed6d84) }

var fileDescriptor_2c0382e93bed6d84 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4b, 0x2d, 0x2f,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x4f, 0x4d, 0x52, 0x72, 0xe5, 0xe2, 0x77,
	0x4f, 0x2d, 0xf1, 0x4b, 0x2d, 0x2f, 0x0e, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12,
	0xe1, 0x62, 0x2d, 0xc9, 0x2c, 0xc9, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70,
	0x84, 0x64, 0xb8, 0x38, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0x24, 0x98, 0x14,
	0x18, 0x35, 0x98, 0x83, 0x10, 0x02, 0x4a, 0x0a, 0x5c, 0x7c, 0x70, 0x63, 0x0a, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0xa0, 0x46, 0x30, 0x65, 0xa6, 0x18, 0x39, 0x73, 0x71,
	0x83, 0xa4, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x4c, 0xb8, 0xd8, 0xa1, 0x1a, 0x84,
	0x84, 0xf5, 0xca, 0x53, 0x93, 0xf4, 0x50, 0xb5, 0x4b, 0x89, 0xa0, 0x0a, 0x42, 0x9c, 0xe6, 0xc4,
	0x16, 0xc5, 0x52, 0x5c, 0x90, 0x9a, 0x9c, 0xc4, 0x06, 0xf6, 0x81, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x89, 0x82, 0xaf, 0x53, 0xcf, 0x00, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/web.NewsService/GetNews", in, out, opts...)
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
		FullMethod: "/web.NewsService/GetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetNews(ctx, req.(*GetNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NewsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "web.NewsService",
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
