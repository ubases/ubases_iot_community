package iotutil

import (
	"math"
	"time"
)

func getPollutionDegree(aqi float64) int {
	var pollutionDegree int = 1
	if aqi <= 50 {
		pollutionDegree = 1
	} else if aqi > 50 && aqi <= 100 {
		pollutionDegree = 2
	} else if aqi > 100 && aqi <= 150 {
		pollutionDegree = 3
	} else if aqi > 150 && aqi <= 200 {
		pollutionDegree = 4
	} else if aqi > 200 && aqi <= 250 {
		pollutionDegree = 5
	} else if aqi > 250 && aqi <= 300 {
		pollutionDegree = 6
	} else if aqi > 300 {
		pollutionDegree = 7
	}
	return pollutionDegree
}

//工业或环境监测用
func getDegree(pollutionDegree int) string {
	if pollutionDegree == 1 {
		return "优"
	} else if pollutionDegree == 2 {
		return "良"
	} else if pollutionDegree == 3 {
		return "轻微污染"
	} else if pollutionDegree == 4 {
		return "轻度污染"
	} else if pollutionDegree == 5 {
		return "中度污染"
	} else if pollutionDegree == 6 {
		return "中度重污染"
	} else if pollutionDegree == 7 {
		return "重度污染"
	}
	return "良"
}

//民用，获取空气质量
// 1 优；  2 良；  3 差
func getDegreeEx(pollutionDegree int) int {
	if pollutionDegree >= 1 && pollutionDegree <= 2 {
		return pollutionDegree
	}
	return 3
}

func countPerIaqi(cp float64, r int) float64 {
	var bph float64 = 0   // 与 cp相近的污染物浓度限值的高位值
	var bpl float64 = 0   // 与 cp相近的污染物浓度限值的低位值
	var iaqih float64 = 0 // 与 bph对应的空气质量分指数
	var iaqil float64 = 0 // 与 bpl对应的空气质量分指数
	var iaqip float64 = 0 // 当前污染物项目P的空气质量分指数

	// 空气质量分指数及对应的污染物项目浓度限值
	var aqiArr [3][8]float64 = [3][8]float64{{0, 50, 100, 150, 200, 300, 400, 500},
		{0, 50, 150, 250, 350, 420, 500, 600},
		{0, 35, 75, 115, 150, 250, 350, 500}}

	var min float64 = aqiArr[r][0]
	var index int = len(aqiArr[r]) - 1
	var max float64 = aqiArr[r][index]
	if cp <= min || cp >= max {
		return 0.0
	} else {
		// 对每种污染物的bph、bpl、iaqih、iaqil进行赋值
		for i := r; i < (r + 1); i++ {
			for j := 0; j < len(aqiArr[0]); j++ {
				if cp < aqiArr[i][j] {
					bph = aqiArr[i][j]
					bpl = aqiArr[i][j-1]
					iaqih = aqiArr[0][j]
					iaqil = aqiArr[0][j-1]
					break
				}
			}
		}
		// 计算污染物项目P的空气质量分指数
		iaqip = (iaqih-iaqil)/(bph-bpl)*(cp-bpl) + iaqil
		return iaqip
	}
}

func getPm10IAQI(pmte float64) float64 {
	if pmte > 0 {
		return countPerIaqi(pmte, 1)
	}
	return 0
}

func getPm25IAQI(pmtw float64) float64 {
	if pmtw > 0 {
		return countPerIaqi(pmtw, 2)
	}
	return 0
}

func CountAqi(pmtw, pmte float64) float64 {
	var pmtwIaqi float64 = getPm25IAQI(pmtw)
	var pmteIaqi float64 = getPm10IAQI(pmte)
	return math.Max(pmteIaqi, pmtwIaqi)
}

func GetDegree(pmtw, pmte float64) string {
	return getDegree(getPollutionDegree(CountAqi(pmtw, pmte)))
}

func GetAqiAndDegree(pmtw, pmte float64) (float64, string) {
	aqi := CountAqi(pmtw, pmte)
	degree := getDegree(getPollutionDegree(aqi))
	return aqi, degree
}

func GetDegreeByAqi(aqi float64) int {
	return getDegreeEx(getPollutionDegree(aqi))
}

func GetTimeLocation(timezone string) *time.Location {
	locStr := timezone
	loc, err := time.LoadLocation(locStr)
	if err != nil {
		loc, _ = time.LoadLocation("Local")
	}
	return loc
}
