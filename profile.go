// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/apiquery"
	shimjson "github.com/swagger-girly/go/internal/encoding/json"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
)

// Pet owner profile and compliance operations
//
// ProfileService contains methods and other services that help with interacting
// with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.options = opts
	return
}

// Creates a pet owner profile with nested preferences, metadata maps, and
// compliance-only write fields.
func (r *ProfileService) New(ctx context.Context, params ProfileNewParams, opts ...option.RequestOption) (res *Profile, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a profile with read-only version fields and nullable contact
// preferences.
func (r *ProfileService) Get(ctx context.Context, profileID string, opts ...option.RequestOption) (res *Profile, err error) {
	opts = slices.Concat(r.options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("profiles/%s", url.PathEscape(profileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates either public profile details or archival state.
func (r *ProfileService) Update(ctx context.Context, profileID string, params ProfileUpdateParams, opts ...option.RequestOption) (res *Profile, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("profiles/%s", url.PathEscape(profileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Deprecated profile search endpoint kept for compatibility.
//
// Deprecated: deprecated
func (r *ProfileService) LegacySearch(ctx context.Context, query ProfileLegacySearchParams, opts ...option.RequestOption) (res *[]Profile, err error) {
	opts = slices.Concat(r.options, opts)
	path := "profiles/legacy-search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Profile struct {
	ID          string    `json:"id" api:"required"`
	CreatedAt   time.Time `json:"createdAt" api:"required" format:"date-time"`
	DisplayName string    `json:"displayName" api:"required"`
	UpdatedAt   time.Time `json:"updatedAt" api:"required" format:"date-time"`
	Version     int64     `json:"version" api:"required"`
	// Free-form audit metadata retained from integrations.
	AuditTrail map[string]any `json:"auditTrail"`
	Email      string         `json:"email" api:"nullable"`
	// Deprecated external profile code.
	//
	// Deprecated: deprecated
	LegacyCode   string             `json:"legacyCode"`
	Metadata     map[string]string  `json:"metadata"`
	PhoneNumbers []string           `json:"phoneNumbers" api:"nullable"`
	Preferences  ProfilePreferences `json:"preferences"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		DisplayName  respjson.Field
		UpdatedAt    respjson.Field
		Version      respjson.Field
		AuditTrail   respjson.Field
		Email        respjson.Field
		LegacyCode   respjson.Field
		Metadata     respjson.Field
		PhoneNumbers respjson.Field
		Preferences  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Profile) RawJSON() string { return r.JSON.raw }
func (r *Profile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfilePreferences struct {
	Alerts        map[string]bool `json:"alerts"`
	Newsletter    bool            `json:"newsletter"`
	ScoreByRegion any             `json:"scoreByRegion"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alerts        respjson.Field
		Newsletter    respjson.Field
		ScoreByRegion respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfilePreferences) RawJSON() string { return r.JSON.raw }
func (r *ProfilePreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property DisplayName is required.
type ProfileCreateRequestParam struct {
	DisplayName string            `json:"displayName" api:"required"`
	Email       param.Opt[string] `json:"email,omitzero"`
	// Compliance identifier used only during profile creation.
	GovernmentID param.Opt[string] `json:"governmentId,omitzero"`
	// Internal note that must not be emitted in profile responses.
	SecretNote  param.Opt[string]                    `json:"secretNote,omitzero"`
	Metadata    map[string]string                    `json:"metadata,omitzero"`
	Preferences ProfileCreateRequestPreferencesParam `json:"preferences,omitzero"`
	paramObj
}

func (r ProfileCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ProfileCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileCreateRequestPreferencesParam struct {
	Newsletter    param.Opt[bool] `json:"newsletter,omitzero"`
	Alerts        map[string]bool `json:"alerts,omitzero"`
	ScoreByRegion any             `json:"scoreByRegion,omitzero"`
	paramObj
}

func (r ProfileCreateRequestPreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow ProfileCreateRequestPreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileCreateRequestPreferencesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileNewParams struct {
	ProfileCreateRequest ProfileCreateRequestParam
	IdempotencyKey       param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r ProfileNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ProfileCreateRequest)
}
func (r *ProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfProfileDetailsUpdate *ProfileUpdateParamsBodyProfileDetailsUpdate `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfProfileArchiveUpdate *ProfileUpdateParamsBodyProfileArchiveUpdate `json:",inline"`

	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (u ProfileUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProfileDetailsUpdate, u.OfProfileArchiveUpdate)
}
func (r *ProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateParamsBodyProfileDetailsUpdate struct {
	Email       param.Opt[string]                                      `json:"email,omitzero"`
	DisplayName param.Opt[string]                                      `json:"displayName,omitzero"`
	Metadata    map[string]string                                      `json:"metadata,omitzero"`
	Preferences ProfileUpdateParamsBodyProfileDetailsUpdatePreferences `json:"preferences,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsBodyProfileDetailsUpdate) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsBodyProfileDetailsUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsBodyProfileDetailsUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateParamsBodyProfileDetailsUpdatePreferences struct {
	Newsletter    param.Opt[bool] `json:"newsletter,omitzero"`
	Alerts        map[string]bool `json:"alerts,omitzero"`
	ScoreByRegion any             `json:"scoreByRegion,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsBodyProfileDetailsUpdatePreferences) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsBodyProfileDetailsUpdatePreferences
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsBodyProfileDetailsUpdatePreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ArchivedAt, ArchivedReason are required.
type ProfileUpdateParamsBodyProfileArchiveUpdate struct {
	ArchivedAt time.Time `json:"archivedAt" api:"required" format:"date-time"`
	// Any of "duplicate", "requested", "policy_violation".
	ArchivedReason string `json:"archivedReason,omitzero" api:"required"`
	paramObj
}

func (r ProfileUpdateParamsBodyProfileArchiveUpdate) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsBodyProfileArchiveUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsBodyProfileArchiveUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfileUpdateParamsBodyProfileArchiveUpdate](
		"archivedReason", "duplicate", "requested", "policy_violation",
	)
}

type ProfileLegacySearchParams struct {
	// Whether archived profiles should be included.
	IncludeArchived param.Opt[bool] `query:"include_archived,omitzero" json:"-"`
	// Legacy free-text search string.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProfileLegacySearchParams]'s query parameters as
// `url.Values`.
func (r ProfileLegacySearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
