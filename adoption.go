// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/param"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/shared"
	"github.com/tidwall/gjson"
)

// Adoption policies and applications
//
// AdoptionService contains methods and other services that help with interacting
// with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdoptionService] method instead.
type AdoptionService struct {
	Options []option.RequestOption
	// Adoption policies and applications
	Policies *AdoptionPolicyService
}

// NewAdoptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdoptionService(opts ...option.RequestOption) (r *AdoptionService) {
	r = &AdoptionService{}
	r.Options = opts
	r.Policies = NewAdoptionPolicyService(opts...)
	return
}

// Submits an adoption application for an individual or an organization applicant.
func (r *AdoptionService) New(ctx context.Context, body AdoptionNewParams, opts ...option.RequestOption) (res *Application, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "adoptions/applications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single adoption application.
func (r *AdoptionService) Get(ctx context.Context, applicationID string, opts ...option.RequestOption) (res *Application, err error) {
	opts = slices.Concat(r.Options, opts)
	if applicationID == "" {
		err = errors.New("missing required applicationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("adoptions/applications/%s", applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the decision for an application. The decision shape depends on the
// outcome and carries no discriminator mapping.
func (r *AdoptionService) GetDecision(ctx context.Context, applicationID string, opts ...option.RequestOption) (res *AdoptionGetDecisionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if applicationID == "" {
		err = errors.New("missing required applicationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("adoptions/applications/%s/decision", applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Application struct {
	ID          string               `json:"id" api:"required"`
	Applicant   ApplicationApplicant `json:"applicant" api:"required"`
	Status      ApplicationStatus    `json:"status" api:"required"`
	SubmittedAt time.Time            `json:"submittedAt" api:"required" format:"date-time"`
	// The decision for an application; the shape depends on the outcome.
	Decision        ApplicationDecision        `json:"decision"`
	Fee             shared.Money               `json:"fee"`
	History         []ApplicationHistory       `json:"history"`
	LatestRejection ApplicationLatestRejection `json:"latestRejection"`
	Pet             Pet                        `json:"pet"`
	JSON            applicationJSON            `json:"-"`
}

// applicationJSON contains the JSON metadata for the struct [Application]
type applicationJSON struct {
	ID              apijson.Field
	Applicant       apijson.Field
	Status          apijson.Field
	SubmittedAt     apijson.Field
	Decision        apijson.Field
	Fee             apijson.Field
	History         apijson.Field
	LatestRejection apijson.Field
	Pet             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Application) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationJSON) RawJSON() string {
	return r.raw
}

type ApplicationApplicant struct {
	ApplicantType ApplicationApplicantApplicantType `json:"applicant_type" api:"required"`
	Name          string                            `json:"name" api:"required"`
	Address       shared.Address                    `json:"address"`
	Age           int64                             `json:"age"`
	Email         string                            `json:"email"`
	// This field can have the runtime type of
	// [ApplicationApplicantOrganizationApplicantShelter].
	Shelter interface{}              `json:"shelter"`
	TaxID   string                   `json:"taxId"`
	JSON    applicationApplicantJSON `json:"-"`
	union   ApplicationApplicantUnion
}

// applicationApplicantJSON contains the JSON metadata for the struct
// [ApplicationApplicant]
type applicationApplicantJSON struct {
	ApplicantType apijson.Field
	Name          apijson.Field
	Address       apijson.Field
	Age           apijson.Field
	Email         apijson.Field
	Shelter       apijson.Field
	TaxID         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r applicationApplicantJSON) RawJSON() string {
	return r.raw
}

func (r *ApplicationApplicant) UnmarshalJSON(data []byte) (err error) {
	*r = ApplicationApplicant{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ApplicationApplicantUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ApplicationApplicantIndividualApplicant],
// [ApplicationApplicantOrganizationApplicant].
func (r ApplicationApplicant) AsUnion() ApplicationApplicantUnion {
	return r.union
}

// Union satisfied by [ApplicationApplicantIndividualApplicant] or
// [ApplicationApplicantOrganizationApplicant].
type ApplicationApplicantUnion interface {
	implementsApplicationApplicant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ApplicationApplicantUnion)(nil)).Elem(),
		"applicant_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ApplicationApplicantIndividualApplicant{}),
			DiscriminatorValue: "individual",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ApplicationApplicantOrganizationApplicant{}),
			DiscriminatorValue: "organization",
		},
	)
}

type ApplicationApplicantIndividualApplicant struct {
	ApplicantType ApplicationApplicantIndividualApplicantApplicantType `json:"applicant_type" api:"required"`
	Name          string                                               `json:"name" api:"required"`
	Address       shared.Address                                       `json:"address"`
	Age           int64                                                `json:"age"`
	Email         string                                               `json:"email"`
	JSON          applicationApplicantIndividualApplicantJSON          `json:"-"`
}

// applicationApplicantIndividualApplicantJSON contains the JSON metadata for the
// struct [ApplicationApplicantIndividualApplicant]
type applicationApplicantIndividualApplicantJSON struct {
	ApplicantType apijson.Field
	Name          apijson.Field
	Address       apijson.Field
	Age           apijson.Field
	Email         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ApplicationApplicantIndividualApplicant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationApplicantIndividualApplicantJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationApplicantIndividualApplicant) implementsApplicationApplicant() {}

type ApplicationApplicantIndividualApplicantApplicantType string

const (
	ApplicationApplicantIndividualApplicantApplicantTypeIndividual ApplicationApplicantIndividualApplicantApplicantType = "individual"
)

func (r ApplicationApplicantIndividualApplicantApplicantType) IsKnown() bool {
	switch r {
	case ApplicationApplicantIndividualApplicantApplicantTypeIndividual:
		return true
	}
	return false
}

type ApplicationApplicantOrganizationApplicant struct {
	ApplicantType ApplicationApplicantOrganizationApplicantApplicantType `json:"applicant_type" api:"required"`
	Name          string                                                 `json:"name" api:"required"`
	Shelter       ApplicationApplicantOrganizationApplicantShelter       `json:"shelter"`
	TaxID         string                                                 `json:"taxId"`
	JSON          applicationApplicantOrganizationApplicantJSON          `json:"-"`
}

// applicationApplicantOrganizationApplicantJSON contains the JSON metadata for the
// struct [ApplicationApplicantOrganizationApplicant]
type applicationApplicantOrganizationApplicantJSON struct {
	ApplicantType apijson.Field
	Name          apijson.Field
	Shelter       apijson.Field
	TaxID         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ApplicationApplicantOrganizationApplicant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationApplicantOrganizationApplicantJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationApplicantOrganizationApplicant) implementsApplicationApplicant() {}

type ApplicationApplicantOrganizationApplicantApplicantType string

const (
	ApplicationApplicantOrganizationApplicantApplicantTypeOrganization ApplicationApplicantOrganizationApplicantApplicantType = "organization"
)

func (r ApplicationApplicantOrganizationApplicantApplicantType) IsKnown() bool {
	switch r {
	case ApplicationApplicantOrganizationApplicantApplicantTypeOrganization:
		return true
	}
	return false
}

type ApplicationApplicantOrganizationApplicantShelter struct {
	Name            string                                                   `json:"name" api:"required"`
	Address         *shared.Address                                          `json:"address"`
	Contact         ApplicationApplicantOrganizationApplicantShelterContact  `json:"contact"`
	Location        ApplicationApplicantOrganizationApplicantShelterLocation `json:"location"`
	RelatedAddress  *shared.Address                                          `json:"relatedAddress"`
	RelatedCategory interface{}                                              `json:"relatedCategory"`
	RelatedCustomer interface{}                                              `json:"relatedCustomer"`
	RelatedMoney    *shared.Money                                            `json:"relatedMoney"`
	RelatedOrder    interface{}                                              `json:"relatedOrder"`
	RelatedPet      *Pet                                                     `json:"relatedPet"`
	RelatedTag      interface{}                                              `json:"relatedTag"`
	RelatedUser     interface{}                                              `json:"relatedUser"`
	JSON            applicationApplicantOrganizationApplicantShelterJSON     `json:"-"`
}

// applicationApplicantOrganizationApplicantShelterJSON contains the JSON metadata
// for the struct [ApplicationApplicantOrganizationApplicantShelter]
type applicationApplicantOrganizationApplicantShelterJSON struct {
	Name            apijson.Field
	Address         apijson.Field
	Contact         apijson.Field
	Location        apijson.Field
	RelatedAddress  apijson.Field
	RelatedCategory apijson.Field
	RelatedCustomer apijson.Field
	RelatedMoney    apijson.Field
	RelatedOrder    apijson.Field
	RelatedPet      apijson.Field
	RelatedTag      apijson.Field
	RelatedUser     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ApplicationApplicantOrganizationApplicantShelter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationApplicantOrganizationApplicantShelterJSON) RawJSON() string {
	return r.raw
}

type ApplicationApplicantOrganizationApplicantShelterContact struct {
	Primary ApplicationApplicantOrganizationApplicantShelterContactPrimary `json:"primary"`
	JSON    applicationApplicantOrganizationApplicantShelterContactJSON    `json:"-"`
}

// applicationApplicantOrganizationApplicantShelterContactJSON contains the JSON
// metadata for the struct
// [ApplicationApplicantOrganizationApplicantShelterContact]
type applicationApplicantOrganizationApplicantShelterContactJSON struct {
	Primary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationApplicantOrganizationApplicantShelterContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationApplicantOrganizationApplicantShelterContactJSON) RawJSON() string {
	return r.raw
}

type ApplicationApplicantOrganizationApplicantShelterContactPrimary struct {
	Availability ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailability `json:"availability"`
	Name         string                                                                     `json:"name"`
	Phone        string                                                                     `json:"phone"`
	JSON         applicationApplicantOrganizationApplicantShelterContactPrimaryJSON         `json:"-"`
}

// applicationApplicantOrganizationApplicantShelterContactPrimaryJSON contains the
// JSON metadata for the struct
// [ApplicationApplicantOrganizationApplicantShelterContactPrimary]
type applicationApplicantOrganizationApplicantShelterContactPrimaryJSON struct {
	Availability apijson.Field
	Name         apijson.Field
	Phone        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ApplicationApplicantOrganizationApplicantShelterContactPrimary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationApplicantOrganizationApplicantShelterContactPrimaryJSON) RawJSON() string {
	return r.raw
}

type ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailability struct {
	Weekday ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekday `json:"weekday"`
	Weekend ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekend `json:"weekend" api:"nullable"`
	JSON    applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityJSON    `json:"-"`
}

// applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityJSON
// contains the JSON metadata for the struct
// [ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailability]
type applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityJSON struct {
	Weekday     apijson.Field
	Weekend     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityJSON) RawJSON() string {
	return r.raw
}

type ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekday struct {
	Close string                                                                                `json:"close"`
	Open  string                                                                                `json:"open"`
	JSON  applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekdayJSON `json:"-"`
}

// applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekdayJSON
// contains the JSON metadata for the struct
// [ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekday]
type applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekdayJSON struct {
	Close       apijson.Field
	Open        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekday) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekdayJSON) RawJSON() string {
	return r.raw
}

type ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekend struct {
	Close string                                                                                `json:"close"`
	Open  string                                                                                `json:"open"`
	JSON  applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekendJSON `json:"-"`
}

// applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekendJSON
// contains the JSON metadata for the struct
// [ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekend]
type applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekendJSON struct {
	Close       apijson.Field
	Open        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationApplicantOrganizationApplicantShelterContactPrimaryAvailabilityWeekendJSON) RawJSON() string {
	return r.raw
}

type ApplicationApplicantOrganizationApplicantShelterLocation struct {
	Latitude  float64                                                      `json:"latitude" api:"required"`
	Longitude float64                                                      `json:"longitude" api:"required"`
	JSON      applicationApplicantOrganizationApplicantShelterLocationJSON `json:"-"`
}

// applicationApplicantOrganizationApplicantShelterLocationJSON contains the JSON
// metadata for the struct
// [ApplicationApplicantOrganizationApplicantShelterLocation]
type applicationApplicantOrganizationApplicantShelterLocationJSON struct {
	Latitude    apijson.Field
	Longitude   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationApplicantOrganizationApplicantShelterLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationApplicantOrganizationApplicantShelterLocationJSON) RawJSON() string {
	return r.raw
}

type ApplicationApplicantApplicantType string

const (
	ApplicationApplicantApplicantTypeIndividual   ApplicationApplicantApplicantType = "individual"
	ApplicationApplicantApplicantTypeOrganization ApplicationApplicantApplicantType = "organization"
)

func (r ApplicationApplicantApplicantType) IsKnown() bool {
	switch r {
	case ApplicationApplicantApplicantTypeIndividual, ApplicationApplicantApplicantTypeOrganization:
		return true
	}
	return false
}

