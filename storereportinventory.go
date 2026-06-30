// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
)

// Access to Petstore orders
//
// StoreReportInventoryService contains methods and other services that help with
// interacting with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreReportInventoryService] method instead.
type StoreReportInventoryService struct {
	Options []option.RequestOption
	// Access to Petstore orders
	Daily *StoreReportInventoryDailyService
}

// NewStoreReportInventoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStoreReportInventoryService(opts ...option.RequestOption) (r *StoreReportInventoryService) {
	r = &StoreReportInventoryService{}
	r.Options = opts
	r.Daily = NewStoreReportInventoryDailyService(opts...)
	return
}

// Returns inventory rows associated with a store report.
func (r *StoreReportInventoryService) List(ctx context.Context, reportID string, opts ...option.RequestOption) (res *InventoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if reportID == "" {
		err = errors.New("missing required reportId parameter")
		return nil, err
	}
	path := fmt.Sprintf("store/reports/%s/inventory", reportID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
