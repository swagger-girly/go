// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/apiquery"
	"github.com/swagger-girly/go/internal/param"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/pagination"
	"github.com/swagger-girly/go/shared"
	"github.com/tidwall/gjson"
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
	Options []option.RequestOption
}

// NewAdoptionPolicyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdoptionPolicyService(opts ...option.RequestOption) (r *AdoptionPolicyService) {
	r = &AdoptionPolicyService{}
	r.Options = opts
	return
}

// Creates an adoption policy from a set of eligibility rules, regional overrides,
// and an optional escalation rule group.
func (r *AdoptionPolicyService) New(ctx context.Context, body AdoptionPolicyNewParams, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "adoptions/policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single adoption policy.
func (r *AdoptionPolicyService) Get(ctx context.Context, policyID string, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("adoptions/policies/%s", policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates either the policy details or its lifecycle state.
func (r *AdoptionPolicyService) Update(ctx context.Context, policyID string, body AdoptionPolicyUpdateParams, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("adoptions/policies/%s", policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a cursor-paginated list of adoption policies.
func (r *AdoptionPolicyService) List(ctx context.Context, query AdoptionPolicyListParams, opts ...option.RequestOption) (res *pagination.CustomCursorPage[Policy], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	Attribute RuleAttribute         `json:"attribute" api:"required"`
	Operation AdoptionRuleOperation `json:"operation" api:"required"`
	// A regex string, to be used with `MATCHES`
	Value ConditionValueUnion `json:"value"`
	JSON  adoptionRuleJSON    `json:"-"`
}

// adoptionRuleJSON contains the JSON metadata for the struct [AdoptionRule]
type adoptionRuleJSON struct {
	Attribute   apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdoptionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adoptionRuleJSON) RawJSON() string {
	return r.raw
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

func (r AdoptionRuleOperation) IsKnown() bool {
	switch r {
	case AdoptionRuleOperationIsOneOf, AdoptionRuleOperationIsNotOneOf, AdoptionRuleOperationMatches, AdoptionRuleOperationIsGreaterThan, AdoptionRuleOperationIsLessThan, AdoptionRuleOperationIsBefore, AdoptionRuleOperationIsAfter:
		return true
	}
	return false
}

type AdoptionRuleParam struct {
	// The applicant or pet attribute the rule targets.
	Attribute param.Field[RuleAttribute]         `json:"attribute" api:"required"`
	Operation param.Field[AdoptionRuleOperation] `json:"operation" api:"required"`
	// A regex string, to be used with `MATCHES`
	Value param.Field[ConditionValueUnionParam] `json:"value"`
}

func (r AdoptionRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A regex string, to be used with `MATCHES`
//
// Union satisfied by [shared.UnionString], [shared.UnionInt], [shared.UnionFloat],
// [ConditionValueListOfStrings] or [shared.UnionTime].
type ConditionValueUnion interface {
	ImplementsConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConditionValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConditionValueListOfStrings{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
	)
}

type ConditionValueListOfStrings []string

func (r ConditionValueListOfStrings) ImplementsConditionValueUnion() {}

// A regex string, to be used with `MATCHES`
//
// Satisfied by [shared.UnionString], [shared.UnionInt], [shared.UnionFloat],
// [ConditionValueListOfStringsParam], [shared.UnionTime].
type ConditionValueUnionParam interface {
	ImplementsConditionValueUnionParam()
}

type ConditionValueListOfStringsParam []string

func (r ConditionValueListOfStringsParam) ImplementsConditionValueUnionParam() {}

type Policy struct {
	ID         string           `json:"id" api:"required"`
	CreatedAt  time.Time        `json:"createdAt" api:"required" format:"date-time"`
	Name       string           `json:"name" api:"required"`
	Rules      []AdoptionRule   `json:"rules" api:"required"`
	Status     PolicyStatus     `json:"status" api:"required"`
	Escalation PolicyEscalation `json:"escalation"`
	FeeCap     shared.Money     `json:"feeCap"`
	// Where policy decisions are delivered: a plain email address or a webhook sink.
	NotificationTarget PolicyNotificationTargetUnion `json:"notificationTarget"`
	// Per-region override values keyed by region code.
	RegionalOverrides map[string]ConditionValueUnion `json:"regionalOverrides"`
	JSON              policyJSON                     `json:"-"`
}

// policyJSON contains the JSON metadata for the struct [Policy]
type policyJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Name               apijson.Field
	Rules              apijson.Field
	Status             apijson.Field
	Escalation         apijson.Field
	FeeCap             apijson.Field
	NotificationTarget apijson.Field
	RegionalOverrides  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Policy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyJSON) RawJSON() string {
	return r.raw
}

type PolicyStatus string

const (
	PolicyStatusDraft    PolicyStatus = "draft"
	PolicyStatusActive   PolicyStatus = "active"
	PolicyStatusArchived PolicyStatus = "archived"
)

func (r PolicyStatus) IsKnown() bool {
	switch r {
	case PolicyStatusDraft, PolicyStatusActive, PolicyStatusArchived:
		return true
	}
	return false
}

type PolicyEscalation struct {
	Operator PolicyEscalationOperator `json:"operator" api:"required"`
	// Nested rule groups; groups can recurse arbitrarily deep.
	Groups []interface{}        `json:"groups"`
	Rules  []AdoptionRule       `json:"rules"`
	JSON   policyEscalationJSON `json:"-"`
}

// policyEscalationJSON contains the JSON metadata for the struct
// [PolicyEscalation]
type policyEscalationJSON struct {
	Operator    apijson.Field
	Groups      apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyEscalation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyEscalationJSON) RawJSON() string {
	return r.raw
}

type PolicyEscalationOperator string

const (
	PolicyEscalationOperatorAll PolicyEscalationOperator = "ALL"
	PolicyEscalationOperatorAny PolicyEscalationOperator = "ANY"
)

func (r PolicyEscalationOperator) IsKnown() bool {
	switch r {
	case PolicyEscalationOperatorAll, PolicyEscalationOperatorAny:
		return true
	}
	return false
}

// Where policy decisions are delivered: a plain email address or a webhook sink.
//
// Union satisfied by [shared.UnionString] or
// [PolicyNotificationTargetWebhookSink].
type PolicyNotificationTargetUnion interface {
	ImplementsPolicyNotificationTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PolicyNotificationTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PolicyNotificationTargetWebhookSink{}),
		},
	)
}

type PolicyNotificationTargetWebhookSink struct {
	URL    string                                     `json:"url" api:"required" format:"uri"`
	Events []PolicyNotificationTargetWebhookSinkEvent `json:"events"`
	JSON   policyNotificationTargetWebhookSinkJSON    `json:"-"`
}

// policyNotificationTargetWebhookSinkJSON contains the JSON metadata for the
// struct [PolicyNotificationTargetWebhookSink]
type policyNotificationTargetWebhookSinkJSON struct {
	URL         apijson.Field
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyNotificationTargetWebhookSink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyNotificationTargetWebhookSinkJSON) RawJSON() string {
	return r.raw
}

func (r PolicyNotificationTargetWebhookSink) ImplementsPolicyNotificationTargetUnion() {}

type PolicyNotificationTargetWebhookSinkEvent string

const (
	PolicyNotificationTargetWebhookSinkEventApproved  PolicyNotificationTargetWebhookSinkEvent = "approved"
	PolicyNotificationTargetWebhookSinkEventRejected  PolicyNotificationTargetWebhookSinkEvent = "rejected"
	PolicyNotificationTargetWebhookSinkEventEscalated PolicyNotificationTargetWebhookSinkEvent = "escalated"
)

func (r PolicyNotificationTargetWebhookSinkEvent) IsKnown() bool {
	switch r {
	case PolicyNotificationTargetWebhookSinkEventApproved, PolicyNotificationTargetWebhookSinkEventRejected, PolicyNotificationTargetWebhookSinkEventEscalated:
		return true
	}
	return false
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

func (r RuleAttribute) IsKnown() bool {
	switch r {
	case RuleAttributePetStatus, RuleAttributeSpecies, RuleAttributeAdopterAge, RuleAttributeRegion, RuleAttributePriorAdoptions, RuleAttributeApplicationDate:
		return true
	}
	return false
}

type AdoptionPolicyNewParams struct {
	Name       param.Field[string]                            `json:"name" api:"required"`
	Rules      param.Field[[]AdoptionRuleParam]               `json:"rules" api:"required"`
	Escalation param.Field[AdoptionPolicyNewParamsEscalation] `json:"escalation"`
	FeeCap     param.Field[shared.MoneyParam]                 `json:"feeCap"`
	// Where policy decisions are delivered: a plain email address or a webhook sink.
	NotificationTarget param.Field[AdoptionPolicyNewParamsNotificationTargetUnion] `json:"notificationTarget"`
	RegionalOverrides  param.Field[map[string]ConditionValueUnionParam]            `json:"regionalOverrides"`
}

func (r AdoptionPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdoptionPolicyNewParamsEscalation struct {
	Operator param.Field[AdoptionPolicyNewParamsEscalationOperator] `json:"operator" api:"required"`
	// Nested rule groups; groups can recurse arbitrarily deep.
	Groups param.Field[[]interface{}]       `json:"groups"`
	Rules  param.Field[[]AdoptionRuleParam] `json:"rules"`
}

func (r AdoptionPolicyNewParamsEscalation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdoptionPolicyNewParamsEscalationOperator string

const (
	AdoptionPolicyNewParamsEscalationOperatorAll AdoptionPolicyNewParamsEscalationOperator = "ALL"
	AdoptionPolicyNewParamsEscalationOperatorAny AdoptionPolicyNewParamsEscalationOperator = "ANY"
)

func (r AdoptionPolicyNewParamsEscalationOperator) IsKnown() bool {
	switch r {
	case AdoptionPolicyNewParamsEscalationOperatorAll, AdoptionPolicyNewParamsEscalationOperatorAny:
		return true
	}
	return false
}

// Where policy decisions are delivered: a plain email address or a webhook sink.
//
// Satisfied by [shared.UnionString],
// [AdoptionPolicyNewParamsNotificationTargetWebhookSink].
type AdoptionPolicyNewParamsNotificationTargetUnion interface {
	ImplementsAdoptionPolicyNewParamsNotificationTargetUnion()
}

type AdoptionPolicyNewParamsNotificationTargetWebhookSink struct {
	URL    param.Field[string]                                                      `json:"url" api:"required" format:"uri"`
	Events param.Field[[]AdoptionPolicyNewParamsNotificationTargetWebhookSinkEvent] `json:"events"`
	Secret param.Field[string]                                                      `json:"secret"`
}

func (r AdoptionPolicyNewParamsNotificationTargetWebhookSink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdoptionPolicyNewParamsNotificationTargetWebhookSink) ImplementsAdoptionPolicyNewParamsNotificationTargetUnion() {
}

type AdoptionPolicyNewParamsNotificationTargetWebhookSinkEvent string

const (
	AdoptionPolicyNewParamsNotificationTargetWebhookSinkEventApproved  AdoptionPolicyNewParamsNotificationTargetWebhookSinkEvent = "approved"
	AdoptionPolicyNewParamsNotificationTargetWebhookSinkEventRejected  AdoptionPolicyNewParamsNotificationTargetWebhookSinkEvent = "rejected"
	AdoptionPolicyNewParamsNotificationTargetWebhookSinkEventEscalated AdoptionPolicyNewParamsNotificationTargetWebhookSinkEvent = "escalated"
)

func (r AdoptionPolicyNewParamsNotificationTargetWebhookSinkEvent) IsKnown() bool {
	switch r {
	case AdoptionPolicyNewParamsNotificationTargetWebhookSinkEventApproved, AdoptionPolicyNewParamsNotificationTargetWebhookSinkEventRejected, AdoptionPolicyNewParamsNotificationTargetWebhookSinkEventEscalated:
		return true
	}
	return false
}

type AdoptionPolicyUpdateParams struct {
	Body AdoptionPolicyUpdateParamsBodyUnion `json:"body" api:"required"`
}

func (r AdoptionPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AdoptionPolicyUpdateParamsBody struct {
	ActedAt            param.Field[time.Time]                            `json:"actedAt" format:"date-time"`
	Actor              param.Field[string]                               `json:"actor"`
	Escalation         param.Field[interface{}]                          `json:"escalation"`
	Name               param.Field[string]                               `json:"name"`
	NotificationTarget param.Field[interface{}]                          `json:"notificationTarget"`
	Reason             param.Field[string]                               `json:"reason"`
	Status             param.Field[AdoptionPolicyUpdateParamsBodyStatus] `json:"status"`
}

func (r AdoptionPolicyUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdoptionPolicyUpdateParamsBody) implementsAdoptionPolicyUpdateParamsBodyUnion() {}

// Satisfied by [AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdate],
// [AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdate],
// [AdoptionPolicyUpdateParamsBody].
type AdoptionPolicyUpdateParamsBodyUnion interface {
	implementsAdoptionPolicyUpdateParamsBodyUnion()
}

type AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdate struct {
	Escalation param.Field[AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalation] `json:"escalation"`
	Name       param.Field[string]                                                              `json:"name"`
	// Where policy decisions are delivered: a plain email address or a webhook sink.
	NotificationTarget param.Field[AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetUnion] `json:"notificationTarget"`
}

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdate) implementsAdoptionPolicyUpdateParamsBodyUnion() {
}

type AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalation struct {
	Operator param.Field[AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalationOperator] `json:"operator" api:"required"`
	// Nested rule groups; groups can recurse arbitrarily deep.
	Groups param.Field[[]interface{}]       `json:"groups"`
	Rules  param.Field[[]AdoptionRuleParam] `json:"rules"`
}

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalationOperator string

const (
	AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalationOperatorAll AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalationOperator = "ALL"
	AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalationOperatorAny AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalationOperator = "ANY"
)

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalationOperator) IsKnown() bool {
	switch r {
	case AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalationOperatorAll, AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalationOperatorAny:
		return true
	}
	return false
}

// Where policy decisions are delivered: a plain email address or a webhook sink.
//
// Satisfied by [shared.UnionString],
// [AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSink].
type AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetUnion interface {
	ImplementsAdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetUnion()
}

type AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSink struct {
	URL    param.Field[string]                                                                                        `json:"url" api:"required" format:"uri"`
	Events param.Field[[]AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEvent] `json:"events"`
	Secret param.Field[string]                                                                                        `json:"secret"`
}

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSink) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSink) ImplementsAdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetUnion() {
}

type AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEvent string

const (
	AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEventApproved  AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEvent = "approved"
	AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEventRejected  AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEvent = "rejected"
	AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEventEscalated AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEvent = "escalated"
)

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEvent) IsKnown() bool {
	switch r {
	case AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEventApproved, AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEventRejected, AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetWebhookSinkEventEscalated:
		return true
	}
	return false
}

type AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdate struct {
	Actor   param.Field[string]                                                            `json:"actor" api:"required"`
	Status  param.Field[AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdateStatus] `json:"status" api:"required"`
	ActedAt param.Field[time.Time]                                                         `json:"actedAt" format:"date-time"`
	Reason  param.Field[string]                                                            `json:"reason"`
}

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdate) implementsAdoptionPolicyUpdateParamsBodyUnion() {
}

type AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdateStatus string

const (
	AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdateStatusActive   AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdateStatus = "active"
	AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdateStatusArchived AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdateStatus = "archived"
)

func (r AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdateStatus) IsKnown() bool {
	switch r {
	case AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdateStatusActive, AdoptionPolicyUpdateParamsBodyAdoptionPolicyLifecycleUpdateStatusArchived:
		return true
	}
	return false
}

type AdoptionPolicyUpdateParamsBodyStatus string

const (
	AdoptionPolicyUpdateParamsBodyStatusActive   AdoptionPolicyUpdateParamsBodyStatus = "active"
	AdoptionPolicyUpdateParamsBodyStatusArchived AdoptionPolicyUpdateParamsBodyStatus = "archived"
)

func (r AdoptionPolicyUpdateParamsBodyStatus) IsKnown() bool {
	switch r {
	case AdoptionPolicyUpdateParamsBodyStatusActive, AdoptionPolicyUpdateParamsBodyStatusArchived:
		return true
	}
	return false
}

type AdoptionPolicyListParams struct {
	// Cursor from a previous response used to fetch the next page.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of policies to return.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [AdoptionPolicyListParams]'s query parameters as
// `url.Values`.
func (r AdoptionPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
