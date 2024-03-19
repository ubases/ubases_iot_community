package iotutil

import (
	"fmt"
	"strings"
	"time"
)

func TimeFullFormat(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
func DateFormat(time time.Time) string {
	return time.Format("2006-01-02")
}

func TimeFormat(time time.Time) string {
	return time.Format("2006-01-02")
}

func TimeFormatNew(time time.Time) string {
	return time.Format("2006.01.02")
}

func TimeFormatYear(time time.Time) string {
	return time.Format("2006")
}

func DateTimeFormat(time time.Time) string {
	return time.Format("15:04")
}

func DateTimeMonth(time time.Time) string {
	return ToString(time.Month()) //月
}

func DateTimeDay(time time.Time) string {
	return ToString(time.Day()) //日
}

// 获取某一天的0点时间
func GetTodaySartTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 获取某一天的23t点时间
func GetTodayLastTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

// 获取月初月末日期
func GetMouthStartAndEndTime() (time.Time, time.Time) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return firstOfMonth, lastOfMonth
}

// 时间戳转换时间
func GetTimeByUnit(unixTime int64) time.Time {
	//utcTime := time.Now().UTC()
	// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	return time.Unix(unixTime, 0)
}

// 字符串转换Time类型
func GetTimeByStr(timeStr string) (time.Time, error) {
	//utcTime := time.Now().UTC()
	// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	TIME_LAYOUT := "2006-01-02 15:04:05"
	return time.ParseInLocation(TIME_LAYOUT, timeStr, time.Now().Location())
}

func GetTimeByStrToMintue(timeStr string) (time.Time, error) {
	//utcTime := time.Now().UTC()
	// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	TIME_LAYOUT := "2006-01-02 15:04"
	return time.ParseInLocation(TIME_LAYOUT, timeStr, time.Now().Location())
}

func GetDateByStr(timeStr string) (time.Time, error) {
	//utcTime := time.Now().UTC()
	// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	TIME_LAYOUT := "2006-01-02"
	return time.ParseInLocation(TIME_LAYOUT, timeStr, time.Now().Location())
}

func GetLocalTimeStr(t time.Time) string {
	return t.Local().Format("2006-01-02 15:04:05")
}

func GetTodayTime(hms string) time.Time {
	d := time.Now()
	times := strings.Split(hms, ":")
	if len(times) == 2 {
		hour := ToInt(times[0])
		min := ToInt(times[1])
		return time.Date(d.Year(), d.Month(), d.Day(), hour, min, 0, 0, d.Location())
	} else if len(times) == 3 {
		hour := ToInt(times[0])
		min := ToInt(times[1])
		second := ToInt(times[2])
		return time.Date(d.Year(), d.Month(), d.Day(), hour, min, second, 0, d.Location())
	}
	return d
}

func GetTodayStrTime(hms string) string {
	d := time.Now()
	return fmt.Sprintf("%v-%v-%v %s", d.Year(), d.Format("01"), d.Day(), hms)
}
