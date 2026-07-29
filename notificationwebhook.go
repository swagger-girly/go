// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"errors"
	"net/http"
	"reflect"
	"slices"
	"time"

	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/shared"
	"github.com/tidwall/gjson"
)

// NotificationWebhookService contains methods and other services that help with
// interacting with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationWebhookService] method instead.
type NotificationWebhookService struct {
	Options []option.RequestOption
}

// NewNotificationWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationWebhookService(opts ...option.RequestOption) (r *NotificationWebhookService) {
	r = &NotificationWebhookService{}
	r.Options = opts
	return
}

func (r *NotificationWebhookService) ParsePet(payload []byte, headers http.Header, opts ...option.RequestOption) (*ParsePetWebhookEvent, error) {
	opts = slices.Concat(r.Options, opts)
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
	res := &ParsePetWebhookEvent{}
	err = res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

type PetCreatedWebhookEvent struct {
	Pet  Pet                        `json:"pet" api:"required"`
	Type PetCreatedWebhookEventType `json:"type" api:"required"`
	JSON petCreatedWebhookEventJSON `json:"-"`
}

// petCreatedWebhookEventJSON contains the JSON metadata for the struct
// [PetCreatedWebhookEvent]
type petCreatedWebhookEventJSON struct {
	Pet         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PetCreatedWebhookEvent) implementsParsedWebhookEvent() {}

func (r PetCreatedWebhookEvent) implementsParsePetWebhookEvent() {}

type PetCreatedWebhookEventType string

const (
	PetCreatedWebhookEventTypePetCreated PetCreatedWebhookEventType = "pet.created"
)

func (r PetCreatedWebhookEventType) IsKnown() bool {
	switch r {
	case PetCreatedWebhookEventTypePetCreated:
		return true
	}
	return false
}

type PetUpdatedWebhookEvent struct {
	ChangedFields []PetUpdatedWebhookEventChangedField `json:"changed_fields" api:"required"`
	Pet           Pet                                  `json:"pet" api:"required"`
	Type          PetUpdatedWebhookEventType           `json:"type" api:"required"`
	Metadata      map[string]string                    `json:"metadata"`
	// pet status in the store
	PreviousStatus PetStatus                  `json:"previous_status" api:"nullable"`
	JSON           petUpdatedWebhookEventJSON `json:"-"`
}

// petUpdatedWebhookEventJSON contains the JSON metadata for the struct
// [PetUpdatedWebhookEvent]
type petUpdatedWebhookEventJSON struct {
	ChangedFields  apijson.Field
	Pet            apijson.Field
	Type           apijson.Field
	Metadata       apijson.Field
	PreviousStatus apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PetUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PetUpdatedWebhookEvent) implementsParsedWebhookEvent() {}

func (r PetUpdatedWebhookEvent) implementsParsePetWebhookEvent() {}

type PetUpdatedWebhookEventChangedField string

const (
	PetUpdatedWebhookEventChangedFieldName     PetUpdatedWebhookEventChangedField = "name"
	PetUpdatedWebhookEventChangedFieldStatus   PetUpdatedWebhookEventChangedField = "status"
	PetUpdatedWebhookEventChangedFieldCategory PetUpdatedWebhookEventChangedField = "category"
	PetUpdatedWebhookEventChangedFieldTags     PetUpdatedWebhookEventChangedField = "tags"
)

func (r PetUpdatedWebhookEventChangedField) IsKnown() bool {
	switch r {
	case PetUpdatedWebhookEventChangedFieldName, PetUpdatedWebhookEventChangedFieldStatus, PetUpdatedWebhookEventChangedFieldCategory, PetUpdatedWebhookEventChangedFieldTags:
		return true
	}
	return false
}

type PetUpdatedWebhookEventType string

const (
	PetUpdatedWebhookEventTypePetUpdated PetUpdatedWebhookEventType = "pet.updated"
)

func (r PetUpdatedWebhookEventType) IsKnown() bool {
	switch r {
	case PetUpdatedWebhookEventTypePetUpdated:
		return true
	}
	return false
}

type PetInventoryLowWebhookEvent struct {
	Pet       Pet                                  `json:"pet" api:"required"`
	Quantity  int64                                `json:"quantity" api:"required"`
	Threshold int64                                `json:"threshold" api:"required"`
	Type      PetInventoryLowWebhookEventType      `json:"type" api:"required"`
	LastOrder PetInventoryLowWebhookEventLastOrder `json:"last_order"`
	Locations []shared.Address                     `json:"locations"`
	JSON      petInventoryLowWebhookEventJSON      `json:"-"`
}

// petInventoryLowWebhookEventJSON contains the JSON metadata for the struct
// [PetInventoryLowWebhookEvent]
type petInventoryLowWebhookEventJSON struct {
	Pet         apijson.Field
	Quantity    apijson.Field
	Threshold   apijson.Field
	Type        apijson.Field
	LastOrder   apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetInventoryLowWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petInventoryLowWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PetInventoryLowWebhookEvent) implementsParsedWebhookEvent() {}

func (r PetInventoryLowWebhookEvent) implementsParsePetWebhookEvent() {}

type PetInventoryLowWebhookEventType string

const (
	PetInventoryLowWebhookEventTypePetInventoryLow PetInventoryLowWebhookEventType = "pet.inventory-low"
)

func (r PetInventoryLowWebhookEventType) IsKnown() bool {
	switch r {
	case PetInventoryLowWebhookEventTypePetInventoryLow:
		return true
	}
	return false
}

type PetInventoryLowWebhookEventLastOrder struct {
	ID              int64          `json:"id"`
	Complete        bool           `json:"complete"`
	PetID           int64          `json:"petId"`
	Quantity        int64          `json:"quantity"`
	RelatedAddress  shared.Address `json:"relatedAddress"`
	RelatedCategory interface{}    `json:"relatedCategory"`
	RelatedCustomer interface{}    `json:"relatedCustomer"`
	RelatedMoney    shared.Money   `json:"relatedMoney"`
	RelatedPet      Pet            `json:"relatedPet"`
	RelatedShelter  interface{}    `json:"relatedShelter"`
	RelatedTag      interface{}    `json:"relatedTag"`
	RelatedUser     interface{}    `json:"relatedUser"`
	ShipDate        time.Time      `json:"shipDate" format:"date-time"`
	// Order Status
	Status PetInventoryLowWebhookEventLastOrderStatus `json:"status"`
	Total  shared.Money                               `json:"total"`
	JSON   petInventoryLowWebhookEventLastOrderJSON   `json:"-"`
}

// petInventoryLowWebhookEventLastOrderJSON contains the JSON metadata for the
// struct [PetInventoryLowWebhookEventLastOrder]
type petInventoryLowWebhookEventLastOrderJSON struct {
	ID              apijson.Field
	Complete        apijson.Field
	PetID           apijson.Field
	Quantity        apijson.Field
	RelatedAddress  apijson.Field
	RelatedCategory apijson.Field
	RelatedCustomer apijson.Field
	RelatedMoney    apijson.Field
	RelatedPet      apijson.Field
	RelatedShelter  apijson.Field
	RelatedTag      apijson.Field
	RelatedUser     apijson.Field
	ShipDate        apijson.Field
	Status          apijson.Field
	Total           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PetInventoryLowWebhookEventLastOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petInventoryLowWebhookEventLastOrderJSON) RawJSON() string {
	return r.raw
}

// Order Status
type PetInventoryLowWebhookEventLastOrderStatus string

const (
	PetInventoryLowWebhookEventLastOrderStatusPlaced    PetInventoryLowWebhookEventLastOrderStatus = "placed"
	PetInventoryLowWebhookEventLastOrderStatusApproved  PetInventoryLowWebhookEventLastOrderStatus = "approved"
	PetInventoryLowWebhookEventLastOrderStatusDelivered PetInventoryLowWebhookEventLastOrderStatus = "delivered"
)

func (r PetInventoryLowWebhookEventLastOrderStatus) IsKnown() bool {
	switch r {
	case PetInventoryLowWebhookEventLastOrderStatusPlaced, PetInventoryLowWebhookEventLastOrderStatusApproved, PetInventoryLowWebhookEventLastOrderStatusDelivered:
		return true
	}
	return false
}

type PetModerationWebhookEvent struct {
	Pet        Pet                             `json:"pet" api:"required"`
	Type       PetModerationWebhookEventType   `json:"type" api:"required"`
	ApprovedAt time.Time                       `json:"approved_at" format:"date-time"`
	Reason     PetModerationWebhookEventReason `json:"reason"`
	// This field can have the runtime type of
	// [[]PetModerationWebhookEventPetModerationRejectedEventReviewNote].
	ReviewNotes interface{}                   `json:"review_notes"`
	JSON        petModerationWebhookEventJSON `json:"-"`
	union       PetModerationWebhookEventUnion
}

// petModerationWebhookEventJSON contains the JSON metadata for the struct
// [PetModerationWebhookEvent]
type petModerationWebhookEventJSON struct {
	Pet         apijson.Field
	Type        apijson.Field
	ApprovedAt  apijson.Field
	Reason      apijson.Field
	ReviewNotes apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r petModerationWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r *PetModerationWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	*r = PetModerationWebhookEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PetModerationWebhookEventUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PetModerationWebhookEventPetModerationApprovedEvent],
// [PetModerationWebhookEventPetModerationRejectedEvent].
func (r PetModerationWebhookEvent) AsUnion() PetModerationWebhookEventUnion {
	return r.union
}

// Union satisfied by [PetModerationWebhookEventPetModerationApprovedEvent] or
// [PetModerationWebhookEventPetModerationRejectedEvent].
type PetModerationWebhookEventUnion interface {
	implementsPetModerationWebhookEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PetModerationWebhookEventUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PetModerationWebhookEventPetModerationApprovedEvent{}),
			DiscriminatorValue: "pet.moderation.approved",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PetModerationWebhookEventPetModerationRejectedEvent{}),
			DiscriminatorValue: "pet.moderation.rejected",
		},
	)
}

