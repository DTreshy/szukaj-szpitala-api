package geocoding

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strings"
)

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Resp struct {
	Tp       any       `json:"type"`
	Features []Feature `json:"features"`
	Query    any       `json:"query"`
}

type Feature struct {
	Tp         any      `json:"type"`
	Properties any      `json:"properties"`
	Geometry   Geometry `json:"geometry"`
	Bbox       any      `json:"bbox"`
}

type Geometry struct {
	Tp          string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func GetCoordinates(street, city, country string) (*Coordinates, error) {
	method := "GET"
	client := &http.Client{}
	response := &Resp{}
	key := "e2980109bb0c4ae88ce4f3d9d7473df7"
	baseUrl := "https://api.geoapify.com/v1/geocode/search?text="
	housenumber := ""

	stSlice := strings.Split(street, "/")
	if len(stSlice) == 2 {
		housenumber = stSlice[1] + " "
	}

	street = stSlice[0]
	searchUrl := url.PathEscape(fmt.Sprintf("%s%s,%s,%s", housenumber, street, city, country))
	fullUrl := fmt.Sprintf("%s%s&apiKey=%s", baseUrl, searchUrl, key)

	req, err := http.NewRequest(method, fullUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to uncode response: %w", err)
	}

	if err := json.Unmarshal(body, response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &Coordinates{
		Longitude: response.Features[0].Geometry.Coordinates[0],
		Latitude:  response.Features[0].Geometry.Coordinates[1],
	}, nil
}

// Returns distance beetween 2 coordinates in km
func CalculateDistance(coords1, coords2 Coordinates) float64 {
	radlat1 := float64(math.Pi * coords1.Latitude / 180)
	radlat2 := float64(math.Pi * coords2.Latitude / 180)

	theta := float64(coords1.Longitude - coords2.Longitude)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.853159616

	return roundFloat(dist, 3)
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
