package fhir

// Patient represents a FHIR patient resource
type Patient struct {
	ResourceType         string           `json:"resourceType"`
	ID                   string           `json:"id"`
	Identifier           []Identifier     `json:"identifier,omitempty"`
	Active               bool             `json:"active,omitempty"`
	Name                 []HumanName      `json:"name,omitempty"`
	Telecom              []ContactPoint   `json:"telecom,omitempty"`
	Gender               string           `json:"gender,omitempty"`
	BirthDate            string           `json:"birthDate,omitempty"`
	DeceasedBoolean      bool             `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     string           `json:"deceasedDateTime,omitempty"`
	Address              []Address        `json:"address,omitempty"`
	MaritalStatus        CodeableConcept  `json:"maritalStatus,omitempty"`
	Photo                []Attachment     `json:"photo,omitempty"`
	Contact              []PatientContact `json:"contact,omitempty"`
	Communication        []Communication  `json:"communication,omitempty"`
	GeneralPractitioner  []Reference      `json:"generalPractitioner,omitempty"`
	ManagingOrganization []Reference      `json:"managingOrganization,omitempty"`
	Link                 []Link           `json:"link,omitempty"`
}

// PatientContact represents a contact person for the patient
type PatientContact struct {
	Relationship []CodeableConcept `json:"relationship,omitempty"`
	Name         HumanName         `json:"name,omitempty"`
	Telecom      []ContactPoint    `json:"telecom,omitempty"`
	Address      Address           `json:"address,omitempty"`
	Gender       string            `json:"gender,omitempty"`
	Organization Reference         `json:"organization,omitempty"`
	Period       Period            `json:"period,omitempty"`
}

// HumanName represents a person's name with the ability to specify use, text, family, and given names
type HumanName struct {
	Use    string   `json:"use,omitempty"`
	Text   string   `json:"text,omitempty"`
	Family string   `json:"family,omitempty"`
	Given  []string `json:"given,omitempty"`
}

// ContactPoint specifies contact information with system, value, and use
type ContactPoint struct {
	System string `json:"system,omitempty"`
	Value  string `json:"value,omitempty"`
	Use    string `json:"use,omitempty"`
}

// Communication represents a language that the patient can communicate in
type Communication struct {
	Language  CodeableConcept `json:"language,omitempty"`
	Preferred bool            `json:"preferred,omitempty"`
}
