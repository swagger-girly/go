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
	shimjson "github.com/swagger-girly/go/internal/encoding/json"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/pagination"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
	"github.com/swagger-girly/go/shared"
	"github.com/swagger-girly/go/shared/constant"
)

// Post-adoption placement tracking
//
// PlacementService contains methods and other services that help with interacting
// with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlacementService] method instead.
type PlacementService struct {
	options []option.RequestOption
}

// NewPlacementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlacementService(opts ...option.RequestOption) (r PlacementService) {
	r = PlacementService{}
	r.options = opts
	return
}

// Creates a placement for an approved adoption application.
func (r *PlacementService) New(ctx context.Context, body PlacementNewParams, opts ...option.RequestOption) (res *Placement, err error) {
	opts = slices.Concat(r.options, opts)
	path := "placements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single placement, including its adoption application and event
// history.
func (r *PlacementService) Get(ctx context.Context, placementID string, opts ...option.RequestOption) (res *Placement, err error) {
	opts = slices.Concat(r.options, opts)
	if placementID == "" {
		err = errors.New("missing required placementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("placements/%s", url.PathEscape(placementID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a cursor-paginated list of placements.
func (r *PlacementService) List(ctx context.Context, query PlacementListParams, opts ...option.RequestOption) (res *pagination.CustomCursorPage[Placement], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "placements"
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

// Returns a cursor-paginated list of placements.
func (r *PlacementService) ListAutoPaging(ctx context.Context, query PlacementListParams, opts ...option.RequestOption) *pagination.CustomCursorPageAutoPager[Placement] {
	return pagination.NewCustomCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Appends a transfer, checkup, or disruption event to a placement's history.
func (r *PlacementService) RecordEvent(ctx context.Context, placementID string, body PlacementRecordEventParams, opts ...option.RequestOption) (res *Placement, err error) {
	opts = slices.Concat(r.options, opts)
	if placementID == "" {
		err = errors.New("missing required placementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("placements/%s/events", url.PathEscape(placementID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A post-adoption placement tracking the pet's transition to its new home.
type Placement struct {
	ID          string      `json:"id" api:"required"`
	Application Application `json:"application" api:"required"`
	CreatedAt   time.Time   `json:"createdAt" api:"required" format:"date-time"`
	// Any of "pending", "active", "completed", "disrupted".
	Status PlacementStatus `json:"status" api:"required"`
	// Unified activity feed mixing event, milestone, and note entries.
	Activity      []PlacementActivityUnion `json:"activity"`
	Events        []PlacementEventUnion    `json:"events"`
	FollowUpAfter time.Time                `json:"followUpAfter" api:"nullable" format:"date-time"`
	// Transport plan for a placement; pickup and delivery share the transfer-leg
	// shape.
	Logistics        PlacementLogistics `json:"logistics"`
	MedicalClearance MedicalSummary     `json:"medicalClearance"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Application      respjson.Field
		CreatedAt        respjson.Field
		Status           respjson.Field
		Activity         respjson.Field
		Events           respjson.Field
		FollowUpAfter    respjson.Field
		Logistics        respjson.Field
		MedicalClearance respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Placement) RawJSON() string { return r.JSON.raw }
func (r *Placement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementStatus string

const (
	PlacementStatusPending   PlacementStatus = "pending"
	PlacementStatusActive    PlacementStatus = "active"
	PlacementStatusCompleted PlacementStatus = "completed"
	PlacementStatusDisrupted PlacementStatus = "disrupted"
)

// PlacementActivityUnion contains all possible properties and values from
// [PlacementActivityPlacementTransferEvent],
// [PlacementActivityPlacementCheckupEvent],
// [PlacementActivityPlacementDisruptionEvent], [PlacementActivityNote],
// [PlacementActivityAdopted], [PlacementActivityReturned].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PlacementActivityUnion struct {
	ID string `json:"id"`
	// This field is from variant [PlacementActivityPlacementTransferEvent].
	Leg        TransferLeg `json:"leg"`
	OccurredAt time.Time   `json:"occurredAt"`
	Type       string      `json:"type"`
	Note       string      `json:"note"`
	// This field is from variant [PlacementActivityPlacementCheckupEvent].
	FollowUp PlacementActivityPlacementCheckupEventFollowUp `json:"followUp"`
	// This field is from variant [PlacementActivityPlacementCheckupEvent].
	Record VaccinationRecord `json:"record"`
	// This field is from variant [PlacementActivityPlacementDisruptionEvent].
	Severity PlacementActivityPlacementDisruptionEventSeverityUnion `json:"severity"`
	// This field is from variant [PlacementActivityNote].
	Body string `json:"body"`
	Kind string `json:"kind"`
	// This field is from variant [PlacementActivityAdopted].
	AdoptedAt time.Time `json:"adoptedAt"`
	// This field is from variant [PlacementActivityReturned].
	Reason string `json:"reason"`
	JSON   struct {
		ID         respjson.Field
		Leg        respjson.Field
		OccurredAt respjson.Field
		Type       respjson.Field
		Note       respjson.Field
		FollowUp   respjson.Field
		Record     respjson.Field
		Severity   respjson.Field
		Body       respjson.Field
		Kind       respjson.Field
		AdoptedAt  respjson.Field
		Reason     respjson.Field
		raw        string
	} `json:"-"`
}

func (u PlacementActivityUnion) AsPlacementActivityPlacementTransferEvent() (v PlacementActivityPlacementTransferEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlacementActivityUnion) AsPlacementActivityPlacementCheckupEvent() (v PlacementActivityPlacementCheckupEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlacementActivityUnion) AsPlacementActivityPlacementDisruptionEvent() (v PlacementActivityPlacementDisruptionEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlacementActivityUnion) AsNote() (v PlacementActivityNote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlacementActivityUnion) AsAdopted() (v PlacementActivityAdopted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlacementActivityUnion) AsReturned() (v PlacementActivityReturned) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlacementActivityUnion) RawJSON() string { return u.JSON.raw }

func (r *PlacementActivityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementActivityPlacementTransferEvent struct {
	ID         string      `json:"id" api:"required"`
	Leg        TransferLeg `json:"leg" api:"required"`
	OccurredAt time.Time   `json:"occurredAt" api:"required" format:"date-time"`
	// Any of "transfer".
	Type string `json:"type" api:"required"`
	Note string `json:"note" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Leg         respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementActivityPlacementTransferEvent) RawJSON() string { return r.JSON.raw }
func (r *PlacementActivityPlacementTransferEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementActivityPlacementCheckupEvent struct {
	ID         string    `json:"id" api:"required"`
	OccurredAt time.Time `json:"occurredAt" api:"required" format:"date-time"`
	// Any of "checkup".
	Type     string                                         `json:"type" api:"required"`
	FollowUp PlacementActivityPlacementCheckupEventFollowUp `json:"followUp"`
	Note     string                                         `json:"note" api:"nullable"`
	Record   VaccinationRecord                              `json:"record"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		FollowUp    respjson.Field
		Note        respjson.Field
		Record      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementActivityPlacementCheckupEvent) RawJSON() string { return r.JSON.raw }
func (r *PlacementActivityPlacementCheckupEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementActivityPlacementCheckupEventFollowUp struct {
	Due    time.Time `json:"due" format:"date-time"`
	Reason string    `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Due         respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementActivityPlacementCheckupEventFollowUp) RawJSON() string { return r.JSON.raw }
func (r *PlacementActivityPlacementCheckupEventFollowUp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementActivityPlacementDisruptionEvent struct {
	ID         string    `json:"id" api:"required"`
	OccurredAt time.Time `json:"occurredAt" api:"required" format:"date-time"`
	// A numeric severity score or a structured assessment.
	Severity PlacementActivityPlacementDisruptionEventSeverityUnion `json:"severity" api:"required"`
	// Any of "disruption".
	Type string `json:"type" api:"required"`
	Note string `json:"note" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		OccurredAt  respjson.Field
		Severity    respjson.Field
		Type        respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementActivityPlacementDisruptionEvent) RawJSON() string { return r.JSON.raw }
func (r *PlacementActivityPlacementDisruptionEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PlacementActivityPlacementDisruptionEventSeverityUnion contains all possible
// properties and values from [int64],
// [PlacementActivityPlacementDisruptionEventSeverityAssessment].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt]
type PlacementActivityPlacementDisruptionEventSeverityUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field is from variant
	// [PlacementActivityPlacementDisruptionEventSeverityAssessment].
	Level string `json:"level"`
	// This field is from variant
	// [PlacementActivityPlacementDisruptionEventSeverityAssessment].
	Reviewer string `json:"reviewer"`
	JSON     struct {
		OfInt    respjson.Field
		Level    respjson.Field
		Reviewer respjson.Field
		raw      string
	} `json:"-"`
}

func (u PlacementActivityPlacementDisruptionEventSeverityUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlacementActivityPlacementDisruptionEventSeverityUnion) AsAssessment() (v PlacementActivityPlacementDisruptionEventSeverityAssessment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlacementActivityPlacementDisruptionEventSeverityUnion) RawJSON() string { return u.JSON.raw }

func (r *PlacementActivityPlacementDisruptionEventSeverityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementActivityPlacementDisruptionEventSeverityAssessment struct {
	// Any of "low", "high", "critical".
	Level    string `json:"level" api:"required"`
	Reviewer string `json:"reviewer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		Reviewer    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementActivityPlacementDisruptionEventSeverityAssessment) RawJSON() string {
	return r.JSON.raw
}
func (r *PlacementActivityPlacementDisruptionEventSeverityAssessment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementActivityNote struct {
	Body string        `json:"body" api:"required"`
	Kind constant.Note `json:"kind" default:"note"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementActivityNote) RawJSON() string { return r.JSON.raw }
func (r *PlacementActivityNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementActivityAdopted struct {
	AdoptedAt time.Time        `json:"adoptedAt" api:"required" format:"date-time"`
	Kind      constant.Adopted `json:"kind" default:"adopted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdoptedAt   respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementActivityAdopted) RawJSON() string { return r.JSON.raw }
func (r *PlacementActivityAdopted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementActivityReturned struct {
	Kind   constant.Returned `json:"kind" default:"returned"`
	Reason string            `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementActivityReturned) RawJSON() string { return r.JSON.raw }
func (r *PlacementActivityReturned) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transport plan for a placement; pickup and delivery share the transfer-leg
// shape.
type PlacementLogistics struct {
	Delivery TransferLeg               `json:"delivery"`
	Handler  PlacementLogisticsHandler `json:"handler"`
	Pickup   TransferLeg               `json:"pickup"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delivery    respjson.Field
		Handler     respjson.Field
		Pickup      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementLogistics) RawJSON() string { return r.JSON.raw }
func (r *PlacementLogistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementLogisticsHandler struct {
	Certification PlacementLogisticsHandlerCertification `json:"certification"`
	Name          string                                 `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Certification respjson.Field
		Name          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementLogisticsHandler) RawJSON() string { return r.JSON.raw }
func (r *PlacementLogisticsHandler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementLogisticsHandlerCertification struct {
	// Any of "standard", "advanced", "specialist".
	Level string `json:"level" api:"required"`
	// Provenance stamp attached to records across unrelated resources.
	Audit    PlacementLogisticsHandlerCertificationAudit `json:"audit"`
	IssuedBy string                                      `json:"issuedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		Audit       respjson.Field
		IssuedBy    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementLogisticsHandlerCertification) RawJSON() string { return r.JSON.raw }
func (r *PlacementLogisticsHandlerCertification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provenance stamp attached to records across unrelated resources.
type PlacementLogisticsHandlerCertificationAudit struct {
	RecordedAt time.Time `json:"recordedAt" api:"required" format:"date-time"`
	RecordedBy string    `json:"recordedBy"`
	// Any of "manual", "import", "system".
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordedAt  respjson.Field
		RecordedBy  respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementLogisticsHandlerCertificationAudit) RawJSON() string { return r.JSON.raw }
func (r *PlacementLogisticsHandlerCertificationAudit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PlacementEventUnion contains all possible properties and values from
// [PlacementEventPlacementTransferEvent], [PlacementEventPlacementCheckupEvent],
// [PlacementEventPlacementDisruptionEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PlacementEventUnion struct {
	ID string `json:"id"`
	// This field is from variant [PlacementEventPlacementTransferEvent].
	Leg        TransferLeg `json:"leg"`
	OccurredAt time.Time   `json:"occurredAt"`
	Type       string      `json:"type"`
	Note       string      `json:"note"`
	// This field is from variant [PlacementEventPlacementCheckupEvent].
	FollowUp PlacementEventPlacementCheckupEventFollowUp `json:"followUp"`
	// This field is from variant [PlacementEventPlacementCheckupEvent].
	Record VaccinationRecord `json:"record"`
	// This field is from variant [PlacementEventPlacementDisruptionEvent].
	Severity PlacementEventPlacementDisruptionEventSeverityUnion `json:"severity"`
	JSON     struct {
		ID         respjson.Field
		Leg        respjson.Field
		OccurredAt respjson.Field
		Type       respjson.Field
		Note       respjson.Field
		FollowUp   respjson.Field
		Record     respjson.Field
		Severity   respjson.Field
		raw        string
	} `json:"-"`
}

func (u PlacementEventUnion) AsPlacementEventPlacementTransferEvent() (v PlacementEventPlacementTransferEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlacementEventUnion) AsPlacementEventPlacementCheckupEvent() (v PlacementEventPlacementCheckupEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlacementEventUnion) AsPlacementEventPlacementDisruptionEvent() (v PlacementEventPlacementDisruptionEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlacementEventUnion) RawJSON() string { return u.JSON.raw }

func (r *PlacementEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PlacementEventUnion to a PlacementEventUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PlacementEventUnionParam.Overrides()
func (r PlacementEventUnion) ToParam() PlacementEventUnionParam {
	return param.Override[PlacementEventUnionParam](json.RawMessage(r.RawJSON()))
}

type PlacementEventPlacementTransferEvent struct {
	ID         string      `json:"id" api:"required"`
	Leg        TransferLeg `json:"leg" api:"required"`
	OccurredAt time.Time   `json:"occurredAt" api:"required" format:"date-time"`
	// Any of "transfer".
	Type string `json:"type" api:"required"`
	Note string `json:"note" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Leg         respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementEventPlacementTransferEvent) RawJSON() string { return r.JSON.raw }
func (r *PlacementEventPlacementTransferEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementEventPlacementCheckupEvent struct {
	ID         string    `json:"id" api:"required"`
	OccurredAt time.Time `json:"occurredAt" api:"required" format:"date-time"`
	// Any of "checkup".
	Type     string                                      `json:"type" api:"required"`
	FollowUp PlacementEventPlacementCheckupEventFollowUp `json:"followUp"`
	Note     string                                      `json:"note" api:"nullable"`
	Record   VaccinationRecord                           `json:"record"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		OccurredAt  respjson.Field
		Type        respjson.Field
		FollowUp    respjson.Field
		Note        respjson.Field
		Record      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementEventPlacementCheckupEvent) RawJSON() string { return r.JSON.raw }
func (r *PlacementEventPlacementCheckupEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementEventPlacementCheckupEventFollowUp struct {
	Due    time.Time `json:"due" format:"date-time"`
	Reason string    `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Due         respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementEventPlacementCheckupEventFollowUp) RawJSON() string { return r.JSON.raw }
func (r *PlacementEventPlacementCheckupEventFollowUp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementEventPlacementDisruptionEvent struct {
	ID         string    `json:"id" api:"required"`
	OccurredAt time.Time `json:"occurredAt" api:"required" format:"date-time"`
	// A numeric severity score or a structured assessment.
	Severity PlacementEventPlacementDisruptionEventSeverityUnion `json:"severity" api:"required"`
	// Any of "disruption".
	Type string `json:"type" api:"required"`
	Note string `json:"note" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		OccurredAt  respjson.Field
		Severity    respjson.Field
		Type        respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementEventPlacementDisruptionEvent) RawJSON() string { return r.JSON.raw }
func (r *PlacementEventPlacementDisruptionEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PlacementEventPlacementDisruptionEventSeverityUnion contains all possible
// properties and values from [int64],
// [PlacementEventPlacementDisruptionEventSeverityAssessment].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt]
type PlacementEventPlacementDisruptionEventSeverityUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field is from variant
	// [PlacementEventPlacementDisruptionEventSeverityAssessment].
	Level string `json:"level"`
	// This field is from variant
	// [PlacementEventPlacementDisruptionEventSeverityAssessment].
	Reviewer string `json:"reviewer"`
	JSON     struct {
		OfInt    respjson.Field
		Level    respjson.Field
		Reviewer respjson.Field
		raw      string
	} `json:"-"`
}

func (u PlacementEventPlacementDisruptionEventSeverityUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlacementEventPlacementDisruptionEventSeverityUnion) AsAssessment() (v PlacementEventPlacementDisruptionEventSeverityAssessment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlacementEventPlacementDisruptionEventSeverityUnion) RawJSON() string { return u.JSON.raw }

func (r *PlacementEventPlacementDisruptionEventSeverityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementEventPlacementDisruptionEventSeverityAssessment struct {
	// Any of "low", "high", "critical".
	Level    string `json:"level" api:"required"`
	Reviewer string `json:"reviewer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		Reviewer    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementEventPlacementDisruptionEventSeverityAssessment) RawJSON() string { return r.JSON.raw }
func (r *PlacementEventPlacementDisruptionEventSeverityAssessment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func PlacementEventParamOfPlacementEventPlacementCheckupEvent(id string, occurredAt time.Time, type_ string) PlacementEventUnionParam {
	var variant PlacementEventPlacementCheckupEventParam
	variant.ID = id
	variant.OccurredAt = occurredAt
	variant.Type = type_
	return PlacementEventUnionParam{OfPlacementEventPlacementCheckupEvent: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PlacementEventUnionParam struct {
	OfPlacementEventPlacementTransferEvent   *PlacementEventPlacementTransferEventParam   `json:",omitzero,inline"`
	OfPlacementEventPlacementCheckupEvent    *PlacementEventPlacementCheckupEventParam    `json:",omitzero,inline"`
	OfPlacementEventPlacementDisruptionEvent *PlacementEventPlacementDisruptionEventParam `json:",omitzero,inline"`
	paramUnion
}

func (u PlacementEventUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPlacementEventPlacementTransferEvent, u.OfPlacementEventPlacementCheckupEvent, u.OfPlacementEventPlacementDisruptionEvent)
}
func (u *PlacementEventUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ID, Leg, OccurredAt, Type are required.
type PlacementEventPlacementTransferEventParam struct {
	ID         string           `json:"id" api:"required"`
	Leg        TransferLegParam `json:"leg,omitzero" api:"required"`
	OccurredAt time.Time        `json:"occurredAt" api:"required" format:"date-time"`
	// Any of "transfer".
	Type string            `json:"type,omitzero" api:"required"`
	Note param.Opt[string] `json:"note,omitzero"`
	paramObj
}

func (r PlacementEventPlacementTransferEventParam) MarshalJSON() (data []byte, err error) {
	type shadow PlacementEventPlacementTransferEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementEventPlacementTransferEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PlacementEventPlacementTransferEventParam](
		"type", "transfer",
	)
}

// The properties ID, OccurredAt, Type are required.
type PlacementEventPlacementCheckupEventParam struct {
	ID         string    `json:"id" api:"required"`
	OccurredAt time.Time `json:"occurredAt" api:"required" format:"date-time"`
	// Any of "checkup".
	Type     string                                           `json:"type,omitzero" api:"required"`
	Note     param.Opt[string]                                `json:"note,omitzero"`
	FollowUp PlacementEventPlacementCheckupEventFollowUpParam `json:"followUp,omitzero"`
	Record   VaccinationRecordParam                           `json:"record,omitzero"`
	paramObj
}

func (r PlacementEventPlacementCheckupEventParam) MarshalJSON() (data []byte, err error) {
	type shadow PlacementEventPlacementCheckupEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementEventPlacementCheckupEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PlacementEventPlacementCheckupEventParam](
		"type", "checkup",
	)
}

type PlacementEventPlacementCheckupEventFollowUpParam struct {
	Due    param.Opt[time.Time] `json:"due,omitzero" format:"date-time"`
	Reason param.Opt[string]    `json:"reason,omitzero"`
	paramObj
}

func (r PlacementEventPlacementCheckupEventFollowUpParam) MarshalJSON() (data []byte, err error) {
	type shadow PlacementEventPlacementCheckupEventFollowUpParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementEventPlacementCheckupEventFollowUpParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, OccurredAt, Severity, Type are required.
type PlacementEventPlacementDisruptionEventParam struct {
	ID         string    `json:"id" api:"required"`
	OccurredAt time.Time `json:"occurredAt" api:"required" format:"date-time"`
	// A numeric severity score or a structured assessment.
	Severity PlacementEventPlacementDisruptionEventSeverityUnionParam `json:"severity,omitzero" api:"required"`
	// Any of "disruption".
	Type string            `json:"type,omitzero" api:"required"`
	Note param.Opt[string] `json:"note,omitzero"`
	paramObj
}

func (r PlacementEventPlacementDisruptionEventParam) MarshalJSON() (data []byte, err error) {
	type shadow PlacementEventPlacementDisruptionEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementEventPlacementDisruptionEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PlacementEventPlacementDisruptionEventParam](
		"type", "disruption",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PlacementEventPlacementDisruptionEventSeverityUnionParam struct {
	OfInt        param.Opt[int64]                                               `json:",omitzero,inline"`
	OfAssessment *PlacementEventPlacementDisruptionEventSeverityAssessmentParam `json:",omitzero,inline"`
	paramUnion
}

func (u PlacementEventPlacementDisruptionEventSeverityUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfAssessment)
}
func (u *PlacementEventPlacementDisruptionEventSeverityUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property Level is required.
type PlacementEventPlacementDisruptionEventSeverityAssessmentParam struct {
	// Any of "low", "high", "critical".
	Level    string            `json:"level,omitzero" api:"required"`
	Reviewer param.Opt[string] `json:"reviewer,omitzero"`
	paramObj
}

func (r PlacementEventPlacementDisruptionEventSeverityAssessmentParam) MarshalJSON() (data []byte, err error) {
	type shadow PlacementEventPlacementDisruptionEventSeverityAssessmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementEventPlacementDisruptionEventSeverityAssessmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PlacementEventPlacementDisruptionEventSeverityAssessmentParam](
		"level", "low", "high", "critical",
	)
}

type TransferLeg struct {
	Location shared.Address     `json:"location" api:"required"`
	Contact  TransferLegContact `json:"contact"`
	Cost     shared.Money       `json:"cost"`
	Window   TransferLegWindow  `json:"window"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Location    respjson.Field
		Contact     respjson.Field
		Cost        respjson.Field
		Window      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferLeg) RawJSON() string { return r.JSON.raw }
func (r *TransferLeg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TransferLeg to a TransferLegParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TransferLegParam.Overrides()
func (r TransferLeg) ToParam() TransferLegParam {
	return param.Override[TransferLegParam](json.RawMessage(r.RawJSON()))
}

type TransferLegContact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferLegContact) RawJSON() string { return r.JSON.raw }
func (r *TransferLegContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferLegWindow struct {
	End   time.Time `json:"end" format:"date-time"`
	Start time.Time `json:"start" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferLegWindow) RawJSON() string { return r.JSON.raw }
func (r *TransferLegWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Location is required.
type TransferLegParam struct {
	Location shared.AddressParam     `json:"location,omitzero" api:"required"`
	Contact  TransferLegContactParam `json:"contact,omitzero"`
	Cost     shared.MoneyParam       `json:"cost,omitzero"`
	Window   TransferLegWindowParam  `json:"window,omitzero"`
	paramObj
}

func (r TransferLegParam) MarshalJSON() (data []byte, err error) {
	type shadow TransferLegParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferLegParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferLegContactParam struct {
	Name  param.Opt[string] `json:"name,omitzero"`
	Phone param.Opt[string] `json:"phone,omitzero"`
	paramObj
}

func (r TransferLegContactParam) MarshalJSON() (data []byte, err error) {
	type shadow TransferLegContactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferLegContactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferLegWindowParam struct {
	End   param.Opt[time.Time] `json:"end,omitzero" format:"date-time"`
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	paramObj
}

func (r TransferLegWindowParam) MarshalJSON() (data []byte, err error) {
	type shadow TransferLegWindowParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferLegWindowParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementNewParams struct {
	ApplicationID string               `json:"applicationId" api:"required"`
	ScheduledFor  param.Opt[time.Time] `json:"scheduledFor,omitzero" format:"date-time"`
	// A placement history entry; the shape depends on the event kind and carries no
	// discriminator mapping.
	InitialEvent PlacementEventUnionParam `json:"initialEvent,omitzero"`
	// Transport plan for a placement; pickup and delivery share the transfer-leg
	// shape.
	Logistics PlacementNewParamsLogistics `json:"logistics,omitzero"`
	paramObj
}

func (r PlacementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PlacementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transport plan for a placement; pickup and delivery share the transfer-leg
// shape.
type PlacementNewParamsLogistics struct {
	Delivery TransferLegParam                   `json:"delivery,omitzero"`
	Handler  PlacementNewParamsLogisticsHandler `json:"handler,omitzero"`
	Pickup   TransferLegParam                   `json:"pickup,omitzero"`
	paramObj
}

func (r PlacementNewParamsLogistics) MarshalJSON() (data []byte, err error) {
	type shadow PlacementNewParamsLogistics
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementNewParamsLogistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementNewParamsLogisticsHandler struct {
	Name          param.Opt[string]                               `json:"name,omitzero"`
	Certification PlacementNewParamsLogisticsHandlerCertification `json:"certification,omitzero"`
	paramObj
}

func (r PlacementNewParamsLogisticsHandler) MarshalJSON() (data []byte, err error) {
	type shadow PlacementNewParamsLogisticsHandler
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementNewParamsLogisticsHandler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Level is required.
type PlacementNewParamsLogisticsHandlerCertification struct {
	// Any of "standard", "advanced", "specialist".
	Level    string            `json:"level,omitzero" api:"required"`
	IssuedBy param.Opt[string] `json:"issuedBy,omitzero"`
	// Provenance stamp attached to records across unrelated resources.
	Audit PlacementNewParamsLogisticsHandlerCertificationAudit `json:"audit,omitzero"`
	paramObj
}

func (r PlacementNewParamsLogisticsHandlerCertification) MarshalJSON() (data []byte, err error) {
	type shadow PlacementNewParamsLogisticsHandlerCertification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementNewParamsLogisticsHandlerCertification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PlacementNewParamsLogisticsHandlerCertification](
		"level", "standard", "advanced", "specialist",
	)
}

// Provenance stamp attached to records across unrelated resources.
//
// The property RecordedAt is required.
type PlacementNewParamsLogisticsHandlerCertificationAudit struct {
	RecordedAt time.Time         `json:"recordedAt" api:"required" format:"date-time"`
	RecordedBy param.Opt[string] `json:"recordedBy,omitzero"`
	// Any of "manual", "import", "system".
	Source string `json:"source,omitzero"`
	paramObj
}

func (r PlacementNewParamsLogisticsHandlerCertificationAudit) MarshalJSON() (data []byte, err error) {
	type shadow PlacementNewParamsLogisticsHandlerCertificationAudit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementNewParamsLogisticsHandlerCertificationAudit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PlacementNewParamsLogisticsHandlerCertificationAudit](
		"source", "manual", "import", "system",
	)
}

type PlacementListParams struct {
	// Cursor from a previous response used to fetch the next page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of placements to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PlacementListParams]'s query parameters as `url.Values`.
func (r PlacementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlacementRecordEventParams struct {
	// A placement history entry; the shape depends on the event kind and carries no
	// discriminator mapping.
	PlacementEvent PlacementEventUnionParam
	paramObj
}

func (r PlacementRecordEventParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PlacementEvent)
}
func (r *PlacementRecordEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
