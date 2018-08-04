// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deviceMessage.proto

package protoc_gen_go

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Platform int32

const (
	Platform_PlatformNone    Platform = 0
	Platform_PlatformIos     Platform = 1
	Platform_PlatformAndroid Platform = 2
)

var Platform_name = map[int32]string{
	0: "PlatformNone",
	1: "PlatformIos",
	2: "PlatformAndroid",
}
var Platform_value = map[string]int32{
	"PlatformNone":    0,
	"PlatformIos":     1,
	"PlatformAndroid": 2,
}

func (x Platform) String() string {
	return proto.EnumName(Platform_name, int32(x))
}
func (Platform) EnumDescriptor() ([]byte, []int) { return fileDescriptorDeviceMessage, []int{0} }

type Device struct {
	UserID               string   `protobuf:"bytes,11,opt,name=user_id,json=userId,proto3" json:"userId" db:"user_id,notnull"`
	Platform             Platform `protobuf:"varint,12,opt,name=platform,proto3,enum=swagchat.protobuf.Platform" json:"platform,omitempty" db:"platform,notnull"`
	Token                string   `protobuf:"bytes,13,opt,name=token,proto3" json:"token" db:"token"`
	NotificationDeviceID string   `protobuf:"bytes,14,opt,name=notification_device_id,json=notificationDeviceId,proto3" json:"notificationDeviceId" db:"notification_device_id"`
	Deleted              int64    `protobuf:"varint,15,opt,name=deleted,proto3" json:"deleted,omitempty" db:"deleted,notnull"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptorDeviceMessage, []int{0} }

func (m *Device) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Device) GetPlatform() Platform {
	if m != nil {
		return m.Platform
	}
	return Platform_PlatformNone
}

func (m *Device) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Device) GetNotificationDeviceID() string {
	if m != nil {
		return m.NotificationDeviceID
	}
	return ""
}

func (m *Device) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

type CreateDeviceRequest struct {
	UserID   string   `protobuf:"bytes,11,opt,name=user_id,json=userId,proto3" json:"userId"`
	Platform Platform `protobuf:"varint,12,opt,name=platform,proto3,enum=swagchat.protobuf.Platform" json:"platform,omitempty"`
	Token    string   `protobuf:"bytes,13,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *CreateDeviceRequest) Reset()                    { *m = CreateDeviceRequest{} }
func (m *CreateDeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDeviceRequest) ProtoMessage()               {}
func (*CreateDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceMessage, []int{1} }

func (m *CreateDeviceRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *CreateDeviceRequest) GetPlatform() Platform {
	if m != nil {
		return m.Platform
	}
	return Platform_PlatformNone
}

func (m *CreateDeviceRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetDevicesRequest struct {
	UserID string `protobuf:"bytes,11,opt,name=user_id,json=userId,proto3" json:"userId"`
}

func (m *GetDevicesRequest) Reset()                    { *m = GetDevicesRequest{} }
func (m *GetDevicesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDevicesRequest) ProtoMessage()               {}
func (*GetDevicesRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceMessage, []int{2} }

func (m *GetDevicesRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type DevicesResponse struct {
	Devices []*Device `protobuf:"bytes,11,rep,name=devices" json:"devices,omitempty"`
}

func (m *DevicesResponse) Reset()                    { *m = DevicesResponse{} }
func (m *DevicesResponse) String() string            { return proto.CompactTextString(m) }
func (*DevicesResponse) ProtoMessage()               {}
func (*DevicesResponse) Descriptor() ([]byte, []int) { return fileDescriptorDeviceMessage, []int{3} }

func (m *DevicesResponse) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type UpdateDeviceRequest struct {
	UserID   string   `protobuf:"bytes,11,opt,name=user_id,json=userId,proto3" json:"userId"`
	Platform Platform `protobuf:"varint,12,opt,name=platform,proto3,enum=swagchat.protobuf.Platform" json:"platform,omitempty"`
	Token    string   `protobuf:"bytes,13,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *UpdateDeviceRequest) Reset()                    { *m = UpdateDeviceRequest{} }
func (m *UpdateDeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDeviceRequest) ProtoMessage()               {}
func (*UpdateDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceMessage, []int{4} }

func (m *UpdateDeviceRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdateDeviceRequest) GetPlatform() Platform {
	if m != nil {
		return m.Platform
	}
	return Platform_PlatformNone
}

func (m *UpdateDeviceRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DeleteDeviceRequest struct {
	UserID   string   `protobuf:"bytes,11,opt,name=user_id,json=userId,proto3" json:"userId"`
	Platform Platform `protobuf:"varint,12,opt,name=platform,proto3,enum=swagchat.protobuf.Platform" json:"platform,omitempty"`
}

func (m *DeleteDeviceRequest) Reset()                    { *m = DeleteDeviceRequest{} }
func (m *DeleteDeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDeviceRequest) ProtoMessage()               {}
func (*DeleteDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceMessage, []int{5} }

func (m *DeleteDeviceRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *DeleteDeviceRequest) GetPlatform() Platform {
	if m != nil {
		return m.Platform
	}
	return Platform_PlatformNone
}

func init() {
	proto.RegisterType((*Device)(nil), "swagchat.protobuf.Device")
	proto.RegisterType((*CreateDeviceRequest)(nil), "swagchat.protobuf.CreateDeviceRequest")
	proto.RegisterType((*GetDevicesRequest)(nil), "swagchat.protobuf.GetDevicesRequest")
	proto.RegisterType((*DevicesResponse)(nil), "swagchat.protobuf.DevicesResponse")
	proto.RegisterType((*UpdateDeviceRequest)(nil), "swagchat.protobuf.UpdateDeviceRequest")
	proto.RegisterType((*DeleteDeviceRequest)(nil), "swagchat.protobuf.DeleteDeviceRequest")
	proto.RegisterEnum("swagchat.protobuf.Platform", Platform_name, Platform_value)
}

func init() { proto.RegisterFile("deviceMessage.proto", fileDescriptorDeviceMessage) }

var fileDescriptorDeviceMessage = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x51, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x09, 0xa4, 0x65, 0x52, 0x9a, 0x74, 0x63, 0xc0, 0xa5, 0x07, 0x47, 0x3e, 0x45, 0x55,
	0x7e, 0xa4, 0xf4, 0x80, 0xd4, 0x1b, 0x26, 0x08, 0x22, 0x41, 0x41, 0x96, 0x7a, 0xe1, 0x52, 0x39,
	0xde, 0x89, 0x6b, 0x91, 0xee, 0x06, 0xef, 0x1a, 0x1e, 0x00, 0xf1, 0x06, 0x5c, 0x38, 0xf2, 0x46,
	0x3c, 0x81, 0x0f, 0x3d, 0xf6, 0x98, 0x27, 0x40, 0xde, 0xf5, 0xa6, 0x81, 0xf8, 0x02, 0x17, 0x7a,
	0x9b, 0xf9, 0x66, 0xe6, 0x9b, 0x6f, 0xbe, 0x81, 0x36, 0xc5, 0x4f, 0x49, 0x84, 0x6f, 0x50, 0x88,
	0x30, 0xc6, 0xc1, 0x22, 0xe5, 0x92, 0x93, 0x7d, 0xf1, 0x39, 0x8c, 0xa3, 0x8b, 0x50, 0xea, 0x7c,
	0x9a, 0xcd, 0x9e, 0xd8, 0x31, 0x8f, 0xb9, 0xca, 0x86, 0x45, 0xa4, 0x0b, 0xde, 0xb7, 0x1a, 0xd4,
	0xc7, 0x8a, 0x80, 0xbc, 0x80, 0xed, 0x4c, 0x60, 0x7a, 0x9e, 0x50, 0xa7, 0xd1, 0xb1, 0xba, 0xf7,
	0xfd, 0xde, 0x55, 0xee, 0xd6, 0xcf, 0x04, 0xa6, 0x93, 0xf1, 0x75, 0xee, 0xd6, 0x8b, 0xe2, 0x84,
	0x2e, 0x73, 0xd7, 0xa6, 0xd3, 0x13, 0xaf, 0x6c, 0xed, 0x31, 0x2e, 0x59, 0x36, 0x9f, 0x7b, 0x41,
	0x59, 0x27, 0x01, 0xec, 0x2c, 0xe6, 0xa1, 0x9c, 0xf1, 0xf4, 0xd2, 0xd9, 0xed, 0x58, 0xdd, 0xbd,
	0xd1, 0xe1, 0x60, 0x43, 0xcd, 0xe0, 0x5d, 0xd9, 0xe2, 0x1f, 0x2c, 0x73, 0xf7, 0x61, 0x41, 0x68,
	0x86, 0x6e, 0x18, 0x57, 0x3c, 0xa4, 0x0f, 0xf7, 0x24, 0xff, 0x80, 0xcc, 0x79, 0xa0, 0x84, 0x3d,
	0xbe, 0xce, 0x5d, 0x0d, 0x2c, 0x73, 0x17, 0x8a, 0x61, 0x95, 0x78, 0x81, 0x06, 0xc9, 0x57, 0x0b,
	0x1e, 0x31, 0x2e, 0x93, 0x59, 0x12, 0x85, 0x32, 0xe1, 0xec, 0x5c, 0x5b, 0x54, 0x5c, 0xb6, 0xa7,
	0x08, 0xde, 0x5e, 0xe5, 0xae, 0x7d, 0xba, 0xd6, 0xa1, 0x2d, 0x50, 0x77, 0xda, 0x6c, 0x13, 0x2f,
	0xae, 0x3e, 0x2c, 0xf6, 0x54, 0xb3, 0x7a, 0x41, 0xe5, 0x10, 0x19, 0xc1, 0x36, 0xc5, 0x39, 0x4a,
	0xa4, 0x4e, 0xb3, 0x63, 0x75, 0x6b, 0xbe, 0x63, 0xdc, 0x2b, 0xe1, 0x9b, 0x5b, 0x4d, 0xe3, 0xc9,
	0xdd, 0x9f, 0x3f, 0xdc, 0x2d, 0xef, 0xbb, 0x05, 0xed, 0xe7, 0x29, 0x86, 0x12, 0x35, 0x59, 0x80,
	0x1f, 0x33, 0x14, 0x92, 0xf4, 0xff, 0xfc, 0x91, 0x5d, 0xf5, 0xa3, 0xd5, 0x2f, 0x9e, 0xfe, 0xd5,
	0x2f, 0xd6, 0x0c, 0xb7, 0x7f, 0x33, 0xbc, 0xf4, 0xb5, 0xd4, 0xf6, 0x0a, 0xf6, 0x5f, 0xa2, 0xd4,
	0xba, 0xc4, 0xbf, 0x09, 0x2b, 0x99, 0x5e, 0x43, 0x73, 0x45, 0x23, 0x16, 0x9c, 0x09, 0x24, 0xc7,
	0x85, 0x65, 0x0a, 0x72, 0x1a, 0x9d, 0x5a, 0xb7, 0x31, 0x3a, 0xa8, 0x10, 0x5c, 0x7a, 0x62, 0x3a,
	0xd7, 0x3c, 0x3b, 0x5b, 0xd0, 0x5b, 0xe9, 0xd9, 0x17, 0x0b, 0xda, 0x63, 0xf5, 0xe1, 0xff, 0xa2,
	0x4d, 0xab, 0x38, 0xf2, 0x61, 0xc7, 0xd4, 0x48, 0x0b, 0x76, 0x4d, 0x7c, 0xca, 0x19, 0xb6, 0xb6,
	0x48, 0x13, 0x1a, 0x06, 0x99, 0x70, 0xd1, 0xb2, 0x48, 0x1b, 0x9a, 0x06, 0x78, 0xc6, 0x68, 0xca,
	0x13, 0xda, 0xba, 0xe3, 0xf7, 0xde, 0x1f, 0xc5, 0x89, 0xbc, 0xc8, 0xa6, 0x83, 0x88, 0x5f, 0x0e,
	0xcd, 0xf2, 0xa1, 0x59, 0xae, 0x83, 0xa8, 0x1f, 0x23, 0xeb, 0xc7, 0x7c, 0x5a, 0x57, 0xe9, 0xf1,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xec, 0xa4, 0xf2, 0xa5, 0x04, 0x00, 0x00,
}