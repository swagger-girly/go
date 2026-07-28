// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/shared"
	"github.com/tidwall/gjson"
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
	Options []option.RequestOption
	// Access to Petstore orders
	Order *StoreOrderService
	// Access to Petstore orders
	Reports *StoreReportService
}

// NewStoreService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStoreService(opts ...option.RequestOption) (r *StoreService) {
	r = &StoreService{}
	r.Options = opts
	r.Order = NewStoreOrderService(opts...)
	r.Reports = NewStoreReportService(opts...)
	return
}

// Returns a map of status codes to quantities
func (r *StoreService) ListInventory(ctx context.Context, opts ...option.RequestOption) (res *StoreListInventoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "store/inventory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a single activity feed entry. The response is a discriminated union
// whose arms each compose the public `Pet` model, so every arm is minted under the
// response's namespace and must inherit rather than flatten.
func (r *StoreService) GetActivity(ctx context.Context, activityID string, opts ...option.RequestOption) (res *StoreGetActivityResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if activityID == "" {
		err = errors.New("missing required activityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("store/activity/%s", activityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type StoreListInventoryResponse map[string]int64

// A store activity feed entry. Every arm composes the public `Pet` model, so each
// one is minted under the response's namespace and inherits from `Pet` instead of
// re-minting its inline type tree.
type StoreGetActivityResponse struct {
	ID int64 `json:"id"`
	// This field can have the runtime type of [PetAcquisitionChannel].
	AcquisitionChannel interface{} `json:"acquisitionChannel"`
	ActivityID         string      `json:"activityId"`
	ArchivedReason     string      `json:"archivedReason"`
	// This field can have the runtime type of [PetCategory].
	Category interface{}                    `json:"category"`
	Family   StoreGetActivityResponseFamily `json:"family"`
	// This field can have the runtime type of [PetMicrochipIDUnion].
	MicrochipID interface{} `json:"microchipId"`
	Name        string      `json:"name"`
	// This field can have the runtime type of [[]string].
	PhotoURLs  interface{} `json:"photoUrls"`
	RecordedAt time.Time   `json:"recordedAt" format:"date-time"`
	// pet status in the store
	Status PetStatus `json:"status" api:"nullable"`
	// This field can have the runtime type of [[]PetTag].
	Tags interface{} `json:"tags"`
	// This field can have the runtime type of
	// [StoreGetActivityResponseStoreActivityTransferredTransferredTo].
	TransferredTo interface{}                  `json:"transferredTo"`
	JSON          storeGetActivityResponseJSON `json:"-"`
	union         StoreGetActivityResponseUnion
}

// storeGetActivityResponseJSON contains the JSON metadata for the struct
// [StoreGetActivityResponse]
type storeGetActivityResponseJSON struct {
	ID                 apijson.Field
	AcquisitionChannel apijson.Field
	ActivityID         apijson.Field
	ArchivedReason     apijson.Field
	Category           apijson.Field
	Family             apijson.Field
	MicrochipID        apijson.Field
	Name               apijson.Field
	PhotoURLs          apijson.Field
	RecordedAt         apijson.Field
	Status             apijson.Field
	Tags               apijson.Field
	TransferredTo      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r storeGetActivityResponseJSON) RawJSON() string {
	return r.raw
}

func (r *StoreGetActivityResponse) UnmarshalJSON(data []byte) (err error) {
	*r = StoreGetActivityResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StoreGetActivityResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StoreGetActivityResponseStoreActivityListed],
// [StoreGetActivityResponseStoreActivitySold],
// [StoreGetActivityResponseStoreActivityArchived],
// [StoreGetActivityResponseStoreActivityTransferred].
func (r StoreGetActivityResponse) AsUnion() StoreGetActivityResponseUnion {
	return r.union
}

// A store activity feed entry. Every arm composes the public `Pet` model, so each
// one is minted under the response's namespace and inherits from `Pet` instead of
// re-minting its inline type tree.
//
// Union satisfied by [StoreGetActivityResponseStoreActivityListed],
// [StoreGetActivityResponseStoreActivitySold],
// [StoreGetActivityResponseStoreActivityArchived] or
// [StoreGetActivityResponseStoreActivityTransferred].
type StoreGetActivityResponseUnion interface {
	implementsStoreGetActivityResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StoreGetActivityResponseUnion)(nil)).Elem(),
		"family",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StoreGetActivityResponseStoreActivityListed{}),
			DiscriminatorValue: "listed",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StoreGetActivityResponseStoreActivitySold{}),
			DiscriminatorValue: "sold",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StoreGetActivityResponseStoreActivityArchived{}),
			DiscriminatorValue: "archived",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StoreGetActivityResponseStoreActivityTransferred{}),
			DiscriminatorValue: "transferred",
		},
	)
}

