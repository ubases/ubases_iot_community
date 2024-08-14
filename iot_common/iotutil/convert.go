/**
 * @Author: hogan
 * @Date: 2022/3/23 20:19
 */
package iotutil

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

var regAccount *regexp.Regexp = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,15}")

func IsEmpty(value interface{}) bool {
	var isempty bool
	switch value.(type) {
	case float64, float32, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		isempty = value == 0
	case string:
		isempty = value == ""
	case []byte:
		isempty = value == ""
	default:
		isempty = value == nil
	}
	return isempty
}

func ToIntErr(value interface{}) (val int, err error) {
	switch value.(type) {
	case string:
		v2, err := strconv.Atoi(value.(string))
		if err != nil {
			return 0, err
		}
		val = int(v2)
	case int:
		val = int(value.(int))
	case int8:
		val = int(value.(int8))
	case int32:
		val = int(value.(int32))
	case int64:
		val = int(value.(int64))
	case uint8:
		val = int(value.(uint8))
	case uint32:
		val = int(value.(uint32))
	case uint64:
		val = int(value.(uint64))
	case float64:
		val = int(value.(float64))
	}
	return
}

func ToInt(value interface{}) int {
	var val int
	switch value.(type) {
	case string:
		v2, err := strconv.Atoi(value.(string))
		if err != nil {
			panic(err)
		}
		val = v2
	case int:
		val = int(value.(int))
	case int8:
		val = int(value.(int8))
	case int32:
		val = int(value.(int32))
	case int64:
		val = int(value.(int64))
	case uint8:
		val = int(value.(uint8))
	case uint32:
		val = int(value.(uint32))
	case uint64:
		val = int(value.(uint64))
	case float64:
		val = int(value.(float64))
	}
	return val
}

// ToInt64 调用该方法之前，请先判断value是否为有效的数值value !="" && value
func ToInt64(value interface{}) int64 {
	var val int64
	switch value.(type) {
	case string:
		v2, err := strconv.Atoi(value.(string))
		if err != nil {
			panic(err)
		}
		val = int64(v2)
		break
	case int:
		val = int64(value.(int))
	case int8:
		val = int64(value.(int8))
	case int32:
		val = int64(value.(int32))
	case int64:
		val = value.(int64)
	case uint8:
		val = int64(value.(uint8))
	case uint32:
		val = int64(value.(uint32))
	case uint64:
		val = int64(value.(uint64))
	case float64:
		val = int64(value.(float64))
	}
	return val
}

func ToInt64AndErr(value interface{}) (int64, error) {
	if value == nil {
		return 0, errors.New("value is empty")
	}
	var val int64
	var err error
	switch value.(type) {
	case string:
		if value.(string) == "" {
			return 0, errors.New("value is null")
		}
		v2, err := strconv.Atoi(value.(string))
		if err != nil {
			return 0, err
		}
		val = int64(v2)
		break
	case int:
		val = int64(value.(int))
	case int8:
		val = int64(value.(int8))
	case int32:
		val = int64(value.(int32))
	case int64:
		val = int64(value.(int64))
	case uint8:
		val = int64(value.(uint8))
	case uint32:
		val = int64(value.(uint32))
	case uint64:
		val = int64(value.(uint64))
	case float64:
		val = int64(value.(float64))
		break
	}
	return val, err
}

func ToInt32(value interface{}) int32 {
	var val int32
	switch value.(type) {
	case string:
		v2, err := strconv.Atoi(value.(string))
		if err != nil {
			panic(err)
		}
		val = int32(v2)
	case int:
		val = int32(value.(int))
	case int8:
		val = int32(value.(int8))
	case int32:
		val = int32(value.(int32))
	case int64:
		val = int32(value.(int64))
	case uint8:
		val = int32(value.(uint8))
	case uint32:
		val = int32(value.(uint32))
	case uint64:
		val = int32(value.(uint64))
	case float64:
		val = int32(value.(float64))
	}
	return val
}

func ToInt32Err(value interface{}) (val int32, err error) {
	switch value.(type) {
	case string:
		if value.(string) == "" {
			return 0, errors.New("value is null")
		}
		v2, err := strconv.Atoi(value.(string))
		if err == nil {
			val = int32(v2)
		}
	case int:
		val = int32(value.(int))
	case int8:
		val = int32(value.(int8))
	case int32:
		val = int32(value.(int32))
	case int64:
		val = int32(value.(int64))
	case uint8:
		val = int32(value.(uint8))
	case uint32:
		val = int32(value.(uint32))
	case uint64:
		val = int32(value.(uint64))
	case float64:
		val = int32(value.(float64))
	}
	return
}