type PetModerationWebhookEventPetModerationApprovedEvent struct {
	ApprovedAt time.Time                                               `json:"approved_at" api:"required" format:"date-time"`
	Pet        Pet                                                     `json:"pet" api:"required"`
	Type       PetModerationWebhookEventPetModerationApprovedEventType `json:"type" api:"required"`
	JSON       petModerationWebhookEventPetModerationApprovedEventJSON `json:"-"`
}

// petModerationWebhookEventPetModerationApprovedEventJSON contains the JSON
// metadata for the struct [PetModerationWebhookEventPetModerationApprovedEvent]
type petModerationWebhookEventPetModerationApprovedEventJSON struct {
	ApprovedAt  apijson.Field
	Pet         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetModerationWebhookEventPetModerationApprovedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petModerationWebhookEventPetModerationApprovedEventJSON) RawJSON() string {
	return r.raw
}

func (r PetModerationWebhookEventPetModerationApprovedEvent) implementsPetModerationWebhookEvent() {}

type PetModerationWebhookEventPetModerationApprovedEventType string

const (
	PetModerationWebhookEventPetModerationApprovedEventTypePetModerationApproved PetModerationWebhookEventPetModerationApprovedEventType = "pet.moderation.approved"
)

func (r PetModerationWebhookEventPetModerationApprovedEventType) IsKnown() bool {
	switch r {
	case PetModerationWebhookEventPetModerationApprovedEventTypePetModerationApproved:
		return true
	}
	return false
}

