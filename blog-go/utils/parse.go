package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseDuration converts a duration string into time.Duration.
// The string should contain numeric values followed by units:
// "d" (days), "h" (hours), "m" (minutes), "s" (seconds).
// Example: "1d2h30m" → 1 day, 2 hours, 30 minutes.
// Returns an error if the string is empty or invalid.
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d) // trim spaces
	if len(d) == 0 {
		return 0, fmt.Errorf("empty duration string")
	}

	// unit → duration mapping
	unitPattern := map[string]time.Duration{
		"d": 24 * time.Hour,
		"h": time.Hour,
		"m": time.Minute,
		"s": time.Second,
	}

	var totalDuration time.Duration
	// iterate over units in order
	for _, unit := range []string{"d", "h", "m", "s"} {
		for strings.Contains(d, unit) {
			unitIndex := strings.Index(d, unit)
			part := d[:unitIndex]
			if part == "" {
				part = "0"
			}
			val, err := strconv.Atoi(part)
			if err != nil {
				return 0, fmt.Errorf("invalid duration part: %v", err)
			}
			totalDuration += time.Duration(val) * unitPattern[unit]
			d = d[unitIndex+len(unit):]
		}
	}

	// check for leftover characters
	if len(d) > 0 {
		return 0, fmt.Errorf("unrecognized duration format")
	}

	return totalDuration, nil
}
