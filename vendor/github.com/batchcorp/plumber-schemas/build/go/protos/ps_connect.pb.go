// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_connect.proto

package protos

import (
	fmt "fmt"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
	opts "github.com/batchcorp/plumber-schemas/build/go/protos/opts"
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

type GetAllConnectionsRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetAllConnectionsRequest) Reset()         { *m = GetAllConnectionsRequest{} }
func (m *GetAllConnectionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllConnectionsRequest) ProtoMessage()    {}
func (*GetAllConnectionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{0}
}

func (m *GetAllConnectionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllConnectionsRequest.Unmarshal(m, b)
}
func (m *GetAllConnectionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllConnectionsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllConnectionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllConnectionsRequest.Merge(m, src)
}
func (m *GetAllConnectionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllConnectionsRequest.Size(m)
}
func (m *GetAllConnectionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllConnectionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllConnectionsRequest proto.InternalMessageInfo

func (m *GetAllConnectionsRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type GetAllConnectionsResponse struct {
	Options              []*opts.ConnectionOptions `protobuf:"bytes,1,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetAllConnectionsResponse) Reset()         { *m = GetAllConnectionsResponse{} }
func (m *GetAllConnectionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllConnectionsResponse) ProtoMessage()    {}
func (*GetAllConnectionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{1}
}

func (m *GetAllConnectionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllConnectionsResponse.Unmarshal(m, b)
}
func (m *GetAllConnectionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllConnectionsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllConnectionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllConnectionsResponse.Merge(m, src)
}
func (m *GetAllConnectionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllConnectionsResponse.Size(m)
}
func (m *GetAllConnectionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllConnectionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllConnectionsResponse proto.InternalMessageInfo

func (m *GetAllConnectionsResponse) GetOptions() []*opts.ConnectionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type GetConnectionRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ConnectionId         string       `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetConnectionRequest) Reset()         { *m = GetConnectionRequest{} }
func (m *GetConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*GetConnectionRequest) ProtoMessage()    {}
func (*GetConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{2}
}

func (m *GetConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConnectionRequest.Unmarshal(m, b)
}
func (m *GetConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConnectionRequest.Marshal(b, m, deterministic)
}
func (m *GetConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConnectionRequest.Merge(m, src)
}
func (m *GetConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_GetConnectionRequest.Size(m)
}
func (m *GetConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConnectionRequest proto.InternalMessageInfo

func (m *GetConnectionRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *GetConnectionRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

type GetConnectionResponse struct {
	Options              *opts.ConnectionOptions `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GetConnectionResponse) Reset()         { *m = GetConnectionResponse{} }
func (m *GetConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*GetConnectionResponse) ProtoMessage()    {}
func (*GetConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{3}
}

func (m *GetConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConnectionResponse.Unmarshal(m, b)
}
func (m *GetConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConnectionResponse.Marshal(b, m, deterministic)
}
func (m *GetConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConnectionResponse.Merge(m, src)
}
func (m *GetConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_GetConnectionResponse.Size(m)
}
func (m *GetConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetConnectionResponse proto.InternalMessageInfo

func (m *GetConnectionResponse) GetOptions() *opts.ConnectionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type CreateConnectionRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth            `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Options              *opts.ConnectionOptions `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CreateConnectionRequest) Reset()         { *m = CreateConnectionRequest{} }
func (m *CreateConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateConnectionRequest) ProtoMessage()    {}
func (*CreateConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{4}
}

func (m *CreateConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConnectionRequest.Unmarshal(m, b)
}
func (m *CreateConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConnectionRequest.Marshal(b, m, deterministic)
}
func (m *CreateConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConnectionRequest.Merge(m, src)
}
func (m *CreateConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateConnectionRequest.Size(m)
}
func (m *CreateConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConnectionRequest proto.InternalMessageInfo

func (m *CreateConnectionRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *CreateConnectionRequest) GetOptions() *opts.ConnectionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type CreateConnectionResponse struct {
	// Set with uuid that represents connection if create is successful
	ConnectionId         string   `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConnectionResponse) Reset()         { *m = CreateConnectionResponse{} }
func (m *CreateConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateConnectionResponse) ProtoMessage()    {}
func (*CreateConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{5}
}

func (m *CreateConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConnectionResponse.Unmarshal(m, b)
}
func (m *CreateConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConnectionResponse.Marshal(b, m, deterministic)
}
func (m *CreateConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConnectionResponse.Merge(m, src)
}
func (m *CreateConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateConnectionResponse.Size(m)
}
func (m *CreateConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConnectionResponse proto.InternalMessageInfo

func (m *CreateConnectionResponse) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

type TestConnectionRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth            `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Options              *opts.ConnectionOptions `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TestConnectionRequest) Reset()         { *m = TestConnectionRequest{} }
func (m *TestConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*TestConnectionRequest) ProtoMessage()    {}
func (*TestConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{6}
}

func (m *TestConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestConnectionRequest.Unmarshal(m, b)
}
func (m *TestConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestConnectionRequest.Marshal(b, m, deterministic)
}
func (m *TestConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestConnectionRequest.Merge(m, src)
}
func (m *TestConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_TestConnectionRequest.Size(m)
}
func (m *TestConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestConnectionRequest proto.InternalMessageInfo

func (m *TestConnectionRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *TestConnectionRequest) GetOptions() *opts.ConnectionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type TestConnectionResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TestConnectionResponse) Reset()         { *m = TestConnectionResponse{} }
func (m *TestConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*TestConnectionResponse) ProtoMessage()    {}
func (*TestConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{7}
}

func (m *TestConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestConnectionResponse.Unmarshal(m, b)
}
func (m *TestConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestConnectionResponse.Marshal(b, m, deterministic)
}
func (m *TestConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestConnectionResponse.Merge(m, src)
}
func (m *TestConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_TestConnectionResponse.Size(m)
}
func (m *TestConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestConnectionResponse proto.InternalMessageInfo

func (m *TestConnectionResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type UpdateConnectionRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth            `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ConnectionId         string                  `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	Options              *opts.ConnectionOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UpdateConnectionRequest) Reset()         { *m = UpdateConnectionRequest{} }
func (m *UpdateConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateConnectionRequest) ProtoMessage()    {}
func (*UpdateConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{8}
}

func (m *UpdateConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConnectionRequest.Unmarshal(m, b)
}
func (m *UpdateConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConnectionRequest.Marshal(b, m, deterministic)
}
func (m *UpdateConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConnectionRequest.Merge(m, src)
}
func (m *UpdateConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateConnectionRequest.Size(m)
}
func (m *UpdateConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConnectionRequest proto.InternalMessageInfo

func (m *UpdateConnectionRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *UpdateConnectionRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *UpdateConnectionRequest) GetOptions() *opts.ConnectionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type UpdateConnectionResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateConnectionResponse) Reset()         { *m = UpdateConnectionResponse{} }
func (m *UpdateConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateConnectionResponse) ProtoMessage()    {}
func (*UpdateConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{9}
}

func (m *UpdateConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConnectionResponse.Unmarshal(m, b)
}
func (m *UpdateConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConnectionResponse.Marshal(b, m, deterministic)
}
func (m *UpdateConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConnectionResponse.Merge(m, src)
}
func (m *UpdateConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateConnectionResponse.Size(m)
}
func (m *UpdateConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConnectionResponse proto.InternalMessageInfo

func (m *UpdateConnectionResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type DeleteConnectionRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ConnectionId         string       `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteConnectionRequest) Reset()         { *m = DeleteConnectionRequest{} }
func (m *DeleteConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectionRequest) ProtoMessage()    {}
func (*DeleteConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{10}
}

func (m *DeleteConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectionRequest.Unmarshal(m, b)
}
func (m *DeleteConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectionRequest.Marshal(b, m, deterministic)
}
func (m *DeleteConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectionRequest.Merge(m, src)
}
func (m *DeleteConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectionRequest.Size(m)
}
func (m *DeleteConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectionRequest proto.InternalMessageInfo

func (m *DeleteConnectionRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *DeleteConnectionRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

type DeleteConnectionResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteConnectionResponse) Reset()         { *m = DeleteConnectionResponse{} }
func (m *DeleteConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectionResponse) ProtoMessage()    {}
func (*DeleteConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb90b631f082347, []int{11}
}

func (m *DeleteConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectionResponse.Unmarshal(m, b)
}
func (m *DeleteConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectionResponse.Marshal(b, m, deterministic)
}
func (m *DeleteConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectionResponse.Merge(m, src)
}
func (m *DeleteConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectionResponse.Size(m)
}
func (m *DeleteConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectionResponse proto.InternalMessageInfo

func (m *DeleteConnectionResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAllConnectionsRequest)(nil), "protos.GetAllConnectionsRequest")
	proto.RegisterType((*GetAllConnectionsResponse)(nil), "protos.GetAllConnectionsResponse")
	proto.RegisterType((*GetConnectionRequest)(nil), "protos.GetConnectionRequest")
	proto.RegisterType((*GetConnectionResponse)(nil), "protos.GetConnectionResponse")
	proto.RegisterType((*CreateConnectionRequest)(nil), "protos.CreateConnectionRequest")
	proto.RegisterType((*CreateConnectionResponse)(nil), "protos.CreateConnectionResponse")
	proto.RegisterType((*TestConnectionRequest)(nil), "protos.TestConnectionRequest")
	proto.RegisterType((*TestConnectionResponse)(nil), "protos.TestConnectionResponse")
	proto.RegisterType((*UpdateConnectionRequest)(nil), "protos.UpdateConnectionRequest")
	proto.RegisterType((*UpdateConnectionResponse)(nil), "protos.UpdateConnectionResponse")
	proto.RegisterType((*DeleteConnectionRequest)(nil), "protos.DeleteConnectionRequest")
	proto.RegisterType((*DeleteConnectionResponse)(nil), "protos.DeleteConnectionResponse")
}

func init() { proto.RegisterFile("ps_connect.proto", fileDescriptor_dcb90b631f082347) }

var fileDescriptor_dcb90b631f082347 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xcf, 0x4a, 0xeb, 0x40,
	0x18, 0xc5, 0xc9, 0xbd, 0x97, 0x96, 0xfb, 0x55, 0x41, 0xa2, 0xb5, 0xb1, 0xa2, 0x94, 0xb8, 0xc9,
	0xc6, 0x04, 0xaa, 0x88, 0x3b, 0xa9, 0x2d, 0x54, 0x5d, 0x28, 0x46, 0xbb, 0x71, 0x53, 0xf2, 0x67,
	0x68, 0x02, 0x49, 0x66, 0xcc, 0x7c, 0x59, 0x89, 0xcf, 0xe0, 0x3b, 0xf8, 0x74, 0x3e, 0x86, 0x4c,
	0xa6, 0x31, 0xd2, 0x74, 0xd1, 0x3f, 0x88, 0xab, 0x09, 0x73, 0xce, 0x7c, 0xe7, 0x77, 0x26, 0x21,
	0xb0, 0xc5, 0xf8, 0xd8, 0xa3, 0x49, 0x42, 0x3c, 0x34, 0x59, 0x4a, 0x91, 0xaa, 0xb5, 0x7c, 0xe1,
	0xed, 0x7d, 0x8f, 0xc6, 0x31, 0x4d, 0xac, 0xdc, 0x20, 0x9e, 0xc6, 0x4e, 0x86, 0x81, 0x34, 0xb5,
	0x0f, 0x2a, 0x22, 0x47, 0x07, 0x33, 0x5e, 0xc8, 0x94, 0x21, 0x17, 0xa2, 0x58, 0x8b, 0xf9, 0x21,
	0x4d, 0xa4, 0xac, 0x0f, 0x40, 0x1b, 0x12, 0xec, 0x45, 0x51, 0xff, 0x4b, 0xe1, 0x36, 0x79, 0xce,
	0x08, 0x47, 0xd5, 0x80, 0x7f, 0x22, 0x47, 0x7b, 0xbb, 0xed, 0x28, 0x46, 0xa3, 0xbb, 0x2d, 0x8f,
	0x70, 0x53, 0xc6, 0x98, 0xbd, 0x0c, 0x03, 0x3b, 0x77, 0xe8, 0x23, 0xd8, 0x9b, 0x33, 0x85, 0x33,
	0x9a, 0x70, 0xa2, 0x9e, 0x43, 0x9d, 0xb2, 0x7c, 0x4b, 0x53, 0x3a, 0x7f, 0x8d, 0x46, 0xf7, 0xb0,
	0x18, 0x24, 0x90, 0xcc, 0xf2, 0xc8, 0x9d, 0x74, 0xd9, 0x85, 0x5d, 0x27, 0xb0, 0x33, 0x24, 0x58,
	0x1a, 0x96, 0x06, 0x53, 0x8f, 0x60, 0xb3, 0xac, 0x3c, 0x0e, 0x7d, 0x4d, 0xe9, 0x28, 0xc6, 0x7f,
	0x7b, 0xa3, 0xdc, 0xbc, 0xf6, 0xf5, 0x7b, 0x68, 0xce, 0xc4, 0xcc, 0x23, 0x57, 0x96, 0x21, 0x7f,
	0x85, 0x56, 0x3f, 0x25, 0x0e, 0x92, 0x75, 0xe0, 0x57, 0x8f, 0xbf, 0x00, 0xad, 0x1a, 0x3f, 0x2d,
	0xb5, 0xd0, 0x95, 0xbc, 0x40, 0xf3, 0x91, 0x70, 0xfc, 0x1d, 0xfa, 0x2b, 0xd8, 0x9d, 0x0d, 0x9f,
	0xb2, 0x9b, 0x50, 0x93, 0x1f, 0xb7, 0xf6, 0x51, 0xcf, 0x67, 0x36, 0x67, 0xf2, 0x1f, 0x72, 0xd5,
	0x9e, 0xba, 0xf4, 0x77, 0x05, 0x5a, 0x23, 0xe6, 0xaf, 0xf9, 0x1e, 0x16, 0xb9, 0xb1, 0xef, 0x75,
	0xff, 0x2c, 0x57, 0xf7, 0x06, 0xb4, 0x2a, 0xe3, 0x8a, 0x85, 0x03, 0x68, 0x0d, 0x48, 0x44, 0x7e,
	0xbe, 0xaf, 0xa0, 0xae, 0x26, 0xad, 0x46, 0x7d, 0x79, 0xf6, 0x74, 0x3a, 0x09, 0x31, 0xc8, 0x5c,
	0xa1, 0x5b, 0xae, 0x83, 0x5e, 0xe0, 0xd1, 0x94, 0x59, 0x2c, 0xca, 0x62, 0x97, 0xa4, 0xc7, 0xdc,
	0x0b, 0x48, 0xec, 0x70, 0xcb, 0xcd, 0xc2, 0xc8, 0xb7, 0x26, 0xd4, 0x92, 0xd3, 0x5c, 0xf9, 0x7f,
	0x3c, 0xf9, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x34, 0xb0, 0x3a, 0x3a, 0x05, 0x00, 0x00,
}