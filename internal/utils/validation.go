package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseTime parse time from HH:MM
func ParseTime(t string) (int, int, error) {
	parts := strings.Split(t, ":")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("incorrect format")
	}

	hour, err1 := strconv.Atoi(parts[0])
	minute, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil || hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return 0, 0, fmt.Errorf("incorrect time: %s", t)
	}

	return hour, minute, nil
}
