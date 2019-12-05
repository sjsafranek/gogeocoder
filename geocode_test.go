package gogeocoder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIP(t *testing.T) {
	ip := "97.126.102.4"
	location, err := IP(ip)
	assert.Nil(t, err)
	assert.Equal(t, "Seattle", location.City)
	fmt.Printf("\n%+v\n", location)
}

func TestSearch(t *testing.T) {
	search := "850 republican street, seattle"
	location, err := Search(search)
	assert.Nil(t, err)
	assert.Equal(t, "Seattle", location.City)
	fmt.Printf("\n%+v\n", location)
}
