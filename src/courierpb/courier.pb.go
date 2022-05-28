// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/courierpb/courier.proto

package courierpb

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

type Parcel struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Weight               float32  `protobuf:"fixed32,4,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parcel) Reset()         { *m = Parcel{} }
func (m *Parcel) String() string { return proto.CompactTextString(m) }
func (*Parcel) ProtoMessage()    {}
func (*Parcel) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dc59718de1fc3ee, []int{0}
}

func (m *Parcel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parcel.Unmarshal(m, b)
}
func (m *Parcel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parcel.Marshal(b, m, deterministic)
}
func (m *Parcel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parcel.Merge(m, src)
}
func (m *Parcel) XXX_Size() int {
	return xxx_messageInfo_Parcel.Size(m)
}
func (m *Parcel) XXX_DiscardUnknown() {
	xxx_messageInfo_Parcel.DiscardUnknown(m)
}

var xxx_messageInfo_Parcel proto.InternalMessageInfo

func (m *Parcel) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Parcel) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Parcel) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Parcel) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type ProcessParcelsRequest struct {
	Date                 string    `protobuf:"bytes,1,opt,name=Date,proto3" json:"Date,omitempty"`
	Parcels              []*Parcel `protobuf:"bytes,2,rep,name=parcels,proto3" json:"parcels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ProcessParcelsRequest) Reset()         { *m = ProcessParcelsRequest{} }
func (m *ProcessParcelsRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessParcelsRequest) ProtoMessage()    {}
func (*ProcessParcelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dc59718de1fc3ee, []int{1}
}

func (m *ProcessParcelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessParcelsRequest.Unmarshal(m, b)
}
func (m *ProcessParcelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessParcelsRequest.Marshal(b, m, deterministic)
}
func (m *ProcessParcelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessParcelsRequest.Merge(m, src)
}
func (m *ProcessParcelsRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessParcelsRequest.Size(m)
}
func (m *ProcessParcelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessParcelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessParcelsRequest proto.InternalMessageInfo

func (m *ProcessParcelsRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *ProcessParcelsRequest) GetParcels() []*Parcel {
	if m != nil {
		return m.Parcels
	}
	return nil
}

type ProcessParcelsResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessParcelsResponse) Reset()         { *m = ProcessParcelsResponse{} }
func (m *ProcessParcelsResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessParcelsResponse) ProtoMessage()    {}
func (*ProcessParcelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dc59718de1fc3ee, []int{2}
}

func (m *ProcessParcelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessParcelsResponse.Unmarshal(m, b)
}
func (m *ProcessParcelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessParcelsResponse.Marshal(b, m, deterministic)
}
func (m *ProcessParcelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessParcelsResponse.Merge(m, src)
}
func (m *ProcessParcelsResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessParcelsResponse.Size(m)
}
func (m *ProcessParcelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessParcelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessParcelsResponse proto.InternalMessageInfo

func (m *ProcessParcelsResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Parcel)(nil), "courierpb.Parcel")
	proto.RegisterType((*ProcessParcelsRequest)(nil), "courierpb.ProcessParcelsRequest")
	proto.RegisterType((*ProcessParcelsResponse)(nil), "courierpb.ProcessParcelsResponse")
}

func init() {
	proto.RegisterFile("src/courierpb/courier.proto", fileDescriptor_0dc59718de1fc3ee)
}

