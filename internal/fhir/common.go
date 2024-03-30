package fhir

// Identifier represents a generic identifier used throughout FHIR resources
type Identifier struct {
	Use    string `json:"use,omitempty"`
	System string `json:"system,omitempty"`
	Value  string `json:"value,omitempty"`
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

// Address represents a postal address for a person or organization, conforming to FHIR standards.
type Address struct {
	Use        string   `json:"use,omitempty"`        // The use of the address (e.g., home, work).
	Type       string   `json:"type,omitempty"`       // The type of the address (e.g., physical, postal).
	Text       string   `json:"text,omitempty"`       // A full text representation of the address.
	Line       []string `json:"line,omitempty"`       // Address lines (e.g., street, PO Box, care of).
	City       string   `json:"city,omitempty"`       // The city name.
	District   string   `json:"district,omitempty"`   // District name (e.g., county).
	State      string   `json:"state,omitempty"`      // Sub-unit of a country (abbreviations ok).
	PostalCode string   `json:"postalCode,omitempty"` // Postal code for area.
	Country    string   `json:"country,omitempty"`    // Country (can be ISO 3166 3 letter code).
}

type Attachment struct {
	// Fields according to the schema
}

type Link struct {
	Other Reference `json:"other,omitempty"`
	Type  string    `json:"type,omitempty"`
}
