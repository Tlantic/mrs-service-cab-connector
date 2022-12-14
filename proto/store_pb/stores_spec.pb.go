// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stores_spec.proto

package store_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Error_Code int32

const (
	Error_OK                        Error_Code = 0
	Error_UNKNOWN                   Error_Code = 2000000
	Error_STORELOCKED               Error_Code = 2010100
	Error_STOREALREADYEXISTS        Error_Code = 2010200
	Error_STORENOTFOUND             Error_Code = 2010300
	Error_STOREINVALIDARGUMENT      Error_Code = 2010400
	Error_STOREREQUIREDFIELDMISSING Error_Code = 2010401
	Error_STOREINVALIDLATITUDE      Error_Code = 2010402
	Error_STOREINVALIDLONGITUDE     Error_Code = 2010403
	Error_STOREINVALIDTIMEZONE      Error_Code = 2010404
)

var Error_Code_name = map[int32]string{
	0:       "OK",
	2000000: "UNKNOWN",
	2010100: "STORELOCKED",
	2010200: "STOREALREADYEXISTS",
	2010300: "STORENOTFOUND",
	2010400: "STOREINVALIDARGUMENT",
	2010401: "STOREREQUIREDFIELDMISSING",
	2010402: "STOREINVALIDLATITUDE",
	2010403: "STOREINVALIDLONGITUDE",
	2010404: "STOREINVALIDTIMEZONE",
}

var Error_Code_value = map[string]int32{
	"OK":                        0,
	"UNKNOWN":                   2000000,
	"STORELOCKED":               2010100,
	"STOREALREADYEXISTS":        2010200,
	"STORENOTFOUND":             2010300,
	"STOREINVALIDARGUMENT":      2010400,
	"STOREREQUIREDFIELDMISSING": 2010401,
	"STOREINVALIDLATITUDE":      2010402,
	"STOREINVALIDLONGITUDE":     2010403,
	"STOREINVALIDTIMEZONE":      2010404,
}

func (x Error_Code) String() string {
	return proto.EnumName(Error_Code_name, int32(x))
}

func (Error_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ca6c893021d80aab, []int{0, 0}
}

type Error struct {
	Code                 Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=mrs.protobuf.stores.Error_Code" json:"code,omitempty"`
	Cause                string     `protobuf:"bytes,2,opt,name=cause,proto3" json:"cause,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6c893021d80aab, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() Error_Code {
	if m != nil {
		return m.Code
	}
	return Error_OK
}

func (m *Error) GetCause() string {
	if m != nil {
		return m.Cause
	}
	return ""
}

type ValidationError struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationError) Reset()         { *m = ValidationError{} }
func (m *ValidationError) String() string { return proto.CompactTextString(m) }
func (*ValidationError) ProtoMessage()    {}
func (*ValidationError) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6c893021d80aab, []int{1}
}

func (m *ValidationError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationError.Unmarshal(m, b)
}
func (m *ValidationError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationError.Marshal(b, m, deterministic)
}
func (m *ValidationError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationError.Merge(m, src)
}
func (m *ValidationError) XXX_Size() int {
	return xxx_messageInfo_ValidationError.Size(m)
}
func (m *ValidationError) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationError.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationError proto.InternalMessageInfo

func (m *ValidationError) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ValidationError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ListStoresRequest struct {
	Fields               []*ListStoresRequest_FieldFilter `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	Offset               uint64                           `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint64                           `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ListStoresRequest) Reset()         { *m = ListStoresRequest{} }
func (m *ListStoresRequest) String() string { return proto.CompactTextString(m) }
func (*ListStoresRequest) ProtoMessage()    {}
func (*ListStoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6c893021d80aab, []int{2}
}

