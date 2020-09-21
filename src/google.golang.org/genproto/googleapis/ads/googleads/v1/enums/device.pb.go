// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/device.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enumerates Google Ads devices available for targeting.
type DeviceEnum_Device int32

const (
	// Not specified.
	DeviceEnum_UNSPECIFIED DeviceEnum_Device = 0
	// The value is unknown in this version.
	DeviceEnum_UNKNOWN DeviceEnum_Device = 1
	// Mobile devices with full browsers.
	DeviceEnum_MOBILE DeviceEnum_Device = 2
	// Tablets with full browsers.
	DeviceEnum_TABLET DeviceEnum_Device = 3
	// Computers.
	DeviceEnum_DESKTOP DeviceEnum_Device = 4
	// Smart TVs and game consoles.
	DeviceEnum_CONNECTED_TV DeviceEnum_Device = 6
	// Other device types.
	DeviceEnum_OTHER DeviceEnum_Device = 5
)

var DeviceEnum_Device_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MOBILE",
	3: "TABLET",
	4: "DESKTOP",
	6: "CONNECTED_TV",
	5: "OTHER",
}

var DeviceEnum_Device_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"MOBILE":       2,
	"TABLET":       3,
	"DESKTOP":      4,
	"CONNECTED_TV": 6,
	"OTHER":        5,
}

func (x DeviceEnum_Device) String() string {
	return proto.EnumName(DeviceEnum_Device_name, int32(x))
}

func (DeviceEnum_Device) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f3653a34a029c3f, []int{0, 0}
}

// Container for enumeration of Google Ads devices available for targeting.
type DeviceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceEnum) Reset()         { *m = DeviceEnum{} }
func (m *DeviceEnum) String() string { return proto.CompactTextString(m) }
func (*DeviceEnum) ProtoMessage()    {}
func (*DeviceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f3653a34a029c3f, []int{0}
}

func (m *DeviceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceEnum.Unmarshal(m, b)
}
func (m *DeviceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceEnum.Marshal(b, m, deterministic)
}
func (m *DeviceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceEnum.Merge(m, src)
}
func (m *DeviceEnum) XXX_Size() int {
	return xxx_messageInfo_DeviceEnum.Size(m)
}
func (m *DeviceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.DeviceEnum_Device", DeviceEnum_Device_name, DeviceEnum_Device_value)
	proto.RegisterType((*DeviceEnum)(nil), "google.ads.googleads.v1.enums.DeviceEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/device.proto", fileDescriptor_2f3653a34a029c3f)
}

var fileDescriptor_2f3653a34a029c3f = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4f, 0xfa, 0x30,
	0x00, 0xc5, 0xff, 0x8c, 0x3f, 0x33, 0x16, 0x13, 0x9b, 0x1d, 0x8d, 0x1c, 0xe0, 0xe8, 0xa1, 0xcb,
	0xe2, 0xad, 0x9e, 0x36, 0x56, 0x91, 0x80, 0xdb, 0x22, 0x65, 0x26, 0x66, 0x89, 0x99, 0x74, 0xa9,
	0x4b, 0xa0, 0x25, 0x14, 0xf6, 0x81, 0x3c, 0xfa, 0x51, 0xfc, 0x14, 0x9e, 0xfd, 0x14, 0xa6, 0xad,
	0xec, 0xa6, 0x97, 0xe5, 0xad, 0xef, 0xd7, 0x97, 0xd7, 0x07, 0xae, 0xb8, 0x94, 0x7c, 0x5d, 0xf9,
	0x25, 0x53, 0xbe, 0x95, 0x5a, 0x35, 0x81, 0x5f, 0x89, 0xc3, 0x46, 0xf9, 0xac, 0x6a, 0xea, 0x55,
	0x85, 0xb6, 0x3b, 0xb9, 0x97, 0xde, 0xc0, 0x02, 0xa8, 0x64, 0x0a, 0xb5, 0x2c, 0x6a, 0x02, 0x64,
	0xd8, 0x8b, 0xcb, 0x63, 0xd4, 0xb6, 0xf6, 0x4b, 0x21, 0xe4, 0xbe, 0xdc, 0xd7, 0x52, 0x28, 0x7b,
	0x79, 0xd4, 0x00, 0x10, 0x9b, 0x30, 0x22, 0x0e, 0x9b, 0xd1, 0x2b, 0x70, 0xed, 0x9f, 0x77, 0x0e,
	0xfa, 0xcb, 0x64, 0x91, 0x91, 0xf1, 0xf4, 0x76, 0x4a, 0x62, 0xf8, 0xcf, 0xeb, 0x83, 0x93, 0x65,
	0x32, 0x4b, 0xd2, 0xc7, 0x04, 0x76, 0x3c, 0x00, 0xdc, 0xfb, 0x34, 0x9a, 0xce, 0x09, 0x74, 0xb4,
	0xa6, 0x61, 0x34, 0x27, 0x14, 0x76, 0x35, 0x14, 0x93, 0xc5, 0x8c, 0xa6, 0x19, 0xfc, 0xef, 0x41,
	0x70, 0x36, 0x4e, 0x93, 0x84, 0x8c, 0x29, 0x89, 0x9f, 0x69, 0x0e, 0x5d, 0xef, 0x14, 0xf4, 0x52,
	0x7a, 0x47, 0x1e, 0x60, 0x2f, 0xfa, 0xec, 0x80, 0xe1, 0x4a, 0x6e, 0xd0, 0x9f, 0xdd, 0xa3, 0xbe,
	0x6d, 0x93, 0xe9, 0xaa, 0x59, 0xe7, 0x29, 0xfa, 0xa1, 0xb9, 0x5c, 0x97, 0x82, 0x23, 0xb9, 0xe3,
	0x3e, 0xaf, 0x84, 0x79, 0xc8, 0x71, 0xa5, 0x6d, 0xad, 0x7e, 0x19, 0xed, 0xc6, 0x7c, 0xdf, 0x9c,
	0xee, 0x24, 0x0c, 0xdf, 0x9d, 0xc1, 0xc4, 0x46, 0x85, 0x4c, 0x21, 0x2b, 0xb5, 0xca, 0x03, 0xa4,
	0x77, 0x50, 0x1f, 0x47, 0xbf, 0x08, 0x99, 0x2a, 0x5a, 0xbf, 0xc8, 0x83, 0xc2, 0xf8, 0x5f, 0xce,
	0xd0, 0x1e, 0x62, 0x1c, 0x32, 0x85, 0x71, 0x4b, 0x60, 0x9c, 0x07, 0x18, 0x1b, 0xe6, 0xc5, 0x35,
	0xc5, 0xae, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x78, 0xaa, 0x06, 0xcc, 0x01, 0x00, 0x00,
}
