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
	"github.com/tidwall/gjson"
)

// WebhookService contains methods and other services that help with interacting
// with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

func (r *WebhookService) Parsed(payload []byte, headers http.Header, opts ...option.RequestOption) (*ParsedWebhookEvent, error) {
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
	res := &ParsedWebhookEvent{}
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
	ChangedFields  []PetUpdatedWebhookEventChangedField `json:"changed_fields" api:"required"`
	Pet            Pet                                  `json:"pet" api:"required"`
	Type           PetUpdatedWebhookEventType           `json:"type" api:"required"`
	Metadata       map[string]string                    `json:"metadata"`
	PreviousStatus PetUpdatedWebhookEventPreviousStatus `json:"previous_status" api:"nullable"`
	JSON           petUpdatedWebhookEventJSON           `json:"-"`
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

type PetUpdatedWebhookEventPreviousStatus string

const (
	PetUpdatedWebhookEventPreviousStatusAvailable PetUpdatedWebhookEventPreviousStatus = "available"
	PetUpdatedWebhookEventPreviousStatusPending   PetUpdatedWebhookEventPreviousStatus = "pending"
	PetUpdatedWebhookEventPreviousStatusSold      PetUpdatedWebhookEventPreviousStatus = "sold"
)

func (r PetUpdatedWebhookEventPreviousStatus) IsKnown() bool {
	switch r {
	case PetUpdatedWebhookEventPreviousStatusAvailable, PetUpdatedWebhookEventPreviousStatusPending, PetUpdatedWebhookEventPreviousStatusSold:
		return true
	}
	return false
}

type PetInventoryLowWebhookEvent struct {
	Pet       Pet                                   `json:"pet" api:"required"`
	Quantity  int64                                 `json:"quantity" api:"required"`
	Threshold int64                                 `json:"threshold" api:"required"`
	Type      PetInventoryLowWebhookEventType       `json:"type" api:"required"`
	LastOrder PetInventoryLowWebhookEventLastOrder  `json:"last_order"`
	Locations []PetInventoryLowWebhookEventLocation `json:"locations"`
	JSON      petInventoryLowWebhookEventJSON       `json:"-"`
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
	ID       int64     `json:"id"`
	Complete bool      `json:"complete"`
	PetID    int64     `json:"petId"`
	Quantity int64     `json:"quantity"`
	ShipDate time.Time `json:"shipDate" format:"date-time"`
	// Order Status
	Status PetInventoryLowWebhookEventLastOrderStatus `json:"status"`
	JSON   petInventoryLowWebhookEventLastOrderJSON   `json:"-"`
}

// petInventoryLowWebhookEventLastOrderJSON contains the JSON metadata for the
// struct [PetInventoryLowWebhookEventLastOrder]
type petInventoryLowWebhookEventLastOrderJSON struct {
	ID          apijson.Field
	Complete    apijson.Field
	PetID       apijson.Field
	Quantity    apijson.Field
	ShipDate    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

type PetInventoryLowWebhookEventLocation struct {
	City   string                                  `json:"city"`
	State  string                                  `json:"state"`
	Street string                                  `json:"street"`
	Zip    string                                  `json:"zip"`
	JSON   petInventoryLowWebhookEventLocationJSON `json:"-"`
}

// petInventoryLowWebhookEventLocationJSON contains the JSON metadata for the
// struct [PetInventoryLowWebhookEventLocation]
type petInventoryLowWebhookEventLocationJSON struct {
	City        apijson.Field
	State       apijson.Field
	Street      apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetInventoryLowWebhookEventLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petInventoryLowWebhookEventLocationJSON) RawJSON() string {
	return r.raw
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
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
	// User Status
	UserStatus int64                                                                      `json:"userStatus"`
	JSON       petModerationWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON `json:"-"`
}

// petModerationWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON
// contains the JSON metadata for the struct
// [PetModerationWebhookEventPetModerationRejectedEventReviewNotesReviewer]
type petModerationWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	Password    apijson.Field
	Phone       apijson.Field
	Username    apijson.Field
	UserStatus  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

type StoreReportGeneratedWebhookEvent struct {
	ID        string                               `json:"id" api:"required"`
	CreatedAt time.Time                            `json:"created_at" api:"required" format:"date-time"`
	Metrics   interface{}                          `json:"metrics" api:"required"`
	Report    Report                               `json:"report" api:"required"`
	Type      StoreReportGeneratedWebhookEventType `json:"type" api:"required"`
	JSON      storeReportGeneratedWebhookEventJSON `json:"-"`
}

// storeReportGeneratedWebhookEventJSON contains the JSON metadata for the struct
// [StoreReportGeneratedWebhookEvent]
type storeReportGeneratedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Metrics     apijson.Field
	Report      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreReportGeneratedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeReportGeneratedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r StoreReportGeneratedWebhookEvent) implementsParsedWebhookEvent() {}

type StoreReportGeneratedWebhookEventType string

const (
	StoreReportGeneratedWebhookEventTypeStoreReportGenerated StoreReportGeneratedWebhookEventType = "store.report.generated"
)

func (r StoreReportGeneratedWebhookEventType) IsKnown() bool {
	switch r {
	case StoreReportGeneratedWebhookEventTypeStoreReportGenerated:
		return true
	}
	return false
}

type ParsedWebhookEvent struct {
	Type       ParsedWebhookEventType `json:"type" api:"required"`
	ID         string                 `json:"id"`
	ApprovedAt time.Time              `json:"approved_at" format:"date-time"`
	// This field can have the runtime type of [[]PetUpdatedWebhookEventChangedField].
	ChangedFields interface{} `json:"changed_fields"`
	CreatedAt     time.Time   `json:"created_at" format:"date-time"`
	// This field can have the runtime type of [PetInventoryLowWebhookEventLastOrder].
	LastOrder interface{} `json:"last_order"`
	// This field can have the runtime type of [[]PetInventoryLowWebhookEventLocation].
	Locations interface{} `json:"locations"`
	// This field can have the runtime type of [map[string]string].
	Metadata interface{} `json:"metadata"`
	// This field can have the runtime type of [interface{}].
	Metrics        interface{}                      `json:"metrics"`
	Pet            Pet                              `json:"pet"`
	PreviousStatus ParsedWebhookEventPreviousStatus `json:"previous_status" api:"nullable"`
	Quantity       int64                            `json:"quantity"`
	Reason         ParsedWebhookEventReason         `json:"reason"`
	Report         Report                           `json:"report"`
	// This field can have the runtime type of
	// [[]ParsedWebhookEventPetModerationRejectedEventReviewNote].
	ReviewNotes interface{}            `json:"review_notes"`
	Threshold   int64                  `json:"threshold"`
	JSON        parsedWebhookEventJSON `json:"-"`
	union       ParsedWebhookEventUnion
}

// parsedWebhookEventJSON contains the JSON metadata for the struct
// [ParsedWebhookEvent]
type parsedWebhookEventJSON struct {
	Type           apijson.Field
	ID             apijson.Field
	ApprovedAt     apijson.Field
	ChangedFields  apijson.Field
	CreatedAt      apijson.Field
	LastOrder      apijson.Field
	Locations      apijson.Field
	Metadata       apijson.Field
	Metrics        apijson.Field
	Pet            apijson.Field
	PreviousStatus apijson.Field
	Quantity       apijson.Field
	Reason         apijson.Field
	Report         apijson.Field
	ReviewNotes    apijson.Field
	Threshold      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r parsedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r *ParsedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	*r = ParsedWebhookEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ParsedWebhookEventUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [PetCreatedWebhookEvent],
// [PetUpdatedWebhookEvent], [PetInventoryLowWebhookEvent],
// [ParsedWebhookEventPetModerationApprovedEvent],
// [ParsedWebhookEventPetModerationRejectedEvent],
// [StoreReportGeneratedWebhookEvent].
func (r ParsedWebhookEvent) AsUnion() ParsedWebhookEventUnion {
	return r.union
}

// Union satisfied by [PetCreatedWebhookEvent], [PetUpdatedWebhookEvent],
// [PetInventoryLowWebhookEvent], [ParsedWebhookEventPetModerationApprovedEvent],
// [ParsedWebhookEventPetModerationRejectedEvent] or
// [StoreReportGeneratedWebhookEvent].
type ParsedWebhookEventUnion interface {
	implementsParsedWebhookEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ParsedWebhookEventUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(ParsedWebhookEventPetModerationApprovedEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParsedWebhookEventPetModerationRejectedEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StoreReportGeneratedWebhookEvent{}),
		},
	)
}

