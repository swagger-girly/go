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
	"slices"

	"github.com/stainless-sdks/hello-world-testingggg-go/internal/apiform"
	"github.com/stainless-sdks/hello-world-testingggg-go/internal/apijson"
	"github.com/stainless-sdks/hello-world-testingggg-go/internal/apiquery"
	"github.com/stainless-sdks/hello-world-testingggg-go/internal/param"
	"github.com/stainless-sdks/hello-world-testingggg-go/internal/requestconfig"
	"github.com/stainless-sdks/hello-world-testingggg-go/option"
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
func (r *PetService) New(ctx context.Context, body PetNewParams, opts ...option.RequestOption) (res *PetNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single pet
func (r *PetService) Get(ctx context.Context, petID int64, opts ...option.RequestOption) (res *PetGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing pet by Id
func (r *PetService) Update(ctx context.Context, body PetUpdateParams, opts ...option.RequestOption) (res *PetUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
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
func (r *PetService) FindByStatus(ctx context.Context, query PetFindByStatusParams, opts ...option.RequestOption) (res *[]PetFindByStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pet/findByStatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3
// for testing.
func (r *PetService) FindByTags(ctx context.Context, query PetFindByTagsParams, opts ...option.RequestOption) (res *[]PetFindByTagsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pet/findByTags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

type PetNewResponse struct {
	Name      string                 `json:"name" api:"required"`
	PhotoURLs []string               `json:"photoUrls" api:"required"`
	ID        int64                  `json:"id"`
	Category  PetNewResponseCategory `json:"category"`
	// pet status in the store
	Status PetNewResponseStatus `json:"status"`
	Tags   []PetNewResponseTag  `json:"tags"`
	JSON   petNewResponseJSON   `json:"-"`
}

// petNewResponseJSON contains the JSON metadata for the struct [PetNewResponse]
type petNewResponseJSON struct {
	Name        apijson.Field
	PhotoURLs   apijson.Field
	ID          apijson.Field
	Category    apijson.Field
	Status      apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petNewResponseJSON) RawJSON() string {
	return r.raw
}

type PetNewResponseCategory struct {
	ID   int64                      `json:"id"`
	Name string                     `json:"name"`
	JSON petNewResponseCategoryJSON `json:"-"`
}

// petNewResponseCategoryJSON contains the JSON metadata for the struct
// [PetNewResponseCategory]
type petNewResponseCategoryJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetNewResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petNewResponseCategoryJSON) RawJSON() string {
	return r.raw
}

// pet status in the store
type PetNewResponseStatus string

const (
	PetNewResponseStatusAvailable PetNewResponseStatus = "available"
	PetNewResponseStatusPending   PetNewResponseStatus = "pending"
	PetNewResponseStatusSold      PetNewResponseStatus = "sold"
)

func (r PetNewResponseStatus) IsKnown() bool {
	switch r {
	case PetNewResponseStatusAvailable, PetNewResponseStatusPending, PetNewResponseStatusSold:
		return true
	}
	return false
}

type PetNewResponseTag struct {
	ID   int64                 `json:"id"`
	Name string                `json:"name"`
	JSON petNewResponseTagJSON `json:"-"`
}

// petNewResponseTagJSON contains the JSON metadata for the struct
// [PetNewResponseTag]
type petNewResponseTagJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetNewResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petNewResponseTagJSON) RawJSON() string {
	return r.raw
}

type PetGetResponse struct {
	Name      string                 `json:"name" api:"required"`
	PhotoURLs []string               `json:"photoUrls" api:"required"`
	ID        int64                  `json:"id"`
	Category  PetGetResponseCategory `json:"category"`
	// pet status in the store
	Status PetGetResponseStatus `json:"status"`
	Tags   []PetGetResponseTag  `json:"tags"`
	JSON   petGetResponseJSON   `json:"-"`
}

// petGetResponseJSON contains the JSON metadata for the struct [PetGetResponse]
type petGetResponseJSON struct {
	Name        apijson.Field
	PhotoURLs   apijson.Field
	ID          apijson.Field
	Category    apijson.Field
	Status      apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petGetResponseJSON) RawJSON() string {
	return r.raw
}

type PetGetResponseCategory struct {
	ID   int64                      `json:"id"`
	Name string                     `json:"name"`
	JSON petGetResponseCategoryJSON `json:"-"`
}

// petGetResponseCategoryJSON contains the JSON metadata for the struct
// [PetGetResponseCategory]
type petGetResponseCategoryJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetGetResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petGetResponseCategoryJSON) RawJSON() string {
	return r.raw
}

// pet status in the store
type PetGetResponseStatus string

const (
	PetGetResponseStatusAvailable PetGetResponseStatus = "available"
	PetGetResponseStatusPending   PetGetResponseStatus = "pending"
	PetGetResponseStatusSold      PetGetResponseStatus = "sold"
)

func (r PetGetResponseStatus) IsKnown() bool {
	switch r {
	case PetGetResponseStatusAvailable, PetGetResponseStatusPending, PetGetResponseStatusSold:
		return true
	}
	return false
}

type PetGetResponseTag struct {
	ID   int64                 `json:"id"`
	Name string                `json:"name"`
	JSON petGetResponseTagJSON `json:"-"`
}

// petGetResponseTagJSON contains the JSON metadata for the struct
// [PetGetResponseTag]
type petGetResponseTagJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetGetResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petGetResponseTagJSON) RawJSON() string {
	return r.raw
}

type PetUpdateResponse struct {
	Name      string                    `json:"name" api:"required"`
	PhotoURLs []string                  `json:"photoUrls" api:"required"`
	ID        int64                     `json:"id"`
	Category  PetUpdateResponseCategory `json:"category"`
	// pet status in the store
	Status PetUpdateResponseStatus `json:"status"`
	Tags   []PetUpdateResponseTag  `json:"tags"`
	JSON   petUpdateResponseJSON   `json:"-"`
}

// petUpdateResponseJSON contains the JSON metadata for the struct
// [PetUpdateResponse]
type petUpdateResponseJSON struct {
	Name        apijson.Field
	PhotoURLs   apijson.Field
	ID          apijson.Field
	Category    apijson.Field
	Status      apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PetUpdateResponseCategory struct {
	ID   int64                         `json:"id"`
	Name string                        `json:"name"`
	JSON petUpdateResponseCategoryJSON `json:"-"`
}

// petUpdateResponseCategoryJSON contains the JSON metadata for the struct
// [PetUpdateResponseCategory]
type petUpdateResponseCategoryJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetUpdateResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petUpdateResponseCategoryJSON) RawJSON() string {
	return r.raw
}

// pet status in the store
type PetUpdateResponseStatus string

const (
	PetUpdateResponseStatusAvailable PetUpdateResponseStatus = "available"
	PetUpdateResponseStatusPending   PetUpdateResponseStatus = "pending"
	PetUpdateResponseStatusSold      PetUpdateResponseStatus = "sold"
)

func (r PetUpdateResponseStatus) IsKnown() bool {
	switch r {
	case PetUpdateResponseStatusAvailable, PetUpdateResponseStatusPending, PetUpdateResponseStatusSold:
		return true
	}
	return false
}

type PetUpdateResponseTag struct {
	ID   int64                    `json:"id"`
	Name string                   `json:"name"`
	JSON petUpdateResponseTagJSON `json:"-"`
}

// petUpdateResponseTagJSON contains the JSON metadata for the struct
// [PetUpdateResponseTag]
type petUpdateResponseTagJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetUpdateResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petUpdateResponseTagJSON) RawJSON() string {
	return r.raw
}