func (m *ListStoresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStoresRequest.Unmarshal(m, b)
}
func (m *ListStoresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStoresRequest.Marshal(b, m, deterministic)
}
func (m *ListStoresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStoresRequest.Merge(m, src)
}
func (m *ListStoresRequest) XXX_Size() int {
	return xxx_messageInfo_ListStoresRequest.Size(m)
}
func (m *ListStoresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStoresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListStoresRequest proto.InternalMessageInfo

func (m *ListStoresRequest) GetFields() []*ListStoresRequest_FieldFilter {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *ListStoresRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListStoresRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListStoresRequest_FieldFilter struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListStoresRequest_FieldFilter) Reset()         { *m = ListStoresRequest_FieldFilter{} }
func (m *ListStoresRequest_FieldFilter) String() string { return proto.CompactTextString(m) }
func (*ListStoresRequest_FieldFilter) ProtoMessage()    {}
func (*ListStoresRequest_FieldFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6c893021d80aab, []int{2, 0}
}

func (m *ListStoresRequest_FieldFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStoresRequest_FieldFilter.Unmarshal(m, b)
}
func (m *ListStoresRequest_FieldFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStoresRequest_FieldFilter.Marshal(b, m, deterministic)
}
func (m *ListStoresRequest_FieldFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStoresRequest_FieldFilter.Merge(m, src)
}
func (m *ListStoresRequest_FieldFilter) XXX_Size() int {
	return xxx_messageInfo_ListStoresRequest_FieldFilter.Size(m)
}
func (m *ListStoresRequest_FieldFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStoresRequest_FieldFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ListStoresRequest_FieldFilter proto.InternalMessageInfo

func (m *ListStoresRequest_FieldFilter) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ListStoresRequest_FieldFilter) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type ListStoresResponse struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 int32      `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Result               []*Store   `protobuf:"bytes,4,rep,name=result,proto3" json:"result,omitempty"`
	Details              []*any.Any `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListStoresResponse) Reset()         { *m = ListStoresResponse{} }
func (m *ListStoresResponse) String() string { return proto.CompactTextString(m) }
func (*ListStoresResponse) ProtoMessage()    {}
func (*ListStoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6c893021d80aab, []int{3}
}

