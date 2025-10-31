package dto

import "time"

type AudienceResponse struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Gender          string    `json:"gender"`
	BirthCountry    string    `json:"birth_country"`
	AgeGroup        string    `json:"age_group"`
	HoursSocial     int       `json:"hours_social"`
	PurchasesLastMo int       `json:"purchases_last_mo"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
