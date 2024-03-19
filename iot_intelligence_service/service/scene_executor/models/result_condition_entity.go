package models

import "strings"

type ResultBools []bool

func (b ResultBools) AllTrue() bool {
	isAllTrue := true
	for _, item := range b {
		if !item {
			isAllTrue = false
			break
		}
	}
	return isAllTrue
}

func (b ResultBools) EveryOneTrue() bool {
	isAllTrue := false
	for _, item := range b {
		if item {
			isAllTrue = true
			break
		}
	}
	return isAllTrue
}

type ResultCondition struct {
	Weather     bool //天气
	Delayed     bool //延时
	StateChange bool //状态变化
	Position    bool //位置变化
}

func (s ResultCondition) Def(defVal bool, vals string) ResultCondition {
	s.Delayed = defVal || strings.Index(vals, "1") != -1
	s.Weather = defVal || strings.Index(vals, "1") != -1
	s.StateChange = defVal || strings.Index(vals, "1") != -1
	s.Position = defVal || strings.Index(vals, "1") != -1
	return s
}

// AllTrue 所有条件为true
func (s ResultCondition) AllTrue() bool {
	return s.Delayed == true &&
		s.Weather == true &&
		s.StateChange == true &&
		s.Position == true
}

// EveryOneTrue 任意条件为true
func (s ResultCondition) EveryOneTrue() bool {
	return s.Delayed == true ||
		s.Weather == true ||
		s.StateChange == true ||
		s.Position == true
}