type PetModerationWebhookEventPetModerationRejectedEvent struct {
	Pet         Pet                                                             `json:"pet" api:"required"`
	Reason      PetModerationWebhookEventPetModerationRejectedEventReason       `json:"reason" api:"required"`
	Type        PetModerationWebhookEventPetModerationRejectedEventType         `json:"type" api:"required"`
	ReviewNotes []PetModerationWebhookEventPetModerationRejectedEventReviewNote `json:"review_notes"`
	JSON        petModerationWebhookEventPetModerationRejectedEventJSON         `json:"-"`
}

// petModerationWebhookEventPetModerationRejectedEventJSON contains the JSON
// metadata for the struct [PetModerationWebhookEventPetModerationRejectedEvent]
type petModerationWebhookEventPetModerationRejectedEventJSON struct {
	Pet         apijson.Field
	Reason      apijson.Field
	Type        apijson.Field
	ReviewNotes apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetModerationWebhookEventPetModerationRejectedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petModerationWebhookEventPetModerationRejectedEventJSON) RawJSON() string {
	return r.raw
}

func (r PetModerationWebhookEventPetModerationRejectedEvent) implementsPetModerationWebhookEvent() {}

type PetModerationWebhookEventPetModerationRejectedEventReason string

const (
	PetModerationWebhookEventPetModerationRejectedEventReasonPolicyViolation PetModerationWebhookEventPetModerationRejectedEventReason = "policy_violation"
	PetModerationWebhookEventPetModerationRejectedEventReasonDuplicate       PetModerationWebhookEventPetModerationRejectedEventReason = "duplicate"
	PetModerationWebhookEventPetModerationRejectedEventReasonUnsafeContent   PetModerationWebhookEventPetModerationRejectedEventReason = "unsafe-content"
)

