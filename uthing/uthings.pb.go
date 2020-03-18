// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uthing/uthings.proto

package uthing

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

type SmartCastCommand int32

const (
	SmartCastCommand_UNKNOWN                     SmartCastCommand = 0
	SmartCastCommand_PING                        SmartCastCommand = 1
	SmartCastCommand_GETALLSWNIS                 SmartCastCommand = 2
	SmartCastCommand_GETSWNI                     SmartCastCommand = 3
	SmartCastCommand_IDENTIFY                    SmartCastCommand = 4
	SmartCastCommand_GET_MY_NETWORK              SmartCastCommand = 5
	SmartCastCommand_LIST_ALL_NETWORKS           SmartCastCommand = 6
	SmartCastCommand_NETWORK_RESET               SmartCastCommand = 7
	SmartCastCommand_GET_MODE                    SmartCastCommand = 8
	SmartCastCommand_ACTUATOR                    SmartCastCommand = 9
	SmartCastCommand_GET_DEVICE_TABLE            SmartCastCommand = 10
	SmartCastCommand_REFRESH_DEVICE_TABLE        SmartCastCommand = 11
	SmartCastCommand_SET_DHB_CONFIG              SmartCastCommand = 12
	SmartCastCommand_DATA_HEARTBEAT              SmartCastCommand = 13
	SmartCastCommand_GET_POWER                   SmartCastCommand = 14
	SmartCastCommand_REBOOT                      SmartCastCommand = 15
	SmartCastCommand_SET_CHANNEL                 SmartCastCommand = 16
	SmartCastCommand_FIRMWARE_UPDATE             SmartCastCommand = 17
	SmartCastCommand_SET_PAN_ID                  SmartCastCommand = 18
	SmartCastCommand_SET_LED_PATTERN             SmartCastCommand = 19
	SmartCastCommand_AUTO_ADDRESS                SmartCastCommand = 20
	SmartCastCommand_GET_DHB_CONFIG              SmartCastCommand = 21
	SmartCastCommand_GET_NODE_COUNT              SmartCastCommand = 22
	SmartCastCommand_OP_MODE_EVENT               SmartCastCommand = 23
	SmartCastCommand_SET_DEV_CONFIG              SmartCastCommand = 24
	SmartCastCommand_GET_DEV_CONFIG              SmartCastCommand = 25
	SmartCastCommand_APPLY_SCENE                 SmartCastCommand = 26
	SmartCastCommand_APPLY_DYNAMIC_PROFILE       SmartCastCommand = 27
	SmartCastCommand_GET_LIGHTING_STATUS         SmartCastCommand = 28
	SmartCastCommand_SWG_TEST                    SmartCastCommand = 29
	SmartCastCommand_SET_DYNAMIC_PROFILE_ID      SmartCastCommand = 30
	SmartCastCommand_TOGGLE                      SmartCastCommand = 31
	SmartCastCommand_PAUSE_PROFILES              SmartCastCommand = 32
	SmartCastCommand_GET_DYNAMIC_LIGHTING_STATUS SmartCastCommand = 33
	//offset spacing
	SmartCastCommand_DATA_HEARTBEAT_0 SmartCastCommand = 1000
	SmartCastCommand_DATA_HEARTBEAT_1 SmartCastCommand = 1001
	SmartCastCommand_DATA_HEARTBEAT_2 SmartCastCommand = 1002
	SmartCastCommand_DATA_HEARTBEAT_3 SmartCastCommand = 1003
	SmartCastCommand_DATA_HEARTBEAT_4 SmartCastCommand = 1004
	SmartCastCommand_DATA_HEARTBEAT_5 SmartCastCommand = 1005
	SmartCastCommand_DATA_HEARTBEAT_6 SmartCastCommand = 1006
	SmartCastCommand_DATA_HEARTBEAT_7 SmartCastCommand = 1007
)

