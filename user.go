// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/apiquery"
	"github.com/swagger-girly/go/internal/param"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/shared"
	"github.com/tidwall/gjson"
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
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	return
}

// This can only be done by the logged in user.
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *UserNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get user by user name
func (r *UserService) Get(ctx context.Context, username string, opts ...option.RequestOption) (res *UserGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if username == "" {
		err = errors.New("missing required username parameter")
		return nil, err
	}
	path := fmt.Sprintf("user/%s", username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This can only be done by the logged in user.
func (r *UserService) Update(ctx context.Context, username string, body UserUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return err
	}
	path := fmt.Sprintf("user/%s", username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// This can only be done by the logged in user.
func (r *UserService) Delete(ctx context.Context, username string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return err
	}
	path := fmt.Sprintf("user/%s", username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Creates list of users with given input array
func (r *UserService) NewWithList(ctx context.Context, body UserNewWithListParams, opts ...option.RequestOption) (res *UserNewWithListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "user/createWithList"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Logs user into the system
func (r *UserService) Login(ctx context.Context, query UserLoginParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "user/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Logs out current logged in user session
func (r *UserService) Logout(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "user/logout"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Casing probe replicating lithic: inline oneOf response whose variant is a
// kebab-named component (kyb-kyc-verification) so adjacent default initialisms
// glue into KYBKYC in the operation-scoped variant name
func (r *UserService) VerifyIdentity(ctx context.Context, username string, opts ...option.RequestOption) (res *UserVerifyIdentityResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if username == "" {
		err = errors.New("missing required username parameter")
		return nil, err
	}
	path := fmt.Sprintf("user/%s/verifyIdentity", username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type UserNewResponse struct {
	ID              int64          `json:"id"`
	Email           string         `json:"email"`
	FirstName       string         `json:"firstName"`
	LastName        string         `json:"lastName"`
	Password        string         `json:"password"`
	Phone           string         `json:"phone"`
	RelatedAddress  shared.Address `json:"relatedAddress"`
	RelatedCategory interface{}    `json:"relatedCategory"`
	RelatedCustomer interface{}    `json:"relatedCustomer"`
	RelatedMoney    shared.Money   `json:"relatedMoney"`
	RelatedOrder    interface{}    `json:"relatedOrder"`
	RelatedPet      Pet            `json:"relatedPet"`
	RelatedShelter  interface{}    `json:"relatedShelter"`
	RelatedTag      interface{}    `json:"relatedTag"`
	Username        string         `json:"username"`
	// User Status
	UserStatus int64               `json:"userStatus"`
	JSON       userNewResponseJSON `json:"-"`
}

// userNewResponseJSON contains the JSON metadata for the struct [UserNewResponse]
type userNewResponseJSON struct {
	ID              apijson.Field
	Email           apijson.Field
	FirstName       apijson.Field
	LastName        apijson.Field
	Password        apijson.Field
	Phone           apijson.Field
	RelatedAddress  apijson.Field
	RelatedCategory apijson.Field
	RelatedCustomer apijson.Field
	RelatedMoney    apijson.Field
	RelatedOrder    apijson.Field
	RelatedPet      apijson.Field
	RelatedShelter  apijson.Field
	RelatedTag      apijson.Field
	Username        apijson.Field
	UserStatus      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userNewResponseJSON) RawJSON() string {
	return r.raw
}

type UserGetResponse struct {
	ID              int64          `json:"id"`
	Email           string         `json:"email"`
	FirstName       string         `json:"firstName"`
	LastName        string         `json:"lastName"`
	Password        string         `json:"password"`
	Phone           string         `json:"phone"`
	RelatedAddress  shared.Address `json:"relatedAddress"`
	RelatedCategory interface{}    `json:"relatedCategory"`
	RelatedCustomer interface{}    `json:"relatedCustomer"`
	RelatedMoney    shared.Money   `json:"relatedMoney"`
	RelatedOrder    interface{}    `json:"relatedOrder"`
	RelatedPet      Pet            `json:"relatedPet"`
	RelatedShelter  interface{}    `json:"relatedShelter"`
	RelatedTag      interface{}    `json:"relatedTag"`
	Username        string         `json:"username"`
	// User Status
	UserStatus int64               `json:"userStatus"`
	JSON       userGetResponseJSON `json:"-"`
}

// userGetResponseJSON contains the JSON metadata for the struct [UserGetResponse]
type userGetResponseJSON struct {
	ID              apijson.Field
	Email           apijson.Field
	FirstName       apijson.Field
	LastName        apijson.Field
	Password        apijson.Field
	Phone           apijson.Field
	RelatedAddress  apijson.Field
	RelatedCategory apijson.Field
	RelatedCustomer apijson.Field
	RelatedMoney    apijson.Field
	RelatedOrder    apijson.Field
	RelatedPet      apijson.Field
	RelatedShelter  apijson.Field
	RelatedTag      apijson.Field
	Username        apijson.Field
	UserStatus      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseJSON) RawJSON() string {
	return r.raw
}

type UserNewWithListResponse struct {
	ID              int64          `json:"id"`
	Email           string         `json:"email"`
	FirstName       string         `json:"firstName"`
	LastName        string         `json:"lastName"`
	Password        string         `json:"password"`
	Phone           string         `json:"phone"`
	RelatedAddress  shared.Address `json:"relatedAddress"`
	RelatedCategory interface{}    `json:"relatedCategory"`
	RelatedCustomer interface{}    `json:"relatedCustomer"`
	RelatedMoney    shared.Money   `json:"relatedMoney"`
	RelatedOrder    interface{}    `json:"relatedOrder"`
	RelatedPet      Pet            `json:"relatedPet"`
	RelatedShelter  interface{}    `json:"relatedShelter"`
	RelatedTag      interface{}    `json:"relatedTag"`
	Username        string         `json:"username"`
	// User Status
	UserStatus int64                       `json:"userStatus"`
	JSON       userNewWithListResponseJSON `json:"-"`
}

// userNewWithListResponseJSON contains the JSON metadata for the struct
// [UserNewWithListResponse]
type userNewWithListResponseJSON struct {
	ID              apijson.Field
	Email           apijson.Field
	FirstName       apijson.Field
	LastName        apijson.Field
	Password        apijson.Field
	Phone           apijson.Field
	RelatedAddress  apijson.Field
	RelatedCategory apijson.Field
	RelatedCustomer apijson.Field
	RelatedMoney    apijson.Field
	RelatedOrder    apijson.Field
	RelatedPet      apijson.Field
	RelatedShelter  apijson.Field
	RelatedTag      apijson.Field
	Username        apijson.Field
	UserStatus      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserNewWithListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userNewWithListResponseJSON) RawJSON() string {
	return r.raw
}

type UserVerifyIdentityResponse struct {
	Token string `json:"token" format:"uuid"`
	// This field can have the runtime type of
	// [[]UserVerifyIdentityResponseKybkycVerificationBeneficialOwnerIndividual].
	BeneficialOwnerIndividuals interface{} `json:"beneficial_owner_individuals"`
	// This field can have the runtime type of
	// [[]UserVerifyIdentityResponseKybkycVerificationStatusReason].
	StatusReasons interface{} `json:"status_reasons"`
	// This field can have the runtime type of
	// [UserVerifyIdentityResponseKYBKYCVerificationVerificationApplication].
	VerificationApplication interface{} `json:"verification_application"`
	// This field can have the runtime type of
	// [[]UserVerifyIdentityResponseKybkycVerificationVerificationDocument].
	VerificationDocument interface{}                    `json:"verification_document"`
	Verified             bool                           `json:"verified"`
	JSON                 userVerifyIdentityResponseJSON `json:"-"`
	union                UserVerifyIdentityResponseUnion
}

// userVerifyIdentityResponseJSON contains the JSON metadata for the struct
// [UserVerifyIdentityResponse]
type userVerifyIdentityResponseJSON struct {
	Token                      apijson.Field
	BeneficialOwnerIndividuals apijson.Field
	StatusReasons              apijson.Field
	VerificationApplication    apijson.Field
	VerificationDocument       apijson.Field
	Verified                   apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r userVerifyIdentityResponseJSON) RawJSON() string {
	return r.raw
}

func (r *UserVerifyIdentityResponse) UnmarshalJSON(data []byte) (err error) {
	*r = UserVerifyIdentityResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserVerifyIdentityResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserVerifyIdentityResponseKYBKYCVerification],
// [UserVerifyIdentityResponseBasicVerification].
func (r UserVerifyIdentityResponse) AsUnion() UserVerifyIdentityResponseUnion {
	return r.union
}

// Union satisfied by [UserVerifyIdentityResponseKYBKYCVerification] or
// [UserVerifyIdentityResponseBasicVerification].
type UserVerifyIdentityResponseUnion interface {
	implementsUserVerifyIdentityResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserVerifyIdentityResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserVerifyIdentityResponseKYBKYCVerification{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserVerifyIdentityResponseBasicVerification{}),
		},
	)
}

type UserVerifyIdentityResponseKYBKYCVerification struct {
	Token string `json:"token" format:"uuid"`
	// Plural array: replicates lithic's re-cased Kybkyc element mint (control probe)
	BeneficialOwnerIndividuals []UserVerifyIdentityResponseKybkycVerificationBeneficialOwnerIndividual `json:"beneficial_owner_individuals"`
	// Plural array: second lithic-exact control probe
	StatusReasons []UserVerifyIdentityResponseKybkycVerificationStatusReason `json:"status_reasons"`
	// Nested object mirroring lithic's deeper plural mint
	VerificationApplication UserVerifyIdentityResponseKYBKYCVerificationVerificationApplication `json:"verification_application"`
	// Non-plural array: the case lithic never exercises (decisive probe)
	VerificationDocument []UserVerifyIdentityResponseKybkycVerificationVerificationDocument `json:"verification_document"`
	JSON                 userVerifyIdentityResponseKybkycVerificationJSON                   `json:"-"`
}

// userVerifyIdentityResponseKybkycVerificationJSON contains the JSON metadata for
// the struct [UserVerifyIdentityResponseKYBKYCVerification]
type userVerifyIdentityResponseKybkycVerificationJSON struct {
	Token                      apijson.Field
	BeneficialOwnerIndividuals apijson.Field
	StatusReasons              apijson.Field
	VerificationApplication    apijson.Field
	VerificationDocument       apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UserVerifyIdentityResponseKYBKYCVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userVerifyIdentityResponseKybkycVerificationJSON) RawJSON() string {
	return r.raw
}

func (r UserVerifyIdentityResponseKYBKYCVerification) implementsUserVerifyIdentityResponse() {}

type UserVerifyIdentityResponseKybkycVerificationBeneficialOwnerIndividual struct {
	OwnerName string                                                                    `json:"owner_name"`
	JSON      userVerifyIdentityResponseKybkycVerificationBeneficialOwnerIndividualJSON `json:"-"`
}

// userVerifyIdentityResponseKybkycVerificationBeneficialOwnerIndividualJSON
// contains the JSON metadata for the struct
// [UserVerifyIdentityResponseKybkycVerificationBeneficialOwnerIndividual]
type userVerifyIdentityResponseKybkycVerificationBeneficialOwnerIndividualJSON struct {
	OwnerName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserVerifyIdentityResponseKybkycVerificationBeneficialOwnerIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userVerifyIdentityResponseKybkycVerificationBeneficialOwnerIndividualJSON) RawJSON() string {
	return r.raw
}

type UserVerifyIdentityResponseKybkycVerificationStatusReason struct {
	ReasonCode string                                                       `json:"reason_code"`
	JSON       userVerifyIdentityResponseKybkycVerificationStatusReasonJSON `json:"-"`
}

// userVerifyIdentityResponseKybkycVerificationStatusReasonJSON contains the JSON
// metadata for the struct
// [UserVerifyIdentityResponseKybkycVerificationStatusReason]
type userVerifyIdentityResponseKybkycVerificationStatusReasonJSON struct {
	ReasonCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserVerifyIdentityResponseKybkycVerificationStatusReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userVerifyIdentityResponseKybkycVerificationStatusReasonJSON) RawJSON() string {
	return r.raw
}

// Nested object mirroring lithic's deeper plural mint
type UserVerifyIdentityResponseKYBKYCVerificationVerificationApplication struct {
	StatusReasons []UserVerifyIdentityResponseKybkycVerificationVerificationApplicationStatusReason `json:"status_reasons"`
	JSON          userVerifyIdentityResponseKybkycVerificationVerificationApplicationJSON           `json:"-"`
}

// userVerifyIdentityResponseKybkycVerificationVerificationApplicationJSON contains
// the JSON metadata for the struct
// [UserVerifyIdentityResponseKYBKYCVerificationVerificationApplication]
type userVerifyIdentityResponseKybkycVerificationVerificationApplicationJSON struct {
	StatusReasons apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserVerifyIdentityResponseKYBKYCVerificationVerificationApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userVerifyIdentityResponseKybkycVerificationVerificationApplicationJSON) RawJSON() string {
	return r.raw
}

type UserVerifyIdentityResponseKybkycVerificationVerificationApplicationStatusReason struct {
	ApplicationNote string                                                                              `json:"application_note"`
	JSON            userVerifyIdentityResponseKybkycVerificationVerificationApplicationStatusReasonJSON `json:"-"`
}

// userVerifyIdentityResponseKybkycVerificationVerificationApplicationStatusReasonJSON
// contains the JSON metadata for the struct
// [UserVerifyIdentityResponseKybkycVerificationVerificationApplicationStatusReason]
type userVerifyIdentityResponseKybkycVerificationVerificationApplicationStatusReasonJSON struct {
	ApplicationNote apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserVerifyIdentityResponseKybkycVerificationVerificationApplicationStatusReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userVerifyIdentityResponseKybkycVerificationVerificationApplicationStatusReasonJSON) RawJSON() string {
	return r.raw
}

type UserVerifyIdentityResponseKybkycVerificationVerificationDocument struct {
	DocumentID string                                                               `json:"document_id"`
	JSON       userVerifyIdentityResponseKybkycVerificationVerificationDocumentJSON `json:"-"`
}

// userVerifyIdentityResponseKybkycVerificationVerificationDocumentJSON contains
// the JSON metadata for the struct
// [UserVerifyIdentityResponseKybkycVerificationVerificationDocument]
type userVerifyIdentityResponseKybkycVerificationVerificationDocumentJSON struct {
	DocumentID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserVerifyIdentityResponseKybkycVerificationVerificationDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userVerifyIdentityResponseKybkycVerificationVerificationDocumentJSON) RawJSON() string {
	return r.raw
}

type UserVerifyIdentityResponseBasicVerification struct {
	Token    string                                          `json:"token" format:"uuid"`
	Verified bool                                            `json:"verified"`
	JSON     userVerifyIdentityResponseBasicVerificationJSON `json:"-"`
}

// userVerifyIdentityResponseBasicVerificationJSON contains the JSON metadata for
// the struct [UserVerifyIdentityResponseBasicVerification]
type userVerifyIdentityResponseBasicVerificationJSON struct {
	Token       apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserVerifyIdentityResponseBasicVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userVerifyIdentityResponseBasicVerificationJSON) RawJSON() string {
	return r.raw
}

func (r UserVerifyIdentityResponseBasicVerification) implementsUserVerifyIdentityResponse() {}

type UserNewParams struct {
	ID              param.Field[int64]               `json:"id"`
	Email           param.Field[string]              `json:"email"`
	FirstName       param.Field[string]              `json:"firstName"`
	LastName        param.Field[string]              `json:"lastName"`
	Password        param.Field[string]              `json:"password"`
	Phone           param.Field[string]              `json:"phone"`
	RelatedAddress  param.Field[shared.AddressParam] `json:"relatedAddress"`
	RelatedCategory param.Field[interface{}]         `json:"relatedCategory"`
	RelatedCustomer param.Field[interface{}]         `json:"relatedCustomer"`
	RelatedMoney    param.Field[shared.MoneyParam]   `json:"relatedMoney"`
	RelatedOrder    param.Field[interface{}]         `json:"relatedOrder"`
	RelatedPet      param.Field[PetParam]            `json:"relatedPet"`
	RelatedShelter  param.Field[interface{}]         `json:"relatedShelter"`
	RelatedTag      param.Field[interface{}]         `json:"relatedTag"`
	Username        param.Field[string]              `json:"username"`
	// User Status
	UserStatus param.Field[int64] `json:"userStatus"`
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParams struct {
	ID              param.Field[int64]               `json:"id"`
	Email           param.Field[string]              `json:"email"`
	FirstName       param.Field[string]              `json:"firstName"`
	LastName        param.Field[string]              `json:"lastName"`
	Password        param.Field[string]              `json:"password"`
	Phone           param.Field[string]              `json:"phone"`
	RelatedAddress  param.Field[shared.AddressParam] `json:"relatedAddress"`
	RelatedCategory param.Field[interface{}]         `json:"relatedCategory"`
	RelatedCustomer param.Field[interface{}]         `json:"relatedCustomer"`
	RelatedMoney    param.Field[shared.MoneyParam]   `json:"relatedMoney"`
	RelatedOrder    param.Field[interface{}]         `json:"relatedOrder"`
	RelatedPet      param.Field[PetParam]            `json:"relatedPet"`
	RelatedShelter  param.Field[interface{}]         `json:"relatedShelter"`
	RelatedTag      param.Field[interface{}]         `json:"relatedTag"`
	Username        param.Field[string]              `json:"username"`
	// User Status
	UserStatus param.Field[int64] `json:"userStatus"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserNewWithListParams struct {
	Body []UserNewWithListParamsBody `json:"body"`
}

func (r UserNewWithListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type UserNewWithListParamsBody struct {
	ID              param.Field[int64]               `json:"id"`
	Email           param.Field[string]              `json:"email"`
	FirstName       param.Field[string]              `json:"firstName"`
	LastName        param.Field[string]              `json:"lastName"`
	Password        param.Field[string]              `json:"password"`
	Phone           param.Field[string]              `json:"phone"`
	RelatedAddress  param.Field[shared.AddressParam] `json:"relatedAddress"`
	RelatedCategory param.Field[interface{}]         `json:"relatedCategory"`
	RelatedCustomer param.Field[interface{}]         `json:"relatedCustomer"`
	RelatedMoney    param.Field[shared.MoneyParam]   `json:"relatedMoney"`
	RelatedOrder    param.Field[interface{}]         `json:"relatedOrder"`
	RelatedPet      param.Field[PetParam]            `json:"relatedPet"`
	RelatedShelter  param.Field[interface{}]         `json:"relatedShelter"`
	RelatedTag      param.Field[interface{}]         `json:"relatedTag"`
	Username        param.Field[string]              `json:"username"`
	// User Status
	UserStatus param.Field[int64] `json:"userStatus"`
}

func (r UserNewWithListParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserLoginParams struct {
	// The password for login in clear text
	Password param.Field[string] `query:"password"`
	// The user name for login
	Username param.Field[string] `query:"username"`
}

// URLQuery serializes [UserLoginParams]'s query parameters as `url.Values`.
func (r UserLoginParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