func (r PetModerationWebhookEventPetModerationRejectedEventReason) IsKnown() bool {
	switch r {
	case PetModerationWebhookEventPetModerationRejectedEventReasonPolicyViolation, PetModerationWebhookEventPetModerationRejectedEventReasonDuplicate, PetModerationWebhookEventPetModerationRejectedEventReasonUnsafeContent:
		return true
	}
	return false
}

type PetModerationWebhookEventPetModerationRejectedEventType string

const (
	PetModerationWebhookEventPetModerationRejectedEventTypePetModerationRejected PetModerationWebhookEventPetModerationRejectedEventType = "pet.moderation.rejected"
)

func (r PetModerationWebhookEventPetModerationRejectedEventType) IsKnown() bool {
	switch r {
	case PetModerationWebhookEventPetModerationRejectedEventTypePetModerationRejected:
		return true
	}
	return false
}

type PetModerationWebhookEventPetModerationRejectedEventReviewNote struct {
	Message  string                                                                 `json:"message" api:"required"`
	Reviewer PetModerationWebhookEventPetModerationRejectedEventReviewNotesReviewer `json:"reviewer"`
	JSON     petModerationWebhookEventPetModerationRejectedEventReviewNoteJSON      `json:"-"`
}

// petModerationWebhookEventPetModerationRejectedEventReviewNoteJSON contains the
// JSON metadata for the struct
// [PetModerationWebhookEventPetModerationRejectedEventReviewNote]
type petModerationWebhookEventPetModerationRejectedEventReviewNoteJSON struct {
	Message     apijson.Field
	Reviewer    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetModerationWebhookEventPetModerationRejectedEventReviewNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petModerationWebhookEventPetModerationRejectedEventReviewNoteJSON) RawJSON() string {
	return r.raw
}

type PetModerationWebhookEventPetModerationRejectedEventReviewNotesReviewer struct {
	ID              int64          `json:"id"`
	Email           string         `json:"email"`
	FirstName       string         `json:"firstName"`
	LastName        string         `json:"lastName"`
	Password        string         `json:"password"`
	Phone           string         `json:"phone"`
	RelatedAddress  shared.Address `json:"relatedAddress"`
	RelatedCategory interface{}    `json:"relatedCategory"`
	RelatedCustomer interface{}    `json:"relatedCustomer"`
	RelatedMoney    shared.Money   `json:"relatedMoney"`
	RelatedOrder    interface{}    `json:"relatedOrder"`
	RelatedPet      Pet            `json:"relatedPet"`
	RelatedShelter  interface{}    `json:"relatedShelter"`
	RelatedTag      interface{}    `json:"relatedTag"`
	Username        string         `json:"username"`
	// User Status
	UserStatus int64                                                                      `json:"userStatus"`
	JSON       petModerationWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON `json:"-"`
}

// petModerationWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON
// contains the JSON metadata for the struct
// [PetModerationWebhookEventPetModerationRejectedEventReviewNotesReviewer]
type petModerationWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON struct {
	ID              apijson.Field
	Email           apijson.Field
	FirstName       apijson.Field
	LastName        apijson.Field
	Password        apijson.Field
	Phone           apijson.Field
	RelatedAddress  apijson.Field
	RelatedCategory apijson.Field
	RelatedCustomer apijson.Field
	RelatedMoney    apijson.Field
	RelatedOrder    apijson.Field
	RelatedPet      apijson.Field
	RelatedShelter  apijson.Field
	RelatedTag      apijson.Field
	Username        apijson.Field
	UserStatus      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PetModerationWebhookEventPetModerationRejectedEventReviewNotesReviewer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petModerationWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON) RawJSON() string {
	return r.raw
}

