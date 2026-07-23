// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"reflect"
	"time"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/param"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/shared"
	"github.com/tidwall/gjson"
)

// VeterinaryService contains methods and other services that help with interacting
// with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVeterinaryService] method instead.
type VeterinaryService struct {
	Options []option.RequestOption
}

// NewVeterinaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVeterinaryService(opts ...option.RequestOption) (r *VeterinaryService) {
	r = &VeterinaryService{}
	r.Options = opts
	return
}

type MedicalSummary struct {
	PetID int64 `json:"petId" api:"required"`
	// Provenance stamp attached to records across unrelated resources.
	Audit        MedicalSummaryAudit       `json:"audit"`
	Conditions   []MedicalSummaryCondition `json:"conditions"`
	Vaccinations []VaccinationRecord       `json:"vaccinations"`
	JSON         medicalSummaryJSON        `json:"-"`
}

// medicalSummaryJSON contains the JSON metadata for the struct [MedicalSummary]
type medicalSummaryJSON struct {
	PetID        apijson.Field
	Audit        apijson.Field
	Conditions   apijson.Field
	Vaccinations apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MedicalSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r medicalSummaryJSON) RawJSON() string {
	return r.raw
}

// Provenance stamp attached to records across unrelated resources.
type MedicalSummaryAudit struct {
	RecordedAt time.Time                 `json:"recordedAt" api:"required" format:"date-time"`
	RecordedBy string                    `json:"recordedBy"`
	Source     MedicalSummaryAuditSource `json:"source"`
	JSON       medicalSummaryAuditJSON   `json:"-"`
}

// medicalSummaryAuditJSON contains the JSON metadata for the struct
// [MedicalSummaryAudit]
type medicalSummaryAuditJSON struct {
	RecordedAt  apijson.Field
	RecordedBy  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MedicalSummaryAudit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r medicalSummaryAuditJSON) RawJSON() string {
	return r.raw
}

type MedicalSummaryAuditSource string

const (
	MedicalSummaryAuditSourceManual MedicalSummaryAuditSource = "manual"
	MedicalSummaryAuditSourceImport MedicalSummaryAuditSource = "import"
	MedicalSummaryAuditSourceSystem MedicalSummaryAuditSource = "system"
)

func (r MedicalSummaryAuditSource) IsKnown() bool {
	switch r {
	case MedicalSummaryAuditSourceManual, MedicalSummaryAuditSourceImport, MedicalSummaryAuditSourceSystem:
		return true
	}
	return false
}

type MedicalSummaryCondition struct {
	Code      string                            `json:"code" api:"required"`
	Severity  MedicalSummaryConditionsSeverity  `json:"severity"`
	Treatment MedicalSummaryConditionsTreatment `json:"treatment"`
	JSON      medicalSummaryConditionJSON       `json:"-"`
}

// medicalSummaryConditionJSON contains the JSON metadata for the struct
// [MedicalSummaryCondition]
type medicalSummaryConditionJSON struct {
	Code        apijson.Field
	Severity    apijson.Field
	Treatment   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MedicalSummaryCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r medicalSummaryConditionJSON) RawJSON() string {
	return r.raw
}

type MedicalSummaryConditionsSeverity string

const (
	MedicalSummaryConditionsSeverityMild     MedicalSummaryConditionsSeverity = "mild"
	MedicalSummaryConditionsSeverityModerate MedicalSummaryConditionsSeverity = "moderate"
	MedicalSummaryConditionsSeveritySevere   MedicalSummaryConditionsSeverity = "severe"
)

func (r MedicalSummaryConditionsSeverity) IsKnown() bool {
	switch r {
	case MedicalSummaryConditionsSeverityMild, MedicalSummaryConditionsSeverityModerate, MedicalSummaryConditionsSeveritySevere:
		return true
	}
	return false
}

type MedicalSummaryConditionsTreatment struct {
	Medications []MedicalSummaryConditionsTreatmentMedication `json:"medications"`
	Plan        string                                        `json:"plan"`
	JSON        medicalSummaryConditionsTreatmentJSON         `json:"-"`
}

// medicalSummaryConditionsTreatmentJSON contains the JSON metadata for the struct
// [MedicalSummaryConditionsTreatment]
type medicalSummaryConditionsTreatmentJSON struct {
	Medications apijson.Field
	Plan        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MedicalSummaryConditionsTreatment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r medicalSummaryConditionsTreatmentJSON) RawJSON() string {
	return r.raw
}

