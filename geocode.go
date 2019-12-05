package gogeocoder

import (
	"encoding/json"
	// "log"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func Reverse(longitude, latitude float64) (Location, error) {
	var location Location

	searchURL := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&addressdetails=1&lon=%v&lat=%v", longitude, latitude)

	resp, err := http.Get(searchURL)
	if err != nil {
		err = errors.Wrap(err, searchURL)
		return location, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		err = errors.Wrap(err, searchURL)
		return location, err
	}

	return location, err
}

func Geocode(search string) (Location, error) {
	var location []Location

	searchURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&addressdetails=1", search)

	resp, err := http.Get(searchURL)
	if err != nil {
		err = errors.Wrap(err, searchURL)
		return Location{}, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		err = errors.Wrap(err, searchURL)
		return Location{}, err
	}

	return location[0], err
}

type Location struct {
	PlaceID int     `json:"place_id,omitempty"`
	OsmID	int	`json:"osm_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Lat     string  `json:"lat,omitempty"`
	Lon     string  `json:"lon,omitempty"`
	Class     string  `json:"class,omitempty"`
	Type     string  `json:"type,omitempty"`
	Address Address `json:"address,omitempty"`
	BoundingBox []string `json:"boundingbox,omitempty"`
}

type Address struct {
	Road string `json:"road,omitempty"`
	Suburb string `json:"suburb,omitempty"`
	PostCode string `json:"postcode,omitempty"`
	Village string `json:"village,omitempty"`
	City    string `json:"city,omitempty"`
	County  string `json:"county,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}
