// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"errors"
	"net/http"
	"slices"

	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
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

type ParsedWebhookEvent struct {
	Pet  Pet                    `json:"pet" api:"required"`
	Type ParsedWebhookEventType `json:"type" api:"required"`
	JSON parsedWebhookEventJSON `json:"-"`
}

// parsedWebhookEventJSON contains the JSON metadata for the struct
// [ParsedWebhookEvent]
type parsedWebhookEventJSON struct {
	Pet         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParsedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parsedWebhookEventJSON) RawJSON() string {
	return r.raw
}

type ParsedWebhookEventType string

const (
	ParsedWebhookEventTypePetCreated ParsedWebhookEventType = "pet.created"
)

func (r ParsedWebhookEventType) IsKnown() bool {
	switch r {
	case ParsedWebhookEventTypePetCreated:
		return true
	}
	return false
}
