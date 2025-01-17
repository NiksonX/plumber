// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/ps_records_base.proto

package records

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

// Returned for read requests (server & cli)
type ReadRecord struct {
	// Unique id automatically created by plumber
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// Plumber counts the number of messages it reads; this number represents
	// the message number (useful for CLI).
	Num int64 `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	// Metadata may contain properties that cannot be found in the Raw message.
	// For example: read lag in Kafka.
	//
	// Metadata may also contain data such as "count" which is an incremental
	// number that plumber assigns to each message it receives. This is used
	// with read via CLI functionality to allow the user to quickly discern
	// whether this is message #1 or #500, etc.
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// UTC unix timestamp of when plumber received the message; a backend record
	// entry might have its own timestamp as well. This should be seconds.
	ReceivedAtUnixTsUtc int64 `protobuf:"varint,6,opt,name=received_at_unix_ts_utc,json=receivedAtUnixTsUtc,proto3" json:"received_at_unix_ts_utc,omitempty"`
	// Set _outside_ the backend; will contain the final value, regardless of
	// whether decoding options were specified for a read.
	// _This_ is what both CLI and desktop should display for the payload.
	Payload []byte `protobuf:"bytes,99,opt,name=payload,proto3" json:"payload,omitempty"`
	// Types that are valid to be assigned to Record:
	//	*ReadRecord_Kafka
	//	*ReadRecord_Rabbit
	//	*ReadRecord_Activemq
	//	*ReadRecord_AwsSqs
	//	*ReadRecord_AzureEventHub
	//	*ReadRecord_AzureServiceBus
	//	*ReadRecord_GcpPubsub
	//	*ReadRecord_Kubemq
	//	*ReadRecord_Mongo
	//	*ReadRecord_Mqtt
	//	*ReadRecord_Nats
	//	*ReadRecord_NatsStreaming
	//	*ReadRecord_Nsq
	//	*ReadRecord_Postgres
	//	*ReadRecord_Pulsar
	//	*ReadRecord_RabbitStreams
	//	*ReadRecord_RedisPubsub
	//	*ReadRecord_RedisStreams
	Record isReadRecord_Record `protobuf_oneof:"Record"`
	// Original backend message (encoded with gob, ie. *skafka.Message, etc.).
	// In most cases, you should use the oneof record instead of the raw message.
	XRaw []byte `protobuf:"bytes,1000,opt,name=_raw,json=Raw,proto3" json:"_raw,omitempty"`
	// Identifies which plumber instance received the event (set outside the backend)
	XPlumberId           string   `protobuf:"bytes,1001,opt,name=_plumber_id,json=PlumberId,proto3" json:"_plumber_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRecord) Reset()         { *m = ReadRecord{} }
func (m *ReadRecord) String() string { return proto.CompactTextString(m) }
func (*ReadRecord) ProtoMessage()    {}
func (*ReadRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_79a95b759e1a93f6, []int{0}
}

func (m *ReadRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRecord.Unmarshal(m, b)
}
func (m *ReadRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRecord.Marshal(b, m, deterministic)
}
func (m *ReadRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRecord.Merge(m, src)
}
func (m *ReadRecord) XXX_Size() int {
	return xxx_messageInfo_ReadRecord.Size(m)
}
func (m *ReadRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRecord proto.InternalMessageInfo

func (m *ReadRecord) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *ReadRecord) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *ReadRecord) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ReadRecord) GetReceivedAtUnixTsUtc() int64 {
	if m != nil {
		return m.ReceivedAtUnixTsUtc
	}
	return 0
}

func (m *ReadRecord) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type isReadRecord_Record interface {
	isReadRecord_Record()
}

type ReadRecord_Kafka struct {
	Kafka *Kafka `protobuf:"bytes,100,opt,name=kafka,proto3,oneof"`
}

type ReadRecord_Rabbit struct {
	Rabbit *Rabbit `protobuf:"bytes,101,opt,name=rabbit,proto3,oneof"`
}

type ReadRecord_Activemq struct {
	Activemq *ActiveMQ `protobuf:"bytes,102,opt,name=activemq,proto3,oneof"`
}

type ReadRecord_AwsSqs struct {
	AwsSqs *AWSSQS `protobuf:"bytes,103,opt,name=aws_sqs,json=awsSqs,proto3,oneof"`
}