type PetModerationWebhookEventType string

const (
	PetModerationWebhookEventTypePetModerationApproved PetModerationWebhookEventType = "pet.moderation.approved"
	PetModerationWebhookEventTypePetModerationRejected PetModerationWebhookEventType = "pet.moderation.rejected"
)

func (r PetModerationWebhookEventType) IsKnown() bool {
	switch r {
	case PetModerationWebhookEventTypePetModerationApproved, PetModerationWebhookEventTypePetModerationRejected:
		return true
	}
	return false
}

type PetModerationWebhookEventReason string

const (
	PetModerationWebhookEventReasonPolicyViolation PetModerationWebhookEventReason = "policy_violation"
	PetModerationWebhookEventReasonDuplicate       PetModerationWebhookEventReason = "duplicate"
	PetModerationWebhookEventReasonUnsafeContent   PetModerationWebhookEventReason = "unsafe-content"
)

func (r PetModerationWebhookEventReason) IsKnown() bool {
	switch r {
	case PetModerationWebhookEventReasonPolicyViolation, PetModerationWebhookEventReasonDuplicate, PetModerationWebhookEventReasonUnsafeContent:
		return true
	}
	return false
}

type ParsePetWebhookEvent struct {
	Pet        Pet                      `json:"pet" api:"required"`
	Type       ParsePetWebhookEventType `json:"type" api:"required"`
	ApprovedAt time.Time                `json:"approved_at" format:"date-time"`
	// This field can have the runtime type of [[]PetUpdatedWebhookEventChangedField].
	ChangedFields interface{} `json:"changed_fields"`
	// This field can have the runtime type of [PetInventoryLowWebhookEventLastOrder].
	LastOrder interface{} `json:"last_order"`
	// This field can have the runtime type of [[]shared.Address].
	Locations interface{} `json:"locations"`
	// This field can have the runtime type of [map[string]string].
	Metadata interface{} `json:"metadata"`
	// pet status in the store
	PreviousStatus PetStatus                  `json:"previous_status" api:"nullable"`
	Quantity       int64                      `json:"quantity"`
	Reason         ParsePetWebhookEventReason `json:"reason"`
	// This field can have the runtime type of
	// [[]ParsePetWebhookEventPetModerationRejectedEventReviewNote].
	ReviewNotes interface{}              `json:"review_notes"`
	Threshold   int64                    `json:"threshold"`
	JSON        parsePetWebhookEventJSON `json:"-"`
	union       ParsePetWebhookEventUnion
}

// parsePetWebhookEventJSON contains the JSON metadata for the struct
// [ParsePetWebhookEvent]
type parsePetWebhookEventJSON struct {
	Pet            apijson.Field
	Type           apijson.Field
	ApprovedAt     apijson.Field
	ChangedFields  apijson.Field
	LastOrder      apijson.Field
	Locations      apijson.Field
	Metadata       apijson.Field
	PreviousStatus apijson.Field
	Quantity       apijson.Field
	Reason         apijson.Field
	ReviewNotes    apijson.Field
	Threshold      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r parsePetWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r *ParsePetWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	*r = ParsePetWebhookEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ParsePetWebhookEventUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [PetCreatedWebhookEvent],
// [PetUpdatedWebhookEvent], [PetInventoryLowWebhookEvent],
// [ParsePetWebhookEventPetModerationApprovedEvent],
// [ParsePetWebhookEventPetModerationRejectedEvent].
func (r ParsePetWebhookEvent) AsUnion() ParsePetWebhookEventUnion {
	return r.union
}

// Union satisfied by [PetCreatedWebhookEvent], [PetUpdatedWebhookEvent],
// [PetInventoryLowWebhookEvent], [ParsePetWebhookEventPetModerationApprovedEvent]
// or [ParsePetWebhookEventPetModerationRejectedEvent].
type ParsePetWebhookEventUnion interface {
	implementsParsePetWebhookEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ParsePetWebhookEventUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PetCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PetUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PetInventoryLowWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParsePetWebhookEventPetModerationApprovedEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParsePetWebhookEventPetModerationRejectedEvent{}),
		},
	)
}