func (m *ListStoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStoresResponse.Unmarshal(m, b)
}
func (m *ListStoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStoresResponse.Marshal(b, m, deterministic)
}
func (m *ListStoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStoresResponse.Merge(m, src)
}
func (m *ListStoresResponse) XXX_Size() int {
	return xxx_messageInfo_ListStoresResponse.Size(m)
}
func (m *ListStoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListStoresResponse proto.InternalMessageInfo

func (m *ListStoresResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListStoresResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListStoresResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ListStoresResponse) GetResult() []*Store {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ListStoresResponse) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

type GetStoreRequest struct {
	StoreId              string   `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStoreRequest) Reset()         { *m = GetStoreRequest{} }
func (m *GetStoreRequest) String() string { return proto.CompactTextString(m) }
func (*GetStoreRequest) ProtoMessage()    {}
func (*GetStoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6c893021d80aab, []int{4}
}

func (m *GetStoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoreRequest.Unmarshal(m, b)
}
func (m *GetStoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoreRequest.Marshal(b, m, deterministic)
}
func (m *GetStoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoreRequest.Merge(m, src)
}
func (m *GetStoreRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoreRequest.Size(m)
}
func (m *GetStoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoreRequest proto.InternalMessageInfo

func (m *GetStoreRequest) GetStoreId() string {
	if m != nil {
		return m.StoreId
	}
	return ""
}

type GetStoreResponse struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 int32      `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Result               *Store     `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	Details              []*any.Any `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetStoreResponse) Reset()         { *m = GetStoreResponse{} }
func (m *GetStoreResponse) String() string { return proto.CompactTextString(m) }
func (*GetStoreResponse) ProtoMessage()    {}
func (*GetStoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6c893021d80aab, []int{5}
}

func (m *GetStoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoreResponse.Unmarshal(m, b)
}
func (m *GetStoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoreResponse.Marshal(b, m, deterministic)
}
func (m *GetStoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoreResponse.Merge(m, src)
}
func (m *GetStoreResponse) XXX_Size() int {
	return xxx_messageInfo_GetStoreResponse.Size(m)
}
func (m *GetStoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoreResponse proto.InternalMessageInfo

func (m *GetStoreResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetStoreResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetStoreResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetStoreResponse) GetResult() *Store {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetStoreResponse) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterEnum("mrs.protobuf.stores.Error_Code", Error_Code_name, Error_Code_value)
	proto.RegisterType((*Error)(nil), "mrs.protobuf.stores.Error")
	proto.RegisterType((*ValidationError)(nil), "mrs.protobuf.stores.ValidationError")
	proto.RegisterType((*ListStoresRequest)(nil), "mrs.protobuf.stores.ListStoresRequest")
	proto.RegisterType((*ListStoresRequest_FieldFilter)(nil), "mrs.protobuf.stores.ListStoresRequest.FieldFilter")
	proto.RegisterType((*ListStoresResponse)(nil), "mrs.protobuf.stores.ListStoresResponse")
	proto.RegisterType((*GetStoreRequest)(nil), "mrs.protobuf.stores.GetStoreRequest")
	proto.RegisterType((*GetStoreResponse)(nil), "mrs.protobuf.stores.GetStoreResponse")
}

func init() { proto.RegisterFile("stores_spec.proto", fileDescriptor_ca6c893021d80aab) }

var fileDescriptor_ca6c893021d80aab = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xad, 0xf3, 0x6c, 0x6f, 0x69, 0xeb, 0xde, 0x86, 0xca, 0x35, 0x8b, 0x46, 0x16, 0x8f, 0x2c,
	0x90, 0x2b, 0xa5, 0x4b, 0x56, 0x56, 0xed, 0x44, 0xa6, 0xa9, 0x2d, 0x26, 0x4e, 0xa1, 0x6c, 0x2a,
	0xb7, 0x9e, 0x54, 0x96, 0xdc, 0xb8, 0x78, 0x1c, 0xa4, 0x7a, 0xc5, 0x96, 0xbf, 0xe0, 0xf5, 0x09,
	0x88, 0x15, 0x4b, 0xc4, 0x82, 0x1f, 0xa0, 0xe4, 0x2f, 0x90, 0xf8, 0x00, 0x94, 0x99, 0x98, 0xb6,
	0x34, 0xa8, 0x2c, 0x60, 0xe7, 0x73, 0xef, 0x39, 0x67, 0xee, 0x9c, 0x3b, 0x86, 0x65, 0x96, 0xc6,
	0x09, 0x65, 0xfb, 0xec, 0x84, 0x1e, 0xea, 0x27, 0x49, 0x9c, 0xc6, 0xb8, 0x72, 0x9c, 0x30, 0xf1,
	0x79, 0x30, 0xec, 0xeb, 0xa2, 0xaf, 0xae, 0x1d, 0xc5, 0xf1, 0x51, 0x44, 0x37, 0xf2, 0xfa, 0x86,
	0x3f, 0x38, 0x15, 0x24, 0xf5, 0x86, 0xa0, 0x08, 0xa4, 0x7d, 0x2e, 0x40, 0xd9, 0x4a, 0x92, 0x38,
	0xc1, 0x4d, 0x28, 0x1d, 0xc6, 0x01, 0x55, 0xa4, 0xba, 0xd4, 0x58, 0x6c, 0xae, 0xeb, 0x53, 0x6c,
	0x75, 0xce, 0xd4, 0xb7, 0xe2, 0x80, 0x12, 0x4e, 0xc6, 0x1a, 0x94, 0x0f, 0xfd, 0x21, 0xa3, 0x4a,
	0xa1, 0x2e, 0x35, 0xe6, 0x88, 0x00, 0xda, 0x77, 0x09, 0x4a, 0x63, 0x12, 0x56, 0xa0, 0xe0, 0x6e,
	0xcb, 0x33, 0xb8, 0x00, 0xd5, 0x9e, 0xb3, 0xed, 0xb8, 0x8f, 0x1d, 0xf9, 0xc5, 0xcb, 0x0c, 0x97,
	0x61, 0xbe, 0xeb, 0xb9, 0xc4, 0xea, 0xb8, 0x5b, 0xdb, 0x96, 0x29, 0xff, 0xf8, 0x9a, 0xa1, 0x02,
	0xc8, 0x4b, 0x46, 0x87, 0x58, 0x86, 0xb9, 0x67, 0x3d, 0xb1, 0xbb, 0x5e, 0x57, 0x3e, 0x3b, 0xcb,
	0x70, 0x05, 0x16, 0x78, 0xc7, 0x71, 0xbd, 0x96, 0xdb, 0x73, 0x4c, 0xf9, 0xe3, 0xb7, 0x0c, 0x55,
	0xa8, 0xf1, 0xa2, 0xed, 0xec, 0x1a, 0x1d, 0xdb, 0x34, 0x48, 0xbb, 0xb7, 0x63, 0x39, 0x9e, 0xfc,
	0x6a, 0x94, 0xe1, 0x3a, 0xac, 0xf1, 0x1e, 0xb1, 0x1e, 0xf5, 0x6c, 0x62, 0x99, 0x2d, 0xdb, 0xea,
	0x98, 0x3b, 0x76, 0xb7, 0x6b, 0x3b, 0x6d, 0xf9, 0xf5, 0xe8, 0x8a, 0xb8, 0x63, 0x78, 0xb6, 0xd7,
	0x33, 0x2d, 0xf9, 0xcd, 0x28, 0xc3, 0x5b, 0x70, 0xf3, 0x52, 0xcf, 0x75, 0xda, 0xa2, 0xf9, 0xf6,
	0xaa, 0xd0, 0xb3, 0x77, 0xac, 0xa7, 0xae, 0x63, 0xc9, 0xef, 0x46, 0x99, 0x66, 0xc0, 0xd2, 0xae,
	0x1f, 0x85, 0x81, 0x9f, 0x86, 0xf1, 0x40, 0x24, 0x5a, 0x83, 0x72, 0x3f, 0xa4, 0x51, 0xc0, 0x23,
	0x9d, 0x23, 0x02, 0xa0, 0x02, 0xd5, 0x63, 0xca, 0x98, 0x7f, 0x94, 0x87, 0x96, 0x43, 0xed, 0x8b,
	0x04, 0xcb, 0x9d, 0x90, 0xa5, 0x5d, 0x1e, 0x36, 0xa1, 0xcf, 0x86, 0x94, 0xa5, 0xf8, 0x10, 0x2a,
	0x5c, 0xc8, 0x14, 0xa9, 0x5e, 0x6c, 0xcc, 0x37, 0x9b, 0x53, 0x37, 0x73, 0x45, 0xa7, 0xb7, 0xc6,
	0xa2, 0x56, 0x18, 0xa5, 0x34, 0x21, 0x13, 0x07, 0x5c, 0x85, 0x4a, 0xdc, 0xef, 0x33, 0x9a, 0xf2,
	0xa3, 0x4b, 0x64, 0x82, 0xc6, 0x93, 0x46, 0xe1, 0x71, 0x98, 0x2a, 0x45, 0x5e, 0x16, 0x40, 0x7d,
	0x00, 0xf3, 0x17, 0x4c, 0xfe, 0x70, 0x9d, 0x55, 0xa8, 0x3c, 0xf7, 0xa3, 0x21, 0x65, 0x4a, 0xa1,
	0x5e, 0x6c, 0xcc, 0x91, 0x09, 0xd2, 0x3e, 0x48, 0x80, 0x17, 0x87, 0x62, 0x27, 0xf1, 0x80, 0x51,
	0x5c, 0x84, 0x42, 0x98, 0x3b, 0x14, 0xc2, 0x00, 0x71, 0xf2, 0xea, 0xc6, 0xf3, 0x94, 0x27, 0x8f,
	0xea, 0x42, 0x42, 0xc5, 0x4b, 0x09, 0x61, 0x13, 0x2a, 0x09, 0x65, 0xc3, 0x28, 0x55, 0x4a, 0x3c,
	0x0b, 0x75, 0x6a, 0x16, 0xfc, 0x48, 0x32, 0x61, 0xa2, 0x0e, 0xd5, 0x80, 0xa6, 0x7e, 0x18, 0x31,
	0xa5, 0xcc, 0x45, 0x35, 0x5d, 0xfc, 0x1c, 0xe7, 0x3a, 0x63, 0x70, 0x4a, 0x72, 0x92, 0x76, 0x1f,
	0x96, 0xda, 0x54, 0x8c, 0x9d, 0xaf, 0x60, 0x0d, 0x66, 0xb9, 0xf5, 0xfe, 0xaf, 0xd1, 0xab, 0x1c,
	0xdb, 0x81, 0xf6, 0x5e, 0x02, 0xf9, 0x9c, 0xfe, 0xcf, 0x2f, 0x29, 0xfd, 0x9f, 0x4b, 0x36, 0x3f,
	0x49, 0x50, 0x11, 0x9b, 0xc1, 0x7d, 0x80, 0xf3, 0x3d, 0xe1, 0xdd, 0xbf, 0x7b, 0x5d, 0xea, 0xbd,
	0x6b, 0x79, 0x22, 0x0b, 0x6d, 0x06, 0xf7, 0x60, 0x36, 0x4f, 0x08, 0x6f, 0x4f, 0x95, 0xfd, 0x96,
	0xb7, 0x7a, 0xe7, 0x1a, 0x56, 0x6e, 0x7d, 0x50, 0xe1, 0x9c, 0xcd, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xd8, 0x53, 0xf1, 0xf2, 0x17, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoresClient is the client API for Stores service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoresClient interface {
	ListStores(ctx context.Context, in *ListStoresRequest, opts ...grpc.CallOption) (*ListStoresResponse, error)
	GetStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*GetStoreResponse, error)
}