type MedicalSummaryConditionsTreatmentMedication struct {
	Name string `json:"name" api:"required"`
	// Free-form dosage text or a structured dosage.
	Dosage MedicalSummaryConditionsTreatmentMedicationsDosageUnion `json:"dosage"`
	JSON   medicalSummaryConditionsTreatmentMedicationJSON         `json:"-"`
}

// medicalSummaryConditionsTreatmentMedicationJSON contains the JSON metadata for
// the struct [MedicalSummaryConditionsTreatmentMedication]
type medicalSummaryConditionsTreatmentMedicationJSON struct {
	Name        apijson.Field
	Dosage      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MedicalSummaryConditionsTreatmentMedication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r medicalSummaryConditionsTreatmentMedicationJSON) RawJSON() string {
	return r.raw
}

// Free-form dosage text or a structured dosage.
//
// Union satisfied by [shared.UnionString] or
// [MedicalSummaryConditionsTreatmentMedicationsDosageStructuredDosage].
type MedicalSummaryConditionsTreatmentMedicationsDosageUnion interface {
	ImplementsMedicalSummaryConditionsTreatmentMedicationsDosageUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MedicalSummaryConditionsTreatmentMedicationsDosageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MedicalSummaryConditionsTreatmentMedicationsDosageStructuredDosage{}),
		},
	)
}

type MedicalSummaryConditionsTreatmentMedicationsDosageStructuredDosage struct {
	Amount    float64                                                                `json:"amount" api:"required"`
	Unit      string                                                                 `json:"unit" api:"required"`
	Frequency string                                                                 `json:"frequency"`
	JSON      medicalSummaryConditionsTreatmentMedicationsDosageStructuredDosageJSON `json:"-"`
}

// medicalSummaryConditionsTreatmentMedicationsDosageStructuredDosageJSON contains
// the JSON metadata for the struct
// [MedicalSummaryConditionsTreatmentMedicationsDosageStructuredDosage]
type medicalSummaryConditionsTreatmentMedicationsDosageStructuredDosageJSON struct {
	Amount      apijson.Field
	Unit        apijson.Field
	Frequency   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MedicalSummaryConditionsTreatmentMedicationsDosageStructuredDosage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r medicalSummaryConditionsTreatmentMedicationsDosageStructuredDosageJSON) RawJSON() string {
	return r.raw
}

func (r MedicalSummaryConditionsTreatmentMedicationsDosageStructuredDosage) ImplementsMedicalSummaryConditionsTreatmentMedicationsDosageUnion() {
}

type VaccinationRecord struct {
	ID             string    `json:"id" api:"required"`
	AdministeredAt time.Time `json:"administeredAt" api:"required" format:"date-time"`
	Vaccine        string    `json:"vaccine" api:"required"`
	// Provenance stamp attached to records across unrelated resources.
	Audit           VaccinationRecordAudit             `json:"audit"`
	BoosterSchedule []VaccinationRecordBoosterSchedule `json:"boosterSchedule"`
	Veterinarian    VaccinationRecordVeterinarian      `json:"veterinarian"`
	JSON            vaccinationRecordJSON              `json:"-"`
}

// vaccinationRecordJSON contains the JSON metadata for the struct
// [VaccinationRecord]
type vaccinationRecordJSON struct {
	ID              apijson.Field
	AdministeredAt  apijson.Field
	Vaccine         apijson.Field
	Audit           apijson.Field
	BoosterSchedule apijson.Field
	Veterinarian    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VaccinationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaccinationRecordJSON) RawJSON() string {
	return r.raw
}

// Provenance stamp attached to records across unrelated resources.
type VaccinationRecordAudit struct {
	RecordedAt time.Time                    `json:"recordedAt" api:"required" format:"date-time"`
	RecordedBy string                       `json:"recordedBy"`
	Source     VaccinationRecordAuditSource `json:"source"`
	JSON       vaccinationRecordAuditJSON   `json:"-"`
}

// vaccinationRecordAuditJSON contains the JSON metadata for the struct
// [VaccinationRecordAudit]
type vaccinationRecordAuditJSON struct {
	RecordedAt  apijson.Field
	RecordedBy  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaccinationRecordAudit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaccinationRecordAuditJSON) RawJSON() string {
	return r.raw
}

type VaccinationRecordAuditSource string

const (
	VaccinationRecordAuditSourceManual VaccinationRecordAuditSource = "manual"
	VaccinationRecordAuditSourceImport VaccinationRecordAuditSource = "import"
	VaccinationRecordAuditSourceSystem VaccinationRecordAuditSource = "system"
)

