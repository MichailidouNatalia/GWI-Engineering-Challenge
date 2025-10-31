package dto

import "time"

type ChartResponse struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	AxesTitles  []string    `json:"axes_titles"`
	Data        [][]float64 `json:"data"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
