package iotenums

// ToStatusName 状态 =1 正常 =2 禁用
func ToStatusName(status int32) string {
	res := "其它"
	switch status {
	case 1:
		res = "正常"
	case 2:
		res = "禁用"
	}
	return res
}

// ToDisabledName 状态 =1 启用 =2 停用
func ToDisabledName(status int32) string {
	res := "其它"
	switch status {
	case 1:
		res = "启用"
	case 2:
		res = "禁用"
	}
	return res
}

// ToShelfName 状态 =1 上架 =2 禁用
func ToShelfName(status int32) string {
	res := "其它"
	switch status {
	case 1:
		res = "上架"
	case 2:
		res = "禁用"
	}
	return res
}
