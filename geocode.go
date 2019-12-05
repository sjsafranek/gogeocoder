package gogeocoder

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func Reverse(longitude, latitude float64) (Location, error) {
	var n Location

	searchURL := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&addressdetails=1&lon=%v&lat=%v", longitude, latitude)

	resp, err := http.Get(searchURL)
	if err != nil {
		err = errors.Wrap(err, searchURL)
		return n, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&n)
	if err != nil {
		err = errors.Wrap(err, searchURL)
		return n, err
	}

	return n, err
}

func Geocode(search string) (Location, error) {
	var n []Location

	searchURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&addressdetails=1", search)

	resp, err := http.Get(searchURL)
	if err != nil {
		err = errors.Wrap(err, searchURL)
		return Location{}, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&n)
	if err != nil {
		err = errors.Wrap(err, searchURL)
		return Location{}, err
	}

	return n[0], err
}

type Location struct {
	PlaceID int     `json:"place_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Lat     string  `json:"lat,omitempty"`
	Lon     string  `json:"lon,omitempty"`
	Address Address `json:"address,omitempty"`
}

type Address struct {
	City    string `json:"city,omitempty"`
	County  string `json:"county,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
}