type ParsedWebhookEventPetModerationApprovedEvent struct {
	ApprovedAt time.Time                                        `json:"approved_at" api:"required" format:"date-time"`
	Pet        Pet                                              `json:"pet" api:"required"`
	Type       ParsedWebhookEventPetModerationApprovedEventType `json:"type" api:"required"`
	JSON       parsedWebhookEventPetModerationApprovedEventJSON `json:"-"`
}

// parsedWebhookEventPetModerationApprovedEventJSON contains the JSON metadata for
// the struct [ParsedWebhookEventPetModerationApprovedEvent]
type parsedWebhookEventPetModerationApprovedEventJSON struct {
	ApprovedAt  apijson.Field
	Pet         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEventPetModerationApprovedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventPetModerationApprovedEventJSON) RawJSON() string {
	return r.raw
}

func (r ParsedWebhookEventPetModerationApprovedEvent) implementsParsedWebhookEvent() {}

type ParsedWebhookEventPetModerationApprovedEventType string

const (
	ParsedWebhookEventPetModerationApprovedEventTypePetModerationApproved ParsedWebhookEventPetModerationApprovedEventType = "pet.moderation.approved"
)

func (r ParsedWebhookEventPetModerationApprovedEventType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventPetModerationApprovedEventTypePetModerationApproved:
		return true
	}
	return false
}

