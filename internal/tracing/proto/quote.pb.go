// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quote.proto

package proto

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

type QuotesRequest struct {
	SKU                  string   `protobuf:"bytes,1,opt,name=SKU,proto3" json:"SKU,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotesRequest) Reset()         { *m = QuotesRequest{} }
func (m *QuotesRequest) String() string { return proto.CompactTextString(m) }
func (*QuotesRequest) ProtoMessage()    {}
func (*QuotesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12e4b46109c2c385, []int{0}
}

func (m *QuotesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotesRequest.Unmarshal(m, b)
}
func (m *QuotesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotesRequest.Marshal(b, m, deterministic)
}
func (m *QuotesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotesRequest.Merge(m, src)
}
func (m *QuotesRequest) XXX_Size() int {
	return xxx_messageInfo_QuotesRequest.Size(m)
}
func (m *QuotesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuotesRequest proto.InternalMessageInfo

func (m *QuotesRequest) GetSKU() string {
	if m != nil {
		return m.SKU
	}
	return ""
}

type QuotesResponse struct {
	Quotes               []*Quote `protobuf:"bytes,1,rep,name=quotes,proto3" json:"quotes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotesResponse) Reset()         { *m = QuotesResponse{} }
func (m *QuotesResponse) String() string { return proto.CompactTextString(m) }
func (*QuotesResponse) ProtoMessage()    {}
func (*QuotesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12e4b46109c2c385, []int{1}
}

func (m *QuotesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotesResponse.Unmarshal(m, b)
}
func (m *QuotesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotesResponse.Marshal(b, m, deterministic)
}
func (m *QuotesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotesResponse.Merge(m, src)
}
func (m *QuotesResponse) XXX_Size() int {
	return xxx_messageInfo_QuotesResponse.Size(m)
}
func (m *QuotesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuotesResponse proto.InternalMessageInfo

func (m *QuotesResponse) GetQuotes() []*Quote {
	if m != nil {
		return m.Quotes
	}
	return nil
}

type Quote struct {
	Vendor               string   `protobuf:"bytes,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Price                float64  `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Quote) Reset()         { *m = Quote{} }
func (m *Quote) String() string { return proto.CompactTextString(m) }
func (*Quote) ProtoMessage()    {}
func (*Quote) Descriptor() ([]byte, []int) {
	return fileDescriptor_12e4b46109c2c385, []int{2}
}

func (m *Quote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quote.Unmarshal(m, b)
}
func (m *Quote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quote.Marshal(b, m, deterministic)
}
func (m *Quote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quote.Merge(m, src)
}
func (m *Quote) XXX_Size() int {
	return xxx_messageInfo_Quote.Size(m)
}
func (m *Quote) XXX_DiscardUnknown() {
	xxx_messageInfo_Quote.DiscardUnknown(m)
}

var xxx_messageInfo_Quote proto.InternalMessageInfo

func (m *Quote) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *Quote) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*QuotesRequest)(nil), "proto.QuotesRequest")
	proto.RegisterType((*QuotesResponse)(nil), "proto.QuotesResponse")
	proto.RegisterType((*Quote)(nil), "proto.Quote")
}

func init() {
	proto.RegisterFile("quote.proto", fileDescriptor_12e4b46109c2c385)
}

var fileDescriptor_12e4b46109c2c385 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2c, 0xcd, 0x2f,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x8a, 0x5c, 0xbc, 0x81,
	0x20, 0xd1, 0xe2, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x01, 0x2e, 0xe6, 0x60, 0xef,
	0x50, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0xc9, 0x8c, 0x8b, 0x0f, 0xa6, 0xa4,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x85, 0x8b, 0x0d, 0x6c, 0x54, 0xb1, 0x04, 0xa3, 0x02,
	0xb3, 0x06, 0xb7, 0x11, 0x0f, 0xc4, 0x4c, 0x3d, 0xb0, 0xb2, 0x20, 0xa8, 0x9c, 0x92, 0x29, 0x17,
	0x2b, 0x58, 0x40, 0x48, 0x8c, 0x8b, 0xad, 0x2c, 0x35, 0x2f, 0x25, 0xbf, 0x08, 0x6a, 0x2a, 0x94,
	0x27, 0x24, 0xc2, 0xc5, 0x5a, 0x50, 0x94, 0x99, 0x9c, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x18,
	0x04, 0xe1, 0x18, 0xb9, 0x73, 0xf1, 0x80, 0xb5, 0x05, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a,
	0x99, 0x73, 0x71, 0xb8, 0xa7, 0x96, 0x40, 0x4c, 0x12, 0x41, 0xb6, 0x08, 0xe6, 0x64, 0x29, 0x51,
	0x34, 0x51, 0x88, 0x2b, 0x93, 0xd8, 0xc0, 0xa2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a,
	0x7e, 0x85, 0x66, 0xf7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QuoteServiceClient is the client API for QuoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuoteServiceClient interface {
	GetQuote(ctx context.Context, in *QuotesRequest, opts ...grpc.CallOption) (*QuotesResponse, error)
}

type quoteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuoteServiceClient(cc grpc.ClientConnInterface) QuoteServiceClient {
	return &quoteServiceClient{cc}
}

func (c *quoteServiceClient) GetQuote(ctx context.Context, in *QuotesRequest, opts ...grpc.CallOption) (*QuotesResponse, error) {
	out := new(QuotesResponse)
	err := c.cc.Invoke(ctx, "/proto.QuoteService/GetQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuoteServiceServer is the server API for QuoteService service.
type QuoteServiceServer interface {
	GetQuote(context.Context, *QuotesRequest) (*QuotesResponse, error)
}

// UnimplementedQuoteServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQuoteServiceServer struct {
}

func (*UnimplementedQuoteServiceServer) GetQuote(ctx context.Context, req *QuotesRequest) (*QuotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuote not implemented")
}

func RegisterQuoteServiceServer(s *grpc.Server, srv QuoteServiceServer) {
	s.RegisterService(&_QuoteService_serviceDesc, srv)
}

func _QuoteService_GetQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).GetQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.QuoteService/GetQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).GetQuote(ctx, req.(*QuotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.QuoteService",
	HandlerType: (*QuoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuote",
			Handler:    _QuoteService_GetQuote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quote.proto",
}
