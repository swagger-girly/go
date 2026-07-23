// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsConditionValueUnionParam() {}
func (UnionTime) ImplementsConditionValueUnion()      {}

type UnionString string

func (UnionString) ImplementsSystemHealthServicesLastIncidentResolutionUnion() {}
func (UnionString) ImplementsPetMicrochipIDUnionParam()                        {}
func (UnionString) ImplementsPetMicrochipIDUnion()                             {}
func (UnionString) ImplementsConditionValueUnionParam()                        {}
func (UnionString) ImplementsConditionValueUnion()                             {}
func (UnionString) ImplementsPolicyNotificationTargetUnion()                   {}
func (UnionString) ImplementsAdoptionPolicyNewParamsNotificationTargetUnion()  {}
func (UnionString) ImplementsAdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetUnion() {
}
func (UnionString) ImplementsVaccinationRecordBoosterScheduleDoseUnionParam()                    {}
func (UnionString) ImplementsVaccinationRecordBoosterScheduleDoseUnion()                         {}
func (UnionString) ImplementsMedicalSummaryConditionsTreatmentMedicationsDosageUnion()           {}
func (UnionString) ImplementsAdoptionsPolicyChangedWebhookEventPreviousNotificationTargetUnion() {}

type UnionInt int64

func (UnionInt) ImplementsPetMicrochipIDUnionParam()                                 {}
func (UnionInt) ImplementsPetMicrochipIDUnion()                                      {}
func (UnionInt) ImplementsPetGetPremiumResponseCoverageLimitUnion()                  {}
func (UnionInt) ImplementsConditionValueUnionParam()                                 {}
func (UnionInt) ImplementsConditionValueUnion()                                      {}
func (UnionInt) ImplementsPlacementActivityPlacementDisruptionEventSeverityUnion()   {}
func (UnionInt) ImplementsPlacementEventPlacementDisruptionEventSeverityUnionParam() {}
func (UnionInt) ImplementsPlacementEventPlacementDisruptionEventSeverityUnion()      {}

type UnionFloat float64

func (UnionFloat) ImplementsConditionValueUnionParam()                       {}
func (UnionFloat) ImplementsConditionValueUnion()                            {}
func (UnionFloat) ImplementsVaccinationRecordBoosterScheduleDoseUnionParam() {}
func (UnionFloat) ImplementsVaccinationRecordBoosterScheduleDoseUnion()      {}
