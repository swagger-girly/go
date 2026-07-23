// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"net/http"
	"slices"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/param"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
)

// AIService contains methods and other services that help with interacting with
// the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIService] method instead.
type AIService struct {
	Options []option.RequestOption
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r *AIService) {
	r = &AIService{}
	r.Options = opts
	return
}

// Casing probe replicating context.dev: glued AIAI prefix (resource ai + method
// ai_query) over plural and non-plural synthesized element mints
func (r *AIService) AIQuery(ctx context.Context, body AIAIQueryParams, opts ...option.RequestOption) (res *AiaiQueryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AiaiQueryResponse struct {
	// Non-plural array on the response side
	DataExtracted []AiaiQueryResponseDataExtracted `json:"data_extracted"`
	// Plural array on the response side
	MatchedRules []AiaiQueryResponseMatchedRule `json:"matched_rules"`
	JSON         aiaiQueryResponseJSON          `json:"-"`
}

// aiaiQueryResponseJSON contains the JSON metadata for the struct
// [AiaiQueryResponse]
type aiaiQueryResponseJSON struct {
	DataExtracted apijson.Field
	MatchedRules  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AiaiQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiaiQueryResponseJSON) RawJSON() string {
	return r.raw
}

type AiaiQueryResponseDataExtracted struct {
	DatapointValue string                             `json:"datapoint_value"`
	JSON           aiaiQueryResponseDataExtractedJSON `json:"-"`
}

// aiaiQueryResponseDataExtractedJSON contains the JSON metadata for the struct
// [AiaiQueryResponseDataExtracted]
type aiaiQueryResponseDataExtractedJSON struct {
	DatapointValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AiaiQueryResponseDataExtracted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiaiQueryResponseDataExtractedJSON) RawJSON() string {
	return r.raw
}

type AiaiQueryResponseMatchedRule struct {
	MatchScore float64                          `json:"match_score"`
	JSON       aiaiQueryResponseMatchedRuleJSON `json:"-"`
}

// aiaiQueryResponseMatchedRuleJSON contains the JSON metadata for the struct
// [AiaiQueryResponseMatchedRule]
type aiaiQueryResponseMatchedRuleJSON struct {
	MatchScore  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AiaiQueryResponseMatchedRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiaiQueryResponseMatchedRuleJSON) RawJSON() string {
	return r.raw
}

type AIAIQueryParams struct {
	// Non-plural array property: singularization is a no-op (decisive probe)
	DataToExtract param.Field[[]AiaiQueryParamsDataToExtract] `json:"data_to_extract" api:"required"`
	// The domain name to analyze
	Domain param.Field[string] `json:"domain" api:"required"`
	// Digit-seam acronym parent for nested element mints
	Authentication3DS param.Field[AiaiQueryParamsAuthentication3DS] `json:"authentication_3ds"`
	// Non-plural map property with enum values (map-mint probe)
	DatapointObjectSchema param.Field[map[string]AiaiQueryParamsDatapointObjectSchema] `json:"datapoint_object_schema"`
	// Plural array property: singularized element mint (control probe)
	ExtractionRules param.Field[[]AiaiQueryParamsExtractionRule] `json:"extraction_rules"`
}

func (r AIAIQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AiaiQueryParamsDataToExtract struct {
	DatapointName param.Field[string]                                    `json:"datapoint_name" api:"required"`
	DatapointType param.Field[AiaiQueryParamsDataToExtractDatapointType] `json:"datapoint_type"`
}

func (r AiaiQueryParamsDataToExtract) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AiaiQueryParamsDataToExtractDatapointType string

const (
	AiaiQueryParamsDataToExtractDatapointTypeText   AiaiQueryParamsDataToExtractDatapointType = "text"
	AiaiQueryParamsDataToExtractDatapointTypeNumber AiaiQueryParamsDataToExtractDatapointType = "number"
)

func (r AiaiQueryParamsDataToExtractDatapointType) IsKnown() bool {
	switch r {
	case AiaiQueryParamsDataToExtractDatapointTypeText, AiaiQueryParamsDataToExtractDatapointTypeNumber:
		return true
	}
	return false
}

// Digit-seam acronym parent for nested element mints
type AiaiQueryParamsAuthentication3DS struct {
	// Plural array under the 3DS seam
	ChallengeRules param.Field[[]AiaiQueryParamsAuthentication3DsChallengeRule] `json:"challenge_rules"`
	// Non-plural array under the 3DS seam
	FallbackCheck param.Field[[]AiaiQueryParamsAuthentication3DsFallbackCheck] `json:"fallback_check"`
}

func (r AiaiQueryParamsAuthentication3DS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AiaiQueryParamsAuthentication3DsChallengeRule struct {
	ChallengeKind param.Field[string] `json:"challenge_kind"`
}

func (r AiaiQueryParamsAuthentication3DsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AiaiQueryParamsAuthentication3DsFallbackCheck struct {
	CheckName param.Field[string] `json:"check_name"`
}

func (r AiaiQueryParamsAuthentication3DsFallbackCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AiaiQueryParamsDatapointObjectSchema string

const (
	AiaiQueryParamsDatapointObjectSchemaString  AiaiQueryParamsDatapointObjectSchema = "string"
	AiaiQueryParamsDatapointObjectSchemaNumber  AiaiQueryParamsDatapointObjectSchema = "number"
	AiaiQueryParamsDatapointObjectSchemaBoolean AiaiQueryParamsDatapointObjectSchema = "boolean"
)

func (r AiaiQueryParamsDatapointObjectSchema) IsKnown() bool {
	switch r {
	case AiaiQueryParamsDatapointObjectSchemaString, AiaiQueryParamsDatapointObjectSchemaNumber, AiaiQueryParamsDatapointObjectSchemaBoolean:
		return true
	}
	return false
}

type AiaiQueryParamsExtractionRule struct {
	RuleName     param.Field[string]                                     `json:"rule_name"`
	RuleSeverity param.Field[AiaiQueryParamsExtractionRulesRuleSeverity] `json:"rule_severity"`
}

func (r AiaiQueryParamsExtractionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AiaiQueryParamsExtractionRulesRuleSeverity string

const (
	AiaiQueryParamsExtractionRulesRuleSeverityLow  AiaiQueryParamsExtractionRulesRuleSeverity = "low"
	AiaiQueryParamsExtractionRulesRuleSeverityHigh AiaiQueryParamsExtractionRulesRuleSeverity = "high"
)

func (r AiaiQueryParamsExtractionRulesRuleSeverity) IsKnown() bool {
	switch r {
	case AiaiQueryParamsExtractionRulesRuleSeverityLow, AiaiQueryParamsExtractionRulesRuleSeverityHigh:
		return true
	}
	return false
}