type ApplicationStatus string

const (
	ApplicationStatusReceived    ApplicationStatus = "received"
	ApplicationStatusUnderReview ApplicationStatus = "under_review"
	ApplicationStatusApproved    ApplicationStatus = "approved"
	ApplicationStatusRejected    ApplicationStatus = "rejected"
)

func (r ApplicationStatus) IsKnown() bool {
	switch r {
	case ApplicationStatusReceived, ApplicationStatusUnderReview, ApplicationStatusApproved, ApplicationStatusRejected:
		return true
	}
	return false
}

// The decision for an application; the shape depends on the outcome.
type ApplicationDecision struct {
	Outcome        ApplicationDecisionOutcome `json:"outcome" api:"required"`
	AppealDeadline time.Time                  `json:"appealDeadline" api:"nullable" format:"date-time"`
	ApprovedAt     time.Time                  `json:"approvedAt" format:"date-time"`
	// This field can have the runtime type of [[]string].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of
	// [ApplicationDecisionDecisionEscalatedEscalatedTo].
	EscalatedTo interface{}                    `json:"escalatedTo"`
	Reason      ApplicationDecisionReason      `json:"reason"`
	ReviewAfter time.Time                      `json:"reviewAfter" format:"date-time"`
	WithdrawnAt time.Time                      `json:"withdrawnAt" format:"date-time"`
	WithdrawnBy ApplicationDecisionWithdrawnBy `json:"withdrawnBy"`
	JSON        applicationDecisionJSON        `json:"-"`
	union       ApplicationDecisionUnion
}

// applicationDecisionJSON contains the JSON metadata for the struct
// [ApplicationDecision]
type applicationDecisionJSON struct {
	Outcome        apijson.Field
	AppealDeadline apijson.Field
	ApprovedAt     apijson.Field
	Conditions     apijson.Field
	EscalatedTo    apijson.Field
	Reason         apijson.Field
	ReviewAfter    apijson.Field
	WithdrawnAt    apijson.Field
	WithdrawnBy    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r applicationDecisionJSON) RawJSON() string {
	return r.raw
}

func (r *ApplicationDecision) UnmarshalJSON(data []byte) (err error) {
	*r = ApplicationDecision{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ApplicationDecisionUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ApplicationDecisionDecisionApproved],
// [ApplicationDecisionDecisionRejected], [ApplicationDecisionDecisionEscalated],
// [ApplicationDecisionDecisionWithdrawn].
func (r ApplicationDecision) AsUnion() ApplicationDecisionUnion {
	return r.union
}

// The decision for an application; the shape depends on the outcome.
//
// Union satisfied by [ApplicationDecisionDecisionApproved],
// [ApplicationDecisionDecisionRejected], [ApplicationDecisionDecisionEscalated] or
// [ApplicationDecisionDecisionWithdrawn].
type ApplicationDecisionUnion interface {
	implementsApplicationDecision()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ApplicationDecisionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationDecisionDecisionApproved{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationDecisionDecisionRejected{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationDecisionDecisionEscalated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ApplicationDecisionDecisionWithdrawn{}),
		},
	)
}

type ApplicationDecisionDecisionApproved struct {
	ApprovedAt time.Time                                  `json:"approvedAt" api:"required" format:"date-time"`
	Outcome    ApplicationDecisionDecisionApprovedOutcome `json:"outcome" api:"required"`
	Conditions []string                                   `json:"conditions"`
	JSON       applicationDecisionDecisionApprovedJSON    `json:"-"`
}

// applicationDecisionDecisionApprovedJSON contains the JSON metadata for the
// struct [ApplicationDecisionDecisionApproved]
type applicationDecisionDecisionApprovedJSON struct {
	ApprovedAt  apijson.Field
	Outcome     apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationDecisionDecisionApproved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDecisionDecisionApprovedJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationDecisionDecisionApproved) implementsApplicationDecision() {}

type ApplicationDecisionDecisionApprovedOutcome string

const (
	ApplicationDecisionDecisionApprovedOutcomeApproved ApplicationDecisionDecisionApprovedOutcome = "approved"
)

func (r ApplicationDecisionDecisionApprovedOutcome) IsKnown() bool {
	switch r {
	case ApplicationDecisionDecisionApprovedOutcomeApproved:
		return true
	}
	return false
}

type ApplicationDecisionDecisionRejected struct {
	Outcome        ApplicationDecisionDecisionRejectedOutcome `json:"outcome" api:"required"`
	Reason         ApplicationDecisionDecisionRejectedReason  `json:"reason" api:"required"`
	AppealDeadline time.Time                                  `json:"appealDeadline" api:"nullable" format:"date-time"`
	JSON           applicationDecisionDecisionRejectedJSON    `json:"-"`
}

// applicationDecisionDecisionRejectedJSON contains the JSON metadata for the
// struct [ApplicationDecisionDecisionRejected]
type applicationDecisionDecisionRejectedJSON struct {
	Outcome        apijson.Field
	Reason         apijson.Field
	AppealDeadline apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ApplicationDecisionDecisionRejected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDecisionDecisionRejectedJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationDecisionDecisionRejected) implementsApplicationDecision() {}

type ApplicationDecisionDecisionRejectedOutcome string

const (
	ApplicationDecisionDecisionRejectedOutcomeRejected ApplicationDecisionDecisionRejectedOutcome = "rejected"
)

func (r ApplicationDecisionDecisionRejectedOutcome) IsKnown() bool {
	switch r {
	case ApplicationDecisionDecisionRejectedOutcomeRejected:
		return true
	}
	return false
}

type ApplicationDecisionDecisionRejectedReason string

const (
	ApplicationDecisionDecisionRejectedReasonPolicy     ApplicationDecisionDecisionRejectedReason = "policy"
	ApplicationDecisionDecisionRejectedReasonCapacity   ApplicationDecisionDecisionRejectedReason = "capacity"
	ApplicationDecisionDecisionRejectedReasonIncomplete ApplicationDecisionDecisionRejectedReason = "incomplete"
)

func (r ApplicationDecisionDecisionRejectedReason) IsKnown() bool {
	switch r {
	case ApplicationDecisionDecisionRejectedReasonPolicy, ApplicationDecisionDecisionRejectedReasonCapacity, ApplicationDecisionDecisionRejectedReasonIncomplete:
		return true
	}
	return false
}

