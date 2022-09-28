package stores

import (
	"encoding/json"

	"github.com/Tlantic/mrs-service-cab-connector/proto/store_pb"
	"github.com/shopspring/decimal"
)

// Store represents an MRS store
type Store struct {
	ID     string `json:"id,omitempty" validate:"required"`
	Status string `json:"status,omitempty" validate:"omitempty,eq=A|eq=I"`

	Name  string `json:"name,omitempty" validate:"required"`
	Group Group  `json:"group,omitempty"`
	Chain Chain  `json:"chain,omitempty" validate:"omitempty,dive"`

	Locale             string   `json:"locale,omitempty"`
	Language           string   `json:"language,omitempty"`
	Contact            Contact  `json:"contact,omitempty" validate:"omitempty,gt=0,dive"`
	Location           Location `json:"location,omitempty" validate:"omitempty,gt=0,dive"`
	Timezone           Timezone `json:"timezone,omitempty" validate:"required,gt=0,dive"`
	OpeningHourWeek    string   `json:"opening_hour_week,omitempty"`
	ClosingHourWeek    string   `json:"closing_hour_week,omitempty"`
	OpeningHourWeekend string   `json:"opening_hour_weekend,omitempty"`
	ClosingHourWeekend string   `json:"closing_hour_weekend,omitempty"`

	WeightUnits              string `json:"weight_units,omitempty"`
	WeightDecimalPlaces      string `json:"weight_decimal_places,omitempty"`
	WeightDecimalSeparator   string `json:"weight_decimal_separator,omitempty"`
	WeightThousandsSeparator string `json:"weight_thousands_separator,omitempty"`

	DimensionUnits              string `json:"dimension_units,omitempty"`
	DimensionDecimalPlaces      string `json:"dimension_decimal_places,omitempty"`
	DimensionDecimalSeparator   string `json:"dimension_decimal_separator,omitempty"`
	DimensionThousandsSeparator string `json:"dimension_thousands_separator,omitempty"`

	// Currency code used by the store (@SEE: https://www.iso.org/iso-4217-currency-codes.html)
	Currency                   string `json:"currency,omitempty"`
	CurrencySymbol             string `json:"currency_symbol,omitempty"`
	CurrencySymbolLocation     string `json:"currency_symbol_location,omitempty"`
	CurrencyDecimalPlaces      string `json:"currency_decimal_places,omitempty"`
	CurrencyDecimalSeparator   string `json:"currency_decimal_separator,omitempty"`
	CurrencyThousandsSeparator string `json:"currency_thousands_separator,omitempty"`

	DecimalPlaces      string `json:"decimal_places,omitempty"`      // DecimalPlaces @DEPRECATED: use CurrencyDecimalPlaces instead
	DecimalSeparator   string `json:"decimal_separator,omitempty"`   // DecimalSeparator @DEPRECATED: use CurrencyDecimalSeparator instead
	ThousandsSeparator string `json:"thousands_separator,omitempty"` // ThousandsSeparator @DEPRECATED: use CurrencyThousandsSeparator instead

	ExternalCode string `json:"external_code,omitempty"`
}

// NewStore returns a new instance
func NewStore() Store {
	return Store{}
}

// NewStoreWithID returns a new instance with an id already defined
func NewStoreWithID(id string) Store {
	return Store{ID: id}
}

