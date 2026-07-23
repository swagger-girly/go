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
			DisplayName:  helloworldtestingggg.F("Ada Lovelace"),
			Email:        helloworldtestingggg.F("ada@example.com"),
			GovernmentID: helloworldtestingggg.F("governmentId"),
			Metadata: helloworldtestingggg.F(map[string]string{
				"crm_id": "crm_123",
			}),
			Preferences: helloworldtestingggg.F(helloworldtestingggg.ProfileCreateRequestPreferencesParam{
				Alerts: helloworldtestingggg.F(map[string]bool{
					"inventory": true,
				}),
				Newsletter:    helloworldtestingggg.F(true),
				ScoreByRegion: helloworldtestingggg.F[any](map[string]interface{}{}),
			}),
			SecretNote: helloworldtestingggg.F("verified manually"),
		},
		IdempotencyKey: helloworldtestingggg.F("profile-create-123"),
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
			Body: helloworldtestingggg.ProfileUpdateParamsBodyProfileDetailsUpdate{
				DisplayName: helloworldtestingggg.F("displayName"),
				Email:       helloworldtestingggg.F("email"),
				Metadata: helloworldtestingggg.F(map[string]string{
					"foo": "string",
				}),
				Preferences: helloworldtestingggg.F(helloworldtestingggg.ProfileUpdateParamsBodyProfileDetailsUpdatePreferences{
					Alerts: helloworldtestingggg.F(map[string]bool{
						"foo": true,
					}),
					Newsletter:    helloworldtestingggg.F(true),
					ScoreByRegion: helloworldtestingggg.F[any](map[string]interface{}{}),
				}),
			},
			IdempotencyKey: helloworldtestingggg.F("profile-update-123"),
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
		IncludeArchived: helloworldtestingggg.F(true),
		Q:               helloworldtestingggg.F("q"),
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
