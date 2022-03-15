package bigduration

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type BigDuration uint64

func ParseDuration(s string) (BigDuration, error) {
	r := regexp.MustCompile(`^(([0-9|\.]+)y)?(([0-9|\.]+)d)?(([0-9|\.]+)h)?(([0-9|.]+)m)?(([0-9|\.]+)s)?$`)
	matches := r.FindStringSubmatch(s)
	if len(matches) != 11 {
		return 0, fmt.Errorf("Failed to parse %s", s)
	}
	years, err := toFloat64(matches[2])
	if err != nil {
		return 0, err
	}
	days, err := toFloat64(matches[4])
	if err != nil {
		return 0, err
	}
	hours, err := toFloat64(matches[6])
	if err != nil {
		return 0, err
	}
	minutes, err := toFloat64(matches[8])
	if err != nil {
		return 0, err
	}
	seconds, err := toFloat64(matches[10])
	if err != nil {
		return 0, err
	}
	totalSeconds := seconds + 60*minutes + 3600*hours + 86400*days + 31536000*years

	f := math.Round(totalSeconds)
	return BigDuration(uint64(f)), nil
}

func (b BigDuration) Seconds() float64 {
	return float64(b)
}

func toFloat64(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseFloat(s, 64)
}
