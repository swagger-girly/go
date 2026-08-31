// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"encoding/json"
	"time"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
)

// VeterinaryService contains methods and other services that help with interacting
// with the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVeterinaryService] method instead.
type VeterinaryService struct {
	options []option.RequestOption
}

// NewVeterinaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVeterinaryService(opts ...option.RequestOption) (r VeterinaryService) {
	r = VeterinaryService{}
	r.options = opts
	return
}

type MedicalSummary struct {
	PetID int64 `json:"petId" api:"required"`
	// Provenance stamp attached to records across unrelated resources.
	Audit        MedicalSummaryAudit       `json:"audit"`
	Conditions   []MedicalSummaryCondition `json:"conditions"`
	Vaccinations []VaccinationRecord       `json:"vaccinations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PetID        respjson.Field
		Audit        respjson.Field
		Conditions   respjson.Field
		Vaccinations respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MedicalSummary) RawJSON() string { return r.JSON.raw }
func (r *MedicalSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provenance stamp attached to records across unrelated resources.
type MedicalSummaryAudit struct {
	RecordedAt time.Time `json:"recordedAt" api:"required" format:"date-time"`
	RecordedBy string    `json:"recordedBy"`
	// Any of "manual", "import", "system".
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordedAt  respjson.Field
		RecordedBy  respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MedicalSummaryAudit) RawJSON() string { return r.JSON.raw }
func (r *MedicalSummaryAudit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MedicalSummaryCondition struct {
	Code string `json:"code" api:"required"`
	// Any of "mild", "moderate", "severe".
	Severity  string                           `json:"severity"`
	Treatment MedicalSummaryConditionTreatment `json:"treatment"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Severity    respjson.Field
		Treatment   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MedicalSummaryCondition) RawJSON() string { return r.JSON.raw }