// Listing activity. Widens the base's `status` enum with a member `Pet` does not
// declare, so the heritage clause has to omit the redeclared key instead of
// extending the base plainly.
type StoreGetActivityResponseStoreActivityListed struct {
	ActivityID string                                            `json:"activityId" api:"required"`
	Family     StoreGetActivityResponseStoreActivityListedFamily `json:"family" api:"required"`
	RecordedAt time.Time                                         `json:"recordedAt" api:"required" format:"date-time"`
	Status     StoreGetActivityResponseStoreActivityListedStatus `json:"status" api:"required"`
	JSON       storeGetActivityResponseStoreActivityListedJSON   `json:"-"`
	Pet
}

// storeGetActivityResponseStoreActivityListedJSON contains the JSON metadata for
// the struct [StoreGetActivityResponseStoreActivityListed]
type storeGetActivityResponseStoreActivityListedJSON struct {
	ActivityID  apijson.Field
	Family      apijson.Field
	RecordedAt  apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreGetActivityResponseStoreActivityListed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeGetActivityResponseStoreActivityListedJSON) RawJSON() string {
	return r.raw
}

func (r StoreGetActivityResponseStoreActivityListed) implementsStoreGetActivityResponse() {}

type StoreGetActivityResponseStoreActivityListedFamily string

const (
	StoreGetActivityResponseStoreActivityListedFamilyListed StoreGetActivityResponseStoreActivityListedFamily = "listed"
)

func (r StoreGetActivityResponseStoreActivityListedFamily) IsKnown() bool {
	switch r {
	case StoreGetActivityResponseStoreActivityListedFamilyListed:
		return true
	}
	return false
}

type StoreGetActivityResponseStoreActivityListedStatus string

const (
	StoreGetActivityResponseStoreActivityListedStatusAvailable StoreGetActivityResponseStoreActivityListedStatus = "available"
	StoreGetActivityResponseStoreActivityListedStatusPending   StoreGetActivityResponseStoreActivityListedStatus = "pending"
	StoreGetActivityResponseStoreActivityListedStatusSold      StoreGetActivityResponseStoreActivityListedStatus = "sold"
	StoreGetActivityResponseStoreActivityListedStatusArchived  StoreGetActivityResponseStoreActivityListedStatus = "archived"
)

func (r StoreGetActivityResponseStoreActivityListedStatus) IsKnown() bool {
	switch r {
	case StoreGetActivityResponseStoreActivityListedStatusAvailable, StoreGetActivityResponseStoreActivityListedStatusPending, StoreGetActivityResponseStoreActivityListedStatusSold, StoreGetActivityResponseStoreActivityListedStatusArchived:
		return true
	}
	return false
}

// Sale activity. Pins the base's `status` to a single member — a provable
// narrowing, so this arm keeps a plain `extends` with no omitted keys.
type StoreGetActivityResponseStoreActivitySold struct {
	ActivityID string                                          `json:"activityId" api:"required"`
	Family     StoreGetActivityResponseStoreActivitySoldFamily `json:"family" api:"required"`
	RecordedAt time.Time                                       `json:"recordedAt" api:"required" format:"date-time"`
	Status     StoreGetActivityResponseStoreActivitySoldStatus `json:"status" api:"required"`
	JSON       storeGetActivityResponseStoreActivitySoldJSON   `json:"-"`
	Pet
}

// storeGetActivityResponseStoreActivitySoldJSON contains the JSON metadata for the
// struct [StoreGetActivityResponseStoreActivitySold]
type storeGetActivityResponseStoreActivitySoldJSON struct {
	ActivityID  apijson.Field
	Family      apijson.Field
	RecordedAt  apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreGetActivityResponseStoreActivitySold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeGetActivityResponseStoreActivitySoldJSON) RawJSON() string {
	return r.raw
}

func (r StoreGetActivityResponseStoreActivitySold) implementsStoreGetActivityResponse() {}

type StoreGetActivityResponseStoreActivitySoldFamily string

const (
	StoreGetActivityResponseStoreActivitySoldFamilySold StoreGetActivityResponseStoreActivitySoldFamily = "sold"
)

func (r StoreGetActivityResponseStoreActivitySoldFamily) IsKnown() bool {
	switch r {
	case StoreGetActivityResponseStoreActivitySoldFamilySold:
		return true
	}
	return false
}

type StoreGetActivityResponseStoreActivitySoldStatus string

const (
	StoreGetActivityResponseStoreActivitySoldStatusSold StoreGetActivityResponseStoreActivitySoldStatus = "sold"
)

