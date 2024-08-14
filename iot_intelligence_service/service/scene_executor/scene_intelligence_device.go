package scene_executor

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_intelligence_service/service/scene_executor/valscene"
	"sync"
)

// ObserverDeviceStatusItems 设备状态管理
type ObserverDeviceStatusItems struct {
	observerList   []observerDevice
	observerHasMap sync.Map
}

func (s *ObserverDeviceStatusItems) initSub() {
	s.observerHasMap = sync.Map{}
	go func() {
		for {
			select {
			case job := <-DeviceChan:
				//start := time.Now()
				//iotlogger.LogHelper.Info("chan sub data " + iotutil.ToString(job))
				s.notifyByDid(job.DeviceId, &job)
				//iotlogger.LogHelper.Infof("DeviceChan剩余%d,notifyByDid 处理时间:%v", len(DeviceChan), time.Since(start).Microseconds())
			}
		}
	}()
}
func (s *ObserverDeviceStatusItems) register(o observerDevice) (bool, error) {
	s.observerList = append(s.observerList, o)
	s.observerHasMap.Store(o.getKey(), o)
	return true, nil
}
func (s *ObserverDeviceStatusItems) deregister(o observerDevice) (bool, error) {
	s.removeFormSlice(o)
	return true, nil
}
func (s *ObserverDeviceStatusItems) removeFormSlice(o observerDevice) {
	s.observerHasMap.Delete(o.getKey())
}
func (s *ObserverDeviceStatusItems) notifyAll() {
	//天气发送变化通知所有观察者
	s.observerHasMap.Range(func(key, value interface{}) bool {
		o := value.(observerDevice)
		if o != nil {
			o.run(nil)
		}
		return true
	})
}
func (s *ObserverDeviceStatusItems) notifyByDid(did string, data *iotstruct.DeviceRedisData) {
	if !s.check(did, data) {
		//当前天气无变化
		return
	}
	mapKey, _ := s.observerHasMap.Load(did)
	if mapKey == nil {
		return
	}
	go func() {
		iotlogger.LogHelper.Infof("设备条件管理器 设备Id：%s", did)
		ob := mapKey.(observerDevice)
		ob.run(data)
	}()
}
func (s *ObserverDeviceStatusItems) check(city string, status *iotstruct.DeviceRedisData) bool {
	//检查是否为当前属性
	//data := status.Data.(map[string]interface{})
	return true
}

// 设备状态变化
type observerDevice interface {
	run(data *iotstruct.DeviceRedisData) bool
	getRuleId() string
	getKey() string
}

// DeviceObserver 设备状态观察者
type DeviceObserver struct {
	id        string
	did       string
	funcKey   string
	funcValue interface{}
}

func (w DeviceObserver) run(data *iotstruct.DeviceRedisData) bool {
	if valscene.Gengine == nil {
		return false
	}
	if valscene.DeviceRuleBuilder == nil {
		return false
	}
	status := data.Data.(map[string]interface{})
	v, ok := status[w.funcKey]
	iotlogger.LogHelper.Infof("开始执行场景任务（设备状态编号）：status['%v']: %v, funcValue:%v", w.funcKey, v, w.funcValue)
	if ok && iotutil.ToString(v) == iotutil.ToString(w.funcValue) {
		err := valscene.Gengine.ExecuteSelectedRules(valscene.DeviceRuleBuilder, []string{w.getRuleId()})
		if err != nil {
			return false
		}
	}
	return true
}
func (w DeviceObserver) getRuleId() string {
	return w.id
}
func (w DeviceObserver) getKey() string {
	//设备状态中的Key是设备Id
	return w.did
}
