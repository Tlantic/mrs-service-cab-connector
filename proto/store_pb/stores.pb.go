// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stores.proto

package store_pb

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

type Store struct {
	Id                          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status                      string    `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Name                        string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Group                       *Group    `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	Chain                       *Chain    `protobuf:"bytes,5,opt,name=chain,proto3" json:"chain,omitempty"`
	Locale                      string    `protobuf:"bytes,6,opt,name=locale,proto3" json:"locale,omitempty"`
	Language                    string    `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	Contact                     *Contact  `protobuf:"bytes,8,opt,name=contact,proto3" json:"contact,omitempty"`
	Location                    *Location `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	Timezone                    *Timezone `protobuf:"bytes,10,opt,name=timezone,proto3" json:"timezone,omitempty"`
	OpeningHourWeek             string    `protobuf:"bytes,11,opt,name=opening_hour_week,json=openingHourWeek,proto3" json:"opening_hour_week,omitempty"`
	ClosingHourWeek             string    `protobuf:"bytes,12,opt,name=closing_hour_week,json=closingHourWeek,proto3" json:"closing_hour_week,omitempty"`
	OpeningHourWeekend          string    `protobuf:"bytes,13,opt,name=opening_hour_weekend,json=openingHourWeekend,proto3" json:"opening_hour_weekend,omitempty"`
	ClosingHourWeekend          string    `protobuf:"bytes,14,opt,name=closing_hour_weekend,json=closingHourWeekend,proto3" json:"closing_hour_weekend,omitempty"`
	WeightUnits                 string    `protobuf:"bytes,15,opt,name=weight_units,json=weightUnits,proto3" json:"weight_units,omitempty"`
	WeightDecimalPlaces         string    `protobuf:"bytes,16,opt,name=weight_decimal_places,json=weightDecimalPlaces,proto3" json:"weight_decimal_places,omitempty"`
	WeightDecimalSeparator      string    `protobuf:"bytes,17,opt,name=weight_decimal_separator,json=weightDecimalSeparator,proto3" json:"weight_decimal_separator,omitempty"`
	WeightThousandsSeparator    string    `protobuf:"bytes,18,opt,name=weight_thousands_separator,json=weightThousandsSeparator,proto3" json:"weight_thousands_separator,omitempty"`
	DimensionUnits              string    `protobuf:"bytes,19,opt,name=dimension_units,json=dimensionUnits,proto3" json:"dimension_units,omitempty"`
	DimensionDecimalPlaces      string    `protobuf:"bytes,20,opt,name=dimension_decimal_places,json=dimensionDecimalPlaces,proto3" json:"dimension_decimal_places,omitempty"`
	DimensionDecimalSeparator   string    `protobuf:"bytes,21,opt,name=dimension_decimal_separator,json=dimensionDecimalSeparator,proto3" json:"dimension_decimal_separator,omitempty"`
	DimensionThousandsSeparator string    `protobuf:"bytes,22,opt,name=dimension_thousands_separator,json=dimensionThousandsSeparator,proto3" json:"dimension_thousands_separator,omitempty"`
	Currency                    string    `protobuf:"bytes,23,opt,name=currency,proto3" json:"currency,omitempty"`
	CurrencySymbol              string    `protobuf:"bytes,24,opt,name=currency_symbol,json=currencySymbol,proto3" json:"currency_symbol,omitempty"`
	CurrencySymbolLocation      string    `protobuf:"bytes,25,opt,name=currency_symbol_location,json=currencySymbolLocation,proto3" json:"currency_symbol_location,omitempty"`
	CurrencyDecimalPlaces       string    `protobuf:"bytes,26,opt,name=currency_decimal_places,json=currencyDecimalPlaces,proto3" json:"currency_decimal_places,omitempty"`
	CurrencyDecimalSeparator    string    `protobuf:"bytes,27,opt,name=currency_decimal_separator,json=currencyDecimalSeparator,proto3" json:"currency_decimal_separator,omitempty"`
	CurrencyThousandsSeparator  string    `protobuf:"bytes,28,opt,name=currency_thousands_separator,json=currencyThousandsSeparator,proto3" json:"currency_thousands_separator,omitempty"`
	DecimalPlaces               string    `protobuf:"bytes,29,opt,name=decimal_places,json=decimalPlaces,proto3" json:"decimal_places,omitempty"`                // Deprecated: Do not use.
	DecimalSeparator            string    `protobuf:"bytes,30,opt,name=decimal_separator,json=decimalSeparator,proto3" json:"decimal_separator,omitempty"`       // Deprecated: Do not use.
	ThousandsSeparator          string    `protobuf:"bytes,31,opt,name=thousands_separator,json=thousandsSeparator,proto3" json:"thousands_separator,omitempty"` // Deprecated: Do not use.
	ExternalCode                string    `protobuf:"bytes,32,opt,name=external_code,json=externalCode,proto3" json:"external_code,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}  `json:"-"`
	XXX_unrecognized            []byte    `json:"-"`
	XXX_sizecache               int32     `json:"-"`
}

func (m *Store) Reset()         { *m = Store{} }
func (m *Store) String() string { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()    {}
func (*Store) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bd58a4c4375fca, []int{0}
}

func (m *Store) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Store.Unmarshal(m, b)
}
func (m *Store) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Store.Marshal(b, m, deterministic)
}
func (m *Store) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Store.Merge(m, src)
}
func (m *Store) XXX_Size() int {
	return xxx_messageInfo_Store.Size(m)
}
func (m *Store) XXX_DiscardUnknown() {
	xxx_messageInfo_Store.DiscardUnknown(m)
}

var xxx_messageInfo_Store proto.InternalMessageInfo

func (m *Store) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Store) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Store) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Store) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *Store) GetChain() *Chain {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *Store) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *Store) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Store) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

func (m *Store) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Store) GetTimezone() *Timezone {
	if m != nil {
		return m.Timezone
	}
	return nil
}

func (m *Store) GetOpeningHourWeek() string {
	if m != nil {
		return m.OpeningHourWeek
	}
	return ""
}

func (m *Store) GetClosingHourWeek() string {
	if m != nil {
		return m.ClosingHourWeek
	}
	return ""
}

func (m *Store) GetOpeningHourWeekend() string {
	if m != nil {
		return m.OpeningHourWeekend
	}
	return ""
}

func (m *Store) GetClosingHourWeekend() string {
	if m != nil {
		return m.ClosingHourWeekend
	}
	return ""
}

func (m *Store) GetWeightUnits() string {
	if m != nil {
		return m.WeightUnits
	}
	return ""
}

func (m *Store) GetWeightDecimalPlaces() string {
	if m != nil {
		return m.WeightDecimalPlaces
	}
	return ""
}

func (m *Store) GetWeightDecimalSeparator() string {
	if m != nil {
		return m.WeightDecimalSeparator
	}
	return ""
}

func (m *Store) GetWeightThousandsSeparator() string {
	if m != nil {
		return m.WeightThousandsSeparator
	}
	return ""
}

func (m *Store) GetDimensionUnits() string {
	if m != nil {
		return m.DimensionUnits
	}
	return ""
}

func (m *Store) GetDimensionDecimalPlaces() string {
	if m != nil {
		return m.DimensionDecimalPlaces
	}
	return ""
}

func (m *Store) GetDimensionDecimalSeparator() string {
	if m != nil {
		return m.DimensionDecimalSeparator
	}
	return ""
}

func (m *Store) GetDimensionThousandsSeparator() string {
	if m != nil {
		return m.DimensionThousandsSeparator
	}
	return ""
}

func (m *Store) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Store) GetCurrencySymbol() string {
	if m != nil {
		return m.CurrencySymbol
	}
	return ""
}

func (m *Store) GetCurrencySymbolLocation() string {
	if m != nil {
		return m.CurrencySymbolLocation
	}
	return ""
}

func (m *Store) GetCurrencyDecimalPlaces() string {
	if m != nil {
		return m.CurrencyDecimalPlaces
	}
	return ""
}

func (m *Store) GetCurrencyDecimalSeparator() string {
	if m != nil {
		return m.CurrencyDecimalSeparator
	}
	return ""
}

func (m *Store) GetCurrencyThousandsSeparator() string {
	if m != nil {
		return m.CurrencyThousandsSeparator
	}
	return ""
}

// Deprecated: Do not use.
func (m *Store) GetDecimalPlaces() string {
	if m != nil {
		return m.DecimalPlaces
	}
	return ""
}

// Deprecated: Do not use.
func (m *Store) GetDecimalSeparator() string {
	if m != nil {
		return m.DecimalSeparator
	}
	return ""
}

// Deprecated: Do not use.
func (m *Store) GetThousandsSeparator() string {
	if m != nil {
		return m.ThousandsSeparator
	}
	return ""
}

func (m *Store) GetExternalCode() string {
	if m != nil {
		return m.ExternalCode
	}
	return ""
}

type Contact struct {
	FormattedPhone       string   `protobuf:"bytes,1,opt,name=formatted_phone,json=formattedPhone,proto3" json:"formatted_phone,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bd58a4c4375fca, []int{1}
}

