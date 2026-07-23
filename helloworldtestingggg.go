// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"encoding/json"
	"time"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/packages/respjson"
)

type SystemHealth struct {
	CheckedAt time.Time `json:"checkedAt" api:"required" format:"date-time"`
	// Any of "ok", "degraded", "down".
	Status SystemHealthStatus `json:"status" api:"required"`
	// Per-service health keyed by service name.
	Services map[string]SystemHealthService `json:"services"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckedAt   respjson.Field
		Status      respjson.Field
		Services    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SystemHealth) RawJSON() string { return r.JSON.raw }
func (r *SystemHealth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemHealthStatus string

const (
	SystemHealthStatusOk       SystemHealthStatus = "ok"
	SystemHealthStatusDegraded SystemHealthStatus = "degraded"
	SystemHealthStatusDown     SystemHealthStatus = "down"
)

type SystemHealthService struct {
	// Any of "ok", "degraded", "down".
	Status       string                          `json:"status" api:"required"`
	LastIncident SystemHealthServiceLastIncident `json:"lastIncident"`
	LatencyMs    int64                           `json:"latencyMs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status       respjson.Field
		LastIncident respjson.Field
		LatencyMs    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SystemHealthService) RawJSON() string { return r.JSON.raw }
func (r *SystemHealthService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemHealthServiceLastIncident struct {
	OpenedAt   time.Time                                      `json:"openedAt" format:"date-time"`
	Resolution SystemHealthServiceLastIncidentResolutionUnion `json:"resolution"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpenedAt    respjson.Field
		Resolution  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SystemHealthServiceLastIncident) RawJSON() string { return r.JSON.raw }
func (r *SystemHealthServiceLastIncident) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SystemHealthServiceLastIncidentResolutionUnion contains all possible properties
// and values from [string],
// [SystemHealthServiceLastIncidentResolutionIncidentReference].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type SystemHealthServiceLastIncidentResolutionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [SystemHealthServiceLastIncidentResolutionIncidentReference].
	IncidentID string `json:"incidentId"`
	// This field is from variant
	// [SystemHealthServiceLastIncidentResolutionIncidentReference].
	URL  string `json:"url"`
	JSON struct {
		OfString   respjson.Field
		IncidentID respjson.Field
		URL        respjson.Field
		raw        string
	} `json:"-"`
}

func (u SystemHealthServiceLastIncidentResolutionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SystemHealthServiceLastIncidentResolutionUnion) AsIncidentReference() (v SystemHealthServiceLastIncidentResolutionIncidentReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SystemHealthServiceLastIncidentResolutionUnion) RawJSON() string { return u.JSON.raw }

func (r *SystemHealthServiceLastIncidentResolutionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemHealthServiceLastIncidentResolutionIncidentReference struct {
	IncidentID string `json:"incidentId" api:"required"`
	URL        string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncidentID  respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SystemHealthServiceLastIncidentResolutionIncidentReference) RawJSON() string {
	return r.JSON.raw
}
func (r *SystemHealthServiceLastIncidentResolutionIncidentReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetRateLimitsResponse struct {
	Remaining int64                       `json:"remaining" api:"required"`
	ResetAt   time.Time                   `json:"resetAt" api:"required" format:"date-time"`
	Tiers     []GetRateLimitsResponseTier `json:"tiers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Remaining   respjson.Field
		ResetAt     respjson.Field
		Tiers       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetRateLimitsResponse) RawJSON() string { return r.JSON.raw }
func (r *GetRateLimitsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetRateLimitsResponseTier struct {
	Name  string                         `json:"name" api:"required"`
	Quota GetRateLimitsResponseTierQuota `json:"quota"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Quota       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetRateLimitsResponseTier) RawJSON() string { return r.JSON.raw }
func (r *GetRateLimitsResponseTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetRateLimitsResponseTierQuota struct {
	Requests int64  `json:"requests"`
	Window   string `json:"window"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Requests    respjson.Field
		Window      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetRateLimitsResponseTierQuota) RawJSON() string { return r.JSON.raw }
func (r *GetRateLimitsResponseTierQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
