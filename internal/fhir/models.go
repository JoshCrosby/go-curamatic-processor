package main

// Patient represents a FHIR patient resource.
type Patient struct {
	ID        string `json:"id"`
	FullName  string `json:"full_name"`
	BirthDate string `json:"birth_date"`
}

// Claim represents a FHIR claim resource.
type Claim struct {
	ID        string  `json:"id"`
	PatientID string  `json:"patient_id"`
	Type      string  `json:"type"`
	Amount    float64 `json:"amount"`
}
