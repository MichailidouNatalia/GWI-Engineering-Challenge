package dto

import "time"

// AssetRequest represents a request to create or update an asset.
// swagger:model AssetRequest
type AssetRequest struct {
	// Unique identifier of the asset
	// example: 123e4567-e89b-12d3-a456-426614174000
	ID string `json:"id" validate:"required"`

	// Type of the asset (audience, chart, or insight)
	// required: true
	// enum: audience,chart,insight
	// example: chart
	Type string `json:"type" validate:"required,oneof=audience chart insight"`

	// Title of the asset
	// required: true
	// example: Customer Demographics Overview
	Title string `json:"title" validate:"required"`

	// Description of the asset
	// example: This chart shows demographic breakdown by age and gender.
	Description string `json:"description"`

	// Timestamp when the asset was created
	// example: 2025-01-01T12:00:00Z
	CreatedAt time.Time `json:"created_at"`

	// Timestamp when the asset was last updated
	// example: 2025-01-02T15:30:00Z
	UpdatedAt time.Time `json:"updated_at"`

	// Gender associated with the audience (optional)
	// example: female
	Gender *string `json:"gender,omitempty"`

	// Birth country of the audience (optional)
	// example: Canada
	BirthCountry *string `json:"birth_country,omitempty"`

	// Age group of the audience (optional)
	// example: 25-34
	AgeGroup *string `json:"age_group,omitempty"`

	// Average hours spent on social media per day (optional)
	// example: 4
	HoursSocial *int `json:"hours_social,omitempty"`

	// Number of purchases made in the last month (optional)
	// example: 12
	PurchasesLastMo *int `json:"purchases_last_month,omitempty"`

	// Titles for the chart axes (for chart-type assets)
	// example: ["Age Group", "Average Purchases"]
	AxesTitles []string `json:"axes_titles,omitempty"`

	// Data points for the chart (for chart-type assets)
	// example: [[18, 2.3], [25, 3.5], [34, 4.1]]
	Data [][]float64 `json:"data,omitempty"`
}