var fileDescriptor_0dc59718de1fc3ee = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x4d, 0x3a, 0x3b, 0x7a, 0x85, 0x82, 0x17, 0x1d, 0x41, 0x5f, 0x62, 0x9f, 0x02, 0x42,
	0x95, 0xfa, 0x13, 0xf4, 0x55, 0x18, 0xf1, 0x45, 0xd0, 0x97, 0x2e, 0xbb, 0x6c, 0x81, 0x6d, 0x89,
	0x49, 0xa7, 0x7f, 0x5f, 0x6c, 0x6c, 0x51, 0x91, 0xbd, 0xdd, 0x73, 0x73, 0xf2, 0x9d, 0xc3, 0x85,
	0xcb, 0x18, 0xcc, 0x8d, 0x71, 0xfb, 0x60, 0x29, 0xf8, 0xc5, 0x30, 0xd5, 0x3e, 0xb8, 0xce, 0x61,
	0x31, 0x3e, 0x54, 0xaf, 0x90, 0xcf, 0xdb, 0x60, 0x68, 0x83, 0x25, 0x70, 0xbb, 0x14, 0x4c, 0x32,
	0x95, 0x69, 0x6e, 0x97, 0x78, 0x06, 0xc7, 0xb4, 0x6d, 0xed, 0x46, 0x70, 0xc9, 0x54, 0xa1, 0x93,
	0xf8, 0xda, 0xfa, 0xb5, 0xdb, 0x91, 0xc8, 0xd2, 0xb6, 0x17, 0x38, 0x83, 0xfc, 0x83, 0xec, 0x6a,
	0xdd, 0x89, 0x89, 0x64, 0x8a, 0xeb, 0x6f, 0x55, 0x3d, 0xc3, 0xf9, 0x3c, 0x38, 0x43, 0x31, 0xa6,
	0x90, 0xa8, 0xe9, 0x6d, 0x4f, 0xb1, 0x43, 0x84, 0xc9, 0x43, 0xdb, 0x51, 0x1f, 0x57, 0xe8, 0x7e,
	0xc6, 0x6b, 0x98, 0xfa, 0xe4, 0x12, 0x5c, 0x66, 0xea, 0xa4, 0x39, 0xad, 0xc7, 0x9e, 0x75, 0xfa,
	0xaf, 0x07, 0x47, 0xd5, 0xc0, 0xec, 0x2f, 0x39, 0x7a, 0xb7, 0x8b, 0x84, 0x02, 0xa6, 0x8f, 0x14,
	0x63, 0xbb, 0x1a, 0xe8, 0x83, 0x6c, 0xb6, 0x50, 0xde, 0x27, 0xe0, 0x13, 0x85, 0x77, 0x6b, 0x08,
	0x5f, 0xa0, 0xfc, 0x4d, 0x41, 0xf9, 0x33, 0xf3, 0xbf, 0xea, 0x17, 0x57, 0x07, 0x1c, 0xa9, 0x42,
	0x75, 0xa4, 0xd8, 0x2d, 0x5b, 0xe4, 0xfd, 0xb1, 0xef, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0b,
	0xb1, 0x0e, 0x4a, 0x8b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CourierServiceClient is the client API for CourierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CourierServiceClient interface {
	ProcessParcels(ctx context.Context, opts ...grpc.CallOption) (CourierService_ProcessParcelsClient, error)
}

type courierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourierServiceClient(cc grpc.ClientConnInterface) CourierServiceClient {
	return &courierServiceClient{cc}
}

func (c *courierServiceClient) ProcessParcels(ctx context.Context, opts ...grpc.CallOption) (CourierService_ProcessParcelsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CourierService_serviceDesc.Streams[0], "/courierpb.CourierService/ProcessParcels", opts...)
	if err != nil {
		return nil, err
	}
	x := &courierServiceProcessParcelsClient{stream}
	return x, nil
}

type CourierService_ProcessParcelsClient interface {
	Send(*ProcessParcelsRequest) error
	Recv() (*ProcessParcelsResponse, error)
	grpc.ClientStream
}

type courierServiceProcessParcelsClient struct {
	grpc.ClientStream
}

func (x *courierServiceProcessParcelsClient) Send(m *ProcessParcelsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *courierServiceProcessParcelsClient) Recv() (*ProcessParcelsResponse, error) {
	m := new(ProcessParcelsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CourierServiceServer is the server API for CourierService service.
type CourierServiceServer interface {
	ProcessParcels(CourierService_ProcessParcelsServer) error
}

// UnimplementedCourierServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCourierServiceServer struct {
}

func (*UnimplementedCourierServiceServer) ProcessParcels(srv CourierService_ProcessParcelsServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessParcels not implemented")
}

func RegisterCourierServiceServer(s *grpc.Server, srv CourierServiceServer) {
	s.RegisterService(&_CourierService_serviceDesc, srv)
}

func _CourierService_ProcessParcels_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CourierServiceServer).ProcessParcels(&courierServiceProcessParcelsServer{stream})
}

type CourierService_ProcessParcelsServer interface {
	Send(*ProcessParcelsResponse) error
	Recv() (*ProcessParcelsRequest, error)
	grpc.ServerStream
}

type courierServiceProcessParcelsServer struct {
	grpc.ServerStream
}

func (x *courierServiceProcessParcelsServer) Send(m *ProcessParcelsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *courierServiceProcessParcelsServer) Recv() (*ProcessParcelsRequest, error) {
	m := new(ProcessParcelsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CourierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "courierpb.CourierService",
	HandlerType: (*CourierServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessParcels",
			Handler:       _CourierService_ProcessParcels_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "src/courierpb/courier.proto",
}
