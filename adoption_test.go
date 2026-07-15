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
	)
	_, err := client.Adoptions.New(context.TODO(), helloworldtestingggg.AdoptionNewParams{
		Body: helloworldtestingggg.AdoptionNewParamsBodyIndividualApplicant{
			ApplicantType: helloworldtestingggg.F(helloworldtestingggg.AdoptionNewParamsBodyIndividualApplicantApplicantTypeIndividual),
			Name:          helloworldtestingggg.F("name"),
			Address: helloworldtestingggg.F(shared.AddressParam{
				City: helloworldtestingggg.F("Palo Alto"),
				Geo: helloworldtestingggg.F(shared.AddressGeoParam{
					Latitude:  helloworldtestingggg.F(37.444300),
					Longitude: helloworldtestingggg.F(-122.159800),
				}),
				State:  helloworldtestingggg.F("CA"),
				Street: helloworldtestingggg.F("437 Lytton"),
				Zip:    helloworldtestingggg.F("94301"),
			}),
			Age:   helloworldtestingggg.F(int64(0)),
			Email: helloworldtestingggg.F("email"),
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