// NewStoreFromProto converts object from a proto buffer definition
func NewStoreFromProto(pbv *store_pb.Store) Store {
	if pbv == nil {
		return Store{}
	}
	return Store{
		ID:     pbv.GetId(),
		Status: pbv.GetStatus(),

		Name:  pbv.GetName(),
		Group: NewGroupFromProto(pbv.GetGroup()),
		Chain: NewChainFromProto(pbv.GetChain()),

		Locale:             pbv.GetLocale(),
		Language:           pbv.GetLanguage(),
		Contact:            NewContactFromProto(pbv.GetContact()),
		Location:           NewLocationFromProto(pbv.GetLocation()),
		Timezone:           NewTimezoneFromProto(pbv.GetTimezone()),
		OpeningHourWeek:    pbv.GetOpeningHourWeek(),
		ClosingHourWeek:    pbv.GetClosingHourWeek(),
		OpeningHourWeekend: pbv.GetOpeningHourWeekend(),
		ClosingHourWeekend: pbv.GetClosingHourWeekend(),

		WeightUnits:              pbv.GetWeightUnits(),
		WeightDecimalPlaces:      pbv.GetWeightDecimalPlaces(),
		WeightDecimalSeparator:   pbv.GetWeightDecimalSeparator(),
		WeightThousandsSeparator: pbv.GetWeightThousandsSeparator(),

		DimensionUnits:              pbv.GetDimensionUnits(),
		DimensionDecimalPlaces:      pbv.GetDimensionDecimalPlaces(),
		DimensionDecimalSeparator:   pbv.GetDimensionDecimalSeparator(),
		DimensionThousandsSeparator: pbv.GetDimensionThousandsSeparator(),

		Currency:                   pbv.GetCurrency(),
		CurrencySymbol:             pbv.GetCurrencySymbol(),
		CurrencySymbolLocation:     pbv.GetCurrencySymbolLocation(),
		CurrencyDecimalPlaces:      pbv.GetCurrencyDecimalPlaces(),
		CurrencyDecimalSeparator:   pbv.GetCurrencyDecimalSeparator(),
		CurrencyThousandsSeparator: pbv.GetCurrencyThousandsSeparator(),

		//lint:ignore SA1019 we want to keep this property until all clients stop using it
		DecimalPlaces: pbv.GetDecimalPlaces(),
		//lint:ignore SA1019 we want to keep this property until all clients stop using it
		DecimalSeparator: pbv.GetDecimalSeparator(),
		//lint:ignore SA1019 we want to keep this property until all clients stop using it
		ThousandsSeparator: pbv.GetThousandsSeparator(),

		ExternalCode: pbv.GetExternalCode(),
	}
}

// Patch sets any non-zero value to the correspondent entity value
func (m Store) Patch(dst *Store) error {
	if dst == nil {
		return nil
	}
	if m.ID != "" {
		dst.ID = m.ID

	}
	if m.Status != "" {
		dst.Status = m.Status
	}

	if m.Name != "" {
		dst.Name = m.Name
	}
	if err := m.Group.Patch(&dst.Group); err != nil {
		return err
	}
	if err := m.Chain.Patch(&dst.Chain); err != nil {
		return err
	}

	if m.Locale != "" {
		dst.Locale = m.Locale
	}
	if m.Language != "" {
		dst.Language = m.Language
	}

	if err := m.Contact.Patch(&dst.Contact); err != nil {
		return err
	}
	if err := m.Location.Patch(&dst.Location); err != nil {
		return err
	}

	if m.OpeningHourWeek != "" {
		dst.OpeningHourWeek = m.OpeningHourWeek
	}
	if m.ClosingHourWeek != "" {
		dst.ClosingHourWeek = m.ClosingHourWeek
	}
	if m.OpeningHourWeekend != "" {
		dst.OpeningHourWeekend = m.OpeningHourWeekend
	}
	if m.ClosingHourWeekend != "" {
		dst.ClosingHourWeekend = m.ClosingHourWeekend
	}

	if m.WeightUnits != "" {
		dst.WeightUnits = m.WeightUnits
	}
	if m.WeightDecimalPlaces != "" {
		dst.WeightDecimalPlaces = m.WeightDecimalPlaces
	}
	if m.WeightDecimalSeparator != "" {
		dst.WeightDecimalSeparator = m.WeightDecimalSeparator
	}
	if m.WeightThousandsSeparator != "" {
		dst.WeightThousandsSeparator = m.WeightThousandsSeparator
	}

	if m.DimensionUnits != "" {
		dst.DimensionUnits = m.DimensionUnits
	}
	if m.DimensionDecimalPlaces != "" {
		dst.DimensionDecimalPlaces = m.DimensionDecimalPlaces
	}
	if m.DimensionDecimalSeparator != "" {
		dst.DimensionDecimalSeparator = m.DimensionDecimalSeparator
	}
	if m.DimensionThousandsSeparator != "" {
		dst.DimensionThousandsSeparator = m.DimensionThousandsSeparator
	}

	// Currency code used by the store (@SEE: https://www.iso.org/iso-4217-currency-codes.html)
	if m.Currency != "" {
		dst.Currency = m.Currency
	}
	if m.CurrencySymbol != "" {
		dst.CurrencySymbol = m.CurrencySymbol
	}
	if m.CurrencySymbolLocation != "" {
		dst.CurrencySymbolLocation = m.CurrencySymbolLocation
	}
	if m.CurrencyDecimalPlaces != "" {
		dst.CurrencyDecimalPlaces = m.CurrencyDecimalPlaces
	}
	if m.CurrencyDecimalSeparator != "" {
		dst.CurrencyDecimalSeparator = m.CurrencyDecimalSeparator
	}
	if m.CurrencyThousandsSeparator != "" {
		dst.CurrencyThousandsSeparator = m.CurrencyThousandsSeparator
	}

	if m.DecimalPlaces != "" {
		dst.DecimalPlaces = m.DecimalPlaces
	}
	if m.DecimalSeparator != "" {
		dst.DecimalSeparator = m.DecimalSeparator
	}
	if m.ThousandsSeparator != "" {
		dst.ThousandsSeparator = m.ThousandsSeparator
	}
	if m.ExternalCode != "" {
		dst.ExternalCode = m.ExternalCode
	}
	return nil
}

