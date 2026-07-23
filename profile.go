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
	"github.com/swagger-girly/go/internal/param"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
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
	Options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r *ProfileService) {
	r = &ProfileService{}
	r.Options = opts
	return
}

// Creates a pet owner profile with nested preferences, metadata maps, and
// compliance-only write fields.
func (r *ProfileService) New(ctx context.Context, params ProfileNewParams, opts ...option.RequestOption) (res *Profile, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a profile with read-only version fields and nullable contact
// preferences.
func (r *ProfileService) Get(ctx context.Context, profileID string, opts ...option.RequestOption) (res *Profile, err error) {
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("profiles/%s", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates either public profile details or archival state.
func (r *ProfileService) Update(ctx context.Context, profileID string, params ProfileUpdateParams, opts ...option.RequestOption) (res *Profile, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("profiles/%s", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Deprecated profile search endpoint kept for compatibility.
//
// Deprecated: deprecated
func (r *ProfileService) LegacySearch(ctx context.Context, query ProfileLegacySearchParams, opts ...option.RequestOption) (res *[]Profile, err error) {
	opts = slices.Concat(r.Options, opts)
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
	AuditTrail map[string]interface{} `json:"auditTrail"`
	Email      string                 `json:"email" api:"nullable"`
	// Deprecated external profile code.
	//
	// Deprecated: deprecated
	LegacyCode   string             `json:"legacyCode"`
	Metadata     map[string]string  `json:"metadata"`
	PhoneNumbers []string           `json:"phoneNumbers" api:"nullable"`
	Preferences  ProfilePreferences `json:"preferences"`
	JSON         profileJSON        `json:"-"`
}

// profileJSON contains the JSON metadata for the struct [Profile]
type profileJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	DisplayName  apijson.Field
	UpdatedAt    apijson.Field
	Version      apijson.Field
	AuditTrail   apijson.Field
	Email        apijson.Field
	LegacyCode   apijson.Field
	Metadata     apijson.Field
	PhoneNumbers apijson.Field
	Preferences  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Profile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileJSON) RawJSON() string {
	return r.raw
}

type ProfilePreferences struct {
	Alerts        map[string]bool        `json:"alerts"`
	Newsletter    bool                   `json:"newsletter"`
	ScoreByRegion interface{}            `json:"scoreByRegion"`
	JSON          profilePreferencesJSON `json:"-"`
}

// profilePreferencesJSON contains the JSON metadata for the struct
// [ProfilePreferences]
type profilePreferencesJSON struct {
	Alerts        apijson.Field
	Newsletter    apijson.Field
	ScoreByRegion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfilePreferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePreferencesJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateRequestParam struct {
	DisplayName param.Field[string] `json:"displayName" api:"required"`
	Email       param.Field[string] `json:"email"`
	// Compliance identifier used only during profile creation.
	GovernmentID param.Field[string]                               `json:"governmentId"`
	Metadata     param.Field[map[string]string]                    `json:"metadata"`
	Preferences  param.Field[ProfileCreateRequestPreferencesParam] `json:"preferences"`
	// Internal note that must not be emitted in profile responses.
	SecretNote param.Field[string] `json:"secretNote"`
}

func (r ProfileCreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProfileCreateRequestPreferencesParam struct {
	Alerts        param.Field[map[string]bool] `json:"alerts"`
	Newsletter    param.Field[bool]            `json:"newsletter"`
	ScoreByRegion param.Field[interface{}]     `json:"scoreByRegion"`
}

func (r ProfileCreateRequestPreferencesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProfileNewParams struct {
	ProfileCreateRequest ProfileCreateRequestParam `json:"profile_create_request" api:"required"`
	IdempotencyKey       param.Field[string]       `header:"Idempotency-Key"`
}

func (r ProfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ProfileCreateRequest)
}

type ProfileUpdateParams struct {
	Body           ProfileUpdateParamsBodyUnion `json:"body" api:"required"`
	IdempotencyKey param.Field[string]          `header:"Idempotency-Key"`
}

func (r ProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProfileUpdateParamsBody struct {
	ArchivedAt     param.Field[time.Time]                             `json:"archivedAt" format:"date-time"`
	ArchivedReason param.Field[ProfileUpdateParamsBodyArchivedReason] `json:"archivedReason"`
	DisplayName    param.Field[string]                                `json:"displayName"`
	Email          param.Field[string]                                `json:"email"`
	Metadata       param.Field[interface{}]                           `json:"metadata"`
	Preferences    param.Field[interface{}]                           `json:"preferences"`
}

func (r ProfileUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProfileUpdateParamsBody) implementsProfileUpdateParamsBodyUnion() {}

// Satisfied by [ProfileUpdateParamsBodyProfileDetailsUpdate],
// [ProfileUpdateParamsBodyProfileArchiveUpdate], [ProfileUpdateParamsBody].
type ProfileUpdateParamsBodyUnion interface {
	implementsProfileUpdateParamsBodyUnion()
}

type ProfileUpdateParamsBodyProfileDetailsUpdate struct {
	DisplayName param.Field[string]                                                 `json:"displayName"`
	Email       param.Field[string]                                                 `json:"email"`
	Metadata    param.Field[map[string]string]                                      `json:"metadata"`
	Preferences param.Field[ProfileUpdateParamsBodyProfileDetailsUpdatePreferences] `json:"preferences"`
}

func (r ProfileUpdateParamsBodyProfileDetailsUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProfileUpdateParamsBodyProfileDetailsUpdate) implementsProfileUpdateParamsBodyUnion() {}

type ProfileUpdateParamsBodyProfileDetailsUpdatePreferences struct {
	Alerts        param.Field[map[string]bool] `json:"alerts"`
	Newsletter    param.Field[bool]            `json:"newsletter"`
	ScoreByRegion param.Field[interface{}]     `json:"scoreByRegion"`
}

func (r ProfileUpdateParamsBodyProfileDetailsUpdatePreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProfileUpdateParamsBodyProfileArchiveUpdate struct {
	ArchivedAt     param.Field[time.Time]                                                 `json:"archivedAt" api:"required" format:"date-time"`
	ArchivedReason param.Field[ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReason] `json:"archivedReason" api:"required"`
}

func (r ProfileUpdateParamsBodyProfileArchiveUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProfileUpdateParamsBodyProfileArchiveUpdate) implementsProfileUpdateParamsBodyUnion() {}

type ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReason string

const (
	ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReasonDuplicate       ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReason = "duplicate"
	ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReasonRequested       ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReason = "requested"
	ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReasonPolicyViolation ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReason = "policy_violation"
)

func (r ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReason) IsKnown() bool {
	switch r {
	case ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReasonDuplicate, ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReasonRequested, ProfileUpdateParamsBodyProfileArchiveUpdateArchivedReasonPolicyViolation:
		return true
	}
	return false
}

type ProfileUpdateParamsBodyArchivedReason string

const (
	ProfileUpdateParamsBodyArchivedReasonDuplicate       ProfileUpdateParamsBodyArchivedReason = "duplicate"
	ProfileUpdateParamsBodyArchivedReasonRequested       ProfileUpdateParamsBodyArchivedReason = "requested"
	ProfileUpdateParamsBodyArchivedReasonPolicyViolation ProfileUpdateParamsBodyArchivedReason = "policy_violation"
)

func (r ProfileUpdateParamsBodyArchivedReason) IsKnown() bool {
	switch r {
	case ProfileUpdateParamsBodyArchivedReasonDuplicate, ProfileUpdateParamsBodyArchivedReasonRequested, ProfileUpdateParamsBodyArchivedReasonPolicyViolation:
		return true
	}
	return false
}

type ProfileLegacySearchParams struct {
	// Whether archived profiles should be included.
	IncludeArchived param.Field[bool] `query:"include_archived"`
	// Legacy free-text search string.
	Q param.Field[string] `query:"q"`
}

// URLQuery serializes [ProfileLegacySearchParams]'s query parameters as
// `url.Values`.
func (r ProfileLegacySearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
