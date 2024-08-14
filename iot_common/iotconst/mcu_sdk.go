package iotconst

// 用于物模型数据类型到MCU SDK的数据类型的转换
var (
	DataTypeMap = map[string]string{
		"BOOL":   "BOOL",
		"INT":    "VALUE",
		"TEXT":   "STRING",
		"ENUM":   "ENUM",
		"DOUBLE": "VALUE",
		"FLOAT":  "VALUE",
		"FAULT":  "ENUM",
	}

	VarValueType = map[string]string{
		"BOOL":   "BOOL",
		"INT":    "unsigned long",
		"TEXT":   "unsigned char",
		"ENUM":   "unsigned char",
		"DOUBLE": "unsigned long long",
		"FLOAT":  "unsigned long",
		"FAULT":  "unsigned long",
	}

	VarDefaultValue = map[string]string{
		"BOOL":   "FALSE",
		"INT":    "0",
		"TEXT":   "",
		"ENUM":   "0",
		"DOUBLE": "0",
		"FLOAT":  "0",
		"FAULT":  "0",
	}
)
