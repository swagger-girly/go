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

// Post-adoption placement tracking
//
// PlacementService contains methods and other services that help with interacting
// with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlacementService] method instead.
type PlacementService struct {
	Options []option.RequestOption
}

// NewPlacementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlacementService(opts ...option.RequestOption) (r *PlacementService) {
	r = &PlacementService{}
	r.Options = opts
	return
}

// Creates a placement for an approved adoption application.
func (r *PlacementService) New(ctx context.Context, body PlacementNewParams, opts ...option.RequestOption) (res *Placement, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "placements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single placement, including its adoption application and event
// history.
func (r *PlacementService) Get(ctx context.Context, placementID string, opts ...option.RequestOption) (res *Placement, err error) {
	opts = slices.Concat(r.Options, opts)
	if placementID == "" {
		err = errors.New("missing required placementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("placements/%s", placementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a cursor-paginated list of placements.
func (r *PlacementService) List(ctx context.Context, query PlacementListParams, opts ...option.RequestOption) (res *pagination.CustomCursorPage[Placement], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	if placementID == "" {
		err = errors.New("missing required placementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("placements/%s/events", placementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A post-adoption placement tracking the pet's transition to its new home.
type Placement struct {
	ID            string           `json:"id" api:"required"`
	Application   Application      `json:"application" api:"required"`
	CreatedAt     time.Time        `json:"createdAt" api:"required" format:"date-time"`
	Status        PlacementStatus  `json:"status" api:"required"`
	Events        []PlacementEvent `json:"events"`
	FollowUpAfter time.Time        `json:"followUpAfter" api:"nullable" format:"date-time"`
	// Transport plan for a placement; pickup and delivery share the transfer-leg
	// shape.
	Logistics        PlacementLogistics `json:"logistics"`
	MedicalClearance MedicalSummary     `json:"medicalClearance"`
	JSON             placementJSON      `json:"-"`
}

// placementJSON contains the JSON metadata for the struct [Placement]
type placementJSON struct {
	ID               apijson.Field
	Application      apijson.Field
	CreatedAt        apijson.Field
	Status           apijson.Field
	Events           apijson.Field
	FollowUpAfter    apijson.Field
	Logistics        apijson.Field
	MedicalClearance apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Placement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementJSON) RawJSON() string {
	return r.raw
}

type PlacementStatus string

const (
	PlacementStatusPending   PlacementStatus = "pending"
	PlacementStatusActive    PlacementStatus = "active"
	PlacementStatusCompleted PlacementStatus = "completed"
	PlacementStatusDisrupted PlacementStatus = "disrupted"
)

func (r PlacementStatus) IsKnown() bool {
	switch r {
	case PlacementStatusPending, PlacementStatusActive, PlacementStatusCompleted, PlacementStatusDisrupted:
		return true
	}
	return false
}

// Transport plan for a placement; pickup and delivery share the transfer-leg
// shape.
type PlacementLogistics struct {
	Delivery TransferLeg               `json:"delivery"`
	Handler  PlacementLogisticsHandler `json:"handler"`
	Pickup   TransferLeg               `json:"pickup"`
	JSON     placementLogisticsJSON    `json:"-"`
}

// placementLogisticsJSON contains the JSON metadata for the struct
// [PlacementLogistics]
type placementLogisticsJSON struct {
	Delivery    apijson.Field
	Handler     apijson.Field
	Pickup      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementLogistics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementLogisticsJSON) RawJSON() string {
	return r.raw
}

type PlacementLogisticsHandler struct {
	Certification PlacementLogisticsHandlerCertification `json:"certification"`
	Name          string                                 `json:"name"`
	JSON          placementLogisticsHandlerJSON          `json:"-"`
}

// placementLogisticsHandlerJSON contains the JSON metadata for the struct
// [PlacementLogisticsHandler]
type placementLogisticsHandlerJSON struct {
	Certification apijson.Field
	Name          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PlacementLogisticsHandler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementLogisticsHandlerJSON) RawJSON() string {
	return r.raw
}

type PlacementLogisticsHandlerCertification struct {
	Level PlacementLogisticsHandlerCertificationLevel `json:"level" api:"required"`
	// Provenance stamp attached to records across unrelated resources.
	Audit    PlacementLogisticsHandlerCertificationAudit `json:"audit"`
	IssuedBy string                                      `json:"issuedBy"`
	JSON     placementLogisticsHandlerCertificationJSON  `json:"-"`
}

// placementLogisticsHandlerCertificationJSON contains the JSON metadata for the
// struct [PlacementLogisticsHandlerCertification]
type placementLogisticsHandlerCertificationJSON struct {
	Level       apijson.Field
	Audit       apijson.Field
	IssuedBy    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementLogisticsHandlerCertification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementLogisticsHandlerCertificationJSON) RawJSON() string {
	return r.raw
}

type PlacementLogisticsHandlerCertificationLevel string

const (
	PlacementLogisticsHandlerCertificationLevelStandard   PlacementLogisticsHandlerCertificationLevel = "standard"
	PlacementLogisticsHandlerCertificationLevelAdvanced   PlacementLogisticsHandlerCertificationLevel = "advanced"
	PlacementLogisticsHandlerCertificationLevelSpecialist PlacementLogisticsHandlerCertificationLevel = "specialist"
)

func (r PlacementLogisticsHandlerCertificationLevel) IsKnown() bool {
	switch r {
	case PlacementLogisticsHandlerCertificationLevelStandard, PlacementLogisticsHandlerCertificationLevelAdvanced, PlacementLogisticsHandlerCertificationLevelSpecialist:
		return true
	}
	return false
}

// Provenance stamp attached to records across unrelated resources.
type PlacementLogisticsHandlerCertificationAudit struct {
	RecordedAt time.Time                                         `json:"recordedAt" api:"required" format:"date-time"`
	RecordedBy string                                            `json:"recordedBy"`
	Source     PlacementLogisticsHandlerCertificationAuditSource `json:"source"`
	JSON       placementLogisticsHandlerCertificationAuditJSON   `json:"-"`
}

// placementLogisticsHandlerCertificationAuditJSON contains the JSON metadata for
// the struct [PlacementLogisticsHandlerCertificationAudit]
type placementLogisticsHandlerCertificationAuditJSON struct {
	RecordedAt  apijson.Field
	RecordedBy  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementLogisticsHandlerCertificationAudit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementLogisticsHandlerCertificationAuditJSON) RawJSON() string {
	return r.raw
}

type PlacementLogisticsHandlerCertificationAuditSource string

const (
	PlacementLogisticsHandlerCertificationAuditSourceManual PlacementLogisticsHandlerCertificationAuditSource = "manual"
	PlacementLogisticsHandlerCertificationAuditSourceImport PlacementLogisticsHandlerCertificationAuditSource = "import"
	PlacementLogisticsHandlerCertificationAuditSourceSystem PlacementLogisticsHandlerCertificationAuditSource = "system"
)

func (r PlacementLogisticsHandlerCertificationAuditSource) IsKnown() bool {
	switch r {
	case PlacementLogisticsHandlerCertificationAuditSourceManual, PlacementLogisticsHandlerCertificationAuditSourceImport, PlacementLogisticsHandlerCertificationAuditSourceSystem:
		return true
	}
	return false
}

// A placement history entry; the shape depends on the event kind and carries no
// discriminator mapping.
type PlacementEvent struct {
	ID         string             `json:"id" api:"required"`
	OccurredAt time.Time          `json:"occurredAt" api:"required" format:"date-time"`
	Type       PlacementEventType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [PlacementEventPlacementCheckupEventFollowUp].
	FollowUp interface{}       `json:"followUp"`
	Leg      TransferLeg       `json:"leg"`
	Note     string            `json:"note" api:"nullable"`
	Record   VaccinationRecord `json:"record"`
	// This field can have the runtime type of
	// [PlacementEventPlacementDisruptionEventSeverityUnion].
	Severity interface{}        `json:"severity"`
	JSON     placementEventJSON `json:"-"`
	union    PlacementEventUnion
}

// placementEventJSON contains the JSON metadata for the struct [PlacementEvent]
type placementEventJSON struct {
	ID          apijson.Field
	OccurredAt  apijson.Field
	Type        apijson.Field
	FollowUp    apijson.Field
	Leg         apijson.Field
	Note        apijson.Field
	Record      apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r placementEventJSON) RawJSON() string {
	return r.raw
}

func (r *PlacementEvent) UnmarshalJSON(data []byte) (err error) {
	*r = PlacementEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PlacementEventUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [PlacementEventPlacementTransferEvent],
// [PlacementEventPlacementCheckupEvent], [PlacementEventPlacementDisruptionEvent].
func (r PlacementEvent) AsUnion() PlacementEventUnion {
	return r.union
}

// A placement history entry; the shape depends on the event kind and carries no
// discriminator mapping.
//
// Union satisfied by [PlacementEventPlacementTransferEvent],
// [PlacementEventPlacementCheckupEvent] or
// [PlacementEventPlacementDisruptionEvent].
type PlacementEventUnion interface {
	implementsPlacementEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlacementEventUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PlacementEventPlacementTransferEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PlacementEventPlacementCheckupEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PlacementEventPlacementDisruptionEvent{}),
		},
	)
}

type PlacementEventPlacementTransferEvent struct {
	ID         string                                   `json:"id" api:"required"`
	Leg        TransferLeg                              `json:"leg" api:"required"`
	OccurredAt time.Time                                `json:"occurredAt" api:"required" format:"date-time"`
	Type       PlacementEventPlacementTransferEventType `json:"type" api:"required"`
	Note       string                                   `json:"note" api:"nullable"`
	JSON       placementEventPlacementTransferEventJSON `json:"-"`
}

// placementEventPlacementTransferEventJSON contains the JSON metadata for the
// struct [PlacementEventPlacementTransferEvent]
type placementEventPlacementTransferEventJSON struct {
	ID          apijson.Field
	Leg         apijson.Field
	OccurredAt  apijson.Field
	Type        apijson.Field
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementEventPlacementTransferEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementEventPlacementTransferEventJSON) RawJSON() string {
	return r.raw
}

func (r PlacementEventPlacementTransferEvent) implementsPlacementEvent() {}

type PlacementEventPlacementTransferEventType string

const (
	PlacementEventPlacementTransferEventTypeTransfer PlacementEventPlacementTransferEventType = "transfer"
)

func (r PlacementEventPlacementTransferEventType) IsKnown() bool {
	switch r {
	case PlacementEventPlacementTransferEventTypeTransfer:
		return true
	}
	return false
}

type PlacementEventPlacementCheckupEvent struct {
	ID         string                                      `json:"id" api:"required"`
	OccurredAt time.Time                                   `json:"occurredAt" api:"required" format:"date-time"`
	Type       PlacementEventPlacementCheckupEventType     `json:"type" api:"required"`
	FollowUp   PlacementEventPlacementCheckupEventFollowUp `json:"followUp"`
	Note       string                                      `json:"note" api:"nullable"`
	Record     VaccinationRecord                           `json:"record"`
	JSON       placementEventPlacementCheckupEventJSON     `json:"-"`
}

// placementEventPlacementCheckupEventJSON contains the JSON metadata for the
// struct [PlacementEventPlacementCheckupEvent]
type placementEventPlacementCheckupEventJSON struct {
	ID          apijson.Field
	OccurredAt  apijson.Field
	Type        apijson.Field
	FollowUp    apijson.Field
	Note        apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementEventPlacementCheckupEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementEventPlacementCheckupEventJSON) RawJSON() string {
	return r.raw
}

func (r PlacementEventPlacementCheckupEvent) implementsPlacementEvent() {}

type PlacementEventPlacementCheckupEventType string

const (
	PlacementEventPlacementCheckupEventTypeCheckup PlacementEventPlacementCheckupEventType = "checkup"
)

func (r PlacementEventPlacementCheckupEventType) IsKnown() bool {
	switch r {
	case PlacementEventPlacementCheckupEventTypeCheckup:
		return true
	}
	return false
}

type PlacementEventPlacementCheckupEventFollowUp struct {
	Due    time.Time                                       `json:"due" format:"date-time"`
	Reason string                                          `json:"reason"`
	JSON   placementEventPlacementCheckupEventFollowUpJSON `json:"-"`
}

// placementEventPlacementCheckupEventFollowUpJSON contains the JSON metadata for
// the struct [PlacementEventPlacementCheckupEventFollowUp]
type placementEventPlacementCheckupEventFollowUpJSON struct {
	Due         apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementEventPlacementCheckupEventFollowUp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementEventPlacementCheckupEventFollowUpJSON) RawJSON() string {
	return r.raw
}

type PlacementEventPlacementDisruptionEvent struct {
	ID         string    `json:"id" api:"required"`
	OccurredAt time.Time `json:"occurredAt" api:"required" format:"date-time"`
	// A numeric severity score or a structured assessment.
	Severity PlacementEventPlacementDisruptionEventSeverityUnion `json:"severity" api:"required"`
	Type     PlacementEventPlacementDisruptionEventType          `json:"type" api:"required"`
	Note     string                                              `json:"note" api:"nullable"`
	JSON     placementEventPlacementDisruptionEventJSON          `json:"-"`
}

// placementEventPlacementDisruptionEventJSON contains the JSON metadata for the
// struct [PlacementEventPlacementDisruptionEvent]
type placementEventPlacementDisruptionEventJSON struct {
	ID          apijson.Field
	OccurredAt  apijson.Field
	Severity    apijson.Field
	Type        apijson.Field
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementEventPlacementDisruptionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementEventPlacementDisruptionEventJSON) RawJSON() string {
	return r.raw
}

func (r PlacementEventPlacementDisruptionEvent) implementsPlacementEvent() {}

// A numeric severity score or a structured assessment.
//
// Union satisfied by [shared.UnionInt] or
// [PlacementEventPlacementDisruptionEventSeverityAssessment].
type PlacementEventPlacementDisruptionEventSeverityUnion interface {
	ImplementsPlacementEventPlacementDisruptionEventSeverityUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlacementEventPlacementDisruptionEventSeverityUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PlacementEventPlacementDisruptionEventSeverityAssessment{}),
		},
	)
}

type PlacementEventPlacementDisruptionEventSeverityAssessment struct {
	Level    PlacementEventPlacementDisruptionEventSeverityAssessmentLevel `json:"level" api:"required"`
	Reviewer string                                                        `json:"reviewer"`
	JSON     placementEventPlacementDisruptionEventSeverityAssessmentJSON  `json:"-"`
}

// placementEventPlacementDisruptionEventSeverityAssessmentJSON contains the JSON
// metadata for the struct
// [PlacementEventPlacementDisruptionEventSeverityAssessment]
type placementEventPlacementDisruptionEventSeverityAssessmentJSON struct {
	Level       apijson.Field
	Reviewer    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementEventPlacementDisruptionEventSeverityAssessment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementEventPlacementDisruptionEventSeverityAssessmentJSON) RawJSON() string {
	return r.raw
}

func (r PlacementEventPlacementDisruptionEventSeverityAssessment) ImplementsPlacementEventPlacementDisruptionEventSeverityUnion() {
}

type PlacementEventPlacementDisruptionEventSeverityAssessmentLevel string

const (
	PlacementEventPlacementDisruptionEventSeverityAssessmentLevelLow      PlacementEventPlacementDisruptionEventSeverityAssessmentLevel = "low"
	PlacementEventPlacementDisruptionEventSeverityAssessmentLevelHigh     PlacementEventPlacementDisruptionEventSeverityAssessmentLevel = "high"
	PlacementEventPlacementDisruptionEventSeverityAssessmentLevelCritical PlacementEventPlacementDisruptionEventSeverityAssessmentLevel = "critical"
)

func (r PlacementEventPlacementDisruptionEventSeverityAssessmentLevel) IsKnown() bool {
	switch r {
	case PlacementEventPlacementDisruptionEventSeverityAssessmentLevelLow, PlacementEventPlacementDisruptionEventSeverityAssessmentLevelHigh, PlacementEventPlacementDisruptionEventSeverityAssessmentLevelCritical:
		return true
	}
	return false
}

type PlacementEventPlacementDisruptionEventType string

const (
	PlacementEventPlacementDisruptionEventTypeDisruption PlacementEventPlacementDisruptionEventType = "disruption"
)

func (r PlacementEventPlacementDisruptionEventType) IsKnown() bool {
	switch r {
	case PlacementEventPlacementDisruptionEventTypeDisruption:
		return true
	}
	return false
}

type PlacementEventType string

const (
	PlacementEventTypeTransfer   PlacementEventType = "transfer"
	PlacementEventTypeCheckup    PlacementEventType = "checkup"
	PlacementEventTypeDisruption PlacementEventType = "disruption"
)

func (r PlacementEventType) IsKnown() bool {
	switch r {
	case PlacementEventTypeTransfer, PlacementEventTypeCheckup, PlacementEventTypeDisruption:
		return true
	}
	return false
}

// A placement history entry; the shape depends on the event kind and carries no
// discriminator mapping.
type PlacementEventParam struct {
	ID         param.Field[string]                 `json:"id" api:"required"`
	OccurredAt param.Field[time.Time]              `json:"occurredAt" api:"required" format:"date-time"`
	Type       param.Field[PlacementEventType]     `json:"type" api:"required"`
	FollowUp   param.Field[interface{}]            `json:"followUp"`
	Leg        param.Field[TransferLegParam]       `json:"leg"`
	Note       param.Field[string]                 `json:"note"`
	Record     param.Field[VaccinationRecordParam] `json:"record"`
	Severity   param.Field[interface{}]            `json:"severity"`
}

func (r PlacementEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlacementEventParam) implementsPlacementEventUnionParam() {}

// A placement history entry; the shape depends on the event kind and carries no
// discriminator mapping.
//
// Satisfied by [PlacementEventPlacementTransferEventParam],
// [PlacementEventPlacementCheckupEventParam],
// [PlacementEventPlacementDisruptionEventParam], [PlacementEventParam].
type PlacementEventUnionParam interface {
	implementsPlacementEventUnionParam()
}

type PlacementEventPlacementTransferEventParam struct {
	ID         param.Field[string]                                   `json:"id" api:"required"`
	Leg        param.Field[TransferLegParam]                         `json:"leg" api:"required"`
	OccurredAt param.Field[time.Time]                                `json:"occurredAt" api:"required" format:"date-time"`
	Type       param.Field[PlacementEventPlacementTransferEventType] `json:"type" api:"required"`
	Note       param.Field[string]                                   `json:"note"`
}

func (r PlacementEventPlacementTransferEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlacementEventPlacementTransferEventParam) implementsPlacementEventUnionParam() {}

type PlacementEventPlacementCheckupEventParam struct {
	ID         param.Field[string]                                           `json:"id" api:"required"`
	OccurredAt param.Field[time.Time]                                        `json:"occurredAt" api:"required" format:"date-time"`
	Type       param.Field[PlacementEventPlacementCheckupEventType]          `json:"type" api:"required"`
	FollowUp   param.Field[PlacementEventPlacementCheckupEventFollowUpParam] `json:"followUp"`
	Note       param.Field[string]                                           `json:"note"`
	Record     param.Field[VaccinationRecordParam]                           `json:"record"`
}

func (r PlacementEventPlacementCheckupEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlacementEventPlacementCheckupEventParam) implementsPlacementEventUnionParam() {}

type PlacementEventPlacementCheckupEventFollowUpParam struct {
	Due    param.Field[time.Time] `json:"due" format:"date-time"`
	Reason param.Field[string]    `json:"reason"`
}

func (r PlacementEventPlacementCheckupEventFollowUpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlacementEventPlacementDisruptionEventParam struct {
	ID         param.Field[string]    `json:"id" api:"required"`
	OccurredAt param.Field[time.Time] `json:"occurredAt" api:"required" format:"date-time"`
	// A numeric severity score or a structured assessment.
	Severity param.Field[PlacementEventPlacementDisruptionEventSeverityUnionParam] `json:"severity" api:"required"`
	Type     param.Field[PlacementEventPlacementDisruptionEventType]               `json:"type" api:"required"`
	Note     param.Field[string]                                                   `json:"note"`
}

func (r PlacementEventPlacementDisruptionEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlacementEventPlacementDisruptionEventParam) implementsPlacementEventUnionParam() {}

// A numeric severity score or a structured assessment.
//
// Satisfied by [shared.UnionInt],
// [PlacementEventPlacementDisruptionEventSeverityAssessmentParam].
type PlacementEventPlacementDisruptionEventSeverityUnionParam interface {
	ImplementsPlacementEventPlacementDisruptionEventSeverityUnionParam()
}

type PlacementEventPlacementDisruptionEventSeverityAssessmentParam struct {
	Level    param.Field[PlacementEventPlacementDisruptionEventSeverityAssessmentLevel] `json:"level" api:"required"`
	Reviewer param.Field[string]                                                        `json:"reviewer"`
}

func (r PlacementEventPlacementDisruptionEventSeverityAssessmentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlacementEventPlacementDisruptionEventSeverityAssessmentParam) ImplementsPlacementEventPlacementDisruptionEventSeverityUnionParam() {
}

type TransferLeg struct {
	Location shared.Address     `json:"location" api:"required"`
	Contact  TransferLegContact `json:"contact"`
	Cost     shared.Money       `json:"cost"`
	Window   TransferLegWindow  `json:"window"`
	JSON     transferLegJSON    `json:"-"`
}

// transferLegJSON contains the JSON metadata for the struct [TransferLeg]
type transferLegJSON struct {
	Location    apijson.Field
	Contact     apijson.Field
	Cost        apijson.Field
	Window      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferLeg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLegJSON) RawJSON() string {
	return r.raw
}

type TransferLegContact struct {
	Name  string                 `json:"name"`
	Phone string                 `json:"phone"`
	JSON  transferLegContactJSON `json:"-"`
}

// transferLegContactJSON contains the JSON metadata for the struct
// [TransferLegContact]
type transferLegContactJSON struct {
	Name        apijson.Field
	Phone       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferLegContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLegContactJSON) RawJSON() string {
	return r.raw
}

type TransferLegWindow struct {
	End   time.Time             `json:"end" format:"date-time"`
	Start time.Time             `json:"start" format:"date-time"`
	JSON  transferLegWindowJSON `json:"-"`
}

// transferLegWindowJSON contains the JSON metadata for the struct
// [TransferLegWindow]
type transferLegWindowJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferLegWindow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLegWindowJSON) RawJSON() string {
	return r.raw
}

type TransferLegParam struct {
	Location param.Field[shared.AddressParam]     `json:"location" api:"required"`
	Contact  param.Field[TransferLegContactParam] `json:"contact"`
	Cost     param.Field[shared.MoneyParam]       `json:"cost"`
	Window   param.Field[TransferLegWindowParam]  `json:"window"`
}

func (r TransferLegParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransferLegContactParam struct {
	Name  param.Field[string] `json:"name"`
	Phone param.Field[string] `json:"phone"`
}

func (r TransferLegContactParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransferLegWindowParam struct {
	End   param.Field[time.Time] `json:"end" format:"date-time"`
	Start param.Field[time.Time] `json:"start" format:"date-time"`
}

func (r TransferLegWindowParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlacementNewParams struct {
	ApplicationID param.Field[string] `json:"applicationId" api:"required"`
	// A placement history entry; the shape depends on the event kind and carries no
	// discriminator mapping.
	InitialEvent param.Field[PlacementEventUnionParam] `json:"initialEvent"`
	// Transport plan for a placement; pickup and delivery share the transfer-leg
	// shape.
	Logistics    param.Field[PlacementNewParamsLogistics] `json:"logistics"`
	ScheduledFor param.Field[time.Time]                   `json:"scheduledFor" format:"date-time"`
}

func (r PlacementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Transport plan for a placement; pickup and delivery share the transfer-leg
// shape.
type PlacementNewParamsLogistics struct {
	Delivery param.Field[TransferLegParam]                   `json:"delivery"`
	Handler  param.Field[PlacementNewParamsLogisticsHandler] `json:"handler"`
	Pickup   param.Field[TransferLegParam]                   `json:"pickup"`
}

func (r PlacementNewParamsLogistics) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlacementNewParamsLogisticsHandler struct {
	Certification param.Field[PlacementNewParamsLogisticsHandlerCertification] `json:"certification"`
	Name          param.Field[string]                                          `json:"name"`
}

func (r PlacementNewParamsLogisticsHandler) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlacementNewParamsLogisticsHandlerCertification struct {
	Level param.Field[PlacementNewParamsLogisticsHandlerCertificationLevel] `json:"level" api:"required"`
	// Provenance stamp attached to records across unrelated resources.
	Audit    param.Field[PlacementNewParamsLogisticsHandlerCertificationAudit] `json:"audit"`
	IssuedBy param.Field[string]                                               `json:"issuedBy"`
}

func (r PlacementNewParamsLogisticsHandlerCertification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlacementNewParamsLogisticsHandlerCertificationLevel string

const (
	PlacementNewParamsLogisticsHandlerCertificationLevelStandard   PlacementNewParamsLogisticsHandlerCertificationLevel = "standard"
	PlacementNewParamsLogisticsHandlerCertificationLevelAdvanced   PlacementNewParamsLogisticsHandlerCertificationLevel = "advanced"
	PlacementNewParamsLogisticsHandlerCertificationLevelSpecialist PlacementNewParamsLogisticsHandlerCertificationLevel = "specialist"
)

func (r PlacementNewParamsLogisticsHandlerCertificationLevel) IsKnown() bool {
	switch r {
	case PlacementNewParamsLogisticsHandlerCertificationLevelStandard, PlacementNewParamsLogisticsHandlerCertificationLevelAdvanced, PlacementNewParamsLogisticsHandlerCertificationLevelSpecialist:
		return true
	}
	return false
}

// Provenance stamp attached to records across unrelated resources.
type PlacementNewParamsLogisticsHandlerCertificationAudit struct {
	RecordedAt param.Field[time.Time]                                                  `json:"recordedAt" api:"required" format:"date-time"`
	RecordedBy param.Field[string]                                                     `json:"recordedBy"`
	Source     param.Field[PlacementNewParamsLogisticsHandlerCertificationAuditSource] `json:"source"`
}

func (r PlacementNewParamsLogisticsHandlerCertificationAudit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlacementNewParamsLogisticsHandlerCertificationAuditSource string

const (
	PlacementNewParamsLogisticsHandlerCertificationAuditSourceManual PlacementNewParamsLogisticsHandlerCertificationAuditSource = "manual"
	PlacementNewParamsLogisticsHandlerCertificationAuditSourceImport PlacementNewParamsLogisticsHandlerCertificationAuditSource = "import"
	PlacementNewParamsLogisticsHandlerCertificationAuditSourceSystem PlacementNewParamsLogisticsHandlerCertificationAuditSource = "system"
)

func (r PlacementNewParamsLogisticsHandlerCertificationAuditSource) IsKnown() bool {
	switch r {
	case PlacementNewParamsLogisticsHandlerCertificationAuditSourceManual, PlacementNewParamsLogisticsHandlerCertificationAuditSourceImport, PlacementNewParamsLogisticsHandlerCertificationAuditSourceSystem:
		return true
	}
	return false
}

type PlacementListParams struct {
	// Cursor from a previous response used to fetch the next page.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of placements to return.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [PlacementListParams]'s query parameters as `url.Values`.
func (r PlacementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlacementRecordEventParams struct {
	// A placement history entry; the shape depends on the event kind and carries no
	// discriminator mapping.
	PlacementEvent PlacementEventUnionParam `json:"placement_event" api:"required"`
}

func (r PlacementRecordEventParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PlacementEvent)
}
