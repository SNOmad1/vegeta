package vegeta

import "time"

func maxTime(t1 time.Time, t2 time.Time) time.Time {
	diff := t1.Sub(t2).Seconds()
	if diff <= 0 {
		return t2
	}
	return t1
}

func minNonzeroTime(t1 time.Time, t2 time.Time) time.Time {
	if t1.IsZero() {
		return t2
	}
	if t2.IsZero() {
		return t1
	}
	diff := t1.Sub(t2).Seconds()
	if diff >= 0 {
		return t2
	}
	return t1
}
