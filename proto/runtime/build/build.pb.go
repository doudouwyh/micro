// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/micro/proto/runtime/build/build.proto

package build

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type BuildRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Options              *Options `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildRequest) Reset()         { *m = BuildRequest{} }
func (m *BuildRequest) String() string { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()    {}
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_368089b70e1f990d, []int{0}
}

func (m *BuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildRequest.Unmarshal(m, b)
}
func (m *BuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildRequest.Marshal(b, m, deterministic)
}
func (m *BuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildRequest.Merge(m, src)
}
func (m *BuildRequest) XXX_Size() int {
	return xxx_messageInfo_BuildRequest.Size(m)
}
func (m *BuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildRequest proto.InternalMessageInfo

func (m *BuildRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BuildRequest) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

type Options struct {
	Archive              string   `protobuf:"bytes,1,opt,name=archive,proto3" json:"archive,omitempty"`
	Entrypoint           string   `protobuf:"bytes,2,opt,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_368089b70e1f990d, []int{1}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetArchive() string {
	if m != nil {
		return m.Archive
	}
	return ""
}

func (m *Options) GetEntrypoint() string {
	if m != nil {
		return m.Entrypoint
	}
	return ""
}

type Result struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_368089b70e1f990d, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildRequest)(nil), "runtime.build.BuildRequest")
	proto.RegisterType((*Options)(nil), "runtime.build.Options")
	proto.RegisterType((*Result)(nil), "runtime.build.Result")
}

func init() {
	proto.RegisterFile("github.com/micro/micro/proto/runtime/build/build.proto", fileDescriptor_368089b70e1f990d)
}

var fileDescriptor_368089b70e1f990d = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0x87, 0x92, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0xfa, 0x45, 0xa5, 0x79, 0x25, 0x99, 0xb9, 0xa9, 0xfa, 0x49, 0xa5, 0x99, 0x39,
	0x29, 0x10, 0x52, 0x0f, 0x2c, 0x23, 0xc4, 0x0b, 0x95, 0xd2, 0x03, 0x0b, 0x2a, 0x85, 0x70, 0xf1,
	0x38, 0x81, 0x18, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89,
	0x25, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x90, 0x01, 0x17, 0x7b, 0x7e,
	0x41, 0x49, 0x66, 0x7e, 0x5e, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x98, 0x1e, 0x8a,
	0x21, 0x7a, 0xfe, 0x10, 0xd9, 0x20, 0x98, 0x32, 0x25, 0x67, 0x2e, 0x76, 0xa8, 0x98, 0x90, 0x04,
	0x17, 0x7b, 0x62, 0x51, 0x72, 0x46, 0x66, 0x59, 0x2a, 0xd8, 0x4c, 0xce, 0x20, 0x18, 0x57, 0x48,
	0x8e, 0x8b, 0x2b, 0x35, 0xaf, 0xa4, 0xa8, 0xb2, 0x20, 0x3f, 0x33, 0xaf, 0x04, 0x6c, 0x32, 0x67,
	0x10, 0x92, 0x88, 0x92, 0x0c, 0x17, 0x5b, 0x50, 0x6a, 0x71, 0x69, 0x0e, 0x56, 0x47, 0x19, 0x79,
	0x71, 0xb1, 0x82, 0x1d, 0x2e, 0xe4, 0x08, 0x63, 0x48, 0xa3, 0xb9, 0x0a, 0xd9, 0x5f, 0x52, 0xa2,
	0x68, 0x92, 0x10, 0x93, 0x95, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x9d, 0x4c, 0xa3, 0x8c, 0x71, 0x84,
	0x66, 0x99, 0x31, 0xb6, 0x00, 0xb5, 0x06, 0x93, 0x49, 0x6c, 0x60, 0x29, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x9b, 0xe8, 0xc8, 0x19, 0x8b, 0x01, 0x00, 0x00,
}