package evgen

import (
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	now := time.Now()
	const shortForm = "2006-Jan-02"

	aDay, _ := time.Parse(shortForm, "2020-Mar-08")

	thisWednesday, _ := time.Parse(shortForm, "2020-Mar-11")
	nextSunday, _ := time.Parse(shortForm, "2020-Mar-15")
	nextWednesday, _ := time.Parse(shortForm, "2020-Mar-18")

	nxtNxtSunday, _ := time.Parse(shortForm, "2020-Mar-22")
	nxtNxtWednesday, _ := time.Parse(shortForm, "2020-Mar-25")
	tables := []struct {
		repeatType string
		startAt    time.Time
		dayOfWeek  []int
		frequency  int
		quantity   int
		endAt      *time.Time
		result     []time.Time
	}{
		{"daily", now, nil, 1, 1, nil, []time.Time{now}},
		{"daily", now, nil, 2, 2, nil, []time.Time{now, now.AddDate(0, 0, 2)}},
		{"daily", now, nil, 1, 0, daysLater(now, 3), []time.Time{now, *daysLater(now, 1), *daysLater(now, 2)}},
		{"weekly", now, nil, 1, 2, nil, []time.Time{now, now.AddDate(0, 0, 7)}},
		{"weekly", aDay, []int{0, 3}, 1, 4, nil, []time.Time{aDay, thisWednesday, nextSunday, nextWednesday}},
		{"weekly", aDay, []int{0, 3}, 2, 4, nil, []time.Time{aDay, thisWednesday, nxtNxtSunday, nxtNxtWednesday}},
		{"weekly", thisWednesday, []int{0, 3}, 2, 3, nil, []time.Time{thisWednesday, nxtNxtSunday, nxtNxtWednesday}},
		{"monthly", now, nil, 1, 1, nil, []time.Time{now}},
		{"monthly", now, nil, 1, 2, nil, []time.Time{now, now.AddDate(0, 1, 0)}},
		{"monthly", now, nil, 2, 3, nil, []time.Time{now, now.AddDate(0, 2, 0), now.AddDate(0, 4, 0)}},
	}

	for _, ta := range tables {
		result, err := Generate(ta.repeatType, ta.startAt, ta.dayOfWeek, ta.frequency, ta.quantity, ta.endAt)
		if err != nil {
			t.Error("Event gneration failed with error: " + err.Error())
		}
		if len(result) != len(ta.result) {
			t.Error("schedule failed to return correct events. Expected:", ta.result, " got:", result, "inputs:", ta)
		}
		for i := range result {
			if !result[i].Equal(ta.result[i]) {
				t.Error("scheduler returned incorrect event. Expected:", ta.result[i], " got:", result[i])

			}
		}
	}
}

func daysLater(now time.Time, d int) *time.Time {
	day := now.AddDate(0, 0, d)
	return &day
}
