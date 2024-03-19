package entitys

import (
	"time"
)

func FillTimeData(srcList []TimeData, flag int, begin, end time.Time) []TimeData {
	src := srcList
	curr := begin
	var next time.Time
	strTime := ""
	i := 0
	for curr.Before(end) || curr == end {
		if flag == 1 { //月
			strTime = curr.Format("2006-01")
			next = curr.AddDate(0, 1, 0)
		} else if flag == 2 { //日
			strTime = curr.Format("2006-01-02")
			next = curr.AddDate(0, 0, 1)
		} else if flag == 3 { //时
			strTime = curr.Format("15:04")
			next = curr.Add(time.Hour)
		}
		if len(src) > i {
			if src[i].Time != strTime {
				obj := TimeData{Time: strTime, Total: 0}
				src = append(src[:i], append([]TimeData{obj}, src[i:]...)...)
			}
		} else {
			obj := TimeData{Time: strTime, Total: 0}
			src = append(src, obj)
		}
		curr = next
		i++
	}
	return src
}
