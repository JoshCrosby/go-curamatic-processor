package fhir

import "time"

// Claim structure to match the provided JSON schema
type Claim struct {
	ResourceType         string           `json:"resourceType"`
	Identifier           []Identifier     `json:"identifier"`
	Status               string           `json:"status"`
	Type                 CodeableConcept  `json:"type"`
	SubType              CodeableConcept  `json:"subType"`
	Use                  string           `json:"use"`
	Patient              Patient          `json:"patient"`
	BillablePeriod       Period           `json:"billablePeriod"`
	Created              time.Time        `json:"created"`
	Enterer              Reference        `json:"enterer"`
	Insurer              Reference        `json:"insurer"`
	Provider             Reference        `json:"provider"`
	Priority             CodeableConcept  `json:"priority"`
	FundsReserve         CodeableConcept  `json:"fundsReserve"`
	Related              []RelatedClaim   `json:"related"`
	Prescription         Reference        `json:"prescription"`
	OriginalPrescription Reference        `json:"originalPrescription"`
	Payee                Payee            `json:"payee"`
	Referral             Reference        `json:"referral"`
	Facility             Reference        `json:"facility"`
	CareTeam             []CareTeamMember `json:"careTeam"`
	SupportingInfo       []SupportingInfo `json:"supportingInfo"`
	Diagnosis            []Diagnosis      `json:"diagnosis"`
	Procedure            []Procedure      `json:"procedure"`
	Insurance            []Insurance      `json:"insurance"`
	Accident             Accident         `json:"accident"`
	Item                 []Item           `json:"item"`
	Total                Money            `json:"total"`
}

// Reference represents a reference to another resource
type Reference struct {
	Reference string `json:"reference,omitempty"`
}

// Period represents a time period with a start and an end
type Period struct {
	Start time.Time `json:"start,omitempty"`
	End   time.Time `json:"end,omitempty"`
}

// Quantity represents a measured amount
type Quantity struct {
	Value  float64 `json:"value,omitempty"`
	Unit   string  `json:"unit,omitempty"`
	System string  `json:"system,omitempty"`
	Code   string  `json:"code,omitempty"`
}

// Money represents a monetary value
type Money struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// RelatedClaim represents a related claim
type RelatedClaim struct {
	Claim        Reference       `json:"claim"`
	Relationship CodeableConcept `json:"relationship"`
	Reference    Identifier      `json:"reference"`
}

// Payee represents the payee of the claim
type Payee struct {
	Type  CodeableConcept `json:"type"`
	Party Reference       `json:"party"`
}

// CareTeamMember represents a member of the care team
type CareTeamMember struct {
	Sequence      int             `json:"sequence"`
	Provider      Reference       `json:"provider"`
	Responsible   bool            `json:"responsible,omitempty"`
	Role          CodeableConcept `json:"role,omitempty"`
	Qualification CodeableConcept `json:"qualification,omitempty"`
}

// SupportingInfo represents supporting information for the claim
type SupportingInfo struct {
	Sequence        int             `json:"sequence"`
	Category        CodeableConcept `json:"category"`
	Code            CodeableConcept `json:"code,omitempty"`
	TimingDate      time.Time       `json:"timingDate,omitempty"`
	TimingPeriod    Period          `json:"timingPeriod,omitempty"`
	ValueBoolean    bool            `json:"valueBoolean,omitempty"`
	ValueString     string          `json:"valueString,omitempty"`
	ValueQuantity   Quantity        `json:"valueQuantity,omitempty"`
	ValueAttachment Attachment      `json:"valueAttachment,omitempty"`
	ValueReference  Reference       `json:"valueReference,omitempty"`
	Reason          CodeableConcept `json:"reason,omitempty"`
}

// Diagnosis represents a diagnosis for the claim
type Diagnosis struct {
	Sequence                 int               `json:"sequence"`
	DiagnosisCodeableConcept CodeableConcept   `json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       Reference         `json:"diagnosisReference,omitempty"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	OnAdmission              CodeableConcept   `json:"onAdmission,omitempty"`
	PackageCode              CodeableConcept   `json:"packageCode,omitempty"`
}

// Procedure represents a procedure for the claim
type Procedure struct {
	Sequence                 int               `json:"sequence"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	Date                     time.Time         `json:"date,omitempty"`
	ProcedureCodeableConcept CodeableConcept   `json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       Reference         `json:"procedureReference,omitempty"`
	UDI                      []Reference       `json:"udi,omitempty"`
}

