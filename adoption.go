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
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
	"github.com/swagger-girly/go/shared"
	"github.com/swagger-girly/go/shared/constant"
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
	options []option.RequestOption
	// Adoption policies and applications
	Policies AdoptionPolicyService
}

// NewAdoptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdoptionService(opts ...option.RequestOption) (r AdoptionService) {
	r = AdoptionService{}
	r.options = opts
	r.Policies = NewAdoptionPolicyService(opts...)
	return
}

// Submits an adoption application for an individual or an organization applicant.
func (r *AdoptionService) New(ctx context.Context, body AdoptionNewParams, opts ...option.RequestOption) (res *Application, err error) {
	opts = slices.Concat(r.options, opts)
	path := "adoptions/applications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single adoption application.
func (r *AdoptionService) Get(ctx context.Context, applicationID string, opts ...option.RequestOption) (res *Application, err error) {
	opts = slices.Concat(r.options, opts)
	if applicationID == "" {
		err = errors.New("missing required applicationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("adoptions/applications/%s", url.PathEscape(applicationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the decision for an application. The decision shape depends on the
// outcome and carries no discriminator mapping.
func (r *AdoptionService) GetDecision(ctx context.Context, applicationID string, opts ...option.RequestOption) (res *AdoptionGetDecisionResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if applicationID == "" {
		err = errors.New("missing required applicationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("adoptions/applications/%s/decision", url.PathEscape(applicationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Application struct {
	ID        string                    `json:"id" api:"required"`
	Applicant ApplicationApplicantUnion `json:"applicant" api:"required"`
	// Any of "received", "under_review", "approved", "rejected".
	Status      ApplicationStatus `json:"status" api:"required"`
	SubmittedAt time.Time         `json:"submittedAt" api:"required" format:"date-time"`
	// The decision for an application; the shape depends on the outcome.
	Decision        ApplicationDecisionUnion   `json:"decision"`
	Fee             shared.Money               `json:"fee"`
	History         []ApplicationHistory       `json:"history"`
	LatestRejection ApplicationLatestRejection `json:"latestRejection"`
	Pet             Pet                        `json:"pet"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Applicant       respjson.Field
		Status          respjson.Field
		SubmittedAt     respjson.Field
		Decision        respjson.Field
		Fee             respjson.Field
		History         respjson.Field
		LatestRejection respjson.Field
		Pet             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Application) RawJSON() string { return r.JSON.raw }
func (r *Application) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ApplicationApplicantUnion contains all possible properties and values from
// [ApplicationApplicantIndividual], [ApplicationApplicantOrganization].
//
// Use the [ApplicationApplicantUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ApplicationApplicantUnion struct {
	// Any of "individual", "organization".
	ApplicantType string `json:"applicant_type"`
	Name          string `json:"name"`
	// This field is from variant [ApplicationApplicantIndividual].
	Address shared.Address `json:"address"`
	// This field is from variant [ApplicationApplicantIndividual].
	Age int64 `json:"age"`
	// This field is from variant [ApplicationApplicantIndividual].
	Email string `json:"email"`
	// This field is from variant [ApplicationApplicantOrganization].
	Shelter ApplicationApplicantOrganizationShelter `json:"shelter"`
	// This field is from variant [ApplicationApplicantOrganization].
	TaxID string `json:"taxId"`
	JSON  struct {
		ApplicantType respjson.Field
		Name          respjson.Field
		Address       respjson.Field
		Age           respjson.Field
		Email         respjson.Field
		Shelter       respjson.Field
		TaxID         respjson.Field
		raw           string
	} `json:"-"`
}

// anyApplicationApplicant is implemented by each variant of
// [ApplicationApplicantUnion] to add type safety for the return type of
// [ApplicationApplicantUnion.AsAny]
type anyApplicationApplicant interface {
	implApplicationApplicantUnion()
}

func (ApplicationApplicantIndividual) implApplicationApplicantUnion()   {}
func (ApplicationApplicantOrganization) implApplicationApplicantUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ApplicationApplicantUnion.AsAny().(type) {
//	case helloworldtestingggg.ApplicationApplicantIndividual:
//	case helloworldtestingggg.ApplicationApplicantOrganization:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ApplicationApplicantUnion) AsAny() anyApplicationApplicant {
	switch u.ApplicantType {
	case "individual":
		return u.AsIndividual()
	case "organization":
		return u.AsOrganization()
	}
	return nil
}

func (u ApplicationApplicantUnion) AsIndividual() (v ApplicationApplicantIndividual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ApplicationApplicantUnion) AsOrganization() (v ApplicationApplicantOrganization) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ApplicationApplicantUnion) RawJSON() string { return u.JSON.raw }

func (r *ApplicationApplicantUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationApplicantIndividual struct {
	ApplicantType constant.Individual `json:"applicant_type" default:"individual"`
	Name          string              `json:"name" api:"required"`
	Address       shared.Address      `json:"address"`
	Age           int64               `json:"age"`
	Email         string              `json:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicantType respjson.Field
		Name          respjson.Field
		Address       respjson.Field
		Age           respjson.Field
		Email         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationApplicantIndividual) RawJSON() string { return r.JSON.raw }
func (r *ApplicationApplicantIndividual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationApplicantOrganization struct {
	ApplicantType constant.Organization                   `json:"applicant_type" default:"organization"`
	Name          string                                  `json:"name" api:"required"`
	Shelter       ApplicationApplicantOrganizationShelter `json:"shelter"`
	TaxID         string                                  `json:"taxId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicantType respjson.Field
		Name          respjson.Field
		Shelter       respjson.Field
		TaxID         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationApplicantOrganization) RawJSON() string { return r.JSON.raw }
func (r *ApplicationApplicantOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationApplicantOrganizationShelter struct {
	Name     string                                          `json:"name" api:"required"`
	Address  shared.Address                                  `json:"address"`
	Contact  ApplicationApplicantOrganizationShelterContact  `json:"contact"`
	Location ApplicationApplicantOrganizationShelterLocation `json:"location"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Address     respjson.Field
		Contact     respjson.Field
		Location    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationApplicantOrganizationShelter) RawJSON() string { return r.JSON.raw }
func (r *ApplicationApplicantOrganizationShelter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationApplicantOrganizationShelterContact struct {
	Primary ApplicationApplicantOrganizationShelterContactPrimary `json:"primary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationApplicantOrganizationShelterContact) RawJSON() string { return r.JSON.raw }
func (r *ApplicationApplicantOrganizationShelterContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationApplicantOrganizationShelterContactPrimary struct {
	Availability ApplicationApplicantOrganizationShelterContactPrimaryAvailability `json:"availability"`
	Name         string                                                            `json:"name"`
	Phone        string                                                            `json:"phone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Availability respjson.Field
		Name         respjson.Field
		Phone        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationApplicantOrganizationShelterContactPrimary) RawJSON() string { return r.JSON.raw }
func (r *ApplicationApplicantOrganizationShelterContactPrimary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationApplicantOrganizationShelterContactPrimaryAvailability struct {
	Weekday ApplicationApplicantOrganizationShelterContactPrimaryAvailabilityWeekday `json:"weekday"`
	Weekend ApplicationApplicantOrganizationShelterContactPrimaryAvailabilityWeekend `json:"weekend" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Weekday     respjson.Field
		Weekend     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationApplicantOrganizationShelterContactPrimaryAvailability) RawJSON() string {
	return r.JSON.raw
}
func (r *ApplicationApplicantOrganizationShelterContactPrimaryAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationApplicantOrganizationShelterContactPrimaryAvailabilityWeekday struct {
	Close string `json:"close"`
	Open  string `json:"open"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close       respjson.Field
		Open        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationApplicantOrganizationShelterContactPrimaryAvailabilityWeekday) RawJSON() string {
	return r.JSON.raw
}
func (r *ApplicationApplicantOrganizationShelterContactPrimaryAvailabilityWeekday) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationApplicantOrganizationShelterContactPrimaryAvailabilityWeekend struct {
	Close string `json:"close"`
	Open  string `json:"open"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close       respjson.Field
		Open        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationApplicantOrganizationShelterContactPrimaryAvailabilityWeekend) RawJSON() string {
	return r.JSON.raw
}
func (r *ApplicationApplicantOrganizationShelterContactPrimaryAvailabilityWeekend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationApplicantOrganizationShelterLocation struct {
	Latitude  float64 `json:"latitude" api:"required"`
	Longitude float64 `json:"longitude" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationApplicantOrganizationShelterLocation) RawJSON() string { return r.JSON.raw }
func (r *ApplicationApplicantOrganizationShelterLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationStatus string

const (
	ApplicationStatusReceived    ApplicationStatus = "received"
	ApplicationStatusUnderReview ApplicationStatus = "under_review"
	ApplicationStatusApproved    ApplicationStatus = "approved"
	ApplicationStatusRejected    ApplicationStatus = "rejected"
)

// ApplicationDecisionUnion contains all possible properties and values from
// [ApplicationDecisionDecisionApproved], [ApplicationDecisionDecisionRejected],
// [ApplicationDecisionDecisionEscalated], [ApplicationDecisionDecisionWithdrawn].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ApplicationDecisionUnion struct {
	// This field is from variant [ApplicationDecisionDecisionApproved].
	ApprovedAt time.Time `json:"approvedAt"`
	Outcome    string    `json:"outcome"`
	// This field is from variant [ApplicationDecisionDecisionApproved].
	Conditions []string `json:"conditions"`
	// This field is from variant [ApplicationDecisionDecisionRejected].
	Reason string `json:"reason"`
	// This field is from variant [ApplicationDecisionDecisionRejected].
	AppealDeadline time.Time `json:"appealDeadline"`
	// This field is from variant [ApplicationDecisionDecisionEscalated].
	EscalatedTo ApplicationDecisionDecisionEscalatedEscalatedTo `json:"escalatedTo"`
	// This field is from variant [ApplicationDecisionDecisionEscalated].
	ReviewAfter time.Time `json:"reviewAfter"`
	// This field is from variant [ApplicationDecisionDecisionWithdrawn].
	WithdrawnBy string `json:"withdrawnBy"`
	// This field is from variant [ApplicationDecisionDecisionWithdrawn].
	WithdrawnAt time.Time `json:"withdrawnAt"`
	JSON        struct {
		ApprovedAt     respjson.Field
		Outcome        respjson.Field
		Conditions     respjson.Field
		Reason         respjson.Field
		AppealDeadline respjson.Field
		EscalatedTo    respjson.Field
		ReviewAfter    respjson.Field
		WithdrawnBy    respjson.Field
		WithdrawnAt    respjson.Field
		raw            string
	} `json:"-"`
}

func (u ApplicationDecisionUnion) AsApplicationDecisionDecisionApproved() (v ApplicationDecisionDecisionApproved) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ApplicationDecisionUnion) AsApplicationDecisionDecisionRejected() (v ApplicationDecisionDecisionRejected) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ApplicationDecisionUnion) AsApplicationDecisionDecisionEscalated() (v ApplicationDecisionDecisionEscalated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ApplicationDecisionUnion) AsApplicationDecisionDecisionWithdrawn() (v ApplicationDecisionDecisionWithdrawn) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ApplicationDecisionUnion) RawJSON() string { return u.JSON.raw }

func (r *ApplicationDecisionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationDecisionDecisionApproved struct {
	ApprovedAt time.Time         `json:"approvedAt" api:"required" format:"date-time"`
	Outcome    constant.Approved `json:"outcome" default:"approved"`
	Conditions []string          `json:"conditions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovedAt  respjson.Field
		Outcome     respjson.Field
		Conditions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationDecisionDecisionApproved) RawJSON() string { return r.JSON.raw }
func (r *ApplicationDecisionDecisionApproved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationDecisionDecisionRejected struct {
	Outcome constant.Rejected `json:"outcome" default:"rejected"`
	// Any of "policy", "capacity", "incomplete".
	Reason         string    `json:"reason" api:"required"`
	AppealDeadline time.Time `json:"appealDeadline" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Outcome        respjson.Field
		Reason         respjson.Field
		AppealDeadline respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationDecisionDecisionRejected) RawJSON() string { return r.JSON.raw }
func (r *ApplicationDecisionDecisionRejected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationDecisionDecisionEscalated struct {
	EscalatedTo ApplicationDecisionDecisionEscalatedEscalatedTo `json:"escalatedTo" api:"required"`
	Outcome     constant.Escalated                              `json:"outcome" default:"escalated"`
	ReviewAfter time.Time                                       `json:"reviewAfter" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EscalatedTo respjson.Field
		Outcome     respjson.Field
		ReviewAfter respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationDecisionDecisionEscalated) RawJSON() string { return r.JSON.raw }
func (r *ApplicationDecisionDecisionEscalated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationDecisionDecisionEscalatedEscalatedTo struct {
	Team    string                                                 `json:"team" api:"required"`
	Contact ApplicationDecisionDecisionEscalatedEscalatedToContact `json:"contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Team        respjson.Field
		Contact     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationDecisionDecisionEscalatedEscalatedTo) RawJSON() string { return r.JSON.raw }
func (r *ApplicationDecisionDecisionEscalatedEscalatedTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationDecisionDecisionEscalatedEscalatedToContact struct {
	Email string `json:"email"`
	Hours string `json:"hours"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Hours       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationDecisionDecisionEscalatedEscalatedToContact) RawJSON() string { return r.JSON.raw }
func (r *ApplicationDecisionDecisionEscalatedEscalatedToContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The applicant or shelter withdrew before a decision was finalized.
type ApplicationDecisionDecisionWithdrawn struct {
	// Any of "withdrawn", "expired".
	Outcome string `json:"outcome" api:"required"`
	// Any of "applicant", "shelter", "system".
	WithdrawnBy string    `json:"withdrawnBy" api:"required"`
	WithdrawnAt time.Time `json:"withdrawnAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Outcome     respjson.Field
		WithdrawnBy respjson.Field
		WithdrawnAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationDecisionDecisionWithdrawn) RawJSON() string { return r.JSON.raw }
func (r *ApplicationDecisionDecisionWithdrawn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationHistory struct {
	ChangedAt time.Time `json:"changedAt" api:"required" format:"date-time"`
	// Any of "received", "under_review", "approved", "rejected".
	Status string `json:"status" api:"required"`
	Note   string `json:"note" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedAt   respjson.Field
		Status      respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationHistory) RawJSON() string { return r.JSON.raw }
func (r *ApplicationHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationLatestRejection struct {
	Outcome constant.Rejected `json:"outcome" default:"rejected"`
	// Any of "policy", "capacity", "incomplete".
	Reason         string    `json:"reason" api:"required"`
	AppealDeadline time.Time `json:"appealDeadline" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Outcome        respjson.Field
		Reason         respjson.Field
		AppealDeadline respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationLatestRejection) RawJSON() string { return r.JSON.raw }
func (r *ApplicationLatestRejection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AdoptionGetDecisionResponseUnion contains all possible properties and values
// from [AdoptionGetDecisionResponseDecisionApproved],
// [AdoptionGetDecisionResponseDecisionRejected],
// [AdoptionGetDecisionResponseDecisionEscalated],
// [AdoptionGetDecisionResponseDecisionWithdrawn].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AdoptionGetDecisionResponseUnion struct {
	// This field is from variant [AdoptionGetDecisionResponseDecisionApproved].
	ApprovedAt time.Time `json:"approvedAt"`
	Outcome    string    `json:"outcome"`
	// This field is from variant [AdoptionGetDecisionResponseDecisionApproved].
	Conditions []string `json:"conditions"`
	// This field is from variant [AdoptionGetDecisionResponseDecisionRejected].
	Reason string `json:"reason"`
	// This field is from variant [AdoptionGetDecisionResponseDecisionRejected].
	AppealDeadline time.Time `json:"appealDeadline"`
	// This field is from variant [AdoptionGetDecisionResponseDecisionEscalated].
	EscalatedTo AdoptionGetDecisionResponseDecisionEscalatedEscalatedTo `json:"escalatedTo"`
	// This field is from variant [AdoptionGetDecisionResponseDecisionEscalated].
	ReviewAfter time.Time `json:"reviewAfter"`
	// This field is from variant [AdoptionGetDecisionResponseDecisionWithdrawn].
	WithdrawnBy string `json:"withdrawnBy"`
	// This field is from variant [AdoptionGetDecisionResponseDecisionWithdrawn].
	WithdrawnAt time.Time `json:"withdrawnAt"`
	JSON        struct {
		ApprovedAt     respjson.Field
		Outcome        respjson.Field
		Conditions     respjson.Field
		Reason         respjson.Field
		AppealDeadline respjson.Field
		EscalatedTo    respjson.Field
		ReviewAfter    respjson.Field
		WithdrawnBy    respjson.Field
		WithdrawnAt    respjson.Field
		raw            string
	} `json:"-"`
}

func (u AdoptionGetDecisionResponseUnion) AsAdoptionGetDecisionResponseDecisionApproved() (v AdoptionGetDecisionResponseDecisionApproved) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AdoptionGetDecisionResponseUnion) AsAdoptionGetDecisionResponseDecisionRejected() (v AdoptionGetDecisionResponseDecisionRejected) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AdoptionGetDecisionResponseUnion) AsAdoptionGetDecisionResponseDecisionEscalated() (v AdoptionGetDecisionResponseDecisionEscalated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AdoptionGetDecisionResponseUnion) AsAdoptionGetDecisionResponseDecisionWithdrawn() (v AdoptionGetDecisionResponseDecisionWithdrawn) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AdoptionGetDecisionResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AdoptionGetDecisionResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionGetDecisionResponseDecisionApproved struct {
	ApprovedAt time.Time         `json:"approvedAt" api:"required" format:"date-time"`
	Outcome    constant.Approved `json:"outcome" default:"approved"`
	Conditions []string          `json:"conditions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovedAt  respjson.Field
		Outcome     respjson.Field
		Conditions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdoptionGetDecisionResponseDecisionApproved) RawJSON() string { return r.JSON.raw }
func (r *AdoptionGetDecisionResponseDecisionApproved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionGetDecisionResponseDecisionRejected struct {
	Outcome constant.Rejected `json:"outcome" default:"rejected"`
	// Any of "policy", "capacity", "incomplete".
	Reason         string    `json:"reason" api:"required"`
	AppealDeadline time.Time `json:"appealDeadline" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Outcome        respjson.Field
		Reason         respjson.Field
		AppealDeadline respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdoptionGetDecisionResponseDecisionRejected) RawJSON() string { return r.JSON.raw }
func (r *AdoptionGetDecisionResponseDecisionRejected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionGetDecisionResponseDecisionEscalated struct {
	EscalatedTo AdoptionGetDecisionResponseDecisionEscalatedEscalatedTo `json:"escalatedTo" api:"required"`
	Outcome     constant.Escalated                                      `json:"outcome" default:"escalated"`
	ReviewAfter time.Time                                               `json:"reviewAfter" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EscalatedTo respjson.Field
		Outcome     respjson.Field
		ReviewAfter respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdoptionGetDecisionResponseDecisionEscalated) RawJSON() string { return r.JSON.raw }
func (r *AdoptionGetDecisionResponseDecisionEscalated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionGetDecisionResponseDecisionEscalatedEscalatedTo struct {
	Team    string                                                         `json:"team" api:"required"`
	Contact AdoptionGetDecisionResponseDecisionEscalatedEscalatedToContact `json:"contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Team        respjson.Field
		Contact     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdoptionGetDecisionResponseDecisionEscalatedEscalatedTo) RawJSON() string { return r.JSON.raw }
func (r *AdoptionGetDecisionResponseDecisionEscalatedEscalatedTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionGetDecisionResponseDecisionEscalatedEscalatedToContact struct {
	Email string `json:"email"`
	Hours string `json:"hours"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Hours       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdoptionGetDecisionResponseDecisionEscalatedEscalatedToContact) RawJSON() string {
	return r.JSON.raw
}
func (r *AdoptionGetDecisionResponseDecisionEscalatedEscalatedToContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The applicant or shelter withdrew before a decision was finalized.
type AdoptionGetDecisionResponseDecisionWithdrawn struct {
	// Any of "withdrawn", "expired".
	Outcome string `json:"outcome" api:"required"`
	// Any of "applicant", "shelter", "system".
	WithdrawnBy string    `json:"withdrawnBy" api:"required"`
	WithdrawnAt time.Time `json:"withdrawnAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Outcome     respjson.Field
		WithdrawnBy respjson.Field
		WithdrawnAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdoptionGetDecisionResponseDecisionWithdrawn) RawJSON() string { return r.JSON.raw }
func (r *AdoptionGetDecisionResponseDecisionWithdrawn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfIndividual *AdoptionNewParamsBodyIndividual `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfOrganization *AdoptionNewParamsBodyOrganization `json:",inline"`

	paramObj
}

func (u AdoptionNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfIndividual, u.OfOrganization)
}
func (r *AdoptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ApplicantType, Name are required.
type AdoptionNewParamsBodyIndividual struct {
	Name    string              `json:"name" api:"required"`
	Age     param.Opt[int64]    `json:"age,omitzero"`
	Email   param.Opt[string]   `json:"email,omitzero"`
	Address shared.AddressParam `json:"address,omitzero"`
	// This field can be elided, and will marshal its zero value as "individual".
	ApplicantType constant.Individual `json:"applicant_type" default:"individual"`
	paramObj
}

func (r AdoptionNewParamsBodyIndividual) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionNewParamsBodyIndividual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionNewParamsBodyIndividual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ApplicantType, Name are required.
type AdoptionNewParamsBodyOrganization struct {
	Name    string                                   `json:"name" api:"required"`
	TaxID   param.Opt[string]                        `json:"taxId,omitzero"`
	Shelter AdoptionNewParamsBodyOrganizationShelter `json:"shelter,omitzero"`
	// This field can be elided, and will marshal its zero value as "organization".
	ApplicantType constant.Organization `json:"applicant_type" default:"organization"`
	paramObj
}

func (r AdoptionNewParamsBodyOrganization) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionNewParamsBodyOrganization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionNewParamsBodyOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type AdoptionNewParamsBodyOrganizationShelter struct {
	Name     string                                           `json:"name" api:"required"`
	Address  shared.AddressParam                              `json:"address,omitzero"`
	Contact  AdoptionNewParamsBodyOrganizationShelterContact  `json:"contact,omitzero"`
	Location AdoptionNewParamsBodyOrganizationShelterLocation `json:"location,omitzero"`
	paramObj
}

func (r AdoptionNewParamsBodyOrganizationShelter) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionNewParamsBodyOrganizationShelter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionNewParamsBodyOrganizationShelter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionNewParamsBodyOrganizationShelterContact struct {
	Primary AdoptionNewParamsBodyOrganizationShelterContactPrimary `json:"primary,omitzero"`
	paramObj
}

func (r AdoptionNewParamsBodyOrganizationShelterContact) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionNewParamsBodyOrganizationShelterContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionNewParamsBodyOrganizationShelterContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionNewParamsBodyOrganizationShelterContactPrimary struct {
	Name         param.Opt[string]                                                  `json:"name,omitzero"`
	Phone        param.Opt[string]                                                  `json:"phone,omitzero"`
	Availability AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailability `json:"availability,omitzero"`
	paramObj
}

func (r AdoptionNewParamsBodyOrganizationShelterContactPrimary) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionNewParamsBodyOrganizationShelterContactPrimary
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionNewParamsBodyOrganizationShelterContactPrimary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailability struct {
	Weekend AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailabilityWeekend `json:"weekend,omitzero"`
	Weekday AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailabilityWeekday `json:"weekday,omitzero"`
	paramObj
}

func (r AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailability) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailability
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailabilityWeekday struct {
	Close param.Opt[string] `json:"close,omitzero"`
	Open  param.Opt[string] `json:"open,omitzero"`
	paramObj
}

func (r AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailabilityWeekday) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailabilityWeekday
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailabilityWeekday) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailabilityWeekend struct {
	Close param.Opt[string] `json:"close,omitzero"`
	Open  param.Opt[string] `json:"open,omitzero"`
	paramObj
}

func (r AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailabilityWeekend) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailabilityWeekend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionNewParamsBodyOrganizationShelterContactPrimaryAvailabilityWeekend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Latitude, Longitude are required.
type AdoptionNewParamsBodyOrganizationShelterLocation struct {
	Latitude  float64 `json:"latitude" api:"required"`
	Longitude float64 `json:"longitude" api:"required"`
	paramObj
}

func (r AdoptionNewParamsBodyOrganizationShelterLocation) MarshalJSON() (data []byte, err error) {
	type shadow AdoptionNewParamsBodyOrganizationShelterLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdoptionNewParamsBodyOrganizationShelterLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
