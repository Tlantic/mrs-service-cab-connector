// Code generated by protoc-gen-go. DO NOT EDIT.
// source: price_types.proto

package proto

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

// Request
type GetPriceRequest struct {
	Store                string   `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
	Sku                  string   `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Ean                  string   `protobuf:"bytes,3,opt,name=ean,proto3" json:"ean,omitempty"`
	ExternalCode         string   `protobuf:"bytes,3,opt,name=ean,proto3" json:"ean,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPriceRequest) Reset()         { *m = GetPriceRequest{} }
func (m *GetPriceRequest) String() string { return proto.CompactTextString(m) }
func (*GetPriceRequest) ProtoMessage()    {}
func (*GetPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c610788ab1ad1e96, []int{0}
}

func (m *GetPriceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPriceRequest.Unmarshal(m, b)
}
func (m *GetPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPriceRequest.Marshal(b, m, deterministic)
}
func (m *GetPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPriceRequest.Merge(m, src)
}
func (m *GetPriceRequest) XXX_Size() int {
	return xxx_messageInfo_GetPriceRequest.Size(m)
}
func (m *GetPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPriceRequest proto.InternalMessageInfo

func (m *GetPriceRequest) GetStore() string {
	if m != nil {
		return m.Store
	}
	return ""
}

func (m *GetPriceRequest) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *GetPriceRequest) GetEan() string {
	if m != nil {
		return m.Ean
	}
	return ""
}

// Response
type PriceResponse struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Sku                  string   `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Ean                  string   `protobuf:"bytes,3,opt,name=ean,proto3" json:"ean,omitempty"`
	Store                string   `protobuf:"bytes,4,opt,name=store,proto3" json:"store,omitempty"`
	Value                string   `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	PromotionValue       string   `protobuf:"bytes,6,opt,name=promotion_value,json=promotionValue,proto3" json:"promotion_value,omitempty"`
	GrossMargin          string   `protobuf:"bytes,7,opt,name=gross_margin,json=grossMargin,proto3" json:"gross_margin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PriceResponse) Reset()         { *m = PriceResponse{} }
func (m *PriceResponse) String() string { return proto.CompactTextString(m) }
func (*PriceResponse) ProtoMessage()    {}
func (*PriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c610788ab1ad1e96, []int{1}
}

func (m *PriceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriceResponse.Unmarshal(m, b)
}
func (m *PriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriceResponse.Marshal(b, m, deterministic)
}
func (m *PriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceResponse.Merge(m, src)
}
func (m *PriceResponse) XXX_Size() int {
	return xxx_messageInfo_PriceResponse.Size(m)
}
func (m *PriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PriceResponse proto.InternalMessageInfo

func (m *PriceResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PriceResponse) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *PriceResponse) GetEan() string {
	if m != nil {
		return m.Ean
	}
	return ""
}

func (m *PriceResponse) GetStore() string {
	if m != nil {
		return m.Store
	}
	return ""
}

func (m *PriceResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PriceResponse) GetPromotionValue() string {
	if m != nil {
		return m.PromotionValue
	}
	return ""
}

func (m *PriceResponse) GetGrossMargin() string {
	if m != nil {
		return m.GrossMargin
	}
	return ""
}

type GetPriceResult struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool           `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message              string         `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Result               *PriceResponse `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetPriceResult) Reset()         { *m = GetPriceResult{} }
func (m *GetPriceResult) String() string { return proto.CompactTextString(m) }
func (*GetPriceResult) ProtoMessage()    {}
func (*GetPriceResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c610788ab1ad1e96, []int{2}
}

func (m *GetPriceResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPriceResult.Unmarshal(m, b)
}
func (m *GetPriceResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPriceResult.Marshal(b, m, deterministic)
}
func (m *GetPriceResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPriceResult.Merge(m, src)
}
func (m *GetPriceResult) XXX_Size() int {
	return xxx_messageInfo_GetPriceResult.Size(m)
}
func (m *GetPriceResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPriceResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetPriceResult proto.InternalMessageInfo

func (m *GetPriceResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPriceResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetPriceResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetPriceResult) GetResult() *PriceResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

// Request
type GetPricePOSRequest struct {
	Store                string   `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
	Sku                  string   `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Ean                  string   `protobuf:"bytes,3,opt,name=ean,proto3" json:"ean,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPricePOSRequest) Reset()         { *m = GetPricePOSRequest{} }
func (m *GetPricePOSRequest) String() string { return proto.CompactTextString(m) }
func (*GetPricePOSRequest) ProtoMessage()    {}
func (*GetPricePOSRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c610788ab1ad1e96, []int{3}
}

func (m *GetPricePOSRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricePOSRequest.Unmarshal(m, b)
}
func (m *GetPricePOSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricePOSRequest.Marshal(b, m, deterministic)
}
func (m *GetPricePOSRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricePOSRequest.Merge(m, src)
}
func (m *GetPricePOSRequest) XXX_Size() int {
	return xxx_messageInfo_GetPricePOSRequest.Size(m)
}
func (m *GetPricePOSRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricePOSRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricePOSRequest proto.InternalMessageInfo

func (m *GetPricePOSRequest) GetStore() string {
	if m != nil {
		return m.Store
	}
	return ""
}

func (m *GetPricePOSRequest) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *GetPricePOSRequest) GetEan() string {
	if m != nil {
		return m.Ean
	}
	return ""
}

type GetPricePOSResult struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool           `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message              string         `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Result               *PriceResponse `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetPricePOSResult) Reset()         { *m = GetPricePOSResult{} }
func (m *GetPricePOSResult) String() string { return proto.CompactTextString(m) }
func (*GetPricePOSResult) ProtoMessage()    {}
func (*GetPricePOSResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c610788ab1ad1e96, []int{4}
}

func (m *GetPricePOSResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricePOSResult.Unmarshal(m, b)
}
func (m *GetPricePOSResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricePOSResult.Marshal(b, m, deterministic)
}
func (m *GetPricePOSResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricePOSResult.Merge(m, src)
}
func (m *GetPricePOSResult) XXX_Size() int {
	return xxx_messageInfo_GetPricePOSResult.Size(m)
}
func (m *GetPricePOSResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricePOSResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricePOSResult proto.InternalMessageInfo

func (m *GetPricePOSResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPricePOSResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetPricePOSResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetPricePOSResult) GetResult() *PriceResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

// Request
type GetPricesRequest struct {
	Store                string   `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
	Sku                  string   `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Ean                  string   `protobuf:"bytes,3,opt,name=ean,proto3" json:"ean,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPricesRequest) Reset()         { *m = GetPricesRequest{} }
func (m *GetPricesRequest) String() string { return proto.CompactTextString(m) }
func (*GetPricesRequest) ProtoMessage()    {}
func (*GetPricesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c610788ab1ad1e96, []int{5}
}

func (m *GetPricesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricesRequest.Unmarshal(m, b)
}
func (m *GetPricesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricesRequest.Marshal(b, m, deterministic)
}
func (m *GetPricesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricesRequest.Merge(m, src)
}
func (m *GetPricesRequest) XXX_Size() int {
	return xxx_messageInfo_GetPricesRequest.Size(m)
}
func (m *GetPricesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricesRequest proto.InternalMessageInfo

func (m *GetPricesRequest) GetStore() string {
	if m != nil {
		return m.Store
	}
	return ""
}

func (m *GetPricesRequest) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *GetPricesRequest) GetEan() string {
	if m != nil {
		return m.Ean
	}
	return ""
}

// Response
type PricesResponse struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Sku                  string   `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Ean                  string   `protobuf:"bytes,3,opt,name=ean,proto3" json:"ean,omitempty"`
	Store                string   `protobuf:"bytes,4,opt,name=store,proto3" json:"store,omitempty"`
	Value                string   `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	PromotionValue       string   `protobuf:"bytes,6,opt,name=promotion_value,json=promotionValue,proto3" json:"promotion_value,omitempty"`
	GrossMargin          string   `protobuf:"bytes,7,opt,name=gross_margin,json=grossMargin,proto3" json:"gross_margin,omitempty"`
	Group                string   `protobuf:"bytes,8,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PricesResponse) Reset()         { *m = PricesResponse{} }
func (m *PricesResponse) String() string { return proto.CompactTextString(m) }
func (*PricesResponse) ProtoMessage()    {}
func (*PricesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c610788ab1ad1e96, []int{6}
}

func (m *PricesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PricesResponse.Unmarshal(m, b)
}
func (m *PricesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PricesResponse.Marshal(b, m, deterministic)
}
func (m *PricesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PricesResponse.Merge(m, src)
}
func (m *PricesResponse) XXX_Size() int {
	return xxx_messageInfo_PricesResponse.Size(m)
}
func (m *PricesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PricesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PricesResponse proto.InternalMessageInfo

func (m *PricesResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PricesResponse) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *PricesResponse) GetEan() string {
	if m != nil {
		return m.Ean
	}
	return ""
}

func (m *PricesResponse) GetStore() string {
	if m != nil {
		return m.Store
	}
	return ""
}

func (m *PricesResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PricesResponse) GetPromotionValue() string {
	if m != nil {
		return m.PromotionValue
	}
	return ""
}

func (m *PricesResponse) GetGrossMargin() string {
	if m != nil {
		return m.GrossMargin
	}
	return ""
}

func (m *PricesResponse) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

type GetPricesResult struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool              `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message              string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Result               []*PricesResponse `protobuf:"bytes,4,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetPricesResult) Reset()         { *m = GetPricesResult{} }
func (m *GetPricesResult) String() string { return proto.CompactTextString(m) }
func (*GetPricesResult) ProtoMessage()    {}
func (*GetPricesResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c610788ab1ad1e96, []int{7}
}

func (m *GetPricesResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricesResult.Unmarshal(m, b)
}
func (m *GetPricesResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricesResult.Marshal(b, m, deterministic)
}
func (m *GetPricesResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricesResult.Merge(m, src)
}
func (m *GetPricesResult) XXX_Size() int {
	return xxx_messageInfo_GetPricesResult.Size(m)
}
func (m *GetPricesResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricesResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricesResult proto.InternalMessageInfo

func (m *GetPricesResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetPricesResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetPricesResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetPricesResult) GetResult() []*PricesResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*GetPriceRequest)(nil), "proto.GetPriceRequest")
	proto.RegisterType((*PriceResponse)(nil), "proto.PriceResponse")
	proto.RegisterType((*GetPriceResult)(nil), "proto.GetPriceResult")
	proto.RegisterType((*GetPricePOSRequest)(nil), "proto.GetPricePOSRequest")
	proto.RegisterType((*GetPricePOSResult)(nil), "proto.GetPricePOSResult")
	proto.RegisterType((*GetPricesRequest)(nil), "proto.GetPricesRequest")
	proto.RegisterType((*PricesResponse)(nil), "proto.PricesResponse")
	proto.RegisterType((*GetPricesResult)(nil), "proto.GetPricesResult")
}

func init() { proto.RegisterFile("price_types.proto", fileDescriptor_c610788ab1ad1e96) }

var fileDescriptor_c610788ab1ad1e96 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0xcd, 0x02, 0x05, 0x7e, 0xc3, 0xcf, 0x02, 0x1b, 0x4c, 0xf6, 0xa8, 0xbd, 0xc8, 0x41, 0x39,
	0xe0, 0x87, 0xf0, 0xe0, 0x3f, 0x52, 0x13, 0xaf, 0xa4, 0xc2, 0xa4, 0x69, 0x84, 0x6e, 0xdd, 0xd9,
	0x9a, 0x78, 0x35, 0xfa, 0xb9, 0xfc, 0x16, 0x7e, 0x1e, 0xd3, 0xd9, 0x56, 0x8a, 0x27, 0x13, 0x4c,
	0x8c, 0xa7, 0xee, 0xbc, 0x37, 0x79, 0x33, 0x6f, 0x5f, 0x17, 0x86, 0x99, 0x49, 0x16, 0x38, 0xb7,
	0x4f, 0x19, 0xd2, 0x24, 0x33, 0xda, 0x6a, 0xe9, 0xf1, 0x27, 0x38, 0x87, 0xfe, 0x19, 0xda, 0x59,
	0x41, 0x87, 0xf8, 0x90, 0x23, 0x59, 0x39, 0x02, 0x8f, 0xac, 0x36, 0xa8, 0xc4, 0x81, 0x18, 0xff,
	0x0b, 0x5d, 0x21, 0x07, 0xd0, 0xa4, 0xfb, 0x5c, 0x35, 0x18, 0x2b, 0x8e, 0x05, 0x82, 0x51, 0xaa,
	0x9a, 0x0e, 0xc1, 0x28, 0x0d, 0xde, 0x04, 0xec, 0x95, 0x52, 0x94, 0xe9, 0x94, 0x50, 0x4a, 0x68,
	0x15, 0x43, 0x4b, 0x29, 0x3e, 0x7f, 0x47, 0x69, 0xb3, 0x43, 0xab, 0xbe, 0xc3, 0x08, 0xbc, 0xc7,
	0x68, 0x95, 0xa3, 0xf2, 0x1c, 0xca, 0x85, 0x3c, 0x82, 0x7e, 0x66, 0xf4, 0x5a, 0xdb, 0x44, 0xa7,
	0x73, 0xc7, 0xb7, 0x99, 0xf7, 0x3f, 0xe1, 0x5b, 0x6e, 0x3c, 0x84, 0xff, 0xb1, 0xd1, 0x44, 0xf3,
	0x75, 0x64, 0xe2, 0x24, 0x55, 0x1d, 0xee, 0xea, 0x31, 0x76, 0xc9, 0x50, 0xf0, 0x2c, 0xc0, 0xdf,
	0xdc, 0x07, 0xe5, 0x2b, 0x2b, 0x7d, 0x68, 0x24, 0xcb, 0xd2, 0x40, 0x23, 0x59, 0x4a, 0x05, 0x1d,
	0xca, 0x17, 0x0b, 0x24, 0x62, 0x0b, 0xdd, 0xb0, 0x2a, 0x0b, 0x66, 0x8d, 0x44, 0x51, 0x8c, 0xa5,
	0x95, 0xaa, 0x94, 0xc7, 0xd0, 0x36, 0xac, 0xc6, 0x7e, 0x7a, 0xd3, 0x91, 0x0b, 0x61, 0xb2, 0x75,
	0x59, 0x61, 0xd9, 0x13, 0x5c, 0x81, 0xac, 0x76, 0x98, 0x5d, 0xdf, 0xec, 0x1e, 0xcb, 0xab, 0x80,
	0xe1, 0x96, 0xe0, 0x2f, 0xf9, 0xba, 0x80, 0x41, 0xb5, 0x06, 0xed, 0xee, 0xea, 0x5d, 0x80, 0x5f,
	0x69, 0xfd, 0xcd, 0xbf, 0xad, 0x98, 0x10, 0x1b, 0x9d, 0x67, 0xaa, 0xeb, 0x26, 0x70, 0x11, 0xbc,
	0x88, 0xcd, 0x9b, 0xa4, 0x1f, 0x0c, 0xeb, 0xa4, 0x16, 0x56, 0x73, 0xdc, 0x9b, 0xee, 0xd7, 0xc3,
	0xa2, 0xaf, 0x69, 0xdd, 0xb5, 0x99, 0x3d, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x7f, 0x44,
	0xa6, 0x3c, 0x04, 0x00, 0x00,
}
