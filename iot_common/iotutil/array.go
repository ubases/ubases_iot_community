package iotutil

import (
	"sort"
)

// 求并集
func union(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// 求交集
func intersect(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int)
	nn := make([]int64, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// 求差集 slice1-并集
func difference(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int)
	nn := make([]int64, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

type ArraySortList []ArraySortObj

type ArraySortObj struct {
	Id   interface{}
	Sort int32
}

func (o ArraySortList) Len() int {
	return len(o)
}

func (o ArraySortList) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

// Less 根据Sort字段进行排序
func (o ArraySortList) Less(i, j int) bool {
	return o[i].Sort < o[j].Sort
}

func MoveUp(objList ArraySortList, index int) {
	if index <= 0 || index >= len(objList) {
		return
	}
	objList[index].Sort, objList[index-1].Sort = objList[index-1].Sort, objList[index].Sort
	sort.Sort(objList)
}

func MoveDown(objList ArraySortList, index int) {
	if index < 0 || index >= len(objList)-1 {
		return
	}
	objList[index].Sort, objList[index+1].Sort = objList[index+1].Sort, objList[index].Sort
	sort.Sort(objList)
}
