package iotutil

import (
	"sort"
	"strconv"
	"strings"
)

//notice 支持3节点分法版本号比较
//版本格式类似 V1.1.1 , v1.1.1 , 2.1.2, V1.1.1-Beta , v1.1.1-Stable , 2.1.2-Alpha , 2.1.2-Alpha1
const segment = 3

//v1=v2 : 0
//v1>v2 : 1
//v1<v2 : -1
func VerCompare(v1 string, v2 string) (int, error) {
	va, err := extract(v1)
	if err != nil {
		return -2, err
	}
	vb, err2 := extract(v2)
	if err2 != nil {
		return -2, err2
	}
	return va.compare(vb), nil
}

type Version struct {
	list   [segment]int //major,minor,build
	suffix string
	src    string
}

func compareInt(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
func (v *Version) compare(vn *Version) int {
	for i, vv := range v.list {
		if r := compareInt(vv, vn.list[i]); r != 0 {
			return r
		}
	}
	if v.suffix == vn.suffix {
		return 0
	}
	list := []string{v.suffix, vn.suffix}
	sort.Strings(list)
	if v.suffix == list[0] {
		return -1
	}
	return 1
}
func extract(in string) (*Version, error) {
	//取短横线之前的比较，之后的不比较
	var ver Version
	ver.src = in
	inList := strings.SplitN(in, "-", 2)
	in = inList[0]
	if len(inList) >= 2 {
		ver.suffix = inList[1]
	}
	in = strings.ReplaceAll(in, "v", "") //可能带v开头，去掉v
	in = strings.ReplaceAll(in, "V", "") //可能带大写V开头，去掉V
	slist := strings.Split(in, ".")
	for i, v := range slist {
		if i >= segment {
			break
		}
		x, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ver.list[i] = x
	}
	return &ver, nil
}

type VerList []*Version

func (vl VerList) Len() int {
	return len(vl)
}

func (vl VerList) Less(i, j int) bool {
	n := vl[i].compare(vl[j])
	if n == 1 {
		return true
	}
	return false
}

func (vl VerList) Swap(i, j int) {
	temp := vl[i]
	vl[i] = vl[j]
	vl[j] = temp
}

func SortVerList(verList []string, desc bool) ([]string, error) {
	var arr VerList
	for _, v := range verList {
		va, err := extract(v)
		if err != nil {
			return nil, err
		}
		arr = append(arr, va)
	}
	sort.Sort(arr)
	retList := make([]string, len(arr))

	if desc {
		for i, v := range arr {
			retList[i] = v.src
		}
	} else {
		for i, v := range arr {
			retList[len(arr)-i-1] = v.src
		}
	}
	return retList, nil
}
