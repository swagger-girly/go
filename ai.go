// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"net/http"
	"slices"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
)

// AIService contains methods and other services that help with interacting with
// the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIService] method instead.
type AIService struct {
	options []option.RequestOption
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r AIService) {
	r = AIService{}
	r.options = opts
	return
}

// Casing probe replicating context.dev: glued AIAI prefix (resource ai + method
// ai_query) over plural and non-plural synthesized element mints
func (r *AIService) AIQuery(ctx context.Context, body AIAIQueryParams, opts ...option.RequestOption) (res *AiaiQueryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "ai/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AiaiQueryResponse struct {
	// Non-plural array on the response side
	DataExtracted []AiaiQueryResponseDataExtracted `json:"data_extracted"`
	// Plural array on the response side
	MatchedRules []AiaiQueryResponseMatchedRule `json:"matched_rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataExtracted respjson.Field
		MatchedRules  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AiaiQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *AiaiQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AiaiQueryResponseDataExtracted struct {
	DatapointValue string `json:"datapoint_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DatapointValue respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AiaiQueryResponseDataExtracted) RawJSON() string { return r.JSON.raw }
func (r *AiaiQueryResponseDataExtracted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AiaiQueryResponseMatchedRule struct {
	MatchScore float64 `json:"match_score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MatchScore  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AiaiQueryResponseMatchedRule) RawJSON() string { return r.JSON.raw }
func (r *AiaiQueryResponseMatchedRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAIQueryParams struct {
	// Non-plural array property: singularization is a no-op (decisive probe)
	DataToExtract []AIAIQueryParamsDataToExtract `json:"data_to_extract,omitzero" api:"required"`
	// The domain name to analyze
	Domain string `json:"domain" api:"required"`
	// Digit-seam acronym parent for nested element mints
	Authentication3DS AIAIQueryParamsAuthentication3DS `json:"authentication_3ds,omitzero"`
	// Non-plural map property with enum values (map-mint probe)
	//
	// Any of "string", "number", "boolean".
	DatapointObjectSchema map[string]string `json:"datapoint_object_schema,omitzero"`
	// Plural array property: singularized element mint (control probe)
	ExtractionRules []AIAIQueryParamsExtractionRule `json:"extraction_rules,omitzero"`
	paramObj
}

func (r AIAIQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAIQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAIQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property DatapointName is required.
type AIAIQueryParamsDataToExtract struct {
	DatapointName string `json:"datapoint_name" api:"required"`
	// Any of "text", "number".
	DatapointType string `json:"datapoint_type,omitzero"`
	paramObj
}

func (r AIAIQueryParamsDataToExtract) MarshalJSON() (data []byte, err error) {
	type shadow AIAIQueryParamsDataToExtract
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAIQueryParamsDataToExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIAIQueryParamsDataToExtract](
		"datapoint_type", "text", "number",
	)
}

// Digit-seam acronym parent for nested element mints
type AIAIQueryParamsAuthentication3DS struct {
	// Plural array under the 3DS seam
	ChallengeRules []AIAIQueryParamsAuthentication3DSChallengeRule `json:"challenge_rules,omitzero"`
	// Non-plural array under the 3DS seam
	FallbackCheck []AIAIQueryParamsAuthentication3DSFallbackCheck `json:"fallback_check,omitzero"`
	paramObj
}

func (r AIAIQueryParamsAuthentication3DS) MarshalJSON() (data []byte, err error) {
	type shadow AIAIQueryParamsAuthentication3DS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAIQueryParamsAuthentication3DS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAIQueryParamsAuthentication3DSChallengeRule struct {
	ChallengeKind param.Opt[string] `json:"challenge_kind,omitzero"`
	paramObj
}

func (r AIAIQueryParamsAuthentication3DSChallengeRule) MarshalJSON() (data []byte, err error) {
	type shadow AIAIQueryParamsAuthentication3DSChallengeRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAIQueryParamsAuthentication3DSChallengeRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAIQueryParamsAuthentication3DSFallbackCheck struct {
	CheckName param.Opt[string] `json:"check_name,omitzero"`
	paramObj
}

func (r AIAIQueryParamsAuthentication3DSFallbackCheck) MarshalJSON() (data []byte, err error) {
	type shadow AIAIQueryParamsAuthentication3DSFallbackCheck
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAIQueryParamsAuthentication3DSFallbackCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAIQueryParamsExtractionRule struct {
	RuleName param.Opt[string] `json:"rule_name,omitzero"`
	// Any of "low", "high".
	RuleSeverity string `json:"rule_severity,omitzero"`
	paramObj
}

func (r AIAIQueryParamsExtractionRule) MarshalJSON() (data []byte, err error) {
	type shadow AIAIQueryParamsExtractionRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAIQueryParamsExtractionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIAIQueryParamsExtractionRule](
		"rule_severity", "low", "high",
	)
}