func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (m *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(m, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetFormattedPhone() string {
	if m != nil {
		return m.FormattedPhone
	}
	return ""
}

func (m *Contact) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type Location struct {
	Cc                   string       `protobuf:"bytes,1,opt,name=cc,proto3" json:"cc,omitempty"`
	Country              string       `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	State                string       `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	City                 string       `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	PostalCode           string       `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	CrossStreet          string       `protobuf:"bytes,6,opt,name=cross_street,json=crossStreet,proto3" json:"cross_street,omitempty"`
	Address              string       `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Coords               *Coordinates `protobuf:"bytes,8,opt,name=coords,proto3" json:"coords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bd58a4c4375fca, []int{2}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetCc() string {
	if m != nil {
		return m.Cc
	}
	return ""
}

func (m *Location) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Location) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Location) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Location) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *Location) GetCrossStreet() string {
	if m != nil {
		return m.CrossStreet
	}
	return ""
}

func (m *Location) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Location) GetCoords() *Coordinates {
	if m != nil {
		return m.Coords
	}
	return nil
}

type Coordinates struct {
	Lat                  string   `protobuf:"bytes,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon                  string   `protobuf:"bytes,2,opt,name=lon,proto3" json:"lon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinates) Reset()         { *m = Coordinates{} }
func (m *Coordinates) String() string { return proto.CompactTextString(m) }
func (*Coordinates) ProtoMessage()    {}
func (*Coordinates) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bd58a4c4375fca, []int{3}
}

