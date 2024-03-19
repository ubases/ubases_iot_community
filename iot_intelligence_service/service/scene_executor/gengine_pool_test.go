package scene_executor

import (
	"cloud_platform/iot_intelligence_service/cached"
	"cloud_platform/iot_intelligence_service/config"
	"cloud_platform/iot_intelligence_service/service/scene_executor/models"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	context2 "context"
	"fmt"
	"testing"

	"github.com/bilibili/gengine/engine"
)

// 业务规则
const service_rules string = `
rule "1" "1"
begin
	resp.At = room.GetAttention()
	println("rule 1...")
end 

rule "2" "2"
begin
	resp.Num = room.GetNum()
	println("rule 2...")
end
`

// 业务接口
type MyService struct {
	//gengine pool
	Pool *engine.GenginePool

	//other params
}

// request
type Request struct {
	Rid       int64
	RuleNames []string
	//other params
}

// resp
type Response struct {
	At  int64
	Num int64
	//other params
}

// 特定的场景服务
type Room struct {
}

func (r *Room) GetAttention( /*params*/ ) int64 {
	// logic
	return 100
}

func (r *Room) GetNum( /*params*/ ) int64 {
	//logic
	return 111
}

// 初始化业务服务
// apiOuter这里最好仅注入一些无状态函数，方便应用中的状态管理
func NewMyService(poolMinLen, poolMaxLen int64, em int, rulesStr string, apiOuter map[string]interface{}) *MyService {
	pool, e := engine.NewGenginePool(poolMinLen, poolMaxLen, em, rulesStr, apiOuter)
	if e != nil {
		panic(fmt.Sprintf("初始化gengine失败，err:%+v", e))
	}

	myService := &MyService{Pool: pool}
	return myService
}

// service
func (ms *MyService) Service(req *Request) (*Response, error) {

	resp := &Response{}

	//基于需要注入接口或数据,data这里最好仅注入与本次请求相关的结构体或数据，便于状态管理
	data := make(map[string]interface{})
	data["req"] = req
	data["resp"] = resp

	//模块化业务逻辑,api
	room := &Room{}
	data["room"] = room

	//
	e, _ := ms.Pool.ExecuteSelectedRules(data, req.RuleNames)
	if e != nil {
		println(fmt.Sprintf("pool execute rules error: %+v", e))
		return nil, e
	}

	return resp, nil
}

// 模拟调用
func Test_run(t *testing.T) {
	//初始化
	//注入api，请确保注入的API属于并发安全
	apis := make(map[string]interface{})
	apis["println"] = fmt.Println
	msr := NewMyService(10, 20, 1, service_rules, apis)

	//调用
	req := &Request{
		Rid:       123,
		RuleNames: []string{"1", "2"},
	}
	response, e := msr.Service(req)
	if e != nil {
		println(fmt.Sprintf("service err:%+v", e))
		return
	}

	println("resp result = ", response.At, response.Num)
}

func Test_Scan(t *testing.T) {
	config.InitTest()
	cached.InitCache()

	obj := models.SceneIntelligenceTaskForm{
		Id:             iotutil.GetNextSeqInt64(),
		IntelligenceId: 123,
		TaskImg:        "123123",
		TaskDesc:       "123123",
		TaskType:       0,
		ObjectId:       "12312",
		ObjectDesc:     "43242",
		FuncKey:        "123",
		FuncDesc:       "4324",
		FuncValue:      "1231",
	}

	err := iotredis.GetClient().Set(context2.Background(), "长沙", iotutil.ToString(obj), 0)
	if err != nil && err.Err() != nil {
		t.Error(err.Err())
		return
	}
	var newobj models.SceneIntelligenceTaskForm
	weatherCmd := iotredis.GetClient().Get(context2.Background(), "长沙").Val()
	iotutil.JsonToStruct(weatherCmd, &newobj)
	if weatherCmd == "" {
		t.Error(weatherCmd)
		return
	}
	fmt.Println(newobj)

	iotredis.GetClient().HSet(context2.Background(), "test001", "aaa", 1)
	iotredis.GetClient().HSet(context2.Background(), "test001", "bbb", 2)
	iotredis.GetClient().HSet(context2.Background(), "test001", "ccc", 1)
	iotredis.GetClient().HSet(context2.Background(), "test001", "ddd", 14)

	deviceCmd := iotredis.GetClient().HGetAll(context2.Background(), "test001")
	fmt.Println(iotutil.ToString(deviceCmd.Val()), deviceCmd.Val()["aaa"])

}
