package models

type CarModel struct {
	Name     string   `json:"name"`
	Url      string   `json:"url"`
	FileName string   `json:"fileName"`
	Models   []string `json:"models"`
	Brand    string   `json:"brand"`
}