// Clone creates a new instance copying any pointer value to new memory space
func (m Store) Clone() Store {
	m.Group = m.Group.Clone()
	m.Chain = m.Chain.Clone()
	m.Contact = m.Contact.Clone()
	m.Location = m.Location.Clone()
	m.Timezone = m.Timezone.Clone()
	return m
}

// Proto returns the proto buffer equivalent of this object
func (m Store) Proto() *store_pb.Store {
	return &store_pb.Store{
		Id:     m.ID,
		Status: m.Status,

		Name:  m.Name,
		Group: m.Group.Proto(),
		Chain: m.Chain.Proto(),

		Locale:             m.Locale,
		Language:           m.Language,
		Contact:            m.Contact.Proto(),
		Location:           m.Location.Proto(),
		Timezone:           m.Timezone.Proto(),
		OpeningHourWeek:    m.OpeningHourWeek,
		ClosingHourWeek:    m.ClosingHourWeek,
		OpeningHourWeekend: m.OpeningHourWeekend,
		ClosingHourWeekend: m.ClosingHourWeekend,

		WeightUnits:              m.WeightUnits,
		WeightDecimalPlaces:      m.WeightDecimalPlaces,
		WeightDecimalSeparator:   m.WeightDecimalSeparator,
		WeightThousandsSeparator: m.WeightThousandsSeparator,

		DimensionUnits:              m.DimensionUnits,
		DimensionDecimalPlaces:      m.DimensionDecimalPlaces,
		DimensionDecimalSeparator:   m.DimensionDecimalSeparator,
		DimensionThousandsSeparator: m.DimensionThousandsSeparator,

		Currency:                   m.Currency,
		CurrencySymbol:             m.CurrencySymbol,
		CurrencySymbolLocation:     m.CurrencySymbolLocation,
		CurrencyDecimalPlaces:      m.CurrencyDecimalPlaces,
		CurrencyDecimalSeparator:   m.CurrencyDecimalSeparator,
		CurrencyThousandsSeparator: m.CurrencyThousandsSeparator,

		DecimalPlaces:      m.DecimalPlaces,
		DecimalSeparator:   m.DecimalSeparator,
		ThousandsSeparator: m.ThousandsSeparator,

		ExternalCode: m.ExternalCode,
	}
}

type StoreList []Store

// NewStoreListWithIDs returns a new entity list with defined identity fields
func NewStoreListWithIDs(ids ...string) StoreList {
	list := make([]Store, 0, len(ids))
	for _, id := range ids {
		list = append(list, NewStoreWithID(id))
	}
	return list
}

// NewStoreListFromProto converts object from a proto buffer definition
func NewStoreListFromProto(list []*store_pb.Store) StoreList {
	if list == nil {
		return nil
	}
	value := make(StoreList, 0, len(list))
	for _, v := range list {
		value = append(value, NewStoreFromProto(v))
	}
	return value
}

// Patch merges object values with another of the same type
func (m StoreList) Patch(dst StoreList) (StoreList, error) {
	if m == nil {
		return dst, nil
	}
	if dst == nil {
		return m.Clone(), nil
	}
	dict := map[string]*Store{}
	for i, v := range dst {
		dict[v.ID] = &dst[i]
	}
	for _, v := range m {
		if r, ok := dict[v.ID]; !ok {
			dst = append(dst, r.Clone())
		} else if err := v.Patch(r); err != nil {
			return nil, err
		}
	}
	return dst, nil
}

