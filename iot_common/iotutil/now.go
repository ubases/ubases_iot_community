package iotutil

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

func (now *Now) BeginningOfMinute() time.Time {
	return now.Truncate(time.Minute)
}

func (now *Now) BeginningOfHour() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, now.Time.Hour(), 0, 0, 0, now.Time.Location())
}

func (now *Now) BeginningOfDay() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Time.Location())
}

func (now *Now) BeginningOfWeek() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())

	if WeekStartDay != time.Sunday {
		weekStartDayInt := int(WeekStartDay)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	return t.AddDate(0, 0, -weekday)
}

func (now *Now) BeginningOfMonth() time.Time {
	y, m, _ := now.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
}

func (now *Now) BeginningOfQuarter() time.Time {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

func (now *Now) BeginningOfYear() time.Time {
	y, _, _ := now.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, now.Location())
}

func (now *Now) EndOfMinute() time.Time {
	return now.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

func (now *Now) EndOfHour() time.Time {
	return now.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

func (now *Now) EndOfDay() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
}

func (now *Now) EndOfWeek() time.Time {
	return now.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

func (now *Now) EndOfMonth() time.Time {
	return now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func (now *Now) EndOfQuarter() time.Time {
	return now.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

func (now *Now) EndOfYear() time.Time {
	return now.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

func (now *Now) Monday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return t.AddDate(0, 0, -weekday+1)
}

func (now *Now) Sunday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		return t
	}
	return t.AddDate(0, 0, (7 - weekday))
}

func (now *Now) EndOfSunday() time.Time {
	return New(now.Sunday()).EndOfDay()
}

func parseWithFormat(str string) (t time.Time, err error) {
	for _, format := range TimeFormats {
		t, err = time.Parse(format, str)
		if err == nil {
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}

var hasTimeRegexp = regexp.MustCompile(`(\s+|^\s*)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)
var onlyTimeRegexp = regexp.MustCompile(`^\s*\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)

func (now *Now) Parse(strs ...string) (t time.Time, err error) {
	var (
		setCurrentTime  bool
		parseTime       []int
		currentTime     = []int{now.Nanosecond(), now.Second(), now.Minute(), now.Hour(), now.Day(), int(now.Month()), now.Year()}
		currentLocation = now.Location()
		onlyTimeInStr   = true
	)

	for _, str := range strs {
		hasTimeInStr := hasTimeRegexp.MatchString(str)
		onlyTimeInStr = hasTimeInStr && onlyTimeInStr && onlyTimeRegexp.MatchString(str)
		if t, err = parseWithFormat(str); err == nil {
			location := t.Location()
			if location.String() == "UTC" {
				location = currentLocation
			}

			parseTime = []int{t.Nanosecond(), t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}

			for i, v := range parseTime {
				if hasTimeInStr && i <= 3 {
					continue
				}

				if v == 0 {
					if setCurrentTime {
						parseTime[i] = currentTime[i]
					}
				} else {
					setCurrentTime = true
				}

				if onlyTimeInStr {
					if i == 4 || i == 5 {
						parseTime[i] = currentTime[i]
						continue
					}
				}
			}

			t = time.Date(parseTime[6], time.Month(parseTime[5]), parseTime[4], parseTime[3], parseTime[2], parseTime[1], parseTime[0], location)
			currentTime = []int{t.Nanosecond(), t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
		}
	}
	return
}

func (now *Now) MustParse(strs ...string) (t time.Time) {
	t, err := now.Parse(strs...)
	if err != nil {
		panic(err)
	}
	return t
}

func (now *Now) Between(begin, end string) bool {
	beginTime := now.MustParse(begin)
	endTime := now.MustParse(end)
	return now.After(beginTime) && now.Before(endTime)
}

var WeekStartDay = time.Sunday

var TimeFormats = []string{"1/2/2006", "1/2/2006 15:4:5", "2006", "2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5", "1-2", "15:4:5", "15:4", "15", "15:4:5 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST"}

type Now struct {
	time.Time
}

func New(t time.Time) *Now {
	return &Now{t}
}

func BeginningOfMinute() time.Time {
	return New(time.Now()).BeginningOfMinute()
}

func BeginningOfHour() time.Time {
	return New(time.Now()).BeginningOfHour()
}

func BeginningOfDay() time.Time {
	return New(time.Now()).BeginningOfDay()
}

//指定天数开始
func BeginningOfPointDay(days int64) time.Time {
	return New(time.Now().Add(time.Duration(days) * 24 * time.Hour)).BeginningOfDay()
}

//指定天数结束
func EndOfPointDay(days int64) time.Time {
	return New(time.Now().Add(time.Duration(days) * 24 * time.Hour)).EndOfDay()
}

func BeginningOfWeek() time.Time {
	return New(time.Now()).BeginningOfWeek()
}

func BeginningOfMonth() time.Time {
	return New(time.Now()).BeginningOfMonth()
}

func BeginningOfQuarter() time.Time {
	return New(time.Now()).BeginningOfQuarter()
}

func BeginningOfYear() time.Time {
	return New(time.Now()).BeginningOfYear()
}

func EndOfMinute() time.Time {
	return New(time.Now()).EndOfMinute()
}

func EndOfHour() time.Time {
	return New(time.Now()).EndOfHour()
}

func EndOfDay() time.Time {
	return New(time.Now()).EndOfDay()
}

func EndOfWeek() time.Time {
	return New(time.Now()).EndOfWeek()
}

func EndOfMonth() time.Time {
	return New(time.Now()).EndOfMonth()
}

func EndOfQuarter() time.Time {
	return New(time.Now()).EndOfQuarter()
}

func EndOfYear() time.Time {
	return New(time.Now()).EndOfYear()
}

func Monday() time.Time {
	return New(time.Now()).Monday()
}

func Sunday() time.Time {
	return New(time.Now()).Sunday()
}

func EndOfSunday() time.Time {
	return New(time.Now()).EndOfSunday()
}

func Parse(strs ...string) (time.Time, error) {
	return New(time.Now()).Parse(strs...)
}

func ParseInLocation(loc *time.Location, strs ...string) (time.Time, error) {
	return New(time.Now().In(loc)).Parse(strs...)
}

func MustParse(strs ...string) time.Time {
	return New(time.Now()).MustParse(strs...)
}

func MustParseInLocation(loc *time.Location, strs ...string) time.Time {
	return New(time.Now().In(loc)).MustParse(strs...)
}
func Between(time1, time2 string) bool {
	return New(time.Now()).Between(time1, time2)
}

//字符串2021-07-13T15:56:52+08:00 转时间
func TStringToTime(str string) time.Time {
	arrStr := strings.Split(str, "T")
	//zIndex := strings.Index(str, "Z")
	if len(arrStr) > 1 {
		strDate := arrStr[0] //2021-07-13
		strTime := arrStr[1] //15:56:52+08:00 (可能存在-时区和+时区2种情况)
		//找到时区开始的位置
		idx := strings.Index(strTime, "+")
		if idx == -1 {
			strTime = strings.Split(strTime, "-")[0]
		} else {
			strTime = strings.Split(strTime, "+")[0]
		}
		//00:00:00Z
		strTime2 := strings.Replace(strTime, "Z", "", 1)
		return MustParse(strDate + " " + strTime2)
	}

	return MustParse("0001-01-01 00:00:00")
}
