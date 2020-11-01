package business

import (
	"time"

	"github.com/jinzhu/now"
)

var dates = []time.Time{}

// IsBusiness - Check if day is business day.
func IsBusiness(date time.Time, weekDays ...int) bool {
	if len(weekDays) <= 0 {
		return false
	}
	day := int(date.Weekday())
	return contains(day, weekDays...)
}

// InWeek - Return total and business dates in month.
func InWeek(date time.Time, weekDays ...int) (int, []time.Time) {
	count := 0
	end := now.With(date).EndOfWeek()
	begin := now.With(date).BeginningOfWeek()

	days := int(end.Sub(begin).Hours() / 24)

	for i := 0; i < days; i++ {
		date := begin.Add(time.Duration(i) * time.Hour)
		if contains(int(date.Weekday()), weekDays...) {
			count++
			dates = append(dates, date)
		}
	}
	return count, dates
}

// InMonth - Return total and business dates in month.
func InMonth(date time.Time, weekDays ...int) (int, []time.Time) {
	count := 0
	end := now.With(date).EndOfMonth()
	begin := now.With(date).BeginningOfMonth()

	days := int(end.Sub(begin).Hours() / 24)

	for i := 0; i < days; i++ {
		date := begin.Add(time.Duration(i) * time.Hour)
		if contains(int(date.Weekday()), weekDays...) {
			count++
			dates = append(dates, date)
		}
	}
	return count, dates
}

// InInterval - Return total and business datae in week.
func InInterval(begin, end time.Time, weekDays ...int) (int, []time.Time) {
	count := 0
	days := int(end.Sub(begin).Hours() / 24)

	for i := 0; i < days; i++ {
		date := begin.Add(time.Duration(i) * time.Hour)
		if contains(int(date.Weekday()), weekDays...) {
			count++
			dates = append(dates, date)
		}
	}
	return count, dates
}

func contains(weekDay int, weekDays ...int) bool {
	for _, current := range weekDays {
		if current == weekDay {
			return true
		}
	}
	return false
}