// Clone creates a new object copying any pointer values to a new memory space
func (m StoreList) Clone() StoreList {
	if m == nil {
		return nil
	}
	list := make(StoreList, 0, len(m))
	for _, v := range m {
		list = append(list, v.Clone())
	}
	return list
}

// Proto returns the proto buffer equivalent of this object
func (m StoreList) Proto() []*store_pb.Store {
	if m == nil {
		return nil
	}
	pbList := make([]*store_pb.Store, len(m))
	for i, v := range m {
		pbList[i] = v.Proto()
	}
	return pbList
}

type Group struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name"`
}

// NewGroupWithID returns a new instance with an id already defined
func NewGroupWithID(id string) Group {
	return Group{ID: id}
}

// NewGroupFromProto converts object from a proto buffer definition
func NewGroupFromProto(pbv *store_pb.Group) Group {
	if pbv == nil {
		return Group{}
	}
	return Group{
		ID:   pbv.GetId(),
		Name: pbv.GetName(),
	}
}

// UnmarshalJSON custom unmarshaler to allow only strings as ids
func (m *Group) UnmarshalJSON(data []byte) error {
	values := map[string]string{}
	if err := json.Unmarshal(data, &values); err != nil {
		if e, ok := err.(*json.UnmarshalTypeError); ok && e.Value == "string" {
			return json.Unmarshal(data, &m.ID)
		}
		return err
	}
	m.ID = values["id"]
	m.Name = values["name"]
	return nil
}

// Patch sets any non-zero value to the correspondent entity value
func (m Group) Patch(dst *Group) error {
	if dst == nil {
		return nil
	}
	if m.ID != "" {
		dst.ID = m.ID
	}
	if m.Name != "" {
		dst.Name = m.Name
	}
	return nil
}

// Clone creates a new instance copying any pointer value to new memory space
func (m Group) Clone() Group {
	return m
}

// Proto returns the proto buffer equivalent of this object
func (m Group) Proto() *store_pb.Group {
	return &store_pb.Group{
		Id:   m.ID,
		Name: m.Name,
	}
}

type Chain struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name"`
}

// NewChainWithID returns a new instance with an id already defined
func NewChainWithID(id string) Chain {
	return Chain{ID: id}
}

// NewChainFromProto converts object from a proto buffer definition
func NewChainFromProto(pbv *store_pb.Chain) Chain {
	if pbv == nil {
		return Chain{}
	}
	return Chain{
		ID:   pbv.GetId(),
		Name: pbv.GetName(),
	}
}

// UnmarshalJSON custom unmarshaler to allow only strings as ids
func (m *Chain) UnmarshalJSON(data []byte) error {
	values := map[string]string{}
	if err := json.Unmarshal(data, &values); err != nil {
		if e, ok := err.(*json.UnmarshalTypeError); ok && e.Value == "string" {
			return json.Unmarshal(data, &m.ID)
		}
		return err
	}
	m.ID = values["id"]
	m.Name = values["name"]
	return nil
}

// Patch sets any non-zero value to the correspondent entity value
func (m Chain) Patch(dst *Chain) error {
	if dst == nil {
		return nil
	}

	if m.ID != "" {
		dst.ID = m.ID
	}
	if m.Name != "" {
		dst.Name = m.Name
	}
	return nil
}

// Clone creates a new instance copying any pointer value to new memory space
func (m Chain) Clone() Chain {
	return m
}

// Proto returns the proto buffer equivalent of this object
func (m Chain) Proto() *store_pb.Chain {
	return &store_pb.Chain{
		Id:   m.ID,
		Name: m.Name,
	}
}

// Contact represents a contact information of a store
type Contact struct {
	Phone          string `json:"phone,omitempty"`           // Phone is the phone number of a store
	FormattedPhone string `json:"formatted_phone,omitempty"` // FormattedPhone is a user friendly phone number
}

// NewContact returns a new contact
func NewContact() Contact {
	return Contact{}
}

