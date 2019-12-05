package gogeocoder

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv"

	"github.com/pkg/errors"
	// log "github.com/schollz/logger"
)

// type Location struct {
// 	IP          string  `json:"ip,omitempty"`
// 	CountryCode string  `json:"country_code,omitempty"`
// 	CountryName string  `json:"country_name,omitempty"`
// 	RegionCode  string  `json:"region_code,omitempty"`
// 	RegionName  string  `json:"region_name,omitempty"`
// 	City        string  `json:"city,omitempty"`
// 	ZipCode     string  `json:"zip_code,omitempty"`
// 	TimeZone    string  `json:"time_zone,omitempty"`
// 	Latitude    float64 `json:"latitude,omitempty"`
// 	Longitude   float64 `json:"longitude,omitempty"`
// 	MetroCode   int     `json:"metro_code,omitempty"`
// }

// func IP(ip string) (location Location, err error) {
// 	resp, err := http.Get("https://geoip.pianos.travel/json/" + ip)
// 	if err != nil {
// 		return
// 	}
// 	defer resp.Body.Close()
//
// 	err = json.NewDecoder(resp.Body).Decode(&location)
// 	return
// }

// func Search(s string) (location Location, err error) {
// 	searchURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&addressdetails=1", s)
// 	resp, err := http.Get(searchURL)
// 	if err != nil {
// 		err = errors.Wrap(err, searchURL)
// 		return
// 	}
// 	defer resp.Body.Close()
//
// 	var n []Nominatim
// 	err = json.NewDecoder(resp.Body).Decode(&n)
// 	if err != nil {
// 		err = errors.Wrap(err, searchURL)
// 		return
// 	}
// 	if len(n) == 0 {
// 		err = fmt.Errorf("could not find any result")
// 		return
// 	}
// 	log.Debugf("found: %+v", n)
// 	location.CountryName = n[0].Address.Country
// 	location.City = n[0].Address.City
// 	location.RegionName = n[0].Address.State
// 	location.Latitude, _ = strconv.ParseFloat(n[0].Lat, 64)
// 	location.Longitude, _ = strconv.ParseFloat(n[0].Lon, 64)
// 	return
// }

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
