// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/param"
)

type Address struct {
	City   string      `json:"city"`
	Geo    AddressGeo  `json:"geo"`
	State  string      `json:"state"`
	Street string      `json:"street"`
	Zip    string      `json:"zip"`
	JSON   addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	City        apijson.Field
	Geo         apijson.Field
	State       apijson.Field
	Street      apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

type AddressGeo struct {
	Latitude  float64        `json:"latitude" api:"required"`
	Longitude float64        `json:"longitude" api:"required"`
	JSON      addressGeoJSON `json:"-"`
}

// addressGeoJSON contains the JSON metadata for the struct [AddressGeo]
type addressGeoJSON struct {
	Latitude    apijson.Field
	Longitude   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressGeoJSON) RawJSON() string {
	return r.raw
}

type AddressParam struct {
	City   param.Field[string]          `json:"city"`
	Geo    param.Field[AddressGeoParam] `json:"geo"`
	State  param.Field[string]          `json:"state"`
	Street param.Field[string]          `json:"street"`
	Zip    param.Field[string]          `json:"zip"`
}

func (r AddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressGeoParam struct {
	Latitude  param.Field[float64] `json:"latitude" api:"required"`
	Longitude param.Field[float64] `json:"longitude" api:"required"`
}

func (r AddressGeoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Money struct {
	// Amount in minor currency units (e.g. cents).
	Amount int64 `json:"amount" api:"required"`
	// ISO 4217 currency code.
	Currency string    `json:"currency" api:"required"`
	JSON     moneyJSON `json:"-"`
}

// moneyJSON contains the JSON metadata for the struct [Money]
type moneyJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Money) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r moneyJSON) RawJSON() string {
	return r.raw
}

type MoneyParam struct {
	// Amount in minor currency units (e.g. cents).
	Amount param.Field[int64] `json:"amount" api:"required"`
	// ISO 4217 currency code.
	Currency param.Field[string] `json:"currency" api:"required"`
}

func (r MoneyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
