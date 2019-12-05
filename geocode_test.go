package gogeocoder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	search := "850 republican street, seattle"
	location, err := Geocode(search)
	assert.Nil(t, err)
	assert.Equal(t, "Seattle", location.City)
	fmt.Printf("\n%+v\n", location)
}
