// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/swagger-girly/go"
	"github.com/swagger-girly/go/internal/testutil"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/shared"
)

func TestAdoptionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Adoptions.New(context.TODO(), helloworldtestingggg.AdoptionNewParams{
		OfIndividual: &helloworldtestingggg.AdoptionNewParamsBodyIndividual{
			Name: "name",
			Address: shared.AddressParam{
				City: helloworldtestingggg.String("Palo Alto"),
				Geo: shared.AddressGeoParam{
					Latitude:  37.4443,
					Longitude: -122.1598,
				},
				RelatedCategory: map[string]any{},
				RelatedCustomer: map[string]any{},
				RelatedMoney: shared.MoneyParam{
					Amount:          2500,
					Currency:        "USD",
					RelatedCategory: map[string]any{},
					RelatedCustomer: map[string]any{},
					RelatedOrder:    map[string]any{},
					RelatedPet: helloworldtestingggg.PetParam{
						Name:               "doggie",
						PhotoURLs:          []string{"string"},
						ID:                 helloworldtestingggg.Int(10),
						AcquisitionChannel: helloworldtestingggg.PetAcquisitionChannelBreeder,
						Category:           map[string]any{},
						MicrochipID: helloworldtestingggg.PetMicrochipIDUnionParam{
							OfString: helloworldtestingggg.String("string"),
						},
						RelatedCategory: map[string]any{},
						RelatedCustomer: map[string]any{},
						RelatedOrder:    map[string]any{},
						RelatedShelter:  map[string]any{},
						RelatedUser:     map[string]any{},
						Status:          helloworldtestingggg.PetStatusAvailable,
						Tags:            []any{map[string]any{}},
					},
					RelatedShelter: map[string]any{},
					RelatedTag:     map[string]any{},
					RelatedUser:    map[string]any{},
				},
				RelatedOrder: map[string]any{},
				RelatedPet: helloworldtestingggg.PetParam{
					Name:               "doggie",
					PhotoURLs:          []string{"string"},
					ID:                 helloworldtestingggg.Int(10),
					AcquisitionChannel: helloworldtestingggg.PetAcquisitionChannelBreeder,
					Category:           map[string]any{},
					MicrochipID: helloworldtestingggg.PetMicrochipIDUnionParam{
						OfString: helloworldtestingggg.String("string"),
					},
					RelatedCategory: map[string]any{},
					RelatedCustomer: map[string]any{},
					RelatedOrder:    map[string]any{},
					RelatedShelter:  map[string]any{},
					RelatedUser:     map[string]any{},
					Status:          helloworldtestingggg.PetStatusAvailable,
					Tags:            []any{map[string]any{}},
				},
				RelatedShelter: map[string]any{},
				RelatedTag:     map[string]any{},
				RelatedUser:    map[string]any{},
				State:          helloworldtestingggg.String("CA"),
				Street:         helloworldtestingggg.String("437 Lytton"),
				Zip:            helloworldtestingggg.String("94301"),
			},
			Age:   helloworldtestingggg.Int(0),
			Email: helloworldtestingggg.String("email"),
		},
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdoptionGet(t *testing.T) {
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
	_, err := client.Adoptions.Get(context.TODO(), "applicationId")
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdoptionGetDecision(t *testing.T) {
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
	_, err := client.Adoptions.GetDecision(context.TODO(), "applicationId")
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