var SmartCastCommand_name = map[int32]string{
	0:    "UNKNOWN",
	1:    "PING",
	2:    "GETALLSWNIS",
	3:    "GETSWNI",
	4:    "IDENTIFY",
	5:    "GET_MY_NETWORK",
	6:    "LIST_ALL_NETWORKS",
	7:    "NETWORK_RESET",
	8:    "GET_MODE",
	9:    "ACTUATOR",
	10:   "GET_DEVICE_TABLE",
	11:   "REFRESH_DEVICE_TABLE",
	12:   "SET_DHB_CONFIG",
	13:   "DATA_HEARTBEAT",
	14:   "GET_POWER",
	15:   "REBOOT",
	16:   "SET_CHANNEL",
	17:   "FIRMWARE_UPDATE",
	18:   "SET_PAN_ID",
	19:   "SET_LED_PATTERN",
	20:   "AUTO_ADDRESS",
	21:   "GET_DHB_CONFIG",
	22:   "GET_NODE_COUNT",
	23:   "OP_MODE_EVENT",
	24:   "SET_DEV_CONFIG",
	25:   "GET_DEV_CONFIG",
	26:   "APPLY_SCENE",
	27:   "APPLY_DYNAMIC_PROFILE",
	28:   "GET_LIGHTING_STATUS",
	29:   "SWG_TEST",
	30:   "SET_DYNAMIC_PROFILE_ID",
	31:   "TOGGLE",
	32:   "PAUSE_PROFILES",
	33:   "GET_DYNAMIC_LIGHTING_STATUS",
	1000: "DATA_HEARTBEAT_0",
	1001: "DATA_HEARTBEAT_1",
	1002: "DATA_HEARTBEAT_2",
	1003: "DATA_HEARTBEAT_3",
	1004: "DATA_HEARTBEAT_4",
	1005: "DATA_HEARTBEAT_5",
	1006: "DATA_HEARTBEAT_6",
	1007: "DATA_HEARTBEAT_7",
}

var SmartCastCommand_value = map[string]int32{
	"UNKNOWN":                     0,
	"PING":                        1,
	"GETALLSWNIS":                 2,
	"GETSWNI":                     3,
	"IDENTIFY":                    4,
	"GET_MY_NETWORK":              5,
	"LIST_ALL_NETWORKS":           6,
	"NETWORK_RESET":               7,
	"GET_MODE":                    8,
	"ACTUATOR":                    9,
	"GET_DEVICE_TABLE":            10,
	"REFRESH_DEVICE_TABLE":        11,
	"SET_DHB_CONFIG":              12,
	"DATA_HEARTBEAT":              13,
	"GET_POWER":                   14,
	"REBOOT":                      15,
	"SET_CHANNEL":                 16,
	"FIRMWARE_UPDATE":             17,
	"SET_PAN_ID":                  18,
	"SET_LED_PATTERN":             19,
	"AUTO_ADDRESS":                20,
	"GET_DHB_CONFIG":              21,
	"GET_NODE_COUNT":              22,
	"OP_MODE_EVENT":               23,
	"SET_DEV_CONFIG":              24,
	"GET_DEV_CONFIG":              25,
	"APPLY_SCENE":                 26,
	"APPLY_DYNAMIC_PROFILE":       27,
	"GET_LIGHTING_STATUS":         28,
	"SWG_TEST":                    29,
	"SET_DYNAMIC_PROFILE_ID":      30,
	"TOGGLE":                      31,
	"PAUSE_PROFILES":              32,
	"GET_DYNAMIC_LIGHTING_STATUS": 33,
	"DATA_HEARTBEAT_0":            1000,
	"DATA_HEARTBEAT_1":            1001,
	"DATA_HEARTBEAT_2":            1002,
	"DATA_HEARTBEAT_3":            1003,
	"DATA_HEARTBEAT_4":            1004,
	"DATA_HEARTBEAT_5":            1005,
	"DATA_HEARTBEAT_6":            1006,
	"DATA_HEARTBEAT_7":            1007,
}

func (x SmartCastCommand) String() string {
	return proto.EnumName(SmartCastCommand_name, int32(x))
}

func (SmartCastCommand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bc8651075e02dc82, []int{0}
}

