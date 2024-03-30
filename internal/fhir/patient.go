package fhir

// Identifier for a FHIR resource
type Identifier struct {
	Use    string `json:"use,omitempty"`
	System string `json:"system,omitempty"`
	Value  string `json:"value,omitempty"`
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

// Address for a FHIR resource
type Address struct {
	Use        string   `json:"use,omitempty"`
	Type       string   `json:"type,omitempty"`
	Text       string   `json:"text,omitempty"`
	Line       []string `json:"line,omitempty"`
	City       string   `json:"city,omitempty"`
	District   string   `json:"district,omitempty"`
	State      string   `json:"state,omitempty"`
	PostalCode string   `json:"postalCode,omitempty"`
	Country    string   `json:"country,omitempty"`
}

// CodeableConcept specifies a coded value with optional coding and text
type CodeableConcept struct {
	Coding []Coding `json:"coding,omitempty"`
	Text   string   `json:"text,omitempty"`
}

// Coding defines a code with system, version, and code
type Coding struct {
	System  string `json:"system,omitempty"`
	Version string `json:"version,omitempty"`
	Code    string `json:"code,omitempty"`
	Display string `json:"display,omitempty"`
}

// Patient represents a FHIR patient resource
type Patient struct {
	ResourceType     string           `json:"resourceType"`
	ID               string           `json:"id"`
	Identifier       []Identifier     `json:"identifier,omitempty"`
	Active           bool             `json:"active,omitempty"`
	Name             []HumanName      `json:"name,omitempty"`
	Telecom          []ContactPoint   `json:"telecom,omitempty"`
	Gender           string           `json:"gender,omitempty"`
	BirthDate        string           `json:"birthDate,omitempty"`
	DeceasedBoolean  bool             `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime string           `json:"deceasedDateTime,omitempty"`
	Address          []Address        `json:"address,omitempty"`
	MaritalStatus    CodeableConcept  `json:"maritalStatus,omitempty"`
	Photo            []Attachment     `json:"photo,omitempty"`
	Contact          []PatientContact `json:"contact,omitempty"`
	Communication    []Communication  `json:"communication,omitempty"`
	// Add other fields as necessary
}

// Simplified for demonstration, should be expanded to match FHIR specification
type Attachment struct{}
type PatientContact struct{}
type Communication struct{}