type PetFindByStatusResponse struct {
	Name      string                          `json:"name" api:"required"`
	PhotoURLs []string                        `json:"photoUrls" api:"required"`
	ID        int64                           `json:"id"`
	Category  PetFindByStatusResponseCategory `json:"category"`
	// pet status in the store
	Status PetFindByStatusResponseStatus `json:"status"`
	Tags   []PetFindByStatusResponseTag  `json:"tags"`
	JSON   petFindByStatusResponseJSON   `json:"-"`
}

// petFindByStatusResponseJSON contains the JSON metadata for the struct
// [PetFindByStatusResponse]
type petFindByStatusResponseJSON struct {
	Name        apijson.Field
	PhotoURLs   apijson.Field
	ID          apijson.Field
	Category    apijson.Field
	Status      apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetFindByStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petFindByStatusResponseJSON) RawJSON() string {
	return r.raw
}

type PetFindByStatusResponseCategory struct {
	ID   int64                               `json:"id"`
	Name string                              `json:"name"`
	JSON petFindByStatusResponseCategoryJSON `json:"-"`
}

// petFindByStatusResponseCategoryJSON contains the JSON metadata for the struct
// [PetFindByStatusResponseCategory]
type petFindByStatusResponseCategoryJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetFindByStatusResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petFindByStatusResponseCategoryJSON) RawJSON() string {
	return r.raw
}

// pet status in the store
type PetFindByStatusResponseStatus string

const (
	PetFindByStatusResponseStatusAvailable PetFindByStatusResponseStatus = "available"
	PetFindByStatusResponseStatusPending   PetFindByStatusResponseStatus = "pending"
	PetFindByStatusResponseStatusSold      PetFindByStatusResponseStatus = "sold"
)

func (r PetFindByStatusResponseStatus) IsKnown() bool {
	switch r {
	case PetFindByStatusResponseStatusAvailable, PetFindByStatusResponseStatusPending, PetFindByStatusResponseStatusSold:
		return true
	}
	return false
}

type PetFindByStatusResponseTag struct {
	ID   int64                          `json:"id"`
	Name string                         `json:"name"`
	JSON petFindByStatusResponseTagJSON `json:"-"`
}

// petFindByStatusResponseTagJSON contains the JSON metadata for the struct
// [PetFindByStatusResponseTag]
type petFindByStatusResponseTagJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetFindByStatusResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petFindByStatusResponseTagJSON) RawJSON() string {
	return r.raw
}

