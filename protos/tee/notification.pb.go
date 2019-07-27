// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification.proto

package tee

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthStatus int32

const (
	AuthStatus_UnAuthorized AuthStatus = 0
	AuthStatus_Authorized   AuthStatus = 1
	AuthStatus_Refused      AuthStatus = 2
)

var AuthStatus_name = map[int32]string{
	0: "UnAuthorized",
	1: "Authorized",
	2: "Refused",
}
var AuthStatus_value = map[string]int32{
	"UnAuthorized": 0,
	"Authorized":   1,
	"Refused":      2,
}

func (x AuthStatus) String() string {
	return proto.EnumName(AuthStatus_name, int32(x))
}
func (AuthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_notification_9fd4f870b637ffaf, []int{0}
}

type Notification struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *SharedData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Requester            string      `protobuf:"bytes,3,opt,name=requester,proto3" json:"requester,omitempty"`
	RequestSeconds       int64       `protobuf:"varint,4,opt,name=request_seconds,json=requestSeconds,proto3" json:"request_seconds,omitempty"`
	Status               AuthStatus  `protobuf:"varint,5,opt,name=status,proto3,enum=tee.AuthStatus" json:"status,omitempty"`
	AuthSeconds          int64       `protobuf:"varint,6,opt,name=auth_seconds,json=authSeconds,proto3" json:"auth_seconds,omitempty"`
	RefusedReason        string      `protobuf:"bytes,7,opt,name=refused_reason,json=refusedReason,proto3" json:"refused_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_notification_9fd4f870b637ffaf, []int{0}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (dst *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(dst, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Notification) GetData() *SharedData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Notification) GetRequester() string {
	if m != nil {
		return m.Requester
	}
	return ""
}

func (m *Notification) GetRequestSeconds() int64 {
	if m != nil {
		return m.RequestSeconds
	}
	return 0
}

func (m *Notification) GetStatus() AuthStatus {
	if m != nil {
		return m.Status
	}
	return AuthStatus_UnAuthorized
}

func (m *Notification) GetAuthSeconds() int64 {
	if m != nil {
		return m.AuthSeconds
	}
	return 0
}

func (m *Notification) GetRefusedReason() string {
	if m != nil {
		return m.RefusedReason
	}
	return ""
}

func init() {
	proto.RegisterType((*Notification)(nil), "tee.Notification")
	proto.RegisterEnum("tee.AuthStatus", AuthStatus_name, AuthStatus_value)
}

func init() { proto.RegisterFile("notification.proto", fileDescriptor_notification_9fd4f870b637ffaf) }

var fileDescriptor_notification_9fd4f870b637ffaf = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0x5a, 0x5a, 0xf5, 0x12, 0xd2, 0xe8, 0x26, 0x0b, 0x31, 0x04, 0x10, 0x6a, 0xc4,
	0x90, 0x01, 0x46, 0x26, 0x24, 0x66, 0x06, 0x47, 0xcc, 0x95, 0xa9, 0xaf, 0x8a, 0x97, 0x18, 0xec,
	0xcb, 0xc2, 0x7f, 0x47, 0x42, 0x71, 0x0c, 0x65, 0x7c, 0xdf, 0x7b, 0xfe, 0x6c, 0x19, 0x70, 0x70,
	0x6c, 0x8f, 0xf6, 0xa0, 0xd9, 0xba, 0xa1, 0xfd, 0xf0, 0x8e, 0x1d, 0x2e, 0x98, 0xe8, 0xb2, 0x0a,
	0xbd, 0xf6, 0x64, 0x8c, 0x66, 0x3d, 0xe3, 0x9b, 0x6f, 0x01, 0xc5, 0xeb, 0xbf, 0x35, 0x96, 0x90,
	0x59, 0x23, 0x45, 0x2d, 0x9a, 0x8d, 0xca, 0xac, 0xc1, 0x5b, 0x58, 0x4e, 0x73, 0x99, 0xd5, 0xa2,
	0xc9, 0x1f, 0xb6, 0x2d, 0x13, 0xb5, 0x5d, 0xb4, 0xbc, 0x68, 0xd6, 0x2a, 0x96, 0x78, 0x05, 0x1b,
	0x4f, 0x9f, 0x23, 0x05, 0x26, 0x2f, 0x17, 0xf1, 0xec, 0x09, 0xe0, 0x0e, 0xb6, 0x29, 0xec, 0x03,
	0x1d, 0xdc, 0x60, 0x82, 0x5c, 0xd6, 0xa2, 0x59, 0xa8, 0x32, 0xe1, 0x6e, 0xa6, 0xb8, 0x83, 0x55,
	0x60, 0xcd, 0x63, 0x90, 0xe7, 0xb5, 0x68, 0xca, 0x74, 0xdb, 0xf3, 0xc8, 0x7d, 0x17, 0xb1, 0x4a,
	0x35, 0x5e, 0x43, 0xa1, 0x47, 0xee, 0xff, 0x74, 0xab, 0xa8, 0xcb, 0x27, 0xf6, 0xeb, 0xba, 0x83,
	0xd2, 0xd3, 0x71, 0x0c, 0x64, 0xf6, 0x9e, 0x74, 0x70, 0x83, 0x5c, 0xc7, 0x77, 0x5d, 0x24, 0xaa,
	0x22, 0xbc, 0x7f, 0x02, 0x38, 0xf9, 0xb1, 0x82, 0xe2, 0x6d, 0x98, 0xb2, 0xf3, 0xf6, 0x8b, 0x4c,
	0x75, 0x86, 0xe5, 0xdc, 0xa7, 0x2c, 0x30, 0x87, 0xb5, 0x9a, 0x05, 0x55, 0xf6, 0xbe, 0x8a, 0x7f,
	0xf8, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x7d, 0xcc, 0x15, 0x70, 0x01, 0x00, 0x00,
}
