// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
)

// Access to Petstore orders
//
// StoreOrderService contains methods and other services that help with interacting
// with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreOrderService] method instead.
type StoreOrderService struct {
	Options []option.RequestOption
}

// NewStoreOrderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStoreOrderService(opts ...option.RequestOption) (r StoreOrderService) {
	r = StoreOrderService{}
	r.Options = opts
	return
}

// Place a new order in the store
func (r *StoreOrderService) New(ctx context.Context, body StoreOrderNewParams, opts ...option.RequestOption) (res *StoreOrderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "store/order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// For valid response try integer IDs with value <= 5 or > 10. Other values will
// generate exceptions.
func (r *StoreOrderService) Get(ctx context.Context, orderID int64, opts ...option.RequestOption) (res *StoreOrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("store/order/%v", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// For valid response try integer IDs with value < 1000. Anything above 1000 or
// nonintegers will generate API errors
func (r *StoreOrderService) Delete(ctx context.Context, orderID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("store/order/%v", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type StoreOrderNewResponse struct {
	ID       int64     `json:"id"`
	Complete bool      `json:"complete"`
	PetID    int64     `json:"petId"`
	Quantity int64     `json:"quantity"`
	ShipDate time.Time `json:"shipDate" format:"date-time"`
	// Order Status
	//
	// Any of "placed", "approved", "delivered".
	Status StoreOrderNewResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Complete    respjson.Field
		PetID       respjson.Field
		Quantity    respjson.Field
		ShipDate    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoreOrderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StoreOrderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Order Status
type StoreOrderNewResponseStatus string

const (
	StoreOrderNewResponseStatusPlaced    StoreOrderNewResponseStatus = "placed"
	StoreOrderNewResponseStatusApproved  StoreOrderNewResponseStatus = "approved"
	StoreOrderNewResponseStatusDelivered StoreOrderNewResponseStatus = "delivered"
)

type StoreOrderGetResponse struct {
	ID       int64     `json:"id"`
	Complete bool      `json:"complete"`
	PetID    int64     `json:"petId"`
	Quantity int64     `json:"quantity"`
	ShipDate time.Time `json:"shipDate" format:"date-time"`
	// Order Status
	//
	// Any of "placed", "approved", "delivered".
	Status StoreOrderGetResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Complete    respjson.Field
		PetID       respjson.Field
		Quantity    respjson.Field
		ShipDate    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoreOrderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StoreOrderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Order Status
type StoreOrderGetResponseStatus string

const (
	StoreOrderGetResponseStatusPlaced    StoreOrderGetResponseStatus = "placed"
	StoreOrderGetResponseStatusApproved  StoreOrderGetResponseStatus = "approved"
	StoreOrderGetResponseStatusDelivered StoreOrderGetResponseStatus = "delivered"
)

type StoreOrderNewParams struct {
	ID       param.Opt[int64]     `json:"id,omitzero"`
	Complete param.Opt[bool]      `json:"complete,omitzero"`
	PetID    param.Opt[int64]     `json:"petId,omitzero"`
	Quantity param.Opt[int64]     `json:"quantity,omitzero"`
	ShipDate param.Opt[time.Time] `json:"shipDate,omitzero" format:"date-time"`
	// Order Status
	//
	// Any of "placed", "approved", "delivered".
	Status StoreOrderNewParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r StoreOrderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StoreOrderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StoreOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Order Status
type StoreOrderNewParamsStatus string

const (
	StoreOrderNewParamsStatusPlaced    StoreOrderNewParamsStatus = "placed"
	StoreOrderNewParamsStatusApproved  StoreOrderNewParamsStatus = "approved"
	StoreOrderNewParamsStatusDelivered StoreOrderNewParamsStatus = "delivered"
)