func (r StoreGetActivityResponseStoreActivitySoldStatus) IsKnown() bool {
	switch r {
	case StoreGetActivityResponseStoreActivitySoldStatusSold:
		return true
	}
	return false
}

// Archival activity. Narrows the base's `status` by value but also makes it
// nullable, which `Pet` does not accept — a value subset alone is not enough to
// keep a plain `extends` here.
type StoreGetActivityResponseStoreActivityArchived struct {
	ActivityID     string                                              `json:"activityId" api:"required"`
	Family         StoreGetActivityResponseStoreActivityArchivedFamily `json:"family" api:"required"`
	RecordedAt     time.Time                                           `json:"recordedAt" api:"required" format:"date-time"`
	ArchivedReason string                                              `json:"archivedReason"`
	Status         StoreGetActivityResponseStoreActivityArchivedStatus `json:"status" api:"nullable"`
	JSON           storeGetActivityResponseStoreActivityArchivedJSON   `json:"-"`
	Pet
}

// storeGetActivityResponseStoreActivityArchivedJSON contains the JSON metadata for
// the struct [StoreGetActivityResponseStoreActivityArchived]
type storeGetActivityResponseStoreActivityArchivedJSON struct {
	ActivityID     apijson.Field
	Family         apijson.Field
	RecordedAt     apijson.Field
	ArchivedReason apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StoreGetActivityResponseStoreActivityArchived) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeGetActivityResponseStoreActivityArchivedJSON) RawJSON() string {
	return r.raw
}

func (r StoreGetActivityResponseStoreActivityArchived) implementsStoreGetActivityResponse() {}

type StoreGetActivityResponseStoreActivityArchivedFamily string

const (
	StoreGetActivityResponseStoreActivityArchivedFamilyArchived StoreGetActivityResponseStoreActivityArchivedFamily = "archived"
)

func (r StoreGetActivityResponseStoreActivityArchivedFamily) IsKnown() bool {
	switch r {
	case StoreGetActivityResponseStoreActivityArchivedFamilyArchived:
		return true
	}
	return false
}

type StoreGetActivityResponseStoreActivityArchivedStatus string

const (
	StoreGetActivityResponseStoreActivityArchivedStatusSold StoreGetActivityResponseStoreActivityArchivedStatus = "sold"
)

func (r StoreGetActivityResponseStoreActivityArchivedStatus) IsKnown() bool {
	switch r {
	case StoreGetActivityResponseStoreActivityArchivedStatusSold:
		return true
	}
	return false
}

// Transfer activity. Reaches `Pet` only through the pass-through wrapper branch,
// so the base is resolvable only transitively.
type StoreGetActivityResponseStoreActivityTransferred struct {
	Family        StoreGetActivityResponseStoreActivityTransferredFamily        `json:"family" api:"required"`
	TransferredTo StoreGetActivityResponseStoreActivityTransferredTransferredTo `json:"transferredTo"`
	JSON          storeGetActivityResponseStoreActivityTransferredJSON          `json:"-"`
	Pet
}

// storeGetActivityResponseStoreActivityTransferredJSON contains the JSON metadata
// for the struct [StoreGetActivityResponseStoreActivityTransferred]
type storeGetActivityResponseStoreActivityTransferredJSON struct {
	Family        apijson.Field
	TransferredTo apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StoreGetActivityResponseStoreActivityTransferred) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeGetActivityResponseStoreActivityTransferredJSON) RawJSON() string {
	return r.raw
}

func (r StoreGetActivityResponseStoreActivityTransferred) implementsStoreGetActivityResponse() {}

type StoreGetActivityResponseStoreActivityTransferredFamily string

const (
	StoreGetActivityResponseStoreActivityTransferredFamilyTransferred StoreGetActivityResponseStoreActivityTransferredFamily = "transferred"
)

func (r StoreGetActivityResponseStoreActivityTransferredFamily) IsKnown() bool {
	switch r {
	case StoreGetActivityResponseStoreActivityTransferredFamilyTransferred:
		return true
	}
	return false
}

type StoreGetActivityResponseStoreActivityTransferredTransferredTo struct {
	Name     string                                                                `json:"name" api:"required"`
	Address  shared.Address                                                        `json:"address"`
	Contact  StoreGetActivityResponseStoreActivityTransferredTransferredToContact  `json:"contact"`
	Location StoreGetActivityResponseStoreActivityTransferredTransferredToLocation `json:"location"`
	JSON     storeGetActivityResponseStoreActivityTransferredTransferredToJSON     `json:"-"`
}

