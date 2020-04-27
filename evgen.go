package evgen

import (
	"errors"
	"time"
)

func Generate(repeatType string, startAt time.Time, dayOfWeek []int, frequency int, quantity int, endAt *time.Time) ([]time.Time, error) {
	if quantity == 0 && endAt == nil {
		return nil, errors.New("Either quantity or endAt should be non zero")
	}
	switch repeatType {
	case "once":
		return generateOnce(startAt), nil
	case "daily":
		return generateDaily(startAt, frequency, quantity, endAt), nil
	case "weekly":
		return generateWeekly(startAt, frequency, dayOfWeek, quantity, endAt), nil
	case "monthly":
		return generateMonthly(startAt, frequency, quantity, endAt), nil
	case "yearly":
		return generateYearly(startAt, frequency, quantity, endAt), nil
	default:
		return nil, errors.New("Unknown repeat type")

	}
}

func generateOnce(startAt time.Time) []time.Time {
	return []time.Time{startAt}
}

func generateDaily(startAt time.Time, frequency, quantity int, endAt *time.Time) []time.Time {
	eventSeries := make([]time.Time, 0)
	i := 1
	nextEvent := startAt
	check := func() bool {
		if quantity != 0 {
			return i <= quantity
		} else {
			return nextEvent.Before(*endAt) || nextEvent.Equal(*endAt)
		}
	}
	for check() {
		eventSeries = append(eventSeries, nextEvent)
		nextEvent = startAt.AddDate(0, 0, i*frequency)
		i++
	}
	return eventSeries
}

func generateWeekly(startAt time.Time, frequency int, dayOfWeek []int, quantity int, endAt *time.Time) []time.Time {
	eventSeries := make([]time.Time, 0)
	i := 1
	nextDay := startAt
	if len(dayOfWeek) == 0 {
		dayOfWeek = []int{int(nextDay.Weekday())}
	}
	_, startWeek := startAt.ISOWeek()
	// adjustment to make sunday move to the next week
	if startAt.Weekday() == 0 {
		startWeek++
	}
	check := func() bool {
		if quantity != 0 {
			return i <= quantity
		} else {
			return nextDay.Before(*endAt) || nextDay.Equal(*endAt)
		}
	}
	for check() {
		correctWeek := false
		_, currentWeek := nextDay.ISOWeek()
		// adjustment to make sunday move to the next week
		if nextDay.Weekday() == 0 {
			currentWeek++
		}
		if (currentWeek-startWeek)%frequency == 0 {
			correctWeek = true
		}
		if correctWeek && intContains(int(nextDay.Weekday()), dayOfWeek) {
			eventSeries = append(eventSeries, nextDay)
			i++
		}
		nextDay = nextDay.Add(24 * time.Hour)
	}
	return eventSeries

}

func generateMonthly(startAt time.Time, frequency, quantity int, endAt *time.Time) []time.Time {
	eventSeries := make([]time.Time, 0)
	i := 1
	nextEvent := startAt
	check := func() bool {
		if quantity != 0 {
			return i <= quantity
		} else {
			return nextEvent.Before(*endAt) || nextEvent.Equal(*endAt)
		}
	}
	for check() {
		eventSeries = append(eventSeries, nextEvent)
		nextEvent = startAt.AddDate(0, i*frequency, 0)
		i++
	}
	return eventSeries
}

func generateYearly(startAt time.Time, frequency, quantity int, endAt *time.Time) []time.Time {
	eventSeries := make([]time.Time, 0)
	i := 1
	nextEvent := startAt
	check := func() bool {
		if quantity != 0 {
			return i <= quantity
		} else {
			return nextEvent.Before(*endAt) || nextEvent.Equal(*endAt)
		}
	}
	for check() {
		eventSeries = append(eventSeries, nextEvent)
		nextEvent = startAt.AddDate(0, 0, i*frequency)
		i++
	}
	return eventSeries
}

func intContains(needle int, hay []int) bool {
	for _, value := range hay {
		if needle == value {
			return true
		}
	}
	return false
}