// NewContactFromProto converts object from a proto buffer definition
func NewContactFromProto(pbv *store_pb.Contact) Contact {
	if pbv == nil {
		return Contact{}
	}
	return Contact{
		Phone:          pbv.GetPhone(),
		FormattedPhone: pbv.GetFormattedPhone(),
	}
}

// Patch sets any non-zero value to the correspondent entity value
func (m Contact) Patch(dst *Contact) error {
	if dst == nil {
		return nil
	}
	if m.Phone != "" {
		dst.Phone = m.Phone
	}
	if m.FormattedPhone != "" {
		dst.FormattedPhone = m.FormattedPhone
	}
	return nil
}

// Clone creates a new instance copying any pointer value to new memory space
func (m Contact) Clone() Contact {
	return m
}

// Proto returns the proto buffer equivalent of this object
func (m Contact) Proto() *store_pb.Contact {
	return &store_pb.Contact{
		Phone:          m.Phone,
		FormattedPhone: m.FormattedPhone,
	}
}

// Coordinates represents the location information of a store
type Coordinates struct {
	Lat decimal.Decimal `json:"lat,omitempty"`
	Lon decimal.Decimal `json:"lon,omitempty"`
}

var _ GeoPoint = (*Location)(nil)

// NewCoordinates returns a new instance
func NewCoordinates(lat, lon decimal.Decimal) Coordinates {
	return Coordinates{Lat: lat, Lon: lon}
}

// NewCoordinatesFromPoint returns a new coordinates instance from an existing geo point
func NewCoordinatesFromPoint(point GeoPoint) Coordinates {
	return Coordinates{Lat: point.Latitude(), Lon: point.Longitude()}
}

// NewCoordinatesFromProto converts object from a proto buffer definition
func NewCoordinatesFromProto(pbv *store_pb.Coordinates) Coordinates {
	c := Coordinates{}
	if val, err := decimal.NewFromString(pbv.GetLat()); err == nil {
		c.Lat = val
	}
	if val, err := decimal.NewFromString(pbv.GetLon()); err == nil {
		c.Lon = val
	}
	return c
}

// Latitude returns the latitude value
func (m Coordinates) Latitude() decimal.Decimal {
	return m.Lat
}

// Longitude returns the longitude value
func (m Coordinates) Longitude() decimal.Decimal {
	return m.Lon
}

// Patch sets any non-zero value to the correspondent entity value
func (m Coordinates) Patch(dst *Coordinates) error {
	if dst == nil {
		return nil
	}
	if !m.Lat.IsZero() {
		dst.Lat = decimal.New(m.Lat.Coefficient().Int64(), m.Lat.Exponent())
	}
	if !m.Lon.IsZero() {
		dst.Lon = decimal.New(m.Lon.Coefficient().Int64(), m.Lon.Exponent())
	}
	return nil
}

// Clone creates a new instance copying any pointer value to new memory space
func (m Coordinates) Clone() Coordinates {
	m.Lat = decimal.New(m.Lat.Coefficient().Int64(), m.Lat.Exponent())
	m.Lon = decimal.New(m.Lon.Coefficient().Int64(), m.Lon.Exponent())
	return m
}

// Proto returns the proto buffer equivalent of this object
func (m Coordinates) Proto() *store_pb.Coordinates {
	return &store_pb.Coordinates{
		Lat: m.Longitude().String(),
		Lon: m.Longitude().String(),
	}
}

// Location represents a location information of a store
type Location struct {
	Cc          string      `json:"cc,omitempty"`
	Country     string      `json:"country,omitempty"`
	State       string      `json:"state,omitempty"`
	City        string      `json:"city,omitempty"`
	PostalCode  string      `json:"postal_code,omitempty"`
	CrossStreet string      `json:"cross_street,omitempty"`
	Address     string      `json:"address,omitempty"`
	Coords      Coordinates `json:"coords,omitempty"`
}

var _ GeoPoint = (*Location)(nil)

// NewLocation returns a new instance
func NewLocation() Location {
	return Location{}
}

// NewLocationFromProto converts object from a proto buffer definition
func NewLocationFromProto(pbv *store_pb.Location) Location {
	if pbv == nil {
		return Location{}
	}
	return Location{
		Cc:          pbv.GetCc(),
		Country:     pbv.GetCountry(),
		State:       pbv.GetState(),
		City:        pbv.GetCity(),
		PostalCode:  pbv.GetPostalCode(),
		CrossStreet: pbv.GetCrossStreet(),
		Address:     pbv.GetAddress(),
		Coords:      NewCoordinatesFromProto(pbv.GetCoords()),
	}
}

