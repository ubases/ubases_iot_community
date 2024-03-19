package valscene

import (
	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/engine"
)

//"time":        ret.Data.UpdatedAt,
//"weather":     ret.Data.Weather,
//"sun":         isSunrise,
//"temperature": ret.Data.Temperature,
//"humidity":    ret.Data.Humidity,
//"pm_2_5":      ret.Data.Pm25,
//"quality":     ret.Data.Quality,
//"windspeed":   ret.Data.WindSpeed,

var (
	RuleBuilder *builder.RuleBuilder

	WeatherRuleBuilder *builder.RuleBuilder
	DeviceRuleBuilder  *builder.RuleBuilder
	TimerRuleBuilder   *builder.RuleBuilder
	Gengine            *engine.Gengine
	WeatherType        = map[int32]string{
		1: "temperature",
		2: "humidity",
		3: "weather",
		4: "pm_2_5",
		5: "quality",
		6: "sun",
		7: "windspeed",
	}
)