type ParsedWebhookEventPetModerationRejectedEvent struct {
	Pet         Pet                                                      `json:"pet" api:"required"`
	Reason      ParsedWebhookEventPetModerationRejectedEventReason       `json:"reason" api:"required"`
	Type        ParsedWebhookEventPetModerationRejectedEventType         `json:"type" api:"required"`
	ReviewNotes []ParsedWebhookEventPetModerationRejectedEventReviewNote `json:"review_notes"`
	JSON        parsedWebhookEventPetModerationRejectedEventJSON         `json:"-"`
}

// parsedWebhookEventPetModerationRejectedEventJSON contains the JSON metadata for
// the struct [ParsedWebhookEventPetModerationRejectedEvent]
type parsedWebhookEventPetModerationRejectedEventJSON struct {
	Pet         apijson.Field
	Reason      apijson.Field
	Type        apijson.Field
	ReviewNotes apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEventPetModerationRejectedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventPetModerationRejectedEventJSON) RawJSON() string {
	return r.raw
}

func (r ParsedWebhookEventPetModerationRejectedEvent) implementsParsedWebhookEvent() {}

type ParsedWebhookEventPetModerationRejectedEventReason string

const (
	ParsedWebhookEventPetModerationRejectedEventReasonPolicyViolation ParsedWebhookEventPetModerationRejectedEventReason = "policy_violation"
	ParsedWebhookEventPetModerationRejectedEventReasonDuplicate       ParsedWebhookEventPetModerationRejectedEventReason = "duplicate"
	ParsedWebhookEventPetModerationRejectedEventReasonUnsafeContent   ParsedWebhookEventPetModerationRejectedEventReason = "unsafe-content"
)

func (r ParsedWebhookEventPetModerationRejectedEventReason) IsKnown() bool {
	switch r {
	case ParsedWebhookEventPetModerationRejectedEventReasonPolicyViolation, ParsedWebhookEventPetModerationRejectedEventReasonDuplicate, ParsedWebhookEventPetModerationRejectedEventReasonUnsafeContent:
		return true
	}
	return false
}

type ParsedWebhookEventPetModerationRejectedEventType string

const (
	ParsedWebhookEventPetModerationRejectedEventTypePetModerationRejected ParsedWebhookEventPetModerationRejectedEventType = "pet.moderation.rejected"
)

