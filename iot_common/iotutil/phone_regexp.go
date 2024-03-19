package iotutil

import "regexp"

var AreaPhoneNumberRegexp = map[string]string{
	"86": `^(\+?0?86\-?)?1[345789]\d{9}$`,
	"1":  `^(\+?1)?[2-9]\d{2}[2-9]\d{6}$`,
}

// 检查国内国外手机号码格式
func CheckAllPhone(areaPhone, phone string) bool {
	switch areaPhone {
	case "":
		for _, v := range AreaPhoneNumberRegexp {
			if ok := regexp.MustCompile(v).MatchString(phone); ok {
				return ok
			}
		}
	default:
		return CheckAreaPhone(areaPhone, phone)
	}
	return false
}

// 检查手机号码格式
func CheckAreaPhone(areaPhone, phone string) bool {
	reg, ok := AreaPhoneNumberRegexp[areaPhone]
	if !ok {
		return false
	}
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}
