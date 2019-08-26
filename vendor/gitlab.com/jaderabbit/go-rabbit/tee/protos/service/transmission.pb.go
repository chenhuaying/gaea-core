// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transmission.proto

package service

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

type UploadFileRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileRequest) Reset()         { *m = UploadFileRequest{} }
func (m *UploadFileRequest) String() string { return proto.CompactTextString(m) }
func (*UploadFileRequest) ProtoMessage()    {}
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e45df04459fa574, []int{0}
}

func (m *UploadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileRequest.Unmarshal(m, b)
}
func (m *UploadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileRequest.Marshal(b, m, deterministic)
}
func (m *UploadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileRequest.Merge(m, src)
}
func (m *UploadFileRequest) XXX_Size() int {
	return xxx_messageInfo_UploadFileRequest.Size(m)
}
func (m *UploadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileRequest proto.InternalMessageInfo

func (m *UploadFileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UploadFileRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type UploadFileResponse struct {
	FileId               string   `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileResponse) Reset()         { *m = UploadFileResponse{} }
func (m *UploadFileResponse) String() string { return proto.CompactTextString(m) }
func (*UploadFileResponse) ProtoMessage()    {}
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e45df04459fa574, []int{1}
}

func (m *UploadFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileResponse.Unmarshal(m, b)
}
func (m *UploadFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileResponse.Marshal(b, m, deterministic)
}
func (m *UploadFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileResponse.Merge(m, src)
}
func (m *UploadFileResponse) XXX_Size() int {
	return xxx_messageInfo_UploadFileResponse.Size(m)
}
func (m *UploadFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileResponse proto.InternalMessageInfo

func (m *UploadFileResponse) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

type DownloadFileRequest struct {
	FileId               string   `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadFileRequest) Reset()         { *m = DownloadFileRequest{} }
func (m *DownloadFileRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadFileRequest) ProtoMessage()    {}
func (*DownloadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e45df04459fa574, []int{2}
}

func (m *DownloadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadFileRequest.Unmarshal(m, b)
}
func (m *DownloadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadFileRequest.Marshal(b, m, deterministic)
}
func (m *DownloadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadFileRequest.Merge(m, src)
}
func (m *DownloadFileRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadFileRequest.Size(m)
}
func (m *DownloadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadFileRequest proto.InternalMessageInfo

func (m *DownloadFileRequest) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

func (m *DownloadFileRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type DownloadFileResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadFileResponse) Reset()         { *m = DownloadFileResponse{} }
func (m *DownloadFileResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadFileResponse) ProtoMessage()    {}
func (*DownloadFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e45df04459fa574, []int{3}
}

func (m *DownloadFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadFileResponse.Unmarshal(m, b)
}
func (m *DownloadFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadFileResponse.Marshal(b, m, deterministic)
}
func (m *DownloadFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadFileResponse.Merge(m, src)
}
func (m *DownloadFileResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadFileResponse.Size(m)
}
func (m *DownloadFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadFileResponse proto.InternalMessageInfo

func (m *DownloadFileResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadFileRequest)(nil), "service.UploadFileRequest")
	proto.RegisterType((*UploadFileResponse)(nil), "service.UploadFileResponse")
	proto.RegisterType((*DownloadFileRequest)(nil), "service.DownloadFileRequest")
	proto.RegisterType((*DownloadFileResponse)(nil), "service.DownloadFileResponse")
}

func init() { proto.RegisterFile("transmission.proto", fileDescriptor_5e45df04459fa574) }

var fileDescriptor_5e45df04459fa574 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x29, 0x4a, 0xcc,
	0x2b, 0xce, 0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x55, 0x72, 0xe0, 0x12, 0x0c, 0x2d, 0xc8, 0xc9, 0x4f,
	0x4c, 0x71, 0xcb, 0xcc, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62,
	0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x85, 0xc4, 0xb9,
	0xd8, 0x4b, 0x8b, 0x53, 0x8b, 0xe2, 0x33, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xd8,
	0x40, 0x5c, 0xcf, 0x14, 0x25, 0x5d, 0x2e, 0x21, 0x64, 0x13, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53,
	0x41, 0xca, 0xd3, 0x32, 0x73, 0x52, 0x41, 0xca, 0x19, 0x21, 0xca, 0x41, 0x5c, 0xcf, 0x14, 0x25,
	0x77, 0x2e, 0x61, 0x97, 0xfc, 0xf2, 0x3c, 0x74, 0x2b, 0x71, 0xa9, 0xc7, 0x6d, 0xaf, 0x16, 0x97,
	0x08, 0xaa, 0x41, 0x50, 0x9b, 0xb1, 0x38, 0xde, 0x68, 0x11, 0x23, 0x17, 0x4f, 0x08, 0x52, 0x28,
	0x08, 0xb9, 0x72, 0x71, 0x21, 0x1c, 0x2d, 0x24, 0xa5, 0x07, 0x0d, 0x0e, 0x3d, 0x8c, 0xb0, 0x90,
	0x92, 0xc6, 0x2a, 0x07, 0xb5, 0xcb, 0x9b, 0x8b, 0x07, 0xd9, 0x0d, 0x42, 0x32, 0x70, 0xc5, 0x58,
	0xfc, 0x28, 0x25, 0x8b, 0x43, 0x16, 0x62, 0x58, 0x12, 0x1b, 0x38, 0x6a, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xbc, 0x6a, 0xef, 0xf6, 0xb0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransmissionClient is the client API for Transmission service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransmissionClient interface {
	UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error)
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error)
}

type transmissionClient struct {
	cc *grpc.ClientConn
}

func NewTransmissionClient(cc *grpc.ClientConn) TransmissionClient {
	return &transmissionClient{cc}
}

func (c *transmissionClient) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error) {
	out := new(UploadFileResponse)
	err := c.cc.Invoke(ctx, "/service.Transmission/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transmissionClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error) {
	out := new(DownloadFileResponse)
	err := c.cc.Invoke(ctx, "/service.Transmission/DownloadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransmissionServer is the server API for Transmission service.
type TransmissionServer interface {
	UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error)
	DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error)
}

// UnimplementedTransmissionServer can be embedded to have forward compatible implementations.
type UnimplementedTransmissionServer struct {
}

func (*UnimplementedTransmissionServer) UploadFile(ctx context.Context, req *UploadFileRequest) (*UploadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedTransmissionServer) DownloadFile(ctx context.Context, req *DownloadFileRequest) (*DownloadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}

func RegisterTransmissionServer(s *grpc.Server, srv TransmissionServer) {
	s.RegisterService(&_Transmission_serviceDesc, srv)
}

func _Transmission_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransmissionServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Transmission/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransmissionServer).UploadFile(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transmission_DownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransmissionServer).DownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Transmission/DownloadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransmissionServer).DownloadFile(ctx, req.(*DownloadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transmission_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Transmission",
	HandlerType: (*TransmissionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadFile",
			Handler:    _Transmission_UploadFile_Handler,
		},
		{
			MethodName: "DownloadFile",
			Handler:    _Transmission_DownloadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transmission.proto",
}
