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

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/apiquery"
	shimjson "github.com/swagger-girly/go/internal/encoding/json"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
)

// Operations about user
//
// UserService contains methods and other services that help with interacting with
// the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.options = opts
	return
}

// This can only be done by the logged in user.
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *UserNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get user by user name
func (r *UserService) Get(ctx context.Context, username string, opts ...option.RequestOption) (res *UserGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if username == "" {
		err = errors.New("missing required username parameter")
		return nil, err
	}
	path := fmt.Sprintf("user/%s", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This can only be done by the logged in user.
func (r *UserService) Update(ctx context.Context, username string, body UserUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return err
	}
	path := fmt.Sprintf("user/%s", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// This can only be done by the logged in user.
func (r *UserService) Delete(ctx context.Context, username string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return err
	}
	path := fmt.Sprintf("user/%s", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Creates list of users with given input array
func (r *UserService) NewWithList(ctx context.Context, body UserNewWithListParams, opts ...option.RequestOption) (res *UserNewWithListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "user/createWithList"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Logs user into the system
func (r *UserService) Login(ctx context.Context, query UserLoginParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "user/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Logs out current logged in user session
func (r *UserService) Logout(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "user/logout"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Casing probe replicating lithic: inline oneOf response whose variant is a
// kebab-named component (kyb-kyc-verification) so adjacent default initialisms
// glue into KYBKYC in the operation-scoped variant name
func (r *UserService) VerifyIdentity(ctx context.Context, username string, opts ...option.RequestOption) (res *UserVerifyIdentityResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if username == "" {
		err = errors.New("missing required username parameter")
		return nil, err
	}
	path := fmt.Sprintf("user/%s/verifyIdentity", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type UserNewResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
	// User Status
	UserStatus int64 `json:"userStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		Password    respjson.Field
		Phone       respjson.Field
		Username    respjson.Field
		UserStatus  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserNewResponse) RawJSON() string { return r.JSON.raw }
func (r *UserNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
	// User Status
	UserStatus int64 `json:"userStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		Password    respjson.Field
		Phone       respjson.Field
		Username    respjson.Field
		UserStatus  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewWithListResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
	// User Status
	UserStatus int64 `json:"userStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		Password    respjson.Field
		Phone       respjson.Field
		Username    respjson.Field
		UserStatus  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserNewWithListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserNewWithListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UserVerifyIdentityResponseUnion contains all possible properties and values from
// [UserVerifyIdentityResponseBusinessIndividualVerification],
// [UserVerifyIdentityResponseBasicVerification].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UserVerifyIdentityResponseUnion struct {
	Token string `json:"token"`
	// This field is from variant
	// [UserVerifyIdentityResponseBusinessIndividualVerification].
	BeneficialOwnerIndividuals []UserVerifyIdentityResponseBusinessIndividualVerificationBeneficialOwnerIndividual `json:"beneficial_owner_individuals"`
	// This field is from variant
	// [UserVerifyIdentityResponseBusinessIndividualVerification].
	StatusReasons []UserVerifyIdentityResponseBusinessIndividualVerificationStatusReason `json:"status_reasons"`
	// This field is from variant
	// [UserVerifyIdentityResponseBusinessIndividualVerification].
	VerificationApplication UserVerifyIdentityResponseBusinessIndividualVerificationVerificationApplication `json:"verification_application"`
	// This field is from variant
	// [UserVerifyIdentityResponseBusinessIndividualVerification].
	VerificationDocument []UserVerifyIdentityResponseBusinessIndividualVerificationVerificationDocument `json:"verification_document"`
	// This field is from variant [UserVerifyIdentityResponseBasicVerification].
	Verified bool `json:"verified"`
	JSON     struct {
		Token                      respjson.Field
		BeneficialOwnerIndividuals respjson.Field
		StatusReasons              respjson.Field
		VerificationApplication    respjson.Field
		VerificationDocument       respjson.Field
		Verified                   respjson.Field
		raw                        string
	} `json:"-"`
}

func (u UserVerifyIdentityResponseUnion) AsBusinessIndividualVerification() (v UserVerifyIdentityResponseBusinessIndividualVerification) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserVerifyIdentityResponseUnion) AsBasicVerification() (v UserVerifyIdentityResponseBasicVerification) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserVerifyIdentityResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *UserVerifyIdentityResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserVerifyIdentityResponseBusinessIndividualVerification struct {
	Token string `json:"token" format:"uuid"`
	// Plural array: replicates lithic's re-cased Kybkyc element mint (control probe)
	BeneficialOwnerIndividuals []UserVerifyIdentityResponseBusinessIndividualVerificationBeneficialOwnerIndividual `json:"beneficial_owner_individuals"`
	// Plural array: second lithic-exact control probe
	StatusReasons []UserVerifyIdentityResponseBusinessIndividualVerificationStatusReason `json:"status_reasons"`
	// Nested object mirroring lithic's deeper plural mint
	VerificationApplication UserVerifyIdentityResponseBusinessIndividualVerificationVerificationApplication `json:"verification_application"`
	// Non-plural array: the case lithic never exercises (decisive probe)
	VerificationDocument []UserVerifyIdentityResponseBusinessIndividualVerificationVerificationDocument `json:"verification_document"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token                      respjson.Field
		BeneficialOwnerIndividuals respjson.Field
		StatusReasons              respjson.Field
		VerificationApplication    respjson.Field
		VerificationDocument       respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserVerifyIdentityResponseBusinessIndividualVerification) RawJSON() string { return r.JSON.raw }
func (r *UserVerifyIdentityResponseBusinessIndividualVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserVerifyIdentityResponseBusinessIndividualVerificationBeneficialOwnerIndividual struct {
	OwnerName string `json:"owner_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OwnerName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserVerifyIdentityResponseBusinessIndividualVerificationBeneficialOwnerIndividual) RawJSON() string {
	return r.JSON.raw
}
func (r *UserVerifyIdentityResponseBusinessIndividualVerificationBeneficialOwnerIndividual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserVerifyIdentityResponseBusinessIndividualVerificationStatusReason struct {
	ReasonCode string `json:"reason_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReasonCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserVerifyIdentityResponseBusinessIndividualVerificationStatusReason) RawJSON() string {
	return r.JSON.raw
}
func (r *UserVerifyIdentityResponseBusinessIndividualVerificationStatusReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Nested object mirroring lithic's deeper plural mint
type UserVerifyIdentityResponseBusinessIndividualVerificationVerificationApplication struct {
	StatusReasons []UserVerifyIdentityResponseBusinessIndividualVerificationVerificationApplicationStatusReason `json:"status_reasons"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StatusReasons respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserVerifyIdentityResponseBusinessIndividualVerificationVerificationApplication) RawJSON() string {
	return r.JSON.raw
}
func (r *UserVerifyIdentityResponseBusinessIndividualVerificationVerificationApplication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserVerifyIdentityResponseBusinessIndividualVerificationVerificationApplicationStatusReason struct {
	ApplicationNote string `json:"application_note"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationNote respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserVerifyIdentityResponseBusinessIndividualVerificationVerificationApplicationStatusReason) RawJSON() string {
	return r.JSON.raw
}
func (r *UserVerifyIdentityResponseBusinessIndividualVerificationVerificationApplicationStatusReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserVerifyIdentityResponseBusinessIndividualVerificationVerificationDocument struct {
	DocumentID string `json:"document_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserVerifyIdentityResponseBusinessIndividualVerificationVerificationDocument) RawJSON() string {
	return r.JSON.raw
}
func (r *UserVerifyIdentityResponseBusinessIndividualVerificationVerificationDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserVerifyIdentityResponseBasicVerification struct {
	Token    string `json:"token" format:"uuid"`
	Verified bool   `json:"verified"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Verified    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserVerifyIdentityResponseBasicVerification) RawJSON() string { return r.JSON.raw }
func (r *UserVerifyIdentityResponseBasicVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	ID        param.Opt[int64]  `json:"id,omitzero"`
	Email     param.Opt[string] `json:"email,omitzero"`
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	LastName  param.Opt[string] `json:"lastName,omitzero"`
	Password  param.Opt[string] `json:"password,omitzero"`
	Phone     param.Opt[string] `json:"phone,omitzero"`
	Username  param.Opt[string] `json:"username,omitzero"`
	// User Status
	UserStatus param.Opt[int64] `json:"userStatus,omitzero"`
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateParams struct {
	ID        param.Opt[int64]  `json:"id,omitzero"`
	Email     param.Opt[string] `json:"email,omitzero"`
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	LastName  param.Opt[string] `json:"lastName,omitzero"`
	Password  param.Opt[string] `json:"password,omitzero"`
	Phone     param.Opt[string] `json:"phone,omitzero"`
	Username  param.Opt[string] `json:"username,omitzero"`
	// User Status
	UserStatus param.Opt[int64] `json:"userStatus,omitzero"`
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewWithListParams struct {
	Body []UserNewWithListParamsBody
	paramObj
}

func (r UserNewWithListParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *UserNewWithListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewWithListParamsBody struct {
	ID        param.Opt[int64]  `json:"id,omitzero"`
	Email     param.Opt[string] `json:"email,omitzero"`
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	LastName  param.Opt[string] `json:"lastName,omitzero"`
	Password  param.Opt[string] `json:"password,omitzero"`
	Phone     param.Opt[string] `json:"phone,omitzero"`
	Username  param.Opt[string] `json:"username,omitzero"`
	// User Status
	UserStatus param.Opt[int64] `json:"userStatus,omitzero"`
	paramObj
}

func (r UserNewWithListParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow UserNewWithListParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewWithListParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoginParams struct {
	// The password for login in clear text
	Password param.Opt[string] `query:"password,omitzero" json:"-"`
	// The user name for login
	Username param.Opt[string] `query:"username,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserLoginParams]'s query parameters as `url.Values`.
func (r UserLoginParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
