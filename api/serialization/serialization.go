package serialization

import (
	"github.com/DTreshy/szukaj-szpitala-api/api/models"
	"github.com/DTreshy/szukaj-szpitala-api/pkg/geocoding"
)

func NewHospital(name, address, city, phone, email string) (*models.Hospital, error) {
	place, err := NewPlace(address, city)
	if err != nil {
		return nil, err
	}

	return &models.Hospital{
		Name:  name,
		Phone: phone,
		Email: email,
		Place: place,
	}, nil
}

func NewPlace(address, city string) (*models.Place, error) {
	coords, err := geocoding.GetCoordinates(address, city, "poland")
	if err != nil {
		return nil, err
	}

	return &models.Place{
		Address: address,
		City:    city,
		Coordinates: &geocoding.Coordinates{
			Latitude:  coords.Latitude,
			Longitude: coords.Longitude,
		},
	}, nil
}
