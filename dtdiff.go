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

// CalculateDiffUntil calculate time diff between now and to(future)
func CalculateDiffUntil(to time.Time) time.Duration {
	d := time.Until(to)
	return d
}

// CalculateDiffs calculate time diff between pair of args
func CalculateDiffs(args ...time.Time) time.Duration {
	// initialize result
	result := time.Nanosecond * 0
	var t1 time.Time
	for i, t := range args {
		if i%2 == 1 {
			t2 := t
			result += CalculateDiff(t1, t2)
			continue
		}
		t1 = t
	}
	return result
}
