package dtdiff

import "time"

// CalculateDiff calculate time diff between from and to
func CalculateDiff(from time.Time, to time.Time) time.Duration {
	d := to.Sub(from)
	return d
}

// CalculateDiffSince calculate time diff between from and now
func CalculateDiffSince(from time.Time) time.Duration {
	d := time.Since(from)
	return d
}
