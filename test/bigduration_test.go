package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	bigduration "github.com/massn/bigduration/pkg"
)

func TestLittleSeconds(t *testing.T) {
	r, err := bigduration.ParseDuration("32s")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(32))
}

func TestBigSeconds(t *testing.T) {
	r, err := bigduration.ParseDuration("244s")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(244))
}

func TestLittleMinutes(t *testing.T) {
	r, err := bigduration.ParseDuration("32m")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(32*60))
}

func TestBigMinutes(t *testing.T) {
	r, err := bigduration.ParseDuration("1034m")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(1034*60))
}

func TestMinutesAndSeconds(t *testing.T) {
	r, err := bigduration.ParseDuration("32m43s")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(32*60+43))
}

func TestHours(t *testing.T) {
	r, err := bigduration.ParseDuration("37h")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(37*60*60))
}

func TestDays(t *testing.T) {
	r, err := bigduration.ParseDuration("21d")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(21*24*60*60))
}

func TestYears(t *testing.T) {
	r, err := bigduration.ParseDuration("120312y")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(120312*365*24*60*60))
}

func TestMix(t *testing.T) {
	r, err := bigduration.ParseDuration("122y8d9h22m21s")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(122*365*24*60*60+8*24*60*60+9*60*60+22*60+21))
}

func TestParsingError(t *testing.T) {
	r, err := bigduration.ParseDuration("21sv")
	assert.EqualError(t, err, "Failed to parse 21sv")
	assert.Equal(t, r, bigduration.BigDuration(0))
}
