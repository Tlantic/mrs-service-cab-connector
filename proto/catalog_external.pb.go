// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog_external.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type GetDamagesDestinationsRequest struct {
	StoreId              string   `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDamagesDestinationsRequest) Reset()         { *m = GetDamagesDestinationsRequest{} }
func (m *GetDamagesDestinationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetDamagesDestinationsRequest) ProtoMessage()    {}
func (*GetDamagesDestinationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e092cd7f7404ca7, []int{0}
}

func (m *GetDamagesDestinationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDamagesDestinationsRequest.Unmarshal(m, b)
}
func (m *GetDamagesDestinationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDamagesDestinationsRequest.Marshal(b, m, deterministic)
}
func (m *GetDamagesDestinationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDamagesDestinationsRequest.Merge(m, src)
}
func (m *GetDamagesDestinationsRequest) XXX_Size() int {
	return xxx_messageInfo_GetDamagesDestinationsRequest.Size(m)
}
func (m *GetDamagesDestinationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDamagesDestinationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDamagesDestinationsRequest proto.InternalMessageInfo

func (m *GetDamagesDestinationsRequest) GetStoreId() string {
	if m != nil {
		return m.StoreId
	}
	return ""
}

type GetDamagesDestinationsResult struct {
	Id                   string                            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool                              `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message              string                            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Result               []*GetDamagesDestinationsResponse `protobuf:"bytes,4,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *GetDamagesDestinationsResult) Reset()         { *m = GetDamagesDestinationsResult{} }
func (m *GetDamagesDestinationsResult) String() string { return proto.CompactTextString(m) }
func (*GetDamagesDestinationsResult) ProtoMessage()    {}
func (*GetDamagesDestinationsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e092cd7f7404ca7, []int{1}
}

func (m *GetDamagesDestinationsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDamagesDestinationsResult.Unmarshal(m, b)
}
func (m *GetDamagesDestinationsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDamagesDestinationsResult.Marshal(b, m, deterministic)
}
func (m *GetDamagesDestinationsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDamagesDestinationsResult.Merge(m, src)
}
func (m *GetDamagesDestinationsResult) XXX_Size() int {
	return xxx_messageInfo_GetDamagesDestinationsResult.Size(m)
}
func (m *GetDamagesDestinationsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDamagesDestinationsResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetDamagesDestinationsResult proto.InternalMessageInfo

func (m *GetDamagesDestinationsResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetDamagesDestinationsResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetDamagesDestinationsResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetDamagesDestinationsResult) GetResult() []*GetDamagesDestinationsResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

type GetDamagesDestinationsResponse struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	OrgCode              string   `protobuf:"bytes,5,opt,name=org_code,json=orgCode,proto3" json:"org_code,omitempty"`
	Store                string   `protobuf:"bytes,6,opt,name=store,proto3" json:"store,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDamagesDestinationsResponse) Reset()         { *m = GetDamagesDestinationsResponse{} }
func (m *GetDamagesDestinationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetDamagesDestinationsResponse) ProtoMessage()    {}
func (*GetDamagesDestinationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e092cd7f7404ca7, []int{2}
}

