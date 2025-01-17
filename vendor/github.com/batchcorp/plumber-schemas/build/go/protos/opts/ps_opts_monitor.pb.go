// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opts/ps_opts_monitor.proto

package opts

import (
	fmt "fmt"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
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

type MonitorServiceSchemaVersionType int32

const (
	// Default to "latest" check
	MonitorServiceSchemaVersionType_SERVICE_SCHEMA_VERSION_CONDITION_TYPE_LATEST MonitorServiceSchemaVersionType = 0
	// If this is used, consumer knows that they should look into
	// condition_args["version"]
	MonitorServiceSchemaVersionType_SERVICE_SCHEMA_VERSION_CONDITION_TYPE_EXACT MonitorServiceSchemaVersionType = 1
)

var MonitorServiceSchemaVersionType_name = map[int32]string{
	0: "SERVICE_SCHEMA_VERSION_CONDITION_TYPE_LATEST",
	1: "SERVICE_SCHEMA_VERSION_CONDITION_TYPE_EXACT",
}

var MonitorServiceSchemaVersionType_value = map[string]int32{
	"SERVICE_SCHEMA_VERSION_CONDITION_TYPE_LATEST": 0,
	"SERVICE_SCHEMA_VERSION_CONDITION_TYPE_EXACT":  1,
}

func (x MonitorServiceSchemaVersionType) String() string {
	return proto.EnumName(MonitorServiceSchemaVersionType_name, int32(x))
}

func (MonitorServiceSchemaVersionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad36ed35e180ed92, []int{0}
}

type ReadOperator int32

const (
	ReadOperator_READ_OPERATOR_UNSET ReadOperator = 0
	ReadOperator_READ_OPERATOR_EQ    ReadOperator = 1
	ReadOperator_READ_OPERATOR_LT    ReadOperator = 2
	ReadOperator_READ_OPERATOR_GT    ReadOperator = 3
	ReadOperator_READ_OPERATOR_LTE   ReadOperator = 4
	ReadOperator_READ_OPERATOR_GTE   ReadOperator = 5
)

var ReadOperator_name = map[int32]string{
	0: "READ_OPERATOR_UNSET",
	1: "READ_OPERATOR_EQ",
	2: "READ_OPERATOR_LT",
	3: "READ_OPERATOR_GT",
	4: "READ_OPERATOR_LTE",
	5: "READ_OPERATOR_GTE",
}

var ReadOperator_value = map[string]int32{
	"READ_OPERATOR_UNSET": 0,
	"READ_OPERATOR_EQ":    1,
	"READ_OPERATOR_LT":    2,
	"READ_OPERATOR_GT":    3,
	"READ_OPERATOR_LTE":   4,
	"READ_OPERATOR_GTE":   5,
}

func (x ReadOperator) String() string {
	return proto.EnumName(ReadOperator_name, int32(x))
}

func (ReadOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad36ed35e180ed92, []int{1}
}

type MonitorOptions struct {
	XId             string `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	XCreatedBy      string `protobuf:"bytes,2,opt,name=_created_by,json=CreatedBy,proto3" json:"_created_by,omitempty"`
	XCreatedAtTsUtc int64  `protobuf:"varint,3,opt,name=_created_at_ts_utc,json=CreatedAtTsUtc,proto3" json:"_created_at_ts_utc,omitempty"`
	Notes           string `protobuf:"bytes,4,opt,name=notes,proto3" json:"notes,omitempty"`
	// Types that are valid to be assigned to Config:
	//	*MonitorOptions_SchemaValidation
	//	*MonitorOptions_ServiceSchemaVersion
	//	*MonitorOptions_ReadHasData
	Config               isMonitorOptions_Config `protobuf_oneof:"config"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MonitorOptions) Reset()         { *m = MonitorOptions{} }
func (m *MonitorOptions) String() string { return proto.CompactTextString(m) }
func (*MonitorOptions) ProtoMessage()    {}
func (*MonitorOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad36ed35e180ed92, []int{0}
}

func (m *MonitorOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorOptions.Unmarshal(m, b)
}
func (m *MonitorOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorOptions.Marshal(b, m, deterministic)
}
func (m *MonitorOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorOptions.Merge(m, src)
}
func (m *MonitorOptions) XXX_Size() int {
	return xxx_messageInfo_MonitorOptions.Size(m)
}
func (m *MonitorOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorOptions proto.InternalMessageInfo

func (m *MonitorOptions) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *MonitorOptions) GetXCreatedBy() string {
	if m != nil {
		return m.XCreatedBy
	}
	return ""
}

func (m *MonitorOptions) GetXCreatedAtTsUtc() int64 {
	if m != nil {
		return m.XCreatedAtTsUtc
	}
	return 0
}

func (m *MonitorOptions) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

type isMonitorOptions_Config interface {
	isMonitorOptions_Config()
}

type MonitorOptions_SchemaValidation struct {
	SchemaValidation *MonitorSchemaValidationConfig `protobuf:"bytes,100,opt,name=schema_validation,json=schemaValidation,proto3,oneof"`
}

type MonitorOptions_ServiceSchemaVersion struct {
	ServiceSchemaVersion *MonitorServiceSchemaVersionConfig `protobuf:"bytes,101,opt,name=service_schema_version,json=serviceSchemaVersion,proto3,oneof"`
}

type MonitorOptions_ReadHasData struct {
	ReadHasData *MonitorReadHasDataConfig `protobuf:"bytes,102,opt,name=read_has_data,json=readHasData,proto3,oneof"`
}

func (*MonitorOptions_SchemaValidation) isMonitorOptions_Config() {}

func (*MonitorOptions_ServiceSchemaVersion) isMonitorOptions_Config() {}

func (*MonitorOptions_ReadHasData) isMonitorOptions_Config() {}

func (m *MonitorOptions) GetConfig() isMonitorOptions_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *MonitorOptions) GetSchemaValidation() *MonitorSchemaValidationConfig {
	if x, ok := m.GetConfig().(*MonitorOptions_SchemaValidation); ok {
		return x.SchemaValidation
	}
	return nil
}