type ReadRecord_AzureEventHub struct {
	AzureEventHub *AzureEventHub `protobuf:"bytes,104,opt,name=azure_event_hub,json=azureEventHub,proto3,oneof"`
}

type ReadRecord_AzureServiceBus struct {
	AzureServiceBus *AzureServiceBus `protobuf:"bytes,105,opt,name=azure_service_bus,json=azureServiceBus,proto3,oneof"`
}

type ReadRecord_GcpPubsub struct {
	GcpPubsub *GCPPubSub `protobuf:"bytes,106,opt,name=gcp_pubsub,json=gcpPubsub,proto3,oneof"`
}

type ReadRecord_Kubemq struct {
	Kubemq *KubeMQ `protobuf:"bytes,107,opt,name=kubemq,proto3,oneof"`
}

type ReadRecord_Mongo struct {
	Mongo *Mongo `protobuf:"bytes,108,opt,name=mongo,proto3,oneof"`
}

type ReadRecord_Mqtt struct {
	Mqtt *MQTT `protobuf:"bytes,109,opt,name=mqtt,proto3,oneof"`
}

type ReadRecord_Nats struct {
	Nats *Nats `protobuf:"bytes,110,opt,name=nats,proto3,oneof"`
}

type ReadRecord_NatsStreaming struct {
	NatsStreaming *NatsStreaming `protobuf:"bytes,111,opt,name=nats_streaming,json=natsStreaming,proto3,oneof"`
}

type ReadRecord_Nsq struct {
	Nsq *NSQ `protobuf:"bytes,112,opt,name=nsq,proto3,oneof"`
}

type ReadRecord_Postgres struct {
	Postgres *Postgres `protobuf:"bytes,113,opt,name=postgres,proto3,oneof"`
}

type ReadRecord_Pulsar struct {
	Pulsar *Pulsar `protobuf:"bytes,114,opt,name=pulsar,proto3,oneof"`
}

type ReadRecord_RabbitStreams struct {
	RabbitStreams *RabbitStreams `protobuf:"bytes,115,opt,name=rabbit_streams,json=rabbitStreams,proto3,oneof"`
}

type ReadRecord_RedisPubsub struct {
	RedisPubsub *RedisPubsub `protobuf:"bytes,116,opt,name=redis_pubsub,json=redisPubsub,proto3,oneof"`
}

type ReadRecord_RedisStreams struct {
	RedisStreams *RedisStreams `protobuf:"bytes,117,opt,name=redis_streams,json=redisStreams,proto3,oneof"`
}

func (*ReadRecord_Kafka) isReadRecord_Record() {}

func (*ReadRecord_Rabbit) isReadRecord_Record() {}

func (*ReadRecord_Activemq) isReadRecord_Record() {}

func (*ReadRecord_AwsSqs) isReadRecord_Record() {}

func (*ReadRecord_AzureEventHub) isReadRecord_Record() {}

func (*ReadRecord_AzureServiceBus) isReadRecord_Record() {}

func (*ReadRecord_GcpPubsub) isReadRecord_Record() {}

func (*ReadRecord_Kubemq) isReadRecord_Record() {}

func (*ReadRecord_Mongo) isReadRecord_Record() {}

func (*ReadRecord_Mqtt) isReadRecord_Record() {}

func (*ReadRecord_Nats) isReadRecord_Record() {}

func (*ReadRecord_NatsStreaming) isReadRecord_Record() {}

func (*ReadRecord_Nsq) isReadRecord_Record() {}

func (*ReadRecord_Postgres) isReadRecord_Record() {}

func (*ReadRecord_Pulsar) isReadRecord_Record() {}

func (*ReadRecord_RabbitStreams) isReadRecord_Record() {}

func (*ReadRecord_RedisPubsub) isReadRecord_Record() {}

func (*ReadRecord_RedisStreams) isReadRecord_Record() {}

func (m *ReadRecord) GetRecord() isReadRecord_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *ReadRecord) GetKafka() *Kafka {
	if x, ok := m.GetRecord().(*ReadRecord_Kafka); ok {
		return x.Kafka
	}
	return nil
}

func (m *ReadRecord) GetRabbit() *Rabbit {
	if x, ok := m.GetRecord().(*ReadRecord_Rabbit); ok {
		return x.Rabbit
	}
	return nil
}

