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
)

func TestProfileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.New(context.TODO(), helloworldtestingggg.ProfileNewParams{
		ProfileCreateRequest: helloworldtestingggg.ProfileCreateRequestParam{
			DisplayName:  "Ada Lovelace",
			Email:        helloworldtestingggg.String("ada@example.com"),
			GovernmentID: helloworldtestingggg.String("governmentId"),
			Metadata: map[string]string{
				"crm_id": "crm_123",
			},
			Preferences: helloworldtestingggg.ProfileCreateRequestPreferencesParam{
				Alerts: map[string]bool{
					"inventory": true,
				},
				Newsletter:    helloworldtestingggg.Bool(true),
				ScoreByRegion: map[string]any{},
			},
			SecretNote: helloworldtestingggg.String("verified manually"),
		},
		IdempotencyKey: helloworldtestingggg.String("profile-create-123"),
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProfileGet(t *testing.T) {
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
	_, err := client.Profiles.Get(context.TODO(), "profileId")
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProfileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.Update(
		context.TODO(),
		"profileId",
		helloworldtestingggg.ProfileUpdateParams{
			OfProfileDetailsUpdate: &helloworldtestingggg.ProfileUpdateParamsBodyProfileDetailsUpdate{
				DisplayName: helloworldtestingggg.String("displayName"),
				Email:       helloworldtestingggg.String("email"),
				Metadata: map[string]string{
					"foo": "string",
				},
				Preferences: helloworldtestingggg.ProfileUpdateParamsBodyProfileDetailsUpdatePreferences{
					Alerts: map[string]bool{
						"foo": true,
					},
					Newsletter:    helloworldtestingggg.Bool(true),
					ScoreByRegion: map[string]any{},
				},
			},
			IdempotencyKey: helloworldtestingggg.String("profile-update-123"),
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

func TestProfileLegacySearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.LegacySearch(context.TODO(), helloworldtestingggg.ProfileLegacySearchParams{
		IncludeArchived: helloworldtestingggg.Bool(true),
		Q:               helloworldtestingggg.String("q"),
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
