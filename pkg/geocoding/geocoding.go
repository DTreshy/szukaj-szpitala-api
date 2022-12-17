package geocoding

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang/geo/s1"
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
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

// Radius of earth in kilometers. Use 3956 for miles
var r int = 6371

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
func CalculateDistance(coords1, coords2 *Coordinates) float64 {
	lon1 := s1.Angle(coords1.Longitude).Radians()
	lon2 := s1.Angle(coords2.Longitude).Radians()
	lat1 := s1.Angle(coords1.Latitude).Radians()
	lat2 := s1.Angle(coords2.Latitude).Radians()
	dlon := lon2 - lon1
	dlat := lat2 - lat1
	a := math.Pow(math.Sin(dlat/2), 2) + math.Pow(math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2), 2)
	c := 2 * math.Asin(math.Sqrt(a))

	return c * float64(r)
}