func (m *ReadRecord) GetActivemq() *ActiveMQ {
	if x, ok := m.GetRecord().(*ReadRecord_Activemq); ok {
		return x.Activemq
	}
	return nil
}

func (m *ReadRecord) GetAwsSqs() *AWSSQS {
	if x, ok := m.GetRecord().(*ReadRecord_AwsSqs); ok {
		return x.AwsSqs
	}
	return nil
}

func (m *ReadRecord) GetAzureEventHub() *AzureEventHub {
	if x, ok := m.GetRecord().(*ReadRecord_AzureEventHub); ok {
		return x.AzureEventHub
	}
	return nil
}

func (m *ReadRecord) GetAzureServiceBus() *AzureServiceBus {
	if x, ok := m.GetRecord().(*ReadRecord_AzureServiceBus); ok {
		return x.AzureServiceBus
	}
	return nil
}

func (m *ReadRecord) GetGcpPubsub() *GCPPubSub {
	if x, ok := m.GetRecord().(*ReadRecord_GcpPubsub); ok {
		return x.GcpPubsub
	}
	return nil
}

func (m *ReadRecord) GetKubemq() *KubeMQ {
	if x, ok := m.GetRecord().(*ReadRecord_Kubemq); ok {
		return x.Kubemq
	}
	return nil
}

func (m *ReadRecord) GetMongo() *Mongo {
	if x, ok := m.GetRecord().(*ReadRecord_Mongo); ok {
		return x.Mongo
	}
	return nil
}

func (m *ReadRecord) GetMqtt() *MQTT {
	if x, ok := m.GetRecord().(*ReadRecord_Mqtt); ok {
		return x.Mqtt
	}
	return nil
}

func (m *ReadRecord) GetNats() *Nats {
	if x, ok := m.GetRecord().(*ReadRecord_Nats); ok {
		return x.Nats
	}
	return nil
}

func (m *ReadRecord) GetNatsStreaming() *NatsStreaming {
	if x, ok := m.GetRecord().(*ReadRecord_NatsStreaming); ok {
		return x.NatsStreaming
	}
	return nil
}

func (m *ReadRecord) GetNsq() *NSQ {
	if x, ok := m.GetRecord().(*ReadRecord_Nsq); ok {
		return x.Nsq
	}
	return nil
}

func (m *ReadRecord) GetPostgres() *Postgres {
	if x, ok := m.GetRecord().(*ReadRecord_Postgres); ok {
		return x.Postgres
	}
	return nil
}

func (m *ReadRecord) GetPulsar() *Pulsar {
	if x, ok := m.GetRecord().(*ReadRecord_Pulsar); ok {
		return x.Pulsar
	}
	return nil
}

func (m *ReadRecord) GetRabbitStreams() *RabbitStreams {
	if x, ok := m.GetRecord().(*ReadRecord_RabbitStreams); ok {
		return x.RabbitStreams
	}
	return nil
}

func (m *ReadRecord) GetRedisPubsub() *RedisPubsub {
	if x, ok := m.GetRecord().(*ReadRecord_RedisPubsub); ok {
		return x.RedisPubsub
	}
	return nil
}

func (m *ReadRecord) GetRedisStreams() *RedisStreams {
	if x, ok := m.GetRecord().(*ReadRecord_RedisStreams); ok {
		return x.RedisStreams
	}
	return nil
}

func (m *ReadRecord) GetXRaw() []byte {
	if m != nil {
		return m.XRaw
	}
	return nil
}

func (m *ReadRecord) GetXPlumberId() string {
	if m != nil {
		return m.XPlumberId
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ReadRecord) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ReadRecord_Kafka)(nil),
		(*ReadRecord_Rabbit)(nil),
		(*ReadRecord_Activemq)(nil),
		(*ReadRecord_AwsSqs)(nil),
		(*ReadRecord_AzureEventHub)(nil),
		(*ReadRecord_AzureServiceBus)(nil),
		(*ReadRecord_GcpPubsub)(nil),
		(*ReadRecord_Kubemq)(nil),
		(*ReadRecord_Mongo)(nil),
		(*ReadRecord_Mqtt)(nil),
		(*ReadRecord_Nats)(nil),
		(*ReadRecord_NatsStreaming)(nil),
		(*ReadRecord_Nsq)(nil),
		(*ReadRecord_Postgres)(nil),
		(*ReadRecord_Pulsar)(nil),
		(*ReadRecord_RabbitStreams)(nil),
		(*ReadRecord_RedisPubsub)(nil),
		(*ReadRecord_RedisStreams)(nil),
	}
}