// Latitude returns the latitude a store
func (m Location) Latitude() decimal.Decimal {
	return m.Coords.Latitude()
}

// Longitude returns the longitude a store
func (m Location) Longitude() decimal.Decimal {
	return m.Coords.Longitude()
}

// Patch sets any non-zero value to the correspondent entity value
func (m Location) Patch(dst *Location) error {
	if dst == nil {
		return nil
	}

	if m.Cc != "" {
		dst.Cc = m.Cc
	}

	if m.Country != "" {
		dst.Country = m.Country
	}

	if m.State != "" {
		dst.State = m.State
	}

	if m.City != "" {
		dst.City = m.City
	}
	if m.PostalCode != "" {
		dst.PostalCode = m.PostalCode
	}
	if m.CrossStreet != "" {
		dst.CrossStreet = m.CrossStreet
	}
	if m.Address != "" {
		dst.Address = m.Address
	}
	if err := m.Coords.Patch(&dst.Coords); err != nil {
		return err
	}
	return nil
}

// Clone creates a new instance copying any pointer value to new memory space
func (m Location) Clone() Location {
	m.Coords = m.Coords.Clone()
	return m
}

// Proto returns the proto buffer equivalent of this object
func (m Location) Proto() *store_pb.Location {
	return &store_pb.Location{
		Cc:          m.Cc,
		Country:     m.Country,
		State:       m.State,
		City:        m.City,
		PostalCode:  m.PostalCode,
		CrossStreet: m.CrossStreet,
		Address:     m.Address,
		Coords:      m.Coords.Proto(),
	}
}

// Timezone contains timezone information about a store
type Timezone struct {
	Value      string     `json:"value,omitempty" validate:"required"`
	DateFormat DateFormat `json:"date_format,omitempty" validate:"dive"`
}

// NewTimezone returns a new instance
func NewTimezone() Timezone {
	return Timezone{}
}

// NewTimezoneFromProto converts object from a proto buffer definition
func NewTimezoneFromProto(pbv *store_pb.Timezone) Timezone {
	if pbv == nil {
		return Timezone{}
	}
	return Timezone{
		Value:      pbv.GetValue(),
		DateFormat: NewDateFormatFromProto(pbv.GetDateFormat()),
	}
}

// Patch sets any non-zero value to the correspondent entity value
func (m Timezone) Patch(dst *Timezone) error {
	if dst == nil {
		return nil
	}
	if m.Value != "" {
		dst.Value = m.Value
	}
	if err := m.DateFormat.Patch(&dst.DateFormat); err != nil {
		return err
	}
	return nil
}

// Clone creates a new instance copying any pointer value to new memory space
func (m Timezone) Clone() Timezone {
	m.DateFormat = m.DateFormat.Clone()
	return m
}

// Proto returns the proto buffer equivalent of this object
func (m Timezone) Proto() *store_pb.Timezone {
	return &store_pb.Timezone{
		Value:      m.Value,
		DateFormat: m.DateFormat.Proto(),
	}
}

// DateFormat contains instruction on how to format a date for a store
type DateFormat struct {
	Display string `json:"display,omitempty" validate:"required"`
}

// NewDateFormat returns a new instance
func NewDateFormat() DateFormat {
	return DateFormat{}
}

// NewDateFormatFromProto converts object from a proto buffer definition
func NewDateFormatFromProto(pbv *store_pb.DateFormat) DateFormat {
	if pbv == nil {
		return DateFormat{}
	}
	return DateFormat{
		Display: pbv.GetDisplay(),
	}
}

// Patch sets any non-zero value to the correspondent entity value
func (m DateFormat) Patch(dst *DateFormat) error {
	if dst == nil {
		return nil
	}
	if m.Display != "" {
		dst.Display = m.Display
	}
	return nil
}

// Clone creates a new instance copying any pointer value to new memory space
func (m DateFormat) Clone() DateFormat {
	return m
}

// Proto returns the proto buffer equivalent of this object
func (m DateFormat) Proto() *store_pb.DateFormat {
	return &store_pb.DateFormat{
		Display: m.Display,
	}
}
