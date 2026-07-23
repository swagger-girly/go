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
	"github.com/swagger-girly/go/packages/pagination"
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
	Options []option.RequestOption
	// Access to Petstore orders
	Inventory *StoreReportInventoryService
}

// NewStoreReportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStoreReportService(opts ...option.RequestOption) (r *StoreReportService) {
	r = &StoreReportService{}
	r.Options = opts
	r.Inventory = NewStoreReportInventoryService(opts...)
	return
}

// Returns a generated store report.
func (r *StoreReportService) Get(ctx context.Context, reportID string, opts ...option.RequestOption) (res *Report, err error) {
	opts = slices.Concat(r.Options, opts)
	if reportID == "" {
		err = errors.New("missing required reportId parameter")
		return nil, err
	}
	path := fmt.Sprintf("store/reports/%s", reportID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a cursor-id paginated list of generated store reports.
func (r *StoreReportService) List(ctx context.Context, query StoreReportListParams, opts ...option.RequestOption) (res *pagination.ReportCursorPage[Report], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/html")}, opts...)
	if reportID == "" {
		err = errors.New("missing required reportId parameter")
		return nil, err
	}
	path := fmt.Sprintf("store/reports/%s/embed", reportID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Pauses a report generation job and returns an empty success response.
func (r *StoreReportService) Pause(ctx context.Context, reportID string, body StoreReportPauseParams, opts ...option.RequestOption) (err error) {
	if body.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", body.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if reportID == "" {
		err = errors.New("missing required reportId parameter")
		return err
	}
	path := fmt.Sprintf("store/reports/%s/pause", reportID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type DailyInventory struct {
	Date    time.Time          `json:"date" api:"required" format:"date"`
	Metrics interface{}        `json:"metrics" api:"required"`
	Notes   []string           `json:"notes" api:"nullable"`
	JSON    dailyInventoryJSON `json:"-"`
}

// dailyInventoryJSON contains the JSON metadata for the struct [DailyInventory]
type dailyInventoryJSON struct {
	Date        apijson.Field
	Metrics     apijson.Field
	Notes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DailyInventory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dailyInventoryJSON) RawJSON() string {
	return r.raw
}

type InventoryResponse struct {
	Data []InventoryResponseData `json:"data" api:"required"`
	JSON inventoryResponseJSON   `json:"-"`
}

// inventoryResponseJSON contains the JSON metadata for the struct
// [InventoryResponse]
type inventoryResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InventoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inventoryResponseJSON) RawJSON() string {
	return r.raw
}

type InventoryResponseData struct {
	Quantity   int64                     `json:"quantity" api:"required"`
	SKU        string                    `json:"sku" api:"required"`
	Attributes map[string]string         `json:"attributes"`
	PetID      int64                     `json:"petId" api:"nullable"`
	JSON       inventoryResponseDataJSON `json:"-"`
}

// inventoryResponseDataJSON contains the JSON metadata for the struct
// [InventoryResponseData]
type inventoryResponseDataJSON struct {
	Quantity    apijson.Field
	SKU         apijson.Field
	Attributes  apijson.Field
	PetID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InventoryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inventoryResponseDataJSON) RawJSON() string {
	return r.raw
}

type Report struct {
	ID          string                 `json:"id" api:"required"`
	CreatedAt   time.Time              `json:"createdAt" api:"required" format:"date-time"`
	Status      ReportStatus           `json:"status" api:"required"`
	CompletedAt time.Time              `json:"completedAt" api:"nullable" format:"date-time"`
	Filters     map[string]interface{} `json:"filters"`
	Totals      map[string]int64       `json:"totals"`
	JSON        reportJSON             `json:"-"`
}

// reportJSON contains the JSON metadata for the struct [Report]
type reportJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Status      apijson.Field
	CompletedAt apijson.Field
	Filters     apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Report) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportJSON) RawJSON() string {
	return r.raw
}

type ReportStatus string

const (
	ReportStatusQueued   ReportStatus = "queued"
	ReportStatusRunning  ReportStatus = "running"
	ReportStatusComplete ReportStatus = "complete"
	ReportStatusFailed   ReportStatus = "failed"
)

func (r ReportStatus) IsKnown() bool {
	switch r {
	case ReportStatusQueued, ReportStatusRunning, ReportStatusComplete, ReportStatusFailed:
		return true
	}
	return false
}

type ReportList struct {
	Data    []Report       `json:"data" api:"required"`
	HasMore bool           `json:"has_more" api:"required"`
	JSON    reportListJSON `json:"-"`
}

// reportListJSON contains the JSON metadata for the struct [ReportList]
type reportListJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportListJSON) RawJSON() string {
	return r.raw
}

type StoreReportListParams struct {
	// Return reports before this report id.
	EndingBefore param.Field[string] `query:"ending_before"`
	// Maximum number of reports to return.
	PageSize param.Field[int64] `query:"page_size"`
	// Return reports after this report id.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [StoreReportListParams]'s query parameters as `url.Values`.
func (r StoreReportListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StoreReportPauseParams struct {
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}
