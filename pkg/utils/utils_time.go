package utils

import (
	"strconv"
	"time"
)

const (
	ExtsTimeYearMonthDay                = "2006-01-02"
	ExtsTimeYearMonth                   = "2006-01"
	ExtsTimeStandard                    = "2006-01-02 15:04:05"
	ExtsTimeSlashYearMonthDayHourMinute = "2006/01/02 15:04"
	ExtsTimeHourMinuteSec               = "15:04:05"
	ExtsOneDaySecond                    = 86400 // 24小时
	ExtsTwelveHoursSecond               = 43200 // 12小时
	ExtTimeZoneNewYork                  = "America/New_York"
	ExtTimeZoneBeijing                  = "Asia/Shanghai"
	ExtsSqlDefaultTime                  = "1001-01-01T00:00:00+08:00"
	ExtsSqlDefaultDate                  = "1001-01-01"
	ExtsTimeChineseStandardTime         = "2006-01-02T15:04:05+08:00"
)

func DaysTime(days int, layout string) string {
	days = days - 2
	start := time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 0, days)
	if layout == "" {
		layout = ExtsTimeYearMonthDay
	}
	return end.Format(layout)
}

func DifferenceDays(minuend time.Time, subtracted time.Time) int {
	return int((minuend.Unix() - subtracted.Unix()) / ExtsOneDaySecond)
}

func TimeNow(layout string) string {
	if layout == "" {
		layout = ExtsTimeStandard
	}
	return time.Now().Format(layout)
}

func TimestampTime(timestamp int64, layout string) string {
	if layout == "" {
		layout = ExtsTimeYearMonthDay
	}
	return time.Unix(timestamp, 0).Format(layout)
}

func ToTimestamp(date string, layout string) int64 {
	if layout == "" {
		layout = ExtsTimeStandard
	}
	stamp, _ := time.ParseInLocation(layout, date, time.Local)
	return stamp.Unix()
}

func NextWeek(currentDate string, layout string) (string, string) {
	if layout == "" {
		layout = ExtsTimeYearMonthDay
	}
	stamp := ToTimestamp(currentDate, layout)
	start := stamp + ExtsOneDaySecond
	end := stamp + 7*ExtsOneDaySecond
	return TimestampTime(start, ""), TimestampTime(end, "")
}

func ElteWeeksLoc() []map[string]string {
	start := "2020-01-04"
	end := "2020-01-04"
	today := TimeNow(ExtsTimeYearMonthDay)
	var weeks []map[string]string
	for true {
		start, end = NextWeek(end, "")
		if start < today {
			week := make(map[string]string)
			week["start"] = start
			week["end"] = end
			weeks = append(weeks, week)
		} else {
			return weeks
		}
	}
	return weeks
}

func MonthsFirstAndLast(startDate string, endDate string) (res []map[string]string) {
	layout := "2006-01-02"
	var err error
	var start time.Time
	var end time.Time

	start, err = time.Parse(layout, startDate)
	if err != nil {
		return
	}
	end, err = time.Parse(layout, endDate)
	if err != nil {
		return
	}
	var month = start
	for month.Unix() < end.Unix() {
		firstDay := month.AddDate(0, 0, -month.Day()+1)
		lastDay := firstDay.AddDate(0, 1, -1)
		m := make(map[string]string)
		m["start"] = firstDay.Format(layout)
		m["end"] = lastDay.Format(layout)
		res = append(res, m)
		if lastDay.Unix() > end.Unix() {
			break
		}
		month = month.AddDate(0, 1, 0)
	}
	return
}

func ToLocationDate(lays string, layt string, loc string, date string) (string, error) {
	var location *time.Location
	location, err := time.LoadLocation(loc)
	if err != nil {
		return "", err
	}
	timestamp, err := time.ParseInLocation(lays, date, location)
	if err != nil {
		return "", err
	}
	target := timestamp.In(location).Format(layt)
	return target, err
}

func ToTime(date string, layout string) (time.Time, error) {
	if layout == "" {
		layout = ExtsTimeYearMonthDay
	}
	return time.ParseInLocation(layout, date, time.Local)
}