type storesClient struct {
	cc *grpc.ClientConn
}

func NewStoresClient(cc *grpc.ClientConn) StoresClient {
	return &storesClient{cc}
}

func (c *storesClient) ListStores(ctx context.Context, in *ListStoresRequest, opts ...grpc.CallOption) (*ListStoresResponse, error) {
	out := new(ListStoresResponse)
	err := c.cc.Invoke(ctx, "/mrs.protobuf.stores.Stores/ListStores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesClient) GetStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*GetStoreResponse, error) {
	out := new(GetStoreResponse)
	err := c.cc.Invoke(ctx, "/mrs.protobuf.stores.Stores/GetStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoresServer is the server API for Stores service.
type StoresServer interface {
	ListStores(context.Context, *ListStoresRequest) (*ListStoresResponse, error)
	GetStore(context.Context, *GetStoreRequest) (*GetStoreResponse, error)
}

func RegisterStoresServer(s *grpc.Server, srv StoresServer) {
	s.RegisterService(&_Stores_serviceDesc, srv)
}

func _Stores_ListStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServer).ListStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mrs.protobuf.stores.Stores/ListStores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServer).ListStores(ctx, req.(*ListStoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stores_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoresServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mrs.protobuf.stores.Stores/GetStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoresServer).GetStore(ctx, req.(*GetStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stores_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mrs.protobuf.stores.Stores",
	HandlerType: (*StoresServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListStores",
			Handler:    _Stores_ListStores_Handler,
		},
		{
			MethodName: "GetStore",
			Handler:    _Stores_GetStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stores_spec.proto",
}
