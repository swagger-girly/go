// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/swagger-girly/go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Adopted string                // Always "adopted"
type AdoptionsPolicyChanged string // Always "adoptions.policy.changed"
type Approved string               // Always "approved"
type Escalated string              // Always "escalated"
type Individual string             // Always "individual"
type Note string                   // Always "note"
type Organization string           // Always "organization"
type PetModerationApproved string  // Always "pet.moderation.approved"
type PetModerationRejected string  // Always "pet.moderation.rejected"
type Ping string                   // Always "ping"
type Pong string                   // Always "pong"
type Rejected string               // Always "rejected"
type Returned string               // Always "returned"
type Status string                 // Always "status"
type StoreReportGenerated string   // Always "store.report.generated"
type Subscribe string              // Always "subscribe"

func (c Adopted) Default() Adopted                               { return "adopted" }
func (c AdoptionsPolicyChanged) Default() AdoptionsPolicyChanged { return "adoptions.policy.changed" }
func (c Approved) Default() Approved                             { return "approved" }
func (c Escalated) Default() Escalated                           { return "escalated" }
func (c Individual) Default() Individual                         { return "individual" }
func (c Note) Default() Note                                     { return "note" }
func (c Organization) Default() Organization                     { return "organization" }
func (c PetModerationApproved) Default() PetModerationApproved   { return "pet.moderation.approved" }
func (c PetModerationRejected) Default() PetModerationRejected   { return "pet.moderation.rejected" }
func (c Ping) Default() Ping                                     { return "ping" }
func (c Pong) Default() Pong                                     { return "pong" }
func (c Rejected) Default() Rejected                             { return "rejected" }
func (c Returned) Default() Returned                             { return "returned" }
func (c Status) Default() Status                                 { return "status" }
func (c StoreReportGenerated) Default() StoreReportGenerated     { return "store.report.generated" }
func (c Subscribe) Default() Subscribe                           { return "subscribe" }

func (c Adopted) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c AdoptionsPolicyChanged) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Approved) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Escalated) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Individual) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Note) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Organization) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c PetModerationApproved) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c PetModerationRejected) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c Ping) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Pong) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Rejected) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Returned) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Status) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c StoreReportGenerated) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Subscribe) MarshalJSON() ([]byte, error)              { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
