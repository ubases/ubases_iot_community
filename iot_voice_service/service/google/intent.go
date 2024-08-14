package google

const IntentSync = "action.devices.SYNC"
const IntentExecute = "action.devices.EXECUTE"
const IntentQuery = "action.devices.QUERY"
const IntentDISCONNECT = "action.devices.DISCONNECT"

type IntentAspect struct {
	Intent string
	Func   func(userId string)
}