type ParsePetWebhookEventPetModerationApprovedEvent struct {
	ApprovedAt time.Time                                          `json:"approved_at" api:"required" format:"date-time"`
	Pet        Pet                                                `json:"pet" api:"required"`
	Type       ParsePetWebhookEventPetModerationApprovedEventType `json:"type" api:"required"`
	JSON       parsePetWebhookEventPetModerationApprovedEventJSON `json:"-"`
}

// parsePetWebhookEventPetModerationApprovedEventJSON contains the JSON metadata
// for the struct [ParsePetWebhookEventPetModerationApprovedEvent]
type parsePetWebhookEventPetModerationApprovedEventJSON struct {
	ApprovedAt  apijson.Field
	Pet         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsePetWebhookEventPetModerationApprovedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsePetWebhookEventPetModerationApprovedEventJSON) RawJSON() string {
	return r.raw
}

func (r ParsePetWebhookEventPetModerationApprovedEvent) implementsParsePetWebhookEvent() {}

type ParsePetWebhookEventPetModerationApprovedEventType string

const (
	ParsePetWebhookEventPetModerationApprovedEventTypePetModerationApproved ParsePetWebhookEventPetModerationApprovedEventType = "pet.moderation.approved"
)

func (r ParsePetWebhookEventPetModerationApprovedEventType) IsKnown() bool {
	switch r {
	case ParsePetWebhookEventPetModerationApprovedEventTypePetModerationApproved:
		return true
	}
	return false
}

type ParsePetWebhookEventPetModerationRejectedEvent struct {
	Pet         Pet                                                        `json:"pet" api:"required"`
	Reason      ParsePetWebhookEventPetModerationRejectedEventReason       `json:"reason" api:"required"`
	Type        ParsePetWebhookEventPetModerationRejectedEventType         `json:"type" api:"required"`
	ReviewNotes []ParsePetWebhookEventPetModerationRejectedEventReviewNote `json:"review_notes"`
	JSON        parsePetWebhookEventPetModerationRejectedEventJSON         `json:"-"`
}

// parsePetWebhookEventPetModerationRejectedEventJSON contains the JSON metadata
// for the struct [ParsePetWebhookEventPetModerationRejectedEvent]
type parsePetWebhookEventPetModerationRejectedEventJSON struct {
	Pet         apijson.Field
	Reason      apijson.Field
	Type        apijson.Field
	ReviewNotes apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsePetWebhookEventPetModerationRejectedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsePetWebhookEventPetModerationRejectedEventJSON) RawJSON() string {
	return r.raw
}

func (r ParsePetWebhookEventPetModerationRejectedEvent) implementsParsePetWebhookEvent() {}

type ParsePetWebhookEventPetModerationRejectedEventReason string

const (
	ParsePetWebhookEventPetModerationRejectedEventReasonPolicyViolation ParsePetWebhookEventPetModerationRejectedEventReason = "policy_violation"
	ParsePetWebhookEventPetModerationRejectedEventReasonDuplicate       ParsePetWebhookEventPetModerationRejectedEventReason = "duplicate"
	ParsePetWebhookEventPetModerationRejectedEventReasonUnsafeContent   ParsePetWebhookEventPetModerationRejectedEventReason = "unsafe-content"
)

func (r ParsePetWebhookEventPetModerationRejectedEventReason) IsKnown() bool {
	switch r {
	case ParsePetWebhookEventPetModerationRejectedEventReasonPolicyViolation, ParsePetWebhookEventPetModerationRejectedEventReasonDuplicate, ParsePetWebhookEventPetModerationRejectedEventReasonUnsafeContent:
		return true
	}
	return false
}

type ParsePetWebhookEventPetModerationRejectedEventType string

const (
	ParsePetWebhookEventPetModerationRejectedEventTypePetModerationRejected ParsePetWebhookEventPetModerationRejectedEventType = "pet.moderation.rejected"
)

func (r ParsePetWebhookEventPetModerationRejectedEventType) IsKnown() bool {
	switch r {
	case ParsePetWebhookEventPetModerationRejectedEventTypePetModerationRejected:
		return true
	}
	return false
}

type ParsePetWebhookEventPetModerationRejectedEventReviewNote struct {
	Message  string                                                            `json:"message" api:"required"`
	Reviewer ParsePetWebhookEventPetModerationRejectedEventReviewNotesReviewer `json:"reviewer"`
	JSON     parsePetWebhookEventPetModerationRejectedEventReviewNoteJSON      `json:"-"`
}

