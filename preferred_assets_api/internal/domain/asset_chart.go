package domain

type Chart struct {
	Asset
	AxesTitles []string
	Data       [][]float64
}