func (m *GetDamagesDestinationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDamagesDestinationsResponse.Unmarshal(m, b)
}
func (m *GetDamagesDestinationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDamagesDestinationsResponse.Marshal(b, m, deterministic)
}
func (m *GetDamagesDestinationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDamagesDestinationsResponse.Merge(m, src)
}
func (m *GetDamagesDestinationsResponse) XXX_Size() int {
	return xxx_messageInfo_GetDamagesDestinationsResponse.Size(m)
}
func (m *GetDamagesDestinationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDamagesDestinationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDamagesDestinationsResponse proto.InternalMessageInfo

func (m *GetDamagesDestinationsResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *GetDamagesDestinationsResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetDamagesDestinationsResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetDamagesDestinationsResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetDamagesDestinationsResponse) GetOrgCode() string {
	if m != nil {
		return m.OrgCode
	}
	return ""
}

func (m *GetDamagesDestinationsResponse) GetStore() string {
	if m != nil {
		return m.Store
	}
	return ""
}

type GetDamageReasonsRequest struct {
	StoreId              string   `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDamageReasonsRequest) Reset()         { *m = GetDamageReasonsRequest{} }
func (m *GetDamageReasonsRequest) String() string { return proto.CompactTextString(m) }
func (*GetDamageReasonsRequest) ProtoMessage()    {}
func (*GetDamageReasonsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e092cd7f7404ca7, []int{3}
}

func (m *GetDamageReasonsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDamageReasonsRequest.Unmarshal(m, b)
}
func (m *GetDamageReasonsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDamageReasonsRequest.Marshal(b, m, deterministic)
}
func (m *GetDamageReasonsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDamageReasonsRequest.Merge(m, src)
}
func (m *GetDamageReasonsRequest) XXX_Size() int {
	return xxx_messageInfo_GetDamageReasonsRequest.Size(m)
}
func (m *GetDamageReasonsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDamageReasonsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDamageReasonsRequest proto.InternalMessageInfo

func (m *GetDamageReasonsRequest) GetStoreId() string {
	if m != nil {
		return m.StoreId
	}
	return ""
}

type GetDamageReasonsResult struct {
	Id                   string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool                        `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message              string                      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Result               []*GetDamageReasonsResponse `protobuf:"bytes,4,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *GetDamageReasonsResult) Reset()         { *m = GetDamageReasonsResult{} }
func (m *GetDamageReasonsResult) String() string { return proto.CompactTextString(m) }
func (*GetDamageReasonsResult) ProtoMessage()    {}
func (*GetDamageReasonsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e092cd7f7404ca7, []int{4}
}

func (m *GetDamageReasonsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDamageReasonsResult.Unmarshal(m, b)
}
func (m *GetDamageReasonsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDamageReasonsResult.Marshal(b, m, deterministic)
}
func (m *GetDamageReasonsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDamageReasonsResult.Merge(m, src)
}
func (m *GetDamageReasonsResult) XXX_Size() int {
	return xxx_messageInfo_GetDamageReasonsResult.Size(m)
}
func (m *GetDamageReasonsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDamageReasonsResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetDamageReasonsResult proto.InternalMessageInfo

func (m *GetDamageReasonsResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetDamageReasonsResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetDamageReasonsResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetDamageReasonsResult) GetResult() []*GetDamageReasonsResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

type GetDamageReasonsResponse struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	OrgCode              string   `protobuf:"bytes,5,opt,name=org_code,json=orgCode,proto3" json:"org_code,omitempty"`
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDamageReasonsResponse) Reset()         { *m = GetDamageReasonsResponse{} }
func (m *GetDamageReasonsResponse) String() string { return proto.CompactTextString(m) }
func (*GetDamageReasonsResponse) ProtoMessage()    {}
func (*GetDamageReasonsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e092cd7f7404ca7, []int{5}
}

func (m *GetDamageReasonsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDamageReasonsResponse.Unmarshal(m, b)
}
func (m *GetDamageReasonsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDamageReasonsResponse.Marshal(b, m, deterministic)
}
func (m *GetDamageReasonsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDamageReasonsResponse.Merge(m, src)
}
func (m *GetDamageReasonsResponse) XXX_Size() int {
	return xxx_messageInfo_GetDamageReasonsResponse.Size(m)
}
func (m *GetDamageReasonsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDamageReasonsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDamageReasonsResponse proto.InternalMessageInfo

func (m *GetDamageReasonsResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *GetDamageReasonsResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetDamageReasonsResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetDamageReasonsResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetDamageReasonsResponse) GetOrgCode() string {
	if m != nil {
		return m.OrgCode
	}
	return ""
}

func (m *GetDamageReasonsResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDamagesDestinationsRequest)(nil), "proto.GetDamagesDestinationsRequest")
	proto.RegisterType((*GetDamagesDestinationsResult)(nil), "proto.GetDamagesDestinationsResult")
	proto.RegisterType((*GetDamagesDestinationsResponse)(nil), "proto.GetDamagesDestinationsResponse")
	proto.RegisterType((*GetDamageReasonsRequest)(nil), "proto.GetDamageReasonsRequest")
	proto.RegisterType((*GetDamageReasonsResult)(nil), "proto.GetDamageReasonsResult")
	proto.RegisterType((*GetDamageReasonsResponse)(nil), "proto.GetDamageReasonsResponse")
}

func init() { proto.RegisterFile("catalog_external.proto", fileDescriptor_0e092cd7f7404ca7) }

var fileDescriptor_0e092cd7f7404ca7 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xf3, 0xd7, 0x64, 0x22, 0x55, 0x74, 0x05, 0xa9, 0x9b, 0xfe, 0x10, 0x05, 0x90, 0x72,
	0xca, 0xa1, 0x20, 0x21, 0x81, 0xe0, 0x92, 0xa0, 0x0a, 0x89, 0xaa, 0x91, 0xcd, 0x3d, 0xda, 0x3a,
	0xa3, 0xc8, 0xaa, 0xeb, 0x0d, 0x3b, 0x6b, 0x09, 0x1e, 0x84, 0x07, 0xe0, 0xc4, 0x91, 0x07, 0xe4,
	0x82, 0xbc, 0xbb, 0x4e, 0xec, 0xc4, 0xa6, 0x1c, 0x90, 0x7a, 0x8a, 0xf7, 0x9b, 0xf9, 0x66, 0xe6,
	0x9b, 0xc9, 0x0c, 0xf4, 0x02, 0xae, 0x78, 0x24, 0x96, 0x73, 0xfc, 0xaa, 0x50, 0xc6, 0x3c, 0x1a,
	0xaf, 0xa4, 0x50, 0x82, 0x35, 0xf5, 0x4f, 0xff, 0x90, 0x94, 0x08, 0x6e, 0xe7, 0xea, 0xdb, 0x0a,
	0x69, 0x6c, 0xa1, 0x95, 0x0c, 0x03, 0x2c, 0x42, 0x11, 0xbf, 0xc1, 0x28, 0x0f, 0x0d, 0xdf, 0xc0,
	0xd9, 0x25, 0xaa, 0x29, 0xbf, 0xe3, 0x4b, 0xa4, 0x29, 0x92, 0x0a, 0x63, 0xae, 0x42, 0x11, 0x93,
	0x87, 0x5f, 0x12, 0x24, 0xc5, 0x8e, 0xa1, 0x4d, 0x4a, 0x48, 0x9c, 0x87, 0x0b, 0xd7, 0x19, 0x38,
	0xa3, 0x8e, 0xb7, 0xaf, 0xdf, 0x1f, 0x17, 0xc3, 0x1f, 0x0e, 0x9c, 0x56, 0x91, 0x29, 0x89, 0x14,
	0x3b, 0x80, 0xda, 0x9a, 0x55, 0x0b, 0x17, 0xcc, 0x85, 0x7d, 0x4a, 0x82, 0x00, 0x89, 0xdc, 0xda,
	0xc0, 0x19, 0xb5, 0xbd, 0xec, 0x99, 0x5a, 0xee, 0x90, 0x88, 0x2f, 0xd1, 0xad, 0x9b, 0x24, 0xf6,
	0xc9, 0xde, 0x41, 0x4b, 0xea, 0x68, 0x6e, 0x63, 0x50, 0x1f, 0x75, 0x2f, 0x5e, 0x98, 0xc2, 0xc7,
	0x95, 0x89, 0x57, 0x22, 0x26, 0xf4, 0x2c, 0x69, 0xf8, 0xcb, 0x81, 0xf3, 0xbf, 0xbb, 0x32, 0x06,
	0x8d, 0x40, 0x2c, 0xd0, 0xd6, 0xa9, 0xbf, 0xd9, 0x00, 0xba, 0x0b, 0xa4, 0x40, 0x86, 0xab, 0xd4,
	0x57, 0x57, 0xdb, 0xf1, 0xf2, 0x90, 0xd5, 0x56, 0x5f, 0x6b, 0xeb, 0x41, 0x8b, 0x14, 0x57, 0x09,
	0xb9, 0x0d, 0x8d, 0xd9, 0x57, 0xda, 0x3f, 0x21, 0x97, 0x73, 0x9d, 0xa1, 0x69, 0xa4, 0x09, 0xb9,
	0x9c, 0xa4, 0x49, 0x1e, 0x43, 0x53, 0xb7, 0xd2, 0x6d, 0x69, 0xdc, 0x3c, 0x86, 0xaf, 0xe0, 0x68,
	0x5d, 0xb0, 0x87, 0x9c, 0xfe, 0x6d, 0x16, 0xdf, 0x1d, 0xe8, 0xed, 0xd2, 0xfe, 0xdb, 0x14, 0x5e,
	0x6f, 0x4d, 0xe1, 0xe9, 0xf6, 0x14, 0x36, 0x29, 0x8b, 0xfd, 0xff, 0xe9, 0x80, 0x5b, 0xe5, 0xf4,
	0x70, 0x9d, 0x67, 0xd0, 0x48, 0x97, 0xc0, 0x36, 0x5e, 0x7f, 0x5f, 0xfc, 0x6e, 0xc2, 0xc9, 0x95,
	0xe7, 0x4f, 0xcc, 0x9e, 0x4d, 0x44, 0x1c, 0x63, 0xa0, 0x84, 0xfc, 0x60, 0xf7, 0x8d, 0xbd, 0x85,
	0xf6, 0x25, 0x2a, 0x3f, 0xdd, 0x33, 0xd6, 0xdb, 0xc8, 0xd7, 0x80, 0x1d, 0x50, 0xff, 0xc9, 0x0e,
	0xae, 0x9b, 0xb0, 0x67, 0xc9, 0xb3, 0x74, 0x23, 0xf3, 0x64, 0x0d, 0x94, 0x90, 0x2d, 0x6e, 0xc9,
	0x53, 0xe8, 0x66, 0xd8, 0xec, 0xda, 0x67, 0xc7, 0x5b, 0x7e, 0xb3, 0x6b, 0x3f, 0x0b, 0xe1, 0x96,
	0x99, 0x6c, 0x94, 0xf7, 0xd0, 0xc9, 0x60, 0x62, 0x47, 0x5b, 0x8e, 0xd9, 0x5f, 0xac, 0xdf, 0xdb,
	0x35, 0x14, 0x24, 0x7c, 0x4a, 0x2f, 0x48, 0x5e, 0x82, 0x06, 0x4a, 0x24, 0x58, 0xdc, 0x92, 0xaf,
	0xe0, 0xc0, 0xf4, 0x44, 0x22, 0x99, 0x16, 0x9e, 0x16, 0x5a, 0x95, 0xc1, 0x59, 0xa0, 0x93, 0x0a,
	0x6b, 0x41, 0x8b, 0xc6, 0x0a, 0x5a, 0x0c, 0x52, 0xa2, 0x25, 0x33, 0x58, 0xbe, 0x0f, 0x8f, 0x52,
	0xf0, 0x36, 0xd1, 0xc1, 0x7d, 0x1e, 0x21, 0xb1, 0xf3, 0x9c, 0x77, 0xde, 0x90, 0x45, 0x3b, 0xab,
	0xb4, 0xdb, 0xa0, 0x98, 0xdb, 0xc0, 0xc2, 0xa5, 0x61, 0xcf, 0xef, 0xb9, 0x59, 0x26, 0xc1, 0xb3,
	0xfb, 0x2e, 0x9b, 0x49, 0xf3, 0x19, 0x0e, 0x37, 0x1e, 0x76, 0xa3, 0xf2, 0xc5, 0x97, 0x5d, 0x8e,
	0x7c, 0xf1, 0x25, 0x27, 0x62, 0xb8, 0x77, 0xd3, 0xd2, 0xf6, 0x97, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xc0, 0xf3, 0xff, 0xbd, 0x68, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MRSCatalogConnectorExternalClient is the client API for MRSCatalogConnectorExternal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MRSCatalogConnectorExternalClient interface {
	GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*GetStockResult, error)
	GetPrice(ctx context.Context, in *GetPriceRequest, opts ...grpc.CallOption) (*GetPriceResult, error)
	GetPricePOS(ctx context.Context, in *GetPricePOSRequest, opts ...grpc.CallOption) (*GetPricePOSResult, error)
	GetPrices(ctx context.Context, in *GetPricesRequest, opts ...grpc.CallOption) (*GetPricesResult, error)
	GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*GetLabelResult, error)
	GetStoresStock(ctx context.Context, in *GetStoresStockRequest, opts ...grpc.CallOption) (*GetStoresStockResult, error)
	GetStocks(ctx context.Context, in *GetStocksRequest, opts ...grpc.CallOption) (*GetStocksResult, error)
	GetSkuStoreSales(ctx context.Context, in *GetSkuStoreSalesRequest, opts ...grpc.CallOption) (*GetSkuStoreSalesResult, error)
	GetDamagesDestinations(ctx context.Context, in *GetDamagesDestinationsRequest, opts ...grpc.CallOption) (*GetDamagesDestinationsResult, error)
	GetDamagesReasons(ctx context.Context, in *GetDamageReasonsRequest, opts ...grpc.CallOption) (*GetDamageReasonsResult, error)
}

type mRSCatalogConnectorExternalClient struct {
	cc *grpc.ClientConn
}

func NewMRSCatalogConnectorExternalClient(cc *grpc.ClientConn) MRSCatalogConnectorExternalClient {
	return &mRSCatalogConnectorExternalClient{cc}
}

func (c *mRSCatalogConnectorExternalClient) GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*GetStockResult, error) {
	out := new(GetStockResult)
	err := c.cc.Invoke(ctx, "/proto.MRSCatalogConnectorExternal/GetStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRSCatalogConnectorExternalClient) GetPrice(ctx context.Context, in *GetPriceRequest, opts ...grpc.CallOption) (*GetPriceResult, error) {
	out := new(GetPriceResult)
	err := c.cc.Invoke(ctx, "/proto.MRSCatalogConnectorExternal/GetPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRSCatalogConnectorExternalClient) GetPricePOS(ctx context.Context, in *GetPricePOSRequest, opts ...grpc.CallOption) (*GetPricePOSResult, error) {
	out := new(GetPricePOSResult)
	err := c.cc.Invoke(ctx, "/proto.MRSCatalogConnectorExternal/GetPricePOS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRSCatalogConnectorExternalClient) GetPrices(ctx context.Context, in *GetPricesRequest, opts ...grpc.CallOption) (*GetPricesResult, error) {
	out := new(GetPricesResult)
	err := c.cc.Invoke(ctx, "/proto.MRSCatalogConnectorExternal/GetPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRSCatalogConnectorExternalClient) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*GetLabelResult, error) {
	out := new(GetLabelResult)
	err := c.cc.Invoke(ctx, "/proto.MRSCatalogConnectorExternal/GetLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRSCatalogConnectorExternalClient) GetStoresStock(ctx context.Context, in *GetStoresStockRequest, opts ...grpc.CallOption) (*GetStoresStockResult, error) {
	out := new(GetStoresStockResult)
	err := c.cc.Invoke(ctx, "/proto.MRSCatalogConnectorExternal/GetStoresStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRSCatalogConnectorExternalClient) GetStocks(ctx context.Context, in *GetStocksRequest, opts ...grpc.CallOption) (*GetStocksResult, error) {
	out := new(GetStocksResult)
	err := c.cc.Invoke(ctx, "/proto.MRSCatalogConnectorExternal/GetStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRSCatalogConnectorExternalClient) GetSkuStoreSales(ctx context.Context, in *GetSkuStoreSalesRequest, opts ...grpc.CallOption) (*GetSkuStoreSalesResult, error) {
	out := new(GetSkuStoreSalesResult)
	err := c.cc.Invoke(ctx, "/proto.MRSCatalogConnectorExternal/GetSkuStoreSales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRSCatalogConnectorExternalClient) GetDamagesDestinations(ctx context.Context, in *GetDamagesDestinationsRequest, opts ...grpc.CallOption) (*GetDamagesDestinationsResult, error) {
	out := new(GetDamagesDestinationsResult)
	err := c.cc.Invoke(ctx, "/proto.MRSCatalogConnectorExternal/GetDamagesDestinations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRSCatalogConnectorExternalClient) GetDamagesReasons(ctx context.Context, in *GetDamageReasonsRequest, opts ...grpc.CallOption) (*GetDamageReasonsResult, error) {
	out := new(GetDamageReasonsResult)
	err := c.cc.Invoke(ctx, "/proto.MRSCatalogConnectorExternal/GetDamagesReasons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MRSCatalogConnectorExternalServer is the server API for MRSCatalogConnectorExternal service.
type MRSCatalogConnectorExternalServer interface {
	GetStock(context.Context, *GetStockRequest) (*GetStockResult, error)
	GetPrice(context.Context, *GetPriceRequest) (*GetPriceResult, error)
	GetPricePOS(context.Context, *GetPricePOSRequest) (*GetPricePOSResult, error)
	GetPrices(context.Context, *GetPricesRequest) (*GetPricesResult, error)
	GetLabel(context.Context, *GetLabelRequest) (*GetLabelResult, error)
	GetStoresStock(context.Context, *GetStoresStockRequest) (*GetStoresStockResult, error)
	GetStocks(context.Context, *GetStocksRequest) (*GetStocksResult, error)
	GetSkuStoreSales(context.Context, *GetSkuStoreSalesRequest) (*GetSkuStoreSalesResult, error)
	GetDamagesDestinations(context.Context, *GetDamagesDestinationsRequest) (*GetDamagesDestinationsResult, error)
	GetDamagesReasons(context.Context, *GetDamageReasonsRequest) (*GetDamageReasonsResult, error)
}

func RegisterMRSCatalogConnectorExternalServer(s *grpc.Server, srv MRSCatalogConnectorExternalServer) {
	s.RegisterService(&_MRSCatalogConnectorExternal_serviceDesc, srv)
}

func _MRSCatalogConnectorExternal_GetStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRSCatalogConnectorExternalServer).GetStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MRSCatalogConnectorExternal/GetStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRSCatalogConnectorExternalServer).GetStock(ctx, req.(*GetStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRSCatalogConnectorExternal_GetPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRSCatalogConnectorExternalServer).GetPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MRSCatalogConnectorExternal/GetPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRSCatalogConnectorExternalServer).GetPrice(ctx, req.(*GetPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRSCatalogConnectorExternal_GetPricePOS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPricePOSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRSCatalogConnectorExternalServer).GetPricePOS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MRSCatalogConnectorExternal/GetPricePOS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRSCatalogConnectorExternalServer).GetPricePOS(ctx, req.(*GetPricePOSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRSCatalogConnectorExternal_GetPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRSCatalogConnectorExternalServer).GetPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MRSCatalogConnectorExternal/GetPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRSCatalogConnectorExternalServer).GetPrices(ctx, req.(*GetPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRSCatalogConnectorExternal_GetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRSCatalogConnectorExternalServer).GetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MRSCatalogConnectorExternal/GetLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRSCatalogConnectorExternalServer).GetLabel(ctx, req.(*GetLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRSCatalogConnectorExternal_GetStoresStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoresStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRSCatalogConnectorExternalServer).GetStoresStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MRSCatalogConnectorExternal/GetStoresStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRSCatalogConnectorExternalServer).GetStoresStock(ctx, req.(*GetStoresStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRSCatalogConnectorExternal_GetStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRSCatalogConnectorExternalServer).GetStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MRSCatalogConnectorExternal/GetStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRSCatalogConnectorExternalServer).GetStocks(ctx, req.(*GetStocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRSCatalogConnectorExternal_GetSkuStoreSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkuStoreSalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRSCatalogConnectorExternalServer).GetSkuStoreSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MRSCatalogConnectorExternal/GetSkuStoreSales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRSCatalogConnectorExternalServer).GetSkuStoreSales(ctx, req.(*GetSkuStoreSalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRSCatalogConnectorExternal_GetDamagesDestinations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDamagesDestinationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRSCatalogConnectorExternalServer).GetDamagesDestinations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MRSCatalogConnectorExternal/GetDamagesDestinations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRSCatalogConnectorExternalServer).GetDamagesDestinations(ctx, req.(*GetDamagesDestinationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRSCatalogConnectorExternal_GetDamagesReasons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDamageReasonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRSCatalogConnectorExternalServer).GetDamagesReasons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MRSCatalogConnectorExternal/GetDamagesReasons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRSCatalogConnectorExternalServer).GetDamagesReasons(ctx, req.(*GetDamageReasonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MRSCatalogConnectorExternal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MRSCatalogConnectorExternal",
	HandlerType: (*MRSCatalogConnectorExternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStock",
			Handler:    _MRSCatalogConnectorExternal_GetStock_Handler,
		},
		{
			MethodName: "GetPrice",
			Handler:    _MRSCatalogConnectorExternal_GetPrice_Handler,
		},
		{
			MethodName: "GetPricePOS",
			Handler:    _MRSCatalogConnectorExternal_GetPricePOS_Handler,
		},
		{
			MethodName: "GetPrices",
			Handler:    _MRSCatalogConnectorExternal_GetPrices_Handler,
		},
		{
			MethodName: "GetLabel",
			Handler:    _MRSCatalogConnectorExternal_GetLabel_Handler,
		},
		{
			MethodName: "GetStoresStock",
			Handler:    _MRSCatalogConnectorExternal_GetStoresStock_Handler,
		},
		{
			MethodName: "GetStocks",
			Handler:    _MRSCatalogConnectorExternal_GetStocks_Handler,
		},
		{
			MethodName: "GetSkuStoreSales",
			Handler:    _MRSCatalogConnectorExternal_GetSkuStoreSales_Handler,
		},
		{
			MethodName: "GetDamagesDestinations",
			Handler:    _MRSCatalogConnectorExternal_GetDamagesDestinations_Handler,
		},
		{
			MethodName: "GetDamagesReasons",
			Handler:    _MRSCatalogConnectorExternal_GetDamagesReasons_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog_external.proto",
}