// parsePetWebhookEventPetModerationRejectedEventReviewNoteJSON contains the JSON
// metadata for the struct
// [ParsePetWebhookEventPetModerationRejectedEventReviewNote]
type parsePetWebhookEventPetModerationRejectedEventReviewNoteJSON struct {
	Message     apijson.Field
	Reviewer    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsePetWebhookEventPetModerationRejectedEventReviewNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsePetWebhookEventPetModerationRejectedEventReviewNoteJSON) RawJSON() string {
	return r.raw
}

type ParsePetWebhookEventPetModerationRejectedEventReviewNotesReviewer struct {
	ID              int64          `json:"id"`
	Email           string         `json:"email"`
	FirstName       string         `json:"firstName"`
	LastName        string         `json:"lastName"`
	Password        string         `json:"password"`
	Phone           string         `json:"phone"`
	RelatedAddress  shared.Address `json:"relatedAddress"`
	RelatedCategory interface{}    `json:"relatedCategory"`
	RelatedCustomer interface{}    `json:"relatedCustomer"`
	RelatedMoney    shared.Money   `json:"relatedMoney"`
	RelatedOrder    interface{}    `json:"relatedOrder"`
	RelatedPet      Pet            `json:"relatedPet"`
	RelatedShelter  interface{}    `json:"relatedShelter"`
	RelatedTag      interface{}    `json:"relatedTag"`
	Username        string         `json:"username"`
	// User Status
	UserStatus int64                                                                 `json:"userStatus"`
	JSON       parsePetWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON `json:"-"`
}

// parsePetWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON contains
// the JSON metadata for the struct
// [ParsePetWebhookEventPetModerationRejectedEventReviewNotesReviewer]
type parsePetWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON struct {
	ID              apijson.Field
	Email           apijson.Field
	FirstName       apijson.Field
	LastName        apijson.Field
	Password        apijson.Field
	Phone           apijson.Field
	RelatedAddress  apijson.Field
	RelatedCategory apijson.Field
	RelatedCustomer apijson.Field
	RelatedMoney    apijson.Field
	RelatedOrder    apijson.Field
	RelatedPet      apijson.Field
	RelatedShelter  apijson.Field
	RelatedTag      apijson.Field
	Username        apijson.Field
	UserStatus      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ParsePetWebhookEventPetModerationRejectedEventReviewNotesReviewer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsePetWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON) RawJSON() string {
	return r.raw
}

type ParsePetWebhookEventType string

const (
	ParsePetWebhookEventTypePetCreated            ParsePetWebhookEventType = "pet.created"
	ParsePetWebhookEventTypePetUpdated            ParsePetWebhookEventType = "pet.updated"
	ParsePetWebhookEventTypePetInventoryLow       ParsePetWebhookEventType = "pet.inventory-low"
	ParsePetWebhookEventTypePetModerationApproved ParsePetWebhookEventType = "pet.moderation.approved"
	ParsePetWebhookEventTypePetModerationRejected ParsePetWebhookEventType = "pet.moderation.rejected"
)

func (r ParsePetWebhookEventType) IsKnown() bool {
	switch r {
	case ParsePetWebhookEventTypePetCreated, ParsePetWebhookEventTypePetUpdated, ParsePetWebhookEventTypePetInventoryLow, ParsePetWebhookEventTypePetModerationApproved, ParsePetWebhookEventTypePetModerationRejected:
		return true
	}
	return false
}

type ParsePetWebhookEventReason string

const (
	ParsePetWebhookEventReasonPolicyViolation ParsePetWebhookEventReason = "policy_violation"
	ParsePetWebhookEventReasonDuplicate       ParsePetWebhookEventReason = "duplicate"
	ParsePetWebhookEventReasonUnsafeContent   ParsePetWebhookEventReason = "unsafe-content"
)

func (r ParsePetWebhookEventReason) IsKnown() bool {
	switch r {
	case ParsePetWebhookEventReasonPolicyViolation, ParsePetWebhookEventReasonDuplicate, ParsePetWebhookEventReasonUnsafeContent:
		return true
	}
	return false
}
