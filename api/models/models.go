package models

import "github.com/DTreshy/szukaj-szpitala-api/pkg/geocoding"

type Hospital struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Place Place  `json:"place"`
}

type Place struct {
	Address     string                `json:"address"`
	City        string                `json:"city"`
	Coordinates geocoding.Coordinates `json:"coordinates"`
}

type HospitalInfo struct {
	Hospital *Hospital
	Distance float64
}