// Used as an arg for write requests (server & cli)
type WriteRecord struct {
	// If encoding options are provided, this value will be updated by plumber
	// to contain the encoded payload _before_ passing it to the backend.
	// @gotags: kong:"help='Input string',name=input,xor=input,default"
	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty" kong:"help='Input string',name=input,xor=input,default"`
	// @gotags: kong:"help='Optional metadata a backend may use (key1:value,key2:value,etc)',name=input-metadata"
	InputMetadata        map[string]string `protobuf:"bytes,2,rep,name=input_metadata,json=inputMetadata,proto3" json:"input_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" kong:"help='Optional metadata a backend may use (key1:value,key2:value,etc)',name=input-metadata"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WriteRecord) Reset()         { *m = WriteRecord{} }
func (m *WriteRecord) String() string { return proto.CompactTextString(m) }
func (*WriteRecord) ProtoMessage()    {}
func (*WriteRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_79a95b759e1a93f6, []int{1}
}

func (m *WriteRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRecord.Unmarshal(m, b)
}
func (m *WriteRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRecord.Marshal(b, m, deterministic)
}
func (m *WriteRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRecord.Merge(m, src)
}
func (m *WriteRecord) XXX_Size() int {
	return xxx_messageInfo_WriteRecord.Size(m)
}
func (m *WriteRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRecord.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRecord proto.InternalMessageInfo

func (m *WriteRecord) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *WriteRecord) GetInputMetadata() map[string]string {
	if m != nil {
		return m.InputMetadata
	}
	return nil
}

// Used for communicating errors that occur during a read, write, relay, etc.
type ErrorRecord struct {
	OccurredAtUnixTsUtc  int64             `protobuf:"varint,1,opt,name=occurred_at_unix_ts_utc,json=occurredAtUnixTsUtc,proto3" json:"occurred_at_unix_ts_utc,omitempty"`
	Error                string            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Metadata             map[string][]byte `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ErrorRecord) Reset()         { *m = ErrorRecord{} }
func (m *ErrorRecord) String() string { return proto.CompactTextString(m) }
func (*ErrorRecord) ProtoMessage()    {}
func (*ErrorRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_79a95b759e1a93f6, []int{2}
}

func (m *ErrorRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorRecord.Unmarshal(m, b)
}
func (m *ErrorRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorRecord.Marshal(b, m, deterministic)
}
func (m *ErrorRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorRecord.Merge(m, src)
}
func (m *ErrorRecord) XXX_Size() int {
	return xxx_messageInfo_ErrorRecord.Size(m)
}
func (m *ErrorRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorRecord proto.InternalMessageInfo

func (m *ErrorRecord) GetOccurredAtUnixTsUtc() int64 {
	if m != nil {
		return m.OccurredAtUnixTsUtc
	}
	return 0
}

func (m *ErrorRecord) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ErrorRecord) GetMetadata() map[string][]byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadRecord)(nil), "protos.records.ReadRecord")
	proto.RegisterMapType((map[string]string)(nil), "protos.records.ReadRecord.MetadataEntry")
	proto.RegisterType((*WriteRecord)(nil), "protos.records.WriteRecord")
	proto.RegisterMapType((map[string]string)(nil), "protos.records.WriteRecord.InputMetadataEntry")
	proto.RegisterType((*ErrorRecord)(nil), "protos.records.ErrorRecord")
	proto.RegisterMapType((map[string][]byte)(nil), "protos.records.ErrorRecord.MetadataEntry")
}

func init() { proto.RegisterFile("records/ps_records_base.proto", fileDescriptor_79a95b759e1a93f6) }