// Insurance represents insurance information for the claim
type Insurance struct {
	Sequence            int        `json:"sequence"`
	Focal               bool       `json:"focal"`
	Identifier          Identifier `json:"identifier,omitempty"`
	Coverage            Reference  `json:"coverage"`
	BusinessArrangement string     `json:"businessArrangement,omitempty"`
	PreAuthRef          []string   `json:"preAuthRef,omitempty"`
	ClaimResponse       Reference  `json:"claimResponse,omitempty"`
}

// Accident represents an accident
type Accident struct {
	Date              time.Time       `json:"date"`
	Type              CodeableConcept `json:"type,omitempty"`
	LocationAddress   Address         `json:"locationAddress,omitempty"`
	LocationReference Reference       `json:"locationReference,omitempty"`
}

// Item represents a product or service provided under the claim.
type Item struct {
	Sequence                int               `json:"sequence"`
	CareTeamSequence        []int             `json:"careTeamSequence,omitempty"`
	DiagnosisSequence       []int             `json:"diagnosisSequence,omitempty"`
	ProcedureSequence       []int             `json:"procedureSequence,omitempty"`
	InformationSequence     []int             `json:"informationSequence,omitempty"`
	Revenue                 CodeableConcept   `json:"revenue,omitempty"`
	Category                CodeableConcept   `json:"category,omitempty"`
	ProductOrService        CodeableConcept   `json:"productOrService"`
	Modifier                []CodeableConcept `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept `json:"programCode,omitempty"`
	ServicedDate            *time.Time        `json:"servicedDate,omitempty"` // Pointer to time.Time for optional field
	ServicedPeriod          Period            `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept CodeableConcept   `json:"locationCodeableConcept,omitempty"`
	LocationAddress         Address           `json:"locationAddress,omitempty"`
	LocationReference       Reference         `json:"locationReference,omitempty"`
	Quantity                Quantity          `json:"quantity,omitempty"`
	UnitPrice               Money             `json:"unitPrice,omitempty"`
	Factor                  float64           `json:"factor,omitempty"`
	Net                     Money             `json:"net,omitempty"`
	UDI                     []Reference       `json:"udi,omitempty"` // Unique device identifier
	BodySite                CodeableConcept   `json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept `json:"subSite,omitempty"`
	Encounter               []Reference       `json:"encounter,omitempty"`
	Detail                  []ItemDetail      `json:"detail,omitempty"`
}

// ItemDetail represents the details of a provided product or service within an item.
type ItemDetail struct {
	Sequence         int               `json:"sequence"`
	Revenue          CodeableConcept   `json:"revenue,omitempty"`
	Category         CodeableConcept   `json:"category,omitempty"`
	ProductOrService CodeableConcept   `json:"productOrService"`
	Modifier         []CodeableConcept `json:"modifier,omitempty"`
	ProgramCode      []CodeableConcept `json:"programCode,omitempty"`
	Quantity         Quantity          `json:"quantity,omitempty"`
	UnitPrice        Money             `json:"unitPrice,omitempty"`
	Factor           float64           `json:"factor,omitempty"`
	Net              Money             `json:"net,omitempty"`
	UDI              []Reference       `json:"udi,omitempty"` // Unique device identifier
	SubDetail        []ItemSubDetail   `json:"subDetail,omitempty"`
}

// ItemSubDetail represents further granularity in the detail of a provided product or service.
type ItemSubDetail struct {
	Sequence         int               `json:"sequence"`
	Revenue          CodeableConcept   `json:"revenue,omitempty"`
	Category         CodeableConcept   `json:"category,omitempty"`
	ProductOrService CodeableConcept   `json:"productOrService"`
	Modifier         []CodeableConcept `json:"modifier,omitempty"`
	ProgramCode      []CodeableConcept `json:"programCode,omitempty"`
	Quantity         Quantity          `json:"quantity,omitempty"`
	UnitPrice        Money             `json:"unitPrice,omitempty"`
	Factor           float64           `json:"factor,omitempty"`
	Net              Money             `json:"net,omitempty"`
	UDI              []Reference       `json:"udi,omitempty"` // Unique device identifier
}
