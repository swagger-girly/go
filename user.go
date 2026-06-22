// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hello-world-testingggg-go/internal/apijson"
	"github.com/stainless-sdks/hello-world-testingggg-go/internal/apiquery"
	"github.com/stainless-sdks/hello-world-testingggg-go/internal/param"
	"github.com/stainless-sdks/hello-world-testingggg-go/internal/requestconfig"
	"github.com/stainless-sdks/hello-world-testingggg-go/option"
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

type UserNewResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
	// User Status
	UserStatus int64               `json:"userStatus"`
	JSON       userNewResponseJSON `json:"-"`
}

// userNewResponseJSON contains the JSON metadata for the struct [UserNewResponse]
type userNewResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	Password    apijson.Field
	Phone       apijson.Field
	Username    apijson.Field
	UserStatus  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userNewResponseJSON) RawJSON() string {
	return r.raw
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
	UserStatus int64               `json:"userStatus"`
	JSON       userGetResponseJSON `json:"-"`
}

// userGetResponseJSON contains the JSON metadata for the struct [UserGetResponse]
type userGetResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	Password    apijson.Field
	Phone       apijson.Field
	Username    apijson.Field
	UserStatus  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseJSON) RawJSON() string {
	return r.raw
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
	UserStatus int64                       `json:"userStatus"`
	JSON       userNewWithListResponseJSON `json:"-"`
}

// userNewWithListResponseJSON contains the JSON metadata for the struct
// [UserNewWithListResponse]
type userNewWithListResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	Password    apijson.Field
	Phone       apijson.Field
	Username    apijson.Field
	UserStatus  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserNewWithListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userNewWithListResponseJSON) RawJSON() string {
	return r.raw
}

type UserNewParams struct {
	ID        param.Field[int64]  `json:"id"`
	Email     param.Field[string] `json:"email"`
	FirstName param.Field[string] `json:"firstName"`
	LastName  param.Field[string] `json:"lastName"`
	Password  param.Field[string] `json:"password"`
	Phone     param.Field[string] `json:"phone"`
	Username  param.Field[string] `json:"username"`
	// User Status
	UserStatus param.Field[int64] `json:"userStatus"`
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParams struct {
	ID        param.Field[int64]  `json:"id"`
	Email     param.Field[string] `json:"email"`
	FirstName param.Field[string] `json:"firstName"`
	LastName  param.Field[string] `json:"lastName"`
	Password  param.Field[string] `json:"password"`
	Phone     param.Field[string] `json:"phone"`
	Username  param.Field[string] `json:"username"`
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
	ID        param.Field[int64]  `json:"id"`
	Email     param.Field[string] `json:"email"`
	FirstName param.Field[string] `json:"firstName"`
	LastName  param.Field[string] `json:"lastName"`
	Password  param.Field[string] `json:"password"`
	Phone     param.Field[string] `json:"phone"`
	Username  param.Field[string] `json:"username"`
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
