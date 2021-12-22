// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ccl/streamingccl/streampb/stream.proto

package streampb

import (
	fmt "fmt"
	roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
	hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// StreamPartitionSpec is the stream partition specification.
type StreamPartitionSpec struct {
	// start_from specifies the starting point for all spans.  If its empty,
	// an initial scan is performed.
	StartFrom hlc.Timestamp `protobuf:"bytes,1,opt,name=start_from,json=startFrom,proto3" json:"start_from"`
	// List of spans to stream.
	Spans  []roachpb.Span                      `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans"`
	Config StreamPartitionSpec_ExecutionConfig `protobuf:"bytes,3,opt,name=config,proto3" json:"config"`
}

func (m *StreamPartitionSpec) Reset()         { *m = StreamPartitionSpec{} }
func (m *StreamPartitionSpec) String() string { return proto.CompactTextString(m) }
func (*StreamPartitionSpec) ProtoMessage()    {}
func (*StreamPartitionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_44eb106be4866a3a, []int{0}
}
func (m *StreamPartitionSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamPartitionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *StreamPartitionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamPartitionSpec.Merge(m, src)
}
func (m *StreamPartitionSpec) XXX_Size() int {
	return m.Size()
}
func (m *StreamPartitionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamPartitionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_StreamPartitionSpec proto.InternalMessageInfo

// ExecutionConfig describes various knobs to control execution behavior
// of the stream.  If unspecified, reasonable defaults will be set.
type StreamPartitionSpec_ExecutionConfig struct {
	// Controls the number of concurrent scan requests issued during initial scan.
	InitialScanParallelism int32 `protobuf:"varint,1,opt,name=initial_scan_parallelism,json=initialScanParallelism,proto3" json:"initial_scan_parallelism,omitempty"`
	// Controls how often checkpoint records are published.
	MinCheckpointFrequency time.Duration `protobuf:"bytes,2,opt,name=min_checkpoint_frequency,json=minCheckpointFrequency,proto3,stdduration" json:"min_checkpoint_frequency"`
	// Controls batch size in bytes.
	BatchByteSize int64 `protobuf:"varint,3,opt,name=batch_byte_size,json=batchByteSize,proto3" json:"batch_byte_size,omitempty"`
}

func (m *StreamPartitionSpec_ExecutionConfig) Reset()         { *m = StreamPartitionSpec_ExecutionConfig{} }
func (m *StreamPartitionSpec_ExecutionConfig) String() string { return proto.CompactTextString(m) }
func (*StreamPartitionSpec_ExecutionConfig) ProtoMessage()    {}
func (*StreamPartitionSpec_ExecutionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_44eb106be4866a3a, []int{0, 0}
}
func (m *StreamPartitionSpec_ExecutionConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamPartitionSpec_ExecutionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *StreamPartitionSpec_ExecutionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamPartitionSpec_ExecutionConfig.Merge(m, src)
}
func (m *StreamPartitionSpec_ExecutionConfig) XXX_Size() int {
	return m.Size()
}
func (m *StreamPartitionSpec_ExecutionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamPartitionSpec_ExecutionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StreamPartitionSpec_ExecutionConfig proto.InternalMessageInfo

// StreamEvent describes a replication stream event
type StreamEvent struct {
	// Only 1 field ought to be set.
	Batch      *StreamEvent_Batch            `protobuf:"bytes,1,opt,name=batch,proto3" json:"batch,omitempty"`
	Checkpoint *StreamEvent_StreamCheckpoint `protobuf:"bytes,2,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (m *StreamEvent) Reset()         { *m = StreamEvent{} }
func (m *StreamEvent) String() string { return proto.CompactTextString(m) }
func (*StreamEvent) ProtoMessage()    {}
func (*StreamEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_44eb106be4866a3a, []int{1}
}
func (m *StreamEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *StreamEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEvent.Merge(m, src)
}
func (m *StreamEvent) XXX_Size() int {
	return m.Size()
}
func (m *StreamEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEvent.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEvent proto.InternalMessageInfo

type StreamEvent_Batch struct {
	KeyValues []roachpb.KeyValue `protobuf:"bytes,1,rep,name=key_values,json=keyValues,proto3" json:"key_values"`
}

func (m *StreamEvent_Batch) Reset()         { *m = StreamEvent_Batch{} }
func (m *StreamEvent_Batch) String() string { return proto.CompactTextString(m) }
func (*StreamEvent_Batch) ProtoMessage()    {}
func (*StreamEvent_Batch) Descriptor() ([]byte, []int) {
	return fileDescriptor_44eb106be4866a3a, []int{1, 0}
}
func (m *StreamEvent_Batch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamEvent_Batch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *StreamEvent_Batch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEvent_Batch.Merge(m, src)
}
func (m *StreamEvent_Batch) XXX_Size() int {
	return m.Size()
}
func (m *StreamEvent_Batch) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEvent_Batch.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEvent_Batch proto.InternalMessageInfo

// SpanCheckpoint represents a checkpoint record for completed span.
// All events up to timestamp must already have been emitted.
type StreamEvent_SpanCheckpoint struct {
	Span      roachpb.Span  `protobuf:"bytes,1,opt,name=span,proto3" json:"span"`
	Timestamp hlc.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp"`
}

func (m *StreamEvent_SpanCheckpoint) Reset()         { *m = StreamEvent_SpanCheckpoint{} }
func (m *StreamEvent_SpanCheckpoint) String() string { return proto.CompactTextString(m) }
func (*StreamEvent_SpanCheckpoint) ProtoMessage()    {}
func (*StreamEvent_SpanCheckpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_44eb106be4866a3a, []int{1, 1}
}
func (m *StreamEvent_SpanCheckpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamEvent_SpanCheckpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *StreamEvent_SpanCheckpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEvent_SpanCheckpoint.Merge(m, src)
}
func (m *StreamEvent_SpanCheckpoint) XXX_Size() int {
	return m.Size()
}
func (m *StreamEvent_SpanCheckpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEvent_SpanCheckpoint.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEvent_SpanCheckpoint proto.InternalMessageInfo

// Checkpoint represents stream checkpoint.
type StreamEvent_StreamCheckpoint struct {
	Spans []StreamEvent_SpanCheckpoint `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans"`
}

func (m *StreamEvent_StreamCheckpoint) Reset()         { *m = StreamEvent_StreamCheckpoint{} }
func (m *StreamEvent_StreamCheckpoint) String() string { return proto.CompactTextString(m) }
func (*StreamEvent_StreamCheckpoint) ProtoMessage()    {}
func (*StreamEvent_StreamCheckpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_44eb106be4866a3a, []int{1, 2}
}
func (m *StreamEvent_StreamCheckpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamEvent_StreamCheckpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *StreamEvent_StreamCheckpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEvent_StreamCheckpoint.Merge(m, src)
}
func (m *StreamEvent_StreamCheckpoint) XXX_Size() int {
	return m.Size()
}
func (m *StreamEvent_StreamCheckpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEvent_StreamCheckpoint.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEvent_StreamCheckpoint proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StreamPartitionSpec)(nil), "cockroach.ccl.streamingccl.StreamPartitionSpec")
	proto.RegisterType((*StreamPartitionSpec_ExecutionConfig)(nil), "cockroach.ccl.streamingccl.StreamPartitionSpec.ExecutionConfig")
	proto.RegisterType((*StreamEvent)(nil), "cockroach.ccl.streamingccl.StreamEvent")
	proto.RegisterType((*StreamEvent_Batch)(nil), "cockroach.ccl.streamingccl.StreamEvent.Batch")
	proto.RegisterType((*StreamEvent_SpanCheckpoint)(nil), "cockroach.ccl.streamingccl.StreamEvent.SpanCheckpoint")
	proto.RegisterType((*StreamEvent_StreamCheckpoint)(nil), "cockroach.ccl.streamingccl.StreamEvent.StreamCheckpoint")
}

func init() {
	proto.RegisterFile("ccl/streamingccl/streampb/stream.proto", fileDescriptor_44eb106be4866a3a)
}

var fileDescriptor_44eb106be4866a3a = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x8e, 0x9b, 0xa6, 0x6a, 0xb7, 0xfa, 0xfd, 0x8a, 0x16, 0x54, 0x8c, 0x11, 0x6e, 0xd5, 0x43,
	0x55, 0x21, 0xb1, 0x16, 0xad, 0x84, 0x7a, 0x03, 0x5c, 0x5a, 0x09, 0x71, 0xa9, 0x1c, 0x84, 0x10,
	0x52, 0x65, 0xad, 0xb7, 0x1b, 0x67, 0x95, 0xf5, 0xae, 0xb1, 0xd7, 0x15, 0xe9, 0x03, 0x70, 0xe6,
	0xc8, 0xf3, 0x70, 0xca, 0xb1, 0xe2, 0xd4, 0x13, 0x7f, 0x92, 0x0b, 0x8f, 0x81, 0xbc, 0x5e, 0xc7,
	0x69, 0x55, 0xd4, 0xdc, 0x66, 0x76, 0xe6, 0x9b, 0x99, 0x6f, 0xe6, 0xb3, 0xc1, 0x36, 0x21, 0xdc,
	0xcb, 0x55, 0x46, 0x71, 0xc2, 0x44, 0xdc, 0x38, 0x69, 0x64, 0x0c, 0x94, 0x66, 0x52, 0x49, 0xe8,
	0x10, 0x49, 0x06, 0x99, 0xc4, 0xa4, 0x8f, 0x08, 0xe1, 0x68, 0x16, 0xe1, 0x40, 0xfd, 0x9e, 0x46,
	0xde, 0x29, 0x56, 0xb8, 0xca, 0x77, 0xec, 0x42, 0x31, 0xee, 0xf5, 0x39, 0xf1, 0x14, 0x4b, 0x68,
	0xae, 0x70, 0x92, 0x9a, 0xc8, 0xbd, 0x58, 0xc6, 0x52, 0x9b, 0x5e, 0x69, 0x99, 0x57, 0x37, 0x96,
	0x32, 0xe6, 0xd4, 0xd3, 0x5e, 0x54, 0xf4, 0xbc, 0xd3, 0x22, 0xc3, 0x8a, 0x49, 0x51, 0xc5, 0xb7,
	0xbe, 0xb5, 0xc1, 0xdd, 0xae, 0x6e, 0x7a, 0x8c, 0x33, 0xc5, 0xca, 0x48, 0x37, 0xa5, 0x04, 0xfa,
	0x00, 0xe4, 0x0a, 0x67, 0x2a, 0xec, 0x65, 0x32, 0xb1, 0xad, 0x4d, 0x6b, 0x67, 0x75, 0xf7, 0x11,
	0x6a, 0x86, 0x2d, 0xc7, 0x40, 0x7d, 0x4e, 0xd0, 0xdb, 0x7a, 0x0c, 0x7f, 0x71, 0xf4, 0x63, 0xa3,
	0x15, 0xac, 0x68, 0xd8, 0x51, 0x26, 0x13, 0xb8, 0x07, 0x3a, 0x79, 0x8a, 0x45, 0x6e, 0x2f, 0x6c,
	0xb6, 0x77, 0x56, 0x77, 0xef, 0xcf, 0xc0, 0x0d, 0x33, 0xd4, 0x4d, 0xb1, 0x30, 0xc0, 0x2a, 0x17,
	0x9e, 0x80, 0x25, 0x22, 0x45, 0x8f, 0xc5, 0x76, 0x5b, 0x37, 0x7d, 0x8e, 0xfe, 0xbd, 0x21, 0x74,
	0xc3, 0xe4, 0xe8, 0xf0, 0x13, 0x25, 0x45, 0xe9, 0x1d, 0xe8, 0x32, 0xa6, 0xba, 0x29, 0xea, 0x7c,
	0xb7, 0xc0, 0xda, 0xb5, 0x0c, 0xb8, 0x0f, 0x6c, 0x26, 0x98, 0x62, 0x98, 0x87, 0x39, 0xc1, 0x22,
	0x4c, 0x71, 0x86, 0x39, 0xa7, 0x9c, 0xe5, 0x15, 0xf3, 0x4e, 0xb0, 0x6e, 0xe2, 0x5d, 0x82, 0xc5,
	0x71, 0x13, 0x85, 0x27, 0xc0, 0x4e, 0x98, 0x08, 0x49, 0x9f, 0x92, 0x41, 0x2a, 0x99, 0x28, 0xd7,
	0x45, 0x3f, 0x16, 0x54, 0x90, 0xa1, 0xbd, 0xa0, 0xc7, 0x7f, 0x80, 0xaa, 0x03, 0xa0, 0xfa, 0x00,
	0xe8, 0x95, 0x39, 0x80, 0xbf, 0x5c, 0x0e, 0xf6, 0xf5, 0xe7, 0x86, 0x15, 0xac, 0x27, 0x4c, 0x1c,
	0x4c, 0x6b, 0x1c, 0xd5, 0x25, 0xe0, 0x36, 0x58, 0x8b, 0xb0, 0x22, 0xfd, 0x30, 0x1a, 0x2a, 0x1a,
	0xe6, 0xec, 0x9c, 0xea, 0xa5, 0xb4, 0x83, 0xff, 0xf4, 0xb3, 0x3f, 0x54, 0xb4, 0xcb, 0xce, 0xe9,
	0xd6, 0x9f, 0x36, 0x58, 0xad, 0x56, 0x71, 0x78, 0x46, 0x85, 0x82, 0x07, 0xa0, 0xa3, 0x13, 0xcc,
	0xdd, 0x9e, 0xdc, 0xbe, 0x42, 0x8d, 0x43, 0x7e, 0x09, 0x0a, 0x2a, 0x2c, 0x7c, 0x0f, 0x40, 0xc3,
	0xcb, 0xb0, 0xd9, 0x9f, 0xb7, 0x52, 0x65, 0x37, 0x9c, 0x82, 0x99, 0x5a, 0xce, 0x6b, 0xd0, 0xd1,
	0x9d, 0xe0, 0x0b, 0x00, 0x06, 0x74, 0x18, 0x9e, 0x61, 0x5e, 0xd0, 0xdc, 0xb6, 0xb4, 0x4a, 0x1e,
	0xde, 0xa0, 0x92, 0x37, 0x74, 0xf8, 0xae, 0xcc, 0xa9, 0x25, 0x36, 0x30, 0x7e, 0xee, 0x7c, 0xb6,
	0xc0, 0xff, 0xa5, 0x86, 0x9a, 0x4e, 0xf0, 0x29, 0x58, 0x2c, 0x95, 0x64, 0xb8, 0xdf, 0x22, 0x3a,
	0x9d, 0x0a, 0x5f, 0x82, 0x95, 0xe9, 0xd7, 0x64, 0x98, 0xce, 0xa7, 0xf5, 0x29, 0xca, 0xe9, 0x81,
	0x3b, 0xd7, 0x39, 0xc3, 0xa0, 0xd6, 0x7f, 0xc5, 0xec, 0xd9, 0xdc, 0xcb, 0xbb, 0x42, 0xe8, 0xca,
	0xe7, 0xe1, 0x3f, 0x1e, 0xfd, 0x76, 0x5b, 0xa3, 0xb1, 0x6b, 0x5d, 0x8c, 0x5d, 0xeb, 0x72, 0xec,
	0x5a, 0xbf, 0xc6, 0xae, 0xf5, 0x65, 0xe2, 0xb6, 0x2e, 0x26, 0x6e, 0xeb, 0x72, 0xe2, 0xb6, 0x3e,
	0x2c, 0xd7, 0x3f, 0x9a, 0x68, 0x49, 0x6b, 0x6e, 0xef, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5,
	0xf9, 0xbe, 0xc6, 0x8c, 0x04, 0x00, 0x00,
}

func (m *StreamPartitionSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamPartitionSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamPartitionSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStream(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Spans) > 0 {
		for iNdEx := len(m.Spans) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Spans[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStream(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.StartFrom.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStream(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *StreamPartitionSpec_ExecutionConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamPartitionSpec_ExecutionConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamPartitionSpec_ExecutionConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BatchByteSize != 0 {
		i = encodeVarintStream(dAtA, i, uint64(m.BatchByteSize))
		i--
		dAtA[i] = 0x18
	}
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MinCheckpointFrequency, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MinCheckpointFrequency):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintStream(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if m.InitialScanParallelism != 0 {
		i = encodeVarintStream(dAtA, i, uint64(m.InitialScanParallelism))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StreamEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Checkpoint != nil {
		{
			size, err := m.Checkpoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStream(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Batch != nil {
		{
			size, err := m.Batch.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStream(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StreamEvent_Batch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamEvent_Batch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamEvent_Batch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeyValues) > 0 {
		for iNdEx := len(m.KeyValues) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KeyValues[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStream(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *StreamEvent_SpanCheckpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamEvent_SpanCheckpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamEvent_SpanCheckpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStream(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Span.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStream(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *StreamEvent_StreamCheckpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamEvent_StreamCheckpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamEvent_StreamCheckpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Spans) > 0 {
		for iNdEx := len(m.Spans) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Spans[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStream(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintStream(dAtA []byte, offset int, v uint64) int {
	offset -= sovStream(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StreamPartitionSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.StartFrom.Size()
	n += 1 + l + sovStream(uint64(l))
	if len(m.Spans) > 0 {
		for _, e := range m.Spans {
			l = e.Size()
			n += 1 + l + sovStream(uint64(l))
		}
	}
	l = m.Config.Size()
	n += 1 + l + sovStream(uint64(l))
	return n
}

func (m *StreamPartitionSpec_ExecutionConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitialScanParallelism != 0 {
		n += 1 + sovStream(uint64(m.InitialScanParallelism))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MinCheckpointFrequency)
	n += 1 + l + sovStream(uint64(l))
	if m.BatchByteSize != 0 {
		n += 1 + sovStream(uint64(m.BatchByteSize))
	}
	return n
}

func (m *StreamEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Batch != nil {
		l = m.Batch.Size()
		n += 1 + l + sovStream(uint64(l))
	}
	if m.Checkpoint != nil {
		l = m.Checkpoint.Size()
		n += 1 + l + sovStream(uint64(l))
	}
	return n
}

func (m *StreamEvent_Batch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.KeyValues) > 0 {
		for _, e := range m.KeyValues {
			l = e.Size()
			n += 1 + l + sovStream(uint64(l))
		}
	}
	return n
}

func (m *StreamEvent_SpanCheckpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Span.Size()
	n += 1 + l + sovStream(uint64(l))
	l = m.Timestamp.Size()
	n += 1 + l + sovStream(uint64(l))
	return n
}

func (m *StreamEvent_StreamCheckpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Spans) > 0 {
		for _, e := range m.Spans {
			l = e.Size()
			n += 1 + l + sovStream(uint64(l))
		}
	}
	return n
}

func sovStream(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStream(x uint64) (n int) {
	return sovStream(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamPartitionSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStream
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamPartitionSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamPartitionSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartFrom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StartFrom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spans = append(m.Spans, roachpb.Span{})
			if err := m.Spans[len(m.Spans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStream(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStream
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamPartitionSpec_ExecutionConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStream
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExecutionConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecutionConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialScanParallelism", wireType)
			}
			m.InitialScanParallelism = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialScanParallelism |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCheckpointFrequency", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MinCheckpointFrequency, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchByteSize", wireType)
			}
			m.BatchByteSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchByteSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStream(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStream
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStream
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Batch == nil {
				m.Batch = &StreamEvent_Batch{}
			}
			if err := m.Batch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Checkpoint == nil {
				m.Checkpoint = &StreamEvent_StreamCheckpoint{}
			}
			if err := m.Checkpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStream(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStream
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamEvent_Batch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStream
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Batch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Batch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyValues", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyValues = append(m.KeyValues, roachpb.KeyValue{})
			if err := m.KeyValues[len(m.KeyValues)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStream(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStream
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamEvent_SpanCheckpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStream
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SpanCheckpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpanCheckpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Span", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Span.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStream(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStream
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamEvent_StreamCheckpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStream
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamCheckpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamCheckpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spans = append(m.Spans, StreamEvent_SpanCheckpoint{})
			if err := m.Spans[len(m.Spans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStream(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStream
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStream(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStream
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStream
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStream
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStream
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStream
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStream
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStream        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStream          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStream = fmt.Errorf("proto: unexpected end of group")
)