func (m *MonitorOptions) GetServiceSchemaVersion() *MonitorServiceSchemaVersionConfig {
	if x, ok := m.GetConfig().(*MonitorOptions_ServiceSchemaVersion); ok {
		return x.ServiceSchemaVersion
	}
	return nil
}

func (m *MonitorOptions) GetReadHasData() *MonitorReadHasDataConfig {
	if x, ok := m.GetConfig().(*MonitorOptions_ReadHasData); ok {
		return x.ReadHasData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MonitorOptions) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MonitorOptions_SchemaValidation)(nil),
		(*MonitorOptions_ServiceSchemaVersion)(nil),
		(*MonitorOptions_ReadHasData)(nil),
	}
}

type MonitorSchemaValidationConfig struct {
	// SchemaID is the read's schema
	SchemaId []string `protobuf:"bytes,1,rep,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	// Array of validations to perform on the payload
	Validations          []*common.Validation `protobuf:"bytes,2,rep,name=validations,proto3" json:"validations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MonitorSchemaValidationConfig) Reset()         { *m = MonitorSchemaValidationConfig{} }
func (m *MonitorSchemaValidationConfig) String() string { return proto.CompactTextString(m) }
func (*MonitorSchemaValidationConfig) ProtoMessage()    {}
func (*MonitorSchemaValidationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad36ed35e180ed92, []int{1}
}

func (m *MonitorSchemaValidationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorSchemaValidationConfig.Unmarshal(m, b)
}
func (m *MonitorSchemaValidationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorSchemaValidationConfig.Marshal(b, m, deterministic)
}
func (m *MonitorSchemaValidationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorSchemaValidationConfig.Merge(m, src)
}
func (m *MonitorSchemaValidationConfig) XXX_Size() int {
	return xxx_messageInfo_MonitorSchemaValidationConfig.Size(m)
}
func (m *MonitorSchemaValidationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorSchemaValidationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorSchemaValidationConfig proto.InternalMessageInfo

func (m *MonitorSchemaValidationConfig) GetSchemaId() []string {
	if m != nil {
		return m.SchemaId
	}
	return nil
}

func (m *MonitorSchemaValidationConfig) GetValidations() []*common.Validation {
	if m != nil {
		return m.Validations
	}
	return nil
}

type MonitorServiceSchemaVersionConfig struct {
	ServiceId string                          `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Type      MonitorServiceSchemaVersionType `protobuf:"varint,2,opt,name=type,proto3,enum=protos.opts.MonitorServiceSchemaVersionType" json:"type,omitempty"`
	// Potentially set, depending on condition
	Args                 map[string]string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MonitorServiceSchemaVersionConfig) Reset()         { *m = MonitorServiceSchemaVersionConfig{} }
func (m *MonitorServiceSchemaVersionConfig) String() string { return proto.CompactTextString(m) }
func (*MonitorServiceSchemaVersionConfig) ProtoMessage()    {}
func (*MonitorServiceSchemaVersionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad36ed35e180ed92, []int{2}
}

func (m *MonitorServiceSchemaVersionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorServiceSchemaVersionConfig.Unmarshal(m, b)
}
func (m *MonitorServiceSchemaVersionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorServiceSchemaVersionConfig.Marshal(b, m, deterministic)
}
func (m *MonitorServiceSchemaVersionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorServiceSchemaVersionConfig.Merge(m, src)
}
func (m *MonitorServiceSchemaVersionConfig) XXX_Size() int {
	return xxx_messageInfo_MonitorServiceSchemaVersionConfig.Size(m)
}
func (m *MonitorServiceSchemaVersionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorServiceSchemaVersionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorServiceSchemaVersionConfig proto.InternalMessageInfo

func (m *MonitorServiceSchemaVersionConfig) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *MonitorServiceSchemaVersionConfig) GetType() MonitorServiceSchemaVersionType {
	if m != nil {
		return m.Type
	}
	return MonitorServiceSchemaVersionType_SERVICE_SCHEMA_VERSION_CONDITION_TYPE_LATEST
}

func (m *MonitorServiceSchemaVersionConfig) GetArgs() map[string]string {
	if m != nil {
		return m.Args
	}
	return nil
}

type MonitorReadHasDataConfig struct {
	// Which reads does this alert config pertain to?
	ReadIds              []string     `protobuf:"bytes,1,rep,name=read_ids,json=readIds,proto3" json:"read_ids,omitempty"`
	Operator             ReadOperator `protobuf:"varint,2,opt,name=operator,proto3,enum=protos.opts.ReadOperator" json:"operator,omitempty"`
	Value                int32        `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	IntervalSeconds      int32        `protobuf:"varint,4,opt,name=interval_seconds,json=intervalSeconds,proto3" json:"interval_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MonitorReadHasDataConfig) Reset()         { *m = MonitorReadHasDataConfig{} }
func (m *MonitorReadHasDataConfig) String() string { return proto.CompactTextString(m) }
func (*MonitorReadHasDataConfig) ProtoMessage()    {}
func (*MonitorReadHasDataConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad36ed35e180ed92, []int{3}
}

func (m *MonitorReadHasDataConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorReadHasDataConfig.Unmarshal(m, b)
}
func (m *MonitorReadHasDataConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorReadHasDataConfig.Marshal(b, m, deterministic)
}
func (m *MonitorReadHasDataConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorReadHasDataConfig.Merge(m, src)
}
func (m *MonitorReadHasDataConfig) XXX_Size() int {
	return xxx_messageInfo_MonitorReadHasDataConfig.Size(m)
}
func (m *MonitorReadHasDataConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorReadHasDataConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorReadHasDataConfig proto.InternalMessageInfo

func (m *MonitorReadHasDataConfig) GetReadIds() []string {
	if m != nil {
		return m.ReadIds
	}
	return nil
}

func (m *MonitorReadHasDataConfig) GetOperator() ReadOperator {
	if m != nil {
		return m.Operator
	}
	return ReadOperator_READ_OPERATOR_UNSET
}

func (m *MonitorReadHasDataConfig) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MonitorReadHasDataConfig) GetIntervalSeconds() int32 {
	if m != nil {
		return m.IntervalSeconds
	}
	return 0
}

func init() {
	proto.RegisterEnum("protos.opts.MonitorServiceSchemaVersionType", MonitorServiceSchemaVersionType_name, MonitorServiceSchemaVersionType_value)
	proto.RegisterEnum("protos.opts.ReadOperator", ReadOperator_name, ReadOperator_value)
	proto.RegisterType((*MonitorOptions)(nil), "protos.opts.MonitorOptions")
	proto.RegisterType((*MonitorSchemaValidationConfig)(nil), "protos.opts.MonitorSchemaValidationConfig")
	proto.RegisterType((*MonitorServiceSchemaVersionConfig)(nil), "protos.opts.MonitorServiceSchemaVersionConfig")
	proto.RegisterMapType((map[string]string)(nil), "protos.opts.MonitorServiceSchemaVersionConfig.ArgsEntry")
	proto.RegisterType((*MonitorReadHasDataConfig)(nil), "protos.opts.MonitorReadHasDataConfig")
}

func init() { proto.RegisterFile("opts/ps_opts_monitor.proto", fileDescriptor_ad36ed35e180ed92) }

var fileDescriptor_ad36ed35e180ed92 = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x4f, 0x13, 0x4b,
	0x18, 0x66, 0xbb, 0x2d, 0xa7, 0x7d, 0x7b, 0x0e, 0x2c, 0x73, 0x38, 0xc7, 0x05, 0x83, 0x96, 0x26,
	0x26, 0xb5, 0x62, 0xd7, 0x60, 0x8c, 0x28, 0x37, 0x96, 0xb2, 0x81, 0x46, 0xa0, 0x38, 0x5d, 0x88,
	0x78, 0x33, 0x99, 0xee, 0x0e, 0xed, 0xc6, 0x76, 0x67, 0x33, 0x33, 0x6d, 0xd2, 0x7b, 0xef, 0xfc,
	0x01, 0x5e, 0xf9, 0x13, 0xfc, 0x8f, 0x66, 0x3f, 0xfa, 0xc1, 0x87, 0x82, 0x57, 0x33, 0xef, 0xf3,
	0x3e, 0xf3, 0xcc, 0xfb, 0x09, 0xeb, 0x3c, 0x54, 0xd2, 0x0a, 0x25, 0x89, 0x4e, 0x32, 0xe0, 0x81,
	0xaf, 0xb8, 0xa8, 0x85, 0x82, 0x2b, 0x8e, 0x8a, 0xf1, 0x21, 0x6b, 0x91, 0x6b, 0x7d, 0xd3, 0xe5,
	0x83, 0x01, 0x0f, 0x22, 0x6a, 0x72, 0x23, 0x23, 0xda, 0xf7, 0x3d, 0xaa, 0x7c, 0x1e, 0x24, 0xfc,
	0xf2, 0x77, 0x1d, 0x96, 0x8e, 0x13, 0x85, 0x56, 0x18, 0xe1, 0x12, 0x2d, 0x83, 0x4e, 0x7c, 0xcf,
	0xd4, 0x4a, 0x5a, 0xa5, 0x80, 0x33, 0x4d, 0x0f, 0x3d, 0x82, 0x22, 0x71, 0x05, 0xa3, 0x8a, 0x79,
	0xa4, 0x33, 0x36, 0x33, 0xb1, 0xa3, 0xd0, 0x48, 0x90, 0xbd, 0x31, 0xaa, 0x02, 0x9a, 0xfa, 0xa9,
	0x22, 0x4a, 0x92, 0xa1, 0x72, 0x4d, 0xbd, 0xa4, 0x55, 0x74, 0xbc, 0x94, 0xd2, 0xea, 0xca, 0x91,
	0x67, 0xca, 0x45, 0xab, 0x90, 0x0b, 0xb8, 0x62, 0xd2, 0xcc, 0xc6, 0x2a, 0x89, 0x81, 0x2e, 0x60,
	0x45, 0xba, 0x3d, 0x36, 0xa0, 0x73, 0x01, 0x9a, 0x5e, 0x49, 0xab, 0x14, 0xb7, 0xab, 0xb5, 0xb9,
	0x8c, 0x6a, 0x69, 0xa8, 0xed, 0x98, 0x7c, 0x3e, 0xe5, 0x36, 0x78, 0x70, 0xe9, 0x77, 0x0f, 0x17,
	0xb0, 0x21, 0xaf, 0x79, 0xd0, 0x25, 0xfc, 0x2f, 0x99, 0x18, 0xf9, 0x2e, 0x23, 0x93, 0x2f, 0x98,
	0x90, 0x91, 0x3e, 0x8b, 0xf5, 0x6b, 0xb7, 0xea, 0x27, 0x2f, 0xd2, 0x6f, 0x12, 0xfe, 0xf4, 0x8f,
	0x55, 0x79, 0x8b, 0x17, 0xbd, 0x87, 0x7f, 0x04, 0xa3, 0x1e, 0xe9, 0x51, 0x49, 0x3c, 0xaa, 0xa8,
	0x79, 0x19, 0xcb, 0x3f, 0xb9, 0x4d, 0x1e, 0x33, 0xea, 0x1d, 0x52, 0xb9, 0x4f, 0x15, 0x9d, 0xaa,
	0x16, 0xc5, 0x0c, 0xdc, 0xcb, 0xc3, 0xa2, 0x1b, 0x3b, 0xca, 0x63, 0xd8, 0xf8, 0x6d, 0xce, 0xe8,
	0x21, 0x14, 0xd2, 0xbc, 0xe2, 0x9e, 0xe9, 0x95, 0x02, 0xce, 0x27, 0x40, 0xd3, 0x43, 0xbb, 0x50,
	0x9c, 0x15, 0x54, 0x9a, 0x99, 0x92, 0x5e, 0x29, 0x6e, 0xaf, 0x4d, 0x42, 0x4a, 0x66, 0xa2, 0x36,
	0x93, 0xc4, 0xf3, 0xec, 0xf2, 0xd7, 0x0c, 0x6c, 0xde, 0x59, 0x0f, 0xb4, 0x01, 0x30, 0xa9, 0xef,
	0x74, 0x68, 0x0a, 0x29, 0xd2, 0xf4, 0xd0, 0x3b, 0xc8, 0xaa, 0x71, 0xc8, 0xe2, 0xa1, 0x59, 0xda,
	0xde, 0xba, 0x6f, 0xb1, 0x9d, 0x71, 0xc8, 0x70, 0xfc, 0x12, 0x1d, 0x41, 0x96, 0x8a, 0xae, 0x34,
	0xf5, 0x38, 0xf8, 0x9d, 0x3f, 0x6b, 0x57, 0xad, 0x2e, 0xba, 0xd2, 0x0e, 0x94, 0x18, 0xe3, 0x58,
	0x65, 0xfd, 0x35, 0x14, 0xa6, 0x10, 0x32, 0x40, 0xff, 0xcc, 0xc6, 0x69, 0xd0, 0xd1, 0x35, 0x1a,
	0xcf, 0x11, 0xed, 0x0f, 0x59, 0x3a, 0xe4, 0x89, 0xf1, 0x36, 0xb3, 0xa3, 0x95, 0x7f, 0x68, 0x60,
	0xfe, 0xaa, 0x7d, 0x68, 0x0d, 0xf2, 0x71, 0xf3, 0x7d, 0x4f, 0xa6, 0x3d, 0xf8, 0x2b, 0xb2, 0x9b,
	0x9e, 0x44, 0xaf, 0x20, 0xcf, 0x43, 0x26, 0xa8, 0xe2, 0x22, 0x2d, 0xc2, 0xda, 0x95, 0x14, 0x22,
	0xb1, 0x56, 0x4a, 0xc0, 0x53, 0xea, 0x2c, 0x90, 0x68, 0x8d, 0x72, 0x69, 0x20, 0xe8, 0x29, 0x18,
	0x7e, 0xa0, 0x98, 0x18, 0xd1, 0x3e, 0x91, 0xcc, 0xe5, 0x81, 0x97, 0x2c, 0x52, 0x0e, 0x2f, 0x4f,
	0xf0, 0x76, 0x02, 0x57, 0xbf, 0x68, 0xf0, 0xf8, 0x8e, 0x02, 0xa3, 0x17, 0xb0, 0xd5, 0xb6, 0xf1,
	0x79, 0xb3, 0x61, 0x93, 0x76, 0xe3, 0xd0, 0x3e, 0xae, 0x93, 0x73, 0x1b, 0xb7, 0x9b, 0xad, 0x13,
	0xd2, 0x68, 0x9d, 0xec, 0x37, 0x9d, 0xe8, 0xe6, 0x5c, 0x9c, 0xda, 0xe4, 0xa8, 0xee, 0xd8, 0x6d,
	0xc7, 0x58, 0x40, 0x16, 0x3c, 0xbb, 0xdf, 0x0b, 0xfb, 0x63, 0xbd, 0xe1, 0x18, 0x5a, 0xf5, 0x9b,
	0x06, 0x7f, 0xcf, 0xa7, 0x88, 0x1e, 0xc0, 0xbf, 0xd8, 0xae, 0xef, 0x93, 0xd6, 0xa9, 0x8d, 0xeb,
	0x4e, 0x0b, 0x93, 0xb3, 0x93, 0xb6, 0x1d, 0x49, 0xaf, 0x82, 0x71, 0xd5, 0x61, 0x7f, 0x30, 0xb4,
	0x9b, 0xe8, 0x91, 0x63, 0x64, 0x6e, 0xa2, 0x07, 0x8e, 0xa1, 0xa3, 0xff, 0x60, 0xe5, 0x3a, 0xd7,
	0x36, 0xb2, 0x37, 0xe1, 0x03, 0xc7, 0x36, 0x72, 0x7b, 0xbb, 0x9f, 0xde, 0x74, 0x7d, 0xd5, 0x1b,
	0x76, 0xa2, 0x3d, 0xb0, 0x3a, 0x54, 0xb9, 0x3d, 0x97, 0x8b, 0xd0, 0x0a, 0xfb, 0xc3, 0x41, 0x87,
	0x89, 0xe7, 0xc9, 0x12, 0x49, 0xab, 0x33, 0xf4, 0xfb, 0x9e, 0xd5, 0xe5, 0x56, 0xd2, 0x35, 0x2b,
	0xea, 0x5a, 0x67, 0x31, 0x36, 0x5e, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x3a, 0x80, 0x4e,
	0x8b, 0x05, 0x00, 0x00,
}