func (r *MedicalSummaryCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MedicalSummaryConditionTreatment struct {
	Medications []MedicalSummaryConditionTreatmentMedication `json:"medications"`
	Plan        string                                       `json:"plan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Medications respjson.Field
		Plan        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MedicalSummaryConditionTreatment) RawJSON() string { return r.JSON.raw }
func (r *MedicalSummaryConditionTreatment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MedicalSummaryConditionTreatmentMedication struct {
	Name string `json:"name" api:"required"`
	// Free-form dosage text or a structured dosage.
	Dosage MedicalSummaryConditionTreatmentMedicationDosageUnion `json:"dosage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Dosage      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MedicalSummaryConditionTreatmentMedication) RawJSON() string { return r.JSON.raw }
func (r *MedicalSummaryConditionTreatmentMedication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MedicalSummaryConditionTreatmentMedicationDosageUnion contains all possible
// properties and values from [string],
// [MedicalSummaryConditionTreatmentMedicationDosageStructuredDosage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type MedicalSummaryConditionTreatmentMedicationDosageUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [MedicalSummaryConditionTreatmentMedicationDosageStructuredDosage].
	Amount float64 `json:"amount"`
	// This field is from variant
	// [MedicalSummaryConditionTreatmentMedicationDosageStructuredDosage].
	Unit string `json:"unit"`
	// This field is from variant
	// [MedicalSummaryConditionTreatmentMedicationDosageStructuredDosage].
	Frequency string `json:"frequency"`
	JSON      struct {
		OfString  respjson.Field
		Amount    respjson.Field
		Unit      respjson.Field
		Frequency respjson.Field
		raw       string
	} `json:"-"`
}

func (u MedicalSummaryConditionTreatmentMedicationDosageUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MedicalSummaryConditionTreatmentMedicationDosageUnion) AsStructuredDosage() (v MedicalSummaryConditionTreatmentMedicationDosageStructuredDosage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MedicalSummaryConditionTreatmentMedicationDosageUnion) RawJSON() string { return u.JSON.raw }

func (r *MedicalSummaryConditionTreatmentMedicationDosageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MedicalSummaryConditionTreatmentMedicationDosageStructuredDosage struct {
	Amount    float64 `json:"amount" api:"required"`
	Unit      string  `json:"unit" api:"required"`
	Frequency string  `json:"frequency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		Frequency   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MedicalSummaryConditionTreatmentMedicationDosageStructuredDosage) RawJSON() string {
	return r.JSON.raw
}
func (r *MedicalSummaryConditionTreatmentMedicationDosageStructuredDosage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VaccinationRecord struct {
	ID             string    `json:"id" api:"required"`
	AdministeredAt time.Time `json:"administeredAt" api:"required" format:"date-time"`
	Vaccine        string    `json:"vaccine" api:"required"`
	// Provenance stamp attached to records across unrelated resources.
	Audit           VaccinationRecordAudit             `json:"audit"`
	BoosterSchedule []VaccinationRecordBoosterSchedule `json:"boosterSchedule"`
	Veterinarian    VaccinationRecordVeterinarian      `json:"veterinarian"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AdministeredAt  respjson.Field
		Vaccine         respjson.Field
		Audit           respjson.Field
		BoosterSchedule respjson.Field
		Veterinarian    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VaccinationRecord) RawJSON() string { return r.JSON.raw }
func (r *VaccinationRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VaccinationRecord to a VaccinationRecordParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VaccinationRecordParam.Overrides()
func (r VaccinationRecord) ToParam() VaccinationRecordParam {
	return param.Override[VaccinationRecordParam](json.RawMessage(r.RawJSON()))
}

// Provenance stamp attached to records across unrelated resources.
type VaccinationRecordAudit struct {
	RecordedAt time.Time `json:"recordedAt" api:"required" format:"date-time"`
	RecordedBy string    `json:"recordedBy"`
	// Any of "manual", "import", "system".
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordedAt  respjson.Field
		RecordedBy  respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VaccinationRecordAudit) RawJSON() string { return r.JSON.raw }
func (r *VaccinationRecordAudit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VaccinationRecordBoosterSchedule struct {
	Due       time.Time                                 `json:"due" api:"required" format:"date-time"`
	Completed bool                                      `json:"completed"`
	Dose      VaccinationRecordBoosterScheduleDoseUnion `json:"dose"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Due         respjson.Field
		Completed   respjson.Field
		Dose        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VaccinationRecordBoosterSchedule) RawJSON() string { return r.JSON.raw }
func (r *VaccinationRecordBoosterSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VaccinationRecordBoosterScheduleDoseUnion contains all possible properties and
// values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type VaccinationRecordBoosterScheduleDoseUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u VaccinationRecordBoosterScheduleDoseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VaccinationRecordBoosterScheduleDoseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VaccinationRecordBoosterScheduleDoseUnion) RawJSON() string { return u.JSON.raw }

func (r *VaccinationRecordBoosterScheduleDoseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VaccinationRecordVeterinarian struct {
	License VaccinationRecordVeterinarianLicense `json:"license"`
	Name    string                               `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		License     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VaccinationRecordVeterinarian) RawJSON() string { return r.JSON.raw }
func (r *VaccinationRecordVeterinarian) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VaccinationRecordVeterinarianLicense struct {
	Number string `json:"number"`
	Region string `json:"region"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VaccinationRecordVeterinarianLicense) RawJSON() string { return r.JSON.raw }
func (r *VaccinationRecordVeterinarianLicense) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, AdministeredAt, Vaccine are required.
type VaccinationRecordParam struct {
	ID             string    `json:"id" api:"required"`
	AdministeredAt time.Time `json:"administeredAt" api:"required" format:"date-time"`
	Vaccine        string    `json:"vaccine" api:"required"`
	// Provenance stamp attached to records across unrelated resources.
	Audit           VaccinationRecordAuditParam             `json:"audit,omitzero"`
	BoosterSchedule []VaccinationRecordBoosterScheduleParam `json:"boosterSchedule,omitzero"`
	Veterinarian    VaccinationRecordVeterinarianParam      `json:"veterinarian,omitzero"`
	paramObj
}

func (r VaccinationRecordParam) MarshalJSON() (data []byte, err error) {
	type shadow VaccinationRecordParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VaccinationRecordParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provenance stamp attached to records across unrelated resources.
//
// The property RecordedAt is required.
type VaccinationRecordAuditParam struct {
	RecordedAt time.Time         `json:"recordedAt" api:"required" format:"date-time"`
	RecordedBy param.Opt[string] `json:"recordedBy,omitzero"`
	// Any of "manual", "import", "system".
	Source string `json:"source,omitzero"`
	paramObj
}

func (r VaccinationRecordAuditParam) MarshalJSON() (data []byte, err error) {
	type shadow VaccinationRecordAuditParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VaccinationRecordAuditParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VaccinationRecordAuditParam](
		"source", "manual", "import", "system",
	)
}

// The property Due is required.
type VaccinationRecordBoosterScheduleParam struct {
	Due       time.Time                                      `json:"due" api:"required" format:"date-time"`
	Completed param.Opt[bool]                                `json:"completed,omitzero"`
	Dose      VaccinationRecordBoosterScheduleDoseUnionParam `json:"dose,omitzero"`
	paramObj
}

func (r VaccinationRecordBoosterScheduleParam) MarshalJSON() (data []byte, err error) {
	type shadow VaccinationRecordBoosterScheduleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VaccinationRecordBoosterScheduleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VaccinationRecordBoosterScheduleDoseUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u VaccinationRecordBoosterScheduleDoseUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *VaccinationRecordBoosterScheduleDoseUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type VaccinationRecordVeterinarianParam struct {
	Name    param.Opt[string]                         `json:"name,omitzero"`
	License VaccinationRecordVeterinarianLicenseParam `json:"license,omitzero"`
	paramObj
}

func (r VaccinationRecordVeterinarianParam) MarshalJSON() (data []byte, err error) {
	type shadow VaccinationRecordVeterinarianParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VaccinationRecordVeterinarianParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VaccinationRecordVeterinarianLicenseParam struct {
	Number param.Opt[string] `json:"number,omitzero"`
	Region param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r VaccinationRecordVeterinarianLicenseParam) MarshalJSON() (data []byte, err error) {
	type shadow VaccinationRecordVeterinarianLicenseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VaccinationRecordVeterinarianLicenseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
