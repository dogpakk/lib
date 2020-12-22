package datetime

import "time"

const (
	BritishDate = "02/01/2006"
)

func CalcUTCOffset(timeZone string) (int, error) {
	// If there is no specified timezone (blank string), Go defaults to UTC,
	specifiedTimeZone, err := time.LoadLocation(timeZone)

	// Location loading can fail on some systems due to OS level stuff,
	// so we can only set the offset if there is no error above
	if err != nil {
		return 0, err
	}

	// Specify the offset in minutes, some timezone such as Indian standard time
	// are not a whole number of hours.
	_, offsetSeconds := time.Now().In(specifiedTimeZone).Zone()
	return offsetSeconds / 60, nil
}

func FormatBritishDate(dt time.Time) string {
	return dt.Format(BritishDate)
}

func AddDays(t time.Time, days int) time.Time {
	return t.Add(time.Hour * 24 * time.Duration(days))
}

// DaysFromNowIgnoringTime works out the day differential from now but based on actual days, rather than units of 24hrs
func DaysFromNowIgnoringTime(t time.Time, days int) int {
	futureDate := AddDays(t, days)
	futureDay := futureDate.YearDay()
	futureYear := futureDate.Year()

	daysDiff := time.Now().YearDay() - futureDay
	yearsDiff := time.Now().Year() - futureYear
	totalDaysDiff := (yearsDiff * 365) + daysDiff

	if totalDaysDiff <= 0 {
		totalDaysDiff = 0
	}

	return totalDaysDiff
}
