// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"net/http"
	"slices"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/respjson"
)

// MediaService contains methods and other services that help with interacting with
// the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaService] method instead.
type MediaService struct {
	options []option.RequestOption
}

// NewMediaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMediaService(opts ...option.RequestOption) (r MediaService) {
	r = MediaService{}
	r.options = opts
	return
}

// A binary body alongside a JSON media type.
func (r *MediaService) BinaryAndJson(ctx context.Context, opts ...option.RequestOption) (res *MediaBinaryAndJsonResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "media/binary-and-json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// JSON alongside an XML media type.
func (r *MediaService) JsonAndXml(ctx context.Context, opts ...option.RequestOption) (res *MediaJsonAndXmlResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "media/json-and-xml"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// JSON listed before a text media type.
func (r *MediaService) JsonFirst(ctx context.Context, opts ...option.RequestOption) (res *MediaJsonFirstResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "media/json-first"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Schema present on application/json but absent on a text/plain listed first.
func (r *MediaService) SchemaOnJson(ctx context.Context, opts ...option.RequestOption) (res *MediaSchemaOnJsonResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "media/schema-on-json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Schema present on text/plain but absent on application/json.
func (r *MediaService) SchemaOnText(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "media/schema-on-text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Text media type listed before JSON (both with a schema).
func (r *MediaService) TextFirst(ctx context.Context, opts ...option.RequestOption) (res *MediaTextFirstResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "media/text-first"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A single text/plain body.
func (r *MediaService) TextOnly(ctx context.Context, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "media/text-only"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A vendor application/vnd.pet.v2+json media type.
func (r *MediaService) VendorJson(ctx context.Context, opts ...option.RequestOption) (res *MediaVendorJsonResponse, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.pet.v2+json")}, opts...)
	path := "media/vendor-json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MediaBinaryAndJsonResponse struct {
	ID    string `json:"id" api:"required"`
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBinaryAndJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBinaryAndJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaJsonAndXmlResponse struct {
	ID    string `json:"id" api:"required"`
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaJsonAndXmlResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaJsonAndXmlResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaJsonFirstResponse struct {
	ID    string `json:"id" api:"required"`
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaJsonFirstResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaJsonFirstResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaSchemaOnJsonResponse struct {
	ID    string `json:"id" api:"required"`
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaSchemaOnJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaSchemaOnJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaTextFirstResponse struct {
	ID    string `json:"id" api:"required"`
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaTextFirstResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaTextFirstResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaVendorJsonResponse struct {
	ID    string `json:"id" api:"required"`
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaVendorJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaVendorJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