var fileDescriptor_79a95b759e1a93f6 = []byte{
	// 957 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0x6b, 0x6f, 0x1b, 0x45,
	0x17, 0xc7, 0xed, 0x38, 0xd7, 0xe3, 0x38, 0xcf, 0xc3, 0x34, 0xd0, 0x21, 0x34, 0xc4, 0x04, 0x41,
	0x4d, 0xa5, 0xda, 0x50, 0x2a, 0x84, 0x0a, 0x42, 0x4d, 0x4b, 0xa8, 0xa3, 0x2a, 0xc8, 0x59, 0x27,
	0xaa, 0xc4, 0x9b, 0xd5, 0xec, 0xee, 0x74, 0xb3, 0xd8, 0x7b, 0xf1, 0x5c, 0x9c, 0x86, 0xcf, 0xc6,
	0xa7, 0xe0, 0x3d, 0xef, 0xe1, 0x5b, 0xa0, 0xb9, 0x6c, 0xbc, 0xb6, 0x27, 0x91, 0xfa, 0x2a, 0x3b,
	0x73, 0x7e, 0x73, 0xe6, 0x5c, 0xe6, 0xfc, 0x1d, 0xd8, 0x67, 0x34, 0xcc, 0x59, 0xc4, 0x7b, 0x05,
	0xf7, 0xed, 0xa7, 0x1f, 0x10, 0x4e, 0xbb, 0x05, 0xcb, 0x45, 0x8e, 0x76, 0xf4, 0x1f, 0xde, 0xb5,
	0xa6, 0xbd, 0x4f, 0x1d, 0xf8, 0x88, 0xbc, 0x1d, 0x11, 0xc3, 0xef, 0x1d, 0x38, 0xec, 0x8c, 0x04,
	0x41, 0x22, 0x2c, 0xf0, 0x99, 0x03, 0x20, 0xa1, 0x48, 0xa6, 0x34, 0x9d, 0x58, 0xa4, 0xed, 0x42,
	0xae, 0xb8, 0xcf, 0x27, 0xdc, 0x12, 0x1d, 0x17, 0xf1, 0x87, 0x64, 0xd4, 0xa7, 0x53, 0x9a, 0x09,
	0xff, 0x52, 0x06, 0x96, 0x7c, 0x74, 0x2b, 0xc9, 0x29, 0x9b, 0x26, 0x21, 0xf5, 0x03, 0x59, 0x7a,
	0xfd, 0xdc, 0xc1, 0xc6, 0x61, 0xe1, 0x17, 0x32, 0xe0, 0x37, 0x0e, 0x5d, 0x09, 0x8e, 0x64, 0x30,
	0x8b, 0xde, 0x55, 0xa1, 0x34, 0xcf, 0xe2, 0xdc, 0xda, 0x5d, 0x05, 0x4f, 0x27, 0x42, 0xdc, 0x61,
	0xce, 0x88, 0x28, 0x63, 0x7c, 0x78, 0x8b, 0xd9, 0xe7, 0x82, 0x51, 0x92, 0x26, 0x59, 0x6c, 0xc1,
	0x07, 0x2e, 0x90, 0x4f, 0xee, 0xe8, 0x42, 0x91, 0x73, 0x11, 0x33, 0xca, 0xef, 0x48, 0xb4, 0x90,
	0x63, 0x4e, 0xd8, 0x1d, 0xa1, 0x98, 0x56, 0xdb, 0x60, 0x4a, 0x4f, 0x5f, 0xb8, 0x40, 0x1a, 0x25,
	0x7c, 0xbe, 0xb2, 0x5f, 0xde, 0x8a, 0xcd, 0xb9, 0x3b, 0xfc, 0x0b, 0x00, 0x3c, 0x4a, 0x22, 0x4f,
	0x33, 0x68, 0x1f, 0x20, 0xa5, 0x9c, 0x93, 0x98, 0xfa, 0x49, 0x84, 0xeb, 0xed, 0x7a, 0x67, 0xcb,
	0xdb, 0xb2, 0x3b, 0x27, 0x11, 0xfa, 0x3f, 0x34, 0x32, 0x99, 0xe2, 0x46, 0xbb, 0xde, 0x69, 0x78,
	0xea, 0x13, 0xfd, 0x0c, 0x9b, 0x29, 0x15, 0x24, 0x22, 0x82, 0xe0, 0xd5, 0x76, 0xa3, 0xd3, 0x7c,
	0xd2, 0xe9, 0xce, 0xbf, 0xf2, 0xee, 0xcc, 0x7d, 0xf7, 0xd4, 0xa2, 0xc7, 0x99, 0x60, 0xd7, 0xde,
	0xcd, 0x49, 0xf4, 0x14, 0xee, 0x33, 0x1a, 0xd2, 0x64, 0x4a, 0x23, 0x9f, 0x08, 0x5f, 0x66, 0xc9,
	0x3b, 0x5f, 0x70, 0x5f, 0x8a, 0x10, 0xaf, 0xeb, 0xbb, 0xee, 0x95, 0xe6, 0x23, 0x71, 0x91, 0x25,
	0xef, 0xce, 0xf9, 0x85, 0x08, 0x11, 0x86, 0x8d, 0x82, 0x5c, 0x8f, 0x73, 0x12, 0xe1, 0xb0, 0x5d,
	0xef, 0x6c, 0x7b, 0xe5, 0x12, 0x3d, 0x86, 0x35, 0x3d, 0x47, 0x38, 0x6a, 0xd7, 0x3b, 0xcd, 0x27,
	0x1f, 0x2e, 0x86, 0xf4, 0x5a, 0x19, 0xfb, 0x35, 0xcf, 0x50, 0xe8, 0x6b, 0x58, 0x37, 0xb5, 0xc6,
	0x54, 0xf3, 0x1f, 0x2d, 0xa5, 0xa0, 0xad, 0xfd, 0x9a, 0x67, 0x39, 0xf4, 0x1d, 0x6c, 0x96, 0x73,
	0x86, 0xdf, 0xea, 0x33, 0x78, 0xf1, 0xcc, 0x91, 0xb6, 0x9f, 0x9e, 0xf5, 0x6b, 0xde, 0x0d, 0x8b,
	0xbe, 0x81, 0x0d, 0x3b, 0x7c, 0x38, 0x76, 0x5f, 0x75, 0xf4, 0x66, 0x38, 0x3c, 0x1b, 0xaa, 0xab,
	0xc8, 0x15, 0x1f, 0x4e, 0x38, 0x7a, 0x05, 0xff, 0x5b, 0x98, 0x46, 0x7c, 0xa9, 0x8f, 0xee, 0x2f,
	0x1d, 0x55, 0xd8, 0xb1, 0xa2, 0xfa, 0x32, 0xe8, 0xd7, 0xbc, 0x16, 0xa9, 0x6e, 0xa0, 0x53, 0xf8,
	0x60, 0x69, 0x58, 0x71, 0xa2, 0x5d, 0x1d, 0x38, 0x5d, 0x0d, 0x0d, 0xf7, 0x42, 0xf2, 0x7e, 0xcd,
	0x33, 0x41, 0xcc, 0xb6, 0xd0, 0x33, 0x80, 0xd9, 0x3c, 0xe3, 0xdf, 0xb5, 0x9f, 0x8f, 0x17, 0xfd,
	0xbc, 0x7a, 0x39, 0x18, 0xc8, 0x60, 0xa8, 0xc3, 0xd9, 0x8a, 0xc3, 0x62, 0xa0, 0x69, 0x55, 0x70,
	0x33, 0xe6, 0x78, 0xe4, 0xae, 0xc2, 0x6b, 0x19, 0x98, 0xd2, 0x59, 0x4e, 0x75, 0x54, 0xcf, 0x3d,
	0x1e, 0xbb, 0x3b, 0x7a, 0xaa, 0x8c, 0xaa, 0xa3, 0x9a, 0x42, 0x8f, 0x60, 0x55, 0xc9, 0x00, 0x4e,
	0x35, 0xbd, 0xbb, 0x44, 0x9f, 0x9d, 0x9f, 0xf7, 0x6b, 0x9e, 0x66, 0x14, 0xab, 0x86, 0x1e, 0x67,
	0x6e, 0xf6, 0x57, 0x22, 0x54, 0xfe, 0x9a, 0x41, 0xbf, 0xc0, 0xce, 0xbc, 0x40, 0xe0, 0xdc, 0xdd,
	0x0b, 0x75, 0x6a, 0x58, 0x42, 0xaa, 0x17, 0x59, 0x75, 0x03, 0x3d, 0x84, 0x46, 0xc6, 0x27, 0xb8,
	0xd0, 0x87, 0xef, 0x2d, 0x1d, 0x1e, 0xaa, 0xd4, 0x15, 0xa1, 0x1e, 0x5a, 0x29, 0x25, 0x78, 0xe2,
	0x7e, 0x68, 0x03, 0x6b, 0x57, 0x0f, 0xad, 0x64, 0x55, 0x85, 0x8d, 0xbe, 0x60, 0xe6, 0xae, 0xf0,
	0x40, 0x5b, 0x55, 0x85, 0x0d, 0xa7, 0x52, 0x9b, 0x17, 0x1c, 0xcc, 0xdd, 0xa9, 0x99, 0x61, 0x30,
	0xb9, 0xa8, 0x4b, 0x5b, 0xac, 0xba, 0x81, 0x9e, 0xc3, 0x76, 0x55, 0x8f, 0xb0, 0xd0, 0x5e, 0x3e,
	0x59, 0x56, 0x85, 0x28, 0xe1, 0xe6, 0x39, 0xf4, 0x6b, 0x5e, 0x93, 0xcd, 0x96, 0xe8, 0x25, 0xb4,
	0xe6, 0xa4, 0x0a, 0x4b, 0xed, 0xe2, 0x81, 0xd3, 0xc5, 0x2c, 0x0e, 0x73, 0x6d, 0x19, 0x06, 0x82,
	0x55, 0x9f, 0x91, 0x2b, 0xfc, 0xcf, 0x86, 0x96, 0x86, 0x86, 0x47, 0xae, 0xd0, 0x01, 0x34, 0xfd,
	0x62, 0x2c, 0xd3, 0x80, 0x32, 0x25, 0x6f, 0xff, 0x6e, 0x18, 0x7d, 0x1b, 0x98, 0xad, 0x93, 0x68,
	0xef, 0x07, 0x68, 0xcd, 0x49, 0x94, 0x12, 0xbc, 0x11, 0xbd, 0xb6, 0x42, 0xa8, 0x3e, 0xd1, 0x2e,
	0xac, 0x4d, 0xc9, 0x58, 0x52, 0xbc, 0xa2, 0xf7, 0xcc, 0xe2, 0xd9, 0xca, 0xf7, 0xf5, 0x17, 0x9b,
	0xb0, 0x6e, 0x64, 0xee, 0xf0, 0xcf, 0x3a, 0x34, 0xdf, 0xb0, 0x44, 0x50, 0xab, 0xaa, 0xbb, 0xb0,
	0x96, 0x64, 0x85, 0x14, 0xd6, 0x8f, 0x59, 0xa0, 0x0b, 0xd8, 0xd1, 0x1f, 0xfe, 0x8d, 0x80, 0xae,
	0x68, 0x01, 0xed, 0x2e, 0xe6, 0x59, 0x71, 0xd5, 0x3d, 0x51, 0x27, 0xe6, 0x65, 0xb4, 0x95, 0x54,
	0xf7, 0xf6, 0x9e, 0x03, 0x5a, 0x86, 0xde, 0x27, 0x91, 0xc3, 0xbf, 0xeb, 0xd0, 0x3c, 0x66, 0x2c,
	0x67, 0x36, 0xfc, 0xa7, 0x70, 0x3f, 0x0f, 0x43, 0xc9, 0xd8, 0xb2, 0x3a, 0xd7, 0x8d, 0x3a, 0x97,
	0xe6, 0xaa, 0x3a, 0xef, 0xc2, 0x1a, 0x55, 0x4e, 0x4a, 0xff, 0x7a, 0x81, 0x8e, 0x2b, 0xbf, 0x17,
	0x0d, 0x9d, 0xee, 0x57, 0x8b, 0xe9, 0x56, 0xae, 0xbe, 0xed, 0x07, 0xe3, 0xbd, 0x1b, 0xb5, 0x5d,
	0x6d, 0xd4, 0x4f, 0xbf, 0xfd, 0x18, 0x27, 0x42, 0xfd, 0x5b, 0x13, 0xe6, 0x69, 0x2f, 0x20, 0x22,
	0xbc, 0x0c, 0x73, 0x56, 0xf4, 0xec, 0xd3, 0x78, 0xcc, 0xc3, 0x4b, 0x9a, 0x12, 0xde, 0x0b, 0x64,
	0x32, 0x8e, 0x7a, 0x71, 0xde, 0x33, 0x01, 0xf6, 0x6c, 0x80, 0xc1, 0xba, 0x5e, 0x7f, 0xfb, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0xff, 0xeb, 0x42, 0xee, 0x09, 0x00, 0x00,
}
