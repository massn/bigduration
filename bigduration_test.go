package bigduration_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/massn/bigduration"
	"github.com/stretchr/testify/assert"
)

func TestParseForLittleSeconds(t *testing.T) {
	r, err := bigduration.ParseDuration("32s")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(32))
}

func TestParseForSubSeconds(t *testing.T) {
	r, err := bigduration.ParseDuration("3.2s")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(3))
}

func TestParseForBigSeconds(t *testing.T) {
	r, err := bigduration.ParseDuration("244s")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(244))
}

func TestParseForLittleMinutes(t *testing.T) {
	r, err := bigduration.ParseDuration("32m")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(32*60))
}

func TestParseForSubMinutes(t *testing.T) {
	r, err := bigduration.ParseDuration("32.23m")
	assert.Nil(t, err)
	f := math.Floor(32.23 * 60)
	assert.Equal(t, r, bigduration.BigDuration(int64(f)))
}

func TestParseForBigMinutes(t *testing.T) {
	r, err := bigduration.ParseDuration("1034m")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(1034*60))
}

func TestParseForMinutesAndSeconds(t *testing.T) {
	r, err := bigduration.ParseDuration("32m43s")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(32*60+43))
}

func TestParseForHours(t *testing.T) {
	r, err := bigduration.ParseDuration("37h")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(37*60*60))
}

func TestParseForDays(t *testing.T) {
	r, err := bigduration.ParseDuration("21d")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(21*24*60*60))
}

func TestParseForYears(t *testing.T) {
	r, err := bigduration.ParseDuration("120312y")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(120312*365*24*60*60))
}

func TestParseForMix(t *testing.T) {
	r, err := bigduration.ParseDuration("122y8d9h22m21s")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(122*365*24*60*60+8*24*60*60+9*60*60+22*60+21))
}

func TestParseForMaxInt64(t *testing.T) {
	maxInt64string := strconv.FormatInt(math.MaxInt64, 10)
	r, err := bigduration.ParseDuration(maxInt64string + "s")
	assert.Nil(t, err)
	assert.Equal(t, r, bigduration.BigDuration(9223372036854775807))
}

func TestParseError(t *testing.T) {
	r, err := bigduration.ParseDuration("21sv")
	assert.EqualError(t, err, "Failed to parse 21sv")
	assert.Equal(t, r, bigduration.BigDuration(0))
}

func TestSeconds(t *testing.T) {
	b := bigduration.BigDuration(2897)
	s := b.Seconds()
	assert.Equal(t, s, float64(2897))
}
