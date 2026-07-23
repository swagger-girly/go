// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
)

// Access to Petstore orders
//
// StoreReportInventoryDailyService contains methods and other services that help
// with interacting with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreReportInventoryDailyService] method instead.
type StoreReportInventoryDailyService struct {
	Options []option.RequestOption
}

// NewStoreReportInventoryDailyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewStoreReportInventoryDailyService(opts ...option.RequestOption) (r *StoreReportInventoryDailyService) {
	r = &StoreReportInventoryDailyService{}
	r.Options = opts
	return
}

// Returns daily inventory metrics for a generated report.
func (r *StoreReportInventoryDailyService) Get(ctx context.Context, reportID string, date time.Time, opts ...option.RequestOption) (res *DailyInventory, err error) {
	opts = slices.Concat(r.Options, opts)
	if reportID == "" {
		err = errors.New("missing required reportId parameter")
		return nil, err
	}
	path := fmt.Sprintf("store/reports/%s/inventory/daily/%s", reportID, date.Format("2006-01-02"))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
