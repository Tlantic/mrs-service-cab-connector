package stores

import (
	"strings"

	"github.com/shopspring/decimal"
)

type GeoPoint interface {
	Latitude() decimal.Decimal
	Longitude() decimal.Decimal
}

type GeoPointString string

func (m GeoPointString) Latitude() decimal.Decimal {
	parts := strings.Split(string(m), ",")
	if len(parts) != 2 {
		return decimal.Zero
	}
	c, err := decimal.NewFromString(parts[0])
	if err != nil {
		return decimal.Zero
	}
	return c
}
func (m GeoPointString) Longitude() decimal.Decimal {
	parts := strings.Split(string(m), ",")
	if len(parts) != 2 {
		return decimal.Zero
	}
	c, err := decimal.NewFromString(parts[1])
	if err != nil {
		return decimal.Zero
	}
	return c
}

type GeoPointStringSlice []string

func (m GeoPointStringSlice) Latitude() decimal.Decimal {
	if len(m) != 2 {
		return decimal.Zero
	}

	c, err := decimal.NewFromString(m[0])
	if err != nil {
		return decimal.Zero
	}
	return c
}
func (m GeoPointStringSlice) Longitude() decimal.Decimal {
	if len(m) != 2 {
		return decimal.Zero
	}

	c, err := decimal.NewFromString(m[1])
	if err != nil {
		return decimal.Zero
	}
	return c
}

type GeoPointFloat [2]float64

func (m GeoPointFloat) Latitude() decimal.Decimal {
	return decimal.NewFromFloat(m[0])
}
func (m GeoPointFloat) Longitude() decimal.Decimal {
	return decimal.NewFromFloat(m[1])
}

type GeoPointDecimal [2]decimal.Decimal

func (m GeoPointDecimal) Latitude() decimal.Decimal {
	return m[0]
}
func (m GeoPointDecimal) Longitude() decimal.Decimal {
	return m[1]
}