func (m *Coordinates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinates.Unmarshal(m, b)
}
func (m *Coordinates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinates.Marshal(b, m, deterministic)
}
func (m *Coordinates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinates.Merge(m, src)
}
func (m *Coordinates) XXX_Size() int {
	return xxx_messageInfo_Coordinates.Size(m)
}
func (m *Coordinates) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinates.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinates proto.InternalMessageInfo

func (m *Coordinates) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *Coordinates) GetLon() string {
	if m != nil {
		return m.Lon
	}
	return ""
}

type Timezone struct {
	Value                string      `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	DateFormat           *DateFormat `protobuf:"bytes,2,opt,name=date_format,json=dateFormat,proto3" json:"date_format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Timezone) Reset()         { *m = Timezone{} }
func (m *Timezone) String() string { return proto.CompactTextString(m) }
func (*Timezone) ProtoMessage()    {}
func (*Timezone) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bd58a4c4375fca, []int{4}
}

func (m *Timezone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timezone.Unmarshal(m, b)
}
func (m *Timezone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timezone.Marshal(b, m, deterministic)
}
func (m *Timezone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timezone.Merge(m, src)
}
func (m *Timezone) XXX_Size() int {
	return xxx_messageInfo_Timezone.Size(m)
}
func (m *Timezone) XXX_DiscardUnknown() {
	xxx_messageInfo_Timezone.DiscardUnknown(m)
}

var xxx_messageInfo_Timezone proto.InternalMessageInfo

func (m *Timezone) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Timezone) GetDateFormat() *DateFormat {
	if m != nil {
		return m.DateFormat
	}
	return nil
}