func ToUint64(value interface{}) uint64 {
	return uint64(ToInt64(value))
}

func ToFloat64(inter interface{}) float64 {
	var val float64
	switch inter.(type) {
	case string:
		v2, err := strconv.ParseFloat(inter.(string), 64)
		if err != nil {
			panic(err)
		}
		val = v2
		break
	case int:
		val = float64(inter.(int))
		break
	case int64:
		val = float64(inter.(int64))
		break
	case float64:
		val = inter.(float64)
		break
	default:
		val = inter.(float64)
	}
	return val
}

func ToFloat64Err(inter interface{}) (val float64, err error) {
	switch inter.(type) {
	case string:
		v2, fErr := strconv.ParseFloat(inter.(string), 64)
		if fErr != nil {
			err = fErr
			return
		}
		val = v2
		break
	case int:
		val = float64(inter.(int))
		break
	case int64:
		val = float64(inter.(int64))
		break
	case float64:
		val = inter.(float64)
		break
	}
	return
}

func ToFloat32(inter interface{}) float32 {
	var val float32
	switch inter.(type) {
	case string:
		v2, err := strconv.ParseFloat(inter.(string), 32)
		if err != nil {
			panic(err)
		}
		val = float32(v2)
		break
	case int:
		val = float32(inter.(int))
		break
	case int64:
		val = float32(inter.(int64))
		break
	case float32:
		val = inter.(float32)
		break
	}
	return val
}

// StructToMap 利用反射将结构体转化为map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

func Convert16To10(val string) int {
	val10, _ := strconv.ParseInt(val, 16, 32)
	return int(val10)
}

// ArrayToString 将数组格式化为字符串
func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func Int64ArrayToString(array []int64) []string {
	newArray := make([]string, 0)
	for _, id := range array {
		newArray = append(newArray, ToString(id))
	}
	return newArray
}

// interface{} convert to map[string]interface{}
func ToMap(value interface{}) map[string]interface{} {
	switch value.(type) {
	case map[string]interface{}:
		return value.(map[string]interface{})
	default:
		res, err := JsonToMapErr(ToString(value))
		if err != nil {
			return nil
		}
		return res
	}
}

// map转字符串
func ObjToString(obj interface{}) string {
	str, _ := json.Marshal(obj)
	return string(str)
}

// json string 转换为map
func JsonToMap(jsonStr string) (resObj map[string]interface{}) {
	err := json.Unmarshal([]byte(jsonStr), &resObj)
	if err != nil {
		panic(err)
	}
	return
}

func JsonToMapArray(jsonStr string) (resObj []map[string]interface{}) {
	err := json.Unmarshal([]byte(jsonStr), &resObj)
	if err != nil {
		panic(err)
	}
	return
}

func JsonToMapArrayErr(jsonStr string) (resObj []map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(jsonStr), &resObj)
	return
}

func JsonToStringArray(jsonStr string) (resObj []string) {
	err := json.Unmarshal([]byte(jsonStr), &resObj)
	if err != nil {
		panic(err)
	}
	return
}

func JsonToMapErr(jsonStr string) (resObj map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(jsonStr), &resObj)
	return
}

// json string convert to struct
// 调用示例
// var resObj TestObj
// convert.JsonToMap2(string(str), &resObj)
func JsonToStruct(jsonStr string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return err
	}
	return nil
}

// 相同接口之间的转换
func StructToStruct(origin interface{}, target interface{}) {
	err := json.Unmarshal([]byte(ObjToString(origin)), &target)
	if err != nil {
		panic(err)
	}
}

func StructToStructErr(origin interface{}, target interface{}) error {
	return json.Unmarshal([]byte(ToString(origin)), &target)
}

// 解析对象json
func JsonParse(jsonStr string, fields string) string {
	var fieldArr []string = strings.Split(fields, ".")
	var reqParams map[string]interface{}
	var fieldByte []byte = []byte(jsonStr)
	for _, v := range fieldArr {
		json.Unmarshal(fieldByte, &reqParams)
		fieldByte, _ = json.Marshal(reqParams[v])
	}

	return strings.Trim(string(fieldByte), "\"")
}