type ApplicationDecisionDecisionEscalated struct {
	EscalatedTo ApplicationDecisionDecisionEscalatedEscalatedTo `json:"escalatedTo" api:"required"`
	Outcome     ApplicationDecisionDecisionEscalatedOutcome     `json:"outcome" api:"required"`
	ReviewAfter time.Time                                       `json:"reviewAfter" format:"date-time"`
	JSON        applicationDecisionDecisionEscalatedJSON        `json:"-"`
}

// applicationDecisionDecisionEscalatedJSON contains the JSON metadata for the
// struct [ApplicationDecisionDecisionEscalated]
type applicationDecisionDecisionEscalatedJSON struct {
	EscalatedTo apijson.Field
	Outcome     apijson.Field
	ReviewAfter apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationDecisionDecisionEscalated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDecisionDecisionEscalatedJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationDecisionDecisionEscalated) implementsApplicationDecision() {}

type ApplicationDecisionDecisionEscalatedEscalatedTo struct {
	Team    string                                                 `json:"team" api:"required"`
	Contact ApplicationDecisionDecisionEscalatedEscalatedToContact `json:"contact"`
	JSON    applicationDecisionDecisionEscalatedEscalatedToJSON    `json:"-"`
}

// applicationDecisionDecisionEscalatedEscalatedToJSON contains the JSON metadata
// for the struct [ApplicationDecisionDecisionEscalatedEscalatedTo]
type applicationDecisionDecisionEscalatedEscalatedToJSON struct {
	Team        apijson.Field
	Contact     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationDecisionDecisionEscalatedEscalatedTo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDecisionDecisionEscalatedEscalatedToJSON) RawJSON() string {
	return r.raw
}

type ApplicationDecisionDecisionEscalatedEscalatedToContact struct {
	Email string                                                     `json:"email"`
	Hours string                                                     `json:"hours"`
	JSON  applicationDecisionDecisionEscalatedEscalatedToContactJSON `json:"-"`
}

// applicationDecisionDecisionEscalatedEscalatedToContactJSON contains the JSON
// metadata for the struct [ApplicationDecisionDecisionEscalatedEscalatedToContact]
type applicationDecisionDecisionEscalatedEscalatedToContactJSON struct {
	Email       apijson.Field
	Hours       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationDecisionDecisionEscalatedEscalatedToContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDecisionDecisionEscalatedEscalatedToContactJSON) RawJSON() string {
	return r.raw
}

type ApplicationDecisionDecisionEscalatedOutcome string

const (
	ApplicationDecisionDecisionEscalatedOutcomeEscalated ApplicationDecisionDecisionEscalatedOutcome = "escalated"
)

func (r ApplicationDecisionDecisionEscalatedOutcome) IsKnown() bool {
	switch r {
	case ApplicationDecisionDecisionEscalatedOutcomeEscalated:
		return true
	}
	return false
}

// The applicant or shelter withdrew before a decision was finalized.
type ApplicationDecisionDecisionWithdrawn struct {
	Outcome     ApplicationDecisionDecisionWithdrawnOutcome     `json:"outcome" api:"required"`
	WithdrawnBy ApplicationDecisionDecisionWithdrawnWithdrawnBy `json:"withdrawnBy" api:"required"`
	WithdrawnAt time.Time                                       `json:"withdrawnAt" format:"date-time"`
	JSON        applicationDecisionDecisionWithdrawnJSON        `json:"-"`
}

// applicationDecisionDecisionWithdrawnJSON contains the JSON metadata for the
// struct [ApplicationDecisionDecisionWithdrawn]
type applicationDecisionDecisionWithdrawnJSON struct {
	Outcome     apijson.Field
	WithdrawnBy apijson.Field
	WithdrawnAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationDecisionDecisionWithdrawn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationDecisionDecisionWithdrawnJSON) RawJSON() string {
	return r.raw
}

func (r ApplicationDecisionDecisionWithdrawn) implementsApplicationDecision() {}

type ApplicationDecisionDecisionWithdrawnOutcome string

const (
	ApplicationDecisionDecisionWithdrawnOutcomeWithdrawn ApplicationDecisionDecisionWithdrawnOutcome = "withdrawn"
	ApplicationDecisionDecisionWithdrawnOutcomeExpired   ApplicationDecisionDecisionWithdrawnOutcome = "expired"
)

func (r ApplicationDecisionDecisionWithdrawnOutcome) IsKnown() bool {
	switch r {
	case ApplicationDecisionDecisionWithdrawnOutcomeWithdrawn, ApplicationDecisionDecisionWithdrawnOutcomeExpired:
		return true
	}
	return false
}

type ApplicationDecisionDecisionWithdrawnWithdrawnBy string

const (
	ApplicationDecisionDecisionWithdrawnWithdrawnByApplicant ApplicationDecisionDecisionWithdrawnWithdrawnBy = "applicant"
	ApplicationDecisionDecisionWithdrawnWithdrawnByShelter   ApplicationDecisionDecisionWithdrawnWithdrawnBy = "shelter"
	ApplicationDecisionDecisionWithdrawnWithdrawnBySystem    ApplicationDecisionDecisionWithdrawnWithdrawnBy = "system"
)

func (r ApplicationDecisionDecisionWithdrawnWithdrawnBy) IsKnown() bool {
	switch r {
	case ApplicationDecisionDecisionWithdrawnWithdrawnByApplicant, ApplicationDecisionDecisionWithdrawnWithdrawnByShelter, ApplicationDecisionDecisionWithdrawnWithdrawnBySystem:
		return true
	}
	return false
}

type ApplicationDecisionOutcome string

const (
	ApplicationDecisionOutcomeApproved  ApplicationDecisionOutcome = "approved"
	ApplicationDecisionOutcomeRejected  ApplicationDecisionOutcome = "rejected"
	ApplicationDecisionOutcomeEscalated ApplicationDecisionOutcome = "escalated"
	ApplicationDecisionOutcomeWithdrawn ApplicationDecisionOutcome = "withdrawn"
	ApplicationDecisionOutcomeExpired   ApplicationDecisionOutcome = "expired"
)

func (r ApplicationDecisionOutcome) IsKnown() bool {
	switch r {
	case ApplicationDecisionOutcomeApproved, ApplicationDecisionOutcomeRejected, ApplicationDecisionOutcomeEscalated, ApplicationDecisionOutcomeWithdrawn, ApplicationDecisionOutcomeExpired:
		return true
	}
	return false
}

type ApplicationDecisionReason string