type SmartCastRequest struct {
	NetworkID            []string         `protobuf:"bytes,1,rep,name=NetworkID,proto3" json:"NetworkID,omitempty"`
	SCCommand            SmartCastCommand `protobuf:"varint,2,opt,name=SCCommand,proto3,enum=uthing.SmartCastCommand" json:"SCCommand,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SmartCastRequest) Reset()         { *m = SmartCastRequest{} }
func (m *SmartCastRequest) String() string { return proto.CompactTextString(m) }
func (*SmartCastRequest) ProtoMessage()    {}
func (*SmartCastRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc8651075e02dc82, []int{0}
}

func (m *SmartCastRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmartCastRequest.Unmarshal(m, b)
}
func (m *SmartCastRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmartCastRequest.Marshal(b, m, deterministic)
}
func (m *SmartCastRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartCastRequest.Merge(m, src)
}
func (m *SmartCastRequest) XXX_Size() int {
	return xxx_messageInfo_SmartCastRequest.Size(m)
}
func (m *SmartCastRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartCastRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SmartCastRequest proto.InternalMessageInfo

func (m *SmartCastRequest) GetNetworkID() []string {
	if m != nil {
		return m.NetworkID
	}
	return nil
}

func (m *SmartCastRequest) GetSCCommand() SmartCastCommand {
	if m != nil {
		return m.SCCommand
	}
	return SmartCastCommand_UNKNOWN
}

type SmartCastResponse struct {
	NetworkID            string           `protobuf:"bytes,1,opt,name=NetworkID,proto3" json:"NetworkID,omitempty"`
	MID                  int32            `protobuf:"varint,2,opt,name=MID,proto3" json:"MID,omitempty"`
	SmartcastCommand     SmartCastCommand `protobuf:"varint,3,opt,name=smartcastCommand,proto3,enum=uthing.SmartCastCommand" json:"smartcastCommand,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SmartCastResponse) Reset()         { *m = SmartCastResponse{} }
func (m *SmartCastResponse) String() string { return proto.CompactTextString(m) }
func (*SmartCastResponse) ProtoMessage()    {}
func (*SmartCastResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc8651075e02dc82, []int{1}
}

func (m *SmartCastResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmartCastResponse.Unmarshal(m, b)
}
func (m *SmartCastResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmartCastResponse.Marshal(b, m, deterministic)
}
func (m *SmartCastResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartCastResponse.Merge(m, src)
}
func (m *SmartCastResponse) XXX_Size() int {
	return xxx_messageInfo_SmartCastResponse.Size(m)
}
func (m *SmartCastResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartCastResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SmartCastResponse proto.InternalMessageInfo

func (m *SmartCastResponse) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *SmartCastResponse) GetMID() int32 {
	if m != nil {
		return m.MID
	}
	return 0
}

func (m *SmartCastResponse) GetSmartcastCommand() SmartCastCommand {
	if m != nil {
		return m.SmartcastCommand
	}
	return SmartCastCommand_UNKNOWN
}

type CreeResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreeResponse) Reset()         { *m = CreeResponse{} }
func (m *CreeResponse) String() string { return proto.CompactTextString(m) }
func (*CreeResponse) ProtoMessage()    {}
func (*CreeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc8651075e02dc82, []int{2}
}

