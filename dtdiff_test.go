package dtdiff

import (
	"testing"
	"time"
)

func TestCalcDiff(t *testing.T) {
	t1 := time.Date(2018, time.January, 1, 2, 6, 24, 0, time.UTC)
	t2 := time.Date(2018, time.January, 1, 4, 33, 41, 0, time.UTC)
	t3 := time.Date(2018, time.January, 1, 5, 52, 32, 0, time.UTC)
	t4 := time.Date(2018, time.January, 5, 2, 11, 2, 0, time.UTC)
	t5 := time.Date(2018, time.January, 5, 12, 40, 18, 0, time.UTC)
	t6 := time.Date(2018, time.January, 5, 12, 40, 18, 0, time.UTC)
	t7 := time.Date(2018, time.February, 5, 1, 38, 8, 0, time.UTC)
	t8 := time.Date(2018, time.March, 4, 21, 24, 59, 0, time.UTC)
	cases := []struct {
		in   []time.Time
		want string
	}{
		{[]time.Time{}, "0s"},
		{[]time.Time{t5, t6}, "0s"},
		{[]time.Time{t1, t2}, "2h27m17s"},
		{[]time.Time{t1, t2, t3}, "2h27m17s"}, // same as between t1 and t2
		{[]time.Time{t1, t2, t3, t4}, "94h45m47s"},
		{[]time.Time{t1, t2, t3, t4, t5}, "94h45m47s"},
		{[]time.Time{t1, t2, t3, t4, t5, t6}, "94h45m47s"},
		{[]time.Time{t1, t2, t3, t4, t5, t6, t7}, "94h45m47s"},
		{[]time.Time{t1, t2, t3, t4, t5, t6, t7, t8}, "762h32m38s"},
		{[]time.Time{t8, t1}, "-1507h18m35s"},
	}
	for _, c := range cases {
		r := CalculateDiffs(c.in...)
		if r.String() != c.want {
			t.Errorf("CalculateDiffs(%q).String == %q, want %q", c.in, r.String(), c.want)
		}
	}
}