func ToNewYork(timeStamp int64, layout string) (res string) {
	if layout == "" {
		layout = ExtsTimeStandard
	}
	res = time.Unix(timeStamp, 0).In(NewYorkLocation()).Format(layout)
	return
}

func ToLocationTime(date string, layout string, local *time.Location) (time.Time, error) {
	if layout == "" {
		layout = ExtsTimeYearMonthDay
	}
	return time.ParseInLocation(layout, date, local)
}

func MidnightToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func NowTime() time.Time {
	return time.Now()
}

func WeekMonth(daytime time.Time) int {
	dayInMonth := daytime.Day()
	//本月第一天
	firstDay := daytime.AddDate(0, 0, 1-daytime.Day())
	//本月第一天是周几
	tw := int(firstDay.Weekday())
	weekday := tw
	if tw == 0 {
		weekday = 7
	}
	//本月第一周有几天
	firstWeekEndDay := 7 - (weekday - 1)
	//当前日期和第一周之差
	diffday := dayInMonth - firstWeekEndDay
	if diffday <= 0 {
		diffday = 1
	}
	//当前是第几周
	var WeekInMonth = 0
	if (diffday % 7) == 0 {
		WeekInMonth = diffday/7 + 1
	} else {
		if dayInMonth > firstWeekEndDay {
			WeekInMonth = (diffday / 7) + 1 + 1
		} else {
			WeekInMonth = (diffday / 7) + 1
		}
	}
	return WeekInMonth
}

func Quarter(date time.Time) int {
	month := int(date.Month())
	switch month {
	case 1, 2, 3:
		return 1
	case 4, 5, 6:
		return 2
	case 7, 8, 9:
		return 3
	case 10, 11, 12:
		return 4
	default:
		return 1
	}
}

func Semiannual(date time.Time) int {
	month := int(date.Month())
	switch {
	case month <= 6:
		return 1
	case month <= 12:
		return 2
	default:
		return 1
	}
}

func NowSubYears(date string, layout string) (years int) {
	now := time.Now()
	d, _ := ToTime(date, layout)
	years = SubYears(now, d)
	return
}

func NowSubMonths(date string, layout string) (months int) {
	now := time.Now()
	d, _ := ToTime(date, layout)
	months = SubMonth(now, d)
	return
}

// 计算日期相差多少年
func SubYears(t1, t2 time.Time) (yearInterval int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval = y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	return
}

// 计算日期相差多少月
func SubMonth(t1, t2 time.Time) (month int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	return
}

func ExcelDateFormat(excelDate string) (res time.Time) {
	if excelDate == "" {
		return
	}
	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days, _ = strconv.ParseFloat(excelDate, 64)
	res = excelTime.Add(time.Second * time.Duration(days*86400))
	return
}

func BeijingLocation() (loc *time.Location) {
	loc, _ = time.LoadLocation(ExtTimeZoneBeijing)
	return
}

func NewYorkLocation() (loc *time.Location) {
	loc, _ = time.LoadLocation(ExtTimeZoneNewYork)
	return
}

func IntervalDays(startTime time.Time, endTime time.Time, layout string) (days []string) {
	for endTime.After(startTime) {
		if layout == "" {
			layout = ExtsTimeYearMonthDay
		}
		days = append(days, startTime.Format(layout))
		startTime = startTime.AddDate(0, 0, 1)
	}
	return
}

func IntervalTimeDays(startTime time.Time, endTime time.Time) (days []time.Time) {
	for endTime.After(startTime) {
		days = append(days, startTime)
		startTime = startTime.AddDate(0, 0, 1)
	}
	return
}

func MonthStartAndEndDate(date time.Time) (start time.Time, end time.Time) {
	if date.IsZero() {
		return
	}
	year, month, _ := date.Date()
	start = time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	end = start.AddDate(0, 1, -1)
	return
}

func WeekStartAndEndDate(date time.Time) (start time.Time, end time.Time) {
	if date.IsZero() {
		return
	}
	start = WeekFirstTime(date)
	end = start.Add(60*60*24*7*time.Second - 1*time.Second)
	return
}

