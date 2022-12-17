package models

import "github.com/DTreshy/szukaj-szpitala-api/pkg/geocoding"

type Hospital struct {
	Name  string
	Phone string
	Email string
	Place *Place
}

type Place struct {
	Address     string
	City        string
	Coordinates *geocoding.Coordinates
}

type HospitalInfo struct {
	Hospital *Hospital
	Distance float64
}