// storeGetActivityResponseStoreActivityTransferredTransferredToJSON contains the
// JSON metadata for the struct
// [StoreGetActivityResponseStoreActivityTransferredTransferredTo]
type storeGetActivityResponseStoreActivityTransferredTransferredToJSON struct {
	Name        apijson.Field
	Address     apijson.Field
	Contact     apijson.Field
	Location    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreGetActivityResponseStoreActivityTransferredTransferredTo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeGetActivityResponseStoreActivityTransferredTransferredToJSON) RawJSON() string {
	return r.raw
}

type StoreGetActivityResponseStoreActivityTransferredTransferredToContact struct {
	Primary StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimary `json:"primary"`
	JSON    storeGetActivityResponseStoreActivityTransferredTransferredToContactJSON    `json:"-"`
}

// storeGetActivityResponseStoreActivityTransferredTransferredToContactJSON
// contains the JSON metadata for the struct
// [StoreGetActivityResponseStoreActivityTransferredTransferredToContact]
type storeGetActivityResponseStoreActivityTransferredTransferredToContactJSON struct {
	Primary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreGetActivityResponseStoreActivityTransferredTransferredToContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeGetActivityResponseStoreActivityTransferredTransferredToContactJSON) RawJSON() string {
	return r.raw
}

type StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimary struct {
	Availability StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailability `json:"availability"`
	Name         string                                                                                  `json:"name"`
	Phone        string                                                                                  `json:"phone"`
	JSON         storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryJSON         `json:"-"`
}

// storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryJSON
// contains the JSON metadata for the struct
// [StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimary]
type storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryJSON struct {
	Availability apijson.Field
	Name         apijson.Field
	Phone        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryJSON) RawJSON() string {
	return r.raw
}

type StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailability struct {
	Weekday StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekday `json:"weekday"`
	Weekend StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekend `json:"weekend" api:"nullable"`
	JSON    storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityJSON    `json:"-"`
}

// storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityJSON
// contains the JSON metadata for the struct
// [StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailability]
type storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityJSON struct {
	Weekday     apijson.Field
	Weekend     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityJSON) RawJSON() string {
	return r.raw
}

type StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekday struct {
	Close string                                                                                             `json:"close"`
	Open  string                                                                                             `json:"open"`
	JSON  storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekdayJSON `json:"-"`
}

// storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekdayJSON
// contains the JSON metadata for the struct
// [StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekday]
type storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekdayJSON struct {
	Close       apijson.Field
	Open        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekday) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekdayJSON) RawJSON() string {
	return r.raw
}

type StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekend struct {
	Close string                                                                                             `json:"close"`
	Open  string                                                                                             `json:"open"`
	JSON  storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekendJSON `json:"-"`
}

// storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekendJSON
// contains the JSON metadata for the struct
// [StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekend]
type storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekendJSON struct {
	Close       apijson.Field
	Open        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeGetActivityResponseStoreActivityTransferredTransferredToContactPrimaryAvailabilityWeekendJSON) RawJSON() string {
	return r.raw
}

type StoreGetActivityResponseStoreActivityTransferredTransferredToLocation struct {
	Latitude  float64                                                                   `json:"latitude" api:"required"`
	Longitude float64                                                                   `json:"longitude" api:"required"`
	JSON      storeGetActivityResponseStoreActivityTransferredTransferredToLocationJSON `json:"-"`
}

// storeGetActivityResponseStoreActivityTransferredTransferredToLocationJSON
// contains the JSON metadata for the struct
// [StoreGetActivityResponseStoreActivityTransferredTransferredToLocation]
type storeGetActivityResponseStoreActivityTransferredTransferredToLocationJSON struct {
	Latitude    apijson.Field
	Longitude   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreGetActivityResponseStoreActivityTransferredTransferredToLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeGetActivityResponseStoreActivityTransferredTransferredToLocationJSON) RawJSON() string {
	return r.raw
}

type StoreGetActivityResponseFamily string

const (
	StoreGetActivityResponseFamilyListed      StoreGetActivityResponseFamily = "listed"
	StoreGetActivityResponseFamilySold        StoreGetActivityResponseFamily = "sold"
	StoreGetActivityResponseFamilyArchived    StoreGetActivityResponseFamily = "archived"
	StoreGetActivityResponseFamilyTransferred StoreGetActivityResponseFamily = "transferred"
)

func (r StoreGetActivityResponseFamily) IsKnown() bool {
	switch r {
	case StoreGetActivityResponseFamilyListed, StoreGetActivityResponseFamilySold, StoreGetActivityResponseFamilyArchived, StoreGetActivityResponseFamilyTransferred:
		return true
	}
	return false
}