const (
	ApplicationDecisionReasonPolicy     ApplicationDecisionReason = "policy"
	ApplicationDecisionReasonCapacity   ApplicationDecisionReason = "capacity"
	ApplicationDecisionReasonIncomplete ApplicationDecisionReason = "incomplete"
)

func (r ApplicationDecisionReason) IsKnown() bool {
	switch r {
	case ApplicationDecisionReasonPolicy, ApplicationDecisionReasonCapacity, ApplicationDecisionReasonIncomplete:
		return true
	}
	return false
}

type ApplicationDecisionWithdrawnBy string

const (
	ApplicationDecisionWithdrawnByApplicant ApplicationDecisionWithdrawnBy = "applicant"
	ApplicationDecisionWithdrawnByShelter   ApplicationDecisionWithdrawnBy = "shelter"
	ApplicationDecisionWithdrawnBySystem    ApplicationDecisionWithdrawnBy = "system"
)

func (r ApplicationDecisionWithdrawnBy) IsKnown() bool {
	switch r {
	case ApplicationDecisionWithdrawnByApplicant, ApplicationDecisionWithdrawnByShelter, ApplicationDecisionWithdrawnBySystem:
		return true
	}
	return false
}

type ApplicationHistory struct {
	ChangedAt time.Time                `json:"changedAt" api:"required" format:"date-time"`
	Status    ApplicationHistoryStatus `json:"status" api:"required"`
	Note      string                   `json:"note" api:"nullable"`
	JSON      applicationHistoryJSON   `json:"-"`
}

// applicationHistoryJSON contains the JSON metadata for the struct
// [ApplicationHistory]
type applicationHistoryJSON struct {
	ChangedAt   apijson.Field
	Status      apijson.Field
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationHistoryJSON) RawJSON() string {
	return r.raw
}

type ApplicationHistoryStatus string

const (
	ApplicationHistoryStatusReceived    ApplicationHistoryStatus = "received"
	ApplicationHistoryStatusUnderReview ApplicationHistoryStatus = "under_review"
	ApplicationHistoryStatusApproved    ApplicationHistoryStatus = "approved"
	ApplicationHistoryStatusRejected    ApplicationHistoryStatus = "rejected"
)

func (r ApplicationHistoryStatus) IsKnown() bool {
	switch r {
	case ApplicationHistoryStatusReceived, ApplicationHistoryStatusUnderReview, ApplicationHistoryStatusApproved, ApplicationHistoryStatusRejected:
		return true
	}
	return false
}

type ApplicationLatestRejection struct {
	Outcome        ApplicationLatestRejectionOutcome `json:"outcome" api:"required"`
	Reason         ApplicationLatestRejectionReason  `json:"reason" api:"required"`
	AppealDeadline time.Time                         `json:"appealDeadline" api:"nullable" format:"date-time"`
	JSON           applicationLatestRejectionJSON    `json:"-"`
}

// applicationLatestRejectionJSON contains the JSON metadata for the struct
// [ApplicationLatestRejection]
type applicationLatestRejectionJSON struct {
	Outcome        apijson.Field
	Reason         apijson.Field
	AppealDeadline apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ApplicationLatestRejection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationLatestRejectionJSON) RawJSON() string {
	return r.raw
}

type ApplicationLatestRejectionOutcome string

const (
	ApplicationLatestRejectionOutcomeRejected ApplicationLatestRejectionOutcome = "rejected"
)

func (r ApplicationLatestRejectionOutcome) IsKnown() bool {
	switch r {
	case ApplicationLatestRejectionOutcomeRejected:
		return true
	}
	return false
}

type ApplicationLatestRejectionReason string

const (
	ApplicationLatestRejectionReasonPolicy     ApplicationLatestRejectionReason = "policy"
	ApplicationLatestRejectionReasonCapacity   ApplicationLatestRejectionReason = "capacity"
	ApplicationLatestRejectionReasonIncomplete ApplicationLatestRejectionReason = "incomplete"
)

func (r ApplicationLatestRejectionReason) IsKnown() bool {
	switch r {
	case ApplicationLatestRejectionReasonPolicy, ApplicationLatestRejectionReasonCapacity, ApplicationLatestRejectionReasonIncomplete:
		return true
	}
	return false
}

// The decision for an application; the shape depends on the outcome.
type AdoptionGetDecisionResponse struct {
	Outcome        AdoptionGetDecisionResponseOutcome `json:"outcome" api:"required"`
	AppealDeadline time.Time                          `json:"appealDeadline" api:"nullable" format:"date-time"`
	ApprovedAt     time.Time                          `json:"approvedAt" format:"date-time"`
	// This field can have the runtime type of [[]string].
	Conditions interface{} `json:"conditions"`
	// This field can have the runtime type of
	// [AdoptionGetDecisionResponseDecisionEscalatedEscalatedTo].
	EscalatedTo interface{}                            `json:"escalatedTo"`
	Reason      AdoptionGetDecisionResponseReason      `json:"reason"`
	ReviewAfter time.Time                              `json:"reviewAfter" format:"date-time"`
	WithdrawnAt time.Time                              `json:"withdrawnAt" format:"date-time"`
	WithdrawnBy AdoptionGetDecisionResponseWithdrawnBy `json:"withdrawnBy"`
	JSON        adoptionGetDecisionResponseJSON        `json:"-"`
	union       AdoptionGetDecisionResponseUnion
}

// adoptionGetDecisionResponseJSON contains the JSON metadata for the struct
// [AdoptionGetDecisionResponse]
type adoptionGetDecisionResponseJSON struct {
	Outcome        apijson.Field
	AppealDeadline apijson.Field
	ApprovedAt     apijson.Field
	Conditions     apijson.Field
	EscalatedTo    apijson.Field
	Reason         apijson.Field
	ReviewAfter    apijson.Field
	WithdrawnAt    apijson.Field
	WithdrawnBy    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r adoptionGetDecisionResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AdoptionGetDecisionResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AdoptionGetDecisionResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AdoptionGetDecisionResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AdoptionGetDecisionResponseDecisionApproved],
// [AdoptionGetDecisionResponseDecisionRejected],
// [AdoptionGetDecisionResponseDecisionEscalated],
// [AdoptionGetDecisionResponseDecisionWithdrawn].
func (r AdoptionGetDecisionResponse) AsUnion() AdoptionGetDecisionResponseUnion {
	return r.union
}