func (r ParsedWebhookEventPetModerationRejectedEventType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventPetModerationRejectedEventTypePetModerationRejected:
		return true
	}
	return false
}

type ParsedWebhookEventPetModerationRejectedEventReviewNote struct {
	Message  string                                                          `json:"message" api:"required"`
	Reviewer ParsedWebhookEventPetModerationRejectedEventReviewNotesReviewer `json:"reviewer"`
	JSON     parsedWebhookEventPetModerationRejectedEventReviewNoteJSON      `json:"-"`
}

// parsedWebhookEventPetModerationRejectedEventReviewNoteJSON contains the JSON
// metadata for the struct [ParsedWebhookEventPetModerationRejectedEventReviewNote]
type parsedWebhookEventPetModerationRejectedEventReviewNoteJSON struct {
	Message     apijson.Field
	Reviewer    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEventPetModerationRejectedEventReviewNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventPetModerationRejectedEventReviewNoteJSON) RawJSON() string {
	return r.raw
}

type ParsedWebhookEventPetModerationRejectedEventReviewNotesReviewer struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
	// User Status
	UserStatus int64                                                               `json:"userStatus"`
	JSON       parsedWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON `json:"-"`
}

// parsedWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON contains the
// JSON metadata for the struct
// [ParsedWebhookEventPetModerationRejectedEventReviewNotesReviewer]
type parsedWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	Password    apijson.Field
	Phone       apijson.Field
	Username    apijson.Field
	UserStatus  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEventPetModerationRejectedEventReviewNotesReviewer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventPetModerationRejectedEventReviewNotesReviewerJSON) RawJSON() string {
	return r.raw
}

type ParsedWebhookEventType string

const (
	ParsedWebhookEventTypePetCreated            ParsedWebhookEventType = "pet.created"
	ParsedWebhookEventTypePetUpdated            ParsedWebhookEventType = "pet.updated"
	ParsedWebhookEventTypePetInventoryLow       ParsedWebhookEventType = "pet.inventory-low"
	ParsedWebhookEventTypePetModerationApproved ParsedWebhookEventType = "pet.moderation.approved"
	ParsedWebhookEventTypePetModerationRejected ParsedWebhookEventType = "pet.moderation.rejected"
	ParsedWebhookEventTypeStoreReportGenerated  ParsedWebhookEventType = "store.report.generated"
)

func (r ParsedWebhookEventType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventTypePetCreated, ParsedWebhookEventTypePetUpdated, ParsedWebhookEventTypePetInventoryLow, ParsedWebhookEventTypePetModerationApproved, ParsedWebhookEventTypePetModerationRejected, ParsedWebhookEventTypeStoreReportGenerated:
		return true
	}
	return false
}

type ParsedWebhookEventPreviousStatus string

const (
	ParsedWebhookEventPreviousStatusAvailable ParsedWebhookEventPreviousStatus = "available"
	ParsedWebhookEventPreviousStatusPending   ParsedWebhookEventPreviousStatus = "pending"
	ParsedWebhookEventPreviousStatusSold      ParsedWebhookEventPreviousStatus = "sold"
)

func (r ParsedWebhookEventPreviousStatus) IsKnown() bool {
	switch r {
	case ParsedWebhookEventPreviousStatusAvailable, ParsedWebhookEventPreviousStatusPending, ParsedWebhookEventPreviousStatusSold:
		return true
	}
	return false
}

type ParsedWebhookEventReason string

const (
	ParsedWebhookEventReasonPolicyViolation ParsedWebhookEventReason = "policy_violation"
	ParsedWebhookEventReasonDuplicate       ParsedWebhookEventReason = "duplicate"
	ParsedWebhookEventReasonUnsafeContent   ParsedWebhookEventReason = "unsafe-content"
)

func (r ParsedWebhookEventReason) IsKnown() bool {
	switch r {
	case ParsedWebhookEventReasonPolicyViolation, ParsedWebhookEventReasonDuplicate, ParsedWebhookEventReasonUnsafeContent:
		return true
	}
	return false
}