// convert.ToString 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func ToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// json里面如果带url使用此方法转字符串合适. 不用将特殊字符转义[比如&转义成\u0026]
func ToStringByUrl(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	newValue, _ := marshal_inner(value)
	key = string(newValue)
	return key
}

func marshal_inner(data interface{}) ([]byte, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(data); err != nil {
		return nil, err
	}
	return bf.Bytes(), nil
}

// 四舍五入
func Round4s5r(x float64) int64 {
	return int64(math.Floor(x + 0.5))
}

// 解析对象json
// 读取json路径
// ReadJsonFieldPath(obj, "field1.field2")
func ReadJsonFieldPath(jsonStr interface{}, fields string) string {
	if fields == "" {
		return ""
	}
	var fieldArr []string = strings.Split(fields, ".")
	var reqParams map[string]interface{}
	var fieldByte []byte = []byte(jsonStr.(string))
	for _, v := range fieldArr {
		json.Unmarshal(fieldByte, &reqParams)
		fieldByte, _ = json.Marshal(reqParams[v])
	}
	return strings.Trim(string(fieldByte), "\"")
}

// 解析对象json
// Map中通过字段路径读取
// ReadMapFieldPath(obj, "field1.field2")
func ReadMapFieldPath(reqParams map[string]interface{}, fields string) interface{} {
	if fields == "" {
		return nil
	}
	var fieldArr []string = strings.Split(fields, ".")
	var tempParams interface{} = reqParams
	for _, v := range fieldArr {
		if tempParams != nil {
			tempParams = ToMap(tempParams)[v]
		}
	}
	return tempParams
}

// 向字符串增加字符
// strs 使用dep隔开的字符串
// str 需要添加的字符串
// dep 分隔符
func JoinStringArr(strs, str string, dep string) string {
	if len(strs) == 0 {
		return str
	}
	arr := strings.Split(strs, dep)
	arr = append(arr, str)
	return strings.Join(arr, dep)
}

// IsPhone 是否为手机 (重构方法名称)
func IsPhone(str string) bool {
	return CheckPhone(str)
}

// IsEmail 是否为邮箱 (重构方法名称)
func IsEmail(str string) bool {
	return VerifyEmailFormat(str)
}