// The decision for an application; the shape depends on the outcome.
//
// Union satisfied by [AdoptionGetDecisionResponseDecisionApproved],
// [AdoptionGetDecisionResponseDecisionRejected],
// [AdoptionGetDecisionResponseDecisionEscalated] or
// [AdoptionGetDecisionResponseDecisionWithdrawn].
type AdoptionGetDecisionResponseUnion interface {
	implementsAdoptionGetDecisionResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AdoptionGetDecisionResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AdoptionGetDecisionResponseDecisionApproved{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AdoptionGetDecisionResponseDecisionRejected{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AdoptionGetDecisionResponseDecisionEscalated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AdoptionGetDecisionResponseDecisionWithdrawn{}),
		},
	)
}

type AdoptionGetDecisionResponseDecisionApproved struct {
	ApprovedAt time.Time                                          `json:"approvedAt" api:"required" format:"date-time"`
	Outcome    AdoptionGetDecisionResponseDecisionApprovedOutcome `json:"outcome" api:"required"`
	Conditions []string                                           `json:"conditions"`
	JSON       adoptionGetDecisionResponseDecisionApprovedJSON    `json:"-"`
}

// adoptionGetDecisionResponseDecisionApprovedJSON contains the JSON metadata for
// the struct [AdoptionGetDecisionResponseDecisionApproved]
type adoptionGetDecisionResponseDecisionApprovedJSON struct {
	ApprovedAt  apijson.Field
	Outcome     apijson.Field
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdoptionGetDecisionResponseDecisionApproved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adoptionGetDecisionResponseDecisionApprovedJSON) RawJSON() string {
	return r.raw
}

func (r AdoptionGetDecisionResponseDecisionApproved) implementsAdoptionGetDecisionResponse() {}

type AdoptionGetDecisionResponseDecisionApprovedOutcome string

const (
	AdoptionGetDecisionResponseDecisionApprovedOutcomeApproved AdoptionGetDecisionResponseDecisionApprovedOutcome = "approved"
)

func (r AdoptionGetDecisionResponseDecisionApprovedOutcome) IsKnown() bool {
	switch r {
	case AdoptionGetDecisionResponseDecisionApprovedOutcomeApproved:
		return true
	}
	return false
}

type AdoptionGetDecisionResponseDecisionRejected struct {
	Outcome        AdoptionGetDecisionResponseDecisionRejectedOutcome `json:"outcome" api:"required"`
	Reason         AdoptionGetDecisionResponseDecisionRejectedReason  `json:"reason" api:"required"`
	AppealDeadline time.Time                                          `json:"appealDeadline" api:"nullable" format:"date-time"`
	JSON           adoptionGetDecisionResponseDecisionRejectedJSON    `json:"-"`
}

// adoptionGetDecisionResponseDecisionRejectedJSON contains the JSON metadata for
// the struct [AdoptionGetDecisionResponseDecisionRejected]
type adoptionGetDecisionResponseDecisionRejectedJSON struct {
	Outcome        apijson.Field
	Reason         apijson.Field
	AppealDeadline apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AdoptionGetDecisionResponseDecisionRejected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adoptionGetDecisionResponseDecisionRejectedJSON) RawJSON() string {
	return r.raw
}

func (r AdoptionGetDecisionResponseDecisionRejected) implementsAdoptionGetDecisionResponse() {}

type AdoptionGetDecisionResponseDecisionRejectedOutcome string

const (
	AdoptionGetDecisionResponseDecisionRejectedOutcomeRejected AdoptionGetDecisionResponseDecisionRejectedOutcome = "rejected"
)

func (r AdoptionGetDecisionResponseDecisionRejectedOutcome) IsKnown() bool {
	switch r {
	case AdoptionGetDecisionResponseDecisionRejectedOutcomeRejected:
		return true
	}
	return false
}

type AdoptionGetDecisionResponseDecisionRejectedReason string

const (
	AdoptionGetDecisionResponseDecisionRejectedReasonPolicy     AdoptionGetDecisionResponseDecisionRejectedReason = "policy"
	AdoptionGetDecisionResponseDecisionRejectedReasonCapacity   AdoptionGetDecisionResponseDecisionRejectedReason = "capacity"
	AdoptionGetDecisionResponseDecisionRejectedReasonIncomplete AdoptionGetDecisionResponseDecisionRejectedReason = "incomplete"
)

func (r AdoptionGetDecisionResponseDecisionRejectedReason) IsKnown() bool {
	switch r {
	case AdoptionGetDecisionResponseDecisionRejectedReasonPolicy, AdoptionGetDecisionResponseDecisionRejectedReasonCapacity, AdoptionGetDecisionResponseDecisionRejectedReasonIncomplete:
		return true
	}
	return false
}

type AdoptionGetDecisionResponseDecisionEscalated struct {
	EscalatedTo AdoptionGetDecisionResponseDecisionEscalatedEscalatedTo `json:"escalatedTo" api:"required"`
	Outcome     AdoptionGetDecisionResponseDecisionEscalatedOutcome     `json:"outcome" api:"required"`
	ReviewAfter time.Time                                               `json:"reviewAfter" format:"date-time"`
	JSON        adoptionGetDecisionResponseDecisionEscalatedJSON        `json:"-"`
}

// adoptionGetDecisionResponseDecisionEscalatedJSON contains the JSON metadata for
// the struct [AdoptionGetDecisionResponseDecisionEscalated]
type adoptionGetDecisionResponseDecisionEscalatedJSON struct {
	EscalatedTo apijson.Field
	Outcome     apijson.Field
	ReviewAfter apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdoptionGetDecisionResponseDecisionEscalated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adoptionGetDecisionResponseDecisionEscalatedJSON) RawJSON() string {
	return r.raw
}

func (r AdoptionGetDecisionResponseDecisionEscalated) implementsAdoptionGetDecisionResponse() {}

type AdoptionGetDecisionResponseDecisionEscalatedEscalatedTo struct {
	Team    string                                                         `json:"team" api:"required"`
	Contact AdoptionGetDecisionResponseDecisionEscalatedEscalatedToContact `json:"contact"`
	JSON    adoptionGetDecisionResponseDecisionEscalatedEscalatedToJSON    `json:"-"`
}

// adoptionGetDecisionResponseDecisionEscalatedEscalatedToJSON contains the JSON
// metadata for the struct
// [AdoptionGetDecisionResponseDecisionEscalatedEscalatedTo]
type adoptionGetDecisionResponseDecisionEscalatedEscalatedToJSON struct {
	Team        apijson.Field
	Contact     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdoptionGetDecisionResponseDecisionEscalatedEscalatedTo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adoptionGetDecisionResponseDecisionEscalatedEscalatedToJSON) RawJSON() string {
	return r.raw
}

