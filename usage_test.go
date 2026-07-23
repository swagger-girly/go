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

func TestUsage(t *testing.T) {
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
	pet, err := client.Pet.Update(context.TODO(), helloworldtestingggg.PetUpdateParams{
		Pet: helloworldtestingggg.PetParam{
			Name:      helloworldtestingggg.F("doggie"),
			PhotoURLs: helloworldtestingggg.F([]string{"string"}),
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", pet.ID)
}
