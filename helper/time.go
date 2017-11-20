package helper

import "time"

// Date Time Format
const (
	yyyy_mm_dd_hh_ii_ss = "2006-01-02 15:04:05"
	yyyy_mm_dd_hh_ii    = "2006-01-02 15:04"
	yyyy_mm_dd          = "2006-01-02"
	RFC3339             = time.RFC3339
)

// time.Now()
func TimeNow() time.Time {
	return time.Now()
}

func TodayDate() string {
	return time.Now().Format(yyyy_mm_dd)
}

func TodayUnix() int64 {
	dateTime, _ := time.ParseInLocation(yyyy_mm_dd, TodayDate(), time.Local)

	return dateTime.Unix()
}

// TimeUnix
func TimeUnix() int64 {
	return time.Now().Unix()
}

// YYYY-MM-DD HH:II:SS
func TimeDate() string {
	return time.Now().Format(yyyy_mm_dd_hh_ii_ss)
}

// date to time.Time
func DateToTime(date, dateFormat string) time.Time {
	timeTime, _ := time.ParseInLocation(dateFormat, date, time.Local)

	return timeTime
}

// Timestamp to time.Time
func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// TimeSince
func TimeSince(start time.Time) time.Duration {
	return time.Since(start)
}

// []byte = [Y Y Y Y M M D D H H I I S S]
func DateByte() []byte {
	var dateSlice []byte

	dateStr := TimeDate()
	dateStrSlice := []byte(dateStr)

	for _, r := range dateStrSlice {
		if r >= '0' && r <= '9' {
			dateSlice = append(dateSlice, r)
		}
	}

	return dateSlice
}