type AdoptionGetDecisionResponseDecisionEscalatedEscalatedToContact struct {
	Email string                                                             `json:"email"`
	Hours string                                                             `json:"hours"`
	JSON  adoptionGetDecisionResponseDecisionEscalatedEscalatedToContactJSON `json:"-"`
}

// adoptionGetDecisionResponseDecisionEscalatedEscalatedToContactJSON contains the
// JSON metadata for the struct
// [AdoptionGetDecisionResponseDecisionEscalatedEscalatedToContact]
type adoptionGetDecisionResponseDecisionEscalatedEscalatedToContactJSON struct {
	Email       apijson.Field
	Hours       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdoptionGetDecisionResponseDecisionEscalatedEscalatedToContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adoptionGetDecisionResponseDecisionEscalatedEscalatedToContactJSON) RawJSON() string {
	return r.raw
}

type AdoptionGetDecisionResponseDecisionEscalatedOutcome string

const (
	AdoptionGetDecisionResponseDecisionEscalatedOutcomeEscalated AdoptionGetDecisionResponseDecisionEscalatedOutcome = "escalated"
)

func (r AdoptionGetDecisionResponseDecisionEscalatedOutcome) IsKnown() bool {
	switch r {
	case AdoptionGetDecisionResponseDecisionEscalatedOutcomeEscalated:
		return true
	}
	return false
}

// The applicant or shelter withdrew before a decision was finalized.
type AdoptionGetDecisionResponseDecisionWithdrawn struct {
	Outcome     AdoptionGetDecisionResponseDecisionWithdrawnOutcome     `json:"outcome" api:"required"`
	WithdrawnBy AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnBy `json:"withdrawnBy" api:"required"`
	WithdrawnAt time.Time                                               `json:"withdrawnAt" format:"date-time"`
	JSON        adoptionGetDecisionResponseDecisionWithdrawnJSON        `json:"-"`
}

// adoptionGetDecisionResponseDecisionWithdrawnJSON contains the JSON metadata for
// the struct [AdoptionGetDecisionResponseDecisionWithdrawn]
type adoptionGetDecisionResponseDecisionWithdrawnJSON struct {
	Outcome     apijson.Field
	WithdrawnBy apijson.Field
	WithdrawnAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdoptionGetDecisionResponseDecisionWithdrawn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adoptionGetDecisionResponseDecisionWithdrawnJSON) RawJSON() string {
	return r.raw
}

func (r AdoptionGetDecisionResponseDecisionWithdrawn) implementsAdoptionGetDecisionResponse() {}

type AdoptionGetDecisionResponseDecisionWithdrawnOutcome string

const (
	AdoptionGetDecisionResponseDecisionWithdrawnOutcomeWithdrawn AdoptionGetDecisionResponseDecisionWithdrawnOutcome = "withdrawn"
	AdoptionGetDecisionResponseDecisionWithdrawnOutcomeExpired   AdoptionGetDecisionResponseDecisionWithdrawnOutcome = "expired"
)

func (r AdoptionGetDecisionResponseDecisionWithdrawnOutcome) IsKnown() bool {
	switch r {
	case AdoptionGetDecisionResponseDecisionWithdrawnOutcomeWithdrawn, AdoptionGetDecisionResponseDecisionWithdrawnOutcomeExpired:
		return true
	}
	return false
}

type AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnBy string

const (
	AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnByApplicant AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnBy = "applicant"
	AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnByShelter   AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnBy = "shelter"
	AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnBySystem    AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnBy = "system"
)

func (r AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnBy) IsKnown() bool {
	switch r {
	case AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnByApplicant, AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnByShelter, AdoptionGetDecisionResponseDecisionWithdrawnWithdrawnBySystem:
		return true
	}
	return false
}

type AdoptionGetDecisionResponseOutcome string

const (
	AdoptionGetDecisionResponseOutcomeApproved  AdoptionGetDecisionResponseOutcome = "approved"
	AdoptionGetDecisionResponseOutcomeRejected  AdoptionGetDecisionResponseOutcome = "rejected"
	AdoptionGetDecisionResponseOutcomeEscalated AdoptionGetDecisionResponseOutcome = "escalated"
	AdoptionGetDecisionResponseOutcomeWithdrawn AdoptionGetDecisionResponseOutcome = "withdrawn"
	AdoptionGetDecisionResponseOutcomeExpired   AdoptionGetDecisionResponseOutcome = "expired"
)

func (r AdoptionGetDecisionResponseOutcome) IsKnown() bool {
	switch r {
	case AdoptionGetDecisionResponseOutcomeApproved, AdoptionGetDecisionResponseOutcomeRejected, AdoptionGetDecisionResponseOutcomeEscalated, AdoptionGetDecisionResponseOutcomeWithdrawn, AdoptionGetDecisionResponseOutcomeExpired:
		return true
	}
	return false
}

type AdoptionGetDecisionResponseReason string

const (
	AdoptionGetDecisionResponseReasonPolicy     AdoptionGetDecisionResponseReason = "policy"
	AdoptionGetDecisionResponseReasonCapacity   AdoptionGetDecisionResponseReason = "capacity"
	AdoptionGetDecisionResponseReasonIncomplete AdoptionGetDecisionResponseReason = "incomplete"
)

func (r AdoptionGetDecisionResponseReason) IsKnown() bool {
	switch r {
	case AdoptionGetDecisionResponseReasonPolicy, AdoptionGetDecisionResponseReasonCapacity, AdoptionGetDecisionResponseReasonIncomplete:
		return true
	}
	return false
}

type AdoptionGetDecisionResponseWithdrawnBy string

const (
	AdoptionGetDecisionResponseWithdrawnByApplicant AdoptionGetDecisionResponseWithdrawnBy = "applicant"
	AdoptionGetDecisionResponseWithdrawnByShelter   AdoptionGetDecisionResponseWithdrawnBy = "shelter"
	AdoptionGetDecisionResponseWithdrawnBySystem    AdoptionGetDecisionResponseWithdrawnBy = "system"
)

func (r AdoptionGetDecisionResponseWithdrawnBy) IsKnown() bool {
	switch r {
	case AdoptionGetDecisionResponseWithdrawnByApplicant, AdoptionGetDecisionResponseWithdrawnByShelter, AdoptionGetDecisionResponseWithdrawnBySystem:
		return true
	}
	return false
}

type AdoptionNewParams struct {
	Body AdoptionNewParamsBodyUnion `json:"body" api:"required"`
}