func DayFirst(date time.Time) (startTime time.Time) {
	startTime = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	return
}

func DayLast(date time.Time) (endTime time.Time) {
	endTime = time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())
	return
}

func MonthLastDay(t time.Time) (lastDay time.Time) {
	lastDay = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).AddDate(0, 1, -1)
	return
}

func MonthLastTime(t time.Time) (lastTime time.Time) {
	lastTime = time.Date(t.Year(), t.Month(), 1, 23, 59, 59, 0, t.Location()).AddDate(0, 1, -1)
	return
}

func MonthFirstTime(t time.Time) (firstTime time.Time) {
	firstTime = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return
}

/**
获取本周周一的日期
*/
func WeekFirstTime(date time.Time) (weekMonday time.Time) {
	offset := int(time.Monday - date.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekMonday = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location()).AddDate(0, 0, offset)
	return
}

/**
获取上周的周一日期
*/
func PastWeekMondays(weeks int, date time.Time) (weekMondays []time.Time) {
	nowWeekMonday := WeekFirstTime(date)
	weekMondays = append(weekMondays, nowWeekMonday)
	for i := 1; i < weeks; i++ {
		weekMonday := nowWeekMonday.AddDate(0, 0, -7*i)
		weekMondays = append(weekMondays, weekMonday)
	}
	return
}

func FrontWeeksMondays(weeks int, date time.Time) (currentMonday time.Time, frontMonday time.Time) {
	currentMonday = WeekFirstTime(date)
	frontMonday = currentMonday.AddDate(0, 0, -7*weeks)
	return
}

func MonthDays(year int, month int) (days []time.Time) {
	var first = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	var last = first.AddDate(0, 1, 0).Add(-1 * time.Second)
	days = IntervalTimeDays(first, last)
	return
}

func FixedZone(offset int) (location *time.Location) {
	location = time.FixedZone("UTC", offset)
	return
}

const (
	TimeOffset = 8 * 3600  //8 hour offset
	HalfOffset = 12 * 3600 //Half-day hourly offset
)

//Get the current timestamp by Second
func GetCurrentTimestampBySecond() int64 {
	return time.Now().Unix()
}

//Convert timestamp to time.Time type
func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

//Convert nano timestamp to time.Time type
func UnixNanoSecondToTime(nanoSecond int64) time.Time {
	return time.Unix(0, nanoSecond)
}
func UnixMillSecondToTime(millSecond int64) time.Time {
	return time.Unix(0, millSecond*1e6)
}

//Get the current timestamp by Nano
func GetCurrentTimestampByNano() int64 {
	return time.Now().UnixNano()
}

//Get the current timestamp by Mill
func GetCurrentTimestampByMill() int64 {
	return time.Now().UnixNano() / 1e6
}

//Get the timestamp at 0 o'clock of the day
func GetCurDayZeroTimestamp() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	return t.Unix() - TimeOffset
}

//Get the timestamp at 12 o'clock on the day
func GetCurDayHalfTimestamp() int64 {
	return GetCurDayZeroTimestamp() + HalfOffset

}

//Get the formatted time at 0 o'clock of the day, the format is "2006-01-02_00-00-00"
func GetCurDayZeroTimeFormat() string {
	return time.Unix(GetCurDayZeroTimestamp(), 0).Format("2006-01-02_15-04-05")
}

//Get the formatted time at 12 o'clock of the day, the format is "2006-01-02_12-00-00"
func GetCurDayHalfTimeFormat() string {
	return time.Unix(GetCurDayZeroTimestamp()+HalfOffset, 0).Format("2006-01-02_15-04-05")
}
func GetTimeStampByFormat(datetime string) string {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix()
	return strconv.FormatInt(timestamp, 10)
}

func TimeStringFormatTimeUnix(timeFormat string, timeSrc string) int64 {
	tm, _ := time.Parse(timeFormat, timeSrc)
	return tm.Unix()
}

func TimeStringToTime(timeString string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", timeString)
	return t, err
}

func NowTimestamp() int64{
	return time.Now().Unix()
}