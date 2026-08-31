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

func TestAIAIQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.AIQuery(context.TODO(), helloworldtestingggg.AIAIQueryParams{
		DataToExtract: []helloworldtestingggg.AIAIQueryParamsDataToExtract{{
			DatapointName: "datapoint_name",
			DatapointType: "text",
		}},
		Domain: "domain",
		Authentication3DS: helloworldtestingggg.AIAIQueryParamsAuthentication3DS{
			ChallengeRules: []helloworldtestingggg.AIAIQueryParamsAuthentication3DSChallengeRule{{
				ChallengeKind: helloworldtestingggg.String("challenge_kind"),
			}},
			FallbackCheck: []helloworldtestingggg.AIAIQueryParamsAuthentication3DSFallbackCheck{{
				CheckName: helloworldtestingggg.String("check_name"),
			}},
		},
		DatapointObjectSchema: map[string]string{
			"foo": "string",
		},
		ExtractionRules: []helloworldtestingggg.AIAIQueryParamsExtractionRule{{
			RuleName:     helloworldtestingggg.String("rule_name"),
			RuleSeverity: "low",
		}},
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