func (r AdoptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AdoptionNewParamsBody struct {
	ApplicantType param.Field[AdoptionNewParamsBodyApplicantType] `json:"applicant_type" api:"required"`
	Name          param.Field[string]                             `json:"name" api:"required"`
	Address       param.Field[shared.AddressParam]                `json:"address"`
	Age           param.Field[int64]                              `json:"age"`
	Email         param.Field[string]                             `json:"email"`
	Shelter       param.Field[interface{}]                        `json:"shelter"`
	TaxID         param.Field[string]                             `json:"taxId"`
}

func (r AdoptionNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdoptionNewParamsBody) implementsAdoptionNewParamsBodyUnion() {}

// Satisfied by [AdoptionNewParamsBodyIndividualApplicant],
// [AdoptionNewParamsBodyOrganizationApplicant], [AdoptionNewParamsBody].
type AdoptionNewParamsBodyUnion interface {
	implementsAdoptionNewParamsBodyUnion()
}

type AdoptionNewParamsBodyIndividualApplicant struct {
	ApplicantType param.Field[AdoptionNewParamsBodyIndividualApplicantApplicantType] `json:"applicant_type" api:"required"`
	Name          param.Field[string]                                                `json:"name" api:"required"`
	Address       param.Field[shared.AddressParam]                                   `json:"address"`
	Age           param.Field[int64]                                                 `json:"age"`
	Email         param.Field[string]                                                `json:"email"`
}

func (r AdoptionNewParamsBodyIndividualApplicant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdoptionNewParamsBodyIndividualApplicant) implementsAdoptionNewParamsBodyUnion() {}

type AdoptionNewParamsBodyIndividualApplicantApplicantType string

const (
	AdoptionNewParamsBodyIndividualApplicantApplicantTypeIndividual AdoptionNewParamsBodyIndividualApplicantApplicantType = "individual"
)

func (r AdoptionNewParamsBodyIndividualApplicantApplicantType) IsKnown() bool {
	switch r {
	case AdoptionNewParamsBodyIndividualApplicantApplicantTypeIndividual:
		return true
	}
	return false
}

type AdoptionNewParamsBodyOrganizationApplicant struct {
	ApplicantType param.Field[AdoptionNewParamsBodyOrganizationApplicantApplicantType] `json:"applicant_type" api:"required"`
	Name          param.Field[string]                                                  `json:"name" api:"required"`
	Shelter       param.Field[AdoptionNewParamsBodyOrganizationApplicantShelter]       `json:"shelter"`
	TaxID         param.Field[string]                                                  `json:"taxId"`
}

func (r AdoptionNewParamsBodyOrganizationApplicant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdoptionNewParamsBodyOrganizationApplicant) implementsAdoptionNewParamsBodyUnion() {}

type AdoptionNewParamsBodyOrganizationApplicantApplicantType string

const (
	AdoptionNewParamsBodyOrganizationApplicantApplicantTypeOrganization AdoptionNewParamsBodyOrganizationApplicantApplicantType = "organization"
)

func (r AdoptionNewParamsBodyOrganizationApplicantApplicantType) IsKnown() bool {
	switch r {
	case AdoptionNewParamsBodyOrganizationApplicantApplicantTypeOrganization:
		return true
	}
	return false
}

type AdoptionNewParamsBodyOrganizationApplicantShelter struct {
	Name            param.Field[string]                                                    `json:"name" api:"required"`
	Address         param.Field[shared.AddressParam]                                       `json:"address"`
	Contact         param.Field[AdoptionNewParamsBodyOrganizationApplicantShelterContact]  `json:"contact"`
	Location        param.Field[AdoptionNewParamsBodyOrganizationApplicantShelterLocation] `json:"location"`
	RelatedAddress  param.Field[shared.AddressParam]                                       `json:"relatedAddress"`
	RelatedCategory param.Field[interface{}]                                               `json:"relatedCategory"`
	RelatedCustomer param.Field[interface{}]                                               `json:"relatedCustomer"`
	RelatedMoney    param.Field[shared.MoneyParam]                                         `json:"relatedMoney"`
	RelatedOrder    param.Field[interface{}]                                               `json:"relatedOrder"`
	RelatedPet      param.Field[PetParam]                                                  `json:"relatedPet"`
	RelatedTag      param.Field[interface{}]                                               `json:"relatedTag"`
	RelatedUser     param.Field[interface{}]                                               `json:"relatedUser"`
}

func (r AdoptionNewParamsBodyOrganizationApplicantShelter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdoptionNewParamsBodyOrganizationApplicantShelterContact struct {
	Primary param.Field[AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimary] `json:"primary"`
}

func (r AdoptionNewParamsBodyOrganizationApplicantShelterContact) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimary struct {
	Availability param.Field[AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimaryAvailability] `json:"availability"`
	Name         param.Field[string]                                                                      `json:"name"`
	Phone        param.Field[string]                                                                      `json:"phone"`
}

func (r AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimary) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimaryAvailability struct {
	Weekday param.Field[AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimaryAvailabilityWeekday] `json:"weekday"`
	Weekend param.Field[AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimaryAvailabilityWeekend] `json:"weekend"`
}

func (r AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimaryAvailability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimaryAvailabilityWeekday struct {
	Close param.Field[string] `json:"close"`
	Open  param.Field[string] `json:"open"`
}

func (r AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimaryAvailabilityWeekday) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimaryAvailabilityWeekend struct {
	Close param.Field[string] `json:"close"`
	Open  param.Field[string] `json:"open"`
}

func (r AdoptionNewParamsBodyOrganizationApplicantShelterContactPrimaryAvailabilityWeekend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdoptionNewParamsBodyOrganizationApplicantShelterLocation struct {
	Latitude  param.Field[float64] `json:"latitude" api:"required"`
	Longitude param.Field[float64] `json:"longitude" api:"required"`
}

func (r AdoptionNewParamsBodyOrganizationApplicantShelterLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdoptionNewParamsBodyApplicantType string

const (
	AdoptionNewParamsBodyApplicantTypeIndividual   AdoptionNewParamsBodyApplicantType = "individual"
	AdoptionNewParamsBodyApplicantTypeOrganization AdoptionNewParamsBodyApplicantType = "organization"
)

func (r AdoptionNewParamsBodyApplicantType) IsKnown() bool {
	switch r {
	case AdoptionNewParamsBodyApplicantTypeIndividual, AdoptionNewParamsBodyApplicantTypeOrganization:
		return true
	}
	return false
}