func (m *CreeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreeResponse.Unmarshal(m, b)
}
func (m *CreeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreeResponse.Marshal(b, m, deterministic)
}
func (m *CreeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreeResponse.Merge(m, src)
}
func (m *CreeResponse) XXX_Size() int {
	return xxx_messageInfo_CreeResponse.Size(m)
}
func (m *CreeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreeResponse proto.InternalMessageInfo

func (m *CreeResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterEnum("uthing.SmartCastCommand", SmartCastCommand_name, SmartCastCommand_value)
	proto.RegisterType((*SmartCastRequest)(nil), "uthing.SmartCastRequest")
	proto.RegisterType((*SmartCastResponse)(nil), "uthing.SmartCastResponse")
	proto.RegisterType((*CreeResponse)(nil), "uthing.CreeResponse")
}

func init() {
	proto.RegisterFile("uthing/uthings.proto", fileDescriptor_bc8651075e02dc82)
}

var fileDescriptor_bc8651075e02dc82 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x8d, 0xa2, 0x58, 0x97, 0xb1, 0x2c, 0x8f, 0xd6, 0x92, 0x23, 0x3b, 0x69, 0xe3, 0xea, 0xa1,
	0x30, 0xfa, 0xe0, 0xa6, 0x4e, 0x9b, 0x3e, 0xaf, 0xc8, 0x11, 0x45, 0x84, 0x5a, 0x12, 0xbb, 0x4b,
	0x0b, 0xca, 0xcb, 0x42, 0x4d, 0x89, 0x3a, 0x68, 0x4d, 0xa5, 0x22, 0x8d, 0x7c, 0x43, 0x7f, 0xad,
	0x5f, 0xd2, 0x7b, 0x7f, 0xa1, 0x58, 0xdd, 0x95, 0x10, 0xe8, 0x13, 0x39, 0x67, 0x66, 0xce, 0x39,
	0x3b, 0x8b, 0x1d, 0x68, 0xdf, 0xe7, 0xb7, 0x6f, 0xd3, 0x1f, 0xbe, 0x5c, 0x7e, 0xb2, 0xab, 0x77,
	0xf3, 0x59, 0x3e, 0x63, 0x95, 0x65, 0xd8, 0xbb, 0x05, 0x54, 0x77, 0xd3, 0x79, 0xee, 0x4c, 0xb3,
	0x5c, 0x26, 0x3f, 0xdf, 0x27, 0x59, 0xce, 0x9e, 0x42, 0x5d, 0x24, 0xf9, 0xfb, 0xd9, 0xfc, 0x47,
	0xdf, 0xed, 0x96, 0x2e, 0xca, 0x97, 0x75, 0xb9, 0x05, 0xd8, 0x4b, 0xa8, 0x2b, 0xc7, 0x99, 0xdd,
	0xdd, 0x4d, 0xd3, 0xef, 0xbb, 0x0f, 0x2f, 0x4a, 0x97, 0xcd, 0xeb, 0xee, 0xd5, 0x92, 0xed, 0x6a,
	0x43, 0xb5, 0xca, 0xcb, 0x6d, 0x69, 0xef, 0x97, 0x12, 0xb4, 0x76, 0xa4, 0xb2, 0x77, 0xb3, 0x34,
	0x4b, 0x3e, 0xd4, 0x2a, 0xed, 0x6b, 0x21, 0x94, 0x47, 0xbe, 0xbb, 0x50, 0x39, 0x90, 0xf6, 0x97,
	0xb9, 0x80, 0x99, 0x25, 0x79, 0xb3, 0x15, 0xe9, 0x96, 0xff, 0xc7, 0xc4, 0x47, 0x1d, 0xbd, 0xcf,
	0xa1, 0xe1, 0xcc, 0x93, 0x64, 0xe3, 0xe2, 0x14, 0x2a, 0x59, 0x3e, 0xcd, 0xef, 0xb3, 0x85, 0x85,
	0x03, 0xb9, 0x8a, 0xbe, 0xf8, 0xb5, 0xb2, 0x33, 0x9e, 0x55, 0x33, 0x3b, 0x84, 0x6a, 0x2c, 0x5e,
	0x89, 0x70, 0x2c, 0xf0, 0x01, 0xab, 0xc1, 0xa3, 0xc8, 0x17, 0x1e, 0x96, 0xd8, 0x31, 0x1c, 0x7a,
	0xa4, 0x79, 0x10, 0xa8, 0xb1, 0xf0, 0x15, 0x3e, 0xb4, 0x75, 0x1e, 0x69, 0x1b, 0x61, 0x99, 0x35,
	0xa0, 0xe6, 0xbb, 0x24, 0xb4, 0x3f, 0x98, 0xe0, 0x23, 0xc6, 0xa0, 0xe9, 0x91, 0x36, 0xa3, 0x89,
	0x11, 0xa4, 0xc7, 0xa1, 0x7c, 0x85, 0x07, 0xac, 0x03, 0xad, 0xc0, 0x57, 0xda, 0xf0, 0x20, 0x58,
	0xa3, 0x0a, 0x2b, 0xac, 0x05, 0x47, 0xab, 0xc8, 0x48, 0x52, 0xa4, 0xb1, 0x6a, 0xb9, 0x16, 0xdd,
	0xa1, 0x4b, 0x58, 0xb3, 0x11, 0x77, 0x74, 0xcc, 0x75, 0x28, 0xb1, 0xce, 0xda, 0x80, 0x36, 0xe7,
	0xd2, 0x8d, 0xef, 0x90, 0xd1, 0xbc, 0x1f, 0x10, 0x02, 0xeb, 0x42, 0x5b, 0xd2, 0x40, 0x92, 0x1a,
	0xee, 0x67, 0x0e, 0xad, 0x13, 0x65, 0xeb, 0x87, 0x7d, 0xe3, 0x84, 0x62, 0xe0, 0x7b, 0xd8, 0xb0,
	0x98, 0xcb, 0x35, 0x37, 0x43, 0xe2, 0x52, 0xf7, 0x89, 0x6b, 0x3c, 0x62, 0x47, 0x50, 0xb7, 0xbc,
	0x51, 0x38, 0x26, 0x89, 0x4d, 0x06, 0x50, 0x91, 0xd4, 0x0f, 0x43, 0x8d, 0xc7, 0xf6, 0xe0, 0x96,
	0xc2, 0x19, 0x72, 0x21, 0x28, 0x40, 0x64, 0x27, 0x70, 0x3c, 0xf0, 0xe5, 0x68, 0xcc, 0x25, 0x99,
	0x38, 0x72, 0xb9, 0x26, 0x6c, 0xb1, 0x26, 0x80, 0xad, 0x8a, 0xb8, 0x30, 0xbe, 0x8b, 0xcc, 0x16,
	0xd9, 0x38, 0x20, 0xd7, 0x44, 0x5c, 0x6b, 0x92, 0x02, 0x4f, 0x18, 0x42, 0x83, 0xc7, 0x3a, 0x34,
	0xdc, 0x75, 0x25, 0x29, 0x85, 0xed, 0xf5, 0xa4, 0x76, 0xfc, 0x75, 0xd6, 0x98, 0x08, 0x5d, 0x32,
	0x4e, 0x18, 0x0b, 0x8d, 0xa7, 0x76, 0x4c, 0x61, 0xb4, 0x18, 0x89, 0xa1, 0x1b, 0x12, 0x1a, 0x1f,
	0x6f, 0x8e, 0x46, 0x37, 0xeb, 0xd6, 0xee, 0x86, 0x6e, 0x8b, 0x9d, 0x59, 0xff, 0x3c, 0x8a, 0x82,
	0x89, 0x51, 0x0e, 0x09, 0xc2, 0x73, 0x76, 0x06, 0x9d, 0x25, 0xe0, 0x4e, 0x04, 0x1f, 0xf9, 0x8e,
	0x89, 0x64, 0x38, 0xf0, 0x03, 0xc2, 0x27, 0xec, 0x31, 0x9c, 0xd8, 0xfe, 0xc0, 0xf7, 0x86, 0xda,
	0x17, 0x9e, 0x51, 0x9a, 0xeb, 0x58, 0xe1, 0x53, 0x7b, 0x0b, 0x6a, 0xec, 0x19, 0x4d, 0x4a, 0xe3,
	0x27, 0xec, 0x1c, 0x4e, 0x17, 0xd2, 0xfb, 0xfd, 0xf6, 0xe0, 0x9f, 0xda, 0xd1, 0xe9, 0xd0, 0xf3,
	0x02, 0xc2, 0x67, 0xd6, 0x4e, 0xc4, 0x63, 0x45, 0xeb, 0x0a, 0x85, 0x17, 0xec, 0x19, 0x3c, 0xf1,
	0x76, 0x7a, 0x3f, 0x94, 0xfa, 0x8c, 0x75, 0x00, 0xf7, 0xaf, 0xc7, 0x3c, 0xc7, 0xdf, 0xaa, 0x05,
	0xf0, 0x57, 0xf8, 0x7b, 0x11, 0x7c, 0x8d, 0x7f, 0x14, 0xc1, 0x2f, 0xf0, 0xcf, 0x22, 0xf8, 0x6b,
	0xfc, 0xab, 0x08, 0xfe, 0x06, 0xff, 0x2e, 0x82, 0x5f, 0xe2, 0x3f, 0x45, 0xf0, 0xb7, 0xf8, 0x6f,
	0xf5, 0xfa, 0x35, 0x54, 0xe3, 0xe5, 0x0e, 0x62, 0x21, 0x30, 0x67, 0x96, 0xa6, 0xc9, 0x9b, 0x5c,
	0xbd, 0x4f, 0xdf, 0x3a, 0xb7, 0xd3, 0x34, 0x4d, 0x7e, 0x62, 0x1f, 0xbf, 0xe0, 0xd5, 0x46, 0x3a,
	0x3f, 0x2b, 0xc8, 0x2c, 0x9f, 0x6e, 0xef, 0xc1, 0x65, 0xe9, 0x79, 0xa9, 0x5f, 0x7b, 0xbd, 0x5a,
	0x68, 0xdf, 0x55, 0x16, 0xfb, 0xed, 0xc5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x48, 0x5e, 0x41,
	0xe2, 0xf7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UthingsClient is the client API for Uthings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UthingsClient interface {
	// stubbed in
	ConnectSwniChannel(ctx context.Context, opts ...grpc.CallOption) (Uthings_ConnectSwniChannelClient, error)
}

type uthingsClient struct {
	cc grpc.ClientConnInterface
}

func NewUthingsClient(cc grpc.ClientConnInterface) UthingsClient {
	return &uthingsClient{cc}
}

func (c *uthingsClient) ConnectSwniChannel(ctx context.Context, opts ...grpc.CallOption) (Uthings_ConnectSwniChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Uthings_serviceDesc.Streams[0], "/uthing.Uthings/ConnectSwniChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &uthingsConnectSwniChannelClient{stream}
	return x, nil
}

type Uthings_ConnectSwniChannelClient interface {
	Send(*SmartCastRequest) error
	Recv() (*SmartCastResponse, error)
	grpc.ClientStream
}

type uthingsConnectSwniChannelClient struct {
	grpc.ClientStream
}

func (x *uthingsConnectSwniChannelClient) Send(m *SmartCastRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uthingsConnectSwniChannelClient) Recv() (*SmartCastResponse, error) {
	m := new(SmartCastResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UthingsServer is the server API for Uthings service.
type UthingsServer interface {
	// stubbed in
	ConnectSwniChannel(Uthings_ConnectSwniChannelServer) error
}

// UnimplementedUthingsServer can be embedded to have forward compatible implementations.
type UnimplementedUthingsServer struct {
}

func (*UnimplementedUthingsServer) ConnectSwniChannel(srv Uthings_ConnectSwniChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectSwniChannel not implemented")
}

func RegisterUthingsServer(s *grpc.Server, srv UthingsServer) {
	s.RegisterService(&_Uthings_serviceDesc, srv)
}

func _Uthings_ConnectSwniChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UthingsServer).ConnectSwniChannel(&uthingsConnectSwniChannelServer{stream})
}

type Uthings_ConnectSwniChannelServer interface {
	Send(*SmartCastResponse) error
	Recv() (*SmartCastRequest, error)
	grpc.ServerStream
}

type uthingsConnectSwniChannelServer struct {
	grpc.ServerStream
}

func (x *uthingsConnectSwniChannelServer) Send(m *SmartCastResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uthingsConnectSwniChannelServer) Recv() (*SmartCastRequest, error) {
	m := new(SmartCastRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Uthings_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uthing.Uthings",
	HandlerType: (*UthingsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectSwniChannel",
			Handler:       _Uthings_ConnectSwniChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "uthing/uthings.proto",
}