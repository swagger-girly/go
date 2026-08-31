// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/respjson"
	"github.com/swagger-girly/go/shared"
)

// Access to Petstore orders
//
// StoreService contains methods and other services that help with interacting with
// the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreService] method instead.
type StoreService struct {
	options []option.RequestOption
	// Access to Petstore orders
	Order StoreOrderService
	// Access to Petstore orders
	Reports StoreReportService
}

// NewStoreService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStoreService(opts ...option.RequestOption) (r StoreService) {
	r = StoreService{}
	r.options = opts
	r.Order = NewStoreOrderService(opts...)
	r.Reports = NewStoreReportService(opts...)
	return
}

// Returns a map of status codes to quantities
func (r *StoreService) ListInventory(ctx context.Context, opts ...option.RequestOption) (res *StoreListInventoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "store/inventory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a single activity feed entry. The response is a discriminated union
// whose arms each compose the public `Pet` model, so every arm is minted under the
// response's namespace and must inherit rather than flatten.
func (r *StoreService) GetActivity(ctx context.Context, activityID string, opts ...option.RequestOption) (res *StoreGetActivityResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if activityID == "" {
		err = errors.New("missing required activityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("store/activity/%s", url.PathEscape(activityID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type StoreListInventoryResponse map[string]int64

// StoreGetActivityResponseUnion contains all possible properties and values from
// [StoreGetActivityResponseListed], [StoreGetActivityResponseSold],
// [StoreGetActivityResponseArchived], [StoreGetActivityResponseTransferred].
//
// Use the [StoreGetActivityResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type StoreGetActivityResponseUnion struct {
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	Name string `json:"name"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	PhotoURLs []string `json:"photoUrls"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	ID int64 `json:"id"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	AcquisitionChannel PetAcquisitionChannel `json:"acquisitionChannel"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	Category any `json:"category"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	MicrochipID PetMicrochipIDUnion `json:"microchipId"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	RelatedAddress shared.Address `json:"relatedAddress"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	RelatedCategory any `json:"relatedCategory"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	RelatedCustomer any `json:"relatedCustomer"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	RelatedMoney shared.Money `json:"relatedMoney"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	RelatedOrder any `json:"relatedOrder"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	RelatedPet Pet `json:"relatedPet"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	RelatedShelter any `json:"relatedShelter"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	RelatedUser any    `json:"relatedUser"`
	Status      string `json:"status"`
	// This field is from variant [StoreGetActivityResponseListed],
	// [StoreGetActivityResponseSold], [StoreGetActivityResponseArchived],
	// [StoreGetActivityResponseTransferred].
	Tags       []any  `json:"tags"`
	ActivityID string `json:"activityId"`
	// Any of "listed", "sold", "archived", "transferred".
	Family     string    `json:"family"`
	RecordedAt time.Time `json:"recordedAt"`
	// This field is from variant [StoreGetActivityResponseArchived].
	ArchivedReason string `json:"archivedReason"`
	// This field is from variant [StoreGetActivityResponseTransferred].
	TransferredTo StoreGetActivityResponseTransferredTransferredTo `json:"transferredTo"`
	JSON          struct {
		Name               respjson.Field
		PhotoURLs          respjson.Field
		ID                 respjson.Field
		AcquisitionChannel respjson.Field
		Category           respjson.Field
		MicrochipID        respjson.Field
		RelatedAddress     respjson.Field
		RelatedCategory    respjson.Field
		RelatedCustomer    respjson.Field
		RelatedMoney       respjson.Field
		RelatedOrder       respjson.Field
		RelatedPet         respjson.Field
		RelatedShelter     respjson.Field
		RelatedUser        respjson.Field
		Status             respjson.Field
		Tags               respjson.Field
		ActivityID         respjson.Field
		Family             respjson.Field
		RecordedAt         respjson.Field
		ArchivedReason     respjson.Field
		TransferredTo      respjson.Field
		raw                string
	} `json:"-"`
}

// anyStoreGetActivityResponse is implemented by each variant of
// [StoreGetActivityResponseUnion] to add type safety for the return type of
// [StoreGetActivityResponseUnion.AsAny]
type anyStoreGetActivityResponse interface {
	implStoreGetActivityResponseUnion()
}

func (StoreGetActivityResponseListed) implStoreGetActivityResponseUnion()      {}
func (StoreGetActivityResponseSold) implStoreGetActivityResponseUnion()        {}
func (StoreGetActivityResponseArchived) implStoreGetActivityResponseUnion()    {}
func (StoreGetActivityResponseTransferred) implStoreGetActivityResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := StoreGetActivityResponseUnion.AsAny().(type) {
//	case helloworldtestingggg.StoreGetActivityResponseListed:
//	case helloworldtestingggg.StoreGetActivityResponseSold:
//	case helloworldtestingggg.StoreGetActivityResponseArchived:
//	case helloworldtestingggg.StoreGetActivityResponseTransferred:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u StoreGetActivityResponseUnion) AsAny() anyStoreGetActivityResponse {
	switch u.Family {
	case "listed":
		return u.AsListed()
	case "sold":
		return u.AsSold()
	case "archived":
		return u.AsArchived()
	case "transferred":
		return u.AsTransferred()
	}
	return nil
}

func (u StoreGetActivityResponseUnion) AsListed() (v StoreGetActivityResponseListed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StoreGetActivityResponseUnion) AsSold() (v StoreGetActivityResponseSold) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StoreGetActivityResponseUnion) AsArchived() (v StoreGetActivityResponseArchived) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StoreGetActivityResponseUnion) AsTransferred() (v StoreGetActivityResponseTransferred) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u StoreGetActivityResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *StoreGetActivityResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Listing activity. Widens the base's `status` enum with a member `Pet` does not
// declare, so the heritage clause has to omit the redeclared key instead of
// extending the base plainly.
type StoreGetActivityResponseListed struct {
	ActivityID string `json:"activityId" api:"required"`
	// Any of "listed".
	Family     string    `json:"family" api:"required"`
	RecordedAt time.Time `json:"recordedAt" api:"required" format:"date-time"`
	// Any of "available", "pending", "sold", "archived".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActivityID  respjson.Field
		Family      respjson.Field
		RecordedAt  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Pet
}

// Returns the unmodified JSON received from the API
func (r StoreGetActivityResponseListed) RawJSON() string { return r.JSON.raw }
func (r *StoreGetActivityResponseListed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sale activity. Pins the base's `status` to a single member — a provable
// narrowing, so this arm keeps a plain `extends` with no omitted keys.
type StoreGetActivityResponseSold struct {
	ActivityID string `json:"activityId" api:"required"`
	// Any of "sold".
	Family     string    `json:"family" api:"required"`
	RecordedAt time.Time `json:"recordedAt" api:"required" format:"date-time"`
	// Any of "sold".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActivityID  respjson.Field
		Family      respjson.Field
		RecordedAt  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Pet
}

// Returns the unmodified JSON received from the API
func (r StoreGetActivityResponseSold) RawJSON() string { return r.JSON.raw }
func (r *StoreGetActivityResponseSold) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Archival activity. Narrows the base's `status` by value but also makes it
// nullable, which `Pet` does not accept — a value subset alone is not enough to
// keep a plain `extends` here.
type StoreGetActivityResponseArchived struct {
	ActivityID string `json:"activityId" api:"required"`
	// Any of "archived".
	Family         string    `json:"family" api:"required"`
	RecordedAt     time.Time `json:"recordedAt" api:"required" format:"date-time"`
	ArchivedReason string    `json:"archivedReason"`
	// Any of "sold".
	Status string `json:"status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActivityID     respjson.Field
		Family         respjson.Field
		RecordedAt     respjson.Field
		ArchivedReason respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
	Pet
}

// Returns the unmodified JSON received from the API
func (r StoreGetActivityResponseArchived) RawJSON() string { return r.JSON.raw }
func (r *StoreGetActivityResponseArchived) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transfer activity. Reaches `Pet` only through the pass-through wrapper branch,
// so the base is resolvable only transitively.
type StoreGetActivityResponseTransferred struct {
	// Any of "transferred".
	Family        string                                           `json:"family" api:"required"`
	TransferredTo StoreGetActivityResponseTransferredTransferredTo `json:"transferredTo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Family        respjson.Field
		TransferredTo respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	Pet
}

// Returns the unmodified JSON received from the API
func (r StoreGetActivityResponseTransferred) RawJSON() string { return r.JSON.raw }
func (r *StoreGetActivityResponseTransferred) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoreGetActivityResponseTransferredTransferredTo struct {
	Name            string                                                   `json:"name" api:"required"`
	Address         shared.Address                                           `json:"address"`
	Contact         StoreGetActivityResponseTransferredTransferredToContact  `json:"contact"`
	Location        StoreGetActivityResponseTransferredTransferredToLocation `json:"location"`
	RelatedAddress  shared.Address                                           `json:"relatedAddress"`
	RelatedCategory any                                                      `json:"relatedCategory"`
	RelatedCustomer any                                                      `json:"relatedCustomer"`
	RelatedMoney    shared.Money                                             `json:"relatedMoney"`
	RelatedOrder    any                                                      `json:"relatedOrder"`
	RelatedPet      Pet                                                      `json:"relatedPet"`
	RelatedTag      any                                                      `json:"relatedTag"`
	RelatedUser     any                                                      `json:"relatedUser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		Address         respjson.Field
		Contact         respjson.Field
		Location        respjson.Field
		RelatedAddress  respjson.Field
		RelatedCategory respjson.Field
		RelatedCustomer respjson.Field
		RelatedMoney    respjson.Field
		RelatedOrder    respjson.Field
		RelatedPet      respjson.Field
		RelatedTag      respjson.Field
		RelatedUser     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoreGetActivityResponseTransferredTransferredTo) RawJSON() string { return r.JSON.raw }
func (r *StoreGetActivityResponseTransferredTransferredTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoreGetActivityResponseTransferredTransferredToContact struct {
	Primary StoreGetActivityResponseTransferredTransferredToContactPrimary `json:"primary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoreGetActivityResponseTransferredTransferredToContact) RawJSON() string { return r.JSON.raw }
func (r *StoreGetActivityResponseTransferredTransferredToContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoreGetActivityResponseTransferredTransferredToContactPrimary struct {
	Availability StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailability `json:"availability"`
	Name         string                                                                     `json:"name"`
	Phone        string                                                                     `json:"phone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Availability respjson.Field
		Name         respjson.Field
		Phone        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoreGetActivityResponseTransferredTransferredToContactPrimary) RawJSON() string {
	return r.JSON.raw
}
func (r *StoreGetActivityResponseTransferredTransferredToContactPrimary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailability struct {
	Weekday StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailabilityWeekday `json:"weekday"`
	Weekend StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailabilityWeekend `json:"weekend" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Weekday     respjson.Field
		Weekend     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailability) RawJSON() string {
	return r.JSON.raw
}
func (r *StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailabilityWeekday struct {
	Close string `json:"close"`
	Open  string `json:"open"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close       respjson.Field
		Open        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailabilityWeekday) RawJSON() string {
	return r.JSON.raw
}
func (r *StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailabilityWeekday) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailabilityWeekend struct {
	Close string `json:"close"`
	Open  string `json:"open"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close       respjson.Field
		Open        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailabilityWeekend) RawJSON() string {
	return r.JSON.raw
}
func (r *StoreGetActivityResponseTransferredTransferredToContactPrimaryAvailabilityWeekend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoreGetActivityResponseTransferredTransferredToLocation struct {
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
func (r StoreGetActivityResponseTransferredTransferredToLocation) RawJSON() string { return r.JSON.raw }
func (r *StoreGetActivityResponseTransferredTransferredToLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