func IsNumeric(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

// 检查手机号码格式
func CheckPhone(value string) bool {
	// reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	reg := `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4(?:[14]0\d{3}|[68]\d{4}|[579]\d{2}))\d{6}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(value)
}

// // 检查手机号码格式
// func CheckForeignPhone(value string) bool {
// 	// reg := `\+(9[976]\d|8[987530]\d|6[987]\d|5[90]\d|42\d|3[875]\d|
// 	// 	2[98654321]\d|9[8543210]|8[6421]|6[6543210]|5[87654321]|
// 	// 	4[987654310]|3[9643210]|2[70]|7|1)\d{1,14}$`
// 	// reg := `^(((\\+\\d{2}-)?0\\d{2,3}-\\d{7,8})|((\\+\\d{2}-)?(\\d{2,3}-)?([1][3,4,5,7,8][0-9]\\d{8})))$`
// 	reg := `/^(\+?1)?[2-9]\d{2}[2-9](?!11)\d{6}$/`
// 	rgx := regexp.MustCompile(reg)
// 	return rgx.MatchString(value)
// }

// 检查邮箱格式
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyAccount(account string) bool {
	return regAccount.MatchString(account)
}

// 截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func Getcode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

func UrlEncode(in string) string {
	r := strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")
	return r.Replace(url.QueryEscape(in))
}

func Sign(stringToSign string) string {
	h := hmac.New(sha1.New, []byte(fmt.Sprintf("%s&", "och1oekoXH4hzdyA8PShQKzZAyJWiM")))
	h.Write([]byte(stringToSign))
	return UrlEncode(base64.StdEncoding.EncodeToString(h.Sum(nil)))
}
func Keys(data map[string]string) []string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// 将map转换位键值对字符串
func SortedString(data map[string]string) string {
	var sortQueryString string
	for _, v := range Keys(data) {
		sortQueryString = fmt.Sprintf("%s&%s=%s", sortQueryString, v, UrlEncode(data[v]))
	}
	return sortQueryString
}

// 判断数组里面是否包含值
func ArraysExistsInt64(list []int64, val int64) bool {
	res := false
	for _, v := range list {
		if v == val {
			res = true
			break
		}
	}
	return res
}
func ArraysExistsInt32(list []int32, val int32) bool {
	res := false
	for _, v := range list {
		if v == val {
			res = true
			break
		}
	}
	return res
}

// 判断数组里面是否包含值
func ArraysExistsString(list []string, val string) bool {
	res := false
	for _, v := range list {
		if v == val {
			res = true
			break
		}
	}
	return res
}

func DeleteStringElement(list []string, ele string) []string {
	result := make([]string, 0)
	for _, v := range list {
		if v != ele {
			result = append(result, v)
		}
	}
	return result
}

func DeleteInt32Element(list []int32, ele int32) []int32 {
	result := make([]int32, 0)
	for _, v := range list {
		if v != ele {
			result = append(result, v)
		}
	}
	return result
}

// 随机生成指定位数的字母和数字的组合
func GetRandomStringCombination(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func ToArrayInt64(ids []string) []int64 {
	newIds := make([]int64, 0)
	for _, fid := range ids {
		newIds = append(newIds, ToInt64(fid))
	}
	return newIds
}

func ToInterface(value interface{}, valueType int) interface{} {
	var dictValue interface{}
	var err error
	switch valueType {
	case 1:
		dictValue, err = ToIntErr(value)
		if err != nil {
			return value
		}
		break
	case 2:
		dictValue, err = ToFloat64Err(value)
		if err != nil {
			return value
		}
		break
	case 3:
		dictValue = ToString(value)
		break
	default:
		dictValue = value
	}
	return dictValue
}

func UUID() string {
	uuid := uuid.NewV4().String()
	//uuidHash   := int(crc32.ChecksumIEEE([]byte(uuid.String())))
	return uuid
}

func IfInt32(condition bool, trueVal, falseVal int32) int32 {
	if condition {
		return trueVal
	}
	return falseVal
}

func IfString(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}
func IfStringEmpty(inVal, defaultVal string) string {
	if inVal != "" {
		return inVal
	}
	return defaultVal
}

func IntToBoolean(v int32) bool {
	if v == 1 {
		return true
	}
	return false
}

func GetMapVal(res map[int64]string, key interface{}) string {
	if val, ok := res[ToInt64(key)]; ok {
		return val
	}
	return ""
}

func GetMapStringKey(res map[string]string, key interface{}) string {
	if val, ok := res[ToString(key)]; ok {
		return val
	}
	return ""
}

// 浮点数
var FloatReg = regexp.MustCompile(`^(-?\d+)+(\.\d+)?$`)

// 整数
var IntReg = regexp.MustCompile(`^(\-)?\d+$`)

// 用于redis hashmap获取
func MapStringToInterface(in map[string]string) map[string]interface{} {
	retMap := make(map[string]interface{}, len(in))
	for k, v := range in {
		retMap[k] = v
		if IntReg.MatchString(v) {
			i, err := strconv.ParseInt(v, 10, 64)
			if err == nil {
				retMap[k] = i
			}
		} else if FloatReg.MatchString(v) {
			f, err := strconv.ParseFloat(v, 64)
			if err == nil {
				retMap[k] = f
			}
		} else if v == "true" {
			retMap[k] = true
		} else if v == "false" {
			retMap[k] = false
		}
	}
	return retMap
}

func StringToBool(str string) bool {
	//todo :string to bool
	//接受 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False 等字符串；
	//其他形式的字符串会返回错误
	b, _ := strconv.ParseBool(str)
	return b
}

// 数组去重
func RemoveRepeatElement(slice []string) []string {
	sort.Strings(slice)
	i := 0
	var j int
	for {
		if i >= len(slice)-1 {
			break
		}

		for j = i + 1; j < len(slice) && slice[i] == slice[j]; j++ {
		}
		slice = append(slice[:i+1], slice[j:]...)
		i++
	}
	return slice
}

// err 转 string
func ResErrToString(err error) string {
	var result string
	if strings.Index(err.Error(), "go.micro.client") != -1 {
		errMap := JsonToMap(err.Error())
		if msg, ok := errMap["detail"]; ok {
			result = msg.(string)
		} else {
			result = err.Error()
		}
	} else {
		result = err.Error()
	}
	return result
}

// HToSTemperature 华氏温度转换为摄氏温度的函数
func HToSTemperature(f float64) float64 {
	return (f - 32) * 5 / 9
}

// SToHTemperature 摄氏温度转换为华氏温度的函数
func SToHTemperature(c float64) float64 {
	return (9*c + 160) / 5
}

func MapGetStringVal(mapData interface{}, defaultVal interface{}) string {
	if mapData != nil && mapData != "" {
		return ToString(mapData)
	}
	return ToString(defaultVal)
}

func MapGetInt64Val(mapData interface{}, defaultVal interface{}) int64 {
	if mapData != nil && mapData != "" {
		v, err := ToInt64AndErr(mapData)
		if err != nil {
			return ToInt64(defaultVal)
		}
		return v
	}
	return ToInt64(defaultVal)
}

func MapGetInt32Val(mapData interface{}, defaultVal interface{}) int32 {
	if mapData != nil && mapData != "" {
		return ToInt32(mapData)
	}
	return ToInt32(defaultVal)
}

func MapGetInt32ValExt(mapData *int32, defaultVal *int32) int32 {
	if mapData != nil {
		return ToInt32(*mapData)
	}
	return ToInt32(*defaultVal)
}

// ArrayUnionInterfaces 将两个数组合并未一个map，参数1作为键，参数2作为值
func ArrayUnionInterfaces(keys []string, slice []interface{}) (resMap map[string]string) {
	resMap = make(map[string]string)
	if keys == nil || len(keys) == 0 {
		return
	}
	if slice == nil || len(slice) == 0 {
		return
	}
	for i, v := range keys {
		resMap[v] = ToString(slice[i])
	}
	return
}

// checkUserName 检查用户名称，判断为手机、邮箱
func CheckUserName(userName string) (phone, email string, err error) {
	if CheckAllPhone("", userName) {
		phone = userName
	} else if IsEmail(userName) {
		email = userName
	} else {
		err = errors.New("无法识别账号格式")
	}
	return
}

func ToInt32ErrNew(value interface{}) (val int32, err error) {
	switch value.(type) {
	case string:
		v2, err := strconv.Atoi(value.(string))
		if err != nil {
			return 0, err
		}
		val = int32(v2)
		break
	case int:
		val = int32(value.(int))
	case int8:
		val = int32(value.(int8))
	case int32:
		val = int32(value.(int32))
	case int64:
		val = int32(value.(int64))
	case uint8:
		val = int32(value.(uint8))
	case uint32:
		val = int32(value.(uint32))
	case uint64:
		val = int32(value.(uint64))
	case float64:
		val = int32(value.(float64))
	}
	return
}

func IsContainsByList(list []string, param string) bool {
	for _, v := range list {
		if v == param {
			return true
		}
	}
	return false
}

// 判断字符是否存在与数组中
func InArray(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

// Int64数组去重
func RemoveRepeatInt64Element(slice []int64) []int64 {
	i := 0
	var j int
	for {
		if i >= len(slice)-1 {
			break
		}

		for j = i + 1; j < len(slice) && slice[i] == slice[j]; j++ {
		}
		slice = append(slice[:i+1], slice[j:]...)
		i++
	}
	return slice
}

func GetInt32AndDef(val, minValue int32) int32 {
	if val == 0 {
		return 1
	}
	return val
}

func GetStringAndDef(val, defValue string) string {
	if val == "" {
		return defValue
	}
	return val
}

func FindMissingElements(A []string, B []string) []string {
	result := []string{}
	// 创建一个 map，用于存储数组 B 中的元素
	elements := make(map[string]bool)
	for _, elem := range B {
		elements[elem] = true
	}
	// 遍历数组 A，将不在元素 map 中的元素加入结果数组
	for _, elem := range A {
		if _, found := elements[elem]; !found {
			result = append(result, elem)
		}
	}
	return result
}

// BoolArrayIsVal 检查布尔数组中的所有值是否与指定的布尔值相等。
func BoolArrayIsVal(vals []bool, eqVal bool) bool {
	if len(vals) == 0 {
		return true
	}
	for _, v := range vals {
		if v != eqVal {
			return false
		}
	}
	return true
}
// RemoveEmptyString 移除字符串切片中的空字符串。
// 参数 arr 是一个字符串切片，可能包含空字符串。
func RemoveEmptyString(arr []string) []string {
	result := []string{}
	for _, str := range arr {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}
func IsJSON(data []byte) (map[string]interface{}, error) {
	var js map[string]interface{}
	err := json.Unmarshal(data, &js)
	return js, err
}
