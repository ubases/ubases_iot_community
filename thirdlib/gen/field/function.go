package field

import (
	"gorm.io/gorm/clause"
)

// Func sql functions
var Func = new(function)

type function struct{}

func (f *function) UnixTimestamp(date ...string) Uint64 {
	if len(date) > 0 {
		return Uint64{expr{e: clause.Expr{SQL: "UNIX_TIMESTAMP(?)", Vars: []interface{}{date[0]}}}}
	}
	return Uint64{expr{e: clause.Expr{SQL: "UNIX_TIMESTAMP()"}}}
}

// 对应MySQL产品库iot_product的getProductTypeFullName函数
func (f *function) GetProductTypeFullName(id Int64) String {
	return String{expr{e: clause.Expr{SQL: "getProductTypeFullName(?)", Vars: []interface{}{id.RawExpr()}}}}
}

// 自定义版本号排序 000.000.000
func (f *function) VersionOrder(version String) Int {
	//INET_ATON(CONCAT(`t_pm_firmware_version`.`version`, '.0'))
	return Int{expr{e: clause.Expr{SQL: "INET_ATON(CONCAT(?, '.0'))", Vars: []interface{}{version.RawExpr()}}}}
}

// 自定义版本号转换
func (f *function) UnVersionOrder(version Float64) Int {
	//SUBSTRING_INDEX(INET_NTOA(max(INET_ATON(CONCAT(`t_pm_firmware_version`.`version`, '.0')))), '.', 3)
	return Int{expr{e: clause.Expr{SQL: "SUBSTRING_INDEX(INET_NTOA(?), '.', 3)", Vars: []interface{}{version.RawExpr()}}}}
}

// 自定义版本号排序, 通过填充0的方式
func (f *function) VersionOrderByZeroFill(version String) Int {
	return Int{expr{e: clause.Expr{
		SQL:  `CONCAT(LPAD(SUBSTRING_INDEX(SUBSTRING_INDEX(?, '.', 1), '.', -1), 6, '0'),LPAD(SUBSTRING_INDEX(SUBSTRING_INDEX(?, '.', 2), '.', -1), 6, '0'),LPAD(SUBSTRING_INDEX(SUBSTRING_INDEX(?, '.', 3), '.', -1), 6, '0'))`,
		Vars: []interface{}{version.RawExpr(), version.RawExpr(), version.RawExpr()},
	}}}
}
