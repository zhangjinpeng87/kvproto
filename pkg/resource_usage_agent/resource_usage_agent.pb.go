// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resource_usage_agent.proto

package resource_usage_agent

import (
	"fmt"
	"io"
	"math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResourceMeteringRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceMeteringRequest) Reset()         { *m = ResourceMeteringRequest{} }
func (m *ResourceMeteringRequest) String() string { return proto.CompactTextString(m) }
func (*ResourceMeteringRequest) ProtoMessage()    {}
func (*ResourceMeteringRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_usage_agent_8426a9d1c2864a7c, []int{0}
}
func (m *ResourceMeteringRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceMeteringRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceMeteringRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ResourceMeteringRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceMeteringRequest.Merge(dst, src)
}
func (m *ResourceMeteringRequest) XXX_Size() int {
	return m.Size()
}
func (m *ResourceMeteringRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceMeteringRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceMeteringRequest proto.InternalMessageInfo

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_usage_agent_8426a9d1c2864a7c, []int{1}
}
func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(dst, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return m.Size()
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type ResourceUsageRecord struct {
	// Types that are valid to be assigned to RecordOneof:
	//	*ResourceUsageRecord_Record
	RecordOneof          isResourceUsageRecord_RecordOneof `protobuf_oneof:"record_oneof"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ResourceUsageRecord) Reset()         { *m = ResourceUsageRecord{} }
func (m *ResourceUsageRecord) String() string { return proto.CompactTextString(m) }
func (*ResourceUsageRecord) ProtoMessage()    {}
func (*ResourceUsageRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_usage_agent_8426a9d1c2864a7c, []int{2}
}
func (m *ResourceUsageRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceUsageRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceUsageRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ResourceUsageRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceUsageRecord.Merge(dst, src)
}
func (m *ResourceUsageRecord) XXX_Size() int {
	return m.Size()
}
func (m *ResourceUsageRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceUsageRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceUsageRecord proto.InternalMessageInfo

type isResourceUsageRecord_RecordOneof interface {
	isResourceUsageRecord_RecordOneof()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ResourceUsageRecord_Record struct {
	Record *GroupTagRecord `protobuf:"bytes,1,opt,name=record,oneof"`
}

func (*ResourceUsageRecord_Record) isResourceUsageRecord_RecordOneof() {}

func (m *ResourceUsageRecord) GetRecordOneof() isResourceUsageRecord_RecordOneof {
	if m != nil {
		return m.RecordOneof
	}
	return nil
}

func (m *ResourceUsageRecord) GetRecord() *GroupTagRecord {
	if x, ok := m.GetRecordOneof().(*ResourceUsageRecord_Record); ok {
		return x.Record
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ResourceUsageRecord) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ResourceUsageRecord_OneofMarshaler, _ResourceUsageRecord_OneofUnmarshaler, _ResourceUsageRecord_OneofSizer, []interface{}{
		(*ResourceUsageRecord_Record)(nil),
	}
}

func _ResourceUsageRecord_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ResourceUsageRecord)
	// record_oneof
	switch x := m.RecordOneof.(type) {
	case *ResourceUsageRecord_Record:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Record); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ResourceUsageRecord.RecordOneof has unexpected type %T", x)
	}
	return nil
}

func _ResourceUsageRecord_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ResourceUsageRecord)
	switch tag {
	case 1: // record_oneof.record
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GroupTagRecord)
		err := b.DecodeMessage(msg)
		m.RecordOneof = &ResourceUsageRecord_Record{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ResourceUsageRecord_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ResourceUsageRecord)
	// record_oneof
	switch x := m.RecordOneof.(type) {
	case *ResourceUsageRecord_Record:
		s := proto.Size(x.Record)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// GroupTagRecord is a set of resource usage data grouped by resource_group_tag.
type GroupTagRecord struct {
	ResourceGroupTag     []byte                `protobuf:"bytes,1,opt,name=resource_group_tag,json=resourceGroupTag,proto3" json:"resource_group_tag,omitempty"`
	Items                []*GroupTagRecordItem `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GroupTagRecord) Reset()         { *m = GroupTagRecord{} }
func (m *GroupTagRecord) String() string { return proto.CompactTextString(m) }
func (*GroupTagRecord) ProtoMessage()    {}
func (*GroupTagRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_usage_agent_8426a9d1c2864a7c, []int{3}
}
func (m *GroupTagRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupTagRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupTagRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GroupTagRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupTagRecord.Merge(dst, src)
}
func (m *GroupTagRecord) XXX_Size() int {
	return m.Size()
}
func (m *GroupTagRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupTagRecord.DiscardUnknown(m)
}

var xxx_messageInfo_GroupTagRecord proto.InternalMessageInfo

func (m *GroupTagRecord) GetResourceGroupTag() []byte {
	if m != nil {
		return m.ResourceGroupTag
	}
	return nil
}

func (m *GroupTagRecord) GetItems() []*GroupTagRecordItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type GroupTagRecordItem struct {
	TimestampSec         uint64   `protobuf:"varint,1,opt,name=timestamp_sec,json=timestampSec,proto3" json:"timestamp_sec,omitempty"`
	CpuTimeMs            uint32   `protobuf:"varint,2,opt,name=cpu_time_ms,json=cpuTimeMs,proto3" json:"cpu_time_ms,omitempty"`
	ReadKeys             uint32   `protobuf:"varint,3,opt,name=read_keys,json=readKeys,proto3" json:"read_keys,omitempty"`
	WriteKeys            uint32   `protobuf:"varint,4,opt,name=write_keys,json=writeKeys,proto3" json:"write_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupTagRecordItem) Reset()         { *m = GroupTagRecordItem{} }
func (m *GroupTagRecordItem) String() string { return proto.CompactTextString(m) }
func (*GroupTagRecordItem) ProtoMessage()    {}
func (*GroupTagRecordItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_usage_agent_8426a9d1c2864a7c, []int{4}
}
func (m *GroupTagRecordItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupTagRecordItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupTagRecordItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GroupTagRecordItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupTagRecordItem.Merge(dst, src)
}
func (m *GroupTagRecordItem) XXX_Size() int {
	return m.Size()
}
func (m *GroupTagRecordItem) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupTagRecordItem.DiscardUnknown(m)
}

var xxx_messageInfo_GroupTagRecordItem proto.InternalMessageInfo

func (m *GroupTagRecordItem) GetTimestampSec() uint64 {
	if m != nil {
		return m.TimestampSec
	}
	return 0
}

func (m *GroupTagRecordItem) GetCpuTimeMs() uint32 {
	if m != nil {
		return m.CpuTimeMs
	}
	return 0
}

func (m *GroupTagRecordItem) GetReadKeys() uint32 {
	if m != nil {
		return m.ReadKeys
	}
	return 0
}

func (m *GroupTagRecordItem) GetWriteKeys() uint32 {
	if m != nil {
		return m.WriteKeys
	}
	return 0
}

func init() {
	proto.RegisterType((*ResourceMeteringRequest)(nil), "resource_usage_agent.ResourceMeteringRequest")
	proto.RegisterType((*EmptyResponse)(nil), "resource_usage_agent.EmptyResponse")
	proto.RegisterType((*ResourceUsageRecord)(nil), "resource_usage_agent.ResourceUsageRecord")
	proto.RegisterType((*GroupTagRecord)(nil), "resource_usage_agent.GroupTagRecord")
	proto.RegisterType((*GroupTagRecordItem)(nil), "resource_usage_agent.GroupTagRecordItem")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceUsageAgent service

type ResourceUsageAgentClient interface {
	// Report the resource usage records. By default, the records with the same
	// resource group tag will be batched by minute.
	Report(ctx context.Context, opts ...grpc.CallOption) (ResourceUsageAgent_ReportClient, error)
}

type resourceUsageAgentClient struct {
	cc *grpc.ClientConn
}

func NewResourceUsageAgentClient(cc *grpc.ClientConn) ResourceUsageAgentClient {
	return &resourceUsageAgentClient{cc}
}

func (c *resourceUsageAgentClient) Report(ctx context.Context, opts ...grpc.CallOption) (ResourceUsageAgent_ReportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ResourceUsageAgent_serviceDesc.Streams[0], "/resource_usage_agent.ResourceUsageAgent/Report", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceUsageAgentReportClient{stream}
	return x, nil
}

type ResourceUsageAgent_ReportClient interface {
	Send(*ResourceUsageRecord) error
	CloseAndRecv() (*EmptyResponse, error)
	grpc.ClientStream
}

type resourceUsageAgentReportClient struct {
	grpc.ClientStream
}

func (x *resourceUsageAgentReportClient) Send(m *ResourceUsageRecord) error {
	return x.ClientStream.SendMsg(m)
}

func (x *resourceUsageAgentReportClient) CloseAndRecv() (*EmptyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ResourceUsageAgent service

type ResourceUsageAgentServer interface {
	// Report the resource usage records. By default, the records with the same
	// resource group tag will be batched by minute.
	Report(ResourceUsageAgent_ReportServer) error
}

func RegisterResourceUsageAgentServer(s *grpc.Server, srv ResourceUsageAgentServer) {
	s.RegisterService(&_ResourceUsageAgent_serviceDesc, srv)
}

func _ResourceUsageAgent_Report_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResourceUsageAgentServer).Report(&resourceUsageAgentReportServer{stream})
}

type ResourceUsageAgent_ReportServer interface {
	SendAndClose(*EmptyResponse) error
	Recv() (*ResourceUsageRecord, error)
	grpc.ServerStream
}

type resourceUsageAgentReportServer struct {
	grpc.ServerStream
}

func (x *resourceUsageAgentReportServer) SendAndClose(m *EmptyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *resourceUsageAgentReportServer) Recv() (*ResourceUsageRecord, error) {
	m := new(ResourceUsageRecord)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ResourceUsageAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resource_usage_agent.ResourceUsageAgent",
	HandlerType: (*ResourceUsageAgentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Report",
			Handler:       _ResourceUsageAgent_Report_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "resource_usage_agent.proto",
}

// Client API for ResourceMeteringPubSub service

type ResourceMeteringPubSubClient interface {
	// Clients subscribe to resource metering records through this RPC, and TiKV periodically (e.g. per minute)
	// publishes resource metering records to clients via gRPC stream.
	Subscribe(ctx context.Context, in *ResourceMeteringRequest, opts ...grpc.CallOption) (ResourceMeteringPubSub_SubscribeClient, error)
}

type resourceMeteringPubSubClient struct {
	cc *grpc.ClientConn
}

func NewResourceMeteringPubSubClient(cc *grpc.ClientConn) ResourceMeteringPubSubClient {
	return &resourceMeteringPubSubClient{cc}
}

func (c *resourceMeteringPubSubClient) Subscribe(ctx context.Context, in *ResourceMeteringRequest, opts ...grpc.CallOption) (ResourceMeteringPubSub_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ResourceMeteringPubSub_serviceDesc.Streams[0], "/resource_usage_agent.ResourceMeteringPubSub/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceMeteringPubSubSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ResourceMeteringPubSub_SubscribeClient interface {
	Recv() (*ResourceUsageRecord, error)
	grpc.ClientStream
}

type resourceMeteringPubSubSubscribeClient struct {
	grpc.ClientStream
}

func (x *resourceMeteringPubSubSubscribeClient) Recv() (*ResourceUsageRecord, error) {
	m := new(ResourceUsageRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ResourceMeteringPubSub service

type ResourceMeteringPubSubServer interface {
	// Clients subscribe to resource metering records through this RPC, and TiKV periodically (e.g. per minute)
	// publishes resource metering records to clients via gRPC stream.
	Subscribe(*ResourceMeteringRequest, ResourceMeteringPubSub_SubscribeServer) error
}

func RegisterResourceMeteringPubSubServer(s *grpc.Server, srv ResourceMeteringPubSubServer) {
	s.RegisterService(&_ResourceMeteringPubSub_serviceDesc, srv)
}

func _ResourceMeteringPubSub_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResourceMeteringRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourceMeteringPubSubServer).Subscribe(m, &resourceMeteringPubSubSubscribeServer{stream})
}

type ResourceMeteringPubSub_SubscribeServer interface {
	Send(*ResourceUsageRecord) error
	grpc.ServerStream
}

type resourceMeteringPubSubSubscribeServer struct {
	grpc.ServerStream
}

func (x *resourceMeteringPubSubSubscribeServer) Send(m *ResourceUsageRecord) error {
	return x.ServerStream.SendMsg(m)
}

var _ResourceMeteringPubSub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resource_usage_agent.ResourceMeteringPubSub",
	HandlerType: (*ResourceMeteringPubSubServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _ResourceMeteringPubSub_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "resource_usage_agent.proto",
}

func (m *ResourceMeteringRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceMeteringRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *EmptyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmptyResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ResourceUsageRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceUsageRecord) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RecordOneof != nil {
		nn1, err := m.RecordOneof.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ResourceUsageRecord_Record) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Record != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResourceUsageAgent(dAtA, i, uint64(m.Record.Size()))
		n2, err := m.Record.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *GroupTagRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupTagRecord) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ResourceGroupTag) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResourceUsageAgent(dAtA, i, uint64(len(m.ResourceGroupTag)))
		i += copy(dAtA[i:], m.ResourceGroupTag)
	}
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintResourceUsageAgent(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GroupTagRecordItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupTagRecordItem) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TimestampSec != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintResourceUsageAgent(dAtA, i, uint64(m.TimestampSec))
	}
	if m.CpuTimeMs != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintResourceUsageAgent(dAtA, i, uint64(m.CpuTimeMs))
	}
	if m.ReadKeys != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintResourceUsageAgent(dAtA, i, uint64(m.ReadKeys))
	}
	if m.WriteKeys != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintResourceUsageAgent(dAtA, i, uint64(m.WriteKeys))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintResourceUsageAgent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResourceMeteringRequest) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EmptyResponse) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResourceUsageRecord) Size() (n int) {
	var l int
	_ = l
	if m.RecordOneof != nil {
		n += m.RecordOneof.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResourceUsageRecord_Record) Size() (n int) {
	var l int
	_ = l
	if m.Record != nil {
		l = m.Record.Size()
		n += 1 + l + sovResourceUsageAgent(uint64(l))
	}
	return n
}
func (m *GroupTagRecord) Size() (n int) {
	var l int
	_ = l
	l = len(m.ResourceGroupTag)
	if l > 0 {
		n += 1 + l + sovResourceUsageAgent(uint64(l))
	}
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovResourceUsageAgent(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GroupTagRecordItem) Size() (n int) {
	var l int
	_ = l
	if m.TimestampSec != 0 {
		n += 1 + sovResourceUsageAgent(uint64(m.TimestampSec))
	}
	if m.CpuTimeMs != 0 {
		n += 1 + sovResourceUsageAgent(uint64(m.CpuTimeMs))
	}
	if m.ReadKeys != 0 {
		n += 1 + sovResourceUsageAgent(uint64(m.ReadKeys))
	}
	if m.WriteKeys != 0 {
		n += 1 + sovResourceUsageAgent(uint64(m.WriteKeys))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovResourceUsageAgent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozResourceUsageAgent(x uint64) (n int) {
	return sovResourceUsageAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceMeteringRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceUsageAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourceMeteringRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceMeteringRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipResourceUsageAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourceUsageAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EmptyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceUsageAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmptyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipResourceUsageAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourceUsageAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResourceUsageRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceUsageAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourceUsageRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceUsageRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceUsageAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResourceUsageAgent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &GroupTagRecord{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.RecordOneof = &ResourceUsageRecord_Record{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResourceUsageAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourceUsageAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GroupTagRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceUsageAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupTagRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupTagRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceGroupTag", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceUsageAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResourceUsageAgent
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceGroupTag = append(m.ResourceGroupTag[:0], dAtA[iNdEx:postIndex]...)
			if m.ResourceGroupTag == nil {
				m.ResourceGroupTag = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceUsageAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResourceUsageAgent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &GroupTagRecordItem{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResourceUsageAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourceUsageAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GroupTagRecordItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceUsageAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupTagRecordItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupTagRecordItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimestampSec", wireType)
			}
			m.TimestampSec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceUsageAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimestampSec |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CpuTimeMs", wireType)
			}
			m.CpuTimeMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceUsageAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CpuTimeMs |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadKeys", wireType)
			}
			m.ReadKeys = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceUsageAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReadKeys |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WriteKeys", wireType)
			}
			m.WriteKeys = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceUsageAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WriteKeys |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipResourceUsageAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourceUsageAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipResourceUsageAgent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourceUsageAgent
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
					return 0, ErrIntOverflowResourceUsageAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowResourceUsageAgent
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthResourceUsageAgent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowResourceUsageAgent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipResourceUsageAgent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthResourceUsageAgent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourceUsageAgent   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("resource_usage_agent.proto", fileDescriptor_resource_usage_agent_8426a9d1c2864a7c)
}

var fileDescriptor_resource_usage_agent_8426a9d1c2864a7c = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6b, 0xd4, 0x40,
	0x18, 0xde, 0xb1, 0x75, 0xe9, 0xbe, 0xbb, 0xdb, 0x96, 0xb1, 0xe8, 0x1a, 0x31, 0x94, 0x54, 0x24,
	0x82, 0xae, 0x12, 0xef, 0x05, 0x0b, 0xa2, 0x22, 0x05, 0x99, 0xad, 0x37, 0x61, 0x48, 0xd2, 0xd7,
	0x61, 0x58, 0x92, 0x19, 0xe7, 0xa3, 0xb2, 0x17, 0x2f, 0xfe, 0x06, 0xc1, 0x9f, 0xe0, 0x4f, 0xf1,
	0xe8, 0xd1, 0xa3, 0xac, 0x7f, 0x44, 0x32, 0xe9, 0x16, 0xb6, 0x86, 0xd2, 0x53, 0x5e, 0x9e, 0x8f,
	0x3c, 0xcf, 0xcc, 0xbc, 0x10, 0x19, 0xb4, 0xca, 0x9b, 0x12, 0xb9, 0xb7, 0xb9, 0x40, 0x9e, 0x0b,
	0xac, 0xdd, 0x54, 0x1b, 0xe5, 0x14, 0xdd, 0xeb, 0xe2, 0xa2, 0x3d, 0xa1, 0x84, 0x0a, 0x82, 0xa7,
	0xcd, 0xd4, 0x6a, 0xa3, 0x1d, 0xe3, 0xad, 0x0b, 0x63, 0x0b, 0x24, 0x77, 0xe1, 0x0e, 0x3b, 0xb7,
	0x1f, 0xa3, 0x43, 0x23, 0x6b, 0xc1, 0xf0, 0x93, 0x47, 0xeb, 0x92, 0x1d, 0x18, 0xbf, 0xac, 0xb4,
	0x5b, 0x30, 0xb4, 0x5a, 0xd5, 0x16, 0x13, 0x84, 0x5b, 0x2b, 0xed, 0xfb, 0x26, 0x89, 0x61, 0xa9,
	0xcc, 0x29, 0x3d, 0x84, 0xbe, 0x09, 0xd3, 0x84, 0xec, 0x93, 0x74, 0x98, 0x3d, 0x98, 0x76, 0x96,
	0x7d, 0x65, 0x94, 0xd7, 0x27, 0xb9, 0x68, 0x5d, 0xaf, 0x7b, 0xec, 0xdc, 0x75, 0xb4, 0x0d, 0xa3,
	0x76, 0xe2, 0xaa, 0x46, 0xf5, 0x31, 0xf9, 0x02, 0xdb, 0xeb, 0x5a, 0xfa, 0x18, 0xe8, 0xc5, 0x2f,
	0x45, 0x43, 0x71, 0x97, 0x8b, 0x90, 0x36, 0x62, 0xbb, 0x2b, 0x66, 0xe5, 0xa1, 0x87, 0x70, 0x53,
	0x3a, 0xac, 0xec, 0xe4, 0xc6, 0xfe, 0x46, 0x3a, 0xcc, 0xd2, 0xeb, 0xd4, 0x79, 0xe3, 0xb0, 0x62,
	0xad, 0x2d, 0xf9, 0x46, 0x80, 0xfe, 0xcf, 0xd2, 0x03, 0x18, 0x3b, 0x59, 0xa1, 0x75, 0x79, 0xa5,
	0xb9, 0xc5, 0x32, 0xe4, 0x6f, 0xb2, 0xd1, 0x05, 0x38, 0xc3, 0x92, 0xc6, 0x30, 0x2c, 0xb5, 0xe7,
	0x0d, 0xc6, 0x43, 0x03, 0x92, 0x8e, 0xd9, 0xa0, 0xd4, 0xfe, 0x44, 0x56, 0x78, 0x6c, 0xe9, 0x3d,
	0x18, 0x18, 0xcc, 0x4f, 0xf9, 0x1c, 0x17, 0x76, 0xb2, 0x11, 0xd8, 0xad, 0x06, 0x78, 0x8b, 0x0b,
	0x4b, 0xef, 0x03, 0x7c, 0x36, 0xd2, 0x61, 0xcb, 0x6e, 0xb6, 0xde, 0x80, 0x34, 0x74, 0x66, 0x80,
	0xae, 0x5d, 0xff, 0x8b, 0xe6, 0x1c, 0xf4, 0x03, 0xf4, 0x19, 0x6a, 0x65, 0x1c, 0x7d, 0xd4, 0x7d,
	0xd0, 0x8e, 0x27, 0x8b, 0x0e, 0xba, 0xa5, 0xeb, 0xcf, 0xdd, 0x4b, 0x49, 0xf6, 0x95, 0xc0, 0xed,
	0xcb, 0xfb, 0xf1, 0xce, 0x17, 0x33, 0x5f, 0x50, 0x09, 0x83, 0x99, 0x2f, 0x6c, 0x69, 0x64, 0x81,
	0xf4, 0xc9, 0xd5, 0xd9, 0x97, 0x56, 0x2b, 0xba, 0x7e, 0xd5, 0xa4, 0xf7, 0x8c, 0x1c, 0x3d, 0xfc,
	0xfd, 0x63, 0x8b, 0xfc, 0x5c, 0xc6, 0xe4, 0xd7, 0x32, 0x26, 0x7f, 0x96, 0x31, 0xf9, 0xfe, 0x37,
	0xee, 0xc1, 0xae, 0x32, 0x62, 0xea, 0xe4, 0xfc, 0x6c, 0x3a, 0x3f, 0x0b, 0xcb, 0x5c, 0xf4, 0xc3,
	0xe7, 0xf9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0xb8, 0xf4, 0x12, 0x2e, 0x03, 0x00, 0x00,
}