type PetFindByTagsResponse struct {
	Name      string                        `json:"name" api:"required"`
	PhotoURLs []string                      `json:"photoUrls" api:"required"`
	ID        int64                         `json:"id"`
	Category  PetFindByTagsResponseCategory `json:"category"`
	// pet status in the store
	Status PetFindByTagsResponseStatus `json:"status"`
	Tags   []PetFindByTagsResponseTag  `json:"tags"`
	JSON   petFindByTagsResponseJSON   `json:"-"`
}

// petFindByTagsResponseJSON contains the JSON metadata for the struct
// [PetFindByTagsResponse]
type petFindByTagsResponseJSON struct {
	Name        apijson.Field
	PhotoURLs   apijson.Field
	ID          apijson.Field
	Category    apijson.Field
	Status      apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetFindByTagsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petFindByTagsResponseJSON) RawJSON() string {
	return r.raw
}

type PetFindByTagsResponseCategory struct {
	ID   int64                             `json:"id"`
	Name string                            `json:"name"`
	JSON petFindByTagsResponseCategoryJSON `json:"-"`
}

// petFindByTagsResponseCategoryJSON contains the JSON metadata for the struct
// [PetFindByTagsResponseCategory]
type petFindByTagsResponseCategoryJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetFindByTagsResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petFindByTagsResponseCategoryJSON) RawJSON() string {
	return r.raw
}

// pet status in the store
type PetFindByTagsResponseStatus string

const (
	PetFindByTagsResponseStatusAvailable PetFindByTagsResponseStatus = "available"
	PetFindByTagsResponseStatusPending   PetFindByTagsResponseStatus = "pending"
	PetFindByTagsResponseStatusSold      PetFindByTagsResponseStatus = "sold"
)

func (r PetFindByTagsResponseStatus) IsKnown() bool {
	switch r {
	case PetFindByTagsResponseStatusAvailable, PetFindByTagsResponseStatusPending, PetFindByTagsResponseStatusSold:
		return true
	}
	return false
}

type PetFindByTagsResponseTag struct {
	ID   int64                        `json:"id"`
	Name string                       `json:"name"`
	JSON petFindByTagsResponseTagJSON `json:"-"`
}

// petFindByTagsResponseTagJSON contains the JSON metadata for the struct
// [PetFindByTagsResponseTag]
type petFindByTagsResponseTagJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetFindByTagsResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petFindByTagsResponseTagJSON) RawJSON() string {
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

type PetNewParams struct {
	Name      param.Field[string]               `json:"name" api:"required"`
	PhotoURLs param.Field[[]string]             `json:"photoUrls" api:"required"`
	ID        param.Field[int64]                `json:"id"`
	Category  param.Field[PetNewParamsCategory] `json:"category"`
	// pet status in the store
	Status param.Field[PetNewParamsStatus] `json:"status"`
	Tags   param.Field[[]PetNewParamsTag]  `json:"tags"`
}

func (r PetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PetNewParamsCategory struct {
	ID   param.Field[int64]  `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r PetNewParamsCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pet status in the store
type PetNewParamsStatus string

const (
	PetNewParamsStatusAvailable PetNewParamsStatus = "available"
	PetNewParamsStatusPending   PetNewParamsStatus = "pending"
	PetNewParamsStatusSold      PetNewParamsStatus = "sold"
)

func (r PetNewParamsStatus) IsKnown() bool {
	switch r {
	case PetNewParamsStatusAvailable, PetNewParamsStatusPending, PetNewParamsStatusSold:
		return true
	}
	return false
}

type PetNewParamsTag struct {
	ID   param.Field[int64]  `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r PetNewParamsTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PetUpdateParams struct {
	Name      param.Field[string]                  `json:"name" api:"required"`
	PhotoURLs param.Field[[]string]                `json:"photoUrls" api:"required"`
	ID        param.Field[int64]                   `json:"id"`
	Category  param.Field[PetUpdateParamsCategory] `json:"category"`
	// pet status in the store
	Status param.Field[PetUpdateParamsStatus] `json:"status"`
	Tags   param.Field[[]PetUpdateParamsTag]  `json:"tags"`
}

func (r PetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PetUpdateParamsCategory struct {
	ID   param.Field[int64]  `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r PetUpdateParamsCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// pet status in the store
type PetUpdateParamsStatus string

const (
	PetUpdateParamsStatusAvailable PetUpdateParamsStatus = "available"
	PetUpdateParamsStatusPending   PetUpdateParamsStatus = "pending"
	PetUpdateParamsStatusSold      PetUpdateParamsStatus = "sold"
)

func (r PetUpdateParamsStatus) IsKnown() bool {
	switch r {
	case PetUpdateParamsStatusAvailable, PetUpdateParamsStatusPending, PetUpdateParamsStatusSold:
		return true
	}
	return false
}

type PetUpdateParamsTag struct {
	ID   param.Field[int64]  `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r PetUpdateParamsTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
