// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"time"

	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/respjson"
	"github.com/swagger-girly/go/shared"
	"github.com/swagger-girly/go/shared/constant"
)

// NotificationWebhookService contains methods and other services that help with
// interacting with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationWebhookService] method instead.
type NotificationWebhookService struct {
	options []option.RequestOption
}

// NewNotificationWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationWebhookService(opts ...option.RequestOption) (r NotificationWebhookService) {
	r = NotificationWebhookService{}
	r.options = opts
	return
}

func (r *NotificationWebhookService) ParsePet(payload []byte, headers http.Header, opts ...option.RequestOption) (*ParsePetWebhookEventUnion, error) {
	opts = slices.Concat(r.options, opts)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	key := cfg.WebhookSecret
	if key == "" {
		return nil, errors.New("The WebhookSecret option must be set in order to verify webhook headers")
	}
	wh, err := standardwebhooks.NewWebhook(key)
	if err != nil {
		return nil, err
	}
	err = wh.Verify(payload, headers)
	if err != nil {
		return nil, err
	}
	res := &ParsePetWebhookEventUnion{}
	err = res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

type PetCreatedWebhookEvent struct {
	Pet Pet `json:"pet" api:"required"`
	// Any of "pet.created".
	Type PetCreatedWebhookEventType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pet         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetCreatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PetCreatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetCreatedWebhookEventType string

const (
	PetCreatedWebhookEventTypePetCreated PetCreatedWebhookEventType = "pet.created"
)

type PetUpdatedWebhookEvent struct {
	// Any of "name", "status", "category", "tags".
	ChangedFields []string `json:"changed_fields" api:"required"`
	Pet           Pet      `json:"pet" api:"required"`
	// Any of "pet.updated".
	Type     PetUpdatedWebhookEventType `json:"type" api:"required"`
	Metadata map[string]string          `json:"metadata"`
	// pet status in the store
	//
	// Any of "available", "pending", "sold".
	PreviousStatus PetStatus `json:"previous_status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangedFields  respjson.Field
		Pet            respjson.Field
		Type           respjson.Field
		Metadata       respjson.Field
		PreviousStatus respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetUpdatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PetUpdatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetUpdatedWebhookEventType string

const (
	PetUpdatedWebhookEventTypePetUpdated PetUpdatedWebhookEventType = "pet.updated"
)

type PetInventoryLowWebhookEvent struct {
	Pet       Pet   `json:"pet" api:"required"`
	Quantity  int64 `json:"quantity" api:"required"`
	Threshold int64 `json:"threshold" api:"required"`
	// Any of "pet.inventory-low".
	Type      PetInventoryLowWebhookEventType      `json:"type" api:"required"`
	LastOrder PetInventoryLowWebhookEventLastOrder `json:"last_order"`
	Locations []shared.Address                     `json:"locations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pet         respjson.Field
		Quantity    respjson.Field
		Threshold   respjson.Field
		Type        respjson.Field
		LastOrder   respjson.Field
		Locations   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetInventoryLowWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PetInventoryLowWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetInventoryLowWebhookEventType string

const (
	PetInventoryLowWebhookEventTypePetInventoryLow PetInventoryLowWebhookEventType = "pet.inventory-low"
)

type PetInventoryLowWebhookEventLastOrder struct {
	ID              int64          `json:"id"`
	Complete        bool           `json:"complete"`
	PetID           int64          `json:"petId"`
	Quantity        int64          `json:"quantity"`
	RelatedAddress  shared.Address `json:"relatedAddress"`
	RelatedCategory any            `json:"relatedCategory"`
	RelatedCustomer any            `json:"relatedCustomer"`
	RelatedMoney    shared.Money   `json:"relatedMoney"`
	RelatedPet      Pet            `json:"relatedPet"`
	RelatedShelter  any            `json:"relatedShelter"`
	RelatedTag      any            `json:"relatedTag"`
	RelatedUser     any            `json:"relatedUser"`
	ShipDate        time.Time      `json:"shipDate" format:"date-time"`
	// Order Status
	//
	// Any of "placed", "approved", "delivered".
	Status string       `json:"status"`
	Total  shared.Money `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Complete        respjson.Field
		PetID           respjson.Field
		Quantity        respjson.Field
		RelatedAddress  respjson.Field
		RelatedCategory respjson.Field
		RelatedCustomer respjson.Field
		RelatedMoney    respjson.Field
		RelatedPet      respjson.Field
		RelatedShelter  respjson.Field
		RelatedTag      respjson.Field
		RelatedUser     respjson.Field
		ShipDate        respjson.Field
		Status          respjson.Field
		Total           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetInventoryLowWebhookEventLastOrder) RawJSON() string { return r.JSON.raw }
func (r *PetInventoryLowWebhookEventLastOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PetModerationWebhookEventUnion contains all possible properties and values from
// [PetModerationWebhookEventPetModerationApproved],
// [PetModerationWebhookEventPetModerationRejected].
//
// Use the [PetModerationWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PetModerationWebhookEventUnion struct {
	// This field is from variant [PetModerationWebhookEventPetModerationApproved].
	ApprovedAt time.Time `json:"approved_at"`
	// This field is from variant [PetModerationWebhookEventPetModerationApproved].
	Pet Pet `json:"pet"`
	// Any of "pet.moderation.approved", "pet.moderation.rejected".
	Type string `json:"type"`
	// This field is from variant [PetModerationWebhookEventPetModerationRejected].
	Reason string `json:"reason"`
	// This field is from variant [PetModerationWebhookEventPetModerationRejected].
	ReviewNotes []PetModerationWebhookEventPetModerationRejectedReviewNote `json:"review_notes"`
	JSON        struct {
		ApprovedAt  respjson.Field
		Pet         respjson.Field
		Type        respjson.Field
		Reason      respjson.Field
		ReviewNotes respjson.Field
		raw         string
	} `json:"-"`
}

// anyPetModerationWebhookEvent is implemented by each variant of
// [PetModerationWebhookEventUnion] to add type safety for the return type of
// [PetModerationWebhookEventUnion.AsAny]
type anyPetModerationWebhookEvent interface {
	implPetModerationWebhookEventUnion()
}

func (PetModerationWebhookEventPetModerationApproved) implPetModerationWebhookEventUnion() {}
func (PetModerationWebhookEventPetModerationRejected) implPetModerationWebhookEventUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PetModerationWebhookEventUnion.AsAny().(type) {
//	case helloworldtestingggg.PetModerationWebhookEventPetModerationApproved:
//	case helloworldtestingggg.PetModerationWebhookEventPetModerationRejected:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PetModerationWebhookEventUnion) AsAny() anyPetModerationWebhookEvent {
	switch u.Type {
	case "pet.moderation.approved":
		return u.AsPetModerationApproved()
	case "pet.moderation.rejected":
		return u.AsPetModerationRejected()
	}
	return nil
}

func (u PetModerationWebhookEventUnion) AsPetModerationApproved() (v PetModerationWebhookEventPetModerationApproved) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PetModerationWebhookEventUnion) AsPetModerationRejected() (v PetModerationWebhookEventPetModerationRejected) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PetModerationWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *PetModerationWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetModerationWebhookEventPetModerationApproved struct {
	ApprovedAt time.Time                      `json:"approved_at" api:"required" format:"date-time"`
	Pet        Pet                            `json:"pet" api:"required"`
	Type       constant.PetModerationApproved `json:"type" default:"pet.moderation.approved"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovedAt  respjson.Field
		Pet         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetModerationWebhookEventPetModerationApproved) RawJSON() string { return r.JSON.raw }
func (r *PetModerationWebhookEventPetModerationApproved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetModerationWebhookEventPetModerationRejected struct {
	Pet Pet `json:"pet" api:"required"`
	// Any of "policy_violation", "duplicate", "unsafe-content".
	Reason      string                                                     `json:"reason" api:"required"`
	Type        constant.PetModerationRejected                             `json:"type" default:"pet.moderation.rejected"`
	ReviewNotes []PetModerationWebhookEventPetModerationRejectedReviewNote `json:"review_notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pet         respjson.Field
		Reason      respjson.Field
		Type        respjson.Field
		ReviewNotes respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetModerationWebhookEventPetModerationRejected) RawJSON() string { return r.JSON.raw }
func (r *PetModerationWebhookEventPetModerationRejected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetModerationWebhookEventPetModerationRejectedReviewNote struct {
	Message  string                                                           `json:"message" api:"required"`
	Reviewer PetModerationWebhookEventPetModerationRejectedReviewNoteReviewer `json:"reviewer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Reviewer    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetModerationWebhookEventPetModerationRejectedReviewNote) RawJSON() string { return r.JSON.raw }
func (r *PetModerationWebhookEventPetModerationRejectedReviewNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetModerationWebhookEventPetModerationRejectedReviewNoteReviewer struct {
	ID              int64          `json:"id"`
	Email           string         `json:"email"`
	FirstName       string         `json:"firstName"`
	LastName        string         `json:"lastName"`
	Password        string         `json:"password"`
	Phone           string         `json:"phone"`
	RelatedAddress  shared.Address `json:"relatedAddress"`
	RelatedCategory any            `json:"relatedCategory"`
	RelatedCustomer any            `json:"relatedCustomer"`
	RelatedMoney    shared.Money   `json:"relatedMoney"`
	RelatedOrder    any            `json:"relatedOrder"`
	RelatedPet      Pet            `json:"relatedPet"`
	RelatedShelter  any            `json:"relatedShelter"`
	RelatedTag      any            `json:"relatedTag"`
	Username        string         `json:"username"`
	// User Status
	UserStatus int64 `json:"userStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Email           respjson.Field
		FirstName       respjson.Field
		LastName        respjson.Field
		Password        respjson.Field
		Phone           respjson.Field
		RelatedAddress  respjson.Field
		RelatedCategory respjson.Field
		RelatedCustomer respjson.Field
		RelatedMoney    respjson.Field
		RelatedOrder    respjson.Field
		RelatedPet      respjson.Field
		RelatedShelter  respjson.Field
		RelatedTag      respjson.Field
		Username        respjson.Field
		UserStatus      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetModerationWebhookEventPetModerationRejectedReviewNoteReviewer) RawJSON() string {
	return r.JSON.raw
}
func (r *PetModerationWebhookEventPetModerationRejectedReviewNoteReviewer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParsePetWebhookEventUnion contains all possible properties and values from
// [PetCreatedWebhookEvent], [PetUpdatedWebhookEvent],
// [PetInventoryLowWebhookEvent], [ParsePetWebhookEventPetModerationApprovedEvent],
// [ParsePetWebhookEventPetModerationRejectedEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ParsePetWebhookEventUnion struct {
	// This field is from variant [PetCreatedWebhookEvent].
	Pet  Pet    `json:"pet"`
	Type string `json:"type"`
	// This field is from variant [PetUpdatedWebhookEvent].
	ChangedFields []string `json:"changed_fields"`
	// This field is from variant [PetUpdatedWebhookEvent].
	Metadata map[string]string `json:"metadata"`
	// This field is from variant [PetUpdatedWebhookEvent].
	PreviousStatus PetStatus `json:"previous_status"`
	// This field is from variant [PetInventoryLowWebhookEvent].
	Quantity int64 `json:"quantity"`
	// This field is from variant [PetInventoryLowWebhookEvent].
	Threshold int64 `json:"threshold"`
	// This field is from variant [PetInventoryLowWebhookEvent].
	LastOrder PetInventoryLowWebhookEventLastOrder `json:"last_order"`
	// This field is from variant [PetInventoryLowWebhookEvent].
	Locations []shared.Address `json:"locations"`
	// This field is from variant [ParsePetWebhookEventPetModerationApprovedEvent].
	ApprovedAt time.Time `json:"approved_at"`
	// This field is from variant [ParsePetWebhookEventPetModerationRejectedEvent].
	Reason string `json:"reason"`
	// This field is from variant [ParsePetWebhookEventPetModerationRejectedEvent].
	ReviewNotes []ParsePetWebhookEventPetModerationRejectedEventReviewNote `json:"review_notes"`
	JSON        struct {
		Pet            respjson.Field
		Type           respjson.Field
		ChangedFields  respjson.Field
		Metadata       respjson.Field
		PreviousStatus respjson.Field
		Quantity       respjson.Field
		Threshold      respjson.Field
		LastOrder      respjson.Field
		Locations      respjson.Field
		ApprovedAt     respjson.Field
		Reason         respjson.Field
		ReviewNotes    respjson.Field
		raw            string
	} `json:"-"`
}

func (u ParsePetWebhookEventUnion) AsPetCreatedWebhookEvent() (v PetCreatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsePetWebhookEventUnion) AsPetUpdatedWebhookEvent() (v PetUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsePetWebhookEventUnion) AsPetInventoryLowWebhookEvent() (v PetInventoryLowWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsePetWebhookEventUnion) AsParsePetWebhookEventPetModerationApprovedEvent() (v ParsePetWebhookEventPetModerationApprovedEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsePetWebhookEventUnion) AsParsePetWebhookEventPetModerationRejectedEvent() (v ParsePetWebhookEventPetModerationRejectedEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParsePetWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *ParsePetWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParsePetWebhookEventPetModerationApprovedEvent struct {
	ApprovedAt time.Time `json:"approved_at" api:"required" format:"date-time"`
	Pet        Pet       `json:"pet" api:"required"`
	// Any of "pet.moderation.approved".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovedAt  respjson.Field
		Pet         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsePetWebhookEventPetModerationApprovedEvent) RawJSON() string { return r.JSON.raw }
func (r *ParsePetWebhookEventPetModerationApprovedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParsePetWebhookEventPetModerationRejectedEvent struct {
	Pet Pet `json:"pet" api:"required"`
	// Any of "policy_violation", "duplicate", "unsafe-content".
	Reason string `json:"reason" api:"required"`
	// Any of "pet.moderation.rejected".
	Type        string                                                     `json:"type" api:"required"`
	ReviewNotes []ParsePetWebhookEventPetModerationRejectedEventReviewNote `json:"review_notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pet         respjson.Field
		Reason      respjson.Field
		Type        respjson.Field
		ReviewNotes respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsePetWebhookEventPetModerationRejectedEvent) RawJSON() string { return r.JSON.raw }
func (r *ParsePetWebhookEventPetModerationRejectedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParsePetWebhookEventPetModerationRejectedEventReviewNote struct {
	Message  string                                                           `json:"message" api:"required"`
	Reviewer ParsePetWebhookEventPetModerationRejectedEventReviewNoteReviewer `json:"reviewer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Reviewer    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsePetWebhookEventPetModerationRejectedEventReviewNote) RawJSON() string { return r.JSON.raw }
func (r *ParsePetWebhookEventPetModerationRejectedEventReviewNote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParsePetWebhookEventPetModerationRejectedEventReviewNoteReviewer struct {
	ID              int64          `json:"id"`
	Email           string         `json:"email"`
	FirstName       string         `json:"firstName"`
	LastName        string         `json:"lastName"`
	Password        string         `json:"password"`
	Phone           string         `json:"phone"`
	RelatedAddress  shared.Address `json:"relatedAddress"`
	RelatedCategory any            `json:"relatedCategory"`
	RelatedCustomer any            `json:"relatedCustomer"`
	RelatedMoney    shared.Money   `json:"relatedMoney"`
	RelatedOrder    any            `json:"relatedOrder"`
	RelatedPet      Pet            `json:"relatedPet"`
	RelatedShelter  any            `json:"relatedShelter"`
	RelatedTag      any            `json:"relatedTag"`
	Username        string         `json:"username"`
	// User Status
	UserStatus int64 `json:"userStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Email           respjson.Field
		FirstName       respjson.Field
		LastName        respjson.Field
		Password        respjson.Field
		Phone           respjson.Field
		RelatedAddress  respjson.Field
		RelatedCategory respjson.Field
		RelatedCustomer respjson.Field
		RelatedMoney    respjson.Field
		RelatedOrder    respjson.Field
		RelatedPet      respjson.Field
		RelatedShelter  respjson.Field
		RelatedTag      respjson.Field
		Username        respjson.Field
		UserStatus      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsePetWebhookEventPetModerationRejectedEventReviewNoteReviewer) RawJSON() string {
	return r.JSON.raw
}
func (r *ParsePetWebhookEventPetModerationRejectedEventReviewNoteReviewer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
