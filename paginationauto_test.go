// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg_test

import (
	"context"
	"os"
	"testing"

	"github.com/swagger-girly/go"
	"github.com/swagger-girly/go/internal/testutil"
	"github.com/swagger-girly/go/option"
)

func TestAutoPagination(t *testing.T) {
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
	)
	iter := client.Pet.ListAutoPaging(context.TODO(), helloworldtestingggg.PetListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		pet := iter.Current()
		t.Logf("%+v\n", pet.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
