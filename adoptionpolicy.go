// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/apiquery"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/pagination"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
	"github.com/swagger-girly/go/shared"
)

// Adoption policies and applications
//
// AdoptionPolicyService contains methods and other services that help with
// interacting with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdoptionPolicyService] method instead.
type AdoptionPolicyService struct {
	options []option.RequestOption
}

// NewAdoptionPolicyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdoptionPolicyService(opts ...option.RequestOption) (r AdoptionPolicyService) {
	r = AdoptionPolicyService{}
	r.options = opts
	return
}

// Creates an adoption policy from a set of eligibility rules, regional overrides,
// and an optional escalation rule group.
func (r *AdoptionPolicyService) New(ctx context.Context, body AdoptionPolicyNewParams, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.options, opts)
	path := "adoptions/policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single adoption policy.
func (r *AdoptionPolicyService) Get(ctx context.Context, policyID string, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.options, opts)
	if policyID == "" {
		err = errors.New("missing required policyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("adoptions/policies/%s", url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates either the policy details or its lifecycle state.
func (r *AdoptionPolicyService) Update(ctx context.Context, policyID string, body AdoptionPolicyUpdateParams, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.options, opts)
	if policyID == "" {
		err = errors.New("missing required policyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("adoptions/policies/%s", url.PathEscape(policyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a cursor-paginated list of adoption policies.
func (r *AdoptionPolicyService) List(ctx context.Context, query AdoptionPolicyListParams, opts ...option.RequestOption) (res *pagination.CustomCursorPage[Policy], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "adoptions/policies"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a cursor-paginated list of adoption policies.
func (r *AdoptionPolicyService) ListAutoPaging(ctx context.Context, query AdoptionPolicyListParams, opts ...option.RequestOption) *pagination.CustomCursorPageAutoPager[Policy] {
	return pagination.NewCustomCursorPageAutoPager(r.List(ctx, query, opts...))
}

type AdoptionRule struct {
	// The applicant or pet attribute the rule targets.
	//
	// Any of "PET_STATUS", "SPECIES", "ADOPTER_AGE", "REGION", "PRIOR_ADOPTIONS",
	// "APPLICATION_DATE".
	Attribute RuleAttribute `json:"attribute" api:"required"`
	// Any of "IS_ONE_OF", "IS_NOT_ONE_OF", "MATCHES", "IS_GREATER_THAN",
	// "IS_LESS_THAN", "IS_BEFORE", "IS_AFTER".
	Operation AdoptionRuleOperation `json:"operation" api:"required"`
	// A regex string, to be used with `MATCHES`
	Value ConditionValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attribute   respjson.Field
		Operation   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdoptionRule) RawJSON() string { return r.JSON.raw }
func (r *AdoptionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AdoptionRule to a AdoptionRuleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AdoptionRuleParam.Overrides()
func (r AdoptionRule) ToParam() AdoptionRuleParam {
	return param.Override[AdoptionRuleParam](json.RawMessage(r.RawJSON()))
}

type AdoptionRuleOperation string

const (
	AdoptionRuleOperationIsOneOf       AdoptionRuleOperation = "IS_ONE_OF"
	AdoptionRuleOperationIsNotOneOf    AdoptionRuleOperation = "IS_NOT_ONE_OF"
	AdoptionRuleOperationMatches       AdoptionRuleOperation = "MATCHES"
	AdoptionRuleOperationIsGreaterThan AdoptionRuleOperation = "IS_GREATER_THAN"
	AdoptionRuleOperationIsLessThan    AdoptionRuleOperation = "IS_LESS_THAN"
	AdoptionRuleOperationIsBefore      AdoptionRuleOperation = "IS_BEFORE"
	AdoptionRuleOperationIsAfter       AdoptionRuleOperation = "IS_AFTER"
)

// The properties Attribute, Operation are required.
type AdoptionRuleParam struct {
	// The applicant or pet attribute the rule targets.
	//
	// Any of "PET_STATUS", "SPECIES", "ADOPTER_AGE", "REGION", "PRIOR_ADOPTIONS",
	// "APPLICATION_DATE".
	Attribute RuleAttribute `json:"attribute,omitzero" api:"required"`
	// Any of "IS_ONE_OF", "IS_NOT_ONE_OF", "MATCHES", "IS_GREATER_THAN",
	// "IS_LESS_THAN", "IS_BEFORE", "IS_AFTER".
	Operation AdoptionRuleOperation `json:"operation,omitzero" api:"required"`
	// A regex string, to be used with `MATCHES`
	Value ConditionValueUnionParam `json:"value,omitzero"`
	paramObj
}

func (r AdoptionRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConditionValueUnion contains all possible properties and values from [string],
// [int64], [float64], [[]string], [time.Time].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfFloat OfListOfStrings OfTime]
type ConditionValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfListOfStrings []string `json:",inline"`
	// This field will be present if the value is a [time.Time] instead of an object.
	OfTime time.Time `json:",inline"`
	JSON   struct {
		OfString        respjson.Field
		OfInt           respjson.Field
		OfFloat         respjson.Field
		OfListOfStrings respjson.Field
		OfTime          respjson.Field
		raw             string
	} `json:"-"`
}

func (u ConditionValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConditionValueUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConditionValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConditionValueUnion) AsListOfStrings() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConditionValueUnion) AsTime() (v time.Time) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConditionValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ConditionValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConditionValueUnion to a ConditionValueUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConditionValueUnionParam.Overrides()
func (r ConditionValueUnion) ToParam() ConditionValueUnionParam {
	return param.Override[ConditionValueUnionParam](json.RawMessage(r.RawJSON()))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConditionValueUnionParam struct {
	OfString        param.Opt[string]    `json:",omitzero,inline"`
	OfInt           param.Opt[int64]     `json:",omitzero,inline"`
	OfFloat         param.Opt[float64]   `json:",omitzero,inline"`
	OfListOfStrings []string             `json:",omitzero,inline"`
	OfTime          param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u ConditionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfInt,
		u.OfFloat,
		u.OfListOfStrings,
		u.OfTime)
}
func (u *ConditionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type Policy struct {
	ID        string         `json:"id" api:"required"`
	CreatedAt time.Time      `json:"createdAt" api:"required" format:"date-time"`
	Name      string         `json:"name" api:"required"`
	Rules     []AdoptionRule `json:"rules" api:"required"`
	// Any of "draft", "active", "archived".
	Status     PolicyStatus     `json:"status" api:"required"`
	Escalation PolicyEscalation `json:"escalation"`
	FeeCap     shared.Money     `json:"feeCap"`
	// Where policy decisions are delivered: a plain email address or a webhook sink.
	NotificationTarget PolicyNotificationTargetUnion `json:"notificationTarget"`
	// Per-region override values keyed by region code.
	RegionalOverrides map[string]ConditionValueUnion `json:"regionalOverrides"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Name               respjson.Field
		Rules              respjson.Field
		Status             respjson.Field
		Escalation         respjson.Field
		FeeCap             respjson.Field
		NotificationTarget respjson.Field
		RegionalOverrides  respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Policy) RawJSON() string { return r.JSON.raw }
func (r *Policy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyStatus string

const (
	PolicyStatusDraft    PolicyStatus = "draft"
	PolicyStatusActive   PolicyStatus = "active"
	PolicyStatusArchived PolicyStatus = "archived"
)

type PolicyEscalation struct {
	// Any of "ALL", "ANY".
	Operator string `json:"operator" api:"required"`
	// Nested rule groups; groups can recurse arbitrarily deep.
	Groups []any          `json:"groups"`
	Rules  []AdoptionRule `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Groups      respjson.Field
		Rules       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyEscalation) RawJSON() string { return r.JSON.raw }
func (r *PolicyEscalation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyNotificationTargetUnion contains all possible properties and values from
// [string], [PolicyNotificationTargetWebhookSink].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type PolicyNotificationTargetUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [PolicyNotificationTargetWebhookSink].
	URL string `json:"url"`
	// This field is from variant [PolicyNotificationTargetWebhookSink].
	Events []string `json:"events"`
	// This field is from variant [PolicyNotificationTargetWebhookSink].
	Secret string `json:"secret"`
	JSON   struct {
		OfString respjson.Field
		URL      respjson.Field
		Events   respjson.Field
		Secret   respjson.Field
		raw      string
	} `json:"-"`
}

func (u PolicyNotificationTargetUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolicyNotificationTargetUnion) AsWebhookSink() (v PolicyNotificationTargetWebhookSink) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolicyNotificationTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *PolicyNotificationTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicyNotificationTargetWebhookSink struct {
	URL string `json:"url" api:"required" format:"uri"`
	// Any of "approved", "rejected", "escalated".
	Events []string `json:"events"`
	Secret string   `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Events      respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyNotificationTargetWebhookSink) RawJSON() string { return r.JSON.raw }
func (r *PolicyNotificationTargetWebhookSink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The applicant or pet attribute the rule targets.
type RuleAttribute string

const (
	RuleAttributePetStatus       RuleAttribute = "PET_STATUS"
	RuleAttributeSpecies         RuleAttribute = "SPECIES"
	RuleAttributeAdopterAge      RuleAttribute = "ADOPTER_AGE"
	RuleAttributeRegion          RuleAttribute = "REGION"
	RuleAttributePriorAdoptions  RuleAttribute = "PRIOR_ADOPTIONS"
	RuleAttributeApplicationDate RuleAttribute = "APPLICATION_DATE"
)

type AdoptionPolicyNewParams struct {
	Name       string                            `json:"name" api:"required"`
	Rules      []AdoptionRuleParam               `json:"rules,omitzero" api:"required"`
	Escalation AdoptionPolicyNewParamsEscalation `json:"escalation,omitzero"`
	FeeCap     shared.MoneyParam                 `json:"feeCap,omitzero"`
	// Where policy decisions are delivered: a plain email address or a webhook sink.
	NotificationTarget AdoptionPolicyNewParamsNotificationTargetUnion `json:"notificationTarget,omitzero"`
	RegionalOverrides  map[string]ConditionValueUnionParam            `json:"regionalOverrides,omitzero"`
	paramObj
}

func (r AdoptionPolicyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionPolicyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionPolicyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Operator is required.
type AdoptionPolicyNewParamsEscalation struct {
	// Any of "ALL", "ANY".
	Operator string `json:"operator,omitzero" api:"required"`
	// Nested rule groups; groups can recurse arbitrarily deep.
	Groups []any               `json:"groups,omitzero"`
	Rules  []AdoptionRuleParam `json:"rules,omitzero"`
	paramObj
}

func (r AdoptionPolicyNewParamsEscalation) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionPolicyNewParamsEscalation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionPolicyNewParamsEscalation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AdoptionPolicyNewParamsEscalation](
		"operator", "ALL", "ANY",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AdoptionPolicyNewParamsNotificationTargetUnion struct {
	OfString      param.Opt[string]                                     `json:",omitzero,inline"`
	OfWebhookSink *AdoptionPolicyNewParamsNotificationTargetWebhookSink `json:",omitzero,inline"`
	paramUnion
}

func (u AdoptionPolicyNewParamsNotificationTargetUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfWebhookSink)
}
func (u *AdoptionPolicyNewParamsNotificationTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property URL is required.
type AdoptionPolicyNewParamsNotificationTargetWebhookSink struct {
	URL    string            `json:"url" api:"required" format:"uri"`
	Secret param.Opt[string] `json:"secret,omitzero"`
	// Any of "approved", "rejected", "escalated".
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r AdoptionPolicyNewParamsNotificationTargetWebhookSink) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionPolicyNewParamsNotificationTargetWebhookSink
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionPolicyNewParamsNotificationTargetWebhookSink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionPolicyUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfAdoptionPolicyDetailsUpdate *AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdate `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAdoptionPolicyLifecycleUpdate *AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdate `json:",inline"`

	paramObj
}

func (u AdoptionPolicyUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAdoptionPolicyDetailsUpdate, u.OfAdoptionPolicyLifecycleUpdate)
}
func (r *AdoptionPolicyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdate struct {
	Name       param.Opt[string]                                                   `json:"name,omitzero"`
	Escalation AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalation `json:"escalation,omitzero"`
	// Where policy decisions are delivered: a plain email address or a webhook sink.
	NotificationTarget AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetUnion `json:"notificationTarget,omitzero"`
	paramObj
}

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdate) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Operator is required.
type AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalation struct {
	// Any of "ALL", "ANY".
	Operator string `json:"operator,omitzero" api:"required"`
	// Nested rule groups; groups can recurse arbitrarily deep.
	Groups []any               `json:"groups,omitzero"`
	Rules  []AdoptionRuleParam `json:"rules,omitzero"`
	paramObj
}

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalation) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalation](
		"operator", "ALL", "ANY",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetUnion struct {
	OfString      param.Opt[string]                                                                       `json:",omitzero,inline"`
	OfWebhookSink *AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSink `json:",omitzero,inline"`
	paramUnion
}

func (u AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfWebhookSink)
}
func (u *AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property URL is required.
type AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSink struct {
	URL    string            `json:"url" api:"required" format:"uri"`
	Secret param.Opt[string] `json:"secret,omitzero"`
	// Any of "approved", "rejected", "escalated".
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSink) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSink
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Actor, Status are required.
type AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdate struct {
	Actor string `json:"actor" api:"required"`
	// Any of "active", "archived".
	Status  string               `json:"status,omitzero" api:"required"`
	ActedAt param.Opt[time.Time] `json:"actedAt,omitzero" format:"date-time"`
	Reason  param.Opt[string]    `json:"reason,omitzero"`
	paramObj
}

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdate) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdate](
		"status", "active", "archived",
	)
}

type AdoptionPolicyListParams struct {
	// Cursor from a previous response used to fetch the next page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of policies to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AdoptionPolicyListParams]'s query parameters as
// `url.Values`.
func (r AdoptionPolicyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