func (r VaccinationRecordAuditSource) IsKnown() bool {
	switch r {
	case VaccinationRecordAuditSourceManual, VaccinationRecordAuditSourceImport, VaccinationRecordAuditSourceSystem:
		return true
	}
	return false
}

type VaccinationRecordBoosterSchedule struct {
	Due       time.Time                                 `json:"due" api:"required" format:"date-time"`
	Completed bool                                      `json:"completed"`
	Dose      VaccinationRecordBoosterScheduleDoseUnion `json:"dose"`
	JSON      vaccinationRecordBoosterScheduleJSON      `json:"-"`
}

// vaccinationRecordBoosterScheduleJSON contains the JSON metadata for the struct
// [VaccinationRecordBoosterSchedule]
type vaccinationRecordBoosterScheduleJSON struct {
	Due         apijson.Field
	Completed   apijson.Field
	Dose        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaccinationRecordBoosterSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaccinationRecordBoosterScheduleJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type VaccinationRecordBoosterScheduleDoseUnion interface {
	ImplementsVaccinationRecordBoosterScheduleDoseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VaccinationRecordBoosterScheduleDoseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type VaccinationRecordVeterinarian struct {
	License VaccinationRecordVeterinarianLicense `json:"license"`
	Name    string                               `json:"name"`
	JSON    vaccinationRecordVeterinarianJSON    `json:"-"`
}

// vaccinationRecordVeterinarianJSON contains the JSON metadata for the struct
// [VaccinationRecordVeterinarian]
type vaccinationRecordVeterinarianJSON struct {
	License     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaccinationRecordVeterinarian) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaccinationRecordVeterinarianJSON) RawJSON() string {
	return r.raw
}

type VaccinationRecordVeterinarianLicense struct {
	Number string                                   `json:"number"`
	Region string                                   `json:"region"`
	JSON   vaccinationRecordVeterinarianLicenseJSON `json:"-"`
}

// vaccinationRecordVeterinarianLicenseJSON contains the JSON metadata for the
// struct [VaccinationRecordVeterinarianLicense]
type vaccinationRecordVeterinarianLicenseJSON struct {
	Number      apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaccinationRecordVeterinarianLicense) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaccinationRecordVeterinarianLicenseJSON) RawJSON() string {
	return r.raw
}

type VaccinationRecordParam struct {
	ID             param.Field[string]    `json:"id" api:"required"`
	AdministeredAt param.Field[time.Time] `json:"administeredAt" api:"required" format:"date-time"`
	Vaccine        param.Field[string]    `json:"vaccine" api:"required"`
	// Provenance stamp attached to records across unrelated resources.
	Audit           param.Field[VaccinationRecordAuditParam]             `json:"audit"`
	BoosterSchedule param.Field[[]VaccinationRecordBoosterScheduleParam] `json:"boosterSchedule"`
	Veterinarian    param.Field[VaccinationRecordVeterinarianParam]      `json:"veterinarian"`
}

func (r VaccinationRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Provenance stamp attached to records across unrelated resources.
type VaccinationRecordAuditParam struct {
	RecordedAt param.Field[time.Time]                    `json:"recordedAt" api:"required" format:"date-time"`
	RecordedBy param.Field[string]                       `json:"recordedBy"`
	Source     param.Field[VaccinationRecordAuditSource] `json:"source"`
}

func (r VaccinationRecordAuditParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaccinationRecordBoosterScheduleParam struct {
	Due       param.Field[time.Time]                                      `json:"due" api:"required" format:"date-time"`
	Completed param.Field[bool]                                           `json:"completed"`
	Dose      param.Field[VaccinationRecordBoosterScheduleDoseUnionParam] `json:"dose"`
}

func (r VaccinationRecordBoosterScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionFloat], [shared.UnionString].
type VaccinationRecordBoosterScheduleDoseUnionParam interface {
	ImplementsVaccinationRecordBoosterScheduleDoseUnionParam()
}

type VaccinationRecordVeterinarianParam struct {
	License param.Field[VaccinationRecordVeterinarianLicenseParam] `json:"license"`
	Name    param.Field[string]                                    `json:"name"`
}

func (r VaccinationRecordVeterinarianParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaccinationRecordVeterinarianLicenseParam struct {
	Number param.Field[string] `json:"number"`
	Region param.Field[string] `json:"region"`
}

func (r VaccinationRecordVeterinarianLicenseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
