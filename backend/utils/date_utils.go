// Package utils – date utility helpers exposed to the rest of the application.
package utils

import "time"

// DaysBetween returns the number of calendar days from a to b.
// A positive result means b is after a (i.e. a delay when b = delivered, a = expected).
func DaysBetween(a, b time.Time) float64 {
	return b.Sub(a).Hours() / 24
}

// IsDelayed returns true when the delivered date is strictly after the expected date.
func IsDelayed(expected, delivered time.Time) bool {
	return delivered.After(expected)
}
