// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/swagger-girly/go"
	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Address struct {
	City            string                    `json:"city"`
	Geo             AddressGeo                `json:"geo"`
	RelatedCategory any                       `json:"relatedCategory"`
	RelatedCustomer any                       `json:"relatedCustomer"`
	RelatedMoney    *Money                    `json:"relatedMoney"`
	RelatedOrder    any                       `json:"relatedOrder"`
	RelatedPet      *helloworldtestingggg.Pet `json:"relatedPet"`
	RelatedShelter  any                       `json:"relatedShelter"`
	RelatedTag      any                       `json:"relatedTag"`
	RelatedUser     any                       `json:"relatedUser"`
	State           string                    `json:"state"`
	Street          string                    `json:"street"`
	Zip             string                    `json:"zip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City            respjson.Field
		Geo             respjson.Field
		RelatedCategory respjson.Field
		RelatedCustomer respjson.Field
		RelatedMoney    respjson.Field
		RelatedOrder    respjson.Field
		RelatedPet      respjson.Field
		RelatedShelter  respjson.Field
		RelatedTag      respjson.Field
		RelatedUser     respjson.Field
		State           respjson.Field
		Street          respjson.Field
		Zip             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Address) RawJSON() string { return r.JSON.raw }
func (r *Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Address to a AddressParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AddressParam.Overrides()
func (r Address) ToParam() AddressParam {
	return param.Override[AddressParam](json.RawMessage(r.RawJSON()))
}

type AddressGeo struct {
	Latitude  float64 `json:"latitude" api:"required"`
	Longitude float64 `json:"longitude" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressGeo) RawJSON() string { return r.JSON.raw }
func (r *AddressGeo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressParam struct {
	City            param.Opt[string]             `json:"city,omitzero"`
	State           param.Opt[string]             `json:"state,omitzero"`
	Street          param.Opt[string]             `json:"street,omitzero"`
	Zip             param.Opt[string]             `json:"zip,omitzero"`
	Geo             AddressGeoParam               `json:"geo,omitzero"`
	RelatedCategory any                           `json:"relatedCategory,omitzero"`
	RelatedCustomer any                           `json:"relatedCustomer,omitzero"`
	RelatedMoney    MoneyParam                    `json:"relatedMoney,omitzero"`
	RelatedOrder    any                           `json:"relatedOrder,omitzero"`
	RelatedPet      helloworldtestingggg.PetParam `json:"relatedPet,omitzero"`
	RelatedShelter  any                           `json:"relatedShelter,omitzero"`
	RelatedTag      any                           `json:"relatedTag,omitzero"`
	RelatedUser     any                           `json:"relatedUser,omitzero"`
	paramObj
}

func (r AddressParam) MarshalJSON() (data []byte, err error) {
	type shadow AddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Latitude, Longitude are required.
type AddressGeoParam struct {
	Latitude  float64 `json:"latitude" api:"required"`
	Longitude float64 `json:"longitude" api:"required"`
	paramObj
}

func (r AddressGeoParam) MarshalJSON() (data []byte, err error) {
	type shadow AddressGeoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressGeoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Money struct {
	// Amount in minor currency units (e.g. cents).
	Amount int64 `json:"amount" api:"required"`
	// ISO 4217 currency code.
	Currency        string                    `json:"currency" api:"required"`
	RelatedAddress  *Address                  `json:"relatedAddress"`
	RelatedCategory any                       `json:"relatedCategory"`
	RelatedCustomer any                       `json:"relatedCustomer"`
	RelatedOrder    any                       `json:"relatedOrder"`
	RelatedPet      *helloworldtestingggg.Pet `json:"relatedPet"`
	RelatedShelter  any                       `json:"relatedShelter"`
	RelatedTag      any                       `json:"relatedTag"`
	RelatedUser     any                       `json:"relatedUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount          respjson.Field
		Currency        respjson.Field
		RelatedAddress  respjson.Field
		RelatedCategory respjson.Field
		RelatedCustomer respjson.Field
		RelatedOrder    respjson.Field
		RelatedPet      respjson.Field
		RelatedShelter  respjson.Field
		RelatedTag      respjson.Field
		RelatedUser     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Money) RawJSON() string { return r.JSON.raw }
func (r *Money) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Money to a MoneyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MoneyParam.Overrides()
func (r Money) ToParam() MoneyParam {
	return param.Override[MoneyParam](json.RawMessage(r.RawJSON()))
}

// The properties Amount, Currency are required.
type MoneyParam struct {
	// Amount in minor currency units (e.g. cents).
	Amount int64 `json:"amount" api:"required"`
	// ISO 4217 currency code.
	Currency        string                        `json:"currency" api:"required"`
	RelatedAddress  AddressParam                  `json:"relatedAddress,omitzero"`
	RelatedCategory any                           `json:"relatedCategory,omitzero"`
	RelatedCustomer any                           `json:"relatedCustomer,omitzero"`
	RelatedOrder    any                           `json:"relatedOrder,omitzero"`
	RelatedPet      helloworldtestingggg.PetParam `json:"relatedPet,omitzero"`
	RelatedShelter  any                           `json:"relatedShelter,omitzero"`
	RelatedTag      any                           `json:"relatedTag,omitzero"`
	RelatedUser     any                           `json:"relatedUser,omitzero"`
	paramObj
}

func (r MoneyParam) MarshalJSON() (data []byte, err error) {
	type shadow MoneyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MoneyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
