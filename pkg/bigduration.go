package bigduration

import (
	"fmt"
	"regexp"
	"strconv"
)

type BigDuration int64

func ParseDuration(s string) (BigDuration, error) {
	r := regexp.MustCompile(`^(([\d]+)y)?(([\d]+)d)?(([\d]+)h)?(([\d]+)+m)?(([\d]+)s)?$`)
	matches := r.FindStringSubmatch(s)
	if len(matches) != 11 {
		return 0, fmt.Errorf("Failed to parse %s", s)
	}
	years, err := toInt64(matches[2])
	if err != nil {
		return 0, err
	}
	days, err := toInt64(matches[4])
	if err != nil {
		return 0, err
	}
	hours, err := toInt64(matches[6])
	if err != nil {
		return 0, err
	}
	minutes, err := toInt64(matches[8])
	if err != nil {
		return 0, err
	}
	seconds, err := toInt64(matches[10])
	if err != nil {
		return 0, err
	}
	totalSeconds := seconds + 60*minutes + 3600*hours + 86400*days + 31536000*years

	return BigDuration(totalSeconds), nil
}

func (b BigDuration) Seconds() float64 {
	return float64(b)
}

func toInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseInt(s, 10, 64)
}
