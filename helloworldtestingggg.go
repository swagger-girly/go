// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"reflect"
	"time"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/shared"
	"github.com/tidwall/gjson"
)

type SystemHealth struct {
	CheckedAt time.Time          `json:"checkedAt" api:"required" format:"date-time"`
	Status    SystemHealthStatus `json:"status" api:"required"`
	// Per-service health keyed by service name.
	Services map[string]SystemHealthService `json:"services"`
	JSON     systemHealthJSON               `json:"-"`
}

// systemHealthJSON contains the JSON metadata for the struct [SystemHealth]
type systemHealthJSON struct {
	CheckedAt   apijson.Field
	Status      apijson.Field
	Services    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SystemHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemHealthJSON) RawJSON() string {
	return r.raw
}

type SystemHealthStatus string

const (
	SystemHealthStatusOk       SystemHealthStatus = "ok"
	SystemHealthStatusDegraded SystemHealthStatus = "degraded"
	SystemHealthStatusDown     SystemHealthStatus = "down"
)

func (r SystemHealthStatus) IsKnown() bool {
	switch r {
	case SystemHealthStatusOk, SystemHealthStatusDegraded, SystemHealthStatusDown:
		return true
	}
	return false
}

type SystemHealthService struct {
	Status       SystemHealthServicesStatus       `json:"status" api:"required"`
	LastIncident SystemHealthServicesLastIncident `json:"lastIncident"`
	LatencyMs    int64                            `json:"latencyMs"`
	JSON         systemHealthServiceJSON          `json:"-"`
}

// systemHealthServiceJSON contains the JSON metadata for the struct
// [SystemHealthService]
type systemHealthServiceJSON struct {
	Status       apijson.Field
	LastIncident apijson.Field
	LatencyMs    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SystemHealthService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemHealthServiceJSON) RawJSON() string {
	return r.raw
}

type SystemHealthServicesStatus string

const (
	SystemHealthServicesStatusOk       SystemHealthServicesStatus = "ok"
	SystemHealthServicesStatusDegraded SystemHealthServicesStatus = "degraded"
	SystemHealthServicesStatusDown     SystemHealthServicesStatus = "down"
)

func (r SystemHealthServicesStatus) IsKnown() bool {
	switch r {
	case SystemHealthServicesStatusOk, SystemHealthServicesStatusDegraded, SystemHealthServicesStatusDown:
		return true
	}
	return false
}

type SystemHealthServicesLastIncident struct {
	OpenedAt   time.Time                                       `json:"openedAt" format:"date-time"`
	Resolution SystemHealthServicesLastIncidentResolutionUnion `json:"resolution"`
	JSON       systemHealthServicesLastIncidentJSON            `json:"-"`
}

// systemHealthServicesLastIncidentJSON contains the JSON metadata for the struct
// [SystemHealthServicesLastIncident]
type systemHealthServicesLastIncidentJSON struct {
	OpenedAt    apijson.Field
	Resolution  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SystemHealthServicesLastIncident) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemHealthServicesLastIncidentJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [SystemHealthServicesLastIncidentResolutionIncidentReference].
type SystemHealthServicesLastIncidentResolutionUnion interface {
	ImplementsSystemHealthServicesLastIncidentResolutionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SystemHealthServicesLastIncidentResolutionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SystemHealthServicesLastIncidentResolutionIncidentReference{}),
		},
	)
}

type SystemHealthServicesLastIncidentResolutionIncidentReference struct {
	IncidentID string                                                          `json:"incidentId" api:"required"`
	URL        string                                                          `json:"url" format:"uri"`
	JSON       systemHealthServicesLastIncidentResolutionIncidentReferenceJSON `json:"-"`
}

// systemHealthServicesLastIncidentResolutionIncidentReferenceJSON contains the
// JSON metadata for the struct
// [SystemHealthServicesLastIncidentResolutionIncidentReference]
type systemHealthServicesLastIncidentResolutionIncidentReferenceJSON struct {
	IncidentID  apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SystemHealthServicesLastIncidentResolutionIncidentReference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemHealthServicesLastIncidentResolutionIncidentReferenceJSON) RawJSON() string {
	return r.raw
}

func (r SystemHealthServicesLastIncidentResolutionIncidentReference) ImplementsSystemHealthServicesLastIncidentResolutionUnion() {
}

type GetRateLimitsResponse struct {
	Remaining int64                       `json:"remaining" api:"required"`
	ResetAt   time.Time                   `json:"resetAt" api:"required" format:"date-time"`
	Tiers     []GetRateLimitsResponseTier `json:"tiers"`
	JSON      getRateLimitsResponseJSON   `json:"-"`
}

// getRateLimitsResponseJSON contains the JSON metadata for the struct
// [GetRateLimitsResponse]
type getRateLimitsResponseJSON struct {
	Remaining   apijson.Field
	ResetAt     apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetRateLimitsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getRateLimitsResponseJSON) RawJSON() string {
	return r.raw
}

type GetRateLimitsResponseTier struct {
	Name  string                          `json:"name" api:"required"`
	Quota GetRateLimitsResponseTiersQuota `json:"quota"`
	JSON  getRateLimitsResponseTierJSON   `json:"-"`
}

// getRateLimitsResponseTierJSON contains the JSON metadata for the struct
// [GetRateLimitsResponseTier]
type getRateLimitsResponseTierJSON struct {
	Name        apijson.Field
	Quota       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetRateLimitsResponseTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getRateLimitsResponseTierJSON) RawJSON() string {
	return r.raw
}

type GetRateLimitsResponseTiersQuota struct {
	Requests int64                               `json:"requests"`
	Window   string                              `json:"window"`
	JSON     getRateLimitsResponseTiersQuotaJSON `json:"-"`
}

// getRateLimitsResponseTiersQuotaJSON contains the JSON metadata for the struct
// [GetRateLimitsResponseTiersQuota]
type getRateLimitsResponseTiersQuotaJSON struct {
	Requests    apijson.Field
	Window      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetRateLimitsResponseTiersQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getRateLimitsResponseTiersQuotaJSON) RawJSON() string {
	return r.raw
}
