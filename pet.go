// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/swagger-girly/go/internal/apiform"
	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/apiquery"
	"github.com/swagger-girly/go/internal/param"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/pagination"
	"github.com/swagger-girly/go/packages/ssestream"
	"github.com/swagger-girly/go/shared"
	"github.com/tidwall/gjson"
)

// Everything about your Pets
//
// PetService contains methods and other services that help with interacting with
// the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPetService] method instead.
type PetService struct {
	Options []option.RequestOption
}

// NewPetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPetService(opts ...option.RequestOption) (r *PetService) {
	r = &PetService{}
	r.Options = opts
	return
}

// Add a new pet to the store
func (r *PetService) New(ctx context.Context, body PetNewParams, opts ...option.RequestOption) (res *Pet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single pet
func (r *PetService) Get(ctx context.Context, petID int64, opts ...option.RequestOption) (res *Pet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing pet by Id
func (r *PetService) Update(ctx context.Context, body PetUpdateParams, opts ...option.RequestOption) (res *Pet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns a cursor-paginated list of pets.
func (r *PetService) List(ctx context.Context, query PetListParams, opts ...option.RequestOption) (res *pagination.CustomCursorPage[Pet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "pet"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a cursor-paginated list of pets.
func (r *PetService) ListAutoPaging(ctx context.Context, query PetListParams, opts ...option.RequestOption) *pagination.CustomCursorPageAutoPager[Pet] {
	return pagination.NewCustomCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a pet
func (r *PetService) Delete(ctx context.Context, petID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Multiple status values can be provided with comma separated strings
func (r *PetService) FindByStatus(ctx context.Context, query PetFindByStatusParams, opts ...option.RequestOption) (res *[]Pet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pet/findByStatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3
// for testing.
func (r *PetService) FindByTags(ctx context.Context, query PetFindByTagsParams, opts ...option.RequestOption) (res *[]Pet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pet/findByTags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a single page-shaped pet response without SDK pagination helpers.
func (r *PetService) ListFakePage(ctx context.Context, opts ...option.RequestOption) (res *PetListFakePageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pet/fake-page"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a single page-shaped pet response whose fake pagination behavior is
// inferred from the config scheme.
func (r *PetService) ListFakePageInferred(ctx context.Context, opts ...option.RequestOption) (res *pagination.XFakeSinglePage[Pet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "pet/fake-page-inferred"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a single page-shaped pet response whose fake pagination behavior is
// inferred from the config scheme.
func (r *PetService) ListFakePageInferredAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.XFakeSinglePageAutoPager[Pet] {
	return pagination.NewXFakeSinglePageAutoPager(r.ListFakePageInferred(ctx, opts...))
}

// Returns the same cursor-shaped pet list response without enabling SDK pagination
// helpers.
func (r *PetService) ListUnpaginated(ctx context.Context, query PetListUnpaginatedParams, opts ...option.RequestOption) (res *PetListUnpaginatedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pet/unpaginated"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the premium profile for a pet, extending the base pet with pedigree and
// insurance details.
func (r *PetService) GetPremium(ctx context.Context, petID int64, opts ...option.RequestOption) (res *PetGetPremiumResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("pet/%v/premium", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a pet in the store with form data
func (r *PetService) UpdateWithForm(ctx context.Context, petID int64, body PetUpdateWithFormParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// uploads an image
func (r *PetService) UploadImage(ctx context.Context, petID int64, body io.Reader, params PetUploadImageParams, opts ...option.RequestOption) (res *PetUploadImageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", body)}, opts...)
	path := fmt.Sprintf("pet/%v/uploadImage", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Streams pet status updates over Server-Sent Events. Each `status` event contains
// a full `Pet` payload.
func (r *PetService) WatchStatusStreaming(ctx context.Context, petID int64, query PetWatchStatusParams, opts ...option.RequestOption) (stream *ssestream.Stream[Pet]) {
	var (
		raw *http.Response
		err error
	)
	if query.LastEventID.Present {
		opts = append(opts, option.WithHeader("Last-Event-ID", fmt.Sprintf("%v", query.LastEventID)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := fmt.Sprintf("pet/%v/status/stream", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[Pet](ssestream.NewDecoder(raw), err)
}

type Pet struct {
	Name      string   `json:"name" api:"required"`
	PhotoURLs []string `json:"photoUrls" api:"required"`
	ID        int64    `json:"id"`
	// How the pet entered the store. Open enum: known channels plus forward-compatible
	// free-form strings.
	AcquisitionChannel PetAcquisitionChannel `json:"acquisitionChannel"`
	Category           PetCategory           `json:"category"`
	// Microchip identifier; legacy chips used numeric identifiers.
	MicrochipID PetMicrochipIDUnion `json:"microchipId"`
	// pet status in the store
	Status PetStatus `json:"status"`
	Tags   []PetTag  `json:"tags"`
	JSON   petJSON   `json:"-"`
}

// petJSON contains the JSON metadata for the struct [Pet]
type petJSON struct {
	Name               apijson.Field
	PhotoURLs          apijson.Field
	ID                 apijson.Field
	AcquisitionChannel apijson.Field
	Category           apijson.Field
	MicrochipID        apijson.Field
	Status             apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Pet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petJSON) RawJSON() string {
	return r.raw
}

// How the pet entered the store. Open enum: known channels plus forward-compatible
// free-form strings.
type PetAcquisitionChannel string

const (
	PetAcquisitionChannelBreeder   PetAcquisitionChannel = "breeder"
	PetAcquisitionChannelShelter   PetAcquisitionChannel = "shelter"
	PetAcquisitionChannelSurrender PetAcquisitionChannel = "surrender"
	PetAcquisitionChannelTransfer  PetAcquisitionChannel = "transfer"
)

func (r PetAcquisitionChannel) IsKnown() bool {
	switch r {
	case PetAcquisitionChannelBreeder, PetAcquisitionChannelShelter, PetAcquisitionChannelSurrender, PetAcquisitionChannelTransfer:
		return true
	}
	return false
}

type PetCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// Nested subcategories; the tree can recurse arbitrarily deep.
	Subcategories []interface{}   `json:"subcategories"`
	JSON          petCategoryJSON `json:"-"`
}

// petCategoryJSON contains the JSON metadata for the struct [PetCategory]
type petCategoryJSON struct {
	ID            apijson.Field
	Name          apijson.Field
	Subcategories apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PetCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petCategoryJSON) RawJSON() string {
	return r.raw
}

// Microchip identifier; legacy chips used numeric identifiers.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type PetMicrochipIDUnion interface {
	ImplementsPetMicrochipIDUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PetMicrochipIDUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

type PetTag struct {
	ID   int64      `json:"id"`
	Name string     `json:"name"`
	JSON petTagJSON `json:"-"`
}

// petTagJSON contains the JSON metadata for the struct [PetTag]
type petTagJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petTagJSON) RawJSON() string {
	return r.raw
}

type PetParam struct {
	Name      param.Field[string]   `json:"name" api:"required"`
	PhotoURLs param.Field[[]string] `json:"photoUrls" api:"required"`
	ID        param.Field[int64]    `json:"id"`
	// How the pet entered the store. Open enum: known channels plus forward-compatible
	// free-form strings.
	AcquisitionChannel param.Field[PetAcquisitionChannel] `json:"acquisitionChannel"`
	Category           param.Field[PetCategoryParam]      `json:"category"`
	// Microchip identifier; legacy chips used numeric identifiers.
	MicrochipID param.Field[PetMicrochipIDUnionParam] `json:"microchipId"`
	// pet status in the store
	Status param.Field[PetStatus]     `json:"status"`
	Tags   param.Field[[]PetTagParam] `json:"tags"`
}

func (r PetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PetCategoryParam struct {
	ID   param.Field[int64]  `json:"id"`
	Name param.Field[string] `json:"name"`
	// Nested subcategories; the tree can recurse arbitrarily deep.
	Subcategories param.Field[[]interface{}] `json:"subcategories"`
}

func (r PetCategoryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Microchip identifier; legacy chips used numeric identifiers.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type PetMicrochipIDUnionParam interface {
	ImplementsPetMicrochipIDUnionParam()
}

type PetTagParam struct {
	ID   param.Field[int64]  `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r PetTagParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pet status in the store
type PetStatus string

const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)

func (r PetStatus) IsKnown() bool {
	switch r {
	case PetStatusAvailable, PetStatusPending, PetStatusSold:
		return true
	}
	return false
}

type PetListFakePageResponse struct {
	Data    []Pet                       `json:"data" api:"required"`
	HasMore bool                        `json:"has_more" api:"required"`
	JSON    petListFakePageResponseJSON `json:"-"`
}

// petListFakePageResponseJSON contains the JSON metadata for the struct
// [PetListFakePageResponse]
type petListFakePageResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetListFakePageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petListFakePageResponseJSON) RawJSON() string {
	return r.raw
}

type PetListUnpaginatedResponse struct {
	Items      []Pet                          `json:"items" api:"required"`
	NextCursor string                         `json:"next_cursor" api:"nullable"`
	JSON       petListUnpaginatedResponseJSON `json:"-"`
}

// petListUnpaginatedResponseJSON contains the JSON metadata for the struct
// [PetListUnpaginatedResponse]
type petListUnpaginatedResponseJSON struct {
	Items       apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetListUnpaginatedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petListUnpaginatedResponseJSON) RawJSON() string {
	return r.raw
}

// A pet extended with premium-tier pedigree and insurance details.
type PetGetPremiumResponse struct {
	PremiumSince time.Time `json:"premiumSince" api:"required" format:"date-time"`
	// Coverage limit in minor units, or a negotiated custom limit.
	CoverageLimit PetGetPremiumResponseCoverageLimitUnion `json:"coverageLimit"`
	Insurance     PetGetPremiumResponseInsurance          `json:"insurance"`
	Pedigree      PetGetPremiumResponsePedigree           `json:"pedigree"`
	JSON          petGetPremiumResponseJSON               `json:"-"`
	Pet
}

// petGetPremiumResponseJSON contains the JSON metadata for the struct
// [PetGetPremiumResponse]
type petGetPremiumResponseJSON struct {
	PremiumSince  apijson.Field
	CoverageLimit apijson.Field
	Insurance     apijson.Field
	Pedigree      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PetGetPremiumResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petGetPremiumResponseJSON) RawJSON() string {
	return r.raw
}

// Coverage limit in minor units, or a negotiated custom limit.
//
// Union satisfied by [shared.UnionInt] or
// [PetGetPremiumResponseCoverageLimitCustomLimit].
type PetGetPremiumResponseCoverageLimitUnion interface {
	ImplementsPetGetPremiumResponseCoverageLimitUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PetGetPremiumResponseCoverageLimitUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PetGetPremiumResponseCoverageLimitCustomLimit{}),
		},
	)
}

type PetGetPremiumResponseCoverageLimitCustomLimit struct {
	Amount     int64                                             `json:"amount" api:"required"`
	Currency   string                                            `json:"currency"`
	Negotiated bool                                              `json:"negotiated"`
	JSON       petGetPremiumResponseCoverageLimitCustomLimitJSON `json:"-"`
}

// petGetPremiumResponseCoverageLimitCustomLimitJSON contains the JSON metadata for
// the struct [PetGetPremiumResponseCoverageLimitCustomLimit]
type petGetPremiumResponseCoverageLimitCustomLimitJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Negotiated  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetGetPremiumResponseCoverageLimitCustomLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petGetPremiumResponseCoverageLimitCustomLimitJSON) RawJSON() string {
	return r.raw
}

func (r PetGetPremiumResponseCoverageLimitCustomLimit) ImplementsPetGetPremiumResponseCoverageLimitUnion() {
}

type PetGetPremiumResponseInsurance struct {
	PlanID     string                             `json:"planId" api:"required"`
	Provider   string                             `json:"provider" api:"required"`
	Deductible int64                              `json:"deductible"`
	Premium    shared.Money                       `json:"premium"`
	Tier       PetGetPremiumResponseInsuranceTier `json:"tier"`
	JSON       petGetPremiumResponseInsuranceJSON `json:"-"`
}

// petGetPremiumResponseInsuranceJSON contains the JSON metadata for the struct
// [PetGetPremiumResponseInsurance]
type petGetPremiumResponseInsuranceJSON struct {
	PlanID      apijson.Field
	Provider    apijson.Field
	Deductible  apijson.Field
	Premium     apijson.Field
	Tier        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetGetPremiumResponseInsurance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petGetPremiumResponseInsuranceJSON) RawJSON() string {
	return r.raw
}

type PetGetPremiumResponseInsuranceTier string

const (
	PetGetPremiumResponseInsuranceTierBasic    PetGetPremiumResponseInsuranceTier = "basic"
	PetGetPremiumResponseInsuranceTierPlus     PetGetPremiumResponseInsuranceTier = "plus"
	PetGetPremiumResponseInsuranceTierPlatinum PetGetPremiumResponseInsuranceTier = "platinum"
)

func (r PetGetPremiumResponseInsuranceTier) IsKnown() bool {
	switch r {
	case PetGetPremiumResponseInsuranceTierBasic, PetGetPremiumResponseInsuranceTierPlus, PetGetPremiumResponseInsuranceTierPlatinum:
		return true
	}
	return false
}

type PetGetPremiumResponsePedigree struct {
	Certified bool                              `json:"certified"`
	Lineage   []string                          `json:"lineage"`
	Registry  string                            `json:"registry"`
	JSON      petGetPremiumResponsePedigreeJSON `json:"-"`
}

// petGetPremiumResponsePedigreeJSON contains the JSON metadata for the struct
// [PetGetPremiumResponsePedigree]
type petGetPremiumResponsePedigreeJSON struct {
	Certified   apijson.Field
	Lineage     apijson.Field
	Registry    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetGetPremiumResponsePedigree) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petGetPremiumResponsePedigreeJSON) RawJSON() string {
	return r.raw
}

type PetUploadImageResponse struct {
	Code    int64                      `json:"code"`
	Message string                     `json:"message"`
	Type    string                     `json:"type"`
	JSON    petUploadImageResponseJSON `json:"-"`
}

// petUploadImageResponseJSON contains the JSON metadata for the struct
// [PetUploadImageResponse]
type petUploadImageResponseJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetUploadImageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petUploadImageResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectClientEvent struct {
	Type ConnectClientEventType `json:"type" api:"required"`
	// pet status in the store
	Status PetStatus              `json:"status"`
	JSON   connectClientEventJSON `json:"-"`
	union  ConnectClientEventUnion
}

// connectClientEventJSON contains the JSON metadata for the struct
// [ConnectClientEvent]
type connectClientEventJSON struct {
	Type        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r connectClientEventJSON) RawJSON() string {
	return r.raw
}

func (r *ConnectClientEvent) UnmarshalJSON(data []byte) (err error) {
	*r = ConnectClientEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConnectClientEventUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ConnectClientEventPetClientPingEvent],
// [ConnectClientEventPetClientSubscribeEvent].
func (r ConnectClientEvent) AsUnion() ConnectClientEventUnion {
	return r.union
}

// Union satisfied by [ConnectClientEventPetClientPingEvent] or
// [ConnectClientEventPetClientSubscribeEvent].
type ConnectClientEventUnion interface {
	implementsConnectClientEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectClientEventUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectClientEventPetClientPingEvent{}),
			DiscriminatorValue: "ping",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectClientEventPetClientSubscribeEvent{}),
			DiscriminatorValue: "subscribe",
		},
	)
}

type ConnectClientEventPetClientPingEvent struct {
	Type ConnectClientEventPetClientPingEventType `json:"type" api:"required"`
	JSON connectClientEventPetClientPingEventJSON `json:"-"`
}

// connectClientEventPetClientPingEventJSON contains the JSON metadata for the
// struct [ConnectClientEventPetClientPingEvent]
type connectClientEventPetClientPingEventJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectClientEventPetClientPingEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectClientEventPetClientPingEventJSON) RawJSON() string {
	return r.raw
}

func (r ConnectClientEventPetClientPingEvent) implementsConnectClientEvent() {}

type ConnectClientEventPetClientPingEventType string

const (
	ConnectClientEventPetClientPingEventTypePing ConnectClientEventPetClientPingEventType = "ping"
)

func (r ConnectClientEventPetClientPingEventType) IsKnown() bool {
	switch r {
	case ConnectClientEventPetClientPingEventTypePing:
		return true
	}
	return false
}

type ConnectClientEventPetClientSubscribeEvent struct {
	// pet status in the store
	Status PetStatus                                     `json:"status" api:"required"`
	Type   ConnectClientEventPetClientSubscribeEventType `json:"type" api:"required"`
	JSON   connectClientEventPetClientSubscribeEventJSON `json:"-"`
}

// connectClientEventPetClientSubscribeEventJSON contains the JSON metadata for the
// struct [ConnectClientEventPetClientSubscribeEvent]
type connectClientEventPetClientSubscribeEventJSON struct {
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectClientEventPetClientSubscribeEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectClientEventPetClientSubscribeEventJSON) RawJSON() string {
	return r.raw
}

func (r ConnectClientEventPetClientSubscribeEvent) implementsConnectClientEvent() {}

type ConnectClientEventPetClientSubscribeEventType string

const (
	ConnectClientEventPetClientSubscribeEventTypeSubscribe ConnectClientEventPetClientSubscribeEventType = "subscribe"
)

func (r ConnectClientEventPetClientSubscribeEventType) IsKnown() bool {
	switch r {
	case ConnectClientEventPetClientSubscribeEventTypeSubscribe:
		return true
	}
	return false
}

type ConnectClientEventType string

const (
	ConnectClientEventTypePing      ConnectClientEventType = "ping"
	ConnectClientEventTypeSubscribe ConnectClientEventType = "subscribe"
)

func (r ConnectClientEventType) IsKnown() bool {
	switch r {
	case ConnectClientEventTypePing, ConnectClientEventTypeSubscribe:
		return true
	}
	return false
}

type ConnectClientEventParam struct {
	Type param.Field[ConnectClientEventType] `json:"type" api:"required"`
	// pet status in the store
	Status param.Field[PetStatus] `json:"status"`
}

func (r ConnectClientEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectClientEventParam) implementsConnectClientEventUnionParam() {}

// Satisfied by [ConnectClientEventPetClientPingEventParam],
// [ConnectClientEventPetClientSubscribeEventParam], [ConnectClientEventParam].
type ConnectClientEventUnionParam interface {
	implementsConnectClientEventUnionParam()
}

type ConnectClientEventPetClientPingEventParam struct {
	Type param.Field[ConnectClientEventPetClientPingEventType] `json:"type" api:"required"`
}

func (r ConnectClientEventPetClientPingEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectClientEventPetClientPingEventParam) implementsConnectClientEventUnionParam() {}

type ConnectClientEventPetClientSubscribeEventParam struct {
	// pet status in the store
	Status param.Field[PetStatus]                                     `json:"status" api:"required"`
	Type   param.Field[ConnectClientEventPetClientSubscribeEventType] `json:"type" api:"required"`
}

func (r ConnectClientEventPetClientSubscribeEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectClientEventPetClientSubscribeEventParam) implementsConnectClientEventUnionParam() {}

type ConnectServerEvent struct {
	Type  ConnectServerEventType `json:"type" api:"required"`
	Pet   Pet                    `json:"pet"`
	JSON  connectServerEventJSON `json:"-"`
	union ConnectServerEventUnion
}

// connectServerEventJSON contains the JSON metadata for the struct
// [ConnectServerEvent]
type connectServerEventJSON struct {
	Type        apijson.Field
	Pet         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r connectServerEventJSON) RawJSON() string {
	return r.raw
}

func (r *ConnectServerEvent) UnmarshalJSON(data []byte) (err error) {
	*r = ConnectServerEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConnectServerEventUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ConnectServerEventPetServerPongEvent],
// [ConnectServerEventPetServerStatusEvent].
func (r ConnectServerEvent) AsUnion() ConnectServerEventUnion {
	return r.union
}

// Union satisfied by [ConnectServerEventPetServerPongEvent] or
// [ConnectServerEventPetServerStatusEvent].
type ConnectServerEventUnion interface {
	implementsConnectServerEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectServerEventUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectServerEventPetServerPongEvent{}),
			DiscriminatorValue: "pong",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectServerEventPetServerStatusEvent{}),
			DiscriminatorValue: "status",
		},
	)
}

type ConnectServerEventPetServerPongEvent struct {
	Type ConnectServerEventPetServerPongEventType `json:"type" api:"required"`
	JSON connectServerEventPetServerPongEventJSON `json:"-"`
}

// connectServerEventPetServerPongEventJSON contains the JSON metadata for the
// struct [ConnectServerEventPetServerPongEvent]
type connectServerEventPetServerPongEventJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectServerEventPetServerPongEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectServerEventPetServerPongEventJSON) RawJSON() string {
	return r.raw
}

func (r ConnectServerEventPetServerPongEvent) implementsConnectServerEvent() {}

type ConnectServerEventPetServerPongEventType string

const (
	ConnectServerEventPetServerPongEventTypePong ConnectServerEventPetServerPongEventType = "pong"
)

func (r ConnectServerEventPetServerPongEventType) IsKnown() bool {
	switch r {
	case ConnectServerEventPetServerPongEventTypePong:
		return true
	}
	return false
}

type ConnectServerEventPetServerStatusEvent struct {
	Pet  Pet                                        `json:"pet" api:"required"`
	Type ConnectServerEventPetServerStatusEventType `json:"type" api:"required"`
	JSON connectServerEventPetServerStatusEventJSON `json:"-"`
}

// connectServerEventPetServerStatusEventJSON contains the JSON metadata for the
// struct [ConnectServerEventPetServerStatusEvent]
type connectServerEventPetServerStatusEventJSON struct {
	Pet         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectServerEventPetServerStatusEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectServerEventPetServerStatusEventJSON) RawJSON() string {
	return r.raw
}

func (r ConnectServerEventPetServerStatusEvent) implementsConnectServerEvent() {}

type ConnectServerEventPetServerStatusEventType string

const (
	ConnectServerEventPetServerStatusEventTypeStatus ConnectServerEventPetServerStatusEventType = "status"
)

func (r ConnectServerEventPetServerStatusEventType) IsKnown() bool {
	switch r {
	case ConnectServerEventPetServerStatusEventTypeStatus:
		return true
	}
	return false
}

type ConnectServerEventType string

const (
	ConnectServerEventTypePong   ConnectServerEventType = "pong"
	ConnectServerEventTypeStatus ConnectServerEventType = "status"
)

func (r ConnectServerEventType) IsKnown() bool {
	switch r {
	case ConnectServerEventTypePong, ConnectServerEventTypeStatus:
		return true
	}
	return false
}

type PetNewParams struct {
	Pet PetParam `json:"pet" api:"required"`
}

func (r PetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Pet)
}

type PetUpdateParams struct {
	Pet PetParam `json:"pet" api:"required"`
}

func (r PetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Pet)
}

type PetListParams struct {
	// Cursor from a previous response used to fetch the next page.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of pets to return.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [PetListParams]'s query parameters as `url.Values`.
func (r PetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetFindByStatusParams struct {
	// Status values that need to be considered for filter
	Status param.Field[PetFindByStatusParamsStatus] `query:"status"`
}

// URLQuery serializes [PetFindByStatusParams]'s query parameters as `url.Values`.
func (r PetFindByStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Status values that need to be considered for filter
type PetFindByStatusParamsStatus string

const (
	PetFindByStatusParamsStatusAvailable PetFindByStatusParamsStatus = "available"
	PetFindByStatusParamsStatusPending   PetFindByStatusParamsStatus = "pending"
	PetFindByStatusParamsStatusSold      PetFindByStatusParamsStatus = "sold"
)

func (r PetFindByStatusParamsStatus) IsKnown() bool {
	switch r {
	case PetFindByStatusParamsStatusAvailable, PetFindByStatusParamsStatusPending, PetFindByStatusParamsStatusSold:
		return true
	}
	return false
}

type PetFindByTagsParams struct {
	// Tags to filter by
	Tags param.Field[[]string] `query:"tags"`
}

// URLQuery serializes [PetFindByTagsParams]'s query parameters as `url.Values`.
func (r PetFindByTagsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetListUnpaginatedParams struct {
	// Cursor from a previous response used to fetch the next page.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of pets to return.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [PetListUnpaginatedParams]'s query parameters as
// `url.Values`.
func (r PetListUnpaginatedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetUpdateWithFormParams struct {
	// Name of pet that needs to be updated
	Name param.Field[string] `query:"name"`
	// Status of pet that needs to be updated
	Status param.Field[string] `query:"status"`
}

// URLQuery serializes [PetUpdateWithFormParams]'s query parameters as
// `url.Values`.
func (r PetUpdateWithFormParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetUploadImageParams struct {
	// Additional Metadata
	AdditionalMetadata param.Field[string] `query:"additionalMetadata"`
}

func (r PetUploadImageParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [PetUploadImageParams]'s query parameters as `url.Values`.
func (r PetUploadImageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetWatchStatusParams struct {
	LastEventID param.Field[string] `header:"Last-Event-ID"`
}
