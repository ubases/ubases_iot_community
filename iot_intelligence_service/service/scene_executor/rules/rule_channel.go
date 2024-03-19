package rules

type CreateRuleChanData struct {
	Id     string
	Desc   string
	Status int32
}

var CreateRuleChan chan CreateRuleChanData
