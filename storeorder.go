// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hello-world-testingggg-go/internal/apijson"
	"github.com/stainless-sdks/hello-world-testingggg-go/internal/param"
	"github.com/stainless-sdks/hello-world-testingggg-go/internal/requestconfig"
	"github.com/stainless-sdks/hello-world-testingggg-go/option"
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
func NewStoreOrderService(opts ...option.RequestOption) (r *StoreOrderService) {
	r = &StoreOrderService{}
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
	Status StoreOrderNewResponseStatus `json:"status"`
	JSON   storeOrderNewResponseJSON   `json:"-"`
}

// storeOrderNewResponseJSON contains the JSON metadata for the struct
// [StoreOrderNewResponse]
type storeOrderNewResponseJSON struct {
	ID          apijson.Field
	Complete    apijson.Field
	PetID       apijson.Field
	Quantity    apijson.Field
	ShipDate    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreOrderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeOrderNewResponseJSON) RawJSON() string {
	return r.raw
}

// Order Status
type StoreOrderNewResponseStatus string

const (
	StoreOrderNewResponseStatusPlaced    StoreOrderNewResponseStatus = "placed"
	StoreOrderNewResponseStatusApproved  StoreOrderNewResponseStatus = "approved"
	StoreOrderNewResponseStatusDelivered StoreOrderNewResponseStatus = "delivered"
)

func (r StoreOrderNewResponseStatus) IsKnown() bool {
	switch r {
	case StoreOrderNewResponseStatusPlaced, StoreOrderNewResponseStatusApproved, StoreOrderNewResponseStatusDelivered:
		return true
	}
	return false
}

type StoreOrderGetResponse struct {
	ID       int64     `json:"id"`
	Complete bool      `json:"complete"`
	PetID    int64     `json:"petId"`
	Quantity int64     `json:"quantity"`
	ShipDate time.Time `json:"shipDate" format:"date-time"`
	// Order Status
	Status StoreOrderGetResponseStatus `json:"status"`
	JSON   storeOrderGetResponseJSON   `json:"-"`
}

// storeOrderGetResponseJSON contains the JSON metadata for the struct
// [StoreOrderGetResponse]
type storeOrderGetResponseJSON struct {
	ID          apijson.Field
	Complete    apijson.Field
	PetID       apijson.Field
	Quantity    apijson.Field
	ShipDate    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreOrderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeOrderGetResponseJSON) RawJSON() string {
	return r.raw
}

// Order Status
type StoreOrderGetResponseStatus string

const (
	StoreOrderGetResponseStatusPlaced    StoreOrderGetResponseStatus = "placed"
	StoreOrderGetResponseStatusApproved  StoreOrderGetResponseStatus = "approved"
	StoreOrderGetResponseStatusDelivered StoreOrderGetResponseStatus = "delivered"
)

func (r StoreOrderGetResponseStatus) IsKnown() bool {
	switch r {
	case StoreOrderGetResponseStatusPlaced, StoreOrderGetResponseStatusApproved, StoreOrderGetResponseStatusDelivered:
		return true
	}
	return false
}

type StoreOrderNewParams struct {
	ID       param.Field[int64]     `json:"id"`
	Complete param.Field[bool]      `json:"complete"`
	PetID    param.Field[int64]     `json:"petId"`
	Quantity param.Field[int64]     `json:"quantity"`
	ShipDate param.Field[time.Time] `json:"shipDate" format:"date-time"`
	// Order Status
	Status param.Field[StoreOrderNewParamsStatus] `json:"status"`
}

func (r StoreOrderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Order Status
type StoreOrderNewParamsStatus string

const (
	StoreOrderNewParamsStatusPlaced    StoreOrderNewParamsStatus = "placed"
	StoreOrderNewParamsStatusApproved  StoreOrderNewParamsStatus = "approved"
	StoreOrderNewParamsStatusDelivered StoreOrderNewParamsStatus = "delivered"
)

func (r StoreOrderNewParamsStatus) IsKnown() bool {
	switch r {
	case StoreOrderNewParamsStatusPlaced, StoreOrderNewParamsStatusApproved, StoreOrderNewParamsStatusDelivered:
		return true
	}
	return false
}