type DateFormat struct {
	Display              string   `protobuf:"bytes,1,opt,name=display,proto3" json:"display,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateFormat) Reset()         { *m = DateFormat{} }
func (m *DateFormat) String() string { return proto.CompactTextString(m) }
func (*DateFormat) ProtoMessage()    {}
func (*DateFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bd58a4c4375fca, []int{5}
}

func (m *DateFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateFormat.Unmarshal(m, b)
}
func (m *DateFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateFormat.Marshal(b, m, deterministic)
}
func (m *DateFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateFormat.Merge(m, src)
}
func (m *DateFormat) XXX_Size() int {
	return xxx_messageInfo_DateFormat.Size(m)
}
func (m *DateFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_DateFormat.DiscardUnknown(m)
}

var xxx_messageInfo_DateFormat proto.InternalMessageInfo

func (m *DateFormat) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

type Group struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bd58a4c4375fca, []int{6}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Chain struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chain) Reset()         { *m = Chain{} }
func (m *Chain) String() string { return proto.CompactTextString(m) }
func (*Chain) ProtoMessage()    {}
func (*Chain) Descriptor() ([]byte, []int) {
	return fileDescriptor_78bd58a4c4375fca, []int{7}
}

func (m *Chain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chain.Unmarshal(m, b)
}
func (m *Chain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chain.Marshal(b, m, deterministic)
}
func (m *Chain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chain.Merge(m, src)
}
func (m *Chain) XXX_Size() int {
	return xxx_messageInfo_Chain.Size(m)
}
func (m *Chain) XXX_DiscardUnknown() {
	xxx_messageInfo_Chain.DiscardUnknown(m)
}

var xxx_messageInfo_Chain proto.InternalMessageInfo

func (m *Chain) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Chain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Store)(nil), "mrs.protobuf.stores.Store")
	proto.RegisterType((*Contact)(nil), "mrs.protobuf.stores.Contact")
	proto.RegisterType((*Location)(nil), "mrs.protobuf.stores.Location")
	proto.RegisterType((*Coordinates)(nil), "mrs.protobuf.stores.Coordinates")
	proto.RegisterType((*Timezone)(nil), "mrs.protobuf.stores.Timezone")
	proto.RegisterType((*DateFormat)(nil), "mrs.protobuf.stores.DateFormat")
	proto.RegisterType((*Group)(nil), "mrs.protobuf.stores.Group")
	proto.RegisterType((*Chain)(nil), "mrs.protobuf.stores.Chain")
}

func init() { proto.RegisterFile("stores.proto", fileDescriptor_78bd58a4c4375fca) }

var fileDescriptor_78bd58a4c4375fca = []byte{
	// 854 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0x5b, 0x6f, 0xdb, 0x36,
	0x14, 0xc7, 0x61, 0xb7, 0x4e, 0x9c, 0xa3, 0xc4, 0x49, 0x98, 0x4b, 0x59, 0x37, 0x59, 0x32, 0x0f,
	0xd8, 0xba, 0x0d, 0xf0, 0xba, 0x14, 0x28, 0x32, 0xa0, 0x18, 0x8a, 0xb6, 0xd8, 0xfa, 0xb0, 0x87,
	0x42, 0xee, 0xb0, 0x47, 0x81, 0x26, 0x59, 0x5b, 0xa8, 0x4c, 0x0a, 0x24, 0xb5, 0xce, 0xfb, 0xc4,
	0xfb, 0x04, 0x7b, 0x1e, 0x78, 0x93, 0x6a, 0x59, 0x0b, 0xfa, 0x76, 0x2e, 0xff, 0x1f, 0x79, 0xce,
	0x11, 0xc4, 0x03, 0xfb, 0xda, 0x48, 0xc5, 0xf5, 0xb4, 0x54, 0xd2, 0x48, 0x74, 0xb2, 0x52, 0xc1,
	0x9c, 0x57, 0xef, 0xa7, 0x3e, 0x35, 0xf9, 0x27, 0x81, 0xc1, 0xcc, 0x9a, 0x68, 0x04, 0xfd, 0x9c,
	0xe1, 0xde, 0x75, 0xef, 0xf1, 0x5e, 0xda, 0xcf, 0x19, 0x3a, 0x87, 0x1d, 0x6d, 0x88, 0xa9, 0x34,
	0xee, 0xbb, 0x58, 0xf0, 0x10, 0x82, 0xfb, 0x82, 0xac, 0x38, 0xbe, 0xe7, 0xa2, 0xce, 0x46, 0x4f,
	0x60, 0xb0, 0x50, 0xb2, 0x2a, 0xf1, 0xfd, 0xeb, 0xde, 0xe3, 0xe4, 0x66, 0x3c, 0xed, 0xb8, 0x6a,
	0xfa, 0xab, 0x55, 0xa4, 0x5e, 0x68, 0x09, 0xba, 0x24, 0xb9, 0xc0, 0x83, 0x3b, 0x88, 0x57, 0x56,
	0x91, 0x7a, 0xa1, 0xad, 0xa7, 0x90, 0x94, 0x14, 0x1c, 0xef, 0xf8, 0x7a, 0xbc, 0x87, 0xc6, 0x30,
	0x2c, 0x88, 0x58, 0x54, 0x64, 0xc1, 0xf1, 0xae, 0xcb, 0xd4, 0x3e, 0x7a, 0x06, 0xbb, 0x54, 0x0a,
	0x43, 0xa8, 0xc1, 0x43, 0x77, 0xcf, 0x45, 0xf7, 0x3d, 0x5e, 0x93, 0x46, 0x31, 0xfa, 0x09, 0x86,
	0xf6, 0x74, 0x93, 0x4b, 0x81, 0xf7, 0x1c, 0x78, 0xd9, 0x09, 0xfe, 0x16, 0x44, 0x69, 0x2d, 0xb7,
	0xa8, 0xc9, 0x57, 0xfc, 0x6f, 0x29, 0x38, 0x86, 0x3b, 0xd0, 0x77, 0x41, 0x94, 0xd6, 0x72, 0xf4,
	0x1d, 0x1c, 0xcb, 0x92, 0x8b, 0x5c, 0x2c, 0xb2, 0xa5, 0xac, 0x54, 0xf6, 0x91, 0xf3, 0x0f, 0x38,
	0x71, 0x2d, 0x1d, 0x86, 0xc4, 0x1b, 0x59, 0xa9, 0x3f, 0x38, 0xff, 0x60, 0xb5, 0xb4, 0x90, 0x7a,
	0x53, 0xbb, 0xef, 0xb5, 0x21, 0x51, 0x6b, 0x9f, 0xc0, 0xe9, 0xd6, 0xb9, 0x5c, 0x30, 0x7c, 0xe0,
	0xe4, 0xa8, 0x75, 0x34, 0x17, 0xcc, 0x12, 0x5b, 0xa7, 0x5b, 0x62, 0xe4, 0x89, 0xd6, 0x05, 0x96,
	0xf8, 0x12, 0xf6, 0x3f, 0xf2, 0x7c, 0xb1, 0x34, 0x59, 0x25, 0x72, 0xa3, 0xf1, 0xa1, 0x53, 0x26,
	0x3e, 0xf6, 0xbb, 0x0d, 0xa1, 0x1b, 0x38, 0x0b, 0x12, 0xc6, 0x69, 0xbe, 0x22, 0x45, 0x56, 0x16,
	0x84, 0x72, 0x8d, 0x8f, 0x9c, 0xf6, 0xc4, 0x27, 0x5f, 0xfb, 0xdc, 0x5b, 0x97, 0x42, 0xb7, 0x80,
	0x5b, 0x8c, 0xe6, 0x25, 0x51, 0xc4, 0x48, 0x85, 0x8f, 0x1d, 0x76, 0xbe, 0x81, 0xcd, 0x62, 0x16,
	0x3d, 0x87, 0x71, 0x20, 0xcd, 0x52, 0x56, 0x9a, 0x08, 0xa6, 0x3f, 0x61, 0x91, 0x63, 0xc3, 0xd9,
	0xef, 0xa2, 0xa0, 0xa1, 0xbf, 0x81, 0x43, 0x96, 0xaf, 0xb8, 0xd0, 0xb9, 0x14, 0xa1, 0xa3, 0x13,
	0x87, 0x8c, 0xea, 0xb0, 0x6f, 0xea, 0x16, 0x70, 0x23, 0x6c, 0xf5, 0x75, 0xea, 0x0b, 0xac, 0xf3,
	0x9b, 0xad, 0xfd, 0x0c, 0x8f, 0xb6, 0xc9, 0xa6, 0xc2, 0x33, 0x07, 0x3f, 0x6c, 0xc3, 0x4d, 0x89,
	0x2f, 0xe1, 0xb2, 0xe1, 0xbb, 0x7a, 0x3c, 0x77, 0x27, 0x34, 0x97, 0x74, 0xb4, 0x39, 0x86, 0x21,
	0xad, 0x94, 0xe2, 0x82, 0xae, 0xf1, 0x03, 0xff, 0xef, 0x44, 0xdf, 0x8e, 0x20, 0xda, 0x99, 0x5e,
	0xaf, 0xe6, 0xb2, 0xc0, 0xd8, 0x8f, 0x20, 0x86, 0x67, 0x2e, 0x6a, 0x47, 0xd0, 0x12, 0x66, 0xf5,
	0xcf, 0xf3, 0xd0, 0x8f, 0x60, 0x93, 0x88, 0x7f, 0x0d, 0x7a, 0x06, 0x0f, 0x6a, 0xb2, 0x35, 0xbb,
	0xb1, 0x03, 0xcf, 0x62, 0x7a, 0x73, 0x74, 0xcf, 0x61, 0xbc, 0xc5, 0x35, 0x7d, 0x3f, 0xf2, 0xdf,
	0xb6, 0x85, 0x36, 0x4d, 0xbf, 0x80, 0x8b, 0x9a, 0xee, 0x9a, 0xdb, 0x85, 0xe3, 0xeb, 0x1b, 0x3a,
	0xc6, 0xf6, 0x2d, 0x8c, 0x5a, 0xe5, 0x5e, 0x5a, 0xe6, 0x65, 0x1f, 0xf7, 0xd2, 0x03, 0xb6, 0x51,
	0xea, 0x0f, 0x70, 0xbc, 0x5d, 0xe1, 0x17, 0xb5, 0xfa, 0x88, 0xb5, 0xab, 0x7b, 0x0a, 0x27, 0x5d,
	0x45, 0x5d, 0xd5, 0x08, 0x32, 0xdb, 0x05, 0x7d, 0x05, 0x07, 0xfc, 0x2f, 0xc3, 0x95, 0x20, 0x45,
	0x46, 0x25, 0xe3, 0xf8, 0xda, 0xf5, 0xb0, 0x1f, 0x83, 0xaf, 0x24, 0xe3, 0x93, 0x37, 0xb0, 0x1b,
	0x1e, 0x3a, 0xfb, 0x6d, 0xdf, 0x4b, 0xb5, 0x22, 0xc6, 0x70, 0x96, 0x95, 0x4b, 0xfb, 0x56, 0xf9,
	0x87, 0x7f, 0x54, 0x87, 0xdf, 0xda, 0x28, 0x3a, 0x85, 0x81, 0x4f, 0xfb, 0x1d, 0xe0, 0x9d, 0xc9,
	0xbf, 0x3d, 0x18, 0xd6, 0x1f, 0x71, 0x04, 0x7d, 0x4a, 0xe3, 0xde, 0xa0, 0x14, 0x61, 0xfb, 0xe6,
	0x56, 0xc2, 0xa8, 0x75, 0x80, 0xa2, 0x6b, 0x0f, 0xb3, 0x3b, 0x24, 0xae, 0x0e, 0xef, 0xd8, 0x7d,
	0x42, 0x73, 0xb3, 0x76, 0xab, 0x63, 0x2f, 0x75, 0x36, 0xba, 0x82, 0xa4, 0x94, 0xda, 0xc4, 0x6e,
	0x06, 0x2e, 0x05, 0x3e, 0x64, 0x7b, 0xb1, 0xcf, 0x0d, 0x55, 0x52, 0xeb, 0x4c, 0x1b, 0xc5, 0xb9,
	0x09, 0x2b, 0x21, 0x71, 0xb1, 0x99, 0x0b, 0xd9, 0x3a, 0x08, 0x63, 0x8a, 0x6b, 0x1d, 0xd6, 0x42,
	0x74, 0xd1, 0x2d, 0xec, 0x50, 0x29, 0x15, 0xd3, 0x61, 0x29, 0x5c, 0xff, 0xcf, 0x52, 0x90, 0x8a,
	0xe5, 0x82, 0x18, 0xae, 0xd3, 0xa0, 0x9f, 0xfc, 0x08, 0xc9, 0x27, 0x61, 0x74, 0x04, 0xf7, 0x0a,
	0x62, 0x42, 0xef, 0xd6, 0x74, 0x11, 0x29, 0x42, 0xe3, 0xd6, 0x9c, 0xcc, 0x61, 0x18, 0x9f, 0x7a,
	0x3b, 0x80, 0x3f, 0x49, 0x51, 0xc5, 0x61, 0x7b, 0x07, 0xbd, 0x80, 0x84, 0x11, 0xc3, 0x33, 0x3f,
	0x7a, 0xc7, 0x26, 0x37, 0x57, 0x9d, 0x35, 0xbd, 0x26, 0x86, 0xff, 0xe2, 0x64, 0x29, 0xb0, 0xda,
	0x9e, 0x7c, 0x0d, 0xd0, 0x64, 0x6c, 0xe3, 0x2c, 0xd7, 0x65, 0x41, 0xd6, 0xe1, 0x9e, 0xe8, 0x4e,
	0xbe, 0x87, 0x81, 0x5b, 0xc2, 0x5b, 0xbb, 0x3e, 0xee, 0xf4, 0x7e, 0xb3, 0xd3, 0xad, 0xd8, 0xed,
	0xdf, 0xcf, 0x11, 0xcf, 0x77, 0x5c, 0xa5, 0x4f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xbf,
	0x59, 0x42, 0x72, 0x08, 0x00, 0x00,
}
