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
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/pagination"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
)

// Access to Petstore orders
//
// StoreReportService contains methods and other services that help with
// interacting with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreReportService] method instead.
type StoreReportService struct {
	options []option.RequestOption
	// Access to Petstore orders
	Inventory StoreReportInventoryService
}

// NewStoreReportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStoreReportService(opts ...option.RequestOption) (r StoreReportService) {
	r = StoreReportService{}
	r.options = opts
	r.Inventory = NewStoreReportInventoryService(opts...)
	return
}

// Returns a generated store report.
func (r *StoreReportService) Get(ctx context.Context, reportID string, opts ...option.RequestOption) (res *Report, err error) {
	opts = slices.Concat(r.options, opts)
	if reportID == "" {
		err = errors.New("missing required reportId parameter")
		return nil, err
	}
	path := fmt.Sprintf("store/reports/%s", url.PathEscape(reportID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a cursor-id paginated list of generated store reports.
func (r *StoreReportService) List(ctx context.Context, query StoreReportListParams, opts ...option.RequestOption) (res *pagination.ReportCursorPage[Report], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "store/reports"
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

// Returns a cursor-id paginated list of generated store reports.
func (r *StoreReportService) ListAutoPaging(ctx context.Context, query StoreReportListParams, opts ...option.RequestOption) *pagination.ReportCursorPageAutoPager[Report] {
	return pagination.NewReportCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Returns an embeddable HTML representation of a store report.
func (r *StoreReportService) Embed(ctx context.Context, reportID string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/html")}, opts...)
	if reportID == "" {
		err = errors.New("missing required reportId parameter")
		return nil, err
	}
	path := fmt.Sprintf("store/reports/%s/embed", url.PathEscape(reportID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Pauses a report generation job and returns an empty success response.
func (r *StoreReportService) Pause(ctx context.Context, reportID string, body StoreReportPauseParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", body.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if reportID == "" {
		err = errors.New("missing required reportId parameter")
		return err
	}
	path := fmt.Sprintf("store/reports/%s/pause", url.PathEscape(reportID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type DailyInventory struct {
	Date    time.Time `json:"date" api:"required" format:"date"`
	Metrics any       `json:"metrics" api:"required"`
	Notes   []string  `json:"notes" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Metrics     respjson.Field
		Notes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DailyInventory) RawJSON() string { return r.JSON.raw }
func (r *DailyInventory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InventoryResponse struct {
	Data []InventoryResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InventoryResponse) RawJSON() string { return r.JSON.raw }
func (r *InventoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InventoryResponseData struct {
	Quantity   int64             `json:"quantity" api:"required"`
	SKU        string            `json:"sku" api:"required"`
	Attributes map[string]string `json:"attributes"`
	PetID      int64             `json:"petId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		SKU         respjson.Field
		Attributes  respjson.Field
		PetID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InventoryResponseData) RawJSON() string { return r.JSON.raw }
func (r *InventoryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Report struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "queued", "running", "complete", "failed".
	Status      ReportStatus     `json:"status" api:"required"`
	CompletedAt time.Time        `json:"completedAt" api:"nullable" format:"date-time"`
	Filters     map[string]any   `json:"filters"`
	Totals      map[string]int64 `json:"totals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Status      respjson.Field
		CompletedAt respjson.Field
		Filters     respjson.Field
		Totals      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Report) RawJSON() string { return r.JSON.raw }
func (r *Report) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportStatus string

const (
	ReportStatusQueued   ReportStatus = "queued"
	ReportStatusRunning  ReportStatus = "running"
	ReportStatusComplete ReportStatus = "complete"
	ReportStatusFailed   ReportStatus = "failed"
)

type ReportList struct {
	Data    []Report `json:"data" api:"required"`
	HasMore bool     `json:"has_more" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportList) RawJSON() string { return r.JSON.raw }
func (r *ReportList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoreReportListParams struct {
	// Return reports before this report id.
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" json:"-"`
	// Maximum number of reports to return.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Return reports after this report id.
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StoreReportListParams]'s query parameters as `url.Values`.
func (r StoreReportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StoreReportPauseParams struct {
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}
