// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/swagger-girly/go"
	"github.com/swagger-girly/go/internal/testutil"
	"github.com/swagger-girly/go/option"
)

func TestStoreReportInventoryDailyGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := helloworldtestingggg.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBasicAuthUsername("My Basic Auth Username"),
		option.WithBasicAuthPassword("My Basic Auth Password"),
	)
	_, err := client.Store.Reports.Inventory.Daily.Get(
		context.TODO(),
		time.Now(),
		helloworldtestingggg.StoreReportInventoryDailyGetParams{
			ReportID: "reportId",
		},
	)
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
