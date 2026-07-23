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
	)
	_, err := client.AI.AIQuery(context.TODO(), helloworldtestingggg.AIAIQueryParams{
		DataToExtract: helloworldtestingggg.F([]helloworldtestingggg.AiaiQueryParamsDataToExtract{{
			DatapointName: helloworldtestingggg.F("datapoint_name"),
			DatapointType: helloworldtestingggg.F(helloworldtestingggg.AiaiQueryParamsDataToExtractDatapointTypeText),
		}}),
		Domain: helloworldtestingggg.F("domain"),
		Authentication3DS: helloworldtestingggg.F(helloworldtestingggg.AiaiQueryParamsAuthentication3DS{
			ChallengeRules: helloworldtestingggg.F([]helloworldtestingggg.AiaiQueryParamsAuthentication3DsChallengeRule{{
				ChallengeKind: helloworldtestingggg.F("challenge_kind"),
			}}),
			FallbackCheck: helloworldtestingggg.F([]helloworldtestingggg.AiaiQueryParamsAuthentication3DsFallbackCheck{{
				CheckName: helloworldtestingggg.F("check_name"),
			}}),
		}),
		DatapointObjectSchema: helloworldtestingggg.F(map[string]helloworldtestingggg.AiaiQueryParamsDatapointObjectSchema{
			"foo": helloworldtestingggg.AiaiQueryParamsDatapointObjectSchemaString,
		}),
		ExtractionRules: helloworldtestingggg.F([]helloworldtestingggg.AiaiQueryParamsExtractionRule{{
			RuleName:     helloworldtestingggg.F("rule_name"),
			RuleSeverity: helloworldtestingggg.F(helloworldtestingggg.AiaiQueryParamsExtractionRulesRuleSeverityLow),
		}}),
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
