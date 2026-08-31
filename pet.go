// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/swagger-girly/go/internal/apiform"
	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/apiquery"
	shimjson "github.com/swagger-girly/go/internal/encoding/json"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/pagination"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
	"github.com/swagger-girly/go/packages/ssestream"
	"github.com/swagger-girly/go/shared"
	"github.com/swagger-girly/go/shared/constant"
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
	options []option.RequestOption
}

// NewPetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPetService(opts ...option.RequestOption) (r PetService) {
	r = PetService{}
	r.options = opts
	return
}

// Add a new pet to the store
func (r *PetService) New(ctx context.Context, body PetNewParams, opts ...option.RequestOption) (res *Pet, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single pet
func (r *PetService) Get(ctx context.Context, petID int64, opts ...option.RequestOption) (res *Pet, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing pet by Id
func (r *PetService) Update(ctx context.Context, body PetUpdateParams, opts ...option.RequestOption) (res *Pet, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns a cursor-paginated list of pets.
func (r *PetService) List(ctx context.Context, query PetListParams, opts ...option.RequestOption) (res *pagination.CustomCursorPage[Pet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
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
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Multiple status values can be provided with comma separated strings
func (r *PetService) FindByStatus(ctx context.Context, query PetFindByStatusParams, opts ...option.RequestOption) (res *[]Pet, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet/findByStatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3
// for testing.
func (r *PetService) FindByTags(ctx context.Context, query PetFindByTagsParams, opts ...option.RequestOption) (res *[]Pet, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet/findByTags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a single page-shaped pet response without SDK pagination helpers.
func (r *PetService) ListFakePage(ctx context.Context, opts ...option.RequestOption) (res *PetListFakePageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet/fake-page"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a single page-shaped pet response whose fake pagination behavior is
// inferred from the config scheme.
func (r *PetService) ListFakePageInferred(ctx context.Context, opts ...option.RequestOption) (res *pagination.XFakeSinglePage[Pet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
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

// Returns a bare top-level array of inline objects so generators must mint a
// distinct element type instead of reusing the response alias name.
func (r *PetService) ListLeaderboard(ctx context.Context, opts ...option.RequestOption) (res *[]PetListLeaderboardResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet/leaderboard"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the same cursor-shaped pet list response without enabling SDK pagination
// helpers.
func (r *PetService) ListUnpaginated(ctx context.Context, query PetListUnpaginatedParams, opts ...option.RequestOption) (res *PetListUnpaginatedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet/unpaginated"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the premium profile for a pet, extending the base pet with pedigree and
// insurance details.
func (r *PetService) GetPremium(ctx context.Context, petID int64, opts ...option.RequestOption) (res *PetGetPremiumResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("pet/%v/premium", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Typed query-parameter probe matrix: an object-schema query param mints a typed
// params model, an array-of-object query param mints a singularized element type,
// an empty object (additionalProperties:false) stays a bare object, and a scalar
// stays scalar. Isolates the emitter query-parameter type-resolution branches so
// object/array-of-object/empty-object params are each exercised.
func (r *PetService) Search(ctx context.Context, query PetSearchParams, opts ...option.RequestOption) (res *[]Pet, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates a pet in the store with form data
func (r *PetService) UpdateWithForm(ctx context.Context, petID int64, body PetUpdateWithFormParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// uploads an image
func (r *PetService) UploadImage(ctx context.Context, petID int64, body io.Reader, params PetUploadImageParams, opts ...option.RequestOption) (res *PetUploadImageResponse, err error) {
	opts = slices.Concat(r.options, opts)
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
	if !param.IsOmitted(query.LastEventID) {
		opts = append(opts, option.WithHeader("Last-Event-ID", fmt.Sprintf("%v", query.LastEventID.Value)))
	}
	opts = slices.Concat(r.options, opts)
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
	Category           any                   `json:"category"`
	// Microchip identifier; legacy chips used numeric identifiers.
	MicrochipID     PetMicrochipIDUnion `json:"microchipId"`
	RelatedAddress  *shared.Address     `json:"relatedAddress"`
	RelatedCategory any                 `json:"relatedCategory"`
	RelatedCustomer any                 `json:"relatedCustomer"`
	RelatedMoney    *shared.Money       `json:"relatedMoney"`
	RelatedOrder    any                 `json:"relatedOrder"`
	RelatedPet      *Pet                `json:"relatedPet"`
	RelatedShelter  any                 `json:"relatedShelter"`
	RelatedUser     any                 `json:"relatedUser"`
	// pet status in the store
	//
	// Any of "available", "pending", "sold".
	Status PetStatus `json:"status"`
	Tags   []any     `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pet) RawJSON() string { return r.JSON.raw }
func (r *Pet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Pet to a PetParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PetParam.Overrides()
func (r Pet) ToParam() PetParam {
	return param.Override[PetParam](json.RawMessage(r.RawJSON()))
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

// PetMicrochipIDUnion contains all possible properties and values from [string],
// [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type PetMicrochipIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u PetMicrochipIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PetMicrochipIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PetMicrochipIDUnion) RawJSON() string { return u.JSON.raw }

func (r *PetMicrochipIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, PhotoURLs are required.
type PetParam struct {
	Name      string           `json:"name" api:"required"`
	PhotoURLs []string         `json:"photoUrls,omitzero" api:"required"`
	ID        param.Opt[int64] `json:"id,omitzero"`
	// How the pet entered the store. Open enum: known channels plus forward-compatible
	// free-form strings.
	AcquisitionChannel PetAcquisitionChannel `json:"acquisitionChannel,omitzero"`
	Category           any                   `json:"category,omitzero"`
	// Microchip identifier; legacy chips used numeric identifiers.
	MicrochipID     PetMicrochipIDUnionParam `json:"microchipId,omitzero"`
	RelatedAddress  shared.AddressParam      `json:"relatedAddress,omitzero"`
	RelatedCategory any                      `json:"relatedCategory,omitzero"`
	RelatedCustomer any                      `json:"relatedCustomer,omitzero"`
	RelatedMoney    shared.MoneyParam        `json:"relatedMoney,omitzero"`
	RelatedOrder    any                      `json:"relatedOrder,omitzero"`
	RelatedPet      *PetParam                `json:"relatedPet,omitzero"`
	RelatedShelter  any                      `json:"relatedShelter,omitzero"`
	RelatedUser     any                      `json:"relatedUser,omitzero"`
	// pet status in the store
	//
	// Any of "available", "pending", "sold".
	Status PetStatus `json:"status,omitzero"`
	Tags   []any     `json:"tags,omitzero"`
	paramObj
}

func (r PetParam) MarshalJSON() (data []byte, err error) {
	type shadow PetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PetParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PetMicrochipIDUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u PetMicrochipIDUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *PetMicrochipIDUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// pet status in the store
type PetStatus string

const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)

type PetListFakePageResponse struct {
	Data    []Pet `json:"data" api:"required"`
	HasMore bool  `json:"has_more" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetListFakePageResponse) RawJSON() string { return r.JSON.raw }
func (r *PetListFakePageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetListLeaderboardResponse struct {
	// Ranked pet ID
	PetID int64                           `json:"petId" api:"required"`
	Badge PetListLeaderboardResponseBadge `json:"badge"`
	// Docs anchor for the badge a leaderboard entry was awarded.
	//
	// Any of "https://docs.petstore.example/leaderboard#top-adopter",
	// "https://docs.petstore.example/leaderboard#rising-star".
	DocsRef PetListLeaderboardResponseDocsRef `json:"docsRef"`
	// Leaderboard position
	Rank int64 `json:"rank"`
	// Adoption return on investment
	Roi PetListLeaderboardResponseRoi `json:"roi"`
	// Consecutive days on the leaderboard.
	SevenDayStreak int64 `json:"streak7d"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PetID          respjson.Field
		Badge          respjson.Field
		DocsRef        respjson.Field
		Rank           respjson.Field
		Roi            respjson.Field
		SevenDayStreak respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetListLeaderboardResponse) RawJSON() string { return r.JSON.raw }
func (r *PetListLeaderboardResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetListLeaderboardResponseBadge struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetListLeaderboardResponseBadge) RawJSON() string { return r.JSON.raw }
func (r *PetListLeaderboardResponseBadge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Docs anchor for the badge a leaderboard entry was awarded.
type PetListLeaderboardResponseDocsRef string

const (
	PetListLeaderboardResponseDocsRefTopAdopter PetListLeaderboardResponseDocsRef = "https://docs.petstore.example/leaderboard#top-adopter"
	PetListLeaderboardResponseDocsRefRisingStar PetListLeaderboardResponseDocsRef = "https://docs.petstore.example/leaderboard#rising-star"
)

// Adoption return on investment
type PetListLeaderboardResponseRoi struct {
	Currency string  `json:"currency"`
	Times    float64 `json:"times"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Times       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetListLeaderboardResponseRoi) RawJSON() string { return r.JSON.raw }
func (r *PetListLeaderboardResponseRoi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetListUnpaginatedResponse struct {
	Items      []Pet  `json:"items" api:"required"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetListUnpaginatedResponse) RawJSON() string { return r.JSON.raw }
func (r *PetListUnpaginatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A pet extended with premium-tier pedigree and insurance details.
type PetGetPremiumResponse struct {
	PremiumSince time.Time `json:"premiumSince" api:"required" format:"date-time"`
	// Coverage limit in minor units, or a negotiated custom limit.
	CoverageLimit PetGetPremiumResponseCoverageLimitUnion `json:"coverageLimit"`
	Insurance     PetGetPremiumResponseInsurance          `json:"insurance"`
	Pedigree      PetGetPremiumResponsePedigree           `json:"pedigree"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PremiumSince  respjson.Field
		CoverageLimit respjson.Field
		Insurance     respjson.Field
		Pedigree      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	Pet
}

// Returns the unmodified JSON received from the API
func (r PetGetPremiumResponse) RawJSON() string { return r.JSON.raw }
func (r *PetGetPremiumResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PetGetPremiumResponseCoverageLimitUnion contains all possible properties and
// values from [int64], [PetGetPremiumResponseCoverageLimitCustomLimit].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt]
type PetGetPremiumResponseCoverageLimitUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field is from variant [PetGetPremiumResponseCoverageLimitCustomLimit].
	Amount int64 `json:"amount"`
	// This field is from variant [PetGetPremiumResponseCoverageLimitCustomLimit].
	Currency string `json:"currency"`
	// This field is from variant [PetGetPremiumResponseCoverageLimitCustomLimit].
	Negotiated bool `json:"negotiated"`
	JSON       struct {
		OfInt      respjson.Field
		Amount     respjson.Field
		Currency   respjson.Field
		Negotiated respjson.Field
		raw        string
	} `json:"-"`
}

func (u PetGetPremiumResponseCoverageLimitUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PetGetPremiumResponseCoverageLimitUnion) AsCustomLimit() (v PetGetPremiumResponseCoverageLimitCustomLimit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PetGetPremiumResponseCoverageLimitUnion) RawJSON() string { return u.JSON.raw }

func (r *PetGetPremiumResponseCoverageLimitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetGetPremiumResponseCoverageLimitCustomLimit struct {
	Amount     int64  `json:"amount" api:"required"`
	Currency   string `json:"currency"`
	Negotiated bool   `json:"negotiated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		Negotiated  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetGetPremiumResponseCoverageLimitCustomLimit) RawJSON() string { return r.JSON.raw }
func (r *PetGetPremiumResponseCoverageLimitCustomLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetGetPremiumResponseInsurance struct {
	PlanID     string       `json:"planId" api:"required"`
	Provider   string       `json:"provider" api:"required"`
	Deductible int64        `json:"deductible"`
	Premium    shared.Money `json:"premium"`
	// Any of "basic", "plus", "platinum".
	Tier string `json:"tier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PlanID      respjson.Field
		Provider    respjson.Field
		Deductible  respjson.Field
		Premium     respjson.Field
		Tier        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetGetPremiumResponseInsurance) RawJSON() string { return r.JSON.raw }
func (r *PetGetPremiumResponseInsurance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetGetPremiumResponsePedigree struct {
	Certified bool     `json:"certified"`
	Lineage   []string `json:"lineage"`
	Registry  string   `json:"registry"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Certified   respjson.Field
		Lineage     respjson.Field
		Registry    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetGetPremiumResponsePedigree) RawJSON() string { return r.JSON.raw }
func (r *PetGetPremiumResponsePedigree) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetUploadImageResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetUploadImageResponse) RawJSON() string { return r.JSON.raw }
func (r *PetUploadImageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConnectClientEventUnion contains all possible properties and values from
// [ConnectClientEventPing], [ConnectClientEventSubscribe].
//
// Use the [ConnectClientEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConnectClientEventUnion struct {
	// Any of "ping", "subscribe".
	Type string `json:"type"`
	// This field is from variant [ConnectClientEventSubscribe].
	Status PetStatus `json:"status"`
	JSON   struct {
		Type   respjson.Field
		Status respjson.Field
		raw    string
	} `json:"-"`
}

// anyConnectClientEvent is implemented by each variant of
// [ConnectClientEventUnion] to add type safety for the return type of
// [ConnectClientEventUnion.AsAny]
type anyConnectClientEvent interface {
	implConnectClientEventUnion()
}

func (ConnectClientEventPing) implConnectClientEventUnion()      {}
func (ConnectClientEventSubscribe) implConnectClientEventUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConnectClientEventUnion.AsAny().(type) {
//	case helloworldtestingggg.ConnectClientEventPing:
//	case helloworldtestingggg.ConnectClientEventSubscribe:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConnectClientEventUnion) AsAny() anyConnectClientEvent {
	switch u.Type {
	case "ping":
		return u.AsPing()
	case "subscribe":
		return u.AsSubscribe()
	}
	return nil
}

func (u ConnectClientEventUnion) AsPing() (v ConnectClientEventPing) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConnectClientEventUnion) AsSubscribe() (v ConnectClientEventSubscribe) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConnectClientEventUnion) RawJSON() string { return u.JSON.raw }

func (r *ConnectClientEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConnectClientEventUnion to a ConnectClientEventUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConnectClientEventUnionParam.Overrides()
func (r ConnectClientEventUnion) ToParam() ConnectClientEventUnionParam {
	return param.Override[ConnectClientEventUnionParam](json.RawMessage(r.RawJSON()))
}

type ConnectClientEventPing struct {
	Type constant.Ping `json:"type" default:"ping"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectClientEventPing) RawJSON() string { return r.JSON.raw }
func (r *ConnectClientEventPing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectClientEventSubscribe struct {
	// pet status in the store
	//
	// Any of "available", "pending", "sold".
	Status PetStatus          `json:"status" api:"required"`
	Type   constant.Subscribe `json:"type" default:"subscribe"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectClientEventSubscribe) RawJSON() string { return r.JSON.raw }
func (r *ConnectClientEventSubscribe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ConnectClientEventParamOfSubscribe(status PetStatus) ConnectClientEventUnionParam {
	var subscribe ConnectClientEventSubscribeParam
	subscribe.Status = status
	return ConnectClientEventUnionParam{OfSubscribe: &subscribe}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConnectClientEventUnionParam struct {
	OfPing      *ConnectClientEventPingParam      `json:",omitzero,inline"`
	OfSubscribe *ConnectClientEventSubscribeParam `json:",omitzero,inline"`
	paramUnion
}

func (u ConnectClientEventUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPing, u.OfSubscribe)
}
func (u *ConnectClientEventUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ConnectClientEventUnionParam](
		"type",
		apijson.Discriminator[ConnectClientEventPingParam]("ping"),
		apijson.Discriminator[ConnectClientEventSubscribeParam]("subscribe"),
	)
}

func NewConnectClientEventPingParam() ConnectClientEventPingParam {
	return ConnectClientEventPingParam{
		Type: "ping",
	}
}

// This struct has a constant value, construct it with
// [NewConnectClientEventPingParam].
type ConnectClientEventPingParam struct {
	Type constant.Ping `json:"type" default:"ping"`
	paramObj
}

func (r ConnectClientEventPingParam) MarshalJSON() (data []byte, err error) {
	type shadow ConnectClientEventPingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectClientEventPingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Status, Type are required.
type ConnectClientEventSubscribeParam struct {
	// pet status in the store
	//
	// Any of "available", "pending", "sold".
	Status PetStatus `json:"status,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "subscribe".
	Type constant.Subscribe `json:"type" default:"subscribe"`
	paramObj
}

func (r ConnectClientEventSubscribeParam) MarshalJSON() (data []byte, err error) {
	type shadow ConnectClientEventSubscribeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectClientEventSubscribeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConnectServerEventUnion contains all possible properties and values from
// [ConnectServerEventPong], [ConnectServerEventStatus].
//
// Use the [ConnectServerEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConnectServerEventUnion struct {
	// Any of "pong", "status".
	Type string `json:"type"`
	// This field is from variant [ConnectServerEventStatus].
	Pet  Pet `json:"pet"`
	JSON struct {
		Type respjson.Field
		Pet  respjson.Field
		raw  string
	} `json:"-"`
}

// anyConnectServerEvent is implemented by each variant of
// [ConnectServerEventUnion] to add type safety for the return type of
// [ConnectServerEventUnion.AsAny]
type anyConnectServerEvent interface {
	implConnectServerEventUnion()
}

func (ConnectServerEventPong) implConnectServerEventUnion()   {}
func (ConnectServerEventStatus) implConnectServerEventUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConnectServerEventUnion.AsAny().(type) {
//	case helloworldtestingggg.ConnectServerEventPong:
//	case helloworldtestingggg.ConnectServerEventStatus:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConnectServerEventUnion) AsAny() anyConnectServerEvent {
	switch u.Type {
	case "pong":
		return u.AsPong()
	case "status":
		return u.AsStatus()
	}
	return nil
}

func (u ConnectServerEventUnion) AsPong() (v ConnectServerEventPong) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConnectServerEventUnion) AsStatus() (v ConnectServerEventStatus) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConnectServerEventUnion) RawJSON() string { return u.JSON.raw }

func (r *ConnectServerEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectServerEventPong struct {
	Type constant.Pong `json:"type" default:"pong"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectServerEventPong) RawJSON() string { return r.JSON.raw }
func (r *ConnectServerEventPong) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectServerEventStatus struct {
	Pet  Pet             `json:"pet" api:"required"`
	Type constant.Status `json:"type" default:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pet         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectServerEventStatus) RawJSON() string { return r.JSON.raw }
func (r *ConnectServerEventStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetNewParams struct {
	Pet PetParam
	paramObj
}

func (r PetNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Pet)
}
func (r *PetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetUpdateParams struct {
	Pet PetParam
	paramObj
}

func (r PetUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Pet)
}
func (r *PetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetListParams struct {
	// Cursor from a previous response used to fetch the next page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of pets to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by created_at timestamp range in UTC. Accepts gt/gte/lt/lte.
	CreatedAt PetListParamsCreatedAt `query:"created_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PetListParams]'s query parameters as `url.Values`.
func (r PetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by created_at timestamp range in UTC. Accepts gt/gte/lt/lte.
type PetListParamsCreatedAt struct {
	// Minimum value to filter by (exclusive).
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// Minimum value to filter by (inclusive).
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Maximum value to filter by (exclusive).
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	// Maximum value to filter by (inclusive).
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [PetListParamsCreatedAt]'s query parameters as `url.Values`.
func (r PetListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetFindByStatusParams struct {
	// Status values that need to be considered for filter
	//
	// Any of "available", "pending", "sold".
	Status PetFindByStatusParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PetFindByStatusParams]'s query parameters as `url.Values`.
func (r PetFindByStatusParams) URLQuery() (v url.Values, err error) {
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

type PetFindByTagsParams struct {
	// Tags to filter by
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PetFindByTagsParams]'s query parameters as `url.Values`.
func (r PetFindByTagsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetListUnpaginatedParams struct {
	// Cursor from a previous response used to fetch the next page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of pets to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PetListUnpaginatedParams]'s query parameters as
// `url.Values`.
func (r PetListUnpaginatedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetSearchParams struct {
	// Scalar query parameter: stays a plain scalar (control probe).
	MaxResults param.Opt[int64] `query:"max_results,omitzero" json:"-"`
	// Object-schema query parameter: mints a typed params model instead of collapsing
	// to a bare object.
	Filters PetSearchParamsFilters `query:"filters,omitzero" json:"-"`
	// Empty-object query parameter (additionalProperties:false): stays a bare object,
	// exercising the empty-object branch.
	RawFilter any `query:"raw_filter,omitzero" json:"-"`
	// Array-of-object query parameter: emitters mint a singularized element type for
	// each item.
	TagFilters []PetSearchParamsTagFilter `query:"tag_filters,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PetSearchParams]'s query parameters as `url.Values`.
func (r PetSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Object-schema query parameter: mints a typed params model instead of collapsing
// to a bare object.
type PetSearchParamsFilters struct {
	Color param.Opt[string] `query:"color,omitzero" json:"-"`
	Size  param.Opt[int64]  `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PetSearchParamsFilters]'s query parameters as `url.Values`.
func (r PetSearchParamsFilters) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetSearchParamsTagFilter struct {
	Key param.Opt[string] `query:"key,omitzero" json:"-"`
	// Any of "exact", "prefix".
	Match string `query:"match,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PetSearchParamsTagFilter]'s query parameters as
// `url.Values`.
func (r PetSearchParamsTagFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetUpdateWithFormParams struct {
	// Name of pet that needs to be updated
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Status of pet that needs to be updated
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PetUpdateWithFormParams]'s query parameters as
// `url.Values`.
func (r PetUpdateWithFormParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetUploadImageParams struct {
	// Additional Metadata
	AdditionalMetadata param.Opt[string] `query:"additionalMetadata,omitzero" json:"-"`
	paramObj
}

func (r PetUploadImageParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
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
func (r PetUploadImageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetWatchStatusParams struct {
	LastEventID param.Opt[string] `header:"Last-Event-ID,omitzero" json:"-"`
	paramObj
}
