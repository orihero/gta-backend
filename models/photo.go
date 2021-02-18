package models

type Photo struct {
	Path   string `json:"path" validate:"required"`
	CarID  uint
	ClientID uint
}
